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

// SchemaRecord is an auto generated low-level Go binding around an user-defined struct.
type SchemaRecord struct {
	Uid       [32]byte
	Resolver  common.Address
	Revocable bool
	Schema    string
}

// SchemaRegistryMetaData contains all meta data concerning the SchemaRegistry contract.
var SchemaRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getSchema\",\"inputs\":[{\"name\":\"uid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSchemaRecord\",\"components\":[{\"name\":\"uid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\",\"internalType\":\"contractISchemaResolver\"},{\"name\":\"revocable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"schema\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"schema\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resolver\",\"type\":\"address\",\"internalType\":\"contractISchemaResolver\"},{\"name\":\"revocable\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Registered\",\"inputs\":[{\"name\":\"uid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registerer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"schema\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSchemaRecord\",\"components\":[{\"name\":\"uid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\",\"internalType\":\"contractISchemaResolver\"},{\"name\":\"revocable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"schema\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyExists\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b50610856806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806354fd4d501461004657806360d7a27814610080578063a2ea7c6e146100a0575b600080fd5b61006a604051806040016040528060058152602001640312e332e360dc1b81525081565b60405161007791906103bb565b60405180910390f35b61009361008e366004610478565b6100c0565b60405161007791906104ef565b6100b36100ae36600461050e565b61021b565b60405161007791906105dc565b60008060405180608001604052806000801b8152602001856001600160a01b03168152602001841515815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093909452509293509150610135905082610325565b600081815260208190526040902054909150156101655760405163119b4fd360e11b815260040160405180910390fd5b808252600081815260208181526040918290208451815590840151600182018054938601511515600160a01b026001600160a81b03199094166001600160a01b039092169190911792909217909155606083015183919060028201906101cb90826106e6565b50905050336001600160a01b0316817fd0b86852e21f9e5fa4bc3b0cff9757ffe243d50c4b43968a42202153d651ea5e8460405161020991906105dc565b60405180910390a39695505050505050565b60408051608081018252600080825260208201819052918101919091526060808201526000828152602081815260409182902082516080810184528154815260018201546001600160a01b03811693820193909352600160a01b90920460ff1615159282019290925260028201805491929160608401919061029c90610619565b80601f01602080910402602001604051908101604052809291908181526020018280546102c890610619565b80156103155780601f106102ea57610100808354040283529160200191610315565b820191906000526020600020905b8154815290600101906020018083116102f857829003601f168201915b5050505050815250509050919050565b600081606001518260200151836040015160405160200161034893929190610816565b604051602081830303815290604052805190602001209050919050565b60005b83811015610380578181015183820152602001610368565b50506000910152565b6000610393825190565b8084526020840193506103aa818560208601610365565b601f01601f19169290920192915050565b602080825281016103cc8184610389565b9392505050565b60008083601f8401126103e8576103e8600080fd5b50813567ffffffffffffffff81111561040357610403600080fd5b60208301915083600182028301111561041e5761041e600080fd5b9250929050565b60006001600160a01b0382165b92915050565b600061043282610425565b61044c81610438565b811461045757600080fd5b50565b803561043281610443565b80151561044c565b803561043281610465565b6000806000806060858703121561049157610491600080fd5b843567ffffffffffffffff8111156104ab576104ab600080fd5b6104b7878288016103d3565b945094505060206104ca8782880161045a565b92505060406104db8782880161046d565b91505092959194509250565b805b82525050565b6020810161043282846104e7565b8061044c565b8035610432816104fd565b60006020828403121561052357610523600080fd5b600061052f8484610503565b949350505050565b60006104326001600160a01b03831661054e565b90565b6001600160a01b031690565b600061043282610537565b60006104328261055a565b6104e981610565565b8015156104e9565b8051600090608084019061059585826104e7565b5060208301516105a86020860182610570565b5060408301516105bb6040860182610579565b50606083015184820360608601526105d38282610389565b95945050505050565b602080825281016103cc8184610581565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052602260045260246000fd5b60028104600182168061062d57607f821691505b60208210810361063f5761063f610603565b50919050565b600061043261054b8381565b61065a83610645565b815460001960089490940293841b1916921b91909117905550565b6000610682818484610651565b505050565b818110156106a25761069a600082610675565b600101610687565b5050565b601f821115610682576000818152602090206020601f850104810160208510156106cd5750805b6106df6020601f860104830182610687565b5050505050565b815167ffffffffffffffff811115610700576107006105ed565b61070a8254610619565b6107158282856106a6565b6020601f83116001811461074957600084156107315750858201515b600019600886021c19811660028602178655506107a2565b600085815260208120601f198616915b828110156107795788850151825560209485019460019092019101610759565b868310156107955784890151600019601f89166008021c191682555b6001600288020188555050505b505050505050565b60006107b4825190565b6107c2818560208601610365565b9290920192915050565b60006104328260601b90565b6000610432826107cc565b6104e96107ef82610565565b6107d8565b60006104328260f81b90565b6000610432826107f4565b6104e9811515610800565b600061082282866107aa565b915061082e82856107e3565b60148201915061083e828461080b565b50600101939250505056fea164736f6c6343000814000a",
}

// SchemaRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use SchemaRegistryMetaData.ABI instead.
var SchemaRegistryABI = SchemaRegistryMetaData.ABI

// SchemaRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SchemaRegistryMetaData.Bin instead.
var SchemaRegistryBin = SchemaRegistryMetaData.Bin

// DeploySchemaRegistry deploys a new Ethereum contract, binding an instance of SchemaRegistry to it.
func DeploySchemaRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SchemaRegistry, error) {
	parsed, err := SchemaRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SchemaRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SchemaRegistry{SchemaRegistryCaller: SchemaRegistryCaller{contract: contract}, SchemaRegistryTransactor: SchemaRegistryTransactor{contract: contract}, SchemaRegistryFilterer: SchemaRegistryFilterer{contract: contract}}, nil
}

// SchemaRegistry is an auto generated Go binding around an Ethereum contract.
type SchemaRegistry struct {
	SchemaRegistryCaller     // Read-only binding to the contract
	SchemaRegistryTransactor // Write-only binding to the contract
	SchemaRegistryFilterer   // Log filterer for contract events
}

// SchemaRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SchemaRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SchemaRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SchemaRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SchemaRegistrySession struct {
	Contract     *SchemaRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SchemaRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SchemaRegistryCallerSession struct {
	Contract *SchemaRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SchemaRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SchemaRegistryTransactorSession struct {
	Contract     *SchemaRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SchemaRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SchemaRegistryRaw struct {
	Contract *SchemaRegistry // Generic contract binding to access the raw methods on
}

// SchemaRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SchemaRegistryCallerRaw struct {
	Contract *SchemaRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SchemaRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SchemaRegistryTransactorRaw struct {
	Contract *SchemaRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSchemaRegistry creates a new instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistry(address common.Address, backend bind.ContractBackend) (*SchemaRegistry, error) {
	contract, err := bindSchemaRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistry{SchemaRegistryCaller: SchemaRegistryCaller{contract: contract}, SchemaRegistryTransactor: SchemaRegistryTransactor{contract: contract}, SchemaRegistryFilterer: SchemaRegistryFilterer{contract: contract}}, nil
}

// NewSchemaRegistryCaller creates a new read-only instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistryCaller(address common.Address, caller bind.ContractCaller) (*SchemaRegistryCaller, error) {
	contract, err := bindSchemaRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryCaller{contract: contract}, nil
}

// NewSchemaRegistryTransactor creates a new write-only instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SchemaRegistryTransactor, error) {
	contract, err := bindSchemaRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryTransactor{contract: contract}, nil
}

// NewSchemaRegistryFilterer creates a new log filterer instance of SchemaRegistry, bound to a specific deployed contract.
func NewSchemaRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SchemaRegistryFilterer, error) {
	contract, err := bindSchemaRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryFilterer{contract: contract}, nil
}

// bindSchemaRegistry binds a generic wrapper to an already deployed contract.
func bindSchemaRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SchemaRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SchemaRegistry *SchemaRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SchemaRegistry.Contract.SchemaRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SchemaRegistry *SchemaRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.SchemaRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SchemaRegistry *SchemaRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.SchemaRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SchemaRegistry *SchemaRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SchemaRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SchemaRegistry *SchemaRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SchemaRegistry *SchemaRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetSchema is a free data retrieval call binding the contract method 0xa2ea7c6e.
//
// Solidity: function getSchema(bytes32 uid) view returns((bytes32,address,bool,string))
func (_SchemaRegistry *SchemaRegistryCaller) GetSchema(opts *bind.CallOpts, uid [32]byte) (SchemaRecord, error) {
	var out []interface{}
	err := _SchemaRegistry.contract.Call(opts, &out, "getSchema", uid)

	if err != nil {
		return *new(SchemaRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(SchemaRecord)).(*SchemaRecord)

	return out0, err

}

// GetSchema is a free data retrieval call binding the contract method 0xa2ea7c6e.
//
// Solidity: function getSchema(bytes32 uid) view returns((bytes32,address,bool,string))
func (_SchemaRegistry *SchemaRegistrySession) GetSchema(uid [32]byte) (SchemaRecord, error) {
	return _SchemaRegistry.Contract.GetSchema(&_SchemaRegistry.CallOpts, uid)
}

// GetSchema is a free data retrieval call binding the contract method 0xa2ea7c6e.
//
// Solidity: function getSchema(bytes32 uid) view returns((bytes32,address,bool,string))
func (_SchemaRegistry *SchemaRegistryCallerSession) GetSchema(uid [32]byte) (SchemaRecord, error) {
	return _SchemaRegistry.Contract.GetSchema(&_SchemaRegistry.CallOpts, uid)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SchemaRegistry *SchemaRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SchemaRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SchemaRegistry *SchemaRegistrySession) Version() (string, error) {
	return _SchemaRegistry.Contract.Version(&_SchemaRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SchemaRegistry *SchemaRegistryCallerSession) Version() (string, error) {
	return _SchemaRegistry.Contract.Version(&_SchemaRegistry.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0x60d7a278.
//
// Solidity: function register(string schema, address resolver, bool revocable) returns(bytes32)
func (_SchemaRegistry *SchemaRegistryTransactor) Register(opts *bind.TransactOpts, schema string, resolver common.Address, revocable bool) (*types.Transaction, error) {
	return _SchemaRegistry.contract.Transact(opts, "register", schema, resolver, revocable)
}

// Register is a paid mutator transaction binding the contract method 0x60d7a278.
//
// Solidity: function register(string schema, address resolver, bool revocable) returns(bytes32)
func (_SchemaRegistry *SchemaRegistrySession) Register(schema string, resolver common.Address, revocable bool) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Register(&_SchemaRegistry.TransactOpts, schema, resolver, revocable)
}

// Register is a paid mutator transaction binding the contract method 0x60d7a278.
//
// Solidity: function register(string schema, address resolver, bool revocable) returns(bytes32)
func (_SchemaRegistry *SchemaRegistryTransactorSession) Register(schema string, resolver common.Address, revocable bool) (*types.Transaction, error) {
	return _SchemaRegistry.Contract.Register(&_SchemaRegistry.TransactOpts, schema, resolver, revocable)
}

// SchemaRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the SchemaRegistry contract.
type SchemaRegistryRegisteredIterator struct {
	Event *SchemaRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *SchemaRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SchemaRegistryRegistered)
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
		it.Event = new(SchemaRegistryRegistered)
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
func (it *SchemaRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SchemaRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SchemaRegistryRegistered represents a Registered event raised by the SchemaRegistry contract.
type SchemaRegistryRegistered struct {
	Uid        [32]byte
	Registerer common.Address
	Schema     SchemaRecord
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xd0b86852e21f9e5fa4bc3b0cff9757ffe243d50c4b43968a42202153d651ea5e.
//
// Solidity: event Registered(bytes32 indexed uid, address indexed registerer, (bytes32,address,bool,string) schema)
func (_SchemaRegistry *SchemaRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, uid [][32]byte, registerer []common.Address) (*SchemaRegistryRegisteredIterator, error) {

	var uidRule []interface{}
	for _, uidItem := range uid {
		uidRule = append(uidRule, uidItem)
	}
	var registererRule []interface{}
	for _, registererItem := range registerer {
		registererRule = append(registererRule, registererItem)
	}

	logs, sub, err := _SchemaRegistry.contract.FilterLogs(opts, "Registered", uidRule, registererRule)
	if err != nil {
		return nil, err
	}
	return &SchemaRegistryRegisteredIterator{contract: _SchemaRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xd0b86852e21f9e5fa4bc3b0cff9757ffe243d50c4b43968a42202153d651ea5e.
//
// Solidity: event Registered(bytes32 indexed uid, address indexed registerer, (bytes32,address,bool,string) schema)
func (_SchemaRegistry *SchemaRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *SchemaRegistryRegistered, uid [][32]byte, registerer []common.Address) (event.Subscription, error) {

	var uidRule []interface{}
	for _, uidItem := range uid {
		uidRule = append(uidRule, uidItem)
	}
	var registererRule []interface{}
	for _, registererItem := range registerer {
		registererRule = append(registererRule, registererItem)
	}

	logs, sub, err := _SchemaRegistry.contract.WatchLogs(opts, "Registered", uidRule, registererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SchemaRegistryRegistered)
				if err := _SchemaRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0xd0b86852e21f9e5fa4bc3b0cff9757ffe243d50c4b43968a42202153d651ea5e.
//
// Solidity: event Registered(bytes32 indexed uid, address indexed registerer, (bytes32,address,bool,string) schema)
func (_SchemaRegistry *SchemaRegistryFilterer) ParseRegistered(log types.Log) (*SchemaRegistryRegistered, error) {
	event := new(SchemaRegistryRegistered)
	if err := _SchemaRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
