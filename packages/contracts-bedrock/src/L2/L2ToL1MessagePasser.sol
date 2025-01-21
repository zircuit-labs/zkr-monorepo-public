// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { Types } from "src/libraries/Types.sol";
import { Hashing } from "src/libraries/Hashing.sol";
import { Encoding } from "src/libraries/Encoding.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Burn } from "src/libraries/Burn.sol";
import { ISemver } from "src/universal/ISemver.sol";
import { AccessControlPausable } from "src/universal/AccessControlPausable.sol";
import { TransferThrottle } from "src/universal/TransferThrottle.sol";
import { Initializable } from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

/// @custom:proxied
/// @custom:predeploy 0x4200000000000000000000000000000000000016
/// @title L2ToL1MessagePasser
/// @notice The L2ToL1MessagePasser is a dedicated contract where messages that are being sent from
///         L2 to L1 can be stored. The storage root of this contract is pulled up to the top level
///         of the L2 output to reduce the cost of proving the existence of sent messages.
contract L2ToL1MessagePasser is Initializable, TransferThrottle, ISemver {
    /// @notice The L1 gas limit set when eth is withdrawn using the receive() function.
    uint256 internal constant RECEIVE_DEFAULT_GAS_LIMIT = 100_000;

    /// @notice The current message version identifier.
    uint16 public constant MESSAGE_VERSION = 1;

    /// @notice Includes the message hashes for all withdrawals
    mapping(bytes32 => bool) public sentMessages;

    /// @notice A unique value hashed with each withdrawal.
    uint240 internal msgNonce;

    /// @notice Throttle for eth withdrawals
    TransferThrottle.Throttle public ethThrottleWithdrawals;

    /// @notice Contract that controls whether pausing is enabled and which addresses are allowed to
    ///         adjust the throttle configuration
    AccessControlPausable public accessController;

    /// @notice Emitted any time a withdrawal is initiated.
    /// @param nonce          Unique value corresponding to each withdrawal.
    /// @param sender         The L2 account address which initiated the withdrawal.
    /// @param target         The L1 account address the call will be send to.
    /// @param value          The ETH value submitted for withdrawal, to be forwarded to the target.
    /// @param gasLimit       The minimum amount of gas that must be provided when withdrawing.
    /// @param data           The data to be forwarded to the target on L1.
    /// @param withdrawalHash The hash of the withdrawal.
    event MessagePassed(
        uint256 indexed nonce,
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 gasLimit,
        bytes data,
        bytes32 withdrawalHash
    );

    /// @notice Emitted when the balance of this contract is burned.
    /// @param amount Amount of ETh that was burned.
    event WithdrawerBalanceBurnt(uint256 indexed amount);

    /// @custom:semver 1.1.0
    string public constant version = "1.1.0";

    /// @notice Constructs the SuperchainConfig contract.
    constructor() { initialize(); }

    /// @notice Initializer.
    function initialize() public initializer {
        accessController = AccessControlPausable(Predeploys.L2_CONTROLLER);
    }

    /// @notice Allows users to withdraw ETH by sending directly to this contract.
    receive() external payable {
        initiateWithdrawal(msg.sender, RECEIVE_DEFAULT_GAS_LIMIT, bytes(""));
    }

    /// @notice Removes all ETH held by this contract from the state. Used to prevent the amount of
    ///         ETH on L2 inflating when ETH is withdrawn. Currently only way to do this is to
    ///         create a contract and self-destruct it to itself. Anyone can call this function. Not
    ///         incentivized since this function is very cheap.
    function burn() external {
        uint256 balance = address(this).balance;
        Burn.eth(balance);
        emit WithdrawerBalanceBurnt(balance);
    }

    /// @notice This function should return true if the contract is paused.
    function paused() public view returns (bool) {
        return accessController.paused();
    }

    /// @notice Sends a message from L2 to L1.
    /// @param _target   Address to call on L1 execution.
    /// @param _gasLimit Minimum gas limit for executing the message on L1.
    /// @param _data     Data to forward to L1 target.
    function initiateWithdrawal(address _target, uint256 _gasLimit, bytes memory _data) public payable {
        if (paused()) {
            revert("L2ToL1MessagePasser: paused");
        }

        // record amount bridged and revert in case it exceeds the allowed value
        // existing value is 0 since a maximum total amount is not supported for withdrawals
        _transferThrottling(ethThrottleWithdrawals, _throttleGlobalUser, 0, msg.value);

        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: messageNonce(),
                sender: msg.sender,
                target: _target,
                value: msg.value,
                gasLimit: _gasLimit,
                data: _data
            })
        );

        sentMessages[withdrawalHash] = true;

        emit MessagePassed(messageNonce(), msg.sender, _target, msg.value, _gasLimit, _data, withdrawalHash);

        unchecked {
            ++msgNonce;
        }
    }

    /// @notice Retrieves the next message nonce. Message version will be added to the upper two
    ///         bytes of the message nonce. Message version allows us to treat messages as having
    ///         different structures.
    /// @return Nonce of the next message to be sent, with added message version.
    function messageNonce() public view returns (uint256) {
        return Encoding.encodeVersionedNonce(msgNonce, MESSAGE_VERSION);
    }

    /// @notice Checks whether `_address` is allowed to change all throttle
    ///         configurations (including disabling it).
    function _transferThrottleHasAdminAccess(address _address) internal view override {
        require(accessController.hasOperatorCapabilities(_address), "L2ToL1MessagePasser: sender is not throttle admin");
    }

    /// @notice Checks whether `_address` is allowed to decrease the amount
    ///         that is allowed within the configured period. The function needs
    ///         to revert if `_address` is not allowed.
    function _transferThrottleHasThrottleAccess(address _address) internal view override {
        require(accessController.hasMonitorCapabilities(_address), "L2ToL1MessagePasser: sender not allowed to throttle");
    }

    /// @notice Returns the amount of eth that can be withdrawn before being throttled
    function getEthThrottleWithdrawalsCredits() external view returns (uint256 availableCredits) {
        availableCredits = _throttleUserAvailableCredits(_throttleGlobalUser, ethThrottleWithdrawals);
    }

    /// @notice Updates the max amount per period for the withdrawals throttle, impacting the current period
    function setEthThrottleWithdrawalsMaxAmount(uint208 maxAmountPerPeriod, uint256 maxAmountTotal) external {
        // setting a maximum amount for withdrawals doesn't make any sense
        require(maxAmountTotal == 0, "L2ToL1MessagePasser: max total amount not supported");
        _setThrottle(maxAmountPerPeriod, maxAmountTotal, ethThrottleWithdrawals);
    }

    /// @notice Sets the length of the withdrawals throttle period to `_periodLength`, which
    ///         immediately affects the speed of credit accumulation.
    function setEthThrottleWithdrawalsPeriodLength(uint48 _periodLength) external {
        _setPeriodLength(_periodLength, ethThrottleWithdrawals);
    }
}
