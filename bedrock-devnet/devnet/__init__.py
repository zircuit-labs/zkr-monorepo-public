import argparse
import logging
import os
import subprocess
import json
import socket
import calendar
import datetime
import time
import shutil
import http.client
import gzip
from multiprocessing import Process, Queue
import concurrent.futures
from collections import namedtuple


import devnet.log_setup

pjoin = os.path.join

parser = argparse.ArgumentParser(description='Bedrock devnet launcher')
parser.add_argument('--monorepo-dir', help='Directory of the monorepo', default=os.getcwd())
parser.add_argument('--allocs', help='Only create the allocs and exit', type=bool, action=argparse.BooleanOptionalAction)
parser.add_argument('--genesis-l1', help='Only create the l1 genesis and exit. Requires allocs to exist', type=bool, action=argparse.BooleanOptionalAction)
parser.add_argument('--genesis-l2', help='Only create the l2 genesis and exit. Requires l1 deployments and starting block to exist', type=bool, action=argparse.BooleanOptionalAction)

log = logging.getLogger()

class Bunch:
    def __init__(self, **kwds):
        self.__dict__.update(kwds)

class ChildProcess:
    def __init__(self, func, *args):
        self.errq = Queue()
        self.process = Process(target=self._func, args=(func, args))

    def _func(self, func, args):
        try:
            func(*args)
        except Exception as e:
            self.errq.put(str(e))

    def start(self):
        self.process.start()

    def join(self):
        self.process.join()

    def get_error(self):
        return self.errq.get() if not self.errq.empty() else None


def main():
    args = parser.parse_args()

    monorepo_dir = os.path.abspath(args.monorepo_dir)
    devnet_dir = pjoin(monorepo_dir, '.devnet')
    contracts_bedrock_dir = pjoin(monorepo_dir, 'packages', 'contracts-bedrock')
    deployment_dir = pjoin(contracts_bedrock_dir, 'deployments', 'devnetL1')
    forge_l1_dump_path = pjoin(contracts_bedrock_dir, 'state-dump-900.json')
    op_node_dir = pjoin(args.monorepo_dir, 'op-node')
    ops_bedrock_dir = pjoin(monorepo_dir, 'ops-bedrock')
    deploy_config_dir = pjoin(contracts_bedrock_dir, 'deploy-config')
    devnet_config_path = pjoin(deploy_config_dir, 'devnetL1.json')
    devnet_config_template_path = pjoin(deploy_config_dir, 'devnetL1-template.json')
    ops_chain_ops = pjoin(monorepo_dir, 'op-chain-ops')
    sdk_dir = pjoin(monorepo_dir, 'packages', 'sdk')

    paths = Bunch(
      mono_repo_dir=monorepo_dir,
      devnet_dir=devnet_dir,
      contracts_bedrock_dir=contracts_bedrock_dir,
      deployment_dir=deployment_dir,
      forge_l1_dump_path=forge_l1_dump_path,
      l1_deployments_path=pjoin(deployment_dir, '.deploy'),
      deploy_config_dir=deploy_config_dir,
      devnet_config_path=devnet_config_path,
      devnet_config_template_path=devnet_config_template_path,
      op_node_dir=op_node_dir,
      ops_bedrock_dir=ops_bedrock_dir,
      ops_chain_ops=ops_chain_ops,
      sdk_dir=sdk_dir,
      genesis_l1_path=pjoin(devnet_dir, 'genesis-l1.json'),
      genesis_l2_path=pjoin(devnet_dir, 'genesis-l2.json'),
      allocs_l1_path=pjoin(devnet_dir, 'allocs-l1.json'),
      addresses_json_path=pjoin(devnet_dir, 'addresses.json'),
      sdk_addresses_json_path=pjoin(devnet_dir, 'sdk-addresses.json'),
      rollup_config_path=pjoin(devnet_dir, 'rollup.json')
    )

    os.makedirs(devnet_dir, exist_ok=True)

    if args.allocs:
        devnet_l1_allocs(paths)
        devnet_l2_allocs(paths)
        return

    if args.genesis_l1:
        devnet_l1_genesis(paths)
        return

    if args.genesis_l2:
        devnet_l2_genesis(paths)
        return

    raise Exception("Need to specify at least one of --alloc, --l1-genesis, --l2-genesis")


def init_devnet_l1_deploy_config(paths, update_timestamp=False):
    deploy_config = read_json(paths.devnet_config_template_path)
    if update_timestamp:
        deploy_config['l1GenesisBlockTimestamp'] = '{:#x}'.format(int(time.time()))
    write_json(paths.devnet_config_path, deploy_config)

def devnet_l1_allocs(paths):
    log.info('Generating L1 genesis allocs')
    init_devnet_l1_deploy_config(paths)

    fqn = 'scripts/Deploy.s.sol:Deploy'
    # Use foundry pre-funded account #1 for the deployer
    run_command([
        'forge', 'script', '--chain-id', '900', fqn, "--sig", "runWithStateDump()", "--private-key", "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
    ], env={}, cwd=paths.contracts_bedrock_dir)

    shutil.move(src=paths.forge_l1_dump_path, dst=paths.allocs_l1_path)

    shutil.copy(paths.l1_deployments_path, paths.addresses_json_path)

    # call sync to create hardhat artifacts
    hardhat_artifacts(paths)

def hardhat_artifacts(paths):
    fqn = 'scripts/Deploy.s.sol:Deploy'
    # call sync to create hardhat artifacts
    run_command([
        'forge', 'script', '--chain-id', '900', fqn,
        '--unlocked', '--with-gas-price', '100000000000',
        '--sig', 'sync()'
    ], env={}, cwd=paths.contracts_bedrock_dir)


def devnet_l2_allocs(paths):
    log.info('Generating L2 genesis allocs, with L1 addresses: '+paths.l1_deployments_path)

    fqn = 'scripts/L2Genesis.s.sol:L2Genesis'
    # Use foundry pre-funded account #1 for the deployer
    run_command([
        'forge', 'script', '--chain-id', '901', fqn, "--sig", "runWithAllUpgrades()", "--private-key", "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
    ], env={
      'CONTRACT_ADDRESSES_PATH': paths.l1_deployments_path,
    }, cwd=paths.contracts_bedrock_dir)

    # For the previous forks, and the latest fork (default, thus empty prefix),
    # move the forge-dumps into place as .devnet allocs.
    for suffix in ["-delta", ""]:
        input_path = pjoin(paths.contracts_bedrock_dir, f"state-dump-901{suffix}.json")
        output_path = pjoin(paths.devnet_dir, f'allocs-l2{suffix}.json')
        shutil.move(src=input_path, dst=output_path)
        log.info("Generated L2 allocs: "+output_path)

def devnet_l1_genesis(paths):
    log.info('Generating L1 genesis.')
    # require that allocs have been generated via earthly first
    if os.path.exists(paths.allocs_l1_path) == False:
        raise Exception(f"{paths.allocs_l1_path} does not exist, generate allocs first")

    # fill in the devnetL1-template with the current timestamp
    init_devnet_l1_deploy_config(paths, update_timestamp=True)

    run_command([
        'go', 'run', 'cmd/main.go', 'genesis', 'l1',
        '--deploy-config', paths.devnet_config_path,
        '--l1-allocs', paths.allocs_l1_path,
        '--l1-deployments', paths.addresses_json_path,
        '--outfile.l1', paths.genesis_l1_path,
    ], cwd=paths.op_node_dir)

    genesis_consensus_path = pjoin(paths.ops_bedrock_dir, 'genesis_consensus.json')
    # combining these is likely also possible in python but using jq for now
    # since it works
    cmd = ['jq', '-s', '.[0] * .[1]',
                 genesis_consensus_path,
                 paths.genesis_l1_path]
    pjq = run_command(cmd, capture_output=True, check=False)
    if pjq.returncode != 0:
        print(pjq.stdout)
        print(pjq.stderr)
        raise Exception(f"{cmd} did not complete successfully. Exit code {pjq.returncode}")
    genesis_l1 = json.loads(pjq.stdout)
    genesis_consensus = read_json(genesis_consensus_path)
    # overwrite the extra data since merging them doesn't do what we want
    genesis_l1['extraData'] = genesis_consensus['extraData']
    # write out the genesis l1
    write_json(paths.genesis_l1_path, genesis_l1)

# Bring up the devnet where the contracts are deployed to L1
def devnet_l2_genesis(paths):
    log.info('Generating L2 genesis and rollup configs.')
    l2_allocs_path = pjoin(paths.devnet_dir, 'allocs-l2.json')
    run_command([
        'go', 'run', 'cmd/main.go', 'genesis', 'l2',
        '--l1-starting-block', pjoin(paths.devnet_dir, 'l1-starting-block.json'),
        '--deploy-config', paths.devnet_config_path,
        '--l2-allocs', l2_allocs_path,
        '--l1-deployments', paths.addresses_json_path,
        '--outfile.l2', paths.genesis_l2_path,
        '--outfile.rollup', paths.rollup_config_path
    ], cwd=paths.op_node_dir)

    rollup_config = read_json(paths.rollup_config_path)
    addresses = read_json(paths.addresses_json_path)


def eth_accounts(url):
    log.info(f'Fetch eth_accounts {url}')
    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":2, "jsonrpc":"2.0", "method": "eth_accounts", "params":[]}'
    conn.request('POST', '/', body, headers)
    response = conn.getresponse()
    data = response.read().decode()
    conn.close()
    return data


def anvil_dumpState(url):
    log.info(f'Fetch debug_dumpBlock {url}')
    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":3, "jsonrpc":"2.0", "method": "anvil_dumpState", "params":[]}'
    conn.request('POST', '/', body, headers)
    data = conn.getresponse().read()
    # Anvil returns a JSON-RPC response with a hex-encoded "result" field
    result = json.loads(data.decode('utf-8'))['result']
    result_bytes = bytes.fromhex(result[2:])
    uncompressed = gzip.decompress(result_bytes).decode()
    return json.loads(uncompressed)

def convert_anvil_dump(dump):
    accounts = dump['accounts']

    for account in accounts.values():
        bal = account['balance']
        account['balance'] = str(int(bal, 16))

        if 'storage' in account:
            storage = account['storage']
            storage_keys = list(storage.keys())
            for key in storage_keys:
                value = storage[key]
                del storage[key]
                storage[pad_hex(key)] = pad_hex(value)

    return dump

def pad_hex(input):
    return '0x' + input.replace('0x', '').zfill(64)

def wait_for_rpc_server(url):
    log.info(f'Waiting for RPC server at {url}')

    headers = {'Content-type': 'application/json'}
    body = '{"id":1, "jsonrpc":"2.0", "method": "eth_chainId", "params":[]}'

    while True:
        try:
            conn = http.client.HTTPConnection(url)
            conn.request('POST', '/', body, headers)
            response = conn.getresponse()
            if response.status < 300:
                log.info(f'RPC server at {url} ready')
                return
        except Exception as e:
            log.info(f'Waiting for RPC server at {url}')
            time.sleep(1)
        finally:
            if conn:
                conn.close()


CommandPreset = namedtuple('Command', ['name', 'args', 'cwd', 'timeout'])

def run_commands(commands: list[CommandPreset], max_workers=2):
    with concurrent.futures.ThreadPoolExecutor(max_workers=max_workers) as executor:
        futures = [executor.submit(run_command_preset, cmd) for cmd in commands]

        for future in concurrent.futures.as_completed(futures):
            result = future.result()
            if result:
                print(result.stdout)


def run_command_preset(command: CommandPreset):
    with subprocess.Popen(command.args, cwd=command.cwd,
                          stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True) as proc:
        try:
            # Live output processing
            for line in proc.stdout:
                # Annotate and print the line with timestamp and command name
                timestamp = datetime.datetime.utcnow().strftime('%H:%M:%S.%f')
                # Annotate and print the line with the timestamp
                print(f"[{timestamp}][{command.name}] {line}", end='')

            stdout, stderr = proc.communicate(timeout=command.timeout)

            if proc.returncode != 0:
                raise RuntimeError(f"Command '{' '.join(command.args)}' failed with return code {proc.returncode}: {stderr}")

        except subprocess.TimeoutExpired:
            raise RuntimeError(f"Command '{' '.join(command.args)}' timed out!")

        except Exception as e:
            raise RuntimeError(f"Error executing '{' '.join(command.args)}': {e}")

        finally:
            # Ensure process is terminated
            proc.kill()
    return proc.returncode


def run_command(args, check=True, shell=False, cwd=None, env=None, timeout=None, capture_output=False) -> subprocess.CompletedProcess:
    env = env if env else {}
    return subprocess.run(
        args,
        check=check,
        shell=shell,
        env={
            **os.environ,
            **env
        },
        cwd=cwd,
        timeout=timeout,
        capture_output=capture_output
    )


def wait_up(port, retries=10, wait_secs=1):
    for i in range(0, retries):
        log.info(f'Trying 127.0.0.1:{port}')
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        try:
            s.connect(('127.0.0.1', int(port)))
            s.shutdown(2)
            log.info(f'Connected 127.0.0.1:{port}')
            return True
        except Exception:
            time.sleep(wait_secs)

    raise Exception(f'Timed out waiting for port {port}.')


def write_json(path, data):
    with open(path, 'w+') as f:
        json.dump(data, f, indent='  ')


def read_json(path):
    with open(path, 'r') as f:
        return json.load(f)
