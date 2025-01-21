import { exec } from 'child_process';

import { ethers } from 'ethers';

const EXPLORER_URL = process.env.EXPLORER_URL;
const API_KEY = process.env.API_KEY;
const provider = new ethers.providers.JsonRpcProvider(process.env.RPC_URL);
const implementationStorageSlot =
  '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'; // see src/universal/Proxy.sol:Proxy

type Contract = {
  path: string;
  address: string;
  optimizerRuns: number;
};

const proxyContracts: Contract[] = [
  {
    address: '0x4200000000000000000000000000000000000016',
    optimizerRuns: 200,
    path: 'src/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser',
  },
  {
    address: '0x4200000000000000000000000000000000000018',
    optimizerRuns: 200,
    path: 'src/universal/ProxyAdmin.sol:ProxyAdmin',
  },
  {
    address: '0x4200000000000000000000000000000000000007',
    optimizerRuns: 200,
    path: 'src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger',
  },
  {
    address: '0x4200000000000000000000000000000000000010',
    optimizerRuns: 200,
    path: 'src/L2/L2StandardBridge.sol:L2StandardBridge',
  },
  {
    address: '0x4200000000000000000000000000000000000011',
    optimizerRuns: 200,
    path: 'src/L2/SequencerFeeVault.sol:SequencerFeeVault',
  },
  {
    address: '0x4200000000000000000000000000000000000012',
    optimizerRuns: 200,
    path: 'src/universal/OptimismMintableERC20Factory.sol:OptimismMintableERC20Factory',
  },
  {
    address: '0x420000000000000000000000000000000000000F',
    optimizerRuns: 200,
    path: 'src/L2/GasPriceOracle.sol:GasPriceOracle',
  },
  {
    address: '0x4200000000000000000000000000000000000015',
    optimizerRuns: 200,
    path: 'src/L2/L1Block.sol:L1Block',
  },
  {
    address: '0x4200000000000000000000000000000000000014',
    optimizerRuns: 200,
    path: 'src/L2/L2ERC721Bridge.sol:L2ERC721Bridge',
  },
  {
    address: '0x4200000000000000000000000000000000000017',
    optimizerRuns: 200,
    path: 'src/universal/OptimismMintableERC721Factory.sol:OptimismMintableERC721Factory',
  },
  {
    address: '0x4200000000000000000000000000000000000019',
    optimizerRuns: 10000,
    path: 'src/L2/BaseFeeVault.sol:BaseFeeVault',
  },
  {
    address: '0x420000000000000000000000000000000000001a',
    optimizerRuns: 200,
    path: 'src/L2/L1FeeVault.sol:L1FeeVault',
  },
];

const nonProxyContracts: Contract[] = [
  {
    path: 'src/vendor/WETH9.sol:WETH9',
    address: '0x4200000000000000000000000000000000000006',
    optimizerRuns: 200,
  },
];

const verifyContract = async ({ address, path, optimizerRuns }: Contract) => {
  console.log('Verifying contract: ', address, path);
  const execString = `forge verify-contract \
        --verifier-url ${EXPLORER_URL} \
        ${address} \
        ${path} \
        --root . \
        --optimizer-runs ${optimizerRuns} \
        --etherscan-api-key ${API_KEY}
    `;
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  exec(execString, (err, output) => {
    if (err) {
      console.log('Failed verification: ', address, path);
    } else {
      console.log('Successfully verified: ', address, path);
    }
  });
};

const main = async () => {
  const proxyContractsImplementations = await Promise.all(
    proxyContracts.map(async (contract: Contract) => ({
      ...contract,
      address: ethers.utils.hexStripZeros(
        await provider.getStorageAt(contract.address, implementationStorageSlot)
      ),
    }))
  );

  const contracts: Contract[] = proxyContracts
    .map((contract) => ({
      path: 'src/universal/Proxy.sol:Proxy',
      address: contract.address,
      optimizerRuns: 200,
    }))
    .concat(proxyContractsImplementations)
    .concat(nonProxyContracts);

  contracts.map(verifyContract);
};
main();
