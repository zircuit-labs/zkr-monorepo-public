#!/usr/bin/env bash
# This script starts everything except the L1 components.

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

export INDEXER_CONSUMER_L1CROSSDOMAINMESSENGER=$(cat ../.devnet/addresses.json | jq -r ".L1CrossDomainMessengerProxy")
export PROPOSER_L2OUTPUTORACLEADDRESS=$(cat ../.devnet/addresses.json | jq -r ".L2OutputOracleProxy")
export PROPOSER_OPTIMISMPORTAL_ADDRESS=$(cat ../.devnet/addresses.json | jq -r ".OptimismPortalProxy")
# if the replica is enabled, this also starts the replica op-node
docker compose up $(docker compose config --services | grep -E "\<(op-node)\>") -d
wait_up localhost:9545 # Waiting for op-node to start

docker compose up $(docker compose config --services | grep -Ev "\<(l1|create-beacon-chain-genesis|geth-genesis|beacon-chain|validator|op-node)\>") -d
