const { ethers } = require("ethers");
require('dotenv').config();
const contractInfo = require("./L1GasCalculator.json")

const baseFeeScalar = ethers.BigNumber.from("2735") // from excel
const l1BaseFee = ethers.BigNumber.from("30000000000") // indicative selected value of 30 GWEI
const blobBaseFeeScalar = ethers.BigNumber.from("819492") // from excel
const blobBaseFee = ethers.BigNumber.from("263")
const ethPrice = 3800;

async function main() {
    // Replace these values with the actual values
    const provider = new ethers.providers.JsonRpcProvider(process.env.RPC_URL)

    // If the contract is not deployed, deploy again using the following 2 lines
    // const contractFactory = new ethers.ContractFactory(contractInfo.abi, contractInfo.bytecode, wallet)
    // const l1GasCalculator = await contractFactory.deploy()
    const l1GasCalculator = new ethers.Contract("0xf3D16573DD89413cbCd649142411C1144C337904", contractInfo.abi, provider) // comment this line if you're deploying
    console.log("Contract: ", l1GasCalculator.address)

    // Define the transaction object
    const tx = {
        nonce: await provider.getTransactionCount("0x13aa3dfF556D04F4c4530b6F05E88b9a900145C4"),
        gasPrice: await provider.getGasPrice(),
        gasLimit: ethers.utils.hexlify(21000),
        to: "0x000000000000000000000000000000000000dead",
        value: ethers.utils.parseEther("0.1"),
        data: "0x",
        chainId: (await provider.getNetwork()).chainId
    };

    const uniswapTx = {
        nonce: await provider.getTransactionCount("0x13aa3dfF556D04F4c4530b6F05E88b9a900145C4"),
        gasPrice: await provider.getGasPrice(),
        gasLimit: ethers.utils.hexlify(144323),
        to: "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
        data: "0x38ed1739000000000000000000000000000000000000000000000b95e6d7347ae6d30000000000000000000000000000000000000000000000000000000000006029b74000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000b1b2d032aa2f52347fbcfd08e5c3cc55216e840400000000000000000000000000000000000000000000000000000000667aa74000000000000000000000000000000000000000000000000000000000000000020000000000000000000000006468e79a80c0eab0f9a2b574c8d5bc374af59414000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
        chainId: (await provider.getNetwork()).chainId,
        value: 0
    }

    // Serialize the transaction
    {
        console.log("NORMAL TRANSFER")
        const serializedTransaction = ethers.utils.serializeTransaction(tx)
        console.log("Serialized Transaction:", serializedTransaction)

        const l1Fee = await l1GasCalculator.getL1FeeEcotone(serializedTransaction, l1BaseFee, blobBaseFee, baseFeeScalar, blobBaseFeeScalar)
        console.log("L1 fee in wei: ", l1Fee.toString())
        const l1FeeToEth = ethers.utils.formatEther(l1Fee)
        console.log("L1 fee in ETH", l1FeeToEth)
        const l1FeeToDollars = l1FeeToEth * ethPrice
        console.log("L1 fee in dollars", l1FeeToDollars)
    }

    {
        console.log("UNISWAP")
        const serializedTransaction = ethers.utils.serializeTransaction(uniswapTx)
        console.log("Serialized Transaction:", serializedTransaction)

        const l1Fee = await l1GasCalculator.getL1FeeEcotone(serializedTransaction, l1BaseFee, blobBaseFee, baseFeeScalar, blobBaseFeeScalar)
        console.log("L1 fee in wei: ", l1Fee.toString())
        const l1FeeToEth = ethers.utils.formatEther(l1Fee)
        console.log("L1 fee in ETH", l1FeeToEth)
        const l1FeeToDollars = l1FeeToEth * ethPrice
        console.log("L1 fee in dollars", l1FeeToDollars)
    }
}

main().catch(error => {
    console.error("Error:", error);
});
