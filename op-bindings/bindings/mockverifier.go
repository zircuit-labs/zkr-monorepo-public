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

// MockVerifierMetaData contains all meta data concerning the MockVerifier contract.
var MockVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061010b806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806354fd4d50146039575b604080516020810190915260009052005b603f6053565b604051604a919060c0565b60405180910390f35b6040518060600160405280602881526020016100d76028913981565b60005b8381101560885781810151838201526020016072565b50506000910152565b6000609a825190565b80845260208401935060af818560208601606f565b601f01601f19169290920192915050565b6020808252810160cf81846091565b939250505056fe30303030303030303030303030303030303030303030303030303030303030303030303030303030a164736f6c6343000814000a",
}

// MockVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use MockVerifierMetaData.ABI instead.
var MockVerifierABI = MockVerifierMetaData.ABI

// MockVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockVerifierMetaData.Bin instead.
var MockVerifierBin = MockVerifierMetaData.Bin

// DeployMockVerifier deploys a new Ethereum contract, binding an instance of MockVerifier to it.
func DeployMockVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockVerifier, error) {
	parsed, err := MockVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockVerifier{MockVerifierCaller: MockVerifierCaller{contract: contract}, MockVerifierTransactor: MockVerifierTransactor{contract: contract}, MockVerifierFilterer: MockVerifierFilterer{contract: contract}}, nil
}

// MockVerifier is an auto generated Go binding around an Ethereum contract.
type MockVerifier struct {
	MockVerifierCaller     // Read-only binding to the contract
	MockVerifierTransactor // Write-only binding to the contract
	MockVerifierFilterer   // Log filterer for contract events
}

// MockVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockVerifierSession struct {
	Contract     *MockVerifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockVerifierCallerSession struct {
	Contract *MockVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MockVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockVerifierTransactorSession struct {
	Contract     *MockVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MockVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockVerifierRaw struct {
	Contract *MockVerifier // Generic contract binding to access the raw methods on
}

// MockVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockVerifierCallerRaw struct {
	Contract *MockVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// MockVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockVerifierTransactorRaw struct {
	Contract *MockVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockVerifier creates a new instance of MockVerifier, bound to a specific deployed contract.
func NewMockVerifier(address common.Address, backend bind.ContractBackend) (*MockVerifier, error) {
	contract, err := bindMockVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockVerifier{MockVerifierCaller: MockVerifierCaller{contract: contract}, MockVerifierTransactor: MockVerifierTransactor{contract: contract}, MockVerifierFilterer: MockVerifierFilterer{contract: contract}}, nil
}

// NewMockVerifierCaller creates a new read-only instance of MockVerifier, bound to a specific deployed contract.
func NewMockVerifierCaller(address common.Address, caller bind.ContractCaller) (*MockVerifierCaller, error) {
	contract, err := bindMockVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockVerifierCaller{contract: contract}, nil
}

// NewMockVerifierTransactor creates a new write-only instance of MockVerifier, bound to a specific deployed contract.
func NewMockVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*MockVerifierTransactor, error) {
	contract, err := bindMockVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockVerifierTransactor{contract: contract}, nil
}

// NewMockVerifierFilterer creates a new log filterer instance of MockVerifier, bound to a specific deployed contract.
func NewMockVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*MockVerifierFilterer, error) {
	contract, err := bindMockVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockVerifierFilterer{contract: contract}, nil
}

// bindMockVerifier binds a generic wrapper to an already deployed contract.
func bindMockVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockVerifier *MockVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockVerifier.Contract.MockVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockVerifier *MockVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockVerifier.Contract.MockVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockVerifier *MockVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockVerifier.Contract.MockVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockVerifier *MockVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockVerifier *MockVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockVerifier *MockVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockVerifier.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MockVerifier *MockVerifierCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockVerifier.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MockVerifier *MockVerifierSession) Version() (string, error) {
	return _MockVerifier.Contract.Version(&_MockVerifier.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MockVerifier *MockVerifierCallerSession) Version() (string, error) {
	return _MockVerifier.Contract.Version(&_MockVerifier.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_MockVerifier *MockVerifierTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MockVerifier.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_MockVerifier *MockVerifierSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MockVerifier.Contract.Fallback(&_MockVerifier.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_MockVerifier *MockVerifierTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MockVerifier.Contract.Fallback(&_MockVerifier.TransactOpts, calldata)
}
