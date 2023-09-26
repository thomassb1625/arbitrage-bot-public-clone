// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapV2

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

// UniswapV2QueryMetaData contains all meta data concerning the UniswapV2Query contract.
var UniswapV2QueryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractUniswapV2Factory\",\"name\":\"_uniswapFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stop\",\"type\":\"uint256\"}],\"name\":\"getPairsByIndexRange\",\"outputs\":[{\"internalType\":\"address[3][]\",\"name\":\"\",\"type\":\"address[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV2Pair[]\",\"name\":\"_pairs\",\"type\":\"address[]\"}],\"name\":\"getReservesByPairs\",\"outputs\":[{\"internalType\":\"uint256[3][]\",\"name\":\"\",\"type\":\"uint256[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV2QueryABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2QueryMetaData.ABI instead.
var UniswapV2QueryABI = UniswapV2QueryMetaData.ABI

// UniswapV2Query is an auto generated Go binding around an Ethereum contract.
type UniswapV2Query struct {
	UniswapV2QueryCaller     // Read-only binding to the contract
	UniswapV2QueryTransactor // Write-only binding to the contract
	UniswapV2QueryFilterer   // Log filterer for contract events
}

// UniswapV2QueryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2QueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2QueryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2QueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2QueryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2QueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2QuerySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2QuerySession struct {
	Contract     *UniswapV2Query   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV2QueryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2QueryCallerSession struct {
	Contract *UniswapV2QueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UniswapV2QueryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2QueryTransactorSession struct {
	Contract     *UniswapV2QueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapV2QueryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2QueryRaw struct {
	Contract *UniswapV2Query // Generic contract binding to access the raw methods on
}

// UniswapV2QueryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2QueryCallerRaw struct {
	Contract *UniswapV2QueryCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2QueryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2QueryTransactorRaw struct {
	Contract *UniswapV2QueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2Query creates a new instance of UniswapV2Query, bound to a specific deployed contract.
func NewUniswapV2Query(address common.Address, backend bind.ContractBackend) (*UniswapV2Query, error) {
	contract, err := bindUniswapV2Query(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Query{UniswapV2QueryCaller: UniswapV2QueryCaller{contract: contract}, UniswapV2QueryTransactor: UniswapV2QueryTransactor{contract: contract}, UniswapV2QueryFilterer: UniswapV2QueryFilterer{contract: contract}}, nil
}

// NewUniswapV2QueryCaller creates a new read-only instance of UniswapV2Query, bound to a specific deployed contract.
func NewUniswapV2QueryCaller(address common.Address, caller bind.ContractCaller) (*UniswapV2QueryCaller, error) {
	contract, err := bindUniswapV2Query(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2QueryCaller{contract: contract}, nil
}

// NewUniswapV2QueryTransactor creates a new write-only instance of UniswapV2Query, bound to a specific deployed contract.
func NewUniswapV2QueryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2QueryTransactor, error) {
	contract, err := bindUniswapV2Query(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2QueryTransactor{contract: contract}, nil
}

// NewUniswapV2QueryFilterer creates a new log filterer instance of UniswapV2Query, bound to a specific deployed contract.
func NewUniswapV2QueryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2QueryFilterer, error) {
	contract, err := bindUniswapV2Query(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2QueryFilterer{contract: contract}, nil
}

// bindUniswapV2Query binds a generic wrapper to an already deployed contract.
func bindUniswapV2Query(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV2QueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Query *UniswapV2QueryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Query.Contract.UniswapV2QueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Query *UniswapV2QueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Query.Contract.UniswapV2QueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Query *UniswapV2QueryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Query.Contract.UniswapV2QueryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Query *UniswapV2QueryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Query.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Query *UniswapV2QueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Query.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Query *UniswapV2QueryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Query.Contract.contract.Transact(opts, method, params...)
}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _uniswapFactory, uint256 _start, uint256 _stop) view returns(address[3][])
func (_UniswapV2Query *UniswapV2QueryCaller) GetPairsByIndexRange(opts *bind.CallOpts, _uniswapFactory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, error) {
	var out []interface{}
	err := _UniswapV2Query.contract.Call(opts, &out, "getPairsByIndexRange", _uniswapFactory, _start, _stop)

	if err != nil {
		return *new([][3]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]common.Address)).(*[][3]common.Address)

	return out0, err

}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _uniswapFactory, uint256 _start, uint256 _stop) view returns(address[3][])
func (_UniswapV2Query *UniswapV2QuerySession) GetPairsByIndexRange(_uniswapFactory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, error) {
	return _UniswapV2Query.Contract.GetPairsByIndexRange(&_UniswapV2Query.CallOpts, _uniswapFactory, _start, _stop)
}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _uniswapFactory, uint256 _start, uint256 _stop) view returns(address[3][])
func (_UniswapV2Query *UniswapV2QueryCallerSession) GetPairsByIndexRange(_uniswapFactory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, error) {
	return _UniswapV2Query.Contract.GetPairsByIndexRange(&_UniswapV2Query.CallOpts, _uniswapFactory, _start, _stop)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_UniswapV2Query *UniswapV2QueryCaller) GetReservesByPairs(opts *bind.CallOpts, _pairs []common.Address) ([][3]*big.Int, error) {
	var out []interface{}
	err := _UniswapV2Query.contract.Call(opts, &out, "getReservesByPairs", _pairs)

	if err != nil {
		return *new([][3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]*big.Int)).(*[][3]*big.Int)

	return out0, err

}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_UniswapV2Query *UniswapV2QuerySession) GetReservesByPairs(_pairs []common.Address) ([][3]*big.Int, error) {
	return _UniswapV2Query.Contract.GetReservesByPairs(&_UniswapV2Query.CallOpts, _pairs)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_UniswapV2Query *UniswapV2QueryCallerSession) GetReservesByPairs(_pairs []common.Address) ([][3]*big.Int, error) {
	return _UniswapV2Query.Contract.GetReservesByPairs(&_UniswapV2Query.CallOpts, _pairs)
}
