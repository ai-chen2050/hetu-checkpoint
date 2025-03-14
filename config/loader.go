package config

import (
	"fmt"

	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/spf13/viper"
)

// LoadDispatcherConfig loads configuration for the dispatcher service
func LoadDispatcherConfig(configFile string) (*DispatcherConfig, error) {
	v := viper.New()

	if configFile != "" {
		v.SetConfigFile(configFile)
	} else {
		v.AddConfigPath(".")
		v.SetConfigName("config")
		v.SetConfigType("json")
	}

	// Read the config first
	if err := v.ReadInConfig(); err != nil {
		logger.Warn("Failed to read config file: %v", err)
		logger.Info("Using default configuration")
	} else {
		logger.Info("Using config file: %s", v.ConfigFileUsed())
	}

	// Set defaults AFTER reading config
	if !v.IsSet("http_port") {
		v.Set("http_port", 8080)
	}
	if !v.IsSet("tcp_port") {
		v.Set("tcp_port", 9090)
	}

	config := &DispatcherConfig{}
	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("unable to decode config: %v", err)
	}

	// Validate ports
	if config.HTTPPort == 0 {
		return nil, fmt.Errorf("HTTP port cannot be 0")
	}
	if config.TCPPort == 0 {
		return nil, fmt.Errorf("TCP port cannot be 0")
	}

	logger.Info("Loaded configuration - HTTP Port: %d, TCP Port: %d",
		config.HTTPPort, config.TCPPort)

	return config, nil
}

// LoadValidatorConfig loads configuration for the validator service
func LoadValidatorConfig(configFile string, cmdPort int) (*ValidatorConfig, error) {
	v := viper.New()

	if configFile != "" {
		v.SetConfigFile(configFile)
	} else {
		v.AddConfigPath(".")
		v.SetConfigName("config")
		v.SetConfigType("json")
	}

	// Set defaults
	v.SetDefault("port", 0) // 0 means random port

	// Read the config
	if err := v.ReadInConfig(); err != nil {
		logger.Warn("Failed to read config file: %v", err)
		logger.Info("Using default configuration")
	} else {
		logger.Info("Using config file: %s", v.ConfigFileUsed())
	}

	config := &ValidatorConfig{}
	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("unable to decode config: %v", err)
	}

	// Command line flags override config file
	if cmdPort != 0 {
		config.Port = cmdPort
	}

	// Validate required fields
	if config.DispatcherTcp == "" {
		return nil, fmt.Errorf("dispatcher address must be specified")
	}

	return config, nil
}
