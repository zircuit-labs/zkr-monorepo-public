#!/usr/bin/env bash
# This script starts the L1 devnet. The creation of the genesis is now done
# in bedrock-devnet/devnet/__init__.py
set -eu

# Helper method that waits for a given URL to be up. Can't use
# cURL's built-in retry logic because connection reset errors
# are ignored unless you're using a very recent version of cURL
function wait_up {
  echo -n "Waiting for $1 to come up..."
  i=0
  until curl -s -f -o /dev/null "$1"
  do
    echo -n .
    sleep 1

    ((i=i+1))
    if [ "$i" -eq 300 ]; then
      echo " Timeout!" >&2
      exit 1
    fi
  done
  echo "Done!"
}

hex_time=$(jq -r '.timestamp' ../.devnet/genesis-l1.json)
export L1_GENESIS_TIME=$((16#${hex_time#0x}))

docker compose up create-beacon-chain-genesis geth-genesis
docker compose up beacon-chain -d
docker compose up l1 validator -d

wait_up localhost:8545 # Waiting for geth to start.
