#!/usr/bin/env bash
# script for transferring the ownership of the proxy admin
set -e -o pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source "${SCRIPT_DIR}/shared_setup.sh"

if [ -z "$1" ]; then
  echo "Missing positional argument. Specify address of the new owner"
  exit 1
fi
NEW_OWNER=$1
shift

export IMPL_SALT="${DEPLOYMENT_NAME}"

SCRIPT_NAME=003-TransferProxyAdminOwnership.s.sol
CONTRACT_NAME=TransferProxyAdminOwnership
export NEW_PROXYADMIN_OWNER="${NEW_OWNER}"
run_forge_script
