const { buildModule } = require('@nomicfoundation/hardhat-ignition/modules');

module.exports = buildModule('OptimismMintableERC20Module', (m) => {
  const deployment = m.contract('OptimismMintableERC20', [
    '0x4200000000000000000000000000000000000010',
    '0x64d76c85Be874cA002F2973E1d66163657c6268a',
    'Zircuit Test WBTC',
    'WBTC',
    8,
  ]);

  return { deployment };
});
