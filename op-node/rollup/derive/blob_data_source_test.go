package derive

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	l1ethereum "github.com/ethereum/go-ethereum"
	l1types "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	l1common "github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/zircuit-labs/l2-geth-public/crypto"

	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/testutils"
)

func TestDataAndHashesFromTxs(t *testing.T) {
	// test setup
	rng := rand.New(rand.NewSource(12345))
	privateKey := testutils.InsecureRandomKey(rng)
	publicKey, _ := privateKey.Public().(*ecdsa.PublicKey)
	batcherAddr := crypto.PubkeyToAddress(*publicKey)
	batchInboxAddr := testutils.RandomAddress(rng)

	chainId := new(big.Int).SetUint64(rng.Uint64())
	signer := l1types.NewCancunSigner(chainId)
	config := DataSourceConfig{
		l1Signer:          signer,
		batchInboxAddress: batchInboxAddr,
	}

	// create a valid non-blob batcher transaction and make sure it's picked up
	toAddr := l1common.Address(batchInboxAddr)
	txData := &l1types.LegacyTx{
		Nonce:    rng.Uint64(),
		GasPrice: new(big.Int).SetUint64(rng.Uint64()),
		Gas:      2_000_000,
		To:       &toAddr,
		Value:    big.NewInt(10),
		Data:     testutils.RandomData(rng, rng.Intn(1000)),
	}
	calldataTx, _ := l1types.SignNewTx(privateKey, signer, txData)
	txs := l1types.Transactions{calldataTx}
	data, blobHashes := dataAndHashesFromTxs(txs, &config, batcherAddr)
	require.Equal(t, 1, len(data))
	require.Equal(t, 0, len(blobHashes))

	// create a valid blob batcher tx and make sure it's picked up
	blobHash := testutils.RandomL1Hash(rng)
	blobTxData := &l1types.BlobTx{
		Nonce:      rng.Uint64(),
		Gas:        2_000_000,
		To:         l1common.Address(batchInboxAddr),
		Data:       testutils.RandomData(rng, rng.Intn(1000)),
		BlobHashes: []l1common.Hash{blobHash},
	}
	blobTx, _ := l1types.SignNewTx(privateKey, signer, blobTxData)
	txs = l1types.Transactions{blobTx}
	data, blobHashes = dataAndHashesFromTxs(txs, &config, batcherAddr)
	require.Equal(t, 1, len(data))
	require.Equal(t, 1, len(blobHashes))
	require.Nil(t, data[0].calldata)

	// try again with both the blob & calldata transactions and make sure both are picked up
	txs = l1types.Transactions{blobTx, calldataTx}
	data, blobHashes = dataAndHashesFromTxs(txs, &config, batcherAddr)
	require.Equal(t, 2, len(data))
	require.Equal(t, 1, len(blobHashes))
	require.NotNil(t, data[1].calldata)

	// make sure blob tx to the batch inbox is ignored if not signed by the batcher
	blobTx, _ = l1types.SignNewTx(testutils.RandomKey(), signer, blobTxData)
	txs = l1types.Transactions{blobTx}
	data, blobHashes = dataAndHashesFromTxs(txs, &config, batcherAddr)
	require.Equal(t, 0, len(data))
	require.Equal(t, 0, len(blobHashes))

	// make sure blob tx ignored if the tx isn't going to the batch inbox addr, even if the
	// signature is valid.
	blobTxData.To = testutils.RandomL1Address(rng)
	blobTx, _ = l1types.SignNewTx(privateKey, signer, blobTxData)
	txs = l1types.Transactions{blobTx}
	data, blobHashes = dataAndHashesFromTxs(txs, &config, batcherAddr)
	require.Equal(t, 0, len(data))
	require.Equal(t, 0, len(blobHashes))
}

func TestFillBlobPointers(t *testing.T) {
	blob := eth.Blob{}
	rng := rand.New(rand.NewSource(1234))
	calldata := eth.Data{}

	for i := 0; i < 100; i++ {
		// create a random length input data array w/ len = [0-10)
		dataLen := rng.Intn(10)
		data := make([]blobOrCalldata, dataLen)

		// pick some subset of those to be blobs, and the rest calldata
		blobLen := 0
		if dataLen != 0 {
			blobLen = rng.Intn(dataLen)
		}
		calldataLen := dataLen - blobLen

		// fill in the calldata entries at random indices
		for j := 0; j < calldataLen; j++ {
			randomIndex := rng.Intn(dataLen)
			for data[randomIndex].calldata != nil {
				randomIndex = (randomIndex + 1) % dataLen
			}
			data[randomIndex].calldata = &calldata
		}

		// create the input blobs array and call fillBlobPointers on it
		blobs := make([]*eth.Blob, blobLen)
		for j := 0; j < blobLen; j++ {
			blobs[j] = &blob
		}
		err := fillBlobPointers(data, blobs)
		require.NoError(t, err)

		// check that we get the expected number of calldata vs blobs results
		blobCount := 0
		calldataCount := 0
		for j := 0; j < dataLen; j++ {
			if data[j].calldata != nil {
				calldataCount++
			}
			if data[j].blob != nil {
				blobCount++
			}
		}
		require.Equal(t, blobLen, blobCount)
		require.Equal(t, calldataLen, calldataCount)
	}
}

func TestBlobDataSourceGetBlobsRetry(t *testing.T) {
	ctx := context.Background()

	// Common setup function for creating mock fetchers, random block info, and blob transactions
	setupMocks := func() (*testutils.MockL1Reader, *testutils.MockBlobsFetcher, eth.L1BlockRef, common.Hash, []eth.IndexedBlobHash, DataSourceConfig, common.Address) {
		rng := rand.New(rand.NewSource(12345))
		ctrl := gomock.NewController(t)
		mockFetcher := testutils.NewMockL1Reader(ctrl)
		mockBlobsFetcher := new(testutils.MockBlobsFetcher)
		mockBlockInfo := testutils.RandomBlockInfo(rng)

		// Ensure that the batchInboxAddress and batcherAddr are correctly set for valid transactions
		batchInboxAddress := testutils.RandomAddress(rng)
		batcherPrivateKey := testutils.InsecureRandomKey(rng)
		batcherAddr := crypto.PubkeyToAddress(batcherPrivateKey.PublicKey)

		// Set up DataSourceConfig with appropriate signer and addresses
		dsConfig := DataSourceConfig{
			l1Signer:          l1types.NewCancunSigner(big.NewInt(1)),
			batchInboxAddress: batchInboxAddress,
		}

		// Create a valid BlobTx transaction with blob hashes
		blobHash := testutils.RandomL1Hash(rng)
		blobTxData := &l1types.BlobTx{
			Nonce:      rng.Uint64(),
			Gas:        2_000_000,
			To:         l1common.Address(batchInboxAddress),
			Data:       testutils.RandomData(rng, rng.Intn(1000)),
			BlobHashes: []l1common.Hash{blobHash},
		}
		blobTx, _ := l1types.SignNewTx(batcherPrivateKey, dsConfig.l1Signer, blobTxData)

		txs := l1types.Transactions{blobTx}

		// Use the hash of the generated mockBlockInfo
		hash := mockBlockInfo.Hash()
		mockFetcher.EXPECT().InfoAndTxsByHash(gomock.Any(), hash).Return(mockBlockInfo, txs, nil).Times(1)
		hashes := []eth.IndexedBlobHash{{Index: 0, Hash: common.Hash(blobHash)}}

		ref := eth.L1BlockRef{
			Hash:       mockBlockInfo.Hash(),
			Number:     mockBlockInfo.NumberU64(),
			ParentHash: mockBlockInfo.ParentHash(),
			Time:       mockBlockInfo.Time(),
		}

		return mockFetcher, mockBlobsFetcher, ref, hash, hashes, dsConfig, batcherAddr
	}

	// Helper function to create BlobDataSource with provided mocks
	createBlobDataSource := func(mockFetcher *testutils.MockL1Reader, mockBlobsFetcher *testutils.MockBlobsFetcher, dsConfig DataSourceConfig, batcherAddr common.Address) *BlobDataSource {
		return &BlobDataSource{
			fetcher:      mockFetcher,
			blobsFetcher: mockBlobsFetcher,
			dsCfg:        dsConfig,
			batcherAddr:  batcherAddr,
		}
	}

	assertExpectations := func(t *testing.T, mockBlobsFetcher *testutils.MockBlobsFetcher) {
		mockBlobsFetcher.AssertExpectations(t)
	}

	// assertDurationInRange asserts that the elapsed time is within the expected range (inclusive).
	assertDurationInRange := func(t *testing.T, elapsedTime, minDuration, maxDuration time.Duration, message string) {
		require.GreaterOrEqual(t, elapsedTime, minDuration, message)
		require.LessOrEqual(t, elapsedTime, maxDuration, message)
	}

	t.Run("Success without errors", func(t *testing.T) {
		t.Parallel()
		mockFetcher, mockBlobsFetcher, ref, hash, hashes, dsConfig, batcherAddr := setupMocks()

		// Mock GetBlobs to return blobs without error
		blobs := []*eth.Blob{{}}
		mockBlobsFetcher.ExpectOnGetBlobs(ctx, ref, hashes, blobs, nil)

		ds := createBlobDataSource(mockFetcher, mockBlobsFetcher, dsConfig, batcherAddr)

		ds.ref = ref
		ds.ref.Hash = hash

		_, err := ds.open(ctx)
		require.NoError(t, err)
		assertExpectations(t, mockBlobsFetcher)
	})

	t.Run("Should limits retries to 3 times when GetBlobs failed with ethereum.NotFound error (404)", func(t *testing.T) {
		t.Parallel()
		mockFetcher, mockBlobsFetcher, ref, hash, hashes, dsConfig, batcherAddr := setupMocks()

		// All 3 calls are failed with 404 error
		mockBlobsFetcher.Mock.On("GetBlobs", ref, hashes).Return(([]*eth.Blob)(nil), l1ethereum.NotFound).Times(3)

		// Set up BlobDataSource with mocks
		ds := createBlobDataSource(mockFetcher, mockBlobsFetcher, dsConfig, batcherAddr)
		ds.ref = ref
		ds.ref.Hash = hash
		startTime := time.Now()
		_, err := ds.open(ctx)
		elapsedTime := time.Since(startTime)

		require.EqualError(t, err, "reset: failed to fetch blobs: not found")
		require.ErrorIs(t, err, NewResetError(fmt.Errorf("failed to fetch blobs: %w", l1ethereum.NotFound)))

		// Assert that the maximum 3 try attempts for base 3 (1, 3, 9 seconds for each retry and total 13 seconds)
		assertDurationInRange(t, elapsedTime, 13*time.Second, 14*time.Second, "The retries should wait for 13 seconds")

		assertExpectations(t, mockBlobsFetcher)
	})

	t.Run("Should break retry loop when fetching blobs fails with an error other than ethereum.NotFound", func(t *testing.T) {
		t.Parallel()
		mockFetcher, mockBlobsFetcher, ref, hash, hashes, dsConfig, batcherAddr := setupMocks()

		// Failed at the first call with 404 error
		mockBlobsFetcher.Mock.On("GetBlobs", ref, hashes).Return(([]*eth.Blob)(nil), l1ethereum.NotFound).Once()
		// Retry failed with a non 404 error
		mockBlobsFetcher.Mock.On("GetBlobs", ref, hashes).Return(([]*eth.Blob)(nil), fmt.Errorf("some temporary error")).Once()

		ds := createBlobDataSource(mockFetcher, mockBlobsFetcher, dsConfig, batcherAddr)
		ds.ref = ref
		ds.ref.Hash = hash

		startTime := time.Now()
		_, err := ds.open(ctx)
		elapsedTime := time.Since(startTime)

		require.EqualError(t, err, "temp: failed to fetch blobs: some temporary error")
		require.ErrorIs(t, err, NewTemporaryError(fmt.Errorf("failed to fetch blobs: %w", fmt.Errorf("some temporary error"))))
		// The first delay is 1 second
		assertDurationInRange(t, elapsedTime, 1*time.Second, 2*time.Second, "Retry loop should break early with non-404 error")

		assertExpectations(t, mockBlobsFetcher)
	})

	t.Run("Should only retry when fetching blobs fails with an ethereum.NotFound error", func(t *testing.T) {
		t.Parallel()
		mockFetcher, mockBlobsFetcher, ref, hash, hashes, dsConfig, batcherAddr := setupMocks()
		// Failed with some random error (non 404)
		mockBlobsFetcher.Mock.On("GetBlobs", ref, hashes).Return(([]*eth.Blob)(nil), fmt.Errorf("some temporary error")).Once()

		ds := createBlobDataSource(mockFetcher, mockBlobsFetcher, dsConfig, batcherAddr)
		ds.ref = ref
		ds.ref.Hash = hash

		startTime := time.Now()
		_, err := ds.open(ctx)
		elapsedTime := time.Since(startTime)

		require.EqualError(t, err, "temp: failed to fetch blobs: some temporary error")
		require.ErrorIs(t, err, NewTemporaryError(fmt.Errorf("failed to fetch blobs: %w", fmt.Errorf("some temporary error"))))
		require.Less(t, elapsedTime, 1*time.Second, "No retries should occur on non-404 error")

		assertExpectations(t, mockBlobsFetcher)
	})
}
