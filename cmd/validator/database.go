package main

import (
	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/store"
)

// initializeDatabase initializes the database connection
func initializeDatabase(cfg *config.ValidatorConfig) {
	logger.Info("Database persistence enabled, initializing connection...")
	var err error
	dbClient, err = store.NewDBClient(store.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		logger.Fatal("Failed to initialize database client: %v", err)
	}

	// Create database tables
	if err := dbClient.CreateValidatorTables(); err != nil {
		logger.Fatal("Failed to create database tables: %v", err)
	}
	logger.Info("Database initialized successfully")
}
