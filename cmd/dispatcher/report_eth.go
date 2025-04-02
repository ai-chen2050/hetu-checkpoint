package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hetu-project/hetu-checkpoint/config"
	CKPTValStaking "github.com/hetu-project/hetu-checkpoint/contracts/ckpt_val_staking"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/proto/types"
)

// ReportCheckpointToEthereum reports the aggregated checkpoint to Ethereum
func ReportCheckpointToEthereum(aggregatedCkpt *types.RawCheckpointWithMeta, cfg *config.DispatcherConfig) error {
	// Check if key pair is loaded
	if keyPair == nil {
		return fmt.Errorf("key pair not loaded, cannot report to Ethereum")
	}

	// Check if contract address is configured
	if cfg.ValidatorStakingAddress == "" {
		return fmt.Errorf("validator staking contract address not configured")
	}

	// Check if Ethereum RPC URL is configured
	if cfg.EthRpcURL == "" {
		return fmt.Errorf("Ethereum RPC URL not configured")
	}

	// Connect to Ethereum client
	client, err := ethclient.Dial(cfg.EthRpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	// Create contract instance
	contractAddress := common.HexToAddress(cfg.ValidatorStakingAddress)
	stakingContract, err := CKPTValStaking.NewCKPTValStaking(contractAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate contract: %v", err)
	}

	privKey, err := crypto.HexToECDSA(keyPair.ETH.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}
	// Prepare transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(cfg.EthChainID))
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	// Set gas price and limit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = 3000000 // Set an appropriate gas limit

	// Convert checkpoint data for contract call
	epochNum := aggregatedCkpt.Ckpt.EpochNum
	blockHash := common.BytesToHash(*aggregatedCkpt.Ckpt.BlockHash)
	powerSum := aggregatedCkpt.PowerSum

	var blsMultiSig []byte
	if aggregatedCkpt.Ckpt.BlsMultiSig != nil {
		blsMultiSig = *aggregatedCkpt.Ckpt.BlsMultiSig
	}

	var blsAggrPk []byte
	if aggregatedCkpt.BlsAggrPk != nil {
		blsAggrPk = *aggregatedCkpt.BlsAggrPk
	}

	bitmap := aggregatedCkpt.Ckpt.Bitmap

	// Log the checkpoint data being submitted
	logger.Info("Submitting checkpoint to Ethereum, Epoch: %d", epochNum)

	// Submit checkpoint to contract
	tx, err := stakingContract.SubmitCheckpoint(
		auth,
		epochNum,
		blockHash,
		bitmap,
		blsMultiSig,
		blsAggrPk,
		powerSum,
	)
	if err != nil {
		return fmt.Errorf("failed to submit checkpoint: %v", err)
	}

	logger.Info("Checkpoint submitted to Ethereum, transaction hash: %s", tx.Hash().Hex())

	// Wait for transaction to be mined
	logger.Info("Waiting for transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	logger.Info("Checkpoint successfully reported to Ethereum")
	return nil
}

// ReportCheckpointWithRetry attempts to report the checkpoint with retries
func ReportCheckpointWithRetry(aggregatedCkpt *types.RawCheckpointWithMeta, cfg *config.DispatcherConfig, maxRetries int) {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = ReportCheckpointToEthereum(aggregatedCkpt, cfg)
		if err == nil {
			return
		}

		logger.Error("Failed to report checkpoint to Ethereum (attempt %d/%d): %v", i+1, maxRetries, err)

		// Wait before retrying
		if i < maxRetries-1 {
			retryDelay := time.Duration(2<<uint(i)) * time.Second // Exponential backoff
			logger.Info("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	logger.Error("Failed to report checkpoint to Ethereum after %d attempts: %v", maxRetries, err)
}

// DistributeRewards distributes rewards for a specific epoch
func DistributeRewards(epochNum uint64, cfg *config.DispatcherConfig) error {
	// Check if key pair is loaded
	if keyPair == nil {
		return fmt.Errorf("key pair not loaded, cannot distribute rewards")
	}

	// Check if contract address is configured
	if cfg.ValidatorStakingAddress == "" {
		return fmt.Errorf("validator staking contract address not configured")
	}

	// Check if Ethereum RPC URL is configured
	if cfg.EthRpcURL == "" {
		return fmt.Errorf("Ethereum RPC URL not configured")
	}

	// Check if rewards for this epoch have already been distributed
	if enableDB {
		distributed, err := dbClient.IsEpochRewardDistributed(epochNum)
		if err != nil {
			logger.Warn("Failed to check if rewards for epoch %d were distributed: %v", epochNum, err)
		} else if distributed {
			logger.Info("Rewards for epoch %d have already been distributed", epochNum)
			return nil
		}
	}

	// Connect to Ethereum client
	client, err := ethclient.Dial(cfg.EthRpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	// Create contract instance
	contractAddress := common.HexToAddress(cfg.ValidatorStakingAddress)
	stakingContract, err := CKPTValStaking.NewCKPTValStaking(contractAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate contract: %v", err)
	}

	privKey, err := crypto.HexToECDSA(keyPair.ETH.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	// Prepare transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(cfg.EthChainID))
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	// Set gas price and limit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = 5000000 // Set a higher gas limit for distribution

	// Log the distribution
	logger.Info("Distributing rewards for epoch %d", epochNum)

	// Call distributeCheckpointRewards
	tx, err := stakingContract.DistributeCheckpointRewards(auth, epochNum)
	if err != nil {
		return fmt.Errorf("failed to distribute rewards: %v", err)
	}

	logger.Info("Reward distribution transaction submitted, hash: %s", tx.Hash().Hex())

	// Wait for transaction to be mined
	logger.Info("Waiting for transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	// Record the distribution in the database
	if enableDB {
		_, err := dbClient.InsertRewardDistribution(
			epochNum,
			tx.Hash().Hex(),
			"SUCCESS",
		)
		if err != nil {
			logger.Error("Failed to record reward distribution: %v", err)
		}
	}

	logger.Info("Rewards successfully distributed for epoch %d", epochNum)
	return nil
}

// DistributeRewardsWithRetry attempts to distribute rewards with retries
func DistributeRewardsWithRetry(epochNum uint64, cfg *config.DispatcherConfig, maxRetries int) {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = DistributeRewards(epochNum, cfg)
		if err == nil {
			return
		}

		logger.Error("Failed to distribute rewards for epoch %d (attempt %d/%d): %v", epochNum, i+1, maxRetries, err)

		// Wait before retrying
		if i < maxRetries-1 {
			retryDelay := time.Duration(2<<uint(i)) * time.Second // Exponential backoff
			logger.Info("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	logger.Error("Failed to distribute rewards for epoch %d after %d attempts: %v", epochNum, maxRetries, err)
}
