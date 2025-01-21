package rpc

import (
	"context"

	"github.com/zircuit-labs/l2-geth-public/log"
	gethrpc "github.com/zircuit-labs/l2-geth-public/rpc"

	"github.com/zircuit-labs/zkr-monorepo-public/op-service/metrics"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/rpc"
)

type BatcherDriver interface {
	StartBatchSubmitting() error
	StopBatchSubmitting(ctx context.Context) error
}

type adminAPI struct {
	*rpc.CommonAdminAPI
	b BatcherDriver
}

func NewAdminAPI(dr BatcherDriver, m metrics.RPCMetricer, log log.Logger) *adminAPI {
	return &adminAPI{
		CommonAdminAPI: rpc.NewCommonAdminAPI(m, log),
		b:              dr,
	}
}

func GetAdminAPI(api *adminAPI) gethrpc.API {
	return gethrpc.API{
		Namespace: "admin",
		Service:   api,
	}
}

func (a *adminAPI) StartBatcher(_ context.Context) error {
	return a.b.StartBatchSubmitting()
}

func (a *adminAPI) StopBatcher(ctx context.Context) error {
	return a.b.StopBatchSubmitting(ctx)
}
