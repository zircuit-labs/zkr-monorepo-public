package sequencing

import (
	"context"
	"errors"
	"time"

	"github.com/zircuit-labs/l2-geth-public/common"

	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup/event"
)

var ErrSequencerNotEnabled = errors.New("sequencer is not enabled")

type DisabledSequencer struct{}

var _ SequencerIface = DisabledSequencer{}

func (ds DisabledSequencer) OnEvent(ev event.Event) bool {
	return false
}

func (ds DisabledSequencer) NextAction() (t time.Time, ok bool) {
	return time.Time{}, false
}

func (ds DisabledSequencer) Active() bool {
	return false
}

func (ds DisabledSequencer) Init(ctx context.Context, active bool) error {
	return ErrSequencerNotEnabled
}

func (ds DisabledSequencer) Start(ctx context.Context, head common.Hash) error {
	return ErrSequencerNotEnabled
}

func (ds DisabledSequencer) Stop(ctx context.Context) (hash common.Hash, err error) {
	return common.Hash{}, ErrSequencerNotEnabled
}

func (ds DisabledSequencer) SetMaxSafeLag(ctx context.Context, v uint64) error {
	return ErrSequencerNotEnabled
}

func (ds DisabledSequencer) OverrideLeader(ctx context.Context) error {
	return ErrSequencerNotEnabled
}

func (ds DisabledSequencer) Close() {}
