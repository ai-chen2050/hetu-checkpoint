package types

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	fmt "fmt"
)

const (
	// HashSize is the size in bytes of a hash
	HashSize = sha256.Size
)

type BlockHash []byte

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
