// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { console2 as console } from "forge-std/console2.sol";
import { CommonTest } from "test/setup/CommonTest.sol";

abstract contract TransferThrottleTest is CommonTest {

    /// @notice which revert to expect
    enum ThrottleRevert {
        None,
        Throughput,
        Total
    }

    /// @notice Return the contract name under test
    function contractAddress() internal view virtual returns (address);

    /// @notice Return the contract name under test
    function contractName() internal pure virtual returns (string memory);

    /// @notice Return the period length of the throttle under test
    function periodLength() internal view virtual returns (uint48 _periodLength);

    /// @notice Return the max amount per period of the throttle under test
    function maxAmountPerPeriod() internal view virtual returns (uint208 _maxAmountPerPeriod);

    /// @notice Return the max amount total of the throttle under test
    function maxAmountTotal() internal view virtual returns (uint256 _maxAmountTotal);

    /// @notice Set the period length of the throttle under test
    function setPeriodLength(uint48 _periodLength) internal virtual;

    /// @notice Set the max amount per period of the throttle under test
    function setMaxAmount(uint208 _maxAmountPerPeriod, uint256 _maxAmountTotal) internal virtual;

    /// @notice Perform transfer for the throttle under test
    function _transferThrottleAsset(uint256 value) internal virtual;

    /// @notice Function to set balance of underlying asset of address, defaults to eth
    function _deal(address addr, uint256 value) internal virtual {
        vm.deal(addr, value);
    }

    /// @notice monitor address for testing
    address constant monitor = address(0x6d6f6e69746f72);
    /// @notice operator address for testing
    address constant operator = address(0x6f70657261746f72);

    /// @notice Perform transfer for the throttle under test
    function _transferThrottleAssetCheck(ThrottleRevert revertExpectation) internal virtual {
        if (revertExpectation == ThrottleRevert.Total) {
            vm.expectRevert("TransferThrottle: maximum allowed total amount exceeded");
        } else if (revertExpectation == ThrottleRevert.Throughput) {
            vm.expectRevert("TransferThrottle: maximum allowed throughput exceeded");
        }
    }

    /// @notice Transfer eth for the throttle under test. If `shouldThrottle`,
    ///         is set, it expects a revert.
    function transferThrottleAsset(uint256 value, ThrottleRevert revertExpectation) internal virtual {
        _transferThrottleAssetCheck(revertExpectation);
        _transferThrottleAsset(value);
    }

    /// @notice Prefix `_str` with the contract name under test
    function prefixContractName(string memory _str) internal pure returns (bytes memory) {
        return bytes(string.concat(contractName(), _str));
    }

    /// @notice forward the timestamp to make sure max credits are available
    function resetPeriod() internal {
        makeCreditsAvailable(100);
    }

    /// @notice forward the timestamp to make `forwardPercentage` credits available,
    ///         maxing out at 100%
    function makeCreditsAvailable(uint256 forwardPercentage) internal {
        vm.warp(block.timestamp + (forwardPercentage * periodLength() / 100));
    }

    /// @dev Setup the system for a ready-to-use state.
    function setUp() public virtual override {
        super.setUp();
    }

    /// @notice Tests that deposits within the allowed limit work
    function test_transfer_succeeds() external {
        uint256 max = maxAmountPerPeriod();

        // single deposits
        transferThrottleAsset(max, ThrottleRevert.None);
        makeCreditsAvailable(50);
        transferThrottleAsset(max / 2, ThrottleRevert.None);
        makeCreditsAvailable(25);
        transferThrottleAsset(max / 4, ThrottleRevert.None);
        makeCreditsAvailable(10);
        transferThrottleAsset(max / 10, ThrottleRevert.None);

        // multiple deposits
        makeCreditsAvailable(40);
        transferThrottleAsset(max / 5, ThrottleRevert.None);
        transferThrottleAsset(max / 5, ThrottleRevert.None);

        makeCreditsAvailable(100);
        transferThrottleAsset(max / 3, ThrottleRevert.None);
        transferThrottleAsset(max / 3, ThrottleRevert.None);
        transferThrottleAsset(max / 3, ThrottleRevert.None);
    }

    /// @notice Tests that single deposits over the limit revert
    function test_single_transfers_reverts() external {
        uint256 max = maxAmountPerPeriod();

        // single deposits
        transferThrottleAsset(max + 1, ThrottleRevert.Throughput);
        transferThrottleAsset(max, ThrottleRevert.None);

        makeCreditsAvailable(50);
        transferThrottleAsset(max / 2 + 1, ThrottleRevert.Throughput);
        transferThrottleAsset(max / 2, ThrottleRevert.None);

        makeCreditsAvailable(25);
        transferThrottleAsset(max / 4 + 1, ThrottleRevert.Throughput);
        transferThrottleAsset(max / 4, ThrottleRevert.None);

        makeCreditsAvailable(10);
        transferThrottleAsset(max / 10 + 1, ThrottleRevert.Throughput);
        transferThrottleAsset(max / 10, ThrottleRevert.None);
    }

    /// @notice Tests that multiple deposits that overall exceed the limit revert
    function test_multiple_transfers_reverts() external {
        uint256 max = maxAmountPerPeriod();
        transferThrottleAsset(max, ThrottleRevert.None);

        // multiple deposits
        makeCreditsAvailable(40);
        transferThrottleAsset(max / 5 + 1, ThrottleRevert.None);
        transferThrottleAsset(max / 5, ThrottleRevert.Throughput);
        transferThrottleAsset(max / 5 - 1, ThrottleRevert.None);

        makeCreditsAvailable(100);
        transferThrottleAsset(max / 3 + 1, ThrottleRevert.None);
        transferThrottleAsset(max / 3 + 1, ThrottleRevert.None);
        transferThrottleAsset(max / 3 + 1, ThrottleRevert.Throughput);

        for (int256 i = 0; i < 10; i++) {
            makeCreditsAvailable(100);
            int256 delta = i - 5;
            transferThrottleAsset(uint256(int256(max) + delta), delta > 0 ? ThrottleRevert.Throughput : ThrottleRevert.None);
        }
    }

    /// @notice Tests that setting the throttle configuration works if the callers
    ///         have the required capabilities
    function test_setThrottle_succeeds() external {
        uint208 max = maxAmountPerPeriod();
        // operator can increase and decrease max amount, as well as set the period length
        vm.startPrank(operator);
        setMaxAmount(max * 2, 0);
        setMaxAmount(max, 0);
        setPeriodLength(1 hours);
        vm.stopPrank();

        // monitor can decrease max amount
        vm.startPrank(monitor);
        setMaxAmount(max / 2, 0);
        vm.stopPrank();
    }

    /// @notice Tests that setting the throttle configuration fails if the callers
    ///         do not have the required capabilities
    function test_setThrottle_reverts() external {
        uint208 max = maxAmountPerPeriod();

        vm.startPrank(operator);
        vm.expectRevert("TransferThrottle: period length cannot be 0");
        setPeriodLength(0);
        vm.stopPrank();

        // monitor cannot increase max amount
        vm.startPrank(monitor);
        vm.expectRevert(prefixContractName(": sender is not throttle admin"));
        setMaxAmount(max * 2, 0);
        vm.expectRevert(prefixContractName(": sender is not throttle admin"));
        setPeriodLength(1 hours);
        vm.stopPrank();

        // random address cannot increase or decrease
        vm.startPrank(alice);
        vm.expectRevert(prefixContractName(": sender is not throttle admin"));
        setMaxAmount(max * 2, 0);

        vm.expectRevert(prefixContractName(": sender not allowed to throttle"));
        setMaxAmount(max / 2, 0);

        vm.expectRevert(prefixContractName(": sender is not throttle admin"));
        setPeriodLength(1 hours);
        vm.stopPrank();
    }

    /// @notice helper for fuzzing, just forwarding arguments
    /// function testFuzz_throttle_totalAmount(...)
    function throttleMaxAmountTotal(uint256 contractBalance, uint256 additionalValue, uint256 maxTotal) internal {
        // make sure adding the balances is actually possible
        unchecked {
            vm.assume(contractBalance + additionalValue >= contractBalance);
        }

        // throttling with just a cap on the total amount
        vm.prank(operator);
        setMaxAmount(0, maxTotal);

        // set up balances
        _deal(contractAddress(), contractBalance);
        _deal(address(this), additionalValue);

        ThrottleRevert revertExpectation = ThrottleRevert.None;
        bool shouldThrottle = maxTotal != 0 && maxTotal < (contractBalance + additionalValue);
        if (shouldThrottle) {
            revertExpectation = ThrottleRevert.Total;
        }
        transferThrottleAsset(additionalValue, revertExpectation);
    }

    /// @notice Override and call one of throttleSetMaxAmountTest{Enabled,Disabled}
    function test_setThrottle_maxAmount() external virtual;

    function throttleSetMaxAmountTestEnabled() internal {
        uint256 max = 1000 ether;
        // operator can increase and decrease max amount, as well as set the period length
        vm.startPrank(operator);
        setMaxAmount(0, max);
        setMaxAmount(0, 2 * max);
        vm.stopPrank();

        // monitor can decrease max amount
        console.log("max total before call:", maxAmountTotal());
        vm.startPrank(monitor);
        setMaxAmount(0, max / 2);
        vm.stopPrank();

        // also test the basic case that we can't transfer more than the max
        vm.startPrank(operator);
        setMaxAmount(0, max);
        vm.stopPrank();

        transferThrottleAsset(max + 1, ThrottleRevert.Total);
    }

    function throttleSetMaxAmountTestDisabled() internal {
        uint256 max = 1000 ether;
        // operator can only set it to zero
        vm.startPrank(operator);
        setMaxAmount(0, 0);

        vm.expectRevert(prefixContractName(": max total amount not supported"));
        setMaxAmount(0, max);
        vm.expectRevert(prefixContractName(": max total amount not supported"));
        setMaxAmount(0, 2 * max);
        vm.stopPrank();

        // monitor can only set it to zero
        vm.startPrank(monitor);
        setMaxAmount(0, 0);

        vm.expectRevert(prefixContractName(": max total amount not supported"));
        setMaxAmount(0, max / 2);
        vm.stopPrank();
    }
}
