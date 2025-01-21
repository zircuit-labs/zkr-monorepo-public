// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { Deploy } from "scripts/Deploy.s.sol";
import { L1UpgradeUtils } from "scripts/upgrades/UpgradeUtils.sol";
import { ChainAssertions } from "scripts/ChainAssertions.sol";
import { L1StandardBridge } from "src/L1/L1StandardBridge.sol";
import { SuperchainConfig } from "src/L1/SuperchainConfig.sol";
import { L1CrossDomainMessenger } from "src/L1/L1CrossDomainMessenger.sol";

/// @notice ad hoc script to upgrade the verifier
contract BridgeThrottlingDeploy is L1UpgradeUtils {

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "109-BridgeThrottling";
    }

    function run() override pure public {
        revert("Use one of --sig with 'runDirectly()' or 'runSimulateFromMultisig()'");
    }

    function runDirectly() public {
        console.log("Running directly from", msg.sender);
        // deploy the new implementations
        deployL1StandardBridge();
        deployOptimismPortal();
        // helper contract for resetting initialized in the second step
        ensureDeployedNoAuthStorageSetter();
    }

    function runSimulateFromMultisig() public {
        console.log("Simulating deployment from", msg.sender);

        // StandardBridge received a new storage value (accessController) so we
        // need to initialize it again
        upgradeToStorageSetter(mustGetAddress("L1StandardBridgeProxy"));
        resetInitialized(mustGetAddress("L1StandardBridgeProxy"), true);
        upgradeL1StandardBridge();

        upgradeProxyViaProxyAdmin(mustGetAddress("OptimismPortalProxy"), mustGetAddress("OptimismPortal"));

        console.log("Performing sanity checks");
        sanityCheckDeployment();
    }

    function upgradeL1StandardBridge() internal {
        upgradeProxyViaProxyAdmin({
            _proxy: mustGetAddress("L1StandardBridgeProxy"),
            _impl: mustGetAddress("L1StandardBridge"),
            _data: abi.encodeCall(L1StandardBridge.initialize,
                (L1CrossDomainMessenger(mustGetAddress("L1CrossDomainMessengerProxy")), SuperchainConfig(mustGetAddress("SuperchainConfigProxy")))
            )
        });
    }

    function sanityCheckDeployment() internal view {
        ChainAssertions.checkL1StandardBridge(_proxies(), true);
    }
}
