// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/zircuit-labs/zkr-monorepo-public/op-bindings/solc"
)

const CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1001,\"contract\":\"src/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"src/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint240\"},{\"astId\":1003,\"contract\":\"src/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1004,\"contract\":\"src/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"otherMessenger\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_contract(CrossDomainMessenger)1006\"},{\"astId\":1005,\"contract\":\"src/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_array(t_uint256)41_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)41_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[41]\",\"numberOfBytes\":\"1312\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(CrossDomainMessenger)1006\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var CrossDomainMessengerDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(CrossDomainMessengerStorageLayoutJSON), CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["CrossDomainMessenger"] = CrossDomainMessengerStorageLayout
	deployedBytecodes["CrossDomainMessenger"] = CrossDomainMessengerDeployedBin
	immutableReferences["CrossDomainMessenger"] = false
}
