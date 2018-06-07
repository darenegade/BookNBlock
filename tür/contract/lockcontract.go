// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// LockContractABI is the input ABI used to generate the binding from.
const LockContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bookings\",\"outputs\":[{\"name\":\"offerID\",\"type\":\"uint256\"},{\"name\":\"checkIn\",\"type\":\"uint256\"},{\"name\":\"checkOut\",\"type\":\"uint256\"},{\"name\":\"tenant\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerID\",\"type\":\"uint256\"},{\"name\":\"checkIn\",\"type\":\"uint256\"},{\"name\":\"checkOut\",\"type\":\"uint256\"}],\"name\":\"rentAnOffer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"offerID\",\"type\":\"uint256\"}],\"name\":\"getOffer\",\"outputs\":[{\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"name\":\"objectName\",\"type\":\"string\"},{\"name\":\"objectAddress\",\"type\":\"string\"},{\"name\":\"ownerName\",\"type\":\"string\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"door\",\"type\":\"address\"},{\"name\":\"validFrom\",\"type\":\"uint256\"},{\"name\":\"validUntil\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"bookingID\",\"type\":\"uint256\"},{\"name\":\"tenant\",\"type\":\"address\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"isAllowedAt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"bookingID\",\"type\":\"uint256\"}],\"name\":\"getBooking\",\"outputs\":[{\"name\":\"offerID\",\"type\":\"uint256\"},{\"name\":\"checkIn\",\"type\":\"uint256\"},{\"name\":\"checkOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOffersLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerID\",\"type\":\"uint256\"}],\"name\":\"deleteOffer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"offerID\",\"type\":\"uint256\"}],\"name\":\"getBookingIDsForOffer\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offers\",\"outputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"name\":\"objectName\",\"type\":\"string\"},{\"name\":\"objectAddress\",\"type\":\"string\"},{\"name\":\"ownerName\",\"type\":\"string\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"validFrom\",\"type\":\"uint256\"},{\"name\":\"validUntil\",\"type\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"door\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerID\",\"type\":\"uint256\"},{\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"name\":\"objectName\",\"type\":\"string\"},{\"name\":\"objectAddress\",\"type\":\"string\"},{\"name\":\"ownerName\",\"type\":\"string\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"door\",\"type\":\"address\"},{\"name\":\"validFrom\",\"type\":\"uint256\"},{\"name\":\"validUntil\",\"type\":\"uint256\"}],\"name\":\"updateOffer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offerIDs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"getFreeOfferIDs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOfferIDs\",\"outputs\":[{\"name\":\"offerIDs\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"name\":\"objectName\",\"type\":\"string\"},{\"name\":\"objectAddress\",\"type\":\"string\"},{\"name\":\"ownerName\",\"type\":\"string\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"door\",\"type\":\"address\"},{\"name\":\"validFrom\",\"type\":\"uint256\"},{\"name\":\"validUntil\",\"type\":\"uint256\"}],\"name\":\"insertOffer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"offerID\",\"type\":\"uint256\"}],\"name\":\"OfferSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"offerID\",\"type\":\"uint256\"}],\"name\":\"OfferDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"offerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"bookingID\",\"type\":\"uint256\"}],\"name\":\"BookingAccepted\",\"type\":\"event\"}]"

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

// Bookings is a free data retrieval call binding the contract method 0x1dab301e.
//
// Solidity: function bookings( uint256) constant returns(offerID uint256, checkIn uint256, checkOut uint256, tenant address)
func (_LockContract *LockContractCaller) Bookings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OfferID  *big.Int
	CheckIn  *big.Int
	CheckOut *big.Int
	Tenant   common.Address
}, error) {
	ret := new(struct {
		OfferID  *big.Int
		CheckIn  *big.Int
		CheckOut *big.Int
		Tenant   common.Address
	})
	out := ret
	err := _LockContract.contract.Call(opts, out, "bookings", arg0)
	return *ret, err
}

// Bookings is a free data retrieval call binding the contract method 0x1dab301e.
//
// Solidity: function bookings( uint256) constant returns(offerID uint256, checkIn uint256, checkOut uint256, tenant address)
func (_LockContract *LockContractSession) Bookings(arg0 *big.Int) (struct {
	OfferID  *big.Int
	CheckIn  *big.Int
	CheckOut *big.Int
	Tenant   common.Address
}, error) {
	return _LockContract.Contract.Bookings(&_LockContract.CallOpts, arg0)
}

// Bookings is a free data retrieval call binding the contract method 0x1dab301e.
//
// Solidity: function bookings( uint256) constant returns(offerID uint256, checkIn uint256, checkOut uint256, tenant address)
func (_LockContract *LockContractCallerSession) Bookings(arg0 *big.Int) (struct {
	OfferID  *big.Int
	CheckIn  *big.Int
	CheckOut *big.Int
	Tenant   common.Address
}, error) {
	return _LockContract.Contract.Bookings(&_LockContract.CallOpts, arg0)
}

// GetBooking is a free data retrieval call binding the contract method 0x6a5c841a.
//
// Solidity: function getBooking(bookingID uint256) constant returns(offerID uint256, checkIn uint256, checkOut uint256)
func (_LockContract *LockContractCaller) GetBooking(opts *bind.CallOpts, bookingID *big.Int) (struct {
	OfferID  *big.Int
	CheckIn  *big.Int
	CheckOut *big.Int
}, error) {
	ret := new(struct {
		OfferID  *big.Int
		CheckIn  *big.Int
		CheckOut *big.Int
	})
	out := ret
	err := _LockContract.contract.Call(opts, out, "getBooking", bookingID)
	return *ret, err
}

// GetBooking is a free data retrieval call binding the contract method 0x6a5c841a.
//
// Solidity: function getBooking(bookingID uint256) constant returns(offerID uint256, checkIn uint256, checkOut uint256)
func (_LockContract *LockContractSession) GetBooking(bookingID *big.Int) (struct {
	OfferID  *big.Int
	CheckIn  *big.Int
	CheckOut *big.Int
}, error) {
	return _LockContract.Contract.GetBooking(&_LockContract.CallOpts, bookingID)
}

// GetBooking is a free data retrieval call binding the contract method 0x6a5c841a.
//
// Solidity: function getBooking(bookingID uint256) constant returns(offerID uint256, checkIn uint256, checkOut uint256)
func (_LockContract *LockContractCallerSession) GetBooking(bookingID *big.Int) (struct {
	OfferID  *big.Int
	CheckIn  *big.Int
	CheckOut *big.Int
}, error) {
	return _LockContract.Contract.GetBooking(&_LockContract.CallOpts, bookingID)
}

// GetBookingIDsForOffer is a free data retrieval call binding the contract method 0x7eef7605.
//
// Solidity: function getBookingIDsForOffer(offerID uint256) constant returns(uint256[])
func (_LockContract *LockContractCaller) GetBookingIDsForOffer(opts *bind.CallOpts, offerID *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _LockContract.contract.Call(opts, out, "getBookingIDsForOffer", offerID)
	return *ret0, err
}

// GetBookingIDsForOffer is a free data retrieval call binding the contract method 0x7eef7605.
//
// Solidity: function getBookingIDsForOffer(offerID uint256) constant returns(uint256[])
func (_LockContract *LockContractSession) GetBookingIDsForOffer(offerID *big.Int) ([]*big.Int, error) {
	return _LockContract.Contract.GetBookingIDsForOffer(&_LockContract.CallOpts, offerID)
}

// GetBookingIDsForOffer is a free data retrieval call binding the contract method 0x7eef7605.
//
// Solidity: function getBookingIDsForOffer(offerID uint256) constant returns(uint256[])
func (_LockContract *LockContractCallerSession) GetBookingIDsForOffer(offerID *big.Int) ([]*big.Int, error) {
	return _LockContract.Contract.GetBookingIDsForOffer(&_LockContract.CallOpts, offerID)
}

// GetFreeOfferIDs is a free data retrieval call binding the contract method 0xd225a22a.
//
// Solidity: function getFreeOfferIDs(from uint256, to uint256) constant returns(uint256[])
func (_LockContract *LockContractCaller) GetFreeOfferIDs(opts *bind.CallOpts, from *big.Int, to *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _LockContract.contract.Call(opts, out, "getFreeOfferIDs", from, to)
	return *ret0, err
}

// GetFreeOfferIDs is a free data retrieval call binding the contract method 0xd225a22a.
//
// Solidity: function getFreeOfferIDs(from uint256, to uint256) constant returns(uint256[])
func (_LockContract *LockContractSession) GetFreeOfferIDs(from *big.Int, to *big.Int) ([]*big.Int, error) {
	return _LockContract.Contract.GetFreeOfferIDs(&_LockContract.CallOpts, from, to)
}

// GetFreeOfferIDs is a free data retrieval call binding the contract method 0xd225a22a.
//
// Solidity: function getFreeOfferIDs(from uint256, to uint256) constant returns(uint256[])
func (_LockContract *LockContractCallerSession) GetFreeOfferIDs(from *big.Int, to *big.Int) ([]*big.Int, error) {
	return _LockContract.Contract.GetFreeOfferIDs(&_LockContract.CallOpts, from, to)
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(offerID uint256) constant returns(priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, door address, validFrom uint256, validUntil uint256)
func (_LockContract *LockContractCaller) GetOffer(opts *bind.CallOpts, offerID *big.Int) (struct {
	PriceInWei    *big.Int
	ObjectName    string
	ObjectAddress string
	OwnerName     string
	Description   string
	Door          common.Address
	ValidFrom     *big.Int
	ValidUntil    *big.Int
}, error) {
	ret := new(struct {
		PriceInWei    *big.Int
		ObjectName    string
		ObjectAddress string
		OwnerName     string
		Description   string
		Door          common.Address
		ValidFrom     *big.Int
		ValidUntil    *big.Int
	})
	out := ret
	err := _LockContract.contract.Call(opts, out, "getOffer", offerID)
	return *ret, err
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(offerID uint256) constant returns(priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, door address, validFrom uint256, validUntil uint256)
func (_LockContract *LockContractSession) GetOffer(offerID *big.Int) (struct {
	PriceInWei    *big.Int
	ObjectName    string
	ObjectAddress string
	OwnerName     string
	Description   string
	Door          common.Address
	ValidFrom     *big.Int
	ValidUntil    *big.Int
}, error) {
	return _LockContract.Contract.GetOffer(&_LockContract.CallOpts, offerID)
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(offerID uint256) constant returns(priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, door address, validFrom uint256, validUntil uint256)
func (_LockContract *LockContractCallerSession) GetOffer(offerID *big.Int) (struct {
	PriceInWei    *big.Int
	ObjectName    string
	ObjectAddress string
	OwnerName     string
	Description   string
	Door          common.Address
	ValidFrom     *big.Int
	ValidUntil    *big.Int
}, error) {
	return _LockContract.Contract.GetOffer(&_LockContract.CallOpts, offerID)
}

// GetOfferIDs is a free data retrieval call binding the contract method 0xf6a80192.
//
// Solidity: function getOfferIDs() constant returns(offerIDs uint256[])
func (_LockContract *LockContractCaller) GetOfferIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _LockContract.contract.Call(opts, out, "getOfferIDs")
	return *ret0, err
}

// GetOfferIDs is a free data retrieval call binding the contract method 0xf6a80192.
//
// Solidity: function getOfferIDs() constant returns(offerIDs uint256[])
func (_LockContract *LockContractSession) GetOfferIDs() ([]*big.Int, error) {
	return _LockContract.Contract.GetOfferIDs(&_LockContract.CallOpts)
}

// GetOfferIDs is a free data retrieval call binding the contract method 0xf6a80192.
//
// Solidity: function getOfferIDs() constant returns(offerIDs uint256[])
func (_LockContract *LockContractCallerSession) GetOfferIDs() ([]*big.Int, error) {
	return _LockContract.Contract.GetOfferIDs(&_LockContract.CallOpts)
}

// GetOffersLength is a free data retrieval call binding the contract method 0x73c2dc12.
//
// Solidity: function getOffersLength() constant returns(uint256)
func (_LockContract *LockContractCaller) GetOffersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockContract.contract.Call(opts, out, "getOffersLength")
	return *ret0, err
}

// GetOffersLength is a free data retrieval call binding the contract method 0x73c2dc12.
//
// Solidity: function getOffersLength() constant returns(uint256)
func (_LockContract *LockContractSession) GetOffersLength() (*big.Int, error) {
	return _LockContract.Contract.GetOffersLength(&_LockContract.CallOpts)
}

// GetOffersLength is a free data retrieval call binding the contract method 0x73c2dc12.
//
// Solidity: function getOffersLength() constant returns(uint256)
func (_LockContract *LockContractCallerSession) GetOffersLength() (*big.Int, error) {
	return _LockContract.Contract.GetOffersLength(&_LockContract.CallOpts)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x614eea3f.
//
// Solidity: function isAllowedAt(bookingID uint256, tenant address, time uint256) constant returns(bool)
func (_LockContract *LockContractCaller) IsAllowedAt(opts *bind.CallOpts, bookingID *big.Int, tenant common.Address, time *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LockContract.contract.Call(opts, out, "isAllowedAt", bookingID, tenant, time)
	return *ret0, err
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x614eea3f.
//
// Solidity: function isAllowedAt(bookingID uint256, tenant address, time uint256) constant returns(bool)
func (_LockContract *LockContractSession) IsAllowedAt(bookingID *big.Int, tenant common.Address, time *big.Int) (bool, error) {
	return _LockContract.Contract.IsAllowedAt(&_LockContract.CallOpts, bookingID, tenant, time)
}

// IsAllowedAt is a free data retrieval call binding the contract method 0x614eea3f.
//
// Solidity: function isAllowedAt(bookingID uint256, tenant address, time uint256) constant returns(bool)
func (_LockContract *LockContractCallerSession) IsAllowedAt(bookingID *big.Int, tenant common.Address, time *big.Int) (bool, error) {
	return _LockContract.Contract.IsAllowedAt(&_LockContract.CallOpts, bookingID, tenant, time)
}

// NextID is a free data retrieval call binding the contract method 0x1e96917d.
//
// Solidity: function nextID() constant returns(uint256)
func (_LockContract *LockContractCaller) NextID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockContract.contract.Call(opts, out, "nextID")
	return *ret0, err
}

// NextID is a free data retrieval call binding the contract method 0x1e96917d.
//
// Solidity: function nextID() constant returns(uint256)
func (_LockContract *LockContractSession) NextID() (*big.Int, error) {
	return _LockContract.Contract.NextID(&_LockContract.CallOpts)
}

// NextID is a free data retrieval call binding the contract method 0x1e96917d.
//
// Solidity: function nextID() constant returns(uint256)
func (_LockContract *LockContractCallerSession) NextID() (*big.Int, error) {
	return _LockContract.Contract.NextID(&_LockContract.CallOpts)
}

// OfferIDs is a free data retrieval call binding the contract method 0xcf38cf76.
//
// Solidity: function offerIDs( uint256) constant returns(uint256)
func (_LockContract *LockContractCaller) OfferIDs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockContract.contract.Call(opts, out, "offerIDs", arg0)
	return *ret0, err
}

// OfferIDs is a free data retrieval call binding the contract method 0xcf38cf76.
//
// Solidity: function offerIDs( uint256) constant returns(uint256)
func (_LockContract *LockContractSession) OfferIDs(arg0 *big.Int) (*big.Int, error) {
	return _LockContract.Contract.OfferIDs(&_LockContract.CallOpts, arg0)
}

// OfferIDs is a free data retrieval call binding the contract method 0xcf38cf76.
//
// Solidity: function offerIDs( uint256) constant returns(uint256)
func (_LockContract *LockContractCallerSession) OfferIDs(arg0 *big.Int) (*big.Int, error) {
	return _LockContract.Contract.OfferIDs(&_LockContract.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers( uint256) constant returns(index uint256, priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, validFrom uint256, validUntil uint256, owner address, door address)
func (_LockContract *LockContractCaller) Offers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index         *big.Int
	PriceInWei    *big.Int
	ObjectName    string
	ObjectAddress string
	OwnerName     string
	Description   string
	ValidFrom     *big.Int
	ValidUntil    *big.Int
	Owner         common.Address
	Door          common.Address
}, error) {
	ret := new(struct {
		Index         *big.Int
		PriceInWei    *big.Int
		ObjectName    string
		ObjectAddress string
		OwnerName     string
		Description   string
		ValidFrom     *big.Int
		ValidUntil    *big.Int
		Owner         common.Address
		Door          common.Address
	})
	out := ret
	err := _LockContract.contract.Call(opts, out, "offers", arg0)
	return *ret, err
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers( uint256) constant returns(index uint256, priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, validFrom uint256, validUntil uint256, owner address, door address)
func (_LockContract *LockContractSession) Offers(arg0 *big.Int) (struct {
	Index         *big.Int
	PriceInWei    *big.Int
	ObjectName    string
	ObjectAddress string
	OwnerName     string
	Description   string
	ValidFrom     *big.Int
	ValidUntil    *big.Int
	Owner         common.Address
	Door          common.Address
}, error) {
	return _LockContract.Contract.Offers(&_LockContract.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers( uint256) constant returns(index uint256, priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, validFrom uint256, validUntil uint256, owner address, door address)
func (_LockContract *LockContractCallerSession) Offers(arg0 *big.Int) (struct {
	Index         *big.Int
	PriceInWei    *big.Int
	ObjectName    string
	ObjectAddress string
	OwnerName     string
	Description   string
	ValidFrom     *big.Int
	ValidUntil    *big.Int
	Owner         common.Address
	Door          common.Address
}, error) {
	return _LockContract.Contract.Offers(&_LockContract.CallOpts, arg0)
}

// DeleteOffer is a paid mutator transaction binding the contract method 0x74268ff2.
//
// Solidity: function deleteOffer(offerID uint256) returns()
func (_LockContract *LockContractTransactor) DeleteOffer(opts *bind.TransactOpts, offerID *big.Int) (*types.Transaction, error) {
	return _LockContract.contract.Transact(opts, "deleteOffer", offerID)
}

// DeleteOffer is a paid mutator transaction binding the contract method 0x74268ff2.
//
// Solidity: function deleteOffer(offerID uint256) returns()
func (_LockContract *LockContractSession) DeleteOffer(offerID *big.Int) (*types.Transaction, error) {
	return _LockContract.Contract.DeleteOffer(&_LockContract.TransactOpts, offerID)
}

// DeleteOffer is a paid mutator transaction binding the contract method 0x74268ff2.
//
// Solidity: function deleteOffer(offerID uint256) returns()
func (_LockContract *LockContractTransactorSession) DeleteOffer(offerID *big.Int) (*types.Transaction, error) {
	return _LockContract.Contract.DeleteOffer(&_LockContract.TransactOpts, offerID)
}

// InsertOffer is a paid mutator transaction binding the contract method 0xf799f97b.
//
// Solidity: function insertOffer(priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, door address, validFrom uint256, validUntil uint256) returns()
func (_LockContract *LockContractTransactor) InsertOffer(opts *bind.TransactOpts, priceInWei *big.Int, objectName string, objectAddress string, ownerName string, description string, door common.Address, validFrom *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _LockContract.contract.Transact(opts, "insertOffer", priceInWei, objectName, objectAddress, ownerName, description, door, validFrom, validUntil)
}

// InsertOffer is a paid mutator transaction binding the contract method 0xf799f97b.
//
// Solidity: function insertOffer(priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, door address, validFrom uint256, validUntil uint256) returns()
func (_LockContract *LockContractSession) InsertOffer(priceInWei *big.Int, objectName string, objectAddress string, ownerName string, description string, door common.Address, validFrom *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _LockContract.Contract.InsertOffer(&_LockContract.TransactOpts, priceInWei, objectName, objectAddress, ownerName, description, door, validFrom, validUntil)
}

// InsertOffer is a paid mutator transaction binding the contract method 0xf799f97b.
//
// Solidity: function insertOffer(priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, door address, validFrom uint256, validUntil uint256) returns()
func (_LockContract *LockContractTransactorSession) InsertOffer(priceInWei *big.Int, objectName string, objectAddress string, ownerName string, description string, door common.Address, validFrom *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _LockContract.Contract.InsertOffer(&_LockContract.TransactOpts, priceInWei, objectName, objectAddress, ownerName, description, door, validFrom, validUntil)
}

// RentAnOffer is a paid mutator transaction binding the contract method 0x2ee9a393.
//
// Solidity: function rentAnOffer(offerID uint256, checkIn uint256, checkOut uint256) returns()
func (_LockContract *LockContractTransactor) RentAnOffer(opts *bind.TransactOpts, offerID *big.Int, checkIn *big.Int, checkOut *big.Int) (*types.Transaction, error) {
	return _LockContract.contract.Transact(opts, "rentAnOffer", offerID, checkIn, checkOut)
}

// RentAnOffer is a paid mutator transaction binding the contract method 0x2ee9a393.
//
// Solidity: function rentAnOffer(offerID uint256, checkIn uint256, checkOut uint256) returns()
func (_LockContract *LockContractSession) RentAnOffer(offerID *big.Int, checkIn *big.Int, checkOut *big.Int) (*types.Transaction, error) {
	return _LockContract.Contract.RentAnOffer(&_LockContract.TransactOpts, offerID, checkIn, checkOut)
}

// RentAnOffer is a paid mutator transaction binding the contract method 0x2ee9a393.
//
// Solidity: function rentAnOffer(offerID uint256, checkIn uint256, checkOut uint256) returns()
func (_LockContract *LockContractTransactorSession) RentAnOffer(offerID *big.Int, checkIn *big.Int, checkOut *big.Int) (*types.Transaction, error) {
	return _LockContract.Contract.RentAnOffer(&_LockContract.TransactOpts, offerID, checkIn, checkOut)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0xbee817a5.
//
// Solidity: function updateOffer(offerID uint256, priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, door address, validFrom uint256, validUntil uint256) returns()
func (_LockContract *LockContractTransactor) UpdateOffer(opts *bind.TransactOpts, offerID *big.Int, priceInWei *big.Int, objectName string, objectAddress string, ownerName string, description string, door common.Address, validFrom *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _LockContract.contract.Transact(opts, "updateOffer", offerID, priceInWei, objectName, objectAddress, ownerName, description, door, validFrom, validUntil)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0xbee817a5.
//
// Solidity: function updateOffer(offerID uint256, priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, door address, validFrom uint256, validUntil uint256) returns()
func (_LockContract *LockContractSession) UpdateOffer(offerID *big.Int, priceInWei *big.Int, objectName string, objectAddress string, ownerName string, description string, door common.Address, validFrom *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _LockContract.Contract.UpdateOffer(&_LockContract.TransactOpts, offerID, priceInWei, objectName, objectAddress, ownerName, description, door, validFrom, validUntil)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0xbee817a5.
//
// Solidity: function updateOffer(offerID uint256, priceInWei uint256, objectName string, objectAddress string, ownerName string, description string, door address, validFrom uint256, validUntil uint256) returns()
func (_LockContract *LockContractTransactorSession) UpdateOffer(offerID *big.Int, priceInWei *big.Int, objectName string, objectAddress string, ownerName string, description string, door common.Address, validFrom *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _LockContract.Contract.UpdateOffer(&_LockContract.TransactOpts, offerID, priceInWei, objectName, objectAddress, ownerName, description, door, validFrom, validUntil)
}

// LockContractBookingAcceptedIterator is returned from FilterBookingAccepted and is used to iterate over the raw logs and unpacked data for BookingAccepted events raised by the LockContract contract.
type LockContractBookingAcceptedIterator struct {
	Event *LockContractBookingAccepted // Event containing the contract specifics and raw log

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
func (it *LockContractBookingAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockContractBookingAccepted)
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
		it.Event = new(LockContractBookingAccepted)
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
func (it *LockContractBookingAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockContractBookingAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockContractBookingAccepted represents a BookingAccepted event raised by the LockContract contract.
type LockContractBookingAccepted struct {
	OfferID   *big.Int
	BookingID *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBookingAccepted is a free log retrieval operation binding the contract event 0xebfdd331456c5e03b6aba4ba4168116ac35f3f243175359d6671790f596ced02.
//
// Solidity: e BookingAccepted(offerID uint256, bookingID uint256)
func (_LockContract *LockContractFilterer) FilterBookingAccepted(opts *bind.FilterOpts) (*LockContractBookingAcceptedIterator, error) {

	logs, sub, err := _LockContract.contract.FilterLogs(opts, "BookingAccepted")
	if err != nil {
		return nil, err
	}
	return &LockContractBookingAcceptedIterator{contract: _LockContract.contract, event: "BookingAccepted", logs: logs, sub: sub}, nil
}

// WatchBookingAccepted is a free log subscription operation binding the contract event 0xebfdd331456c5e03b6aba4ba4168116ac35f3f243175359d6671790f596ced02.
//
// Solidity: e BookingAccepted(offerID uint256, bookingID uint256)
func (_LockContract *LockContractFilterer) WatchBookingAccepted(opts *bind.WatchOpts, sink chan<- *LockContractBookingAccepted) (event.Subscription, error) {

	logs, sub, err := _LockContract.contract.WatchLogs(opts, "BookingAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockContractBookingAccepted)
				if err := _LockContract.contract.UnpackLog(event, "BookingAccepted", log); err != nil {
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

// LockContractOfferDeletedIterator is returned from FilterOfferDeleted and is used to iterate over the raw logs and unpacked data for OfferDeleted events raised by the LockContract contract.
type LockContractOfferDeletedIterator struct {
	Event *LockContractOfferDeleted // Event containing the contract specifics and raw log

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
func (it *LockContractOfferDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockContractOfferDeleted)
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
		it.Event = new(LockContractOfferDeleted)
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
func (it *LockContractOfferDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockContractOfferDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockContractOfferDeleted represents a OfferDeleted event raised by the LockContract contract.
type LockContractOfferDeleted struct {
	OfferID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferDeleted is a free log retrieval operation binding the contract event 0x88686b85d6f2c3ab9a04e4f15a22fcfa025ffd97226dcf0a67cdf682def55676.
//
// Solidity: e OfferDeleted(offerID uint256)
func (_LockContract *LockContractFilterer) FilterOfferDeleted(opts *bind.FilterOpts) (*LockContractOfferDeletedIterator, error) {

	logs, sub, err := _LockContract.contract.FilterLogs(opts, "OfferDeleted")
	if err != nil {
		return nil, err
	}
	return &LockContractOfferDeletedIterator{contract: _LockContract.contract, event: "OfferDeleted", logs: logs, sub: sub}, nil
}

// WatchOfferDeleted is a free log subscription operation binding the contract event 0x88686b85d6f2c3ab9a04e4f15a22fcfa025ffd97226dcf0a67cdf682def55676.
//
// Solidity: e OfferDeleted(offerID uint256)
func (_LockContract *LockContractFilterer) WatchOfferDeleted(opts *bind.WatchOpts, sink chan<- *LockContractOfferDeleted) (event.Subscription, error) {

	logs, sub, err := _LockContract.contract.WatchLogs(opts, "OfferDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockContractOfferDeleted)
				if err := _LockContract.contract.UnpackLog(event, "OfferDeleted", log); err != nil {
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

// LockContractOfferSavedIterator is returned from FilterOfferSaved and is used to iterate over the raw logs and unpacked data for OfferSaved events raised by the LockContract contract.
type LockContractOfferSavedIterator struct {
	Event *LockContractOfferSaved // Event containing the contract specifics and raw log

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
func (it *LockContractOfferSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockContractOfferSaved)
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
		it.Event = new(LockContractOfferSaved)
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
func (it *LockContractOfferSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockContractOfferSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockContractOfferSaved represents a OfferSaved event raised by the LockContract contract.
type LockContractOfferSaved struct {
	OfferID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferSaved is a free log retrieval operation binding the contract event 0x7d8710ba1a688609c154ba81335d0d0ccdffc7aa16d629c15673ec1d8e352dfb.
//
// Solidity: e OfferSaved(offerID uint256)
func (_LockContract *LockContractFilterer) FilterOfferSaved(opts *bind.FilterOpts) (*LockContractOfferSavedIterator, error) {

	logs, sub, err := _LockContract.contract.FilterLogs(opts, "OfferSaved")
	if err != nil {
		return nil, err
	}
	return &LockContractOfferSavedIterator{contract: _LockContract.contract, event: "OfferSaved", logs: logs, sub: sub}, nil
}

// WatchOfferSaved is a free log subscription operation binding the contract event 0x7d8710ba1a688609c154ba81335d0d0ccdffc7aa16d629c15673ec1d8e352dfb.
//
// Solidity: e OfferSaved(offerID uint256)
func (_LockContract *LockContractFilterer) WatchOfferSaved(opts *bind.WatchOpts, sink chan<- *LockContractOfferSaved) (event.Subscription, error) {

	logs, sub, err := _LockContract.contract.WatchLogs(opts, "OfferSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockContractOfferSaved)
				if err := _LockContract.contract.UnpackLog(event, "OfferSaved", log); err != nil {
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
