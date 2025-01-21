// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { StorageSlot } from "scripts/Deployer.sol";
import { Deploy } from "scripts/Deploy.s.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { StorageSetter } from "src/universal/StorageSetter.sol";
import { ChainAssertions } from "scripts/ChainAssertions.sol";
import { Proxy } from "src/universal/Proxy.sol";

/// @notice script to upgrade both the L2OutputOracle and the Verifier contract
contract L2OOVerifierDeploy is Deploy {
    address storageSetter;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "104-L2OOVerifier";
    }

    function run() override public {
        console.log("Deploying verifier from", msg.sender);

        deployL2OutputOracle();
        upgradeL2OutputOracle();

        deployVerifier();
        upgradeToVerifier();
    }

    function upgradeL2OutputOracle() broadcast internal {
        console.log("Upgrading L2OutputOracle");
        address proxyToUpgrade = mustGetAddress("L2OutputOracleProxy");
        address implementation = mustGetAddress("L2OutputOracle");
        console.log("Upgrading", proxyToUpgrade, "to", implementation);

        _upgradeViaSafe({
            _proxy: payable(proxyToUpgrade),
            _implementation: implementation
        });
    }

    function upgradeToVerifier() broadcast internal {
        console.log("Upgrading verifier");
        address proxyToUpgrade = mustGetAddress("VerifierProxy");
        address implementation = mustGetAddress("Verifier");
        console.log("Upgrading", proxyToUpgrade, "to", implementation);

        _upgradeViaSafe({
            _proxy: payable(proxyToUpgrade),
            _implementation: implementation
        });
    }
}
