// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610723806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630a7896021461003b5780638a42ebe91461006b575b600080fd5b61005560048036038101906100509190610271565b610087565b6040516100629190610427565b60405180910390f35b610085600480360381019061008091906102ba565b6100b5565b005b6000818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b60646000836040516100c791906103b2565b90815260200160405180910390205460026100e291906104f0565b6100ec91906104bf565b6101148260008560405161010091906103b2565b9081526020016040518091039020546101bc565b111561017d578060008360405161012b91906103b2565b9081526020016040518091039020819055507fca022b99c592bbe8ea70fa4ded405f1cad24d7b013184a4683787abbf8086c73828242604051610170939291906103c9565b60405180910390a16101b8565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101af90610407565b60405180910390fd5b5050565b6000818310156101d75782826101d2919061054a565b6101e4565b81836101e3919061054a565b5b905092915050565b60006101ff6101fa84610467565b610442565b90508281526020810184848401111561021b5761021a61068d565b5b610226848285610588565b509392505050565b600082601f83011261024357610242610688565b5b81356102538482602086016101ec565b91505092915050565b60008135905061026b816106d6565b92915050565b60006020828403121561028757610286610697565b5b600082013567ffffffffffffffff8111156102a5576102a4610692565b5b6102b18482850161022e565b91505092915050565b600080604083850312156102d1576102d0610697565b5b600083013567ffffffffffffffff8111156102ef576102ee610692565b5b6102fb8582860161022e565b925050602061030c8582860161025c565b9150509250929050565b600061032182610498565b61032b81856104a3565b935061033b818560208601610597565b6103448161069c565b840191505092915050565b600061035a82610498565b61036481856104b4565b9350610374818560208601610597565b80840191505092915050565b600061038d601f836104a3565b9150610398826106ad565b602082019050919050565b6103ac8161057e565b82525050565b60006103be828461034f565b915081905092915050565b600060608201905081810360008301526103e38186610316565b90506103f260208301856103a3565b6103ff60408301846103a3565b949350505050565b6000602082019050818103600083015261042081610380565b9050919050565b600060208201905061043c60008301846103a3565b92915050565b600061044c61045d565b905061045882826105ca565b919050565b6000604051905090565b600067ffffffffffffffff82111561048257610481610659565b5b61048b8261069c565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006104ca8261057e565b91506104d58361057e565b9250826104e5576104e461062a565b5b828204905092915050565b60006104fb8261057e565b91506105068361057e565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561053f5761053e6105fb565b5b828202905092915050565b60006105558261057e565b91506105608361057e565b925082821015610573576105726105fb565b5b828203905092915050565b6000819050919050565b82818337600083830152505050565b60005b838110156105b557808201518184015260208101905061059a565b838111156105c4576000848401525b50505050565b6105d38261069c565b810181811067ffffffffffffffff821117156105f2576105f1610659565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f50726963652073696d696c617220746f20636f6e747261637420707269636500600082015250565b6106df8161057e565b81146106ea57600080fd5b5056fea264697066735822122022c0cbac2778e2d2138d12e2bc4b2233e25a7afc4668bdcd37132095cbf8a0cb64736f6c63430008070033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// Coins is a free data retrieval call binding the contract method 0x0a789602.
//
// Solidity: function coins(string ) view returns(uint256)
func (_Storage *StorageCaller) Coins(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0x0a789602.
//
// Solidity: function coins(string ) view returns(uint256)
func (_Storage *StorageSession) Coins(arg0 string) (*big.Int, error) {
	return _Storage.Contract.Coins(&_Storage.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0x0a789602.
//
// Solidity: function coins(string ) view returns(uint256)
func (_Storage *StorageCallerSession) Coins(arg0 string) (*big.Int, error) {
	return _Storage.Contract.Coins(&_Storage.CallOpts, arg0)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string symbol, uint256 amount) returns()
func (_Storage *StorageTransactor) Set(opts *bind.TransactOpts, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "set", symbol, amount)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string symbol, uint256 amount) returns()
func (_Storage *StorageSession) Set(symbol string, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Set(&_Storage.TransactOpts, symbol, amount)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string symbol, uint256 amount) returns()
func (_Storage *StorageTransactorSession) Set(symbol string, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Set(&_Storage.TransactOpts, symbol, amount)
}

// StorageSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Storage contract.
type StorageSentIterator struct {
	Event *StorageSent // Event containing the contract specifics and raw log

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
func (it *StorageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSent)
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
		it.Event = new(StorageSent)
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
func (it *StorageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSent represents a Sent event raised by the Storage contract.
type StorageSent struct {
	Symbol string
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0xca022b99c592bbe8ea70fa4ded405f1cad24d7b013184a4683787abbf8086c73.
//
// Solidity: event Sent(string symbol, uint256 amount, uint256 time)
func (_Storage *StorageFilterer) FilterSent(opts *bind.FilterOpts) (*StorageSentIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return &StorageSentIterator{contract: _Storage.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0xca022b99c592bbe8ea70fa4ded405f1cad24d7b013184a4683787abbf8086c73.
//
// Solidity: event Sent(string symbol, uint256 amount, uint256 time)
func (_Storage *StorageFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *StorageSent) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSent)
				if err := _Storage.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0xca022b99c592bbe8ea70fa4ded405f1cad24d7b013184a4683787abbf8086c73.
//
// Solidity: event Sent(string symbol, uint256 amount, uint256 time)
func (_Storage *StorageFilterer) ParseSent(log types.Log) (*StorageSent, error) {
	event := new(StorageSent)
	if err := _Storage.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
