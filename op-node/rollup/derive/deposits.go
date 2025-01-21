package derive

import (
	"fmt"

	"github.com/hashicorp/go-multierror"

	l1common "github.com/ethereum/go-ethereum/common"
	l1types "github.com/ethereum/go-ethereum/core/types"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/common/hexutil"
	"github.com/zircuit-labs/l2-geth-public/core/types"
)

// UserDeposits transforms the L2 block-height and L1 receipts into the transaction inputs for a full L2 block
func UserDeposits(receipts []*l1types.Receipt, depositContractAddr common.Address) ([]*types.DepositTx, error) {
	var out []*types.DepositTx
	var result error
	for i, rec := range receipts {
		if rec.Status != types.ReceiptStatusSuccessful {
			continue
		}
		for j, log := range rec.Logs {
			if log.Address == l1common.Address(depositContractAddr) && len(log.Topics) > 0 && log.Topics[0] == l1common.Hash(DepositEventABIHash) {
				dep, err := UnmarshalDepositLogEvent(log)
				if err != nil {
					result = multierror.Append(result, fmt.Errorf("malformatted L1 deposit log in receipt %d, log %d: %w", i, j, err))
				} else {
					out = append(out, dep)
				}
			}
		}
	}
	return out, result
}

func DeriveDeposits(receipts []*l1types.Receipt, depositContractAddr common.Address, exclusions *types.Bitmap) ([]hexutil.Bytes, error) {
	var result error
	userDeposits, err := UserDeposits(receipts, depositContractAddr)
	if err != nil {
		result = multierror.Append(result, err)
	}
	encodedTxs := make([]hexutil.Bytes, 0, len(userDeposits))
	for i, tx := range userDeposits {
		// if applicable, skip any excluded deposits, +1 since the l1 block info deposit
		// is not part of the user deposits but will always be included as first deposit
		if exclusions != nil && exclusions.Test(i+1) {
			continue
		}
		opaqueTx, err := types.NewTx(tx).MarshalBinary()
		if err != nil {
			result = multierror.Append(result, fmt.Errorf("failed to encode user tx %d", i))
		} else {
			encodedTxs = append(encodedTxs, opaqueTx)
		}
	}
	return encodedTxs, result
}
