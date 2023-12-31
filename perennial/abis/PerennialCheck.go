// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package perennial

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

// PerennialCheckMetaData contains all meta data concerning the PerennialCheck contract.
var PerennialCheckMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"products\",\"type\":\"address[]\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PerennialCheckABI is the input ABI used to generate the binding from.
// Deprecated: Use PerennialCheckMetaData.ABI instead.
var PerennialCheckABI = PerennialCheckMetaData.ABI

// PerennialCheck is an auto generated Go binding around an Ethereum contract.
type PerennialCheck struct {
	PerennialCheckCaller     // Read-only binding to the contract
	PerennialCheckTransactor // Write-only binding to the contract
	PerennialCheckFilterer   // Log filterer for contract events
}

// PerennialCheckCaller is an auto generated read-only Go binding around an Ethereum contract.
type PerennialCheckCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerennialCheckTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PerennialCheckTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerennialCheckFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PerennialCheckFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerennialCheckSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PerennialCheckSession struct {
	Contract     *PerennialCheck   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PerennialCheckCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PerennialCheckCallerSession struct {
	Contract *PerennialCheckCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PerennialCheckTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PerennialCheckTransactorSession struct {
	Contract     *PerennialCheckTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PerennialCheckRaw is an auto generated low-level Go binding around an Ethereum contract.
type PerennialCheckRaw struct {
	Contract *PerennialCheck // Generic contract binding to access the raw methods on
}

// PerennialCheckCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PerennialCheckCallerRaw struct {
	Contract *PerennialCheckCaller // Generic read-only contract binding to access the raw methods on
}

// PerennialCheckTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PerennialCheckTransactorRaw struct {
	Contract *PerennialCheckTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerennialCheck creates a new instance of PerennialCheck, bound to a specific deployed contract.
func NewPerennialCheck(address common.Address, backend bind.ContractBackend) (*PerennialCheck, error) {
	contract, err := bindPerennialCheck(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PerennialCheck{PerennialCheckCaller: PerennialCheckCaller{contract: contract}, PerennialCheckTransactor: PerennialCheckTransactor{contract: contract}, PerennialCheckFilterer: PerennialCheckFilterer{contract: contract}}, nil
}

// NewPerennialCheckCaller creates a new read-only instance of PerennialCheck, bound to a specific deployed contract.
func NewPerennialCheckCaller(address common.Address, caller bind.ContractCaller) (*PerennialCheckCaller, error) {
	contract, err := bindPerennialCheck(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PerennialCheckCaller{contract: contract}, nil
}

// NewPerennialCheckTransactor creates a new write-only instance of PerennialCheck, bound to a specific deployed contract.
func NewPerennialCheckTransactor(address common.Address, transactor bind.ContractTransactor) (*PerennialCheckTransactor, error) {
	contract, err := bindPerennialCheck(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PerennialCheckTransactor{contract: contract}, nil
}

// NewPerennialCheckFilterer creates a new log filterer instance of PerennialCheck, bound to a specific deployed contract.
func NewPerennialCheckFilterer(address common.Address, filterer bind.ContractFilterer) (*PerennialCheckFilterer, error) {
	contract, err := bindPerennialCheck(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PerennialCheckFilterer{contract: contract}, nil
}

// bindPerennialCheck binds a generic wrapper to an already deployed contract.
func bindPerennialCheck(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PerennialCheckABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerennialCheck *PerennialCheckRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerennialCheck.Contract.PerennialCheckCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerennialCheck *PerennialCheckRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerennialCheck.Contract.PerennialCheckTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerennialCheck *PerennialCheckRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerennialCheck.Contract.PerennialCheckTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerennialCheck *PerennialCheckCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerennialCheck.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerennialCheck *PerennialCheckTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerennialCheck.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerennialCheck *PerennialCheckTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerennialCheck.Contract.contract.Transact(opts, method, params...)
}

// Check is a paid mutator transaction binding the contract method 0xcb15de5d.
//
// Solidity: function check(address[] accounts, address[] products) returns(uint256[])
func (_PerennialCheck *PerennialCheckTransactor) Check(opts *bind.TransactOpts, accounts []common.Address, products []common.Address) (*types.Transaction, error) {
	return _PerennialCheck.contract.Transact(opts, "check", accounts, products)
}

// Check is a paid mutator transaction binding the contract method 0xcb15de5d.
//
// Solidity: function check(address[] accounts, address[] products) returns(uint256[])
func (_PerennialCheck *PerennialCheckSession) Check(accounts []common.Address, products []common.Address) (*types.Transaction, error) {
	return _PerennialCheck.Contract.Check(&_PerennialCheck.TransactOpts, accounts, products)
}

// Check is a paid mutator transaction binding the contract method 0xcb15de5d.
//
// Solidity: function check(address[] accounts, address[] products) returns(uint256[])
func (_PerennialCheck *PerennialCheckTransactorSession) Check(accounts []common.Address, products []common.Address) (*types.Transaction, error) {
	return _PerennialCheck.Contract.Check(&_PerennialCheck.TransactOpts, accounts, products)
}
