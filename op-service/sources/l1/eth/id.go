package eth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type BlockID struct {
	Hash   common.Hash `json:"hash"`
	Number uint64      `json:"number"`
}

func (id BlockID) String() string {
	return fmt.Sprintf("%s:%d", id.Hash.String(), id.Number)
}

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (id BlockID) TerminalString() string {
	return fmt.Sprintf("%s:%d", id.Hash.TerminalString(), id.Number)
}

func ReceiptBlockID(r *types.Receipt) BlockID {
	return BlockID{Number: r.BlockNumber.Uint64(), Hash: r.BlockHash}
}

func HeaderBlockID(h *types.Header) BlockID {
	return BlockID{Number: h.Number.Uint64(), Hash: h.Hash()}
}

// hashID implements rpcBlockID for safe block-by-hash fetching
type HashID common.Hash

func (h HashID) Arg() any { return common.Hash(h) }
func (h HashID) CheckID(id BlockID) error {
	if common.Hash(h) != id.Hash {
		return fmt.Errorf("expected block hash %s but got block %s", common.Hash(h), id)
	}
	return nil
}

// numberID implements rpcBlockID for safe block-by-number fetching
type NumberID uint64

func (n NumberID) Arg() any { return hexutil.EncodeUint64(uint64(n)) }
func (n NumberID) CheckID(id BlockID) error {
	if uint64(n) != id.Number {
		return fmt.Errorf("expected block number %d but got block %s", uint64(n), id)
	}
	return nil
}

type L1BlockRef struct {
	Hash       common.Hash `json:"hash"`
	Number     uint64      `json:"number"`
	ParentHash common.Hash `json:"parentHash"`
	Time       uint64      `json:"timestamp"`
}

func (id L1BlockRef) String() string {
	return fmt.Sprintf("%s:%d", id.Hash.String(), id.Number)
}

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (id L1BlockRef) TerminalString() string {
	return fmt.Sprintf("%s:%d", id.Hash.TerminalString(), id.Number)
}

func (id L1BlockRef) ID() BlockID {
	return BlockID{
		Hash:   id.Hash,
		Number: id.Number,
	}
}

func (id L1BlockRef) ParentID() BlockID {
	n := id.ID().Number
	// Saturate at 0 with subtraction
	if n > 0 {
		n -= 1
	}
	return BlockID{
		Hash:   id.ParentHash,
		Number: n,
	}
}
