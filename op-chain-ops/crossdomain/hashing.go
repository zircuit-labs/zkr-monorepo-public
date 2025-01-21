package crossdomain

import (
	"math/big"

	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/crypto"
)

// HashCrossDomainMessageV0 computes the pre bedrock cross domain messaging
// hashing scheme.
func HashCrossDomainMessageV0(
	target common.Address,
	sender common.Address,
	data []byte,
	nonce *big.Int,
) (common.Hash, error) {
	encoded, err := EncodeCrossDomainMessageV0(target, sender, data, nonce)
	if err != nil {
		return common.Hash{}, err
	}
	hash := crypto.Keccak256(encoded)
	return common.BytesToHash(hash), nil
}

// HashCrossDomainMessageV1 computes the first post bedrock cross domain
// messaging hashing scheme.
func HashCrossDomainMessageV1(
	nonce *big.Int,
	sender common.Address,
	target common.Address,
	value *big.Int,
	gasLimit *big.Int,
	data []byte,
) (common.Hash, error) {
	encoded, err := EncodeCrossDomainMessageV1(nonce, sender, target, value, gasLimit, data)
	if err != nil {
		return common.Hash{}, err
	}
	hash := crypto.Keccak256(encoded)
	return common.BytesToHash(hash), nil
}
