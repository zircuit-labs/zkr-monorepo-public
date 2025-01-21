#!/usr/bin/env bash
# script for testing deposit exclusions by sending deposits from two accounts,
# one of them should be configured to be deny-listed
# Use the following command to filter for the bitmaps
#   docker logs ops-bedrock-op-node-1 --follow | grep "deposit_exclusions=0b"
set -eu

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

NUM_DEPOSITS=8
MNEMONIC="test test test test test test test test test test test junk"
MNEMONIC_INDEX_0=7
MNEMONIC_INDEX_1=8
SENDER_0=$(cast wallet addr --mnemonic "$MNEMONIC" --mnemonic-index $MNEMONIC_INDEX_0)
SENDER_1=$(cast wallet addr --mnemonic "$MNEMONIC" --mnemonic-index $MNEMONIC_INDEX_1)

PORTAL_ADDR=$(jq -r ".address" < "${SCRIPT_DIR}/../deployments/devnetL1/OptimismPortalProxy.json")

NONCE_0=$(cast nonce $SENDER_0)
NONCE_1=$(cast nonce $SENDER_1)

EXPECTED_BITMAP="0b0"
for i in $(seq 1 ${NUM_DEPOSITS})
do
  # decide which account to send it from
  # ACCOUNT_IDX=$(($i % 2))
  ACCOUNT_IDX=$(($RANDOM % 2))
  EXPECTED_BITMAP="${EXPECTED_BITMAP}${ACCOUNT_IDX}"

  if [[ $ACCOUNT_IDX == 0 ]];
  then
    MNEMONIC_INDEX=$MNEMONIC_INDEX_0
    NONCE=$NONCE_0
    NONCE_0=$(($NONCE_0 + 1))
  else
    MNEMONIC_INDEX=$MNEMONIC_INDEX_1
    NONCE=$NONCE_1
    NONCE_1=$(($NONCE_1 + 1))
  fi
  echo -n "From ${MNEMONIC_INDEX}: "
  cast send "${PORTAL_ADDR}" --value 1 --async --gas-limit 500000 --legacy --nonce "${NONCE}" --mnemonic "$MNEMONIC" --mnemonic-index $MNEMONIC_INDEX
done

echo -e "\nExpected bitmap: ${EXPECTED_BITMAP}"
