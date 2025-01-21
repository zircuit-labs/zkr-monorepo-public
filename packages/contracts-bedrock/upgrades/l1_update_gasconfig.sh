#!/usr/bin/env bash
# update the gas config for the ecotone upgrade using the values from the deploy-config

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

export IMPL_SALT="${DEPLOYMENT_NAME}"
echo "> Calling from ${SENDER} for ${DEPLOYMENT_CONTEXT}"
forge script -vvv scripts/upgrades/106-EcotoneUpgrade.s.sol:EcotoneUpgrade --sig "updateGasConfig()" --rpc-url "${UPGRADE_RPC}" --broadcast --mnemonics "${MNEMONIC_DEPLOYER}" --sender "${SENDER}"
