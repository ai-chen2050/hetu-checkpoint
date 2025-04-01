package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/proto/types"
	"github.com/hetu-project/hetu-checkpoint/store"
)

// handleRequest processes HTTP requests for BLS signing
func handleRequest(w http.ResponseWriter, r *http.Request, cfg *config.DispatcherConfig) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Parse the request to get checkpoint data
	var req config.Request
	if err := json.Unmarshal(body, &req); err != nil {
		logger.Error("Failed to parse request message: %v", err)
		return
	}

	logger.Info("Received client request: %s", string(body))

	// Generate message to be signed
	msgToBeSigned := config.GetSignBytes(req.CheckpointWithMeta.Ckpt.EpochNum-1, *req.CheckpointWithMeta.Ckpt.BlockHash)

	// Store the request in database if enabled
	var request *store.SignRequest
	if enableDB {
		var err error
		request, err = dbClient.InsertDisSignRequest(string(body))
		if err != nil {
			logger.Error("Failed to store sign request: %v", err)
			// Continue processing even if DB storage fails
		}
	}

	validatorClients, validatorCount := getConnectedValidators()

	if validatorCount == 0 {
		logger.Warn("No validators connected")
		http.Error(w, "No validators connected", http.StatusInternalServerError)
		return
	}

	// Create channels for collecting responses
	results := make(chan config.ValidatorResponse, validatorCount)

	// Create new connections for each request
	for _, vc := range validatorClients {
		go requestValidatorSignature(vc.Addr, msgToBeSigned, results)
	}

	// Collect responses with timeout
	validResponses, errorCount := collectValidatorResponses(validatorClients, results, validatorCount)
	validRespLen := len(validResponses)

	if validRespLen == 0 {
		logger.Error("No valid responses received from validators")
		if enableDB && request != nil {
			_ = dbClient.UpdateDisSignRequestStatus(request.ID, "FAILED")
		}
		http.Error(w, "No valid responses received from validators", http.StatusInternalServerError)
		return
	}

	// Store validator responses if DB is enabled
	if enableDB && request != nil {
		storeValidatorResponses(request.ID, validResponses)
	}

	// Aggregate signatures and update checkpoint
	aggregatedCkpt, err := AggregateSignatures(validResponses, &req, cfg)
	if err != nil {
		logger.Error("Failed to aggregate signatures: %v", err)
	} else {
		logger.Info("Successfully aggregated %d signatures for epoch %d",
			validRespLen, req.CheckpointWithMeta.Ckpt.EpochNum)

		// Update the request with the aggregated checkpoint
		req.CheckpointWithMeta = *aggregatedCkpt

		// Only log bitmap if aggregation was successful
		logger.Info("Aggregated checkpoint's bitmap: %s", hex.EncodeToString(aggregatedCkpt.Ckpt.Bitmap))

		// Store aggregated checkpoint if DB is enabled
		if enableDB && request != nil {
			storeAggregatedCheckpoint(aggregatedCkpt, request, validRespLen)
		}
	}

	// Report BLS signatures if enabled
	if cfg.EnableReport {
		// Report to Ethereum if aggregation was successful
		if aggregatedCkpt != nil {
			go ReportCheckpointWithRetry(aggregatedCkpt, cfg, 3)
		}
	}

	// Convert binary responses to map for JSON response
	jsonResponses := make(map[string]string)
	for ethAddr, response := range validResponses {
		jsonResponses[ethAddr] = hex.EncodeToString(response)
	}

	// Write summary of responses
	w.Header().Set("Content-Type", "application/json")
	response := config.Response{
		TotalValidators:    validatorCount,
		ResponsesReceived:  validRespLen,
		Errors:             errorCount,
		NoResponse:         validatorCount - validRespLen - errorCount,
		ValidatorResponses: jsonResponses,
	}
	json.NewEncoder(w).Encode(response)
}

// storeAggregatedCheckpoint stores the aggregated checkpoint in the database
func storeAggregatedCheckpoint(aggregatedCkpt *types.RawCheckpointWithMeta, request *store.SignRequest, validRespLen int) {
	if dbClient == nil {
		logger.Error("Database client is not initialized")
		return
	}

	// Convert binary data to hex strings for storage
	blockHashHex := hex.EncodeToString(*aggregatedCkpt.Ckpt.BlockHash)
	bitmapHex := hex.EncodeToString(aggregatedCkpt.Ckpt.Bitmap)

	var blsMultiSigHex, blsAggrPkHex string
	if aggregatedCkpt.Ckpt.BlsMultiSig != nil {
		blsMultiSigHex = hex.EncodeToString(*aggregatedCkpt.Ckpt.BlsMultiSig)
	}
	if aggregatedCkpt.BlsAggrPk != nil {
		blsAggrPkHex = hex.EncodeToString(*aggregatedCkpt.BlsAggrPk)
	}

	// Determine status
	status := "PENDING"
	if aggregatedCkpt.Status == types.Sealed {
		status = "SEALED"
	}

	// Store in database
	_, err := dbClient.InsertAggregatedCheckpoint(
		request.ID,
		aggregatedCkpt.Ckpt.EpochNum,
		blockHashHex,
		bitmapHex,
		blsMultiSigHex,
		blsAggrPkHex,
		aggregatedCkpt.PowerSum,
		status,
		validRespLen,
	)
	if err != nil {
		logger.Error("Failed to store aggregated checkpoint: %v", err)
	} else {
		logger.Info("Stored aggregated checkpoint for epoch %d in database", aggregatedCkpt.Ckpt.EpochNum)
	}
}

// handleValidatorConnection manages a connection from a validator
func handleValidatorConnection(conn net.Conn) {
	buffer := make([]byte, 1024)

	// Read validator address information
	n, err := conn.Read(buffer)
	if err != nil {
		logger.Error("Failed to read validator address: %v", err)
		conn.Close()
		return
	}

	addrInfo := string(buffer[:n])
	logger.Info("Received validator address: %s", addrInfo)

	// Skip HTTP/2 protocol messages
	if strings.HasPrefix(addrInfo, "PRI * HTTP/2.0") {
		logger.Warn("Received HTTP/2 protocol message instead of validator address, closing connection")
		conn.Close()
		return
	}

	// Parse address information, format is "listen_addr|eth_addr"
	parts := strings.Split(addrInfo, "|")
	if len(parts) != 2 {
		logger.Error("Invalid validator address format: %s", addrInfo)
		conn.Close()
		return
	}

	validatorAddr := parts[0]
	validatorEthAddr := parts[1]

	// Store connection and address information
	if !registerValidator(conn, validatorAddr, validatorEthAddr) {
		// close the connection
		conn.Write([]byte("ERROR: Maximum validator connections reached"))
		conn.Close()
		return
	}

	// Handle connection until it's closed
	handleValidatorHeartbeat(conn)
}

// handleValidatorHeartbeat manages the heartbeat communication with a validator
func handleValidatorHeartbeat(conn net.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logger.Info("Validator disconnected: %v", err)
			unregisterValidator(conn)
			return
		}

		data := string(buffer[:n])

		// Skip HTTP/2 protocol messages
		if strings.HasPrefix(data, "PRI * HTTP/2.0") {
			logger.Warn("Received HTTP/2 protocol message, ignoring")
			continue
		}

		if data != "ping" {
			logger.Debug("Received data from validator: %s", data)
		}

		// If it's a heartbeat, respond with pong
		if data == "ping" {
			_, err = conn.Write([]byte("pong"))
			if err != nil {
				logger.Error("Failed to send heartbeat response: %v", err)
				unregisterValidator(conn)
				return
			}
		}
	}
}

// requestValidatorSignature sends a signing request to a validator
func requestValidatorSignature(addr string, msgToBeSigned []byte, results chan<- config.ValidatorResponse) {
	// Create a new connection for this request
	reqConn, err := net.Dial("tcp", addr)
	if err != nil {
		results <- config.ValidatorResponse{Error: fmt.Errorf("connection error: %v", err)}
		return
	}
	defer reqConn.Close()

	// Set timeouts
	reqConn.SetDeadline(time.Now().Add(900 * time.Millisecond))

	// Send request
	_, err = reqConn.Write(msgToBeSigned)
	if err != nil {
		results <- config.ValidatorResponse{Error: fmt.Errorf("write error: %v", err)}
		return
	}

	// Read response
	var responseData []byte
	buffer := make([]byte, 1024)

	for {
		n, err := reqConn.Read(buffer)
		if err != nil {
			if len(responseData) == 0 {
				results <- config.ValidatorResponse{Error: fmt.Errorf("read error: %v", err)}
			} else {
				results <- config.ValidatorResponse{Response: responseData}
			}
			return
		}
		responseData = append(responseData, buffer[:n]...)

		// If we received less than buffer size, we've got the complete message
		if n < len(buffer) {
			break
		}
	}

	logger.Debug("Received response from validator: %s", string(responseData))
	results <- config.ValidatorResponse{Response: responseData}
}
