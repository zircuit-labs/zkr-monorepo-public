// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { StorageSlot } from "scripts/Deployer.sol";
import { Deploy } from "scripts/Deploy.s.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { StorageSetter } from "src/universal/StorageSetter.sol";
import { ChainAssertions } from "scripts/ChainAssertions.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { SystemConfig } from "src/L1/SystemConfig.sol";
import { SuperchainConfig } from "src/L1/SuperchainConfig.sol";

/// @notice script to upgrade the L1 contracts for ecotone
contract EcotoneUpgrade is Deploy {
    address storageSetter;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "106-EcotoneUpgrade";
    }

    /// @notice update the gas config on the SystemConfig contract with the values
    ///         from the corresponding deploy-config
    function updateGasConfig() public {
        SystemConfig systemConfigProxy = SystemConfig(mustGetAddress("SystemConfigProxy"));
        vm.startBroadcast();
        _callViaSafe({
            _target: address(systemConfigProxy),
            _data: abi.encodeCall(SystemConfig.setGasConfigEcotone, (cfg.basefeeScalar(), cfg.blobbasefeeScalar()))
        });
        vm.stopBroadcast();
        require(systemConfigProxy.basefeeScalar() == cfg.basefeeScalar(), "basefeeScalar mismatch");
        require(systemConfigProxy.blobbasefeeScalar() == cfg.blobbasefeeScalar(), "basefeeScalar mismatch");
    }

    function run() override public {
        // save the old scalar/overhead values since the initialization will overwrite them with blob gas config values
        SystemConfig systemConfigProxy = SystemConfig(mustGetAddress("SystemConfigProxy"));
        uint256 oldScalar = systemConfigProxy.scalar();
        uint256 oldOverhead = systemConfigProxy.overhead();

        storageSetter = deployStorageSetter(uint16(uint256(keccak256("ecotone"))));

        console.log("Deploying implementations");
        deployL1CrossDomainMessenger();
        deployOptimismMintableERC20Factory();
        deploySystemConfig();
        deployL1StandardBridge();
        deployL1ERC721Bridge();
        deployOptimismPortal();
        deployL2OutputOracle();

        upgradeToStorageSetter("SystemConfig");
        resetInitialized("SystemConfig", true);

        upgradeToStorageSetter("OptimismPortal");
        resetInitialized("OptimismPortal", true);

        upgradeToStorageSetter("L1CrossDomainMessenger");
        resetInitialized("L1CrossDomainMessenger", true);

        upgradeToStorageSetter("L2OutputOracle");
        resetInitialized("L2OutputOracle", true);

        upgradeToStorageSetter("L1StandardBridge");
        resetInitialized("L1StandardBridge", true);

        upgradeToStorageSetter("L1ERC721Bridge");
        resetInitialized("L1ERC721Bridge", true);

        // Factory was not initialized before this upgrade, so no need to reset
        // upgradeToStorageSetter("OptimismMintableERC20Factory");
        // resetInitialized("OptimismMintableERC20Factory", false);

        // not part of the upgrade but let's check that the owner is set correctly
        SuperchainConfig superchainConfigProxy = SuperchainConfig(mustGetAddress("SuperchainConfigProxy"));
        if (superchainConfigProxy.owner() != cfg.superchainConfigGuardian()) {
            console.log("Correcting superchain config owner");
            upgradeToStorageSetter("SuperchainConfig");
            address rightOwner = cfg.superchainConfigGuardian();

            StorageSetter setter = StorageSetter(address(superchainConfigProxy));
            vm.startBroadcast();
            // taken from storage snapshot
            // overwrite current admin
            // will overwrite pendingDelay but that doesn't matter
            setter.setAddress(bytes32(uint256(152)), rightOwner);

            // upgrade back to old implementation
            _upgradeViaSafe({
                _proxy: payable(address(superchainConfigProxy)),
                _implementation: mustGetAddress("SuperchainConfig")
            });
            vm.stopBroadcast();
        }

        console.log("Initializing implementations");
        initializeSystemConfig();

        // restore old (pre-ecotone) gas config so we can upgrade before the hard fork
        vm.startBroadcast();
        _callViaSafe({
            _target: address(systemConfigProxy),
            _data: abi.encodeCall(SystemConfig.setGasConfig, (oldOverhead, oldScalar))
        });
        vm.stopBroadcast();

        initializeL1StandardBridge();
        initializeL1ERC721Bridge();
        initializeOptimismMintableERC20Factory();
        initializeL1CrossDomainMessenger();
        initializeL2OutputOracle();
        initializeOptimismPortal();
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
            console.log("getBytes32 on", address(setter));
            if (setter.getBytes32(bytes32(storageSlot)) == bytes32(0)) {
                revert(string.concat("Storage slot was zero before resetting for ", _name));
            }
        }

        console.log("setUint on", address(setter));
        setter.setUint(bytes32(storageSlot), 0);
    }
}
