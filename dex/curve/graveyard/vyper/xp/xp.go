// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xp

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

// XpMetaData contains all meta data concerning the Xp contract.
var XpMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"_xp\",\"inputs\":[{\"name\":\"balances\",\"type\":\"uint256[8]\"},{\"name\":\"N_COINS\",\"type\":\"uint256\"},{\"name\":\"RATES\",\"type\":\"uint256[8]\"},{\"name\":\"LENDING_PRECISION\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]}]",
}

// XpABI is the input ABI used to generate the binding from.
// Deprecated: Use XpMetaData.ABI instead.
var XpABI = XpMetaData.ABI

// Xp is an auto generated Go binding around an Ethereum contract.
type Xp struct {
	XpCaller     // Read-only binding to the contract
	XpTransactor // Write-only binding to the contract
	XpFilterer   // Log filterer for contract events
}

// XpCaller is an auto generated read-only Go binding around an Ethereum contract.
type XpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XpSession struct {
	Contract     *Xp               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XpCallerSession struct {
	Contract *XpCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// XpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XpTransactorSession struct {
	Contract     *XpTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XpRaw is an auto generated low-level Go binding around an Ethereum contract.
type XpRaw struct {
	Contract *Xp // Generic contract binding to access the raw methods on
}

// XpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XpCallerRaw struct {
	Contract *XpCaller // Generic read-only contract binding to access the raw methods on
}

// XpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XpTransactorRaw struct {
	Contract *XpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXp creates a new instance of Xp, bound to a specific deployed contract.
func NewXp(address common.Address, backend bind.ContractBackend) (*Xp, error) {
	contract, err := bindXp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Xp{XpCaller: XpCaller{contract: contract}, XpTransactor: XpTransactor{contract: contract}, XpFilterer: XpFilterer{contract: contract}}, nil
}

// NewXpCaller creates a new read-only instance of Xp, bound to a specific deployed contract.
func NewXpCaller(address common.Address, caller bind.ContractCaller) (*XpCaller, error) {
	contract, err := bindXp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XpCaller{contract: contract}, nil
}

// NewXpTransactor creates a new write-only instance of Xp, bound to a specific deployed contract.
func NewXpTransactor(address common.Address, transactor bind.ContractTransactor) (*XpTransactor, error) {
	contract, err := bindXp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XpTransactor{contract: contract}, nil
}

// NewXpFilterer creates a new log filterer instance of Xp, bound to a specific deployed contract.
func NewXpFilterer(address common.Address, filterer bind.ContractFilterer) (*XpFilterer, error) {
	contract, err := bindXp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XpFilterer{contract: contract}, nil
}

// bindXp binds a generic wrapper to an already deployed contract.
func bindXp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xp *XpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xp.Contract.XpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xp *XpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xp.Contract.XpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xp *XpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xp.Contract.XpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xp *XpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xp *XpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xp *XpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xp.Contract.contract.Transact(opts, method, params...)
}

// Xp is a free data retrieval call binding the contract method 0x0a957ba6.
//
// Solidity: function _xp(uint256[8] balances, uint256 N_COINS, uint256[8] RATES, uint256 LENDING_PRECISION) view returns(uint256[8])
func (_Xp *XpCaller) Xp(opts *bind.CallOpts, balances [8]*big.Int, N_COINS *big.Int, RATES [8]*big.Int, LENDING_PRECISION *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Xp.contract.Call(opts, &out, "_xp", balances, N_COINS, RATES, LENDING_PRECISION)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// Xp is a free data retrieval call binding the contract method 0x0a957ba6.
//
// Solidity: function _xp(uint256[8] balances, uint256 N_COINS, uint256[8] RATES, uint256 LENDING_PRECISION) view returns(uint256[8])
func (_Xp *XpSession) Xp(balances [8]*big.Int, N_COINS *big.Int, RATES [8]*big.Int, LENDING_PRECISION *big.Int) ([8]*big.Int, error) {
	return _Xp.Contract.Xp(&_Xp.CallOpts, balances, N_COINS, RATES, LENDING_PRECISION)
}

// Xp is a free data retrieval call binding the contract method 0x0a957ba6.
//
// Solidity: function _xp(uint256[8] balances, uint256 N_COINS, uint256[8] RATES, uint256 LENDING_PRECISION) view returns(uint256[8])
func (_Xp *XpCallerSession) Xp(balances [8]*big.Int, N_COINS *big.Int, RATES [8]*big.Int, LENDING_PRECISION *big.Int) ([8]*big.Int, error) {
	return _Xp.Contract.Xp(&_Xp.CallOpts, balances, N_COINS, RATES, LENDING_PRECISION)
}
