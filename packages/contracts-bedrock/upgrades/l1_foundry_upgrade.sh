#!/usr/bin/env bash

set -e -o pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source "${SCRIPT_DIR}/shared_setup.sh"

# Check if MNEMONIC_DEPLOYER is set
if [ -z "$MNEMONIC_DEPLOYER" ]; then
  echo "MNEMONIC_DEPLOYER is not set. Please enter the mnemonic:"
  read -r MNEMONIC_DEPLOYER
  export MNEMONIC_DEPLOYER
fi

# Deploy l1 smart contracts
SENDER=$(cast wallet address --mnemonic "${MNEMONIC_DEPLOYER}")

# use a different salt for each deployment since we are deploying multiple testnets
# to the same L1
export IMPL_SALT="${DEPLOYMENT_NAME}"
echo "> Deploying contracts from ${SENDER} for ${DEPLOYMENT_CONTEXT}"
forge script -vvv scripts/upgrades/101-FoundryMigration.s.sol:FoundryMigrationUpgrade --rpc-url "${UPGRADE_RPC}" --broadcast --mnemonics "${MNEMONIC_DEPLOYER}" --sender "${SENDER}" --slow

echo "> Generating hardhat artifacts"
forge script -vvv scripts/upgrades/101-FoundryMigration.s.sol:FoundryMigrationUpgrade --sig 'sync()' --rpc-url "${UPGRADE_RPC}" --sender "${SENDER}"

