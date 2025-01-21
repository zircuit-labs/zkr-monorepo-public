// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

// Testing utilities
import { EIP1967Helper } from "test/mocks/EIP1967Helper.sol";

// Target contract is imported by the `Bridge_Initializer`
import { Bridge_Initializer } from "test/setup/Bridge_Initializer.sol";
import { stdStorage, StdStorage } from "forge-std/Test.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { L2ToL1MessagePasser } from "src/L2/L2ToL1MessagePasser.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { TransferThrottleTest } from "test/universal/TransferThrottle.t.sol";

// Libraries
import { Hashing } from "src/libraries/Hashing.sol";
import { Types } from "src/libraries/Types.sol";

// Target contract dependencies
import { L2StandardBridge } from "src/L2/L2StandardBridge.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { OptimismMintableERC20 } from "src/universal/OptimismMintableERC20.sol";

contract L2StandardBridge_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    /// @dev Test that the bridge's constructor sets the correct values.
    function test_constructor_succeeds() external view {
        L2StandardBridge impl =
            L2StandardBridge(payable(EIP1967Helper.getImplementation(deploy.mustGetAddress("L2StandardBridge"))));
        // The implementation contract is initialized with a 0 L1 bridge address,
        // but the L2 cross-domain-messenger is always set to the predeploy address for both proxy and implementation.
        assertEq(address(impl.MESSENGER()), Predeploys.L2_CROSS_DOMAIN_MESSENGER, "constructor zero check MESSENGER");
        assertEq(address(impl.messenger()), Predeploys.L2_CROSS_DOMAIN_MESSENGER, "constructor zero check messenger");
        assertEq(address(impl.OTHER_BRIDGE()), address(0), "constructor zero check OTHER_BRIDGE");
        assertEq(address(impl.otherBridge()), address(0), "constructor zero check otherBridge");
    }

    /// @dev Tests that the bridge is initialized correctly.
    function test_initialize_succeeds() external view {
        assertEq(address(l2StandardBridge.MESSENGER()), address(l2CrossDomainMessenger));
        assertEq(address(l2StandardBridge.messenger()), address(l2CrossDomainMessenger));
        assertEq(address(l2StandardBridge.OTHER_BRIDGE()), address(l1StandardBridge));
        assertEq(address(l2StandardBridge.otherBridge()), address(l1StandardBridge));
    }

    /// @dev Ensures that the L2StandardBridge is always not paused. The pausability
    ///      happens on L1 and not L2.
    function test_paused_succeeds() external view {
        assertFalse(l2StandardBridge.paused());
    }

    /// @dev Tests that the bridge receives ETH and successfully initiates a withdrawal.
    function test_receive_succeeds() external {
        assertEq(address(l2ToL1MessagePasser).balance, 0);
        uint256 nonce = l2CrossDomainMessenger.messageNonce();

        bytes memory message =
            abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, alice, 100, hex"");
        uint64 baseGas = l2CrossDomainMessenger.baseGas(message, 200_000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l2StandardBridge),
            address(l1StandardBridge),
            100,
            200_000,
            message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(l2CrossDomainMessenger),
                target: address(l1CrossDomainMessenger),
                value: 100,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        vm.expectEmit(true, true, true, true);
        emit ETHBridgeInitiated(alice, alice, 100, hex"");

        // L2ToL1MessagePasser will emit a MessagePassed event
        vm.expectEmit(true, true, true, true, address(l2ToL1MessagePasser));
        emit MessagePassed(
            nonce,
            address(l2CrossDomainMessenger),
            address(l1CrossDomainMessenger),
            100,
            baseGas,
            withdrawalData,
            withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(l2CrossDomainMessenger));
        emit SentMessage(address(l1StandardBridge), address(l2StandardBridge), message, nonce, 200_000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(l2CrossDomainMessenger));
        emit SentMessageExtension1(address(l2StandardBridge), 100);

        vm.expectCall(
            address(l2CrossDomainMessenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(l1StandardBridge),
                message,
                200_000 // StandardBridge's RECEIVE_DEFAULT_GAS_LIMIT
            )
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector,
                address(l1CrossDomainMessenger),
                baseGas,
                withdrawalData
            )
        );

        vm.prank(alice, alice);
        (bool success,) = address(l2StandardBridge).call{ value: 100 }(hex"");
        assertEq(success, true);
        assertEq(address(l2ToL1MessagePasser).balance, 100);
    }
}

contract PreBridgeERC20 is Bridge_Initializer {
    /// @dev Sets up expected calls and emits for a successful ERC20 withdrawal.
    function _preBridgeERC20(address _l2Token) internal {
        // Alice has 100 L2Token
        deal(_l2Token, alice, 100, true);
        assertEq(ERC20(_l2Token).balanceOf(alice), 100);
        uint256 nonce = l2CrossDomainMessenger.messageNonce();
        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(L1Token), _l2Token, alice, alice, 100, hex""
        );
        uint64 baseGas = l2CrossDomainMessenger.baseGas(message, 1000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l2StandardBridge),
            address(l1StandardBridge),
            0,
            1000,
            message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(l2CrossDomainMessenger),
                target: address(l1CrossDomainMessenger),
                value: 0,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        vm.expectCall(
            address(l2StandardBridge),
            abi.encodeWithSelector(
                l2StandardBridge.bridgeERC20.selector, _l2Token, address(L1Token), 100, 1000, hex""
            )
        );

        vm.expectCall(
            address(l2CrossDomainMessenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l1StandardBridge), message, 1000)
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector,
                address(l1CrossDomainMessenger),
                baseGas,
                withdrawalData
            )
        );

        // The l2StandardBridge should burn the tokens
        vm.expectCall(_l2Token, abi.encodeWithSelector(OptimismMintableERC20.burn.selector, alice, 100));

        vm.expectEmit(true, true, true, true);
        emit ERC20BridgeInitiated(_l2Token, address(L1Token), alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true);
        emit MessagePassed(
            nonce,
            address(l2CrossDomainMessenger),
            address(l1CrossDomainMessenger),
            0,
            baseGas,
            withdrawalData,
            withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true);
        emit SentMessage(address(l1StandardBridge), address(l2StandardBridge), message, nonce, 1000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true);
        emit SentMessageExtension1(address(l2StandardBridge), 0);

        vm.prank(alice, alice);
    }
}

contract L2StandardBridge_BridgeERC20_Test is PreBridgeERC20 {
    // BridgeERC20
    // - token is burned
    // - calls Withdrawer.initiateWithdrawal
    function test_bridgeERC20_succeeds() external {
        _preBridgeERC20({_l2Token: address(L2Token) });
        l2StandardBridge.bridgeERC20(address(L2Token), address(L1Token), 100, 1000, hex"");

        assertEq(L2Token.balanceOf(alice), 0);
    }
}

contract PreBridgeERC20To is Bridge_Initializer {
    // withdrawTo and BridgeERC20To should behave the same when transferring ERC20 tokens
    // so they should share the same setup and expectEmit calls
    function _preBridgeERC20To(address _l2Token) internal {
        deal(_l2Token, alice, 100, true);
        assertEq(ERC20(L2Token).balanceOf(alice), 100);
        uint256 nonce = l2CrossDomainMessenger.messageNonce();
        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(L1Token), _l2Token, alice, bob, 100, hex""
        );
        uint64 baseGas = l2CrossDomainMessenger.baseGas(message, 1000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l2StandardBridge),
            address(l1StandardBridge),
            0,
            1000,
            message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(l2CrossDomainMessenger),
                target: address(l1CrossDomainMessenger),
                value: 0,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        vm.expectEmit(true, true, true, true, address(l2StandardBridge));
        emit ERC20BridgeInitiated(_l2Token, address(L1Token), alice, bob, 100, hex"");

        vm.expectEmit(true, true, true, true, address(l2ToL1MessagePasser));
        emit MessagePassed(
            nonce,
            address(l2CrossDomainMessenger),
            address(l1CrossDomainMessenger),
            0,
            baseGas,
            withdrawalData,
            withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(l2CrossDomainMessenger));
        emit SentMessage(address(l1StandardBridge), address(l2StandardBridge), message, nonce, 1000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(l2CrossDomainMessenger));
        emit SentMessageExtension1(address(l2StandardBridge), 0);

        vm.expectCall(
            address(l2StandardBridge),
            abi.encodeWithSelector(
                l2StandardBridge.bridgeERC20To.selector, _l2Token, address(L1Token), bob, 100, 1000, hex""
            )
        );

        vm.expectCall(
            address(l2CrossDomainMessenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l1StandardBridge), message, 1000)
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector,
                address(l1CrossDomainMessenger),
                baseGas,
                withdrawalData
            )
        );

        // The l2StandardBridge should burn the tokens
        vm.expectCall(address(L2Token), abi.encodeWithSelector(OptimismMintableERC20.burn.selector, alice, 100));

        vm.prank(alice, alice);
    }
}

contract L2StandardBridge_BridgeERC20To_Test is PreBridgeERC20To {
    /// @dev Tests that `bridgeERC20To` burns the tokens, emits `WithdrawalInitiated`,
    ///      and initiates a withdrawal with `Withdrawer.initiateWithdrawal`.
    function test_bridgeERC20To_succeeds() external {
        _preBridgeERC20To({_l2Token: address(L2Token) });
        l2StandardBridge.bridgeERC20To(address(L2Token), address(L1Token), bob, 100, 1000, hex"");
        assertEq(L2Token.balanceOf(alice), 0);
    }
}

contract L2StandardBridge_Bridge_Test is Bridge_Initializer {
    /// @dev Tests that `finalizeDeposit` reverts if the amounts do not match.
    function test_finalizeBridgeETH_incorrectValue_reverts() external {
        vm.mockCall(
            address(l2StandardBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l2StandardBridge.OTHER_BRIDGE()))
        );
        vm.deal(address(l2CrossDomainMessenger), 100);
        vm.prank(address(l2CrossDomainMessenger));
        vm.expectRevert("StandardBridge: amount sent does not match amount required");
        l2StandardBridge.finalizeBridgeETH{ value: 50 }(alice, alice, 100, hex"");
    }

    /// @dev Tests that `finalizeDeposit` reverts if the receipient is the other bridge.
    function test_finalizeBridgeETH_sendToSelf_reverts() external {
        vm.mockCall(
            address(l2StandardBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l2StandardBridge.OTHER_BRIDGE()))
        );
        vm.deal(address(l2CrossDomainMessenger), 100);
        vm.prank(address(l2CrossDomainMessenger));
        vm.expectRevert("StandardBridge: cannot send to self");
        l2StandardBridge.finalizeBridgeETH{ value: 100 }(alice, address(l2StandardBridge), 100, hex"");
    }

    /// @dev Tests that `finalizeDeposit` reverts if the receipient is the messenger.
    function test_finalizeBridgeETH_sendToMessenger_reverts() external {
        vm.mockCall(
            address(l2StandardBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l2StandardBridge.OTHER_BRIDGE()))
        );
        vm.deal(address(l2CrossDomainMessenger), 100);
        vm.prank(address(l2CrossDomainMessenger));
        vm.expectRevert("StandardBridge: cannot send to messenger");
        l2StandardBridge.finalizeBridgeETH{ value: 100 }(alice, address(l2CrossDomainMessenger), 100, hex"");
    }
}

contract L2StandardBridge_FinalizeBridgeETH_Test is Bridge_Initializer {
    /// @dev Tests that `finalizeBridgeETH` succeeds.
    function test_finalizeBridgeETH_succeeds() external {
        address messenger = address(l2StandardBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l2StandardBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);

        vm.expectEmit(true, true, true, true);
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        l2StandardBridge.finalizeBridgeETH{ value: 100 }(alice, alice, 100, hex"");
    }
}

contract L2StandardBridge_WithdrawalThrottling_Test is Bridge_Initializer, TransferThrottleTest {
    using stdStorage for StdStorage;

    function contractAddress() internal view override returns (address) {
        return address(l2StandardBridge);
    }

    function contractName() internal pure override returns (string memory) {
        return "StandardBridge";
    }

    function periodLength() internal view override returns (uint48 _periodLength) {
        (, _periodLength, ) = l2StandardBridge.erc20ThrottleWithdrawals(address(L2Token));
    }

    function maxAmountPerPeriod() internal view override returns (uint208 _maxAmountPerPeriod) {
        (_maxAmountPerPeriod, , ) = l2StandardBridge.erc20ThrottleWithdrawals(address(L2Token));
    }

    function maxAmountTotal() internal view override returns (uint256 _maxAmountTotal) {
        (, , _maxAmountTotal) = l2StandardBridge.erc20ThrottleWithdrawals(address(L2Token));
    }

    function setPeriodLength(uint48 _periodLength) internal override {
        setPeriodLength(_periodLength, address(L2Token));
    }

    function setPeriodLength(uint48 _periodLength, address _token) internal {
        l2StandardBridge.setErc20ThrottleWithdrawalsPeriodLength(_token, _periodLength);
    }

    function setMaxAmount(uint208 _maxAmountPerPeriod, uint256 _maxAmountTotal) internal override {
        setMaxAmount(_maxAmountPerPeriod, _maxAmountTotal, address(L2Token));
    }

    function setMaxAmount(uint208 _maxAmountPerPeriod, uint256 _maxAmountTotal, address _token) internal {
        l2StandardBridge.setErc20ThrottleWithdrawalsMaxAmount(_token, _maxAmountPerPeriod, _maxAmountTotal);
    }

    function _deal(address addr, uint256 value) internal override {
        // if it's any address other than the contract itself, deal it to alice since it will be sent from that address
        deal(address(L2Token), addr != contractAddress() ? alice : addr, value, true);
    }

    function _transferThrottleAsset(uint256) internal pure override {
        revert("Not used");
    }

    function transferThrottleAsset(uint256 value, ThrottleRevert revertExpectation) internal override {
        // Deal Alice's ERC20 State
        deal(address(L2Token), alice, value, true);

        _transferThrottleAssetCheck(revertExpectation);

        vm.prank(alice);
        l2StandardBridge.bridgeERC20(address(L2Token), address(L1Token), value, 10000, hex"");
    }

    /// @dev Setup the system for a ready-to-use state.
    function setUp() public override(Bridge_Initializer, TransferThrottleTest) {
        super.setUp();

        // set up roles for access control testing
        vm.startPrank(l2Controller.defaultAdmin());
        l2Controller.grantRole(l2Controller.MONITOR_ROLE(), monitor);
        l2Controller.grantRole(l2Controller.OPERATOR_ROLE(), operator);
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
