// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// LockContractABI is the input ABI used to generate the binding from.
const LockContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newMessage\",\"type\":\"string\"}],\"name\":\"setMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setBooked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setFree\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"message\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialMessage\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// LockContract is an auto generated Go binding around an Ethereum contract.
type LockContract struct {
	LockContractCaller     // Read-only binding to the contract
	LockContractTransactor // Write-only binding to the contract
	LockContractFilterer   // Log filterer for contract events
}

// LockContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockContractSession struct {
	Contract     *LockContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockContractCallerSession struct {
	Contract *LockContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LockContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockContractTransactorSession struct {
	Contract     *LockContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LockContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockContractRaw struct {
	Contract *LockContract // Generic contract binding to access the raw methods on
}

// LockContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockContractCallerRaw struct {
	Contract *LockContractCaller // Generic read-only contract binding to access the raw methods on
}

// LockContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockContractTransactorRaw struct {
	Contract *LockContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockContract creates a new instance of LockContract, bound to a specific deployed contract.
func NewLockContract(address common.Address, backend bind.ContractBackend) (*LockContract, error) {
	contract, err := bindLockContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockContract{LockContractCaller: LockContractCaller{contract: contract}, LockContractTransactor: LockContractTransactor{contract: contract}, LockContractFilterer: LockContractFilterer{contract: contract}}, nil
}

// NewLockContractCaller creates a new read-only instance of LockContract, bound to a specific deployed contract.
func NewLockContractCaller(address common.Address, caller bind.ContractCaller) (*LockContractCaller, error) {
	contract, err := bindLockContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockContractCaller{contract: contract}, nil
}

// NewLockContractTransactor creates a new write-only instance of LockContract, bound to a specific deployed contract.
func NewLockContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LockContractTransactor, error) {
	contract, err := bindLockContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockContractTransactor{contract: contract}, nil
}

// NewLockContractFilterer creates a new log filterer instance of LockContract, bound to a specific deployed contract.
func NewLockContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LockContractFilterer, error) {
	contract, err := bindLockContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockContractFilterer{contract: contract}, nil
}

// bindLockContract binds a generic wrapper to an already deployed contract.
func bindLockContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockContract *LockContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockContract.Contract.LockContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockContract *LockContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockContract.Contract.LockContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockContract *LockContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockContract.Contract.LockContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockContract *LockContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockContract *LockContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockContract *LockContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockContract.Contract.contract.Transact(opts, method, params...)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_LockContract *LockContractCaller) Message(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LockContract.contract.Call(opts, out, "message")
	return *ret0, err
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_LockContract *LockContractSession) Message() (string, error) {
	return _LockContract.Contract.Message(&_LockContract.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_LockContract *LockContractCallerSession) Message() (string, error) {
	return _LockContract.Contract.Message(&_LockContract.CallOpts)
}

// SetBooked is a paid mutator transaction binding the contract method 0xbb0f7e00.
//
// Solidity: function setBooked() returns()
func (_LockContract *LockContractTransactor) SetBooked(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockContract.contract.Transact(opts, "setBooked")
}

// SetBooked is a paid mutator transaction binding the contract method 0xbb0f7e00.
//
// Solidity: function setBooked() returns()
func (_LockContract *LockContractSession) SetBooked() (*types.Transaction, error) {
	return _LockContract.Contract.SetBooked(&_LockContract.TransactOpts)
}

// SetBooked is a paid mutator transaction binding the contract method 0xbb0f7e00.
//
// Solidity: function setBooked() returns()
func (_LockContract *LockContractTransactorSession) SetBooked() (*types.Transaction, error) {
	return _LockContract.Contract.SetBooked(&_LockContract.TransactOpts)
}

// SetFree is a paid mutator transaction binding the contract method 0xde6cf0ee.
//
// Solidity: function setFree() returns()
func (_LockContract *LockContractTransactor) SetFree(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockContract.contract.Transact(opts, "setFree")
}

// SetFree is a paid mutator transaction binding the contract method 0xde6cf0ee.
//
// Solidity: function setFree() returns()
func (_LockContract *LockContractSession) SetFree() (*types.Transaction, error) {
	return _LockContract.Contract.SetFree(&_LockContract.TransactOpts)
}

// SetFree is a paid mutator transaction binding the contract method 0xde6cf0ee.
//
// Solidity: function setFree() returns()
func (_LockContract *LockContractTransactorSession) SetFree() (*types.Transaction, error) {
	return _LockContract.Contract.SetFree(&_LockContract.TransactOpts)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(newMessage string) returns()
func (_LockContract *LockContractTransactor) SetMessage(opts *bind.TransactOpts, newMessage string) (*types.Transaction, error) {
	return _LockContract.contract.Transact(opts, "setMessage", newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(newMessage string) returns()
func (_LockContract *LockContractSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _LockContract.Contract.SetMessage(&_LockContract.TransactOpts, newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(newMessage string) returns()
func (_LockContract *LockContractTransactorSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _LockContract.Contract.SetMessage(&_LockContract.TransactOpts, newMessage)
}
