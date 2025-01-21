package derive

import (
	"bytes"
	"math/big"
	"math/rand"
	"testing"

	"github.com/zircuit-labs/l2-geth-public/rlp"

	"github.com/stretchr/testify/require"
)

func TestSingularBatchForBatchInterface(t *testing.T) {
	rng := rand.New(rand.NewSource(0x543331))
	chainID := big.NewInt(rng.Int63n(1000))
	txCount := 1 + rng.Intn(8)

	singularBatch := RandomSingularBatch(rng, txCount, chainID)

	require.Equal(t, SingularBatchType, singularBatch.GetBatchType())
	require.Equal(t, singularBatch.Timestamp, singularBatch.GetTimestamp())
	require.Equal(t, singularBatch.EpochNum, singularBatch.GetEpochNum())
}

func TestSingularBatchRLP(t *testing.T) {
	rng := rand.New(rand.NewSource(0x543331))
	chainID := big.NewInt(rng.Int63n(1000))
	txCount := 1 + rng.Intn(8)

	singularBatch := RandomSingularBatch(rng, txCount, chainID)

	var buf bytes.Buffer
	err := rlp.Encode(&buf, singularBatch)
	require.NoError(t, err)

	var decodedSingularBatch SingularBatch
	err = rlp.Decode(&buf, &decodedSingularBatch)
	require.NoError(t, err)

	require.Equal(t, *singularBatch, decodedSingularBatch)

	// also ensure that DepositExclusions is optional field
	singularBatch.DepositExclusions = nil

	err = rlp.Encode(&buf, singularBatch)
	require.NoError(t, err)

	err = rlp.Decode(&buf, &decodedSingularBatch)
	require.NoError(t, err)

	require.Equal(t, *singularBatch, decodedSingularBatch)
}
