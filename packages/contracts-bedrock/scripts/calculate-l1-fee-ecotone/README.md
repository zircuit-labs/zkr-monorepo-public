# ECOTONE GAS FEES CALCULATION BASED ON DEFAULT VALUES

To run the script: ```node index.js```

Check the code before running, you can decide to deploy the contract again, but there's an instance already deoloyed on testnet here: 0xf3D16573DD89413cbCd649142411C1144C337904.

Before running, set the env var: RPC_URL with the right rpc url.

This script estimates the price of a normal eth transfer and a swap under the conditions:
- L1 gas price at 30 gwei
- eth price 3800$
- blob base fee at 263 wei

Other constants were taken from this excel: https://docs.google.com/spreadsheets/d/1jt1sVTA9z1C0bCTHe_y1DzGV1-1RSGh3y2BxvIeyRGM

There's no need to compile the contract, you can use the json L1GasCalculator.json to deploy directly. In case you need to recompile, any compiler will work. I'm using hardhat 2.14.0