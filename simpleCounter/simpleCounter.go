// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpleCounter

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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CountChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b505f5f819055503373ffffffffffffffffffffffffffffffffffffffff167f31295c5079188370bc96e74c957d1bebe436df233302a41c40843701790699525f5460405161005d91906100dc565b60405180910390a2610108565b5f819050919050565b61007c8161006a565b82525050565b5f82825260208201905092915050565b7f496e697469616c697a65000000000000000000000000000000000000000000005f82015250565b5f6100c6600a83610082565b91506100d182610092565b602082019050919050565b5f6040820190506100ef5f830184610073565b8181036020830152610100816100ba565b905092915050565b610696806101155f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80632baeceb71461006457806387db03b71461006e578063a87d942c1461008a578063d09de08a146100a8578063d826f88f146100b2578063e5c19b2d146100bc575b5f5ffd5b61006c6100d8565b005b61008860048036038101906100839190610301565b610141565b005b6100926101aa565b60405161009f919061033b565b60405180910390f35b6100b06101b2565b005b6100ba61021b565b005b6100d660048036038101906100d19190610301565b610272565b005b60015f5f8282546100e99190610381565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f31295c5079188370bc96e74c957d1bebe436df233302a41c40843701790699525f54604051610137919061041b565b60405180910390a2565b805f5f8282546101519190610447565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f31295c5079188370bc96e74c957d1bebe436df233302a41c40843701790699525f5460405161019f91906104d2565b60405180910390a250565b5f5f54905090565b60015f5f8282546101c39190610447565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f31295c5079188370bc96e74c957d1bebe436df233302a41c40843701790699525f546040516102119190610548565b60405180910390a2565b5f5f819055503373ffffffffffffffffffffffffffffffffffffffff167f31295c5079188370bc96e74c957d1bebe436df233302a41c40843701790699525f5460405161026891906105be565b60405180910390a2565b805f819055503373ffffffffffffffffffffffffffffffffffffffff167f31295c5079188370bc96e74c957d1bebe436df233302a41c40843701790699525f546040516102bf9190610634565b60405180910390a250565b5f5ffd5b5f819050919050565b6102e0816102ce565b81146102ea575f5ffd5b50565b5f813590506102fb816102d7565b92915050565b5f60208284031215610316576103156102ca565b5b5f610323848285016102ed565b91505092915050565b610335816102ce565b82525050565b5f60208201905061034e5f83018461032c565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61038b826102ce565b9150610396836102ce565b925082820390508181125f8412168282135f8512151617156103bb576103ba610354565b5b92915050565b5f82825260208201905092915050565b7f44656372656d656e7400000000000000000000000000000000000000000000005f82015250565b5f6104056009836103c1565b9150610410826103d1565b602082019050919050565b5f60408201905061042e5f83018461032c565b818103602083015261043f816103f9565b905092915050565b5f610451826102ce565b915061045c836102ce565b92508282019050828112155f8312168382125f84121516171561048257610481610354565b5b92915050565b7f41646400000000000000000000000000000000000000000000000000000000005f82015250565b5f6104bc6003836103c1565b91506104c782610488565b602082019050919050565b5f6040820190506104e55f83018461032c565b81810360208301526104f6816104b0565b905092915050565b7f496e6372656d656e7400000000000000000000000000000000000000000000005f82015250565b5f6105326009836103c1565b915061053d826104fe565b602082019050919050565b5f60408201905061055b5f83018461032c565b818103602083015261056c81610526565b905092915050565b7f52657365740000000000000000000000000000000000000000000000000000005f82015250565b5f6105a86005836103c1565b91506105b382610574565b602082019050919050565b5f6040820190506105d15f83018461032c565b81810360208301526105e28161059c565b905092915050565b7f53657400000000000000000000000000000000000000000000000000000000005f82015250565b5f61061e6003836103c1565b9150610629826105ea565b602082019050919050565b5f6040820190506106475f83018461032c565b818103602083015261065881610612565b90509291505056fea26469706673582212202e286b3c797fce2000cbf368838c020ba3bae1f541241c336c1a23e69a4c06ab64736f6c634300081e0033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_Counter *CounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_Counter *CounterSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_Counter *CounterCallerSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x87db03b7.
//
// Solidity: function add(int256 _value) returns()
func (_Counter *CounterTransactor) Add(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "add", _value)
}

// Add is a paid mutator transaction binding the contract method 0x87db03b7.
//
// Solidity: function add(int256 _value) returns()
func (_Counter *CounterSession) Add(_value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Add(&_Counter.TransactOpts, _value)
}

// Add is a paid mutator transaction binding the contract method 0x87db03b7.
//
// Solidity: function add(int256 _value) returns()
func (_Counter *CounterTransactorSession) Add(_value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Add(&_Counter.TransactOpts, _value)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Counter *CounterTransactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Counter *CounterSession) Decrement() (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Counter *CounterTransactorSession) Decrement() (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactorSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterTransactorSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0xe5c19b2d.
//
// Solidity: function set(int256 _value) returns()
func (_Counter *CounterTransactor) Set(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "set", _value)
}

// Set is a paid mutator transaction binding the contract method 0xe5c19b2d.
//
// Solidity: function set(int256 _value) returns()
func (_Counter *CounterSession) Set(_value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Set(&_Counter.TransactOpts, _value)
}

// Set is a paid mutator transaction binding the contract method 0xe5c19b2d.
//
// Solidity: function set(int256 _value) returns()
func (_Counter *CounterTransactorSession) Set(_value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Set(&_Counter.TransactOpts, _value)
}

// CounterCountChangedIterator is returned from FilterCountChanged and is used to iterate over the raw logs and unpacked data for CountChanged events raised by the Counter contract.
type CounterCountChangedIterator struct {
	Event *CounterCountChanged // Event containing the contract specifics and raw log

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
func (it *CounterCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCountChanged)
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
		it.Event = new(CounterCountChanged)
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
func (it *CounterCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCountChanged represents a CountChanged event raised by the Counter contract.
type CounterCountChanged struct {
	NewValue  *big.Int
	ChangedBy common.Address
	Operation string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCountChanged is a free log retrieval operation binding the contract event 0x31295c5079188370bc96e74c957d1bebe436df233302a41c4084370179069952.
//
// Solidity: event CountChanged(int256 newValue, address indexed changedBy, string operation)
func (_Counter *CounterFilterer) FilterCountChanged(opts *bind.FilterOpts, changedBy []common.Address) (*CounterCountChangedIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _Counter.contract.FilterLogs(opts, "CountChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return &CounterCountChangedIterator{contract: _Counter.contract, event: "CountChanged", logs: logs, sub: sub}, nil
}

// WatchCountChanged is a free log subscription operation binding the contract event 0x31295c5079188370bc96e74c957d1bebe436df233302a41c4084370179069952.
//
// Solidity: event CountChanged(int256 newValue, address indexed changedBy, string operation)
func (_Counter *CounterFilterer) WatchCountChanged(opts *bind.WatchOpts, sink chan<- *CounterCountChanged, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _Counter.contract.WatchLogs(opts, "CountChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCountChanged)
				if err := _Counter.contract.UnpackLog(event, "CountChanged", log); err != nil {
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

// ParseCountChanged is a log parse operation binding the contract event 0x31295c5079188370bc96e74c957d1bebe436df233302a41c4084370179069952.
//
// Solidity: event CountChanged(int256 newValue, address indexed changedBy, string operation)
func (_Counter *CounterFilterer) ParseCountChanged(log types.Log) (*CounterCountChanged, error) {
	event := new(CounterCountChanged)
	if err := _Counter.contract.UnpackLog(event, "CountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
