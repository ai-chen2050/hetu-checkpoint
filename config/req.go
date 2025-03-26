package config

import (
	"github.com/hetu-project/hetu-checkpoint/proto/types"
)

// Request represents a client request to the dispatcher
type Request struct {
	ValidatorAddress string              `json:"validator_address"`
	Checkpoint       types.RawCheckpoint `json:"checkpoint"`
}

// Response represents the dispatcher's response to a client request
type Response struct {
	TotalValidators    int               `json:"total_validators"`
	ResponsesReceived  int               `json:"responses_received"`
	Errors             int               `json:"errors"`
	NoResponse         int               `json:"no_response"`
	ValidatorResponses map[string]string `json:"validator_responses"`
}

// ValidatorResponse represents a validator's response to a signing request
type ValidatorResponse struct {
	Response []byte
	Error    error
}
