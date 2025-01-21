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

// L2ToL1MessagePasserMetaData contains all meta data concerning the L2ToL1MessagePasser contract.
var L2ToL1MessagePasserMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MESSAGE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accessController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAccessControlPausable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ethThrottleWithdrawals\",\"inputs\":[],\"outputs\":[{\"name\":\"maxAmountPerPeriod\",\"type\":\"uint208\",\"internalType\":\"uint208\"},{\"name\":\"periodLength\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxAmountTotal\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEthThrottleWithdrawalsCredits\",\"inputs\":[],\"outputs\":[{\"name\":\"availableCredits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateWithdrawal\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"messageNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sentMessages\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setEthThrottleWithdrawalsMaxAmount\",\"inputs\":[{\"name\":\"maxAmountPerPeriod\",\"type\":\"uint208\",\"internalType\":\"uint208\"},{\"name\":\"maxAmountTotal\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEthThrottleWithdrawalsPeriodLength\",\"inputs\":[{\"name\":\"_periodLength\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessagePassed\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"withdrawalHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawerBalanceBurnt\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b62000196565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff1615906001600160401b03166000811580156200006d5750825b90506000826001600160401b031660011480156200008a5750303b155b90508115801562000099575080155b15620000b85760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315620000e757845460ff60401b1916680100000000000000001785555b600e80546001600160a01b03191673420000000000000000000000000000000000010017905583156200015b57845460ff60401b191685556040517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290620001529060019062000186565b60405180910390a15b5050505050565b60006001600160401b0382165b92915050565b620001808162000162565b82525050565b602081016200016f828462000175565b61122380620001a66000396000f3fe6080604052600436106100c65760003560e01c80638129fc1c1161007f578063bc43cbaf11610059578063bc43cbaf14610259578063c2b3e5ac14610286578063e07ffaf214610299578063ecc70428146102bb57600080fd5b80638129fc1c146101f457806382e3702d14610209578063b26221701461023957600080fd5b80630915ba01146100ef578063393655241461013d5780633f827a5a1461015d57806344df8e701461017f57806354fd4d50146101945780635c975abb146101d257600080fd5b366100ea576100e833620186a0604051806020016040528060008152506102d0565b005b600080fd5b3480156100fb57600080fd5b50600254600454610125916001600160d01b03811691600160d01b90910465ffffffffffff169083565b60405161013493929190610a94565b60405180910390f35b34801561014957600080fd5b506100e8610158366004610ae4565b6103e5565b34801561016957600080fd5b50610172600181565b6040516101349190610b0f565b34801561018b57600080fd5b506100e86103f3565b3480156101a057600080fd5b506101c5604051806040016040528060058152602001640312e312e360dc1b81525081565b6040516101349190610b73565b3480156101de57600080fd5b506101e761042b565b6040516101349190610b93565b34801561020057600080fd5b506100e861049e565b34801561021557600080fd5b506101e7610224366004610bb2565b60006020819052908152604090205460ff1681565b34801561024557600080fd5b506100e8610254366004610bed565b6105cf565b34801561026557600080fd5b50600e54610279906001600160a01b031681565b6040516101349190610c5a565b6100e8610294366004610d77565b6102d0565b3480156102a557600080fd5b506102ae6105fd565b6040516101349190610de2565b3480156102c757600080fd5b506102ae61060b565b6102d861042b565b156102fe5760405162461bcd60e51b81526004016102f590610df0565b60405180910390fd5b61030c600260008034610620565b60006103536040518060c0016040528061032461060b565b81523360208201526001600160a01b03871660408201523460608201526080810186905260a001849052610722565b6000818152602081905260409020805460ff1916600117905590506001600160a01b0384163361038161060b565b7f02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054348787876040516103b69493929190610e2b565b60405180910390a45050600180546001600160f01b038082168301166001600160f01b03199091161790555050565b6103f081600261076f565b50565b476103fd816107c5565b60405181907f7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f90600090a250565b600e5460408051635c975abb60e01b815290516000926001600160a01b031691635c975abb9160048083019260209291908290030181865afa158015610475573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104999190610e83565b905090565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156104e45750825b905060008267ffffffffffffffff1660011480156105015750303b155b90508115801561050f575080155b1561052d5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561055757845460ff60401b1916600160401b1785555b600e80546001600160a01b03191673420000000000000000000000000000000000010017905583156105c857845460ff60401b191685556040517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2906105bf90600190610ebf565b60405180910390a15b5050505050565b80156105ed5760405162461bcd60e51b81526004016102f590610f20565b6105f9828260026107f4565b5050565b60006104996000600261089b565b6001546001600160f01b0316600160f01b1790565b835460028501546001600160d01b0390911690801580159061064a5750806106488486610f46565b115b156106675760405162461bcd60e51b81526004016102f590610fb3565b8160000361067657505061071c565b6001600160a01b038516600090815260018701602052604090208654815465ffffffffffff600160d01b928390048116926001600160d01b038316920416420382868202816106c7576106c7610fc3565b0482019150858211156106d8578591505b818711156106f85760405162461bcd60e51b81526004016102f59061102b565b508590036001600160d01b0316600160d01b4265ffffffffffff1602179091555050505b50505050565b80516020808301516040808501516060860151608087015160a08801519351600097610752979096959101611044565b604051602081830303815290604052805190602001209050919050565b61077833610949565b8165ffffffffffff166000036107a05760405162461bcd60e51b81526004016102f5906110ec565b805465ffffffffffff909216600160d01b026001600160d01b03909216919091179055565b806040516107d290610a63565b6040518091039082f09050801580156107ef573d6000803e3d6000fd5b505050565b805460028201546001600160d01b0391821691851682108015906108185750808411155b1561082b57610826336109d6565b610834565b61083433610949565b6001600160d01b0385161580159061085a57508254600160d01b900465ffffffffffff16155b156108735782546001600160d01b031660e160d41b1783555b505080546001600160d01b0319166001600160d01b0393909316929092178255600290910155565b80546000906001600160d01b03168082036108bb57600019915050610943565b82546001600160a01b0385166000908152600185016020526040812080546001600160d01b038116955065ffffffffffff600160d01b94859004811694929392610907920416426110fc565b905065ffffffffffff831661091c828661110f565b610926919061112e565b6109309086610f46565b94508385111561093e578394505b505050505b92915050565b600e5460405163ee2a6b8760e01b81526001600160a01b039091169063ee2a6b8790610979908490600401611142565b602060405180830381865afa158015610996573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ba9190610e83565b6103f05760405162461bcd60e51b81526004016102f59061119e565b600e54604051631239bc9960e11b81526001600160a01b0390911690632473793290610a06908490600401611142565b602060405180830381865afa158015610a23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a479190610e83565b6103f05760405162461bcd60e51b81526004016102f5906111fe565b60088061120f83390190565b6001600160d01b0381165b82525050565b65ffffffffffff8116610a7a565b80610a7a565b60608101610aa28286610a6f565b610aaf6020830185610a80565b610abc6040830184610a8e565b949350505050565b65ffffffffffff81165b81146103f057600080fd5b803561094381610ac4565b600060208284031215610af957610af9600080fd5b6000610abc8484610ad9565b61ffff8116610a7a565b602081016109438284610b05565b60005b83811015610b38578181015183820152602001610b20565b50506000910152565b6000610b4b825190565b808452602084019350610b62818560208601610b1d565b601f01601f19169290920192915050565b60208082528101610b848184610b41565b9392505050565b801515610a7a565b602081016109438284610b8b565b80610ace565b803561094381610ba1565b600060208284031215610bc757610bc7600080fd5b6000610abc8484610ba7565b6001600160d01b038116610ace565b803561094381610bd3565b60008060408385031215610c0357610c03600080fd5b6000610c0f8585610be2565b9250506020610c2085828601610ba7565b9150509250929050565b60006001600160a01b038216610943565b600061094382610c2a565b600061094382610c3b565b610a7a81610c46565b602081016109438284610c51565b610ace81610c2a565b803561094381610c68565b634e487b7160e01b600052604160045260246000fd5b601f19601f830116810181811067ffffffffffffffff82111715610cb857610cb8610c7c565b6040525050565b6000610cca60405190565b9050610cd68282610c92565b919050565b600067ffffffffffffffff821115610cf557610cf5610c7c565b601f19601f83011660200192915050565b82818337506000910152565b6000610d25610d2084610cdb565b610cbf565b905082815260208101848484011115610d4057610d40600080fd5b610d4b848285610d06565b509392505050565b600082601f830112610d6757610d67600080fd5b8135610abc848260208601610d12565b600080600060608486031215610d8f57610d8f600080fd5b6000610d9b8686610c71565b9350506020610dac86828701610ba7565b925050604084013567ffffffffffffffff811115610dcc57610dcc600080fd5b610dd886828701610d53565b9150509250925092565b602081016109438284610a8e565b6020808252810161094381601b81527f4c32546f4c314d6573736167655061737365723a207061757365640000000000602082015260400190565b60808101610e398287610a8e565b610e466020830186610a8e565b8181036040830152610e588185610b41565b9050610e676060830184610a8e565b95945050505050565b801515610ace565b805161094381610e70565b600060208284031215610e9857610e98600080fd5b6000610abc8484610e78565b600067ffffffffffffffff8216610943565b610a7a81610ea4565b602081016109438284610eb6565b603381526000602082017f4c32546f4c314d6573736167655061737365723a206d617820746f74616c20618152721b5bdd5b9d081b9bdd081cdd5c1c1bdc9d1959606a1b602082015291505b5060400190565b6020808252810161094381610ecd565b634e487b7160e01b600052601160045260246000fd5b8082018082111561094357610943610f30565b603781526000602082017f5472616e736665725468726f74746c653a206d6178696d756d20616c6c6f776581527f6420746f74616c20616d6f756e7420657863656564656400000000000000000060208201529150610f19565b6020808252810161094381610f59565b634e487b7160e01b600052601260045260246000fd5b603581526000602082017f5472616e736665725468726f74746c653a206d6178696d756d20616c6c6f776581527419081d1a1c9bdd59da1c1d5d08195e18d959591959605a1b60208201529150610f19565b6020808252810161094381610fd9565b610a7a81610c2a565b60c081016110528289610a8e565b61105f602083018861103b565b61106c604083018761103b565b6110796060830186610a8e565b6110866080830185610a8e565b81810360a08301526110988184610b41565b98975050505050505050565b602b81526000602082017f5472616e736665725468726f74746c653a20706572696f64206c656e6774682081526a063616e6e6f7420626520360ac1b60208201529150610f19565b60208082528101610943816110a4565b8181038181111561094357610943610f30565b81810280821583820485141761112757611127610f30565b5092915050565b60008261113d5761113d610fc3565b500490565b60208101610943828461103b565b603181526000602082017f4c32546f4c314d6573736167655061737365723a2073656e646572206973206e81527037ba103a343937ba3a36329030b236b4b760791b60208201529150610f19565b6020808252810161094381611150565b603381526000602082017f4c32546f4c314d6573736167655061737365723a2073656e646572206e6f7420815272616c6c6f77656420746f207468726f74746c6560681b60208201529150610f19565b60208082528101610943816111ae56fe608060405230fffea164736f6c6343000814000a",
}

// L2ToL1MessagePasserABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ToL1MessagePasserMetaData.ABI instead.
var L2ToL1MessagePasserABI = L2ToL1MessagePasserMetaData.ABI

// L2ToL1MessagePasserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ToL1MessagePasserMetaData.Bin instead.
var L2ToL1MessagePasserBin = L2ToL1MessagePasserMetaData.Bin

// DeployL2ToL1MessagePasser deploys a new Ethereum contract, binding an instance of L2ToL1MessagePasser to it.
func DeployL2ToL1MessagePasser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2ToL1MessagePasser, error) {
	parsed, err := L2ToL1MessagePasserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ToL1MessagePasserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// L2ToL1MessagePasser is an auto generated Go binding around an Ethereum contract.
type L2ToL1MessagePasser struct {
	L2ToL1MessagePasserCaller     // Read-only binding to the contract
	L2ToL1MessagePasserTransactor // Write-only binding to the contract
	L2ToL1MessagePasserFilterer   // Log filterer for contract events
}

// L2ToL1MessagePasserCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ToL1MessagePasserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ToL1MessagePasserSession struct {
	Contract     *L2ToL1MessagePasser // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ToL1MessagePasserCallerSession struct {
	Contract *L2ToL1MessagePasserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// L2ToL1MessagePasserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ToL1MessagePasserTransactorSession struct {
	Contract     *L2ToL1MessagePasserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ToL1MessagePasserRaw struct {
	Contract *L2ToL1MessagePasser // Generic contract binding to access the raw methods on
}

// L2ToL1MessagePasserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCallerRaw struct {
	Contract *L2ToL1MessagePasserCaller // Generic read-only contract binding to access the raw methods on
}

// L2ToL1MessagePasserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactorRaw struct {
	Contract *L2ToL1MessagePasserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ToL1MessagePasser creates a new instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasser(address common.Address, backend bind.ContractBackend) (*L2ToL1MessagePasser, error) {
	contract, err := bindL2ToL1MessagePasser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// NewL2ToL1MessagePasserCaller creates a new read-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserCaller(address common.Address, caller bind.ContractCaller) (*L2ToL1MessagePasserCaller, error) {
	contract, err := bindL2ToL1MessagePasser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserCaller{contract: contract}, nil
}

// NewL2ToL1MessagePasserTransactor creates a new write-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ToL1MessagePasserTransactor, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserTransactor{contract: contract}, nil
}

// NewL2ToL1MessagePasserFilterer creates a new log filterer instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ToL1MessagePasserFilterer, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserFilterer{contract: contract}, nil
}

// bindL2ToL1MessagePasser binds a generic wrapper to an already deployed contract.
func bindL2ToL1MessagePasser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2ToL1MessagePasserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transact(opts, method, params...)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MESSAGEVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "MESSAGE_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MESSAGEVERSION() (uint16, error) {
	return _L2ToL1MessagePasser.Contract.MESSAGEVERSION(&_L2ToL1MessagePasser.CallOpts)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MESSAGEVERSION() (uint16, error) {
	return _L2ToL1MessagePasser.Contract.MESSAGEVERSION(&_L2ToL1MessagePasser.CallOpts)
}

// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
//
// Solidity: function accessController() view returns(address)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) AccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "accessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
//
// Solidity: function accessController() view returns(address)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) AccessController() (common.Address, error) {
	return _L2ToL1MessagePasser.Contract.AccessController(&_L2ToL1MessagePasser.CallOpts)
}

// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
//
// Solidity: function accessController() view returns(address)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) AccessController() (common.Address, error) {
	return _L2ToL1MessagePasser.Contract.AccessController(&_L2ToL1MessagePasser.CallOpts)
}

// EthThrottleWithdrawals is a free data retrieval call binding the contract method 0x0915ba01.
//
// Solidity: function ethThrottleWithdrawals() view returns(uint208 maxAmountPerPeriod, uint48 periodLength, uint256 maxAmountTotal)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) EthThrottleWithdrawals(opts *bind.CallOpts) (struct {
	MaxAmountPerPeriod *big.Int
	PeriodLength       *big.Int
	MaxAmountTotal     *big.Int
}, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "ethThrottleWithdrawals")

	outstruct := new(struct {
		MaxAmountPerPeriod *big.Int
		PeriodLength       *big.Int
		MaxAmountTotal     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxAmountPerPeriod = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PeriodLength = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxAmountTotal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EthThrottleWithdrawals is a free data retrieval call binding the contract method 0x0915ba01.
//
// Solidity: function ethThrottleWithdrawals() view returns(uint208 maxAmountPerPeriod, uint48 periodLength, uint256 maxAmountTotal)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) EthThrottleWithdrawals() (struct {
	MaxAmountPerPeriod *big.Int
	PeriodLength       *big.Int
	MaxAmountTotal     *big.Int
}, error) {
	return _L2ToL1MessagePasser.Contract.EthThrottleWithdrawals(&_L2ToL1MessagePasser.CallOpts)
}

// EthThrottleWithdrawals is a free data retrieval call binding the contract method 0x0915ba01.
//
// Solidity: function ethThrottleWithdrawals() view returns(uint208 maxAmountPerPeriod, uint48 periodLength, uint256 maxAmountTotal)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) EthThrottleWithdrawals() (struct {
	MaxAmountPerPeriod *big.Int
	PeriodLength       *big.Int
	MaxAmountTotal     *big.Int
}, error) {
	return _L2ToL1MessagePasser.Contract.EthThrottleWithdrawals(&_L2ToL1MessagePasser.CallOpts)
}

// GetEthThrottleWithdrawalsCredits is a free data retrieval call binding the contract method 0xe07ffaf2.
//
// Solidity: function getEthThrottleWithdrawalsCredits() view returns(uint256 availableCredits)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) GetEthThrottleWithdrawalsCredits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "getEthThrottleWithdrawalsCredits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthThrottleWithdrawalsCredits is a free data retrieval call binding the contract method 0xe07ffaf2.
//
// Solidity: function getEthThrottleWithdrawalsCredits() view returns(uint256 availableCredits)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) GetEthThrottleWithdrawalsCredits() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.GetEthThrottleWithdrawalsCredits(&_L2ToL1MessagePasser.CallOpts)
}

// GetEthThrottleWithdrawalsCredits is a free data retrieval call binding the contract method 0xe07ffaf2.
//
// Solidity: function getEthThrottleWithdrawalsCredits() view returns(uint256 availableCredits)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) GetEthThrottleWithdrawalsCredits() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.GetEthThrottleWithdrawalsCredits(&_L2ToL1MessagePasser.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MessageNonce() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.MessageNonce(&_L2ToL1MessagePasser.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MessageNonce() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.MessageNonce(&_L2ToL1MessagePasser.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Paused() (bool, error) {
	return _L2ToL1MessagePasser.Contract.Paused(&_L2ToL1MessagePasser.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) Paused() (bool, error) {
	return _L2ToL1MessagePasser.Contract.Paused(&_L2ToL1MessagePasser.CallOpts)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) SentMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "sentMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.SentMessages(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.SentMessages(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) Burn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Burn() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Burn(&_L2ToL1MessagePasser.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) Burn() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Burn(&_L2ToL1MessagePasser.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Initialize() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Initialize(&_L2ToL1MessagePasser.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) Initialize() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Initialize(&_L2ToL1MessagePasser.TransactOpts)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) InitiateWithdrawal(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "initiateWithdrawal", _target, _gasLimit, _data)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) InitiateWithdrawal(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateWithdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _gasLimit, _data)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) InitiateWithdrawal(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateWithdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _gasLimit, _data)
}

// SetEthThrottleWithdrawalsMaxAmount is a paid mutator transaction binding the contract method 0xb2622170.
//
// Solidity: function setEthThrottleWithdrawalsMaxAmount(uint208 maxAmountPerPeriod, uint256 maxAmountTotal) returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) SetEthThrottleWithdrawalsMaxAmount(opts *bind.TransactOpts, maxAmountPerPeriod *big.Int, maxAmountTotal *big.Int) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "setEthThrottleWithdrawalsMaxAmount", maxAmountPerPeriod, maxAmountTotal)
}

// SetEthThrottleWithdrawalsMaxAmount is a paid mutator transaction binding the contract method 0xb2622170.
//
// Solidity: function setEthThrottleWithdrawalsMaxAmount(uint208 maxAmountPerPeriod, uint256 maxAmountTotal) returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) SetEthThrottleWithdrawalsMaxAmount(maxAmountPerPeriod *big.Int, maxAmountTotal *big.Int) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.SetEthThrottleWithdrawalsMaxAmount(&_L2ToL1MessagePasser.TransactOpts, maxAmountPerPeriod, maxAmountTotal)
}

// SetEthThrottleWithdrawalsMaxAmount is a paid mutator transaction binding the contract method 0xb2622170.
//
// Solidity: function setEthThrottleWithdrawalsMaxAmount(uint208 maxAmountPerPeriod, uint256 maxAmountTotal) returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) SetEthThrottleWithdrawalsMaxAmount(maxAmountPerPeriod *big.Int, maxAmountTotal *big.Int) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.SetEthThrottleWithdrawalsMaxAmount(&_L2ToL1MessagePasser.TransactOpts, maxAmountPerPeriod, maxAmountTotal)
}

// SetEthThrottleWithdrawalsPeriodLength is a paid mutator transaction binding the contract method 0x39365524.
//
// Solidity: function setEthThrottleWithdrawalsPeriodLength(uint48 _periodLength) returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) SetEthThrottleWithdrawalsPeriodLength(opts *bind.TransactOpts, _periodLength *big.Int) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "setEthThrottleWithdrawalsPeriodLength", _periodLength)
}

// SetEthThrottleWithdrawalsPeriodLength is a paid mutator transaction binding the contract method 0x39365524.
//
// Solidity: function setEthThrottleWithdrawalsPeriodLength(uint48 _periodLength) returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) SetEthThrottleWithdrawalsPeriodLength(_periodLength *big.Int) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.SetEthThrottleWithdrawalsPeriodLength(&_L2ToL1MessagePasser.TransactOpts, _periodLength)
}

// SetEthThrottleWithdrawalsPeriodLength is a paid mutator transaction binding the contract method 0x39365524.
//
// Solidity: function setEthThrottleWithdrawalsPeriodLength(uint48 _periodLength) returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) SetEthThrottleWithdrawalsPeriodLength(_periodLength *big.Int) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.SetEthThrottleWithdrawalsPeriodLength(&_L2ToL1MessagePasser.TransactOpts, _periodLength)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Receive() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Receive(&_L2ToL1MessagePasser.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) Receive() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Receive(&_L2ToL1MessagePasser.TransactOpts)
}

// L2ToL1MessagePasserInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserInitializedIterator struct {
	Event *L2ToL1MessagePasserInitialized // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserInitialized)
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
		it.Event = new(L2ToL1MessagePasserInitialized)
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
func (it *L2ToL1MessagePasserInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserInitialized represents a Initialized event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2ToL1MessagePasserInitializedIterator, error) {

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserInitializedIterator{contract: _L2ToL1MessagePasser.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserInitialized) (event.Subscription, error) {

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserInitialized)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseInitialized(log types.Log) (*L2ToL1MessagePasserInitialized, error) {
	event := new(L2ToL1MessagePasserInitialized)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ToL1MessagePasserMessagePassedIterator is returned from FilterMessagePassed and is used to iterate over the raw logs and unpacked data for MessagePassed events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserMessagePassedIterator struct {
	Event *L2ToL1MessagePasserMessagePassed // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserMessagePassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserMessagePassed)
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
		it.Event = new(L2ToL1MessagePasserMessagePassed)
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
func (it *L2ToL1MessagePasserMessagePassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserMessagePassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserMessagePassed represents a MessagePassed event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserMessagePassed struct {
	Nonce          *big.Int
	Sender         common.Address
	Target         common.Address
	Value          *big.Int
	GasLimit       *big.Int
	Data           []byte
	WithdrawalHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessagePassed is a free log retrieval operation binding the contract event 0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterMessagePassed(opts *bind.FilterOpts, nonce []*big.Int, sender []common.Address, target []common.Address) (*L2ToL1MessagePasserMessagePassedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserMessagePassedIterator{contract: _L2ToL1MessagePasser.contract, event: "MessagePassed", logs: logs, sub: sub}, nil
}

// WatchMessagePassed is a free log subscription operation binding the contract event 0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchMessagePassed(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserMessagePassed, nonce []*big.Int, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserMessagePassed)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "MessagePassed", log); err != nil {
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

// ParseMessagePassed is a log parse operation binding the contract event 0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseMessagePassed(log types.Log) (*L2ToL1MessagePasserMessagePassed, error) {
	event := new(L2ToL1MessagePasserMessagePassed)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "MessagePassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ToL1MessagePasserWithdrawerBalanceBurntIterator is returned from FilterWithdrawerBalanceBurnt and is used to iterate over the raw logs and unpacked data for WithdrawerBalanceBurnt events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserWithdrawerBalanceBurntIterator struct {
	Event *L2ToL1MessagePasserWithdrawerBalanceBurnt // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
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
		it.Event = new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
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
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserWithdrawerBalanceBurnt represents a WithdrawerBalanceBurnt event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserWithdrawerBalanceBurnt struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawerBalanceBurnt is a free log retrieval operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterWithdrawerBalanceBurnt(opts *bind.FilterOpts, amount []*big.Int) (*L2ToL1MessagePasserWithdrawerBalanceBurntIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "WithdrawerBalanceBurnt", amountRule)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserWithdrawerBalanceBurntIterator{contract: _L2ToL1MessagePasser.contract, event: "WithdrawerBalanceBurnt", logs: logs, sub: sub}, nil
}

// WatchWithdrawerBalanceBurnt is a free log subscription operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchWithdrawerBalanceBurnt(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserWithdrawerBalanceBurnt, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "WithdrawerBalanceBurnt", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "WithdrawerBalanceBurnt", log); err != nil {
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

// ParseWithdrawerBalanceBurnt is a log parse operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseWithdrawerBalanceBurnt(log types.Log) (*L2ToL1MessagePasserWithdrawerBalanceBurnt, error) {
	event := new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "WithdrawerBalanceBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
