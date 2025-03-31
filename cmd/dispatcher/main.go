package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/crypto"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/store"
)

var (
	configFile string
	logLevel   string
	enableDB   bool
	keyFile    string
	keyPwd     string // password
	keyPair    *crypto.CombinedKeyPair
	validators struct {
		sync.RWMutex
		connections  map[net.Conn]bool
		addresses    map[net.Conn]string
		ethAddresses map[net.Conn]string
	}
	dbClient *store.DBClient
)

const (
	// MaxValidatorConnections is the maximum number of validator connections
	MaxValidatorConnections = 512
)

func init() {
	validators.connections = make(map[net.Conn]bool)
	validators.addresses = make(map[net.Conn]string)
	validators.ethAddresses = make(map[net.Conn]string)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config.json)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "INFO", "log level (DEBUG, INFO, WARN, ERROR, FATAL)")
	rootCmd.PersistentFlags().BoolVar(&enableDB, "enable-db", false, "enable database persistence")
	rootCmd.PersistentFlags().StringVar(&keyFile, "keys", "", "path to the key file")
	rootCmd.PersistentFlags().StringVar(&keyPwd, "key-password", "", "password for the key file")

	// Create run command
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the dispatcher service",
		Long:  `Run the dispatcher service that coordinates validators for BLS signing operations.`,
		Run:   runDispatcher,
	}

	// Add run command to root
	rootCmd.AddCommand(runCmd)
}

var rootCmd = &cobra.Command{
	Use:   "dispatcher",
	Short: "Dispatcher service for BLS signing",
	Long:  `Dispatcher service that coordinates validators for BLS signing operations.`,
	// Default behavior is to show help
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func runDispatcher(cmd *cobra.Command, args []string) {
	// Set log level
	logger.SetLevel(logger.GetLevelFromString(logLevel))
	logger.Info("Setting log level to %s", logLevel)

	// Load configuration
	cfg, err := config.LoadDispatcherConfig(configFile)
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
	}

	// Initialize database if enabled
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

		// Ensure we don't close the connection prematurely
		// Only close it when the application exits
		defer dbClient.Close()

		// Create database tables
		if err := dbClient.CreateDispatcherTables(); err != nil {
			logger.Fatal("Failed to create database tables: %v", err)
		}
		logger.Info("Database initialized successfully")
	}

	// Initialize gRPC client if endpoint is configured
	if cfg.ChainGRpcURL != "" {
		logger.Info("Initializing gRPC client connection to %s", cfg.ChainGRpcURL)
		if err := InitGRPCClient(cfg.ChainGRpcURL); err != nil {
			logger.Warn("Failed to initialize gRPC client: %v", err)
		} else {
			// Ensure connection is closed when program exits
			defer CloseGRPCClient()
		}
	}

	// Log reporting status
	logReportingStatus(cfg)

	// Start the server
	startServer(cfg)
}

func logReportingStatus(cfg *config.DispatcherConfig) {
	if cfg.EnableReport {
		logger.Info("BLS signature reporting is enabled")
		if cfg.ChainGRpcURL == "" {
			logger.Warn("Chain gRPC URL is not set, reporting will fail")
		}
		if cfg.ChainID == "" {
			logger.Warn("Chain ID is not set, reporting will fail")
		}
	} else {
		logger.Info("BLS signature reporting is disabled")
	}
}

func startServer(cfg *config.DispatcherConfig) {
	httpPort := fmt.Sprintf(":%d", cfg.HTTPPort)
	tcpPort := fmt.Sprintf(":%d", cfg.TCPPort)

	// Calculate Hetu address
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount(config.Bech32PrefixAccAddr, config.Bech32PrefixAccPub)

	// Set up HTTP handler
	http.HandleFunc("/reqblssign", func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, cfg)
	})

	// Start HTTP server
	go func() {
		logger.Info("Starting HTTP server on port %d", cfg.HTTPPort)
		if err := http.ListenAndServe(httpPort, nil); err != nil {
			logger.Fatal("Error starting HTTP server: %v", err)
		}
	}()

	// Start TCP server
	listener, err := net.Listen("tcp", tcpPort)
	if err != nil {
		logger.Fatal("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	logger.Info("The dispatcher is listening on TCP port %d, waiting for connections", cfg.TCPPort)

	// Accept validator connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Error accepting connection: %v", err)
			continue
		}
		go handleValidatorConnection(conn)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
