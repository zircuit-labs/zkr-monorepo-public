// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";
import { Deploy } from "scripts/Deploy.s.sol";
import { NoAuthStorageSetter as StorageSetter } from "src/universal/StorageSetter.sol";
import { ForgeArtifacts } from "scripts/ForgeArtifacts.sol";
import { Constants } from "src/libraries/Constants.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

abstract contract UpgradeUtils is Deploy {

    function getProxyAdmin() virtual internal view returns (ProxyAdmin proxyAdmin);

    /// @notice uses vm.load to get the value of the implementation slot
    function loadImpl(address _proxy) internal view returns (address _impl){
        _impl = address(uint160(uint256(vm.load(_proxy, Constants.PROXY_IMPLEMENTATION_ADDRESS))));
    }

    /// @notice Return the address of the StorageSetter contract and check that
    ///         the address contains code. Reverts if there is no code.
    function getStorageSetterChecked() internal view returns (address setter) {
      setter = getStorageSetter();
      require(setter.code.length > 0, "StorageSetter contract is not present, make sure to deploy it first");
    }

    /// @notice Return the address of the StorageSetter contract
    function getStorageSetter() internal pure returns (address setter) {
        bytes32 salt = bytes32(0);
        setter = vm.computeCreate2Address(salt, keccak256(type(StorageSetter).creationCode));
    }

    /// @notice Return the address of the StorageSetter contract, check if there
    ///         is code at this address and if there is not, deploy it.
    function ensureDeployedNoAuthStorageSetter() internal returns (address storageSetter) {
        storageSetter = getStorageSetter();
        if (storageSetter.code.length == 0) {
          console.log("Deploying StorageSetter contract");
          vm.broadcast();
          storageSetter = address(new StorageSetter{salt: 0}());
          require(storageSetter != address(0), "deployment failed");
          require(storageSetter.code.length > 0, "no code at deployed address");
        }
    }

    /// @notice upgrade the proxy of the contract to the storage setter contract
    ///         and set the slot 0 to 0, allowing initialize to be called again.
    ///         This should only be used on contracts that have the the
    ///         initialized/initializing field in slot 0, otherwise it will
    ///         break the contract.
    function resetInitialized(address _addr, bool _expectInitialized) internal {
        console.log("Resetting initialized");
        console.log("  on   ", _addr);
        bytes32 slot = ForgeArtifacts.getInitializedSlot();
        // reset the storage slot and assert that it was not zero before.
        // if it was, it means we likely picked the wrong slot since all of the
        // contracts should be initialized
        resetStorageSlot(_addr, slot, _expectInitialized);
    }

    /// @notice Upgrade the proxy to the StorageSetter contract
    function upgradeToStorageSetter(address _proxy) broadcast internal {
        address storageSetter = getStorageSetterChecked();
        ProxyAdmin proxyAdmin = getProxyAdmin();
        console.log("Upgrading to StorageSetter");
        console.log("  proxy", _proxy);
        console.log("  from ", loadImpl(_proxy));
        console.log("  to   ", storageSetter);

        proxyAdmin.upgrade({_proxy: payable(_proxy), _implementation: storageSetter});
    }

    /// @notice set the storage slot `storageSlot` to 0. Proxy implementation
    ///         must already be set to the StorageSetter contract.
    function resetStorageSlot(address proxy, bytes32 storageSlot, bool checkNonZero) broadcast internal {
        address storageSetter = getStorageSetterChecked();
        address implementation = address(uint160(uint256(vm.load(proxy, bytes32(uint256(keccak256('eip1967.proxy.implementation')) - 1)))));
        if (implementation != storageSetter) {
            console.log("implementation: ", implementation);
            console.log("storageSetter : ", storageSetter);
            revert("resetStorageSlot requires implementation to be StorageSetter");
        }

        StorageSetter setter = StorageSetter(proxy);
        if (checkNonZero) {
            if (setter.getBytes32(storageSlot) == bytes32(0)) {
                console.log("Storage slot was zero for", proxy);
                revert("Storage slot was zero before resetting");
            }
        }

        setter.setUint(storageSlot, 0);
    }

    function upgradeProxyViaProxyAdmin(address _proxy, address _impl) internal {
        upgradeProxyViaProxyAdmin(_proxy, _impl, "");
    }

    function upgradeProxyViaProxyAdmin(address _proxy, address _impl, bytes memory _data) internal {
        upgradeProxyViaProxyAdmin(_proxy, _impl, _data, true);
    }

    function upgradeProxyViaProxyAdmin(address _proxy, address _impl, bytes memory _data, bool _checkForExistingImpl) broadcast internal {
        address currentImpl = loadImpl(_proxy);

        require(!_checkForExistingImpl || currentImpl.code.length > 0, "Current implementation has no code. Something must be wrong");
        require(_impl.code.length > 0, "New implementation has no code. Did you run the 'runDirectly()` function first?");

        console.log("Upgrading");
        console.log("  proxy", _proxy);
        console.log("  from ", currentImpl);
        console.log("  to   ", _impl);

        ProxyAdmin proxyAdmin = getProxyAdmin();
        if (_data.length > 0) {
            proxyAdmin.upgradeAndCall({_proxy: payable(_proxy), _implementation: _impl, _data: _data});
        } else {
            proxyAdmin.upgrade({_proxy: payable(_proxy), _implementation: _impl});
        }

        address implAfterUpgrade = loadImpl(_proxy);
        require(implAfterUpgrade == _impl, "Implementation was not updated like expected");
    }
}

abstract contract L1UpgradeUtils is UpgradeUtils {
    function getProxyAdmin() override internal view returns (ProxyAdmin proxyAdmin) {
        proxyAdmin = ProxyAdmin(mustGetAddress("ProxyAdmin"));
        require(address(proxyAdmin).code.length > 0, "ProxyAdmin has no code. Are you on the wrong chain?");
    }
}

abstract contract L2UpgradeUtils is UpgradeUtils {
    function getProxyAdmin() override internal view returns (ProxyAdmin proxyAdmin) {
        proxyAdmin = ProxyAdmin(Predeploys.PROXY_ADMIN);
        require(address(proxyAdmin).code.length > 0, "ProxyAdmin has no code. Are you on the wrong chain?");
    }
}

