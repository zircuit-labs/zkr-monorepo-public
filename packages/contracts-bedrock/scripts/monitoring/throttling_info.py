#!/usr/bin/env python3

import os
import argparse
import json
from pathlib import Path
import subprocess
from datetime import datetime
from pprint import pprint
SCRIPT_DIR = Path(os.path.realpath(__file__)).parent

def parse_args():
    parser = argparse.ArgumentParser(description='Print throttling information')
    parser.add_argument("deployment_name", help="Name of the deployment in the deployments directory")
    parser.add_argument('-l1', '--l1-rpc', help="L1 rpc url as passed to cast (foundry aliases are valid)")
    parser.add_argument('-l2', '--l2-rpc', help="L2 rpc url as passed to cast (foundry aliases are valid)")

    config = parser.parse_args()
    config.deployment_path = SCRIPT_DIR / "../../deployments" / config.deployment_name
    if not config.deployment_path.exists():
        print(f'deployment_name does not exist at {config.deployment_path}')
        exit(1)

    if not config.l1_rpc and not config.l2_rpc:
        print("At least one of --l1-rpc or --l2-rpc is required")
        exit(1)

    return config

def get_deployment_address(path, name):
    with open(path / f'{name}.json', 'r') as f:
        return json.load(f)['address']

def foundry_str_to_int(s: str) -> int:
    # output of uint sometimes contains the scientific notation like this
    # 1500000000000000000000 [1.5e21]
    if '[' in s:
        s = s.split(' ')[0]
    return int(s)

def format_ether(num):
    if num > (10**15):
        return f'{(num // (10**15)) / 1000.0} eth'
    return f'{num} wei'

def format_time(seconds):
    hours = seconds // 3600
    minutes = (seconds % 3600) // 60
    seconds = seconds % 60
    return f"{hours}h {minutes}m {seconds}s"

def format_timestamp(timestamp):
    dt = datetime.fromtimestamp(timestamp)
    return f'{dt} ({timestamp})'

def get_throttle_info(addr, throttle_name, rpc_url):
    return_signature = '(uint208 maxAmountPerPeriod, uint48 periodLength, uint256 maxAmountTotal)'
    cmd = ['cast', 'call', addr,
           f'{throttle_name}() returns {return_signature}',
           f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    lines = p.stdout.decode().splitlines()
    if len(lines) != 3:
        raise Exception(f'Expected output of throttle to contain 3 lines but got "{lines}"')
    lines = [foundry_str_to_int(line) for line in lines]
    return [
            (f'{throttle_name}', ''),
            ('  Max amount per period', format_ether(lines[0])),
            ('  Period length', format_time(lines[1])),
            ('  Max amount total', format_ether(lines[2])),
    ]

def ensure_contract_exists(addr, rpc_url):
    cmd = ['cast', 'code', addr, f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    if p.stdout.decode().strip() == '0x':
        raise Exception(f'No code at {addr} for {rpc_url}. Ensure rpc url matches deployment.')

def get_l1_admin(addr, rpc_url):
    cmd = ['cast', 'call', addr, 'guardian() returns (address)', f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    return [('Admin (guardian)', p.stdout.decode().strip())]

def get_l2_admin(addr, rpc_url):
    # get the access controller
    cmd = ['cast', 'call', addr, 'accessController() returns (address)', f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    # get the admin on the controller
    cmd = ['cast', 'call', p.stdout.decode().strip(), 'defaultAdmin() returns (address)', f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    return [('Admin (defaultAdmin)', p.stdout.decode().strip())]

def get_paused(addr, rpc_url):
    cmd = ['cast', 'call', addr, 'paused() returns (bool)', f'--rpc-url={rpc_url}']
    p = subprocess.run(cmd, capture_output=True, check=True)
    return [('Paused', p.stdout.decode().strip())]

def print_info(title, info):
    min_len = max(len(i[0]) for i in info) + 2
    print(title)
    for name, value in info:
        if len(value) == 0:
            print(f'  {name}')
            continue

        padding = ' ' * (min_len - len(name))
        print(f'  {name}{padding}: {value}')

def main():
    config = parse_args()
    l1_contract_addr = get_deployment_address(config.deployment_path, "OptimismPortalProxy")
    l2_contract_addr = '0x4200000000000000000000000000000000000016'

    if config.l1_rpc:
        ensure_contract_exists(l1_contract_addr, config.l1_rpc)
        info = []
        info.extend(get_l1_admin(l1_contract_addr, config.l1_rpc))
        info.extend(get_paused(l1_contract_addr, config.l1_rpc))
        info.extend(get_throttle_info(l1_contract_addr, 'ethThrottleDeposits', config.l1_rpc))
        info.extend(get_throttle_info(l1_contract_addr, 'ethThrottleWithdrawals', config.l1_rpc))
        print_info('L1 info', info)

    if config.l2_rpc:
        ensure_contract_exists(l2_contract_addr, config.l2_rpc)
        info = []
        info.extend(get_l2_admin(l2_contract_addr, config.l2_rpc))
        info.extend(get_paused(l2_contract_addr, config.l2_rpc))
        info.extend(get_throttle_info(l2_contract_addr, 'ethThrottleWithdrawals', config.l2_rpc))
        print_info('L2 info', info)

if __name__ == '__main__':
    main()
