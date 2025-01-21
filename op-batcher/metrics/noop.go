package metrics

import (
	"io"

	l1common "github.com/ethereum/go-ethereum/common"
	l1ethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/zircuit-labs/l2-geth-public/core/types"
	"github.com/zircuit-labs/l2-geth-public/log"

	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup/derive"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
	opmetrics "github.com/zircuit-labs/zkr-monorepo-public/op-service/metrics"
	txmetrics "github.com/zircuit-labs/zkr-monorepo-public/op-service/txmgr/metrics"
)

type noopMetrics struct {
	opmetrics.NoopRefMetrics
	txmetrics.NoopTxMetrics
	opmetrics.NoopRPCMetrics
}

var NoopMetrics Metricer = new(noopMetrics)

func (*noopMetrics) Document() []opmetrics.DocumentedMetric { return nil }

func (*noopMetrics) RecordInfo(version string) {}
func (*noopMetrics) RecordUp()                 {}

func (*noopMetrics) RecordLatestL1Block(l1ref eth.L1BlockRef)               {}
func (*noopMetrics) RecordL2BlocksLoaded(eth.L2BlockRef)                    {}
func (*noopMetrics) RecordChannelOpened(derive.ChannelID, int)              {}
func (*noopMetrics) RecordL2BlocksAdded(eth.L2BlockRef, int, int, int, int) {}
func (*noopMetrics) RecordL2BlockInPendingQueue(*types.Block)               {}
func (*noopMetrics) RecordL2BlockInChannel(*types.Block)                    {}

func (*noopMetrics) RecordChannelClosed(derive.ChannelID, int, int, int, int, error) {}

func (*noopMetrics) RecordChannelFullySubmitted(derive.ChannelID) {}
func (*noopMetrics) RecordChannelTimedOut(derive.ChannelID)       {}

func (*noopMetrics) RecordBatchTxSubmitted() {}
func (*noopMetrics) RecordBatchTxSuccess()   {}
func (*noopMetrics) RecordBatchTxFailed()    {}
func (*noopMetrics) RecordBlobUsedBytes(int) {}
func (*noopMetrics) StartBalanceMetrics(log.Logger, *l1ethclient.Client, l1common.Address) io.Closer {
	return nil
}
