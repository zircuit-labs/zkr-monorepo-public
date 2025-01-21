// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { Deploy } from "scripts/Deploy.s.sol";
import { ThrottlingConfig } from "scripts/throttling/ThrottlingConfig.sol";
import { L2StandardBridge } from "src/L2/L2StandardBridge.sol";
import { L2ToL1MessagePasser } from "src/L2/L2ToL1MessagePasser.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

/// @notice ad hoc script to upgrade the verifier
contract L2ThrottlingScript is Deploy {
    string configName;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "ApplyL2Throttling";
    }

    function run() pure override public {
        revert("Only --sig 'runSimulateFromMultisig(string memory _configName)' is supported");
    }

    function runSimulateFromMultisig(string memory _configName) public {
        // sanity check we are on L1
        if (address(0x4200000000000000000000000000000000000000).code.length == 0)
            revert("L2 address has no code, are you on an L2?");

        console.log("Simulating script from", msg.sender);
        console.log("Using config", _configName);
        configName = _configName;

        applyL2ToL1MessagePasserThrottle();

        applyL2StandardBridgeErc20Throttle();
    }

    function applyL2StandardBridgeErc20Throttle() broadcast public {
        // TODO
    }

    function applyL2ToL1MessagePasserThrottle() broadcast public {
        ThrottlingConfig.ThrottleConfig memory withdrawalConfig;
        (, withdrawalConfig) = ThrottlingConfig.ethThrottleConfig(configName);

        L2ToL1MessagePasser messagePasser = L2ToL1MessagePasser(payable(Predeploys.L2_TO_L1_MESSAGE_PASSER));
        // withdrawal throttle
        {
            {
                (uint208 maxAmountPerPeriod, , uint256 maxAmountTotal) = messagePasser.ethThrottleWithdrawals();
                if (maxAmountPerPeriod != withdrawalConfig.maxAmountPerPeriod || maxAmountTotal != withdrawalConfig.maxAmountTotal) {
                    messagePasser.setEthThrottleWithdrawalsMaxAmount(withdrawalConfig.maxAmountPerPeriod, withdrawalConfig.maxAmountTotal);
                }
            }
            {
                (, uint48 periodLength,) = messagePasser.ethThrottleWithdrawals();
                if (withdrawalConfig.periodLength != 0 && periodLength != withdrawalConfig.periodLength) {
                    messagePasser.setEthThrottleWithdrawalsPeriodLength(withdrawalConfig.periodLength);
                }
            }
        }
    }
}
