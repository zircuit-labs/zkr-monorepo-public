# ZKR-MONOREPO

## Prerequisites

This repository contains git submodules and should be cloned with `git clone --recurse-submodules`.
Alternatively, `git submodule update --init --recursive` can be used on an existing repository
to download the submodules.

### Earthly

Install [Earthly](https://earthly.dev). The instructions can be found in [this file](https://earthly.dev/get-earthly). On MacOS, this is running

```shell
brew install earthly && earthly bootstrap
```

To avoid pull ping error while running the system, set `EARTHLY_DISABLE_REMOTE_REGISTRY_PROXY="true"`.

### Foundry

Install [Foundry](https://github.com/foundry-rs) by running

```shell
curl -L https://foundry.paradigm.xyz | bash && foundryup
./ops/scripts/install-foundry.sh
```

Other installation options are described [here](https://book.getfoundry.sh/getting-started/installation).

### jq

Install jq by running:

```shell
brew install jq
```

## Devnet

### Starting devnet

```shell
make devnet-up
```

### Shutting down devnet

```shell
make devnet-down
```

### Clean devnet state

This command removes images and volumes and any other persisted state for the devnet.

```shell
make devnet-clean
```

### View devnet logs

```shell
make devnet-logs
```

## Enabling ZKTrie

### Method 1: from genesis

This is only suitable for local deployment and hard-fork with ZKTrie enabled, which isn't preferred solution.

It is possible to start l2-geth with ZKTrie enabled from the genesis block. To achieve this, make the following changes in code:

1. In [entrypoint-l2.sh](./ops-bedrock/entrypoint-l2.sh) add `--zktrie-enabled=true` flag into both `geth` commands. Ensure that the flag is supplied right after `geth` command before any subcommands. It should look like that:

```shell
	# ...
	geth \
	  --zktrie-enabled=true \
	  --verbosity="$VERBOSITY" init \
	# ...
exec geth \
	--zktrie-enabled=true \
	--datadir="$GETH_DATA_DIR" \
	--verbosity="$VERBOSITY" \
	# ...
```

2. In [Python generation script](./bedrock-devnet/devnet/__init__.py) in function `devnet_l2_genesis` supply `'--zktrie-enabled', 'true'`, so it looks like that:

```python
    run_command([
        'go', 'run', 'cmd/main.go', 'genesis', 'l2',
        '--l1-starting-block', pjoin(paths.devnet_dir, 'l1-starting-block.json'),
        '--deploy-config', paths.devnet_config_path,
        '--l2-allocs', l2_allocs_path,
        '--l1-deployments', paths.addresses_json_path,
        '--outfile.l2', paths.genesis_l2_path,
        '--outfile.rollup', paths.rollup_config_path,
        '--zktrie-enabled', 'true'
    ], cwd=paths.op_node_dir)
```

3. Next steps are the same: log into ECR on need and run `make devnet-up`.

### Method 2: Block on local network

This method is preferred to test in general due to more similarity to the live network.

1. In [devnetL1-template.json](./packages/contracts-bedrock/deploy-config/devnetL1-template.json) set the desired `zkTrieSwitchBlock` value.
2. Run `make devnet-up` as usual.
3. The migration starts as soon as you launch the node, but ZKTrie state is only used at blocks >= `zkTrieSwitchBlock`.

### Method 3: Block on live network

Add `ZKTrieSwitchBlock` in the live network genesis.json into `.config` along with other blocks.
If you have `gethconfig.toml`, it might be necessary to add there as well.
Restarting node is required.

Ensure to set the big enough block number for migration to finish at that time.
One of the ways is to set very big switch block, set debug log level, monitor how
quickly the blocks are processed and reset the switch block to the desired value.
Better to do it on devnet first.

### Checking

You will have log about ZKTrie enabled if you set the respective flag.
In case you set up block, seek for logs with `zk_state_migrator` entry.
`zktrie` entries will help to understand when ZKTrie logic is executed.

- "Migration finished" log appears when current head reaches the switch block.
- "switching to subscription" log appears when all blocks before the current head have been processed.
  You want this log to understand much time the migration will take to choose the switch block.

### Design

This section describes the implementation details of ZKTrie feature.

First of all, ZKTrie may be enabled via `--zktrie-enabled` flag.
The same underlying DB is used, but the nodes format and hashes are different from Trie.
This can only be enabled on genesis initialization and running on top of ZKTrie genesis.
Trie will not be used at all.

On `geth --zktrie-enabled=true init genesis.json` genesis state is processed and persisted with ZKTrie.
There must be no previous Trie state.

On `geth --zktrie-enabled=true run` all states are processed with ZKTrie.
Genesis must be initialized with ZKTrie to run this, and vice versa: the flag must be specified if genesis has ZKTrie state.

For existing network the migration should be performed, and flag shouldn't be set.
When `ZKTrieSwitchBlock` is set in genesis, the existing state is not broken,
because the chain configuration (`params.ChainConfig`) does not affect hash values.
`gethconfig.toml` currently shouldn't have impact on ZKTrie, but it wasn't tested.

Setting `ZKTrieSwitchBlock` triggers migration when `NewBlockChain` is called (not on `init`, but on `run`).
The migration runs in the background and doesn't affect the node's operation.
If it fails, several retries with period are done, then if neither of them succeeds,
the error log appears and the node continues to work with Trie until restart.

The migrator looks up for migration state. If it's not found, the genesis state is processed first.
After that each block is read from chain and processed in order.
Intermediate ZKTrie root is committed and stored in ZKTrie table.
If migration is interrupted, the migrator will continue from the last processed block.
The only difference in processing compared to Trie is using StateDB instance with underlying ZKTrie.
Nothing is changed in existing Trie blocks, including genesis, despite committing the state into the same underlying `ethdb.Database`.

When current head <= `ZKTrieSwitchBlock` is reached, migrator switches to subscription mode, using `ChainHeadEvent`.
It ensures that all incoming blocks are strictly ordered, failing otherwise.
As soon as `ZKTrieSwitchBlock - 1` is processed, the migrator stops and notifies about migration end.

The existing services retrieve the latest **ZKTrie** root for `ZKTrieSwitchBlock - 1` and use it to process `ZKTrieSwitchBlock` block,
waiting for migration to finish on need. Then all the states become block-dependent:
ZKTrie is used when block number is unavailable and for blocks >= `ZKTrieSwitchBlock`.
Otherwise, old Trie states are used.
Trie states are not changed at this point, unless reconnecting blocks or hard-forks are done (not tested).
