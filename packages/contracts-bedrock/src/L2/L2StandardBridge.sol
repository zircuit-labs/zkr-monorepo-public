// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { Predeploys } from "src/libraries/Predeploys.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { ISemver } from "src/universal/ISemver.sol";
import { OptimismMintableERC20 } from "src/universal/OptimismMintableERC20.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { Constants } from "src/libraries/Constants.sol";
import { AccessControlPausable } from "src/universal/AccessControlPausable.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/// @custom:proxied
/// @custom:predeploy 0x4200000000000000000000000000000000000010
/// @title L2StandardBridge
/// @notice The L2StandardBridge is responsible for transfering ETH and ERC20 tokens between L1 and
///         L2. In the case that an ERC20 token is native to L2, it will be escrowed within this
///         contract. If the ERC20 token is native to L1, it will be burnt.
///         NOTE: this contract is not intended to support all variations of ERC20 tokens. Examples
///         of some token types that may not be properly supported by this contract include, but are
///         not limited to: tokens with transfer fees, rebasing tokens, and tokens with blocklists.
contract L2StandardBridge is StandardBridge, ISemver {
    /// @custom:semver 1.8.0
    string public constant version = "1.8.0";

    // @notice Throttle config for ERC20 withdrawals (L2 => L1) by L2 token address
    mapping(address => Throttle) public erc20ThrottleWithdrawals;

    /// @notice Constructs the L2StandardBridge contract.
    constructor() StandardBridge() {
        initialize({ _otherBridge: StandardBridge(payable(address(0))) });
    }

    /// @notice Initializer.
    /// @param _otherBridge Contract for the corresponding bridge on the other chain.
    function initialize(StandardBridge _otherBridge) public initializer {
        __StandardBridge_init({
            _messenger: CrossDomainMessenger(Predeploys.L2_CROSS_DOMAIN_MESSENGER),
            _otherBridge: _otherBridge,
            _accessController: AccessControlPausable(Predeploys.L2_CONTROLLER)
        });
    }

    /// @notice Allows EOAs to bridge ETH by sending directly to the bridge.
    receive() external payable override onlyEOA {
        _initiateBridgeETH(msg.sender, msg.sender, msg.value, RECEIVE_DEFAULT_GAS_LIMIT, bytes(""));
    }

    function _throttleETHInitiate(address, uint256) internal override {
        // don't throttle eth withdrawals here since this is handled on the L2ToL1MessagePasser already
    }

    function _throttleERC20Initiate(address, address _localToken, uint256 _amount) internal override {
        Throttle storage throttle = erc20ThrottleWithdrawals[_localToken];
        // withdrawals are throttled globally to guard against hacks with large withdrawals
        _transferThrottling(throttle, _throttleGlobalUser, IERC20(_localToken).balanceOf(address(this)), _amount);
    }

    function _throttleERC20Finalize(address, address, uint256) internal override {
        // don't throttle deposits onto L2
    }

    /// @notice Returns the number of erc20 tokens of `token` that can be withdrawn before being throttled
    function getErc20ThrottleWithdrawalsCredits(address token) external view returns (uint256 availableCredits) {
        availableCredits = _throttleUserAvailableCredits(_throttleGlobalUser, erc20ThrottleWithdrawals[token]);
    }

    /// @notice Updates the max amount per period for the withdrawals throttle, impacting the current period
    function setErc20ThrottleWithdrawalsMaxAmount(address token, uint208 maxAmountPerPeriod, uint256 maxAmountTotal) external {
        require(token.code.length != 0, "StandardBridge: token has no code");
        // setting a maximum amount for withdrawals doesn't make any sense
        require(maxAmountTotal == 0, "StandardBridge: max total amount not supported");
        _setThrottle(maxAmountPerPeriod, maxAmountTotal, erc20ThrottleWithdrawals[token]);
    }

    /// @notice Sets the length of the withdrawals throttle period to `_periodLength`, which
    ///         immediately affects the speed of credit accumulation.
    function setErc20ThrottleWithdrawalsPeriodLength(address token, uint48 _periodLength) external {
        require(token.code.length != 0, "StandardBridge: token has no code");
        _setPeriodLength(_periodLength, erc20ThrottleWithdrawals[token]);
    }
}
