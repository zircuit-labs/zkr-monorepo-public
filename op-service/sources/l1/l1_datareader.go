package l1

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	l2common "github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/log"
	l2eth "github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
	l1eth "github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/l1/eth"
)

// L1Reader defines an interface for reading Layer 1 (L1) blockchain data.
// It provides methods to retrieve chain information, block references,
// subscribe to new block headers, fetch receipts, transaction details,
// read storage at specific addresses, and close the reader.
type L1Reader interface {
	ChainID(ctx context.Context) (*big.Int, error)
	L1BlockRefByLabel(ctx context.Context, label l2eth.BlockLabel) (l2eth.L1BlockRef, error)
	L1BlockRefByNumber(ctx context.Context, num uint64) (l2eth.L1BlockRef, error)
	L1BlockRefByHash(context.Context, l2common.Hash) (l2eth.L1BlockRef, error)
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	FetchReceipts(ctx context.Context, blockHash l2common.Hash) (l2eth.BlockInfo, types.Receipts, error)
	InfoAndTxsByHash(ctx context.Context, hash l2common.Hash) (l2eth.BlockInfo, types.Transactions, error)
	InfoByHash(ctx context.Context, hash l2common.Hash) (l2eth.BlockInfo, error)
	InfoByNumber(ctx context.Context, num uint64) (l2eth.BlockInfo, error)
	ReadStorageAt(ctx context.Context, address l2common.Address, storageSlot l2common.Hash, blockHash l2common.Hash) (common.Hash, error)
	Close()
}

type L1DataReader struct {
	l1Client *L1Client
	log      log.Logger
}

// NewL1DataReader creates and returns a new instance of L1DataReader.
// It requires an L1Client for L1 interactions and a logger for logging purposes.
func NewL1DataReader(l1Client *L1Client, log log.Logger) *L1DataReader {
	return &L1DataReader{
		l1Client: l1Client,
		log:      log,
	}
}

func (r *L1DataReader) ChainID(ctx context.Context) (*big.Int, error) {
	return r.l1Client.ChainID(ctx)
}

// L1BlockRefByLabel fetches a L1 block reference based on a block label.
// It converts the L2 block label to an L1-compatible label before querying and return the result
// as a l2eth.L1BlockRef that op-node can understand.
func (r *L1DataReader) L1BlockRefByLabel(ctx context.Context, label l2eth.BlockLabel) (l2eth.L1BlockRef, error) {
	l1Ref, err := r.l1Client.L1BlockRefByLabel(ctx, l1eth.ConvertToL1BlockLabel(label))
	if err != nil {
		return l2eth.L1BlockRef{}, err
	}

	return l1eth.ConvertToL1BlockRef(l1Ref), nil
}

// L1BlockRefByNumber retrieves a block reference by its block number.
// It delegates the call to the underlying L1Client and converts the result
// as a l2eth.L1BlockRef that op-node can understand.
func (r *L1DataReader) L1BlockRefByNumber(ctx context.Context, num uint64) (l2eth.L1BlockRef, error) {
	l1Ref, err := r.l1Client.L1BlockRefByNumber(ctx, num)
	return l1eth.ConvertToL1BlockRef(l1Ref), err
}

// L1BlockRefByHash fetches a block reference using its hash.
// It converts the hash to an L1-compatible hash before querying and return the result
// as a l2eth.L1BlockRef that op-node can understand.
func (r *L1DataReader) L1BlockRefByHash(ctx context.Context, hash l2common.Hash) (l2eth.L1BlockRef, error) {
	l1Ref, err := r.l1Client.L1BlockRefByHash(ctx, l1eth.ConvertHashToL1(hash))
	return l1eth.ConvertToL1BlockRef(l1Ref), err
}

// SubscribeNewHead allows subscription to new L1 block headers.
// It delegates the subscription to the underlying L1Client.
func (r *L1DataReader) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return r.l1Client.SubscribeNewHead(ctx, ch)
}

// FetchReceipts retrieves the receipts for all transactions in a specific L1 block.
// It converts the block hash to an L1-compatible hash before querying.
// The returned receipts are of the L1 type, ensuring compatibility with the L1 data structure.
// The L1BlockInfo is wrapped in a BlockInfoWrapper to make it compatible with op-node's l2eth.BlockInfo.
func (r *L1DataReader) FetchReceipts(ctx context.Context, blockHash l2common.Hash) (l2eth.BlockInfo, types.Receipts, error) {
	l1BlockInfo, l1Receipt, err := r.l1Client.FetchReceipts(ctx, l1eth.ConvertHashToL1(blockHash))
	blockInfo := &l1eth.BlockInfoWrapper{L1BlockInfo: l1BlockInfo}
	return blockInfo, l1Receipt, err
}

// InfoAndTxsByHash retrieves the L1 block info and transactions for a specific block hash.
// It converts the block hash to an L1-compatible hash before querying.
// The returned receipts are of the L1 type, ensuring compatibility with the L1 data structure.
// The L1BlockInfo is wrapped in a BlockInfoWrapper to make it compatible with op-node's l2eth.BlockInfo.
func (r *L1DataReader) InfoAndTxsByHash(ctx context.Context, hash l2common.Hash) (l2eth.BlockInfo, types.Transactions, error) {
	l1BlockInfo, l1Txs, err := r.l1Client.InfoAndTxsByHash(ctx, l1eth.ConvertHashToL1(hash))
	blockInfo := &l1eth.BlockInfoWrapper{L1BlockInfo: l1BlockInfo}
	return blockInfo, l1Txs, err
}

// InfoByHash retrieves the L1 block info for a specific block hash.
// It converts the block hash to an L1-compatible hash before querying.
// Wraps the L1BlockInfo into a BlockInfoWrapper to make it compatible with op-node's l2eth.BlockInfo.
func (r *L1DataReader) InfoByHash(ctx context.Context, hash l2common.Hash) (l2eth.BlockInfo, error) {
	l1BlockInfo, err := r.l1Client.InfoByHash(ctx, l1eth.ConvertHashToL1(hash))
	blockInfo := &l1eth.BlockInfoWrapper{L1BlockInfo: l1BlockInfo}
	return blockInfo, err
}

// InfoByNumber retrieves the block info for a specific block number.
// Wraps the L1BlockInfo into a BlockInfoWrapper to make it compatible with op-node's l2eth.BlockInfo.
func (r *L1DataReader) InfoByNumber(ctx context.Context, num uint64) (l2eth.BlockInfo, error) {
	l1BlockInfo, err := r.l1Client.InfoByNumber(ctx, num)
	blockInfo := &l1eth.BlockInfoWrapper{L1BlockInfo: l1BlockInfo}
	return blockInfo, err
}

// ReadStorageAt reads the storage value at a specific address, storage slot, and block hash.
// It converts the address, storage slot, and block hash to L1-compatible values before querying.
func (r *L1DataReader) ReadStorageAt(ctx context.Context, address l2common.Address, storageSlot l2common.Hash, blockHash l2common.Hash) (common.Hash, error) {
	hash, err := r.l1Client.ReadStorageAt(ctx, l1eth.ConvertAddressToL1(address), l1eth.ConvertHashToL1(storageSlot), l1eth.ConvertHashToL1(blockHash))
	return hash, err
}

func (r *L1DataReader) Close() {
	r.l1Client.Close()
}
