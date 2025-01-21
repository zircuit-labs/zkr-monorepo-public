// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { StandardBridge } from "src/universal/StandardBridge.sol";
import { CommonTest } from "test/setup/CommonTest.sol";
import { OptimismMintableERC20 } from "src/universal/OptimismMintableERC20.sol";
import { OZERC20 as ERC20 } from "test/mocks/OZERC20.sol";

/// @title StandardBridgeTester
/// @notice Simple wrapper around the StandardBridge contract that exposes
///         internal functions so they can be more easily tested directly.
contract StandardBridgeTester is StandardBridge {
    constructor() StandardBridge() { }

    function isOptimismMintableERC20(address _token) external view returns (bool) {
        return _isOptimismMintableERC20(_token);
    }

    function isCorrectTokenPair(address _mintableToken, address _otherToken) external view returns (bool) {
        return _isCorrectTokenPair(_mintableToken, _otherToken);
    }

    function _throttleETHInitiate(address, uint256 _amount) internal override { }

    function _throttleERC20Initiate(address, address, uint256 _amount) internal override { }

    function _throttleERC20Finalize(address, address, uint256 _amount) internal override { }


    receive() external payable override { }
}

/// @title StandardBridge_Stateless_Test
/// @notice Tests internal functions that require no existing state or contract
///         interactions with the messenger.
contract StandardBridge_Stateless_Test is CommonTest {
    StandardBridgeTester internal bridge;
    OptimismMintableERC20 internal mintable;
    ERC20 internal erc20;

    function setUp() public override {
        super.setUp();

        bridge = new StandardBridgeTester();

        mintable = new OptimismMintableERC20({
            _bridge: address(0),
            _remoteToken: address(0),
            _name: "Stonks",
            _symbol: "STONK",
            _decimals: 18
        });

        erc20 = new ERC20("Altcoin", "ALT");
    }

    /// @notice Test coverage for identifying OptimismMintableERC20 tokens.
    ///         This function should return true for both modern and legacy
    ///         OptimismMintableERC20 tokens and false for any accounts that
    ///         do not implement the interface.
    function test_isOptimismMintableERC20_succeeds() external view {
        // Both the modern and legacy mintable tokens should return true
        assertTrue(bridge.isOptimismMintableERC20(address(mintable)));
        // A regular ERC20 should return false
        assertFalse(bridge.isOptimismMintableERC20(address(erc20)));
        // Non existent contracts should return false and not revert
        assertEq(address(0x20).code.length, 0);
        assertFalse(bridge.isOptimismMintableERC20(address(0x20)));
    }

    /// @notice Test coverage of isCorrectTokenPair under different types of
    ///         tokens.
    function test_isCorrectTokenPair_succeeds() external {
        // Modern + known to be correct remote token
        assertTrue(bridge.isCorrectTokenPair(address(mintable), mintable.remoteToken()));
        // Modern + known to be correct l1Token (legacy interface)
        assertTrue(bridge.isCorrectTokenPair(address(mintable), mintable.l1Token()));
        // Modern + known to be incorrect remote token
        assertTrue(mintable.remoteToken() != address(0x20));
        assertFalse(bridge.isCorrectTokenPair(address(mintable), address(0x20)));
        // A token that doesn't support either modern or legacy interface
        // will revert
        vm.expectRevert(bytes(""));
        bridge.isCorrectTokenPair(address(erc20), address(1));
    }
}
