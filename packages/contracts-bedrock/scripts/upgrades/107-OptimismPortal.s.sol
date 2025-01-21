// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";


import { StorageSlot } from "scripts/Deployer.sol";
import { Deploy } from "scripts/Deploy.s.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { StorageSetter } from "src/universal/StorageSetter.sol";
import { ChainAssertions } from "scripts/ChainAssertions.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { L2OutputOracle } from "src/L1/L2OutputOracle.sol";

/// @notice script to upgrade the OptimismPortal contract
contract OptimismPortalDeploy is Deploy {
    address storageSetter;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "107-OptimismPortal";
    }

    function run() override public {
        deployOptimismPortal();
        upgradePortal();
    }

    function upgradePortal() broadcast internal {
        console.log("Upgrading OptimismPortal");
        address proxyToUpgrade = mustGetAddress("OptimismPortalProxy");
        address implementation = mustGetAddress("OptimismPortal");
        console.log("Upgrading", proxyToUpgrade, "to", implementation);

        _upgradeViaSafe({
            _proxy: payable(proxyToUpgrade),
            _implementation: implementation
        });
    }
}
