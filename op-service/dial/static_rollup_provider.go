package dial

import (
	"context"

	"github.com/zircuit-labs/l2-geth-public/log"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources"
)

// RollupProvider is an interface for providing a RollupClient
// It manages the lifecycle of the RollupClient for callers
type RollupProvider interface {
	// RollupClient(ctx) returns the underlying sources.RollupClient pointing to the L2 rollup consensus node
	RollupClient(ctx context.Context) (RollupClientInterface, error)
	// Close() closes the underlying client or clients
	Close()
}

// StaticL2RollupProvider is a RollupProvider that always returns the same static RollupClient
// It is meant for scenarios where a single, unchanging L2 rollup node is used
type StaticL2RollupProvider struct {
	rollupClient *sources.RollupClient
}

func NewStaticL2RollupProvider(ctx context.Context, log log.Logger, rollupClientUrl string) (*StaticL2RollupProvider, error) {
	rollupClient, err := DialRollupClientWithTimeout(ctx, DefaultDialTimeout, log, rollupClientUrl)
	if err != nil {
		return nil, err
	}
	return &StaticL2RollupProvider{
		rollupClient: rollupClient,
	}, nil
}

// The NewStaticL2EndpointProviderFromExistingRollup constructor is used in e2e testing
func NewStaticL2RollupProviderFromExistingRollup(rollupCl *sources.RollupClient) (*StaticL2RollupProvider, error) {
	return &StaticL2RollupProvider{
		rollupClient: rollupCl,
	}, nil
}

func (p *StaticL2RollupProvider) RollupClient(context.Context) (RollupClientInterface, error) {
	return p.rollupClient, nil
}

func (p *StaticL2RollupProvider) Close() {
	if p.rollupClient != nil {
		p.rollupClient.Close()
	}
}
