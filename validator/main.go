package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type Config struct {
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	Dispatcher string `json:"dispatcher"`
}

var config Config

func loadConfig() {
	file, err := os.Open("docs/config/val_config.json")
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

	// Start TCP server to accept connections
	listener, err := net.Listen("tcp", ":0") // Listen on any available port
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	localAddr := listener.Addr().String()
	log.Printf("Validator listening on %s", localAddr)

	// Connect to dispatcher for heartbeat
	go connectToDispatcher(localAddr)

	// Accept connections for signing requests
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleSigningRequest(conn)
	}
}

func connectToDispatcher(localAddr string) {
	for {
		conn, err := net.Dial("tcp", config.Dispatcher)
		if err != nil {
			log.Printf("Error connecting to dispatcher: %v", err)
			log.Println("Retrying connection in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("Connected to dispatcher at %s", config.Dispatcher)

		// Send local address as first message
		_, err = conn.Write([]byte("ADDR:" + localAddr))
		if err != nil {
			log.Printf("Failed to send local address: %v", err)
			conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		// Set a read deadline for the first heartbeat cycle
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))

		// Start heartbeat with connection monitoring
		heartbeatDone := make(chan struct{})
		go func() {
			handleHeartbeat(conn)
			close(heartbeatDone)
		}()

		// Wait for heartbeat to finish (connection lost)
		<-heartbeatDone
		log.Println("Connection lost, attempting to reconnect...")
		time.Sleep(5 * time.Second)
	}
}

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
			log.Printf("Failed to send heartbeat: %v", err)
			return
		}

		// Read response
		buf := make([]byte, 4)
		n, err := conn.Read(buf)
		if err != nil {
			failureCount++
			log.Printf("Error reading heartbeat response: %v (failure %d/%d)",
				err, failureCount, maxFailures)

			if failureCount >= maxFailures {
				log.Printf("Too many heartbeat failures, reconnecting...")
				return
			}
			continue
		}

		// Reset failure counter on successful heartbeat
		failureCount = 0

		if string(buf[:n]) != "pong" {
			log.Printf("Unexpected heartbeat response: %s", string(buf[:n]))
		}
	}
}

func handleSigningRequest(conn net.Conn) {
	defer conn.Close()

	// Read request
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error reading request: %v", err)
		return
	}

	request := buf[:n]
	log.Printf("Received signing request: %s", string(request))

	// simulate signature
	signature := fmt.Sprintf("Signed by validator: %s", string(request))

	// Send response
	_, err = conn.Write([]byte(signature))
	if err != nil {
		log.Printf("Error sending response: %v", err)
		return
	}

	log.Printf("Sent response: %s", signature)
}
