package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup"
	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup/derive"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
)

type BuildTransactionsRejectedEvent struct {
	Attributes           *derive.AttributesWithParent
	RejectedTransactions []common.Hash
}

func (ev BuildTransactionsRejectedEvent) String() string {
	return "build-transactionsRejected"
}

type BuildStartEvent struct {
	Attributes   *derive.AttributesWithParent
	IsSequencing bool
}

func (ev BuildStartEvent) String() string {
	return "build-start"
}

func (eq *EngDeriver) onBuildStart(ev BuildStartEvent) {
	ctx, cancel := context.WithTimeout(eq.ctx, buildStartTimeout)
	defer cancel()

	if ev.Attributes.DerivedFrom != (eth.L1BlockRef{}) &&
		eq.ec.PendingSafeL2Head().Hash != ev.Attributes.Parent.Hash {
		// Warn about small reorgs, happens when pending safe head is getting rolled back
		eq.log.Warn("block-attributes derived from L1 do not build on pending safe head, likely reorg",
			"pending_safe", eq.ec.PendingSafeL2Head(), "attributes_parent", ev.Attributes.Parent)
	}

	fcEvent := ForkchoiceUpdateEvent{
		UnsafeL2Head:    ev.Attributes.Parent,
		SafeL2Head:      eq.ec.safeHead,
		FinalizedL2Head: eq.ec.finalizedHead,
	}
	fc := eth.ForkchoiceState{
		HeadBlockHash:      fcEvent.UnsafeL2Head.Hash,
		SafeBlockHash:      fcEvent.SafeL2Head.Hash,
		FinalizedBlockHash: fcEvent.FinalizedL2Head.Hash,
		CheckTransactions:  ev.IsSequencing,
	}
	buildStartTime := time.Now()
	id, errTyp, rejectedTxs, err := startPayload(ctx, eq.ec.engine, fc, ev.Attributes.Attributes)
	if err != nil {
		switch errTyp {
		case BlockInsertTemporaryErr:
			// RPC errors are recoverable, we can retry the buffered payload attributes later.
			eq.emitter.Emit(rollup.EngineTemporaryErrorEvent{Err: fmt.Errorf("temporarily cannot insert new safe block: %w", err)})
			return
		case BlockInsertPrestateErr:
			eq.emitter.Emit(rollup.ResetEvent{Err: fmt.Errorf("need reset to resolve pre-state problem: %w", err)})
			return
		case BlockInsertPayloadErr:
			eq.emitter.Emit(BuildInvalidEvent{Attributes: ev.Attributes, Err: err})
			return
		case BlockInsertTransactionsRejected:
			eq.emitter.Emit(BuildTransactionsRejectedEvent{Attributes: ev.Attributes, RejectedTransactions: rejectedTxs})
			return
		default:
			eq.emitter.Emit(rollup.CriticalErrorEvent{Err: fmt.Errorf("unknown error type %d: %w", errTyp, err)})
			return
		}
	}
	eq.emitter.Emit(fcEvent)

	eq.emitter.Emit(BuildStartedEvent{
		Info:         eth.PayloadInfo{ID: id, Timestamp: uint64(ev.Attributes.Attributes.Timestamp)},
		Attributes:   ev.Attributes,
		BuildStarted: buildStartTime,
		IsLastInSpan: ev.Attributes.IsLastInSpan,
		DerivedFrom:  ev.Attributes.DerivedFrom,
		Parent:       ev.Attributes.Parent,
	})
}
