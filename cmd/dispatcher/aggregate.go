package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/proto/types"
)

// AggregateSignatures aggregates BLS signatures from validators
func AggregateSignatures(validResponses map[string][]byte, req *config.Request, cfg *config.DispatcherConfig) (*types.RawCheckpointWithMeta, error) {
	if cfg.ChainGRpcURL == "" {
		return nil, fmt.Errorf("chain gRPC URL is not set")
	}

	// Initialize gRPC client if not already initialized
	if queryClient == nil {
		err := InitGRPCClient(cfg.ChainGRpcURL)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize gRPC client: %v", err)
		}
	}

	// Create context with timeout - increase timeout to 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Query validator set for the current epoch
	epochNum := req.CheckpointWithMeta.Ckpt.EpochNum
	logger.Info("Querying validator set for epoch %d", epochNum)

	valSetResp, err := queryClient.BlsPublicKeyList(ctx, &types.QueryBlsPublicKeyListRequest{
		EpochNum: epochNum,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query validator set: %v", err)
	}

	if len(valSetResp.ValidatorWithBlsKeys) == 0 {
		return nil, fmt.Errorf("no validators found for epoch %d", epochNum)
	}

	// Calculate total power
	var totalPower uint64 = 0
	for _, val := range valSetResp.ValidatorWithBlsKeys {
		totalPower += uint64(val.VotingPower)
	}

	// Convert to types.ValidatorSet
	validatorSet := make(types.ValidatorSet, len(valSetResp.ValidatorWithBlsKeys))
	for i, val := range valSetResp.ValidatorWithBlsKeys {
		validatorSet[i] = types.Validator{
			Addr:  common.HexToAddress(val.ValidatorAddress).Bytes(),
			Power: int64(val.VotingPower),
		}
	}

	// Create a copy of the checkpoint with meta
	ckptWithMeta := req.CheckpointWithMeta

	// Process each validator's signature
	for _, val := range valSetResp.ValidatorWithBlsKeys {
		valAddr := common.HexToAddress(val.ValidatorAddress)

		// Parse BLS public key
		blsPubkey, err := hex.DecodeString(val.BlsPubKeyHex)
		if err != nil {
			logger.Error("Failed to parse BLS public key for validator %s: %v", val.ValidatorAddress, err)
			continue
		}

		// Get the BLS signature for the validator address
		blsSigBytes, found := validResponses[val.ValidatorAddress]
		if !found {
			logger.Debug("BLS signature not found for validator %s", val.ValidatorAddress)
			continue
		}

		// Accumulate the signature
		err = ckptWithMeta.Accumulate(validatorSet, valAddr, blsPubkey, blsSigBytes, totalPower)
		if err != nil {
			logger.Error("Failed to accumulate BLS signature for validator %s: %v", val.ValidatorAddress, err)
			continue
		}

		logger.Debug("Accumulated BLS signature for validator %s", val.ValidatorAddress)
	}

	return &ckptWithMeta, nil
}