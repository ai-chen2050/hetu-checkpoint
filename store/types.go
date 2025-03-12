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
