// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.20;

import { Test } from "forge-std/Test.sol";
import { Verifier } from "src/L1/Verifier.sol";

library VerifierHelper {
    function getGapProof() internal pure returns (bytes memory _proof) {
        _proof = abi.encodePacked(bytes2(0xDEAD));
    }

    function getFailingProof() internal pure returns (bytes memory _proof) {
        _proof = abi.encodePacked(bytes2(0xBEEF));
    }
}

contract VerifierTest is Test {
    Verifier private verifier;

    function setUp() public {
        verifier = new Verifier();
    }

    /// @notice Helper function to call the verifier since it does not use a function selector
    function verify(bytes memory _proof) internal returns (bool success) {
        vm.pauseGasMetering();
        (success,) = address(verifier).staticcall(_proof);
        vm.resumeGasMetering();
    }

    function getRealProof() internal view returns (bytes memory _proof) {
        string memory proofPath = string.concat(vm.projectRoot(), "/prover-artifacts/final_proof.json");
        string memory json = vm.readFile(proofPath);
        // the calldata is an array of uint8 values. Since parseJsonBytes expects
        // hex encoded bytes, we parse them as uint256 and convert them to bytes
        uint256[] memory rawBytes = vm.parseJsonUintArray(json, ".FinalProof.meta.calldata");
        _proof = new bytes(rawBytes.length);
        for (uint256 i = 0; i < rawBytes.length; i++) {
            _proof[i] = bytes1(uint8(rawBytes[i]));
        }
    }

    function test_realProof_succeeds() external {
        bytes memory proof = getRealProof();
        require(verify(proof), "Verifier rejected proof");
    }

    function test_tamperedProof_reverts() external {
        bytes memory proof = getRealProof();
        proof[0] ^= 0xff;
        require(verify(proof) == false, "Verifier accepted wrong proof");
    }

    function test_gapProof_succeeds() external {
        bytes memory proof = VerifierHelper.getGapProof();
        require(verify(proof), "Verifier rejected proof");
    }

    function test_tamperedProof_reverts_fuzz(uint256 _byteToFlip) public {
        bytes memory proof = getRealProof();
        uint256 proofLength = proof.length;
        proof[_byteToFlip % proofLength] ^= 0xff;
        require(verify(proof) == false, "Verifier accepted wrong proof");
    }

    function test_tamperedGapProof_reverts() external {
        bytes memory proof = VerifierHelper.getGapProof();
        proof[0] ^= 0xff;
        require(verify(proof) == false, "Verifier accepted wrong proof");
    }
}
