package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zircuit-labs/zkr-monorepo-public/op-bindings/predeploys"
)

// Contract represents the structure of the JSON file.
// We only care about the 'address' field; other fields will be ignored.
type Contract struct {
	Address string `json:"address"`
}

type DeploymentInfo struct {
	// name of the directory in the `deployments` dir
	DirectoryName string
	// name of the testnet as it will appear in the docs
	Name string
	// name of the L1 chain
	L1Name string
	// url where L1 addresses will be concatenated (without trailing slash)
	L1BaseUrl string
	// url where L2 addresses will be concatenated (without trailing slash)
	L2BaseUrl string
}

// Deployments Information about which deployments to parse and in which order
var Deployments = []DeploymentInfo{
	{
		DirectoryName: "testnet-sepolia",
		Name:          "Testnet",
		L1Name:        "Sepolia",
		L1BaseUrl:     "https://sepolia.etherscan.io/address",
		L2BaseUrl:     "https://explorer.testnet.zircuit.com/address",
	},
	{
		DirectoryName: "mainnet",
		Name:          "Mainnet",
		L1Name:        "Ethereum",
		L1BaseUrl:     "https://etherscan.io/address",
		L2BaseUrl:     "https://explorer.zircuit.com/address",
	},
}

type ContractEntry struct {
	contractName  string
	displayString string
}

type contractGenerator func() map[string]string

var (
	L1Entries           = make([]ContractEntry, 0)
	L2Entries           = make([]ContractEntry, 0)
	L2PreinstallEntries = make([]ContractEntry, 0)
)

func init() {
	// only contracts that are included here will be part of the table
	// it's possible to rename them by setting the value to a non-empty string
	// if no name is specified, the name is surrounded by verbatim quotes
	L1Entries = append(L1Entries, ContractEntry{"L1CrossDomainMessengerProxy", ""})
	L1Entries = append(L1Entries, ContractEntry{"L1StandardBridgeProxy", ""})
	L1Entries = append(L1Entries, ContractEntry{"ProxyAdmin", ""})
	L1Entries = append(L1Entries, ContractEntry{"SystemConfigProxy", ""})
	L1Entries = append(L1Entries, ContractEntry{"VerifierProxy", ""})
	L1Entries = append(L1Entries, ContractEntry{"OptimismPortalProxy", ""})
	L1Entries = append(L1Entries, ContractEntry{"OptimismMintableERC20FactoryProxy", ""})
	L1Entries = append(L1Entries, ContractEntry{"L2OutputOracleProxy", ""})
	L1Entries = append(L1Entries, ContractEntry{"L1ERC721BridgeProxy", ""})

	L2Entries = append(L2Entries, ContractEntry{"BaseFeeVault", ""})
	L2Entries = append(L2Entries, ContractEntry{"GasPriceOracle", ""})
	L2Entries = append(L2Entries, ContractEntry{"L1Block", ""})
	L2Entries = append(L2Entries, ContractEntry{"L1FeeVault", ""})
	L2Entries = append(L2Entries, ContractEntry{"L2CrossDomainMessenger", ""})
	L2Entries = append(L2Entries, ContractEntry{"L2ERC721Bridge", ""})
	L2Entries = append(L2Entries, ContractEntry{"L2StandardBridge", ""})
	L2Entries = append(L2Entries, ContractEntry{"L2ToL1MessagePasser", ""})
	L2Entries = append(L2Entries, ContractEntry{"OptimismMintableERC20Factory", ""})
	L2Entries = append(L2Entries, ContractEntry{"OptimismMintableERC721Factory", ""})
	L2Entries = append(L2Entries, ContractEntry{"ProxyAdmin", ""})
	L2Entries = append(L2Entries, ContractEntry{"SequencerFeeVault", ""})
	L2Entries = append(L2Entries, ContractEntry{"WETH9", ""})
	L2Entries = append(L2Entries, ContractEntry{"SchemaRegistry", ""})
	L2Entries = append(L2Entries, ContractEntry{"EAS", ""})

	// Preinstalls
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"Safe_v130", "[`Safe`](https://github.com/safe-global/safe-smart-account/blob/v1.3.0/contracts/GnosisSafe.sol)"})
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"SafeL2_v130", "[`SafeL2`](https://github.com/safe-global/safe-smart-account/blob/v1.3.0/contracts/GnosisSafeL2.sol)"})
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"MultiCall3", "[`Multicall3`](https://github.com/mds1/multicall/tree/main)"})
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"MultiSend_v130", "[`MultiSend`](https://github.com/safe-global/safe-smart-account/blob/v1.3.0/contracts/libraries/MultiSend.sol)"})
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"MultiSendCallOnly_v130", "[`MultiSendCallOnly`](https://github.com/safe-global/safe-smart-account/blob/v1.3.0/contracts/libraries/MultiSendCallOnly.sol)"})

	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"SafeSingletonFactory", "[`SafeSingletonFactory`](https://github.com/safe-global/safe-singleton-factory/blob/main/source/deterministic-deployment-proxy.yul)"})
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"DeterministicDeploymentProxy", "[`DeterministicDeploymentProxy`](https://github.com/Arachnid/deterministic-deployment-proxy)"})
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"Create2Deployer", "[`create2deployer`](https://github.com/pcaversaccio/create2deployer)"})
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"Permit2", "[`permit2`](https://github.com/Uniswap/permit2)"})
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"EntryPoint", "[`ERC-4337 EntryPoint`](https://github.com/eth-infinitism/account-abstraction/blob/v0.6.0/contracts/core/EntryPoint.sol)"})
	L2PreinstallEntries = append(L2PreinstallEntries, ContractEntry{"SenderCreator", "[`ERC-4337 SenderCreator`](https://github.com/eth-infinitism/account-abstraction/blob/v0.6.0/contracts/core/SenderCreator.sol)"})
}

func writeTable(outputBuilder *strings.Builder, tableName string, baseUrl string, includedContracts []ContractEntry, genFn contractGenerator) {
	outputBuilder.WriteString(fmt.Sprintf("### %s\n\n", tableName))

	tableHeader := "| Contract Name    | Contract Address     |\n|------------------|----------------------|\n"
	outputBuilder.WriteString(tableHeader)
	contracts := genFn()

	for _, entry := range includedContracts {
		// check if we should include this contract
		contractAddress, exists := contracts[entry.contractName]
		if !exists {
			fmt.Fprintf(os.Stderr, "Contract '%s' does not exist\n", entry.contractName)
			os.Exit(1)
		}

		var contractName string
		// rename if a different name was specified in the map
		if entry.displayString != "" {
			contractName = entry.displayString
		} else {
			contractName = fmt.Sprintf("`%s`", entry.contractName)
		}
		row := fmt.Sprintf("| %s | [%s](%s/%s) |\n",
			contractName, contractAddress, baseUrl, contractAddress)
		outputBuilder.WriteString(row)
	}
	outputBuilder.WriteString("\n\n")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run script.go <directory> <deployment>")
		os.Exit(1)
	}

	deployment_dir := os.Args[1]
	deployment_name := os.Args[2]
	var matched_deployment *DeploymentInfo

	// Find the deployment that matches the given name
	for _, deployment := range Deployments {
		if deployment.Name == deployment_name {
			matched_deployment = &deployment
			break
		}
	}

	if matched_deployment == nil {
		fmt.Fprintf(os.Stderr, "No deployment matches the name: %s\n", deployment_name)
		os.Exit(1)
	}

	var outputBuilder strings.Builder
	outputBuilder.WriteString("---\n")
	outputBuilder.WriteString("description: >-\n")
	outputBuilder.WriteString("  This reference guide lists all the contract addresses.\n")
	outputBuilder.WriteString("---\n")

	outputBuilder.WriteString("# Contract Addresses\n\n")

	// get all files for the current deployment
	deploymentPath := filepath.Join(deployment_dir, matched_deployment.DirectoryName)
	files, err := os.ReadDir(deploymentPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading directory:", err)
		os.Exit(1)
	}

	// Adding the header with the directory name
	header := fmt.Sprintf("## %s\n\n", matched_deployment.Name)
	outputBuilder.WriteString(header)

	{
		tableName := fmt.Sprintf("%s (L1)", matched_deployment.L1Name)
		writeTable(&outputBuilder, tableName, matched_deployment.L1BaseUrl, L1Entries, func() map[string]string {
			contracts := make(map[string]string, 0)
			for _, file := range files {
				if filepath.Ext(file.Name()) != ".json" {
					continue
				}
				filePath := filepath.Join(deploymentPath, file.Name())
				fileContent, err := os.ReadFile(filePath)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error reading file:", err)
					os.Exit(1)
				}

				var contract Contract
				if err := json.Unmarshal(fileContent, &contract); err != nil {
					fmt.Fprintln(os.Stderr, "Error unmarshalling JSON:", err)
					os.Exit(1)
				}

				contractName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
				contracts[contractName] = contract.Address
			}
			return contracts
		})
	}

	{
		tableName := fmt.Sprintf("Zircuit %s (L2)", matched_deployment.Name)
		writeTable(&outputBuilder, tableName, matched_deployment.L2BaseUrl, L2Entries, func() map[string]string {
			contracts := make(map[string]string, len(predeploys.Predeploys))
			for name, predeploy := range predeploys.Predeploys {
				contracts[name] = predeploy.Address.String()
			}
			return contracts
		})
	}

	{
		tableName := fmt.Sprintf("Zircuit %s (L2 Preinstalls)", matched_deployment.Name)
		writeTable(&outputBuilder, tableName, matched_deployment.L2BaseUrl, L2PreinstallEntries, func() map[string]string {
			contracts := make(map[string]string, len(predeploys.Predeploys))
			for name, predeploy := range predeploys.Predeploys {
				contracts[name] = predeploy.Address.String()
			}
			return contracts
		})
	}

	// Print the complete header and table
	fmt.Print(outputBuilder.String())
}
