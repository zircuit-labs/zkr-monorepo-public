#!/usr/bin/env bash
# script for deploying the bridge throttling changes
set -e -o pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source "${SCRIPT_DIR}/shared_setup.sh"

export IMPL_SALT="${DEPLOYMENT_NAME}"

if [ -z "$MNEMONIC_DEPLOYER" ]; then
  echo "MNEMONIC_DEPLOYER is not set"
  exit 1
fi

if [ -z "$ETHERSCAN_API_KEY" ]; then
  echo "ETHERSCAN_API_KEY is not set"
  exit 1
fi

# Deploy l1 smart contracts from deployer mnemonic directly if needed
SENDER=$(cast wallet address --mnemonic "${MNEMONIC_DEPLOYER}")

# these values are also used in `run_forge_script`
SCRIPT_NAME=109-BridgeThrottling.s.sol
CONTRACT_NAME=BridgeThrottlingDeploy

echo "> Deploying contracts from ${SENDER} for ${DEPLOYMENT_CONTEXT}"
forge script -vvv "scripts/upgrades/${SCRIPT_NAME}:${CONTRACT_NAME}" --sig 'runDirectly()' --rpc-url "${UPGRADE_RPC}" --broadcast --mnemonics "${MNEMONIC_DEPLOYER}" --sender "${SENDER}" --slow --verify

echo "> Generating hardhat artifacts"
forge script -vvv "scripts/upgrades/${SCRIPT_NAME}:${CONTRACT_NAME}" --sig 'sync()' --rpc-url "${UPGRADE_RPC}" --sender "${SENDER}"

# Simulate the upgrade from the multisig to generate the json input for signing
run_forge_script
