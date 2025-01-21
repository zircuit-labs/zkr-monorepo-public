package testutils

import (
	"github.com/stretchr/testify/mock"
	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup/event"
)

type MockEmitter struct {
	mock.Mock
}

func (m *MockEmitter) Emit(ev event.Event) {
	m.Mock.MethodCalled("Emit", ev)
}

func (m *MockEmitter) ExpectOnce(expected event.Event) {
	m.Mock.On("Emit", expected).Once()
}

func (m *MockEmitter) ExpectMaybeRun(fn func(ev event.Event)) {
	m.Mock.On("Emit", mock.Anything).Maybe().Run(func(args mock.Arguments) {
		fn(args.Get(0).(event.Event))
	})
}

func (m *MockEmitter) ExpectOnceType(typ string) {
	m.Mock.On("Emit", mock.AnythingOfType(typ)).Once()
}

func (m *MockEmitter) ExpectOnceRun(fn func(ev event.Event)) {
	m.Mock.On("Emit", mock.Anything).Once().Run(func(args mock.Arguments) {
		fn(args.Get(0).(event.Event))
	})
}

func (m *MockEmitter) AssertExpectations(t mock.TestingT) {
	m.Mock.AssertExpectations(t)
}

var _ event.Emitter = (*MockEmitter)(nil)
