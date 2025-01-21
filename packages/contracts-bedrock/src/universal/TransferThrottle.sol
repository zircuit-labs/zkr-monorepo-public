// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { Storage } from "src/libraries/Storage.sol";

/// @title TransferThrottle
/// @notice Provide common functions for implementing throttling for any value transfers
abstract contract TransferThrottle {
    /// @notice address to use for `user` in `_transferThrottling` if throttling
    ///         should be applied globally
    address constant _throttleGlobalUser = address(0);

    struct ThrottleEntry {
        uint208 lastUpdateCredits;
        uint48 lastUpdateTimestamp;
    }

    /// @notice Contains the configuration and the current values for throttling
    ///         any assets that flow through this contract.
    /// @custom:field maxAmountPerPeriod         Maximum amount that can pass through the contract within `periodLength`.
    /// @custom:field periodLength               Number of seconds each throttling period lasts.
    /// @custom:field lastUpdateCredits          The credits that were available at `lastUpdateTimestamp`.
    /// @custom:field lastUpdateTimestamp        Timestamp of the last update.
    struct Throttle {
        // throughput configuration
        uint208 maxAmountPerPeriod;
        uint48 periodLength;

        // accounting
        mapping(address => ThrottleEntry) entries;

        // maximum amount that can be stored in total
        uint256 maxAmountTotal;

        // unused
        uint256[9] _reserved;
    }

    /// @notice Returns the available credits for `_user`, not taking into account
    ///         the total locked value
    function _throttleUserAvailableCredits(address _user, Throttle storage throttle) internal view returns (uint256 availableCredits) {
        uint256 maxAmountPerPeriod = throttle.maxAmountPerPeriod;
        // no throttle set
        if (maxAmountPerPeriod == 0) {
            return type(uint256).max;
        }
        uint48 periodLength = throttle.periodLength;

        ThrottleEntry storage entry = throttle.entries[_user];
        availableCredits = entry.lastUpdateCredits;
        uint256 secondsSinceLastUpdate = block.timestamp - entry.lastUpdateTimestamp;

        // calculate how many credits we have available by taking the last known credits
        // and adding the amount generated since the last withdrawal
        availableCredits += maxAmountPerPeriod * secondsSinceLastUpdate / periodLength;

        // cap the available credits to the max allowed per period
        if (availableCredits > maxAmountPerPeriod) {
            availableCredits = maxAmountPerPeriod;
        }
    }

    /// @notice Checks whether `_address` is allowed to change all throttle
    ///         configurations (including disabling it).
    function _transferThrottleHasAdminAccess(address _address) internal view virtual;

    /// @notice Checks whether `_address` is allowed to decrease the amount
    ///         that is allowed within the configured period. The function needs
    ///         to revert if `_address` is not allowed.
    function _transferThrottleHasThrottleAccess(address _address) internal view virtual;

    /// @notice Sets the length of the throttle period to `_periodLength`, which
    ///         immediately affects the speed of credit accumulation.
    function _setPeriodLength(uint48 _periodLength, Throttle storage throttle) internal {
        _transferThrottleHasAdminAccess(msg.sender);
        require(_periodLength != 0, "TransferThrottle: period length cannot be 0");
        throttle.periodLength = _periodLength;
    }

    /// @notice Sets the max amount per period of `throttle` to `maxAmountPerPeriod` if the sender is allowed
    ///         to make the update. The required capabilities are determined based on whether
    ///         the new value is an increase (`_transferThrottleHasAdminAccess` only) or a decrease
    ///         (`_transferThrottleHasThrottleAccess`)
    function _setThrottle(uint208 maxAmountPerPeriod, uint256 maxAmountTotal, Throttle storage throttle) internal {
        uint256 currentMaxAmountPerPeriod = throttle.maxAmountPerPeriod;
        uint256 currentMaxAmountTotal = throttle.maxAmountTotal;
        // reductions in max value only require monitor capabilities, whereas
        // increases are limited to operators
        if (maxAmountPerPeriod <= currentMaxAmountPerPeriod && maxAmountTotal <= currentMaxAmountTotal) {
            _transferThrottleHasThrottleAccess(msg.sender);
        } else {
            _transferThrottleHasAdminAccess(msg.sender);
        }

        // ensure the period length is initialized if max per period is set
        if (maxAmountPerPeriod != 0 && throttle.periodLength == 0) {
            throttle.periodLength = 1 hours;
        }

        // updating the max amount per period will require a full period to pass
        // in the worst case (all credits were depleted)
        throttle.maxAmountPerPeriod = maxAmountPerPeriod;
        throttle.maxAmountTotal = maxAmountTotal;
    }

    /// @notice Perform accounting operations and enforce throttling of the total
    ///         `value` that is allowed across the bridge in the configured period, as well
    ///         as the total amount locked in the bridge. These can be configured/enabled independently.
    ///         `existingValue` is the amount stored in the contract before the transfer.
    ///         Contracts can use address(0) for the `user` parameter to apply global transfer throttling.
    function _transferThrottling(Throttle storage throttle, address user, uint256 existingValue, uint256 value) internal {
        uint256 maxAmountPerPeriod = throttle.maxAmountPerPeriod;
        uint256 maxAmountTotal = throttle.maxAmountTotal;

        // check global cap
        if (maxAmountTotal != 0 && (existingValue + value) > maxAmountTotal) {
            revert("TransferThrottle: maximum allowed total amount exceeded");
        }

        // disabled
        if (maxAmountPerPeriod == 0)
            return;

        unchecked {
            ThrottleEntry storage entry = throttle.entries[user];
            uint48 periodLength = throttle.periodLength;
            uint256 availableCredits = entry.lastUpdateCredits;
            uint256 secondsSinceLastUpdate = block.timestamp - entry.lastUpdateTimestamp;

            // calculate how many credits we have available by taking the last known credits
            // and adding the amount generated since the last withdrawal
            // this is safe since the following holds
            // type(uint256).max > (uint256(type(uint208).max) * type(uint48).max + type(uint208).max)
            availableCredits += maxAmountPerPeriod * secondsSinceLastUpdate / periodLength;

            // cap the available credits to the max allowed per period
            if (availableCredits > maxAmountPerPeriod) {
                availableCredits = maxAmountPerPeriod;
            }

            require(value <= availableCredits, "TransferThrottle: maximum allowed throughput exceeded");
            // this cast is safe since availableCredits is lte than maxAmountPerPeriod, which has been
            // assigned from a uint208
            entry.lastUpdateCredits = uint208(availableCredits - value);
            entry.lastUpdateTimestamp = uint48(block.timestamp);
        }
    }
}
