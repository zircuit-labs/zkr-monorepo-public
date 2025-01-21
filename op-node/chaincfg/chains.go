package chaincfg

import (
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/superchain-registry/superchain"

	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup"
)

// mostly empty config for p2p tests
var P2PTestConfig *rollup.Config

func init() {
	P2PTestConfig = &rollup.Config{
		BlockTime: 2,
		L2ChainID: big.NewInt(901),
	}
}

var L2ChainIDToNetworkDisplayName = func() map[string]string {
	out := make(map[string]string)
	for _, netCfg := range superchain.OPChains {
		out[fmt.Sprintf("%d", netCfg.ChainID)] = netCfg.Name
	}
	return out
}()

// AvailableNetworks returns the selection of network configurations that is available by default.
func AvailableNetworks() []string {
	var networks []string
	return networks
}

func GetRollupConfig(name string) (*rollup.Config, error) {
	return nil, fmt.Errorf("Network configuration is currently not supported")
}
