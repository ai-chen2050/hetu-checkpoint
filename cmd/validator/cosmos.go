package main

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	// sdkcrypto "github.com/cosmos/cosmos-sdk/crypto"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hetu-project/hetu-checkpoint/crypto"
	"github.com/hetu-project/hetu-checkpoint/crypto/bls12381"
	"github.com/hetu-project/hetu-checkpoint/crypto/ethsecp256k1"
	hdw "github.com/hetu-project/hetu-checkpoint/crypto/hd"
	"github.com/hetu-project/hetu-checkpoint/encoding"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/proto/types"

	"github.com/ethereum/go-ethereum/common"
)

var (
	chainGRpcURL      string
	chainID           string
	cosmosGasLimit    uint64
	cosmosGasPrice    string
	cosmosMemo        string
	cosmosTimeoutSecs int

	// Global Cosmos client
	cosmosConn   *grpc.ClientConn
	cosmosMsgSvc types.MsgClient
)

const (
	// Bech32Prefix defines the Bech32 prefix used for EthAccounts
	Bech32Prefix = "hetu"

	// keyName is the name of the key in the keyring
	keyName = "validator"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32Prefix

	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32Prefix + sdk.PrefixPublic
)

func init() {
	// Register validator command
	registerValidatorCmd := &cobra.Command{
		Use:     "register-validator",
		Short:   "Register validator on Cosmos chain",
		Long:    `Register validator on Cosmos chain by submitting a transaction to the RegistValidator RPC endpoint.`,
		Example: `./build/validator register-validator --cosmos-rpc=localhost:9090 --chain-id=hetu_7878-1`,
		Run:     registerValidator,
	}

	registerValidatorCmd.Flags().StringVar(&chainGRpcURL, "cosmos-rpc", "localhost:9090", "Cosmos gRPC endpoint")
	registerValidatorCmd.Flags().StringVar(&chainID, "chain-id", "hetu_560002-1", "Cosmos chain ID")
	registerValidatorCmd.Flags().Uint64Var(&cosmosGasLimit, "gas", 200000, "Gas limit for the transaction")
	registerValidatorCmd.Flags().StringVar(&cosmosGasPrice, "gas-price", "0.025uhetu", "Gas price for the transaction")
	registerValidatorCmd.Flags().StringVar(&cosmosMemo, "memo", "", "Memo for the transaction")
	registerValidatorCmd.Flags().IntVar(&cosmosTimeoutSecs, "timeout", 30, "Timeout in seconds for the transaction")

	// Add command to root
	rootCmd.AddCommand(registerValidatorCmd)
}

// InitCosmosClient initializes the Cosmos gRPC client
func InitCosmosClient(endpoint string) error {
	if endpoint == "" {
		return fmt.Errorf("Cosmos gRPC endpoint is empty")
	}

	var err error
	// Close existing connection if any
	if cosmosConn != nil {
		cosmosConn.Close()
	}

	// Create new connection
	cosmosConn, err = grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Cosmos gRPC server: %v", err)
	}

	// Create client
	cosmosMsgSvc = types.NewMsgClient(cosmosConn)
	return nil
}

// CloseCosmosClient closes the Cosmos gRPC client connection
func CloseCosmosClient() {
	if cosmosConn != nil {
		cosmosConn.Close()
		cosmosConn = nil
	}
}

// Convert Ethereum private key to Cosmos secp256k1 private key
func ethPrivKeyFromHex(ethPrivKeyHex string) (cryptotypes.PrivKey, error) {
	privKey := &ethsecp256k1.PrivKey{
		Key: common.FromHex(ethPrivKeyHex),
	}

	return privKey, nil
}

// Get Hetu address from private key
func getHetuAddressFromPrivKey(privKey cryptotypes.PrivKey) string {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	return sdk.AccAddress(privKey.PubKey().Address()).String()
}

func registerValidator(cmd *cobra.Command, args []string) {
	// Set log level
	logger.SetLevel(logger.GetLevelFromString(logLevel))

	// Load configuration if not provided via flags
	if configFile != "" {
		v := viper.New()
		v.SetConfigFile(configFile)
		if err := v.ReadInConfig(); err != nil {
			logger.Fatal("Failed to read config file: %v", err)
		}

		// Load values from config if not set via flags
		if chainGRpcURL == "" {
			chainGRpcURL = v.GetString("chain_grpc_url")
		}
		if chainID == "" {
			chainID = v.GetString("cosmos_chain_id")
		}
	}

	// Validate required parameters
	if chainGRpcURL == "" {
		logger.Fatal("Cosmos gRPC endpoint is required. Use --cosmos-rpc flag or set in config file.")
	}
	if chainID == "" {
		logger.Fatal("Cosmos chain ID is required. Use --chain-id flag or set in config file.")
	}
	if keyFile == "" {
		logger.Fatal("Key file must be specified with --keys flag")
	}

	// Load key pair
	logger.Info("Loading key pair from %s", keyFile)
	keyPair, err := crypto.LoadKeyPair(keyFile, keyPwd)
	if err != nil {
		logger.Fatal("Failed to load key pair: %v", err)
	}
	logger.Info("Loaded key pair with Ethereum address: %s", keyPair.ETH.Address)

	// Initialize Cosmos client
	if err := InitCosmosClient(chainGRpcURL); err != nil {
		logger.Fatal("Failed to initialize Cosmos client: %v", err)
	}
	defer CloseCosmosClient()

	// Convert Ethereum private key from hex string
	hetuPrivKey, err := ethPrivKeyFromHex(keyPair.ETH.PrivateKey)
	if err != nil {
		logger.Fatal("Failed to convert Ethereum private key to Cosmos private key: %v", err)
	}

	// Get Hetu address
	hetuAddress := getHetuAddressFromPrivKey(hetuPrivKey)
	logger.Info("Hetu address: %s", hetuAddress)

	// Parse BLS public key
	blsPubKeyBytes, err := hex.DecodeString(keyPair.BLS.PublicKey)
	if err != nil {
		logger.Fatal("Failed to decode BLS public key: %v", err)
	}

	// Create BLS public key object
	blsPubKey := new(bls12381.BlsPubKey)
	pubkeyBytes := blsPubKey.Deserialize(blsPubKeyBytes).Compress()
	if pubkeyBytes == nil {
		logger.Fatal("Failed to deserialize BLS public key: %v", err)
	}
	pubkey := new(bls12381.PublicKey)
	pubkey.Unmarshal(pubkeyBytes)

	cfg := encoding.MakeConfig()

	// use memory keyring
	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendMemory, "", nil, cfg.Codec)
	if err != nil {
		logger.Fatal("Failed to create keyring: %v", err)
	}

	// Create client context with keyring
	clientCtx := client.Context{
		ChainID:          chainID,
		GRPCClient:       cosmosConn,
		TxConfig:         cfg.TxConfig,
		Codec:            cfg.Codec,
		AccountRetriever: authtypes.AccountRetriever{},
		FromAddress:      sdk.MustAccAddressFromBech32(hetuAddress),
		Keyring:          kb,
		KeyringOptions:   []keyring.Option{hdw.EthSecp256k1Option()},
		FromName:         keyName,
	}

	// Import private key to keyring
	// armor := sdkcrypto.EncryptArmorPrivKey(hetuPrivKey, "", "eth_secp256k1")
	err = clientCtx.Keyring.ImportPrivKeyHex(keyName, keyPair.ETH.PrivateKey, "eth_secp256k1")
	// err = clientCtx.Keyring.ImportPrivKey(keyName, armor, "")
	if err != nil {
		logger.Fatal("Failed to import private key: %v", err)
	}

	// Check if account exists
	// account, err := clientCtx.AccountRetriever.GetAccount(clientCtx, sdk.MustAccAddressFromBech32(cosmosAddress))
	// if err != nil {
	// 	logger.Fatal("Failed to retrieve account: %v", err)
	// }
	// if account == nil {
	// 	logger.Fatal("Account %s not found", cosmosAddress)
	// }

	// Create the message
	msg := &types.MsgRegistValidator{
		BlsPubkey:        pubkey,
		ValidatorAddress: keyPair.ETH.Address,
		Sender:           hetuAddress,
	}

	// Set gas parameters
	gasAdjustment := float64(1.5)
	fees, err := sdk.ParseCoinsNormalized(cosmosGasPrice)
	if err != nil {
		logger.Fatal("Failed to parse gas price: %v", err)
	}

	// Create transaction factory
	txFactory := tx.Factory{}.
		WithChainID(chainID).
		WithGas(cosmosGasLimit).
		WithGasAdjustment(gasAdjustment).
		WithFees(fees.String()).
		WithMemo(cosmosMemo).
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
		WithTxConfig(clientCtx.TxConfig).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithKeybase(clientCtx.Keyring)

	// Broadcast the transaction
	logger.Info("Broadcasting transaction to register validator...")
	err = tx.GenerateOrBroadcastTxWithFactory(clientCtx, txFactory, msg)
	if err != nil {
		logger.Fatal("Failed to broadcast transaction: %v", err)
	}

	logger.Info("Transaction broadcast successful!")
	// logger.Info("Transaction hash: %s", resp.TxHash)
	// logger.Info("Gas used: %d", resp.GasUsed)
	// logger.Info("Height: %d", resp.Height)
}
