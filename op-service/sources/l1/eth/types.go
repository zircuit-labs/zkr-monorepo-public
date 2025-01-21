package eth

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common/hexutil"

	l1common "github.com/ethereum/go-ethereum/common"
	l1types "github.com/ethereum/go-ethereum/core/types"
	l2common "github.com/zircuit-labs/l2-geth-public/common"
	l2types "github.com/zircuit-labs/l2-geth-public/core/types"
	l2eth "github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
)

type Bytes32 [32]byte

func (b *Bytes32) UnmarshalJSON(text []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(b), text, b[:])
}

func (b *Bytes32) UnmarshalText(text []byte) error {
	return hexutil.UnmarshalFixedText("Bytes32", text, b[:])
}

func (b Bytes32) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

func (b Bytes32) String() string {
	return hexutil.Encode(b[:])
}

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (b Bytes32) TerminalString() string {
	return fmt.Sprintf("%x..%x", b[:3], b[29:])
}

type Bytes256 [256]byte

func (b *Bytes256) UnmarshalJSON(text []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(b), text, b[:])
}

func (b *Bytes256) UnmarshalText(text []byte) error {
	return hexutil.UnmarshalFixedText("Bytes32", text, b[:])
}

func (b Bytes256) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

func (b Bytes256) String() string {
	return hexutil.Encode(b[:])
}

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (b Bytes256) TerminalString() string {
	return fmt.Sprintf("%x..%x", b[:3], b[253:])
}

// TransactionsToHashes computes the transaction-hash for every transaction in the input.
func TransactionsToHashes(elems []*l1types.Transaction) []l1common.Hash {
	out := make([]l1common.Hash, len(elems))
	for i, el := range elems {
		out[i] = el.Hash()
	}
	return out
}

func ConvertToL1BlockLabel(label l2eth.BlockLabel) BlockLabel {
	return BlockLabel(label)
}

func ConvertHashToL1(l2Hash l2common.Hash) l1common.Hash {
	return l1common.Hash(l2Hash)
}

func ConvertHashToL2(l1Hash l1common.Hash) l2common.Hash {
	return l2common.Hash(l1Hash)
}

func ConvertAddressToL1(l2Address l2common.Address) l1common.Address {
	return l1common.Address(l2Address)
}

func ConvertAddressToL2(l1Address l1common.Address) l2common.Address {
	return l2common.Address(l1Address)
}

func InfoToBlockRef(info BlockInfo) L1BlockRef {
	return L1BlockRef{
		Hash:       info.Hash(),
		Number:     info.NumberU64(),
		ParentHash: info.ParentHash(),
		Time:       info.Time(),
	}
}

func ConvertToL1BlockRef(blockRef L1BlockRef) l2eth.L1BlockRef {
	return l2eth.L1BlockRef{
		Hash:       ConvertHashToL2(blockRef.Hash),
		Number:     blockRef.Number,
		ParentHash: ConvertHashToL2(blockRef.ParentHash),
		Time:       blockRef.Time,
	}
}

// Wrapper to adapt L1 BlockInfo to L2 BlockInfo interface
type BlockInfoWrapper struct {
	L1BlockInfo BlockInfo
}

func (w *BlockInfoWrapper) Hash() l2common.Hash {
	return l2common.BytesToHash(w.L1BlockInfo.Hash().Bytes())
}

func (w *BlockInfoWrapper) ParentHash() l2common.Hash {
	return l2common.BytesToHash(w.L1BlockInfo.ParentHash().Bytes())
}

func (w *BlockInfoWrapper) Coinbase() l2common.Address {
	return l2common.BytesToAddress(w.L1BlockInfo.Coinbase().Bytes())
}

func (w *BlockInfoWrapper) Root() l2common.Hash {
	return l2common.BytesToHash(w.L1BlockInfo.Root().Bytes())
}

func (w *BlockInfoWrapper) NumberU64() uint64 {
	return w.L1BlockInfo.NumberU64()
}

func (w *BlockInfoWrapper) Time() uint64 {
	return w.L1BlockInfo.Time()
}

func (w *BlockInfoWrapper) MixDigest() l2common.Hash {
	return l2common.BytesToHash(w.L1BlockInfo.MixDigest().Bytes())
}

func (w *BlockInfoWrapper) BaseFee() *big.Int {
	return w.L1BlockInfo.BaseFee()
}

func (w *BlockInfoWrapper) BlobBaseFee() *big.Int {
	return w.L1BlockInfo.BlobBaseFee()
}

func (w *BlockInfoWrapper) ReceiptHash() l2common.Hash {
	return l2common.BytesToHash(w.L1BlockInfo.ReceiptHash().Bytes())
}

func (w *BlockInfoWrapper) GasUsed() uint64 {
	return w.L1BlockInfo.GasUsed()
}

func (w *BlockInfoWrapper) GasLimit() uint64 {
	return w.L1BlockInfo.GasLimit()
}

func (w *BlockInfoWrapper) ParentBeaconRoot() *l2common.Hash {
	if w.L1BlockInfo.ParentBeaconRoot() == nil {
		return nil
	}
	hash := l2common.BytesToHash(w.L1BlockInfo.ParentBeaconRoot().Bytes())
	return &hash
}

func (w *BlockInfoWrapper) HeaderRLP() ([]byte, error) {
	return w.L1BlockInfo.HeaderRLP()
}

// Used for testings & op-e2e
func ConvertL1LogToL2(l1Log l1types.Log) l2types.Log {
	l2Topics := make([]l2common.Hash, len(l1Log.Topics))
	for i, topic := range l1Log.Topics {
		l2Topics[i] = l2common.Hash(topic)
	}
	return l2types.Log{
		Address:     l2common.Address(l1Log.Address),
		Topics:      l2Topics,
		Data:        l1Log.Data,
		BlockNumber: l1Log.BlockNumber,
		TxHash:      l2common.Hash(l1Log.TxHash),
		TxIndex:     l1Log.TxIndex,
		BlockHash:   l2common.Hash(l1Log.BlockHash),
		Index:       l1Log.Index,
		Removed:     l1Log.Removed,
	}
}

func ConvertL2LogToL1(l2Log l2types.Log) l1types.Log {
	l2Topics := make([]l1common.Hash, len(l2Log.Topics))
	for i, topic := range l2Log.Topics {
		l2Topics[i] = l1common.Hash(topic)
	}
	return l1types.Log{
		Address:     l1common.Address(l2Log.Address),
		Topics:      l2Topics,
		Data:        l2Log.Data,
		BlockNumber: l2Log.BlockNumber,
		TxHash:      l1common.Hash(l2Log.TxHash),
		TxIndex:     l2Log.TxIndex,
		BlockHash:   l1common.Hash(l2Log.BlockHash),
		Index:       l2Log.Index,
		Removed:     l2Log.Removed,
	}
}
