package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/logger"
)

var (
	configFile string
	logLevel   string
	port       int
)

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config.json)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "INFO", "log level (DEBUG, INFO, WARN, ERROR, FATAL)")
	rootCmd.PersistentFlags().IntVar(&port, "port", 0, "port to listen on (0 for random port)")
}

var rootCmd = &cobra.Command{
	Use:   "validator",
	Short: "Validator service for BLS signing",
	Long:  `Validator service that performs BLS signing operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set log level
		logger.SetLevel(logger.GetLevelFromString(logLevel))
		logger.Info("Setting log level to %s", logLevel)

		// Load configuration
		cfg, err := config.LoadValidatorConfig(configFile, port)
		if err != nil {
			logger.Fatal("Failed to load configuration: %v", err)
		}

		// Start the server
		startServer(cfg)
	},
}

func startServer(cfg *config.ValidatorConfig) {
	// Start TCP server to accept connections
	var listener net.Listener
	var err error

	if cfg.Port == 0 {
		// Listen on any available port
		listener, err = net.Listen("tcp", ":0")
	} else {
		// Listen on the specified port
		listener, err = net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	}

	if err != nil {
		logger.Fatal("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	localAddr := listener.Addr().String()
	logger.Info("Validator listening on %s", localAddr)

	// Connect to dispatcher for heartbeat
	go connectToDispatcher(localAddr, cfg.Dispatcher)

	// Accept connections for signing requests
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Error accepting connection: %v", err)
			continue
		}
		go handleSigningRequest(conn)
	}
}

func connectToDispatcher(localAddr string, dispatcher string) {
	for {
		conn, err := net.Dial("tcp", dispatcher)
		if err != nil {
			logger.Error("Error connecting to dispatcher: %v", err)
			logger.Info("Retrying connection in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		logger.Info("Connected to dispatcher at %s", dispatcher)

		// Send local address as first message
		_, err = conn.Write([]byte("ADDR:" + localAddr))
		if err != nil {
			logger.Error("Failed to send local address: %v", err)
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
		logger.Warn("Connection lost, attempting to reconnect...")
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

	for {
		select {
		case <-ticker.C:
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
}

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
	logger.Info("Received signing request: %s", string(request))

	// simulate signature
	signature := fmt.Sprintf("Signed by validator: %s", string(request))

	// Send response
	_, err = conn.Write([]byte(signature))
	if err != nil {
		logger.Error("Error sending response: %v", err)
		return
	}

	logger.Debug("Sent response: %s", signature)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}