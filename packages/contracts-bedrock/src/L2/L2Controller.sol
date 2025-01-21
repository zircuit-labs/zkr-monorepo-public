// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { ISemver } from "src/universal/ISemver.sol";
import { AccessControlPausable } from "src/universal/AccessControlPausable.sol";

/// @custom:audit none This contracts is not yet audited.
/// @title L2Controller
/// @notice The L2Controller contract is used to manage access control on L2.
contract L2Controller is AccessControlPausable, ISemver {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Constructs the SuperchainConfig contract.
    constructor() {
        initialize({ _admin: address(0xdead), _paused: false });
    }

    /// @notice Initializer.
    /// @param _admin    Address of the admin, can control access roles.
    /// @param _paused      Initial paused status.
    function initialize(address _admin, bool _paused) public initializer {
        // assign the _admin address the DEFAULT_ADMIN_ROLE
        // changing admin addresses requires 1 day to pass
        __AccessControlPausable_init(1 days, _admin);
        if (_paused) {
            _pause("Initializer paused");
        }
    }
}
