package eth

import (
	"testing"

	l1common "github.com/ethereum/go-ethereum/common"
	l2common "github.com/zircuit-labs/l2-geth-public/common"
	l2eth "github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"

	"github.com/stretchr/testify/assert"
)

func TestConvertToL1BlockLabel(t *testing.T) {
	label := l2eth.BlockLabel("label")
	convertedLabel := ConvertToL1BlockLabel(label)
	assert.Equal(t, BlockLabel(label), convertedLabel)
}

func TestConvertHashToL1(t *testing.T) {
	l2Hash := l2common.BytesToHash([]byte("l2hash"))
	l1Hash := ConvertHashToL1(l2Hash)
	assert.Equal(t, l1common.Hash(l2Hash), l1Hash)
}

func TestConvertHashToL2(t *testing.T) {
	l1Hash := l1common.BytesToHash([]byte("l1hash"))
	l2Hash := ConvertHashToL2(l1Hash)
	assert.Equal(t, l2common.Hash(l1Hash), l2Hash)
}

func TestConvertAddressToL1(t *testing.T) {
	l2Address := l2common.BytesToAddress([]byte("l2address"))
	l1Address := ConvertAddressToL1(l2Address)
	assert.Equal(t, l1common.Address(l2Address), l1Address)
}

func TestConvertAddressToL2(t *testing.T) {
	l1Address := l1common.BytesToAddress([]byte("l1address"))
	l2Address := ConvertAddressToL2(l1Address)
	assert.Equal(t, l2common.Address(l1Address), l2Address)
}

func TestConvertToL1BlockRef(t *testing.T) {
	blockRef := L1BlockRef{
		Hash:       l1common.BytesToHash([]byte("hash")),
		Number:     123,
		ParentHash: l1common.BytesToHash([]byte("parent")),
		Time:       456,
	}
	l2BlockRef := ConvertToL1BlockRef(blockRef)
	assert.Equal(t, ConvertHashToL2(blockRef.Hash), l2BlockRef.Hash)
	assert.Equal(t, blockRef.Number, l2BlockRef.Number)
	assert.Equal(t, ConvertHashToL2(blockRef.ParentHash), l2BlockRef.ParentHash)
	assert.Equal(t, blockRef.Time, l2BlockRef.Time)
}
