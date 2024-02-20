// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lara_factory

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

// LaraFactoryMetaData contains all meta data concerning the LaraFactory contract.
var LaraFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sttaraToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dposContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_apyOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasuryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newCommission\",\"type\":\"uint256\"}],\"name\":\"CommissionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"laraAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"activator\",\"type\":\"address\"}],\"name\":\"LaraActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"laraAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"LaraCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"laraAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deactivator\",\"type\":\"address\"}],\"name\":\"LaraDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"activateLara\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeLaraInstanceCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apyOracle\",\"outputs\":[{\"internalType\":\"contractIApyOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createLara\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"deactivateLara\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dposContract\",\"outputs\":[{\"internalType\":\"contractDposInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"laraActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"laraAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"laraInstanceCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"laraInstances\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorStakeCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_commission\",\"type\":\"uint256\"}],\"name\":\"setCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"name\":\"setEpochDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxValidatorStakeCapacity\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorStakeCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minStakeAmount\",\"type\":\"uint256\"}],\"name\":\"setMinStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddress\",\"type\":\"address\"}],\"name\":\"setTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stTaraToken\",\"outputs\":[{\"internalType\":\"contractIstTara\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LaraFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LaraFactoryMetaData.ABI instead.
var LaraFactoryABI = LaraFactoryMetaData.ABI

// LaraFactory is an auto generated Go binding around an Ethereum contract.
type LaraFactory struct {
	LaraFactoryCaller     // Read-only binding to the contract
	LaraFactoryTransactor // Write-only binding to the contract
	LaraFactoryFilterer   // Log filterer for contract events
}

// LaraFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LaraFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaraFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LaraFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaraFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LaraFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaraFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LaraFactorySession struct {
	Contract     *LaraFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LaraFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LaraFactoryCallerSession struct {
	Contract *LaraFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LaraFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LaraFactoryTransactorSession struct {
	Contract     *LaraFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LaraFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LaraFactoryRaw struct {
	Contract *LaraFactory // Generic contract binding to access the raw methods on
}

// LaraFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LaraFactoryCallerRaw struct {
	Contract *LaraFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LaraFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LaraFactoryTransactorRaw struct {
	Contract *LaraFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLaraFactory creates a new instance of LaraFactory, bound to a specific deployed contract.
func NewLaraFactory(address common.Address, backend bind.ContractBackend) (*LaraFactory, error) {
	contract, err := bindLaraFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LaraFactory{LaraFactoryCaller: LaraFactoryCaller{contract: contract}, LaraFactoryTransactor: LaraFactoryTransactor{contract: contract}, LaraFactoryFilterer: LaraFactoryFilterer{contract: contract}}, nil
}

// NewLaraFactoryCaller creates a new read-only instance of LaraFactory, bound to a specific deployed contract.
func NewLaraFactoryCaller(address common.Address, caller bind.ContractCaller) (*LaraFactoryCaller, error) {
	contract, err := bindLaraFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LaraFactoryCaller{contract: contract}, nil
}

// NewLaraFactoryTransactor creates a new write-only instance of LaraFactory, bound to a specific deployed contract.
func NewLaraFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LaraFactoryTransactor, error) {
	contract, err := bindLaraFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LaraFactoryTransactor{contract: contract}, nil
}

// NewLaraFactoryFilterer creates a new log filterer instance of LaraFactory, bound to a specific deployed contract.
func NewLaraFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LaraFactoryFilterer, error) {
	contract, err := bindLaraFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LaraFactoryFilterer{contract: contract}, nil
}

// bindLaraFactory binds a generic wrapper to an already deployed contract.
func bindLaraFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LaraFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaraFactory *LaraFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaraFactory.Contract.LaraFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaraFactory *LaraFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraFactory.Contract.LaraFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaraFactory *LaraFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaraFactory.Contract.LaraFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaraFactory *LaraFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaraFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaraFactory *LaraFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaraFactory *LaraFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaraFactory.Contract.contract.Transact(opts, method, params...)
}

// ActiveLaraInstanceCount is a free data retrieval call binding the contract method 0xb2a7d9ee.
//
// Solidity: function activeLaraInstanceCount() view returns(uint32)
func (_LaraFactory *LaraFactoryCaller) ActiveLaraInstanceCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "activeLaraInstanceCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActiveLaraInstanceCount is a free data retrieval call binding the contract method 0xb2a7d9ee.
//
// Solidity: function activeLaraInstanceCount() view returns(uint32)
func (_LaraFactory *LaraFactorySession) ActiveLaraInstanceCount() (uint32, error) {
	return _LaraFactory.Contract.ActiveLaraInstanceCount(&_LaraFactory.CallOpts)
}

// ActiveLaraInstanceCount is a free data retrieval call binding the contract method 0xb2a7d9ee.
//
// Solidity: function activeLaraInstanceCount() view returns(uint32)
func (_LaraFactory *LaraFactoryCallerSession) ActiveLaraInstanceCount() (uint32, error) {
	return _LaraFactory.Contract.ActiveLaraInstanceCount(&_LaraFactory.CallOpts)
}

// ApyOracle is a free data retrieval call binding the contract method 0x627ed636.
//
// Solidity: function apyOracle() view returns(address)
func (_LaraFactory *LaraFactoryCaller) ApyOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "apyOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApyOracle is a free data retrieval call binding the contract method 0x627ed636.
//
// Solidity: function apyOracle() view returns(address)
func (_LaraFactory *LaraFactorySession) ApyOracle() (common.Address, error) {
	return _LaraFactory.Contract.ApyOracle(&_LaraFactory.CallOpts)
}

// ApyOracle is a free data retrieval call binding the contract method 0x627ed636.
//
// Solidity: function apyOracle() view returns(address)
func (_LaraFactory *LaraFactoryCallerSession) ApyOracle() (common.Address, error) {
	return _LaraFactory.Contract.ApyOracle(&_LaraFactory.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_LaraFactory *LaraFactoryCaller) Commission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "commission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_LaraFactory *LaraFactorySession) Commission() (*big.Int, error) {
	return _LaraFactory.Contract.Commission(&_LaraFactory.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_LaraFactory *LaraFactoryCallerSession) Commission() (*big.Int, error) {
	return _LaraFactory.Contract.Commission(&_LaraFactory.CallOpts)
}

// Delegators is a free data retrieval call binding the contract method 0x5be612c7.
//
// Solidity: function delegators(uint256 ) view returns(address)
func (_LaraFactory *LaraFactoryCaller) Delegators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "delegators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegators is a free data retrieval call binding the contract method 0x5be612c7.
//
// Solidity: function delegators(uint256 ) view returns(address)
func (_LaraFactory *LaraFactorySession) Delegators(arg0 *big.Int) (common.Address, error) {
	return _LaraFactory.Contract.Delegators(&_LaraFactory.CallOpts, arg0)
}

// Delegators is a free data retrieval call binding the contract method 0x5be612c7.
//
// Solidity: function delegators(uint256 ) view returns(address)
func (_LaraFactory *LaraFactoryCallerSession) Delegators(arg0 *big.Int) (common.Address, error) {
	return _LaraFactory.Contract.Delegators(&_LaraFactory.CallOpts, arg0)
}

// DposContract is a free data retrieval call binding the contract method 0xe1fb9ae2.
//
// Solidity: function dposContract() view returns(address)
func (_LaraFactory *LaraFactoryCaller) DposContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "dposContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DposContract is a free data retrieval call binding the contract method 0xe1fb9ae2.
//
// Solidity: function dposContract() view returns(address)
func (_LaraFactory *LaraFactorySession) DposContract() (common.Address, error) {
	return _LaraFactory.Contract.DposContract(&_LaraFactory.CallOpts)
}

// DposContract is a free data retrieval call binding the contract method 0xe1fb9ae2.
//
// Solidity: function dposContract() view returns(address)
func (_LaraFactory *LaraFactoryCallerSession) DposContract() (common.Address, error) {
	return _LaraFactory.Contract.DposContract(&_LaraFactory.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_LaraFactory *LaraFactoryCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_LaraFactory *LaraFactorySession) EpochDuration() (*big.Int, error) {
	return _LaraFactory.Contract.EpochDuration(&_LaraFactory.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_LaraFactory *LaraFactoryCallerSession) EpochDuration() (*big.Int, error) {
	return _LaraFactory.Contract.EpochDuration(&_LaraFactory.CallOpts)
}

// LaraActive is a free data retrieval call binding the contract method 0x24b2401b.
//
// Solidity: function laraActive(address ) view returns(bool)
func (_LaraFactory *LaraFactoryCaller) LaraActive(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "laraActive", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LaraActive is a free data retrieval call binding the contract method 0x24b2401b.
//
// Solidity: function laraActive(address ) view returns(bool)
func (_LaraFactory *LaraFactorySession) LaraActive(arg0 common.Address) (bool, error) {
	return _LaraFactory.Contract.LaraActive(&_LaraFactory.CallOpts, arg0)
}

// LaraActive is a free data retrieval call binding the contract method 0x24b2401b.
//
// Solidity: function laraActive(address ) view returns(bool)
func (_LaraFactory *LaraFactoryCallerSession) LaraActive(arg0 common.Address) (bool, error) {
	return _LaraFactory.Contract.LaraActive(&_LaraFactory.CallOpts, arg0)
}

// LaraAddresses is a free data retrieval call binding the contract method 0x1848b00b.
//
// Solidity: function laraAddresses(uint256 ) view returns(address)
func (_LaraFactory *LaraFactoryCaller) LaraAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "laraAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LaraAddresses is a free data retrieval call binding the contract method 0x1848b00b.
//
// Solidity: function laraAddresses(uint256 ) view returns(address)
func (_LaraFactory *LaraFactorySession) LaraAddresses(arg0 *big.Int) (common.Address, error) {
	return _LaraFactory.Contract.LaraAddresses(&_LaraFactory.CallOpts, arg0)
}

// LaraAddresses is a free data retrieval call binding the contract method 0x1848b00b.
//
// Solidity: function laraAddresses(uint256 ) view returns(address)
func (_LaraFactory *LaraFactoryCallerSession) LaraAddresses(arg0 *big.Int) (common.Address, error) {
	return _LaraFactory.Contract.LaraAddresses(&_LaraFactory.CallOpts, arg0)
}

// LaraInstanceCount is a free data retrieval call binding the contract method 0xa1a092eb.
//
// Solidity: function laraInstanceCount() view returns(uint32)
func (_LaraFactory *LaraFactoryCaller) LaraInstanceCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "laraInstanceCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LaraInstanceCount is a free data retrieval call binding the contract method 0xa1a092eb.
//
// Solidity: function laraInstanceCount() view returns(uint32)
func (_LaraFactory *LaraFactorySession) LaraInstanceCount() (uint32, error) {
	return _LaraFactory.Contract.LaraInstanceCount(&_LaraFactory.CallOpts)
}

// LaraInstanceCount is a free data retrieval call binding the contract method 0xa1a092eb.
//
// Solidity: function laraInstanceCount() view returns(uint32)
func (_LaraFactory *LaraFactoryCallerSession) LaraInstanceCount() (uint32, error) {
	return _LaraFactory.Contract.LaraInstanceCount(&_LaraFactory.CallOpts)
}

// LaraInstances is a free data retrieval call binding the contract method 0xdb108425.
//
// Solidity: function laraInstances(address ) view returns(address)
func (_LaraFactory *LaraFactoryCaller) LaraInstances(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "laraInstances", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LaraInstances is a free data retrieval call binding the contract method 0xdb108425.
//
// Solidity: function laraInstances(address ) view returns(address)
func (_LaraFactory *LaraFactorySession) LaraInstances(arg0 common.Address) (common.Address, error) {
	return _LaraFactory.Contract.LaraInstances(&_LaraFactory.CallOpts, arg0)
}

// LaraInstances is a free data retrieval call binding the contract method 0xdb108425.
//
// Solidity: function laraInstances(address ) view returns(address)
func (_LaraFactory *LaraFactoryCallerSession) LaraInstances(arg0 common.Address) (common.Address, error) {
	return _LaraFactory.Contract.LaraInstances(&_LaraFactory.CallOpts, arg0)
}

// MaxValidatorStakeCapacity is a free data retrieval call binding the contract method 0x2a8cf87f.
//
// Solidity: function maxValidatorStakeCapacity() view returns(uint256)
func (_LaraFactory *LaraFactoryCaller) MaxValidatorStakeCapacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "maxValidatorStakeCapacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorStakeCapacity is a free data retrieval call binding the contract method 0x2a8cf87f.
//
// Solidity: function maxValidatorStakeCapacity() view returns(uint256)
func (_LaraFactory *LaraFactorySession) MaxValidatorStakeCapacity() (*big.Int, error) {
	return _LaraFactory.Contract.MaxValidatorStakeCapacity(&_LaraFactory.CallOpts)
}

// MaxValidatorStakeCapacity is a free data retrieval call binding the contract method 0x2a8cf87f.
//
// Solidity: function maxValidatorStakeCapacity() view returns(uint256)
func (_LaraFactory *LaraFactoryCallerSession) MaxValidatorStakeCapacity() (*big.Int, error) {
	return _LaraFactory.Contract.MaxValidatorStakeCapacity(&_LaraFactory.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_LaraFactory *LaraFactoryCaller) MinStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "minStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_LaraFactory *LaraFactorySession) MinStakeAmount() (*big.Int, error) {
	return _LaraFactory.Contract.MinStakeAmount(&_LaraFactory.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_LaraFactory *LaraFactoryCallerSession) MinStakeAmount() (*big.Int, error) {
	return _LaraFactory.Contract.MinStakeAmount(&_LaraFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaraFactory *LaraFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaraFactory *LaraFactorySession) Owner() (common.Address, error) {
	return _LaraFactory.Contract.Owner(&_LaraFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaraFactory *LaraFactoryCallerSession) Owner() (common.Address, error) {
	return _LaraFactory.Contract.Owner(&_LaraFactory.CallOpts)
}

// StTaraToken is a free data retrieval call binding the contract method 0x021b7a81.
//
// Solidity: function stTaraToken() view returns(address)
func (_LaraFactory *LaraFactoryCaller) StTaraToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "stTaraToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StTaraToken is a free data retrieval call binding the contract method 0x021b7a81.
//
// Solidity: function stTaraToken() view returns(address)
func (_LaraFactory *LaraFactorySession) StTaraToken() (common.Address, error) {
	return _LaraFactory.Contract.StTaraToken(&_LaraFactory.CallOpts)
}

// StTaraToken is a free data retrieval call binding the contract method 0x021b7a81.
//
// Solidity: function stTaraToken() view returns(address)
func (_LaraFactory *LaraFactoryCallerSession) StTaraToken() (common.Address, error) {
	return _LaraFactory.Contract.StTaraToken(&_LaraFactory.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_LaraFactory *LaraFactoryCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraFactory.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_LaraFactory *LaraFactorySession) TreasuryAddress() (common.Address, error) {
	return _LaraFactory.Contract.TreasuryAddress(&_LaraFactory.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_LaraFactory *LaraFactoryCallerSession) TreasuryAddress() (common.Address, error) {
	return _LaraFactory.Contract.TreasuryAddress(&_LaraFactory.CallOpts)
}

// ActivateLara is a paid mutator transaction binding the contract method 0x09e00417.
//
// Solidity: function activateLara(address delegator) returns()
func (_LaraFactory *LaraFactoryTransactor) ActivateLara(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "activateLara", delegator)
}

// ActivateLara is a paid mutator transaction binding the contract method 0x09e00417.
//
// Solidity: function activateLara(address delegator) returns()
func (_LaraFactory *LaraFactorySession) ActivateLara(delegator common.Address) (*types.Transaction, error) {
	return _LaraFactory.Contract.ActivateLara(&_LaraFactory.TransactOpts, delegator)
}

// ActivateLara is a paid mutator transaction binding the contract method 0x09e00417.
//
// Solidity: function activateLara(address delegator) returns()
func (_LaraFactory *LaraFactoryTransactorSession) ActivateLara(delegator common.Address) (*types.Transaction, error) {
	return _LaraFactory.Contract.ActivateLara(&_LaraFactory.TransactOpts, delegator)
}

// CreateLara is a paid mutator transaction binding the contract method 0x91086a96.
//
// Solidity: function createLara() returns(address)
func (_LaraFactory *LaraFactoryTransactor) CreateLara(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "createLara")
}

// CreateLara is a paid mutator transaction binding the contract method 0x91086a96.
//
// Solidity: function createLara() returns(address)
func (_LaraFactory *LaraFactorySession) CreateLara() (*types.Transaction, error) {
	return _LaraFactory.Contract.CreateLara(&_LaraFactory.TransactOpts)
}

// CreateLara is a paid mutator transaction binding the contract method 0x91086a96.
//
// Solidity: function createLara() returns(address)
func (_LaraFactory *LaraFactoryTransactorSession) CreateLara() (*types.Transaction, error) {
	return _LaraFactory.Contract.CreateLara(&_LaraFactory.TransactOpts)
}

// DeactivateLara is a paid mutator transaction binding the contract method 0xf5edf160.
//
// Solidity: function deactivateLara(address delegator) returns()
func (_LaraFactory *LaraFactoryTransactor) DeactivateLara(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "deactivateLara", delegator)
}

// DeactivateLara is a paid mutator transaction binding the contract method 0xf5edf160.
//
// Solidity: function deactivateLara(address delegator) returns()
func (_LaraFactory *LaraFactorySession) DeactivateLara(delegator common.Address) (*types.Transaction, error) {
	return _LaraFactory.Contract.DeactivateLara(&_LaraFactory.TransactOpts, delegator)
}

// DeactivateLara is a paid mutator transaction binding the contract method 0xf5edf160.
//
// Solidity: function deactivateLara(address delegator) returns()
func (_LaraFactory *LaraFactoryTransactorSession) DeactivateLara(delegator common.Address) (*types.Transaction, error) {
	return _LaraFactory.Contract.DeactivateLara(&_LaraFactory.TransactOpts, delegator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LaraFactory *LaraFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LaraFactory *LaraFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _LaraFactory.Contract.RenounceOwnership(&_LaraFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LaraFactory *LaraFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LaraFactory.Contract.RenounceOwnership(&_LaraFactory.TransactOpts)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(uint256 _commission) returns()
func (_LaraFactory *LaraFactoryTransactor) SetCommission(opts *bind.TransactOpts, _commission *big.Int) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "setCommission", _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(uint256 _commission) returns()
func (_LaraFactory *LaraFactorySession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetCommission(&_LaraFactory.TransactOpts, _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(uint256 _commission) returns()
func (_LaraFactory *LaraFactoryTransactorSession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetCommission(&_LaraFactory.TransactOpts, _commission)
}

// SetEpochDuration is a paid mutator transaction binding the contract method 0x30024dfe.
//
// Solidity: function setEpochDuration(uint256 _epochDuration) returns()
func (_LaraFactory *LaraFactoryTransactor) SetEpochDuration(opts *bind.TransactOpts, _epochDuration *big.Int) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "setEpochDuration", _epochDuration)
}

// SetEpochDuration is a paid mutator transaction binding the contract method 0x30024dfe.
//
// Solidity: function setEpochDuration(uint256 _epochDuration) returns()
func (_LaraFactory *LaraFactorySession) SetEpochDuration(_epochDuration *big.Int) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetEpochDuration(&_LaraFactory.TransactOpts, _epochDuration)
}

// SetEpochDuration is a paid mutator transaction binding the contract method 0x30024dfe.
//
// Solidity: function setEpochDuration(uint256 _epochDuration) returns()
func (_LaraFactory *LaraFactoryTransactorSession) SetEpochDuration(_epochDuration *big.Int) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetEpochDuration(&_LaraFactory.TransactOpts, _epochDuration)
}

// SetMaxValidatorStakeCapacity is a paid mutator transaction binding the contract method 0x6d2d8519.
//
// Solidity: function setMaxValidatorStakeCapacity(uint256 _maxValidatorStakeCapacity) returns()
func (_LaraFactory *LaraFactoryTransactor) SetMaxValidatorStakeCapacity(opts *bind.TransactOpts, _maxValidatorStakeCapacity *big.Int) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "setMaxValidatorStakeCapacity", _maxValidatorStakeCapacity)
}

// SetMaxValidatorStakeCapacity is a paid mutator transaction binding the contract method 0x6d2d8519.
//
// Solidity: function setMaxValidatorStakeCapacity(uint256 _maxValidatorStakeCapacity) returns()
func (_LaraFactory *LaraFactorySession) SetMaxValidatorStakeCapacity(_maxValidatorStakeCapacity *big.Int) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetMaxValidatorStakeCapacity(&_LaraFactory.TransactOpts, _maxValidatorStakeCapacity)
}

// SetMaxValidatorStakeCapacity is a paid mutator transaction binding the contract method 0x6d2d8519.
//
// Solidity: function setMaxValidatorStakeCapacity(uint256 _maxValidatorStakeCapacity) returns()
func (_LaraFactory *LaraFactoryTransactorSession) SetMaxValidatorStakeCapacity(_maxValidatorStakeCapacity *big.Int) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetMaxValidatorStakeCapacity(&_LaraFactory.TransactOpts, _maxValidatorStakeCapacity)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 _minStakeAmount) returns()
func (_LaraFactory *LaraFactoryTransactor) SetMinStakeAmount(opts *bind.TransactOpts, _minStakeAmount *big.Int) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "setMinStakeAmount", _minStakeAmount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 _minStakeAmount) returns()
func (_LaraFactory *LaraFactorySession) SetMinStakeAmount(_minStakeAmount *big.Int) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetMinStakeAmount(&_LaraFactory.TransactOpts, _minStakeAmount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 _minStakeAmount) returns()
func (_LaraFactory *LaraFactoryTransactorSession) SetMinStakeAmount(_minStakeAmount *big.Int) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetMinStakeAmount(&_LaraFactory.TransactOpts, _minStakeAmount)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_LaraFactory *LaraFactoryTransactor) SetTreasuryAddress(opts *bind.TransactOpts, _treasuryAddress common.Address) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "setTreasuryAddress", _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_LaraFactory *LaraFactorySession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetTreasuryAddress(&_LaraFactory.TransactOpts, _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_LaraFactory *LaraFactoryTransactorSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _LaraFactory.Contract.SetTreasuryAddress(&_LaraFactory.TransactOpts, _treasuryAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaraFactory *LaraFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LaraFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaraFactory *LaraFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LaraFactory.Contract.TransferOwnership(&_LaraFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaraFactory *LaraFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LaraFactory.Contract.TransferOwnership(&_LaraFactory.TransactOpts, newOwner)
}

// LaraFactoryCommissionChangedIterator is returned from FilterCommissionChanged and is used to iterate over the raw logs and unpacked data for CommissionChanged events raised by the LaraFactory contract.
type LaraFactoryCommissionChangedIterator struct {
	Event *LaraFactoryCommissionChanged // Event containing the contract specifics and raw log

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
func (it *LaraFactoryCommissionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraFactoryCommissionChanged)
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
		it.Event = new(LaraFactoryCommissionChanged)
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
func (it *LaraFactoryCommissionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraFactoryCommissionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraFactoryCommissionChanged represents a CommissionChanged event raised by the LaraFactory contract.
type LaraFactoryCommissionChanged struct {
	NewCommission *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommissionChanged is a free log retrieval operation binding the contract event 0x839e4456845dbc05c7d8638cf0b0976161331b5f9163980d71d9a6444a326c61.
//
// Solidity: event CommissionChanged(uint256 indexed newCommission)
func (_LaraFactory *LaraFactoryFilterer) FilterCommissionChanged(opts *bind.FilterOpts, newCommission []*big.Int) (*LaraFactoryCommissionChangedIterator, error) {

	var newCommissionRule []interface{}
	for _, newCommissionItem := range newCommission {
		newCommissionRule = append(newCommissionRule, newCommissionItem)
	}

	logs, sub, err := _LaraFactory.contract.FilterLogs(opts, "CommissionChanged", newCommissionRule)
	if err != nil {
		return nil, err
	}
	return &LaraFactoryCommissionChangedIterator{contract: _LaraFactory.contract, event: "CommissionChanged", logs: logs, sub: sub}, nil
}

// WatchCommissionChanged is a free log subscription operation binding the contract event 0x839e4456845dbc05c7d8638cf0b0976161331b5f9163980d71d9a6444a326c61.
//
// Solidity: event CommissionChanged(uint256 indexed newCommission)
func (_LaraFactory *LaraFactoryFilterer) WatchCommissionChanged(opts *bind.WatchOpts, sink chan<- *LaraFactoryCommissionChanged, newCommission []*big.Int) (event.Subscription, error) {

	var newCommissionRule []interface{}
	for _, newCommissionItem := range newCommission {
		newCommissionRule = append(newCommissionRule, newCommissionItem)
	}

	logs, sub, err := _LaraFactory.contract.WatchLogs(opts, "CommissionChanged", newCommissionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraFactoryCommissionChanged)
				if err := _LaraFactory.contract.UnpackLog(event, "CommissionChanged", log); err != nil {
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

// ParseCommissionChanged is a log parse operation binding the contract event 0x839e4456845dbc05c7d8638cf0b0976161331b5f9163980d71d9a6444a326c61.
//
// Solidity: event CommissionChanged(uint256 indexed newCommission)
func (_LaraFactory *LaraFactoryFilterer) ParseCommissionChanged(log types.Log) (*LaraFactoryCommissionChanged, error) {
	event := new(LaraFactoryCommissionChanged)
	if err := _LaraFactory.contract.UnpackLog(event, "CommissionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraFactoryLaraActivatedIterator is returned from FilterLaraActivated and is used to iterate over the raw logs and unpacked data for LaraActivated events raised by the LaraFactory contract.
type LaraFactoryLaraActivatedIterator struct {
	Event *LaraFactoryLaraActivated // Event containing the contract specifics and raw log

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
func (it *LaraFactoryLaraActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraFactoryLaraActivated)
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
		it.Event = new(LaraFactoryLaraActivated)
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
func (it *LaraFactoryLaraActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraFactoryLaraActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraFactoryLaraActivated represents a LaraActivated event raised by the LaraFactory contract.
type LaraFactoryLaraActivated struct {
	LaraAddress common.Address
	Activator   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLaraActivated is a free log retrieval operation binding the contract event 0x2ce08dd8645384c370d444eba21e5a5796f967af52d548ee812be30a847677c3.
//
// Solidity: event LaraActivated(address indexed laraAddress, address indexed activator)
func (_LaraFactory *LaraFactoryFilterer) FilterLaraActivated(opts *bind.FilterOpts, laraAddress []common.Address, activator []common.Address) (*LaraFactoryLaraActivatedIterator, error) {

	var laraAddressRule []interface{}
	for _, laraAddressItem := range laraAddress {
		laraAddressRule = append(laraAddressRule, laraAddressItem)
	}
	var activatorRule []interface{}
	for _, activatorItem := range activator {
		activatorRule = append(activatorRule, activatorItem)
	}

	logs, sub, err := _LaraFactory.contract.FilterLogs(opts, "LaraActivated", laraAddressRule, activatorRule)
	if err != nil {
		return nil, err
	}
	return &LaraFactoryLaraActivatedIterator{contract: _LaraFactory.contract, event: "LaraActivated", logs: logs, sub: sub}, nil
}

// WatchLaraActivated is a free log subscription operation binding the contract event 0x2ce08dd8645384c370d444eba21e5a5796f967af52d548ee812be30a847677c3.
//
// Solidity: event LaraActivated(address indexed laraAddress, address indexed activator)
func (_LaraFactory *LaraFactoryFilterer) WatchLaraActivated(opts *bind.WatchOpts, sink chan<- *LaraFactoryLaraActivated, laraAddress []common.Address, activator []common.Address) (event.Subscription, error) {

	var laraAddressRule []interface{}
	for _, laraAddressItem := range laraAddress {
		laraAddressRule = append(laraAddressRule, laraAddressItem)
	}
	var activatorRule []interface{}
	for _, activatorItem := range activator {
		activatorRule = append(activatorRule, activatorItem)
	}

	logs, sub, err := _LaraFactory.contract.WatchLogs(opts, "LaraActivated", laraAddressRule, activatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraFactoryLaraActivated)
				if err := _LaraFactory.contract.UnpackLog(event, "LaraActivated", log); err != nil {
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

// ParseLaraActivated is a log parse operation binding the contract event 0x2ce08dd8645384c370d444eba21e5a5796f967af52d548ee812be30a847677c3.
//
// Solidity: event LaraActivated(address indexed laraAddress, address indexed activator)
func (_LaraFactory *LaraFactoryFilterer) ParseLaraActivated(log types.Log) (*LaraFactoryLaraActivated, error) {
	event := new(LaraFactoryLaraActivated)
	if err := _LaraFactory.contract.UnpackLog(event, "LaraActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraFactoryLaraCreatedIterator is returned from FilterLaraCreated and is used to iterate over the raw logs and unpacked data for LaraCreated events raised by the LaraFactory contract.
type LaraFactoryLaraCreatedIterator struct {
	Event *LaraFactoryLaraCreated // Event containing the contract specifics and raw log

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
func (it *LaraFactoryLaraCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraFactoryLaraCreated)
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
		it.Event = new(LaraFactoryLaraCreated)
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
func (it *LaraFactoryLaraCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraFactoryLaraCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraFactoryLaraCreated represents a LaraCreated event raised by the LaraFactory contract.
type LaraFactoryLaraCreated struct {
	LaraAddress common.Address
	Creator     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLaraCreated is a free log retrieval operation binding the contract event 0x5a5f4cf1954ae600ec51dc85fdf015142e874cddc7989c712f7d5cac13e66332.
//
// Solidity: event LaraCreated(address indexed laraAddress, address indexed creator)
func (_LaraFactory *LaraFactoryFilterer) FilterLaraCreated(opts *bind.FilterOpts, laraAddress []common.Address, creator []common.Address) (*LaraFactoryLaraCreatedIterator, error) {

	var laraAddressRule []interface{}
	for _, laraAddressItem := range laraAddress {
		laraAddressRule = append(laraAddressRule, laraAddressItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _LaraFactory.contract.FilterLogs(opts, "LaraCreated", laraAddressRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &LaraFactoryLaraCreatedIterator{contract: _LaraFactory.contract, event: "LaraCreated", logs: logs, sub: sub}, nil
}

// WatchLaraCreated is a free log subscription operation binding the contract event 0x5a5f4cf1954ae600ec51dc85fdf015142e874cddc7989c712f7d5cac13e66332.
//
// Solidity: event LaraCreated(address indexed laraAddress, address indexed creator)
func (_LaraFactory *LaraFactoryFilterer) WatchLaraCreated(opts *bind.WatchOpts, sink chan<- *LaraFactoryLaraCreated, laraAddress []common.Address, creator []common.Address) (event.Subscription, error) {

	var laraAddressRule []interface{}
	for _, laraAddressItem := range laraAddress {
		laraAddressRule = append(laraAddressRule, laraAddressItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _LaraFactory.contract.WatchLogs(opts, "LaraCreated", laraAddressRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraFactoryLaraCreated)
				if err := _LaraFactory.contract.UnpackLog(event, "LaraCreated", log); err != nil {
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

// ParseLaraCreated is a log parse operation binding the contract event 0x5a5f4cf1954ae600ec51dc85fdf015142e874cddc7989c712f7d5cac13e66332.
//
// Solidity: event LaraCreated(address indexed laraAddress, address indexed creator)
func (_LaraFactory *LaraFactoryFilterer) ParseLaraCreated(log types.Log) (*LaraFactoryLaraCreated, error) {
	event := new(LaraFactoryLaraCreated)
	if err := _LaraFactory.contract.UnpackLog(event, "LaraCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraFactoryLaraDeactivatedIterator is returned from FilterLaraDeactivated and is used to iterate over the raw logs and unpacked data for LaraDeactivated events raised by the LaraFactory contract.
type LaraFactoryLaraDeactivatedIterator struct {
	Event *LaraFactoryLaraDeactivated // Event containing the contract specifics and raw log

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
func (it *LaraFactoryLaraDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraFactoryLaraDeactivated)
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
		it.Event = new(LaraFactoryLaraDeactivated)
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
func (it *LaraFactoryLaraDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraFactoryLaraDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraFactoryLaraDeactivated represents a LaraDeactivated event raised by the LaraFactory contract.
type LaraFactoryLaraDeactivated struct {
	LaraAddress common.Address
	Deactivator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLaraDeactivated is a free log retrieval operation binding the contract event 0xf87989512134cd9b9d054b28910e48bdc2c1c41a3c28097da69fde1be5edf610.
//
// Solidity: event LaraDeactivated(address indexed laraAddress, address indexed deactivator)
func (_LaraFactory *LaraFactoryFilterer) FilterLaraDeactivated(opts *bind.FilterOpts, laraAddress []common.Address, deactivator []common.Address) (*LaraFactoryLaraDeactivatedIterator, error) {

	var laraAddressRule []interface{}
	for _, laraAddressItem := range laraAddress {
		laraAddressRule = append(laraAddressRule, laraAddressItem)
	}
	var deactivatorRule []interface{}
	for _, deactivatorItem := range deactivator {
		deactivatorRule = append(deactivatorRule, deactivatorItem)
	}

	logs, sub, err := _LaraFactory.contract.FilterLogs(opts, "LaraDeactivated", laraAddressRule, deactivatorRule)
	if err != nil {
		return nil, err
	}
	return &LaraFactoryLaraDeactivatedIterator{contract: _LaraFactory.contract, event: "LaraDeactivated", logs: logs, sub: sub}, nil
}

// WatchLaraDeactivated is a free log subscription operation binding the contract event 0xf87989512134cd9b9d054b28910e48bdc2c1c41a3c28097da69fde1be5edf610.
//
// Solidity: event LaraDeactivated(address indexed laraAddress, address indexed deactivator)
func (_LaraFactory *LaraFactoryFilterer) WatchLaraDeactivated(opts *bind.WatchOpts, sink chan<- *LaraFactoryLaraDeactivated, laraAddress []common.Address, deactivator []common.Address) (event.Subscription, error) {

	var laraAddressRule []interface{}
	for _, laraAddressItem := range laraAddress {
		laraAddressRule = append(laraAddressRule, laraAddressItem)
	}
	var deactivatorRule []interface{}
	for _, deactivatorItem := range deactivator {
		deactivatorRule = append(deactivatorRule, deactivatorItem)
	}

	logs, sub, err := _LaraFactory.contract.WatchLogs(opts, "LaraDeactivated", laraAddressRule, deactivatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraFactoryLaraDeactivated)
				if err := _LaraFactory.contract.UnpackLog(event, "LaraDeactivated", log); err != nil {
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

// ParseLaraDeactivated is a log parse operation binding the contract event 0xf87989512134cd9b9d054b28910e48bdc2c1c41a3c28097da69fde1be5edf610.
//
// Solidity: event LaraDeactivated(address indexed laraAddress, address indexed deactivator)
func (_LaraFactory *LaraFactoryFilterer) ParseLaraDeactivated(log types.Log) (*LaraFactoryLaraDeactivated, error) {
	event := new(LaraFactoryLaraDeactivated)
	if err := _LaraFactory.contract.UnpackLog(event, "LaraDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LaraFactory contract.
type LaraFactoryOwnershipTransferredIterator struct {
	Event *LaraFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LaraFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraFactoryOwnershipTransferred)
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
		it.Event = new(LaraFactoryOwnershipTransferred)
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
func (it *LaraFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the LaraFactory contract.
type LaraFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaraFactory *LaraFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LaraFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LaraFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LaraFactoryOwnershipTransferredIterator{contract: _LaraFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaraFactory *LaraFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LaraFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LaraFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraFactoryOwnershipTransferred)
				if err := _LaraFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaraFactory *LaraFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*LaraFactoryOwnershipTransferred, error) {
	event := new(LaraFactoryOwnershipTransferred)
	if err := _LaraFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraFactoryTreasuryChangedIterator is returned from FilterTreasuryChanged and is used to iterate over the raw logs and unpacked data for TreasuryChanged events raised by the LaraFactory contract.
type LaraFactoryTreasuryChangedIterator struct {
	Event *LaraFactoryTreasuryChanged // Event containing the contract specifics and raw log

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
func (it *LaraFactoryTreasuryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraFactoryTreasuryChanged)
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
		it.Event = new(LaraFactoryTreasuryChanged)
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
func (it *LaraFactoryTreasuryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraFactoryTreasuryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraFactoryTreasuryChanged represents a TreasuryChanged event raised by the LaraFactory contract.
type LaraFactoryTreasuryChanged struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasuryChanged is a free log retrieval operation binding the contract event 0xc714d22a2f08b695f81e7c707058db484aa5b4d6b4c9fd64beb10fe85832f608.
//
// Solidity: event TreasuryChanged(address indexed newTreasury)
func (_LaraFactory *LaraFactoryFilterer) FilterTreasuryChanged(opts *bind.FilterOpts, newTreasury []common.Address) (*LaraFactoryTreasuryChangedIterator, error) {

	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _LaraFactory.contract.FilterLogs(opts, "TreasuryChanged", newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &LaraFactoryTreasuryChangedIterator{contract: _LaraFactory.contract, event: "TreasuryChanged", logs: logs, sub: sub}, nil
}

// WatchTreasuryChanged is a free log subscription operation binding the contract event 0xc714d22a2f08b695f81e7c707058db484aa5b4d6b4c9fd64beb10fe85832f608.
//
// Solidity: event TreasuryChanged(address indexed newTreasury)
func (_LaraFactory *LaraFactoryFilterer) WatchTreasuryChanged(opts *bind.WatchOpts, sink chan<- *LaraFactoryTreasuryChanged, newTreasury []common.Address) (event.Subscription, error) {

	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _LaraFactory.contract.WatchLogs(opts, "TreasuryChanged", newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraFactoryTreasuryChanged)
				if err := _LaraFactory.contract.UnpackLog(event, "TreasuryChanged", log); err != nil {
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

// ParseTreasuryChanged is a log parse operation binding the contract event 0xc714d22a2f08b695f81e7c707058db484aa5b4d6b4c9fd64beb10fe85832f608.
//
// Solidity: event TreasuryChanged(address indexed newTreasury)
func (_LaraFactory *LaraFactoryFilterer) ParseTreasuryChanged(log types.Log) (*LaraFactoryTreasuryChanged, error) {
	event := new(LaraFactoryTreasuryChanged)
	if err := _LaraFactory.contract.UnpackLog(event, "TreasuryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
