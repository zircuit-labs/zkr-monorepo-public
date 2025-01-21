// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { StorageSlot } from "scripts/Deployer.sol";
import { Deploy } from "scripts/Deploy.s.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { StorageSetter } from "src/universal/StorageSetter.sol";
import { ChainAssertions } from "scripts/ChainAssertions.sol";
import { Proxy } from "src/universal/Proxy.sol";

/// @notice This upgrade script is not safe against frontrunning. In the time between
///         upgrading the contracts to the StorageSetter implementation to reset the
///         initialization value and upgrading to the actual implementation, anyone
///         could set any storage slot, e.g. the owner of the contract and then
///         upgrade it themselves for a complete takeover. Since we only have a testnet
///         and did not add the safe bundle code from Optimism yet, this is likely
///         good enough.
contract FoundryMigrationUpgrade is Deploy {
    address storageSetter;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "01-FoundryMigration";
    }

    function runWithStateDiff() public override stateDiff {
        run();
    }

    function run() override public {
        console.log("Deploying new contracts and upgrading the smart contracts from before the foundry migration");

        storageSetter = deployStorageSetter(0);

        // deploy a safe that will be the owner of the ProxyAdmin
        deploySafe();
        transferProxyAdminOwnership();

        // deploy a proxy and the superchain config
        deployERC1967Proxy("SuperchainConfigProxy");
        deploySuperchainConfig();
        initializeSuperchainConfig();

        // deploy the implementations of the new contracts
        deploySystemConfig();
        deployOptimismPortal();
        deployL1CrossDomainMessenger();
        deployL2OutputOracle();
        deployOptimismMintableERC20Factory();
        deployL1StandardBridge();
        deployL1ERC721Bridge();

        // upgrade the implementations
        // contracts that were already using an initializer and thus need to
        // be upgraded in two steps, first resetting initialization and then
        // upgrading + initializing
        upgradeToStorageSetter("SystemConfig");
        resetInitialized("SystemConfig", true);
        initializeSystemConfig();

        upgradeToStorageSetter("OptimismPortal");
        resetInitialized("OptimismPortal", true);
        initializeOptimismPortal();

        upgradeToStorageSetter("L1CrossDomainMessenger");
        resetInitialized("L1CrossDomainMessenger", true);
        initializeL1CrossDomainMessenger();

        upgradeToStorageSetter("L2OutputOracle");
        resetInitialized("L2OutputOracle", true);
        initializeL2OutputOracle();

        upgradeToStorageSetter("L1StandardBridge");
        resetInitialized("L1StandardBridge", true);
        initializeL1StandardBridge();

        upgradeToStorageSetter("L1ERC721Bridge");
        // L1ERC721 bridge was not initialized before this upgrade
        resetInitialized("L1ERC721Bridge", false);
        initializeL1ERC721Bridge();

        // the factory can simply be upgraded without any initialization
        initializeOptimismMintableERC20Factory();
    }

    /// @notice upgrade the proxy of the contract to the storage setter contract
    ///         and set the slot 0 to 0, allowing initialize to be called again.
    ///         This should only be used on contracts that have the the
    ///         initialized/initializing field in slot 0, otherwise it will
    ///         break the contract.
    function resetInitialized(string memory _name, bool expectInitialized) internal {
        console.log("Resetting initialized for", _name);
        StorageSlot memory slot = getInitializedSlot(_name);
        // reset the storage slot and assert that it was not zero before.
        // if it was, it means we likely picked the wrong slot since all of the
        // contracts should be initialized
        resetStorageSlot(_name, vm.parseUint(slot.slot), expectInitialized);
    }

    /// @notice Upgrade the proxy to the StorageSetter contract and claim ownership
    ///         so only the current sender is allowed to make changes
    function upgradeToStorageSetter(string memory _name) broadcast internal {
        console.log("Upgrading", _name, "to StorageSetter");
        address proxyToUpgrade = mustGetAddress(string.concat(_name, "Proxy"));

        _upgradeAndCallViaSafe({
            _proxy: payable(proxyToUpgrade),
            _implementation: storageSetter,
            _innerCallData: abi.encodeCall(StorageSetter.claimOwnership, (msg.sender))
        });
    }

    /// @notice set the storage slot `storageSlot` to 0. Proxy implementation
    ///         must already be set to the StorageSetter contract.
    function resetStorageSlot(string memory _name, uint256 storageSlot, bool checkNonZero) broadcast internal {
        address payable proxy = mustGetAddress(string.concat(_name, "Proxy"));
        address implementation = address(uint160(uint256(vm.load(proxy, bytes32(uint256(keccak256('eip1967.proxy.implementation')) - 1)))));
        if (implementation != storageSetter) {
            console.log("implementation: ", implementation);
            console.log("storageSetter : ", storageSetter);
            revert("resetStorageSlot requires implementation to be StorageSetter");
        }

        StorageSetter setter = StorageSetter(proxy);
        if (checkNonZero) {
            if (setter.getBytes32(bytes32(storageSlot)) == bytes32(0)) {
                revert(string.concat("Storage slot was zero before resetting for ", _name));
            }
        }

        setter.setUint(bytes32(storageSlot), 0);
    }
}
