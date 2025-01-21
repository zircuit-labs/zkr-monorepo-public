// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { StorageSlot } from "scripts/Deployer.sol";
import { Deploy } from "scripts/Deploy.s.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { StorageSetter } from "src/universal/StorageSetter.sol";
import { ChainAssertions } from "scripts/ChainAssertions.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { L2ToL1MessagePasser } from "src/L2/L2ToL1MessagePasser.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

/// @notice Script that upgrades the L2ToL1MessagePasser
contract MessagePasserUpgrade is Deploy {
    address storageSetter;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "201-MessagePasser";
    }

    function run() override public {
        console.log("Running", name(), "from", msg.sender);

        deployL2ToL1MessagePasser();
        upgradeL2ToL1MessagePasser();
    }

    function upgradeL2ToL1MessagePasser() broadcast internal {
        console.log("Upgrading L2ToL1MessagePasser");
        address proxyToUpgrade = Predeploys.L2_TO_L1_MESSAGE_PASSER;
        address implementation = mustGetAddress("L2ToL1MessagePasser");
        console.log("Upgrading", proxyToUpgrade, "to", implementation);
        console.log("Using", msg.sender, "as admin for the L2ToL1MessagePasser contract");

        ProxyAdmin pa = ProxyAdmin(Predeploys.PROXY_ADMIN);
        pa.upgradeAndCall({
            _proxy: payable(proxyToUpgrade),
            _implementation: implementation,
            _data: abi.encodeCall(L2ToL1MessagePasser.initialize, ())
        });
    }

    /// @notice Deploy the Verifier
    function deployL2ToL1MessagePasser() public broadcast returns (address addr_) {
        console.log("Deploying L2ToL1MessagePasser implementation");

        addr_ = address(new L2ToL1MessagePasser());

        save("L2ToL1MessagePasser", addr_);
        console.log("L2ToL1MessagePasser deployed at %s", addr_);
    }
}
