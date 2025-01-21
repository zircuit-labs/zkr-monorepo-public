#!/bin/bash

ENDPOINT="http://localhost:9545"

latestBlockData=$(curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", false],"id":1}' $ENDPOINT)
blockTimestampHex=$(echo $latestBlockData | jq -r '.result.timestamp')
blockTimestampDec=$((16#${blockTimestampHex#0x}))
blockTimestampDec=$((blockTimestampDec + 30)) # it must be in the future

(cd ./ops-bedrock && docker compose stop l2 op-node)

genesisFilePath="./.devnet/genesis-l2.json"
rollupFilePath="./.devnet/rollup.json"
jq --argjson time "$blockTimestampDec" '.config.canyonTime = $time | .config.shanghaiTime = $time' $genesisFilePath > temp.json && mv temp.json $genesisFilePath
jq --argjson time "$blockTimestampDec" '.canyon_time = $time' $rollupFilePath > temp.json && mv temp.json $rollupFilePath

jq ".config" ./.devnet/genesis-l2.json
jq "." ./.devnet/rollup.json

(cd ./ops-bedrock && docker compose up l2 op-node -d)
