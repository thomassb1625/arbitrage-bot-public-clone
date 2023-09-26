// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// CheckTransferTaxMetaData contains all meta data concerning the CheckTransferTax contract.
var CheckTransferTaxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"swapWithoutFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// CheckTransferTaxABI is the input ABI used to generate the binding from.
// Deprecated: Use CheckTransferTaxMetaData.ABI instead.
var CheckTransferTaxABI = CheckTransferTaxMetaData.ABI

// CheckTransferTax is an auto generated Go binding around an Ethereum contract.
type CheckTransferTax struct {
	CheckTransferTaxCaller     // Read-only binding to the contract
	CheckTransferTaxTransactor // Write-only binding to the contract
	CheckTransferTaxFilterer   // Log filterer for contract events
}

// CheckTransferTaxCaller is an auto generated read-only Go binding around an Ethereum contract.
type CheckTransferTaxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckTransferTaxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CheckTransferTaxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckTransferTaxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CheckTransferTaxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckTransferTaxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CheckTransferTaxSession struct {
	Contract     *CheckTransferTax // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CheckTransferTaxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CheckTransferTaxCallerSession struct {
	Contract *CheckTransferTaxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CheckTransferTaxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CheckTransferTaxTransactorSession struct {
	Contract     *CheckTransferTaxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CheckTransferTaxRaw is an auto generated low-level Go binding around an Ethereum contract.
type CheckTransferTaxRaw struct {
	Contract *CheckTransferTax // Generic contract binding to access the raw methods on
}

// CheckTransferTaxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CheckTransferTaxCallerRaw struct {
	Contract *CheckTransferTaxCaller // Generic read-only contract binding to access the raw methods on
}

// CheckTransferTaxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CheckTransferTaxTransactorRaw struct {
	Contract *CheckTransferTaxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCheckTransferTax creates a new instance of CheckTransferTax, bound to a specific deployed contract.
func NewCheckTransferTax(address common.Address, backend bind.ContractBackend) (*CheckTransferTax, error) {
	contract, err := bindCheckTransferTax(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CheckTransferTax{CheckTransferTaxCaller: CheckTransferTaxCaller{contract: contract}, CheckTransferTaxTransactor: CheckTransferTaxTransactor{contract: contract}, CheckTransferTaxFilterer: CheckTransferTaxFilterer{contract: contract}}, nil
}

// NewCheckTransferTaxCaller creates a new read-only instance of CheckTransferTax, bound to a specific deployed contract.
func NewCheckTransferTaxCaller(address common.Address, caller bind.ContractCaller) (*CheckTransferTaxCaller, error) {
	contract, err := bindCheckTransferTax(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CheckTransferTaxCaller{contract: contract}, nil
}

// NewCheckTransferTaxTransactor creates a new write-only instance of CheckTransferTax, bound to a specific deployed contract.
func NewCheckTransferTaxTransactor(address common.Address, transactor bind.ContractTransactor) (*CheckTransferTaxTransactor, error) {
	contract, err := bindCheckTransferTax(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CheckTransferTaxTransactor{contract: contract}, nil
}

// NewCheckTransferTaxFilterer creates a new log filterer instance of CheckTransferTax, bound to a specific deployed contract.
func NewCheckTransferTaxFilterer(address common.Address, filterer bind.ContractFilterer) (*CheckTransferTaxFilterer, error) {
	contract, err := bindCheckTransferTax(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CheckTransferTaxFilterer{contract: contract}, nil
}

// bindCheckTransferTax binds a generic wrapper to an already deployed contract.
func bindCheckTransferTax(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckTransferTaxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckTransferTax *CheckTransferTaxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckTransferTax.Contract.CheckTransferTaxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckTransferTax *CheckTransferTaxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckTransferTax.Contract.CheckTransferTaxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckTransferTax *CheckTransferTaxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckTransferTax.Contract.CheckTransferTaxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckTransferTax *CheckTransferTaxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckTransferTax.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckTransferTax *CheckTransferTaxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckTransferTax.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckTransferTax *CheckTransferTaxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckTransferTax.Contract.contract.Transact(opts, method, params...)
}

// SwapWithoutFlashLoan is a paid mutator transaction binding the contract method 0x5cdffdca.
//
// Solidity: function swapWithoutFlashLoan(address[] _targets, bytes[] _payloads, address token) payable returns(uint256)
func (_CheckTransferTax *CheckTransferTaxTransactor) SwapWithoutFlashLoan(opts *bind.TransactOpts, _targets []common.Address, _payloads [][]byte, token common.Address) (*types.Transaction, error) {
	return _CheckTransferTax.contract.Transact(opts, "swapWithoutFlashLoan", _targets, _payloads, token)
}

// SwapWithoutFlashLoan is a paid mutator transaction binding the contract method 0x5cdffdca.
//
// Solidity: function swapWithoutFlashLoan(address[] _targets, bytes[] _payloads, address token) payable returns(uint256)
func (_CheckTransferTax *CheckTransferTaxSession) SwapWithoutFlashLoan(_targets []common.Address, _payloads [][]byte, token common.Address) (*types.Transaction, error) {
	return _CheckTransferTax.Contract.SwapWithoutFlashLoan(&_CheckTransferTax.TransactOpts, _targets, _payloads, token)
}

// SwapWithoutFlashLoan is a paid mutator transaction binding the contract method 0x5cdffdca.
//
// Solidity: function swapWithoutFlashLoan(address[] _targets, bytes[] _payloads, address token) payable returns(uint256)
func (_CheckTransferTax *CheckTransferTaxTransactorSession) SwapWithoutFlashLoan(_targets []common.Address, _payloads [][]byte, token common.Address) (*types.Transaction, error) {
	return _CheckTransferTax.Contract.SwapWithoutFlashLoan(&_CheckTransferTax.TransactOpts, _targets, _payloads, token)
}
