#!/usr/bin/env bash
# script for deploying the bridge throttling changes
set -e -o pipefail

# Check if cast is available
if ! command -v cast &> /dev/null; then
  echo "Error: 'cast' command not found. Please install cast (part of Foundry) and make sure it's in your PATH. https://github.com/foundry-rs/foundry"
  exit 1
fi

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

if [ -z "$L1_RPC" ]; then
  echo "L1_RPC is not set"
  exit 1
fi

if [[ "$#" -lt 1 ]]; then
  echo "Usage (shared options):"
  echo '       export L1_RPC=https://...'
  echo "       $0 <deployment_name>"
  echo "  e.g. $0 alphanet-sepolia"
  exit 1
fi

# Parse command line argument
DEPLOYMENT_NAME="$1"
shift

# Define the deployments directory
DEPLOYMENTS_DIR="${SCRIPT_DIR}/../../deployments/$DEPLOYMENT_NAME"

if [[ ! -d "$DEPLOYMENTS_DIR" ]] ;
then
  echo "deployment directory ${DEPLOYMENTS_DIR} does not exist."
  echo "Did you supply the correct network?"
  exit 1
fi

export DEPLOYMENT_CONTEXT=${DEPLOYMENT_NAME}
export IMPL_SALT="${DEPLOYMENT_NAME}"

SCRIPT_NAME="ApplyL1ThrottlingConfig.s.sol"
CONTRACT_NAME="L1ThrottlingScript"
THROTTLING_ADMIN=$(jq -r ".address" <"${DEPLOYMENTS_DIR}/SuperchainConfigProxy.json")
SENDER=$(cast call "${THROTTLING_ADMIN}" "defaultAdmin()(address)" --rpc-url "${L1_RPC}")
echo "> Simulating deployment from ${SENDER} for ${DEPLOYMENT_CONTEXT}"
forge script -vvvv "scripts/throttling/${SCRIPT_NAME}:${CONTRACT_NAME}" \
  --rpc-url "${L1_RPC}" --sender "${SENDER}" \
  --sig 'runSimulateFromMultisig(string memory)' "${DEPLOYMENT_NAME}"

FOUNDRY_OUTPUT_PATH="${SCRIPT_DIR}/../../broadcast/${SCRIPT_NAME}/$(cast chain-id --rpc-url ${L1_RPC})/dry-run/runSimulateFromMultisig-latest.json"
OUTPUT_PATH="${SCRIPT_DIR}/input_l1.json"
cp "${FOUNDRY_OUTPUT_PATH}" "${OUTPUT_PATH}"

echo "========================================================"
echo "Input json can be found at $(realpath ${OUTPUT_PATH})"
