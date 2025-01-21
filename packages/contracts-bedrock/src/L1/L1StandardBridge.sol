// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { Predeploys } from "src/libraries/Predeploys.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { ISemver } from "src/universal/ISemver.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { SuperchainConfig } from "src/L1/SuperchainConfig.sol";
import { Constants } from "src/libraries/Constants.sol";
import { AccessControlPausable } from "src/universal/AccessControlPausable.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/// @custom:proxied
/// @title L1StandardBridge
/// @notice The L1StandardBridge is responsible for transfering ETH and ERC20 tokens between L1 and
///         L2. In the case that an ERC20 token is native to L1, it will be escrowed within this
///         contract. If the ERC20 token is native to L2, it will be burnt. Before Bedrock, ETH was
///         stored within this contract. After Bedrock, ETH is instead stored inside the
///         OptimismPortal contract.
///         NOTE: this contract is not intended to support all variations of ERC20 tokens. Examples
///         of some token types that may not be properly supported by this contract include, but are
///         not limited to: tokens with transfer fees, rebasing tokens, and tokens with blocklists.
contract L1StandardBridge is StandardBridge, ISemver {
    /// @notice Semantic version.
    /// @custom:semver 2.1.0
    string public constant version = "2.1.0";

    /// @notice Throttle for eth deposits
    Throttle public ethThrottleDeposits;

    // @notice Throttle config for ERC20 deposits (L1 => L2) by L1 token address
    mapping(address => Throttle) public erc20ThrottleDeposits;

    // @notice Throttle config for ERC20 withdrawals (L2 => L1) by L1 token address
    mapping(address => Throttle) public erc20ThrottleWithdrawals;

    /// @notice Constructs the L1StandardBridge contract.
    constructor() StandardBridge() {
        initialize({ _messenger: CrossDomainMessenger(address(0)), _superchainConfig: SuperchainConfig(address(0)) });
    }

    /// @notice Initializer.
    /// @param _messenger        Contract for the CrossDomainMessenger on this network.
    /// @param _superchainConfig Contract for the SuperchainConfig on this network.
    function initialize(CrossDomainMessenger _messenger, SuperchainConfig _superchainConfig) public initializer {
        __StandardBridge_init({
            _messenger: _messenger,
            _otherBridge: StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE)),
            _accessController: AccessControlPausable(_superchainConfig)
        });
    }


    /// @notice The access controller is also the superchain config. To avoid storing it twice, only use a getter here
    function superchainConfig() external view returns (SuperchainConfig) {
        return SuperchainConfig(address(accessController));
    }

    /// @inheritdoc StandardBridge
    function paused() public view override returns (bool) {
        return accessController.paused();
    }

    /// @notice Allows EOAs to bridge ETH by sending directly to the bridge.
    receive() external payable override onlyEOA {
        _initiateETHDeposit(msg.sender, msg.sender, RECEIVE_DEFAULT_GAS_LIMIT, bytes(""));
    }

    /// @notice Internal function for initiating an ETH deposit.
    /// @param _from        Address of the sender on L1.
    /// @param _to          Address of the recipient on L2.
    /// @param _minGasLimit Minimum gas limit for the deposit message on L2.
    /// @param _extraData   Optional data to forward to L2.
    function _initiateETHDeposit(address _from, address _to, uint32 _minGasLimit, bytes memory _extraData) internal {
        _initiateBridgeETH(_from, _to, msg.value, _minGasLimit, _extraData);
    }

    function _throttleETHInitiate(address _from, uint256 _amount) internal override {
        // perform per-user throttling for ETH deposits through the bridge
        _transferThrottling(ethThrottleDeposits, _from, address(this).balance - _amount, _amount);
    }

    function _throttleERC20Initiate(address _from, address _localToken, uint256 _amount) internal override {
        // perform per-user throttling for ERC20 deposits through the bridge
        Throttle storage throttle = erc20ThrottleDeposits[_localToken];
        _transferThrottling(throttle, _from, IERC20(_localToken).balanceOf(address(this)), _amount);
    }

    function _throttleERC20Finalize(address, address _localToken, uint256 _amount) internal override {
        Throttle storage throttle = erc20ThrottleWithdrawals[_localToken];
        // withdrawals are throttled globally to guard against hacks with large withdrawals
        _transferThrottling(throttle, _throttleGlobalUser, IERC20(_localToken).balanceOf(address(this)), _amount);
    }

    /// @notice Returns the amount of eth that `user` can deposit before being throttled, not taking into account the total locked value
    function getEthThrottleDepositsCredits(address user) external view returns (uint256 availableCredits) {
        availableCredits = _throttleUserAvailableCredits(user, ethThrottleDeposits);
    }

    /// @notice Returns the number of erc20 tokens of `token` that `user` can deposit before being throttled, not taking into account the total locked value
    function getERC20ThrottleDepositsCredits(address token, address user) external view returns (uint256 availableCredits) {
        availableCredits = _throttleUserAvailableCredits(user, erc20ThrottleDeposits[token]);
    }

    /// @notice Returns the number of erc20 tokens of `token` that can be withdrawn before being throttled
    function getERC20ThrottleWithdrawalsCredits(address token) external view returns (uint256 availableCredits) {
        availableCredits = _throttleUserAvailableCredits(_throttleGlobalUser, erc20ThrottleWithdrawals[token]);
    }

    /// @notice Updates the max amount per period for the deposits throttle, impacting the current period
    function setEthThrottleDepositsMaxAmount(uint208 maxAmountPerPeriod, uint256 maxAmountTotal) external {
        // we only perform per-user throttling of eth deposits since the global cap is handled on the OptimismPortal
        require(maxAmountTotal == 0, "StandardBridge: max total amount not supported");
        _setThrottle(maxAmountPerPeriod, maxAmountTotal, ethThrottleDeposits);
    }

    /// @notice Updates the max amount per period and total amount for the deposits throttle, impacting the current period
    function setErc20ThrottleDepositsMaxAmount(address token, uint208 maxAmountPerPeriod, uint256 maxAmountTotal) external {
        require(token.code.length != 0, "StandardBridge: token has no code");
        _setThrottle(maxAmountPerPeriod, maxAmountTotal, erc20ThrottleDeposits[token]);
    }

    /// @notice Updates the max amount per period for the withdrawals throttle, impacting the current period
    function setErc20ThrottleWithdrawalsMaxAmount(address token, uint208 maxAmountPerPeriod, uint256 maxAmountTotal) external {
        require(token.code.length != 0, "StandardBridge: token has no code");
        // setting a maximum amount for withdrawals doesn't make any sense
        require(maxAmountTotal == 0, "StandardBridge: max total amount not supported");
        _setThrottle(maxAmountPerPeriod, maxAmountTotal, erc20ThrottleWithdrawals[token]);
    }

    /// @notice Sets the length of the deposits throttle period to `_periodLength`, which
    ///         immediately affects the speed of credit accumulation.
    function setEthThrottleDepositsPeriodLength(uint48 _periodLength) external {
        _setPeriodLength(_periodLength, ethThrottleDeposits);
    }

    /// @notice Sets the length of the deposits throttle period to `_periodLength`, which
    ///         immediately affects the speed of credit accumulation.
    function setErc20ThrottleDepositsPeriodLength(address token, uint48 _periodLength) external {
        require(token.code.length != 0, "StandardBridge: token has no code");
        _setPeriodLength(_periodLength, erc20ThrottleDeposits[token]);
    }

    /// @notice Sets the length of the withdrawals throttle period to `_periodLength`, which
    ///         immediately affects the speed of credit accumulation.
    function setErc20ThrottleWithdrawalsPeriodLength(address token, uint48 _periodLength) external {
        require(token.code.length != 0, "StandardBridge: token has no code");
        _setPeriodLength(_periodLength, erc20ThrottleWithdrawals[token]);
    }
}
