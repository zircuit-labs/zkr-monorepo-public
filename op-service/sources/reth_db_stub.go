//go:build !rethdb

package sources

import (
	"github.com/zircuit-labs/l2-geth-public/log"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/client"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/caching"
)

const buildRethdb = false

func newRecProviderFromConfig(client client.RPC, log log.Logger, metrics caching.Metrics, config *EthClientConfig) *CachingReceiptsProvider {
	return newRPCRecProviderFromConfig(client, log, metrics, config)
}
