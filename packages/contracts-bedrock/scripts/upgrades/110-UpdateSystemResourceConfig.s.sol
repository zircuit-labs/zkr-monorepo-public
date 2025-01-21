// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {console2 as console} from 'forge-std/console2.sol';

import {L1UpgradeUtils} from 'scripts/upgrades/UpgradeUtils.sol';
import {SystemConfig} from 'src/L1/SystemConfig.sol';
import {ResourceMetering} from 'src/L1/ResourceMetering.sol';

/// @notice ad hoc script to upgrade the verifier
contract UpdateSystemResourceConfig is L1UpgradeUtils {
  uint32 public constant newGasLimit = 7000000;

  function name() public pure virtual override returns (string memory name_) {
    name_ = '110-UpdateSystemResourceConfig';
  }

  function run() public pure override {
    revert("Use --sig with 'runSimulateFromMultisig()'");
  }

  function runSimulateFromMultisig() public broadcast {
    console.log('Simulating transaction from', msg.sender);

    SystemConfig systemConfigProxy = SystemConfig(
      mustGetAddress('SystemConfigProxy')
    );
    ResourceMetering.ResourceConfig memory newConfig = systemConfigProxy
      .resourceConfig();
    console.log('Current config:');
    console.log(
      'baseFeeMaxChangeDenominator',
      newConfig.baseFeeMaxChangeDenominator
    );
    console.log('elasticityMultiplier', newConfig.elasticityMultiplier);
    console.log('maximumBaseFee', newConfig.maximumBaseFee);
    console.log('maxResourceLimit', newConfig.maxResourceLimit);
    console.log('maxTransactionLimit', newConfig.maxTransactionLimit);
    console.log('minimumBaseFee', newConfig.minimumBaseFee);
    console.log('systemTxMaxGas', newConfig.systemTxMaxGas);
    console.log('');

    newConfig.maxResourceLimit = newGasLimit;

    console.log('Changing max resource limit value.');
    systemConfigProxy.setResourceConfig(newConfig);

    require(
      systemConfigProxy.resourceConfig().maxResourceLimit == newGasLimit,
      'resourcelimit not set'
    );

    console.log('Successfully changed max resource limit value.');
  }
}
