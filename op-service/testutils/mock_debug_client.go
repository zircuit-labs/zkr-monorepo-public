package testutils

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/zircuit-labs/l2-geth-public/common"
)

type MockDebugClient struct {
	mock.Mock
}

func (m *MockDebugClient) ExpectNodeByHash(hash common.Hash, res []byte, err error) {
	m.Mock.On("NodeByHash", hash).Once().Return(res, err)
}

func (m *MockDebugClient) NodeByHash(ctx context.Context, hash common.Hash) ([]byte, error) {
	out := m.Mock.Called(hash)
	return out.Get(0).([]byte), out.Error(1)
}

func (m *MockDebugClient) ExpectCodeByHash(hash common.Hash, res []byte, err error) {
	m.Mock.On("CodeByHash", hash).Once().Return(res, err)
}

func (m *MockDebugClient) CodeByHash(ctx context.Context, hash common.Hash) ([]byte, error) {
	out := m.Mock.Called(hash)
	return out.Get(0).([]byte), out.Error(1)
}
