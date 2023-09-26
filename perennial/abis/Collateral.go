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

// CollateralMetaData contains all meta data concerning the Collateral contract.
var CollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"Token18\",\"name\":\"token_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"CollateralAccountLiquidatingError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"totalMaintenance\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"totalCollateral\",\"type\":\"uint256\"}],\"name\":\"CollateralCantLiquidate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralInsufficientCollateralError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralUnderLimitError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralZeroAddressError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fixed18OverflowError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidControllerError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"NotAccountOrMultiInvokerError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCollateralError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"coordinatorId\",\"type\":\"uint256\"}],\"name\":\"NotOwnerError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"}],\"name\":\"NotProductError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"UInitializableAlreadyInitializedError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UInitializableNotInitializingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UInitializableZeroVersionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UReentrancyGuardReentrantCallError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"Fixed18\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newShortfall\",\"type\":\"uint256\"}],\"name\":\"AccountSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Liquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"productFee\",\"type\":\"uint256\"}],\"name\":\"ProductSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ShortfallResolution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"}],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"}],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIController\",\"name\":\"controller_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"}],\"name\":\"liquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"}],\"name\":\"liquidatableNext\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"resolveShortfall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"Fixed18\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"settleAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"settleProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"}],\"name\":\"shortfall\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"Token18\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"},{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use CollateralMetaData.ABI instead.
var CollateralABI = CollateralMetaData.ABI

// Collateral is an auto generated Go binding around an Ethereum contract.
type Collateral struct {
	CollateralCaller     // Read-only binding to the contract
	CollateralTransactor // Write-only binding to the contract
	CollateralFilterer   // Log filterer for contract events
}

// CollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollateralSession struct {
	Contract     *Collateral       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollateralCallerSession struct {
	Contract *CollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollateralTransactorSession struct {
	Contract     *CollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollateralRaw struct {
	Contract *Collateral // Generic contract binding to access the raw methods on
}

// CollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollateralCallerRaw struct {
	Contract *CollateralCaller // Generic read-only contract binding to access the raw methods on
}

// CollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollateralTransactorRaw struct {
	Contract *CollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollateral creates a new instance of Collateral, bound to a specific deployed contract.
func NewCollateral(address common.Address, backend bind.ContractBackend) (*Collateral, error) {
	contract, err := bindCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Collateral{CollateralCaller: CollateralCaller{contract: contract}, CollateralTransactor: CollateralTransactor{contract: contract}, CollateralFilterer: CollateralFilterer{contract: contract}}, nil
}

// NewCollateralCaller creates a new read-only instance of Collateral, bound to a specific deployed contract.
func NewCollateralCaller(address common.Address, caller bind.ContractCaller) (*CollateralCaller, error) {
	contract, err := bindCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralCaller{contract: contract}, nil
}

// NewCollateralTransactor creates a new write-only instance of Collateral, bound to a specific deployed contract.
func NewCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*CollateralTransactor, error) {
	contract, err := bindCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralTransactor{contract: contract}, nil
}

// NewCollateralFilterer creates a new log filterer instance of Collateral, bound to a specific deployed contract.
func NewCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*CollateralFilterer, error) {
	contract, err := bindCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollateralFilterer{contract: contract}, nil
}

// bindCollateral binds a generic wrapper to an already deployed contract.
func bindCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CollateralABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collateral *CollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Collateral.Contract.CollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collateral *CollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.Contract.CollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collateral *CollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collateral.Contract.CollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collateral *CollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Collateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collateral *CollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collateral *CollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collateral.Contract.contract.Transact(opts, method, params...)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address product) view returns(uint256)
func (_Collateral *CollateralCaller) Collateral(opts *bind.CallOpts, product common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "collateral", product)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address product) view returns(uint256)
func (_Collateral *CollateralSession) Collateral(product common.Address) (*big.Int, error) {
	return _Collateral.Contract.Collateral(&_Collateral.CallOpts, product)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address product) view returns(uint256)
func (_Collateral *CollateralCallerSession) Collateral(product common.Address) (*big.Int, error) {
	return _Collateral.Contract.Collateral(&_Collateral.CallOpts, product)
}

// Collateral0 is a free data retrieval call binding the contract method 0xcc218ece.
//
// Solidity: function collateral(address account, address product) view returns(uint256)
func (_Collateral *CollateralCaller) Collateral0(opts *bind.CallOpts, account common.Address, product common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "collateral0", account, product)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral0 is a free data retrieval call binding the contract method 0xcc218ece.
//
// Solidity: function collateral(address account, address product) view returns(uint256)
func (_Collateral *CollateralSession) Collateral0(account common.Address, product common.Address) (*big.Int, error) {
	return _Collateral.Contract.Collateral0(&_Collateral.CallOpts, account, product)
}

// Collateral0 is a free data retrieval call binding the contract method 0xcc218ece.
//
// Solidity: function collateral(address account, address product) view returns(uint256)
func (_Collateral *CollateralCallerSession) Collateral0(account common.Address, product common.Address) (*big.Int, error) {
	return _Collateral.Contract.Collateral0(&_Collateral.CallOpts, account, product)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Collateral *CollateralCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Collateral *CollateralSession) Controller() (common.Address, error) {
	return _Collateral.Contract.Controller(&_Collateral.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Collateral *CollateralCallerSession) Controller() (common.Address, error) {
	return _Collateral.Contract.Controller(&_Collateral.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0xfaaebd21.
//
// Solidity: function fees(address ) view returns(uint256)
func (_Collateral *CollateralCaller) Fees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "fees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0xfaaebd21.
//
// Solidity: function fees(address ) view returns(uint256)
func (_Collateral *CollateralSession) Fees(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.Fees(&_Collateral.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0xfaaebd21.
//
// Solidity: function fees(address ) view returns(uint256)
func (_Collateral *CollateralCallerSession) Fees(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.Fees(&_Collateral.CallOpts, arg0)
}

// Liquidatable is a free data retrieval call binding the contract method 0x82df39de.
//
// Solidity: function liquidatable(address account, address product) view returns(bool)
func (_Collateral *CollateralCaller) Liquidatable(opts *bind.CallOpts, account common.Address, product common.Address) (bool, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "liquidatable", account, product)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Liquidatable is a free data retrieval call binding the contract method 0x82df39de.
//
// Solidity: function liquidatable(address account, address product) view returns(bool)
func (_Collateral *CollateralSession) Liquidatable(account common.Address, product common.Address) (bool, error) {
	return _Collateral.Contract.Liquidatable(&_Collateral.CallOpts, account, product)
}

// Liquidatable is a free data retrieval call binding the contract method 0x82df39de.
//
// Solidity: function liquidatable(address account, address product) view returns(bool)
func (_Collateral *CollateralCallerSession) Liquidatable(account common.Address, product common.Address) (bool, error) {
	return _Collateral.Contract.Liquidatable(&_Collateral.CallOpts, account, product)
}

// LiquidatableNext is a free data retrieval call binding the contract method 0xe4119180.
//
// Solidity: function liquidatableNext(address account, address product) view returns(bool)
func (_Collateral *CollateralCaller) LiquidatableNext(opts *bind.CallOpts, account common.Address, product common.Address) (bool, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "liquidatableNext", account, product)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LiquidatableNext is a free data retrieval call binding the contract method 0xe4119180.
//
// Solidity: function liquidatableNext(address account, address product) view returns(bool)
func (_Collateral *CollateralSession) LiquidatableNext(account common.Address, product common.Address) (bool, error) {
	return _Collateral.Contract.LiquidatableNext(&_Collateral.CallOpts, account, product)
}

// LiquidatableNext is a free data retrieval call binding the contract method 0xe4119180.
//
// Solidity: function liquidatableNext(address account, address product) view returns(bool)
func (_Collateral *CollateralCallerSession) LiquidatableNext(account common.Address, product common.Address) (bool, error) {
	return _Collateral.Contract.LiquidatableNext(&_Collateral.CallOpts, account, product)
}

// Shortfall is a free data retrieval call binding the contract method 0xb4d6f781.
//
// Solidity: function shortfall(address product) view returns(uint256)
func (_Collateral *CollateralCaller) Shortfall(opts *bind.CallOpts, product common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "shortfall", product)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shortfall is a free data retrieval call binding the contract method 0xb4d6f781.
//
// Solidity: function shortfall(address product) view returns(uint256)
func (_Collateral *CollateralSession) Shortfall(product common.Address) (*big.Int, error) {
	return _Collateral.Contract.Shortfall(&_Collateral.CallOpts, product)
}

// Shortfall is a free data retrieval call binding the contract method 0xb4d6f781.
//
// Solidity: function shortfall(address product) view returns(uint256)
func (_Collateral *CollateralCallerSession) Shortfall(product common.Address) (*big.Int, error) {
	return _Collateral.Contract.Shortfall(&_Collateral.CallOpts, product)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Collateral *CollateralCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Collateral *CollateralSession) Token() (common.Address, error) {
	return _Collateral.Contract.Token(&_Collateral.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Collateral *CollateralCallerSession) Token() (common.Address, error) {
	return _Collateral.Contract.Token(&_Collateral.CallOpts)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_Collateral *CollateralTransactor) ClaimFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "claimFee")
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_Collateral *CollateralSession) ClaimFee() (*types.Transaction, error) {
	return _Collateral.Contract.ClaimFee(&_Collateral.TransactOpts)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_Collateral *CollateralTransactorSession) ClaimFee() (*types.Transaction, error) {
	return _Collateral.Contract.ClaimFee(&_Collateral.TransactOpts)
}

// DepositTo is a paid mutator transaction binding the contract method 0xf213159c.
//
// Solidity: function depositTo(address account, address product, uint256 amount) returns()
func (_Collateral *CollateralTransactor) DepositTo(opts *bind.TransactOpts, account common.Address, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "depositTo", account, product, amount)
}

// DepositTo is a paid mutator transaction binding the contract method 0xf213159c.
//
// Solidity: function depositTo(address account, address product, uint256 amount) returns()
func (_Collateral *CollateralSession) DepositTo(account common.Address, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.DepositTo(&_Collateral.TransactOpts, account, product, amount)
}

// DepositTo is a paid mutator transaction binding the contract method 0xf213159c.
//
// Solidity: function depositTo(address account, address product, uint256 amount) returns()
func (_Collateral *CollateralTransactorSession) DepositTo(account common.Address, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.DepositTo(&_Collateral.TransactOpts, account, product, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address controller_) returns()
func (_Collateral *CollateralTransactor) Initialize(opts *bind.TransactOpts, controller_ common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "initialize", controller_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address controller_) returns()
func (_Collateral *CollateralSession) Initialize(controller_ common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Initialize(&_Collateral.TransactOpts, controller_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address controller_) returns()
func (_Collateral *CollateralTransactorSession) Initialize(controller_ common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Initialize(&_Collateral.TransactOpts, controller_)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address account, address product) returns()
func (_Collateral *CollateralTransactor) Liquidate(opts *bind.TransactOpts, account common.Address, product common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "liquidate", account, product)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address account, address product) returns()
func (_Collateral *CollateralSession) Liquidate(account common.Address, product common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Liquidate(&_Collateral.TransactOpts, account, product)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address account, address product) returns()
func (_Collateral *CollateralTransactorSession) Liquidate(account common.Address, product common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Liquidate(&_Collateral.TransactOpts, account, product)
}

// ResolveShortfall is a paid mutator transaction binding the contract method 0xca70d300.
//
// Solidity: function resolveShortfall(address product, uint256 amount) returns()
func (_Collateral *CollateralTransactor) ResolveShortfall(opts *bind.TransactOpts, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "resolveShortfall", product, amount)
}

// ResolveShortfall is a paid mutator transaction binding the contract method 0xca70d300.
//
// Solidity: function resolveShortfall(address product, uint256 amount) returns()
func (_Collateral *CollateralSession) ResolveShortfall(product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.ResolveShortfall(&_Collateral.TransactOpts, product, amount)
}

// ResolveShortfall is a paid mutator transaction binding the contract method 0xca70d300.
//
// Solidity: function resolveShortfall(address product, uint256 amount) returns()
func (_Collateral *CollateralTransactorSession) ResolveShortfall(product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.ResolveShortfall(&_Collateral.TransactOpts, product, amount)
}

// SettleAccount is a paid mutator transaction binding the contract method 0xc9cb48de.
//
// Solidity: function settleAccount(address account, int256 amount) returns()
func (_Collateral *CollateralTransactor) SettleAccount(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "settleAccount", account, amount)
}

// SettleAccount is a paid mutator transaction binding the contract method 0xc9cb48de.
//
// Solidity: function settleAccount(address account, int256 amount) returns()
func (_Collateral *CollateralSession) SettleAccount(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SettleAccount(&_Collateral.TransactOpts, account, amount)
}

// SettleAccount is a paid mutator transaction binding the contract method 0xc9cb48de.
//
// Solidity: function settleAccount(address account, int256 amount) returns()
func (_Collateral *CollateralTransactorSession) SettleAccount(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SettleAccount(&_Collateral.TransactOpts, account, amount)
}

// SettleProduct is a paid mutator transaction binding the contract method 0x7f3dd53b.
//
// Solidity: function settleProduct(uint256 amount) returns()
func (_Collateral *CollateralTransactor) SettleProduct(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "settleProduct", amount)
}

// SettleProduct is a paid mutator transaction binding the contract method 0x7f3dd53b.
//
// Solidity: function settleProduct(uint256 amount) returns()
func (_Collateral *CollateralSession) SettleProduct(amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SettleProduct(&_Collateral.TransactOpts, amount)
}

// SettleProduct is a paid mutator transaction binding the contract method 0x7f3dd53b.
//
// Solidity: function settleProduct(uint256 amount) returns()
func (_Collateral *CollateralTransactorSession) SettleProduct(amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SettleProduct(&_Collateral.TransactOpts, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x26441318.
//
// Solidity: function withdrawFrom(address account, address receiver, address product, uint256 amount) returns()
func (_Collateral *CollateralTransactor) WithdrawFrom(opts *bind.TransactOpts, account common.Address, receiver common.Address, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "withdrawFrom", account, receiver, product, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x26441318.
//
// Solidity: function withdrawFrom(address account, address receiver, address product, uint256 amount) returns()
func (_Collateral *CollateralSession) WithdrawFrom(account common.Address, receiver common.Address, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.WithdrawFrom(&_Collateral.TransactOpts, account, receiver, product, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x26441318.
//
// Solidity: function withdrawFrom(address account, address receiver, address product, uint256 amount) returns()
func (_Collateral *CollateralTransactorSession) WithdrawFrom(account common.Address, receiver common.Address, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.WithdrawFrom(&_Collateral.TransactOpts, account, receiver, product, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address receiver, address product, uint256 amount) returns()
func (_Collateral *CollateralTransactor) WithdrawTo(opts *bind.TransactOpts, receiver common.Address, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "withdrawTo", receiver, product, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address receiver, address product, uint256 amount) returns()
func (_Collateral *CollateralSession) WithdrawTo(receiver common.Address, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.WithdrawTo(&_Collateral.TransactOpts, receiver, product, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address receiver, address product, uint256 amount) returns()
func (_Collateral *CollateralTransactorSession) WithdrawTo(receiver common.Address, product common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.WithdrawTo(&_Collateral.TransactOpts, receiver, product, amount)
}

// CollateralAccountSettleIterator is returned from FilterAccountSettle and is used to iterate over the raw logs and unpacked data for AccountSettle events raised by the Collateral contract.
type CollateralAccountSettleIterator struct {
	Event *CollateralAccountSettle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralAccountSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralAccountSettle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralAccountSettle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralAccountSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralAccountSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralAccountSettle represents a AccountSettle event raised by the Collateral contract.
type CollateralAccountSettle struct {
	Product      common.Address
	Account      common.Address
	Amount       *big.Int
	NewShortfall *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountSettle is a free log retrieval operation binding the contract event 0x40957a7f26dd9f9d628b1aba01d4a3db4c38008974c79a62f6cf0f273a78b3a8.
//
// Solidity: event AccountSettle(address indexed product, address indexed account, int256 amount, uint256 newShortfall)
func (_Collateral *CollateralFilterer) FilterAccountSettle(opts *bind.FilterOpts, product []common.Address, account []common.Address) (*CollateralAccountSettleIterator, error) {

	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "AccountSettle", productRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &CollateralAccountSettleIterator{contract: _Collateral.contract, event: "AccountSettle", logs: logs, sub: sub}, nil
}

// WatchAccountSettle is a free log subscription operation binding the contract event 0x40957a7f26dd9f9d628b1aba01d4a3db4c38008974c79a62f6cf0f273a78b3a8.
//
// Solidity: event AccountSettle(address indexed product, address indexed account, int256 amount, uint256 newShortfall)
func (_Collateral *CollateralFilterer) WatchAccountSettle(opts *bind.WatchOpts, sink chan<- *CollateralAccountSettle, product []common.Address, account []common.Address) (event.Subscription, error) {

	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "AccountSettle", productRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralAccountSettle)
				if err := _Collateral.contract.UnpackLog(event, "AccountSettle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountSettle is a log parse operation binding the contract event 0x40957a7f26dd9f9d628b1aba01d4a3db4c38008974c79a62f6cf0f273a78b3a8.
//
// Solidity: event AccountSettle(address indexed product, address indexed account, int256 amount, uint256 newShortfall)
func (_Collateral *CollateralFilterer) ParseAccountSettle(log types.Log) (*CollateralAccountSettle, error) {
	event := new(CollateralAccountSettle)
	if err := _Collateral.contract.UnpackLog(event, "AccountSettle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Collateral contract.
type CollateralDepositIterator struct {
	Event *CollateralDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralDeposit represents a Deposit event raised by the Collateral contract.
type CollateralDeposit struct {
	User    common.Address
	Product common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed product, uint256 amount)
func (_Collateral *CollateralFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, product []common.Address) (*CollateralDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Deposit", userRule, productRule)
	if err != nil {
		return nil, err
	}
	return &CollateralDepositIterator{contract: _Collateral.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed product, uint256 amount)
func (_Collateral *CollateralFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CollateralDeposit, user []common.Address, product []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Deposit", userRule, productRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralDeposit)
				if err := _Collateral.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed product, uint256 amount)
func (_Collateral *CollateralFilterer) ParseDeposit(log types.Log) (*CollateralDeposit, error) {
	event := new(CollateralDeposit)
	if err := _Collateral.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralFeeClaimIterator is returned from FilterFeeClaim and is used to iterate over the raw logs and unpacked data for FeeClaim events raised by the Collateral contract.
type CollateralFeeClaimIterator struct {
	Event *CollateralFeeClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralFeeClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralFeeClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralFeeClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralFeeClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralFeeClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralFeeClaim represents a FeeClaim event raised by the Collateral contract.
type CollateralFeeClaim struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeClaim is a free log retrieval operation binding the contract event 0x10df095d1434aed409b2f33d2a6a8456f8b0824633cc12a1b43032085aadc41d.
//
// Solidity: event FeeClaim(address indexed account, uint256 amount)
func (_Collateral *CollateralFilterer) FilterFeeClaim(opts *bind.FilterOpts, account []common.Address) (*CollateralFeeClaimIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "FeeClaim", accountRule)
	if err != nil {
		return nil, err
	}
	return &CollateralFeeClaimIterator{contract: _Collateral.contract, event: "FeeClaim", logs: logs, sub: sub}, nil
}

// WatchFeeClaim is a free log subscription operation binding the contract event 0x10df095d1434aed409b2f33d2a6a8456f8b0824633cc12a1b43032085aadc41d.
//
// Solidity: event FeeClaim(address indexed account, uint256 amount)
func (_Collateral *CollateralFilterer) WatchFeeClaim(opts *bind.WatchOpts, sink chan<- *CollateralFeeClaim, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "FeeClaim", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralFeeClaim)
				if err := _Collateral.contract.UnpackLog(event, "FeeClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeClaim is a log parse operation binding the contract event 0x10df095d1434aed409b2f33d2a6a8456f8b0824633cc12a1b43032085aadc41d.
//
// Solidity: event FeeClaim(address indexed account, uint256 amount)
func (_Collateral *CollateralFilterer) ParseFeeClaim(log types.Log) (*CollateralFeeClaim, error) {
	event := new(CollateralFeeClaim)
	if err := _Collateral.contract.UnpackLog(event, "FeeClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Collateral contract.
type CollateralInitializedIterator struct {
	Event *CollateralInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralInitialized represents a Initialized event raised by the Collateral contract.
type CollateralInitialized struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xbe9b076dc5b65990cca9dd9d7366682482e7817a6f6bc7f4faf4dc32af497f32.
//
// Solidity: event Initialized(uint256 version)
func (_Collateral *CollateralFilterer) FilterInitialized(opts *bind.FilterOpts) (*CollateralInitializedIterator, error) {

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CollateralInitializedIterator{contract: _Collateral.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xbe9b076dc5b65990cca9dd9d7366682482e7817a6f6bc7f4faf4dc32af497f32.
//
// Solidity: event Initialized(uint256 version)
func (_Collateral *CollateralFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CollateralInitialized) (event.Subscription, error) {

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralInitialized)
				if err := _Collateral.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xbe9b076dc5b65990cca9dd9d7366682482e7817a6f6bc7f4faf4dc32af497f32.
//
// Solidity: event Initialized(uint256 version)
func (_Collateral *CollateralFilterer) ParseInitialized(log types.Log) (*CollateralInitialized, error) {
	event := new(CollateralInitialized)
	if err := _Collateral.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralLiquidationIterator is returned from FilterLiquidation and is used to iterate over the raw logs and unpacked data for Liquidation events raised by the Collateral contract.
type CollateralLiquidationIterator struct {
	Event *CollateralLiquidation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralLiquidation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralLiquidation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralLiquidation represents a Liquidation event raised by the Collateral contract.
type CollateralLiquidation struct {
	User       common.Address
	Product    common.Address
	Liquidator common.Address
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidation is a free log retrieval operation binding the contract event 0x76f1c54646947e8e8df3eab8141bb3d115d07da2520f8e41b208a17d138a8959.
//
// Solidity: event Liquidation(address indexed user, address indexed product, address liquidator, uint256 fee)
func (_Collateral *CollateralFilterer) FilterLiquidation(opts *bind.FilterOpts, user []common.Address, product []common.Address) (*CollateralLiquidationIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Liquidation", userRule, productRule)
	if err != nil {
		return nil, err
	}
	return &CollateralLiquidationIterator{contract: _Collateral.contract, event: "Liquidation", logs: logs, sub: sub}, nil
}

// WatchLiquidation is a free log subscription operation binding the contract event 0x76f1c54646947e8e8df3eab8141bb3d115d07da2520f8e41b208a17d138a8959.
//
// Solidity: event Liquidation(address indexed user, address indexed product, address liquidator, uint256 fee)
func (_Collateral *CollateralFilterer) WatchLiquidation(opts *bind.WatchOpts, sink chan<- *CollateralLiquidation, user []common.Address, product []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Liquidation", userRule, productRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralLiquidation)
				if err := _Collateral.contract.UnpackLog(event, "Liquidation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidation is a log parse operation binding the contract event 0x76f1c54646947e8e8df3eab8141bb3d115d07da2520f8e41b208a17d138a8959.
//
// Solidity: event Liquidation(address indexed user, address indexed product, address liquidator, uint256 fee)
func (_Collateral *CollateralFilterer) ParseLiquidation(log types.Log) (*CollateralLiquidation, error) {
	event := new(CollateralLiquidation)
	if err := _Collateral.contract.UnpackLog(event, "Liquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralProductSettleIterator is returned from FilterProductSettle and is used to iterate over the raw logs and unpacked data for ProductSettle events raised by the Collateral contract.
type CollateralProductSettleIterator struct {
	Event *CollateralProductSettle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralProductSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralProductSettle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralProductSettle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralProductSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralProductSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralProductSettle represents a ProductSettle event raised by the Collateral contract.
type CollateralProductSettle struct {
	Product     common.Address
	ProtocolFee *big.Int
	ProductFee  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProductSettle is a free log retrieval operation binding the contract event 0xbb1fb7280f9f082e1804d5dc09e8cd7c89debd6b303542f604bcdf7696d9579d.
//
// Solidity: event ProductSettle(address indexed product, uint256 protocolFee, uint256 productFee)
func (_Collateral *CollateralFilterer) FilterProductSettle(opts *bind.FilterOpts, product []common.Address) (*CollateralProductSettleIterator, error) {

	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "ProductSettle", productRule)
	if err != nil {
		return nil, err
	}
	return &CollateralProductSettleIterator{contract: _Collateral.contract, event: "ProductSettle", logs: logs, sub: sub}, nil
}

// WatchProductSettle is a free log subscription operation binding the contract event 0xbb1fb7280f9f082e1804d5dc09e8cd7c89debd6b303542f604bcdf7696d9579d.
//
// Solidity: event ProductSettle(address indexed product, uint256 protocolFee, uint256 productFee)
func (_Collateral *CollateralFilterer) WatchProductSettle(opts *bind.WatchOpts, sink chan<- *CollateralProductSettle, product []common.Address) (event.Subscription, error) {

	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "ProductSettle", productRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralProductSettle)
				if err := _Collateral.contract.UnpackLog(event, "ProductSettle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProductSettle is a log parse operation binding the contract event 0xbb1fb7280f9f082e1804d5dc09e8cd7c89debd6b303542f604bcdf7696d9579d.
//
// Solidity: event ProductSettle(address indexed product, uint256 protocolFee, uint256 productFee)
func (_Collateral *CollateralFilterer) ParseProductSettle(log types.Log) (*CollateralProductSettle, error) {
	event := new(CollateralProductSettle)
	if err := _Collateral.contract.UnpackLog(event, "ProductSettle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralShortfallResolutionIterator is returned from FilterShortfallResolution and is used to iterate over the raw logs and unpacked data for ShortfallResolution events raised by the Collateral contract.
type CollateralShortfallResolutionIterator struct {
	Event *CollateralShortfallResolution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralShortfallResolutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralShortfallResolution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralShortfallResolution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralShortfallResolutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralShortfallResolutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralShortfallResolution represents a ShortfallResolution event raised by the Collateral contract.
type CollateralShortfallResolution struct {
	Product common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShortfallResolution is a free log retrieval operation binding the contract event 0x8ae4dab3d718e6b7019e16378795bb41b3046b9117483430a39f7af4ab9bb893.
//
// Solidity: event ShortfallResolution(address indexed product, uint256 amount)
func (_Collateral *CollateralFilterer) FilterShortfallResolution(opts *bind.FilterOpts, product []common.Address) (*CollateralShortfallResolutionIterator, error) {

	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "ShortfallResolution", productRule)
	if err != nil {
		return nil, err
	}
	return &CollateralShortfallResolutionIterator{contract: _Collateral.contract, event: "ShortfallResolution", logs: logs, sub: sub}, nil
}

// WatchShortfallResolution is a free log subscription operation binding the contract event 0x8ae4dab3d718e6b7019e16378795bb41b3046b9117483430a39f7af4ab9bb893.
//
// Solidity: event ShortfallResolution(address indexed product, uint256 amount)
func (_Collateral *CollateralFilterer) WatchShortfallResolution(opts *bind.WatchOpts, sink chan<- *CollateralShortfallResolution, product []common.Address) (event.Subscription, error) {

	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "ShortfallResolution", productRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralShortfallResolution)
				if err := _Collateral.contract.UnpackLog(event, "ShortfallResolution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShortfallResolution is a log parse operation binding the contract event 0x8ae4dab3d718e6b7019e16378795bb41b3046b9117483430a39f7af4ab9bb893.
//
// Solidity: event ShortfallResolution(address indexed product, uint256 amount)
func (_Collateral *CollateralFilterer) ParseShortfallResolution(log types.Log) (*CollateralShortfallResolution, error) {
	event := new(CollateralShortfallResolution)
	if err := _Collateral.contract.UnpackLog(event, "ShortfallResolution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Collateral contract.
type CollateralWithdrawalIterator struct {
	Event *CollateralWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralWithdrawal represents a Withdrawal event raised by the Collateral contract.
type CollateralWithdrawal struct {
	User    common.Address
	Product common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed user, address indexed product, uint256 amount)
func (_Collateral *CollateralFilterer) FilterWithdrawal(opts *bind.FilterOpts, user []common.Address, product []common.Address) (*CollateralWithdrawalIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Withdrawal", userRule, productRule)
	if err != nil {
		return nil, err
	}
	return &CollateralWithdrawalIterator{contract: _Collateral.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed user, address indexed product, uint256 amount)
func (_Collateral *CollateralFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *CollateralWithdrawal, user []common.Address, product []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Withdrawal", userRule, productRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralWithdrawal)
				if err := _Collateral.contract.UnpackLog(event, "Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawal is a log parse operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed user, address indexed product, uint256 amount)
func (_Collateral *CollateralFilterer) ParseWithdrawal(log types.Log) (*CollateralWithdrawal, error) {
	event := new(CollateralWithdrawal)
	if err := _Collateral.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
