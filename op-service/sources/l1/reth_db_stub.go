//go:build !rethdb

package l1

import (
	"github.com/zircuit-labs/l2-geth-public/log"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/caching"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/l1/client"
)

const buildRethdb = false

func newRecProviderFromConfig(client client.RPC, log log.Logger, metrics caching.Metrics, config *L1EthClientConfig) *CachingReceiptsProvider {
	return newRPCRecProviderFromConfig(client, log, metrics, config)
}
