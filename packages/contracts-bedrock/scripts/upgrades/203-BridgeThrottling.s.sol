// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { Deploy } from "scripts/Deploy.s.sol";
import { L2UpgradeUtils } from "scripts/upgrades/UpgradeUtils.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { L2Controller } from "src/L2/L2Controller.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { L2StandardBridge } from "src/L2/L2StandardBridge.sol";
import { L2ToL1MessagePasser } from "src/L2/L2ToL1MessagePasser.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

/// @notice ad hoc script to upgrade the verifier
contract BridgeThrottlingDeploy is L2UpgradeUtils {
    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "203-BridgeThrottling";
    }

    function run() override pure public {
        revert("Use one of --sig with 'runDirectly()' or 'runSimulateFromMultisig()'");
    }

    /// @notice Deploy contracts from any address
    function runDirectly() public {
        console.log("Running directly from", msg.sender);
        // deploy the new implementations
        deployL2Controller();
        deployL2StandardBridge();
        deployL2ToL1MessagePasser();

        // helper contract for resetting initialized in the second step
        ensureDeployedNoAuthStorageSetter();
    }

    function L2CONTROLLER_DEPLOY_ADDR() internal pure returns (address _addr) {
        _addr = vm.computeCreate2Address(0, keccak256(type(L2Controller).creationCode));
    }

    function deployL2Controller() internal {
        if (L2CONTROLLER_DEPLOY_ADDR().code.length > 0) {
            console.log("L2Controller already deployed");
            return;
        }

        vm.broadcast();
        L2Controller l2Controller = new L2Controller{salt: 0}();
        console.log("L2Controller implementation deployed at", address(l2Controller));
        require(address(l2Controller) == L2CONTROLLER_DEPLOY_ADDR(), "L2Controller was not deployed to expected address");
    }

    function L2STANDARDBRIDGE_DEPLOY_ADDR() internal pure returns (address _addr) {
        _addr = vm.computeCreate2Address(0, keccak256(type(L2StandardBridge).creationCode));
    }

    function deployL2StandardBridge() internal {
        if (L2STANDARDBRIDGE_DEPLOY_ADDR().code.length > 0) {
            console.log("L2StandardBridge already deployed");
            return;
        }
        vm.broadcast();
        L2StandardBridge l2StandardBridge = new L2StandardBridge{salt: 0}();
        console.log("L2StandardBridge implementation deployed at", address(l2StandardBridge));
        require(address(l2StandardBridge) == L2STANDARDBRIDGE_DEPLOY_ADDR(), "L2StandardBridge was not deployed to expected address");
    }

    function L2TOL1MESSAGEPASSER_DEPLOY_ADDR() internal pure returns (address _addr) {
        _addr = vm.computeCreate2Address(0, keccak256(type(L2ToL1MessagePasser).creationCode));
    }

    function deployL2ToL1MessagePasser() internal {
        if (L2TOL1MESSAGEPASSER_DEPLOY_ADDR().code.length > 0) {
            console.log("L2ToL1MessagePasser already deployed");
            return;
        }
        vm.broadcast();
        L2ToL1MessagePasser l2ToL1MessagePasser = new L2ToL1MessagePasser{salt: 0}();
        console.log("L2ToL1MessagePasser implementation deployed at", address(l2ToL1MessagePasser));
        require(address(l2ToL1MessagePasser) == L2TOL1MESSAGEPASSER_DEPLOY_ADDR(), "L2ToL1MessagePasser was not deployed to expected address");
    }

    /// @notice Upgrade that is atomically executed in a single transaction
    function runSimulateFromMultisig() public {
        console.log("Simulating deployment from", msg.sender);

        // L2Controller does not exist yet, so all we need to do is upgrade it
        console.log("Starting with L2Controller");
        upgradeL2Controller();

        // StandardBridge received a new storage value (accessController) so we
        // need to initialize it again
        console.log("Starting with L2StandardBridge");
        StandardBridge otherBridge = L2StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE)).otherBridge();
        upgradeToStorageSetter(Predeploys.L2_STANDARD_BRIDGE);
        resetInitialized({_addr: Predeploys.L2_STANDARD_BRIDGE, _expectInitialized: true});
        upgradeL2StandardBridge(otherBridge);

        // L1ToL2MessagePasser also received a new storage value (accessController)
        // so we need to initialize it again
        console.log("Starting with L2ToL1MessagePasser");
        upgradeToStorageSetter(Predeploys.L2_TO_L1_MESSAGE_PASSER);
        resetInitialized({_addr: Predeploys.L2_TO_L1_MESSAGE_PASSER, _expectInitialized: true});
        upgradeL2ToL1MessagePasser();

        console.log("Performing sanity checks");
        sanityCheckDeployment();
    }

    function upgradeL2Controller() internal {
        upgradeProxyViaProxyAdmin({
            _proxy: Predeploys.L2_CONTROLLER,
            _impl:  L2CONTROLLER_DEPLOY_ADDR(),
            // (admin, paused)
            _data: abi.encodeCall(L2Controller.initialize, (cfg.superchainConfigGuardian(), false)),
            // newly introduce contract so implementation will not have code
            _checkForExistingImpl: false
        });
    }

    function upgradeL2StandardBridge(StandardBridge otherBridge) internal {
        upgradeProxyViaProxyAdmin({
            _proxy: Predeploys.L2_STANDARD_BRIDGE,
            _impl:  L2STANDARDBRIDGE_DEPLOY_ADDR(),
            _data: abi.encodeCall(L2StandardBridge.initialize, (otherBridge))
        });
    }

    function upgradeL2ToL1MessagePasser() internal {
        upgradeProxyViaProxyAdmin({
            _proxy: Predeploys.L2_TO_L1_MESSAGE_PASSER,
            _impl:  L2TOL1MESSAGEPASSER_DEPLOY_ADDR(),
            _data: abi.encodeCall(L2ToL1MessagePasser.initialize, ())
        });
    }

    function sanityCheckDeployment() internal view {
        L2Controller l2Controller = L2Controller(Predeploys.L2_CONTROLLER);
        require(l2Controller.defaultAdmin() == cfg.superchainConfigGuardian(), "l2controller admin");
        require(l2Controller.paused() == false, "l2controller paused");

        L2StandardBridge l2StandardBridge = L2StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE));
        require(l2StandardBridge.accessController() == l2Controller, "l2standardbridge accessController");

        L2ToL1MessagePasser l2ToL1MessagePasser = L2ToL1MessagePasser(payable(Predeploys.L2_TO_L1_MESSAGE_PASSER));
        require(l2ToL1MessagePasser.accessController() == l2Controller, "l2tol1messagepasser accessController");
    }
}
