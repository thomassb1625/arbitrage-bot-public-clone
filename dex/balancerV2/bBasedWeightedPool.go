// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancerV2

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


// BBasedWeightedPoolMetaData contains all meta data concerning the BBasedWeightedPool contract.
var BBasedWeightedPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"RecoveryModeStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInvariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNormalizedWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"contractIProtocolFeesCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getScalingFactors\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inRecoveryMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onExitPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onJoinPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIPoolSwapStructs.SwapRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenOut\",\"type\":\"uint256\"}],\"name\":\"onSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryJoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setSwapFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BBasedWeightedPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use BBasedWeightedPoolMetaData.ABI instead.
var BBasedWeightedPoolABI = BBasedWeightedPoolMetaData.ABI

// BBasedWeightedPool is an auto generated Go binding around an Ethereum contract.
type BBasedWeightedPool struct {
	BBasedWeightedPoolCaller     // Read-only binding to the contract
	BBasedWeightedPoolTransactor // Write-only binding to the contract
	BBasedWeightedPoolFilterer   // Log filterer for contract events
}

// BBasedWeightedPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BBasedWeightedPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BBasedWeightedPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BBasedWeightedPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BBasedWeightedPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BBasedWeightedPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BBasedWeightedPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BBasedWeightedPoolSession struct {
	Contract     *BBasedWeightedPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BBasedWeightedPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BBasedWeightedPoolCallerSession struct {
	Contract *BBasedWeightedPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BBasedWeightedPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BBasedWeightedPoolTransactorSession struct {
	Contract     *BBasedWeightedPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BBasedWeightedPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BBasedWeightedPoolRaw struct {
	Contract *BBasedWeightedPool // Generic contract binding to access the raw methods on
}

// BBasedWeightedPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BBasedWeightedPoolCallerRaw struct {
	Contract *BBasedWeightedPoolCaller // Generic read-only contract binding to access the raw methods on
}

// BBasedWeightedPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BBasedWeightedPoolTransactorRaw struct {
	Contract *BBasedWeightedPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBBasedWeightedPool creates a new instance of BBasedWeightedPool, bound to a specific deployed contract.
func NewBBasedWeightedPool(address common.Address, backend bind.ContractBackend) (*BBasedWeightedPool, error) {
	contract, err := bindBBasedWeightedPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BBasedWeightedPool{BBasedWeightedPoolCaller: BBasedWeightedPoolCaller{contract: contract}, BBasedWeightedPoolTransactor: BBasedWeightedPoolTransactor{contract: contract}, BBasedWeightedPoolFilterer: BBasedWeightedPoolFilterer{contract: contract}}, nil
}

// NewBBasedWeightedPoolCaller creates a new read-only instance of BBasedWeightedPool, bound to a specific deployed contract.
func NewBBasedWeightedPoolCaller(address common.Address, caller bind.ContractCaller) (*BBasedWeightedPoolCaller, error) {
	contract, err := bindBBasedWeightedPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BBasedWeightedPoolCaller{contract: contract}, nil
}

// NewBBasedWeightedPoolTransactor creates a new write-only instance of BBasedWeightedPool, bound to a specific deployed contract.
func NewBBasedWeightedPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BBasedWeightedPoolTransactor, error) {
	contract, err := bindBBasedWeightedPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BBasedWeightedPoolTransactor{contract: contract}, nil
}

// NewBBasedWeightedPoolFilterer creates a new log filterer instance of BBasedWeightedPool, bound to a specific deployed contract.
func NewBBasedWeightedPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BBasedWeightedPoolFilterer, error) {
	contract, err := bindBBasedWeightedPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BBasedWeightedPoolFilterer{contract: contract}, nil
}

// bindBBasedWeightedPool binds a generic wrapper to an already deployed contract.
func bindBBasedWeightedPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BBasedWeightedPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BBasedWeightedPool *BBasedWeightedPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BBasedWeightedPool.Contract.BBasedWeightedPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BBasedWeightedPool *BBasedWeightedPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.BBasedWeightedPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BBasedWeightedPool *BBasedWeightedPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.BBasedWeightedPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BBasedWeightedPool *BBasedWeightedPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BBasedWeightedPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BBasedWeightedPool.Contract.DOMAINSEPARATOR(&_BBasedWeightedPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BBasedWeightedPool.Contract.DOMAINSEPARATOR(&_BBasedWeightedPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BBasedWeightedPool.Contract.Allowance(&_BBasedWeightedPool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BBasedWeightedPool.Contract.Allowance(&_BBasedWeightedPool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BBasedWeightedPool.Contract.BalanceOf(&_BBasedWeightedPool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BBasedWeightedPool.Contract.BalanceOf(&_BBasedWeightedPool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Decimals() (uint8, error) {
	return _BBasedWeightedPool.Contract.Decimals(&_BBasedWeightedPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) Decimals() (uint8, error) {
	return _BBasedWeightedPool.Contract.Decimals(&_BBasedWeightedPool.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BBasedWeightedPool.Contract.GetActionId(&_BBasedWeightedPool.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BBasedWeightedPool.Contract.GetActionId(&_BBasedWeightedPool.CallOpts, selector)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetAuthorizer() (common.Address, error) {
	return _BBasedWeightedPool.Contract.GetAuthorizer(&_BBasedWeightedPool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetAuthorizer() (common.Address, error) {
	return _BBasedWeightedPool.Contract.GetAuthorizer(&_BBasedWeightedPool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetDomainSeparator() ([32]byte, error) {
	return _BBasedWeightedPool.Contract.GetDomainSeparator(&_BBasedWeightedPool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _BBasedWeightedPool.Contract.GetDomainSeparator(&_BBasedWeightedPool.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetInvariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getInvariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetInvariant() (*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetInvariant(&_BBasedWeightedPool.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetInvariant() (*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetInvariant(&_BBasedWeightedPool.CallOpts)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetNextNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getNextNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetNextNonce(&_BBasedWeightedPool.CallOpts, account)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetNextNonce(&_BBasedWeightedPool.CallOpts, account)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetNormalizedWeights(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getNormalizedWeights")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetNormalizedWeights(&_BBasedWeightedPool.CallOpts)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetNormalizedWeights(&_BBasedWeightedPool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetOwner() (common.Address, error) {
	return _BBasedWeightedPool.Contract.GetOwner(&_BBasedWeightedPool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetOwner() (common.Address, error) {
	return _BBasedWeightedPool.Contract.GetOwner(&_BBasedWeightedPool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getPausedState")

	outstruct := new(struct {
		Paused              bool
		PauseWindowEndTime  *big.Int
		BufferPeriodEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PauseWindowEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodEndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BBasedWeightedPool.Contract.GetPausedState(&_BBasedWeightedPool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BBasedWeightedPool.Contract.GetPausedState(&_BBasedWeightedPool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetPoolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getPoolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetPoolId() ([32]byte, error) {
	return _BBasedWeightedPool.Contract.GetPoolId(&_BBasedWeightedPool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetPoolId() ([32]byte, error) {
	return _BBasedWeightedPool.Contract.GetPoolId(&_BBasedWeightedPool.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BBasedWeightedPool.Contract.GetProtocolFeesCollector(&_BBasedWeightedPool.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BBasedWeightedPool.Contract.GetProtocolFeesCollector(&_BBasedWeightedPool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetScalingFactors(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getScalingFactors")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetScalingFactors() ([]*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetScalingFactors(&_BBasedWeightedPool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetScalingFactors() ([]*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetScalingFactors(&_BBasedWeightedPool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetSwapFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getSwapFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetSwapFeePercentage(&_BBasedWeightedPool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BBasedWeightedPool.Contract.GetSwapFeePercentage(&_BBasedWeightedPool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) GetVault() (common.Address, error) {
	return _BBasedWeightedPool.Contract.GetVault(&_BBasedWeightedPool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) GetVault() (common.Address, error) {
	return _BBasedWeightedPool.Contract.GetVault(&_BBasedWeightedPool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) InRecoveryMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "inRecoveryMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) InRecoveryMode() (bool, error) {
	return _BBasedWeightedPool.Contract.InRecoveryMode(&_BBasedWeightedPool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) InRecoveryMode() (bool, error) {
	return _BBasedWeightedPool.Contract.InRecoveryMode(&_BBasedWeightedPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Name() (string, error) {
	return _BBasedWeightedPool.Contract.Name(&_BBasedWeightedPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) Name() (string, error) {
	return _BBasedWeightedPool.Contract.Name(&_BBasedWeightedPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BBasedWeightedPool.Contract.Nonces(&_BBasedWeightedPool.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BBasedWeightedPool.Contract.Nonces(&_BBasedWeightedPool.CallOpts, owner)
}

// QueryExit is a free data retrieval call binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptIn, uint256[] amountsOut)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) QueryExit(opts *bind.CallOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptIn      *big.Int
	AmountsOut []*big.Int
}, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "queryExit", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)

	outstruct := new(struct {
		BptIn      *big.Int
		AmountsOut []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BptIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountsOut = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// QueryExit is a free data retrieval call binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptIn, uint256[] amountsOut)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptIn      *big.Int
	AmountsOut []*big.Int
}, error) {
	return _BBasedWeightedPool.Contract.QueryExit(&_BBasedWeightedPool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a free data retrieval call binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptIn, uint256[] amountsOut)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptIn      *big.Int
	AmountsOut []*big.Int
}, error) {
	return _BBasedWeightedPool.Contract.QueryExit(&_BBasedWeightedPool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a free data retrieval call binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptOut, uint256[] amountsIn)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) QueryJoin(opts *bind.CallOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptOut    *big.Int
	AmountsIn []*big.Int
}, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "queryJoin", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)

	outstruct := new(struct {
		BptOut    *big.Int
		AmountsIn []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BptOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountsIn = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// QueryJoin is a free data retrieval call binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptOut, uint256[] amountsIn)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptOut    *big.Int
	AmountsIn []*big.Int
}, error) {
	return _BBasedWeightedPool.Contract.QueryJoin(&_BBasedWeightedPool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a free data retrieval call binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptOut, uint256[] amountsIn)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptOut    *big.Int
	AmountsIn []*big.Int
}, error) {
	return _BBasedWeightedPool.Contract.QueryJoin(&_BBasedWeightedPool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Symbol() (string, error) {
	return _BBasedWeightedPool.Contract.Symbol(&_BBasedWeightedPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) Symbol() (string, error) {
	return _BBasedWeightedPool.Contract.Symbol(&_BBasedWeightedPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BBasedWeightedPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) TotalSupply() (*big.Int, error) {
	return _BBasedWeightedPool.Contract.TotalSupply(&_BBasedWeightedPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _BBasedWeightedPool.Contract.TotalSupply(&_BBasedWeightedPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Approve(&_BBasedWeightedPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Approve(&_BBasedWeightedPool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.DecreaseAllowance(&_BBasedWeightedPool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.DecreaseAllowance(&_BBasedWeightedPool.TransactOpts, spender, amount)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) DisableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "disableRecoveryMode")
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.DisableRecoveryMode(&_BBasedWeightedPool.TransactOpts)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.DisableRecoveryMode(&_BBasedWeightedPool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) EnableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "enableRecoveryMode")
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.EnableRecoveryMode(&_BBasedWeightedPool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.EnableRecoveryMode(&_BBasedWeightedPool.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.IncreaseAllowance(&_BBasedWeightedPool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.IncreaseAllowance(&_BBasedWeightedPool.TransactOpts, spender, addedValue)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) OnExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "onExitPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.OnExitPool(&_BBasedWeightedPool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.OnExitPool(&_BBasedWeightedPool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) OnJoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "onJoinPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.OnJoinPool(&_BBasedWeightedPool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.OnJoinPool(&_BBasedWeightedPool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) OnSwap(opts *bind.TransactOpts, request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "onSwap", request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.OnSwap(&_BBasedWeightedPool.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.OnSwap(&_BBasedWeightedPool.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Pause() (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Pause(&_BBasedWeightedPool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Pause(&_BBasedWeightedPool.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Permit(&_BBasedWeightedPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Permit(&_BBasedWeightedPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) SetSwapFeePercentage(opts *bind.TransactOpts, swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "setSwapFeePercentage", swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BBasedWeightedPool *BBasedWeightedPoolSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.SetSwapFeePercentage(&_BBasedWeightedPool.TransactOpts, swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.SetSwapFeePercentage(&_BBasedWeightedPool.TransactOpts, swapFeePercentage)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Transfer(&_BBasedWeightedPool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Transfer(&_BBasedWeightedPool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.TransferFrom(&_BBasedWeightedPool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.TransferFrom(&_BBasedWeightedPool.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BBasedWeightedPool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolSession) Unpause() (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Unpause(&_BBasedWeightedPool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BBasedWeightedPool *BBasedWeightedPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _BBasedWeightedPool.Contract.Unpause(&_BBasedWeightedPool.TransactOpts)
}

// BBasedWeightedPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolApprovalIterator struct {
	Event *BBasedWeightedPoolApproval // Event containing the contract specifics and raw log

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
func (it *BBasedWeightedPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BBasedWeightedPoolApproval)
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
		it.Event = new(BBasedWeightedPoolApproval)
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
func (it *BBasedWeightedPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BBasedWeightedPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BBasedWeightedPoolApproval represents a Approval event raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BBasedWeightedPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BBasedWeightedPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BBasedWeightedPoolApprovalIterator{contract: _BBasedWeightedPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BBasedWeightedPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BBasedWeightedPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BBasedWeightedPoolApproval)
				if err := _BBasedWeightedPool.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) ParseApproval(log types.Log) (*BBasedWeightedPoolApproval, error) {
	event := new(BBasedWeightedPoolApproval)
	if err := _BBasedWeightedPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BBasedWeightedPoolPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolPausedStateChangedIterator struct {
	Event *BBasedWeightedPoolPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *BBasedWeightedPoolPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BBasedWeightedPoolPausedStateChanged)
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
		it.Event = new(BBasedWeightedPoolPausedStateChanged)
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
func (it *BBasedWeightedPoolPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BBasedWeightedPoolPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BBasedWeightedPoolPausedStateChanged represents a PausedStateChanged event raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*BBasedWeightedPoolPausedStateChangedIterator, error) {

	logs, sub, err := _BBasedWeightedPool.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &BBasedWeightedPoolPausedStateChangedIterator{contract: _BBasedWeightedPool.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *BBasedWeightedPoolPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _BBasedWeightedPool.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BBasedWeightedPoolPausedStateChanged)
				if err := _BBasedWeightedPool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
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

// ParsePausedStateChanged is a log parse operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) ParsePausedStateChanged(log types.Log) (*BBasedWeightedPoolPausedStateChanged, error) {
	event := new(BBasedWeightedPoolPausedStateChanged)
	if err := _BBasedWeightedPool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BBasedWeightedPoolRecoveryModeStateChangedIterator is returned from FilterRecoveryModeStateChanged and is used to iterate over the raw logs and unpacked data for RecoveryModeStateChanged events raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolRecoveryModeStateChangedIterator struct {
	Event *BBasedWeightedPoolRecoveryModeStateChanged // Event containing the contract specifics and raw log

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
func (it *BBasedWeightedPoolRecoveryModeStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BBasedWeightedPoolRecoveryModeStateChanged)
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
		it.Event = new(BBasedWeightedPoolRecoveryModeStateChanged)
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
func (it *BBasedWeightedPoolRecoveryModeStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BBasedWeightedPoolRecoveryModeStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BBasedWeightedPoolRecoveryModeStateChanged represents a RecoveryModeStateChanged event raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolRecoveryModeStateChanged struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecoveryModeStateChanged is a free log retrieval operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) FilterRecoveryModeStateChanged(opts *bind.FilterOpts) (*BBasedWeightedPoolRecoveryModeStateChangedIterator, error) {

	logs, sub, err := _BBasedWeightedPool.contract.FilterLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return &BBasedWeightedPoolRecoveryModeStateChangedIterator{contract: _BBasedWeightedPool.contract, event: "RecoveryModeStateChanged", logs: logs, sub: sub}, nil
}

// WatchRecoveryModeStateChanged is a free log subscription operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) WatchRecoveryModeStateChanged(opts *bind.WatchOpts, sink chan<- *BBasedWeightedPoolRecoveryModeStateChanged) (event.Subscription, error) {

	logs, sub, err := _BBasedWeightedPool.contract.WatchLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BBasedWeightedPoolRecoveryModeStateChanged)
				if err := _BBasedWeightedPool.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
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

// ParseRecoveryModeStateChanged is a log parse operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) ParseRecoveryModeStateChanged(log types.Log) (*BBasedWeightedPoolRecoveryModeStateChanged, error) {
	event := new(BBasedWeightedPoolRecoveryModeStateChanged)
	if err := _BBasedWeightedPool.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BBasedWeightedPoolSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolSwapFeePercentageChangedIterator struct {
	Event *BBasedWeightedPoolSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *BBasedWeightedPoolSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BBasedWeightedPoolSwapFeePercentageChanged)
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
		it.Event = new(BBasedWeightedPoolSwapFeePercentageChanged)
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
func (it *BBasedWeightedPoolSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BBasedWeightedPoolSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BBasedWeightedPoolSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolSwapFeePercentageChanged struct {
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts) (*BBasedWeightedPoolSwapFeePercentageChangedIterator, error) {

	logs, sub, err := _BBasedWeightedPool.contract.FilterLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &BBasedWeightedPoolSwapFeePercentageChangedIterator{contract: _BBasedWeightedPool.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *BBasedWeightedPoolSwapFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _BBasedWeightedPool.contract.WatchLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BBasedWeightedPoolSwapFeePercentageChanged)
				if err := _BBasedWeightedPool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
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

// ParseSwapFeePercentageChanged is a log parse operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) ParseSwapFeePercentageChanged(log types.Log) (*BBasedWeightedPoolSwapFeePercentageChanged, error) {
	event := new(BBasedWeightedPoolSwapFeePercentageChanged)
	if err := _BBasedWeightedPool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BBasedWeightedPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolTransferIterator struct {
	Event *BBasedWeightedPoolTransfer // Event containing the contract specifics and raw log

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
func (it *BBasedWeightedPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BBasedWeightedPoolTransfer)
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
		it.Event = new(BBasedWeightedPoolTransfer)
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
func (it *BBasedWeightedPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BBasedWeightedPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BBasedWeightedPoolTransfer represents a Transfer event raised by the BBasedWeightedPool contract.
type BBasedWeightedPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BBasedWeightedPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BBasedWeightedPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BBasedWeightedPoolTransferIterator{contract: _BBasedWeightedPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BBasedWeightedPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BBasedWeightedPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BBasedWeightedPoolTransfer)
				if err := _BBasedWeightedPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BBasedWeightedPool *BBasedWeightedPoolFilterer) ParseTransfer(log types.Log) (*BBasedWeightedPoolTransfer, error) {
	event := new(BBasedWeightedPoolTransfer)
	if err := _BBasedWeightedPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
