package batcher

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/zircuit-labs/zkr-monorepo-public/op-batcher/flags"
	opservice "github.com/zircuit-labs/zkr-monorepo-public/op-service"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/cliapp"
	oplog "github.com/zircuit-labs/zkr-monorepo-public/op-service/log"
)

// Main is the entrypoint into the Batch Submitter.
// This method returns a cliapp.LifecycleAction, to create an op-service CLI-lifecycle-managed batch-submitter with.
func Main(version string) cliapp.LifecycleAction {
	return func(cliCtx *cli.Context, closeApp context.CancelCauseFunc) (cliapp.Lifecycle, error) {
		if err := flags.CheckRequired(cliCtx); err != nil {
			return nil, err
		}
		cfg := NewConfig(cliCtx)
		if err := cfg.Check(); err != nil {
			return nil, fmt.Errorf("invalid CLI flags: %w", err)
		}

		l := oplog.NewLogger(oplog.AppOut(cliCtx), cfg.LogConfig)
		oplog.SetGlobalLogHandler(l.Handler())
		opservice.ValidateEnvVars(flags.EnvVarPrefix, flags.Flags, l)

		l.Info("Initializing Batch Submitter")
		return BatcherServiceFromCLIConfig(cliCtx.Context, version, cfg, l)
	}
}
