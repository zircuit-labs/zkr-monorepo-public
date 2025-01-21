#!/bin/bash
# redeploys the L1 smart contracts of a network using the current finalized block as starting point
set -e

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Check if cast is available
if ! command -v cast &> /dev/null; then
  echo "Error: 'cast' command not found. Please install cast (part of Foundry) and make sure it's in your PATH. https://github.com/foundry-rs/foundry"
  exit 1
fi

if [[ "$#" -lt 1 ]]; then
  echo "Usage:"
  echo '       export MNEMONIC_DEPLOYER=$(cat /path_to/deployer_mnemonic.txt)'
  echo '       export L1_RPC=https://...'
  echo '       export ETHERSCAN_API_KEY=...'
  echo "       $0 <deployment_name>"
  echo "  e.g. $0 alphanet-sepolia"
  exit 1
fi

LOCAL_TEST=false
# Parse command line arguments
while [[ "$#" -gt 0 ]]; do
    case "$1" in
        --local-test) LOCAL_TEST=true; shift ;;
        *) DEPLOYMENT_NAME="$1"; shift ;;
    esac
done

# don't run contract verification if testing locally
VERIFY_FLAG="--verify"
if [ $LOCAL_TEST = "true" ]; then
  VERIFY_FLAG=""
fi

# Check if the config JSON file argument is provided
if [ -z "$DEPLOYMENT_NAME" ]; then
  echo -e "Usage: $0 <deployment_name>"
  exit 1
fi

# Check if MNEMONIC_DEPLOYER is set
if [ -z "$MNEMONIC_DEPLOYER" ]; then
  echo "MNEMONIC_DEPLOYER is not set"
  exit 1
fi

# Check if L1_RPC is set
if [ -z "$L1_RPC" ]; then
  echo "L1_RPC is not set"
  exit 1
fi

# Check if L1_RPC is set
if [ -z "$ETHERSCAN_API_KEY" ]; then
  echo "ETHERSCAN_API_KEY is not set"
  exit 1
fi

# Extract the base name of the JSON file (without extension)
base_name=$(basename "$DEPLOYMENT_NAME" .json)

# Define the deployments directory
DEPLOYMENTS_DIR="${SCRIPT_DIR}/../deployments/${DEPLOYMENT_NAME}"
CONFIG_FILE="${SCRIPT_DIR}/../deploy-config/${DEPLOYMENT_NAME}.json"

# Check if the deployments directory exists and delete it if it does
if [ -d "$DEPLOYMENTS_DIR" ]; then
  rm -rf "$DEPLOYMENTS_DIR"
fi

# Get the timestamp and hash
BLOCK_HASH=$(cast block safe --rpc-url "$L1_RPC" --field hash)
BLOCK_TIMESTAMP=$(cast block ${BLOCK_HASH} --rpc-url "$L1_RPC" --field timestamp)

# Print the extracted values (optional)
echo "Hash: ${BLOCK_HASH}"
echo "Timestamp: ${BLOCK_TIMESTAMP} ($(date -d @${BLOCK_TIMESTAMP}))"

# update the config file
cp "${CONFIG_FILE}" "${CONFIG_FILE}.bak"
jq --arg hash "${BLOCK_HASH}" --arg timestamp "${BLOCK_TIMESTAMP}" '. + {"l1StartingBlockTag": $hash, "l2OutputOracleStartingTimestamp": ($timestamp | tonumber) }' < "${CONFIG_FILE}.bak" > "${CONFIG_FILE}"

# Check if the backup file exists and remove it if it does
if [ -f "${CONFIG_FILE}.bak" ]; then
  rm "${CONFIG_FILE}.bak"
fi

# https://docs.optimism.io/builders/chain-operators/tutorials/create-l2-rollup#:~:text=url%20%24L1_RPC_URL-,Deploy%20the%20L1%20contracts,-Once%20you%27ve%20configured
# https://docs.optimism.io/builders/chain-operators/management/troubleshooting
export IMPL_SALT=$(openssl rand -hex 32)
# Deploy l1 smart contracts
SENDER=$(cast wallet address --mnemonic "${MNEMONIC_DEPLOYER}")
export DEPLOYMENT_CONTEXT=$base_name
echo "> Deploying contracts from ${SENDER} for ${DEPLOYMENT_CONTEXT}"
forge script -vvv scripts/Deploy.s.sol:Deploy --rpc-url "${L1_RPC}" --broadcast --mnemonics "${MNEMONIC_DEPLOYER}" --sender "${SENDER}" ${VERIFY_FLAG} --slow

echo "> Generating hardhat artifacts"
forge script -vvv scripts/Deploy.s.sol:Deploy --sig 'sync()' --rpc-url "${L1_RPC}" --broadcast --mnemonics "${MNEMONIC_DEPLOYER}" --sender "${SENDER}"

echo "> Making copy of .deploy for genesis"
cp "${DEPLOYMENTS_DIR}/.deploy" "${DEPLOYMENTS_DIR}/.genesis_l1_deployments"
