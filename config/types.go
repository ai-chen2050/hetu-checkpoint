package config

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Bech32Prefix defines the Bech32 prefix used for EthAccounts
	Bech32Prefix = "hetu"

	// keyName is the name of the key in the keyring
	KeyName = "validator"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32Prefix

	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32Prefix + sdk.PrefixPublic
)

// DispatcherConfig holds configuration for the dispatcher service
type DispatcherConfig struct {
	DBHost       string `json:"db_host"`
	DBPort       int    `json:"db_port"`
	DBUser       string `json:"db_user"`
	DBPassword   string `json:"db_password"`
	DBName       string `json:"db_name"`
	HTTPPort     int    `json:"http_port"`
	TCPPort      int    `json:"tcp_port"`
	GRPCEndpoint string `json:"grpc_endpoint"`
}

// ValidatorConfig holds configuration for the validator service
type ValidatorConfig struct {
	DBHost                  string `json:"db_host"`
	DBPort                  int    `json:"db_port"`
	DBUser                  string `json:"db_user"`
	DBPassword              string `json:"db_password"`
	DBName                  string `json:"db_name"`
	DispatcherTcp           string `json:"dispatcher_tcp"`
	Port                    int    `json:"port"`
	ChainRpcURL             string `json:"chain_rpc_url"`
	StakingTokenAddress     string `json:"staking_token_address"`
	ValidatorStakingAddress string `json:"validator_staking_address"`
	DispatcherURL           string `json:"dispatcher_url"`
	ChainGRpcURL            string `json:"chain_grpc_url"`
	CometBFTSvr             string `json:"comet_bft_svr"`
	ChainID                 string `json:"chain_id"`
}
