package rpc

import (
	"context"
	"fmt"

	"github.com/zircuit-labs/l2-geth-public/log"
	"github.com/zircuit-labs/l2-geth-public/rpc"
	oplog "github.com/zircuit-labs/zkr-monorepo-public/op-service/log"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/metrics"
)

func ToGethAdminAPI(api *CommonAdminAPI) rpc.API {
	return rpc.API{
		Namespace: "admin",
		Service:   api,
	}
}

type CommonAdminAPI struct {
	M   metrics.RPCMetricer
	log log.Logger
}

func NewCommonAdminAPI(m metrics.RPCMetricer, log log.Logger) *CommonAdminAPI {
	return &CommonAdminAPI{
		M:   m,
		log: log,
	}
}

func (n *CommonAdminAPI) SetLogLevel(ctx context.Context, lvlStr string) error {
	recordDur := n.M.RecordRPCServerRequest("admin_setLogLevel")
	defer recordDur()

	lvl, err := oplog.LevelFromString(lvlStr)
	if err != nil {
		return err
	}

	h := n.log.Handler()
	// We set the log level, and do not wrap the handler with an additional filter handler,
	// as the underlying handler would otherwise also still filter with the previous log level.
	lvlSetter, ok := h.(oplog.LvlSetter)
	if !ok {
		return fmt.Errorf("log handler type %T cannot change log level", h)
	}
	lvlSetter.SetLogLevel(lvl)
	return nil
}
