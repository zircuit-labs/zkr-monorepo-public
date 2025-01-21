/** @type import('hardhat/config').HardhatUserConfig */
require('@nomicfoundation/hardhat-foundry');
require('@nomicfoundation/hardhat-ignition');
require('@nomiclabs/hardhat-ethers');

const { HardhatUserConfig } = require('hardhat/config');
require('dotenv').config();

console.log(process.env.mnemonic);

module.exports = {
  solidity: {
    compilers: [
      {
        version: '0.8.20',
      },
      {
        version: '0.8.19',
      },
      {
        version: '0.8.15',
      },
      {
        version: '0.4.22',
      },
    ],
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    'alphanet-sepolia': {
      url: process.env.rpcL2,
      accounts: {
        mnemonic: process.env.mnemonic,
        initialIndex: parseInt(process.env.accountStartIndex ?? '5', 10),
      },
    },
    'betanet-sepolia': {
      url: process.env.rpcL2,
      accounts: {
        mnemonic: process.env.mnemonic,
        initialIndex: parseInt(process.env.accountStartIndex ?? '5', 10),
      },
    },
    'testnet-sepolia': {
      url: process.env.rpcL2,
      accounts: {
        mnemonic: process.env.mnemonic,
        initialIndex: parseInt(process.env.accountStartIndex ?? '5', 10),
      },
    },
    'mainnet-sepolia': {
      url: process.env.rpcL2,
      accounts: {
        mnemonic: process.env.mnemonic,
        initialIndex: parseInt(process.env.accountStartIndex ?? '5', 10),
      },
    },
    mainnet: {
      url: process.env.rpcL2,
      accounts: {
        mnemonic: process.env.mnemonic,
        initialIndex: parseInt(process.env.accountStartIndex ?? '5', 10),
      },
    },
  },
  etherscan: {
    apiKey: {
      'alphanet-sepolia': process.env.explorerApiKey,
      'betanet-sepolia': process.env.explorerApiKey,
      'testnet-sepolia': process.env.explorerApiKey,
      'mainnet-sepolia': process.env.explorerApiKey,
      mainnet: process.env.explorerApiKey,
    },
    customChains: [
      {
        network: 'alphanet-sepolia',
        chainId: 47777,
        urls: {
          apiURL: process.env.explorerVerificationUrl,
        },
      },
      {
        network: 'betanet-sepolia',
        chainId: 48888,
        urls: {
          apiURL: process.env.explorerVerificationUrl,
        },
      },
      {
        network: 'testnet-sepolia',
        chainId: 48899,
        urls: {
          apiURL: process.env.explorerVerificationUrl,
        },
      },
      {
        network: 'mainnet-sepolia',
        chainId: 299792,
        urls: {
          apiURL: process.env.explorerVerificationUrl,
        },
      },
      {
        network: 'mainnet',
        chainId: 48900,
        urls: {
          apiURL: process.env.explorerVerificationUrl,
        },
      },
    ],
  },
};
