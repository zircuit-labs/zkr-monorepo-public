#!/usr/bin/env bash
# script to update the SystemConfig contract's ResourceConfig
set -e -o pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source "${SCRIPT_DIR}/shared_setup.sh"

# these values are also used in `run_forge_script`
SCRIPT_NAME=110-UpdateSystemResourceConfig.s.sol
CONTRACT_NAME=UpdateSystemResourceConfig

# Simulate the transaction from the multisig to generate the json input for signing
run_forge_script
