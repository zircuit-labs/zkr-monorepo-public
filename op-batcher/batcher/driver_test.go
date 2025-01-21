package batcher

import (
	"context"
	"errors"
	"math/big"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/core/types"
	"github.com/zircuit-labs/l2-geth-public/ethclient"
	"github.com/zircuit-labs/l2-geth-public/log"
	"github.com/zircuit-labs/zkr-monorepo-public/op-batcher/metrics"
	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup"
	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup/derive"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/dial"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/testlog"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/testutils"
)

type mockL2EndpointProvider struct {
	ethClient       *testutils.MockL2Client
	ethClientErr    error
	rollupClient    *testutils.MockRollupClient
	rollupClientErr error
}

func newEndpointProvider() *mockL2EndpointProvider {
	return &mockL2EndpointProvider{
		ethClient:    new(testutils.MockL2Client),
		rollupClient: new(testutils.MockRollupClient),
	}
}

func (p *mockL2EndpointProvider) EthClient(context.Context) (dial.EthClientInterface, error) {
	return p.ethClient, p.ethClientErr
}

func (p *mockL2EndpointProvider) RollupClient(context.Context) (dial.RollupClientInterface, error) {
	return p.rollupClient, p.rollupClientErr
}

func (p *mockL2EndpointProvider) Close() {}

const genesisL1Origin = uint64(123)

func setup(t *testing.T) (*BatchSubmitter, *mockL2EndpointProvider) {
	ep := newEndpointProvider()

	cfg := defaultTestRollupConfig
	cfg.Genesis.L1.Number = genesisL1Origin

	return NewBatchSubmitter(DriverSetup{
		Log:              testlog.Logger(t, log.LevelDebug),
		Metr:             metrics.NoopMetrics,
		RollupConfig:     &cfg,
		EndpointProvider: ep,
	}), ep
}

func TestBatchSubmitter_loadBlockIntoState(t *testing.T) {
	rng := rand.New(rand.NewSource(1234))
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
	l2Parent := testutils.RandomL2BlockRef(rng)
	bs, ep := setup(t)
	defer ep.ethClient.AssertExpectations(t)

	block, _ := testutils.RandomBlock(rng, 4)
	bs.state.tip = block.ParentHash()
	mockL1Info := testutils.RandomBlockInfo(rng)
	l1Info, err := derive.L1BlockInfoFromParts(cfg, testSysCfg, 0, mockL1Info, l2Parent.Time, nil)
	require.NoError(t, err)

	t.Run("hyrax use l1Info provided by extended endpoint", func(t *testing.T) {
		ep.ethClient.ExpectBlockByNumberEx(block.Number(), &ethclient.BlockEx{
			Block:  block,
			L1Info: l1Info.L1Info,
		}, nil)

		resBlock, resL1Info, err := bs.loadBlockIntoState(context.Background(), block.NumberU64())
		require.NoError(t, err)
		require.Equal(t, block, resBlock)
		require.Equal(t, l1Info.L1Info, resL1Info)
	})

	t.Run("hyrax fallback to first transaction", func(t *testing.T) {
		dep, err := derive.L1InfoDeposit(cfg, testSysCfg, 0, mockL1Info, l2Parent.Time, nil)
		require.NoError(t, err)
		l1Tx := types.NewTx(dep)

		block = block.WithBody(append([]*types.Transaction{l1Tx}, block.Body().Transactions...), nil)
		bs.state.tip = block.ParentHash()

		ep.ethClient.ExpectBlockByNumberEx(block.Number(), &ethclient.BlockEx{
			Block:  block,
			L1Info: nil,
		}, nil)

		resBlock, resL1Info, err := bs.loadBlockIntoState(context.Background(), block.NumberU64())
		require.NoError(t, err)
		require.Equal(t, block, resBlock)
		require.Equal(t, l1Info.L1Info, resL1Info)
	})
}

func TestBatchSubmitter_SafeL1Origin(t *testing.T) {
	bs, ep := setup(t)

	tests := []struct {
		name                   string
		currentSafeOrigin      uint64
		failsToFetchSyncStatus bool
		expectResult           uint64
		expectErr              bool
	}{
		{
			name:              "ExistingSafeL1Origin",
			currentSafeOrigin: 999,
			expectResult:      999,
		},
		{
			name:              "NoExistingSafeL1OriginUsesGenesis",
			currentSafeOrigin: 0,
			expectResult:      genesisL1Origin,
		},
		{
			name:                   "ErrorFetchingSyncStatus",
			failsToFetchSyncStatus: true,
			expectErr:              true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.failsToFetchSyncStatus {
				ep.rollupClient.ExpectSyncStatus(&eth.SyncStatus{}, errors.New("failed to fetch sync status"))
			} else {
				ep.rollupClient.ExpectSyncStatus(&eth.SyncStatus{
					SafeL2: eth.L2BlockRef{
						L1Origin: eth.BlockID{
							Number: tt.currentSafeOrigin,
						},
					},
				}, nil)
			}

			id, err := bs.safeL1Origin(context.Background())

			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectResult, id.Number)
			}
		})
	}
}

func TestBatchSubmitter_SafeL1Origin_FailsToResolveRollupClient(t *testing.T) {
	bs, ep := setup(t)

	ep.rollupClientErr = errors.New("failed to resolve rollup client")

	_, err := bs.safeL1Origin(context.Background())
	require.Error(t, err)
}
