package driver

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zircuit-labs/l2-geth-public/log"

	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup/event"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/testlog"
)

func TestStepSchedulingDeriver(t *testing.T) {
	logger := testlog.Logger(t, log.LevelError)
	var queued []event.Event
	emitter := event.EmitterFunc(func(ev event.Event) {
		queued = append(queued, ev)
	})
	sched := NewStepSchedulingDeriver(logger)
	sched.AttachEmitter(emitter)
	require.Len(t, sched.NextStep(), 0, "start empty")
	sched.OnEvent(StepReqEvent{})
	require.Len(t, sched.NextStep(), 1, "take request")
	sched.OnEvent(StepReqEvent{})
	require.Len(t, sched.NextStep(), 1, "ignore duplicate request")
	require.Empty(t, queued, "only scheduled so far, no step attempts yet")
	<-sched.NextStep()
	sched.OnEvent(StepAttemptEvent{})
	require.Equal(t, []event.Event{StepEvent{}}, queued, "got step event")
	require.Nil(t, sched.NextDelayedStep(), "no delayed steps yet")
	sched.OnEvent(StepReqEvent{})
	require.NotNil(t, sched.NextDelayedStep(), "2nd attempt before backoff reset causes delayed step to be scheduled")
	sched.OnEvent(StepReqEvent{})
	require.NotNil(t, sched.NextDelayedStep(), "can continue to request attempts")

	sched.OnEvent(StepReqEvent{})
	require.Len(t, sched.NextStep(), 0, "no step requests accepted without delay if backoff is counting")

	sched.OnEvent(StepReqEvent{ResetBackoff: true})
	require.Len(t, sched.NextStep(), 1, "request accepted if backoff is reset")
	<-sched.NextStep()

	sched.OnEvent(StepReqEvent{})
	require.Len(t, sched.NextStep(), 1, "no backoff, no attempt has been made yet")
	<-sched.NextStep()
	sched.OnEvent(StepAttemptEvent{})
	sched.OnEvent(StepReqEvent{})
	require.Len(t, sched.NextStep(), 0, "backoff again")

	sched.OnEvent(ResetStepBackoffEvent{})
	sched.OnEvent(StepReqEvent{})
	require.Len(t, sched.NextStep(), 1, "reset backoff accepted, was able to schedule non-delayed step")
}
