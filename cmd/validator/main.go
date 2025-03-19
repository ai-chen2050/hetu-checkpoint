package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"

	"encoding/hex"

	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/crypto"
	"github.com/hetu-project/hetu-checkpoint/crypto/bls12381"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/store"
)

var (
	configFile string
	logLevel   string
	port       int
	enableDB   bool
	dbClient   *store.DBClient
	keyFile    string
	keyPwd     string // password
	keyPair    *crypto.CombinedKeyPair
)

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config.json)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "INFO", "log level (DEBUG, INFO, WARN, ERROR, FATAL)")
	rootCmd.PersistentFlags().IntVar(&port, "port", 0, "port to listen on (0 for random port)")
	rootCmd.PersistentFlags().BoolVar(&enableDB, "enable-db", false, "enable database persistence")
	rootCmd.PersistentFlags().StringVar(&keyFile, "keys", "", "path to the key file")
	rootCmd.PersistentFlags().StringVar(&keyPwd, "key-password", "", "password for the key file")

	// Create run command
	runCmd := &cobra.Command{
		Use:     "run",
		Short:   "Run the validator service",
		Long:    `Run the validator service that performs BLS signing operations.`,
		Example: `./build/validator run --config docs/config/val_config.json --enable-db --keys=keys/validator.json`,
		Run:     runValidator,
	}

	// Add run command to root
	rootCmd.AddCommand(runCmd)
}

var rootCmd = &cobra.Command{
	Use:     "validator",
	Short:   "Validator service for BLS signing",
	Long:    `Validator service that performs BLS signing operations.`,
	Example: `./build/validator --help`,
	// Default behavior is to show help
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func runValidator(cmd *cobra.Command, args []string) {
	// Set log level
	logger.SetLevel(logger.GetLevelFromString(logLevel))
	logger.Info("Setting log level to %s", logLevel)

	// Load configuration
	cfg, err := config.LoadValidatorConfig(configFile, port)
	if err != nil {
		logger.Fatal("Failed to load configuration: %v", err)
	}

	// Load key pair if specified
	if keyFile != "" {
		logger.Info("Loading key pair from %s", keyFile)
		keyPair, err = crypto.LoadKeyPair(keyFile, keyPwd)
		if err != nil {
			logger.Fatal("Failed to load key pair: %v", err)
		}
		logger.Info("Loaded key pair with Ethereum address: %s", keyPair.ETH.Address)
	} else {
		logger.Fatal("Key file must be specified with --keys flag")
	}

	// Initialize database client only if enabled
	if enableDB {
		logger.Info("Database persistence enabled, initializing connection...")
		dbClient, err = store.NewDBClient(store.Config{
			Host:     cfg.DBHost,
			Port:     cfg.DBPort,
			User:     cfg.DBUser,
			Password: cfg.DBPassword,
			DBName:   cfg.DBName,
		})
		if err != nil {
			logger.Fatal("Failed to initialize database client: %v", err)
		}
		defer dbClient.Close()

		// Create database tables
		if err := dbClient.CreateValidatorTables(); err != nil {
			logger.Fatal("Failed to create database tables: %v", err)
		}
		logger.Info("Database initialized successfully")
	}

	// Start the server
	startServer(cfg)
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
	go connectToDispatcher(localAddr, cfg.DispatcherTcp)

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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
