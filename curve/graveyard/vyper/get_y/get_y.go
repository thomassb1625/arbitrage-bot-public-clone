// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package get_Y

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

// GetYMetaData contains all meta data concerning the GetY contract.
var GetYMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_y\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"xp_\",\"type\":\"uint256[8]\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"N_COINS\",\"type\":\"uint256\"},{\"name\":\"D\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// GetYABI is the input ABI used to generate the binding from.
// Deprecated: Use GetYMetaData.ABI instead.
var GetYABI = GetYMetaData.ABI

// GetY is an auto generated Go binding around an Ethereum contract.
type GetY struct {
	GetYCaller     // Read-only binding to the contract
	GetYTransactor // Write-only binding to the contract
	GetYFilterer   // Log filterer for contract events
}

// GetYCaller is an auto generated read-only Go binding around an Ethereum contract.
type GetYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GetYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GetYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GetYSession struct {
	Contract     *GetY             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GetYCallerSession struct {
	Contract *GetYCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GetYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GetYTransactorSession struct {
	Contract     *GetYTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetYRaw is an auto generated low-level Go binding around an Ethereum contract.
type GetYRaw struct {
	Contract *GetY // Generic contract binding to access the raw methods on
}

// GetYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GetYCallerRaw struct {
	Contract *GetYCaller // Generic read-only contract binding to access the raw methods on
}

// GetYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GetYTransactorRaw struct {
	Contract *GetYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetY creates a new instance of GetY, bound to a specific deployed contract.
func NewGetY(address common.Address, backend bind.ContractBackend) (*GetY, error) {
	contract, err := bindGetY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GetY{GetYCaller: GetYCaller{contract: contract}, GetYTransactor: GetYTransactor{contract: contract}, GetYFilterer: GetYFilterer{contract: contract}}, nil
}

// NewGetYCaller creates a new read-only instance of GetY, bound to a specific deployed contract.
func NewGetYCaller(address common.Address, caller bind.ContractCaller) (*GetYCaller, error) {
	contract, err := bindGetY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GetYCaller{contract: contract}, nil
}

// NewGetYTransactor creates a new write-only instance of GetY, bound to a specific deployed contract.
func NewGetYTransactor(address common.Address, transactor bind.ContractTransactor) (*GetYTransactor, error) {
	contract, err := bindGetY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GetYTransactor{contract: contract}, nil
}

// NewGetYFilterer creates a new log filterer instance of GetY, bound to a specific deployed contract.
func NewGetYFilterer(address common.Address, filterer bind.ContractFilterer) (*GetYFilterer, error) {
	contract, err := bindGetY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GetYFilterer{contract: contract}, nil
}

// bindGetY binds a generic wrapper to an already deployed contract.
func bindGetY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GetYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetY *GetYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetY.Contract.GetYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetY *GetYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetY.Contract.GetYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetY *GetYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetY.Contract.GetYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetY *GetYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetY *GetYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetY *GetYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetY.Contract.contract.Transact(opts, method, params...)
}

// GetY is a free data retrieval call binding the contract method 0x812a07ce.
//
// Solidity: function get_y(int128 i, int128 j, uint256 x, uint256[8] xp_, uint256 A, uint256 N_COINS, uint256 D) view returns(uint256)
func (_GetY *GetYCaller) GetY(opts *bind.CallOpts, i *big.Int, j *big.Int, x *big.Int, xp_ [8]*big.Int, A *big.Int, N_COINS *big.Int, D *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GetY.contract.Call(opts, &out, "get_y", i, j, x, xp_, A, N_COINS, D)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetY is a free data retrieval call binding the contract method 0x812a07ce.
//
// Solidity: function get_y(int128 i, int128 j, uint256 x, uint256[8] xp_, uint256 A, uint256 N_COINS, uint256 D) view returns(uint256)
func (_GetY *GetYSession) GetY(i *big.Int, j *big.Int, x *big.Int, xp_ [8]*big.Int, A *big.Int, N_COINS *big.Int, D *big.Int) (*big.Int, error) {
	return _GetY.Contract.GetY(&_GetY.CallOpts, i, j, x, xp_, A, N_COINS, D)
}

// GetY is a free data retrieval call binding the contract method 0x812a07ce.
//
// Solidity: function get_y(int128 i, int128 j, uint256 x, uint256[8] xp_, uint256 A, uint256 N_COINS, uint256 D) view returns(uint256)
func (_GetY *GetYCallerSession) GetY(i *big.Int, j *big.Int, x *big.Int, xp_ [8]*big.Int, A *big.Int, N_COINS *big.Int, D *big.Int) (*big.Int, error) {
	return _GetY.Contract.GetY(&_GetY.CallOpts, i, j, x, xp_, A, N_COINS, D)
}
