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

// StorageSetterSlot is an auto generated low-level Go binding around an user-defined struct.
type StorageSetterSlot struct {
	Key   [32]byte
	Value [32]byte
}

// StorageSetterMetaData contains all meta data concerning the StorageSetter contract.
var StorageSetterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OWNER_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimOwnership\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBool\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"value_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBytes32\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"value_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUint\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"value_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBool\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBytes32\",\"inputs\":[{\"name\":\"slots\",\"type\":\"tuple[]\",\"internalType\":\"structStorageSetter.Slot[]\",\"components\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBytes32\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUint\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516107af3803806107af83398101604081905261002f9161008a565b61005d61ffff82167f255d19709eaa5389f721d4939b6f0af0f9794e8701d4785985dcd80eb2c48fa66100c9565b608052506100dc565b61ffff8116811461007657600080fd5b50565b805161008481610066565b92915050565b60006020828403121561009f5761009f600080fd5b60006100ab8484610079565b949350505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610084576100846100b3565b6080516106a361010c6000396000818161015b01528181610259015281816102b301526102de01526106a36000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063a6ed563e11610071578063a6ed563e1461018a578063abfdcced14610198578063bd02d0f51461018a578063ca446dd9146101a6578063da92d3ae146101b4578063e2a4853a146100f757600080fd5b80630528afe2146100b957806321f8a721146100ce5780634e91db08146100f757806354fd4d501461010a5780637ae1cfca1461013b578063963949a314610156575b600080fd5b6100cc6100c7366004610396565b6101c7565b005b6100e16100dc3660046103f6565b610232565b6040516100ee919061043f565b60405180910390f35b6100cc61010536600461044d565b610242565b61012e604051806040016040528060058152602001640322e302e360dc1b81525081565b6040516100ee91906104e0565b6101496100dc3660046103f6565b6040516100ee9190610500565b61017d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516100ee9190610514565b61017d6100dc3660046103f6565b6100cc610105366004610535565b6100cc61010536600461057c565b6100cc6101c23660046105af565b610252565b6101cf6102d7565b8060005b8181101561022c5761021a8484838181106101f0576101f06105d0565b9050604002016000013585858481811061020c5761020c6105d0565b905060400201602001359055565b80610224816105fc565b9150506101d3565b50505050565b600061023c825490565b92915050565b61024a6102d7565b9055565b5050565b600061027c7f00000000000000000000000000000000000000000000000000000000000000005490565b90506001600160a01b038116156102ae5760405162461bcd60e51b81526004016102a590610616565b60405180910390fd5b61024e7f0000000000000000000000000000000000000000000000000000000000000000839055565b60006103017f00000000000000000000000000000000000000000000000000000000000000005490565b90506001600160a01b038116331480159061032457506001600160a01b03811615155b156103415760405162461bcd60e51b81526004016102a590610651565b50565b60008083601f84011261035957610359600080fd5b50813567ffffffffffffffff81111561037457610374600080fd5b60208301915083604082028301111561038f5761038f600080fd5b9250929050565b600080602083850312156103ac576103ac600080fd5b823567ffffffffffffffff8111156103c6576103c6600080fd5b6103d285828601610344565b92509250509250929050565b805b811461034157600080fd5b803561023c816103de565b60006020828403121561040b5761040b600080fd5b600061041784846103eb565b949350505050565b60006001600160a01b03821661023c565b6104398161041f565b82525050565b6020810161023c8284610430565b6000806040838503121561046357610463600080fd5b600061046f85856103eb565b9250506020610480858286016103eb565b9150509250929050565b60005b838110156104a557818101518382015260200161048d565b50506000910152565b60006104b8825190565b8084526020840193506104cf81856020860161048a565b601f01601f19169290920192915050565b602080825281016104f181846104ae565b9392505050565b801515610439565b6020810161023c82846104f8565b80610439565b6020810161023c828461050e565b8015156103e0565b803561023c81610522565b6000806040838503121561054b5761054b600080fd5b600061055785856103eb565b92505060206104808582860161052a565b6103e08161041f565b803561023c81610568565b6000806040838503121561059257610592600080fd5b600061059e85856103eb565b925050602061048085828601610571565b6000602082840312156105c4576105c4600080fd5b60006104178484610571565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019820361060f5761060f6105e6565b5060010190565b6020808252810161023c81601981527f4f776e65727368697020616c726561647920636c61696d656400000000000000602082015260400190565b6020808252810161023c81602181527f4f776e657220736574206275742063616c6c6572206973206e6f74206f776e656020820152603960f91b60408201526060019056fea164736f6c6343000814000a",
}

// StorageSetterABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageSetterMetaData.ABI instead.
var StorageSetterABI = StorageSetterMetaData.ABI

// StorageSetterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageSetterMetaData.Bin instead.
var StorageSetterBin = StorageSetterMetaData.Bin

// DeployStorageSetter deploys a new Ethereum contract, binding an instance of StorageSetter to it.
func DeployStorageSetter(auth *bind.TransactOpts, backend bind.ContractBackend, offset uint16) (common.Address, *types.Transaction, *StorageSetter, error) {
	parsed, err := StorageSetterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageSetterBin), backend, offset)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSetter{StorageSetterCaller: StorageSetterCaller{contract: contract}, StorageSetterTransactor: StorageSetterTransactor{contract: contract}, StorageSetterFilterer: StorageSetterFilterer{contract: contract}}, nil
}

// StorageSetter is an auto generated Go binding around an Ethereum contract.
type StorageSetter struct {
	StorageSetterCaller     // Read-only binding to the contract
	StorageSetterTransactor // Write-only binding to the contract
	StorageSetterFilterer   // Log filterer for contract events
}

// StorageSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSetterSession struct {
	Contract     *StorageSetter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageSetterCallerSession struct {
	Contract *StorageSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StorageSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageSetterTransactorSession struct {
	Contract     *StorageSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StorageSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageSetterRaw struct {
	Contract *StorageSetter // Generic contract binding to access the raw methods on
}

// StorageSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageSetterCallerRaw struct {
	Contract *StorageSetterCaller // Generic read-only contract binding to access the raw methods on
}

// StorageSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageSetterTransactorRaw struct {
	Contract *StorageSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageSetter creates a new instance of StorageSetter, bound to a specific deployed contract.
func NewStorageSetter(address common.Address, backend bind.ContractBackend) (*StorageSetter, error) {
	contract, err := bindStorageSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSetter{StorageSetterCaller: StorageSetterCaller{contract: contract}, StorageSetterTransactor: StorageSetterTransactor{contract: contract}, StorageSetterFilterer: StorageSetterFilterer{contract: contract}}, nil
}

// NewStorageSetterCaller creates a new read-only instance of StorageSetter, bound to a specific deployed contract.
func NewStorageSetterCaller(address common.Address, caller bind.ContractCaller) (*StorageSetterCaller, error) {
	contract, err := bindStorageSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSetterCaller{contract: contract}, nil
}

// NewStorageSetterTransactor creates a new write-only instance of StorageSetter, bound to a specific deployed contract.
func NewStorageSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSetterTransactor, error) {
	contract, err := bindStorageSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSetterTransactor{contract: contract}, nil
}

// NewStorageSetterFilterer creates a new log filterer instance of StorageSetter, bound to a specific deployed contract.
func NewStorageSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSetterFilterer, error) {
	contract, err := bindStorageSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSetterFilterer{contract: contract}, nil
}

// bindStorageSetter binds a generic wrapper to an already deployed contract.
func bindStorageSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageSetterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSetter *StorageSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSetter.Contract.StorageSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSetter *StorageSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSetter.Contract.StorageSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSetter *StorageSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSetter.Contract.StorageSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSetter *StorageSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSetter *StorageSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSetter *StorageSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSetter.Contract.contract.Transact(opts, method, params...)
}

// OWNERSLOT is a free data retrieval call binding the contract method 0x963949a3.
//
// Solidity: function OWNER_SLOT() view returns(bytes32)
func (_StorageSetter *StorageSetterCaller) OWNERSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "OWNER_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNERSLOT is a free data retrieval call binding the contract method 0x963949a3.
//
// Solidity: function OWNER_SLOT() view returns(bytes32)
func (_StorageSetter *StorageSetterSession) OWNERSLOT() ([32]byte, error) {
	return _StorageSetter.Contract.OWNERSLOT(&_StorageSetter.CallOpts)
}

// OWNERSLOT is a free data retrieval call binding the contract method 0x963949a3.
//
// Solidity: function OWNER_SLOT() view returns(bytes32)
func (_StorageSetter *StorageSetterCallerSession) OWNERSLOT() ([32]byte, error) {
	return _StorageSetter.Contract.OWNERSLOT(&_StorageSetter.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _slot) view returns(address addr_)
func (_StorageSetter *StorageSetterCaller) GetAddress(opts *bind.CallOpts, _slot [32]byte) (common.Address, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "getAddress", _slot)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _slot) view returns(address addr_)
func (_StorageSetter *StorageSetterSession) GetAddress(_slot [32]byte) (common.Address, error) {
	return _StorageSetter.Contract.GetAddress(&_StorageSetter.CallOpts, _slot)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _slot) view returns(address addr_)
func (_StorageSetter *StorageSetterCallerSession) GetAddress(_slot [32]byte) (common.Address, error) {
	return _StorageSetter.Contract.GetAddress(&_StorageSetter.CallOpts, _slot)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _slot) view returns(bool value_)
func (_StorageSetter *StorageSetterCaller) GetBool(opts *bind.CallOpts, _slot [32]byte) (bool, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "getBool", _slot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _slot) view returns(bool value_)
func (_StorageSetter *StorageSetterSession) GetBool(_slot [32]byte) (bool, error) {
	return _StorageSetter.Contract.GetBool(&_StorageSetter.CallOpts, _slot)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _slot) view returns(bool value_)
func (_StorageSetter *StorageSetterCallerSession) GetBool(_slot [32]byte) (bool, error) {
	return _StorageSetter.Contract.GetBool(&_StorageSetter.CallOpts, _slot)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _slot) view returns(bytes32 value_)
func (_StorageSetter *StorageSetterCaller) GetBytes32(opts *bind.CallOpts, _slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "getBytes32", _slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _slot) view returns(bytes32 value_)
func (_StorageSetter *StorageSetterSession) GetBytes32(_slot [32]byte) ([32]byte, error) {
	return _StorageSetter.Contract.GetBytes32(&_StorageSetter.CallOpts, _slot)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _slot) view returns(bytes32 value_)
func (_StorageSetter *StorageSetterCallerSession) GetBytes32(_slot [32]byte) ([32]byte, error) {
	return _StorageSetter.Contract.GetBytes32(&_StorageSetter.CallOpts, _slot)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _slot) view returns(uint256 value_)
func (_StorageSetter *StorageSetterCaller) GetUint(opts *bind.CallOpts, _slot [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "getUint", _slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _slot) view returns(uint256 value_)
func (_StorageSetter *StorageSetterSession) GetUint(_slot [32]byte) (*big.Int, error) {
	return _StorageSetter.Contract.GetUint(&_StorageSetter.CallOpts, _slot)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _slot) view returns(uint256 value_)
func (_StorageSetter *StorageSetterCallerSession) GetUint(_slot [32]byte) (*big.Int, error) {
	return _StorageSetter.Contract.GetUint(&_StorageSetter.CallOpts, _slot)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StorageSetter *StorageSetterCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StorageSetter.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StorageSetter *StorageSetterSession) Version() (string, error) {
	return _StorageSetter.Contract.Version(&_StorageSetter.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StorageSetter *StorageSetterCallerSession) Version() (string, error) {
	return _StorageSetter.Contract.Version(&_StorageSetter.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0xda92d3ae.
//
// Solidity: function claimOwnership(address _owner) returns()
func (_StorageSetter *StorageSetterTransactor) ClaimOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "claimOwnership", _owner)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0xda92d3ae.
//
// Solidity: function claimOwnership(address _owner) returns()
func (_StorageSetter *StorageSetterSession) ClaimOwnership(_owner common.Address) (*types.Transaction, error) {
	return _StorageSetter.Contract.ClaimOwnership(&_StorageSetter.TransactOpts, _owner)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0xda92d3ae.
//
// Solidity: function claimOwnership(address _owner) returns()
func (_StorageSetter *StorageSetterTransactorSession) ClaimOwnership(_owner common.Address) (*types.Transaction, error) {
	return _StorageSetter.Contract.ClaimOwnership(&_StorageSetter.TransactOpts, _owner)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _slot, address _address) returns()
func (_StorageSetter *StorageSetterTransactor) SetAddress(opts *bind.TransactOpts, _slot [32]byte, _address common.Address) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "setAddress", _slot, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _slot, address _address) returns()
func (_StorageSetter *StorageSetterSession) SetAddress(_slot [32]byte, _address common.Address) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetAddress(&_StorageSetter.TransactOpts, _slot, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 _slot, address _address) returns()
func (_StorageSetter *StorageSetterTransactorSession) SetAddress(_slot [32]byte, _address common.Address) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetAddress(&_StorageSetter.TransactOpts, _slot, _address)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 _slot, bool _value) returns()
func (_StorageSetter *StorageSetterTransactor) SetBool(opts *bind.TransactOpts, _slot [32]byte, _value bool) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "setBool", _slot, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 _slot, bool _value) returns()
func (_StorageSetter *StorageSetterSession) SetBool(_slot [32]byte, _value bool) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBool(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 _slot, bool _value) returns()
func (_StorageSetter *StorageSetterTransactorSession) SetBool(_slot [32]byte, _value bool) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBool(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x0528afe2.
//
// Solidity: function setBytes32((bytes32,bytes32)[] slots) returns()
func (_StorageSetter *StorageSetterTransactor) SetBytes32(opts *bind.TransactOpts, slots []StorageSetterSlot) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "setBytes32", slots)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x0528afe2.
//
// Solidity: function setBytes32((bytes32,bytes32)[] slots) returns()
func (_StorageSetter *StorageSetterSession) SetBytes32(slots []StorageSetterSlot) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBytes32(&_StorageSetter.TransactOpts, slots)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x0528afe2.
//
// Solidity: function setBytes32((bytes32,bytes32)[] slots) returns()
func (_StorageSetter *StorageSetterTransactorSession) SetBytes32(slots []StorageSetterSlot) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBytes32(&_StorageSetter.TransactOpts, slots)
}

// SetBytes320 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 _slot, bytes32 _value) returns()
func (_StorageSetter *StorageSetterTransactor) SetBytes320(opts *bind.TransactOpts, _slot [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "setBytes320", _slot, _value)
}

// SetBytes320 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 _slot, bytes32 _value) returns()
func (_StorageSetter *StorageSetterSession) SetBytes320(_slot [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBytes320(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetBytes320 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 _slot, bytes32 _value) returns()
func (_StorageSetter *StorageSetterTransactorSession) SetBytes320(_slot [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetBytes320(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 _slot, uint256 _value) returns()
func (_StorageSetter *StorageSetterTransactor) SetUint(opts *bind.TransactOpts, _slot [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _StorageSetter.contract.Transact(opts, "setUint", _slot, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 _slot, uint256 _value) returns()
func (_StorageSetter *StorageSetterSession) SetUint(_slot [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetUint(&_StorageSetter.TransactOpts, _slot, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 _slot, uint256 _value) returns()
func (_StorageSetter *StorageSetterTransactorSession) SetUint(_slot [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _StorageSetter.Contract.SetUint(&_StorageSetter.TransactOpts, _slot, _value)
}
