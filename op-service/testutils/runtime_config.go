package testutils

import "github.com/zircuit-labs/l2-geth-public/common"

type MockRuntimeConfig struct {
	P2PSeqAddress common.Address
}

func (m *MockRuntimeConfig) P2PSequencerAddress() common.Address {
	return m.P2PSeqAddress
}
