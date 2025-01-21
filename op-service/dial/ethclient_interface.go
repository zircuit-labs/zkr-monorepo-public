package dial

import (
	"context"
	"math/big"

	"github.com/zircuit-labs/l2-geth-public/ethclient"
)

// EthClientInterface is an interface for providing an ethclient.Client
// It does not describe all of the functions an ethclient.Client has, only the ones used by callers of the L2 Providers
type EthClientInterface interface {
	BlockByNumberEx(ctx context.Context, number *big.Int) (*ethclient.BlockEx, error)

	Close()
}
