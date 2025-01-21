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

// SafeProxyFactoryMetaData contains all meta data concerning the SafeProxyFactory contract.
var SafeProxyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createChainSpecificProxyWithNonce\",\"inputs\":[{\"name\":\"_singleton\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initializer\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"saltNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"contractSafeProxy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createProxyWithCallback\",\"inputs\":[{\"name\":\"_singleton\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initializer\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"saltNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callback\",\"type\":\"address\",\"internalType\":\"contractIProxyCreationCallback\"}],\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"contractSafeProxy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createProxyWithNonce\",\"inputs\":[{\"name\":\"_singleton\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initializer\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"saltNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"contractSafeProxy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyCreationCode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"ProxyCreation\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractSafeProxy\"},{\"name\":\"singleton\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506109af806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631688f0b91461005c5780633408e4701461008557806353e5d93514610093578063d18af54d146100a8578063ec9e80bb146100bb575b600080fd5b61006f61006a366004610485565b6100ce565b60405161007c9190610526565b60405180910390f35b4660405161007c919061053a565b61009b610159565b60405161007c919061059e565b61006f6100b63660046105d5565b610183565b61006f6100c9366004610485565b61023e565b6000808380519060200120836040516020016100eb929190610654565b60405160208183030381529060405280519060200120905061010e858583610263565b9150816001600160a01b03167f4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235866040516101499190610683565b60405180910390a2509392505050565b60606040518060200161016b90610337565b601f1982820381018352601f90910116604052919050565b60008083836040516020016101999291906106b9565b6040516020818303038152906040528051906020012060001c90506101bf8686836100ce565b91506001600160a01b03831615610235576040516303ca56a360e31b81526001600160a01b03841690631e52b518906102029085908a908a908a906004016106df565b600060405180830381600087803b15801561021c57600080fd5b505af1158015610230573d6000803e3d6000fd5b505050505b50949350505050565b6000808380519060200120836102514690565b6040516020016100eb93929190610724565b6000833b61028c5760405162461bcd60e51b815260040161028390610792565b60405180910390fd5b60006040518060200161029e90610337565b601f1982820381018352601f9091011660408190526102cb91906001600160a01b038816906020016107c4565b6040516020818303038152906040529050828151826020016000f591506001600160a01b03821661030e5760405162461bcd60e51b815260040161028390610806565b83511561032f5760008060008651602088016000875af10361032f57600080fd5b509392505050565b61018c8061081783390190565b60006001600160a01b0382165b92915050565b61036081610344565b811461036b57600080fd5b50565b803561035181610357565b634e487b7160e01b600052604160045260246000fd5b601f19601f830116810181811067ffffffffffffffff821117156103b5576103b5610379565b6040525050565b60006103c760405190565b90506103d3828261038f565b919050565b600067ffffffffffffffff8211156103f2576103f2610379565b601f19601f83011660200192915050565b82818337506000910152565b600061042261041d846103d8565b6103bc565b90508281526020810184848401111561043d5761043d600080fd5b61032f848285610403565b600082601f83011261045c5761045c600080fd5b813561046c84826020860161040f565b949350505050565b80610360565b803561035181610474565b60008060006060848603121561049d5761049d600080fd5b60006104a9868661036e565b935050602084013567ffffffffffffffff8111156104c9576104c9600080fd5b6104d586828701610448565b92505060406104e68682870161047a565b9150509250925092565b60006001600160a01b038216610351565b6000610351826104f0565b600061035182610501565b6105208161050c565b82525050565b602081016103518284610517565b80610520565b602081016103518284610534565b60005b8381101561056357818101518382015260200161054b565b50506000910152565b6000610576825190565b80845260208401935061058d818560208601610548565b601f01601f19169290920192915050565b602080825281016105af818461056c565b9392505050565b600061035182610344565b610360816105b6565b8035610351816105c1565b600080600080608085870312156105ee576105ee600080fd5b60006105fa878761036e565b945050602085013567ffffffffffffffff81111561061a5761061a600080fd5b61062687828801610448565b93505060406106378782880161047a565b9250506060610648878288016105ca565b91505092959194509250565b60006106608285610534565b6020820191506106708284610534565b5060200192915050565b61052081610344565b60208101610351828461067a565b60006103518260601b90565b600061035182610691565b6105206106b48261050c565b61069d565b60006106c58285610534565b6020820191506106d582846106a8565b5060140192915050565b608081016106ed8287610517565b6106fa602083018661067a565b818103604083015261070c818561056c565b905061071b6060830184610534565b95945050505050565b60006107308286610534565b6020820191506107408285610534565b6020820191506107508284610534565b506020019392505050565b601f81526000602082017f53696e676c65746f6e20636f6e7472616374206e6f74206465706c6f79656400815291505b5060200190565b602080825281016103518161075b565b60006107ac825190565b6107ba818560208601610548565b9290920192915050565b60006107d082856107a2565b91506106708284610534565b601381526000602082017210dc99585d194c8818d85b1b0819985a5b1959606a1b8152915061078b565b60208082528101610351816107dc56fe608060405234801561001057600080fd5b5060405161018c38038061018c83398101604081905261002f916100b8565b6001600160a01b03811661005e5760405162461bcd60e51b8152600401610055906100e1565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055610127565b60006001600160a01b0382165b92915050565b61009f81610083565b81146100aa57600080fd5b50565b805161009081610096565b6000602082840312156100cd576100cd600080fd5b60006100d984846100ad565b949350505050565b6020808252810161009081602281527f496e76616c69642073696e676c65746f6e20616464726573732070726f766964602082015261195960f21b604082015260600190565b6057806101356000396000f3fe60806040526001600160a01b036000541663530ca43760e11b600035036029578060005260206000f35b3660008037600080366000845af43d6000803e806045573d6000fd5b3d6000f3fea164736f6c6343000814000aa164736f6c6343000814000a",
}

// SafeProxyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeProxyFactoryMetaData.ABI instead.
var SafeProxyFactoryABI = SafeProxyFactoryMetaData.ABI

// SafeProxyFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeProxyFactoryMetaData.Bin instead.
var SafeProxyFactoryBin = SafeProxyFactoryMetaData.Bin

// DeploySafeProxyFactory deploys a new Ethereum contract, binding an instance of SafeProxyFactory to it.
func DeploySafeProxyFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeProxyFactory, error) {
	parsed, err := SafeProxyFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeProxyFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeProxyFactory{SafeProxyFactoryCaller: SafeProxyFactoryCaller{contract: contract}, SafeProxyFactoryTransactor: SafeProxyFactoryTransactor{contract: contract}, SafeProxyFactoryFilterer: SafeProxyFactoryFilterer{contract: contract}}, nil
}

// SafeProxyFactory is an auto generated Go binding around an Ethereum contract.
type SafeProxyFactory struct {
	SafeProxyFactoryCaller     // Read-only binding to the contract
	SafeProxyFactoryTransactor // Write-only binding to the contract
	SafeProxyFactoryFilterer   // Log filterer for contract events
}

// SafeProxyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeProxyFactorySession struct {
	Contract     *SafeProxyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeProxyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeProxyFactoryCallerSession struct {
	Contract *SafeProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SafeProxyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeProxyFactoryTransactorSession struct {
	Contract     *SafeProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SafeProxyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeProxyFactoryRaw struct {
	Contract *SafeProxyFactory // Generic contract binding to access the raw methods on
}

// SafeProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeProxyFactoryCallerRaw struct {
	Contract *SafeProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SafeProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeProxyFactoryTransactorRaw struct {
	Contract *SafeProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeProxyFactory creates a new instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactory(address common.Address, backend bind.ContractBackend) (*SafeProxyFactory, error) {
	contract, err := bindSafeProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactory{SafeProxyFactoryCaller: SafeProxyFactoryCaller{contract: contract}, SafeProxyFactoryTransactor: SafeProxyFactoryTransactor{contract: contract}, SafeProxyFactoryFilterer: SafeProxyFactoryFilterer{contract: contract}}, nil
}

// NewSafeProxyFactoryCaller creates a new read-only instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*SafeProxyFactoryCaller, error) {
	contract, err := bindSafeProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryCaller{contract: contract}, nil
}

// NewSafeProxyFactoryTransactor creates a new write-only instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeProxyFactoryTransactor, error) {
	contract, err := bindSafeProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryTransactor{contract: contract}, nil
}

// NewSafeProxyFactoryFilterer creates a new log filterer instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeProxyFactoryFilterer, error) {
	contract, err := bindSafeProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryFilterer{contract: contract}, nil
}

// bindSafeProxyFactory binds a generic wrapper to an already deployed contract.
func bindSafeProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeProxyFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactory *SafeProxyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactory.Contract.SafeProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactory *SafeProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.SafeProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactory *SafeProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.SafeProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactory *SafeProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactory *SafeProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactory *SafeProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactory *SafeProxyFactoryCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactory *SafeProxyFactorySession) GetChainId() (*big.Int, error) {
	return _SafeProxyFactory.Contract.GetChainId(&_SafeProxyFactory.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) GetChainId() (*big.Int, error) {
	return _SafeProxyFactory.Contract.GetChainId(&_SafeProxyFactory.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactory *SafeProxyFactoryCaller) ProxyCreationCode(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "proxyCreationCode")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactory *SafeProxyFactorySession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactory.Contract.ProxyCreationCode(&_SafeProxyFactory.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactory.Contract.ProxyCreationCode(&_SafeProxyFactory.CallOpts)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactor) CreateChainSpecificProxyWithNonce(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.contract.Transact(opts, "createChainSpecificProxyWithNonce", _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactorySession) CreateChainSpecificProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateChainSpecificProxyWithNonce(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactorSession) CreateChainSpecificProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateChainSpecificProxyWithNonce(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactor) CreateProxyWithCallback(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _SafeProxyFactory.contract.Transact(opts, "createProxyWithCallback", _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactorySession) CreateProxyWithCallback(_singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxyWithCallback(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactorSession) CreateProxyWithCallback(_singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxyWithCallback(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactor) CreateProxyWithNonce(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.contract.Transact(opts, "createProxyWithNonce", _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactorySession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxyWithNonce(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactorSession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxyWithNonce(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// SafeProxyFactoryProxyCreationIterator is returned from FilterProxyCreation and is used to iterate over the raw logs and unpacked data for ProxyCreation events raised by the SafeProxyFactory contract.
type SafeProxyFactoryProxyCreationIterator struct {
	Event *SafeProxyFactoryProxyCreation // Event containing the contract specifics and raw log

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
func (it *SafeProxyFactoryProxyCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeProxyFactoryProxyCreation)
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
		it.Event = new(SafeProxyFactoryProxyCreation)
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
func (it *SafeProxyFactoryProxyCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeProxyFactoryProxyCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeProxyFactoryProxyCreation represents a ProxyCreation event raised by the SafeProxyFactory contract.
type SafeProxyFactoryProxyCreation struct {
	Proxy     common.Address
	Singleton common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProxyCreation is a free log retrieval operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactory *SafeProxyFactoryFilterer) FilterProxyCreation(opts *bind.FilterOpts, proxy []common.Address) (*SafeProxyFactoryProxyCreationIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactory.contract.FilterLogs(opts, "ProxyCreation", proxyRule)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryProxyCreationIterator{contract: _SafeProxyFactory.contract, event: "ProxyCreation", logs: logs, sub: sub}, nil
}

// WatchProxyCreation is a free log subscription operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactory *SafeProxyFactoryFilterer) WatchProxyCreation(opts *bind.WatchOpts, sink chan<- *SafeProxyFactoryProxyCreation, proxy []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactory.contract.WatchLogs(opts, "ProxyCreation", proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeProxyFactoryProxyCreation)
				if err := _SafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
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

// ParseProxyCreation is a log parse operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactory *SafeProxyFactoryFilterer) ParseProxyCreation(log types.Log) (*SafeProxyFactoryProxyCreation, error) {
	event := new(SafeProxyFactoryProxyCreation)
	if err := _SafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
