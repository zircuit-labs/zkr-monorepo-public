#!/usr/bin/env bash
set -eu

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd "$SCRIPT_DIR"

cp ./contract_addresses_mainnet.md zircuit-docs/smart-contracts/contract_addresses.md
cp ./contract_addresses_testnet.md zircuit-docs/testnet/contract_addresses.md
cd zircuit-docs

# check if there are changes
if [ -z "$(git status --porcelain)" ]; then
  # working directory clean
  echo "Nothing to do, docs are identical."
  exit 0
fi

# set config to identify this as a non-human user
git config user.name "DocUpdater"
git config user.email "docupdater@zircuit.com"

echo "Detected changes. Creating commit and pushing new changes"
# we have uncommitted changes in tracked files that we need to push
git add smart-contracts/contract_addresses.md testnet/contract_addresses.md \
    && git commit -m "auto: updated contract addresses page (${GITHUB_RUN_NUMBER})" \
    && git push
