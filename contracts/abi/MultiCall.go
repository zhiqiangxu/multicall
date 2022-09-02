// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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
)

// MultiCallMetaData contains all meta data concerning the MultiCall contract.
var MultiCallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161058638038061058683398101604081905261002f916103d9565b8151815181146100915760405162461bcd60e51b815260206004820152602260248201527f4572726f723a204172726179206c656e6774687320646f206e6f74206d617463604482015261341760f11b60648201526084015b60405180910390fd5b6000816001600160401b038111156100ab576100ab610258565b6040519080825280602002602001820160405280156100de57816020015b60608152602001906001900390816100c95790505b50905060005b82811015610229576000858281518110610100576101006104aa565b60200260200101519050600085838151811061011e5761011e6104aa565b60200260200101519050600080836001600160a01b03168360405161014391906104c0565b6000604051808303816000865af19150503d8060008114610180576040519150601f19603f3d011682016040523d82523d6000602084013e610185565b606091505b5091509150816101f45780516000036101eb5760405162461bcd60e51b815260206004820152602260248201527f4572726f723a2063616c6c206661696c65642077697468206e6f20726561736f604482015261371760f11b6064820152608401610088565b80518082602001fd5b80868681518110610207576102076104aa565b6020026020010181905250505050508080610221906104dc565b9150506100e4565b506000438260405160200161023f929190610503565b6040516020818303038152906040529050805160208201f35b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561029657610296610258565b604052919050565b60006001600160401b038211156102b7576102b7610258565b5060051b60200190565b60005b838110156102dc5781810151838201526020016102c4565b838111156102eb576000848401525b50505050565b6000601f838184011261030357600080fd5b825160206103186103138361029e565b61026e565b82815260059290921b8501810191818101908784111561033757600080fd5b8287015b848110156103cd5780516001600160401b038082111561035b5760008081fd5b818a0191508a603f8301126103705760008081fd5b8582015160408282111561038657610386610258565b610397828b01601f1916890161026e565b92508183528c818386010111156103ae5760008081fd5b6103bd828985018387016102c1565b505084525091830191830161033b565b50979650505050505050565b600080604083850312156103ec57600080fd5b82516001600160401b038082111561040357600080fd5b818501915085601f83011261041757600080fd5b815160206104276103138361029e565b82815260059290921b8401810191818101908984111561044657600080fd5b948201945b8386101561047a5785516001600160a01b038116811461046b5760008081fd5b8252948201949082019061044b565b9188015191965090935050508082111561049357600080fd5b506104a0858286016102f1565b9150509250929050565b634e487b7160e01b600052603260045260246000fd5b600082516104d28184602087016102c1565b9190910192915050565b6000600182016104fc57634e487b7160e01b600052601160045260246000fd5b5060010190565b600060408201848352602060408185015281855180845260608601915060608160051b870101935082870160005b8281101561057757878603605f190184528151805180885261055881888a018985016102c1565b601f01601f191696909601850195509284019290840190600101610531565b50939897505050505050505056fe",
}

// MultiCallABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiCallMetaData.ABI instead.
var MultiCallABI = MultiCallMetaData.ABI

// MultiCallBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiCallMetaData.Bin instead.
var MultiCallBin = MultiCallMetaData.Bin

// DeployMultiCall deploys a new Ethereum contract, binding an instance of MultiCall to it.
func DeployMultiCall(auth *bind.TransactOpts, backend bind.ContractBackend, targets []common.Address, datas [][]byte) (common.Address, *types.Transaction, *MultiCall, error) {
	parsed, err := MultiCallMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiCallBin), backend, targets, datas)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiCall{MultiCallCaller: MultiCallCaller{contract: contract}, MultiCallTransactor: MultiCallTransactor{contract: contract}, MultiCallFilterer: MultiCallFilterer{contract: contract}}, nil
}

// MultiCall is an auto generated Go binding around an Ethereum contract.
type MultiCall struct {
	MultiCallCaller     // Read-only binding to the contract
	MultiCallTransactor // Write-only binding to the contract
	MultiCallFilterer   // Log filterer for contract events
}

// MultiCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiCallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiCallSession struct {
	Contract     *MultiCall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiCallCallerSession struct {
	Contract *MultiCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultiCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiCallTransactorSession struct {
	Contract     *MultiCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultiCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiCallRaw struct {
	Contract *MultiCall // Generic contract binding to access the raw methods on
}

// MultiCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiCallCallerRaw struct {
	Contract *MultiCallCaller // Generic read-only contract binding to access the raw methods on
}

// MultiCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiCallTransactorRaw struct {
	Contract *MultiCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiCall creates a new instance of MultiCall, bound to a specific deployed contract.
func NewMultiCall(address common.Address, backend bind.ContractBackend) (*MultiCall, error) {
	contract, err := bindMultiCall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiCall{MultiCallCaller: MultiCallCaller{contract: contract}, MultiCallTransactor: MultiCallTransactor{contract: contract}, MultiCallFilterer: MultiCallFilterer{contract: contract}}, nil
}

// NewMultiCallCaller creates a new read-only instance of MultiCall, bound to a specific deployed contract.
func NewMultiCallCaller(address common.Address, caller bind.ContractCaller) (*MultiCallCaller, error) {
	contract, err := bindMultiCall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallCaller{contract: contract}, nil
}

// NewMultiCallTransactor creates a new write-only instance of MultiCall, bound to a specific deployed contract.
func NewMultiCallTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiCallTransactor, error) {
	contract, err := bindMultiCall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallTransactor{contract: contract}, nil
}

// NewMultiCallFilterer creates a new log filterer instance of MultiCall, bound to a specific deployed contract.
func NewMultiCallFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiCallFilterer, error) {
	contract, err := bindMultiCall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiCallFilterer{contract: contract}, nil
}

// bindMultiCall binds a generic wrapper to an already deployed contract.
func bindMultiCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiCallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCall *MultiCallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCall.Contract.MultiCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCall *MultiCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCall.Contract.MultiCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCall *MultiCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCall.Contract.MultiCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCall *MultiCallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCall *MultiCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCall *MultiCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCall.Contract.contract.Transact(opts, method, params...)
}

