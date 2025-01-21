package batching

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/batching/rpcblock"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/batching/test"
)

func TestGetBalance(t *testing.T) {
	addr := common.Address{0xab, 0xcd}
	expectedBalance := big.NewInt(248924)

	stub := test.NewRpcStub(t)
	stub.AddExpectedCall(test.NewGetBalanceCall(addr, rpcblock.Latest, expectedBalance))

	caller := NewMultiCaller(stub, DefaultBatchSize)
	result, err := caller.SingleCall(context.Background(), rpcblock.Latest, NewBalanceCall(addr))
	require.NoError(t, err)
	require.Equal(t, expectedBalance, result.GetBigInt(0))
}
