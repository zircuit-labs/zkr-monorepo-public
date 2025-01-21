// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { console } from "forge-std/console.sol";
import { Script } from "forge-std/Script.sol";
import { ProxyAdmin } from "../src/universal/ProxyAdmin.sol";

/**
 * @title RecoverEther
 * @notice A script to recover eth from from a proxy.
 *         The usage is as follows:
 *         $ PROXY_ADMIN=0x... PROXY_ADDRESS=0x... \
 *           forge script scripts/RecoverEther.s.sol:RecoverEther -vvvv \
 *             --rpc-url $ETH_RPC_URL \
 *             --mnemonics $PATH_TO_MNEMONIC \
 *             --sender $SENDER_ADDR \
 * PROXY_ADMIN is the address of the ProxyAdmin contract that owns the proxy.
 * PROXY_ADDRESS is the address of the proxy with the funds that should be recovered.
 * SENDER_ADDR must be set to the owner of the ProxyAdmin contract and the mnemonic
 *   (+optionally a --mnemonic-indexes argument) must correspond to the sender. This
 *   should imo happen automatically but foundry unfortunately does not.
 * Unless executed with `--broadcast`, the execution will only happen locally and
 * print out a trace of the run. Verify in that trace that the eth is being recovered
 * to the correct address as expected. Once everything looks good, execute with
 * `--broadcast` to actually recover the funds.
 */

contract SendBalance {
    // address the funds will be sent to
    address payable constant RECOVER_ADDR = payable(0xAa8032cC6F8aA59651742878Ca017d4d68976E67);

    function recover() external {
        (bool success, ) = RECOVER_ADDR.call{value: address(this).balance}("");
        require(success, "send failed");
    }
}

contract RecoverEther is Script {

    /**
     * @notice The entrypoint function. Determines which FeeVaults can be withdrawn from and then
     *         will send the transaction via Multicall3 to withdraw all FeeVaults.
     */
    function run() external {
        ProxyAdmin admin = ProxyAdmin(vm.envAddress("PROXY_ADMIN"));

        require(admin.owner() == msg.sender, "sender not owner of ProxyAdmin. Make sure to specify --sender");

        address proxy = vm.envAddress("PROXY_ADDRESS");
        try admin.getProxyImplementation(proxy) returns (address proxyImpl) {
            console.log("Old impl address: ", proxyImpl);
        } catch {
            revert("PROXY_ADDRESS not set or not known by PROXY_ADMIN");
        }

        vm.startBroadcast();
        SendBalance impl = new SendBalance{salt: bytes32("adsiuadshfiuadshfiuahs")}();
        admin.upgradeAndCall(payable(proxy), address(impl), abi.encodeWithSelector(SendBalance.recover.selector));
        vm.stopBroadcast();
    }
}
