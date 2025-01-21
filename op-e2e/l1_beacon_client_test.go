package op_e2e

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/stretchr/testify/require"
	"github.com/zircuit-labs/l2-geth-public/log"
	"github.com/zircuit-labs/zkr-monorepo-public/op-e2e/e2eutils/fakebeacon"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/l1"
	l1client "github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/l1/client"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/testlog"
)

func TestGetVersion(t *testing.T) {
	InitParallel(t)

	l := testlog.Logger(t, log.LevelInfo)

	beaconApi := fakebeacon.NewBeacon(l, t.TempDir(), uint64(0), uint64(0))
	t.Cleanup(func() {
		_ = beaconApi.Close()
	})
	require.NoError(t, beaconApi.Start("127.0.0.1:0"))

	beaconCfg := l1.L1BeaconClientConfig{FetchAllSidecars: false}
	cl := l1.NewL1BeaconClient(l1.NewBeaconHTTPClient(l1client.NewBasicHTTPClient(beaconApi.BeaconAddr(), l)), beaconCfg)

	version, err := cl.GetVersion(context.Background())
	require.NoError(t, err)
	require.Equal(t, "fakebeacon 1.2.3", version)
}

func Test404NotFound(t *testing.T) {
	InitParallel(t)

	l := testlog.Logger(t, log.LevelInfo)

	beaconApi := fakebeacon.NewBeacon(l, t.TempDir(), uint64(0), uint64(12))
	t.Cleanup(func() {
		_ = beaconApi.Close()
	})
	require.NoError(t, beaconApi.Start("127.0.0.1:0"))

	beaconCfg := l1.L1BeaconClientConfig{FetchAllSidecars: false}
	cl := l1.NewL1BeaconClient(l1.NewBeaconHTTPClient(l1client.NewBasicHTTPClient(beaconApi.BeaconAddr(), l)), beaconCfg)

	hashes := []eth.IndexedBlobHash{{Index: 1}}
	_, err := cl.GetBlobs(context.Background(), eth.L1BlockRef{Number: 10, Time: 120}, hashes)
	require.ErrorIs(t, err, ethereum.NotFound)
}
