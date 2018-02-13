// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IRMASchemeABI is the input ABI used to generate the binding from.
const IRMASchemeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_issuerId\",\"type\":\"string\"},{\"name\":\"_credId\",\"type\":\"string\"}],\"name\":\"getIssuerCredentialById\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes1[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"issuerIds\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfIssuers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_logoUrl\",\"type\":\"string\"},{\"name\":\"_metadata\",\"type\":\"bytes1[]\"}],\"name\":\"addIssuer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getIssuerById\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes1[]\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_issuerId\",\"type\":\"string\"},{\"name\":\"_key\",\"type\":\"bytes1[]\"}],\"name\":\"addIssuerPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_issuerId\",\"type\":\"string\"},{\"name\":\"_keyIndex\",\"type\":\"uint256\"}],\"name\":\"getIssuerPublicKeyById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes1[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_issuerId\",\"type\":\"string\"},{\"name\":\"_credentialId\",\"type\":\"string\"},{\"name\":\"_logoUrl\",\"type\":\"string\"},{\"name\":\"_issueSpec\",\"type\":\"bytes1[]\"}],\"name\":\"addIssuerCredential\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_metadata\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// IRMASchemeBin is the compiled bytecode used for deploying new contracts.
const IRMASchemeBin = `0x606060405234156200001057600080fd5b604051620019ac380380620019ac8339810160405280805182019190602001805160018054600160a060020a03191633600160a060020a031617905591909101905060008280516200006792916020019062000086565b5060028180516200007d92916020019062000086565b5050506200012b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000c957805160ff1916838001178555620000f9565b82800160010185558215620000f9579182015b82811115620000f9578251825591602001919060010190620000dc565b50620001079291506200010b565b5090565b6200012891905b8082111562000107576000815560010162000112565b90565b611871806200013b6000396000f3006060604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313246ba281146100b3578063392f37e91461026c578063468e9703146102f657806369c05a991461030c5780638da5cb5b1461033157806397a19be814610360578063af640d0f14610447578063b53469701461045a578063b6a9da9f146105f8578063c2c9807714610689578063dafd434214610736575b600080fd5b34156100be57600080fd5b61014660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061084b95505050505050565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561018b578082015183820152602001610173565b50505050905090810190601f1680156101b85780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156101ee5780820151838201526020016101d6565b50505050905090810190601f16801561021b5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019060200280838360005b8381101561025457808201518382015260200161023c565b50505050905001965050505050505060405180910390f35b341561027757600080fd5b61027f610b2f565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102bb5780820151838201526020016102a3565b50505050905090810190601f1680156102e85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561030157600080fd5b61027f600435610bcd565b341561031757600080fd5b61031f610c55565b60405190815260200160405180910390f35b341561033c57600080fd5b610344610c5c565b604051600160a060020a03909116815260200160405180910390f35b341561036b57600080fd5b61043360046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650610c6b95505050505050565b604051901515815260200160405180910390f35b341561045257600080fd5b61027f610ed7565b341561046557600080fd5b6104ab60046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610f4295505050505050565b60405180806020018060200187600160a060020a0316600160a060020a031681526020018060200186815260200185815260200184810384528a818151815260200191508051906020019080838360005b838110156105145780820151838201526020016104fc565b50505050905090810190601f1680156105415780820380516001836020036101000a031916815260200191505b50848103835289818151815260200191508051906020019080838360005b8381101561057757808201518382015260200161055f565b50505050905090810190601f1680156105a45780820380516001836020036101000a031916815260200191505b50848103825287818151815260200191508051906020019060200280838360005b838110156105dd5780820151838201526020016105c5565b50505050905001995050505050505050505060405180910390f35b341561060357600080fd5b61043360046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437509496506111d695505050505050565b341561069457600080fd5b6106dc60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050933593506112d892505050565b60405182815260406020820181815290820183818151815260200191508051906020019060200280838360005b83811015610721578082015183820152602001610709565b50505050905001935050505060405180910390f35b341561074157600080fd5b61043360046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284375094965061141d95505050505050565b6108536115ec565b61085b6115ec565b6108636115ec565b6000806003876040518082805190602001908083835b602083106108985780518252601f199092019160209182019101610879565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490925060ff1615156108df57600080fd5b81600701866040518082805190602001908083835b602083106109135780518252601f1990920191602091820191016108f4565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490915060ff16151561095a57600080fd5b806001018160020182600301828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109fb5780601f106109d0576101008083540402835291602001916109fb565b820191906000526020600020905b8154815290600101906020018083116109de57829003601f168201915b50505050509250818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a975780601f10610a6c57610100808354040283529160200191610a97565b820191906000526020600020905b815481529060010190602001808311610a7a57829003601f168201915b5050505050915080805480602002602001604051908101604052809291908181526020018280548015610b1957602002820191906000526020600020906000905b82829054906101000a900460f860020a02600160f860020a03191681526020019060010190602082600001049283019260010382029150808411610ad85790505b5050505050905094509450945050509250925092565b60028054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bc55780601f10610b9a57610100808354040283529160200191610bc5565b820191906000526020600020905b815481529060010190602001808311610ba857829003601f168201915b505050505081565b6004805482908110610bdb57fe5b90600052602060002090016000915090508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bc55780601f10610b9a57610100808354040283529160200191610bc5565b6004545b90565b600154600160a060020a031681565b600083511515610c7d57506000610ed0565b6003846040518082805190602001908083835b60208310610caf5780518252601f199092019160209182019101610c90565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205460ff1615610cf457506000610ed0565b60e06040519081016040528060011515815260200185815260200184815260200133600160a060020a03168152602001838152602001600081526020016000604051805910610d405750595b908082528060200260200182016040528015610d7657816020015b610d636115ec565b815260200190600190039081610d5b5790505b5090526003856040518082805190602001908083835b60208310610dab5780518252601f199092019160209182019101610d8c565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390208151815460ff1916901515178155602082015181600101908051610e079291602001906115fe565b50604082015181600201908051610e229291602001906115fe565b50606082015160038201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055608082015181600401908051610e7192916020019061167c565b5060a0820151816006015560c082015181600801908051610e96929160200190611722565b50506004805490915060018101610ead838261177a565b6000928352602090922001858051610ec99291602001906115fe565b5050600190505b9392505050565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bc55780601f10610b9a57610100808354040283529160200191610bc5565b610f4a6115ec565b610f526115ec565b6000610f5c6115ec565b60008060006003886040518082805190602001908083835b60208310610f935780518252601f199092019160209182019101610f74565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490915060ff16156111cc5780600101816002018260030160009054906101000a9004600160a060020a03168360040184600601548560080180549050858054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561109a5780601f1061106f5761010080835404028352916020019161109a565b820191906000526020600020905b81548152906001019060200180831161107d57829003601f168201915b50505050509550848054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111365780601f1061110b57610100808354040283529160200191611136565b820191906000526020600020905b81548152906001019060200180831161111957829003601f168201915b50505050509450828054806020026020016040519081016040528092919081815260200182805480156111b857602002820191906000526020600020906000905b82829054906101000a900460f860020a02600160f860020a031916815260200190600101906020826000010492830192600103820291508084116111775790505b505050505092509650965096509650965096505b5091939550919395565b6000806003846040518082805190602001908083835b6020831061120b5780518252601f1990920191602091820191016111ec565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490915060ff16151561125657600091506112d1565b600381015433600160a060020a0390811691161461127757600091506112d1565b60408051908101604090815260068301548083526020808401879052600091825260058501905220815181556020820151816001019080516112bd92916020019061167c565b505050600681018054600190810190915591505b5092915050565b60006112e26115ec565b6000806003866040518082805190602001908083835b602083106113175780518252601f1990920191602091820191016112f8565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490925060ff16151561135e57600080fd5b6006820154851061136e57600080fd5b81600501600086815260200190815260200160002090508060000154816001018080548060200260200160405190810160405280929190818152602001828054801561140957602002820191906000526020600020906000905b82829054906101000a900460f860020a02600160f860020a031916815260200190600101906020826000010492830192600103820291508084116113c85790505b505050505090509350935050509250929050565b6000806003866040518082805190602001908083835b602083106114525780518252601f199092019160209182019101611433565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490915060ff16151561149d57600091506115e3565b600381015433600160a060020a039081169116146114be57600091506115e3565b6080604051908101604052806001151581526020018681526020018581526020018481525081600701866040518082805190602001908083835b602083106115175780518252601f1990920191602091820191016114f8565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390208151815460ff19169015151781556020820151816001019080516115739291602001906115fe565b5060408201518160020190805161158e9291602001906115fe565b506060820151816003019080516115a992916020019061167c565b505050600881018054600181016115c0838261177a565b60009283526020909220018680516115dc9291602001906115fe565b5050600191505b50949350505050565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061163f57805160ff191683800117855561166c565b8280016001018555821561166c579182015b8281111561166c578251825591602001919060010190611651565b506116789291506117a3565b5090565b82805482825590600052602060002090601f016020900481019282156117165791602002820160005b838211156116e757835183826101000a81548160ff021916908360f860020a9004021790555092602001926001016020816000010492830192600103026116a5565b80156117145782816101000a81549060ff02191690556001016020816000010492830192600103026116e7565b505b506116789291506117bd565b82805482825590600052602060002090810192821561176e579160200282015b8281111561176e5782518290805161175e9291602001906115fe565b5091602001919060010190611742565b506116789291506117db565b81548183558181151161179e5760008381526020902061179e9181019083016117db565b505050565b610c5991905b8082111561167857600081556001016117a9565b610c5991905b8082111561167857805460ff191681556001016117c3565b610c5991905b808211156116785760006117f582826117fe565b506001016117e1565b50805460018160011615610100020316600290046000825580601f106118245750611842565b601f01602090049060005260206000209081019061184291906117a3565b505600a165627a7a72305820116a1b7673f6ed6c60f12d3d97e7c7286a9119c11cca8f725dc45ebaa9607fc20029`

// DeployIRMAScheme deploys a new Ethereum contract, binding an instance of IRMAScheme to it.
func DeployIRMAScheme(auth *bind.TransactOpts, backend bind.ContractBackend, _id string, _metadata []byte) (common.Address, *types.Transaction, *IRMAScheme, error) {
	parsed, err := abi.JSON(strings.NewReader(IRMASchemeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IRMASchemeBin), backend, _id, _metadata)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IRMAScheme{IRMASchemeCaller: IRMASchemeCaller{contract: contract}, IRMASchemeTransactor: IRMASchemeTransactor{contract: contract}}, nil
}

// IRMAScheme is an auto generated Go binding around an Ethereum contract.
type IRMAScheme struct {
	IRMASchemeCaller     // Read-only binding to the contract
	IRMASchemeTransactor // Write-only binding to the contract
}

// IRMASchemeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRMASchemeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRMASchemeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRMASchemeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRMASchemeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRMASchemeSession struct {
	Contract     *IRMAScheme       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRMASchemeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRMASchemeCallerSession struct {
	Contract *IRMASchemeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IRMASchemeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRMASchemeTransactorSession struct {
	Contract     *IRMASchemeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IRMASchemeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRMASchemeRaw struct {
	Contract *IRMAScheme // Generic contract binding to access the raw methods on
}

// IRMASchemeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRMASchemeCallerRaw struct {
	Contract *IRMASchemeCaller // Generic read-only contract binding to access the raw methods on
}

// IRMASchemeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRMASchemeTransactorRaw struct {
	Contract *IRMASchemeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRMAScheme creates a new instance of IRMAScheme, bound to a specific deployed contract.
func NewIRMAScheme(address common.Address, backend bind.ContractBackend) (*IRMAScheme, error) {
	contract, err := bindIRMAScheme(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRMAScheme{IRMASchemeCaller: IRMASchemeCaller{contract: contract}, IRMASchemeTransactor: IRMASchemeTransactor{contract: contract}}, nil
}

// NewIRMASchemeCaller creates a new read-only instance of IRMAScheme, bound to a specific deployed contract.
func NewIRMASchemeCaller(address common.Address, caller bind.ContractCaller) (*IRMASchemeCaller, error) {
	contract, err := bindIRMAScheme(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &IRMASchemeCaller{contract: contract}, nil
}

// NewIRMASchemeTransactor creates a new write-only instance of IRMAScheme, bound to a specific deployed contract.
func NewIRMASchemeTransactor(address common.Address, transactor bind.ContractTransactor) (*IRMASchemeTransactor, error) {
	contract, err := bindIRMAScheme(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &IRMASchemeTransactor{contract: contract}, nil
}

// bindIRMAScheme binds a generic wrapper to an already deployed contract.
func bindIRMAScheme(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRMASchemeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRMAScheme *IRMASchemeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRMAScheme.Contract.IRMASchemeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRMAScheme *IRMASchemeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRMAScheme.Contract.IRMASchemeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRMAScheme *IRMASchemeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRMAScheme.Contract.IRMASchemeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRMAScheme *IRMASchemeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRMAScheme.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRMAScheme *IRMASchemeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRMAScheme.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRMAScheme *IRMASchemeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRMAScheme.Contract.contract.Transact(opts, method, params...)
}

// GetIssuerById is a free data retrieval call binding the contract method 0xb5346970.
//
// Solidity: function getIssuerById(_id string) constant returns(string, string, address, bytes1[], uint256, uint256)
func (_IRMAScheme *IRMASchemeCaller) GetIssuerById(opts *bind.CallOpts, _id string) (string, string, common.Address, [][1]byte, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(common.Address)
		ret3 = new([][1]byte)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _IRMAScheme.contract.Call(opts, out, "getIssuerById", _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetIssuerById is a free data retrieval call binding the contract method 0xb5346970.
//
// Solidity: function getIssuerById(_id string) constant returns(string, string, address, bytes1[], uint256, uint256)
func (_IRMAScheme *IRMASchemeSession) GetIssuerById(_id string) (string, string, common.Address, [][1]byte, *big.Int, *big.Int, error) {
	return _IRMAScheme.Contract.GetIssuerById(&_IRMAScheme.CallOpts, _id)
}

// GetIssuerById is a free data retrieval call binding the contract method 0xb5346970.
//
// Solidity: function getIssuerById(_id string) constant returns(string, string, address, bytes1[], uint256, uint256)
func (_IRMAScheme *IRMASchemeCallerSession) GetIssuerById(_id string) (string, string, common.Address, [][1]byte, *big.Int, *big.Int, error) {
	return _IRMAScheme.Contract.GetIssuerById(&_IRMAScheme.CallOpts, _id)
}

// GetIssuerCredentialById is a free data retrieval call binding the contract method 0x13246ba2.
//
// Solidity: function getIssuerCredentialById(_issuerId string, _credId string) constant returns(string, string, bytes1[])
func (_IRMAScheme *IRMASchemeCaller) GetIssuerCredentialById(opts *bind.CallOpts, _issuerId string, _credId string) (string, string, [][1]byte, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([][1]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _IRMAScheme.contract.Call(opts, out, "getIssuerCredentialById", _issuerId, _credId)
	return *ret0, *ret1, *ret2, err
}

// GetIssuerCredentialById is a free data retrieval call binding the contract method 0x13246ba2.
//
// Solidity: function getIssuerCredentialById(_issuerId string, _credId string) constant returns(string, string, bytes1[])
func (_IRMAScheme *IRMASchemeSession) GetIssuerCredentialById(_issuerId string, _credId string) (string, string, [][1]byte, error) {
	return _IRMAScheme.Contract.GetIssuerCredentialById(&_IRMAScheme.CallOpts, _issuerId, _credId)
}

// GetIssuerCredentialById is a free data retrieval call binding the contract method 0x13246ba2.
//
// Solidity: function getIssuerCredentialById(_issuerId string, _credId string) constant returns(string, string, bytes1[])
func (_IRMAScheme *IRMASchemeCallerSession) GetIssuerCredentialById(_issuerId string, _credId string) (string, string, [][1]byte, error) {
	return _IRMAScheme.Contract.GetIssuerCredentialById(&_IRMAScheme.CallOpts, _issuerId, _credId)
}

// GetIssuerPublicKeyById is a free data retrieval call binding the contract method 0xc2c98077.
//
// Solidity: function getIssuerPublicKeyById(_issuerId string, _keyIndex uint256) constant returns(uint256, bytes1[])
func (_IRMAScheme *IRMASchemeCaller) GetIssuerPublicKeyById(opts *bind.CallOpts, _issuerId string, _keyIndex *big.Int) (*big.Int, [][1]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([][1]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IRMAScheme.contract.Call(opts, out, "getIssuerPublicKeyById", _issuerId, _keyIndex)
	return *ret0, *ret1, err
}

// GetIssuerPublicKeyById is a free data retrieval call binding the contract method 0xc2c98077.
//
// Solidity: function getIssuerPublicKeyById(_issuerId string, _keyIndex uint256) constant returns(uint256, bytes1[])
func (_IRMAScheme *IRMASchemeSession) GetIssuerPublicKeyById(_issuerId string, _keyIndex *big.Int) (*big.Int, [][1]byte, error) {
	return _IRMAScheme.Contract.GetIssuerPublicKeyById(&_IRMAScheme.CallOpts, _issuerId, _keyIndex)
}

// GetIssuerPublicKeyById is a free data retrieval call binding the contract method 0xc2c98077.
//
// Solidity: function getIssuerPublicKeyById(_issuerId string, _keyIndex uint256) constant returns(uint256, bytes1[])
func (_IRMAScheme *IRMASchemeCallerSession) GetIssuerPublicKeyById(_issuerId string, _keyIndex *big.Int) (*big.Int, [][1]byte, error) {
	return _IRMAScheme.Contract.GetIssuerPublicKeyById(&_IRMAScheme.CallOpts, _issuerId, _keyIndex)
}

// GetNumberOfIssuers is a free data retrieval call binding the contract method 0x69c05a99.
//
// Solidity: function getNumberOfIssuers() constant returns(uint256)
func (_IRMAScheme *IRMASchemeCaller) GetNumberOfIssuers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRMAScheme.contract.Call(opts, out, "getNumberOfIssuers")
	return *ret0, err
}

// GetNumberOfIssuers is a free data retrieval call binding the contract method 0x69c05a99.
//
// Solidity: function getNumberOfIssuers() constant returns(uint256)
func (_IRMAScheme *IRMASchemeSession) GetNumberOfIssuers() (*big.Int, error) {
	return _IRMAScheme.Contract.GetNumberOfIssuers(&_IRMAScheme.CallOpts)
}

// GetNumberOfIssuers is a free data retrieval call binding the contract method 0x69c05a99.
//
// Solidity: function getNumberOfIssuers() constant returns(uint256)
func (_IRMAScheme *IRMASchemeCallerSession) GetNumberOfIssuers() (*big.Int, error) {
	return _IRMAScheme.Contract.GetNumberOfIssuers(&_IRMAScheme.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(string)
func (_IRMAScheme *IRMASchemeCaller) Id(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IRMAScheme.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(string)
func (_IRMAScheme *IRMASchemeSession) Id() (string, error) {
	return _IRMAScheme.Contract.Id(&_IRMAScheme.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(string)
func (_IRMAScheme *IRMASchemeCallerSession) Id() (string, error) {
	return _IRMAScheme.Contract.Id(&_IRMAScheme.CallOpts)
}

// IssuerIds is a free data retrieval call binding the contract method 0x468e9703.
//
// Solidity: function issuerIds( uint256) constant returns(string)
func (_IRMAScheme *IRMASchemeCaller) IssuerIds(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IRMAScheme.contract.Call(opts, out, "issuerIds", arg0)
	return *ret0, err
}

// IssuerIds is a free data retrieval call binding the contract method 0x468e9703.
//
// Solidity: function issuerIds( uint256) constant returns(string)
func (_IRMAScheme *IRMASchemeSession) IssuerIds(arg0 *big.Int) (string, error) {
	return _IRMAScheme.Contract.IssuerIds(&_IRMAScheme.CallOpts, arg0)
}

// IssuerIds is a free data retrieval call binding the contract method 0x468e9703.
//
// Solidity: function issuerIds( uint256) constant returns(string)
func (_IRMAScheme *IRMASchemeCallerSession) IssuerIds(arg0 *big.Int) (string, error) {
	return _IRMAScheme.Contract.IssuerIds(&_IRMAScheme.CallOpts, arg0)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() constant returns(bytes)
func (_IRMAScheme *IRMASchemeCaller) Metadata(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IRMAScheme.contract.Call(opts, out, "metadata")
	return *ret0, err
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() constant returns(bytes)
func (_IRMAScheme *IRMASchemeSession) Metadata() ([]byte, error) {
	return _IRMAScheme.Contract.Metadata(&_IRMAScheme.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() constant returns(bytes)
func (_IRMAScheme *IRMASchemeCallerSession) Metadata() ([]byte, error) {
	return _IRMAScheme.Contract.Metadata(&_IRMAScheme.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IRMAScheme *IRMASchemeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IRMAScheme.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IRMAScheme *IRMASchemeSession) Owner() (common.Address, error) {
	return _IRMAScheme.Contract.Owner(&_IRMAScheme.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IRMAScheme *IRMASchemeCallerSession) Owner() (common.Address, error) {
	return _IRMAScheme.Contract.Owner(&_IRMAScheme.CallOpts)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x97a19be8.
//
// Solidity: function addIssuer(_id string, _logoUrl string, _metadata bytes1[]) returns(bool)
func (_IRMAScheme *IRMASchemeTransactor) AddIssuer(opts *bind.TransactOpts, _id string, _logoUrl string, _metadata [][1]byte) (*types.Transaction, error) {
	return _IRMAScheme.contract.Transact(opts, "addIssuer", _id, _logoUrl, _metadata)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x97a19be8.
//
// Solidity: function addIssuer(_id string, _logoUrl string, _metadata bytes1[]) returns(bool)
func (_IRMAScheme *IRMASchemeSession) AddIssuer(_id string, _logoUrl string, _metadata [][1]byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuer(&_IRMAScheme.TransactOpts, _id, _logoUrl, _metadata)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x97a19be8.
//
// Solidity: function addIssuer(_id string, _logoUrl string, _metadata bytes1[]) returns(bool)
func (_IRMAScheme *IRMASchemeTransactorSession) AddIssuer(_id string, _logoUrl string, _metadata [][1]byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuer(&_IRMAScheme.TransactOpts, _id, _logoUrl, _metadata)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0xdafd4342.
//
// Solidity: function addIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec bytes1[]) returns(bool)
func (_IRMAScheme *IRMASchemeTransactor) AddIssuerCredential(opts *bind.TransactOpts, _issuerId string, _credentialId string, _logoUrl string, _issueSpec [][1]byte) (*types.Transaction, error) {
	return _IRMAScheme.contract.Transact(opts, "addIssuerCredential", _issuerId, _credentialId, _logoUrl, _issueSpec)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0xdafd4342.
//
// Solidity: function addIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec bytes1[]) returns(bool)
func (_IRMAScheme *IRMASchemeSession) AddIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec [][1]byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuerCredential(&_IRMAScheme.TransactOpts, _issuerId, _credentialId, _logoUrl, _issueSpec)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0xdafd4342.
//
// Solidity: function addIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec bytes1[]) returns(bool)
func (_IRMAScheme *IRMASchemeTransactorSession) AddIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec [][1]byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuerCredential(&_IRMAScheme.TransactOpts, _issuerId, _credentialId, _logoUrl, _issueSpec)
}

// AddIssuerPublicKey is a paid mutator transaction binding the contract method 0xb6a9da9f.
//
// Solidity: function addIssuerPublicKey(_issuerId string, _key bytes1[]) returns(bool)
func (_IRMAScheme *IRMASchemeTransactor) AddIssuerPublicKey(opts *bind.TransactOpts, _issuerId string, _key [][1]byte) (*types.Transaction, error) {
	return _IRMAScheme.contract.Transact(opts, "addIssuerPublicKey", _issuerId, _key)
}

// AddIssuerPublicKey is a paid mutator transaction binding the contract method 0xb6a9da9f.
//
// Solidity: function addIssuerPublicKey(_issuerId string, _key bytes1[]) returns(bool)
func (_IRMAScheme *IRMASchemeSession) AddIssuerPublicKey(_issuerId string, _key [][1]byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuerPublicKey(&_IRMAScheme.TransactOpts, _issuerId, _key)
}

// AddIssuerPublicKey is a paid mutator transaction binding the contract method 0xb6a9da9f.
//
// Solidity: function addIssuerPublicKey(_issuerId string, _key bytes1[]) returns(bool)
func (_IRMAScheme *IRMASchemeTransactorSession) AddIssuerPublicKey(_issuerId string, _key [][1]byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuerPublicKey(&_IRMAScheme.TransactOpts, _issuerId, _key)
}
