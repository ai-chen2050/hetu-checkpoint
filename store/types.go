package store

import "time"

// SignRequest represents a signing request record
type SignRequest struct {
	ID        int64     `db:"id"`
	Message   string    `db:"message"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

// SignResponse represents a validator's response to a signing request
type SignResponse struct {
	ID          int64     `db:"id"`
	RequestID   int64     `db:"request_id"`
	ValidatorID string    `db:"validator_id"`
	Signature   string    `db:"signature"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
}

// AggregatedCheckpoint represents an aggregated checkpoint record
type AggregatedCheckpoint struct {
	ID             int64     `db:"id"`
	RequestID      int64     `db:"request_id"`
	EpochNum       uint64    `db:"epoch_num"`
	BlockHash      string    `db:"block_hash"`
	Bitmap         string    `db:"bitmap"`
	BlsMultiSig    string    `db:"bls_multi_sig"`
	BlsAggrPk      string    `db:"bls_aggr_pk"`
	PowerSum       uint64    `db:"power_sum"`
	Status         string    `db:"status"`
	ValidatorCount int       `db:"validator_count"`
	CreatedAt      time.Time `db:"created_at"`
}

// RewardDistribution represents a reward distribution record
type RewardDistribution struct {
	ID              int64     `db:"id"`
	EpochNum        uint64    `db:"epoch_num"`
	TransactionHash string    `db:"transaction_hash"`
	Status          string    `db:"status"`
	CreatedAt       time.Time `db:"created_at"`
}
