// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { Math } from "@openzeppelin/contracts/utils/math/Math.sol";
import { Burn } from "src/libraries/Burn.sol";
import { Arithmetic } from "src/libraries/Arithmetic.sol";

/// @custom:upgradeable
/// @title ResourceMetering
/// @notice ResourceMetering implements an EIP-1559 style resource metering system where pricing
///         updates automatically based on current demand.
abstract contract ResourceMetering is Initializable {
    /// @notice Error returned when too much gas resource is consumed.
    error OutOfGas();

    /// @notice Represents the various parameters that control the way in which resources are
    ///         metered. Corresponds to the EIP-1559 resource metering system.
    /// @custom:field prevBaseFee   Base fee from the previous block(s).
    /// @custom:field prevBoughtGas Amount of gas bought so far in the current block.
    /// @custom:field prevBlockNum  Last block number that the base fee was updated.
    /// @custom:field prevTxCount  Number of transactions in the current block.
    struct ResourceParams {
        uint128 prevBaseFee;
        uint64 prevBoughtGas;
        uint64 prevBlockNum;
        uint16 prevTxCount;
    }

    /// @notice Represents the configuration for the EIP-1559 based curve for the deposit gas
    ///         market. These values should be set with care as it is possible to set them in
    ///         a way that breaks the deposit gas market. The target resource limit is defined as
    ///         maxResourceLimit / elasticityMultiplier. This struct was designed to fit within a
    ///         single word. There is additional space for additions in the future.
    /// @custom:field maxResourceLimit             Represents the maximum amount of deposit gas that
    ///                                            can be purchased per block.
    /// @custom:field elasticityMultiplier         Determines the target resource limit along with
    ///                                            the resource limit.
    /// @custom:field baseFeeMaxChangeDenominator  Determines max change on fee per block.
    /// @custom:field maxTransactionLimit          Determines max deposit transaction count per block.
    /// @custom:field minimumBaseFee               The min deposit base fee, it is clamped to this
    ///                                            value.
    /// @custom:field systemTxMaxGas               The amount of gas supplied to the system
    ///                                            transaction. This should be set to the same
    ///                                            number that the op-node sets as the gas limit
    ///                                            for the system transaction.
    /// @custom:field maximumBaseFee               The max deposit base fee, it is clamped to this
    ///                                            value.
    struct ResourceConfig {
        uint32 maxResourceLimit;
        uint8 elasticityMultiplier;
        uint8 baseFeeMaxChangeDenominator;
        uint16 maxTransactionLimit;
        uint32 minimumBaseFee;
        uint32 systemTxMaxGas;
        uint128 maximumBaseFee;
    }

    /// @notice EIP-1559 style gas parameters.
    ResourceParams public params;

    /// @notice Reserve extra slots (to a total of 50) in the storage layout for future upgrades.
    uint256[48] private __gap;

    event GasBurned(uint256 gasAmount, address indexed sender);

    /// @notice Meters access to a function based an amount of a requested resource.
    /// @param _amount Amount of the resource requested.
    modifier metered(uint64 _amount) {
        // Record initial gas amount so we can refund for it later.
        uint256 initialGas = gasleft();

        // Run the underlying function.
        _;

        // Run the metering function.
        _metered(_amount, initialGas);
    }

    /// @notice An internal function that holds all of the logic for metering a resource.
    /// @param _amount     Amount of the resource requested.
    /// @param _initialGas The amount of gas before any modifier execution.
    function _metered(uint64 _amount, uint256 _initialGas) internal {
        // Update block number and base fee if necessary.
        uint256 blockDiff = block.number - params.prevBlockNum;

        ResourceConfig memory config = _resourceConfig();
        int256 targetResourceLimit =
            int256(uint256(config.maxResourceLimit)) / int256(uint256(config.elasticityMultiplier));

        if (blockDiff > 0) {
            // Handle updating EIP-1559 style gas parameters. We use EIP-1559 to restrict the rate
            // at which deposits can be created and therefore limit the potential for deposits to
            // spam the L2 system. Fee scheme is very similar to EIP-1559 with minor changes.

            // If limit deposit limit is hit, increase the gas fee by max amount
            uint256 boughtGas = params.prevBoughtGas;
            if (params.prevTxCount == config.maxTransactionLimit) {
                boughtGas = Math.max(uint256(targetResourceLimit * 2), boughtGas);
            }
            int256 gasUsedDelta = int256(boughtGas) - targetResourceLimit;
            int256 baseFeeDelta = (int256(uint256(params.prevBaseFee)) * gasUsedDelta)
                / (targetResourceLimit * int256(uint256(config.baseFeeMaxChangeDenominator)));

            // Update base fee by adding the base fee delta and clamp the resulting value between
            // min and max.
            int256 newBaseFee = Arithmetic.clamp({
                _value: int256(uint256(params.prevBaseFee)) + baseFeeDelta,
                _min: int256(uint256(config.minimumBaseFee)),
                _max: int256(uint256(config.maximumBaseFee))
            });

            // If we skipped more than one block, we also need to account for every empty block.
            // Empty block means there was no demand for deposits in that block, so we should
            // reflect this lack of demand in the fee.
            if (blockDiff > 1) {
                // Update the base fee by repeatedly applying the exponent 1-(1/change_denominator)
                // blockDiff - 1 times. Simulates multiple empty blocks. Clamp the resulting value
                // between min and max.
                newBaseFee = Arithmetic.clamp({
                    _value: Arithmetic.cdexp({
                        _coefficient: newBaseFee,
                        _denominator: int256(uint256(config.baseFeeMaxChangeDenominator)),
                        _exponent: int256(blockDiff - 1)
                    }),
                    _min: int256(uint256(config.minimumBaseFee)),
                    _max: int256(uint256(config.maximumBaseFee))
                });
            }

            // Update new base fee, reset bought gas, and update block number.
            params.prevBaseFee = uint128(uint256(newBaseFee));
            params.prevBoughtGas = 0;
            params.prevTxCount = 0;
            params.prevBlockNum = uint64(block.number);
        }

        params.prevTxCount += 1;
        // If limit is surpassed,
        require(params.prevTxCount <= config.maxTransactionLimit, "ResourceMetering: too many deposits in this block");

        // Make sure we can actually buy the resource amount requested by the user.
        params.prevBoughtGas += _amount;
        if (int256(uint256(params.prevBoughtGas)) > int256(uint256(config.maxResourceLimit))) {
            revert OutOfGas();
        }

        // Determine the amount of ETH to be paid.
        uint256 resourceCost = uint256(_amount) * uint256(params.prevBaseFee);

        // We currently charge for this ETH amount as an L1 gas burn, so we convert the ETH amount
        // into gas by dividing by the L1 base fee. We assume a minimum base fee of 1 gwei to avoid
        // division by zero for L1s that don't support 1559 or to avoid excessive gas burns during
        // periods of extremely low L1 demand. One-day average gas fee hasn't dipped below 1 gwei
        // during any 1 day period in the last 5 years, so should be fine.
        uint256 gasCost = resourceCost / Math.max(block.basefee, 1 gwei);

        // Give the user a refund based on the amount of gas they used to do all of the work up to
        // this point. Since we're at the end of the modifier, this should be pretty accurate. Acts
        // effectively like a dynamic stipend (with a minimum value).
        uint256 usedGas = _initialGas - gasleft();
        if (gasCost > usedGas) {
            // We calculate gasToBurn based on the resourceCosts, but reserve some Gas for an event
            // that keeps track of the GasBurned. There we add the costs for the event because we
            // would be burning it otherwise
            uint256 gasToBurn = gasCost - usedGas;
            // Gas Costs For Event:
            // 375 Per LOG* operation.
            // 375 per indexed parameter
            // 8   Per byte in a LOG* operation's data.
            // 375 + 375 + 256 + 160 = 1166 (uint256 32 bytes, address 20bytes)
            uint256 estimatedEventGasCosts = 1200;

            // Ensure gasToBurn is greater than estimatedEventGasCosts to avoid underflow
            if (gasToBurn > estimatedEventGasCosts) {
                emit GasBurned(gasToBurn - estimatedEventGasCosts, tx.origin);
                // Subtract estimatedEventGasCosts because the event emission accounts for this already
                Burn.gas(gasToBurn - estimatedEventGasCosts);
            } else {
                // Handle case where gasToBurn is not enough to cover estimatedEventGasCosts
                // Emit event with what we have and burn zero gas
                emit GasBurned(estimatedEventGasCosts, tx.origin);
                Burn.gas(0);
            }
        }
    }

    /// @notice Virtual function that returns the resource config.
    ///         Contracts that inherit this contract must implement this function.
    /// @return ResourceConfig
    function _resourceConfig() internal virtual returns (ResourceConfig memory);

    /// @notice Sets initial resource parameter values.
    ///         This function must either be called by the initializer function of an upgradeable
    ///         child contract.
    // solhint-disable-next-line func-name-mixedcase
    function __ResourceMetering_init() internal onlyInitializing {
        if (params.prevBlockNum == 0) {
            params = ResourceParams({ prevBaseFee: 1 gwei, prevBoughtGas: 0, prevBlockNum: uint64(block.number), prevTxCount: 0 });
        }
    }
}
