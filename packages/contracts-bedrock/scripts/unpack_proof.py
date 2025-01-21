#!/usr/bin/env python3

import argparse
from base64 import b64decode
import json

def unpack_uint256_array(data: bytes) -> list[int]:
    output = []
    for i in range(0, len(data), 32):
        slice = data[i:i+32]
        # pad to full length
        if len(slice) < 32:
            slice += b'\x00' * (32 - len(slice))

        num = 0
        for j in range(0, 32):
            num = (num << 8) + slice[31 - j]
        output.append(num)
    return output

def bytes_to_sol_input(proof: bytes, accumulator: bytes):
    unpacked_proof = unpack_uint256_array(proof)
    unpacked_accumulator = unpack_uint256_array(accumulator)

    return f'{array_to_hex(unpacked_proof)} {array_to_hex(unpacked_accumulator)}'

def json_to_sol_input(data):
    proof = b64decode(data['proof'])
    accumulator = b64decode(data['accumulator'])
    return bytes_to_sol_input(proof, accumulator)

def json_to_sol_input_raw(data):
    proof = bytes(data['proof'])
    accumulator = bytes(data['accumulator'])
    return bytes_to_sol_input(proof, accumulator)

def array_to_hex(arr) -> str:
    s = '['
    s += ','.join((hex(num) for num in arr))
    s += ']'
    return s

def main():
    parser = argparse.ArgumentParser(
            description='Read proofs taken from the DB or elsewhere and unpack them to uint256 arrays for usage with cast/solidity',
            epilog='example: cast call $(cat deployments/betanet-sepolia/VerifierProxy.json | jq -r ".address" ) "verify(uint256[] calldata proof, uint256[] calldata target_circuit_final_pair)" --rpc-url sepolia $(python scripts/unpack_proof.py /tmp/proof.json)'
        )
    parser.add_argument('json_file', help='json inputs can be taken from the DB: proof_orchestrator => batches => proof')
    parser.add_argument('--raw', action='store_true', help='json data is not base64 encoded but raw bytes')

    args = parser.parse_args()
    with open(args.json_file, 'r') as f:
        json_input = json.load(f)
    if args.raw:
        print(json_to_sol_input_raw(json_input), end='')
    else:
        print(json_to_sol_input(json_input), end='')

if __name__ == '__main__':
    main()
