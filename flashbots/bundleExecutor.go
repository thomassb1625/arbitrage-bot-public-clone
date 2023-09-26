// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flashbots

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

// BundleExecutorMetaData contains all meta data concerning the BundleExecutor contract.
var BundleExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LENDING_POOL\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"premiums\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_ethAmountToCoinbase\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_flashLoanAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_modes\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"}],\"name\":\"swapWithFlashLoan\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmountToCoinbase\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"}],\"name\":\"swapWithoutFlashLoan\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BundleExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use BundleExecutorMetaData.ABI instead.
var BundleExecutorABI = BundleExecutorMetaData.ABI

// BundleExecutor is an auto generated Go binding around an Ethereum contract.
type BundleExecutor struct {
	BundleExecutorCaller     // Read-only binding to the contract
	BundleExecutorTransactor // Write-only binding to the contract
	BundleExecutorFilterer   // Log filterer for contract events
}

// BundleExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BundleExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BundleExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BundleExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BundleExecutorSession struct {
	Contract     *BundleExecutor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BundleExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BundleExecutorCallerSession struct {
	Contract *BundleExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BundleExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BundleExecutorTransactorSession struct {
	Contract     *BundleExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BundleExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BundleExecutorRaw struct {
	Contract *BundleExecutor // Generic contract binding to access the raw methods on
}

// BundleExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BundleExecutorCallerRaw struct {
	Contract *BundleExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// BundleExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BundleExecutorTransactorRaw struct {
	Contract *BundleExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBundleExecutor creates a new instance of BundleExecutor, bound to a specific deployed contract.
func NewBundleExecutor(address common.Address, backend bind.ContractBackend) (*BundleExecutor, error) {
	contract, err := bindBundleExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BundleExecutor{BundleExecutorCaller: BundleExecutorCaller{contract: contract}, BundleExecutorTransactor: BundleExecutorTransactor{contract: contract}, BundleExecutorFilterer: BundleExecutorFilterer{contract: contract}}, nil
}

// NewBundleExecutorCaller creates a new read-only instance of BundleExecutor, bound to a specific deployed contract.
func NewBundleExecutorCaller(address common.Address, caller bind.ContractCaller) (*BundleExecutorCaller, error) {
	contract, err := bindBundleExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorCaller{contract: contract}, nil
}

// NewBundleExecutorTransactor creates a new write-only instance of BundleExecutor, bound to a specific deployed contract.
func NewBundleExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*BundleExecutorTransactor, error) {
	contract, err := bindBundleExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorTransactor{contract: contract}, nil
}

// NewBundleExecutorFilterer creates a new log filterer instance of BundleExecutor, bound to a specific deployed contract.
func NewBundleExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*BundleExecutorFilterer, error) {
	contract, err := bindBundleExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorFilterer{contract: contract}, nil
}

// bindBundleExecutor binds a generic wrapper to an already deployed contract.
func bindBundleExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BundleExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleExecutor *BundleExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleExecutor.Contract.BundleExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleExecutor *BundleExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutor.Contract.BundleExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleExecutor *BundleExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleExecutor.Contract.BundleExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleExecutor *BundleExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleExecutor *BundleExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleExecutor *BundleExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleExecutor.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_BundleExecutor *BundleExecutorCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BundleExecutor.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_BundleExecutor *BundleExecutorSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _BundleExecutor.Contract.ADDRESSESPROVIDER(&_BundleExecutor.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_BundleExecutor *BundleExecutorCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _BundleExecutor.Contract.ADDRESSESPROVIDER(&_BundleExecutor.CallOpts)
}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_BundleExecutor *BundleExecutorCaller) LENDINGPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BundleExecutor.contract.Call(opts, &out, "LENDING_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_BundleExecutor *BundleExecutorSession) LENDINGPOOL() (common.Address, error) {
	return _BundleExecutor.Contract.LENDINGPOOL(&_BundleExecutor.CallOpts)
}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_BundleExecutor *BundleExecutorCallerSession) LENDINGPOOL() (common.Address, error) {
	return _BundleExecutor.Contract.LENDINGPOOL(&_BundleExecutor.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_BundleExecutor *BundleExecutorCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BundleExecutor.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_BundleExecutor *BundleExecutorSession) Executor() (common.Address, error) {
	return _BundleExecutor.Contract.Executor(&_BundleExecutor.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_BundleExecutor *BundleExecutorCallerSession) Executor() (common.Address, error) {
	return _BundleExecutor.Contract.Executor(&_BundleExecutor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BundleExecutor *BundleExecutorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BundleExecutor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BundleExecutor *BundleExecutorSession) Owner() (common.Address, error) {
	return _BundleExecutor.Contract.Owner(&_BundleExecutor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BundleExecutor *BundleExecutorCallerSession) Owner() (common.Address, error) {
	return _BundleExecutor.Contract.Owner(&_BundleExecutor.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x6dbf2fa0.
//
// Solidity: function call(address _to, uint256 _value, bytes _data) payable returns(bytes)
func (_BundleExecutor *BundleExecutorTransactor) Call(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _BundleExecutor.contract.Transact(opts, "call", _to, _value, _data)
}

// Call is a paid mutator transaction binding the contract method 0x6dbf2fa0.
//
// Solidity: function call(address _to, uint256 _value, bytes _data) payable returns(bytes)
func (_BundleExecutor *BundleExecutorSession) Call(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.Call(&_BundleExecutor.TransactOpts, _to, _value, _data)
}

// Call is a paid mutator transaction binding the contract method 0x6dbf2fa0.
//
// Solidity: function call(address _to, uint256 _value, bytes _data) payable returns(bytes)
func (_BundleExecutor *BundleExecutorTransactorSession) Call(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.Call(&_BundleExecutor.TransactOpts, _to, _value, _data)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_BundleExecutor *BundleExecutorTransactor) ExecuteOperation(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _BundleExecutor.contract.Transact(opts, "executeOperation", assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_BundleExecutor *BundleExecutorSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.ExecuteOperation(&_BundleExecutor.TransactOpts, assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_BundleExecutor *BundleExecutorTransactorSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.ExecuteOperation(&_BundleExecutor.TransactOpts, assets, amounts, premiums, initiator, params)
}

// SwapWithFlashLoan is a paid mutator transaction binding the contract method 0xbc19815f.
//
// Solidity: function swapWithFlashLoan(address[] _tokens, uint256 _ethAmountToCoinbase, uint256[] _flashLoanAmounts, uint256[] _modes, address[] _targets, bytes[] _payloads) payable returns()
func (_BundleExecutor *BundleExecutorTransactor) SwapWithFlashLoan(opts *bind.TransactOpts, _tokens []common.Address, _ethAmountToCoinbase *big.Int, _flashLoanAmounts []*big.Int, _modes []*big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _BundleExecutor.contract.Transact(opts, "swapWithFlashLoan", _tokens, _ethAmountToCoinbase, _flashLoanAmounts, _modes, _targets, _payloads)
}

// SwapWithFlashLoan is a paid mutator transaction binding the contract method 0xbc19815f.
//
// Solidity: function swapWithFlashLoan(address[] _tokens, uint256 _ethAmountToCoinbase, uint256[] _flashLoanAmounts, uint256[] _modes, address[] _targets, bytes[] _payloads) payable returns()
func (_BundleExecutor *BundleExecutorSession) SwapWithFlashLoan(_tokens []common.Address, _ethAmountToCoinbase *big.Int, _flashLoanAmounts []*big.Int, _modes []*big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.SwapWithFlashLoan(&_BundleExecutor.TransactOpts, _tokens, _ethAmountToCoinbase, _flashLoanAmounts, _modes, _targets, _payloads)
}

// SwapWithFlashLoan is a paid mutator transaction binding the contract method 0xbc19815f.
//
// Solidity: function swapWithFlashLoan(address[] _tokens, uint256 _ethAmountToCoinbase, uint256[] _flashLoanAmounts, uint256[] _modes, address[] _targets, bytes[] _payloads) payable returns()
func (_BundleExecutor *BundleExecutorTransactorSession) SwapWithFlashLoan(_tokens []common.Address, _ethAmountToCoinbase *big.Int, _flashLoanAmounts []*big.Int, _modes []*big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.SwapWithFlashLoan(&_BundleExecutor.TransactOpts, _tokens, _ethAmountToCoinbase, _flashLoanAmounts, _modes, _targets, _payloads)
}

// SwapWithoutFlashLoan is a paid mutator transaction binding the contract method 0x9303038d.
//
// Solidity: function swapWithoutFlashLoan(uint256 _ethAmountToCoinbase, address[] _targets, bytes[] _payloads) payable returns()
func (_BundleExecutor *BundleExecutorTransactor) SwapWithoutFlashLoan(opts *bind.TransactOpts, _ethAmountToCoinbase *big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _BundleExecutor.contract.Transact(opts, "swapWithoutFlashLoan", _ethAmountToCoinbase, _targets, _payloads)
}

// SwapWithoutFlashLoan is a paid mutator transaction binding the contract method 0x9303038d.
//
// Solidity: function swapWithoutFlashLoan(uint256 _ethAmountToCoinbase, address[] _targets, bytes[] _payloads) payable returns()
func (_BundleExecutor *BundleExecutorSession) SwapWithoutFlashLoan(_ethAmountToCoinbase *big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.SwapWithoutFlashLoan(&_BundleExecutor.TransactOpts, _ethAmountToCoinbase, _targets, _payloads)
}

// SwapWithoutFlashLoan is a paid mutator transaction binding the contract method 0x9303038d.
//
// Solidity: function swapWithoutFlashLoan(uint256 _ethAmountToCoinbase, address[] _targets, bytes[] _payloads) payable returns()
func (_BundleExecutor *BundleExecutorTransactorSession) SwapWithoutFlashLoan(_ethAmountToCoinbase *big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.SwapWithoutFlashLoan(&_BundleExecutor.TransactOpts, _ethAmountToCoinbase, _targets, _payloads)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleExecutor *BundleExecutorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleExecutor *BundleExecutorSession) Receive() (*types.Transaction, error) {
	return _BundleExecutor.Contract.Receive(&_BundleExecutor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleExecutor *BundleExecutorTransactorSession) Receive() (*types.Transaction, error) {
	return _BundleExecutor.Contract.Receive(&_BundleExecutor.TransactOpts)
}
