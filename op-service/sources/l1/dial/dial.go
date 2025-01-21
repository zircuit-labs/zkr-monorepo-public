package dial

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zircuit-labs/l2-geth-public/log"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/retry"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/l1/client"
)

// DefaultDialTimeout is a default timeout for dialing a client.
const (
	DefaultDialTimeout = 1 * time.Minute
	defaultRetryCount  = 30
	defaultRetryTime   = 2 * time.Second
)

// DialEthClientWithTimeout attempts to dial the L1 provider using the provided
// URL. If the dial doesn't complete within defaultDialTimeout seconds, this
// method will return an error.
func DialEthClientWithTimeout(ctx context.Context, timeout time.Duration, log log.Logger, url string) (*ethclient.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	c, err := dialRPCClientWithBackoff(ctx, log, url)
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(c), nil
}

// DialRPCClientWithTimeout attempts to dial the RPC provider using the provided URL.
// If the dial doesn't complete within timeout seconds, this method will return an error.
func DialRPCClientWithTimeout(ctx context.Context, timeout time.Duration, log log.Logger, url string) (*rpc.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return dialRPCClientWithBackoff(ctx, log, url)
}

// Dials a JSON-RPC endpoint repeatedly, with a backoff, until a client connection is established. Auth is optional.
func dialRPCClientWithBackoff(ctx context.Context, log log.Logger, addr string) (*rpc.Client, error) {
	bOff := retry.Fixed(defaultRetryTime)
	return retry.Do(ctx, defaultRetryCount, bOff, func() (*rpc.Client, error) {
		return dialRPCClient(ctx, log, addr)
	})
}

// Dials a JSON-RPC endpoint once.
func dialRPCClient(ctx context.Context, log log.Logger, addr string) (*rpc.Client, error) {
	if !client.IsURLAvailable(ctx, addr) {
		log.Warn("failed to dial address, but may connect later", "addr", addr)
		return nil, fmt.Errorf("address unavailable (%s)", addr)
	}
	client, err := rpc.DialOptions(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial address (%s): %w", addr, err)
	}
	return client, nil
}
