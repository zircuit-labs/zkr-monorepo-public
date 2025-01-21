#!/usr/bin/env python3

import os
import struct
import argparse
from pathlib import Path
import subprocess
from binascii import unhexlify, hexlify

SCRIPT_DIR = Path(os.path.realpath(__file__)).parent
L1BLOCK_ADDR = '0x4200000000000000000000000000000000000015'
DEBUG = False

def print_dbg(*args):
    if DEBUG:
        print(*args)

def parse_args():
    parser = argparse.ArgumentParser(description='Find the first L2 block based on an l1 origin')
    parser.add_argument('l1_block', help="L1 block to find")
    parser.add_argument('-l2', '--l2-rpc', required=True, help="L2 rpc url as passed to cast (foundry aliases are valid)")
    parser.add_argument('--print-deposit-exclusions', action='store_true', help="After finding the l1 block, print the deposit exclusions for it")

    config = parser.parse_args()

    config.l1_block = to_int(config.l1_block)

    return config

def to_int(num: str) -> int:
    if num.startswith("0x"):
      return int(num, base=16)

    return int(num)

def parse_uint64(s: str) -> int:
    # sample outputs
    # 5954006 [5.954e6]
    # 0
    if '[' not in s:
        return int(s)

    space_pos = s.find(' ')
    if space_pos < 0:
        print(f"Could not parse uint64 from {s}")
    return to_int(s[:space_pos])

def parse_bytes(s: str) -> bytes:
    if not s.startswith('0x'):
        print(f"Could not parse bytes from {s}")
    return unhexlify(s[2:])

def bitmap_to_str(data: bytes) -> str:
    if len(data) < 16 or len(data) % 8 != 0:
        print(f"bitmap_to_str: unexpected length {len(data)}")
    bitmap_len = struct.unpack(">Q", data[:8])[0]
    s = ''
    for chunk, in struct.iter_unpack(">Q", data[8:]):
        s += f'{chunk:064b}'[::-1]
    return '0b' + s[:bitmap_len]

def get_block_num(block, rpc_url):
    cmd = ['cast', 'block', str(block), '--field=number', f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    return p.stdout.decode().strip()

def get_l1block_for_l2(block_num, rpc_url):
    cmd = ['cast', 'call', L1BLOCK_ADDR, 'number()(uint64)', f'--block={block_num}', f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    block = parse_uint64(p.stdout.decode().strip())
    print_dbg(f'l1 block for {block_num} = {block}')
    return block

def get_sequencenum_for_l2(block_num, rpc_url):
    cmd = ['cast', 'call', L1BLOCK_ADDR, 'sequenceNumber()(uint64)', f'--block={block_num}', f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    seq_num = parse_uint64(p.stdout.decode().strip())
    print_dbg(f'sequence number for {block_num} = {seq_num}')
    return seq_num

def get_deposit_exclusions_for_l2(block_num, rpc_url):
    cmd = ['cast', 'call', L1BLOCK_ADDR, 'depositExclusions()(bytes)', f'--block={block_num}', f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    deposit_exclusions = parse_bytes(p.stdout.decode().strip())
    print_dbg(f'deposit exclusions for {block_num} = {deposit_exclusions}')
    return deposit_exclusions

def find_block(target_l1_block, rpc):
    lower = 0
    starting_block = to_int(get_block_num("latest", rpc))
    upper = starting_block

    l1_block_start = get_l1block_for_l2(starting_block, rpc)
    if l1_block_start < target_l1_block:
        print(f"Latest block is lower than target block {l1_block_start} < {target_l1_block}, need to wait")
        return None

    print_dbg(f"Starting search for {target_l1_block} between 0 and {starting_block}")
    while True:
        if upper - lower <= 1:
            print(f"Could not find block between 0 and {starting_block}")
            return None

        cur_l2 = lower + (upper - lower) // 2
        l1_block = get_l1block_for_l2(cur_l2, rpc)
        if l1_block == target_l1_block:
            return cur_l2
        elif l1_block > target_l1_block:
            upper = cur_l2
        else:
            lower = cur_l2

def main():
    config = parse_args()
    block = find_block(config.l1_block, config.l2_rpc)
    if block == None:
        exit(1)
    sequence_number = get_sequencenum_for_l2(block, config.l2_rpc)

    print_dbg(f"Block: {block}")
    print_dbg(f"Sequence number: {sequence_number}")
    final_block = block - sequence_number

    if get_l1block_for_l2(final_block, config.l2_rpc) != config.l1_block:
        print("Block is no longer correct after subtracting sequencer number")
        exit(1)
    print(f"{final_block}")

    if config.print_deposit_exclusions:
        deposit_exclusions = get_deposit_exclusions_for_l2(final_block, config.l2_rpc)
        if not deposit_exclusions:
            print("deposit exclusions empty")
        else:
            print(f'raw deposit exclusions: {hexlify(deposit_exclusions)}')
            print(f'parsed deposit exclusions: {bitmap_to_str(deposit_exclusions)}')



if __name__ == '__main__':
    main()
