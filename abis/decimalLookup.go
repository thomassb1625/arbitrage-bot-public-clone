// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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

// DecimalLookupMetaData contains all meta data concerning the DecimalLookup contract.
var DecimalLookupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"lookupDecimals\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DecimalLookupABI is the input ABI used to generate the binding from.
// Deprecated: Use DecimalLookupMetaData.ABI instead.
var DecimalLookupABI = DecimalLookupMetaData.ABI

// DecimalLookup is an auto generated Go binding around an Ethereum contract.
type DecimalLookup struct {
	DecimalLookupCaller     // Read-only binding to the contract
	DecimalLookupTransactor // Write-only binding to the contract
	DecimalLookupFilterer   // Log filterer for contract events
}

// DecimalLookupCaller is an auto generated read-only Go binding around an Ethereum contract.
type DecimalLookupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecimalLookupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DecimalLookupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecimalLookupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DecimalLookupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecimalLookupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DecimalLookupSession struct {
	Contract     *DecimalLookup    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DecimalLookupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DecimalLookupCallerSession struct {
	Contract *DecimalLookupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DecimalLookupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DecimalLookupTransactorSession struct {
	Contract     *DecimalLookupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DecimalLookupRaw is an auto generated low-level Go binding around an Ethereum contract.
type DecimalLookupRaw struct {
	Contract *DecimalLookup // Generic contract binding to access the raw methods on
}

// DecimalLookupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DecimalLookupCallerRaw struct {
	Contract *DecimalLookupCaller // Generic read-only contract binding to access the raw methods on
}

// DecimalLookupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DecimalLookupTransactorRaw struct {
	Contract *DecimalLookupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDecimalLookup creates a new instance of DecimalLookup, bound to a specific deployed contract.
func NewDecimalLookup(address common.Address, backend bind.ContractBackend) (*DecimalLookup, error) {
	contract, err := bindDecimalLookup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DecimalLookup{DecimalLookupCaller: DecimalLookupCaller{contract: contract}, DecimalLookupTransactor: DecimalLookupTransactor{contract: contract}, DecimalLookupFilterer: DecimalLookupFilterer{contract: contract}}, nil
}

// NewDecimalLookupCaller creates a new read-only instance of DecimalLookup, bound to a specific deployed contract.
func NewDecimalLookupCaller(address common.Address, caller bind.ContractCaller) (*DecimalLookupCaller, error) {
	contract, err := bindDecimalLookup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DecimalLookupCaller{contract: contract}, nil
}

// NewDecimalLookupTransactor creates a new write-only instance of DecimalLookup, bound to a specific deployed contract.
func NewDecimalLookupTransactor(address common.Address, transactor bind.ContractTransactor) (*DecimalLookupTransactor, error) {
	contract, err := bindDecimalLookup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DecimalLookupTransactor{contract: contract}, nil
}

// NewDecimalLookupFilterer creates a new log filterer instance of DecimalLookup, bound to a specific deployed contract.
func NewDecimalLookupFilterer(address common.Address, filterer bind.ContractFilterer) (*DecimalLookupFilterer, error) {
	contract, err := bindDecimalLookup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DecimalLookupFilterer{contract: contract}, nil
}

// bindDecimalLookup binds a generic wrapper to an already deployed contract.
func bindDecimalLookup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DecimalLookupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecimalLookup *DecimalLookupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecimalLookup.Contract.DecimalLookupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecimalLookup *DecimalLookupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecimalLookup.Contract.DecimalLookupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecimalLookup *DecimalLookupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecimalLookup.Contract.DecimalLookupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecimalLookup *DecimalLookupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecimalLookup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecimalLookup *DecimalLookupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecimalLookup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecimalLookup *DecimalLookupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecimalLookup.Contract.contract.Transact(opts, method, params...)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(bool)
func (_DecimalLookup *DecimalLookupCaller) IsContract(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _DecimalLookup.contract.Call(opts, &out, "isContract", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(bool)
func (_DecimalLookup *DecimalLookupSession) IsContract(addr common.Address) (bool, error) {
	return _DecimalLookup.Contract.IsContract(&_DecimalLookup.CallOpts, addr)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(bool)
func (_DecimalLookup *DecimalLookupCallerSession) IsContract(addr common.Address) (bool, error) {
	return _DecimalLookup.Contract.IsContract(&_DecimalLookup.CallOpts, addr)
}

// LookupDecimals is a free data retrieval call binding the contract method 0xd0cbe797.
//
// Solidity: function lookupDecimals(address[] _tokens) view returns(int256[])
func (_DecimalLookup *DecimalLookupCaller) LookupDecimals(opts *bind.CallOpts, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _DecimalLookup.contract.Call(opts, &out, "lookupDecimals", _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// LookupDecimals is a free data retrieval call binding the contract method 0xd0cbe797.
//
// Solidity: function lookupDecimals(address[] _tokens) view returns(int256[])
func (_DecimalLookup *DecimalLookupSession) LookupDecimals(_tokens []common.Address) ([]*big.Int, error) {
	return _DecimalLookup.Contract.LookupDecimals(&_DecimalLookup.CallOpts, _tokens)
}

// LookupDecimals is a free data retrieval call binding the contract method 0xd0cbe797.
//
// Solidity: function lookupDecimals(address[] _tokens) view returns(int256[])
func (_DecimalLookup *DecimalLookupCallerSession) LookupDecimals(_tokens []common.Address) ([]*big.Int, error) {
	return _DecimalLookup.Contract.LookupDecimals(&_DecimalLookup.CallOpts, _tokens)
}
