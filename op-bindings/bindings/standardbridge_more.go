// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/zircuit-labs/zkr-monorepo-public/op-bindings/solc"
)

const StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/universal/StandardBridge.sol:StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1001,\"contract\":\"src/universal/StandardBridge.sol:StandardBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_contract(CrossDomainMessenger)1006\"},{\"astId\":1002,\"contract\":\"src/universal/StandardBridge.sol:StandardBridge\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_contract(StandardBridge)1007\"},{\"astId\":1003,\"contract\":\"src/universal/StandardBridge.sol:StandardBridge\",\"label\":\"accessController\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(AccessControlPausable)1005\"},{\"astId\":1004,\"contract\":\"src/universal/StandardBridge.sol:StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_array(t_uint256)45_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)45_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[45]\",\"numberOfBytes\":\"1440\",\"base\":\"t_uint256\"},\"t_contract(AccessControlPausable)1005\":{\"encoding\":\"inplace\",\"label\":\"contract AccessControlPausable\",\"numberOfBytes\":\"20\"},\"t_contract(CrossDomainMessenger)1006\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(StandardBridge)1007\":{\"encoding\":\"inplace\",\"label\":\"contract StandardBridge\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var StandardBridgeStorageLayout = new(solc.StorageLayout)

var StandardBridgeDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(StandardBridgeStorageLayoutJSON), StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["StandardBridge"] = StandardBridgeStorageLayout
	deployedBytecodes["StandardBridge"] = StandardBridgeDeployedBin
	immutableReferences["StandardBridge"] = false
}
