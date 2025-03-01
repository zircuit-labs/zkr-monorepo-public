package crossdomain

import (
	"github.com/zircuit-labs/l2-geth-public/accounts/abi"
	"github.com/zircuit-labs/l2-geth-public/common"
)

var (
	// Standard ABI types
	Uint256Type, _ = abi.NewType("uint256", "", nil)
	BytesType, _   = abi.NewType("bytes", "", nil)
	AddressType, _ = abi.NewType("address", "", nil)
	Bytes32Type, _ = abi.NewType("bytes32", "", nil)
)

// WithdrawalMessage represents a Withdrawal. The Withdrawal
// and LegacyWithdrawal types must implement this interface.
type WithdrawalMessage interface {
	Encode() ([]byte, error)
	Decode([]byte) error
	Hash() (common.Hash, error)
	StorageSlot() (common.Hash, error)
}
