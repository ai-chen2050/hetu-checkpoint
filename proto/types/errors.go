package types

import (
	"errors"
)

// x/checkpointing module errors
var (
	// NOTE: code 1 is reserved for internal errors
	ErrCkptAlreadyExist         = errors.New("raw checkpoint already exists")
	ErrCkptHashNotEqual         = errors.New("hash does not equal to raw checkpoint")
	ErrCkptDoesNotExist         = errors.New("raw checkpoint does not exist")
	ErrCkptNotAccumulating      = errors.New("raw checkpoint is no longer accumulating BLS sigs")
	ErrCkptAlreadyVoted         = errors.New("raw checkpoint already accumulated the validator")
	ErrInvalidRawCheckpoint     = errors.New("raw checkpoint is invalid")
	ErrInvalidCkptStatus        = errors.New("raw checkpoint's status is invalid")
	ErrInvalidBlsSignature      = errors.New("BLS signature is invalid")
	ErrBlsKeyAlreadyExist       = errors.New( "BLS public key already exists")
	ErrBlsKeyDoesNotExist       = errors.New( "BLS public key does not exist")
	ErrInsufficientVotingPower  = errors.New( "Accumulated voting power is not greater than 2/3 of total power")
	ErrValAddrDoesNotExist      = errors.New( "Validator address does not exist")
	ErrConflictingCheckpoint    = errors.New( "Conflicting checkpoint is found")
	ErrInvalidAppHash           = errors.New( "Provided app hash is Invalid")
	ErrReportBlsSigDoesNotExist = errors.New( "report bls signatures does not exist")
)
