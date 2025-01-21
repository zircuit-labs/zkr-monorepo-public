// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Chain IDs for the various networks.
library Chains {
    uint256 internal constant Mainnet = 1;
    uint256 internal constant Sepolia = 11155111;
    uint256 internal constant SepoliaAlphanet = 47777;
    uint256 internal constant SepoliaBetanet = 48888;
    uint256 internal constant SepoliaBetanet2 = 48898;
    uint256 internal constant SepoliaTestnet = 48899;
    uint256 internal constant LocalDevnet = 900;
    uint256 internal constant OPLocalDevnet = 901;
    uint256 internal constant GethDevnet = 1337;
    uint256 internal constant Hardhat = 31337;
}
