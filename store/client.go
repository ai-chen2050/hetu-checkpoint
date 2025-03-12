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

	// Test the connection
	if err := db.Ping(); err != nil {
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
