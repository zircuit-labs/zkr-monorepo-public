#!/bin/sh
set -eu

# alias for legacy GETH_MINING_ENABLED name
IS_SEQUENCER="${GETH_MINING_ENABLED:-1}"
ETH_STATS_ENABLED="${ETH_STATS_ENABLED:-0}"
#In case we are starting a replica the l2geth sequencer rpc http endpoint must be set
GETH_SEQUENCER_RPC_HTTP="${GETH_SEQUENCER_RPC_HTTP:-"http://localhost:8545"}"
VERBOSITY=${GETH_VERBOSITY:-3}
GETH_DATA_DIR=${GETH_DATA_DIR:-/db}
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
CHAIN_ID=$(cat "$GENESIS_FILE_PATH" | jq -r .config.chainId)
OP_NODE_L2_ENGINE_AUTH="${OP_NODE_L2_ENGINE_AUTH:-/config/jwt-secret.txt}"
RPC_PORT="${RPC_PORT:-8545}"
WS_PORT="${WS_PORT:-8546}"
CONFIG_FILE_PATH="/config/gethconfig.toml"
HTTP_API="web3,debug,eth,txpool,net,engine"
WS_API="debug,eth,txpool,net,engine"
GETH_TXPOOL_GLOBALQUEUE="${GETH_TXPOOL_GLOBALQUEUE:-1024}"

if [ ! -d "$GETH_CHAINDATA_DIR" ] || [ "$FORK_TRIGGER" = "true" ]; then
  echo "$GETH_CHAINDATA_DIR missing, running init"
  echo "Initializing genesis."
  geth --verbosity="$VERBOSITY" init \
    --datadir="$GETH_DATA_DIR" \
    "$GENESIS_FILE_PATH"
else
  echo "$GETH_CHAINDATA_DIR exists."
fi

if [ "${IS_SEQUENCER}" = 1 ]; then
  HTTP_API="${HTTP_API},admin"
  WS_API="${WS_API},admin"
  EXTRA_FLAGS=
else
  EXTRA_FLAGS="--rollup.sequencerhttp=$GETH_SEQUENCER_RPC_HTTP"
fi

#If eth stats is enabled then ETH_STATS_NODE, WS_SECRET, ETH_STATS_SERVER, ETH_STATS_SERVER_PORT must be provided
if [ "${ETH_STATS_ENABLED}" = 1 ]; then
  ETH_STATS_FLAG="--ethstats=${ETH_STATS_NODE}:${WS_SECRET}@${ETH_STATS_SERVER}:${ETH_STATS_SERVER_PORT}"
else
  ETH_STATS_FLAG=""
fi

# Check if the config file exists and set the CONFIG_FLAG accordingly
CONFIG_FLAG=""
if [ -f "$CONFIG_FILE_PATH" ]; then
  CONFIG_FLAG="--config=$CONFIG_FILE_PATH"
fi

# Warning: Archive mode is required, otherwise old trie nodes will be
# pruned within minutes of starting the devnet.

exec geth \
  --datadir="$GETH_DATA_DIR" \
  --verbosity="$VERBOSITY" \
  --http \
  --http.corsdomain="*" \
  --http.vhosts="*" \
  --http.addr=0.0.0.0 \
  --http.port="$RPC_PORT" \
  --http.api="$HTTP_API" \
  --ws \
  --ws.addr=0.0.0.0 \
  --ws.port="$WS_PORT" \
  --ws.origins="*" \
  --ws.api="$WS_API" \
  --syncmode=full \
  --nodiscover \
  --maxpeers=0 \
  --networkid="$CHAIN_ID" \
  --rpc.txfeecap=10 \
  --authrpc.addr="0.0.0.0" \
  --authrpc.port="8551" \
  --authrpc.vhosts="*" \
  --authrpc.jwtsecret=${OP_NODE_L2_ENGINE_AUTH} \
  --gcmode=archive \
  --circuit-capacity-check="false" \
  --txpool.globalqueue="$GETH_TXPOOL_GLOBALQUEUE" \
  ${ETH_STATS_FLAG} \
  ${EXTRA_FLAGS} \
  ${CONFIG_FLAG} \
  "$@"
