// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { Proxy } from "src/universal/Proxy.sol";

/// @title IStaticERC1967Proxy
/// @notice IStaticERC1967Proxy is a static version of the ERC1967 proxy interface.
interface IStaticERC1967Proxy {
    function implementation() external view returns (address);

    function admin() external view returns (address);
}

/// @title ProxyAdmin
/// @notice This is an auxiliary contract meant to be assigned as the admin of an ERC1967 Proxy,
///         based on the OpenZeppelin implementation. It has backwards compatibility logic to work
///         with the various types of proxies that have been deployed by Optimism in the past.
contract ProxyAdmin is Ownable {
    /// @notice The proxy types that the ProxyAdmin can manage.
    /// @custom:value ERC1967    Represents an ERC1967 compliant transparent proxy interface.
    enum ProxyType { ERC1967 }

    /// @notice A mapping of proxy types, used for backwards compatibility.
    mapping(address => ProxyType) public proxyType;

    /// @param _owner Address of the initial owner of this contract.
    constructor(address _owner) Ownable(_owner) { }

    /// @notice Sets the proxy type for a given address. Only required for non-standard (legacy)
    ///         proxy types.
    /// @param _address Address of the proxy.
    /// @param _type    Type of the proxy.
    function setProxyType(address _address, ProxyType _type) external onlyOwner {
        proxyType[_address] = _type;
    }

    /// @notice Returns the implementation of the given proxy address.
    /// @param _proxy Address of the proxy to get the implementation of.
    /// @return Address of the implementation of the proxy.
    function getProxyImplementation(address _proxy) external view returns (address) {
        ProxyType ptype = proxyType[_proxy];
        if (ptype == ProxyType.ERC1967) {
            return IStaticERC1967Proxy(_proxy).implementation();
        } else {
            revert("ProxyAdmin: unknown proxy type");
        }
    }

    /// @notice Returns the admin of the given proxy address.
    /// @param _proxy Address of the proxy to get the admin of.
    /// @return Address of the admin of the proxy.
    function getProxyAdmin(address payable _proxy) external view returns (address) {
        ProxyType ptype = proxyType[_proxy];
        if (ptype == ProxyType.ERC1967) {
            return IStaticERC1967Proxy(_proxy).admin();
        } else {
            revert("ProxyAdmin: unknown proxy type");
        }
    }

    /// @notice Updates the admin of the given proxy address.
    /// @param _proxy    Address of the proxy to update.
    /// @param _newAdmin Address of the new proxy admin.
    function changeProxyAdmin(address payable _proxy, address _newAdmin) external onlyOwner {
        ProxyType ptype = proxyType[_proxy];
        if (ptype == ProxyType.ERC1967) {
            Proxy(_proxy).changeAdmin(_newAdmin);
        } else {
            revert("ProxyAdmin: unknown proxy type");
        }
    }

    /// @notice Changes a proxy's implementation contract.
    /// @param _proxy          Address of the proxy to upgrade.
    /// @param _implementation Address of the new implementation address.
    function upgrade(address payable _proxy, address _implementation) public onlyOwner {
        ProxyType ptype = proxyType[_proxy];
        if (ptype == ProxyType.ERC1967) {
            Proxy(_proxy).upgradeTo(_implementation);
        } else {
            // It should not be possible to retrieve a ProxyType value which is not matched by
            // one of the previous conditions.
            assert(false);
        }
    }

    /// @notice Changes a proxy's implementation contract and delegatecalls the new implementation
    ///         with some given data. Useful for atomic upgrade-and-initialize calls.
    /// @param _proxy          Address of the proxy to upgrade.
    /// @param _implementation Address of the new implementation address.
    /// @param _data           Data to trigger the new implementation with.
    function upgradeAndCall(
        address payable _proxy,
        address _implementation,
        bytes memory _data
    )
        external
        payable
        onlyOwner
    {
        ProxyType ptype = proxyType[_proxy];
        if (ptype == ProxyType.ERC1967) {
            Proxy(_proxy).upgradeToAndCall{ value: msg.value }(_implementation, _data);
        } else {
            // reverts if proxy type is unknown
            upgrade(_proxy, _implementation);
            (bool success,) = _proxy.call{ value: msg.value }(_data);
            require(success, "ProxyAdmin: call to proxy after upgrade failed");
        }
    }
}
