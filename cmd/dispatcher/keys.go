package main

import (
	"github.com/hetu-project/hetu-checkpoint/crypto"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/spf13/cobra"
)

var (
	keyFilePath string
	keyPassword string
)

func init() {
	// Generate key command
	generateKeyCmd := &cobra.Command{
		Use:   "generate-key",
		Short: "Generate new Ethereum and BLS key pairs",
		Long:  `Generate new Ethereum and BLS key pairs and save them to a file.`,
		Run:   generateKey,
	}
	generateKeyCmd.Flags().StringVar(&keyFilePath, "output", "keys/dispatcher.json", "output file path")
	generateKeyCmd.Flags().StringVar(&keyPassword, "password", "", "password to encrypt the key (optional)")

	// Export private key command
	exportKeyCmd := &cobra.Command{
		Use:   "export-key",
		Short: "Export private keys from a key file",
		Long:  `Export the Ethereum and BLS private keys from a key file.`,
		Run:   exportKey,
	}
	exportKeyCmd.Flags().StringVar(&keyFilePath, "input", "keys/dispatcher.json", "input key file path")
	exportKeyCmd.Flags().StringVar(&keyPassword, "password", "", "password to decrypt the key (if encrypted)")

	// Add commands to root
	rootCmd.AddCommand(generateKeyCmd)
	rootCmd.AddCommand(exportKeyCmd)
}

func generateKey(cmd *cobra.Command, args []string) {
	// Set log level
	logger.SetLevel(logger.GetLevelFromString(logLevel))

	// Generate key pair
	keyPair, err := crypto.GenerateKeyPair()
	if err != nil {
		logger.Fatal("Failed to generate key pairs: %v", err)
	}

	// Save key pair
	err = crypto.SaveKeyPair(keyPair, keyFilePath, keyPassword)
	if err != nil {
		logger.Fatal("Failed to save key pairs: %v", err)
	}

	logger.Info("Generated and saved key pairs:")
	logger.Info("Ethereum Address: %s", keyPair.ETH.Address)
	logger.Info("Hetu Address: %s", keyPair.HetuAddress)
	logger.Info("Ethereum Public Key: %s", keyPair.ETH.PublicKey)
	logger.Info("BLS Public Key: %s", keyPair.BLS.PublicKey)

	if keyPassword == "" {
		logger.Info("Ethereum Private Key: %s", keyPair.ETH.PrivateKey)
		logger.Info("BLS Private Key: %s", keyPair.BLS.PrivateKey)
	} else {
		logger.Info("Private Keys: [encrypted with password]")
	}
	logger.Info("Keys saved to: %s", keyFilePath)
}

func exportKey(cmd *cobra.Command, args []string) {
	// Set log level
	logger.SetLevel(logger.GetLevelFromString(logLevel))

	// Export private keys
	ethPrivKey, blsPrivKey, hetuAddress, err := crypto.ExportPrivateKeys(keyFilePath, keyPassword)
	if err != nil {
		logger.Fatal("Failed to export private keys: %v", err)
	}

	logger.Info("Exported private keys:")
	logger.Info("Ethereum Private Key: %s", ethPrivKey)
	logger.Info("BLS Private Key: %s", blsPrivKey)
	logger.Info("Hetu Address: %s", hetuAddress)
}
