// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smart_contract

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
	_ = abi.ConvertType
)

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"name\":\"CounterDecremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"name\":\"CounterIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"name\":\"CounterReset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decrementedCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Store *StoreCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Store *StoreSession) Counter() (*big.Int, error) {
	return _Store.Contract.Counter(&_Store.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Store *StoreCallerSession) Counter() (*big.Int, error) {
	return _Store.Contract.Counter(&_Store.CallOpts)
}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() view returns(uint256)
func (_Store *StoreCaller) GetCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() view returns(uint256)
func (_Store *StoreSession) GetCounter() (*big.Int, error) {
	return _Store.Contract.GetCounter(&_Store.CallOpts)
}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() view returns(uint256)
func (_Store *StoreCallerSession) GetCounter() (*big.Int, error) {
	return _Store.Contract.GetCounter(&_Store.CallOpts)
}

// DecrementedCounter is a paid mutator transaction binding the contract method 0x3fb4d94c.
//
// Solidity: function decrementedCounter() returns()
func (_Store *StoreTransactor) DecrementedCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "decrementedCounter")
}

// DecrementedCounter is a paid mutator transaction binding the contract method 0x3fb4d94c.
//
// Solidity: function decrementedCounter() returns()
func (_Store *StoreSession) DecrementedCounter() (*types.Transaction, error) {
	return _Store.Contract.DecrementedCounter(&_Store.TransactOpts)
}

// DecrementedCounter is a paid mutator transaction binding the contract method 0x3fb4d94c.
//
// Solidity: function decrementedCounter() returns()
func (_Store *StoreTransactorSession) DecrementedCounter() (*types.Transaction, error) {
	return _Store.Contract.DecrementedCounter(&_Store.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_Store *StoreTransactor) IncrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "incrementCounter")
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_Store *StoreSession) IncrementCounter() (*types.Transaction, error) {
	return _Store.Contract.IncrementCounter(&_Store.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_Store *StoreTransactorSession) IncrementCounter() (*types.Transaction, error) {
	return _Store.Contract.IncrementCounter(&_Store.TransactOpts)
}

// StoreCounterDecrementedIterator is returned from FilterCounterDecremented and is used to iterate over the raw logs and unpacked data for CounterDecremented events raised by the Store contract.
type StoreCounterDecrementedIterator struct {
	Event *StoreCounterDecremented // Event containing the contract specifics and raw log

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
func (it *StoreCounterDecrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCounterDecremented)
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
		it.Event = new(StoreCounterDecremented)
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
func (it *StoreCounterDecrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCounterDecrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCounterDecremented represents a CounterDecremented event raised by the Store contract.
type StoreCounterDecremented struct {
	NewCounter *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterDecremented is a free log retrieval operation binding the contract event 0x1a00a27c962d5410357331e1a8cffff62058bd0161ad624818df31152f1eeb45.
//
// Solidity: event CounterDecremented(uint256 newCounter)
func (_Store *StoreFilterer) FilterCounterDecremented(opts *bind.FilterOpts) (*StoreCounterDecrementedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CounterDecremented")
	if err != nil {
		return nil, err
	}
	return &StoreCounterDecrementedIterator{contract: _Store.contract, event: "CounterDecremented", logs: logs, sub: sub}, nil
}

// WatchCounterDecremented is a free log subscription operation binding the contract event 0x1a00a27c962d5410357331e1a8cffff62058bd0161ad624818df31152f1eeb45.
//
// Solidity: event CounterDecremented(uint256 newCounter)
func (_Store *StoreFilterer) WatchCounterDecremented(opts *bind.WatchOpts, sink chan<- *StoreCounterDecremented) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "CounterDecremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCounterDecremented)
				if err := _Store.contract.UnpackLog(event, "CounterDecremented", log); err != nil {
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

// ParseCounterDecremented is a log parse operation binding the contract event 0x1a00a27c962d5410357331e1a8cffff62058bd0161ad624818df31152f1eeb45.
//
// Solidity: event CounterDecremented(uint256 newCounter)
func (_Store *StoreFilterer) ParseCounterDecremented(log types.Log) (*StoreCounterDecremented, error) {
	event := new(StoreCounterDecremented)
	if err := _Store.contract.UnpackLog(event, "CounterDecremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreCounterIncrementedIterator is returned from FilterCounterIncremented and is used to iterate over the raw logs and unpacked data for CounterIncremented events raised by the Store contract.
type StoreCounterIncrementedIterator struct {
	Event *StoreCounterIncremented // Event containing the contract specifics and raw log

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
func (it *StoreCounterIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCounterIncremented)
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
		it.Event = new(StoreCounterIncremented)
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
func (it *StoreCounterIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCounterIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCounterIncremented represents a CounterIncremented event raised by the Store contract.
type StoreCounterIncremented struct {
	NewCounter *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterIncremented is a free log retrieval operation binding the contract event 0x3cf8b50771c17d723f2cb711ca7dadde485b222e13c84ba0730a14093fad6d5c.
//
// Solidity: event CounterIncremented(uint256 newCounter)
func (_Store *StoreFilterer) FilterCounterIncremented(opts *bind.FilterOpts) (*StoreCounterIncrementedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CounterIncremented")
	if err != nil {
		return nil, err
	}
	return &StoreCounterIncrementedIterator{contract: _Store.contract, event: "CounterIncremented", logs: logs, sub: sub}, nil
}

// WatchCounterIncremented is a free log subscription operation binding the contract event 0x3cf8b50771c17d723f2cb711ca7dadde485b222e13c84ba0730a14093fad6d5c.
//
// Solidity: event CounterIncremented(uint256 newCounter)
func (_Store *StoreFilterer) WatchCounterIncremented(opts *bind.WatchOpts, sink chan<- *StoreCounterIncremented) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "CounterIncremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCounterIncremented)
				if err := _Store.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
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

// ParseCounterIncremented is a log parse operation binding the contract event 0x3cf8b50771c17d723f2cb711ca7dadde485b222e13c84ba0730a14093fad6d5c.
//
// Solidity: event CounterIncremented(uint256 newCounter)
func (_Store *StoreFilterer) ParseCounterIncremented(log types.Log) (*StoreCounterIncremented, error) {
	event := new(StoreCounterIncremented)
	if err := _Store.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreCounterResetIterator is returned from FilterCounterReset and is used to iterate over the raw logs and unpacked data for CounterReset events raised by the Store contract.
type StoreCounterResetIterator struct {
	Event *StoreCounterReset // Event containing the contract specifics and raw log

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
func (it *StoreCounterResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCounterReset)
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
		it.Event = new(StoreCounterReset)
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
func (it *StoreCounterResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCounterResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCounterReset represents a CounterReset event raised by the Store contract.
type StoreCounterReset struct {
	NewCounter *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterReset is a free log retrieval operation binding the contract event 0x5ee614ee051809d12a47ecce5391d6753965c7559f89ccf50b59ab97cbd0d70f.
//
// Solidity: event CounterReset(uint256 newCounter)
func (_Store *StoreFilterer) FilterCounterReset(opts *bind.FilterOpts) (*StoreCounterResetIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CounterReset")
	if err != nil {
		return nil, err
	}
	return &StoreCounterResetIterator{contract: _Store.contract, event: "CounterReset", logs: logs, sub: sub}, nil
}

// WatchCounterReset is a free log subscription operation binding the contract event 0x5ee614ee051809d12a47ecce5391d6753965c7559f89ccf50b59ab97cbd0d70f.
//
// Solidity: event CounterReset(uint256 newCounter)
func (_Store *StoreFilterer) WatchCounterReset(opts *bind.WatchOpts, sink chan<- *StoreCounterReset) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "CounterReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCounterReset)
				if err := _Store.contract.UnpackLog(event, "CounterReset", log); err != nil {
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

// ParseCounterReset is a log parse operation binding the contract event 0x5ee614ee051809d12a47ecce5391d6753965c7559f89ccf50b59ab97cbd0d70f.
//
// Solidity: event CounterReset(uint256 newCounter)
func (_Store *StoreFilterer) ParseCounterReset(log types.Log) (*StoreCounterReset, error) {
	event := new(StoreCounterReset)
	if err := _Store.contract.UnpackLog(event, "CounterReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
