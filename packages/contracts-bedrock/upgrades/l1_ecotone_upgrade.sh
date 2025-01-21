#!/usr/bin/env bash
# upgrade all L1 contracts for the ecotone upgrade

set -e -o pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source "${SCRIPT_DIR}/shared_setup.sh"

# Check if MNEMONIC_DEPLOYER is set
if [ -z "$MNEMONIC_DEPLOYER" ]; then
  echo "MNEMONIC_DEPLOYER is not set. Please enter the mnemonic:"
  read -r MNEMONIC_DEPLOYER
  export MNEMONIC_DEPLOYER
fi

# Check if ETHERSCAN_API_KEY is set for verification
if [ -z "$ETHERSCAN_API_KEY" ]; then
  echo "ETHERSCAN_API_KEY is not set. Please enter the mnemonic:"
  read -r ETHERSCAN_API_KEY
  export ETHERSCAN_API_KEY
fi

# Deploy l1 smart contracts
SENDER=$(cast wallet address --mnemonic "${MNEMONIC_DEPLOYER}")

export IMPL_SALT="${DEPLOYMENT_NAME}"
echo "> Deploying contracts from ${SENDER} for ${DEPLOYMENT_CONTEXT}"
forge script -vvv scripts/upgrades/106-EcotoneUpgrade.s.sol:EcotoneUpgrade --rpc-url "${UPGRADE_RPC}" --broadcast --mnemonics "${MNEMONIC_DEPLOYER}" --sender "${SENDER}" --verify

echo "> Generating hardhat artifacts"
forge script -vvv scripts/upgrades/106-EcotoneUpgrade.s.sol:EcotoneUpgrade --sig 'sync()' --rpc-url "${UPGRADE_RPC}" --sender "${SENDER}"
