// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StandardTokenMetaData contains all meta data concerning the StandardToken contract.
var StandardTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addTransferHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"bulkTransferFromEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"delTransferHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"migrateOperatorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintTokenEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"senderAddressEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addTransferHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"bulkTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"delTransferHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listTransferHook\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"migrateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StandardTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use StandardTokenMetaData.ABI instead.
var StandardTokenABI = StandardTokenMetaData.ABI

// StandardTokenBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const StandardTokenBinRuntime = ``

// StandardToken is an auto generated Go binding around a Klaytn contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around a Klaytn contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around a Klaytn contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around a Klaytn contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StandardTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StandardToken *StandardTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StandardToken *StandardTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StandardToken *StandardTokenSession) Decimals() (uint8, error) {
	return _StandardToken.Contract.Decimals(&_StandardToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StandardToken *StandardTokenCallerSession) Decimals() (uint8, error) {
	return _StandardToken.Contract.Decimals(&_StandardToken.CallOpts)
}

// ListTransferHook is a free data retrieval call binding the contract method 0xd0907dae.
//
// Solidity: function listTransferHook() view returns(address[])
func (_StandardToken *StandardTokenCaller) ListTransferHook(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "listTransferHook")
	return *ret0, err
}

// ListTransferHook is a free data retrieval call binding the contract method 0xd0907dae.
//
// Solidity: function listTransferHook() view returns(address[])
func (_StandardToken *StandardTokenSession) ListTransferHook() ([]common.Address, error) {
	return _StandardToken.Contract.ListTransferHook(&_StandardToken.CallOpts)
}

// ListTransferHook is a free data retrieval call binding the contract method 0xd0907dae.
//
// Solidity: function listTransferHook() view returns(address[])
func (_StandardToken *StandardTokenCallerSession) ListTransferHook() ([]common.Address, error) {
	return _StandardToken.Contract.ListTransferHook(&_StandardToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StandardToken *StandardTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StandardToken *StandardTokenSession) Name() (string, error) {
	return _StandardToken.Contract.Name(&_StandardToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StandardToken *StandardTokenCallerSession) Name() (string, error) {
	return _StandardToken.Contract.Name(&_StandardToken.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_StandardToken *StandardTokenCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_StandardToken *StandardTokenSession) Ready() (bool, error) {
	return _StandardToken.Contract.Ready(&_StandardToken.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_StandardToken *StandardTokenCallerSession) Ready() (bool, error) {
	return _StandardToken.Contract.Ready(&_StandardToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_StandardToken *StandardTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_StandardToken *StandardTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StandardToken.Contract.SupportsInterface(&_StandardToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_StandardToken *StandardTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StandardToken.Contract.SupportsInterface(&_StandardToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StandardToken *StandardTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StandardToken *StandardTokenSession) Symbol() (string, error) {
	return _StandardToken.Contract.Symbol(&_StandardToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StandardToken *StandardTokenCallerSession) Symbol() (string, error) {
	return _StandardToken.Contract.Symbol(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// AddTransferHook is a paid mutator transaction binding the contract method 0x6b7e8644.
//
// Solidity: function addTransferHook(address addr) returns()
func (_StandardToken *StandardTokenTransactor) AddTransferHook(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "addTransferHook", addr)
}

// AddTransferHook is a paid mutator transaction binding the contract method 0x6b7e8644.
//
// Solidity: function addTransferHook(address addr) returns()
func (_StandardToken *StandardTokenSession) AddTransferHook(addr common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.AddTransferHook(&_StandardToken.TransactOpts, addr)
}

// AddTransferHook is a paid mutator transaction binding the contract method 0x6b7e8644.
//
// Solidity: function addTransferHook(address addr) returns()
func (_StandardToken *StandardTokenTransactorSession) AddTransferHook(addr common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.AddTransferHook(&_StandardToken.TransactOpts, addr)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, spender, amount)
}

// BulkTransferFrom is a paid mutator transaction binding the contract method 0xba7b52e0.
//
// Solidity: function bulkTransferFrom(address sender, address[] tos, uint256[] values) returns()
func (_StandardToken *StandardTokenTransactor) BulkTransferFrom(opts *bind.TransactOpts, sender common.Address, tos []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "bulkTransferFrom", sender, tos, values)
}

// BulkTransferFrom is a paid mutator transaction binding the contract method 0xba7b52e0.
//
// Solidity: function bulkTransferFrom(address sender, address[] tos, uint256[] values) returns()
func (_StandardToken *StandardTokenSession) BulkTransferFrom(sender common.Address, tos []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.BulkTransferFrom(&_StandardToken.TransactOpts, sender, tos, values)
}

// BulkTransferFrom is a paid mutator transaction binding the contract method 0xba7b52e0.
//
// Solidity: function bulkTransferFrom(address sender, address[] tos, uint256[] values) returns()
func (_StandardToken *StandardTokenTransactorSession) BulkTransferFrom(sender common.Address, tos []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.BulkTransferFrom(&_StandardToken.TransactOpts, sender, tos, values)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StandardToken *StandardTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StandardToken *StandardTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseAllowance(&_StandardToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseAllowance(&_StandardToken.TransactOpts, spender, subtractedValue)
}

// DelTransferHook is a paid mutator transaction binding the contract method 0x20ed083b.
//
// Solidity: function delTransferHook(address addr) returns()
func (_StandardToken *StandardTokenTransactor) DelTransferHook(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "delTransferHook", addr)
}

// DelTransferHook is a paid mutator transaction binding the contract method 0x20ed083b.
//
// Solidity: function delTransferHook(address addr) returns()
func (_StandardToken *StandardTokenSession) DelTransferHook(addr common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.DelTransferHook(&_StandardToken.TransactOpts, addr)
}

// DelTransferHook is a paid mutator transaction binding the contract method 0x20ed083b.
//
// Solidity: function delTransferHook(address addr) returns()
func (_StandardToken *StandardTokenTransactorSession) DelTransferHook(addr common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.DelTransferHook(&_StandardToken.TransactOpts, addr)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StandardToken *StandardTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StandardToken *StandardTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseAllowance(&_StandardToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseAllowance(&_StandardToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc9aba0aa.
//
// Solidity: function initialize(string n, string s, address operator, address , uint256 initialSupply) returns()
func (_StandardToken *StandardTokenTransactor) Initialize(opts *bind.TransactOpts, n string, s string, operator common.Address, arg3 common.Address, initialSupply *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "initialize", n, s, operator, arg3, initialSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0xc9aba0aa.
//
// Solidity: function initialize(string n, string s, address operator, address , uint256 initialSupply) returns()
func (_StandardToken *StandardTokenSession) Initialize(n string, s string, operator common.Address, arg3 common.Address, initialSupply *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Initialize(&_StandardToken.TransactOpts, n, s, operator, arg3, initialSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0xc9aba0aa.
//
// Solidity: function initialize(string n, string s, address operator, address , uint256 initialSupply) returns()
func (_StandardToken *StandardTokenTransactorSession) Initialize(n string, s string, operator common.Address, arg3 common.Address, initialSupply *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Initialize(&_StandardToken.TransactOpts, n, s, operator, arg3, initialSupply)
}

// MigrateOperator is a paid mutator transaction binding the contract method 0x6379808f.
//
// Solidity: function migrateOperator(address newOperator) returns()
func (_StandardToken *StandardTokenTransactor) MigrateOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "migrateOperator", newOperator)
}

// MigrateOperator is a paid mutator transaction binding the contract method 0x6379808f.
//
// Solidity: function migrateOperator(address newOperator) returns()
func (_StandardToken *StandardTokenSession) MigrateOperator(newOperator common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.MigrateOperator(&_StandardToken.TransactOpts, newOperator)
}

// MigrateOperator is a paid mutator transaction binding the contract method 0x6379808f.
//
// Solidity: function migrateOperator(address newOperator) returns()
func (_StandardToken *StandardTokenTransactorSession) MigrateOperator(newOperator common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.MigrateOperator(&_StandardToken.TransactOpts, newOperator)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(address account, uint256 amount) returns()
func (_StandardToken *StandardTokenTransactor) MintToken(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "mintToken", account, amount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(address account, uint256 amount) returns()
func (_StandardToken *StandardTokenSession) MintToken(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.MintToken(&_StandardToken.TransactOpts, account, amount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(address account, uint256 amount) returns()
func (_StandardToken *StandardTokenTransactorSession) MintToken(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.MintToken(&_StandardToken.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, from, to, amount)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StandardToken *StandardTokenFilterer) ParseApproval(log types.Log) (*StandardTokenApproval, error) {
	event := new(StandardTokenApproval)
	if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StandardToken *StandardTokenFilterer) ParseTransfer(log types.Log) (*StandardTokenTransfer, error) {
	event := new(StandardTokenTransfer)
	if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenAddTransferHookEventIterator is returned from FilterAddTransferHookEvent and is used to iterate over the raw logs and unpacked data for AddTransferHookEvent events raised by the StandardToken contract.
type StandardTokenAddTransferHookEventIterator struct {
	Event *StandardTokenAddTransferHookEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenAddTransferHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenAddTransferHookEvent)
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
		it.Event = new(StandardTokenAddTransferHookEvent)
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
func (it *StandardTokenAddTransferHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenAddTransferHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenAddTransferHookEvent represents a AddTransferHookEvent event raised by the StandardToken contract.
type StandardTokenAddTransferHookEvent struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddTransferHookEvent is a free log retrieval operation binding the contract event 0x51f432aaab3c23d4e0ea969adac50edb11043d50540b7d335ca4c5843a5e7628.
//
// Solidity: event addTransferHookEvent(address addr)
func (_StandardToken *StandardTokenFilterer) FilterAddTransferHookEvent(opts *bind.FilterOpts) (*StandardTokenAddTransferHookEventIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "addTransferHookEvent")
	if err != nil {
		return nil, err
	}
	return &StandardTokenAddTransferHookEventIterator{contract: _StandardToken.contract, event: "addTransferHookEvent", logs: logs, sub: sub}, nil
}

// WatchAddTransferHookEvent is a free log subscription operation binding the contract event 0x51f432aaab3c23d4e0ea969adac50edb11043d50540b7d335ca4c5843a5e7628.
//
// Solidity: event addTransferHookEvent(address addr)
func (_StandardToken *StandardTokenFilterer) WatchAddTransferHookEvent(opts *bind.WatchOpts, sink chan<- *StandardTokenAddTransferHookEvent) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "addTransferHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenAddTransferHookEvent)
				if err := _StandardToken.contract.UnpackLog(event, "addTransferHookEvent", log); err != nil {
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

// ParseAddTransferHookEvent is a log parse operation binding the contract event 0x51f432aaab3c23d4e0ea969adac50edb11043d50540b7d335ca4c5843a5e7628.
//
// Solidity: event addTransferHookEvent(address addr)
func (_StandardToken *StandardTokenFilterer) ParseAddTransferHookEvent(log types.Log) (*StandardTokenAddTransferHookEvent, error) {
	event := new(StandardTokenAddTransferHookEvent)
	if err := _StandardToken.contract.UnpackLog(event, "addTransferHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenBulkTransferFromEventIterator is returned from FilterBulkTransferFromEvent and is used to iterate over the raw logs and unpacked data for BulkTransferFromEvent events raised by the StandardToken contract.
type StandardTokenBulkTransferFromEventIterator struct {
	Event *StandardTokenBulkTransferFromEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenBulkTransferFromEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenBulkTransferFromEvent)
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
		it.Event = new(StandardTokenBulkTransferFromEvent)
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
func (it *StandardTokenBulkTransferFromEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenBulkTransferFromEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenBulkTransferFromEvent represents a BulkTransferFromEvent event raised by the StandardToken contract.
type StandardTokenBulkTransferFromEvent struct {
	Sender common.Address
	Tos    []common.Address
	Values []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBulkTransferFromEvent is a free log retrieval operation binding the contract event 0xef8ba5e737fefc545d57c58c314399d9ff5cc2e19a8c95011cddb89036e50b9f.
//
// Solidity: event bulkTransferFromEvent(address sender, address[] tos, uint256[] values)
func (_StandardToken *StandardTokenFilterer) FilterBulkTransferFromEvent(opts *bind.FilterOpts) (*StandardTokenBulkTransferFromEventIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "bulkTransferFromEvent")
	if err != nil {
		return nil, err
	}
	return &StandardTokenBulkTransferFromEventIterator{contract: _StandardToken.contract, event: "bulkTransferFromEvent", logs: logs, sub: sub}, nil
}

// WatchBulkTransferFromEvent is a free log subscription operation binding the contract event 0xef8ba5e737fefc545d57c58c314399d9ff5cc2e19a8c95011cddb89036e50b9f.
//
// Solidity: event bulkTransferFromEvent(address sender, address[] tos, uint256[] values)
func (_StandardToken *StandardTokenFilterer) WatchBulkTransferFromEvent(opts *bind.WatchOpts, sink chan<- *StandardTokenBulkTransferFromEvent) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "bulkTransferFromEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenBulkTransferFromEvent)
				if err := _StandardToken.contract.UnpackLog(event, "bulkTransferFromEvent", log); err != nil {
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

// ParseBulkTransferFromEvent is a log parse operation binding the contract event 0xef8ba5e737fefc545d57c58c314399d9ff5cc2e19a8c95011cddb89036e50b9f.
//
// Solidity: event bulkTransferFromEvent(address sender, address[] tos, uint256[] values)
func (_StandardToken *StandardTokenFilterer) ParseBulkTransferFromEvent(log types.Log) (*StandardTokenBulkTransferFromEvent, error) {
	event := new(StandardTokenBulkTransferFromEvent)
	if err := _StandardToken.contract.UnpackLog(event, "bulkTransferFromEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenDelTransferHookEventIterator is returned from FilterDelTransferHookEvent and is used to iterate over the raw logs and unpacked data for DelTransferHookEvent events raised by the StandardToken contract.
type StandardTokenDelTransferHookEventIterator struct {
	Event *StandardTokenDelTransferHookEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenDelTransferHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenDelTransferHookEvent)
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
		it.Event = new(StandardTokenDelTransferHookEvent)
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
func (it *StandardTokenDelTransferHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenDelTransferHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenDelTransferHookEvent represents a DelTransferHookEvent event raised by the StandardToken contract.
type StandardTokenDelTransferHookEvent struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDelTransferHookEvent is a free log retrieval operation binding the contract event 0x76c88f6301009d9137965385f378e8833decf1a8599d36067ffffd96c6a91207.
//
// Solidity: event delTransferHookEvent(address addr)
func (_StandardToken *StandardTokenFilterer) FilterDelTransferHookEvent(opts *bind.FilterOpts) (*StandardTokenDelTransferHookEventIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "delTransferHookEvent")
	if err != nil {
		return nil, err
	}
	return &StandardTokenDelTransferHookEventIterator{contract: _StandardToken.contract, event: "delTransferHookEvent", logs: logs, sub: sub}, nil
}

// WatchDelTransferHookEvent is a free log subscription operation binding the contract event 0x76c88f6301009d9137965385f378e8833decf1a8599d36067ffffd96c6a91207.
//
// Solidity: event delTransferHookEvent(address addr)
func (_StandardToken *StandardTokenFilterer) WatchDelTransferHookEvent(opts *bind.WatchOpts, sink chan<- *StandardTokenDelTransferHookEvent) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "delTransferHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenDelTransferHookEvent)
				if err := _StandardToken.contract.UnpackLog(event, "delTransferHookEvent", log); err != nil {
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

// ParseDelTransferHookEvent is a log parse operation binding the contract event 0x76c88f6301009d9137965385f378e8833decf1a8599d36067ffffd96c6a91207.
//
// Solidity: event delTransferHookEvent(address addr)
func (_StandardToken *StandardTokenFilterer) ParseDelTransferHookEvent(log types.Log) (*StandardTokenDelTransferHookEvent, error) {
	event := new(StandardTokenDelTransferHookEvent)
	if err := _StandardToken.contract.UnpackLog(event, "delTransferHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the StandardToken contract.
type StandardTokenInitializeEventIterator struct {
	Event *StandardTokenInitializeEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenInitializeEvent)
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
		it.Event = new(StandardTokenInitializeEvent)
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
func (it *StandardTokenInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenInitializeEvent represents a InitializeEvent event raised by the StandardToken contract.
type StandardTokenInitializeEvent struct {
	N             string
	S             string
	Operator      common.Address
	InitialSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0x240c8f2cf78ca68c268d6aebab132a67dda404f67587f2e8fdfe616d517c8e6c.
//
// Solidity: event initializeEvent(string n, string s, address operator, uint256 initialSupply)
func (_StandardToken *StandardTokenFilterer) FilterInitializeEvent(opts *bind.FilterOpts) (*StandardTokenInitializeEventIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &StandardTokenInitializeEventIterator{contract: _StandardToken.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0x240c8f2cf78ca68c268d6aebab132a67dda404f67587f2e8fdfe616d517c8e6c.
//
// Solidity: event initializeEvent(string n, string s, address operator, uint256 initialSupply)
func (_StandardToken *StandardTokenFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *StandardTokenInitializeEvent) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenInitializeEvent)
				if err := _StandardToken.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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

// ParseInitializeEvent is a log parse operation binding the contract event 0x240c8f2cf78ca68c268d6aebab132a67dda404f67587f2e8fdfe616d517c8e6c.
//
// Solidity: event initializeEvent(string n, string s, address operator, uint256 initialSupply)
func (_StandardToken *StandardTokenFilterer) ParseInitializeEvent(log types.Log) (*StandardTokenInitializeEvent, error) {
	event := new(StandardTokenInitializeEvent)
	if err := _StandardToken.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenMigrateOperatorEventIterator is returned from FilterMigrateOperatorEvent and is used to iterate over the raw logs and unpacked data for MigrateOperatorEvent events raised by the StandardToken contract.
type StandardTokenMigrateOperatorEventIterator struct {
	Event *StandardTokenMigrateOperatorEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenMigrateOperatorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenMigrateOperatorEvent)
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
		it.Event = new(StandardTokenMigrateOperatorEvent)
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
func (it *StandardTokenMigrateOperatorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenMigrateOperatorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenMigrateOperatorEvent represents a MigrateOperatorEvent event raised by the StandardToken contract.
type StandardTokenMigrateOperatorEvent struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMigrateOperatorEvent is a free log retrieval operation binding the contract event 0x099d2c6de184bc25866ad9bf894425f5af26d7970ba0154ab19feddaae24a78c.
//
// Solidity: event migrateOperatorEvent(address newOperator)
func (_StandardToken *StandardTokenFilterer) FilterMigrateOperatorEvent(opts *bind.FilterOpts) (*StandardTokenMigrateOperatorEventIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "migrateOperatorEvent")
	if err != nil {
		return nil, err
	}
	return &StandardTokenMigrateOperatorEventIterator{contract: _StandardToken.contract, event: "migrateOperatorEvent", logs: logs, sub: sub}, nil
}

// WatchMigrateOperatorEvent is a free log subscription operation binding the contract event 0x099d2c6de184bc25866ad9bf894425f5af26d7970ba0154ab19feddaae24a78c.
//
// Solidity: event migrateOperatorEvent(address newOperator)
func (_StandardToken *StandardTokenFilterer) WatchMigrateOperatorEvent(opts *bind.WatchOpts, sink chan<- *StandardTokenMigrateOperatorEvent) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "migrateOperatorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenMigrateOperatorEvent)
				if err := _StandardToken.contract.UnpackLog(event, "migrateOperatorEvent", log); err != nil {
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

// ParseMigrateOperatorEvent is a log parse operation binding the contract event 0x099d2c6de184bc25866ad9bf894425f5af26d7970ba0154ab19feddaae24a78c.
//
// Solidity: event migrateOperatorEvent(address newOperator)
func (_StandardToken *StandardTokenFilterer) ParseMigrateOperatorEvent(log types.Log) (*StandardTokenMigrateOperatorEvent, error) {
	event := new(StandardTokenMigrateOperatorEvent)
	if err := _StandardToken.contract.UnpackLog(event, "migrateOperatorEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenMintTokenEventIterator is returned from FilterMintTokenEvent and is used to iterate over the raw logs and unpacked data for MintTokenEvent events raised by the StandardToken contract.
type StandardTokenMintTokenEventIterator struct {
	Event *StandardTokenMintTokenEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenMintTokenEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenMintTokenEvent)
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
		it.Event = new(StandardTokenMintTokenEvent)
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
func (it *StandardTokenMintTokenEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenMintTokenEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenMintTokenEvent represents a MintTokenEvent event raised by the StandardToken contract.
type StandardTokenMintTokenEvent struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintTokenEvent is a free log retrieval operation binding the contract event 0x91bc8230fb99b47a7ef868ae2c20e3226cc7d7470f21f854b82fe8af9285a109.
//
// Solidity: event mintTokenEvent(address account, uint256 amount)
func (_StandardToken *StandardTokenFilterer) FilterMintTokenEvent(opts *bind.FilterOpts) (*StandardTokenMintTokenEventIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "mintTokenEvent")
	if err != nil {
		return nil, err
	}
	return &StandardTokenMintTokenEventIterator{contract: _StandardToken.contract, event: "mintTokenEvent", logs: logs, sub: sub}, nil
}

// WatchMintTokenEvent is a free log subscription operation binding the contract event 0x91bc8230fb99b47a7ef868ae2c20e3226cc7d7470f21f854b82fe8af9285a109.
//
// Solidity: event mintTokenEvent(address account, uint256 amount)
func (_StandardToken *StandardTokenFilterer) WatchMintTokenEvent(opts *bind.WatchOpts, sink chan<- *StandardTokenMintTokenEvent) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "mintTokenEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenMintTokenEvent)
				if err := _StandardToken.contract.UnpackLog(event, "mintTokenEvent", log); err != nil {
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

// ParseMintTokenEvent is a log parse operation binding the contract event 0x91bc8230fb99b47a7ef868ae2c20e3226cc7d7470f21f854b82fe8af9285a109.
//
// Solidity: event mintTokenEvent(address account, uint256 amount)
func (_StandardToken *StandardTokenFilterer) ParseMintTokenEvent(log types.Log) (*StandardTokenMintTokenEvent, error) {
	event := new(StandardTokenMintTokenEvent)
	if err := _StandardToken.contract.UnpackLog(event, "mintTokenEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenSenderAddressEventIterator is returned from FilterSenderAddressEvent and is used to iterate over the raw logs and unpacked data for SenderAddressEvent events raised by the StandardToken contract.
type StandardTokenSenderAddressEventIterator struct {
	Event *StandardTokenSenderAddressEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenSenderAddressEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenSenderAddressEvent)
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
		it.Event = new(StandardTokenSenderAddressEvent)
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
func (it *StandardTokenSenderAddressEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenSenderAddressEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenSenderAddressEvent represents a SenderAddressEvent event raised by the StandardToken contract.
type StandardTokenSenderAddressEvent struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSenderAddressEvent is a free log retrieval operation binding the contract event 0xb76038372324eb6680d39a1933336f5268eb00faac1e6c7cb3b33c9bc4f0418b.
//
// Solidity: event senderAddressEvent(address arg0)
func (_StandardToken *StandardTokenFilterer) FilterSenderAddressEvent(opts *bind.FilterOpts) (*StandardTokenSenderAddressEventIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "senderAddressEvent")
	if err != nil {
		return nil, err
	}
	return &StandardTokenSenderAddressEventIterator{contract: _StandardToken.contract, event: "senderAddressEvent", logs: logs, sub: sub}, nil
}

// WatchSenderAddressEvent is a free log subscription operation binding the contract event 0xb76038372324eb6680d39a1933336f5268eb00faac1e6c7cb3b33c9bc4f0418b.
//
// Solidity: event senderAddressEvent(address arg0)
func (_StandardToken *StandardTokenFilterer) WatchSenderAddressEvent(opts *bind.WatchOpts, sink chan<- *StandardTokenSenderAddressEvent) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "senderAddressEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenSenderAddressEvent)
				if err := _StandardToken.contract.UnpackLog(event, "senderAddressEvent", log); err != nil {
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

// ParseSenderAddressEvent is a log parse operation binding the contract event 0xb76038372324eb6680d39a1933336f5268eb00faac1e6c7cb3b33c9bc4f0418b.
//
// Solidity: event senderAddressEvent(address arg0)
func (_StandardToken *StandardTokenFilterer) ParseSenderAddressEvent(log types.Log) (*StandardTokenSenderAddressEvent, error) {
	event := new(StandardTokenSenderAddressEvent)
	if err := _StandardToken.contract.UnpackLog(event, "senderAddressEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenSetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the StandardToken contract.
type StandardTokenSetReadyEventIterator struct {
	Event *StandardTokenSetReadyEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenSetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenSetReadyEvent)
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
		it.Event = new(StandardTokenSetReadyEvent)
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
func (it *StandardTokenSetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenSetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenSetReadyEvent represents a SetReadyEvent event raised by the StandardToken contract.
type StandardTokenSetReadyEvent struct {
	R   bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool r)
func (_StandardToken *StandardTokenFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*StandardTokenSetReadyEventIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &StandardTokenSetReadyEventIterator{contract: _StandardToken.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool r)
func (_StandardToken *StandardTokenFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *StandardTokenSetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenSetReadyEvent)
				if err := _StandardToken.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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

// ParseSetReadyEvent is a log parse operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool r)
func (_StandardToken *StandardTokenFilterer) ParseSetReadyEvent(log types.Log) (*StandardTokenSetReadyEvent, error) {
	event := new(StandardTokenSetReadyEvent)
	if err := _StandardToken.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
