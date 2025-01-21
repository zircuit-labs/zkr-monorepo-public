#!/usr/bin/env bash

# This script is used to generate the four wallets that are used in the Getting
# Started quickstart guide on the docs site. Simplifies things for users
# slightly while also avoiding the need for users to manually copy/paste a
# bunch of stuff over to the environment file.

mnemonic=$(cast wallet new-mnemonic | grep "Phrase" --after-context 1 | tail --lines 1)

# Grab wallet addresses
address1=$(cast wallet address --mnemonic "${mnemonic}" --mnemonic-index 0)
address2=$(cast wallet address --mnemonic "${mnemonic}" --mnemonic-index 1)
address3=$(cast wallet address --mnemonic "${mnemonic}" --mnemonic-index 2)
address4=$(cast wallet address --mnemonic "${mnemonic}" --mnemonic-index 3)

# Grab wallet private keys
key1=$(cast wallet derive-private-key "${mnemonic}" 0 | awk '/Private key/ { print $3 }')
key2=$(cast wallet derive-private-key "${mnemonic}" 1 | awk '/Private key/ { print $3 }')
key3=$(cast wallet derive-private-key "${mnemonic}" 2 | awk '/Private key/ { print $3 }')
key4=$(cast wallet derive-private-key "${mnemonic}" 3 | awk '/Private key/ { print $3 }')

# Print out the environment variables to copy
echo "# This is your deployment information, save the mnemonic"
echo "export MNEMONIC_DEPLOYER=\"${mnemonic}\""
echo "# paste the output into your terminal to set it up for 'config.sh'"
echo
echo "# Admin account"
echo "export GS_ADMIN_ADDRESS=$address1"
echo "export GS_ADMIN_PRIVATE_KEY=$key1"
echo
echo "# Batcher account"
echo "export GS_BATCHER_ADDRESS=$address2"
echo "export GS_BATCHER_PRIVATE_KEY=$key2"
echo
echo "# Proposer account"
echo "export GS_PROPOSER_ADDRESS=$address3"
echo "export GS_PROPOSER_PRIVATE_KEY=$key3"
echo
echo "# Sequencer account"
echo "export GS_SEQUENCER_ADDRESS=$address4"
echo "export GS_SEQUENCER_PRIVATE_KEY=$key4"
