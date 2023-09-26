// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

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

// SwapMetaData contains all meta data concerning the Swap contract.
var SwapMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_sold\",\"type\":\"address\",\"indexed\":false},{\"name\":\"token_bought\",\"type\":\"address\",\"indexed\":false},{\"name\":\"amount_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"amount_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_address_provider\",\"type\":\"address\"},{\"name\":\"_calculator\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_with_best_rate\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1019563733},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_with_best_rate\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1019563733},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":427142},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":427142},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_multiple\",\"inputs\":[{\"name\":\"_route\",\"type\":\"address[9]\"},{\"name\":\"_swap_params\",\"type\":\"uint256[3][4]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":313422},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_multiple\",\"inputs\":[{\"name\":\"_route\",\"type\":\"address[9]\"},{\"name\":\"_swap_params\",\"type\":\"uint256[3][4]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"},{\"name\":\"_pools\",\"type\":\"address[4]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":313422},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_multiple\",\"inputs\":[{\"name\":\"_route\",\"type\":\"address[9]\"},{\"name\":\"_swap_params\",\"type\":\"uint256[3][4]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\"},{\"name\":\"_pools\",\"type\":\"address[4]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":313422},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_best_rate\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3002213116},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_best_rate\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_exclude_pools\",\"type\":\"address[8]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3002213116},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_exchange_amount\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":30596},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_input_amount\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":34701},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_exchange_amounts\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amounts\",\"type\":\"uint256[100]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[100]\"}],\"gas\":38286},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_exchange_multiple_amount\",\"inputs\":[{\"name\":\"_route\",\"type\":\"address[9]\"},{\"name\":\"_swap_params\",\"type\":\"uint256[3][4]\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":21334},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_exchange_multiple_amount\",\"inputs\":[{\"name\":\"_route\",\"type\":\"address[9]\"},{\"name\":\"_swap_params\",\"type\":\"uint256[3][4]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_pools\",\"type\":\"address[4]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":21334},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_calculator\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":5215},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_registry_address\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":115368},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_calculator\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_calculator\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":40695},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_default_calculator\",\"inputs\":[{\"name\":\"_calculator\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":40459},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_balance\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":41823},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_killed\",\"inputs\":[{\"name\":\"_is_killed\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":40519},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2970},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory_registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3000},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"crypto_registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3030},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"default_calculator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3060},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_killed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3090}]",
}

// SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapMetaData.ABI instead.
var SwapABI = SwapMetaData.ABI

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// CryptoRegistry is a free data retrieval call binding the contract method 0xf3b8f829.
//
// Solidity: function crypto_registry() view returns(address)
func (_Swap *SwapCaller) CryptoRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "crypto_registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CryptoRegistry is a free data retrieval call binding the contract method 0xf3b8f829.
//
// Solidity: function crypto_registry() view returns(address)
func (_Swap *SwapSession) CryptoRegistry() (common.Address, error) {
	return _Swap.Contract.CryptoRegistry(&_Swap.CallOpts)
}

// CryptoRegistry is a free data retrieval call binding the contract method 0xf3b8f829.
//
// Solidity: function crypto_registry() view returns(address)
func (_Swap *SwapCallerSession) CryptoRegistry() (common.Address, error) {
	return _Swap.Contract.CryptoRegistry(&_Swap.CallOpts)
}

// DefaultCalculator is a free data retrieval call binding the contract method 0x3b359fc8.
//
// Solidity: function default_calculator() view returns(address)
func (_Swap *SwapCaller) DefaultCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "default_calculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultCalculator is a free data retrieval call binding the contract method 0x3b359fc8.
//
// Solidity: function default_calculator() view returns(address)
func (_Swap *SwapSession) DefaultCalculator() (common.Address, error) {
	return _Swap.Contract.DefaultCalculator(&_Swap.CallOpts)
}

// DefaultCalculator is a free data retrieval call binding the contract method 0x3b359fc8.
//
// Solidity: function default_calculator() view returns(address)
func (_Swap *SwapCallerSession) DefaultCalculator() (common.Address, error) {
	return _Swap.Contract.DefaultCalculator(&_Swap.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0xf7cbf4c6.
//
// Solidity: function factory_registry() view returns(address)
func (_Swap *SwapCaller) FactoryRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "factory_registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryRegistry is a free data retrieval call binding the contract method 0xf7cbf4c6.
//
// Solidity: function factory_registry() view returns(address)
func (_Swap *SwapSession) FactoryRegistry() (common.Address, error) {
	return _Swap.Contract.FactoryRegistry(&_Swap.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0xf7cbf4c6.
//
// Solidity: function factory_registry() view returns(address)
func (_Swap *SwapCallerSession) FactoryRegistry() (common.Address, error) {
	return _Swap.Contract.FactoryRegistry(&_Swap.CallOpts)
}

// GetBestRate is a free data retrieval call binding the contract method 0x4e21df75.
//
// Solidity: function get_best_rate(address _from, address _to, uint256 _amount) view returns(address, uint256)
func (_Swap *SwapCaller) GetBestRate(opts *bind.CallOpts, _from common.Address, _to common.Address, _amount *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "get_best_rate", _from, _to, _amount)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBestRate is a free data retrieval call binding the contract method 0x4e21df75.
//
// Solidity: function get_best_rate(address _from, address _to, uint256 _amount) view returns(address, uint256)
func (_Swap *SwapSession) GetBestRate(_from common.Address, _to common.Address, _amount *big.Int) (common.Address, *big.Int, error) {
	return _Swap.Contract.GetBestRate(&_Swap.CallOpts, _from, _to, _amount)
}

// GetBestRate is a free data retrieval call binding the contract method 0x4e21df75.
//
// Solidity: function get_best_rate(address _from, address _to, uint256 _amount) view returns(address, uint256)
func (_Swap *SwapCallerSession) GetBestRate(_from common.Address, _to common.Address, _amount *big.Int) (common.Address, *big.Int, error) {
	return _Swap.Contract.GetBestRate(&_Swap.CallOpts, _from, _to, _amount)
}

// GetBestRate0 is a free data retrieval call binding the contract method 0x488de9af.
//
// Solidity: function get_best_rate(address _from, address _to, uint256 _amount, address[8] _exclude_pools) view returns(address, uint256)
func (_Swap *SwapCaller) GetBestRate0(opts *bind.CallOpts, _from common.Address, _to common.Address, _amount *big.Int, _exclude_pools [8]common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "get_best_rate0", _from, _to, _amount, _exclude_pools)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBestRate0 is a free data retrieval call binding the contract method 0x488de9af.
//
// Solidity: function get_best_rate(address _from, address _to, uint256 _amount, address[8] _exclude_pools) view returns(address, uint256)
func (_Swap *SwapSession) GetBestRate0(_from common.Address, _to common.Address, _amount *big.Int, _exclude_pools [8]common.Address) (common.Address, *big.Int, error) {
	return _Swap.Contract.GetBestRate0(&_Swap.CallOpts, _from, _to, _amount, _exclude_pools)
}

// GetBestRate0 is a free data retrieval call binding the contract method 0x488de9af.
//
// Solidity: function get_best_rate(address _from, address _to, uint256 _amount, address[8] _exclude_pools) view returns(address, uint256)
func (_Swap *SwapCallerSession) GetBestRate0(_from common.Address, _to common.Address, _amount *big.Int, _exclude_pools [8]common.Address) (common.Address, *big.Int, error) {
	return _Swap.Contract.GetBestRate0(&_Swap.CallOpts, _from, _to, _amount, _exclude_pools)
}

// GetCalculator is a free data retrieval call binding the contract method 0x5d7dc825.
//
// Solidity: function get_calculator(address _pool) view returns(address)
func (_Swap *SwapCaller) GetCalculator(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "get_calculator", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCalculator is a free data retrieval call binding the contract method 0x5d7dc825.
//
// Solidity: function get_calculator(address _pool) view returns(address)
func (_Swap *SwapSession) GetCalculator(_pool common.Address) (common.Address, error) {
	return _Swap.Contract.GetCalculator(&_Swap.CallOpts, _pool)
}

// GetCalculator is a free data retrieval call binding the contract method 0x5d7dc825.
//
// Solidity: function get_calculator(address _pool) view returns(address)
func (_Swap *SwapCallerSession) GetCalculator(_pool common.Address) (common.Address, error) {
	return _Swap.Contract.GetCalculator(&_Swap.CallOpts, _pool)
}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x3973e834.
//
// Solidity: function get_exchange_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swap *SwapCaller) GetExchangeAmount(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "get_exchange_amount", _pool, _from, _to, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x3973e834.
//
// Solidity: function get_exchange_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swap *SwapSession) GetExchangeAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetExchangeAmount(&_Swap.CallOpts, _pool, _from, _to, _amount)
}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x3973e834.
//
// Solidity: function get_exchange_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swap *SwapCallerSession) GetExchangeAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetExchangeAmount(&_Swap.CallOpts, _pool, _from, _to, _amount)
}

// GetExchangeAmounts is a free data retrieval call binding the contract method 0x4be9ae42.
//
// Solidity: function get_exchange_amounts(address _pool, address _from, address _to, uint256[100] _amounts) view returns(uint256[100])
func (_Swap *SwapCaller) GetExchangeAmounts(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address, _amounts [100]*big.Int) ([100]*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "get_exchange_amounts", _pool, _from, _to, _amounts)

	if err != nil {
		return *new([100]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([100]*big.Int)).(*[100]*big.Int)

	return out0, err

}

// GetExchangeAmounts is a free data retrieval call binding the contract method 0x4be9ae42.
//
// Solidity: function get_exchange_amounts(address _pool, address _from, address _to, uint256[100] _amounts) view returns(uint256[100])
func (_Swap *SwapSession) GetExchangeAmounts(_pool common.Address, _from common.Address, _to common.Address, _amounts [100]*big.Int) ([100]*big.Int, error) {
	return _Swap.Contract.GetExchangeAmounts(&_Swap.CallOpts, _pool, _from, _to, _amounts)
}

// GetExchangeAmounts is a free data retrieval call binding the contract method 0x4be9ae42.
//
// Solidity: function get_exchange_amounts(address _pool, address _from, address _to, uint256[100] _amounts) view returns(uint256[100])
func (_Swap *SwapCallerSession) GetExchangeAmounts(_pool common.Address, _from common.Address, _to common.Address, _amounts [100]*big.Int) ([100]*big.Int, error) {
	return _Swap.Contract.GetExchangeAmounts(&_Swap.CallOpts, _pool, _from, _to, _amounts)
}

// GetExchangeMultipleAmount is a free data retrieval call binding the contract method 0x7b3d22cf.
//
// Solidity: function get_exchange_multiple_amount(address[9] _route, uint256[3][4] _swap_params, uint256 _amount) view returns(uint256)
func (_Swap *SwapCaller) GetExchangeMultipleAmount(opts *bind.CallOpts, _route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "get_exchange_multiple_amount", _route, _swap_params, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeMultipleAmount is a free data retrieval call binding the contract method 0x7b3d22cf.
//
// Solidity: function get_exchange_multiple_amount(address[9] _route, uint256[3][4] _swap_params, uint256 _amount) view returns(uint256)
func (_Swap *SwapSession) GetExchangeMultipleAmount(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetExchangeMultipleAmount(&_Swap.CallOpts, _route, _swap_params, _amount)
}

// GetExchangeMultipleAmount is a free data retrieval call binding the contract method 0x7b3d22cf.
//
// Solidity: function get_exchange_multiple_amount(address[9] _route, uint256[3][4] _swap_params, uint256 _amount) view returns(uint256)
func (_Swap *SwapCallerSession) GetExchangeMultipleAmount(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetExchangeMultipleAmount(&_Swap.CallOpts, _route, _swap_params, _amount)
}

// GetExchangeMultipleAmount0 is a free data retrieval call binding the contract method 0xe6eabf23.
//
// Solidity: function get_exchange_multiple_amount(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, address[4] _pools) view returns(uint256)
func (_Swap *SwapCaller) GetExchangeMultipleAmount0(opts *bind.CallOpts, _route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _pools [4]common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "get_exchange_multiple_amount0", _route, _swap_params, _amount, _pools)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeMultipleAmount0 is a free data retrieval call binding the contract method 0xe6eabf23.
//
// Solidity: function get_exchange_multiple_amount(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, address[4] _pools) view returns(uint256)
func (_Swap *SwapSession) GetExchangeMultipleAmount0(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _pools [4]common.Address) (*big.Int, error) {
	return _Swap.Contract.GetExchangeMultipleAmount0(&_Swap.CallOpts, _route, _swap_params, _amount, _pools)
}

// GetExchangeMultipleAmount0 is a free data retrieval call binding the contract method 0xe6eabf23.
//
// Solidity: function get_exchange_multiple_amount(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, address[4] _pools) view returns(uint256)
func (_Swap *SwapCallerSession) GetExchangeMultipleAmount0(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _pools [4]common.Address) (*big.Int, error) {
	return _Swap.Contract.GetExchangeMultipleAmount0(&_Swap.CallOpts, _route, _swap_params, _amount, _pools)
}

// GetInputAmount is a free data retrieval call binding the contract method 0x7fa5a654.
//
// Solidity: function get_input_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swap *SwapCaller) GetInputAmount(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "get_input_amount", _pool, _from, _to, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInputAmount is a free data retrieval call binding the contract method 0x7fa5a654.
//
// Solidity: function get_input_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swap *SwapSession) GetInputAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetInputAmount(&_Swap.CallOpts, _pool, _from, _to, _amount)
}

// GetInputAmount is a free data retrieval call binding the contract method 0x7fa5a654.
//
// Solidity: function get_input_amount(address _pool, address _from, address _to, uint256 _amount) view returns(uint256)
func (_Swap *SwapCallerSession) GetInputAmount(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetInputAmount(&_Swap.CallOpts, _pool, _from, _to, _amount)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Swap *SwapCaller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Swap *SwapSession) IsKilled() (bool, error) {
	return _Swap.Contract.IsKilled(&_Swap.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Swap *SwapCallerSession) IsKilled() (bool, error) {
	return _Swap.Contract.IsKilled(&_Swap.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Swap *SwapCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Swap *SwapSession) Registry() (common.Address, error) {
	return _Swap.Contract.Registry(&_Swap.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Swap *SwapCallerSession) Registry() (common.Address, error) {
	return _Swap.Contract.Registry(&_Swap.CallOpts)
}

// ClaimBalance is a paid mutator transaction binding the contract method 0x752d53c6.
//
// Solidity: function claim_balance(address _token) returns(bool)
func (_Swap *SwapTransactor) ClaimBalance(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "claim_balance", _token)
}

// ClaimBalance is a paid mutator transaction binding the contract method 0x752d53c6.
//
// Solidity: function claim_balance(address _token) returns(bool)
func (_Swap *SwapSession) ClaimBalance(_token common.Address) (*types.Transaction, error) {
	return _Swap.Contract.ClaimBalance(&_Swap.TransactOpts, _token)
}

// ClaimBalance is a paid mutator transaction binding the contract method 0x752d53c6.
//
// Solidity: function claim_balance(address _token) returns(bool)
func (_Swap *SwapTransactorSession) ClaimBalance(_token common.Address) (*types.Transaction, error) {
	return _Swap.Contract.ClaimBalance(&_Swap.TransactOpts, _token)
}

// Exchange is a paid mutator transaction binding the contract method 0x4798ce5b.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swap *SwapTransactor) Exchange(opts *bind.TransactOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "exchange", _pool, _from, _to, _amount, _expected)
}

// Exchange is a paid mutator transaction binding the contract method 0x4798ce5b.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swap *SwapSession) Exchange(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Exchange(&_Swap.TransactOpts, _pool, _from, _to, _amount, _expected)
}

// Exchange is a paid mutator transaction binding the contract method 0x4798ce5b.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swap *SwapTransactorSession) Exchange(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Exchange(&_Swap.TransactOpts, _pool, _from, _to, _amount, _expected)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x1a4c1ca3.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swap *SwapTransactor) Exchange0(opts *bind.TransactOpts, _pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "exchange0", _pool, _from, _to, _amount, _expected, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x1a4c1ca3.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swap *SwapSession) Exchange0(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swap.Contract.Exchange0(&_Swap.TransactOpts, _pool, _from, _to, _amount, _expected, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x1a4c1ca3.
//
// Solidity: function exchange(address _pool, address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swap *SwapTransactorSession) Exchange0(_pool common.Address, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swap.Contract.Exchange0(&_Swap.TransactOpts, _pool, _from, _to, _amount, _expected, _receiver)
}

// ExchangeMultiple is a paid mutator transaction binding the contract method 0x353ca424.
//
// Solidity: function exchange_multiple(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swap *SwapTransactor) ExchangeMultiple(opts *bind.TransactOpts, _route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "exchange_multiple", _route, _swap_params, _amount, _expected)
}

// ExchangeMultiple is a paid mutator transaction binding the contract method 0x353ca424.
//
// Solidity: function exchange_multiple(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swap *SwapSession) ExchangeMultiple(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeMultiple(&_Swap.TransactOpts, _route, _swap_params, _amount, _expected)
}

// ExchangeMultiple is a paid mutator transaction binding the contract method 0x353ca424.
//
// Solidity: function exchange_multiple(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swap *SwapTransactorSession) ExchangeMultiple(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeMultiple(&_Swap.TransactOpts, _route, _swap_params, _amount, _expected)
}

// ExchangeMultiple0 is a paid mutator transaction binding the contract method 0x9db4f7aa.
//
// Solidity: function exchange_multiple(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, uint256 _expected, address[4] _pools) payable returns(uint256)
func (_Swap *SwapTransactor) ExchangeMultiple0(opts *bind.TransactOpts, _route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _expected *big.Int, _pools [4]common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "exchange_multiple0", _route, _swap_params, _amount, _expected, _pools)
}

// ExchangeMultiple0 is a paid mutator transaction binding the contract method 0x9db4f7aa.
//
// Solidity: function exchange_multiple(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, uint256 _expected, address[4] _pools) payable returns(uint256)
func (_Swap *SwapSession) ExchangeMultiple0(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _expected *big.Int, _pools [4]common.Address) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeMultiple0(&_Swap.TransactOpts, _route, _swap_params, _amount, _expected, _pools)
}

// ExchangeMultiple0 is a paid mutator transaction binding the contract method 0x9db4f7aa.
//
// Solidity: function exchange_multiple(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, uint256 _expected, address[4] _pools) payable returns(uint256)
func (_Swap *SwapTransactorSession) ExchangeMultiple0(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _expected *big.Int, _pools [4]common.Address) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeMultiple0(&_Swap.TransactOpts, _route, _swap_params, _amount, _expected, _pools)
}

// ExchangeMultiple1 is a paid mutator transaction binding the contract method 0x0651cb35.
//
// Solidity: function exchange_multiple(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, uint256 _expected, address[4] _pools, address _receiver) payable returns(uint256)
func (_Swap *SwapTransactor) ExchangeMultiple1(opts *bind.TransactOpts, _route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _expected *big.Int, _pools [4]common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "exchange_multiple1", _route, _swap_params, _amount, _expected, _pools, _receiver)
}

// ExchangeMultiple1 is a paid mutator transaction binding the contract method 0x0651cb35.
//
// Solidity: function exchange_multiple(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, uint256 _expected, address[4] _pools, address _receiver) payable returns(uint256)
func (_Swap *SwapSession) ExchangeMultiple1(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _expected *big.Int, _pools [4]common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeMultiple1(&_Swap.TransactOpts, _route, _swap_params, _amount, _expected, _pools, _receiver)
}

// ExchangeMultiple1 is a paid mutator transaction binding the contract method 0x0651cb35.
//
// Solidity: function exchange_multiple(address[9] _route, uint256[3][4] _swap_params, uint256 _amount, uint256 _expected, address[4] _pools, address _receiver) payable returns(uint256)
func (_Swap *SwapTransactorSession) ExchangeMultiple1(_route [9]common.Address, _swap_params [4][3]*big.Int, _amount *big.Int, _expected *big.Int, _pools [4]common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeMultiple1(&_Swap.TransactOpts, _route, _swap_params, _amount, _expected, _pools, _receiver)
}

// ExchangeWithBestRate is a paid mutator transaction binding the contract method 0x10e5e303.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swap *SwapTransactor) ExchangeWithBestRate(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "exchange_with_best_rate", _from, _to, _amount, _expected)
}

// ExchangeWithBestRate is a paid mutator transaction binding the contract method 0x10e5e303.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swap *SwapSession) ExchangeWithBestRate(_from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeWithBestRate(&_Swap.TransactOpts, _from, _to, _amount, _expected)
}

// ExchangeWithBestRate is a paid mutator transaction binding the contract method 0x10e5e303.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected) payable returns(uint256)
func (_Swap *SwapTransactorSession) ExchangeWithBestRate(_from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeWithBestRate(&_Swap.TransactOpts, _from, _to, _amount, _expected)
}

// ExchangeWithBestRate0 is a paid mutator transaction binding the contract method 0x9f69a6a6.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swap *SwapTransactor) ExchangeWithBestRate0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "exchange_with_best_rate0", _from, _to, _amount, _expected, _receiver)
}

// ExchangeWithBestRate0 is a paid mutator transaction binding the contract method 0x9f69a6a6.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swap *SwapSession) ExchangeWithBestRate0(_from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeWithBestRate0(&_Swap.TransactOpts, _from, _to, _amount, _expected, _receiver)
}

// ExchangeWithBestRate0 is a paid mutator transaction binding the contract method 0x9f69a6a6.
//
// Solidity: function exchange_with_best_rate(address _from, address _to, uint256 _amount, uint256 _expected, address _receiver) payable returns(uint256)
func (_Swap *SwapTransactorSession) ExchangeWithBestRate0(_from common.Address, _to common.Address, _amount *big.Int, _expected *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Swap.Contract.ExchangeWithBestRate0(&_Swap.TransactOpts, _from, _to, _amount, _expected, _receiver)
}

// SetCalculator is a paid mutator transaction binding the contract method 0x188c7ee5.
//
// Solidity: function set_calculator(address _pool, address _calculator) returns(bool)
func (_Swap *SwapTransactor) SetCalculator(opts *bind.TransactOpts, _pool common.Address, _calculator common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "set_calculator", _pool, _calculator)
}

// SetCalculator is a paid mutator transaction binding the contract method 0x188c7ee5.
//
// Solidity: function set_calculator(address _pool, address _calculator) returns(bool)
func (_Swap *SwapSession) SetCalculator(_pool common.Address, _calculator common.Address) (*types.Transaction, error) {
	return _Swap.Contract.SetCalculator(&_Swap.TransactOpts, _pool, _calculator)
}

// SetCalculator is a paid mutator transaction binding the contract method 0x188c7ee5.
//
// Solidity: function set_calculator(address _pool, address _calculator) returns(bool)
func (_Swap *SwapTransactorSession) SetCalculator(_pool common.Address, _calculator common.Address) (*types.Transaction, error) {
	return _Swap.Contract.SetCalculator(&_Swap.TransactOpts, _pool, _calculator)
}

// SetDefaultCalculator is a paid mutator transaction binding the contract method 0xda3fb2ab.
//
// Solidity: function set_default_calculator(address _calculator) returns(bool)
func (_Swap *SwapTransactor) SetDefaultCalculator(opts *bind.TransactOpts, _calculator common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "set_default_calculator", _calculator)
}

// SetDefaultCalculator is a paid mutator transaction binding the contract method 0xda3fb2ab.
//
// Solidity: function set_default_calculator(address _calculator) returns(bool)
func (_Swap *SwapSession) SetDefaultCalculator(_calculator common.Address) (*types.Transaction, error) {
	return _Swap.Contract.SetDefaultCalculator(&_Swap.TransactOpts, _calculator)
}

// SetDefaultCalculator is a paid mutator transaction binding the contract method 0xda3fb2ab.
//
// Solidity: function set_default_calculator(address _calculator) returns(bool)
func (_Swap *SwapTransactorSession) SetDefaultCalculator(_calculator common.Address) (*types.Transaction, error) {
	return _Swap.Contract.SetDefaultCalculator(&_Swap.TransactOpts, _calculator)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns(bool)
func (_Swap *SwapTransactor) SetKilled(opts *bind.TransactOpts, _is_killed bool) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "set_killed", _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns(bool)
func (_Swap *SwapSession) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _Swap.Contract.SetKilled(&_Swap.TransactOpts, _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns(bool)
func (_Swap *SwapTransactorSession) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _Swap.Contract.SetKilled(&_Swap.TransactOpts, _is_killed)
}

// UpdateRegistryAddress is a paid mutator transaction binding the contract method 0x4bbc5b1f.
//
// Solidity: function update_registry_address() returns(bool)
func (_Swap *SwapTransactor) UpdateRegistryAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "update_registry_address")
}

// UpdateRegistryAddress is a paid mutator transaction binding the contract method 0x4bbc5b1f.
//
// Solidity: function update_registry_address() returns(bool)
func (_Swap *SwapSession) UpdateRegistryAddress() (*types.Transaction, error) {
	return _Swap.Contract.UpdateRegistryAddress(&_Swap.TransactOpts)
}

// UpdateRegistryAddress is a paid mutator transaction binding the contract method 0x4bbc5b1f.
//
// Solidity: function update_registry_address() returns(bool)
func (_Swap *SwapTransactorSession) UpdateRegistryAddress() (*types.Transaction, error) {
	return _Swap.Contract.UpdateRegistryAddress(&_Swap.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Swap *SwapTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Swap.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Swap *SwapSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Swap.Contract.Fallback(&_Swap.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Swap *SwapTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Swap.Contract.Fallback(&_Swap.TransactOpts, calldata)
}

// SwapTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Swap contract.
type SwapTokenExchangeIterator struct {
	Event *SwapTokenExchange // Event containing the contract specifics and raw log

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
func (it *SwapTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapTokenExchange)
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
		it.Event = new(SwapTokenExchange)
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
func (it *SwapTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapTokenExchange represents a TokenExchange event raised by the Swap contract.
type SwapTokenExchange struct {
	Buyer        common.Address
	Receiver     common.Address
	Pool         common.Address
	TokenSold    common.Address
	TokenBought  common.Address
	AmountSold   *big.Int
	AmountBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0xbd3eb7bcfdd1721a4eb4f00d0df3ed91bd6f17222f82b2d7bce519d8cab3fe46.
//
// Solidity: event TokenExchange(address indexed buyer, address indexed receiver, address indexed pool, address token_sold, address token_bought, uint256 amount_sold, uint256 amount_bought)
func (_Swap *SwapFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address, receiver []common.Address, pool []common.Address) (*SwapTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "TokenExchange", buyerRule, receiverRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &SwapTokenExchangeIterator{contract: _Swap.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0xbd3eb7bcfdd1721a4eb4f00d0df3ed91bd6f17222f82b2d7bce519d8cab3fe46.
//
// Solidity: event TokenExchange(address indexed buyer, address indexed receiver, address indexed pool, address token_sold, address token_bought, uint256 amount_sold, uint256 amount_bought)
func (_Swap *SwapFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *SwapTokenExchange, buyer []common.Address, receiver []common.Address, pool []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "TokenExchange", buyerRule, receiverRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapTokenExchange)
				if err := _Swap.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0xbd3eb7bcfdd1721a4eb4f00d0df3ed91bd6f17222f82b2d7bce519d8cab3fe46.
//
// Solidity: event TokenExchange(address indexed buyer, address indexed receiver, address indexed pool, address token_sold, address token_bought, uint256 amount_sold, uint256 amount_bought)
func (_Swap *SwapFilterer) ParseTokenExchange(log types.Log) (*SwapTokenExchange, error) {
	event := new(SwapTokenExchange)
	if err := _Swap.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
