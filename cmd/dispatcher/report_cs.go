package main

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/crypto/ethsecp256k1"
	hdw "github.com/hetu-project/hetu-checkpoint/crypto/hd"
	"github.com/hetu-project/hetu-checkpoint/encoding"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/proto/types"
)

var (
	// Global Cosmos client for reporting
	reportConn   *grpc.ClientConn
	reportMsgSvc types.MsgClient
)

// InitReportClient initializes the Cosmos gRPC client for reporting
func InitReportClient(endpoint string) error {
	if endpoint == "" {
		return fmt.Errorf("Cosmos gRPC endpoint is empty")
	}

	var err error
	// Close existing connection if any
	if reportConn != nil {
		reportConn.Close()
	}

	// Create new connection
	reportConn, err = grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Cosmos gRPC server: %v", err)
	}

	// Create client
	reportMsgSvc = types.NewMsgClient(reportConn)
	return nil
}

// CloseReportClient closes the Cosmos gRPC client connection
func CloseReportClient() {
	if reportConn != nil {
		reportConn.Close()
		reportConn = nil
	}
}

// ReportBLSSignaturesByCosmosTx reports BLS signatures to the chain
func ReportBLSSignaturesByCosmosTx(validResponses map[string][]byte, req *config.Request, cfg *config.DispatcherConfig) {
	// Use the epoch number from the checkpoint
	epochNum := req.Checkpoint.EpochNum
	if epochNum == 0 {
		logger.Warn("Epoch number is 0 in the request, using default value 1")
		epochNum = 1
	}

	logger.Info("Reporting BLS signatures for epoch %d", epochNum)

	// Convert binary signatures to hex strings and create AddrSigs array instead of map
	var addrSigs []*types.AddrSig
	for ethAddr, sig := range validResponses {
		addrSigs = append(addrSigs, &types.AddrSig{
			Address:   ethAddr,
			Signature: hex.EncodeToString(sig),
		})
	}

	// Get Hetu address from private key
	if keyPair == nil {
		logger.Error("Key pair not loaded, cannot report BLS signatures")
		return
	}

	if cfg.ChainID == "" {
		logger.Error("Cosmos chain ID is required. Use --chain-id flag or set in config file.")
		return
	}
	if cfg.CometBFTSvr == "" {
		logger.Error("Comet BFT gRPC endpoint is required. Use --comet-bft-svr flag or set in config file.")
		return
	}

	// Create a new gRPC connection for this transaction
	conn, err := grpc.NewClient(cfg.ChainGRpcURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Failed to connect to Cosmos gRPC server: %v", err)
		return
	}
	// Ensure connection is closed when function returns
	defer conn.Close()

	// Use the stored Hetu address instead of recalculating it
	hetuAddress := keyPair.HetuAddress

	// If HetuAddress is not available in the keyPair (for backward compatibility)
	if hetuAddress == "" {
		// Convert Ethereum private key to Cosmos private key
		privKey := &ethsecp256k1.PrivKey{
			Key: common.FromHex(keyPair.ETH.PrivateKey),
		}
		hetuAddress = sdk.AccAddress(privKey.PubKey().Address()).String()

		// Store it for future use
		keyPair.HetuAddress = hetuAddress
	}

	// Create encoding config
	encodingCfg := encoding.MakeConfig()

	// Use memory keyring
	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendMemory, "", nil, encodingCfg.Codec, hdw.EthSecp256k1Option())
	if err != nil {
		logger.Error("Failed to create keyring: %v", err)
		return
	}

	// Import private key to keyring
	err = kb.ImportPrivKeyHex(config.KeyName, keyPair.ETH.PrivateKey, "eth_secp256k1")
	if err != nil {
		logger.Error("Failed to import private key: %v", err)
		return
	}

	rpcClient, err := client.NewClientFromNode("tcp://" + cfg.CometBFTSvr)
	if err != nil {
		logger.Error("Failed to create RPC client: %v", err)
		return
	}

	// Create client context with the new connection
	clientCtx := client.Context{
		ChainID:           cfg.ChainID,
		GRPCClient:        conn,
		Client:            rpcClient,
		TxConfig:          encodingCfg.TxConfig,
		Codec:             encodingCfg.Codec,
		InterfaceRegistry: encodingCfg.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		FromAddress:       sdk.MustAccAddressFromBech32(hetuAddress),
		Keyring:           kb,
		KeyringOptions:    []keyring.Option{hdw.EthSecp256k1Option()},
		FromName:          config.KeyName,
		SkipConfirm:       true,
		BroadcastMode:     flags.BroadcastSync, // Use sync mode for better reliability
	}

	// Create the message with array of AddrSig instead of map
	msg := &types.MsgBLSCallback{
		EpochNum: epochNum,
		AddrSigs: addrSigs,
		Sender:   hetuAddress,
	}

	// Set gas parameters
	gasAdjustment := float64(1.5)
	fees, err := sdk.ParseCoinsNormalized("10gas")
	if err != nil {
		logger.Error("Failed to parse gas price: %v", err)
		return
	}

	// Create transaction factory
	txFactory := tx.Factory{}.
		WithChainID(cfg.ChainID).
		WithGas(2000000).
		WithGasAdjustment(gasAdjustment).
		WithGasPrices(fees.String()).
		WithMemo("").
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
		WithTxConfig(clientCtx.TxConfig).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithKeybase(clientCtx.Keyring)

	// Broadcast the transaction
	logger.Info("Broadcasting transaction to report BLS signatures...")
	err = tx.GenerateOrBroadcastTxWithFactory(clientCtx, txFactory, msg)
	if err != nil {
		logger.Error("Failed to broadcast transaction: %v", err)
		return
	}

	logger.Info("BLS signatures reported successfully for epoch %d", epochNum)
}
