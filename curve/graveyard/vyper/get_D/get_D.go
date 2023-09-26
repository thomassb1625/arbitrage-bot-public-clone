// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package get_D

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

// GetDMetaData contains all meta data concerning the GetD contract.
var GetDMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"get_D\",\"inputs\":[{\"name\":\"xp\",\"type\":\"uint256[8]\"},{\"name\":\"amp\",\"type\":\"uint256\"},{\"name\":\"N_COINS\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// GetDABI is the input ABI used to generate the binding from.
// Deprecated: Use GetDMetaData.ABI instead.
var GetDABI = GetDMetaData.ABI

// GetD is an auto generated Go binding around an Ethereum contract.
type GetD struct {
	GetDCaller     // Read-only binding to the contract
	GetDTransactor // Write-only binding to the contract
	GetDFilterer   // Log filterer for contract events
}

// GetDCaller is an auto generated read-only Go binding around an Ethereum contract.
type GetDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GetDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GetDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GetDSession struct {
	Contract     *GetD             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GetDCallerSession struct {
	Contract *GetDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GetDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GetDTransactorSession struct {
	Contract     *GetDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetDRaw is an auto generated low-level Go binding around an Ethereum contract.
type GetDRaw struct {
	Contract *GetD // Generic contract binding to access the raw methods on
}

// GetDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GetDCallerRaw struct {
	Contract *GetDCaller // Generic read-only contract binding to access the raw methods on
}

// GetDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GetDTransactorRaw struct {
	Contract *GetDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetD creates a new instance of GetD, bound to a specific deployed contract.
func NewGetD(address common.Address, backend bind.ContractBackend) (*GetD, error) {
	contract, err := bindGetD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GetD{GetDCaller: GetDCaller{contract: contract}, GetDTransactor: GetDTransactor{contract: contract}, GetDFilterer: GetDFilterer{contract: contract}}, nil
}

// NewGetDCaller creates a new read-only instance of GetD, bound to a specific deployed contract.
func NewGetDCaller(address common.Address, caller bind.ContractCaller) (*GetDCaller, error) {
	contract, err := bindGetD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GetDCaller{contract: contract}, nil
}

// NewGetDTransactor creates a new write-only instance of GetD, bound to a specific deployed contract.
func NewGetDTransactor(address common.Address, transactor bind.ContractTransactor) (*GetDTransactor, error) {
	contract, err := bindGetD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GetDTransactor{contract: contract}, nil
}

// NewGetDFilterer creates a new log filterer instance of GetD, bound to a specific deployed contract.
func NewGetDFilterer(address common.Address, filterer bind.ContractFilterer) (*GetDFilterer, error) {
	contract, err := bindGetD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GetDFilterer{contract: contract}, nil
}

// bindGetD binds a generic wrapper to an already deployed contract.
func bindGetD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GetDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetD *GetDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetD.Contract.GetDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetD *GetDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetD.Contract.GetDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetD *GetDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetD.Contract.GetDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetD *GetDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetD *GetDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetD *GetDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetD.Contract.contract.Transact(opts, method, params...)
}

// GetD is a free data retrieval call binding the contract method 0x24dbe384.
//
// Solidity: function get_D(uint256[8] xp, uint256 amp, uint256 N_COINS) pure returns(uint256)
func (_GetD *GetDCaller) GetD(opts *bind.CallOpts, xp [8]*big.Int, amp *big.Int, N_COINS *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GetD.contract.Call(opts, &out, "get_D", xp, amp, N_COINS)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetD is a free data retrieval call binding the contract method 0x24dbe384.
//
// Solidity: function get_D(uint256[8] xp, uint256 amp, uint256 N_COINS) pure returns(uint256)
func (_GetD *GetDSession) GetD(xp [8]*big.Int, amp *big.Int, N_COINS *big.Int) (*big.Int, error) {
	return _GetD.Contract.GetD(&_GetD.CallOpts, xp, amp, N_COINS)
}

// GetD is a free data retrieval call binding the contract method 0x24dbe384.
//
// Solidity: function get_D(uint256[8] xp, uint256 amp, uint256 N_COINS) pure returns(uint256)
func (_GetD *GetDCallerSession) GetD(xp [8]*big.Int, amp *big.Int, N_COINS *big.Int) (*big.Int, error) {
	return _GetD.Contract.GetD(&_GetD.CallOpts, xp, amp, N_COINS)
}
