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

// L2ERC721BridgeMetaData contains all meta data concerning the L2ERC721Bridge contract.
var L2ERC721BridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OTHER_BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractStandardBridge\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeERC721\",\"inputs\":[{\"name\":\"_localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgeERC721To\",\"inputs\":[{\"name\":\"_localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeBridgeERC721\",\"inputs\":[{\"name\":\"_localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"otherBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractStandardBridge\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ERC721BridgeFinalized\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC721BridgeInitiated\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001e600062000024565b6200021c565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff1615906001600160401b03166000811580156200006f5750825b90506000826001600160401b031660011480156200008c5750303b155b9050811580156200009b575080155b15620000ba5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315620000e957845460ff60401b1916680100000000000000001785555b62000109734200000000000000000000000000000000000007876200015f565b83156200015757845460ff60401b191685556040517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2906200014e906001906200020c565b60405180910390a15b505050505050565b6200016962000197565b600080546001600160a01b039384166001600160a01b03199182161790915560018054929093169116179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16620001e657604051631afcd79f60e31b815260040160405180910390fd5b565b60006001600160401b0382165b92915050565b6200020681620001e8565b82525050565b60208101620001f58284620001fb565b611249806200022c6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637f46ddb2116100665780637f46ddb214610134578063927ede2d14610145578063aa55745214610156578063c4d66de814610169578063c89701a21461017c57600080fd5b80633687011a146100a35780633cb747bf146100b857806354fd4d50146100e15780635c975abb14610112578063761f449314610121575b600080fd5b6100b66100b1366004610a77565b61018f565b005b6000546100cb906001600160a01b031681565b6040516100d89190610b32565b60405180910390f35b610105604051806040016040528060058152602001640312e372e360dc1b81525081565b6040516100d89190610b98565b60006040516100d89190610bb1565b6100b661012f366004610bbf565b6101cf565b6001546001600160a01b03166100cb565b6000546001600160a01b03166100cb565b6100b6610164366004610c6a565b61042b565b6100b6610177366004610cd8565b61046a565b6001546100cb906001600160a01b031681565b333b156101b75760405162461bcd60e51b81526004016101ae90610d4e565b60405180910390fd5b6101c78686333388888888610586565b505050505050565b6000546001600160a01b031633148015610264575060015460005460408051636e296e4560e01b815290516001600160a01b039384169390921691636e296e45916004808201926020929091908290030181865afa158015610235573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102599190610d69565b6001600160a01b0316145b6102805760405162461bcd60e51b81526004016101ae90610de4565b306001600160a01b038816036102a85760405162461bcd60e51b81526004016101ae90610e3b565b6102b9876374259ebf60e01b61086e565b6102d55760405162461bcd60e51b81526004016101ae90610e9e565b866001600160a01b031663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610313573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103379190610d69565b6001600160a01b0316866001600160a01b0316146103675760405162461bcd60e51b81526004016101ae90610eae565b604051632851206560e21b81526001600160a01b0388169063a1448194906103959087908790600401610f32565b600060405180830381600087803b1580156103af57600080fd5b505af11580156103c3573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b0316886001600160a01b03167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac8787878760405161041a9493929190610f83565b60405180910390a450505050505050565b6001600160a01b0385166104515760405162461bcd60e51b81526004016101ae90611008565b6104618787338888888888610586565b50505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156104b05750825b905060008267ffffffffffffffff1660011480156104cd5750303b155b9050811580156104db575080155b156104f95760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561052357845460ff60401b1916600160401b1785555b6105346007602160991b0187610893565b83156101c757845460ff60401b191685556040517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29061057690600190611033565b60405180910390a1505050505050565b6001600160a01b0387166105ac5760405162461bcd60e51b81526004016101ae9061108f565b6040516331a9108f60e11b81526001600160a01b03891690636352211e906105d890879060040161109f565b602060405180830381865afa1580156105f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106199190610d69565b6001600160a01b0316866001600160a01b0316146106495760405162461bcd60e51b81526004016101ae90611107565b6000886001600160a01b031663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610689573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ad9190610d69565b9050876001600160a01b0316816001600160a01b0316146106e05760405162461bcd60e51b81526004016101ae90611171565b604051632770a7eb60e21b81526001600160a01b038a1690639dc29fac9061070e908a908990600401610f32565b600060405180830381600087803b15801561072857600080fd5b505af115801561073c573d6000803e3d6000fd5b50505050600063761f449360e01b828b8a8a8a89896040516024016107679796959493929190611181565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092526000546001549251633dbb202b60e01b81529193506001600160a01b0390811692633dbb202b926107d592919091169085908a906004016111ef565b600060405180830381600087803b1580156107ef57600080fd5b505af1158015610803573d6000803e3d6000fd5b50505050876001600160a01b0316826001600160a01b03168b6001600160a01b03167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a58a8a898960405161085a9493929190610f83565b60405180910390a450505050505050505050565b6000610879836108c9565b801561088a575061088a83836108fc565b90505b92915050565b61089b61097f565b600080546001600160a01b039384166001600160a01b03199182161790915560018054929093169116179055565b60006108dc826301ffc9a760e01b6108fc565b801561088d57506108f5826001600160e01b03196108fc565b1592915050565b60008082604051602401610910919061122e565b60408051601f19818403018152919052602080820180516001600160e01b03166301ffc9a760e01b178152825192935060009283928392909183918a617530fa92503d91506000519050828015610968575060208210155b80156109745750600081115b979650505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166109c857604051631afcd79f60e31b815260040160405180910390fd5b565b60006001600160a01b03821661088d565b6109e4816109ca565b81146109ef57600080fd5b50565b803561088d816109db565b806109e4565b803561088d816109fd565b63ffffffff81166109e4565b803561088d81610a0e565b60008083601f840112610a3a57610a3a600080fd5b50813567ffffffffffffffff811115610a5557610a55600080fd5b602083019150836001820283011115610a7057610a70600080fd5b9250929050565b60008060008060008060a08789031215610a9357610a93600080fd5b6000610a9f89896109f2565b9650506020610ab089828a016109f2565b9550506040610ac189828a01610a03565b9450506060610ad289828a01610a1a565b935050608087013567ffffffffffffffff811115610af257610af2600080fd5b610afe89828a01610a25565b92509250509295509295509295565b600061088d826109ca565b600061088d82610b0d565b610b2c81610b18565b82525050565b6020810161088d8284610b23565b60005b83811015610b5b578181015183820152602001610b43565b50506000910152565b6000610b6e825190565b808452602084019350610b85818560208601610b40565b601f19601f8201165b9093019392505050565b6020808252810161088a8184610b64565b801515610b2c565b6020810161088d8284610ba9565b600080600080600080600060c0888a031215610bdd57610bdd600080fd5b6000610be98a8a6109f2565b9750506020610bfa8a828b016109f2565b9650506040610c0b8a828b016109f2565b9550506060610c1c8a828b016109f2565b9450506080610c2d8a828b01610a03565b93505060a088013567ffffffffffffffff811115610c4d57610c4d600080fd5b610c598a828b01610a25565b925092505092959891949750929550565b600080600080600080600060c0888a031215610c8857610c88600080fd5b6000610c948a8a6109f2565b9750506020610ca58a828b016109f2565b9650506040610cb68a828b016109f2565b9550506060610cc78a828b01610a03565b9450506080610c2d8a828b01610a1a565b600060208284031215610ced57610ced600080fd5b6000610cf984846109f2565b949350505050565b602d81526000602082017f4552433732314272696467653a206163636f756e74206973206e6f742065787481526c195c9b985b1b1e481bdddb9959609a1b602082015291505b5060400190565b6020808252810161088d81610d01565b805161088d816109db565b600060208284031215610d7e57610d7e600080fd5b6000610cf98484610d5e565b603f81526000602082017f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792081527f62652063616c6c65642066726f6d20746865206f74686572206272696467650060208201529150610d47565b6020808252810161088d81610d8a565b602a81526000602082017f4c324552433732314272696467653a206c6f63616c20746f6b656e2063616e6e81526937ba1031329039b2b63360b11b60208201529150610d47565b6020808252810161088d81610df4565b603681526000602082017f4c324552433732314272696467653a206c6f63616c20746f6b656e20696e74658152751c999858d9481a5cc81b9bdd0818dbdb5c1b1a585b9d60521b60208201529150610d47565b6020808252810161088d81610e4b565b6020808252810161088d81604b81527f4c324552433732314272696467653a2077726f6e672072656d6f746520746f6b60208201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433732312060408201526a3637b1b0b6103a37b5b2b760a91b606082015260800190565b610b2c816109ca565b80610b2c565b60408101610f408285610f23565b610f4d6020830184610f2c565b9392505050565b82818337506000910152565b8183526000602084019350610f76838584610f54565b601f19601f840116610b8e565b60608101610f918287610f23565b610f9e6020830186610f2c565b8181036040830152610fb1818486610f60565b9695505050505050565b603081526000602082017f4552433732314272696467653a206e667420726563697069656e742063616e6e81526f6f74206265206164647265737328302960801b60208201529150610d47565b6020808252810161088d81610fbb565b600067ffffffffffffffff821661088d565b610b2c81611018565b6020810161088d828461102a565b603181526000602082017f4c324552433732314272696467653a2072656d6f746520746f6b656e2063616e8152706e6f74206265206164647265737328302960781b60208201529150610d47565b6020808252810161088d81611041565b6020810161088d8284610f2c565b603e81526000602082017f4c324552433732314272696467653a205769746864726177616c206973206e6f81527f74206265696e6720696e69746961746564206279204e4654206f776e6572000060208201529150610d47565b6020808252810161088d816110ad565b603781526000602082017f4c324552433732314272696467653a2072656d6f746520746f6b656e20646f6581527f73206e6f74206d6174636820676976656e2076616c756500000000000000000060208201529150610d47565b6020808252810161088d81611117565b60c0810161118f828a610f23565b61119c6020830189610f23565b6111a96040830188610f23565b6111b66060830187610f23565b6111c36080830186610f2c565b81810360a08301526111d6818486610f60565b9998505050505050505050565b63ffffffff8116610b2c565b606081016111fd8286610f23565b818103602083015261120f8185610b64565b9050610cf960408301846111e3565b6001600160e01b03198116610b2c565b6020810161088d828461121e56fea164736f6c6343000814000a",
}

// L2ERC721BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ERC721BridgeMetaData.ABI instead.
var L2ERC721BridgeABI = L2ERC721BridgeMetaData.ABI

// L2ERC721BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ERC721BridgeMetaData.Bin instead.
var L2ERC721BridgeBin = L2ERC721BridgeMetaData.Bin

// DeployL2ERC721Bridge deploys a new Ethereum contract, binding an instance of L2ERC721Bridge to it.
func DeployL2ERC721Bridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2ERC721Bridge, error) {
	parsed, err := L2ERC721BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ERC721BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ERC721Bridge{L2ERC721BridgeCaller: L2ERC721BridgeCaller{contract: contract}, L2ERC721BridgeTransactor: L2ERC721BridgeTransactor{contract: contract}, L2ERC721BridgeFilterer: L2ERC721BridgeFilterer{contract: contract}}, nil
}

// L2ERC721Bridge is an auto generated Go binding around an Ethereum contract.
type L2ERC721Bridge struct {
	L2ERC721BridgeCaller     // Read-only binding to the contract
	L2ERC721BridgeTransactor // Write-only binding to the contract
	L2ERC721BridgeFilterer   // Log filterer for contract events
}

// L2ERC721BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ERC721BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC721BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ERC721BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC721BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ERC721BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC721BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ERC721BridgeSession struct {
	Contract     *L2ERC721Bridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ERC721BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ERC721BridgeCallerSession struct {
	Contract *L2ERC721BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// L2ERC721BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ERC721BridgeTransactorSession struct {
	Contract     *L2ERC721BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// L2ERC721BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ERC721BridgeRaw struct {
	Contract *L2ERC721Bridge // Generic contract binding to access the raw methods on
}

// L2ERC721BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ERC721BridgeCallerRaw struct {
	Contract *L2ERC721BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L2ERC721BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ERC721BridgeTransactorRaw struct {
	Contract *L2ERC721BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ERC721Bridge creates a new instance of L2ERC721Bridge, bound to a specific deployed contract.
func NewL2ERC721Bridge(address common.Address, backend bind.ContractBackend) (*L2ERC721Bridge, error) {
	contract, err := bindL2ERC721Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ERC721Bridge{L2ERC721BridgeCaller: L2ERC721BridgeCaller{contract: contract}, L2ERC721BridgeTransactor: L2ERC721BridgeTransactor{contract: contract}, L2ERC721BridgeFilterer: L2ERC721BridgeFilterer{contract: contract}}, nil
}

// NewL2ERC721BridgeCaller creates a new read-only instance of L2ERC721Bridge, bound to a specific deployed contract.
func NewL2ERC721BridgeCaller(address common.Address, caller bind.ContractCaller) (*L2ERC721BridgeCaller, error) {
	contract, err := bindL2ERC721Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeCaller{contract: contract}, nil
}

// NewL2ERC721BridgeTransactor creates a new write-only instance of L2ERC721Bridge, bound to a specific deployed contract.
func NewL2ERC721BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ERC721BridgeTransactor, error) {
	contract, err := bindL2ERC721Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeTransactor{contract: contract}, nil
}

// NewL2ERC721BridgeFilterer creates a new log filterer instance of L2ERC721Bridge, bound to a specific deployed contract.
func NewL2ERC721BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ERC721BridgeFilterer, error) {
	contract, err := bindL2ERC721Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeFilterer{contract: contract}, nil
}

// bindL2ERC721Bridge binds a generic wrapper to an already deployed contract.
func bindL2ERC721Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2ERC721BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC721Bridge *L2ERC721BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC721Bridge.Contract.L2ERC721BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC721Bridge *L2ERC721BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.L2ERC721BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC721Bridge *L2ERC721BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.L2ERC721BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC721Bridge *L2ERC721BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC721Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC721Bridge *L2ERC721BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC721Bridge *L2ERC721BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeSession) MESSENGER() (common.Address, error) {
	return _L2ERC721Bridge.Contract.MESSENGER(&_L2ERC721Bridge.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) MESSENGER() (common.Address, error) {
	return _L2ERC721Bridge.Contract.MESSENGER(&_L2ERC721Bridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) OTHERBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "OTHER_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeSession) OTHERBRIDGE() (common.Address, error) {
	return _L2ERC721Bridge.Contract.OTHERBRIDGE(&_L2ERC721Bridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) OTHERBRIDGE() (common.Address, error) {
	return _L2ERC721Bridge.Contract.OTHERBRIDGE(&_L2ERC721Bridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeSession) Messenger() (common.Address, error) {
	return _L2ERC721Bridge.Contract.Messenger(&_L2ERC721Bridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) Messenger() (common.Address, error) {
	return _L2ERC721Bridge.Contract.Messenger(&_L2ERC721Bridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) OtherBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "otherBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeSession) OtherBridge() (common.Address, error) {
	return _L2ERC721Bridge.Contract.OtherBridge(&_L2ERC721Bridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) OtherBridge() (common.Address, error) {
	return _L2ERC721Bridge.Contract.OtherBridge(&_L2ERC721Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2ERC721Bridge *L2ERC721BridgeSession) Paused() (bool, error) {
	return _L2ERC721Bridge.Contract.Paused(&_L2ERC721Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) Paused() (bool, error) {
	return _L2ERC721Bridge.Contract.Paused(&_L2ERC721Bridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ERC721Bridge *L2ERC721BridgeSession) Version() (string, error) {
	return _L2ERC721Bridge.Contract.Version(&_L2ERC721Bridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) Version() (string, error) {
	return _L2ERC721Bridge.Contract.Version(&_L2ERC721Bridge.CallOpts)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x3687011a.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactor) BridgeERC721(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2ERC721Bridge.contract.Transact(opts, "bridgeERC721", _localToken, _remoteToken, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x3687011a.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeSession) BridgeERC721(_localToken common.Address, _remoteToken common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.BridgeERC721(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x3687011a.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactorSession) BridgeERC721(_localToken common.Address, _remoteToken common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.BridgeERC721(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721To is a paid mutator transaction binding the contract method 0xaa557452.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactor) BridgeERC721To(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _to common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2ERC721Bridge.contract.Transact(opts, "bridgeERC721To", _localToken, _remoteToken, _to, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721To is a paid mutator transaction binding the contract method 0xaa557452.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeSession) BridgeERC721To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.BridgeERC721To(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _to, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721To is a paid mutator transaction binding the contract method 0xaa557452.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactorSession) BridgeERC721To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.BridgeERC721To(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _to, _tokenId, _minGasLimit, _extraData)
}

// FinalizeBridgeERC721 is a paid mutator transaction binding the contract method 0x761f4493.
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactor) FinalizeBridgeERC721(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2ERC721Bridge.contract.Transact(opts, "finalizeBridgeERC721", _localToken, _remoteToken, _from, _to, _tokenId, _extraData)
}

// FinalizeBridgeERC721 is a paid mutator transaction binding the contract method 0x761f4493.
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeSession) FinalizeBridgeERC721(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.FinalizeBridgeERC721(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _from, _to, _tokenId, _extraData)
}

// FinalizeBridgeERC721 is a paid mutator transaction binding the contract method 0x761f4493.
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactorSession) FinalizeBridgeERC721(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.FinalizeBridgeERC721(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _from, _to, _tokenId, _extraData)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _l1ERC721Bridge) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactor) Initialize(opts *bind.TransactOpts, _l1ERC721Bridge common.Address) (*types.Transaction, error) {
	return _L2ERC721Bridge.contract.Transact(opts, "initialize", _l1ERC721Bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _l1ERC721Bridge) returns()
func (_L2ERC721Bridge *L2ERC721BridgeSession) Initialize(_l1ERC721Bridge common.Address) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.Initialize(&_L2ERC721Bridge.TransactOpts, _l1ERC721Bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _l1ERC721Bridge) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactorSession) Initialize(_l1ERC721Bridge common.Address) (*types.Transaction, error) {
	return _L2ERC721Bridge.Contract.Initialize(&_L2ERC721Bridge.TransactOpts, _l1ERC721Bridge)
}

// L2ERC721BridgeERC721BridgeFinalizedIterator is returned from FilterERC721BridgeFinalized and is used to iterate over the raw logs and unpacked data for ERC721BridgeFinalized events raised by the L2ERC721Bridge contract.
type L2ERC721BridgeERC721BridgeFinalizedIterator struct {
	Event *L2ERC721BridgeERC721BridgeFinalized // Event containing the contract specifics and raw log

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
func (it *L2ERC721BridgeERC721BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721BridgeERC721BridgeFinalized)
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
		it.Event = new(L2ERC721BridgeERC721BridgeFinalized)
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
func (it *L2ERC721BridgeERC721BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721BridgeERC721BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721BridgeERC721BridgeFinalized represents a ERC721BridgeFinalized event raised by the L2ERC721Bridge contract.
type L2ERC721BridgeERC721BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	TokenId     *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC721BridgeFinalized is a free log retrieval operation binding the contract event 0x1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) FilterERC721BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L2ERC721BridgeERC721BridgeFinalizedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2ERC721Bridge.contract.FilterLogs(opts, "ERC721BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeERC721BridgeFinalizedIterator{contract: _L2ERC721Bridge.contract, event: "ERC721BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchERC721BridgeFinalized is a free log subscription operation binding the contract event 0x1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) WatchERC721BridgeFinalized(opts *bind.WatchOpts, sink chan<- *L2ERC721BridgeERC721BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2ERC721Bridge.contract.WatchLogs(opts, "ERC721BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721BridgeERC721BridgeFinalized)
				if err := _L2ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeFinalized", log); err != nil {
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

// ParseERC721BridgeFinalized is a log parse operation binding the contract event 0x1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) ParseERC721BridgeFinalized(log types.Log) (*L2ERC721BridgeERC721BridgeFinalized, error) {
	event := new(L2ERC721BridgeERC721BridgeFinalized)
	if err := _L2ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721BridgeERC721BridgeInitiatedIterator is returned from FilterERC721BridgeInitiated and is used to iterate over the raw logs and unpacked data for ERC721BridgeInitiated events raised by the L2ERC721Bridge contract.
type L2ERC721BridgeERC721BridgeInitiatedIterator struct {
	Event *L2ERC721BridgeERC721BridgeInitiated // Event containing the contract specifics and raw log

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
func (it *L2ERC721BridgeERC721BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721BridgeERC721BridgeInitiated)
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
		it.Event = new(L2ERC721BridgeERC721BridgeInitiated)
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
func (it *L2ERC721BridgeERC721BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721BridgeERC721BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721BridgeERC721BridgeInitiated represents a ERC721BridgeInitiated event raised by the L2ERC721Bridge contract.
type L2ERC721BridgeERC721BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	TokenId     *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC721BridgeInitiated is a free log retrieval operation binding the contract event 0xb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a5.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) FilterERC721BridgeInitiated(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L2ERC721BridgeERC721BridgeInitiatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2ERC721Bridge.contract.FilterLogs(opts, "ERC721BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeERC721BridgeInitiatedIterator{contract: _L2ERC721Bridge.contract, event: "ERC721BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchERC721BridgeInitiated is a free log subscription operation binding the contract event 0xb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a5.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) WatchERC721BridgeInitiated(opts *bind.WatchOpts, sink chan<- *L2ERC721BridgeERC721BridgeInitiated, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2ERC721Bridge.contract.WatchLogs(opts, "ERC721BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721BridgeERC721BridgeInitiated)
				if err := _L2ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeInitiated", log); err != nil {
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

// ParseERC721BridgeInitiated is a log parse operation binding the contract event 0xb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a5.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) ParseERC721BridgeInitiated(log types.Log) (*L2ERC721BridgeERC721BridgeInitiated, error) {
	event := new(L2ERC721BridgeERC721BridgeInitiated)
	if err := _L2ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2ERC721Bridge contract.
type L2ERC721BridgeInitializedIterator struct {
	Event *L2ERC721BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *L2ERC721BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721BridgeInitialized)
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
		it.Event = new(L2ERC721BridgeInitialized)
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
func (it *L2ERC721BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721BridgeInitialized represents a Initialized event raised by the L2ERC721Bridge contract.
type L2ERC721BridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2ERC721BridgeInitializedIterator, error) {

	logs, sub, err := _L2ERC721Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeInitializedIterator{contract: _L2ERC721Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2ERC721BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _L2ERC721Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721BridgeInitialized)
				if err := _L2ERC721Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) ParseInitialized(log types.Log) (*L2ERC721BridgeInitialized, error) {
	event := new(L2ERC721BridgeInitialized)
	if err := _L2ERC721Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
