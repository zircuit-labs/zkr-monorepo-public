#!/bin/bash
set -e

export DEPLOYMENT_CONTEXT=$(cat /tmp/network_name)
export DEPLOY_CONFIG_PATH=/src/packages/contracts-bedrock/deploy-config/${DEPLOYMENT_CONTEXT}.json
DEPLOYMENTS_PATH=/src/packages/contracts-bedrock/deployments/${DEPLOYMENT_CONTEXT}

CHAIN_ID=$(jq -r ".l2ChainID" < "${DEPLOY_CONFIG_PATH}")
STARTING_BLOCK_HASH=$(jq -r ".l1StartingBlockTag" < ${DEPLOY_CONFIG_PATH})

export CONTRACT_ADDRESSES_PATH="${DEPLOYMENTS_PATH}/.genesis_l1_deployments"
export STATE_DUMP_PATH=/src/packages/contracts-bedrock/state-dump.json
(cd /src/packages/contracts-bedrock && forge script --chain-id "${CHAIN_ID}" scripts/L2Genesis.s.sol:L2Genesis --sig "runWithAllUpgrades()" --private-key "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")

(cd /src/op-node/cmd && go run main.go genesis l2 \
              --deploy-config "${DEPLOY_CONFIG_PATH}" \
              --l1-deployments "${CONTRACT_ADDRESSES_PATH}" \
              --l2-allocs "${STATE_DUMP_PATH}" \
              --outfile.l2 /workdir/genesis.json \
              --outfile.rollup /workdir/rollup.json \
              --l1-rpc "${L1_RPC_NODE}")

cat "${DEPLOY_CONFIG_PATH}"
echo "generated genesis + rollup config at /workdir/{genesis,rollup}.json"
