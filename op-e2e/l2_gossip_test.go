package op_e2e

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/zkr-monorepo-public/op-e2e/e2eutils/geth"
)

func TestTxGossip(t *testing.T) {
	// InitParallel(t) // TODO:
	cfg := DefaultSystemConfig(t)
	gethOpts := []geth.GethOption{
		geth.WithP2P(),
	}
	cfg.GethOptions["sequencer"] = append(cfg.GethOptions["sequencer"], gethOpts...)
	cfg.GethOptions["verifier"] = append(cfg.GethOptions["verifier"], gethOpts...)
	sys, err := cfg.Start(t)
	require.NoError(t, err, "Start system")

	seqClient := sys.Clients["sequencer"]
	verifClient := sys.Clients["verifier"]
	geth.ConnectP2P(t, seqClient, verifClient)

	// Send a transaction to the verifier and it should be gossiped to the sequencer and included in a block.
	SendL2Tx(t, cfg, verifClient, cfg.Secrets.Alice, func(opts *TxOpts) {
		opts.ToAddr = &common.Address{0xaa}
		opts.Value = common.Big1
		opts.VerifyOnClients(seqClient, verifClient)
	})
}
