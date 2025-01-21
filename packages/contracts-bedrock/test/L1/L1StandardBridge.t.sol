// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

// Testing utilities
import { stdStorage, StdStorage } from "forge-std/Test.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { Bridge_Initializer } from "test/setup/Bridge_Initializer.sol";

// Libraries
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Constants } from "src/libraries/Constants.sol";

// Target contract dependencies
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { L2StandardBridge } from "src/L2/L2StandardBridge.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { SuperchainConfig } from "src/L1/SuperchainConfig.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { AccessControlPausable } from "src/universal/AccessControlPausable.sol";
import { TransferThrottleTest } from "test/universal/TransferThrottle.t.sol";
import { ResourceMetering } from "src/L1/ResourceMetering.sol";

// Target contract
import { OptimismPortal } from "src/L1/OptimismPortal.sol";

contract L1StandardBridge_Getter_Test is Bridge_Initializer {
    /// @dev Test that the accessors return the correct initialized values.
    function test_getters_succeeds() external view {
        assert(l1StandardBridge.OTHER_BRIDGE() == l2StandardBridge);
        assert(l1StandardBridge.messenger() == l1CrossDomainMessenger);
        assert(l1StandardBridge.MESSENGER() == l1CrossDomainMessenger);
    }
}

contract L1StandardBridge_Initialize_Test is Bridge_Initializer {
    /// @dev Test that the initialize function sets the correct values.
    function test_initialize_succeeds() external view {
        assertEq(address(l1StandardBridge.messenger()), address(l1CrossDomainMessenger));
        assertEq(address(l1StandardBridge.OTHER_BRIDGE()), Predeploys.L2_STANDARD_BRIDGE);
        assertEq(address(l2StandardBridge), Predeploys.L2_STANDARD_BRIDGE);
    }
}

contract L1StandardBridge_Pause_Test is Bridge_Initializer {
    /// @dev Verifies that the `paused` accessor returns the same value as the `paused` function of the
    ///      `superchainConfig`.
    function test_paused_succeeds() external view {
        assertEq(l1StandardBridge.paused(), superchainConfig.paused());
    }

    /// @dev Ensures that the `paused` function of the bridge contract actually calls the `paused` function of the
    ///      `superchainConfig`.
    function test_pause_callsSuperchainConfig_succeeds() external {
        vm.expectCall(address(superchainConfig), abi.encodeWithSelector(AccessControlPausable.paused.selector));
        l1StandardBridge.paused();
    }

    /// @dev Checks that the `paused` state of the bridge matches the `paused` state of the `superchainConfig` after
    ///      it's been changed.
    function test_pause_matchesSuperchainConfig_succeeds() external {
        assertFalse(l1StandardBridge.paused());
        assertEq(l1StandardBridge.paused(), superchainConfig.paused());

        vm.prank(superchainConfig.guardian());
        superchainConfig.pause("identifier");

        assertTrue(l1StandardBridge.paused());
        assertEq(l1StandardBridge.paused(), superchainConfig.paused());
    }
}

contract L1StandardBridge_Pause_TestFail is Bridge_Initializer {
    /// @dev Sets up the test by pausing the bridge, giving ether to the bridge and mocking
    ///      the calls to the xDomainMessageSender so that it returns the correct value.
    function setUp() public override {
        super.setUp();
        vm.prank(superchainConfig.guardian());
        superchainConfig.pause("identifier");
        assertTrue(l1StandardBridge.paused());

        vm.deal(address(l1StandardBridge.messenger()), 1 ether);

        vm.mockCall(
            address(l1StandardBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1StandardBridge.otherBridge()))
        );
    }

    /// @dev Confirms that the `finalizeBridgeETH` function reverts when the bridge is paused.
    function test_pause_finalizeBridgeETH_reverts() external {
        vm.prank(address(l1StandardBridge.messenger()));
        vm.expectRevert("StandardBridge: paused");
        l1StandardBridge.finalizeBridgeETH{ value: 100 }({
            _from: address(2),
            _to: address(3),
            _amount: 100,
            _extraData: hex""
        });
    }

    /// @dev Confirms that the `finalizeBridgeERC20` function reverts when the bridge is paused.
    function test_pause_finalizeBridgeERC20_reverts() external {
        vm.prank(address(l1StandardBridge.messenger()));
        vm.expectRevert("StandardBridge: paused");
        l1StandardBridge.finalizeBridgeERC20({
            _localToken: address(0),
            _remoteToken: address(0),
            _from: address(0),
            _to: address(0),
            _amount: 0,
            _extraData: hex""
        });
    }
}

contract L1StandardBridge_Initialize_TestFail is Bridge_Initializer { }

contract L1StandardBridge_Receive_Test is Bridge_Initializer {
    /// @dev Tests receive bridges ETH successfully.
    function test_receive_succeeds() external {
        assertEq(address(optimismPortal).balance, 0);

        vm.expectEmit(address(l1StandardBridge));
        emit ETHBridgeInitiated(alice, alice, 100, hex"");

        vm.expectCall(
        address(l1CrossDomainMessenger),
        abi.encodeWithSelector(
        CrossDomainMessenger.sendMessage.selector,
        address(l2StandardBridge),
        abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, alice, 100, hex""),
        200_000
        )
        );

        vm.prank(alice, alice);
        (bool success,) = address(l1StandardBridge).call{ value: 100 }(hex"");
        assertEq(success, true);
        assertEq(address(optimismPortal).balance, 100);
    }
}

contract L1StandardBridge_Receive_TestFail { }

contract PreBridgeETH is Bridge_Initializer {
    /// @dev Asserts the expected calls and events for bridging ETH depending
    ///      on whether the bridge call is legacy or not.
    function _preBridgeETH() internal {
        assertEq(address(optimismPortal).balance, 0);
        uint256 nonce = l1CrossDomainMessenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));

        bytes memory message =
                            abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, alice, 500, hex"dead");

        vm.expectCall(
            address(l1StandardBridge),
            500,
            abi.encodeWithSelector(l1StandardBridge.bridgeETH.selector, 50000, hex"dead")
        );
        vm.expectCall(
            address(l1CrossDomainMessenger),
            500,
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l2StandardBridge), message, 50000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l1StandardBridge),
            address(l2StandardBridge),
            500,
            50000,
            message
        );

        uint64 baseGas = l1CrossDomainMessenger.baseGas(message, 50000);
        vm.expectCall(
            address(optimismPortal),
            500,
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector,
                address(l2CrossDomainMessenger),
                500,
                baseGas,
                false,
                innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(500), uint256(500), baseGas, false, innerMessage);

        vm.expectEmit(address(l1StandardBridge));
        emit ETHBridgeInitiated(alice, alice, 500, hex"dead");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(address(optimismPortal));
        emit TransactionDeposited(l1MessengerAliased, address(l2CrossDomainMessenger), version, opaqueData);

        vm.expectEmit(true, false, false, false);
        emit GasBurned(0 , alice); // the gas amount is ignored

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(address(l1CrossDomainMessenger));
        emit SentMessage(address(l2StandardBridge), address(l1StandardBridge), message, nonce, 50000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(address(l1CrossDomainMessenger));
        emit SentMessageExtension1(address(l1StandardBridge), 500);

        vm.prank(alice, alice);
    }
}

contract L1StandardBridge_BridgeETH_Test is PreBridgeETH {
    /// @dev Tests that bridging ETH succeeds.
    ///      Emits ETHDepositInitiated and ETHBridgeInitiated events.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call bridgeETH.
    ///      ETH ends up in the optimismPortal.
    function test_bridgeETH_succeeds() external {
        _preBridgeETH();
        l1StandardBridge.bridgeETH{ value: 500 }(50000, hex"dead");
        assertEq(address(optimismPortal).balance, 500);
    }
}

contract PreBridgeETHTo is Bridge_Initializer {
    /// @dev Asserts the expected calls and events for bridging ETH to a different
    ///      address depending on whether the bridge call is legacy or not.
    function _preBridgeETHTo() internal {
        assertEq(address(optimismPortal).balance, 0);
        uint256 nonce = l1CrossDomainMessenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));

        vm.expectCall(
            address(l1StandardBridge),
            600,
            abi.encodeWithSelector(l1StandardBridge.bridgeETHTo.selector, bob, 60000, hex"dead")
        );

        bytes memory message =
                            abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, bob, 600, hex"dead");

        // the L1 bridge should call
        // L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l2StandardBridge), message, 60000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l1StandardBridge),
            address(l2StandardBridge),
            600,
            60000,
            message
        );

        uint64 baseGas = l1CrossDomainMessenger.baseGas(message, 60000);
        vm.expectCall(
            address(optimismPortal),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector,
                address(l2CrossDomainMessenger),
                600,
                baseGas,
                false,
                innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(600), uint256(600), baseGas, false, innerMessage);

        vm.expectEmit(address(l1StandardBridge));
        emit ETHBridgeInitiated(alice, bob, 600, hex"dead");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(address(optimismPortal));
        emit TransactionDeposited(l1MessengerAliased, address(l2CrossDomainMessenger), version, opaqueData);

        vm.expectEmit(true, false, false, false);
        emit GasBurned(0 , alice); // the gas amount is ignored

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(address(l1CrossDomainMessenger));
        emit SentMessage(address(l2StandardBridge), address(l1StandardBridge), message, nonce, 60000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(address(l1CrossDomainMessenger));
        emit SentMessageExtension1(address(l1StandardBridge), 600);

        // deposit eth to bob
        vm.prank(alice, alice);
    }
}

contract L1StandardBridge_BridgeETHTo_Test is PreBridgeETHTo {
    /// @dev Tests that bridging ETH to a different address succeeds.
    ///      Emits ETHDepositInitiated and ETHBridgeInitiated events.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call bridgeETHTo.
    ///      ETH ends up in the optimismPortal.
    function test_bridgeETHTo_succeeds() external {
        _preBridgeETHTo();
        l1StandardBridge.bridgeETHTo{ value: 600 }(bob, 60000, hex"dead");
        assertEq(address(optimismPortal).balance, 600);
    }
}

contract L1StandardBridge_BridgeERC20_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    // bridgeERC20
    // - updates bridge.deposits
    // - emits ERC20BridgeInitiated
    // - calls optimismPortal.depositTransaction
    // - only callable by EOA

    /// @dev Tests that depositing ERC20 to the bridge succeeds.
    ///      Bridge deposits are updated.
    ///      Emits ERC20BridgeInitiated event.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call depositERC20.
    function test_depositERC20_succeeds() external {
        uint256 nonce = l1CrossDomainMessenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));

        // Deal Alice's ERC20 State
        deal(address(L1Token), alice, 100000, true);
        vm.prank(alice);
        L1Token.approve(address(l1StandardBridge), type(uint256).max);

        // The l1StandardBridge should transfer alice's tokens to itself
        vm.expectCall(
            address(L1Token), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(l1StandardBridge), 100)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(L2Token), address(L1Token), alice, alice, 100, hex""
        );

        // the L1 bridge should call L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l2StandardBridge), message, 10000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l1StandardBridge),
            address(l2StandardBridge),
            0,
            10000,
            message
        );

        uint64 baseGas = l1CrossDomainMessenger.baseGas(message, 10000);
        vm.expectCall(
            address(optimismPortal),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector,
                address(l2CrossDomainMessenger),
                0,
                baseGas,
                false,
                innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        vm.expectEmit(address(l1StandardBridge));
        emit ERC20BridgeInitiated(address(L1Token), address(L2Token), alice, alice, 100, hex"");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(address(optimismPortal));
        emit TransactionDeposited(l1MessengerAliased, address(l2CrossDomainMessenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(address(l1CrossDomainMessenger));
        emit SentMessage(address(l2StandardBridge), address(l1StandardBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(address(l1CrossDomainMessenger));
        emit SentMessageExtension1(address(l1StandardBridge), 0);

        vm.prank(alice);
        l1StandardBridge.bridgeERC20(address(L1Token), address(L2Token), 100, 10000, hex"");
        assertEq(l1StandardBridge.deposits(address(L1Token), address(L2Token)), 100);
    }
}

contract L1StandardBridge_BridgeERC20_TestFail is Bridge_Initializer {
    /// @dev Tests that depositing an ERC20 to the bridge reverts
    ///      if the caller is not an EOA.
    function test_depositERC20_notEoa_reverts() external {
        // turn alice into a contract
        vm.etch(alice, hex"ffff");

        vm.expectRevert("StandardBridge: function can only be called from an EOA");
        vm.prank(alice, alice);
        l1StandardBridge.bridgeERC20(address(0), address(0), 100, 100, hex"");
    }
}

contract L1StandardBridge_BridgeERC20To_Test is Bridge_Initializer {
    /// @dev Tests that depositing ERC20 to the bridge succeeds when
    ///      sent to a different address.
    ///      Bridge deposits are updated.
    ///      Emits ERC20BridgeInitiated event.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Contracts can call depositERC20.
    function test_depositERC20To_succeeds() external {
        uint256 nonce = l1CrossDomainMessenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(L2Token), address(L1Token), alice, bob, 1000, hex""
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l1StandardBridge),
            address(l2StandardBridge),
            0,
            10000,
            message
        );

        uint64 baseGas = l1CrossDomainMessenger.baseGas(message, 10000);
        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        deal(address(L1Token), alice, 100000, true);

        vm.prank(alice);
        L1Token.approve(address(l1StandardBridge), type(uint256).max);

        vm.expectEmit(address(l1StandardBridge));
        emit ERC20BridgeInitiated(address(L1Token), address(L2Token), alice, bob, 1000, hex"");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(address(optimismPortal));
        emit TransactionDeposited(l1MessengerAliased, address(l2CrossDomainMessenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(address(l1CrossDomainMessenger));
        emit SentMessage(address(l2StandardBridge), address(l1StandardBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(address(l1CrossDomainMessenger));
        emit SentMessageExtension1(address(l1StandardBridge), 0);

        // the L1 bridge should call L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l2StandardBridge), message, 10000)
        );
        // The L1 XDM should call OptimismPortal.depositTransaction
        vm.expectCall(
            address(optimismPortal),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector,
                address(l2CrossDomainMessenger),
                0,
                baseGas,
                false,
                innerMessage
            )
        );
        vm.expectCall(
            address(L1Token),
            abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(l1StandardBridge), 1000)
        );

        vm.prank(alice);
        l1StandardBridge.bridgeERC20To(address(L1Token), address(L2Token), bob, 1000, 10000, hex"");

        assertEq(l1StandardBridge.deposits(address(L1Token), address(L2Token)), 1000);
    }
}

contract L1StandardBridge_FinalizeERC20Withdrawal_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    /// @dev Tests that finalizing an ERC20 withdrawal succeeds.
    ///      Bridge deposits are updated.
    ///      Emits ERC20WithdrawalFinalized event.
    ///      Only callable by the L2 bridge.
    function test_finalizeERC20Withdrawal_succeeds() external {
        deal(address(L1Token), address(l1StandardBridge), 100, true);

        uint256 slot = stdstore.target(address(l1StandardBridge)).sig("deposits(address,address)").with_key(
            address(L1Token)
        ).with_key(address(L2Token)).find();

        // Give the L1 bridge some ERC20 tokens
        vm.store(address(l1StandardBridge), bytes32(slot), bytes32(uint256(100)));
        assertEq(l1StandardBridge.deposits(address(L1Token), address(L2Token)), 100);

        vm.expectEmit(address(l1StandardBridge));
        emit ERC20BridgeFinalized(address(L1Token), address(L2Token), alice, alice, 100, hex"");

        vm.expectCall(address(L1Token), abi.encodeWithSelector(ERC20.transfer.selector, alice, 100));

        vm.mockCall(
            address(l1StandardBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1StandardBridge.OTHER_BRIDGE()))
        );
        vm.prank(address(l1StandardBridge.messenger()));
        l1StandardBridge.finalizeBridgeERC20(address(L1Token), address(L2Token), alice, alice, 100, hex"");

        assertEq(L1Token.balanceOf(address(l1StandardBridge)), 0);
        assertEq(L1Token.balanceOf(address(alice)), 100);
    }
}

contract L1StandardBridge_FinalizeERC20Withdrawal_TestFail is Bridge_Initializer {
    /// @dev Tests that finalizing an ERC20 withdrawal reverts if the caller is not the L2 bridge.
    function test_finalizeERC20Withdrawal_notMessenger_reverts() external {
        vm.mockCall(
            address(l1StandardBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1StandardBridge.OTHER_BRIDGE()))
        );
        vm.prank(address(28));
        vm.expectRevert("StandardBridge: function can only be called from the other bridge");
        l1StandardBridge.finalizeBridgeERC20(address(L1Token), address(L2Token), alice, alice, 100, hex"");
    }

    /// @dev Tests that finalizing an ERC20 withdrawal reverts if the caller is not the L2 bridge.
    function test_finalizeERC20Withdrawal_notOtherBridge_reverts() external {
        vm.mockCall(
            address(l1StandardBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(address(0)))
        );
        vm.prank(address(l1StandardBridge.messenger()));
        vm.expectRevert("StandardBridge: function can only be called from the other bridge");
        l1StandardBridge.finalizeBridgeERC20(address(L1Token), address(L2Token), alice, alice, 100, hex"");
    }
}

contract L1StandardBridge_FinalizeBridgeETH_Test is Bridge_Initializer {
    /// @dev Tests that finalizing bridged ETH succeeds.
    function test_finalizeBridgeETH_succeeds() external {
        address messenger = address(l1StandardBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1StandardBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);

        vm.expectEmit(address(l1StandardBridge));
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        l1StandardBridge.finalizeBridgeETH{ value: 100 }(alice, alice, 100, hex"");
    }
}

contract L1StandardBridge_FinalizeBridgeETH_TestFail is Bridge_Initializer {
    /// @dev Tests that finalizing bridged ETH reverts if the amount is incorrect.
    function test_finalizeBridgeETH_incorrectValue_reverts() external {
        address messenger = address(l1StandardBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1StandardBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);
        vm.expectRevert("StandardBridge: amount sent does not match amount required");
        l1StandardBridge.finalizeBridgeETH{ value: 50 }(alice, alice, 100, hex"");
    }

    /// @dev Tests that finalizing bridged ETH reverts if the destination is the L1 bridge.
    function test_finalizeBridgeETH_sendToSelf_reverts() external {
        address messenger = address(l1StandardBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1StandardBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);
        vm.expectRevert("StandardBridge: cannot send to self");
        l1StandardBridge.finalizeBridgeETH{ value: 100 }(alice, address(l1StandardBridge), 100, hex"");
    }

    /// @dev Tests that finalizing bridged ETH reverts if the destination is the messenger.
    function test_finalizeBridgeETH_sendToMessenger_reverts() external {
        address messenger = address(l1StandardBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1StandardBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);
        vm.expectRevert("StandardBridge: cannot send to messenger");
        l1StandardBridge.finalizeBridgeETH{ value: 100 }(alice, messenger, 100, hex"");
    }
}

contract L1StandardBridge_EthDepositThrottling_Test is TransferThrottleTest {
    address senderAddress;

    function contractAddress() internal view override returns (address) {
        return address(l1StandardBridge);
    }

    function contractName() internal pure override returns (string memory) {
        return "StandardBridge";
    }

    function periodLength() internal view override returns (uint48 _periodLength) {
        (, _periodLength,) = l1StandardBridge.ethThrottleDeposits();
    }

    function maxAmountPerPeriod() internal view override returns (uint208 _maxAmountPerPeriod) {
        (_maxAmountPerPeriod, ,) = l1StandardBridge.ethThrottleDeposits();
    }

    function maxAmountTotal() internal view override returns (uint256 _maxAmountTotal) {
        ( , , _maxAmountTotal) = l1StandardBridge.ethThrottleDeposits();
    }

    function setPeriodLength(uint48 _periodLength) internal override {
        l1StandardBridge.setEthThrottleDepositsPeriodLength(_periodLength);
    }

    function setMaxAmount(uint208 _maxAmountPerPeriod, uint256 _maxAmountTotal) internal override {
        l1StandardBridge.setEthThrottleDepositsMaxAmount(_maxAmountPerPeriod, _maxAmountTotal);
    }

    function _transferThrottleAsset(uint256 value) internal override {
        vm.deal(senderAddress, value);
        vm.prank(senderAddress);
        l1StandardBridge.bridgeETH{ value: value }(50000, hex"dead");
    }

    /// @dev Setup the system for a ready-to-use state.
    function setUp() public override {
        super.setUp();

        senderAddress = alice;

        // set deposit limit to max so we don't get blocked
        ResourceMetering.ResourceConfig memory config = Constants.DEFAULT_RESOURCE_CONFIG();
        config.maxTransactionLimit = type(uint16).max;
        vm.prank(systemConfig.owner());
        systemConfig.setResourceConfig(config);

        // set up roles for access control testing
        vm.startPrank(superchainConfig.defaultAdmin());
        superchainConfig.grantRole(superchainConfig.MONITOR_ROLE(), monitor);
        superchainConfig.grantRole(superchainConfig.OPERATOR_ROLE(), operator);
        vm.stopPrank();

        // set up throttling parameters
        vm.prank(l1StandardBridge.accessController().defaultAdmin());
        setMaxAmount(100 ether, 0);

        vm.prank(l1StandardBridge.accessController().defaultAdmin());
        setPeriodLength(1 hours);

        resetPeriod();
    }

    function test_setThrottle_maxAmount() external override {
        throttleSetMaxAmountTestDisabled();
    }

    function test_perUser_throttle() external {
        uint256 max = maxAmountPerPeriod();

        // send max amount from alice
        senderAddress = alice;
        transferThrottleAsset(max, ThrottleRevert.None);

        // now send max amount from bob
        senderAddress = bob;
        transferThrottleAsset(max, ThrottleRevert.None);

        // try sending more and make sure it fails
        senderAddress = alice;
        transferThrottleAsset(max / 2, ThrottleRevert.Throughput);
        senderAddress = bob;
        transferThrottleAsset(max / 2, ThrottleRevert.Throughput);
    }
}

contract L1StandardBridge_ERC20DepositThrottling_Test is Bridge_Initializer, TransferThrottleTest {
    address senderAddress;

    function contractAddress() internal view override returns (address) {
        return address(l1StandardBridge);
    }

    function contractName() internal pure override returns (string memory) {
        return "StandardBridge";
    }

    function periodLength() internal view override returns (uint48 _periodLength) {
        (, _periodLength,) = l1StandardBridge.erc20ThrottleDeposits(address(L1Token));
    }

    function maxAmountPerPeriod() internal view override returns (uint208 _maxAmountPerPeriod) {
        (_maxAmountPerPeriod, ,) = l1StandardBridge.erc20ThrottleDeposits(address(L1Token));
    }

    function maxAmountTotal() internal view override returns (uint256 _maxAmountTotal) {
        ( , , _maxAmountTotal) = l1StandardBridge.erc20ThrottleDeposits(address(L1Token));
    }

    function setPeriodLength(uint48 _periodLength) internal override {
        setPeriodLength(_periodLength,  address(L1Token));
    }

    function setPeriodLength(uint48 _periodLength, address _token) internal {
        l1StandardBridge.setErc20ThrottleDepositsPeriodLength(_token, _periodLength);
    }

    function setMaxAmount(uint208 _maxAmountPerPeriod, uint256 _maxAmountTotal) internal override {
        setMaxAmount(_maxAmountPerPeriod, _maxAmountTotal, address(L1Token));
    }

    function setMaxAmount(uint208 _maxAmountPerPeriod, uint256 _maxAmountTotal, address customToken) internal {
        l1StandardBridge.setErc20ThrottleDepositsMaxAmount(customToken, _maxAmountPerPeriod, _maxAmountTotal);
    }

    function _deal(address addr, uint256 value) internal override {
        // if it's any address other than the contract itself, deal it to alice since it will be sent from that address
        deal(address(L1Token), addr != contractAddress() ? alice : addr, value, true);
    }

    function _transferThrottleAsset(uint256) internal pure override {
        revert("Not used");
    }

    function transferThrottleAsset(uint256 value, ThrottleRevert revertExpectation) internal override {
        // Deal Alice's ERC20 State
        deal(address(L1Token), senderAddress, value, true);
        vm.prank(senderAddress);
        L1Token.approve(address(l1StandardBridge), type(uint256).max);

        _transferThrottleAssetCheck(revertExpectation);

        vm.prank(senderAddress);
        l1StandardBridge.bridgeERC20(address(L1Token), address(L2Token), value, 10000, hex"");
    }

    /// @dev Setup the system for a ready-to-use state.
    function setUp() public override(Bridge_Initializer, TransferThrottleTest) {
        super.setUp();

        senderAddress = alice;

        // set deposit limit to max so we don't get blocked
        ResourceMetering.ResourceConfig memory config = Constants.DEFAULT_RESOURCE_CONFIG();
        config.maxTransactionLimit = type(uint16).max;
        vm.prank(systemConfig.owner());
        systemConfig.setResourceConfig(config);

        // set up roles for access control testing
        vm.startPrank(superchainConfig.defaultAdmin());
        superchainConfig.grantRole(superchainConfig.MONITOR_ROLE(), monitor);
        superchainConfig.grantRole(superchainConfig.OPERATOR_ROLE(), operator);
        vm.stopPrank();

        // set up throttling parameters
        vm.prank(optimismPortal.guardian());
        setMaxAmount(100 ether, 0);

        vm.prank(optimismPortal.guardian());
        setPeriodLength(1 hours);

        resetPeriod();
    }

    /// @notice Try to set the throttle config of a token that does not exist
    function test_setThrottle_nocode_reverts() external {
        address noCodeToken = address(0xdeaddead);
        require(noCodeToken.code.length == 0, "address is supposed to have no code");

        // technically doesn't matter who calls it but let's assume we would have the capability
        vm.startPrank(operator);
        vm.expectRevert(prefixContractName(": token has no code"));
        setMaxAmount(0, 0, noCodeToken);
        vm.expectRevert(prefixContractName(": token has no code"));
        setPeriodLength(0, noCodeToken);
        vm.stopPrank();
    }

    function test_setThrottle_maxAmount() external override {
        throttleSetMaxAmountTestEnabled();
    }

    function testFuzz_throttle_totalAmount(uint256 contractBalance, uint256 additionalValue, uint256 maxTotal) external {
        throttleMaxAmountTotal(contractBalance, additionalValue, maxTotal);
    }

    function test_perUser_throttle() external {
        uint256 max = maxAmountPerPeriod();

        // send max amount from alice
        senderAddress = alice;
        transferThrottleAsset(max, ThrottleRevert.None);

        // now send max amount from bob
        senderAddress = bob;
        transferThrottleAsset(max, ThrottleRevert.None);

        // try sending more and make sure it fails
        senderAddress = alice;
        transferThrottleAsset(max / 2, ThrottleRevert.Throughput);
        senderAddress = bob;
        transferThrottleAsset(max / 2, ThrottleRevert.Throughput);
    }
}

contract L1StandardBridge_WithdrawalThrottling_Test is Bridge_Initializer, TransferThrottleTest {
    using stdStorage for StdStorage;

    function contractAddress() internal view override returns (address) {
        return address(l1StandardBridge);
    }

    function contractName() internal pure override returns (string memory) {
        return "StandardBridge";
    }

    function periodLength() internal view override returns (uint48 _periodLength) {
        (, _periodLength,) = l1StandardBridge.erc20ThrottleWithdrawals(address(L1Token));
    }

    function maxAmountPerPeriod() internal view override returns (uint208 _maxAmountPerPeriod) {
        (_maxAmountPerPeriod, ,) = l1StandardBridge.erc20ThrottleWithdrawals(address(L1Token));
    }

    function maxAmountTotal() internal view override returns (uint256 _maxAmountTotal) {
        ( , , _maxAmountTotal) = l1StandardBridge.erc20ThrottleWithdrawals(address(L1Token));
    }

    function setPeriodLength(uint48 _periodLength) internal override {
        setPeriodLength(_periodLength, address(L1Token));
    }

    function setPeriodLength(uint48 _periodLength, address _token) internal {
        l1StandardBridge.setErc20ThrottleWithdrawalsPeriodLength(_token, _periodLength);
    }

    function setMaxAmount(uint208 _maxAmountPerPeriod, uint256 _maxAmountTotal) internal override {
        setMaxAmount(_maxAmountPerPeriod, _maxAmountTotal, address(L1Token));
    }

    function setMaxAmount(uint208 _maxAmountPerPeriod, uint256 _maxAmountTotal, address _token) internal {
        l1StandardBridge.setErc20ThrottleWithdrawalsMaxAmount(_token, _maxAmountPerPeriod, _maxAmountTotal);
    }

    function _deal(address addr, uint256 value) internal override {
        // if it's any address other than the contract itself, deal it to alice since it will be sent from that address
        deal(address(L1Token), addr != contractAddress() ? alice : addr, value, true);
    }

    function _transferThrottleAsset(uint256) internal pure override {
        revert("Not used");
    }

    function transferThrottleAsset(uint256 value, ThrottleRevert revertExpectation) internal override {
        // make sure the bridge has enough tokens to finalize the transfer
        deal(address(L1Token), address(l1StandardBridge), value, true);

        // set up the deposits
        uint256 slot = stdstore.target(address(l1StandardBridge)).sig("deposits(address,address)").with_key(
            address(L1Token)
        ).with_key(address(L2Token)).find();

        // Give the L1 bridge some ERC20 tokens
        vm.store(address(l1StandardBridge), bytes32(slot), bytes32(uint256(value)));
        assertEq(l1StandardBridge.deposits(address(L1Token), address(L2Token)), value);

        vm.mockCall(
            address(l1StandardBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1StandardBridge.OTHER_BRIDGE()))
        );

        address messenger = address(l1StandardBridge.messenger());
        _transferThrottleAssetCheck(revertExpectation);

        vm.prank(messenger);
        l1StandardBridge.finalizeBridgeERC20(address(L1Token), address(L2Token), alice, alice, value, hex"");
    }

    /// @dev Setup the system for a ready-to-use state.
    function setUp() public override(Bridge_Initializer, TransferThrottleTest) {
        super.setUp();

        // set deposit limit to max so we don't get blocked
        ResourceMetering.ResourceConfig memory config = Constants.DEFAULT_RESOURCE_CONFIG();
        config.maxTransactionLimit = type(uint16).max;
        vm.prank(systemConfig.owner());
        systemConfig.setResourceConfig(config);

        // set up roles for access control testing
        vm.startPrank(superchainConfig.defaultAdmin());
        superchainConfig.grantRole(superchainConfig.MONITOR_ROLE(), monitor);
        superchainConfig.grantRole(superchainConfig.OPERATOR_ROLE(), operator);
        vm.stopPrank();

        // set up throttling parameters
        vm.prank(optimismPortal.guardian());
        setMaxAmount(100 ether, 0);

        vm.prank(optimismPortal.guardian());
        setPeriodLength(1 hours);

        resetPeriod();
    }

    /// @notice Try to set the throttle config of a token that does not exist
    function test_setThrottle_nocode_reverts() external {
        address noCodeToken = address(0xdeaddead);
        require(noCodeToken.code.length == 0, "address is supposed to have no code");

        // technically doesn't matter who calls it but let's assume we would have the capability
        vm.startPrank(operator);
        vm.expectRevert(prefixContractName(": token has no code"));
        setMaxAmount(0, 0, noCodeToken);
        vm.expectRevert(prefixContractName(": token has no code"));
        setPeriodLength(0, noCodeToken);
        vm.stopPrank();
    }

    function test_setThrottle_maxAmount() external override {
        throttleSetMaxAmountTestDisabled();
    }
}
