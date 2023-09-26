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

// BWeightedFactoryMetaData contains all meta data concerning the BWeightedFactory contract.
var BWeightedFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"contractIProtocolFeePercentagesProvider\",\"name\":\"protocolFeeProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FactoryDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"normalizedWeights\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIRateProvider[]\",\"name\":\"rateProviders\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreationCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreationCodeContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractB\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPauseConfiguration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeePercentagesProvider\",\"outputs\":[{\"internalType\":\"contractIProtocolFeePercentagesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolFromFactory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BWeightedFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use BWeightedFactoryMetaData.ABI instead.
var BWeightedFactoryABI = BWeightedFactoryMetaData.ABI

// BWeightedFactory is an auto generated Go binding around an Ethereum contract.
type BWeightedFactory struct {
	BWeightedFactoryCaller     // Read-only binding to the contract
	BWeightedFactoryTransactor // Write-only binding to the contract
	BWeightedFactoryFilterer   // Log filterer for contract events
}

// BWeightedFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BWeightedFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BWeightedFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BWeightedFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BWeightedFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BWeightedFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BWeightedFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BWeightedFactorySession struct {
	Contract     *BWeightedFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BWeightedFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BWeightedFactoryCallerSession struct {
	Contract *BWeightedFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BWeightedFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BWeightedFactoryTransactorSession struct {
	Contract     *BWeightedFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BWeightedFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BWeightedFactoryRaw struct {
	Contract *BWeightedFactory // Generic contract binding to access the raw methods on
}

// BWeightedFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BWeightedFactoryCallerRaw struct {
	Contract *BWeightedFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BWeightedFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BWeightedFactoryTransactorRaw struct {
	Contract *BWeightedFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBWeightedFactory creates a new instance of BWeightedFactory, bound to a specific deployed contract.
func NewBWeightedFactory(address common.Address, backend bind.ContractBackend) (*BWeightedFactory, error) {
	contract, err := bindBWeightedFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BWeightedFactory{BWeightedFactoryCaller: BWeightedFactoryCaller{contract: contract}, BWeightedFactoryTransactor: BWeightedFactoryTransactor{contract: contract}, BWeightedFactoryFilterer: BWeightedFactoryFilterer{contract: contract}}, nil
}

// NewBWeightedFactoryCaller creates a new read-only instance of BWeightedFactory, bound to a specific deployed contract.
func NewBWeightedFactoryCaller(address common.Address, caller bind.ContractCaller) (*BWeightedFactoryCaller, error) {
	contract, err := bindBWeightedFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BWeightedFactoryCaller{contract: contract}, nil
}

// NewBWeightedFactoryTransactor creates a new write-only instance of BWeightedFactory, bound to a specific deployed contract.
func NewBWeightedFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BWeightedFactoryTransactor, error) {
	contract, err := bindBWeightedFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BWeightedFactoryTransactor{contract: contract}, nil
}

// NewBWeightedFactoryFilterer creates a new log filterer instance of BWeightedFactory, bound to a specific deployed contract.
func NewBWeightedFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BWeightedFactoryFilterer, error) {
	contract, err := bindBWeightedFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BWeightedFactoryFilterer{contract: contract}, nil
}

// bindBWeightedFactory binds a generic wrapper to an already deployed contract.
func bindBWeightedFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BWeightedFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BWeightedFactory *BWeightedFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BWeightedFactory.Contract.BWeightedFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BWeightedFactory *BWeightedFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedFactory.Contract.BWeightedFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BWeightedFactory *BWeightedFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BWeightedFactory.Contract.BWeightedFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BWeightedFactory *BWeightedFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BWeightedFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BWeightedFactory *BWeightedFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BWeightedFactory *BWeightedFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BWeightedFactory.Contract.contract.Transact(opts, method, params...)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BWeightedFactory *BWeightedFactoryCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _BWeightedFactory.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BWeightedFactory *BWeightedFactorySession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BWeightedFactory.Contract.GetActionId(&_BWeightedFactory.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BWeightedFactory *BWeightedFactoryCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BWeightedFactory.Contract.GetActionId(&_BWeightedFactory.CallOpts, selector)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BWeightedFactory *BWeightedFactoryCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BWeightedFactory.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BWeightedFactory *BWeightedFactorySession) GetAuthorizer() (common.Address, error) {
	return _BWeightedFactory.Contract.GetAuthorizer(&_BWeightedFactory.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BWeightedFactory *BWeightedFactoryCallerSession) GetAuthorizer() (common.Address, error) {
	return _BWeightedFactory.Contract.GetAuthorizer(&_BWeightedFactory.CallOpts)
}

// GetCreationCode is a free data retrieval call binding the contract method 0x00c194db.
//
// Solidity: function getCreationCode() view returns(bytes)
func (_BWeightedFactory *BWeightedFactoryCaller) GetCreationCode(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BWeightedFactory.contract.Call(opts, &out, "getCreationCode")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCreationCode is a free data retrieval call binding the contract method 0x00c194db.
//
// Solidity: function getCreationCode() view returns(bytes)
func (_BWeightedFactory *BWeightedFactorySession) GetCreationCode() ([]byte, error) {
	return _BWeightedFactory.Contract.GetCreationCode(&_BWeightedFactory.CallOpts)
}

// GetCreationCode is a free data retrieval call binding the contract method 0x00c194db.
//
// Solidity: function getCreationCode() view returns(bytes)
func (_BWeightedFactory *BWeightedFactoryCallerSession) GetCreationCode() ([]byte, error) {
	return _BWeightedFactory.Contract.GetCreationCode(&_BWeightedFactory.CallOpts)
}

// GetCreationCodeContracts is a free data retrieval call binding the contract method 0x174481fa.
//
// Solidity: function getCreationCodeContracts() view returns(address contractA, address contractB)
func (_BWeightedFactory *BWeightedFactoryCaller) GetCreationCodeContracts(opts *bind.CallOpts) (struct {
	ContractA common.Address
	ContractB common.Address
}, error) {
	var out []interface{}
	err := _BWeightedFactory.contract.Call(opts, &out, "getCreationCodeContracts")

	outstruct := new(struct {
		ContractA common.Address
		ContractB common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ContractA = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ContractB = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetCreationCodeContracts is a free data retrieval call binding the contract method 0x174481fa.
//
// Solidity: function getCreationCodeContracts() view returns(address contractA, address contractB)
func (_BWeightedFactory *BWeightedFactorySession) GetCreationCodeContracts() (struct {
	ContractA common.Address
	ContractB common.Address
}, error) {
	return _BWeightedFactory.Contract.GetCreationCodeContracts(&_BWeightedFactory.CallOpts)
}

// GetCreationCodeContracts is a free data retrieval call binding the contract method 0x174481fa.
//
// Solidity: function getCreationCodeContracts() view returns(address contractA, address contractB)
func (_BWeightedFactory *BWeightedFactoryCallerSession) GetCreationCodeContracts() (struct {
	ContractA common.Address
	ContractB common.Address
}, error) {
	return _BWeightedFactory.Contract.GetCreationCodeContracts(&_BWeightedFactory.CallOpts)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_BWeightedFactory *BWeightedFactoryCaller) GetPauseConfiguration(opts *bind.CallOpts) (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	var out []interface{}
	err := _BWeightedFactory.contract.Call(opts, &out, "getPauseConfiguration")

	outstruct := new(struct {
		PauseWindowDuration  *big.Int
		BufferPeriodDuration *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PauseWindowDuration = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodDuration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_BWeightedFactory *BWeightedFactorySession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _BWeightedFactory.Contract.GetPauseConfiguration(&_BWeightedFactory.CallOpts)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_BWeightedFactory *BWeightedFactoryCallerSession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _BWeightedFactory.Contract.GetPauseConfiguration(&_BWeightedFactory.CallOpts)
}

// GetProtocolFeePercentagesProvider is a free data retrieval call binding the contract method 0x739238d6.
//
// Solidity: function getProtocolFeePercentagesProvider() view returns(address)
func (_BWeightedFactory *BWeightedFactoryCaller) GetProtocolFeePercentagesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BWeightedFactory.contract.Call(opts, &out, "getProtocolFeePercentagesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeePercentagesProvider is a free data retrieval call binding the contract method 0x739238d6.
//
// Solidity: function getProtocolFeePercentagesProvider() view returns(address)
func (_BWeightedFactory *BWeightedFactorySession) GetProtocolFeePercentagesProvider() (common.Address, error) {
	return _BWeightedFactory.Contract.GetProtocolFeePercentagesProvider(&_BWeightedFactory.CallOpts)
}

// GetProtocolFeePercentagesProvider is a free data retrieval call binding the contract method 0x739238d6.
//
// Solidity: function getProtocolFeePercentagesProvider() view returns(address)
func (_BWeightedFactory *BWeightedFactoryCallerSession) GetProtocolFeePercentagesProvider() (common.Address, error) {
	return _BWeightedFactory.Contract.GetProtocolFeePercentagesProvider(&_BWeightedFactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BWeightedFactory *BWeightedFactoryCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BWeightedFactory.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BWeightedFactory *BWeightedFactorySession) GetVault() (common.Address, error) {
	return _BWeightedFactory.Contract.GetVault(&_BWeightedFactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BWeightedFactory *BWeightedFactoryCallerSession) GetVault() (common.Address, error) {
	return _BWeightedFactory.Contract.GetVault(&_BWeightedFactory.CallOpts)
}

// IsDisabled is a free data retrieval call binding the contract method 0x6c57f5a9.
//
// Solidity: function isDisabled() view returns(bool)
func (_BWeightedFactory *BWeightedFactoryCaller) IsDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BWeightedFactory.contract.Call(opts, &out, "isDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDisabled is a free data retrieval call binding the contract method 0x6c57f5a9.
//
// Solidity: function isDisabled() view returns(bool)
func (_BWeightedFactory *BWeightedFactorySession) IsDisabled() (bool, error) {
	return _BWeightedFactory.Contract.IsDisabled(&_BWeightedFactory.CallOpts)
}

// IsDisabled is a free data retrieval call binding the contract method 0x6c57f5a9.
//
// Solidity: function isDisabled() view returns(bool)
func (_BWeightedFactory *BWeightedFactoryCallerSession) IsDisabled() (bool, error) {
	return _BWeightedFactory.Contract.IsDisabled(&_BWeightedFactory.CallOpts)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_BWeightedFactory *BWeightedFactoryCaller) IsPoolFromFactory(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _BWeightedFactory.contract.Call(opts, &out, "isPoolFromFactory", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_BWeightedFactory *BWeightedFactorySession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _BWeightedFactory.Contract.IsPoolFromFactory(&_BWeightedFactory.CallOpts, pool)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_BWeightedFactory *BWeightedFactoryCallerSession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _BWeightedFactory.Contract.IsPoolFromFactory(&_BWeightedFactory.CallOpts, pool)
}

// Create is a paid mutator transaction binding the contract method 0x80773a93.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] normalizedWeights, address[] rateProviders, uint256 swapFeePercentage, address owner) returns(address)
func (_BWeightedFactory *BWeightedFactoryTransactor) Create(opts *bind.TransactOpts, name string, symbol string, tokens []common.Address, normalizedWeights []*big.Int, rateProviders []common.Address, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _BWeightedFactory.contract.Transact(opts, "create", name, symbol, tokens, normalizedWeights, rateProviders, swapFeePercentage, owner)
}

// Create is a paid mutator transaction binding the contract method 0x80773a93.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] normalizedWeights, address[] rateProviders, uint256 swapFeePercentage, address owner) returns(address)
func (_BWeightedFactory *BWeightedFactorySession) Create(name string, symbol string, tokens []common.Address, normalizedWeights []*big.Int, rateProviders []common.Address, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _BWeightedFactory.Contract.Create(&_BWeightedFactory.TransactOpts, name, symbol, tokens, normalizedWeights, rateProviders, swapFeePercentage, owner)
}

// Create is a paid mutator transaction binding the contract method 0x80773a93.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] normalizedWeights, address[] rateProviders, uint256 swapFeePercentage, address owner) returns(address)
func (_BWeightedFactory *BWeightedFactoryTransactorSession) Create(name string, symbol string, tokens []common.Address, normalizedWeights []*big.Int, rateProviders []common.Address, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _BWeightedFactory.Contract.Create(&_BWeightedFactory.TransactOpts, name, symbol, tokens, normalizedWeights, rateProviders, swapFeePercentage, owner)
}

// Disable is a paid mutator transaction binding the contract method 0x2f2770db.
//
// Solidity: function disable() returns()
func (_BWeightedFactory *BWeightedFactoryTransactor) Disable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedFactory.contract.Transact(opts, "disable")
}

// Disable is a paid mutator transaction binding the contract method 0x2f2770db.
//
// Solidity: function disable() returns()
func (_BWeightedFactory *BWeightedFactorySession) Disable() (*types.Transaction, error) {
	return _BWeightedFactory.Contract.Disable(&_BWeightedFactory.TransactOpts)
}

// Disable is a paid mutator transaction binding the contract method 0x2f2770db.
//
// Solidity: function disable() returns()
func (_BWeightedFactory *BWeightedFactoryTransactorSession) Disable() (*types.Transaction, error) {
	return _BWeightedFactory.Contract.Disable(&_BWeightedFactory.TransactOpts)
}

// BWeightedFactoryFactoryDisabledIterator is returned from FilterFactoryDisabled and is used to iterate over the raw logs and unpacked data for FactoryDisabled events raised by the BWeightedFactory contract.
type BWeightedFactoryFactoryDisabledIterator struct {
	Event *BWeightedFactoryFactoryDisabled // Event containing the contract specifics and raw log

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
func (it *BWeightedFactoryFactoryDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BWeightedFactoryFactoryDisabled)
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
		it.Event = new(BWeightedFactoryFactoryDisabled)
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
func (it *BWeightedFactoryFactoryDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BWeightedFactoryFactoryDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BWeightedFactoryFactoryDisabled represents a FactoryDisabled event raised by the BWeightedFactory contract.
type BWeightedFactoryFactoryDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFactoryDisabled is a free log retrieval operation binding the contract event 0x432acbfd662dbb5d8b378384a67159b47ca9d0f1b79f97cf64cf8585fa362d50.
//
// Solidity: event FactoryDisabled()
func (_BWeightedFactory *BWeightedFactoryFilterer) FilterFactoryDisabled(opts *bind.FilterOpts) (*BWeightedFactoryFactoryDisabledIterator, error) {

	logs, sub, err := _BWeightedFactory.contract.FilterLogs(opts, "FactoryDisabled")
	if err != nil {
		return nil, err
	}
	return &BWeightedFactoryFactoryDisabledIterator{contract: _BWeightedFactory.contract, event: "FactoryDisabled", logs: logs, sub: sub}, nil
}

// WatchFactoryDisabled is a free log subscription operation binding the contract event 0x432acbfd662dbb5d8b378384a67159b47ca9d0f1b79f97cf64cf8585fa362d50.
//
// Solidity: event FactoryDisabled()
func (_BWeightedFactory *BWeightedFactoryFilterer) WatchFactoryDisabled(opts *bind.WatchOpts, sink chan<- *BWeightedFactoryFactoryDisabled) (event.Subscription, error) {

	logs, sub, err := _BWeightedFactory.contract.WatchLogs(opts, "FactoryDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BWeightedFactoryFactoryDisabled)
				if err := _BWeightedFactory.contract.UnpackLog(event, "FactoryDisabled", log); err != nil {
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

// ParseFactoryDisabled is a log parse operation binding the contract event 0x432acbfd662dbb5d8b378384a67159b47ca9d0f1b79f97cf64cf8585fa362d50.
//
// Solidity: event FactoryDisabled()
func (_BWeightedFactory *BWeightedFactoryFilterer) ParseFactoryDisabled(log types.Log) (*BWeightedFactoryFactoryDisabled, error) {
	event := new(BWeightedFactoryFactoryDisabled)
	if err := _BWeightedFactory.contract.UnpackLog(event, "FactoryDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BWeightedFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the BWeightedFactory contract.
type BWeightedFactoryPoolCreatedIterator struct {
	Event *BWeightedFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *BWeightedFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BWeightedFactoryPoolCreated)
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
		it.Event = new(BWeightedFactoryPoolCreated)
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
func (it *BWeightedFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BWeightedFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BWeightedFactoryPoolCreated represents a PoolCreated event raised by the BWeightedFactory contract.
type BWeightedFactoryPoolCreated struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_BWeightedFactory *BWeightedFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, pool []common.Address) (*BWeightedFactoryPoolCreatedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BWeightedFactory.contract.FilterLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return &BWeightedFactoryPoolCreatedIterator{contract: _BWeightedFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_BWeightedFactory *BWeightedFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *BWeightedFactoryPoolCreated, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BWeightedFactory.contract.WatchLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BWeightedFactoryPoolCreated)
				if err := _BWeightedFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_BWeightedFactory *BWeightedFactoryFilterer) ParsePoolCreated(log types.Log) (*BWeightedFactoryPoolCreated, error) {
	event := new(BWeightedFactoryPoolCreated)
	if err := _BWeightedFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
