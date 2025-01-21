// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { ISemver } from "src/universal/ISemver.sol";
import { Storage } from "src/libraries/Storage.sol";

/// @title StorageSetter
/// @notice A simple contract that allows setting arbitrary storage slots.
///         WARNING: this contract is not safe to be called by untrusted parties.
///         It is only meant as an intermediate step during upgrades.
///         Ensure that ownership is being claimed atomically after upgrading
///         a proxy to this implementation.
contract StorageSetter is ISemver {
    /// @notice Represents a storage slot key value pair.
    struct Slot {
        bytes32 key;
        bytes32 value;
    }

    /// @notice Semantic version.
    /// @custom:semver 2.0.0
    string public constant version = "2.0.0";

    /// @notice Slot that stores the owner that is allowed set storage
    bytes32 public immutable OWNER_SLOT;

    /// @notice deploy a StorageSetter contract with the specified `offset`.
    ///         The offset can be used to vary the owner slot slightly in case
    ///         multiple upgrades do not use the same owner address.
    constructor(uint16 offset) {
        OWNER_SLOT = bytes32(uint256(keccak256("storagesetter.owner")) - offset);
    }

    /// @notice helper function to set the owner, useful for upgrading to this contract
    ///         and calling this function atomically
    function claimOwnership(address _owner) external {
        address owner = Storage.getAddress(OWNER_SLOT);
        require(owner == address(0), "Ownership already claimed");
        Storage.setAddress(OWNER_SLOT, _owner);
    }

    /// @notice Assert that the caller is authorized to perform state changes
    function assertValidCaller() internal virtual view {
        address owner = Storage.getAddress(OWNER_SLOT);
        if (owner != msg.sender && owner != address(0)) {
            revert("Owner set but caller is not owner");
        }
    }

    /// @notice Stores a bytes32 `_value` at `_slot`. Any storage slots that
    ///         are packed should be set through this interface.
    function setBytes32(bytes32 _slot, bytes32 _value) public {
        assertValidCaller();
        Storage.setBytes32(_slot, _value);
    }

    /// @notice Stores a bytes32 value at each key in `_slots`.
    function setBytes32(Slot[] calldata slots) public {
        assertValidCaller();
        uint256 length = slots.length;
        for (uint256 i; i < length; i++) {
            Storage.setBytes32(slots[i].key, slots[i].value);
        }
    }

    /// @notice Retrieves a bytes32 value from `_slot`.
    function getBytes32(bytes32 _slot) external view returns (bytes32 value_) {
        value_ = Storage.getBytes32(_slot);
    }

    /// @notice Stores a uint256 `_value` at `_slot`.
    function setUint(bytes32 _slot, uint256 _value) public {
        assertValidCaller();
        Storage.setUint(_slot, _value);
    }

    /// @notice Retrieves a uint256 value from `_slot`.
    function getUint(bytes32 _slot) external view returns (uint256 value_) {
        value_ = Storage.getUint(_slot);
    }

    /// @notice Stores an address `_value` at `_slot`.
    function setAddress(bytes32 _slot, address _address) public {
        assertValidCaller();
        Storage.setAddress(_slot, _address);
    }

    /// @notice Retrieves an address value from `_slot`.
    function getAddress(bytes32 _slot) external view returns (address addr_) {
        addr_ = Storage.getAddress(_slot);
    }

    /// @notice Stores a bool `_value` at `_slot`.
    function setBool(bytes32 _slot, bool _value) public {
        assertValidCaller();
        Storage.setBool(_slot, _value);
    }

    /// @notice Retrieves a bool value from `_slot`.
    function getBool(bytes32 _slot) external view returns (bool value_) {
        value_ = Storage.getBool(_slot);
    }
}

contract NoAuthStorageSetter is StorageSetter {
    constructor() StorageSetter(1) { }

    function assertValidCaller() internal override view {
        // allow anyone to perform storage changes
        // when using this contract to change values on a proxy the following
        // must happen atomically in a single transaction
        //  1. the upgrade of the implementation to this contract
        //  2. storage writes using this contract
        //  3. the upgrade of the implementation to the real implementation
    }
}
