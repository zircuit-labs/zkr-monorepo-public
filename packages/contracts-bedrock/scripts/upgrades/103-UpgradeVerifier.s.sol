// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { StorageSlot } from "scripts/Deployer.sol";
import { Deploy } from "scripts/Deploy.s.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { StorageSetter } from "src/universal/StorageSetter.sol";
import { ChainAssertions } from "scripts/ChainAssertions.sol";
import { Proxy } from "src/universal/Proxy.sol";

/// @notice ad hoc script to upgrade the verifier
contract UpgradeVerifierDeploy is Deploy {
    address storageSetter;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "103-UpgradeVerifier";
    }

    function run() override public pure {
        revert("Use one of --sig with 'runDirectly()' or 'runSimulateFromMultisig()'");
    }

    function runDirectly() public {
        console.log("Deploying verifier from", msg.sender);
        deployVerifier();
    }

    function runSimulateFromMultisig() public {
        console.log("Simulating deployment from", msg.sender);
        upgradeToVerifier();
    }

    function upgradeToVerifier() broadcast internal {
        console.log("Upgrading verifier");
        address proxyToUpgrade = mustGetAddress("VerifierProxy");
        address implementation = mustGetAddress("Verifier");
        console.log("Upgrading", proxyToUpgrade, "to", implementation);

        ProxyAdmin proxyAdmin = ProxyAdmin(mustGetAddress("ProxyAdmin"));
        proxyAdmin.upgrade({_proxy: payable(proxyToUpgrade), _implementation: implementation});
    }
}
