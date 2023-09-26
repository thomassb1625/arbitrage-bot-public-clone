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

// IPoolSwapStructsSwapRequest is an auto generated low-level Go binding around an user-defined struct.
type IPoolSwapStructsSwapRequest struct {
	Kind            uint8
	TokenIn         common.Address
	TokenOut        common.Address
	Amount          *big.Int
	PoolId          [32]byte
	LastChangeBlock *big.Int
	From            common.Address
	To              common.Address
	UserData        []byte
}

// WeightedPoolNewPoolParams is an auto generated low-level Go binding around an user-defined struct.
type WeightedPoolNewPoolParams struct {
	Name              string
	Symbol            string
	Tokens            []common.Address
	NormalizedWeights []*big.Int
	RateProviders     []common.Address
	AssetManagers     []common.Address
	SwapFeePercentage *big.Int
}

// BWeightedPoolMetaData contains all meta data concerning the BWeightedPool contract.
var BWeightedPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"normalizedWeights\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIRateProvider[]\",\"name\":\"rateProviders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"internalType\":\"structWeightedPool.NewPoolParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"contractIProtocolFeePercentagesProvider\",\"name\":\"protocolFeeProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeePercentage\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeePercentageCacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"RecoveryModeStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getATHRateProduct\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActualSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInvariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastPostJoinExitInvariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNormalizedWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"}],\"name\":\"getProtocolFeePercentageCache\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"contractIProtocolFeesCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateProviders\",\"outputs\":[{\"internalType\":\"contractIRateProvider[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getScalingFactors\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inRecoveryMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onExitPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onJoinPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIPoolSwapStructs.SwapRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenOut\",\"type\":\"uint256\"}],\"name\":\"onSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryJoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setSwapFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateProtocolFeePercentageCache\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BWeightedPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use BWeightedPoolMetaData.ABI instead.
var BWeightedPoolABI = BWeightedPoolMetaData.ABI

// BWeightedPool is an auto generated Go binding around an Ethereum contract.
type BWeightedPool struct {
	BWeightedPoolCaller     // Read-only binding to the contract
	BWeightedPoolTransactor // Write-only binding to the contract
	BWeightedPoolFilterer   // Log filterer for contract events
}

// BWeightedPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BWeightedPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BWeightedPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BWeightedPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BWeightedPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BWeightedPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BWeightedPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BWeightedPoolSession struct {
	Contract     *BWeightedPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BWeightedPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BWeightedPoolCallerSession struct {
	Contract *BWeightedPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BWeightedPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BWeightedPoolTransactorSession struct {
	Contract     *BWeightedPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BWeightedPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BWeightedPoolRaw struct {
	Contract *BWeightedPool // Generic contract binding to access the raw methods on
}

// BWeightedPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BWeightedPoolCallerRaw struct {
	Contract *BWeightedPoolCaller // Generic read-only contract binding to access the raw methods on
}

// BWeightedPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BWeightedPoolTransactorRaw struct {
	Contract *BWeightedPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBWeightedPool creates a new instance of BWeightedPool, bound to a specific deployed contract.
func NewBWeightedPool(address common.Address, backend bind.ContractBackend) (*BWeightedPool, error) {
	contract, err := bindBWeightedPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BWeightedPool{BWeightedPoolCaller: BWeightedPoolCaller{contract: contract}, BWeightedPoolTransactor: BWeightedPoolTransactor{contract: contract}, BWeightedPoolFilterer: BWeightedPoolFilterer{contract: contract}}, nil
}

// NewBWeightedPoolCaller creates a new read-only instance of BWeightedPool, bound to a specific deployed contract.
func NewBWeightedPoolCaller(address common.Address, caller bind.ContractCaller) (*BWeightedPoolCaller, error) {
	contract, err := bindBWeightedPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BWeightedPoolCaller{contract: contract}, nil
}

// NewBWeightedPoolTransactor creates a new write-only instance of BWeightedPool, bound to a specific deployed contract.
func NewBWeightedPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BWeightedPoolTransactor, error) {
	contract, err := bindBWeightedPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BWeightedPoolTransactor{contract: contract}, nil
}

// NewBWeightedPoolFilterer creates a new log filterer instance of BWeightedPool, bound to a specific deployed contract.
func NewBWeightedPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BWeightedPoolFilterer, error) {
	contract, err := bindBWeightedPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BWeightedPoolFilterer{contract: contract}, nil
}

// bindBWeightedPool binds a generic wrapper to an already deployed contract.
func bindBWeightedPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BWeightedPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BWeightedPool *BWeightedPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BWeightedPool.Contract.BWeightedPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BWeightedPool *BWeightedPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedPool.Contract.BWeightedPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BWeightedPool *BWeightedPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BWeightedPool.Contract.BWeightedPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BWeightedPool *BWeightedPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BWeightedPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BWeightedPool *BWeightedPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BWeightedPool *BWeightedPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BWeightedPool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BWeightedPool *BWeightedPoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BWeightedPool *BWeightedPoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BWeightedPool.Contract.DOMAINSEPARATOR(&_BWeightedPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BWeightedPool *BWeightedPoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BWeightedPool.Contract.DOMAINSEPARATOR(&_BWeightedPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BWeightedPool.Contract.Allowance(&_BWeightedPool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BWeightedPool.Contract.Allowance(&_BWeightedPool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BWeightedPool.Contract.BalanceOf(&_BWeightedPool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BWeightedPool.Contract.BalanceOf(&_BWeightedPool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BWeightedPool *BWeightedPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BWeightedPool *BWeightedPoolSession) Decimals() (uint8, error) {
	return _BWeightedPool.Contract.Decimals(&_BWeightedPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BWeightedPool *BWeightedPoolCallerSession) Decimals() (uint8, error) {
	return _BWeightedPool.Contract.Decimals(&_BWeightedPool.CallOpts)
}

// GetATHRateProduct is a free data retrieval call binding the contract method 0x23ef89ed.
//
// Solidity: function getATHRateProduct() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) GetATHRateProduct(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getATHRateProduct")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetATHRateProduct is a free data retrieval call binding the contract method 0x23ef89ed.
//
// Solidity: function getATHRateProduct() view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) GetATHRateProduct() (*big.Int, error) {
	return _BWeightedPool.Contract.GetATHRateProduct(&_BWeightedPool.CallOpts)
}

// GetATHRateProduct is a free data retrieval call binding the contract method 0x23ef89ed.
//
// Solidity: function getATHRateProduct() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) GetATHRateProduct() (*big.Int, error) {
	return _BWeightedPool.Contract.GetATHRateProduct(&_BWeightedPool.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BWeightedPool *BWeightedPoolCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BWeightedPool *BWeightedPoolSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BWeightedPool.Contract.GetActionId(&_BWeightedPool.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BWeightedPool *BWeightedPoolCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BWeightedPool.Contract.GetActionId(&_BWeightedPool.CallOpts, selector)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) GetActualSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getActualSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) GetActualSupply() (*big.Int, error) {
	return _BWeightedPool.Contract.GetActualSupply(&_BWeightedPool.CallOpts)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) GetActualSupply() (*big.Int, error) {
	return _BWeightedPool.Contract.GetActualSupply(&_BWeightedPool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BWeightedPool *BWeightedPoolCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BWeightedPool *BWeightedPoolSession) GetAuthorizer() (common.Address, error) {
	return _BWeightedPool.Contract.GetAuthorizer(&_BWeightedPool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BWeightedPool *BWeightedPoolCallerSession) GetAuthorizer() (common.Address, error) {
	return _BWeightedPool.Contract.GetAuthorizer(&_BWeightedPool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BWeightedPool *BWeightedPoolCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BWeightedPool *BWeightedPoolSession) GetDomainSeparator() ([32]byte, error) {
	return _BWeightedPool.Contract.GetDomainSeparator(&_BWeightedPool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BWeightedPool *BWeightedPoolCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _BWeightedPool.Contract.GetDomainSeparator(&_BWeightedPool.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) GetInvariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getInvariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) GetInvariant() (*big.Int, error) {
	return _BWeightedPool.Contract.GetInvariant(&_BWeightedPool.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) GetInvariant() (*big.Int, error) {
	return _BWeightedPool.Contract.GetInvariant(&_BWeightedPool.CallOpts)
}

// GetLastPostJoinExitInvariant is a free data retrieval call binding the contract method 0xb1096278.
//
// Solidity: function getLastPostJoinExitInvariant() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) GetLastPostJoinExitInvariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getLastPostJoinExitInvariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPostJoinExitInvariant is a free data retrieval call binding the contract method 0xb1096278.
//
// Solidity: function getLastPostJoinExitInvariant() view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) GetLastPostJoinExitInvariant() (*big.Int, error) {
	return _BWeightedPool.Contract.GetLastPostJoinExitInvariant(&_BWeightedPool.CallOpts)
}

// GetLastPostJoinExitInvariant is a free data retrieval call binding the contract method 0xb1096278.
//
// Solidity: function getLastPostJoinExitInvariant() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) GetLastPostJoinExitInvariant() (*big.Int, error) {
	return _BWeightedPool.Contract.GetLastPostJoinExitInvariant(&_BWeightedPool.CallOpts)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) GetNextNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getNextNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BWeightedPool.Contract.GetNextNonce(&_BWeightedPool.CallOpts, account)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BWeightedPool.Contract.GetNextNonce(&_BWeightedPool.CallOpts, account)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BWeightedPool *BWeightedPoolCaller) GetNormalizedWeights(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getNormalizedWeights")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BWeightedPool *BWeightedPoolSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BWeightedPool.Contract.GetNormalizedWeights(&_BWeightedPool.CallOpts)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BWeightedPool *BWeightedPoolCallerSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BWeightedPool.Contract.GetNormalizedWeights(&_BWeightedPool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BWeightedPool *BWeightedPoolCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BWeightedPool *BWeightedPoolSession) GetOwner() (common.Address, error) {
	return _BWeightedPool.Contract.GetOwner(&_BWeightedPool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BWeightedPool *BWeightedPoolCallerSession) GetOwner() (common.Address, error) {
	return _BWeightedPool.Contract.GetOwner(&_BWeightedPool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BWeightedPool *BWeightedPoolCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getPausedState")

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
func (_BWeightedPool *BWeightedPoolSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BWeightedPool.Contract.GetPausedState(&_BWeightedPool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BWeightedPool *BWeightedPoolCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BWeightedPool.Contract.GetPausedState(&_BWeightedPool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BWeightedPool *BWeightedPoolCaller) GetPoolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getPoolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BWeightedPool *BWeightedPoolSession) GetPoolId() ([32]byte, error) {
	return _BWeightedPool.Contract.GetPoolId(&_BWeightedPool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BWeightedPool *BWeightedPoolCallerSession) GetPoolId() ([32]byte, error) {
	return _BWeightedPool.Contract.GetPoolId(&_BWeightedPool.CallOpts)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) GetProtocolFeePercentageCache(opts *bind.CallOpts, feeType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getProtocolFeePercentageCache", feeType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _BWeightedPool.Contract.GetProtocolFeePercentageCache(&_BWeightedPool.CallOpts, feeType)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _BWeightedPool.Contract.GetProtocolFeePercentageCache(&_BWeightedPool.CallOpts, feeType)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BWeightedPool *BWeightedPoolCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BWeightedPool *BWeightedPoolSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BWeightedPool.Contract.GetProtocolFeesCollector(&_BWeightedPool.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BWeightedPool *BWeightedPoolCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BWeightedPool.Contract.GetProtocolFeesCollector(&_BWeightedPool.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BWeightedPool *BWeightedPoolCaller) GetRateProviders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getRateProviders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BWeightedPool *BWeightedPoolSession) GetRateProviders() ([]common.Address, error) {
	return _BWeightedPool.Contract.GetRateProviders(&_BWeightedPool.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BWeightedPool *BWeightedPoolCallerSession) GetRateProviders() ([]common.Address, error) {
	return _BWeightedPool.Contract.GetRateProviders(&_BWeightedPool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BWeightedPool *BWeightedPoolCaller) GetScalingFactors(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getScalingFactors")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BWeightedPool *BWeightedPoolSession) GetScalingFactors() ([]*big.Int, error) {
	return _BWeightedPool.Contract.GetScalingFactors(&_BWeightedPool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BWeightedPool *BWeightedPoolCallerSession) GetScalingFactors() ([]*big.Int, error) {
	return _BWeightedPool.Contract.GetScalingFactors(&_BWeightedPool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) GetSwapFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getSwapFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BWeightedPool.Contract.GetSwapFeePercentage(&_BWeightedPool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BWeightedPool.Contract.GetSwapFeePercentage(&_BWeightedPool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BWeightedPool *BWeightedPoolCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BWeightedPool *BWeightedPoolSession) GetVault() (common.Address, error) {
	return _BWeightedPool.Contract.GetVault(&_BWeightedPool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BWeightedPool *BWeightedPoolCallerSession) GetVault() (common.Address, error) {
	return _BWeightedPool.Contract.GetVault(&_BWeightedPool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BWeightedPool *BWeightedPoolCaller) InRecoveryMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "inRecoveryMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BWeightedPool *BWeightedPoolSession) InRecoveryMode() (bool, error) {
	return _BWeightedPool.Contract.InRecoveryMode(&_BWeightedPool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BWeightedPool *BWeightedPoolCallerSession) InRecoveryMode() (bool, error) {
	return _BWeightedPool.Contract.InRecoveryMode(&_BWeightedPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BWeightedPool *BWeightedPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BWeightedPool *BWeightedPoolSession) Name() (string, error) {
	return _BWeightedPool.Contract.Name(&_BWeightedPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BWeightedPool *BWeightedPoolCallerSession) Name() (string, error) {
	return _BWeightedPool.Contract.Name(&_BWeightedPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BWeightedPool.Contract.Nonces(&_BWeightedPool.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BWeightedPool.Contract.Nonces(&_BWeightedPool.CallOpts, owner)
}

// QueryExit is a free data retrieval call binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptIn, uint256[] amountsOut)
func (_BWeightedPool *BWeightedPoolCaller) QueryExit(opts *bind.CallOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptIn      *big.Int
	AmountsOut []*big.Int
}, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "queryExit", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)

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
func (_BWeightedPool *BWeightedPoolSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptIn      *big.Int
	AmountsOut []*big.Int
}, error) {
	return _BWeightedPool.Contract.QueryExit(&_BWeightedPool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a free data retrieval call binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptIn, uint256[] amountsOut)
func (_BWeightedPool *BWeightedPoolCallerSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptIn      *big.Int
	AmountsOut []*big.Int
}, error) {
	return _BWeightedPool.Contract.QueryExit(&_BWeightedPool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a free data retrieval call binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptOut, uint256[] amountsIn)
func (_BWeightedPool *BWeightedPoolCaller) QueryJoin(opts *bind.CallOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptOut    *big.Int
	AmountsIn []*big.Int
}, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "queryJoin", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)

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
func (_BWeightedPool *BWeightedPoolSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptOut    *big.Int
	AmountsIn []*big.Int
}, error) {
	return _BWeightedPool.Contract.QueryJoin(&_BWeightedPool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a free data retrieval call binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) view returns(uint256 bptOut, uint256[] amountsIn)
func (_BWeightedPool *BWeightedPoolCallerSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (struct {
	BptOut    *big.Int
	AmountsIn []*big.Int
}, error) {
	return _BWeightedPool.Contract.QueryJoin(&_BWeightedPool.CallOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BWeightedPool *BWeightedPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BWeightedPool *BWeightedPoolSession) Symbol() (string, error) {
	return _BWeightedPool.Contract.Symbol(&_BWeightedPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BWeightedPool *BWeightedPoolCallerSession) Symbol() (string, error) {
	return _BWeightedPool.Contract.Symbol(&_BWeightedPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BWeightedPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) TotalSupply() (*big.Int, error) {
	return _BWeightedPool.Contract.TotalSupply(&_BWeightedPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BWeightedPool *BWeightedPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _BWeightedPool.Contract.TotalSupply(&_BWeightedPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.Approve(&_BWeightedPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.Approve(&_BWeightedPool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.DecreaseAllowance(&_BWeightedPool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.DecreaseAllowance(&_BWeightedPool.TransactOpts, spender, amount)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BWeightedPool *BWeightedPoolTransactor) DisableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "disableRecoveryMode")
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BWeightedPool *BWeightedPoolSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BWeightedPool.Contract.DisableRecoveryMode(&_BWeightedPool.TransactOpts)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BWeightedPool *BWeightedPoolTransactorSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BWeightedPool.Contract.DisableRecoveryMode(&_BWeightedPool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BWeightedPool *BWeightedPoolTransactor) EnableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "enableRecoveryMode")
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BWeightedPool *BWeightedPoolSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BWeightedPool.Contract.EnableRecoveryMode(&_BWeightedPool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BWeightedPool *BWeightedPoolTransactorSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BWeightedPool.Contract.EnableRecoveryMode(&_BWeightedPool.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BWeightedPool *BWeightedPoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.IncreaseAllowance(&_BWeightedPool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.IncreaseAllowance(&_BWeightedPool.TransactOpts, spender, addedValue)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BWeightedPool *BWeightedPoolTransactor) OnExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "onExitPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BWeightedPool *BWeightedPoolSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BWeightedPool.Contract.OnExitPool(&_BWeightedPool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BWeightedPool *BWeightedPoolTransactorSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BWeightedPool.Contract.OnExitPool(&_BWeightedPool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BWeightedPool *BWeightedPoolTransactor) OnJoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "onJoinPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BWeightedPool *BWeightedPoolSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BWeightedPool.Contract.OnJoinPool(&_BWeightedPool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BWeightedPool *BWeightedPoolTransactorSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BWeightedPool.Contract.OnJoinPool(&_BWeightedPool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BWeightedPool *BWeightedPoolTransactor) OnSwap(opts *bind.TransactOpts, request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "onSwap", request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BWeightedPool *BWeightedPoolSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.OnSwap(&_BWeightedPool.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BWeightedPool *BWeightedPoolTransactorSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.OnSwap(&_BWeightedPool.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BWeightedPool *BWeightedPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BWeightedPool *BWeightedPoolSession) Pause() (*types.Transaction, error) {
	return _BWeightedPool.Contract.Pause(&_BWeightedPool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BWeightedPool *BWeightedPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _BWeightedPool.Contract.Pause(&_BWeightedPool.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BWeightedPool *BWeightedPoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BWeightedPool *BWeightedPoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BWeightedPool.Contract.Permit(&_BWeightedPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BWeightedPool *BWeightedPoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BWeightedPool.Contract.Permit(&_BWeightedPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BWeightedPool *BWeightedPoolTransactor) SetSwapFeePercentage(opts *bind.TransactOpts, swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "setSwapFeePercentage", swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BWeightedPool *BWeightedPoolSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.SetSwapFeePercentage(&_BWeightedPool.TransactOpts, swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BWeightedPool *BWeightedPoolTransactorSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.SetSwapFeePercentage(&_BWeightedPool.TransactOpts, swapFeePercentage)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.Transfer(&_BWeightedPool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.Transfer(&_BWeightedPool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.TransferFrom(&_BWeightedPool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BWeightedPool *BWeightedPoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BWeightedPool.Contract.TransferFrom(&_BWeightedPool.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BWeightedPool *BWeightedPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BWeightedPool *BWeightedPoolSession) Unpause() (*types.Transaction, error) {
	return _BWeightedPool.Contract.Unpause(&_BWeightedPool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BWeightedPool *BWeightedPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _BWeightedPool.Contract.Unpause(&_BWeightedPool.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BWeightedPool *BWeightedPoolTransactor) UpdateProtocolFeePercentageCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BWeightedPool.contract.Transact(opts, "updateProtocolFeePercentageCache")
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BWeightedPool *BWeightedPoolSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _BWeightedPool.Contract.UpdateProtocolFeePercentageCache(&_BWeightedPool.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BWeightedPool *BWeightedPoolTransactorSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _BWeightedPool.Contract.UpdateProtocolFeePercentageCache(&_BWeightedPool.TransactOpts)
}

// BWeightedPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BWeightedPool contract.
type BWeightedPoolApprovalIterator struct {
	Event *BWeightedPoolApproval // Event containing the contract specifics and raw log

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
func (it *BWeightedPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BWeightedPoolApproval)
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
		it.Event = new(BWeightedPoolApproval)
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
func (it *BWeightedPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BWeightedPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BWeightedPoolApproval represents a Approval event raised by the BWeightedPool contract.
type BWeightedPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BWeightedPool *BWeightedPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BWeightedPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BWeightedPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BWeightedPoolApprovalIterator{contract: _BWeightedPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BWeightedPool *BWeightedPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BWeightedPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BWeightedPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BWeightedPoolApproval)
				if err := _BWeightedPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BWeightedPool *BWeightedPoolFilterer) ParseApproval(log types.Log) (*BWeightedPoolApproval, error) {
	event := new(BWeightedPoolApproval)
	if err := _BWeightedPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BWeightedPoolPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the BWeightedPool contract.
type BWeightedPoolPausedStateChangedIterator struct {
	Event *BWeightedPoolPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *BWeightedPoolPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BWeightedPoolPausedStateChanged)
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
		it.Event = new(BWeightedPoolPausedStateChanged)
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
func (it *BWeightedPoolPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BWeightedPoolPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BWeightedPoolPausedStateChanged represents a PausedStateChanged event raised by the BWeightedPool contract.
type BWeightedPoolPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BWeightedPool *BWeightedPoolFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*BWeightedPoolPausedStateChangedIterator, error) {

	logs, sub, err := _BWeightedPool.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &BWeightedPoolPausedStateChangedIterator{contract: _BWeightedPool.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BWeightedPool *BWeightedPoolFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *BWeightedPoolPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _BWeightedPool.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BWeightedPoolPausedStateChanged)
				if err := _BWeightedPool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
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
func (_BWeightedPool *BWeightedPoolFilterer) ParsePausedStateChanged(log types.Log) (*BWeightedPoolPausedStateChanged, error) {
	event := new(BWeightedPoolPausedStateChanged)
	if err := _BWeightedPool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BWeightedPoolProtocolFeePercentageCacheUpdatedIterator is returned from FilterProtocolFeePercentageCacheUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeePercentageCacheUpdated events raised by the BWeightedPool contract.
type BWeightedPoolProtocolFeePercentageCacheUpdatedIterator struct {
	Event *BWeightedPoolProtocolFeePercentageCacheUpdated // Event containing the contract specifics and raw log

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
func (it *BWeightedPoolProtocolFeePercentageCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BWeightedPoolProtocolFeePercentageCacheUpdated)
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
		it.Event = new(BWeightedPoolProtocolFeePercentageCacheUpdated)
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
func (it *BWeightedPoolProtocolFeePercentageCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BWeightedPoolProtocolFeePercentageCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BWeightedPoolProtocolFeePercentageCacheUpdated represents a ProtocolFeePercentageCacheUpdated event raised by the BWeightedPool contract.
type BWeightedPoolProtocolFeePercentageCacheUpdated struct {
	FeeType               *big.Int
	ProtocolFeePercentage *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeePercentageCacheUpdated is a free log retrieval operation binding the contract event 0x6bfb689528fa96ec1ad670ad6d6064be1ae96bfd5d2ee35c837fd0fe0c11959a.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(uint256 indexed feeType, uint256 protocolFeePercentage)
func (_BWeightedPool *BWeightedPoolFilterer) FilterProtocolFeePercentageCacheUpdated(opts *bind.FilterOpts, feeType []*big.Int) (*BWeightedPoolProtocolFeePercentageCacheUpdatedIterator, error) {

	var feeTypeRule []interface{}
	for _, feeTypeItem := range feeType {
		feeTypeRule = append(feeTypeRule, feeTypeItem)
	}

	logs, sub, err := _BWeightedPool.contract.FilterLogs(opts, "ProtocolFeePercentageCacheUpdated", feeTypeRule)
	if err != nil {
		return nil, err
	}
	return &BWeightedPoolProtocolFeePercentageCacheUpdatedIterator{contract: _BWeightedPool.contract, event: "ProtocolFeePercentageCacheUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeePercentageCacheUpdated is a free log subscription operation binding the contract event 0x6bfb689528fa96ec1ad670ad6d6064be1ae96bfd5d2ee35c837fd0fe0c11959a.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(uint256 indexed feeType, uint256 protocolFeePercentage)
func (_BWeightedPool *BWeightedPoolFilterer) WatchProtocolFeePercentageCacheUpdated(opts *bind.WatchOpts, sink chan<- *BWeightedPoolProtocolFeePercentageCacheUpdated, feeType []*big.Int) (event.Subscription, error) {

	var feeTypeRule []interface{}
	for _, feeTypeItem := range feeType {
		feeTypeRule = append(feeTypeRule, feeTypeItem)
	}

	logs, sub, err := _BWeightedPool.contract.WatchLogs(opts, "ProtocolFeePercentageCacheUpdated", feeTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BWeightedPoolProtocolFeePercentageCacheUpdated)
				if err := _BWeightedPool.contract.UnpackLog(event, "ProtocolFeePercentageCacheUpdated", log); err != nil {
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

// ParseProtocolFeePercentageCacheUpdated is a log parse operation binding the contract event 0x6bfb689528fa96ec1ad670ad6d6064be1ae96bfd5d2ee35c837fd0fe0c11959a.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(uint256 indexed feeType, uint256 protocolFeePercentage)
func (_BWeightedPool *BWeightedPoolFilterer) ParseProtocolFeePercentageCacheUpdated(log types.Log) (*BWeightedPoolProtocolFeePercentageCacheUpdated, error) {
	event := new(BWeightedPoolProtocolFeePercentageCacheUpdated)
	if err := _BWeightedPool.contract.UnpackLog(event, "ProtocolFeePercentageCacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BWeightedPoolRecoveryModeStateChangedIterator is returned from FilterRecoveryModeStateChanged and is used to iterate over the raw logs and unpacked data for RecoveryModeStateChanged events raised by the BWeightedPool contract.
type BWeightedPoolRecoveryModeStateChangedIterator struct {
	Event *BWeightedPoolRecoveryModeStateChanged // Event containing the contract specifics and raw log

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
func (it *BWeightedPoolRecoveryModeStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BWeightedPoolRecoveryModeStateChanged)
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
		it.Event = new(BWeightedPoolRecoveryModeStateChanged)
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
func (it *BWeightedPoolRecoveryModeStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BWeightedPoolRecoveryModeStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BWeightedPoolRecoveryModeStateChanged represents a RecoveryModeStateChanged event raised by the BWeightedPool contract.
type BWeightedPoolRecoveryModeStateChanged struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecoveryModeStateChanged is a free log retrieval operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BWeightedPool *BWeightedPoolFilterer) FilterRecoveryModeStateChanged(opts *bind.FilterOpts) (*BWeightedPoolRecoveryModeStateChangedIterator, error) {

	logs, sub, err := _BWeightedPool.contract.FilterLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return &BWeightedPoolRecoveryModeStateChangedIterator{contract: _BWeightedPool.contract, event: "RecoveryModeStateChanged", logs: logs, sub: sub}, nil
}

// WatchRecoveryModeStateChanged is a free log subscription operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BWeightedPool *BWeightedPoolFilterer) WatchRecoveryModeStateChanged(opts *bind.WatchOpts, sink chan<- *BWeightedPoolRecoveryModeStateChanged) (event.Subscription, error) {

	logs, sub, err := _BWeightedPool.contract.WatchLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BWeightedPoolRecoveryModeStateChanged)
				if err := _BWeightedPool.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
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
func (_BWeightedPool *BWeightedPoolFilterer) ParseRecoveryModeStateChanged(log types.Log) (*BWeightedPoolRecoveryModeStateChanged, error) {
	event := new(BWeightedPoolRecoveryModeStateChanged)
	if err := _BWeightedPool.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BWeightedPoolSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the BWeightedPool contract.
type BWeightedPoolSwapFeePercentageChangedIterator struct {
	Event *BWeightedPoolSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *BWeightedPoolSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BWeightedPoolSwapFeePercentageChanged)
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
		it.Event = new(BWeightedPoolSwapFeePercentageChanged)
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
func (it *BWeightedPoolSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BWeightedPoolSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BWeightedPoolSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the BWeightedPool contract.
type BWeightedPoolSwapFeePercentageChanged struct {
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BWeightedPool *BWeightedPoolFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts) (*BWeightedPoolSwapFeePercentageChangedIterator, error) {

	logs, sub, err := _BWeightedPool.contract.FilterLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &BWeightedPoolSwapFeePercentageChangedIterator{contract: _BWeightedPool.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BWeightedPool *BWeightedPoolFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *BWeightedPoolSwapFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _BWeightedPool.contract.WatchLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BWeightedPoolSwapFeePercentageChanged)
				if err := _BWeightedPool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
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
func (_BWeightedPool *BWeightedPoolFilterer) ParseSwapFeePercentageChanged(log types.Log) (*BWeightedPoolSwapFeePercentageChanged, error) {
	event := new(BWeightedPoolSwapFeePercentageChanged)
	if err := _BWeightedPool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BWeightedPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BWeightedPool contract.
type BWeightedPoolTransferIterator struct {
	Event *BWeightedPoolTransfer // Event containing the contract specifics and raw log

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
func (it *BWeightedPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BWeightedPoolTransfer)
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
		it.Event = new(BWeightedPoolTransfer)
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
func (it *BWeightedPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BWeightedPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BWeightedPoolTransfer represents a Transfer event raised by the BWeightedPool contract.
type BWeightedPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BWeightedPool *BWeightedPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BWeightedPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BWeightedPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BWeightedPoolTransferIterator{contract: _BWeightedPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BWeightedPool *BWeightedPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BWeightedPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BWeightedPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BWeightedPoolTransfer)
				if err := _BWeightedPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BWeightedPool *BWeightedPoolFilterer) ParseTransfer(log types.Log) (*BWeightedPoolTransfer, error) {
	event := new(BWeightedPoolTransfer)
	if err := _BWeightedPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
