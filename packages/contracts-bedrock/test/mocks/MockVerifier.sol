// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

contract MockVerifier {

    string public constant version = "0000000000000000000000000000000000000000";

    // accept everything
    fallback(bytes calldata) external returns (bytes memory) {
        return bytes("");
    }
}
