// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { Deploy } from "scripts/Deploy.s.sol";
import { ThrottlingConfig } from "scripts/throttling/ThrottlingConfig.sol";
import { L1StandardBridge } from "src/L1/L1StandardBridge.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";

/// @notice ad hoc script to upgrade the verifier
contract L1ThrottlingScript is Deploy {
    string configName;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "ApplyL1Throttling";
    }

    function run() pure override public {
        revert("Only --sig 'runSimulateFromMultisig(string memory _configName)' is supported");
    }

    function runSimulateFromMultisig(string memory _configName) public {
        // sanity check we are on L1
        if (address(0x4200000000000000000000000000000000000000).code.length > 0)
            revert("L2 address has code, are you on an L1?");

        console.log("Simulating script from", msg.sender);
        console.log("Using config", _configName);
        configName = _configName;

        applyOptimismPortalThrottle();

        applyL1StandardBridgeEthThrottle();
        applyL1StandardBridgeErc20Throttle();
    }

    function applyL1StandardBridgeEthThrottle() broadcast public {
        ThrottlingConfig.ThrottleConfig memory config = ThrottlingConfig.ethDepositThrottleL1StandardBridgeConfig(configName);
        L1StandardBridge l1StandardBridge = L1StandardBridge(mustGetAddress("L1StandardBridgeProxy"));
        (uint208 maxAmountPerPeriod, uint48 periodLength, ) = l1StandardBridge.ethThrottleDeposits();
        if (maxAmountPerPeriod != config.maxAmountPerPeriod) {
            l1StandardBridge.setEthThrottleDepositsMaxAmount(config.maxAmountPerPeriod, 0);
        }
        if (config.periodLength != 0 && periodLength != config.periodLength) {
            l1StandardBridge.setEthThrottleDepositsPeriodLength(config.periodLength);
        }
    }

    function applyL1StandardBridgeErc20Throttle() broadcast public {
        // TODO
    }

    function applyOptimismPortalThrottle() broadcast public {
        ThrottlingConfig.ThrottleConfig memory depositConfig;
        ThrottlingConfig.ThrottleConfig memory withdrawalConfig;
        (depositConfig, withdrawalConfig) = ThrottlingConfig.ethThrottleConfig(configName);

        OptimismPortal portal = OptimismPortal(mustGetAddress("OptimismPortalProxy"));
        // deposit throttle
        {
            {
                (uint208 maxAmountPerPeriod, , uint256 maxAmountTotal) = portal.ethThrottleDeposits();
                if (maxAmountPerPeriod != depositConfig.maxAmountPerPeriod || maxAmountTotal != depositConfig.maxAmountTotal) {
                    portal.setEthThrottleDepositsMaxAmount(depositConfig.maxAmountPerPeriod, depositConfig.maxAmountTotal);
                }
            }
            {
                (, uint48 periodLength,) = portal.ethThrottleDeposits();
                if (depositConfig.periodLength != 0 && periodLength != depositConfig.periodLength) {
                    portal.setEthThrottleDepositsPeriodLength(depositConfig.periodLength);
                }
            }
        }
        // withdrawal throttle
        {
            {
                (uint208 maxAmountPerPeriod, , uint256 maxAmountTotal) = portal.ethThrottleWithdrawals();
                if (maxAmountPerPeriod != withdrawalConfig.maxAmountPerPeriod || maxAmountTotal != withdrawalConfig.maxAmountTotal) {
                    portal.setEthThrottleWithdrawalsMaxAmount(withdrawalConfig.maxAmountPerPeriod, withdrawalConfig.maxAmountTotal);
                }
            }
            {
                (, uint48 periodLength,) = portal.ethThrottleWithdrawals();
                if (withdrawalConfig.periodLength != 0 && periodLength != withdrawalConfig.periodLength) {
                    portal.setEthThrottleWithdrawalsPeriodLength(withdrawalConfig.periodLength);
                }
            }
        }
    }
}
