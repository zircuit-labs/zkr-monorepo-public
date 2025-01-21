// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Config } from "scripts/Config.sol";
import { LibString } from "solady/utils/LibString.sol";

/// @title ThrottlingConfig
/// @notice Contains the configuration for throttling on different networks
library ThrottlingConfig {
    using LibString for string;

    /// @notice settings for any throttling config
    struct ThrottleConfig {
        /// max amount that can be transferred in `periodLength` time.
        /// can be per-user or global, depending on the throttle
        uint208 maxAmountPerPeriod;
        uint48  periodLength;

        /// total amount that can be locked, ignored in some configs
        uint256 maxAmountTotal;

        /// ignored for eth configs
        address token;
    }

    /// @notice The config for L1StandardBridge for eth deposits
    function ethDepositThrottleL1StandardBridgeConfig(string memory configName) internal pure returns (ThrottleConfig memory config) {
        if (configName.eq("mainnet")) {
            // disabled: no per-user limit
            config = ThrottleConfig({
                maxAmountPerPeriod: 0,
                periodLength: 0,
                maxAmountTotal: 0, // enforced on OptimismPortal
                token: address(0)
            });
        } else if (configName.eq("alphanet-sepolia")) {
            // disabled: no per-user limit
            config = ThrottleConfig({
                maxAmountPerPeriod: 0,
                periodLength: 0,
                maxAmountTotal: 0, // enforced on OptimismPortal
                token: address(0)
            });
        } else if (configName.eq("betanet-sepolia")) {
            config = ThrottleConfig({
                maxAmountPerPeriod: 5e18, // 5 eth
                periodLength: 2 hours,
                maxAmountTotal: 0, // enforced on OptimismPortal
                token: address(0)
            });
        } else if (configName.eq("betanet2-sepolia")) {
            config = ThrottleConfig({
                maxAmountPerPeriod: 5e18, // 5 eth
                periodLength: 2 hours,
                maxAmountTotal: 0, // enforced on OptimismPortal
                token: address(0)
            });
        }
    }

    /// @notice The config for OptimismPortal and L2ToL1MessagePasser for eth deposits and withdrawals
    function ethThrottleConfig(string memory configName) internal pure returns (ThrottleConfig memory depositConfig, ThrottleConfig memory withdrawalConfig) {
        if (configName.eq("mainnet")) {
            // disabled: no total limit
            depositConfig = ThrottleConfig({
                maxAmountPerPeriod: 0, // enforced per-user on L1StandardBridge
                periodLength: 0,
                maxAmountTotal: 0,
                token: address(0)
            });
            // disabled: no limit
            withdrawalConfig = ThrottleConfig({
                maxAmountPerPeriod: 0, // enforced globally
                periodLength: 0,
                maxAmountTotal: 0, // unused
                token: address(0)
            });
        } else if (configName.eq("alphanet-sepolia")) {
            // disabled: no total limit
            depositConfig = ThrottleConfig({
                maxAmountPerPeriod: 0, // enforced per-user on L1StandardBridge
                periodLength: 0,
                maxAmountTotal: 0,
                token: address(0)
            });
            // allow 1k eth per hour to guard against big hacks being withdrawn at once
            withdrawalConfig = ThrottleConfig({
                maxAmountPerPeriod: 0, // enforced globally
                periodLength: 0,
                maxAmountTotal: 0, // unused
                token: address(0)
            });
        } else if (configName.eq("betanet-sepolia")) {
            depositConfig = ThrottleConfig({
                maxAmountPerPeriod: 0, // enforced per-user on L1StandardBridge
                periodLength: 0,
                maxAmountTotal: 1500 ether,
                token: address(0)
            });
            withdrawalConfig = ThrottleConfig({
                maxAmountPerPeriod: 300 ether, // enforced globally
                periodLength: 1 hours,
                maxAmountTotal: 0, // unused
                token: address(0)
            });
        } else if (configName.eq("betanet2-sepolia")) {
            depositConfig = ThrottleConfig({
                maxAmountPerPeriod: 0, // enforced per-user on L1StandardBridge
                periodLength: 0,
                maxAmountTotal: 1500 ether,
                token: address(0)
            });
            withdrawalConfig = ThrottleConfig({
                maxAmountPerPeriod: 300 ether, // enforced globally
                periodLength: 1 hours,
                maxAmountTotal: 0, // unused
                token: address(0)
            });
        }
    }
}
