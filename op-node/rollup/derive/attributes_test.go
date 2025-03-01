package derive

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	l1common "github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/core/types"

	"github.com/zircuit-labs/zkr-monorepo-public/op-bindings/predeploys"
	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/testutils"
)

func TestPreparePayloadAttributes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	l1Fetcher := testutils.NewMockL1Reader(ctrl)

	// test sysCfg, only init the necessary fields
	cfg := &rollup.Config{
		BlockTime:              2,
		L1ChainID:              big.NewInt(101),
		L2ChainID:              big.NewInt(102),
		DepositContractAddress: common.Address{0xbb},
		L1SystemConfigAddress:  common.Address{0xcc},
	}

	testSysCfg := eth.SystemConfig{
		BatcherAddr: common.Address{42},
		Overhead:    [32]byte{},
		Scalar:      [32]byte{},
	}

	t.Run("inconsistent next height origin", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		l2Parent := testutils.RandomL2BlockRef(rng)
		l1CfgFetcher := &testutils.MockL2Client{}
		l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
		defer l1CfgFetcher.AssertExpectations(t)
		l1Info := testutils.RandomBlockInfo(rng)
		l1Info.InfoNum = l2Parent.L1Origin.Number + 1
		epoch := l1Info.ID()
		l1Fetcher.EXPECT().
			FetchReceipts(gomock.Any(), epoch.Hash).
			Return(l1Info, nil, nil).
			Times(1)
		attrBuilder := NewFetchingAttributesBuilder(cfg, l1Fetcher, l1CfgFetcher)
		_, err := attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
		require.NotNil(t, err, "inconsistent L1 origin error expected")
		require.ErrorIs(t, err, ErrReset, "inconsistent L1 origin transition must be handled like a critical error with reorg")
	})
	t.Run("inconsistent equal height origin", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		l2Parent := testutils.RandomL2BlockRef(rng)
		l1CfgFetcher := &testutils.MockL2Client{}
		l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
		defer l1CfgFetcher.AssertExpectations(t)
		l1Info := testutils.RandomBlockInfo(rng)
		l1Info.InfoNum = l2Parent.L1Origin.Number
		epoch := l1Info.ID()
		attrBuilder := NewFetchingAttributesBuilder(cfg, l1Fetcher, l1CfgFetcher)
		_, err := attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
		require.NotNil(t, err, "inconsistent L1 origin error expected")
		require.ErrorIs(t, err, ErrReset, "inconsistent L1 origin transition must be handled like a critical error with reorg")
	})
	t.Run("rpc fail Fetch", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		l2Parent := testutils.RandomL2BlockRef(rng)
		l1CfgFetcher := &testutils.MockL2Client{}
		l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
		defer l1CfgFetcher.AssertExpectations(t)
		epoch := l2Parent.L1Origin
		epoch.Number += 1
		mockRPCErr := errors.New("mock rpc error")
		l1Fetcher.EXPECT().
			FetchReceipts(gomock.Any(), epoch.Hash).
			Return(nil, nil, mockRPCErr).
			Times(1)
		attrBuilder := NewFetchingAttributesBuilder(cfg, l1Fetcher, l1CfgFetcher)
		_, err := attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
		require.ErrorIs(t, err, mockRPCErr, "mock rpc error expected")
		require.ErrorIs(t, err, ErrTemporary, "rpc errors should not be critical, it is not necessary to reorg")
	})
	t.Run("rpc fail InfoByHash", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		l2Parent := testutils.RandomL2BlockRef(rng)
		l1CfgFetcher := &testutils.MockL2Client{}
		l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
		defer l1CfgFetcher.AssertExpectations(t)
		epoch := l2Parent.L1Origin
		mockRPCErr := errors.New("mock rpc error")
		l1Fetcher.EXPECT().
			InfoByHash(gomock.Any(), epoch.Hash).
			Return(nil, mockRPCErr).
			Times(1)
		attrBuilder := NewFetchingAttributesBuilder(cfg, l1Fetcher, l1CfgFetcher)
		_, err := attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
		require.ErrorIs(t, err, mockRPCErr, "mock rpc error expected")
		require.ErrorIs(t, err, ErrTemporary, "rpc errors should not be critical, it is not necessary to reorg")
	})
	t.Run("next origin without deposits", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		l2Parent := testutils.RandomL2BlockRef(rng)
		l1CfgFetcher := &testutils.MockL2Client{}
		l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
		defer l1CfgFetcher.AssertExpectations(t)
		l1Info := testutils.RandomBlockInfo(rng)
		l1Info.InfoParentHash = l2Parent.L1Origin.Hash
		l1Info.InfoNum = l2Parent.L1Origin.Number + 1
		epoch := l1Info.ID()
		l1InfoTx, err := L1InfoDepositBytes(cfg, testSysCfg, 0, l1Info, nil, 0)
		require.NoError(t, err)
		l1Fetcher.EXPECT().
			FetchReceipts(gomock.Any(), epoch.Hash).
			Return(l1Info, nil, nil).
			Times(1)
		attrBuilder := NewFetchingAttributesBuilder(cfg, l1Fetcher, l1CfgFetcher)
		attrs, err := attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
		require.NoError(t, err)
		require.NotNil(t, attrs)
		require.Equal(t, l2Parent.Time+cfg.BlockTime, uint64(attrs.Timestamp))
		require.Equal(t, eth.Bytes32(l1Info.InfoMixDigest), attrs.PrevRandao)
		require.Equal(t, predeploys.SequencerFeeVaultAddr, attrs.SuggestedFeeRecipient)
		require.Equal(t, 1, len(attrs.Transactions))
		require.Equal(t, l1InfoTx, []byte(attrs.Transactions[0]))
		require.True(t, attrs.NoTxPool)
	})
	t.Run("next origin with deposits", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		l2Parent := testutils.RandomL2BlockRef(rng)
		l1CfgFetcher := &testutils.MockL2Client{}
		l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
		defer l1CfgFetcher.AssertExpectations(t)
		l1Info := testutils.RandomBlockInfo(rng)
		l1Info.InfoParentHash = l2Parent.L1Origin.Hash
		l1Info.InfoNum = l2Parent.L1Origin.Number + 1

		receipts, depositTxs, err := makeReceipts(rng, l1common.Hash(l1Info.InfoHash), cfg.DepositContractAddress, []receiptData{
			{goodReceipt: true, DepositLogs: []bool{true, false}},
			{goodReceipt: true, DepositLogs: []bool{true}},
			{goodReceipt: false, DepositLogs: []bool{true}},
			{goodReceipt: false, DepositLogs: []bool{false}},
		})
		require.NoError(t, err)
		usedDepositTxs, err := encodeDeposits(depositTxs)
		require.NoError(t, err)

		epoch := l1Info.ID()
		l1InfoTx, err := L1InfoDepositBytes(cfg, testSysCfg, 0, l1Info, nil, 0)
		require.NoError(t, err)

		l2Txs := append(append(make([]eth.Data, 0), l1InfoTx), usedDepositTxs...)

		l1Fetcher.EXPECT().
			FetchReceipts(gomock.Any(), epoch.Hash).
			Return(l1Info, receipts, nil).
			Times(1)
		attrBuilder := NewFetchingAttributesBuilder(cfg, l1Fetcher, l1CfgFetcher)
		attrs, err := attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
		require.NoError(t, err)
		require.NotNil(t, attrs)
		require.Equal(t, l2Parent.Time+cfg.BlockTime, uint64(attrs.Timestamp))
		require.Equal(t, eth.Bytes32(l1Info.InfoMixDigest), attrs.PrevRandao)
		require.Equal(t, predeploys.SequencerFeeVaultAddr, attrs.SuggestedFeeRecipient)
		require.Equal(t, len(l2Txs), len(attrs.Transactions), "Expected txs to equal l1 info tx + user deposit txs")
		require.Equal(t, l2Txs, attrs.Transactions)
		require.True(t, attrs.NoTxPool)
	})
	t.Run("same origin again", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		l2Parent := testutils.RandomL2BlockRef(rng)
		l1CfgFetcher := &testutils.MockL2Client{}
		l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
		defer l1CfgFetcher.AssertExpectations(t)
		l1Info := testutils.RandomBlockInfo(rng)
		l1Info.InfoHash = l2Parent.L1Origin.Hash
		l1Info.InfoNum = l2Parent.L1Origin.Number

		epoch := l1Info.ID()
		l1InfoTx, err := L1InfoDepositBytes(cfg, testSysCfg, l2Parent.SequenceNumber+1, l1Info, nil, 0)
		require.NoError(t, err)

		l1Fetcher.EXPECT().
			InfoByHash(gomock.Any(), epoch.Hash).
			Return(l1Info, nil).
			Times(1)
		attrBuilder := NewFetchingAttributesBuilder(cfg, l1Fetcher, l1CfgFetcher)
		attrs, err := attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
		require.NoError(t, err)
		require.NotNil(t, attrs)
		require.Equal(t, l2Parent.Time+cfg.BlockTime, uint64(attrs.Timestamp))
		require.Equal(t, eth.Bytes32(l1Info.InfoMixDigest), attrs.PrevRandao)
		require.Equal(t, predeploys.SequencerFeeVaultAddr, attrs.SuggestedFeeRecipient)
		require.Equal(t, 1, len(attrs.Transactions))
		require.Equal(t, l1InfoTx, []byte(attrs.Transactions[0]))
		require.True(t, attrs.NoTxPool)
	})
	// Test that the payload attributes builder changes the deposit format based on L2-time-based regolith activation
	t.Run("regolith", func(t *testing.T) {
		testCases := []struct {
			name         string
			l1Time       uint64
			l2ParentTime uint64
			regolithTime uint64
			regolith     bool
		}{
			{"exactly", 900, 1000 - cfg.BlockTime, 1000, true},
			{"almost", 900, 1000 - cfg.BlockTime - 1, 1000, false},
			{"inactive", 700, 700, 1000, false},
			{"l1 time before regolith", 1000, 1001, 1001, true},
			{"l1 time way before regolith", 1000, 2000, 2000, true},
			{"l1 time before regoltih and l2 after", 1000, 3000, 2000, true},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				cfgCopy := *cfg // copy, we are making regolith config modifications
				cfg := &cfgCopy
				rng := rand.New(rand.NewSource(1234))
				l2Parent := testutils.RandomL2BlockRef(rng)
				cfg.RegolithTime = &tc.regolithTime
				l2Parent.Time = tc.l2ParentTime

				l1CfgFetcher := &testutils.MockL2Client{}
				l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
				defer l1CfgFetcher.AssertExpectations(t)

				l1Info := testutils.RandomBlockInfo(rng)
				l1Info.InfoParentHash = l2Parent.L1Origin.Hash
				l1Info.InfoNum = l2Parent.L1Origin.Number + 1
				l1Info.InfoTime = tc.l1Time

				epoch := l1Info.ID()
				time := tc.regolithTime
				if !tc.regolith {
					time--
				}
				l1InfoTx, err := L1InfoDepositBytes(cfg, testSysCfg, 0, l1Info, nil, time)
				require.NoError(t, err)
				l1Fetcher.EXPECT().
					FetchReceipts(gomock.Any(), epoch.Hash).
					Return(l1Info, nil, nil).
					Times(1)
				attrBuilder := NewFetchingAttributesBuilder(cfg, l1Fetcher, l1CfgFetcher)
				attrs, err := attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
				require.NoError(t, err)
				require.Equal(t, l1InfoTx, []byte(attrs.Transactions[0]))
			})
		}
	})

	t.Run("hyrax", func(t *testing.T) {
		baseTime := uint64(1000)
		t.Run("l1 origin change", func(t *testing.T) {
			cfgCopy := *cfg // copy, we are making regolith config modifications
			cfg := &cfgCopy
			cfg.HyraxTime = &baseTime

			rng := rand.New(rand.NewSource(1234))
			l1Fetcher := &testutils.MockL1Source{}
			defer l1Fetcher.AssertExpectations(t)

			l2Parent := testutils.RandomL2BlockRef(rng)
			l2Parent.Time = baseTime + 1

			l1CfgFetcher := &testutils.MockL2Client{}
			defer l1CfgFetcher.AssertExpectations(t)

			attrBuilder := NewFetchingAttributesBuilder(cfg, l1Fetcher, l1CfgFetcher)

			l1Info := testutils.RandomBlockInfo(rng)
			l1Info.InfoParentHash = l2Parent.L1Origin.Hash
			l1Info.InfoNum = l2Parent.L1Origin.Number + 1
			l1Info.InfoTime = l2Parent.Time

			epoch := l1Info.ID()

			// New epoch, l1 origin changed and isHyrax
			l1InfoTx, err := L1InfoDepositBytes(cfg, testSysCfg, 0, l1Info, nil, l2Parent.Time)
			require.NoError(t, err)
			l1Fetcher.ExpectFetchReceipts(epoch.Hash, l1Info, nil, nil)
			l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)

			attrs, err := attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
			require.NoError(t, err)
			require.Equal(t, l1InfoTx, []byte(attrs.Transactions[0]))

			// Same epoch, l1 origin didn't changed and isHyrax
			epoch.Number = l2Parent.L1Origin.Number
			epoch.Hash = l2Parent.L1Origin.Hash
			l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
			l1Fetcher.ExpectInfoByHash(epoch.Hash, l1Info, nil)
			attrs, err = attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
			require.NoError(t, err)
			require.Equal(t, 0, len(attrs.Transactions))

			// Same epoch, l1 origin didn't changed and not isHyrax
			*cfg.HyraxTime += 100
			l1CfgFetcher.ExpectSystemConfigByL2Hash(l2Parent.Hash, testSysCfg, nil)
			l1Fetcher.ExpectInfoByHash(epoch.Hash, l1Info, nil)

			l1InfoTx, err = L1InfoDepositBytes(cfg, testSysCfg, l2Parent.SequenceNumber+1, l1Info, nil, l2Parent.Time)
			require.NoError(t, err)

			attrs, err = attrBuilder.PreparePayloadAttributes(context.Background(), l2Parent, epoch, nil)
			require.NoError(t, err)
			require.Equal(t, 1, len(attrs.Transactions))
			require.Equal(t, l1InfoTx, []byte(attrs.Transactions[0]))
		})
	})
}

func encodeDeposits(deposits []*types.DepositTx) (out []eth.Data, err error) {
	for i, tx := range deposits {
		opaqueTx, err := types.NewTx(tx).MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("bad deposit %d: %w", i, err)
		}
		out = append(out, opaqueTx)
	}
	return
}
