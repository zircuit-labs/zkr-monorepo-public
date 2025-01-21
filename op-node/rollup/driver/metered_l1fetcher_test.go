package driver

import (
	"context"
	"errors"
	"testing"
	"time"

	l1types "github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/testutils"
)

func TestDurationRecorded(t *testing.T) {
	num := uint64(1234)
	hash := common.Hash{0xaa}
	ref := eth.L1BlockRef{Number: num}
	info := &testutils.MockBlockInfo{}
	expectedErr := errors.New("test error")

	tests := []struct {
		method string
		expect func(inner *testutils.MockL1Source)
		call   func(t *testing.T, fetcher *MeteredL1Fetcher, inner *testutils.MockL1Reader)
	}{
		{
			method: "L1BlockRefByLabel",
			call: func(t *testing.T, fetcher *MeteredL1Fetcher, inner *testutils.MockL1Reader) {
				inner.EXPECT().L1BlockRefByLabel(gomock.Any(), eth.BlockLabel(eth.Finalized)).Return(ref, expectedErr).Times(1)

				result, err := fetcher.L1BlockRefByLabel(context.Background(), eth.Finalized)
				require.Equal(t, ref, result)
				require.Equal(t, expectedErr, err)
			},
		},
		{
			method: "L1BlockRefByNumber",
			call: func(t *testing.T, fetcher *MeteredL1Fetcher, inner *testutils.MockL1Reader) {
				inner.EXPECT().L1BlockRefByNumber(gomock.Any(), num).Return(ref, expectedErr).Times(1)

				result, err := fetcher.L1BlockRefByNumber(context.Background(), num)
				require.Equal(t, ref, result)
				require.Equal(t, expectedErr, err)
			},
		},
		{
			method: "L1BlockRefByHash",
			call: func(t *testing.T, fetcher *MeteredL1Fetcher, inner *testutils.MockL1Reader) {
				inner.EXPECT().L1BlockRefByHash(gomock.Any(), hash).Return(ref, expectedErr).Times(1)

				result, err := fetcher.L1BlockRefByHash(context.Background(), hash)
				require.Equal(t, ref, result)
				require.Equal(t, expectedErr, err)
			},
		},
		{
			method: "InfoByHash",
			call: func(t *testing.T, fetcher *MeteredL1Fetcher, inner *testutils.MockL1Reader) {
				inner.EXPECT().InfoByHash(gomock.Any(), hash).Return(info, expectedErr).Times(1)

				result, err := fetcher.InfoByHash(context.Background(), hash)
				require.Equal(t, info, result)
				require.Equal(t, expectedErr, err)
			},
		},
		{
			method: "InfoAndTxsByHash",
			call: func(t *testing.T, fetcher *MeteredL1Fetcher, inner *testutils.MockL1Reader) {
				txs := l1types.Transactions{
					&l1types.Transaction{},
				}
				inner.EXPECT().InfoAndTxsByHash(gomock.Any(), hash).Return(info, txs, expectedErr).Times(1)

				actualInfo, actualTxs, err := fetcher.InfoAndTxsByHash(context.Background(), hash)
				require.Equal(t, info, actualInfo)
				require.Equal(t, txs, actualTxs)
				require.Equal(t, expectedErr, err)
			},
		},
		{
			method: "FetchReceipts",
			call: func(t *testing.T, fetcher *MeteredL1Fetcher, inner *testutils.MockL1Reader) {
				rcpts := l1types.Receipts{
					&l1types.Receipt{},
				}
				inner.EXPECT().FetchReceipts(gomock.Any(), hash).Return(info, rcpts, expectedErr).Times(1)

				actualInfo, actualRcpts, err := fetcher.FetchReceipts(context.Background(), hash)
				require.Equal(t, info, actualInfo)
				require.Equal(t, rcpts, actualRcpts)
				require.Equal(t, expectedErr, err)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.method, func(t *testing.T) {
			duration := 200 * time.Millisecond
			fetcher, inner, metrics := createFetcher(t, duration)
			defer metrics.AssertExpectations(t)

			metrics.ExpectRecordRequestTime(test.method, duration)

			test.call(t, fetcher, inner)
		})
	}
}

// createFetcher creates a MeteredL1Fetcher with a mock inner.
// The clock used to calculate the current time will advance by clockIncrement on each call, making it appear as if
// each request takes that amount of time to execute.
func createFetcher(t *testing.T, clockIncrement time.Duration) (*MeteredL1Fetcher, *testutils.MockL1Reader, *mockMetrics) {
	ctrl := gomock.NewController(t)
	l1 := testutils.NewMockL1Reader(ctrl)

	currTime := time.UnixMilli(1294812934000000)
	clock := func() time.Time {
		currTime = currTime.Add(clockIncrement)
		return currTime
	}
	metrics := &mockMetrics{}
	fetcher := MeteredL1Fetcher{
		inner:   l1,
		metrics: metrics,
		now:     clock,
	}
	return &fetcher, l1, metrics
}

type mockMetrics struct {
	mock.Mock
}

func (m *mockMetrics) RecordL1RequestTime(method string, duration time.Duration) {
	m.MethodCalled("RecordL1RequestTime", method, duration)
}

func (m *mockMetrics) ExpectRecordRequestTime(method string, duration time.Duration) {
	m.On("RecordL1RequestTime", method, duration).Once()
}
