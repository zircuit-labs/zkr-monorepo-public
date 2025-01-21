// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/zircuit-labs/l2-geth-public"
	"github.com/zircuit-labs/l2-geth-public/accounts/abi"
	"github.com/zircuit-labs/l2-geth-public/accounts/abi/bind"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/core/types"
	"github.com/zircuit-labs/l2-geth-public/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SuperchainConfigMetaData contains all meta data concerning the SuperchainConfig contract.
var SuperchainConfigMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MONITOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSED_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"guardian\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasMonitorCapabilities\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasOperatorCapabilities\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_paused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"_identifier\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"paused_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"identifier\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60806040523480156200001157600080fd5b506200002161dead600062000027565b62000599565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff1615906001600160401b0316600081158015620000725750825b90506000826001600160401b031660011480156200008f5750303b155b9050811580156200009e575080155b15620000bd5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315620000ec57845460ff60401b1916680100000000000000001785555b620000fb62015180886200018d565b85156200013657604080518082019091526012815271125b9a5d1a585b1a5e995c881c185d5cd95960721b60208201526200013690620001a7565b83156200018457845460ff60401b191685556040517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2906200017b90600190620004bb565b60405180910390a15b50505050505050565b620001976200021a565b620001a382826200026b565b5050565b620001de620001d860017fee35723ac350a69d2a92d3703f17439cbaadf2f093a21ba5bf5f1a53eb2a14d9620004e1565b60019055565b7fc32e6d5d6d1de257f64eac19ddb1f700ba13527983849c9486b1ab007ea28381816040516200020f919062000552565b60405180910390a150565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166200026957604051631afcd79f60e31b815260040160405180910390fd5b565b620002756200021a565b620001a38282620002856200021a565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984006001600160a01b038216620002dc576000604051636116401160e11b8152600401620002d3919062000589565b60405180910390fd5b80546001600160d01b0316600160d01b65ffffffffffff851602178155620003066000836200030c565b50505050565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840083620003af576000620003697feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6001600160a01b0316146200039157604051631fe1e13d60e11b815260040160405180910390fd5b6001810180546001600160a01b0319166001600160a01b0385161790555b620003bb8484620003c5565b9150505b92915050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166200048d576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055620004423390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050620003bf565b6000915050620003bf565b60006001600160401b038216620003bf565b620004b58162000498565b82525050565b60208101620003bf8284620004aa565b634e487b7160e01b600052601160045260246000fd5b81810381811115620003bf57620003bf620004cb565b60005b8381101562000514578181015183820152602001620004fa565b50506000910152565b600062000528825190565b80845260208401935062000541818560208601620004f7565b601f01601f19169290920192915050565b602080825281016200056581846200051d565b9392505050565b60006001600160a01b038216620003bf565b620004b5816200056c565b60208101620003bf82846200057e565b61175e80620005a96000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c8063649a5ec711610104578063a217fddf116100a2578063d547741f11610071578063d547741f146103a2578063d602b9fd146103b5578063ee2a6b87146103bd578063f5b541a6146103d057600080fd5b8063a217fddf14610374578063cc8463c81461037c578063cefc142914610384578063cf6eefb71461038c57600080fd5b806384ef8ffc116100de57806384ef8ffc146103435780638da5cb5b1461028d57806391d148541461034b578063a1eda53c1461035e57600080fd5b8063649a5ec7146103155780636da66355146103285780637fbf7b6a1461033b57600080fd5b80633f4ba83a116101715780634d9b47e21161014b5780634d9b47e2146102a257806354fd4d50146102c95780635c975abb146102fa578063634e93da1461030257600080fd5b80633f4ba83a14610272578063400ada751461027a578063452a93201461028d57600080fd5b806324737932116101ad5780632473793214610219578063248a9ca31461022c5780632f2ff15d1461024c57806336568abe1461025f57600080fd5b806301ffc9a7146101d4578063022d63fb146101fd5780630aa6220b1461020f575b600080fd5b6101e76101e23660046111db565b6103f7565b6040516101f49190611206565b60405180910390f35b620697805b6040516101f49190611222565b610217610422565b005b6101e7610227366004611255565b610438565b61023f61023a366004611287565b6104a5565b6040516101f491906112ae565b61021761025a3660046112bc565b6104c7565b61021761026d3660046112bc565b6104f3565b6102176105b8565b61021761028836600461130c565b61062a565b61029561077d565b6040516101f49190611348565b61023f7f8227712ef8ad39d0f26f06731ef0df8665eb7ada7f41b1ee089adf3c238862a281565b6102ed604051806040016040528060058152602001640312e312e360dc1b81525081565b6040516101f491906113ac565b6101e761078c565b610217610310366004611255565b6107ae565b6102176103233660046113d6565b6107c2565b6102176103363660046114f2565b6107d6565b61023f610804565b610295610820565b6101e76103593660046112bc565b61083c565b610366610874565b6040516101f492919061152d565b61023f600081565b6102026108e7565b610217610965565b6103946109a5565b6040516101f4929190611548565b6102176103b03660046112bc565b6109d3565b6102176109fb565b6101e76103cb366004611255565b610a0e565b61023f7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92981565b60006001600160e01b031982166318a4c3c360e11b148061041c575061041c82610a3a565b92915050565b600061042d81610a6f565b610435610a79565b50565b60006104647f8227712ef8ad39d0f26f06731ef0df8665eb7ada7f41b1ee089adf3c238862a28361083c565b8061049457506104947f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b9298361083c565b8061041c575061041c60008361083c565b6000908152600080516020611712833981519152602052604090206001015490565b816104e557604051631fe1e13d60e11b815260040160405180910390fd5b6104ef8282610a86565b5050565b6000805160206116f2833981519152821580156105285750610513610820565b6001600160a01b0316826001600160a01b0316145b156105a9576000806105386109a5565b90925090506001600160a01b03821615158061055a575065ffffffffffff8116155b8061056d57504265ffffffffffff821610155b1561059657806040516319ca5ebb60e01b815260040161058d9190611222565b60405180910390fd5b5050805465ffffffffffff60a01b191681555b6105b38383610aa8565b505050565b6105c133610a0e565b6105dd5760405162461bcd60e51b815260040161058d9061159d565b6105ff6105f960016000805160206116d28339815191526115c3565b60009055565b6040517fa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d1693390600090a1565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156106705750825b905060008267ffffffffffffffff16600114801561068d5750303b155b90508115801561069b575080155b156106b95760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156106e357845460ff60401b1916600160401b1785555b6106f06201518088610adb565b85156107295761072960405180604001604052806012815260200171125b9a5d1a585b1a5e995c881c185d5cd95960721b815250610aed565b831561077457845460ff60401b191685556040517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29061076b906001906115f1565b60405180910390a15b50505050505050565b6000610787610820565b905090565b60006107876107aa60016000805160206116d28339815191526115c3565b5490565b60006107b981610a6f565b6104ef82610b49565b60006107cd81610a6f565b6104ef82610bb7565b6107df33610438565b6107fb5760405162461bcd60e51b815260040161058d90611640565b61043581610aed565b61081d60016000805160206116d28339815191526115c3565b81565b600080516020611732833981519152546001600160a01b031690565b6000918252600080516020611712833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60008051602061173283398151915254600090600160d01b900465ffffffffffff166000805160206116f283398151915281158015906108bc57504265ffffffffffff831610155b6108c8576000806108de565b6001810154600160a01b900465ffffffffffff16825b92509250509091565b600080516020611732833981519152546000906000805160206116f283398151915290600160d01b900465ffffffffffff16801580159061092f57504265ffffffffffff8216105b610949578154600160d01b900465ffffffffffff1661095e565b6001820154600160a01b900465ffffffffffff165b9250505090565b600061096f6109a5565b509050336001600160a01b0382161461099d5733604051636116401160e11b815260040161058d9190611348565b610435610c1e565b6000805160206116f2833981519152546001600160a01b03811691600160a01b90910465ffffffffffff1690565b816109f157604051631fe1e13d60e11b815260040160405180910390fd5b6104ef8282610cb3565b6000610a0681610a6f565b610435610ccf565b60006104947f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b9298361083c565b60006001600160e01b03198216637965db0b60e01b148061041c57506301ffc9a760e01b6001600160e01b031983161461041c565b6104358133610cda565b610a84600080610d05565b565b610a8f826104a5565b610a9881610a6f565b610aa28383610de0565b50505050565b6001600160a01b0381163314610ad15760405163334bd91960e11b815260040160405180910390fd5b6105b38282610e57565b610ae3610eb0565b6104ef8282610ef9565b610b0f610b0960016000805160206116d28339815191526115c3565b60019055565b7fc32e6d5d6d1de257f64eac19ddb1f700ba13527983849c9486b1ab007ea2838181604051610b3e91906113ac565b60405180910390a150565b6000610b536108e7565b610b5c42610f0b565b610b669190611650565b9050610b728282610f3d565b816001600160a01b03167f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed682604051610bab9190611222565b60405180910390a25050565b6000610bc282610fca565b610bcb42610f0b565b610bd59190611650565b9050610be18282610d05565b7ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b8282604051610c1292919061152d565b60405180910390a15050565b6000805160206116f2833981519152600080610c386109a5565b91509150610c4d8165ffffffffffff16151590565b1580610c6157504265ffffffffffff821610155b15610c8157806040516319ca5ebb60e01b815260040161058d9190611222565b610c936000610c8e610820565b610e57565b50610c9f600083610de0565b505081546001600160d01b03191690915550565b610cbc826104a5565b610cc581610a6f565b610aa28383610e57565b610a84600080610f3d565b610ce4828261083c565b6104ef57808260405163e2517d3f60e01b815260040161058d929190611672565b600080516020611732833981519152546000805160206116f283398151915290600160d01b900465ffffffffffff168015610da2574265ffffffffffff82161015610d7857600182015482546001600160d01b0316600160a01b90910465ffffffffffff16600160d01b02178255610da2565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec590600090a15b5060010180546001600160a01b0316600160a01b65ffffffffffff948516026001600160d01b031617600160d01b9290931691909102919091179055565b60006000805160206116f283398151915283610e45576000610e00610820565b6001600160a01b031614610e2757604051631fe1e13d60e11b815260040160405180910390fd5b6001810180546001600160a01b0319166001600160a01b0385161790555b610e4f8484611019565b949350505050565b60006000805160206116f283398151915283158015610e8e5750610e79610820565b6001600160a01b0316836001600160a01b0316145b15610ea6576001810180546001600160a01b03191690555b610e4f84846110be565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610a8457604051631afcd79f60e31b815260040160405180910390fd5b610f01610eb0565b6104ef828261113a565b600065ffffffffffff821115610f39576030826040516306dfcc6560e41b815260040161058d9291906116a1565b5090565b6000805160206116f28339815191526000610f566109a5565b835465ffffffffffff8616600160a01b026001600160d01b03199091166001600160a01b038816171784559150610f9690508165ffffffffffff16151590565b15610aa2576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510990600090a150505050565b600080610fd56108e7565b90508065ffffffffffff168365ffffffffffff1611610ffd57610ff883826116af565b611012565b61101265ffffffffffff8416620697806111a3565b9392505050565b6000600080516020611712833981519152611034848461083c565b6110b4576000848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561106a3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061041c565b600091505061041c565b60006000805160206117128339815191526110d9848461083c565b156110b4576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061041c565b611142610eb0565b6000805160206116f28339815191526001600160a01b03821661117b576000604051636116401160e11b815260040161058d9190611348565b80546001600160d01b0316600160d01b65ffffffffffff851602178155610aa2600083610de0565b60008183106111b25781611012565b5090919050565b6001600160e01b031981165b811461043557600080fd5b803561041c816111b9565b6000602082840312156111f0576111f0600080fd5b6000610e4f84846111d0565b8015155b82525050565b6020810161041c82846111fc565b65ffffffffffff8116611200565b6020810161041c8284611214565b60006001600160a01b03821661041c565b6111c581611230565b803561041c81611241565b60006020828403121561126a5761126a600080fd5b6000610e4f848461124a565b806111c5565b803561041c81611276565b60006020828403121561129c5761129c600080fd5b6000610e4f848461127c565b80611200565b6020810161041c82846112a8565b600080604083850312156112d2576112d2600080fd5b60006112de858561127c565b92505060206112ef8582860161124a565b9150509250929050565b8015156111c5565b803561041c816112f9565b6000806040838503121561132257611322600080fd5b600061132e858561124a565b92505060206112ef85828601611301565b61120081611230565b6020810161041c828461133f565b60005b83811015611371578181015183820152602001611359565b50506000910152565b6000611384825190565b80845260208401935061139b818560208601611356565b601f01601f19169290920192915050565b60208082528101611012818461137a565b65ffffffffffff81166111c5565b803561041c816113bd565b6000602082840312156113eb576113eb600080fd5b6000610e4f84846113cb565b634e487b7160e01b600052604160045260246000fd5b601f19601f830116810181811067ffffffffffffffff82111715611433576114336113f7565b6040525050565b600061144560405190565b9050611451828261140d565b919050565b600067ffffffffffffffff821115611470576114706113f7565b601f19601f83011660200192915050565b82818337506000910152565b60006114a061149b84611456565b61143a565b9050828152602081018484840111156114bb576114bb600080fd5b6114c6848285611481565b509392505050565b600082601f8301126114e2576114e2600080fd5b8135610e4f84826020860161148d565b60006020828403121561150757611507600080fd5b813567ffffffffffffffff81111561152157611521600080fd5b610e4f848285016114ce565b6040810161153b8285611214565b6110126020830184611214565b6040810161153b828561133f565b602781526000602082017f6f6e6c79204f50455241544f525f524f4c45206f722061646d696e2063616e20815266756e706175736560c81b602082015291505b5060400190565b6020808252810161041c81611556565b634e487b7160e01b600052601160045260246000fd5b8181038181111561041c5761041c6115ad565b600067ffffffffffffffff821661041c565b611200816115d6565b6020810161041c82846115e8565b602481526000602082017f6f6e6c79204d4f4e49544f525f524f4c45206f722061646d696e2063616e20708152636175736560e01b60208201529150611596565b6020808252810161041c816115ff565b65ffffffffffff91821691908116908282019081111561041c5761041c6115ad565b60408101611680828561133f565b61101260208301846112a8565b600060ff821661041c565b6112008161168d565b604081016116808285611698565b65ffffffffffff91821691908116908282039081111561041c5761041c6115ad56feee35723ac350a69d2a92d3703f17439cbaadf2f093a21ba5bf5f1a53eb2a14d9eef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800eef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401a164736f6c6343000814000a",
}

// SuperchainConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use SuperchainConfigMetaData.ABI instead.
var SuperchainConfigABI = SuperchainConfigMetaData.ABI

// SuperchainConfigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SuperchainConfigMetaData.Bin instead.
var SuperchainConfigBin = SuperchainConfigMetaData.Bin

// DeploySuperchainConfig deploys a new Ethereum contract, binding an instance of SuperchainConfig to it.
func DeploySuperchainConfig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SuperchainConfig, error) {
	parsed, err := SuperchainConfigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SuperchainConfigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SuperchainConfig{SuperchainConfigCaller: SuperchainConfigCaller{contract: contract}, SuperchainConfigTransactor: SuperchainConfigTransactor{contract: contract}, SuperchainConfigFilterer: SuperchainConfigFilterer{contract: contract}}, nil
}

// SuperchainConfig is an auto generated Go binding around an Ethereum contract.
type SuperchainConfig struct {
	SuperchainConfigCaller     // Read-only binding to the contract
	SuperchainConfigTransactor // Write-only binding to the contract
	SuperchainConfigFilterer   // Log filterer for contract events
}

// SuperchainConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type SuperchainConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperchainConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SuperchainConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperchainConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SuperchainConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperchainConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SuperchainConfigSession struct {
	Contract     *SuperchainConfig // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SuperchainConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SuperchainConfigCallerSession struct {
	Contract *SuperchainConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SuperchainConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SuperchainConfigTransactorSession struct {
	Contract     *SuperchainConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SuperchainConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type SuperchainConfigRaw struct {
	Contract *SuperchainConfig // Generic contract binding to access the raw methods on
}

// SuperchainConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SuperchainConfigCallerRaw struct {
	Contract *SuperchainConfigCaller // Generic read-only contract binding to access the raw methods on
}

// SuperchainConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SuperchainConfigTransactorRaw struct {
	Contract *SuperchainConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuperchainConfig creates a new instance of SuperchainConfig, bound to a specific deployed contract.
func NewSuperchainConfig(address common.Address, backend bind.ContractBackend) (*SuperchainConfig, error) {
	contract, err := bindSuperchainConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SuperchainConfig{SuperchainConfigCaller: SuperchainConfigCaller{contract: contract}, SuperchainConfigTransactor: SuperchainConfigTransactor{contract: contract}, SuperchainConfigFilterer: SuperchainConfigFilterer{contract: contract}}, nil
}

// NewSuperchainConfigCaller creates a new read-only instance of SuperchainConfig, bound to a specific deployed contract.
func NewSuperchainConfigCaller(address common.Address, caller bind.ContractCaller) (*SuperchainConfigCaller, error) {
	contract, err := bindSuperchainConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigCaller{contract: contract}, nil
}

// NewSuperchainConfigTransactor creates a new write-only instance of SuperchainConfig, bound to a specific deployed contract.
func NewSuperchainConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*SuperchainConfigTransactor, error) {
	contract, err := bindSuperchainConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigTransactor{contract: contract}, nil
}

// NewSuperchainConfigFilterer creates a new log filterer instance of SuperchainConfig, bound to a specific deployed contract.
func NewSuperchainConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*SuperchainConfigFilterer, error) {
	contract, err := bindSuperchainConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigFilterer{contract: contract}, nil
}

// bindSuperchainConfig binds a generic wrapper to an already deployed contract.
func bindSuperchainConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SuperchainConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperchainConfig *SuperchainConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SuperchainConfig.Contract.SuperchainConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperchainConfig *SuperchainConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.SuperchainConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperchainConfig *SuperchainConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.SuperchainConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperchainConfig *SuperchainConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SuperchainConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperchainConfig *SuperchainConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperchainConfig *SuperchainConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SuperchainConfig.Contract.DEFAULTADMINROLE(&_SuperchainConfig.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SuperchainConfig.Contract.DEFAULTADMINROLE(&_SuperchainConfig.CallOpts)
}

// MONITORROLE is a free data retrieval call binding the contract method 0x4d9b47e2.
//
// Solidity: function MONITOR_ROLE() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCaller) MONITORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "MONITOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MONITORROLE is a free data retrieval call binding the contract method 0x4d9b47e2.
//
// Solidity: function MONITOR_ROLE() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigSession) MONITORROLE() ([32]byte, error) {
	return _SuperchainConfig.Contract.MONITORROLE(&_SuperchainConfig.CallOpts)
}

// MONITORROLE is a free data retrieval call binding the contract method 0x4d9b47e2.
//
// Solidity: function MONITOR_ROLE() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCallerSession) MONITORROLE() ([32]byte, error) {
	return _SuperchainConfig.Contract.MONITORROLE(&_SuperchainConfig.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigSession) OPERATORROLE() ([32]byte, error) {
	return _SuperchainConfig.Contract.OPERATORROLE(&_SuperchainConfig.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCallerSession) OPERATORROLE() ([32]byte, error) {
	return _SuperchainConfig.Contract.OPERATORROLE(&_SuperchainConfig.CallOpts)
}

// PAUSEDSLOT is a free data retrieval call binding the contract method 0x7fbf7b6a.
//
// Solidity: function PAUSED_SLOT() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCaller) PAUSEDSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "PAUSED_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEDSLOT is a free data retrieval call binding the contract method 0x7fbf7b6a.
//
// Solidity: function PAUSED_SLOT() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigSession) PAUSEDSLOT() ([32]byte, error) {
	return _SuperchainConfig.Contract.PAUSEDSLOT(&_SuperchainConfig.CallOpts)
}

// PAUSEDSLOT is a free data retrieval call binding the contract method 0x7fbf7b6a.
//
// Solidity: function PAUSED_SLOT() view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCallerSession) PAUSEDSLOT() ([32]byte, error) {
	return _SuperchainConfig.Contract.PAUSEDSLOT(&_SuperchainConfig.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_SuperchainConfig *SuperchainConfigCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_SuperchainConfig *SuperchainConfigSession) DefaultAdmin() (common.Address, error) {
	return _SuperchainConfig.Contract.DefaultAdmin(&_SuperchainConfig.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_SuperchainConfig *SuperchainConfigCallerSession) DefaultAdmin() (common.Address, error) {
	return _SuperchainConfig.Contract.DefaultAdmin(&_SuperchainConfig.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_SuperchainConfig *SuperchainConfigCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_SuperchainConfig *SuperchainConfigSession) DefaultAdminDelay() (*big.Int, error) {
	return _SuperchainConfig.Contract.DefaultAdminDelay(&_SuperchainConfig.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_SuperchainConfig *SuperchainConfigCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _SuperchainConfig.Contract.DefaultAdminDelay(&_SuperchainConfig.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_SuperchainConfig *SuperchainConfigCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_SuperchainConfig *SuperchainConfigSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _SuperchainConfig.Contract.DefaultAdminDelayIncreaseWait(&_SuperchainConfig.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_SuperchainConfig *SuperchainConfigCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _SuperchainConfig.Contract.DefaultAdminDelayIncreaseWait(&_SuperchainConfig.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SuperchainConfig.Contract.GetRoleAdmin(&_SuperchainConfig.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SuperchainConfig *SuperchainConfigCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SuperchainConfig.Contract.GetRoleAdmin(&_SuperchainConfig.CallOpts, role)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_SuperchainConfig *SuperchainConfigCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_SuperchainConfig *SuperchainConfigSession) Guardian() (common.Address, error) {
	return _SuperchainConfig.Contract.Guardian(&_SuperchainConfig.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_SuperchainConfig *SuperchainConfigCallerSession) Guardian() (common.Address, error) {
	return _SuperchainConfig.Contract.Guardian(&_SuperchainConfig.CallOpts)
}

// HasMonitorCapabilities is a free data retrieval call binding the contract method 0x24737932.
//
// Solidity: function hasMonitorCapabilities(address _address) view returns(bool)
func (_SuperchainConfig *SuperchainConfigCaller) HasMonitorCapabilities(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "hasMonitorCapabilities", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMonitorCapabilities is a free data retrieval call binding the contract method 0x24737932.
//
// Solidity: function hasMonitorCapabilities(address _address) view returns(bool)
func (_SuperchainConfig *SuperchainConfigSession) HasMonitorCapabilities(_address common.Address) (bool, error) {
	return _SuperchainConfig.Contract.HasMonitorCapabilities(&_SuperchainConfig.CallOpts, _address)
}

// HasMonitorCapabilities is a free data retrieval call binding the contract method 0x24737932.
//
// Solidity: function hasMonitorCapabilities(address _address) view returns(bool)
func (_SuperchainConfig *SuperchainConfigCallerSession) HasMonitorCapabilities(_address common.Address) (bool, error) {
	return _SuperchainConfig.Contract.HasMonitorCapabilities(&_SuperchainConfig.CallOpts, _address)
}

// HasOperatorCapabilities is a free data retrieval call binding the contract method 0xee2a6b87.
//
// Solidity: function hasOperatorCapabilities(address _address) view returns(bool)
func (_SuperchainConfig *SuperchainConfigCaller) HasOperatorCapabilities(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "hasOperatorCapabilities", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOperatorCapabilities is a free data retrieval call binding the contract method 0xee2a6b87.
//
// Solidity: function hasOperatorCapabilities(address _address) view returns(bool)
func (_SuperchainConfig *SuperchainConfigSession) HasOperatorCapabilities(_address common.Address) (bool, error) {
	return _SuperchainConfig.Contract.HasOperatorCapabilities(&_SuperchainConfig.CallOpts, _address)
}

// HasOperatorCapabilities is a free data retrieval call binding the contract method 0xee2a6b87.
//
// Solidity: function hasOperatorCapabilities(address _address) view returns(bool)
func (_SuperchainConfig *SuperchainConfigCallerSession) HasOperatorCapabilities(_address common.Address) (bool, error) {
	return _SuperchainConfig.Contract.HasOperatorCapabilities(&_SuperchainConfig.CallOpts, _address)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SuperchainConfig *SuperchainConfigCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SuperchainConfig *SuperchainConfigSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SuperchainConfig.Contract.HasRole(&_SuperchainConfig.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SuperchainConfig *SuperchainConfigCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SuperchainConfig.Contract.HasRole(&_SuperchainConfig.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SuperchainConfig *SuperchainConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SuperchainConfig *SuperchainConfigSession) Owner() (common.Address, error) {
	return _SuperchainConfig.Contract.Owner(&_SuperchainConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SuperchainConfig *SuperchainConfigCallerSession) Owner() (common.Address, error) {
	return _SuperchainConfig.Contract.Owner(&_SuperchainConfig.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool paused_)
func (_SuperchainConfig *SuperchainConfigCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool paused_)
func (_SuperchainConfig *SuperchainConfigSession) Paused() (bool, error) {
	return _SuperchainConfig.Contract.Paused(&_SuperchainConfig.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool paused_)
func (_SuperchainConfig *SuperchainConfigCallerSession) Paused() (bool, error) {
	return _SuperchainConfig.Contract.Paused(&_SuperchainConfig.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_SuperchainConfig *SuperchainConfigCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(struct {
		NewAdmin common.Address
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_SuperchainConfig *SuperchainConfigSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _SuperchainConfig.Contract.PendingDefaultAdmin(&_SuperchainConfig.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_SuperchainConfig *SuperchainConfigCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _SuperchainConfig.Contract.PendingDefaultAdmin(&_SuperchainConfig.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_SuperchainConfig *SuperchainConfigCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(struct {
		NewDelay *big.Int
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_SuperchainConfig *SuperchainConfigSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _SuperchainConfig.Contract.PendingDefaultAdminDelay(&_SuperchainConfig.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_SuperchainConfig *SuperchainConfigCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _SuperchainConfig.Contract.PendingDefaultAdminDelay(&_SuperchainConfig.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SuperchainConfig *SuperchainConfigCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SuperchainConfig *SuperchainConfigSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SuperchainConfig.Contract.SupportsInterface(&_SuperchainConfig.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SuperchainConfig *SuperchainConfigCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SuperchainConfig.Contract.SupportsInterface(&_SuperchainConfig.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SuperchainConfig *SuperchainConfigCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SuperchainConfig.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SuperchainConfig *SuperchainConfigSession) Version() (string, error) {
	return _SuperchainConfig.Contract.Version(&_SuperchainConfig.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SuperchainConfig *SuperchainConfigCallerSession) Version() (string, error) {
	return _SuperchainConfig.Contract.Version(&_SuperchainConfig.CallOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_SuperchainConfig *SuperchainConfigTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_SuperchainConfig *SuperchainConfigSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _SuperchainConfig.Contract.AcceptDefaultAdminTransfer(&_SuperchainConfig.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _SuperchainConfig.Contract.AcceptDefaultAdminTransfer(&_SuperchainConfig.TransactOpts)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_SuperchainConfig *SuperchainConfigTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_SuperchainConfig *SuperchainConfigSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.BeginDefaultAdminTransfer(&_SuperchainConfig.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.BeginDefaultAdminTransfer(&_SuperchainConfig.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_SuperchainConfig *SuperchainConfigTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_SuperchainConfig *SuperchainConfigSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _SuperchainConfig.Contract.CancelDefaultAdminTransfer(&_SuperchainConfig.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _SuperchainConfig.Contract.CancelDefaultAdminTransfer(&_SuperchainConfig.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_SuperchainConfig *SuperchainConfigTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_SuperchainConfig *SuperchainConfigSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.ChangeDefaultAdminDelay(&_SuperchainConfig.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.ChangeDefaultAdminDelay(&_SuperchainConfig.TransactOpts, newDelay)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SuperchainConfig *SuperchainConfigTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SuperchainConfig *SuperchainConfigSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.GrantRole(&_SuperchainConfig.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.GrantRole(&_SuperchainConfig.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address _admin, bool _paused) returns()
func (_SuperchainConfig *SuperchainConfigTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _paused bool) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "initialize", _admin, _paused)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address _admin, bool _paused) returns()
func (_SuperchainConfig *SuperchainConfigSession) Initialize(_admin common.Address, _paused bool) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.Initialize(&_SuperchainConfig.TransactOpts, _admin, _paused)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address _admin, bool _paused) returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) Initialize(_admin common.Address, _paused bool) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.Initialize(&_SuperchainConfig.TransactOpts, _admin, _paused)
}

// Pause is a paid mutator transaction binding the contract method 0x6da66355.
//
// Solidity: function pause(string _identifier) returns()
func (_SuperchainConfig *SuperchainConfigTransactor) Pause(opts *bind.TransactOpts, _identifier string) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "pause", _identifier)
}

// Pause is a paid mutator transaction binding the contract method 0x6da66355.
//
// Solidity: function pause(string _identifier) returns()
func (_SuperchainConfig *SuperchainConfigSession) Pause(_identifier string) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.Pause(&_SuperchainConfig.TransactOpts, _identifier)
}

// Pause is a paid mutator transaction binding the contract method 0x6da66355.
//
// Solidity: function pause(string _identifier) returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) Pause(_identifier string) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.Pause(&_SuperchainConfig.TransactOpts, _identifier)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SuperchainConfig *SuperchainConfigTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SuperchainConfig *SuperchainConfigSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.RenounceRole(&_SuperchainConfig.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.RenounceRole(&_SuperchainConfig.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SuperchainConfig *SuperchainConfigTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SuperchainConfig *SuperchainConfigSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.RevokeRole(&_SuperchainConfig.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SuperchainConfig.Contract.RevokeRole(&_SuperchainConfig.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_SuperchainConfig *SuperchainConfigTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_SuperchainConfig *SuperchainConfigSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _SuperchainConfig.Contract.RollbackDefaultAdminDelay(&_SuperchainConfig.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _SuperchainConfig.Contract.RollbackDefaultAdminDelay(&_SuperchainConfig.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SuperchainConfig *SuperchainConfigTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperchainConfig.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SuperchainConfig *SuperchainConfigSession) Unpause() (*types.Transaction, error) {
	return _SuperchainConfig.Contract.Unpause(&_SuperchainConfig.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SuperchainConfig *SuperchainConfigTransactorSession) Unpause() (*types.Transaction, error) {
	return _SuperchainConfig.Contract.Unpause(&_SuperchainConfig.TransactOpts)
}

// SuperchainConfigDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the SuperchainConfig contract.
type SuperchainConfigDefaultAdminDelayChangeCanceledIterator struct {
	Event *SuperchainConfigDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigDefaultAdminDelayChangeCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigDefaultAdminDelayChangeCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the SuperchainConfig contract.
type SuperchainConfigDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_SuperchainConfig *SuperchainConfigFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*SuperchainConfigDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigDefaultAdminDelayChangeCanceledIterator{contract: _SuperchainConfig.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_SuperchainConfig *SuperchainConfigFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *SuperchainConfigDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigDefaultAdminDelayChangeCanceled)
				if err := _SuperchainConfig.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultAdminDelayChangeCanceled is a log parse operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_SuperchainConfig *SuperchainConfigFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*SuperchainConfigDefaultAdminDelayChangeCanceled, error) {
	event := new(SuperchainConfigDefaultAdminDelayChangeCanceled)
	if err := _SuperchainConfig.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperchainConfigDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the SuperchainConfig contract.
type SuperchainConfigDefaultAdminDelayChangeScheduledIterator struct {
	Event *SuperchainConfigDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigDefaultAdminDelayChangeScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigDefaultAdminDelayChangeScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the SuperchainConfig contract.
type SuperchainConfigDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_SuperchainConfig *SuperchainConfigFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*SuperchainConfigDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigDefaultAdminDelayChangeScheduledIterator{contract: _SuperchainConfig.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_SuperchainConfig *SuperchainConfigFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *SuperchainConfigDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigDefaultAdminDelayChangeScheduled)
				if err := _SuperchainConfig.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultAdminDelayChangeScheduled is a log parse operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_SuperchainConfig *SuperchainConfigFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*SuperchainConfigDefaultAdminDelayChangeScheduled, error) {
	event := new(SuperchainConfigDefaultAdminDelayChangeScheduled)
	if err := _SuperchainConfig.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperchainConfigDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the SuperchainConfig contract.
type SuperchainConfigDefaultAdminTransferCanceledIterator struct {
	Event *SuperchainConfigDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigDefaultAdminTransferCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigDefaultAdminTransferCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the SuperchainConfig contract.
type SuperchainConfigDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_SuperchainConfig *SuperchainConfigFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*SuperchainConfigDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigDefaultAdminTransferCanceledIterator{contract: _SuperchainConfig.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_SuperchainConfig *SuperchainConfigFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *SuperchainConfigDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigDefaultAdminTransferCanceled)
				if err := _SuperchainConfig.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultAdminTransferCanceled is a log parse operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_SuperchainConfig *SuperchainConfigFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*SuperchainConfigDefaultAdminTransferCanceled, error) {
	event := new(SuperchainConfigDefaultAdminTransferCanceled)
	if err := _SuperchainConfig.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperchainConfigDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the SuperchainConfig contract.
type SuperchainConfigDefaultAdminTransferScheduledIterator struct {
	Event *SuperchainConfigDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigDefaultAdminTransferScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigDefaultAdminTransferScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the SuperchainConfig contract.
type SuperchainConfigDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_SuperchainConfig *SuperchainConfigFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*SuperchainConfigDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigDefaultAdminTransferScheduledIterator{contract: _SuperchainConfig.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_SuperchainConfig *SuperchainConfigFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *SuperchainConfigDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigDefaultAdminTransferScheduled)
				if err := _SuperchainConfig.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultAdminTransferScheduled is a log parse operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_SuperchainConfig *SuperchainConfigFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*SuperchainConfigDefaultAdminTransferScheduled, error) {
	event := new(SuperchainConfigDefaultAdminTransferScheduled)
	if err := _SuperchainConfig.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperchainConfigInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SuperchainConfig contract.
type SuperchainConfigInitializedIterator struct {
	Event *SuperchainConfigInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigInitialized represents a Initialized event raised by the SuperchainConfig contract.
type SuperchainConfigInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SuperchainConfig *SuperchainConfigFilterer) FilterInitialized(opts *bind.FilterOpts) (*SuperchainConfigInitializedIterator, error) {

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigInitializedIterator{contract: _SuperchainConfig.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SuperchainConfig *SuperchainConfigFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SuperchainConfigInitialized) (event.Subscription, error) {

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigInitialized)
				if err := _SuperchainConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SuperchainConfig *SuperchainConfigFilterer) ParseInitialized(log types.Log) (*SuperchainConfigInitialized, error) {
	event := new(SuperchainConfigInitialized)
	if err := _SuperchainConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperchainConfigPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SuperchainConfig contract.
type SuperchainConfigPausedIterator struct {
	Event *SuperchainConfigPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigPaused represents a Paused event raised by the SuperchainConfig contract.
type SuperchainConfigPaused struct {
	Identifier string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xc32e6d5d6d1de257f64eac19ddb1f700ba13527983849c9486b1ab007ea28381.
//
// Solidity: event Paused(string identifier)
func (_SuperchainConfig *SuperchainConfigFilterer) FilterPaused(opts *bind.FilterOpts) (*SuperchainConfigPausedIterator, error) {

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigPausedIterator{contract: _SuperchainConfig.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xc32e6d5d6d1de257f64eac19ddb1f700ba13527983849c9486b1ab007ea28381.
//
// Solidity: event Paused(string identifier)
func (_SuperchainConfig *SuperchainConfigFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SuperchainConfigPaused) (event.Subscription, error) {

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigPaused)
				if err := _SuperchainConfig.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xc32e6d5d6d1de257f64eac19ddb1f700ba13527983849c9486b1ab007ea28381.
//
// Solidity: event Paused(string identifier)
func (_SuperchainConfig *SuperchainConfigFilterer) ParsePaused(log types.Log) (*SuperchainConfigPaused, error) {
	event := new(SuperchainConfigPaused)
	if err := _SuperchainConfig.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperchainConfigRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SuperchainConfig contract.
type SuperchainConfigRoleAdminChangedIterator struct {
	Event *SuperchainConfigRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigRoleAdminChanged represents a RoleAdminChanged event raised by the SuperchainConfig contract.
type SuperchainConfigRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SuperchainConfig *SuperchainConfigFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SuperchainConfigRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigRoleAdminChangedIterator{contract: _SuperchainConfig.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SuperchainConfig *SuperchainConfigFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SuperchainConfigRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigRoleAdminChanged)
				if err := _SuperchainConfig.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SuperchainConfig *SuperchainConfigFilterer) ParseRoleAdminChanged(log types.Log) (*SuperchainConfigRoleAdminChanged, error) {
	event := new(SuperchainConfigRoleAdminChanged)
	if err := _SuperchainConfig.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperchainConfigRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SuperchainConfig contract.
type SuperchainConfigRoleGrantedIterator struct {
	Event *SuperchainConfigRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigRoleGranted represents a RoleGranted event raised by the SuperchainConfig contract.
type SuperchainConfigRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SuperchainConfig *SuperchainConfigFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SuperchainConfigRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigRoleGrantedIterator{contract: _SuperchainConfig.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SuperchainConfig *SuperchainConfigFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SuperchainConfigRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigRoleGranted)
				if err := _SuperchainConfig.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SuperchainConfig *SuperchainConfigFilterer) ParseRoleGranted(log types.Log) (*SuperchainConfigRoleGranted, error) {
	event := new(SuperchainConfigRoleGranted)
	if err := _SuperchainConfig.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperchainConfigRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SuperchainConfig contract.
type SuperchainConfigRoleRevokedIterator struct {
	Event *SuperchainConfigRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigRoleRevoked represents a RoleRevoked event raised by the SuperchainConfig contract.
type SuperchainConfigRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SuperchainConfig *SuperchainConfigFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SuperchainConfigRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigRoleRevokedIterator{contract: _SuperchainConfig.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SuperchainConfig *SuperchainConfigFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SuperchainConfigRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigRoleRevoked)
				if err := _SuperchainConfig.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SuperchainConfig *SuperchainConfigFilterer) ParseRoleRevoked(log types.Log) (*SuperchainConfigRoleRevoked, error) {
	event := new(SuperchainConfigRoleRevoked)
	if err := _SuperchainConfig.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperchainConfigUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SuperchainConfig contract.
type SuperchainConfigUnpausedIterator struct {
	Event *SuperchainConfigUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperchainConfigUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperchainConfigUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperchainConfigUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperchainConfigUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperchainConfigUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperchainConfigUnpaused represents a Unpaused event raised by the SuperchainConfig contract.
type SuperchainConfigUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_SuperchainConfig *SuperchainConfigFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SuperchainConfigUnpausedIterator, error) {

	logs, sub, err := _SuperchainConfig.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SuperchainConfigUnpausedIterator{contract: _SuperchainConfig.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_SuperchainConfig *SuperchainConfigFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SuperchainConfigUnpaused) (event.Subscription, error) {

	logs, sub, err := _SuperchainConfig.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperchainConfigUnpaused)
				if err := _SuperchainConfig.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_SuperchainConfig *SuperchainConfigFilterer) ParseUnpaused(log types.Log) (*SuperchainConfigUnpaused, error) {
	event := new(SuperchainConfigUnpaused)
	if err := _SuperchainConfig.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
