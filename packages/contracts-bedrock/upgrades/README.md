# Scripts to perform smart contract upgrades

## Multisigs with thresholds of more than 1

### Usage
Set the rpc url `export UPGRADE_RPC=sepolia` (foundry aliases work).

Call one of the scripts and specify the deployment to upgrade, e.g.
```
upgrades/l0_transfer_proxyadmin_owner.sh alphanet-sepolia 0xNEW_OWNER
```

If the execution was successful, there should now be a json file in this
directory called `input.json`. Use this file as input for a new/existing task
in `packages/ops-contracts`.

### Solidity (forge) upgrade scripts
The `packages/contracts-bedrock/scripts/upgrades` directory contains the
solidity code for upgrading contracts. The naming convention is to prefix the
contract name with either 0, 1 or 2, depending on whether the upgrade is
universal or limited to L1 or L2. Since upgrades scripts are generally obsolete
after they have been executed, there should be two digits after the initial one
that is always incremented by one for each script. This gives us a quick
overview on which scripts are the recent ones.

### Considerations when writing solidity scripts
Our proxies are generally owned by a (gnosis) safe multisig wallet. As a result,
we cannot simply run forge scripts to perform upgrades. The scripts should
generally be written in a way that the sender is assumed to be the safe wallet,
so calls to contracts should be made directly. When deploying new contracts,
these should be deployed through a CREATE2 proxy (or separately) since the
smart contract wallet cannot perform regular create calls.

### Executing the upgrades
The upgrade scripts should typically only be used to generate the list of
transactions that will be passed to the multisig tooling. To do so, create a
new script in this directory and use the `run_forge_script` function. It will
simulate the script as if it was sent by the `SystemOwnerSafe` of the specified
deployment and save the transaction data as
`contracts-bedrock/upgrades/input.json`


## OLD/PRE-MULTISIG WAY OF DOING THINGS

### Usage
Set required environment variables.
```bash
# set deployer mnemonic
export MNEMONIC_DEPLOYER=$(cat ~/secrets/sepolia_alphanet_mnemonic.txt)
# set l1 rpc, e.g. sepolia. Has to be an http(s) endpoint, not websocket
# for testing, use http://localhost:8545 and run `anvil --fork-url $L1_RPC`
export L1_RPC=https://eth-sepolia.g.alchemy.com/v2/TODO
```

Call one of the scripts and specify the deployment to upgrade, e.g.
```
upgrades/l1_foundry_upgrade.sh alphanet-sepolia
```

This script will simulate all upgrade transactions and broadcast them to the
network if the upgrade ran successfully. After that, it will serialize the new
artifacts to the deployments directory.

### New upgrade
1. Copy one of the existing `l1_*_upgrade.sh` scripts.
2. Create a new script `scripts/upgrade/??-???.s.sol` similar to the existing
   ones.
