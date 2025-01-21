package batching

import (
	"context"
	"fmt"
	"io"

	"github.com/zircuit-labs/l2-geth-public/rpc"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/batching/rpcblock"
)

var DefaultBatchSize = 100

type EthRpc interface {
	CallContext(ctx context.Context, out interface{}, method string, args ...interface{}) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
}

type MultiCaller struct {
	rpc       EthRpc
	batchSize int
}

func NewMultiCaller(rpc EthRpc, batchSize int) *MultiCaller {
	return &MultiCaller{
		rpc:       rpc,
		batchSize: batchSize,
	}
}

func (m *MultiCaller) BatchSize() int {
	return m.batchSize
}

func (m *MultiCaller) SingleCall(ctx context.Context, block rpcblock.Block, call Call) (*CallResult, error) {
	results, err := m.Call(ctx, block, call)
	if err != nil {
		return nil, err
	}
	return results[0], nil
}

func (m *MultiCaller) Call(ctx context.Context, block rpcblock.Block, calls ...Call) ([]*CallResult, error) {
	keys := make([]BatchElementCreator, len(calls))
	for i := 0; i < len(calls); i++ {
		creator, err := calls[i].ToBatchElemCreator()
		if err != nil {
			return nil, err
		}
		keys[i] = creator
	}
	fetcher := NewIterativeBatchCall[BatchElementCreator, any](
		keys,
		func(key BatchElementCreator) (any, rpc.BatchElem) {
			return key(block)
		},
		m.rpc.BatchCallContext,
		m.rpc.CallContext,
		m.batchSize)
	for {
		if err := fetcher.Fetch(ctx); err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("failed to fetch batch: %w", err)
		}
	}
	results, err := fetcher.Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get batch call results: %w", err)
	}

	callResults := make([]*CallResult, len(results))
	for i, result := range results {
		call := calls[i]
		out, err := call.HandleResult(result)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack result: %w", err)
		}
		callResults[i] = out
	}
	return callResults, nil
}
