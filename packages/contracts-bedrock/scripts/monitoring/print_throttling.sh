#!/usr/bin/env bash

set -e

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Check if cast is available
if ! command -v cast &> /dev/null; then
  echo "Error: 'cast' command not found. Please install cast (part of Foundry) and make sure it's in your PATH. https://github.com/foundry-rs/foundry"
  exit 1
fi

if [[ "$#" -lt 1 ]]; then
  echo "Usage (shared options):"
  echo '       export L1_RPC=https://...'
  echo '       export L2_RPC=https://...'
  echo "       $0 <deployment_name>"
  echo "  e.g. $0 alphanet-sepolia"
  exit 1
fi

# Parse command line argument
deployment_name="$1"
shift

# Define the deployments directory
deployments_dir="${SCRIPT_DIR}/../../deployments/$deployment_name"

if [[ ! -d "$deployments_dir" ]] ;
then
  echo "deployment directory ${deployments_dir} does not exist."
  echo "Did you supply the correct network?"
  exit 1
fi

# Check if L1_RPC is set
if [[ ! -v L1_RPC ]]; then
  echo "L1_RPC is not set. Please enter the L1_RPC URL:"
  read -r L1_RPC
  export L1_RPC
fi

# Check if L1_RPC is set
if [[ ! -v L2_RPC ]]; then
  echo "L2_RPC is not set. Please enter the L2_RPC URL:"
  read -r L2_RPC
  export L2_RPC
fi

# OptimismPortal
L1_CONTRACT=$(jq -r '.address' < "${deployments_dir}/OptimismPortalProxy.json")
# L2ToL1MessagePasser
L2_CONTRACT="0x4200000000000000000000000000000000000016"
RETURN_SIG="returns (uint208 maxAmountPerPeriod, uint48 periodLength, uint256 maxTotal)"

L1_DEPOSIT_CONFIG=$(cast call "$L1_CONTRACT" "ethThrottleDeposits() $RETURN_SIG" --rpc-url $L1_RPC)
# TODO
# L1_DEPOSIT_DATA=$(cast call "$L1_CONTRACT" "ethThrottleDeposits() $RETURN_SIG" --rpc-url $L1_RPC)
L1_WITHDRAWAL_CONFIG=$(cast call "$L1_CONTRACT" "ethThrottleWithdrawals() $RETURN_SIG" --rpc-url $L1_RPC)
L1_ADMIN=$(cast call "$L1_CONTRACT" "guardian() returns (address)" --rpc-url $L1_RPC)
L1_PAUSED=$(cast call "$L1_CONTRACT" "paused() returns (bool)" --rpc-url $L1_RPC)

L2_WITHDRAWAL_CONFIG=$(cast call "$L2_CONTRACT" "ethThrottleWithdrawals() $RETURN_SIG" --rpc-url $L2_RPC)
L2_ADMIN=$(cast call "$L2_CONTRACT" "accessController() returns (address)" --rpc-url $L2_RPC)
L2_PAUSED=$(cast call "$L2_CONTRACT" "paused() returns (bool)" --rpc-url $L2_RPC)

echo "L1 info"
echo "Contract: ${L1_CONTRACT}"
echo "Admin:    ${L1_ADMIN}"
echo "Paused:   ${L1_PAUSED}"

echo "deposits"
echo ${L1_DEPOSIT_CONFIG}
echo "withdrawals"
echo ${L1_WITHDRAWAL_CONFIG}

echo ""
echo "L2 info"
echo "Contract: ${L2_CONTRACT}"
echo "Admin:    ${L2_ADMIN}"
echo "Paused:   ${L2_PAUSED}"

echo "withdrawals"
echo ${L2_WITHDRAWAL_CONFIG}
