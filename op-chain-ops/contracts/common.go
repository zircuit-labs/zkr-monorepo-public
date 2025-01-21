package contracts

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/zircuit-labs/l2-geth-public/common"
)

// parseAddress will parse a [common.Address] from a [cli.Context] and return
// an error if the configured address is not correct.
func parseAddress(ctx *cli.Context, name string) (common.Address, error) {
	value := ctx.String(name)
	if value == "" {
		return common.Address{}, nil
	}
	if !common.IsHexAddress(value) {
		return common.Address{}, fmt.Errorf("invalid address: %s", value)
	}
	return common.HexToAddress(value), nil
}
