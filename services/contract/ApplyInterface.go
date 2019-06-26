// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApplyABI is the input ABI used to generate the binding from.
const ApplyABI = "[{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"finger\",\"type\":\"string\"}],\"name\":\"apply\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"finger\",\"type\":\"string\"}],\"name\":\"sign\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApplicants\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUnSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Apply is an auto generated Go binding around an Ethereum contract.
type Apply struct {
	ApplyCaller     // Read-only binding to the contract
	ApplyTransactor // Write-only binding to the contract
	ApplyFilterer   // Log filterer for contract events
}

// ApplyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApplyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApplyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApplyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApplyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApplyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApplySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApplySession struct {
	Contract     *Apply            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApplyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApplyCallerSession struct {
	Contract *ApplyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApplyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApplyTransactorSession struct {
	Contract     *ApplyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApplyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApplyRaw struct {
	Contract *Apply // Generic contract binding to access the raw methods on
}

// ApplyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApplyCallerRaw struct {
	Contract *ApplyCaller // Generic read-only contract binding to access the raw methods on
}

// ApplyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApplyTransactorRaw struct {
	Contract *ApplyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApply creates a new instance of Apply, bound to a specific deployed contract.
func NewApply(address common.Address, backend bind.ContractBackend) (*Apply, error) {
	contract, err := bindApply(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Apply{ApplyCaller: ApplyCaller{contract: contract}, ApplyTransactor: ApplyTransactor{contract: contract}, ApplyFilterer: ApplyFilterer{contract: contract}}, nil
}

// NewApplyCaller creates a new read-only instance of Apply, bound to a specific deployed contract.
func NewApplyCaller(address common.Address, caller bind.ContractCaller) (*ApplyCaller, error) {
	contract, err := bindApply(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApplyCaller{contract: contract}, nil
}

// NewApplyTransactor creates a new write-only instance of Apply, bound to a specific deployed contract.
func NewApplyTransactor(address common.Address, transactor bind.ContractTransactor) (*ApplyTransactor, error) {
	contract, err := bindApply(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApplyTransactor{contract: contract}, nil
}

// NewApplyFilterer creates a new log filterer instance of Apply, bound to a specific deployed contract.
func NewApplyFilterer(address common.Address, filterer bind.ContractFilterer) (*ApplyFilterer, error) {
	contract, err := bindApply(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApplyFilterer{contract: contract}, nil
}

// bindApply binds a generic wrapper to an already deployed contract.
func bindApply(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApplyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apply *ApplyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Apply.Contract.ApplyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apply *ApplyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apply.Contract.ApplyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apply *ApplyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apply.Contract.ApplyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apply *ApplyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Apply.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apply *ApplyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apply.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apply *ApplyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apply.Contract.contract.Transact(opts, method, params...)
}

// GetApplicants is a free data retrieval call binding the contract method 0x455772d2.
//
// Solidity: function getApplicants() constant returns(address[])
func (_Apply *ApplyCaller) GetApplicants(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Apply.contract.Call(opts, out, "getApplicants")
	return *ret0, err
}

// GetApplicants is a free data retrieval call binding the contract method 0x455772d2.
//
// Solidity: function getApplicants() constant returns(address[])
func (_Apply *ApplySession) GetApplicants() ([]common.Address, error) {
	return _Apply.Contract.GetApplicants(&_Apply.CallOpts)
}

// GetApplicants is a free data retrieval call binding the contract method 0x455772d2.
//
// Solidity: function getApplicants() constant returns(address[])
func (_Apply *ApplyCallerSession) GetApplicants() ([]common.Address, error) {
	return _Apply.Contract.GetApplicants(&_Apply.CallOpts)
}

// GetSigned is a free data retrieval call binding the contract method 0x56be3570.
//
// Solidity: function getSigned() constant returns(address[])
func (_Apply *ApplyCaller) GetSigned(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Apply.contract.Call(opts, out, "getSigned")
	return *ret0, err
}

// GetSigned is a free data retrieval call binding the contract method 0x56be3570.
//
// Solidity: function getSigned() constant returns(address[])
func (_Apply *ApplySession) GetSigned() ([]common.Address, error) {
	return _Apply.Contract.GetSigned(&_Apply.CallOpts)
}

// GetSigned is a free data retrieval call binding the contract method 0x56be3570.
//
// Solidity: function getSigned() constant returns(address[])
func (_Apply *ApplyCallerSession) GetSigned() ([]common.Address, error) {
	return _Apply.Contract.GetSigned(&_Apply.CallOpts)
}

// GetUnSigned is a free data retrieval call binding the contract method 0x79106e58.
//
// Solidity: function getUnSigned() constant returns(address[])
func (_Apply *ApplyCaller) GetUnSigned(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Apply.contract.Call(opts, out, "getUnSigned")
	return *ret0, err
}

// GetUnSigned is a free data retrieval call binding the contract method 0x79106e58.
//
// Solidity: function getUnSigned() constant returns(address[])
func (_Apply *ApplySession) GetUnSigned() ([]common.Address, error) {
	return _Apply.Contract.GetUnSigned(&_Apply.CallOpts)
}

// GetUnSigned is a free data retrieval call binding the contract method 0x79106e58.
//
// Solidity: function getUnSigned() constant returns(address[])
func (_Apply *ApplyCallerSession) GetUnSigned() ([]common.Address, error) {
	return _Apply.Contract.GetUnSigned(&_Apply.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0x897a142a.
//
// Solidity: function apply(address addr, string finger) returns(bool)
func (_Apply *ApplyTransactor) Apply(opts *bind.TransactOpts, addr common.Address, finger string) (*types.Transaction, error) {
	return _Apply.contract.Transact(opts, "apply", addr, finger)
}

// Apply is a paid mutator transaction binding the contract method 0x897a142a.
//
// Solidity: function apply(address addr, string finger) returns(bool)
func (_Apply *ApplySession) Apply(addr common.Address, finger string) (*types.Transaction, error) {
	return _Apply.Contract.Apply(&_Apply.TransactOpts, addr, finger)
}

// Apply is a paid mutator transaction binding the contract method 0x897a142a.
//
// Solidity: function apply(address addr, string finger) returns(bool)
func (_Apply *ApplyTransactorSession) Apply(addr common.Address, finger string) (*types.Transaction, error) {
	return _Apply.Contract.Apply(&_Apply.TransactOpts, addr, finger)
}

// Sign is a paid mutator transaction binding the contract method 0xad2a9fcf.
//
// Solidity: function sign(address addr, string finger) returns(bool)
func (_Apply *ApplyTransactor) Sign(opts *bind.TransactOpts, addr common.Address, finger string) (*types.Transaction, error) {
	return _Apply.contract.Transact(opts, "sign", addr, finger)
}

// Sign is a paid mutator transaction binding the contract method 0xad2a9fcf.
//
// Solidity: function sign(address addr, string finger) returns(bool)
func (_Apply *ApplySession) Sign(addr common.Address, finger string) (*types.Transaction, error) {
	return _Apply.Contract.Sign(&_Apply.TransactOpts, addr, finger)
}

// Sign is a paid mutator transaction binding the contract method 0xad2a9fcf.
//
// Solidity: function sign(address addr, string finger) returns(bool)
func (_Apply *ApplyTransactorSession) Sign(addr common.Address, finger string) (*types.Transaction, error) {
	return _Apply.Contract.Sign(&_Apply.TransactOpts, addr, finger)
}
