// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/zircuit-labs/zkr-monorepo-public/op-bindings/solc"
)

const MockVerifierStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var MockVerifierStorageLayout = new(solc.StorageLayout)

var MockVerifierDeployedBin = "0x6080604052348015600f57600080fd5b506004361060285760003560e01c806354fd4d50146039575b604080516020810190915260009052005b603f6053565b604051604a919060c0565b60405180910390f35b6040518060600160405280602881526020016100d76028913981565b60005b8381101560885781810151838201526020016072565b50506000910152565b6000609a825190565b80845260208401935060af818560208601606f565b601f01601f19169290920192915050565b6020808252810160cf81846091565b939250505056fe30303030303030303030303030303030303030303030303030303030303030303030303030303030a164736f6c6343000814000a"

func init() {
	if err := json.Unmarshal([]byte(MockVerifierStorageLayoutJSON), MockVerifierStorageLayout); err != nil {
		panic(err)
	}

	layouts["MockVerifier"] = MockVerifierStorageLayout
	deployedBytecodes["MockVerifier"] = MockVerifierDeployedBin
	immutableReferences["MockVerifier"] = false
}
