#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Check if cast is available
if ! command -v cast &> /dev/null; then
  echo "Error: 'cast' command not found. Please install cast (part of Foundry) and make sure it's in your PATH. https://github.com/foundry-rs/foundry"
  exit 1
fi

if [[ "$#" -lt 1 ]]; then
  echo "Usage (shared options):"
  echo '       export MNEMONIC_DEPLOYER=$(cat /path_to/deployer_mnemonic.txt)'
  echo '       export UPGRADE_RPC=https://...'
  echo "       $0 <deployment_name>"
  echo "  e.g. $0 alphanet-sepolia"
  exit 1
fi

# Parse command line argument
DEPLOYMENT_NAME="$1"
shift

# Define the deployments directory
DEPLOYMENTS_DIR="${SCRIPT_DIR}/../deployments/$DEPLOYMENT_NAME"

if [[ ! -d "$DEPLOYMENTS_DIR" ]] ;
then
  echo "deployment directory ${DEPLOYMENTS_DIR} does not exist."
  echo "Did you supply the correct network?"
  exit 1
fi

# Check if UPGRADE_RPC is set
if [ -z "$UPGRADE_RPC" ]; then
  echo "UPGRADE_RPC is not set. Please enter the UPGRADE_RPC URL:"
  read -r UPGRADE_RPC
  export UPGRADE_RPC
fi

export DEPLOYMENT_CONTEXT=${DEPLOYMENT_NAME}

# runs a forge script simulation with the SystemOwnerSafe as sender
# any additional arguments are forwarded to the forge script invocation as is
# expects SCRIPT_NAME and CONTRACT_NAME to be set
run_forge_script () {
  if [ -z "${SCRIPT_NAME}" ]; then
    echo "run_forge_script: SCRIPT_NAME is not set"
    exit 1
  fi
  if [ -z "${CONTRACT_NAME}" ]; then
    echo "run_forge_script: CONTRACT_NAME is not set"
    exit 1
  fi

  SENDER=$(jq -r ".address" <"${SCRIPT_DIR}/../deployments/${DEPLOYMENT_NAME}/SystemOwnerSafe.json")
  echo "> Simulating deployment from ${SENDER} for ${DEPLOYMENT_CONTEXT}"
  forge script -vv "scripts/upgrades/${SCRIPT_NAME}:${CONTRACT_NAME}" \
    --rpc-url "${UPGRADE_RPC}" --sender "${SENDER}" \
    --sig 'runSimulateFromMultisig()' \
    $@

  FOUNDRY_OUTPUT_PATH="${SCRIPT_DIR}/../broadcast/${SCRIPT_NAME}/$(cast chain-id --rpc-url ${UPGRADE_RPC})/dry-run/runSimulateFromMultisig-latest.json"
  OUTPUT_PATH="${SCRIPT_DIR}/input.json"
  cp "${FOUNDRY_OUTPUT_PATH}" "${OUTPUT_PATH}"

  echo "========================================================"
  echo "Input json can be found at $(realpath ${OUTPUT_PATH})"
}
