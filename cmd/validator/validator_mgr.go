package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/crypto/bls12381"
	"github.com/hetu-project/hetu-checkpoint/logger"
)

// connectToDispatcher establishes a connection to the dispatcher
func connectToDispatcher(cfg *config.ValidatorConfig) (net.Conn, error) {
	// Connect to dispatcher
	conn, err := net.Dial("tcp", cfg.DispatcherTcp)
	if err != nil {
		return nil, err
	}

	// Send validator address information
	// Format: "listen_addr|eth_addr"
	addrInfo := cfg.ListenAddr + "|" + keyPair.ETH.Address
	_, err = conn.Write([]byte(addrInfo))
	if err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

// maintainDispatcherConnection maintains a connection to the dispatcher with reconnection logic
func maintainDispatcherConnection(cfg *config.ValidatorConfig) {
	for {
		// Connect to dispatcher
		conn, err := connectToDispatcher(cfg)
		if err != nil {
			logger.Error("Failed to connect to dispatcher: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		logger.Info("Connected to dispatcher at %s", cfg.DispatcherTcp)

		// Start heartbeat in a goroutine
		heartbeatDone := make(chan struct{})
		go func() {
			handleHeartbeat(conn)
			close(heartbeatDone)
		}()

		// Wait for heartbeat to finish (connection lost)
		<-heartbeatDone
		logger.Warn("Connection lost, attempting to reconnect...")
		time.Sleep(5 * time.Second)
	}
}

// handleHeartbeat sends periodic heartbeats to the dispatcher
func handleHeartbeat(conn net.Conn) {
	defer conn.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Set up a failure counter
	failureCount := 0
	const maxFailures = 3

	for range ticker.C {
		// Set a deadline for this heartbeat cycle
		conn.SetDeadline(time.Now().Add(2 * time.Second))

		_, err := conn.Write([]byte("ping"))
		if err != nil {
			logger.Error("Failed to send heartbeat: %v", err)
			return
		}

		// Read response
		buf := make([]byte, 4)
		n, err := conn.Read(buf)
		if err != nil {
			failureCount++
			logger.Error("Error reading heartbeat response: %v (failure %d/%d)",
				err, failureCount, maxFailures)

			if failureCount >= maxFailures {
				logger.Warn("Too many heartbeat failures, reconnecting...")
				return
			}
			continue
		}

		// Reset failure counter on successful heartbeat
		failureCount = 0

		if string(buf[:n]) != "pong" {
			logger.Warn("Unexpected heartbeat response: %s", string(buf[:n]))
		}
	}
}

// startListeningServer starts a TCP server to listen for signing requests
func startListeningServer(cfg *config.ValidatorConfig, ready chan struct{}) {
	// Start TCP server to accept connections
	var listener net.Listener
	var err error

	if cfg.Port == 0 {
		// Use random port
		listener, err = net.Listen("tcp", cfg.ListenAddr+":0")
		if err != nil {
			logger.Fatal("Failed to start TCP server: %v", err)
		}
	} else {
		// Use specified port
		portStr := fmt.Sprintf(":%d", cfg.Port)
		listener, err = net.Listen("tcp", cfg.ListenAddr+portStr)
		if err != nil {
			logger.Fatal("Failed to start TCP server: %v", err)
		}
	}

	// Update listen address with actual port
	cfg.ListenAddr = listener.Addr().String()
	logger.Info("Validator listening on %s", cfg.ListenAddr)

	// Signal that the server is ready
	close(ready)

	// Start accepting connections
	acceptConnections(listener)
}

// acceptConnections handles incoming connections
func acceptConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Error accepting connection: %v", err)
			continue
		}

		go handleSigningRequest(conn)
	}
}

// handleSigningRequest processes a signing request from the dispatcher
func handleSigningRequest(conn net.Conn) {
	defer conn.Close()

	// Read request
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		logger.Error("Error reading request: %v", err)
		return
	}

	request := buf[:n]
	logger.Info("Received signing request: %s", hex.EncodeToString(request))

	// Create BLS signature using the loaded key
	var signature []byte
	if keyPair != nil {
		// Convert BLS private key from hex string to bytes
		blsPrivKeyHex := keyPair.BLS.PrivateKey
		blsPrivKeyBytes, err := hex.DecodeString(blsPrivKeyHex)
		if err != nil {
			logger.Error("Failed to decode BLS private key: %v", err)
			return
		}

		// Sign the message using BLS
		blsSig := bls12381.Sign(blsPrivKeyBytes, request)
		signature = blsSig
		logger.Debug("Created BLS signature: %x", signature)
	} else {
		logger.Error("No key pair loaded, cannot sign message")
		return
	}

	// Store the response in database if enabled
	if enableDB {
		validatorID := conn.LocalAddr().String()
		_, err = dbClient.InsertValSignResponse(-1, validatorID, hex.EncodeToString(signature))
		if err != nil {
			logger.Error("Failed to store sign response: %v", err)
		}
	}

	// Send response
	_, err = conn.Write(signature)
	if err != nil {
		logger.Error("Error sending response: %v", err)
		return
	}

	logger.Debug("Sent BLS signature: %x", signature)
}
