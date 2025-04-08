package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/crypto"
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
		initializeDatabase(cfg)
		defer dbClient.Close()
	}

	// Start the validator service
	startValidatorService(cfg)
}

func startValidatorService(cfg *config.ValidatorConfig) {
	// Create a channel to signal when the listening server is ready
	serverReady := make(chan struct{})

	// Start TCP server in a goroutine
	go startListeningServer(cfg, serverReady)

	// Wait for the server to be ready before connecting to dispatcher
	<-serverReady

	// Connect to dispatcher and maintain connection
	maintainDispatcherConnection(cfg)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
