package main

import (
	"encoding/hex"
	"net"
	"time"

	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/logger"
)

// registerValidator adds a validator to the registry
func registerValidator(conn net.Conn, validatorAddr, validatorEthAddr string) {
	validators.Lock()
	if validators.connections == nil {
		validators.connections = make(map[net.Conn]bool)
	}
	validators.connections[conn] = true

	// Store validator's listening address
	if validators.addresses == nil {
		validators.addresses = make(map[net.Conn]string)
	}
	validators.addresses[conn] = validatorAddr

	// Store validator's ETH address
	if validators.ethAddresses == nil {
		validators.ethAddresses = make(map[net.Conn]string)
	}
	validators.ethAddresses[conn] = validatorEthAddr
	validators.Unlock()

	logger.Info("New validator connected. Total validators: %d", len(validators.connections))
}

// unregisterValidator removes a validator from the registry
func unregisterValidator(conn net.Conn) {
	validators.Lock()
	delete(validators.connections, conn)
	delete(validators.addresses, conn)
	delete(validators.ethAddresses, conn)
	conn.Close()
	validators.Unlock()
	logger.Info("Validator removed. Remaining validators: %d", len(validators.connections))
}

// getConnectedValidators returns a list of connected validators
func getConnectedValidators() ([]config.ValidatorClient, int) {
	var validatorClients []config.ValidatorClient

	// Get validator addresses
	validators.RLock()
	for conn := range validators.connections {
		addr := validators.addresses[conn]
		if addr != "" {
			validatorClients = append(validatorClients, config.ValidatorClient{
				Conn: conn,
				Addr: addr,
			})
		}
	}
	validatorCount := len(validatorClients)
	validators.RUnlock()

	return validatorClients, validatorCount
}

// collectValidatorResponses collects responses from validators with a timeout
func collectValidatorResponses(validatorClients []config.ValidatorClient, results chan config.ValidatorResponse, validatorCount int) (map[string][]byte, int) {
	// Collect responses with timeout
	timeout := time.After(1 * time.Second)
	validResponses := make(map[string][]byte)
	errorCount := 0

	// Wait for all responses or timeout
	for i := 0; i < validatorCount; i++ {
		select {
		case result := <-results:
			if result.Error != nil {
				logger.Error("Validator error: %v", result.Error)
				errorCount++
			} else {
				validators.RLock()
				ethAddr := validators.ethAddresses[validatorClients[i].Conn]
				validators.RUnlock()
				validResponses[ethAddr] = result.Response
				logger.Debug("Received valid response: %s", string(result.Response))
			}
		case <-timeout:
			logger.Warn("Timeout reached. Received %d valid responses, %d errors, %d validators did not respond",
				len(validResponses), errorCount, validatorCount-len(validResponses)-errorCount)
			return validResponses, errorCount
		}
	}

	return validResponses, errorCount
}

// storeValidatorResponses stores validator responses in the database
func storeValidatorResponses(requestID int64, validResponses map[string][]byte) {
	for ethAddr, response := range validResponses {
		_, err := dbClient.InsertDisSignResponse(requestID, ethAddr, hex.EncodeToString(response))
		if err != nil {
			logger.Error("Failed to store validator response: %v", err)
		}
	}
	err := dbClient.UpdateDisSignRequestStatus(requestID, "COMPLETED")
	if err != nil {
		logger.Error("Failed to update request status: %v", err)
	}
}
