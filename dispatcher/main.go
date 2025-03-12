package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type Config struct {
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
}

var config Config
var validators struct {
	sync.RWMutex
	connections map[net.Conn]bool
	addresses   map[net.Conn]string
}

func init() {
	validators.connections = make(map[net.Conn]bool)
	validators.addresses = make(map[net.Conn]string)
}

func loadConfig() {
	file, err := os.Open("docs/config/dis_config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
}

func main() {
	loadConfig()

	http.HandleFunc("/reqblssign", handleRequest)
	go http.ListenAndServe(":8080", nil)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	log.Println("The dispatcher is listening now, waiting for connected.")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleValidatorConnection(conn)
	}
}

func handleValidatorConnection(conn net.Conn) {
	// Wait for the validator to send its listening address
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading validator address: %v", err)
		conn.Close()
		return
	}

	data := string(buffer[:n])
	var validatorAddr string

	// Parse the ADDR: message
	if len(data) > 5 && data[:5] == "ADDR:" {
		validatorAddr = data[5:]
		log.Printf("Validator registered with address: %s", validatorAddr)
	} else {
		log.Printf("Invalid address format from validator: %s", data)
		conn.Close()
		return
	}

	validators.Lock()
	validators.connections[conn] = true
	// Store the validator's listening address
	if validators.addresses == nil {
		validators.addresses = make(map[net.Conn]string)
	}
	validators.addresses[conn] = validatorAddr
	validators.Unlock()

	log.Printf("New validator connected. Total validators: %d", len(validators.connections))

	// Handle connection until it's closed
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Validator disconnected: %v", err)
			validators.Lock()
			delete(validators.connections, conn)
			delete(validators.addresses, conn)
			conn.Close()
			validators.Unlock()
			log.Printf("Validator removed. Remaining validators: %d", len(validators.connections))
			return
		}

		data := string(buffer[:n])
		if data != "ping" {
			log.Printf("Received data from validator: %s", data)
		}

		// If it's a heartbeat, respond with pong
		if data == "ping" {
			_, err = conn.Write([]byte("pong"))
			if err != nil {
				log.Printf("Failed to send heartbeat response: %v", err)
				validators.Lock()
				delete(validators.connections, conn)
				delete(validators.addresses, conn)
				conn.Close()
				validators.Unlock()
				return
			}
		}
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Received client request: %s", string(body))

	// Create separate connections for each validator
	type ValidatorClient struct {
		Conn net.Conn
		Addr string
	}

	var validatorClients []ValidatorClient

	// Get validator addresses
	validators.RLock()
	for conn := range validators.connections {
		addr := validators.addresses[conn]
		if addr != "" {
			validatorClients = append(validatorClients, ValidatorClient{
				Conn: conn,
				Addr: addr,
			})
		}
	}
	validatorCount := len(validatorClients)
	validators.RUnlock()

	if validatorCount == 0 {
		http.Error(w, "No validators connected", http.StatusInternalServerError)
		return
	}

	// Create channels for collecting responses
	type ValidatorResponse struct {
		Response []byte
		Error    error
	}
	results := make(chan ValidatorResponse, validatorCount)

	// Create new connections for each request
	for _, vc := range validatorClients {
		go func(addr string) {
			// Create a new connection for this request
			reqConn, err := net.Dial("tcp", addr)
			if err != nil {
				results <- ValidatorResponse{Error: fmt.Errorf("connection error: %v", err)}
				return
			}
			defer reqConn.Close()

			// Set timeouts
			reqConn.SetDeadline(time.Now().Add(900 * time.Millisecond))

			// Send request
			_, err = reqConn.Write(body)
			if err != nil {
				results <- ValidatorResponse{Error: fmt.Errorf("write error: %v", err)}
				return
			}

			// Read response
			var responseData []byte
			buffer := make([]byte, 1024)

			for {
				n, err := reqConn.Read(buffer)
				if err != nil {
					if len(responseData) == 0 {
						results <- ValidatorResponse{Error: fmt.Errorf("read error: %v", err)}
					} else {
						results <- ValidatorResponse{Response: responseData}
					}
					return
				}
				responseData = append(responseData, buffer[:n]...)

				// If we received less than buffer size, we've got the complete message
				if n < len(buffer) {
					break
				}
			}

			log.Printf("Received response from validator: %s", string(responseData))
			results <- ValidatorResponse{Response: responseData}
		}(vc.Addr)
	}

	// Collect responses with timeout
	timeout := time.After(1 * time.Second)
	validResponses := make([][]byte, 0, validatorCount)
	errorCount := 0

	// Wait for all responses or timeout
	for i := 0; i < validatorCount; i++ {
		select {
		case result := <-results:
			if result.Error != nil {
				log.Printf("Validator error: %v", result.Error)
				errorCount++
			} else {
				validResponses = append(validResponses, result.Response)
				log.Printf("Received valid response: %s", string(result.Response))
			}
		case <-timeout:
			log.Printf("Timeout reached. Received %d valid responses, %d errors, %d validators did not respond",
				len(validResponses), errorCount, validatorCount-len(validResponses)-errorCount)
			goto respondToClient
		}
	}

respondToClient:
	if len(validResponses) == 0 {
		http.Error(w, "No valid responses received from validators", http.StatusInternalServerError)
		return
	}

	// Write summary of responses
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"total_validators":    validatorCount,
		"responses_received":  len(validResponses),
		"errors":              errorCount,
		"no_response":         validatorCount - len(validResponses) - errorCount,
		"validator_responses": validResponses,
	}
	json.NewEncoder(w).Encode(response)
}
