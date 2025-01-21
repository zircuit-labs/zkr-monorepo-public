package derive

import (
	"fmt"

	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/core/types"

	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
)

// L2BlockRefSource is a source for the generation of a L2BlockRef. E.g. a
// *types.Block is a L2BlockRefSource.
//
// L2BlockToBlockRef extracts L2BlockRef from a L2BlockRefSource. The first
// transaction of a source must be a Deposit transaction.
type L2BlockRefSource interface {
	Hash() common.Hash
	ParentHash() common.Hash
	NumberU64() uint64
	Time() uint64
	Transactions() types.Transactions
}

// L2BlockToBlockRef extracts the essential L2BlockRef information from an L2
// block ref source, falling back to genesis information if necessary.
func L2BlockToBlockRef(rollupCfg *rollup.Config, block L2BlockRefSource, l1Info *types.L1Info) (eth.L2BlockRef, error) {
	hash, number := block.Hash(), block.NumberU64()

	var l1Origin eth.BlockID
	var sequenceNumber uint64
	genesis := &rollupCfg.Genesis
	if number == genesis.L2.Number {
		if hash != genesis.L2.Hash {
			return eth.L2BlockRef{}, fmt.Errorf("expected L2 genesis hash to match L2 block at genesis block number %d: %s <> %s", genesis.L2.Number, hash, genesis.L2.Hash)
		}
		l1Origin = genesis.L1
		sequenceNumber = 0
	} else {
		l1Origin = eth.BlockID{Hash: l1Info.BlockHash, Number: l1Info.Number}
		sequenceNumber = l1Info.SequenceNumber
	}

	return eth.L2BlockRef{
		Hash:           hash,
		Number:         number,
		ParentHash:     block.ParentHash(),
		Time:           block.Time(),
		L1Origin:       l1Origin,
		SequenceNumber: sequenceNumber,
	}, nil
}
