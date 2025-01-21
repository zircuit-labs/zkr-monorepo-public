// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

/// @notice This contract is a simplified version of
/// packages/contracts-bedrock/src/L2/GasPriceOracle.sol
contract L1GasCalculator {
    /// @notice Number of decimals used in the scalar.
    uint256 public constant DECIMALS = 6;

    /// @notice this function replicates the original one "getL1FeeEcotone"
    /// from the GasPriceOracle.sol contract, allowing to specify some parameters
    /// that are otherwise constant or provided by a system address.
    function getL1FeeEcotone(
        bytes memory _data,
        uint256 l1BaseFee,
        uint256 blobBaseFee,
        uint32 baseFeeScalar,
        uint32 blobBaseFeeScalar
    ) external pure returns (uint256) {
        uint256 l1GasUsed = _getCalldataGas(_data);
        uint256 scaledBaseFee = baseFeeScalar * 16 * l1BaseFee;
        uint256 scaledBlobBaseFee = blobBaseFeeScalar * blobBaseFee;
        uint256 fee = l1GasUsed * (scaledBaseFee + scaledBlobBaseFee);
        return fee / (16 * 10 ** DECIMALS);
    }

    /// @notice L1 gas estimation calculation.
    /// @param _data Unsigned fully RLP-encoded transaction to get the L1 gas for.
    /// @return Amount of L1 gas used to publish the transaction.
    function _getCalldataGas(bytes memory _data) internal pure returns (uint256) {
        uint256 total = 0;
        uint256 length = _data.length;
        for (uint256 i = 0; i < length; i++) {
            if (_data[i] == 0) {
                total += 4;
            } else {
                total += 16;
            }
        }
        return total + (68 * 16);
    }
}
