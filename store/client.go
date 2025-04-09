package store

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DBClient represents a PostgreSQL database client
type DBClient struct {
	db *sqlx.DB
}

// Config holds database configuration
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// NewDBClient creates a new database client
func NewDBClient(cfg Config) (*DBClient, error) {
	// First connect to 'postgres' database to create the target database if needed
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password)

	tempDB, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres database: %v", err)
	}
	defer tempDB.Close()

	// Check if database exists
	var exists bool
	err = tempDB.Get(&exists, "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", cfg.DBName)
	if err != nil {
		return nil, fmt.Errorf("error checking database existence: %v", err)
	}

	// Create database if it doesn't exist
	if !exists {
		_, err = tempDB.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DBName))
		if err != nil {
			return nil, fmt.Errorf("error creating database: %v", err)
		}
	}

	// Connect to the target database
	targetPsqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sqlx.Connect("postgres", targetPsqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close() // Close the connection if ping fails
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return &DBClient{db: db}, nil
}

// Close closes the database connection
func (c *DBClient) Close() error {
	return c.db.Close()
}

// CreateDispatcherTables creates the necessary database tables if they don't exist
func (c *DBClient) CreateDispatcherTables() error {
	// Create sign_requests table
	_, err := c.db.Exec(`
		CREATE TABLE IF NOT EXISTS sign_requests (
			id SERIAL PRIMARY KEY,
			message TEXT NOT NULL,
			status VARCHAR(20) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating sign_requests table: %v", err)
	}

	// Create sign_responses table
	_, err = c.db.Exec(`
		CREATE TABLE IF NOT EXISTS sign_responses (
			id SERIAL PRIMARY KEY,
			request_id INTEGER REFERENCES sign_requests(id),
			validator_id VARCHAR(100) NOT NULL,
			signature TEXT NOT NULL,
			status VARCHAR(20) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating sign_responses table: %v", err)
	}

	// Create aggregated_checkpoints table
	_, err = c.db.Exec(`
		CREATE TABLE IF NOT EXISTS aggregated_checkpoints (
			id SERIAL PRIMARY KEY,
			request_id INTEGER REFERENCES sign_requests(id),
			epoch_num NUMERIC NOT NULL,
			block_hash TEXT NOT NULL,
			bitmap TEXT NOT NULL,
			bls_multi_sig TEXT,
			bls_aggr_pk TEXT,
			power_sum NUMERIC NOT NULL,
			status VARCHAR(20) NOT NULL,
			validator_count INTEGER NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating aggregated_checkpoints table: %v", err)
	}

	// Create reward_distributions table
	_, err = c.db.Exec(`
		CREATE TABLE IF NOT EXISTS reward_distributions (
			id SERIAL PRIMARY KEY,
			epoch_num NUMERIC NOT NULL,
			transaction_hash TEXT NOT NULL,
			status VARCHAR(20) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating reward_distributions table: %v", err)
	}

	return nil
}

// CreateValidatorTables creates the necessary database tables if they don't exist
func (c *DBClient) CreateValidatorTables() error {
	var err error

	// Create sign_responses table if it doesn't exist
	_, err = c.db.Exec(`
		CREATE TABLE IF NOT EXISTS sign_responses (
			id SERIAL PRIMARY KEY,
			request_id INTEGER,
			validator_id VARCHAR(100) NOT NULL,
			signature TEXT NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating sign_responses table: %v", err)
	}

	return nil
}

// InsertDisSignRequest inserts a new signing request to dispatcher db
func (c *DBClient) InsertDisSignRequest(message string) (*SignRequest, error) {
	query := `
		INSERT INTO sign_requests (message, status, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, message, status, created_at
	`
	req := &SignRequest{}
	err := c.db.QueryRowx(query, message, "PENDING", time.Now()).StructScan(req)
	if err != nil {
		return nil, fmt.Errorf("error inserting sign request: %v", err)
	}
	return req, nil
}

// InsertDisSignResponse inserts a validator's response
func (c *DBClient) InsertDisSignResponse(requestID int64, validatorID, signature string) (*SignResponse, error) {
	query := `
		INSERT INTO sign_responses (request_id, validator_id, signature, status, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, request_id, validator_id, signature, status, created_at
	`
	resp := &SignResponse{}
	err := c.db.QueryRowx(query, requestID, validatorID, signature, "COMPLETED", time.Now()).StructScan(resp)
	if err != nil {
		return nil, fmt.Errorf("error inserting sign response: %v", err)
	}
	return resp, nil
}

// UpdateDisSignRequestStatus updates the status of a signing request
func (c *DBClient) UpdateDisSignRequestStatus(requestID int64, status string) error {
	query := `
		UPDATE sign_requests
		SET status = $1
		WHERE id = $2
	`
	result, err := c.db.Exec(query, status, requestID)
	if err != nil {
		return fmt.Errorf("error updating sign request status: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("no request found with ID %d", requestID)
	}

	return nil
}

// InsertValSignResponse inserts a validator's response to validator db
func (c *DBClient) InsertValSignResponse(requestID int64, validatorID, signature string) (*SignResponse, error) {
	query := `
		INSERT INTO sign_responses (request_id, validator_id, signature, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, request_id, validator_id, signature, created_at
	`
	resp := &SignResponse{}
	err := c.db.QueryRowx(query, requestID, validatorID, signature, time.Now()).StructScan(resp)
	if err != nil {
		return nil, fmt.Errorf("error inserting sign response: %v", err)
	}
	return resp, nil
}

// InsertAggregatedCheckpoint inserts a new aggregated checkpoint record
func (c *DBClient) InsertAggregatedCheckpoint(
	requestID int64,
	epochNum uint64,
	blockHash string,
	bitmap string,
	blsMultiSig string,
	blsAggrPk string,
	powerSum string,
	status string,
	validatorCount int,
) (*AggregatedCheckpoint, error) {
	query := `
		INSERT INTO aggregated_checkpoints (
			request_id, epoch_num, block_hash, bitmap, bls_multi_sig, 
			bls_aggr_pk, power_sum, status, validator_count, created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, request_id, epoch_num, block_hash, bitmap, bls_multi_sig, 
		          bls_aggr_pk, power_sum, status, validator_count, created_at
	`
	checkpoint := &AggregatedCheckpoint{}
	err := c.db.QueryRowx(
		query,
		requestID,
		fmt.Sprintf("%d", epochNum), // Convert uint64 to string for NUMERIC
		blockHash,
		bitmap,
		blsMultiSig,
		blsAggrPk,
		powerSum,
		status,
		validatorCount,
		time.Now(),
	).StructScan(checkpoint)
	if err != nil {
		return nil, fmt.Errorf("error inserting aggregated checkpoint: %v", err)
	}
	return checkpoint, nil
}

// GetAggregatedCheckpointByEpoch retrieves an aggregated checkpoint by epoch number
func (c *DBClient) GetAggregatedCheckpointByEpoch(epochNum uint64) (*AggregatedCheckpoint, error) {
	query := `
		SELECT * FROM aggregated_checkpoints
		WHERE epoch_num = $1
		ORDER BY created_at DESC
		LIMIT 1
	`
	checkpoint := &AggregatedCheckpoint{}
	err := c.db.Get(checkpoint, query, epochNum)
	if err != nil {
		return nil, fmt.Errorf("error getting aggregated checkpoint: %v", err)
	}
	return checkpoint, nil
}

// GetLatestAggregatedCheckpoints retrieves the latest aggregated checkpoints
func (c *DBClient) GetLatestAggregatedCheckpoints(limit int) ([]*AggregatedCheckpoint, error) {
	query := `
		SELECT * FROM aggregated_checkpoints
		ORDER BY epoch_num DESC
		LIMIT $1
	`
	checkpoints := []*AggregatedCheckpoint{}
	err := c.db.Select(&checkpoints, query, limit)
	if err != nil {
		return nil, fmt.Errorf("error getting latest aggregated checkpoints: %v", err)
	}
	return checkpoints, nil
}

// InsertRewardDistribution inserts a new reward distribution record
func (c *DBClient) InsertRewardDistribution(
	epochNum uint64,
	transactionHash string,
	status string,
) (*RewardDistribution, error) {
	query := `
		INSERT INTO reward_distributions (
			epoch_num, transaction_hash, status, created_at
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id, epoch_num, transaction_hash, status, created_at
	`
	distribution := &RewardDistribution{}
	err := c.db.QueryRowx(
		query,
		epochNum,
		transactionHash,
		status,
		time.Now(),
	).StructScan(distribution)
	if err != nil {
		return nil, fmt.Errorf("error inserting reward distribution: %v", err)
	}
	return distribution, nil
}

// GetLastRewardDistribution retrieves the last reward distribution record
func (c *DBClient) GetLastRewardDistribution() (*RewardDistribution, error) {
	query := `
		SELECT * FROM reward_distributions
		ORDER BY epoch_num DESC
		LIMIT 1
	`
	distribution := &RewardDistribution{}
	err := c.db.Get(distribution, query)
	if err != nil {
		return nil, fmt.Errorf("error getting last reward distribution: %v", err)
	}
	return distribution, nil
}

// IsEpochRewardDistributed checks if rewards for a specific epoch have been distributed
func (c *DBClient) IsEpochRewardDistributed(epochNum uint64) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM reward_distributions
			WHERE epoch_num = $1 AND status = 'SUCCESS'
		)
	`
	var exists bool
	err := c.db.Get(&exists, query, epochNum)
	if err != nil {
		return false, fmt.Errorf("error checking if epoch reward was distributed: %v", err)
	}
	return exists, nil
}
