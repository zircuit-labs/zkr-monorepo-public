package derive

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zircuit-labs/l2-geth-public/core/types"
)

func TestBitmap(t *testing.T) {
	// create a bitmap that with 10 bits set to 0
	bm := EmptyBitmap(10)
	assert.Equal(t, 10, bm.Len())

	setIndices := bm.Indices()
	assert.Empty(t, setIndices)

	// set some indices in the bitmap
	bm.Set(2)
	bm.Set(5)
	bm.Set(6)
	bm.Set(7)
	assert.Equal(t, 10, bm.Len())  // len is still 10
	bm.Set(13)                     // will cause the bitmap to grow transparently
	assert.Equal(t, 14, bm.Len())  // len is now 14
	assert.Equal(t, 5, bm.Count()) // 5 of the bits are set

	bm.Clear(15)                  // will cause the bitmap to grow transparently
	assert.Equal(t, 16, bm.Len()) // len is now 16
	bm.Clear(5)
	assert.Equal(t, 4, bm.Count()) // 4 of the bits are set

	// validate String method
	assert.Equal(t, "0b0010001100000100", bm.String())

	// validate bits are correctly set
	setIndices = bm.Indices()
	assert.Equal(t, []int{2, 6, 7, 13}, setIndices)

	for i := range bm.Len() {
		assert.Equal(t, slices.Contains(setIndices, i), bm.Test(i))
	}

	// validate encode/decode to hexutil.Bytes
	mapBytes := bm.MustBytes()
	decodedBitmap := MustBitmap(mapBytes)
	assert.True(t, bm.Equal(*decodedBitmap))
	assert.True(t, decodedBitmap.Equal(*bm))
	assert.Equal(t, setIndices, decodedBitmap.Indices())

	// check decoding nil is nil
	emptyBytesBitmap := MustBitmap(nil)
	assert.Nil(t, emptyBytesBitmap)

	// check encoding empty bitmap is nil
	// NOTE: all zeroes != empty
	emptyBitmap := EmptyBitmap(0)
	assert.Equal(t, 0, emptyBitmap.Len())
	assert.Nil(t, emptyBitmap.MustBytes())
	assert.Equal(t, "<empty>", emptyBitmap.String())

	// demonstrate how all zeroes != empty
	zeroedBitmap := EmptyBitmap(0)
	zeroedBitmap.Clear(7)
	assert.Equal(t, 8, zeroedBitmap.Len())
	assert.Equal(t, "0b00000000", zeroedBitmap.String())
}

func RandomBitmap(length int) *types.Bitmap {
	bm := EmptyBitmap(length)
	for i := range length {
		if rand.Intn(2) != 0 {
			bm.Set(i)
		}
	}
	return bm
}

func RandomBitmapNonEmpty(length int) *types.Bitmap {
	bm := EmptyBitmap(length)
	// make sure at least one bit is set
	if bm.Count() == 0 {
		bm.Set(rand.Intn(length))
	}
	return bm
}
