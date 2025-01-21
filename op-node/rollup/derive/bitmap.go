package derive

import (
	"bytes"

	"github.com/bits-and-blooms/bitset"
	"github.com/zircuit-labs/l2-geth-public/common/hexutil"
	"github.com/zircuit-labs/l2-geth-public/core/types"
)

// EmptyBitmap creates a bitmap of the given expected capacity, with all bits set to 0.
func EmptyBitmap(lengthHint int) *types.Bitmap {
	return types.NewBitmap(bitset.New(uint(lengthHint)))
}

// MustBitmap returns a Bitmap from the given portable bytes.
func MustBitmap(b hexutil.Bytes) *types.Bitmap {
	if len(b) == 0 {
		return nil
	}

	bs := bitset.New(uint(len(b)))
	buf := bytes.NewBuffer(b)
	_, err := bs.ReadFrom(buf)
	if err != nil {
		panic("failed to read from buffer")
	}

	return types.NewBitmap(bs)
}
