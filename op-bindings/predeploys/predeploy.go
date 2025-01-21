package predeploys

import (
	"github.com/zircuit-labs/l2-geth-public/common"
)

type DeployConfig interface {
	GovernanceEnabled() bool
	CanyonTime(genesisTime uint64) *uint64
}

type Predeploy struct {
	Address       common.Address
	ProxyDisabled bool
	Enabled       func(config DeployConfig) bool
}
