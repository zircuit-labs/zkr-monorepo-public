// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { Deploy } from "scripts/Deploy.s.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { L2OutputOracle } from "src/L1/L2OutputOracle.sol";
import { NoAuthStorageSetter as StorageSetter } from "src/universal/StorageSetter.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { Constants } from "src/libraries/Constants.sol";

/// @notice Script to change the proposer by upgrading to storage setter and back
contract UpdateProposerDeploy is Deploy {
    address storageSetter;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "108-UpdateProposer";
    }

    function run() override pure public {
        revert("Use one of --sig with 'runDirectly()' or 'runSimulateFromMultisig()'");
    }

    function runDirectly() public {
        console.log("Running directly from", msg.sender);
        ensureDeployedNoAuthStorageSetter();
    }

    function runSimulateFromMultisig() public {
        console.log("Simulating deployment from", msg.sender);

        upgradeToProposer();
    }

    function getStorageSetter() internal pure returns (address setter) {
        bytes32 salt = bytes32(0);
        setter = vm.computeCreate2Address(salt, keccak256(type(StorageSetter).creationCode));
    }

    function ensureDeployedNoAuthStorageSetter() internal {
        storageSetter = getStorageSetter();
        if (storageSetter.code.length > 0) {
            console.log("StorageSetter contract is already deployed");
            return;
        }
        console.log("Deploying StorageSetter contract");
        vm.broadcast();
        storageSetter = address(new StorageSetter{salt: 0}());
        require(storageSetter != address(0), "deployment failed");
        require(storageSetter.code.length > 0, "no code at deployed address");
    }

    function upgradeToProposer() broadcast internal {
        address newProposer = cfg.l2OutputOracleProposer();
        address proxyToUpgrade = mustGetAddress("L2OutputOracleProxy");
        L2OutputOracle l2oo = L2OutputOracle(proxyToUpgrade);
        address oldProposer = l2oo.proposer();
        require(newProposer != oldProposer, "Proposer in config is equal to proposer on contract. Either update the config or upgrade has already been done.");

        address oldImplementation = address(uint160(uint256(vm.load(proxyToUpgrade, Constants.PROXY_IMPLEMENTATION_ADDRESS))));
        address storageSetterImpl = getStorageSetter();

        require(oldImplementation.code.length > 0, "Current implementation of L2OutputOracleProxy has no code. Something must be wrong");
        require(storageSetterImpl.code.length > 0, "StorageSetter implementation has no code. Did you run the 'runDirectly()` function first?");

        console.log("Upgrading to StorageSetter");
        console.log("Upgrading", proxyToUpgrade, "to", storageSetterImpl);

        ProxyAdmin proxyAdmin = ProxyAdmin(mustGetAddress("ProxyAdmin"));
        proxyAdmin.upgrade({_proxy: payable(proxyToUpgrade), _implementation: storageSetterImpl});

        // jq -r '.[] | select(.label == "proposer") | .slot' snapshots/storageLayout/L2OutputOracle.json
        bytes32 proposerSlot = bytes32(uint256(8));
        StorageSetter proxySetter = StorageSetter(proxyToUpgrade);
        proxySetter.setAddress(proposerSlot, newProposer);

        // upgrade back to the previous implementation
        proxyAdmin.upgrade({_proxy: payable(proxyToUpgrade), _implementation: oldImplementation});

        address currentImplementation = address(uint160(uint256(vm.load(proxyToUpgrade, Constants.PROXY_IMPLEMENTATION_ADDRESS))));
        require(currentImplementation == oldImplementation, "Implementation did not revert to what it was set to previously");
        require(l2oo.proposer() == newProposer, "Proposer was not updated");
    }
}
