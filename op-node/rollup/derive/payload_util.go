package derive

import (
	"encoding/binary"
	"fmt"

	"github.com/zircuit-labs/l2-geth-public/core/types"

	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
)

// PayloadToBlockRef extracts the essential L2BlockRef information from an execution payload,
// falling back to genesis information if necessary.
func PayloadToBlockRef(rollupCfg *rollup.Config, payload *eth.ExecutionPayload, l1Info *types.L1Info) (eth.L2BlockRef, error) {
	genesis := &rollupCfg.Genesis
	var l1Origin eth.BlockID
	var sequenceNumber uint64
	if uint64(payload.BlockNumber) == genesis.L2.Number {
		if payload.BlockHash != genesis.L2.Hash {
			return eth.L2BlockRef{}, fmt.Errorf("expected L2 genesis hash to match L2 block at genesis block number %d: %s <> %s", genesis.L2.Number, payload.BlockHash, genesis.L2.Hash)
		}
		l1Origin = genesis.L1
		sequenceNumber = 0
	} else {
		if l1Info == nil {
			info, err := L1InfoFromExecutionPayload(rollupCfg, payload)
			if err != nil {
				return eth.L2BlockRef{}, fmt.Errorf("failed to derive L1 block info from execution payload: %w", err)
			}
			l1Origin = eth.BlockID{Hash: info.BlockHash, Number: info.Number}
			sequenceNumber = info.SequenceNumber
		} else {
			l1Origin = eth.BlockID{Hash: l1Info.BlockHash, Number: l1Info.Number}
			sequenceNumber = l1Info.SequenceNumber
		}
	}

	return eth.L2BlockRef{
		Hash:           payload.BlockHash,
		Number:         uint64(payload.BlockNumber),
		ParentHash:     payload.ParentHash,
		Time:           uint64(payload.Timestamp),
		L1Origin:       l1Origin,
		SequenceNumber: sequenceNumber,
	}, nil
}

func PayloadToSystemConfig(rollupCfg *rollup.Config, payload *eth.ExecutionPayload, l1Info *types.L1Info) (eth.SystemConfig, error) {
	if uint64(payload.BlockNumber) == rollupCfg.Genesis.L2.Number {
		if payload.BlockHash != rollupCfg.Genesis.L2.Hash {
			return eth.SystemConfig{}, fmt.Errorf(
				"expected L2 genesis hash to match L2 block at genesis block number %d: %s <> %s",
				rollupCfg.Genesis.L2.Number, payload.BlockHash, rollupCfg.Genesis.L2.Hash)
		}
		return rollupCfg.Genesis.SystemConfig, nil
	}

	var l1BlockInfo *L1BlockInfo

	if l1Info == nil {
		var err error
		l1BlockInfo, err = L1InfoFromExecutionPayload(rollupCfg, payload)
		if err != nil {
			return eth.SystemConfig{}, fmt.Errorf("failed to derive L1 block info from execution payload: %w", err)
		}
	} else {
		// Copy the L1Info to avoid modifying the original.
		l1Info := *l1Info
		l1BlockInfo = NewL1BlockInfo(&l1Info)
	}

	if isEcotoneButNotFirstBlock(rollupCfg, uint64(payload.Timestamp)) {
		// Translate Ecotone values back into encoded scalar if needed.
		// We do not know if it was derived from a v0 or v1 scalar,
		// but v1 is fine, a 0 blob base fee has the same effect.
		l1BlockInfo.L1FeeScalar[0] = 1
		binary.BigEndian.PutUint32(l1BlockInfo.L1FeeScalar[24:28], l1BlockInfo.BlobBaseFeeScalar)
		binary.BigEndian.PutUint32(l1BlockInfo.L1FeeScalar[28:32], l1BlockInfo.BaseFeeScalar)
	}
	return eth.SystemConfig{
		BatcherAddr: l1BlockInfo.BatcherAddr,
		Overhead:    l1BlockInfo.L1FeeOverhead,
		Scalar:      l1BlockInfo.L1FeeScalar,
		GasLimit:    uint64(payload.GasLimit),
	}, nil
}
