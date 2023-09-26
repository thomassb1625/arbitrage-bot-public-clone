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

// MoonbeamExecutorMetaData contains all meta data concerning the MoonbeamExecutor contract.
var MoonbeamExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wglmr\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wglmrAmountToFirstMarket\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_glmrAmountToCoinbase\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"}],\"name\":\"swapWithoutFlashLoan\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawGLMR\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MoonbeamExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use MoonbeamExecutorMetaData.ABI instead.
var MoonbeamExecutorABI = MoonbeamExecutorMetaData.ABI

// MoonbeamExecutor is an auto generated Go binding around an Ethereum contract.
type MoonbeamExecutor struct {
	MoonbeamExecutorCaller     // Read-only binding to the contract
	MoonbeamExecutorTransactor // Write-only binding to the contract
	MoonbeamExecutorFilterer   // Log filterer for contract events
}

// MoonbeamExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MoonbeamExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonbeamExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MoonbeamExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonbeamExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MoonbeamExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonbeamExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MoonbeamExecutorSession struct {
	Contract     *MoonbeamExecutor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MoonbeamExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MoonbeamExecutorCallerSession struct {
	Contract *MoonbeamExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MoonbeamExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MoonbeamExecutorTransactorSession struct {
	Contract     *MoonbeamExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MoonbeamExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MoonbeamExecutorRaw struct {
	Contract *MoonbeamExecutor // Generic contract binding to access the raw methods on
}

// MoonbeamExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MoonbeamExecutorCallerRaw struct {
	Contract *MoonbeamExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// MoonbeamExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MoonbeamExecutorTransactorRaw struct {
	Contract *MoonbeamExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMoonbeamExecutor creates a new instance of MoonbeamExecutor, bound to a specific deployed contract.
func NewMoonbeamExecutor(address common.Address, backend bind.ContractBackend) (*MoonbeamExecutor, error) {
	contract, err := bindMoonbeamExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MoonbeamExecutor{MoonbeamExecutorCaller: MoonbeamExecutorCaller{contract: contract}, MoonbeamExecutorTransactor: MoonbeamExecutorTransactor{contract: contract}, MoonbeamExecutorFilterer: MoonbeamExecutorFilterer{contract: contract}}, nil
}

// NewMoonbeamExecutorCaller creates a new read-only instance of MoonbeamExecutor, bound to a specific deployed contract.
func NewMoonbeamExecutorCaller(address common.Address, caller bind.ContractCaller) (*MoonbeamExecutorCaller, error) {
	contract, err := bindMoonbeamExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MoonbeamExecutorCaller{contract: contract}, nil
}

// NewMoonbeamExecutorTransactor creates a new write-only instance of MoonbeamExecutor, bound to a specific deployed contract.
func NewMoonbeamExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*MoonbeamExecutorTransactor, error) {
	contract, err := bindMoonbeamExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MoonbeamExecutorTransactor{contract: contract}, nil
}

// NewMoonbeamExecutorFilterer creates a new log filterer instance of MoonbeamExecutor, bound to a specific deployed contract.
func NewMoonbeamExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*MoonbeamExecutorFilterer, error) {
	contract, err := bindMoonbeamExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MoonbeamExecutorFilterer{contract: contract}, nil
}

// bindMoonbeamExecutor binds a generic wrapper to an already deployed contract.
func bindMoonbeamExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MoonbeamExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MoonbeamExecutor *MoonbeamExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MoonbeamExecutor.Contract.MoonbeamExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MoonbeamExecutor *MoonbeamExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.MoonbeamExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MoonbeamExecutor *MoonbeamExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.MoonbeamExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MoonbeamExecutor *MoonbeamExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MoonbeamExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MoonbeamExecutor *MoonbeamExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MoonbeamExecutor *MoonbeamExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.contract.Transact(opts, method, params...)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_MoonbeamExecutor *MoonbeamExecutorCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonbeamExecutor.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_MoonbeamExecutor *MoonbeamExecutorSession) Executor() (common.Address, error) {
	return _MoonbeamExecutor.Contract.Executor(&_MoonbeamExecutor.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_MoonbeamExecutor *MoonbeamExecutorCallerSession) Executor() (common.Address, error) {
	return _MoonbeamExecutor.Contract.Executor(&_MoonbeamExecutor.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _tokenAddress) view returns(uint256)
func (_MoonbeamExecutor *MoonbeamExecutorCaller) GetBalance(opts *bind.CallOpts, _tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MoonbeamExecutor.contract.Call(opts, &out, "getBalance", _tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _tokenAddress) view returns(uint256)
func (_MoonbeamExecutor *MoonbeamExecutorSession) GetBalance(_tokenAddress common.Address) (*big.Int, error) {
	return _MoonbeamExecutor.Contract.GetBalance(&_MoonbeamExecutor.CallOpts, _tokenAddress)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _tokenAddress) view returns(uint256)
func (_MoonbeamExecutor *MoonbeamExecutorCallerSession) GetBalance(_tokenAddress common.Address) (*big.Int, error) {
	return _MoonbeamExecutor.Contract.GetBalance(&_MoonbeamExecutor.CallOpts, _tokenAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MoonbeamExecutor *MoonbeamExecutorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonbeamExecutor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MoonbeamExecutor *MoonbeamExecutorSession) Owner() (common.Address, error) {
	return _MoonbeamExecutor.Contract.Owner(&_MoonbeamExecutor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MoonbeamExecutor *MoonbeamExecutorCallerSession) Owner() (common.Address, error) {
	return _MoonbeamExecutor.Contract.Owner(&_MoonbeamExecutor.CallOpts)
}

// SwapWithoutFlashLoan is a paid mutator transaction binding the contract method 0x4686fb9e.
//
// Solidity: function swapWithoutFlashLoan(uint256 _wglmrAmountToFirstMarket, uint256 _glmrAmountToCoinbase, address[] _targets, bytes[] _payloads) payable returns()
func (_MoonbeamExecutor *MoonbeamExecutorTransactor) SwapWithoutFlashLoan(opts *bind.TransactOpts, _wglmrAmountToFirstMarket *big.Int, _glmrAmountToCoinbase *big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _MoonbeamExecutor.contract.Transact(opts, "swapWithoutFlashLoan", _wglmrAmountToFirstMarket, _glmrAmountToCoinbase, _targets, _payloads)
}

// SwapWithoutFlashLoan is a paid mutator transaction binding the contract method 0x4686fb9e.
//
// Solidity: function swapWithoutFlashLoan(uint256 _wglmrAmountToFirstMarket, uint256 _glmrAmountToCoinbase, address[] _targets, bytes[] _payloads) payable returns()
func (_MoonbeamExecutor *MoonbeamExecutorSession) SwapWithoutFlashLoan(_wglmrAmountToFirstMarket *big.Int, _glmrAmountToCoinbase *big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.SwapWithoutFlashLoan(&_MoonbeamExecutor.TransactOpts, _wglmrAmountToFirstMarket, _glmrAmountToCoinbase, _targets, _payloads)
}

// SwapWithoutFlashLoan is a paid mutator transaction binding the contract method 0x4686fb9e.
//
// Solidity: function swapWithoutFlashLoan(uint256 _wglmrAmountToFirstMarket, uint256 _glmrAmountToCoinbase, address[] _targets, bytes[] _payloads) payable returns()
func (_MoonbeamExecutor *MoonbeamExecutorTransactorSession) SwapWithoutFlashLoan(_wglmrAmountToFirstMarket *big.Int, _glmrAmountToCoinbase *big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.SwapWithoutFlashLoan(&_MoonbeamExecutor.TransactOpts, _wglmrAmountToFirstMarket, _glmrAmountToCoinbase, _targets, _payloads)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _tokenAddress) returns()
func (_MoonbeamExecutor *MoonbeamExecutorTransactor) Withdraw(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _MoonbeamExecutor.contract.Transact(opts, "withdraw", _tokenAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _tokenAddress) returns()
func (_MoonbeamExecutor *MoonbeamExecutorSession) Withdraw(_tokenAddress common.Address) (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.Withdraw(&_MoonbeamExecutor.TransactOpts, _tokenAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _tokenAddress) returns()
func (_MoonbeamExecutor *MoonbeamExecutorTransactorSession) Withdraw(_tokenAddress common.Address) (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.Withdraw(&_MoonbeamExecutor.TransactOpts, _tokenAddress)
}

// WithdrawGLMR is a paid mutator transaction binding the contract method 0xa1c939e2.
//
// Solidity: function withdrawGLMR() returns()
func (_MoonbeamExecutor *MoonbeamExecutorTransactor) WithdrawGLMR(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonbeamExecutor.contract.Transact(opts, "withdrawGLMR")
}

// WithdrawGLMR is a paid mutator transaction binding the contract method 0xa1c939e2.
//
// Solidity: function withdrawGLMR() returns()
func (_MoonbeamExecutor *MoonbeamExecutorSession) WithdrawGLMR() (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.WithdrawGLMR(&_MoonbeamExecutor.TransactOpts)
}

// WithdrawGLMR is a paid mutator transaction binding the contract method 0xa1c939e2.
//
// Solidity: function withdrawGLMR() returns()
func (_MoonbeamExecutor *MoonbeamExecutorTransactorSession) WithdrawGLMR() (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.WithdrawGLMR(&_MoonbeamExecutor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MoonbeamExecutor *MoonbeamExecutorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonbeamExecutor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MoonbeamExecutor *MoonbeamExecutorSession) Receive() (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.Receive(&_MoonbeamExecutor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MoonbeamExecutor *MoonbeamExecutorTransactorSession) Receive() (*types.Transaction, error) {
	return _MoonbeamExecutor.Contract.Receive(&_MoonbeamExecutor.TransactOpts)
}
