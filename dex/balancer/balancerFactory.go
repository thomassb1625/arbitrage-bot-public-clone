// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancer

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

// BalancerFactoryMetaData contains all meta data concerning the BalancerFactory contract.
var BalancerFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blabs\",\"type\":\"address\"}],\"name\":\"LOG_BLABS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LOG_NEW_POOL\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractBPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBLabs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"isBPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newBPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setBLabs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BalancerFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerFactoryMetaData.ABI instead.
var BalancerFactoryABI = BalancerFactoryMetaData.ABI

// BalancerFactory is an auto generated Go binding around an Ethereum contract.
type BalancerFactory struct {
	BalancerFactoryCaller     // Read-only binding to the contract
	BalancerFactoryTransactor // Write-only binding to the contract
	BalancerFactoryFilterer   // Log filterer for contract events
}

// BalancerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerFactorySession struct {
	Contract     *BalancerFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerFactoryCallerSession struct {
	Contract *BalancerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BalancerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerFactoryTransactorSession struct {
	Contract     *BalancerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BalancerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerFactoryRaw struct {
	Contract *BalancerFactory // Generic contract binding to access the raw methods on
}

// BalancerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerFactoryCallerRaw struct {
	Contract *BalancerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerFactoryTransactorRaw struct {
	Contract *BalancerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerFactory creates a new instance of BalancerFactory, bound to a specific deployed contract.
func NewBalancerFactory(address common.Address, backend bind.ContractBackend) (*BalancerFactory, error) {
	contract, err := bindBalancerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerFactory{BalancerFactoryCaller: BalancerFactoryCaller{contract: contract}, BalancerFactoryTransactor: BalancerFactoryTransactor{contract: contract}, BalancerFactoryFilterer: BalancerFactoryFilterer{contract: contract}}, nil
}

// NewBalancerFactoryCaller creates a new read-only instance of BalancerFactory, bound to a specific deployed contract.
func NewBalancerFactoryCaller(address common.Address, caller bind.ContractCaller) (*BalancerFactoryCaller, error) {
	contract, err := bindBalancerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerFactoryCaller{contract: contract}, nil
}

// NewBalancerFactoryTransactor creates a new write-only instance of BalancerFactory, bound to a specific deployed contract.
func NewBalancerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerFactoryTransactor, error) {
	contract, err := bindBalancerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerFactoryTransactor{contract: contract}, nil
}

// NewBalancerFactoryFilterer creates a new log filterer instance of BalancerFactory, bound to a specific deployed contract.
func NewBalancerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerFactoryFilterer, error) {
	contract, err := bindBalancerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerFactoryFilterer{contract: contract}, nil
}

// bindBalancerFactory binds a generic wrapper to an already deployed contract.
func bindBalancerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerFactory *BalancerFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerFactory.Contract.BalancerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerFactory *BalancerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerFactory.Contract.BalancerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerFactory *BalancerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerFactory.Contract.BalancerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerFactory *BalancerFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerFactory *BalancerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerFactory *BalancerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerFactory.Contract.contract.Transact(opts, method, params...)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_BalancerFactory *BalancerFactoryCaller) GetBLabs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerFactory.contract.Call(opts, &out, "getBLabs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_BalancerFactory *BalancerFactorySession) GetBLabs() (common.Address, error) {
	return _BalancerFactory.Contract.GetBLabs(&_BalancerFactory.CallOpts)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_BalancerFactory *BalancerFactoryCallerSession) GetBLabs() (common.Address, error) {
	return _BalancerFactory.Contract.GetBLabs(&_BalancerFactory.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BalancerFactory *BalancerFactoryCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerFactory.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BalancerFactory *BalancerFactorySession) GetColor() ([32]byte, error) {
	return _BalancerFactory.Contract.GetColor(&_BalancerFactory.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BalancerFactory *BalancerFactoryCallerSession) GetColor() ([32]byte, error) {
	return _BalancerFactory.Contract.GetColor(&_BalancerFactory.CallOpts)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_BalancerFactory *BalancerFactoryCaller) IsBPool(opts *bind.CallOpts, b common.Address) (bool, error) {
	var out []interface{}
	err := _BalancerFactory.contract.Call(opts, &out, "isBPool", b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_BalancerFactory *BalancerFactorySession) IsBPool(b common.Address) (bool, error) {
	return _BalancerFactory.Contract.IsBPool(&_BalancerFactory.CallOpts, b)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_BalancerFactory *BalancerFactoryCallerSession) IsBPool(b common.Address) (bool, error) {
	return _BalancerFactory.Contract.IsBPool(&_BalancerFactory.CallOpts, b)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_BalancerFactory *BalancerFactoryTransactor) Collect(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _BalancerFactory.contract.Transact(opts, "collect", pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_BalancerFactory *BalancerFactorySession) Collect(pool common.Address) (*types.Transaction, error) {
	return _BalancerFactory.Contract.Collect(&_BalancerFactory.TransactOpts, pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_BalancerFactory *BalancerFactoryTransactorSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _BalancerFactory.Contract.Collect(&_BalancerFactory.TransactOpts, pool)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_BalancerFactory *BalancerFactoryTransactor) NewBPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerFactory.contract.Transact(opts, "newBPool")
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_BalancerFactory *BalancerFactorySession) NewBPool() (*types.Transaction, error) {
	return _BalancerFactory.Contract.NewBPool(&_BalancerFactory.TransactOpts)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_BalancerFactory *BalancerFactoryTransactorSession) NewBPool() (*types.Transaction, error) {
	return _BalancerFactory.Contract.NewBPool(&_BalancerFactory.TransactOpts)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_BalancerFactory *BalancerFactoryTransactor) SetBLabs(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _BalancerFactory.contract.Transact(opts, "setBLabs", b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_BalancerFactory *BalancerFactorySession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _BalancerFactory.Contract.SetBLabs(&_BalancerFactory.TransactOpts, b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_BalancerFactory *BalancerFactoryTransactorSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _BalancerFactory.Contract.SetBLabs(&_BalancerFactory.TransactOpts, b)
}

// BalancerFactoryLOGBLABSIterator is returned from FilterLOGBLABS and is used to iterate over the raw logs and unpacked data for LOGBLABS events raised by the BalancerFactory contract.
type BalancerFactoryLOGBLABSIterator struct {
	Event *BalancerFactoryLOGBLABS // Event containing the contract specifics and raw log

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
func (it *BalancerFactoryLOGBLABSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerFactoryLOGBLABS)
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
		it.Event = new(BalancerFactoryLOGBLABS)
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
func (it *BalancerFactoryLOGBLABSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerFactoryLOGBLABSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerFactoryLOGBLABS represents a LOGBLABS event raised by the BalancerFactory contract.
type BalancerFactoryLOGBLABS struct {
	Caller common.Address
	Blabs  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGBLABS is a free log retrieval operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_BalancerFactory *BalancerFactoryFilterer) FilterLOGBLABS(opts *bind.FilterOpts, caller []common.Address, blabs []common.Address) (*BalancerFactoryLOGBLABSIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _BalancerFactory.contract.FilterLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return &BalancerFactoryLOGBLABSIterator{contract: _BalancerFactory.contract, event: "LOG_BLABS", logs: logs, sub: sub}, nil
}

// WatchLOGBLABS is a free log subscription operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_BalancerFactory *BalancerFactoryFilterer) WatchLOGBLABS(opts *bind.WatchOpts, sink chan<- *BalancerFactoryLOGBLABS, caller []common.Address, blabs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _BalancerFactory.contract.WatchLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerFactoryLOGBLABS)
				if err := _BalancerFactory.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
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

// ParseLOGBLABS is a log parse operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_BalancerFactory *BalancerFactoryFilterer) ParseLOGBLABS(log types.Log) (*BalancerFactoryLOGBLABS, error) {
	event := new(BalancerFactoryLOGBLABS)
	if err := _BalancerFactory.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerFactoryLOGNEWPOOLIterator is returned from FilterLOGNEWPOOL and is used to iterate over the raw logs and unpacked data for LOGNEWPOOL events raised by the BalancerFactory contract.
type BalancerFactoryLOGNEWPOOLIterator struct {
	Event *BalancerFactoryLOGNEWPOOL // Event containing the contract specifics and raw log

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
func (it *BalancerFactoryLOGNEWPOOLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerFactoryLOGNEWPOOL)
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
		it.Event = new(BalancerFactoryLOGNEWPOOL)
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
func (it *BalancerFactoryLOGNEWPOOLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerFactoryLOGNEWPOOLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerFactoryLOGNEWPOOL represents a LOGNEWPOOL event raised by the BalancerFactory contract.
type BalancerFactoryLOGNEWPOOL struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGNEWPOOL is a free log retrieval operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_BalancerFactory *BalancerFactoryFilterer) FilterLOGNEWPOOL(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*BalancerFactoryLOGNEWPOOLIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BalancerFactory.contract.FilterLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &BalancerFactoryLOGNEWPOOLIterator{contract: _BalancerFactory.contract, event: "LOG_NEW_POOL", logs: logs, sub: sub}, nil
}

// WatchLOGNEWPOOL is a free log subscription operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_BalancerFactory *BalancerFactoryFilterer) WatchLOGNEWPOOL(opts *bind.WatchOpts, sink chan<- *BalancerFactoryLOGNEWPOOL, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BalancerFactory.contract.WatchLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerFactoryLOGNEWPOOL)
				if err := _BalancerFactory.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
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

// ParseLOGNEWPOOL is a log parse operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_BalancerFactory *BalancerFactoryFilterer) ParseLOGNEWPOOL(log types.Log) (*BalancerFactoryLOGNEWPOOL, error) {
	event := new(BalancerFactoryLOGNEWPOOL)
	if err := _BalancerFactory.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
