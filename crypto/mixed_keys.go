package crypto

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/crypto/bls12381"
	"github.com/hetu-project/hetu-checkpoint/crypto/eth"
)

// CombinedKeyPair represents both Ethereum and BLS key pairs
type CombinedKeyPair struct {
	ETH *eth.KeyPair `json:"eth"`
	BLS *BLSKeyPair  `json:"bls"`
	// Add Hetu address field
	HetuAddress string `json:"hetu_address,omitempty"`
}

// BLSKeyPair represents a BLS key pair
type BLSKeyPair struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// GenerateKeyPair generates both Ethereum and BLS key pairs
func GenerateKeyPair() (*CombinedKeyPair, error) {
	// Generate Ethereum key pair
	ethKeyPair, err := eth.GenerateKeyPair()
	if err != nil {
		return nil, fmt.Errorf("failed to generate Ethereum key pair: %v", err)
	}

	// Generate BLS key pair
	blsPrivKey, blsPubKey := bls12381.GenKeyPair()
	blsPubKey = blsPrivKey.PubKeyUncompress()
	blsKeyPair := &BLSKeyPair{
		PrivateKey: fmt.Sprintf("%x", blsPrivKey),
		PublicKey:  fmt.Sprintf("%x", blsPubKey),
	}

	// Calculate Hetu address
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount(config.Bech32PrefixAccAddr, config.Bech32PrefixAccPub)
	hetuAddress := sdk.AccAddress(ethKeyPair.PublicKey).String()

	return &CombinedKeyPair{
		ETH:         ethKeyPair,
		BLS:         blsKeyPair,
		HetuAddress: hetuAddress,
	}, nil
}

// SaveKeyPair saves a combined key pair to a file
func SaveKeyPair(keyPair *CombinedKeyPair, filePath string, password string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// If no password is provided, save as plain JSON
	if password == "" {
		data, err := json.MarshalIndent(keyPair, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal key pair: %v", err)
		}

		if err := os.WriteFile(filePath, data, 0600); err != nil {
			return fmt.Errorf("failed to write key pair to file: %v", err)
		}

		return nil
	}

	// For encrypted storage, we'll use Ethereum's keystore format
	// but include the BLS key in the encrypted data

	// First, save to a temporary JSON file
	tempFile := filePath + ".temp"
	data, err := json.MarshalIndent(keyPair, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal key pair: %v", err)
	}

	if err := os.WriteFile(tempFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write temporary key pair file: %v", err)
	}
	defer os.Remove(tempFile)

	// Now use eth.SaveKeyPair to create an encrypted keystore file
	if err := eth.SaveKeyPair(keyPair.ETH, filePath, password); err != nil {
		return fmt.Errorf("failed to save encrypted key pair: %v", err)
	}

	// Read the encrypted keystore file
	encryptedData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read encrypted keystore file: %v", err)
	}

	// Parse the encrypted keystore file
	var keystoreJSON map[string]interface{}
	if err := json.Unmarshal(encryptedData, &keystoreJSON); err != nil {
		return fmt.Errorf("failed to parse encrypted keystore file: %v", err)
	}

	// Add BLS key information to the keystore file
	keystoreJSON["bls"] = keyPair.BLS
	keystoreJSON["hetu_address"] = keyPair.HetuAddress
	// Write the updated keystore file
	updatedData, err := json.MarshalIndent(keystoreJSON, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal updated keystore file: %v", err)
	}

	if err := os.WriteFile(filePath, updatedData, 0600); err != nil {
		return fmt.Errorf("failed to write updated keystore file: %v", err)
	}

	return nil
}

// LoadKeyPair loads a combined key pair from a file
func LoadKeyPair(filePath string, password string) (*CombinedKeyPair, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("key file does not exist: %s", filePath)
	}

	// Read file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %v", err)
	}

	// Try to parse as plain JSON first
	var keyPair CombinedKeyPair
	if err := json.Unmarshal(data, &keyPair); err == nil && keyPair.ETH != nil && keyPair.BLS != nil {
		return &keyPair, nil
	}

	// If not plain JSON, try to parse as keystore file
	var keystoreJSON map[string]interface{}
	if err := json.Unmarshal(data, &keystoreJSON); err != nil {
		return nil, fmt.Errorf("failed to parse key file: %v", err)
	}

	// Check if it has BLS key information
	blsData, hasBLS := keystoreJSON["bls"]
	if !hasBLS {
		return nil, fmt.Errorf("key file does not contain BLS key information")
	}

	// Load Ethereum key pair
	ethKeyPair, err := eth.LoadKeyPair(filePath, password)
	if err != nil {
		return nil, fmt.Errorf("failed to load Ethereum key pair: %v", err)
	}

	// Check if it has Hetu address information
	hetuAddress, hasHetu := keystoreJSON["hetu_address"]
	if !hasHetu {
		// Calculate Hetu address
		sdkConfig := sdk.GetConfig()
		sdkConfig.SetBech32PrefixForAccount(config.Bech32PrefixAccAddr, config.Bech32PrefixAccPub)
		hetuAddress = sdk.AccAddress(ethKeyPair.PublicKey).String()
	}

	// Extract BLS key pair
	blsJSON, err := json.Marshal(blsData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal BLS key data: %v", err)
	}

	var blsKeyPair BLSKeyPair
	if err := json.Unmarshal(blsJSON, &blsKeyPair); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BLS key data: %v", err)
	}

	return &CombinedKeyPair{
		ETH: ethKeyPair,
		BLS: &blsKeyPair,
		HetuAddress: hetuAddress.(string),
	}, nil
}

// ExportPrivateKeys exports both Ethereum and BLS private keys
func ExportPrivateKeys(filePath string, password string) (string, string, string, error) {
	keyPair, err := LoadKeyPair(filePath, password)
	if err != nil {
		return "", "", "", err
	}
	return keyPair.ETH.PrivateKey, keyPair.BLS.PrivateKey, keyPair.HetuAddress, nil
}
