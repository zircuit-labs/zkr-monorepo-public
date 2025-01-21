// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { Deploy } from "scripts/Deploy.s.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { L1Block } from "src/L2/L1Block.sol";

/// @notice Upgrade the L1Block contract
contract UpgradeL1Block is Deploy {

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "202-UpgradeL1Block";
    }

    function run() override public {
        address l1Block = deployL1Block();
        upgradeL1Block(l1Block);
    }

    /// @notice Deploy the L1Block contract
    function deployL1Block() public broadcast returns (address addr_) {
        console.log("Deploying L1Block implementation");
        L1Block deployedAddr = new L1Block();

        console.log("L1Block deployed at %s", address(deployedAddr));

        addr_ = address(deployedAddr);
    }

    function upgradeL1Block(address implementation) broadcast internal {
        console.log("Upgrading L1Block");
        address proxyToUpgrade = Predeploys.L1_BLOCK_ATTRIBUTES;
        console.log("Upgrading", proxyToUpgrade, "to", implementation);

        ProxyAdmin proxyAdmin = ProxyAdmin(Predeploys.PROXY_ADMIN);
        proxyAdmin.upgrade({_proxy: payable(proxyToUpgrade), _implementation: implementation});
    }
}

