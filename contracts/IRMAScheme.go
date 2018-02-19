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
const IRMASchemeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_issuerId\",\"type\":\"string\"},{\"name\":\"_credId\",\"type\":\"string\"}],\"name\":\"getIssuerCredentialById\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_issuerId\",\"type\":\"string\"},{\"name\":\"_credentialId\",\"type\":\"string\"},{\"name\":\"_logoUrl\",\"type\":\"string\"},{\"name\":\"_issueSpec\",\"type\":\"bytes\"}],\"name\":\"addIssuerCredential\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_issuerId\",\"type\":\"string\"},{\"name\":\"_key\",\"type\":\"bytes\"}],\"name\":\"addIssuerPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"issuerIds\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfIssuers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_logoUrl\",\"type\":\"string\"},{\"name\":\"_metadata\",\"type\":\"bytes\"}],\"name\":\"addIssuer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getIssuerById\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_issuerId\",\"type\":\"string\"},{\"name\":\"_keyIndex\",\"type\":\"uint256\"}],\"name\":\"getIssuerPublicKeyById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_metadata\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// IRMASchemeBin is the compiled bytecode used for deploying new contracts.
const IRMASchemeBin = `0x606060405234156200001057600080fd5b604051620019a8380380620019a88339810160405280805182019190602001805160018054600160a060020a03191633600160a060020a031617905591909101905060008280516200006792916020019062000086565b5060028180516200007d92916020019062000086565b5050506200012b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000c957805160ff1916838001178555620000f9565b82800160010185558215620000f9579182015b82811115620000f9578251825591602001919060010190620000dc565b50620001079291506200010b565b5090565b6200012891905b8082111562000107576000815560010162000112565b90565b61186d806200013b6000396000f3006060604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313246ba281146100b357806316f3b1a51461029057806335ec27c4146103bb578063392f37e91461044e578063468e9703146104d857806369c05a99146104ee5780636cd4b82e146105135780638da5cb5b146105e8578063af640d0f14610617578063b53469701461062a578063c2c98077146107ec575b600080fd5b34156100be57600080fd5b61014660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506108bd95505050505050565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561018b578082015183820152602001610173565b50505050905090810190601f1680156101b85780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156101ee5780820151838201526020016101d6565b50505050905090810190601f16801561021b5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015610251578082015183820152602001610239565b50505050905090810190601f16801561027e5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b341561029b57600080fd5b6103a760046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650610bbb95505050505050565b604051901515815260200160405180910390f35b34156103c657600080fd5b6103a760046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650610d8a95505050505050565b341561045957600080fd5b610461610e8c565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561049d578082015183820152602001610485565b50505050905090810190601f1680156104ca5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156104e357600080fd5b610461600435610f2a565b34156104f957600080fd5b610501610fb2565b60405190815260200160405180910390f35b341561051e57600080fd5b6103a760046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650610fb995505050505050565b34156105f357600080fd5b6105fb611225565b604051600160a060020a03909116815260200160405180910390f35b341561062257600080fd5b610461611234565b341561063557600080fd5b61067b60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061129f95505050505050565b60405180806020018060200187600160a060020a0316600160a060020a031681526020018060200186815260200185815260200184810384528a818151815260200191508051906020019080838360005b838110156106e45780820151838201526020016106cc565b50505050905090810190601f1680156107115780820380516001836020036101000a031916815260200191505b50848103835289818151815260200191508051906020019080838360005b8381101561074757808201518382015260200161072f565b50505050905090810190601f1680156107745780820380516001836020036101000a031916815260200191505b50848103825287818151815260200191508051906020019080838360005b838110156107aa578082015183820152602001610792565b50505050905090810190601f1680156107d75780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b34156107f757600080fd5b61083f60046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650509335935061154d92505050565b60405182815260406020820181815290820183818151815260200191508051906020019080838360005b83811015610881578082015183820152602001610869565b50505050905090810190601f1680156108ae5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6108c56116ac565b6108cd6116ac565b6108d56116ac565b6000806003876040518082805190602001908083835b6020831061090a5780518252601f1990920191602091820191016108eb565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490925060ff16151561095157600080fd5b81600701866040518082805190602001908083835b602083106109855780518252601f199092019160209182019101610966565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490915060ff1615156109cc57600080fd5b806001018160020182600301828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a6d5780601f10610a4257610100808354040283529160200191610a6d565b820191906000526020600020905b815481529060010190602001808311610a5057829003601f168201915b50505050509250818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b095780601f10610ade57610100808354040283529160200191610b09565b820191906000526020600020905b815481529060010190602001808311610aec57829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ba55780601f10610b7a57610100808354040283529160200191610ba5565b820191906000526020600020905b815481529060010190602001808311610b8857829003601f168201915b5050505050905094509450945050509250925092565b6000806003866040518082805190602001908083835b60208310610bf05780518252601f199092019160209182019101610bd1565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490915060ff161515610c3b5760009150610d81565b600381015433600160a060020a03908116911614610c5c5760009150610d81565b6080604051908101604052806001151581526020018681526020018581526020018481525081600701866040518082805190602001908083835b60208310610cb55780518252601f199092019160209182019101610c96565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390208151815460ff1916901515178155602082015181600101908051610d119291602001906116be565b50604082015181600201908051610d2c9291602001906116be565b50606082015181600301908051610d479291602001906116be565b50505060088101805460018101610d5e838261173c565b6000928352602090922001868051610d7a9291602001906116be565b5050600191505b50949350505050565b6000806003846040518082805190602001908083835b60208310610dbf5780518252601f199092019160209182019101610da0565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490915060ff161515610e0a5760009150610e85565b600381015433600160a060020a03908116911614610e2b5760009150610e85565b6040805190810160409081526006830154808352602080840187905260009182526005850190522081518155602082015181600101908051610e719291602001906116be565b505050600681018054600190810190915591505b5092915050565b60028054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f225780601f10610ef757610100808354040283529160200191610f22565b820191906000526020600020905b815481529060010190602001808311610f0557829003601f168201915b505050505081565b6004805482908110610f3857fe5b90600052602060002090016000915090508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f225780601f10610ef757610100808354040283529160200191610f22565b6004545b90565b600083511515610fcb5750600061121e565b6003846040518082805190602001908083835b60208310610ffd5780518252601f199092019160209182019101610fde565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205460ff16156110425750600061121e565b60e06040519081016040528060011515815260200185815260200184815260200133600160a060020a0316815260200183815260200160008152602001600060405180591061108e5750595b9080825280602002602001820160405280156110c457816020015b6110b16116ac565b8152602001906001900390816110a95790505b5090526003856040518082805190602001908083835b602083106110f95780518252601f1990920191602091820191016110da565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390208151815460ff19169015151781556020820151816001019080516111559291602001906116be565b506040820151816002019080516111709291602001906116be565b50606082015160038201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556080820151816004019080516111bf9291602001906116be565b5060a0820151816006015560c0820151816008019080516111e4929160200190611765565b505060048054909150600181016111fb838261173c565b60009283526020909220018580516112179291602001906116be565b5050600190505b9392505050565b600154600160a060020a031681565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f225780601f10610ef757610100808354040283529160200191610f22565b6112a76116ac565b6112af6116ac565b60006112b96116ac565b60008060006003886040518082805190602001908083835b602083106112f05780518252601f1990920191602091820191016112d1565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490915060ff16156115435780600101816002018260030160009054906101000a9004600160a060020a03168360040184600601548560080180549050858054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113f75780601f106113cc576101008083540402835291602001916113f7565b820191906000526020600020905b8154815290600101906020018083116113da57829003601f168201915b50505050509550848054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114935780601f1061146857610100808354040283529160200191611493565b820191906000526020600020905b81548152906001019060200180831161147657829003601f168201915b50505050509450828054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561152f5780601f106115045761010080835404028352916020019161152f565b820191906000526020600020905b81548152906001019060200180831161151257829003601f168201915b505050505092509650965096509650965096505b5091939550919395565b60006115576116ac565b6000806003866040518082805190602001908083835b6020831061158c5780518252601f19909201916020918201910161156d565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805490925060ff1615156115d357600080fd5b600682015485106115e357600080fd5b8160050160008681526020019081526020016000209050806000015481600101808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156116985780601f1061166d57610100808354040283529160200191611698565b820191906000526020600020905b81548152906001019060200180831161167b57829003601f168201915b505050505090509350935050509250929050565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106116ff57805160ff191683800117855561172c565b8280016001018555821561172c579182015b8281111561172c578251825591602001919060010190611711565b506117389291506117bd565b5090565b815481835581811511611760576000838152602090206117609181019083016117d7565b505050565b8280548282559060005260206000209081019282156117b1579160200282015b828111156117b1578251829080516117a19291602001906116be565b5091602001919060010190611785565b506117389291506117d7565b610fb691905b8082111561173857600081556001016117c3565b610fb691905b808211156117385760006117f182826117fa565b506001016117dd565b50805460018160011615610100020316600290046000825580601f10611820575061183e565b601f01602090049060005260206000209081019061183e91906117bd565b505600a165627a7a723058200f97b08f7bb0c76f2f0bb448112cf8c61154f44c70600fa96a4b495c22476cbb0029`

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
// Solidity: function getIssuerById(_id string) constant returns(string, string, address, bytes, uint256, uint256)
func (_IRMAScheme *IRMASchemeCaller) GetIssuerById(opts *bind.CallOpts, _id string) (string, string, common.Address, []byte, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(common.Address)
		ret3 = new([]byte)
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
// Solidity: function getIssuerById(_id string) constant returns(string, string, address, bytes, uint256, uint256)
func (_IRMAScheme *IRMASchemeSession) GetIssuerById(_id string) (string, string, common.Address, []byte, *big.Int, *big.Int, error) {
	return _IRMAScheme.Contract.GetIssuerById(&_IRMAScheme.CallOpts, _id)
}

// GetIssuerById is a free data retrieval call binding the contract method 0xb5346970.
//
// Solidity: function getIssuerById(_id string) constant returns(string, string, address, bytes, uint256, uint256)
func (_IRMAScheme *IRMASchemeCallerSession) GetIssuerById(_id string) (string, string, common.Address, []byte, *big.Int, *big.Int, error) {
	return _IRMAScheme.Contract.GetIssuerById(&_IRMAScheme.CallOpts, _id)
}

// GetIssuerCredentialById is a free data retrieval call binding the contract method 0x13246ba2.
//
// Solidity: function getIssuerCredentialById(_issuerId string, _credId string) constant returns(string, string, bytes)
func (_IRMAScheme *IRMASchemeCaller) GetIssuerCredentialById(opts *bind.CallOpts, _issuerId string, _credId string) (string, string, []byte, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([]byte)
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
// Solidity: function getIssuerCredentialById(_issuerId string, _credId string) constant returns(string, string, bytes)
func (_IRMAScheme *IRMASchemeSession) GetIssuerCredentialById(_issuerId string, _credId string) (string, string, []byte, error) {
	return _IRMAScheme.Contract.GetIssuerCredentialById(&_IRMAScheme.CallOpts, _issuerId, _credId)
}

// GetIssuerCredentialById is a free data retrieval call binding the contract method 0x13246ba2.
//
// Solidity: function getIssuerCredentialById(_issuerId string, _credId string) constant returns(string, string, bytes)
func (_IRMAScheme *IRMASchemeCallerSession) GetIssuerCredentialById(_issuerId string, _credId string) (string, string, []byte, error) {
	return _IRMAScheme.Contract.GetIssuerCredentialById(&_IRMAScheme.CallOpts, _issuerId, _credId)
}

// GetIssuerPublicKeyById is a free data retrieval call binding the contract method 0xc2c98077.
//
// Solidity: function getIssuerPublicKeyById(_issuerId string, _keyIndex uint256) constant returns(uint256, bytes)
func (_IRMAScheme *IRMASchemeCaller) GetIssuerPublicKeyById(opts *bind.CallOpts, _issuerId string, _keyIndex *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
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
// Solidity: function getIssuerPublicKeyById(_issuerId string, _keyIndex uint256) constant returns(uint256, bytes)
func (_IRMAScheme *IRMASchemeSession) GetIssuerPublicKeyById(_issuerId string, _keyIndex *big.Int) (*big.Int, []byte, error) {
	return _IRMAScheme.Contract.GetIssuerPublicKeyById(&_IRMAScheme.CallOpts, _issuerId, _keyIndex)
}

// GetIssuerPublicKeyById is a free data retrieval call binding the contract method 0xc2c98077.
//
// Solidity: function getIssuerPublicKeyById(_issuerId string, _keyIndex uint256) constant returns(uint256, bytes)
func (_IRMAScheme *IRMASchemeCallerSession) GetIssuerPublicKeyById(_issuerId string, _keyIndex *big.Int) (*big.Int, []byte, error) {
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

// AddIssuer is a paid mutator transaction binding the contract method 0x6cd4b82e.
//
// Solidity: function addIssuer(_id string, _logoUrl string, _metadata bytes) returns(bool)
func (_IRMAScheme *IRMASchemeTransactor) AddIssuer(opts *bind.TransactOpts, _id string, _logoUrl string, _metadata []byte) (*types.Transaction, error) {
	return _IRMAScheme.contract.Transact(opts, "addIssuer", _id, _logoUrl, _metadata)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x6cd4b82e.
//
// Solidity: function addIssuer(_id string, _logoUrl string, _metadata bytes) returns(bool)
func (_IRMAScheme *IRMASchemeSession) AddIssuer(_id string, _logoUrl string, _metadata []byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuer(&_IRMAScheme.TransactOpts, _id, _logoUrl, _metadata)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x6cd4b82e.
//
// Solidity: function addIssuer(_id string, _logoUrl string, _metadata bytes) returns(bool)
func (_IRMAScheme *IRMASchemeTransactorSession) AddIssuer(_id string, _logoUrl string, _metadata []byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuer(&_IRMAScheme.TransactOpts, _id, _logoUrl, _metadata)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0x16f3b1a5.
//
// Solidity: function addIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec bytes) returns(bool)
func (_IRMAScheme *IRMASchemeTransactor) AddIssuerCredential(opts *bind.TransactOpts, _issuerId string, _credentialId string, _logoUrl string, _issueSpec []byte) (*types.Transaction, error) {
	return _IRMAScheme.contract.Transact(opts, "addIssuerCredential", _issuerId, _credentialId, _logoUrl, _issueSpec)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0x16f3b1a5.
//
// Solidity: function addIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec bytes) returns(bool)
func (_IRMAScheme *IRMASchemeSession) AddIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec []byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuerCredential(&_IRMAScheme.TransactOpts, _issuerId, _credentialId, _logoUrl, _issueSpec)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0x16f3b1a5.
//
// Solidity: function addIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec bytes) returns(bool)
func (_IRMAScheme *IRMASchemeTransactorSession) AddIssuerCredential(_issuerId string, _credentialId string, _logoUrl string, _issueSpec []byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuerCredential(&_IRMAScheme.TransactOpts, _issuerId, _credentialId, _logoUrl, _issueSpec)
}

// AddIssuerPublicKey is a paid mutator transaction binding the contract method 0x35ec27c4.
//
// Solidity: function addIssuerPublicKey(_issuerId string, _key bytes) returns(bool)
func (_IRMAScheme *IRMASchemeTransactor) AddIssuerPublicKey(opts *bind.TransactOpts, _issuerId string, _key []byte) (*types.Transaction, error) {
	return _IRMAScheme.contract.Transact(opts, "addIssuerPublicKey", _issuerId, _key)
}

// AddIssuerPublicKey is a paid mutator transaction binding the contract method 0x35ec27c4.
//
// Solidity: function addIssuerPublicKey(_issuerId string, _key bytes) returns(bool)
func (_IRMAScheme *IRMASchemeSession) AddIssuerPublicKey(_issuerId string, _key []byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuerPublicKey(&_IRMAScheme.TransactOpts, _issuerId, _key)
}

// AddIssuerPublicKey is a paid mutator transaction binding the contract method 0x35ec27c4.
//
// Solidity: function addIssuerPublicKey(_issuerId string, _key bytes) returns(bool)
func (_IRMAScheme *IRMASchemeTransactorSession) AddIssuerPublicKey(_issuerId string, _key []byte) (*types.Transaction, error) {
	return _IRMAScheme.Contract.AddIssuerPublicKey(&_IRMAScheme.TransactOpts, _issuerId, _key)
}
