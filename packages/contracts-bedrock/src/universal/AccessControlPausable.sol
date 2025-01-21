// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { AccessControlDefaultAdminRulesUpgradeable } from "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import { Storage } from "src/libraries/Storage.sol";

/// @title AccessControlPausable
/// @notice Provide common constants/functions for implementing access control pausability
abstract contract AccessControlPausable is AccessControlDefaultAdminRulesUpgradeable {
    /// @notice Whether or not the contract is paused.
    bytes32 public constant PAUSED_SLOT = bytes32(uint256(keccak256("paused")) - 1);

    /// @notice Addresses with this role are allowed to pause the contract but not unpause it.
    bytes32 public constant MONITOR_ROLE = keccak256("MONITOR_ROLE");

    /// @notice Addresses with this role are allowed to pause and unpause the contract.
    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");

    /// @notice Emitted when the pause is triggered.
    /// @param identifier A string helping to identify provenance of the pause transaction.
    event Paused(string identifier);

    /// @notice Emitted when the pause is lifted.
    event Unpaused();

    /// @notice Returns whether `_address` has at least monitor capabilities
    function hasMonitorCapabilities(address _address) public view returns (bool) {
        return hasRole(MONITOR_ROLE, _address)
            || hasRole(OPERATOR_ROLE, _address)
            || hasRole(DEFAULT_ADMIN_ROLE, _address);
    }

    /// @notice Returns whether `_address` has operator capabilities
    function hasOperatorCapabilities(address _address) public view returns (bool) {
        return hasRole(OPERATOR_ROLE, _address)
            || hasRole(DEFAULT_ADMIN_ROLE, _address);
    }

    /// @notice Getter for the current paused status.
    function paused() public view returns (bool paused_) {
        paused_ = Storage.getBool(PAUSED_SLOT);
    }

    /// @notice Pauses deposits and withdrawals.
    /// @param _identifier (Optional) A string to identify provenance of the pause transaction.
    function pause(string memory _identifier) external {
        require(hasMonitorCapabilities(msg.sender), "only MONITOR_ROLE or admin can pause");
        _pause(_identifier);
    }

    /// @notice Pauses deposits and withdrawals.
    /// @param _identifier (Optional) A string to identify provenance of the pause transaction.
    function _pause(string memory _identifier) internal {
        Storage.setBool(PAUSED_SLOT, true);
        emit Paused(_identifier);
    }

    /// @notice Unpauses deposits and withdrawals.
    function unpause() external {
        require(hasOperatorCapabilities(msg.sender), "only OPERATOR_ROLE or admin can unpause");
        Storage.setBool(PAUSED_SLOT, false);
        emit Unpaused();
    }

    /// @notice Sets the initial values for {defaultAdminDelay} and {defaultAdmin} address.
    function __AccessControlPausable_init(uint48 _initialDelay, address _initialDefaultAdmin) internal onlyInitializing {
        __AccessControlDefaultAdminRules_init(_initialDelay, _initialDefaultAdmin);
    }
}
