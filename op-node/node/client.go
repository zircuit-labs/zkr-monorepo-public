package node

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/client"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/l1"
	l1client "github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/l1/client"

	"github.com/zircuit-labs/l2-geth-public/log"
	gn "github.com/zircuit-labs/l2-geth-public/node"
	"github.com/zircuit-labs/l2-geth-public/rpc"
)

type L2EndpointSetup interface {
	// Setup a RPC client to a L2 execution engine to process rollup blocks with.
	Setup(ctx context.Context, log log.Logger, rollupCfg *rollup.Config) (cl client.RPC, rpcCfg *sources.EngineClientConfig, err error)
	Check() error
}

type L1EndpointSetup interface {
	// Setup a RPC client to a L1 node to pull rollup input-data from.
	// The results of the RPC client may be trusted for faster processing, or strictly validated.
	// The kind of the RPC may be non-basic, to optimize RPC usage.
	Setup(ctx context.Context, log log.Logger, rollupCfg *rollup.Config) (cl l1client.RPC, rpcCfg *l1.L1ClientConfig, err error)
	Check() error
}

type L1BeaconEndpointSetup interface {
	Setup(ctx context.Context, log log.Logger) (cl l1.BeaconClient, fb []l1.BlobSideCarsFetcher, err error)
	// ShouldIgnoreBeaconCheck returns true if the Beacon-node version check should not halt startup.
	ShouldIgnoreBeaconCheck() bool
	ShouldFetchAllSidecars() bool
	Check() error
}

type L2EndpointConfig struct {
	// L2EngineAddr is the address of the L2 Engine JSON-RPC endpoint to use. The engine and eth
	// namespaces must be enabled by the endpoint.
	L2EngineAddr string

	// JWT secrets for L2 Engine API authentication during HTTP or initial Websocket communication.
	// Any value for an IPC connection.
	L2EngineJWTSecret [32]byte
}

var _ L2EndpointSetup = (*L2EndpointConfig)(nil)

func (cfg *L2EndpointConfig) Check() error {
	if cfg.L2EngineAddr == "" {
		return errors.New("empty L2 Engine Address")
	}

	return nil
}

func (cfg *L2EndpointConfig) Setup(ctx context.Context, log log.Logger, rollupCfg *rollup.Config) (client.RPC, *sources.EngineClientConfig, error) {
	if err := cfg.Check(); err != nil {
		return nil, nil, err
	}
	auth := rpc.WithHTTPAuth(gn.NewJWTAuth(cfg.L2EngineJWTSecret))
	opts := []client.RPCOption{
		client.WithGethRPCOptions(auth),
		client.WithDialBackoff(10),
	}
	l2Node, err := client.NewRPC(ctx, log, cfg.L2EngineAddr, opts...)
	if err != nil {
		return nil, nil, err
	}

	return l2Node, sources.EngineClientDefaultConfig(rollupCfg), nil
}

// PreparedL2Endpoints enables testing with in-process pre-setup RPC connections to L2 engines
type PreparedL2Endpoints struct {
	Client client.RPC
}

func (p *PreparedL2Endpoints) Check() error {
	if p.Client == nil {
		return errors.New("client cannot be nil")
	}
	return nil
}

var _ L2EndpointSetup = (*PreparedL2Endpoints)(nil)

func (p *PreparedL2Endpoints) Setup(ctx context.Context, log log.Logger, rollupCfg *rollup.Config) (client.RPC, *sources.EngineClientConfig, error) {
	return p.Client, sources.EngineClientDefaultConfig(rollupCfg), nil
}

type L1EndpointConfig struct {
	L1NodeAddr string // Address of L1 User JSON-RPC endpoint to use (eth namespace required)

	// L1TrustRPC: if we trust the L1 RPC we do not have to validate L1 response contents like headers
	// against block hashes, or cached transaction sender addresses.
	// Thus we can sync faster at the risk of the source RPC being wrong.
	L1TrustRPC bool

	// L1RPCKind identifies the RPC provider kind that serves the RPC,
	// to inform the optimal usage of the RPC for transaction receipts fetching.
	L1RPCKind l1.RPCProviderKind

	// RateLimit specifies a self-imposed rate-limit on L1 requests. 0 is no rate-limit.
	RateLimit float64

	// BatchSize specifies the maximum batch-size, which also applies as L1 rate-limit burst amount (if set).
	BatchSize int

	// MaxConcurrency specifies the maximum number of concurrent requests to the L1 RPC.
	MaxConcurrency int

	// HttpPollInterval specifies the interval between polling for the latest L1 block,
	// when the RPC is detected to be an HTTP type.
	// It is recommended to use websockets or IPC for efficient following of the changing block.
	// Setting this to 0 disables polling.
	HttpPollInterval time.Duration
}

var _ L1EndpointSetup = (*L1EndpointConfig)(nil)

func (cfg *L1EndpointConfig) Check() error {
	if cfg.BatchSize < 1 || cfg.BatchSize > 500 {
		return fmt.Errorf("batch size is invalid or unreasonable: %d", cfg.BatchSize)
	}
	if cfg.RateLimit < 0 {
		return fmt.Errorf("rate limit cannot be negative")
	}
	if cfg.MaxConcurrency < 1 {
		return fmt.Errorf("max concurrent requests cannot be less than 1, was %d", cfg.MaxConcurrency)
	}
	return nil
}

func (cfg *L1EndpointConfig) Setup(ctx context.Context, log log.Logger, rollupCfg *rollup.Config) (l1client.RPC, *l1.L1ClientConfig, error) {
	opts := []l1client.RPCOption{
		l1client.WithHttpPollInterval(cfg.HttpPollInterval),
		l1client.WithDialBackoff(10),
	}
	if cfg.RateLimit != 0 {
		opts = append(opts, l1client.WithRateLimit(cfg.RateLimit, cfg.BatchSize))
	}

	l1Node, err := l1client.NewRPC(ctx, log, cfg.L1NodeAddr, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial L1 address (%s): %w", cfg.L1NodeAddr, err)
	}
	rpcCfg := l1.L1ClientDefaultConfig(rollupCfg, cfg.L1TrustRPC, cfg.L1RPCKind)
	rpcCfg.MaxRequestsPerBatch = cfg.BatchSize
	rpcCfg.MaxConcurrentRequests = cfg.MaxConcurrency
	return l1Node, rpcCfg, nil
}

// PreparedL1Endpoint enables testing with an in-process pre-setup RPC connection to L1
type PreparedL1Endpoint struct {
	Client          l1client.RPC
	TrustRPC        bool
	RPCProviderKind l1.RPCProviderKind
}

var _ L1EndpointSetup = (*PreparedL1Endpoint)(nil)

func (p *PreparedL1Endpoint) Setup(ctx context.Context, log log.Logger, rollupCfg *rollup.Config) (l1client.RPC, *l1.L1ClientConfig, error) {
	return p.Client, l1.L1ClientDefaultConfig(rollupCfg, p.TrustRPC, p.RPCProviderKind), nil
}

func (cfg *PreparedL1Endpoint) Check() error {
	if cfg.Client == nil {
		return errors.New("rpc client cannot be nil")
	}

	return nil
}

type L1BeaconEndpointConfig struct {
	BeaconAddr             string   // Address of L1 User Beacon-API endpoint to use (beacon namespace required)
	BeaconHeader           string   // Optional HTTP header for all requests to L1 Beacon
	BeaconFallbackAddrs    []string // Addresses of L1 Beacon-API fallback endpoints (only for blob sidecars retrieval)
	BeaconCheckIgnore      bool     // When false, halt startup if the beacon version endpoint fails
	BeaconFetchAllSidecars bool     // Whether to fetch all blob sidecars and filter locally
}

var _ L1BeaconEndpointSetup = (*L1BeaconEndpointConfig)(nil)

func (cfg *L1BeaconEndpointConfig) Setup(ctx context.Context, log log.Logger) (cl l1.BeaconClient, fb []l1.BlobSideCarsFetcher, err error) {
	var opts []client.BasicHTTPClientOption
	if cfg.BeaconHeader != "" {
		hdr, err := parseHTTPHeader(cfg.BeaconHeader)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing beacon header: %w", err)
		}
		opts = append(opts, client.WithHeader(hdr))
	}

	for _, addr := range cfg.BeaconFallbackAddrs {
		b := client.NewBasicHTTPClient(addr, log)
		fb = append(fb, l1.NewBeaconHTTPClient(b))
	}

	a := client.NewBasicHTTPClient(cfg.BeaconAddr, log, opts...)
	return l1.NewBeaconHTTPClient(a), fb, nil
}

func (cfg *L1BeaconEndpointConfig) Check() error {
	if cfg.BeaconAddr == "" && !cfg.BeaconCheckIgnore {
		return errors.New("expected L1 Beacon API endpoint, but got none")
	}
	return nil
}

func (cfg *L1BeaconEndpointConfig) ShouldIgnoreBeaconCheck() bool {
	return cfg.BeaconCheckIgnore
}

func (cfg *L1BeaconEndpointConfig) ShouldFetchAllSidecars() bool {
	return cfg.BeaconFetchAllSidecars
}

func parseHTTPHeader(headerStr string) (http.Header, error) {
	h := make(http.Header, 1)
	s := strings.SplitN(headerStr, ": ", 2)
	if len(s) != 2 {
		return nil, errors.New("invalid header format")
	}
	h.Add(s[0], s[1])
	return h, nil
}
