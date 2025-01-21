package derive

import "github.com/zircuit-labs/zkr-monorepo-public/op-service/testutils"

var _ L1Fetcher = (*testutils.MockL1Reader)(nil)

var _ Metrics = (*testutils.TestDerivationMetrics)(nil)
