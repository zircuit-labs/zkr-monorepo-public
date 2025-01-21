package testutils

import (
	"math/big"
	"math/rand"

	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/core/types"

	l1common "github.com/ethereum/go-ethereum/common"
	l1types "github.com/ethereum/go-ethereum/core/types"
)

// Returns a DepositEvent customized on the basis of the id parameter.
func GenerateDeposit(sourceHash common.Hash, rng *rand.Rand) *types.DepositTx {
	dataLen := rng.Int63n(10_000)
	data := make([]byte, dataLen)
	rng.Read(data)

	var to *common.Address
	if rng.Intn(2) == 0 {
		x := RandomAddress(rng)
		to = &x
	}
	var mint *big.Int
	if rng.Intn(2) == 0 {
		mint = RandomETH(rng, 200)
	}

	dep := &types.DepositTx{
		SourceHash:          sourceHash,
		From:                RandomAddress(rng),
		To:                  to,
		Value:               RandomETH(rng, 200),
		Gas:                 uint64(rng.Int63n(10 * 1e6)), // 10 M gas max
		Data:                data,
		Mint:                mint,
		IsSystemTransaction: false,
	}
	return dep
}

// Generates an EVM log entry with the given topics and data.
func GenerateLog(addr l1common.Address, topics []l1common.Hash, data []byte) *l1types.Log {
	return &l1types.Log{
		Address: addr,
		Topics:  topics,
		Data:    data,
		Removed: false,

		// ignored (zeroed):
		BlockNumber: 0,
		TxHash:      l1common.Hash{},
		TxIndex:     0,
		BlockHash:   l1common.Hash{},
		Index:       0,
	}
}
