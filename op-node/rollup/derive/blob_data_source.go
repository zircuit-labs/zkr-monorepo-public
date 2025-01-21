package derive

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/core/types"
	"github.com/zircuit-labs/l2-geth-public/log"

	l1ethereum "github.com/ethereum/go-ethereum"
	l1types "github.com/ethereum/go-ethereum/core/types"
	"github.com/zircuit-labs/zkr-go-common/retry"
	"github.com/zircuit-labs/zkr-go-common/retry/strategy"
	"github.com/zircuit-labs/zkr-go-common/xerrors/errclass"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
)

type blobOrCalldata struct {
	// union type. exactly one of calldata or blob should be non-nil
	blob     *eth.Blob
	calldata *eth.Data
}

// BlobDataSource fetches blobs or calldata as appropriate and transforms them into usable rollup
// data.
type BlobDataSource struct {
	data         []blobOrCalldata
	ref          eth.L1BlockRef
	batcherAddr  common.Address
	dsCfg        DataSourceConfig
	fetcher      L1TransactionFetcher
	blobsFetcher L1BlobsFetcher
	log          log.Logger
}

// NewBlobDataSource creates a new blob data source.
func NewBlobDataSource(ctx context.Context, log log.Logger, dsCfg DataSourceConfig, fetcher L1TransactionFetcher, blobsFetcher L1BlobsFetcher, ref eth.L1BlockRef, batcherAddr common.Address) DataIter {
	return &BlobDataSource{
		ref:          ref,
		dsCfg:        dsCfg,
		fetcher:      fetcher,
		log:          log.New("origin", ref),
		batcherAddr:  batcherAddr,
		blobsFetcher: blobsFetcher,
	}
}

// Next returns the next piece of batcher data, or an io.EOF error if no data remains. It returns
// ResetError if it cannot find the referenced block or a referenced blob, or TemporaryError for
// any other failure to fetch a block or blob.
func (ds *BlobDataSource) Next(ctx context.Context) (eth.Data, error) {
	if ds.data == nil {
		var err error
		if ds.data, err = ds.open(ctx); err != nil {
			return nil, err
		}
	}

	if len(ds.data) == 0 {
		return nil, io.EOF
	}

	next := ds.data[0]
	ds.data = ds.data[1:]
	if next.calldata != nil {
		return *next.calldata, nil
	}

	data, err := next.blob.ToData()
	if err != nil {
		ds.log.Error("ignoring blob due to parse failure", "err", err)
		return ds.Next(ctx)
	}
	return data, nil
}

// open fetches and returns the blob or calldata (as appropriate) from all valid batcher
// transactions in the referenced block. Returns an empty (non-nil) array if no batcher
// transactions are found. It returns ResetError if it cannot find the referenced block or a
// referenced blob, or TemporaryError for any other failure to fetch a block or blob.
func (ds *BlobDataSource) open(ctx context.Context) ([]blobOrCalldata, error) {
	_, txs, err := ds.fetcher.InfoAndTxsByHash(ctx, ds.ref.Hash)
	if err != nil {
		if errors.Is(err, l1ethereum.NotFound) {
			return nil, NewResetError(fmt.Errorf("failed to open blob data source: %w", err))
		}
		return nil, NewTemporaryError(fmt.Errorf("failed to open blob data source: %w", err))
	}

	data, hashes := dataAndHashesFromTxs(txs, &ds.dsCfg, ds.batcherAddr)

	if len(hashes) == 0 {
		// there are no blobs to fetch so we can return immediately
		return data, nil
	}

	// Create a retrier with exponential backoff base 3, initial delay 1 second, max delay 9 seconds (1, 3, 9)
	retrier := retry.NewRetrier(
		retry.WithStrategy(
			strategy.NewExponential(
				time.Second,              // initial delay
				time.Second*9,            // max delay
				strategy.WithBase(3),     // base for exponential backoff
				strategy.WithoutJitter(), // Disable jitter for consistent delays
			),
		),
		retry.WithMaxAttempts(3),                       // Retry up to 3 times
		retry.WithUnknownErrorsAs(errclass.Persistent), // Treat unknown errors as persistent (no retry)
	)

	var blobs []*eth.Blob
	// Use the retrier to execute the function with retries
	err = retrier.Try(ctx, func() error {
		blobs, err = ds.blobsFetcher.GetBlobs(ctx, ds.ref, hashes)
		if err == nil {
			return nil // Successfully fetched blobs, stop retrying
		} else if errors.Is(err, l1ethereum.NotFound) {
			// The beacon node might be lagging behind the L1 node.
			// In such cases, retrying allows the beacon node time to catch up,
			// which may help it successfully fetch the blobs on subsequent attempts.
			// This helps prevent frequent resets of the derivation pipeline,
			// which would otherwise occur if the blobs are not found immediately.
			return errclass.WrapAs(err, errclass.Transient)
		} else {
			// For other errors, treat them as persistent and do not retry
			return errclass.WrapAs(err, errclass.Persistent)
		}
	})

	// If all retries are exhausted and blobs are still not found
	if errors.Is(err, l1ethereum.NotFound) {
		// If the L1 block was available, then the blobs should be available too. The only
		// exception is if the blob retention window has expired, which we will ultimately handle
		// by failing over to a blob archival service.
		return nil, NewResetError(fmt.Errorf("failed to fetch blobs: %w", err))
	} else if err != nil {
		return nil, NewTemporaryError(fmt.Errorf("failed to fetch blobs: %w", err))
	}

	// go back over the data array and populate the blob pointers
	if err := fillBlobPointers(data, blobs); err != nil {
		// this shouldn't happen unless there is a bug in the blobs fetcher
		return nil, NewResetError(fmt.Errorf("failed to fill blob pointers: %w", err))
	}
	return data, nil
}

// dataAndHashesFromTxs extracts calldata and datahashes from the input transactions and returns them. It
// creates a placeholder blobOrCalldata element for each returned blob hash that must be populated
// by fillBlobPointers after blob bodies are retrieved.
func dataAndHashesFromTxs(txs l1types.Transactions, config *DataSourceConfig, batcherAddr common.Address) ([]blobOrCalldata, []eth.IndexedBlobHash) {
	data := []blobOrCalldata{}
	var hashes []eth.IndexedBlobHash
	blobIndex := 0 // index of each blob in the block's blob sidecar
	for _, tx := range txs {
		// skip any non-batcher transactions
		if !isValidBatchTx(tx, config.l1Signer, config.batchInboxAddress, batcherAddr) {
			blobIndex += len(tx.BlobHashes())
			continue
		}
		// handle non-blob batcher transactions by extracting their calldata
		if tx.Type() != types.BlobTxType {
			calldata := eth.Data(tx.Data())
			data = append(data, blobOrCalldata{nil, &calldata})
			continue
		}
		// handle blob batcher transactions by extracting their blob hashes, ignoring any calldata.
		if len(tx.Data()) > 0 {
			log.Warn("blob tx has calldata, which will be ignored", "txhash", tx.Hash())
		}
		for _, h := range tx.BlobHashes() {
			idh := eth.IndexedBlobHash{
				Index: uint64(blobIndex),
				Hash:  common.Hash(h),
			}
			hashes = append(hashes, idh)
			data = append(data, blobOrCalldata{nil, nil}) // will fill in blob pointers after we download them below
			blobIndex += 1
		}
	}
	return data, hashes
}

// fillBlobPointers goes back through the data array and fills in the pointers to the fetched blob
// bodies. There should be exactly one placeholder blobOrCalldata element for each blob, otherwise
// error is returned.
func fillBlobPointers(data []blobOrCalldata, blobs []*eth.Blob) error {
	blobIndex := 0
	for i := range data {
		if data[i].calldata != nil {
			continue
		}
		if blobIndex >= len(blobs) {
			return fmt.Errorf("didn't get enough blobs")
		}
		if blobs[blobIndex] == nil {
			return fmt.Errorf("found a nil blob")
		}
		data[i].blob = blobs[blobIndex]
		blobIndex++
	}
	if blobIndex != len(blobs) {
		return fmt.Errorf("got too many blobs")
	}
	return nil
}
