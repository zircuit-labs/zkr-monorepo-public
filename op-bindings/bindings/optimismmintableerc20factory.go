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

// OptimismMintableERC20FactoryMetaData contains all meta data concerning the OptimismMintableERC20Factory contract.
var OptimismMintableERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createOptimismMintableERC20\",\"inputs\":[{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOptimismMintableERC20WithDecimals\",\"inputs\":[{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createStandardL2Token\",\"inputs\":[{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OptimismMintableERC20Created\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StandardL2TokenCreated\",\"inputs\":[{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001b6000610020565b61017e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff1615906001600160401b031660008115801561006a5750825b90506000826001600160401b031660011480156100865750303b155b905081158015610094575080155b156100b25760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b031916600117855583156100e057845460ff60401b1916680100000000000000001785555b600080546001600160a01b0319166001600160a01b038816179055831561014657845460ff60401b191685556040517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29061013d90600190610170565b60405180910390a15b505050505050565b60006001600160401b0382165b92915050565b61016a8161014e565b82525050565b6020810161015b8284610161565b61197b8061018d6000396000f3fe60806040523480156200001157600080fd5b5060043610620000875760003560e01c8063c4d66de81162000062578063c4d66de81462000106578063ce5ac90f146200011f578063e78cea921462000136578063ee9a31a2146200014a57600080fd5b806354fd4d50146200008c578063896f93d114620000c95780638cf0629c14620000ef575b600080fd5b620000b1604051806040016040528060058152602001640312e392e360dc1b81525081565b604051620000c0919062000477565b60405180910390f35b620000e0620000da366004620005db565b6200015c565b604051620000c091906200067e565b620000e062000100366004620006a5565b62000173565b6200011d620001173660046200074d565b620002ce565b005b620000e062000130366004620005db565b620003fd565b600054620000e0906001600160a01b031681565b6000546001600160a01b0316620000e0565b60006200016b848484620003fd565b949350505050565b60006001600160a01b038516620001a75760405162461bcd60e51b81526004016200019e9062000772565b60405180910390fd5b600085858585604051602001620001c29493929190620007de565b60405160208183030381529060405280519060200120905060008160008054906101000a90046001600160a01b03168888888860405162000203906200040e565b6200021395949392919062000832565b8190604051809103906000f590508015801562000234573d6000803e3d6000fd5b509050806001600160a01b0316876001600160a01b03167fceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf60405160405180910390a3866001600160a01b0316816001600160a01b03167f52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb33604051620002bc91906200067e565b60405180910390a39695505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff16600081158015620003155750825b905060008267ffffffffffffffff166001148015620003335750303b155b90508115801562000342575080155b15620003615760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156200038c57845460ff60401b1916600160401b1785555b600080546001600160a01b0319166001600160a01b0388161790558315620003f557845460ff60401b191685556040517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290620003ec90600190620008b4565b60405180910390a15b505050505050565b60006200016b848484601262000173565b6110aa80620008c583390190565b60005b83811015620004395781810151838201526020016200041f565b50506000910152565b60006200044d825190565b808452602084019350620004668185602086016200041c565b601f01601f19169290920192915050565b602080825281016200048a818462000442565b9392505050565b60006001600160a01b0382165b92915050565b620004af8162000491565b8114620004bb57600080fd5b50565b80356200049e81620004a4565b634e487b7160e01b600052604160045260246000fd5b601f19601f830116810181811067ffffffffffffffff821117156200050a576200050a620004cb565b6040525050565b60006200051d60405190565b90506200052b8282620004e1565b919050565b600067ffffffffffffffff8211156200054d576200054d620004cb565b601f19601f83011660200192915050565b82818337506000910152565b6000620005816200057b8462000530565b62000511565b9050828152602081018484840111156200059e576200059e600080fd5b620005ab8482856200055e565b509392505050565b600082601f830112620005c957620005c9600080fd5b81356200016b8482602086016200056a565b600080600060608486031215620005f557620005f5600080fd5b6000620006038686620004be565b935050602084013567ffffffffffffffff811115620006255762000625600080fd5b6200063386828701620005b3565b925050604084013567ffffffffffffffff811115620006555762000655600080fd5b6200066386828701620005b3565b9150509250925092565b620006788162000491565b82525050565b602081016200049e82846200066d565b60ff8116620004af565b80356200049e816200068e565b60008060008060808587031215620006c057620006c0600080fd5b6000620006ce8787620004be565b945050602085013567ffffffffffffffff811115620006f057620006f0600080fd5b620006fe87828801620005b3565b935050604085013567ffffffffffffffff811115620007205762000720600080fd5b6200072e87828801620005b3565b9250506060620007418782880162000698565b91505092959194509250565b600060208284031215620007645762000764600080fd5b60006200016b8484620004be565b602080825281016200049e81603f81527f4f7074696d69736d4d696e7461626c654552433230466163746f72793a206d7560208201527f73742070726f766964652072656d6f746520746f6b656e206164647265737300604082015260600190565b60ff811662000678565b60808101620007ee82876200066d565b818103602083015262000802818662000442565b9050818103604083015262000818818562000442565b9050620008296060830184620007d4565b95945050505050565b60a081016200084282886200066d565b6200085160208301876200066d565b818103604083015262000865818662000442565b905081810360608301526200087b818562000442565b90506200088c6080830184620007d4565b9695505050505050565b600067ffffffffffffffff82166200049e565b620006788162000896565b602081016200049e8284620008a956fe60e06040523480156200001157600080fd5b50604051620010aa380380620010aa8339810160408190526200003491620001fa565b82826003620000448382620003ae565b506004620000538282620003ae565b5050506001600160a01b039384166080529390921660a052505060ff1660c0526200047e565b60006001600160a01b0382165b92915050565b620000978162000079565b8114620000a357600080fd5b50565b805162000086816200008c565b634e487b7160e01b600052604160045260246000fd5b601f19601f83011681018181106001600160401b0382111715620000f157620000f1620000b3565b6040525050565b60006200010460405190565b9050620001128282620000c9565b919050565b60006001600160401b03821115620001335762000133620000b3565b601f19601f83011660200192915050565b60005b838110156200016157818101518382015260200162000147565b50506000910152565b6000620001816200017b8462000117565b620000f8565b9050828152602081018484840111156200019e576200019e600080fd5b620001ab84828562000144565b509392505050565b600082601f830112620001c957620001c9600080fd5b8151620001db8482602086016200016a565b949350505050565b60ff811662000097565b80516200008681620001e3565b600080600080600060a08688031215620002175762000217600080fd5b6000620002258888620000a6565b95505060206200023888828901620000a6565b94505060408601516001600160401b03811115620002595762000259600080fd5b6200026788828901620001b3565b93505060608601516001600160401b03811115620002885762000288600080fd5b6200029688828901620001b3565b9250506080620002a988828901620001ed565b9150509295509295909350565b634e487b7160e01b600052602260045260246000fd5b600281046001821680620002e157607f821691505b602082108103620002f657620002f6620002b6565b50919050565b6000620000866200030a8381565b90565b6200031883620002fc565b815460001960089490940293841b1916921b91909117905550565b6000620003428184846200030d565b505050565b8181101562000366576200035d60008262000333565b60010162000347565b5050565b601f82111562000342576000818152602090206020601f85010481016020851015620003935750805b620003a76020601f86010483018262000347565b5050505050565b81516001600160401b03811115620003ca57620003ca620000b3565b620003d68254620002cc565b620003e38282856200036a565b6020601f8311600181146200041a5760008415620004015750858201515b600019600886021c198116600286021786555062000476565b600085815260208120601f198616915b828110156200044c57888501518255602094850194600190920191016200042a565b86831015620004695784890151600019601f89166008021c191682555b6001600288020188555050505b505050505050565b60805160a05160c051610be0620004ca60003960006101d101526000818161028f0152818161031701528181610452015261050101526000818161015401526102b50152610be06000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806370a08231116100ad578063c01e1bd611610071578063c01e1bd6146102b3578063d6c0b2c4146102b3578063dd62ed3e146102d9578063e78cea921461028d578063ee9a31a21461031257600080fd5b806370a082311461023657806395d89b411461025f5780639dc29fac14610267578063a9059cbb1461027a578063ae1f6aaf1461028d57600080fd5b806318160ddd116100f457806318160ddd146101ab57806323b872dd146101bc578063313ce567146101cf57806340c10f19146101fd57806354fd4d501461021257600080fd5b806301ffc9a714610126578063033964be1461014f57806306fdde0314610183578063095ea7b314610198575b600080fd5b6101396101343660046108e8565b610339565b6040516101469190610913565b60405180910390f35b6101767f000000000000000000000000000000000000000000000000000000000000000081565b604051610146919061093b565b61018b610377565b604051610146919061099f565b6101396101a63660046109dc565b610409565b6002545b6040516101469190610a1f565b6101396101ca366004610a2d565b610423565b7f00000000000000000000000000000000000000000000000000000000000000006040516101469190610a86565b61021061020b3660046109dc565b610447565b005b61018b604051806040016040528060058152602001640312e332e360dc1b81525081565b6101af610244366004610a94565b6001600160a01b031660009081526020819052604090205490565b61018b6104e7565b6102106102753660046109dc565b6104f6565b6101396102883660046109dc565b610581565b7f0000000000000000000000000000000000000000000000000000000000000000610176565b7f0000000000000000000000000000000000000000000000000000000000000000610176565b6101af6102e7366004610ab5565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6101767f000000000000000000000000000000000000000000000000000000000000000081565b60006301ffc9a760e01b63ec4fc8e360e01b6001600160e01b0319841682148061036f57506001600160e01b0319848116908216145b949350505050565b60606003805461038690610afe565b80601f01602080910402602001604051908101604052809291908181526020018280546103b290610afe565b80156103ff5780601f106103d4576101008083540402835291602001916103ff565b820191906000526020600020905b8154815290600101906020018083116103e257829003601f168201915b5050505050905090565b60003361041781858561058f565b60019150505b92915050565b6000336104318582856105a1565b61043c85858561060c565b506001949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104985760405162461bcd60e51b815260040161048f90610b2a565b60405180910390fd5b6104a2828261066b565b816001600160a01b03167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516104db9190610a1f565b60405180910390a25050565b60606004805461038690610afe565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461053e5760405162461bcd60e51b815260040161048f90610b2a565b61054882826106a5565b816001600160a01b03167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516104db9190610a1f565b60003361041781858561060c565b61059c83838360016106db565b505050565b6001600160a01b03838116600090815260016020908152604080832093861683529290522054600019811461060657818110156105f757828183604051637dc7a0d960e11b815260040161048f93929190610b82565b610606848484840360006106db565b50505050565b6001600160a01b038316610636576000604051634b637e8f60e11b815260040161048f919061093b565b6001600160a01b03821661066057600060405163ec442f0560e01b815260040161048f919061093b565b61059c8383836107ae565b6001600160a01b03821661069557600060405163ec442f0560e01b815260040161048f919061093b565b6106a1600083836107ae565b5050565b6001600160a01b0382166106cf576000604051634b637e8f60e11b815260040161048f919061093b565b6106a1826000836107ae565b6001600160a01b03841661070557600060405163e602df0560e01b815260040161048f919061093b565b6001600160a01b03831661072f576000604051634a1406b160e11b815260040161048f919061093b565b6001600160a01b038085166000908152600160209081526040808320938716835292905220829055801561060657826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516107a09190610a1f565b60405180910390a350505050565b6001600160a01b0383166107d95780600260008282546107ce9190610bc0565b909155506108389050565b6001600160a01b038316600090815260208190526040902054818110156108195783818360405163391434e360e21b815260040161048f93929190610b82565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b03821661085457600280548290039055610873565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516108b69190610a1f565b60405180910390a3505050565b6001600160e01b031981165b81146108da57600080fd5b50565b803561041d816108c3565b6000602082840312156108fd576108fd600080fd5b600061036f84846108dd565b8015155b82525050565b6020810161041d8284610909565b60006001600160a01b03821661041d565b61090d81610921565b6020810161041d8284610932565b60005b8381101561096457818101518382015260200161094c565b50506000910152565b6000610977825190565b80845260208401935061098e818560208601610949565b601f01601f19169290920192915050565b602080825281016109b0818461096d565b9392505050565b6108cf81610921565b803561041d816109b7565b806108cf565b803561041d816109cb565b600080604083850312156109f2576109f2600080fd5b60006109fe85856109c0565b9250506020610a0f858286016109d1565b9150509250929050565b8061090d565b6020810161041d8284610a19565b600080600060608486031215610a4557610a45600080fd5b6000610a5186866109c0565b9350506020610a62868287016109c0565b9250506040610a73868287016109d1565b9150509250925092565b60ff811661090d565b6020810161041d8284610a7d565b600060208284031215610aa957610aa9600080fd5b600061036f84846109c0565b60008060408385031215610acb57610acb600080fd5b6000610ad785856109c0565b9250506020610a0f858286016109c0565b634e487b7160e01b600052602260045260246000fd5b600281046001821680610b1257607f821691505b602082108103610b2457610b24610ae8565b50919050565b6020808252810161041d81603481527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460208201527333b29031b0b71036b4b73a1030b73210313ab93760611b604082015260600190565b60608101610b908286610932565b610b9d6020830185610a19565b61036f6040830184610a19565b634e487b7160e01b600052601160045260246000fd5b8082018082111561041d5761041d610baa56fea164736f6c6343000814000aa164736f6c6343000814000a",
}

// OptimismMintableERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimismMintableERC20FactoryMetaData.ABI instead.
var OptimismMintableERC20FactoryABI = OptimismMintableERC20FactoryMetaData.ABI

// OptimismMintableERC20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptimismMintableERC20FactoryMetaData.Bin instead.
var OptimismMintableERC20FactoryBin = OptimismMintableERC20FactoryMetaData.Bin

// DeployOptimismMintableERC20Factory deploys a new Ethereum contract, binding an instance of OptimismMintableERC20Factory to it.
func DeployOptimismMintableERC20Factory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptimismMintableERC20Factory, error) {
	parsed, err := OptimismMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismMintableERC20FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismMintableERC20Factory{OptimismMintableERC20FactoryCaller: OptimismMintableERC20FactoryCaller{contract: contract}, OptimismMintableERC20FactoryTransactor: OptimismMintableERC20FactoryTransactor{contract: contract}, OptimismMintableERC20FactoryFilterer: OptimismMintableERC20FactoryFilterer{contract: contract}}, nil
}

// OptimismMintableERC20Factory is an auto generated Go binding around an Ethereum contract.
type OptimismMintableERC20Factory struct {
	OptimismMintableERC20FactoryCaller     // Read-only binding to the contract
	OptimismMintableERC20FactoryTransactor // Write-only binding to the contract
	OptimismMintableERC20FactoryFilterer   // Log filterer for contract events
}

// OptimismMintableERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismMintableERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismMintableERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimismMintableERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismMintableERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimismMintableERC20FactorySession struct {
	Contract     *OptimismMintableERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OptimismMintableERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimismMintableERC20FactoryCallerSession struct {
	Contract *OptimismMintableERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// OptimismMintableERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimismMintableERC20FactoryTransactorSession struct {
	Contract     *OptimismMintableERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// OptimismMintableERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryRaw struct {
	Contract *OptimismMintableERC20Factory // Generic contract binding to access the raw methods on
}

// OptimismMintableERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryCallerRaw struct {
	Contract *OptimismMintableERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// OptimismMintableERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryTransactorRaw struct {
	Contract *OptimismMintableERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimismMintableERC20Factory creates a new instance of OptimismMintableERC20Factory, bound to a specific deployed contract.
func NewOptimismMintableERC20Factory(address common.Address, backend bind.ContractBackend) (*OptimismMintableERC20Factory, error) {
	contract, err := bindOptimismMintableERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20Factory{OptimismMintableERC20FactoryCaller: OptimismMintableERC20FactoryCaller{contract: contract}, OptimismMintableERC20FactoryTransactor: OptimismMintableERC20FactoryTransactor{contract: contract}, OptimismMintableERC20FactoryFilterer: OptimismMintableERC20FactoryFilterer{contract: contract}}, nil
}

// NewOptimismMintableERC20FactoryCaller creates a new read-only instance of OptimismMintableERC20Factory, bound to a specific deployed contract.
func NewOptimismMintableERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*OptimismMintableERC20FactoryCaller, error) {
	contract, err := bindOptimismMintableERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryCaller{contract: contract}, nil
}

// NewOptimismMintableERC20FactoryTransactor creates a new write-only instance of OptimismMintableERC20Factory, bound to a specific deployed contract.
func NewOptimismMintableERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismMintableERC20FactoryTransactor, error) {
	contract, err := bindOptimismMintableERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryTransactor{contract: contract}, nil
}

// NewOptimismMintableERC20FactoryFilterer creates a new log filterer instance of OptimismMintableERC20Factory, bound to a specific deployed contract.
func NewOptimismMintableERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismMintableERC20FactoryFilterer, error) {
	contract, err := bindOptimismMintableERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryFilterer{contract: contract}, nil
}

// bindOptimismMintableERC20Factory binds a generic wrapper to an already deployed contract.
func bindOptimismMintableERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismMintableERC20Factory.Contract.OptimismMintableERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.OptimismMintableERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.OptimismMintableERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismMintableERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismMintableERC20Factory.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) BRIDGE() (common.Address, error) {
	return _OptimismMintableERC20Factory.Contract.BRIDGE(&_OptimismMintableERC20Factory.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCallerSession) BRIDGE() (common.Address, error) {
	return _OptimismMintableERC20Factory.Contract.BRIDGE(&_OptimismMintableERC20Factory.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismMintableERC20Factory.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) Bridge() (common.Address, error) {
	return _OptimismMintableERC20Factory.Contract.Bridge(&_OptimismMintableERC20Factory.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCallerSession) Bridge() (common.Address, error) {
	return _OptimismMintableERC20Factory.Contract.Bridge(&_OptimismMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OptimismMintableERC20Factory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) Version() (string, error) {
	return _OptimismMintableERC20Factory.Contract.Version(&_OptimismMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCallerSession) Version() (string, error) {
	return _OptimismMintableERC20Factory.Contract.Version(&_OptimismMintableERC20Factory.CallOpts)
}

// CreateOptimismMintableERC20 is a paid mutator transaction binding the contract method 0xce5ac90f.
//
// Solidity: function createOptimismMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactor) CreateOptimismMintableERC20(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.contract.Transact(opts, "createOptimismMintableERC20", _remoteToken, _name, _symbol)
}

// CreateOptimismMintableERC20 is a paid mutator transaction binding the contract method 0xce5ac90f.
//
// Solidity: function createOptimismMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) CreateOptimismMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateOptimismMintableERC20(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateOptimismMintableERC20 is a paid mutator transaction binding the contract method 0xce5ac90f.
//
// Solidity: function createOptimismMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorSession) CreateOptimismMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateOptimismMintableERC20(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateOptimismMintableERC20WithDecimals is a paid mutator transaction binding the contract method 0x8cf0629c.
//
// Solidity: function createOptimismMintableERC20WithDecimals(address _remoteToken, string _name, string _symbol, uint8 _decimals) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactor) CreateOptimismMintableERC20WithDecimals(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.contract.Transact(opts, "createOptimismMintableERC20WithDecimals", _remoteToken, _name, _symbol, _decimals)
}

// CreateOptimismMintableERC20WithDecimals is a paid mutator transaction binding the contract method 0x8cf0629c.
//
// Solidity: function createOptimismMintableERC20WithDecimals(address _remoteToken, string _name, string _symbol, uint8 _decimals) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) CreateOptimismMintableERC20WithDecimals(_remoteToken common.Address, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateOptimismMintableERC20WithDecimals(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol, _decimals)
}

// CreateOptimismMintableERC20WithDecimals is a paid mutator transaction binding the contract method 0x8cf0629c.
//
// Solidity: function createOptimismMintableERC20WithDecimals(address _remoteToken, string _name, string _symbol, uint8 _decimals) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorSession) CreateOptimismMintableERC20WithDecimals(_remoteToken common.Address, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateOptimismMintableERC20WithDecimals(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol, _decimals)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactor) CreateStandardL2Token(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.contract.Transact(opts, "createStandardL2Token", _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) CreateStandardL2Token(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateStandardL2Token(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorSession) CreateStandardL2Token(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateStandardL2Token(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.contract.Transact(opts, "initialize", _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.Initialize(&_OptimismMintableERC20Factory.TransactOpts, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.Initialize(&_OptimismMintableERC20Factory.TransactOpts, _bridge)
}

// OptimismMintableERC20FactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryInitializedIterator struct {
	Event *OptimismMintableERC20FactoryInitialized // Event containing the contract specifics and raw log

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
func (it *OptimismMintableERC20FactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismMintableERC20FactoryInitialized)
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
		it.Event = new(OptimismMintableERC20FactoryInitialized)
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
func (it *OptimismMintableERC20FactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimismMintableERC20FactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimismMintableERC20FactoryInitialized represents a Initialized event raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*OptimismMintableERC20FactoryInitializedIterator, error) {

	logs, sub, err := _OptimismMintableERC20Factory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryInitializedIterator{contract: _OptimismMintableERC20Factory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OptimismMintableERC20FactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _OptimismMintableERC20Factory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimismMintableERC20FactoryInitialized)
				if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) ParseInitialized(log types.Log) (*OptimismMintableERC20FactoryInitialized, error) {
	event := new(OptimismMintableERC20FactoryInitialized)
	if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator is returned from FilterOptimismMintableERC20Created and is used to iterate over the raw logs and unpacked data for OptimismMintableERC20Created events raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator struct {
	Event *OptimismMintableERC20FactoryOptimismMintableERC20Created // Event containing the contract specifics and raw log

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
func (it *OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismMintableERC20FactoryOptimismMintableERC20Created)
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
		it.Event = new(OptimismMintableERC20FactoryOptimismMintableERC20Created)
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
func (it *OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimismMintableERC20FactoryOptimismMintableERC20Created represents a OptimismMintableERC20Created event raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryOptimismMintableERC20Created struct {
	LocalToken  common.Address
	RemoteToken common.Address
	Deployer    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOptimismMintableERC20Created is a free log retrieval operation binding the contract event 0x52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb.
//
// Solidity: event OptimismMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) FilterOptimismMintableERC20Created(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address) (*OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _OptimismMintableERC20Factory.contract.FilterLogs(opts, "OptimismMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator{contract: _OptimismMintableERC20Factory.contract, event: "OptimismMintableERC20Created", logs: logs, sub: sub}, nil
}

// WatchOptimismMintableERC20Created is a free log subscription operation binding the contract event 0x52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb.
//
// Solidity: event OptimismMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) WatchOptimismMintableERC20Created(opts *bind.WatchOpts, sink chan<- *OptimismMintableERC20FactoryOptimismMintableERC20Created, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _OptimismMintableERC20Factory.contract.WatchLogs(opts, "OptimismMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimismMintableERC20FactoryOptimismMintableERC20Created)
				if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "OptimismMintableERC20Created", log); err != nil {
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

// ParseOptimismMintableERC20Created is a log parse operation binding the contract event 0x52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb.
//
// Solidity: event OptimismMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) ParseOptimismMintableERC20Created(log types.Log) (*OptimismMintableERC20FactoryOptimismMintableERC20Created, error) {
	event := new(OptimismMintableERC20FactoryOptimismMintableERC20Created)
	if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "OptimismMintableERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimismMintableERC20FactoryStandardL2TokenCreatedIterator is returned from FilterStandardL2TokenCreated and is used to iterate over the raw logs and unpacked data for StandardL2TokenCreated events raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryStandardL2TokenCreatedIterator struct {
	Event *OptimismMintableERC20FactoryStandardL2TokenCreated // Event containing the contract specifics and raw log

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
func (it *OptimismMintableERC20FactoryStandardL2TokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismMintableERC20FactoryStandardL2TokenCreated)
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
		it.Event = new(OptimismMintableERC20FactoryStandardL2TokenCreated)
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
func (it *OptimismMintableERC20FactoryStandardL2TokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimismMintableERC20FactoryStandardL2TokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimismMintableERC20FactoryStandardL2TokenCreated represents a StandardL2TokenCreated event raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryStandardL2TokenCreated struct {
	RemoteToken common.Address
	LocalToken  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStandardL2TokenCreated is a free log retrieval operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) FilterStandardL2TokenCreated(opts *bind.FilterOpts, remoteToken []common.Address, localToken []common.Address) (*OptimismMintableERC20FactoryStandardL2TokenCreatedIterator, error) {

	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _OptimismMintableERC20Factory.contract.FilterLogs(opts, "StandardL2TokenCreated", remoteTokenRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryStandardL2TokenCreatedIterator{contract: _OptimismMintableERC20Factory.contract, event: "StandardL2TokenCreated", logs: logs, sub: sub}, nil
}

// WatchStandardL2TokenCreated is a free log subscription operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) WatchStandardL2TokenCreated(opts *bind.WatchOpts, sink chan<- *OptimismMintableERC20FactoryStandardL2TokenCreated, remoteToken []common.Address, localToken []common.Address) (event.Subscription, error) {

	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _OptimismMintableERC20Factory.contract.WatchLogs(opts, "StandardL2TokenCreated", remoteTokenRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimismMintableERC20FactoryStandardL2TokenCreated)
				if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
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

// ParseStandardL2TokenCreated is a log parse operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) ParseStandardL2TokenCreated(log types.Log) (*OptimismMintableERC20FactoryStandardL2TokenCreated, error) {
	event := new(OptimismMintableERC20FactoryStandardL2TokenCreated)
	if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
