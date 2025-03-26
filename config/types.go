package config

import (
	"net"

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

// ValidatorClient represents a connected validator client
type ValidatorClient struct {
	Conn net.Conn
	Addr string
}

// ValidatorInfo represents information about a validator
type ValidatorInfo struct {
	Address    string `json:"address"`
	BlsPubKey  string `json:"bls_pub_key"`
	EthAddress string `json:"eth_address"`
}
