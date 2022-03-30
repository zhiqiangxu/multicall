// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultiCallABI is the input ABI used to generate the binding from.
const MultiCallABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MultiCallBin is the compiled bytecode used for deploying new contracts.
var MultiCallBin = "0x608060405234801561001057600080fd5b506040516105c73803806105c783398101604081905261002f916102f6565b81518151811461005a5760405162461bcd60e51b8152600401610051906104a4565b60405180910390fd5b60608160405190808252806020026020018201604052801561009057816020015b606081526020019060019003908161007b5790505b50905060005b8281101561018b5760008582815181106100ac57fe5b6020026020010151905060608583815181106100c457fe5b6020026020010151905060006060836001600160a01b0316836040516100ea9190610498565b6000604051808303816000865af19150503d8060008114610127576040519150601f19603f3d011682016040523d82523d6000602084013e61012c565b606091505b509150915081610162576040518060200160405280600081525086868151811061015257fe5b602002602001018190525061017b565b8086868151811061016f57fe5b60200260200101819052505b5050600190920191506100969050565b50606043826040516020016101a19291906104b4565b6040516020818303038152906040529050805160208201f35b80516101c5816105af565b92915050565b600082601f8301126101dc57600080fd5b81516101ef6101ea82610502565b6104dc565b9150818183526020840193506020810190508385602084028201111561021457600080fd5b60005b83811015610240578161022a88826101ba565b8452506020928301929190910190600101610217565b5050505092915050565b600082601f83011261025b57600080fd5b81516102696101ea82610502565b81815260209384019390925082018360005b83811015610240578151860161029188826102a7565b845250602092830192919091019060010161027b565b600082601f8301126102b857600080fd5b81516102c66101ea82610522565b915080825260208301602083018583830111156102e257600080fd5b6102ed838284610575565b50505092915050565b6000806040838503121561030957600080fd5b82516001600160401b0381111561031f57600080fd5b61032b858286016101cb565b92505060208301516001600160401b0381111561034757600080fd5b6103538582860161024a565b9150509250929050565b6000610369838361040d565b9392505050565b600061037b8261054f565b6103858185610553565b93508360208202850161039785610549565b8060005b858110156103d157848403895281516103b4858261035d565b94506103bf83610549565b60209a909a019992505060010161039b565b5091979650505050505050565b60006103e98261054f565b6103f3818561055c565b9350610403818560208601610575565b9290920192915050565b60006104188261054f565b6104228185610553565b9350610432818560208601610575565b61043b816105a5565b9093019392505050565b6000610452602283610553565b7f4572726f723a204172726179206c656e6774687320646f206e6f74206d617463815261341760f11b602082015260400192915050565b61049281610572565b82525050565b600061036982846103de565b602080825281016101c581610445565b604081016104c28285610489565b81810360208301526104d48184610370565b949350505050565b6040518181016001600160401b03811182821017156104fa57600080fd5b604052919050565b60006001600160401b0382111561051857600080fd5b5060209081020190565b60006001600160401b0382111561053857600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b60006001600160a01b0382166101c5565b90565b60005b83811015610590578181015183820152602001610578565b8381111561059f576000848401525b50505050565b601f01601f191690565b6105b881610561565b81146105c357600080fd5b5056fe"

// DeployMultiCall deploys a new Ethereum contract, binding an instance of MultiCall to it.
func DeployMultiCall(auth *bind.TransactOpts, backend bind.ContractBackend, targets []common.Address, datas [][]byte) (common.Address, *types.Transaction, *MultiCall, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiCallABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiCallBin), backend, targets, datas)
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

