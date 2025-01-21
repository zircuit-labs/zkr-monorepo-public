package metrics

import (
	"math/big"

	"github.com/zircuit-labs/l2-geth-public/core/types"
)

type NoopTxMetrics struct{}

func (*NoopTxMetrics) RecordNonce(uint64)                {}
func (*NoopTxMetrics) RecordPendingTx(int64)             {}
func (*NoopTxMetrics) RecordGasBumpCount(int)            {}
func (*NoopTxMetrics) RecordTxConfirmationLatency(int64) {}
func (*NoopTxMetrics) TxConfirmed(*types.Receipt)        {}
func (*NoopTxMetrics) TxPublished(string)                {}
func (*NoopTxMetrics) RecordBaseFee(*big.Int)            {}
func (*NoopTxMetrics) RecordBlobBaseFee(*big.Int)        {}
func (*NoopTxMetrics) RecordTipCap(*big.Int)             {}
func (*NoopTxMetrics) RPCError()                         {}
