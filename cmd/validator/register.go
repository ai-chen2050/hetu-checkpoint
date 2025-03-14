package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	CKPTValStaking "github.com/hetu-project/hetu-checkpoint/contracts/ckpt_val_staking"
	erc20 "github.com/hetu-project/hetu-checkpoint/contracts/erc20"
	"github.com/hetu-project/hetu-checkpoint/crypto"
	"github.com/hetu-project/hetu-checkpoint/logger"
)

var (
	rpcURL            string
	stakingAmount     string
	dispatcherURL     string
	stakingTokenAddr  string
	valStakingAddr    string
	gasPrice          int64
	gasLimit          uint64
	waitConfirmations uint64
)

func init() {
	registerCmd := &cobra.Command{
		Use:   "register-and-stake",
		Short: "Register validator and stake tokens",
		Long:  `Register validator in the staking contract and stake tokens.`,
		Example: `validator register-and-stake --rpc-url http://localhost:8545 --amount 100 --dispatcher-url http://localhost:8080`,
		Run:   registerAndStake,
	}

	registerCmd.Flags().StringVar(&rpcURL, "rpc-url", "", "Ethereum RPC URL")
	registerCmd.Flags().StringVar(&stakingAmount, "amount", "100", "Amount to stake (in tokens)")
	registerCmd.Flags().StringVar(&dispatcherURL, "dispatcher-url", "", "Dispatcher URL to register")
	registerCmd.Flags().StringVar(&stakingTokenAddr, "token-address", "", "Staking token contract address")
	registerCmd.Flags().StringVar(&valStakingAddr, "staking-address", "", "Validator staking contract address")
	registerCmd.Flags().Int64Var(&gasPrice, "gas-price", 0, "Gas price in gwei (0 for auto)")
	registerCmd.Flags().Uint64Var(&gasLimit, "gas-limit", 3000000, "Gas limit for transactions")
	registerCmd.Flags().Uint64Var(&waitConfirmations, "confirmations", 1, "Number of confirmations to wait")

	rootCmd.AddCommand(registerCmd)
}

func registerAndStake(cmd *cobra.Command, args []string) {
	// Set log level
	logger.SetLevel(logger.GetLevelFromString(logLevel))

	// Load configuration if not provided via flags
	logger.Info("Loading file: %s", configFile)
	if configFile != "" {
		v := viper.New()
		v.SetConfigFile(configFile)
		if err := v.ReadInConfig(); err != nil {
			logger.Fatal("Failed to read config file: %v", err)
		}

		// Load values from config if not set via flags
		if rpcURL == "" {
			rpcURL = v.GetString("ChainRpcURL")
		}
		if stakingTokenAddr == "" {
			stakingTokenAddr = v.GetString("StakingTokenAddress")
		}
		if valStakingAddr == "" {
			valStakingAddr = v.GetString("ValidatorStakingAddress")
		}
		if dispatcherURL == "" {
			dispatcherURL = v.GetString("DispatcherURL")
		}
	}

	// log
	logger.Info("Using RPC URL: %s", rpcURL)
	logger.Info("Using Staking Token Address: %s", stakingTokenAddr)
	logger.Info("Using Validator Staking Address: %s", valStakingAddr)
	logger.Info("Using Dispatcher URL: %s", dispatcherURL)

	// Validate required parameters
	if rpcURL == "" {
		logger.Fatal("RPC URL is required")
	}
	if stakingTokenAddr == "" {
		logger.Fatal("Staking token address is required")
	}
	if valStakingAddr == "" {
		logger.Fatal("Validator staking address is required")
	}
	if dispatcherURL == "" {
		logger.Fatal("Dispatcher URL is required")
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

	// Connect to Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		logger.Fatal("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	// Create private key for transactions
	privateKeyHex := keyPair.ETH.PrivateKey
	privateKey, err := ethcrypto.HexToECDSA(privateKeyHex)
	if err != nil {
		logger.Fatal("Failed to parse private key: %v", err)
	}

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		logger.Fatal("Failed to get chain ID: %v", err)
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		logger.Fatal("Failed to create transactor: %v", err)
	}

	// Set gas price if provided
	if gasPrice > 0 {
		auth.GasPrice = big.NewInt(gasPrice * 1e9) // Convert gwei to wei
	}
	auth.GasLimit = gasLimit

	// Parse contract addresses
	tokenAddress := common.HexToAddress(stakingTokenAddr)
	stakingAddress := common.HexToAddress(valStakingAddr)

	// Create contract instances
	tokenContract, err := erc20.NewERC20(tokenAddress, client)
	if err != nil {
		logger.Fatal("Failed to instantiate token contract: %v", err)
	}

	stakingContract, err := CKPTValStaking.NewCKPTValStaking(stakingAddress, client)
	if err != nil {
		logger.Fatal("Failed to instantiate staking contract: %v", err)
	}

	// Get token decimals
	decimals, err := tokenContract.Decimals(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to get token decimals: %v", err)
	}
	logger.Info("Token decimals: %d", decimals)

	// Parse staking amount
	amount, ok := new(big.Int).SetString(stakingAmount, 10)
	if !ok {
		logger.Fatal("Failed to parse staking amount")
	}

	// Convert to token units with decimals
	decimalFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	amount = amount.Mul(amount, decimalFactor)
	logger.Info("Staking amount (with decimals): %s", amount.String())

	// Check token balance
	balance, err := tokenContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(keyPair.ETH.Address))
	if err != nil {
		logger.Fatal("Failed to get token balance: %v", err)
	}
	logger.Info("Token balance: %s", balance.String())

	if balance.Cmp(amount) < 0 {
		logger.Fatal("Insufficient token balance. Have %s, need %s", balance.String(), amount.String())
	}

	// Approve tokens for staking
	logger.Info("Approving tokens for staking...")
	tx, err := tokenContract.Approve(auth, stakingAddress, amount)
	if err != nil {
		logger.Fatal("Failed to approve tokens: %v", err)
	}
	logger.Info("Approval transaction sent: %s", tx.Hash().Hex())

	// Wait for approval transaction to be mined
	logger.Info("Waiting for approval transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Failed to wait for approval transaction: %v", err)
	}
	if receipt.Status == 0 {
		logger.Fatal("Approval transaction failed")
	}
	logger.Info("Approval transaction confirmed")

	// Stake tokens
	logger.Info("Staking tokens...")
	tx, err = stakingContract.Stake(auth, amount)
	if err != nil {
		logger.Fatal("Failed to stake tokens: %v", err)
	}
	logger.Info("Staking transaction sent: %s", tx.Hash().Hex())

	// Wait for staking transaction to be mined
	logger.Info("Waiting for staking transaction to be mined...")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Failed to wait for staking transaction: %v", err)
	}
	if receipt.Status == 0 {
		logger.Fatal("Staking transaction failed")
	}
	logger.Info("Staking transaction confirmed")

	// Update validator info
	logger.Info("Updating validator info...")
	tx, err = stakingContract.UpdateValidatorInfo(auth, dispatcherURL, keyPair.BLS.PublicKey)
	if err != nil {
		logger.Fatal("Failed to update validator info: %v", err)
	}
	logger.Info("Update validator info transaction sent: %s", tx.Hash().Hex())

	// Wait for update transaction to be mined
	logger.Info("Waiting for update validator info transaction to be mined...")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Failed to wait for update validator info transaction: %v", err)
	}
	if receipt.Status == 0 {
		logger.Fatal("Update validator info transaction failed")
	}
	logger.Info("Update validator info transaction confirmed")

	logger.Info("Validator successfully registered and staked!")
	logger.Info("Validator address: %s", keyPair.ETH.Address)
	logger.Info("BLS public key: %s", keyPair.BLS.PublicKey)
	logger.Info("Dispatcher URL: %s", dispatcherURL)
	logger.Info("Staked amount: %s tokens", stakingAmount)
}
