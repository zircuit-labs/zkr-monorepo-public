#!/usr/bin/env bash
set -euo pipefail

verify_flag=""
if [ -n "${DEPLOY_VERIFY:-}" ]; then
  verify_flag="--verify"
fi

SENDER=$(cast wallet address --mnemonic "${MNEMONIC_DEPLOYER}")
echo "> Deploying contracts"
forge script -vvv scripts/Deploy.s.sol:Deploy --rpc-url "$DEPLOY_ETH_RPC_URL" --broadcast --mnemonics "${MNEMONIC_DEPLOYER}" --sender "${SENDER}" $verify_flag

if [ -n "${DEPLOY_GENERATE_HARDHAT_ARTIFACTS:-}" ]; then
  echo "> Generating hardhat artifacts"
  forge script -vvv scripts/Deploy.s.sol:Deploy --sig 'sync()' --rpc-url "$DEPLOY_ETH_RPC_URL" --broadcast --mnemonics "${MNEMONIC_DEPLOYER}" --sender "${SENDER}"
fi
