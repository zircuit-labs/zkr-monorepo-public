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

/// @notice script to upgrade both the L2OutputOracle and the OptimismPortal contract
contract L2OOPortalDeploy is Deploy {
    address storageSetter;

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "105-L2OOPortal";
    }

    function run() override public {
        console.log("Deploying L2OO and from", msg.sender);

        deployL2OutputOracle();
        upgradeL2OutputOracle();

        // setFinalizationPeriod through proxy so it's not 0 after upgrade
        address implementation = mustGetAddress("L2OutputOracleProxy"); // this is  0x3883D42E90B204F5a5e2f072f414e68c2C944698
        L2OutputOracle l2OO = L2OutputOracle(implementation);
        uint256 period = cfg.finalizationPeriodSeconds();
        vm.broadcast();
        l2OO.setFinalizationPeriodSeconds(period);

        deployOptimismPortal();
        upgradePortal();
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
