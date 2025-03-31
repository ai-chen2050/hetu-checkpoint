package types

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"github.com/boljen/go-bitmap"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/hetu-project/hetu-checkpoint/crypto/bls12381"
)

const (
	// HashSize is the size in bytes of a hash
	HashSize   = sha256.Size
	BitmapBits = 64 * 8 // 512 bits for 512 validators at top for a epoch
	DefaultValidatorSize = 512
)

// BlsSigner is an interface for signing BLS messages
type BlsSigner interface {
	SignMsgWithBls(msg []byte) (bls12381.Signature, error)
	BlsPubKey() (bls12381.PublicKey, error)
}

type BlockHash []byte

type BlsSigHash []byte

type RawCkptHash []byte

func NewCheckpoint(epochNum uint64, blockHash BlockHash) *RawCheckpoint {
	return &RawCheckpoint{
		EpochNum:    epochNum,
		BlockHash:   &blockHash,
		Bitmap:      bitmap.New(BitmapBits), // 64 bytes, holding 512 validators
		BlsMultiSig: nil,
	}
}

func NewCheckpointWithMeta(ckpt *RawCheckpoint, status CheckpointStatus) *RawCheckpointWithMeta {
	return &RawCheckpointWithMeta{
		Ckpt:      ckpt,
		Status:    status,
		Lifecycle: []*CheckpointStateUpdate{},
	}
}

// Accumulate does the following things
// 1. aggregates the BLS signature
// 2. aggregates the BLS public key
// 3. updates Bitmap
// 4. accumulates voting power
// it returns nil if the checkpoint is updated, otherwise it returns an error
func (cm *RawCheckpointWithMeta) Accumulate(
	vals ValidatorSet,
	signerAddr common.Address,
	signerBlsKey bls12381.PublicKey,
	sig bls12381.Signature,
	totalPower uint64) error {
	// the checkpoint should be accumulating
	if cm.Status != Accumulating {
		// return nil if the checkpoint is no longer accumulating(maybe sealed)
		return nil
	}

	val, index, err := vals.FindValidatorWithIndex(signerAddr)
	if err != nil {
		return err
	}

	// return an error if the validator has already voted
	if bitmap.Get(cm.Ckpt.Bitmap, index) {
		return ErrCkptAlreadyVoted
	}

	// aggregate BLS sig
	if cm.Ckpt.BlsMultiSig != nil {
		aggSig, err := bls12381.AggrSig(*cm.Ckpt.BlsMultiSig, sig)
		if err != nil {
			return err
		}
		cm.Ckpt.BlsMultiSig = &aggSig
	} else {
		cm.Ckpt.BlsMultiSig = &sig
	}

	// aggregate BLS public key
	if cm.BlsAggrPk != nil {
		aggPK, err := bls12381.AggrPK(*cm.BlsAggrPk, signerBlsKey)
		if err != nil {
			return err
		}
		cm.BlsAggrPk = &aggPK
	} else {
		cm.BlsAggrPk = &signerBlsKey
	}

	// update bitmap
	bitmap.Set(cm.Ckpt.Bitmap, index, true)

	// accumulate voting power and update status when the threshold is reached
	cm.PowerSum += uint64(val.Power)
	
	powerSumBig := new(big.Int).SetUint64(cm.PowerSum)
	totalPowerBig := new(big.Int).SetUint64(totalPower)
	// Multiply PowerSum by 3 and TotalPower by 2
	powerSumBig.Mul(powerSumBig, big.NewInt(3))
	totalPowerBig.Mul(totalPowerBig, big.NewInt(2))

	// Compare the results
	if powerSumBig.Cmp(totalPowerBig) > 0 {
		cm.Status = Sealed
	}

	return nil
}

// RecordStateUpdate appends a new state update to the raw ckpt with meta
// where the time/height are captured by the current ctx
func (cm *RawCheckpointWithMeta) RecordStateUpdate(ctx context.Context, status CheckpointStatus) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	height, time := sdkCtx.HeaderInfo().Height, sdkCtx.HeaderInfo().Time
	stateUpdate := &CheckpointStateUpdate{
		State:       status,
		BlockHeight: uint64(height),
		BlockTime:   &time,
	}
	cm.Lifecycle = append(cm.Lifecycle, stateUpdate)
}

func (bh *BlockHash) Unmarshal(bz []byte) error {
	if len(bz) != HashSize {
		return fmt.Errorf(
			"invalid block hash length, expected: %d, got: %d",
			HashSize, len(bz))
	}
	*bh = bz
	return nil
}

func (bh *BlockHash) Size() (n int) {
	if bh == nil {
		return 0
	}
	return len(*bh)
}

func (bh *BlockHash) Equal(l BlockHash) bool {
	return bh.String() == l.String()
}

func (bh *BlockHash) String() string {
	return hex.EncodeToString(*bh)
}

func (bh *BlockHash) MustMarshal() []byte {
	bz, err := bh.Marshal()
	if err != nil {
		panic(err)
	}
	return bz
}

func (bh *BlockHash) Marshal() ([]byte, error) {
	return *bh, nil
}

func (bh BlockHash) MarshalTo(data []byte) (int, error) {
	copy(data, bh)
	return len(data), nil
}

func (bh *BlockHash) ValidateBasic() error {
	if bh == nil {
		return errors.New("invalid block hash")
	}
	if len(*bh) != HashSize {
		return errors.New("invalid block hash")
	}
	return nil
}

// ValidateBasic does sanity checks on a raw checkpoint
func (ckpt RawCheckpoint) ValidateBasic() error {
	err := ckpt.BlockHash.ValidateBasic()
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidRawCheckpoint, "error validating block hash: %s", err.Error())
	}
	err = ckpt.BlsMultiSig.ValidateBasic()
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidRawCheckpoint, "error validating BLS multi-signature: %s", err.Error())
	}

	return nil
}

func CkptWithMetaToBytes(cdc codec.BinaryCodec, ckptWithMeta *RawCheckpointWithMeta) []byte {
	return cdc.MustMarshal(ckptWithMeta)
}

func BytesToCkptWithMeta(cdc codec.BinaryCodec, bz []byte) (*RawCheckpointWithMeta, error) {
	ckptWithMeta := new(RawCheckpointWithMeta)
	err := cdc.Unmarshal(bz, ckptWithMeta)
	return ckptWithMeta, err
}