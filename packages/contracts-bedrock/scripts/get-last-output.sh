#!/bin/bash

set -e

usage () {
  echo "Usage: $0 <deployment-name> <l1-rpc-url>"
  echo "Examples:    $0 alphanet-sepolia https://eth-sepolia.g.alchemy.com/v2/..."
  echo "convenience: $0 devnetL1"
  exit 1
}

# check if args are supplied first
if [[ "$#" -eq 2 ]] ; then
  DEPLOYMENT_NAME=$1
  shift

  RPC_URL=$1
  shift

# special convenience case for devnet
elif [[ "$#" -eq 1 ]] ; then
  if [[ "$1" == "devnetL1" ]]; then
    DEPLOYMENT_NAME="$1"
    RPC_URL="http://localhost:8545"
  else
    usage
  fi
# also check the same environment arguments as the e2e tests
elif [[ ! -z "${deployment}" && ! -z "${rpcL1}" ]] ; then
  RPC_URL="${rpcL1}"
  DEPLOYMENT_NAME="${deployment}"
else
  usage
fi

set -u

if [[ "${DEPLOYMENT_NAME}" == "devnetL1" ]] ; then
  L2OO_ADDR="0x6900000000000000000000000000000000000000"
else
  SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

  # retrieve L2OO address from deployments dir
  L2OO_PROXY_PATH="${SCRIPT_DIR}/../deployments/${DEPLOYMENT_NAME}/L2OutputOracleProxy.json"
  L2OO_ADDR=$(cat "${L2OO_PROXY_PATH}" | jq ".address" | tr -d '"')
fi

# use foundry to query the L2OO for the latest block number
LATEST_BLOCK=$(cast call "${L2OO_ADDR}" "latestBlockNumber()" --rpc-url "${RPC_URL}")
LATEST_BLOCK_B10=$(cast to-base $LATEST_BLOCK 10)

# sanity check that the output is a number
case $LATEST_BLOCK_B10 in
    ''|*[!0-9]*) exit 1 ;;
    *) echo "${LATEST_BLOCK_B10}" ;;
esac

