package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/proto/types"
)

var (
	grpcEndpoint string
	epochNum     uint64
	outputFormat string
	timeout      int

	// Global gRPC connection and client
	grpcConn    *grpc.ClientConn
	queryClient types.QueryClient
)

func init() {
	// Query BLS public keys command
	queryBlsKeysCmd := &cobra.Command{
		Use:     "query-bls-keys",
		Short:   "Query BLS public keys for a specific epoch",
		Long:    `Query the BLS public keys of validators for a specific epoch using gRPC.`,
		Example: `./build/dispatcher query-bls-keys --epoch=1 --grpc-endpoint=localhost:9000`,
		Run:     queryBlsKeys,
	}

	queryBlsKeysCmd.Flags().StringVar(&grpcEndpoint, "grpc-endpoint", "", "gRPC endpoint (host:port)")
	queryBlsKeysCmd.Flags().Uint64Var(&epochNum, "epoch", 0, "Epoch number to query")
	queryBlsKeysCmd.Flags().StringVar(&outputFormat, "output", "json", "Output format (json or text)")
	queryBlsKeysCmd.Flags().IntVar(&timeout, "timeout", 10, "Timeout in seconds for the gRPC call")

	// Add command to root
	rootCmd.AddCommand(queryBlsKeysCmd)
}

// InitGRPCClient initializes the gRPC client connection
func InitGRPCClient(endpoint string) error {
	if endpoint == "" {
		return fmt.Errorf("gRPC endpoint is empty")
	}

	// If we already have a connection, don't create a new one
	if grpcConn != nil && queryClient != nil {
		return nil
	}

	var err error
	// Close existing connection if any
	if grpcConn != nil {
		grpcConn.Close()
	}

	// Create new connection
	grpcConn, err = grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	// Create client
	queryClient = types.NewQueryClient(grpcConn)
	logger.Info("gRPC client initialized to endpoint: %s", endpoint)
	return nil
}

// CloseGRPCClient closes the gRPC client connection
func CloseGRPCClient() {
	if grpcConn != nil {
		grpcConn.Close()
		grpcConn = nil
	}
}

func queryBlsKeys(cmd *cobra.Command, args []string) {
	// Set log level
	logger.SetLevel(logger.GetLevelFromString(logLevel))

	// Load configuration if not provided via flags
	if grpcEndpoint == "" && configFile != "" {
		v := viper.New()
		v.SetConfigFile(configFile)
		if err := v.ReadInConfig(); err != nil {
			logger.Fatal("Failed to read config file: %v", err)
		}
		grpcEndpoint = v.GetString("grpc_endpoint")
	}

	// Validate required parameters
	if grpcEndpoint == "" {
		logger.Fatal("gRPC endpoint is required. Use --grpc-endpoint flag or set in config file.")
	}

	// Initialize gRPC client
	if err := InitGRPCClient(grpcEndpoint); err != nil {
		logger.Fatal("Failed to initialize gRPC client: %v", err)
	}
	defer CloseGRPCClient()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// Prepare request
	request := &types.QueryBlsPublicKeyListRequest{
		EpochNum: epochNum,
	}

	// Make the call
	logger.Info("Querying BLS public keys for epoch %d", epochNum)
	response, err := queryClient.BlsPublicKeyList(ctx, request)
	if err != nil {
		logger.Fatal("Failed to query BLS public keys: %v", err)
	}

	// Process response
	if outputFormat == "json" {
		// Output as JSON
		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			logger.Fatal("Failed to marshal response to JSON: %v", err)
		}
		fmt.Println(string(jsonData))
	} else {
		// Output as text
		fmt.Printf("BLS Public Keys for Epoch %d:\n", epochNum)
		fmt.Println("-----------------------------------")

		if len(response.ValidatorWithBlsKeys) == 0 {
			fmt.Println("No validators found for this epoch.")
		} else {
			for i, validator := range response.ValidatorWithBlsKeys {
				fmt.Printf("Validator #%d:\n", i+1)
				fmt.Printf("  Address: %s\n", validator.ValidatorAddress)
				fmt.Printf("  BLS Public Key: %s\n", validator.BlsPubKeyHex)
				fmt.Printf("  Voting Power: %s\n", strconv.FormatUint(validator.VotingPower, 10))
				fmt.Println("-----------------------------------")
			}
		}

		if response.Pagination != nil {
			fmt.Printf("Total: %d validators\n", response.Pagination.Total)
		}
	}

	logger.Info("Query completed successfully")
}
