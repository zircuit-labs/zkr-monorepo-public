// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { Test } from "forge-std/Test.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { SimpleStorage } from "test/universal/Proxy.t.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

contract ProxyAdmin_Test is Test {
    address alice = address(64);

    Proxy proxy;

    ProxyAdmin admin;

    SimpleStorage implementation;

    function setUp() external {
        // Deploy the proxy admin
        admin = new ProxyAdmin(alice);
        // Deploy the standard proxy
        proxy = new Proxy(address(admin));

        // Impersonate alice for setting up the admin.
        vm.startPrank(alice);

        // Set the proxy types
        admin.setProxyType(address(proxy), ProxyAdmin.ProxyType.ERC1967);
        vm.stopPrank();

        implementation = new SimpleStorage();
    }

    function test_owner_succeeds() external view {
        assertEq(admin.owner(), alice);
    }

    function test_proxyType_succeeds() external view {
        assertEq(uint256(admin.proxyType(address(proxy))), uint256(ProxyAdmin.ProxyType.ERC1967));
    }

    function test_erc1967GetProxyImplementation_succeeds() external {
        getProxyImplementation(payable(proxy));
    }

    function getProxyImplementation(address payable _proxy) internal {
        {
            address impl = admin.getProxyImplementation(_proxy);
            assertEq(impl, address(0));
        }

        vm.prank(alice);
        admin.upgrade(_proxy, address(implementation));

        {
            address impl = admin.getProxyImplementation(_proxy);
            assertEq(impl, address(implementation));
        }
    }

    function test_erc1967GetProxyAdmin_succeeds() external view {
        getProxyAdmin(payable(proxy));
    }

    function getProxyAdmin(address payable _proxy) internal view {
        address owner = admin.getProxyAdmin(_proxy);
        assertEq(owner, address(admin));
    }

    function test_erc1967ChangeProxyAdmin_succeeds() external {
        changeProxyAdmin(payable(proxy));
    }

    function changeProxyAdmin(address payable _proxy) internal {
        ProxyAdmin.ProxyType proxyType = admin.proxyType(address(_proxy));

        vm.prank(alice);
        admin.changeProxyAdmin(_proxy, address(128));

        // The proxy is no longer the admin and can
        // no longer call the proxy interface
        if (proxyType == ProxyAdmin.ProxyType.ERC1967) {
            vm.expectRevert("Proxy: implementation not initialized");
            admin.getProxyAdmin(_proxy);
        } else {
            vm.expectRevert("ProxyAdmin: unknown proxy type");
        }

        // Call the proxy contract directly to get the admin.
        // Different proxy types have different interfaces.
        vm.prank(address(128));
        if (proxyType == ProxyAdmin.ProxyType.ERC1967) {
            assertEq(Proxy(payable(_proxy)).admin(), address(128));
        } else {
            assert(false);
        }
    }

    function test_erc1967Upgrade_succeeds() external {
        upgrade(payable(proxy));
    }

    function upgrade(address payable _proxy) internal {
        vm.prank(alice);
        admin.upgrade(_proxy, address(implementation));

        address impl = admin.getProxyImplementation(_proxy);
        assertEq(impl, address(implementation));
    }

    function test_erc1967UpgradeAndCall_succeeds() external {
        upgradeAndCall(payable(proxy));
    }

    function upgradeAndCall(address payable _proxy) internal {
        vm.prank(alice);
        admin.upgradeAndCall(_proxy, address(implementation), abi.encodeWithSelector(SimpleStorage.set.selector, 1, 1));

        address impl = admin.getProxyImplementation(_proxy);
        assertEq(impl, address(implementation));

        uint256 got = SimpleStorage(address(_proxy)).get(1);
        assertEq(got, 1);
    }

    function test_onlyOwner_notOwner_reverts() external {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, address(this)));
        admin.changeProxyAdmin(payable(proxy), address(0));

        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, address(this)));
        admin.upgrade(payable(proxy), address(implementation));

        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, address(this)));
        admin.upgradeAndCall(payable(proxy), address(implementation), hex"");
    }
}
