package geth

import (
	"fmt"
	"math/big"

	"github.com/zircuit-labs/l2-geth-public/rollup/tracing"

	"github.com/zircuit-labs/l2-geth-public/cmd/utils"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/core"
	"github.com/zircuit-labs/l2-geth-public/core/txpool/blobpool"
	"github.com/zircuit-labs/l2-geth-public/eth"
	"github.com/zircuit-labs/l2-geth-public/eth/catalyst"
	"github.com/zircuit-labs/l2-geth-public/eth/ethconfig"
	"github.com/zircuit-labs/l2-geth-public/eth/tracers"
	"github.com/zircuit-labs/l2-geth-public/log"
	"github.com/zircuit-labs/l2-geth-public/miner"
	"github.com/zircuit-labs/l2-geth-public/node"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/clock"

	// Force-load the tracer engines to trigger registration
	_ "github.com/zircuit-labs/l2-geth-public/eth/tracers/js"
	_ "github.com/zircuit-labs/l2-geth-public/eth/tracers/native"
)

func InitL1(chainID uint64, blockTime uint64, finalizedDistance uint64, genesis *core.Genesis, c clock.Clock, blobPoolDir string, beaconSrv Beacon, opts ...GethOption) (*node.Node, *eth.Ethereum, error) {
	ethConfig := &ethconfig.Config{
		NetworkId: chainID,
		Genesis:   genesis,
		BlobPool: blobpool.Config{
			Datadir:   blobPoolDir,
			Datacap:   blobpool.DefaultConfig.Datacap,
			PriceBump: blobpool.DefaultConfig.PriceBump,
		},
	}
	nodeConfig := &node.Config{
		Name:        "l1-geth",
		HTTPHost:    "127.0.0.1",
		HTTPPort:    0,
		WSHost:      "127.0.0.1",
		WSPort:      0,
		WSModules:   []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal", "engine"},
		HTTPModules: []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal", "engine"},
	}

	l1Node, l1Eth, err := createGethNode(false, nodeConfig, ethConfig, opts...)
	if err != nil {
		return nil, nil, err
	}
	// Activate merge
	l1Eth.Merger().FinalizePoS()
	api, err := catalyst.NewConsensusAPI(l1Eth)
	if err != nil {
		return nil, nil, err
	}

	// Instead of running a whole beacon node, we run this fake-proof-of-stake sidecar that sequences L1 blocks using the Engine API.
	l1Node.RegisterLifecycle(&fakePoS{
		clock:             c,
		eth:               l1Eth,
		log:               log.Root(), // geth logger is global anyway. Would be nice to replace with a local logger though.
		blockTime:         blockTime,
		finalizedDistance: finalizedDistance,
		safeDistance:      4,
		engineAPI:         api,
		beacon:            beaconSrv,
	})

	return l1Node, l1Eth, nil
}

func defaultNodeConfig(name string, jwtPath string) *node.Config {
	return &node.Config{
		Name:        name,
		WSHost:      "127.0.0.1",
		WSPort:      0,
		AuthAddr:    "127.0.0.1",
		AuthPort:    0,
		HTTPHost:    "127.0.0.1",
		HTTPPort:    0,
		WSModules:   []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal", "engine"},
		HTTPModules: []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal", "engine"},
		JWTSecret:   jwtPath,
	}
}

type GethOption func(ethCfg *ethconfig.Config, nodeCfg *node.Config) error

// InitL2 inits a L2 geth node.
func InitL2(name string, l2ChainID *big.Int, genesis *core.Genesis, jwtPath string, opts ...GethOption) (*node.Node, *eth.Ethereum, error) {
	ethConfig := &ethconfig.Config{
		NetworkId: l2ChainID.Uint64(),
		Genesis:   genesis,
		Miner: miner.Config{
			Etherbase:         common.Address{},
			ExtraData:         nil,
			GasFloor:          0,
			GasCeil:           0,
			GasPrice:          nil,
			Recommit:          0,
			NewPayloadTimeout: 0,
		},
	}
	nodeConfig := defaultNodeConfig(fmt.Sprintf("l2-geth-%v", name), jwtPath)
	return createGethNode(true, nodeConfig, ethConfig, opts...)
}

// createGethNode creates an in-memory geth node based on the configuration.
// The private keys are added to the keystore and are unlocked.
// If the node is l2, catalyst is enabled.
// The node should be started and then closed when done.
func createGethNode(l2 bool, nodeCfg *node.Config, ethCfg *ethconfig.Config, opts ...GethOption) (*node.Node, *eth.Ethereum, error) {
	for i, opt := range opts {
		if err := opt(ethCfg, nodeCfg); err != nil {
			return nil, nil, fmt.Errorf("failed to apply geth option %d: %w", i, err)
		}
	}
	ethCfg.NoPruning = true // force everything to be an archive node
	n, err := node.New(nodeCfg)
	if err != nil {
		n.Close()
		return nil, nil, err
	}

	backend, err := eth.New(n, ethCfg)
	if err != nil {
		n.Close()
		return nil, nil, err

	}

	// PR 25459 changed this to only default in CLI, but not in default programmatic RPC selection.
	// PR 25642 fixed it for the mobile version only...
	utils.RegisterFilterAPI(n, backend.APIBackend, ethCfg)

	scrollTracerWrapper := tracing.NewTracerWrapper(backend.BlockChain())
	n.RegisterAPIs(tracers.APIs(backend.APIBackend, scrollTracerWrapper))

	// Enable catalyst if l2
	if l2 {
		if err := catalyst.Register(n, backend); err != nil {
			n.Close()
			return nil, nil, err
		}
	}
	return n, backend, nil
}
