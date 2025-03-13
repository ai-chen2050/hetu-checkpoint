package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

// KeyPair represents an Ethereum key pair
type KeyPair struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// GenerateKeyPair generates a new Ethereum key pair
func GenerateKeyPair() (*KeyPair, error) {
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	// Get the public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}

	// Get the address
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Convert to hex strings
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	return &KeyPair{
		Address:    address.Hex(),
		PrivateKey: hex.EncodeToString(privateKeyBytes),
		PublicKey:  hex.EncodeToString(publicKeyBytes),
	}, nil
}

// SaveKeyPair saves a key pair to a file
func SaveKeyPair(keyPair *KeyPair, filePath string, password string) error {
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

	// Convert hex private key to ECDSA
	privateKeyBytes, err := hex.DecodeString(keyPair.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to decode private key: %v", err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return fmt.Errorf("failed to convert private key to ECDSA: %v", err)
	}

	// Create a keystore and encrypt the key
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)

	// Create account from private key
	_, err = ks.ImportECDSA(privateKey, password)
	if err != nil {
		return fmt.Errorf("failed to import private key: %v", err)
	}

	// Get the keystore file path
	keystoreFiles, err := filepath.Glob(filepath.Join(dir, "UTC--*"))
	if err != nil || len(keystoreFiles) == 0 {
		return fmt.Errorf("failed to find keystore file: %v", err)
	}

	// Find the most recently created file (should be our new keystore)
	var latestFile string
	var latestTime int64
	for _, file := range keystoreFiles {
		info, err := os.Stat(file)
		if err != nil {
			continue
		}
		if info.ModTime().Unix() > latestTime {
			latestTime = info.ModTime().Unix()
			latestFile = file
		}
	}

	// Rename the keystore file to the specified path
	if latestFile != filePath {
		if err := os.Rename(latestFile, filePath); err != nil {
			return fmt.Errorf("failed to rename keystore file: %v", err)
		}
	}

	return nil
}

// LoadKeyPair loads a key pair from a file
func LoadKeyPair(filePath string, password string) (*KeyPair, error) {
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
	var keyPair KeyPair
	if err := json.Unmarshal(data, &keyPair); err == nil && keyPair.PrivateKey != "" {
		return &keyPair, nil
	}

	// If not plain JSON or missing private key, try to decrypt as keystore file
	if password == "" {
		return nil, fmt.Errorf("password required to decrypt keystore file")
	}

	// Parse keystore file
	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt keystore file: %v", err)
	}

	// Convert to KeyPair
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	publicKeyBytes := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)

	return &KeyPair{
		Address:    key.Address.Hex(),
		PrivateKey: hex.EncodeToString(privateKeyBytes),
		PublicKey:  hex.EncodeToString(publicKeyBytes),
	}, nil
}

// ExportPrivateKey exports the private key from a keystore file
func ExportPrivateKey(filePath string, password string) (string, error) {
	keyPair, err := LoadKeyPair(filePath, password)
	if err != nil {
		return "", err
	}
	return keyPair.PrivateKey, nil
}
