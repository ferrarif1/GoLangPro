// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TokenLockMetaData contains all meta data concerning the TokenLock contract.
var TokenLockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UnlockEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"lastUnlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockedToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TokenLockABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenLockMetaData.ABI instead.
var TokenLockABI = TokenLockMetaData.ABI

// TokenLock is an auto generated Go binding around an Ethereum contract.
type TokenLock struct {
	TokenLockCaller     // Read-only binding to the contract
	TokenLockTransactor // Write-only binding to the contract
	TokenLockFilterer   // Log filterer for contract events
}

// TokenLockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenLockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenLockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenLockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenLockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenLockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenLockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenLockSession struct {
	Contract     *TokenLock        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenLockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenLockCallerSession struct {
	Contract *TokenLockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenLockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenLockTransactorSession struct {
	Contract     *TokenLockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenLockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenLockRaw struct {
	Contract *TokenLock // Generic contract binding to access the raw methods on
}

// TokenLockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenLockCallerRaw struct {
	Contract *TokenLockCaller // Generic read-only contract binding to access the raw methods on
}

// TokenLockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenLockTransactorRaw struct {
	Contract *TokenLockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenLock creates a new instance of TokenLock, bound to a specific deployed contract.
func NewTokenLock(address common.Address, backend bind.ContractBackend) (*TokenLock, error) {
	contract, err := bindTokenLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenLock{TokenLockCaller: TokenLockCaller{contract: contract}, TokenLockTransactor: TokenLockTransactor{contract: contract}, TokenLockFilterer: TokenLockFilterer{contract: contract}}, nil
}

// NewTokenLockCaller creates a new read-only instance of TokenLock, bound to a specific deployed contract.
func NewTokenLockCaller(address common.Address, caller bind.ContractCaller) (*TokenLockCaller, error) {
	contract, err := bindTokenLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenLockCaller{contract: contract}, nil
}

// NewTokenLockTransactor creates a new write-only instance of TokenLock, bound to a specific deployed contract.
func NewTokenLockTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenLockTransactor, error) {
	contract, err := bindTokenLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenLockTransactor{contract: contract}, nil
}

// NewTokenLockFilterer creates a new log filterer instance of TokenLock, bound to a specific deployed contract.
func NewTokenLockFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenLockFilterer, error) {
	contract, err := bindTokenLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenLockFilterer{contract: contract}, nil
}

// bindTokenLock binds a generic wrapper to an already deployed contract.
func bindTokenLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenLockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenLock *TokenLockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenLock.Contract.TokenLockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenLock *TokenLockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLock.Contract.TokenLockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenLock *TokenLockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenLock.Contract.TokenLockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenLock *TokenLockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenLock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenLock *TokenLockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenLock *TokenLockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenLock.Contract.contract.Transact(opts, method, params...)
}

// LastUnlockTimestamp is a free data retrieval call binding the contract method 0x21226f40.
//
// Solidity: function lastUnlockTimestamp() view returns(uint256)
func (_TokenLock *TokenLockCaller) LastUnlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenLock.contract.Call(opts, &out, "lastUnlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUnlockTimestamp is a free data retrieval call binding the contract method 0x21226f40.
//
// Solidity: function lastUnlockTimestamp() view returns(uint256)
func (_TokenLock *TokenLockSession) LastUnlockTimestamp() (*big.Int, error) {
	return _TokenLock.Contract.LastUnlockTimestamp(&_TokenLock.CallOpts)
}

// LastUnlockTimestamp is a free data retrieval call binding the contract method 0x21226f40.
//
// Solidity: function lastUnlockTimestamp() view returns(uint256)
func (_TokenLock *TokenLockCallerSession) LastUnlockTimestamp() (*big.Int, error) {
	return _TokenLock.Contract.LastUnlockTimestamp(&_TokenLock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenLock *TokenLockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenLock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenLock *TokenLockSession) Owner() (common.Address, error) {
	return _TokenLock.Contract.Owner(&_TokenLock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenLock *TokenLockCallerSession) Owner() (common.Address, error) {
	return _TokenLock.Contract.Owner(&_TokenLock.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_TokenLock *TokenLockCaller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenLock.contract.Call(opts, &out, "startTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_TokenLock *TokenLockSession) StartTimestamp() (*big.Int, error) {
	return _TokenLock.Contract.StartTimestamp(&_TokenLock.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_TokenLock *TokenLockCallerSession) StartTimestamp() (*big.Int, error) {
	return _TokenLock.Contract.StartTimestamp(&_TokenLock.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenLock *TokenLockCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenLock.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenLock *TokenLockSession) Token() (common.Address, error) {
	return _TokenLock.Contract.Token(&_TokenLock.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenLock *TokenLockCallerSession) Token() (common.Address, error) {
	return _TokenLock.Contract.Token(&_TokenLock.CallOpts)
}

// UnlockedToken is a free data retrieval call binding the contract method 0x79a3821c.
//
// Solidity: function unlockedToken() view returns(uint256)
func (_TokenLock *TokenLockCaller) UnlockedToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenLock.contract.Call(opts, &out, "unlockedToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedToken is a free data retrieval call binding the contract method 0x79a3821c.
//
// Solidity: function unlockedToken() view returns(uint256)
func (_TokenLock *TokenLockSession) UnlockedToken() (*big.Int, error) {
	return _TokenLock.Contract.UnlockedToken(&_TokenLock.CallOpts)
}

// UnlockedToken is a free data retrieval call binding the contract method 0x79a3821c.
//
// Solidity: function unlockedToken() view returns(uint256)
func (_TokenLock *TokenLockCallerSession) UnlockedToken() (*big.Int, error) {
	return _TokenLock.Contract.UnlockedToken(&_TokenLock.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_TokenLock *TokenLockTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenLock.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_TokenLock *TokenLockSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenLock.Contract.Transfer(&_TokenLock.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_TokenLock *TokenLockTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenLock.Contract.Transfer(&_TokenLock.TransactOpts, to, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TokenLock *TokenLockTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLock.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TokenLock *TokenLockSession) Unlock() (*types.Transaction, error) {
	return _TokenLock.Contract.Unlock(&_TokenLock.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TokenLock *TokenLockTransactorSession) Unlock() (*types.Transaction, error) {
	return _TokenLock.Contract.Unlock(&_TokenLock.TransactOpts)
}

// TokenLockTransferEventIterator is returned from FilterTransferEvent and is used to iterate over the raw logs and unpacked data for TransferEvent events raised by the TokenLock contract.
type TokenLockTransferEventIterator struct {
	Event *TokenLockTransferEvent // Event containing the contract specifics and raw log

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
func (it *TokenLockTransferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockTransferEvent)
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
		it.Event = new(TokenLockTransferEvent)
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
func (it *TokenLockTransferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockTransferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockTransferEvent represents a TransferEvent event raised by the TokenLock contract.
type TokenLockTransferEvent struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0xb98a26c1d0427d0e3492e749861c7795bed8f6e7599a65143b5903942e611bb0.
//
// Solidity: event TransferEvent(address indexed to, uint256 indexed amount)
func (_TokenLock *TokenLockFilterer) FilterTransferEvent(opts *bind.FilterOpts, to []common.Address, amount []*big.Int) (*TokenLockTransferEventIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _TokenLock.contract.FilterLogs(opts, "TransferEvent", toRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &TokenLockTransferEventIterator{contract: _TokenLock.contract, event: "TransferEvent", logs: logs, sub: sub}, nil
}

// WatchTransferEvent is a free log subscription operation binding the contract event 0xb98a26c1d0427d0e3492e749861c7795bed8f6e7599a65143b5903942e611bb0.
//
// Solidity: event TransferEvent(address indexed to, uint256 indexed amount)
func (_TokenLock *TokenLockFilterer) WatchTransferEvent(opts *bind.WatchOpts, sink chan<- *TokenLockTransferEvent, to []common.Address, amount []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _TokenLock.contract.WatchLogs(opts, "TransferEvent", toRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockTransferEvent)
				if err := _TokenLock.contract.UnpackLog(event, "TransferEvent", log); err != nil {
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

// ParseTransferEvent is a log parse operation binding the contract event 0xb98a26c1d0427d0e3492e749861c7795bed8f6e7599a65143b5903942e611bb0.
//
// Solidity: event TransferEvent(address indexed to, uint256 indexed amount)
func (_TokenLock *TokenLockFilterer) ParseTransferEvent(log types.Log) (*TokenLockTransferEvent, error) {
	event := new(TokenLockTransferEvent)
	if err := _TokenLock.contract.UnpackLog(event, "TransferEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenLockUnlockEventIterator is returned from FilterUnlockEvent and is used to iterate over the raw logs and unpacked data for UnlockEvent events raised by the TokenLock contract.
type TokenLockUnlockEventIterator struct {
	Event *TokenLockUnlockEvent // Event containing the contract specifics and raw log

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
func (it *TokenLockUnlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockUnlockEvent)
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
		it.Event = new(TokenLockUnlockEvent)
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
func (it *TokenLockUnlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockUnlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockUnlockEvent represents a UnlockEvent event raised by the TokenLock contract.
type TokenLockUnlockEvent struct {
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockEvent is a free log retrieval operation binding the contract event 0x085f3df4cd037be4b0a27fe2cb958e1e523fbdb7bf444c6447ab8542e14f6406.
//
// Solidity: event UnlockEvent(uint256 indexed amount, uint256 indexed timestamp)
func (_TokenLock *TokenLockFilterer) FilterUnlockEvent(opts *bind.FilterOpts, amount []*big.Int, timestamp []*big.Int) (*TokenLockUnlockEventIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _TokenLock.contract.FilterLogs(opts, "UnlockEvent", amountRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &TokenLockUnlockEventIterator{contract: _TokenLock.contract, event: "UnlockEvent", logs: logs, sub: sub}, nil
}

// WatchUnlockEvent is a free log subscription operation binding the contract event 0x085f3df4cd037be4b0a27fe2cb958e1e523fbdb7bf444c6447ab8542e14f6406.
//
// Solidity: event UnlockEvent(uint256 indexed amount, uint256 indexed timestamp)
func (_TokenLock *TokenLockFilterer) WatchUnlockEvent(opts *bind.WatchOpts, sink chan<- *TokenLockUnlockEvent, amount []*big.Int, timestamp []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _TokenLock.contract.WatchLogs(opts, "UnlockEvent", amountRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockUnlockEvent)
				if err := _TokenLock.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
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

// ParseUnlockEvent is a log parse operation binding the contract event 0x085f3df4cd037be4b0a27fe2cb958e1e523fbdb7bf444c6447ab8542e14f6406.
//
// Solidity: event UnlockEvent(uint256 indexed amount, uint256 indexed timestamp)
func (_TokenLock *TokenLockFilterer) ParseUnlockEvent(log types.Log) (*TokenLockUnlockEvent, error) {
	event := new(TokenLockUnlockEvent)
	if err := _TokenLock.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
