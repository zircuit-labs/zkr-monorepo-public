package status

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/testutils"
)

func mockL1BlockRef(num uint64) eth.L1BlockRef {
	return eth.L1BlockRef{Number: num, Hash: common.Hash{byte(num)}, ParentHash: common.Hash{byte(num - 1)}}
}

func newL1HeadEvent(l1Tracker *L1Tracker, head eth.L1BlockRef) {
	l1Tracker.OnEvent(L1UnsafeEvent{
		L1Unsafe: head,
	})
}

func TestCachingHeadReorg(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	l1Fetcher := testutils.NewMockL1Reader(ctrl)
	l1Tracker := NewL1Tracker(l1Fetcher)

	// no blocks added to cache yet
	l1Head := mockL1BlockRef(99)
	l1Fetcher.EXPECT().L1BlockRefByNumber(gomock.Any(), uint64(99)).Return(l1Head, nil).Times(1)
	ret, err := l1Tracker.L1BlockRefByNumber(ctx, 99)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(100)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 100)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(101)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(102)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 102)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// trigger a reorg of block 102
	l1Head = mockL1BlockRef(102)
	l1Head.Hash = common.Hash{0xde, 0xad, 0xbe, 0xef}
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 102)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// confirm that 101 is still in the cache
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, mockL1BlockRef(101), ret)
}

func TestCachingHeadRewind(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	l1Fetcher := testutils.NewMockL1Reader(ctrl)
	l1Tracker := NewL1Tracker(l1Fetcher)

	// no blocks added to cache yet
	l1Head := mockL1BlockRef(99)
	l1Fetcher.EXPECT().L1BlockRefByNumber(gomock.Any(), uint64(99)).Return(l1Head, nil).Times(1)
	ret, err := l1Tracker.L1BlockRefByNumber(ctx, 99)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(100)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 100)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(101)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(102)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 102)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// 101 is the new head, invalidating 102
	l1Head = mockL1BlockRef(101)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// confirm that 102 is no longer in the cache
	l1Head = mockL1BlockRef(102)
	l1Fetcher.EXPECT().L1BlockRefByNumber(gomock.Any(), uint64(102)).Return(l1Head, nil).Times(1)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 102)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// confirm that 101 is still in the cache
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, mockL1BlockRef(101), ret)
}

func TestCachingChainShorteningReorg(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	l1Fetcher := testutils.NewMockL1Reader(ctrl)
	l1Tracker := NewL1Tracker(l1Fetcher)

	// no blocks added to cache yet
	l1Head := mockL1BlockRef(99)
	l1Fetcher.EXPECT().L1BlockRefByNumber(gomock.Any(), uint64(99)).Return(l1Head, nil).Times(1)
	ret, err := l1Tracker.L1BlockRefByNumber(ctx, 99)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(100)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 100)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(101)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(102)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 102)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// trigger a reorg of block 101, invalidating the following cache elements (102)
	l1Head = mockL1BlockRef(101)
	l1Head.Hash = common.Hash{0xde, 0xad, 0xbe, 0xef}
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// confirm that 102 has been removed
	l1Fetcher.EXPECT().L1BlockRefByNumber(gomock.Any(), uint64(102)).Return(mockL1BlockRef(102), nil).Times(1)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 102)
	require.NoError(t, err)
	require.Equal(t, mockL1BlockRef(102), ret)
}

func TestCachingDeepReorg(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	l1Fetcher := testutils.NewMockL1Reader(ctrl)
	l1Tracker := NewL1Tracker(l1Fetcher)

	// from cache
	l1Head := mockL1BlockRef(100)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err := l1Tracker.L1BlockRefByNumber(ctx, 100)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(101)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(102)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 102)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// append a new block 102 based on a different 101, invalidating the entire cache
	parentHash := common.Hash{0xde, 0xad, 0xbe, 0xef}
	l1Head = mockL1BlockRef(102)
	l1Head.ParentHash = parentHash
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 102)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// confirm that the cache contains no 101
	l1Fetcher.EXPECT().L1BlockRefByNumber(gomock.Any(), uint64(101)).Return(mockL1BlockRef(101), nil).Times(1)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, mockL1BlockRef(101), ret)

	// confirm that the cache contains no 100
	l1Fetcher.EXPECT().L1BlockRefByNumber(gomock.Any(), uint64(100)).Return(mockL1BlockRef(100), nil).Times(1)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 100)
	require.NoError(t, err)
	require.Equal(t, mockL1BlockRef(100), ret)
}

func TestCachingSkipAhead(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	l1Fetcher := testutils.NewMockL1Reader(ctrl)
	l1Tracker := NewL1Tracker(l1Fetcher)

	// from cache
	l1Head := mockL1BlockRef(100)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err := l1Tracker.L1BlockRefByNumber(ctx, 100)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// from cache
	l1Head = mockL1BlockRef(101)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, l1Head, ret)

	// head jumps ahead from 101->103, invalidating the entire cache
	l1Head = mockL1BlockRef(103)
	newL1HeadEvent(l1Tracker, l1Head)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 103)
	require.NoError(t, err)
	require.Equal(t, mockL1BlockRef(103), ret)

	// confirm that the cache contains no 101
	l1Fetcher.EXPECT().L1BlockRefByNumber(gomock.Any(), uint64(101)).Return(mockL1BlockRef(101), nil).Times(1)
	ret, err = l1Tracker.L1BlockRefByNumber(ctx, 101)
	require.NoError(t, err)
	require.Equal(t, mockL1BlockRef(101), ret)
}

func TestCacheSizeEviction(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	l1Fetcher := testutils.NewMockL1Reader(ctrl)
	l1Tracker := NewL1Tracker(l1Fetcher)

	// insert 1000 elements into the cache
	for idx := 1000; idx < 2000; idx++ {
		l1Head := mockL1BlockRef(uint64(idx))
		newL1HeadEvent(l1Tracker, l1Head)
	}

	// request each element from cache
	for idx := 1000; idx < 2000; idx++ {
		ret, err := l1Tracker.L1BlockRefByNumber(ctx, uint64(idx))
		require.NoError(t, err)
		require.Equal(t, mockL1BlockRef(uint64(idx)), ret)
	}

	// insert 1001st element, removing the first
	l1Head := mockL1BlockRef(2000)
	newL1HeadEvent(l1Tracker, l1Head)

	// request first element, which now requires a live fetch instead
	l1Fetcher.EXPECT().L1BlockRefByNumber(gomock.Any(), uint64(1000)).Return(mockL1BlockRef(1000), nil).Times(1)
	ret, err := l1Tracker.L1BlockRefByNumber(ctx, 1000)
	require.NoError(t, err)
	require.Equal(t, mockL1BlockRef(1000), ret)
}
