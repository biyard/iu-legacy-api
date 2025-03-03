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

// ProposalSummary is an auto generated low-level Go binding around an user-defined struct.
type ProposalSummary struct {
	ProposalId      *big.Int
	Proposer        common.Address
	Title           string
	ProposalAppName string
	VoteAppName     string
	SubCategory     string
	SubmittedAt     *big.Int
	VoteStatus      uint16
	NumberOfVotes   *big.Int
}

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"stateAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"beforeUpgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"roleNames\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"roleValues\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"}],\"name\":\"changeInitialDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"changeStandardTokenAddressEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"completeSetupEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitAppDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityType\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structProposalSummary\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"emitProposalActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"injectDependancyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setAmountPerOnceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistryEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getRoleList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"getRoleValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStandardTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"handleActivityHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"handleBulkTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"handleTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"injectDependancy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameExt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"to_\",\"type\":\"address[]\"}],\"name\":\"sendTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setAmountPerOnce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transferType\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// BankBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BankBinRuntime = ``

// Bank is an auto generated Go binding around a Klaytn contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around a Klaytn contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around a Klaytn contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around a Klaytn contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// GetRoleList is a free data retrieval call binding the contract method 0x94ed7fd8.
//
// Solidity: function getRoleList() view returns(string[], uint256[])
func (_Bank *BankCaller) GetRoleList(opts *bind.CallOpts) ([]string, []*big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Bank.contract.Call(opts, out, "getRoleList")
	return *ret0, *ret1, err
}

// GetRoleList is a free data retrieval call binding the contract method 0x94ed7fd8.
//
// Solidity: function getRoleList() view returns(string[], uint256[])
func (_Bank *BankSession) GetRoleList() ([]string, []*big.Int, error) {
	return _Bank.Contract.GetRoleList(&_Bank.CallOpts)
}

// GetRoleList is a free data retrieval call binding the contract method 0x94ed7fd8.
//
// Solidity: function getRoleList() view returns(string[], uint256[])
func (_Bank *BankCallerSession) GetRoleList() ([]string, []*big.Int, error) {
	return _Bank.Contract.GetRoleList(&_Bank.CallOpts)
}

// GetRoleValue is a free data retrieval call binding the contract method 0x61a78d6e.
//
// Solidity: function getRoleValue(string role) view returns(uint256)
func (_Bank *BankCaller) GetRoleValue(opts *bind.CallOpts, role string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bank.contract.Call(opts, out, "getRoleValue", role)
	return *ret0, err
}

// GetRoleValue is a free data retrieval call binding the contract method 0x61a78d6e.
//
// Solidity: function getRoleValue(string role) view returns(uint256)
func (_Bank *BankSession) GetRoleValue(role string) (*big.Int, error) {
	return _Bank.Contract.GetRoleValue(&_Bank.CallOpts, role)
}

// GetRoleValue is a free data retrieval call binding the contract method 0x61a78d6e.
//
// Solidity: function getRoleValue(string role) view returns(uint256)
func (_Bank *BankCallerSession) GetRoleValue(role string) (*big.Int, error) {
	return _Bank.Contract.GetRoleValue(&_Bank.CallOpts, role)
}

// GetStandardTokenAddress is a free data retrieval call binding the contract method 0x52dba960.
//
// Solidity: function getStandardTokenAddress() view returns(address)
func (_Bank *BankCaller) GetStandardTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bank.contract.Call(opts, out, "getStandardTokenAddress")
	return *ret0, err
}

// GetStandardTokenAddress is a free data retrieval call binding the contract method 0x52dba960.
//
// Solidity: function getStandardTokenAddress() view returns(address)
func (_Bank *BankSession) GetStandardTokenAddress() (common.Address, error) {
	return _Bank.Contract.GetStandardTokenAddress(&_Bank.CallOpts)
}

// GetStandardTokenAddress is a free data retrieval call binding the contract method 0x52dba960.
//
// Solidity: function getStandardTokenAddress() view returns(address)
func (_Bank *BankCallerSession) GetStandardTokenAddress() (common.Address, error) {
	return _Bank.Contract.GetStandardTokenAddress(&_Bank.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_Bank *BankCaller) GetTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bank.contract.Call(opts, out, "getTotalSupply")
	return *ret0, err
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_Bank *BankSession) GetTotalSupply() (*big.Int, error) {
	return _Bank.Contract.GetTotalSupply(&_Bank.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_Bank *BankCallerSession) GetTotalSupply() (*big.Int, error) {
	return _Bank.Contract.GetTotalSupply(&_Bank.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Bank *BankCaller) NameExt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bank.contract.Call(opts, out, "nameExt")
	return *ret0, err
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Bank *BankSession) NameExt() (string, error) {
	return _Bank.Contract.NameExt(&_Bank.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Bank *BankCallerSession) NameExt() (string, error) {
	return _Bank.Contract.NameExt(&_Bank.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Bank *BankCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bank.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Bank *BankSession) Ready() (bool, error) {
	return _Bank.Contract.Ready(&_Bank.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Bank *BankCallerSession) Ready() (bool, error) {
	return _Bank.Contract.Ready(&_Bank.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Bank *BankCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bank.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Bank *BankSession) Registry() (common.Address, error) {
	return _Bank.Contract.Registry(&_Bank.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Bank *BankCallerSession) Registry() (common.Address, error) {
	return _Bank.Contract.Registry(&_Bank.CallOpts)
}

// HandleActivityHook is a paid mutator transaction binding the contract method 0x9ecd71a7.
//
// Solidity: function handleActivityHook(uint256 activityFlag, bytes data) returns()
func (_Bank *BankTransactor) HandleActivityHook(opts *bind.TransactOpts, activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "handleActivityHook", activityFlag, data)
}

// HandleActivityHook is a paid mutator transaction binding the contract method 0x9ecd71a7.
//
// Solidity: function handleActivityHook(uint256 activityFlag, bytes data) returns()
func (_Bank *BankSession) HandleActivityHook(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _Bank.Contract.HandleActivityHook(&_Bank.TransactOpts, activityFlag, data)
}

// HandleActivityHook is a paid mutator transaction binding the contract method 0x9ecd71a7.
//
// Solidity: function handleActivityHook(uint256 activityFlag, bytes data) returns()
func (_Bank *BankTransactorSession) HandleActivityHook(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _Bank.Contract.HandleActivityHook(&_Bank.TransactOpts, activityFlag, data)
}

// HandleBulkTokenTransfer is a paid mutator transaction binding the contract method 0x8ab47f7e.
//
// Solidity: function handleBulkTokenTransfer(address from, address[] to, uint256[] amount) returns()
func (_Bank *BankTransactor) HandleBulkTokenTransfer(opts *bind.TransactOpts, from common.Address, to []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "handleBulkTokenTransfer", from, to, amount)
}

// HandleBulkTokenTransfer is a paid mutator transaction binding the contract method 0x8ab47f7e.
//
// Solidity: function handleBulkTokenTransfer(address from, address[] to, uint256[] amount) returns()
func (_Bank *BankSession) HandleBulkTokenTransfer(from common.Address, to []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Bank.Contract.HandleBulkTokenTransfer(&_Bank.TransactOpts, from, to, amount)
}

// HandleBulkTokenTransfer is a paid mutator transaction binding the contract method 0x8ab47f7e.
//
// Solidity: function handleBulkTokenTransfer(address from, address[] to, uint256[] amount) returns()
func (_Bank *BankTransactorSession) HandleBulkTokenTransfer(from common.Address, to []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Bank.Contract.HandleBulkTokenTransfer(&_Bank.TransactOpts, from, to, amount)
}

// HandleTokenTransfer is a paid mutator transaction binding the contract method 0xc919a6e6.
//
// Solidity: function handleTokenTransfer(address from, address to, uint256 amount) returns()
func (_Bank *BankTransactor) HandleTokenTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "handleTokenTransfer", from, to, amount)
}

// HandleTokenTransfer is a paid mutator transaction binding the contract method 0xc919a6e6.
//
// Solidity: function handleTokenTransfer(address from, address to, uint256 amount) returns()
func (_Bank *BankSession) HandleTokenTransfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.HandleTokenTransfer(&_Bank.TransactOpts, from, to, amount)
}

// HandleTokenTransfer is a paid mutator transaction binding the contract method 0xc919a6e6.
//
// Solidity: function handleTokenTransfer(address from, address to, uint256 amount) returns()
func (_Bank *BankTransactorSession) HandleTokenTransfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.HandleTokenTransfer(&_Bank.TransactOpts, from, to, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Bank *BankTransactor) Initialize(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "initialize", addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Bank *BankSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Bank.Contract.Initialize(&_Bank.TransactOpts, addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Bank *BankTransactorSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Bank.Contract.Initialize(&_Bank.TransactOpts, addrs)
}

// InjectDependancy is a paid mutator transaction binding the contract method 0x35d07ec4.
//
// Solidity: function injectDependancy(string depName, address addr) returns()
func (_Bank *BankTransactor) InjectDependancy(opts *bind.TransactOpts, depName string, addr common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "injectDependancy", depName, addr)
}

// InjectDependancy is a paid mutator transaction binding the contract method 0x35d07ec4.
//
// Solidity: function injectDependancy(string depName, address addr) returns()
func (_Bank *BankSession) InjectDependancy(depName string, addr common.Address) (*types.Transaction, error) {
	return _Bank.Contract.InjectDependancy(&_Bank.TransactOpts, depName, addr)
}

// InjectDependancy is a paid mutator transaction binding the contract method 0x35d07ec4.
//
// Solidity: function injectDependancy(string depName, address addr) returns()
func (_Bank *BankTransactorSession) InjectDependancy(depName string, addr common.Address) (*types.Transaction, error) {
	return _Bank.Contract.InjectDependancy(&_Bank.TransactOpts, depName, addr)
}

// SendTokens is a paid mutator transaction binding the contract method 0x48fbcaa8.
//
// Solidity: function sendTokens(string role, address[] to_) returns()
func (_Bank *BankTransactor) SendTokens(opts *bind.TransactOpts, role string, to_ []common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "sendTokens", role, to_)
}

// SendTokens is a paid mutator transaction binding the contract method 0x48fbcaa8.
//
// Solidity: function sendTokens(string role, address[] to_) returns()
func (_Bank *BankSession) SendTokens(role string, to_ []common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SendTokens(&_Bank.TransactOpts, role, to_)
}

// SendTokens is a paid mutator transaction binding the contract method 0x48fbcaa8.
//
// Solidity: function sendTokens(string role, address[] to_) returns()
func (_Bank *BankTransactorSession) SendTokens(role string, to_ []common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SendTokens(&_Bank.TransactOpts, role, to_)
}

// SetAmountPerOnce is a paid mutator transaction binding the contract method 0x02dc984e.
//
// Solidity: function setAmountPerOnce(uint256 amount_) returns()
func (_Bank *BankTransactor) SetAmountPerOnce(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "setAmountPerOnce", amount_)
}

// SetAmountPerOnce is a paid mutator transaction binding the contract method 0x02dc984e.
//
// Solidity: function setAmountPerOnce(uint256 amount_) returns()
func (_Bank *BankSession) SetAmountPerOnce(amount_ *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.SetAmountPerOnce(&_Bank.TransactOpts, amount_)
}

// SetAmountPerOnce is a paid mutator transaction binding the contract method 0x02dc984e.
//
// Solidity: function setAmountPerOnce(uint256 amount_) returns()
func (_Bank *BankTransactorSession) SetAmountPerOnce(amount_ *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.SetAmountPerOnce(&_Bank.TransactOpts, amount_)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Bank *BankTransactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "setForwarder", forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Bank *BankSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SetForwarder(&_Bank.TransactOpts, forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Bank *BankTransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SetForwarder(&_Bank.TransactOpts, forwarder)
}

// TransferToken is a paid mutator transaction binding the contract method 0xdbb5313f.
//
// Solidity: function transferToken(uint256 transferType, address addr, uint256 token) returns()
func (_Bank *BankTransactor) TransferToken(opts *bind.TransactOpts, transferType *big.Int, addr common.Address, token *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "transferToken", transferType, addr, token)
}

// TransferToken is a paid mutator transaction binding the contract method 0xdbb5313f.
//
// Solidity: function transferToken(uint256 transferType, address addr, uint256 token) returns()
func (_Bank *BankSession) TransferToken(transferType *big.Int, addr common.Address, token *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.TransferToken(&_Bank.TransactOpts, transferType, addr, token)
}

// TransferToken is a paid mutator transaction binding the contract method 0xdbb5313f.
//
// Solidity: function transferToken(uint256 transferType, address addr, uint256 token) returns()
func (_Bank *BankTransactorSession) TransferToken(transferType *big.Int, addr common.Address, token *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.TransferToken(&_Bank.TransactOpts, transferType, addr, token)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Bank *BankTransactor) Upgrade(opts *bind.TransactOpts, newVersion common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "upgrade", newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Bank *BankSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Bank.Contract.Upgrade(&_Bank.TransactOpts, newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Bank *BankTransactorSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Bank.Contract.Upgrade(&_Bank.TransactOpts, newVersion)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Bank *BankTransactor) UpgradeRegistry(opts *bind.TransactOpts, reg common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "upgradeRegistry", reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Bank *BankSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Bank.Contract.UpgradeRegistry(&_Bank.TransactOpts, reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Bank *BankTransactorSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Bank.Contract.UpgradeRegistry(&_Bank.TransactOpts, reg)
}

// BankBeforeUpgradeHookEventIterator is returned from FilterBeforeUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for BeforeUpgradeHookEvent events raised by the Bank contract.
type BankBeforeUpgradeHookEventIterator struct {
	Event *BankBeforeUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *BankBeforeUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankBeforeUpgradeHookEvent)
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
		it.Event = new(BankBeforeUpgradeHookEvent)
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
func (it *BankBeforeUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankBeforeUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankBeforeUpgradeHookEvent represents a BeforeUpgradeHookEvent event raised by the Bank contract.
type BankBeforeUpgradeHookEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBeforeUpgradeHookEvent is a free log retrieval operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Bank *BankFilterer) FilterBeforeUpgradeHookEvent(opts *bind.FilterOpts) (*BankBeforeUpgradeHookEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &BankBeforeUpgradeHookEventIterator{contract: _Bank.contract, event: "beforeUpgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchBeforeUpgradeHookEvent is a free log subscription operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Bank *BankFilterer) WatchBeforeUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *BankBeforeUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankBeforeUpgradeHookEvent)
				if err := _Bank.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
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

// ParseBeforeUpgradeHookEvent is a log parse operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Bank *BankFilterer) ParseBeforeUpgradeHookEvent(log types.Log) (*BankBeforeUpgradeHookEvent, error) {
	event := new(BankBeforeUpgradeHookEvent)
	if err := _Bank.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankChangeInitialDataEventIterator is returned from FilterChangeInitialDataEvent and is used to iterate over the raw logs and unpacked data for ChangeInitialDataEvent events raised by the Bank contract.
type BankChangeInitialDataEventIterator struct {
	Event *BankChangeInitialDataEvent // Event containing the contract specifics and raw log

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
func (it *BankChangeInitialDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankChangeInitialDataEvent)
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
		it.Event = new(BankChangeInitialDataEvent)
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
func (it *BankChangeInitialDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankChangeInitialDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankChangeInitialDataEvent represents a ChangeInitialDataEvent event raised by the Bank contract.
type BankChangeInitialDataEvent struct {
	RoleNames  []string
	RoleValues []*big.Int
	Ready      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChangeInitialDataEvent is a free log retrieval operation binding the contract event 0x64ef8276728c55b073f0b0c5c0959f0e98649ab9b86eaa3a85d6e74455feaf4e.
//
// Solidity: event changeInitialDataEvent(string[] roleNames, uint256[] roleValues, bool ready)
func (_Bank *BankFilterer) FilterChangeInitialDataEvent(opts *bind.FilterOpts) (*BankChangeInitialDataEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return &BankChangeInitialDataEventIterator{contract: _Bank.contract, event: "changeInitialDataEvent", logs: logs, sub: sub}, nil
}

// WatchChangeInitialDataEvent is a free log subscription operation binding the contract event 0x64ef8276728c55b073f0b0c5c0959f0e98649ab9b86eaa3a85d6e74455feaf4e.
//
// Solidity: event changeInitialDataEvent(string[] roleNames, uint256[] roleValues, bool ready)
func (_Bank *BankFilterer) WatchChangeInitialDataEvent(opts *bind.WatchOpts, sink chan<- *BankChangeInitialDataEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankChangeInitialDataEvent)
				if err := _Bank.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
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

// ParseChangeInitialDataEvent is a log parse operation binding the contract event 0x64ef8276728c55b073f0b0c5c0959f0e98649ab9b86eaa3a85d6e74455feaf4e.
//
// Solidity: event changeInitialDataEvent(string[] roleNames, uint256[] roleValues, bool ready)
func (_Bank *BankFilterer) ParseChangeInitialDataEvent(log types.Log) (*BankChangeInitialDataEvent, error) {
	event := new(BankChangeInitialDataEvent)
	if err := _Bank.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankChangeStandardTokenAddressEventIterator is returned from FilterChangeStandardTokenAddressEvent and is used to iterate over the raw logs and unpacked data for ChangeStandardTokenAddressEvent events raised by the Bank contract.
type BankChangeStandardTokenAddressEventIterator struct {
	Event *BankChangeStandardTokenAddressEvent // Event containing the contract specifics and raw log

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
func (it *BankChangeStandardTokenAddressEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankChangeStandardTokenAddressEvent)
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
		it.Event = new(BankChangeStandardTokenAddressEvent)
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
func (it *BankChangeStandardTokenAddressEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankChangeStandardTokenAddressEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankChangeStandardTokenAddressEvent represents a ChangeStandardTokenAddressEvent event raised by the Bank contract.
type BankChangeStandardTokenAddressEvent struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChangeStandardTokenAddressEvent is a free log retrieval operation binding the contract event 0x7b5835f7d2baac239ac74e0051f234107a6d642e952e9558118259907be9a153.
//
// Solidity: event changeStandardTokenAddressEvent(address addr)
func (_Bank *BankFilterer) FilterChangeStandardTokenAddressEvent(opts *bind.FilterOpts) (*BankChangeStandardTokenAddressEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "changeStandardTokenAddressEvent")
	if err != nil {
		return nil, err
	}
	return &BankChangeStandardTokenAddressEventIterator{contract: _Bank.contract, event: "changeStandardTokenAddressEvent", logs: logs, sub: sub}, nil
}

// WatchChangeStandardTokenAddressEvent is a free log subscription operation binding the contract event 0x7b5835f7d2baac239ac74e0051f234107a6d642e952e9558118259907be9a153.
//
// Solidity: event changeStandardTokenAddressEvent(address addr)
func (_Bank *BankFilterer) WatchChangeStandardTokenAddressEvent(opts *bind.WatchOpts, sink chan<- *BankChangeStandardTokenAddressEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "changeStandardTokenAddressEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankChangeStandardTokenAddressEvent)
				if err := _Bank.contract.UnpackLog(event, "changeStandardTokenAddressEvent", log); err != nil {
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

// ParseChangeStandardTokenAddressEvent is a log parse operation binding the contract event 0x7b5835f7d2baac239ac74e0051f234107a6d642e952e9558118259907be9a153.
//
// Solidity: event changeStandardTokenAddressEvent(address addr)
func (_Bank *BankFilterer) ParseChangeStandardTokenAddressEvent(log types.Log) (*BankChangeStandardTokenAddressEvent, error) {
	event := new(BankChangeStandardTokenAddressEvent)
	if err := _Bank.contract.UnpackLog(event, "changeStandardTokenAddressEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankCompleteSetupEventIterator is returned from FilterCompleteSetupEvent and is used to iterate over the raw logs and unpacked data for CompleteSetupEvent events raised by the Bank contract.
type BankCompleteSetupEventIterator struct {
	Event *BankCompleteSetupEvent // Event containing the contract specifics and raw log

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
func (it *BankCompleteSetupEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankCompleteSetupEvent)
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
		it.Event = new(BankCompleteSetupEvent)
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
func (it *BankCompleteSetupEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankCompleteSetupEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankCompleteSetupEvent represents a CompleteSetupEvent event raised by the Bank contract.
type BankCompleteSetupEvent struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCompleteSetupEvent is a free log retrieval operation binding the contract event 0xefb63b2500e859546a2f900fe8b455236a475f013dc4d91dfa1ee50ddbd78f61.
//
// Solidity: event completeSetupEvent()
func (_Bank *BankFilterer) FilterCompleteSetupEvent(opts *bind.FilterOpts) (*BankCompleteSetupEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "completeSetupEvent")
	if err != nil {
		return nil, err
	}
	return &BankCompleteSetupEventIterator{contract: _Bank.contract, event: "completeSetupEvent", logs: logs, sub: sub}, nil
}

// WatchCompleteSetupEvent is a free log subscription operation binding the contract event 0xefb63b2500e859546a2f900fe8b455236a475f013dc4d91dfa1ee50ddbd78f61.
//
// Solidity: event completeSetupEvent()
func (_Bank *BankFilterer) WatchCompleteSetupEvent(opts *bind.WatchOpts, sink chan<- *BankCompleteSetupEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "completeSetupEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankCompleteSetupEvent)
				if err := _Bank.contract.UnpackLog(event, "completeSetupEvent", log); err != nil {
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

// ParseCompleteSetupEvent is a log parse operation binding the contract event 0xefb63b2500e859546a2f900fe8b455236a475f013dc4d91dfa1ee50ddbd78f61.
//
// Solidity: event completeSetupEvent()
func (_Bank *BankFilterer) ParseCompleteSetupEvent(log types.Log) (*BankCompleteSetupEvent, error) {
	event := new(BankCompleteSetupEvent)
	if err := _Bank.contract.UnpackLog(event, "completeSetupEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankEmitActivityEventIterator is returned from FilterEmitActivityEvent and is used to iterate over the raw logs and unpacked data for EmitActivityEvent events raised by the Bank contract.
type BankEmitActivityEventIterator struct {
	Event *BankEmitActivityEvent // Event containing the contract specifics and raw log

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
func (it *BankEmitActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEmitActivityEvent)
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
		it.Event = new(BankEmitActivityEvent)
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
func (it *BankEmitActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEmitActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEmitActivityEvent represents a EmitActivityEvent event raised by the Bank contract.
type BankEmitActivityEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitActivityEvent is a free log retrieval operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Bank *BankFilterer) FilterEmitActivityEvent(opts *bind.FilterOpts) (*BankEmitActivityEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return &BankEmitActivityEventIterator{contract: _Bank.contract, event: "emitActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitActivityEvent is a free log subscription operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Bank *BankFilterer) WatchEmitActivityEvent(opts *bind.WatchOpts, sink chan<- *BankEmitActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEmitActivityEvent)
				if err := _Bank.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
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

// ParseEmitActivityEvent is a log parse operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Bank *BankFilterer) ParseEmitActivityEvent(log types.Log) (*BankEmitActivityEvent, error) {
	event := new(BankEmitActivityEvent)
	if err := _Bank.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankEmitAppDataEventIterator is returned from FilterEmitAppDataEvent and is used to iterate over the raw logs and unpacked data for EmitAppDataEvent events raised by the Bank contract.
type BankEmitAppDataEventIterator struct {
	Event *BankEmitAppDataEvent // Event containing the contract specifics and raw log

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
func (it *BankEmitAppDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEmitAppDataEvent)
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
		it.Event = new(BankEmitAppDataEvent)
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
func (it *BankEmitAppDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEmitAppDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEmitAppDataEvent represents a EmitAppDataEvent event raised by the Bank contract.
type BankEmitAppDataEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitAppDataEvent is a free log retrieval operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Bank *BankFilterer) FilterEmitAppDataEvent(opts *bind.FilterOpts) (*BankEmitAppDataEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return &BankEmitAppDataEventIterator{contract: _Bank.contract, event: "emitAppDataEvent", logs: logs, sub: sub}, nil
}

// WatchEmitAppDataEvent is a free log subscription operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Bank *BankFilterer) WatchEmitAppDataEvent(opts *bind.WatchOpts, sink chan<- *BankEmitAppDataEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEmitAppDataEvent)
				if err := _Bank.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
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

// ParseEmitAppDataEvent is a log parse operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Bank *BankFilterer) ParseEmitAppDataEvent(log types.Log) (*BankEmitAppDataEvent, error) {
	event := new(BankEmitAppDataEvent)
	if err := _Bank.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankEmitProposalActivityEventIterator is returned from FilterEmitProposalActivityEvent and is used to iterate over the raw logs and unpacked data for EmitProposalActivityEvent events raised by the Bank contract.
type BankEmitProposalActivityEventIterator struct {
	Event *BankEmitProposalActivityEvent // Event containing the contract specifics and raw log

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
func (it *BankEmitProposalActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEmitProposalActivityEvent)
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
		it.Event = new(BankEmitProposalActivityEvent)
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
func (it *BankEmitProposalActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEmitProposalActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEmitProposalActivityEvent represents a EmitProposalActivityEvent event raised by the Bank contract.
type BankEmitProposalActivityEvent struct {
	ActivityType *big.Int
	Proposal     ProposalSummary
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitProposalActivityEvent is a free log retrieval operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Bank *BankFilterer) FilterEmitProposalActivityEvent(opts *bind.FilterOpts) (*BankEmitProposalActivityEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return &BankEmitProposalActivityEventIterator{contract: _Bank.contract, event: "emitProposalActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitProposalActivityEvent is a free log subscription operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Bank *BankFilterer) WatchEmitProposalActivityEvent(opts *bind.WatchOpts, sink chan<- *BankEmitProposalActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEmitProposalActivityEvent)
				if err := _Bank.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
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

// ParseEmitProposalActivityEvent is a log parse operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Bank *BankFilterer) ParseEmitProposalActivityEvent(log types.Log) (*BankEmitProposalActivityEvent, error) {
	event := new(BankEmitProposalActivityEvent)
	if err := _Bank.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the Bank contract.
type BankInitializeEventIterator struct {
	Event *BankInitializeEvent // Event containing the contract specifics and raw log

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
func (it *BankInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankInitializeEvent)
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
		it.Event = new(BankInitializeEvent)
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
func (it *BankInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankInitializeEvent represents a InitializeEvent event raised by the Bank contract.
type BankInitializeEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Bank *BankFilterer) FilterInitializeEvent(opts *bind.FilterOpts) (*BankInitializeEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &BankInitializeEventIterator{contract: _Bank.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Bank *BankFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *BankInitializeEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankInitializeEvent)
				if err := _Bank.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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

// ParseInitializeEvent is a log parse operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Bank *BankFilterer) ParseInitializeEvent(log types.Log) (*BankInitializeEvent, error) {
	event := new(BankInitializeEvent)
	if err := _Bank.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankInitializeHookEventIterator is returned from FilterInitializeHookEvent and is used to iterate over the raw logs and unpacked data for InitializeHookEvent events raised by the Bank contract.
type BankInitializeHookEventIterator struct {
	Event *BankInitializeHookEvent // Event containing the contract specifics and raw log

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
func (it *BankInitializeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankInitializeHookEvent)
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
		it.Event = new(BankInitializeHookEvent)
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
func (it *BankInitializeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankInitializeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankInitializeHookEvent represents a InitializeHookEvent event raised by the Bank contract.
type BankInitializeHookEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeHookEvent is a free log retrieval operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Bank *BankFilterer) FilterInitializeHookEvent(opts *bind.FilterOpts) (*BankInitializeHookEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return &BankInitializeHookEventIterator{contract: _Bank.contract, event: "initializeHookEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeHookEvent is a free log subscription operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Bank *BankFilterer) WatchInitializeHookEvent(opts *bind.WatchOpts, sink chan<- *BankInitializeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankInitializeHookEvent)
				if err := _Bank.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
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

// ParseInitializeHookEvent is a log parse operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Bank *BankFilterer) ParseInitializeHookEvent(log types.Log) (*BankInitializeHookEvent, error) {
	event := new(BankInitializeHookEvent)
	if err := _Bank.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankInjectDependancyEventIterator is returned from FilterInjectDependancyEvent and is used to iterate over the raw logs and unpacked data for InjectDependancyEvent events raised by the Bank contract.
type BankInjectDependancyEventIterator struct {
	Event *BankInjectDependancyEvent // Event containing the contract specifics and raw log

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
func (it *BankInjectDependancyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankInjectDependancyEvent)
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
		it.Event = new(BankInjectDependancyEvent)
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
func (it *BankInjectDependancyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankInjectDependancyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankInjectDependancyEvent represents a InjectDependancyEvent event raised by the Bank contract.
type BankInjectDependancyEvent struct {
	DepName string
	Addr    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInjectDependancyEvent is a free log retrieval operation binding the contract event 0xf8f5332d13b15e842f4d0ce6d8e8adfadaafb61027e41d328afeaa8e1e3cc6cd.
//
// Solidity: event injectDependancyEvent(string depName, address addr)
func (_Bank *BankFilterer) FilterInjectDependancyEvent(opts *bind.FilterOpts) (*BankInjectDependancyEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "injectDependancyEvent")
	if err != nil {
		return nil, err
	}
	return &BankInjectDependancyEventIterator{contract: _Bank.contract, event: "injectDependancyEvent", logs: logs, sub: sub}, nil
}

// WatchInjectDependancyEvent is a free log subscription operation binding the contract event 0xf8f5332d13b15e842f4d0ce6d8e8adfadaafb61027e41d328afeaa8e1e3cc6cd.
//
// Solidity: event injectDependancyEvent(string depName, address addr)
func (_Bank *BankFilterer) WatchInjectDependancyEvent(opts *bind.WatchOpts, sink chan<- *BankInjectDependancyEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "injectDependancyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankInjectDependancyEvent)
				if err := _Bank.contract.UnpackLog(event, "injectDependancyEvent", log); err != nil {
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

// ParseInjectDependancyEvent is a log parse operation binding the contract event 0xf8f5332d13b15e842f4d0ce6d8e8adfadaafb61027e41d328afeaa8e1e3cc6cd.
//
// Solidity: event injectDependancyEvent(string depName, address addr)
func (_Bank *BankFilterer) ParseInjectDependancyEvent(log types.Log) (*BankInjectDependancyEvent, error) {
	event := new(BankInjectDependancyEvent)
	if err := _Bank.contract.UnpackLog(event, "injectDependancyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankSetAmountPerOnceEventIterator is returned from FilterSetAmountPerOnceEvent and is used to iterate over the raw logs and unpacked data for SetAmountPerOnceEvent events raised by the Bank contract.
type BankSetAmountPerOnceEventIterator struct {
	Event *BankSetAmountPerOnceEvent // Event containing the contract specifics and raw log

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
func (it *BankSetAmountPerOnceEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankSetAmountPerOnceEvent)
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
		it.Event = new(BankSetAmountPerOnceEvent)
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
func (it *BankSetAmountPerOnceEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankSetAmountPerOnceEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankSetAmountPerOnceEvent represents a SetAmountPerOnceEvent event raised by the Bank contract.
type BankSetAmountPerOnceEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetAmountPerOnceEvent is a free log retrieval operation binding the contract event 0x368cfa2a9a5bcb34f51c0a687e17e96fb0beda8cd5ad308204e199af53308257.
//
// Solidity: event setAmountPerOnceEvent(uint256 amount_)
func (_Bank *BankFilterer) FilterSetAmountPerOnceEvent(opts *bind.FilterOpts) (*BankSetAmountPerOnceEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "setAmountPerOnceEvent")
	if err != nil {
		return nil, err
	}
	return &BankSetAmountPerOnceEventIterator{contract: _Bank.contract, event: "setAmountPerOnceEvent", logs: logs, sub: sub}, nil
}

// WatchSetAmountPerOnceEvent is a free log subscription operation binding the contract event 0x368cfa2a9a5bcb34f51c0a687e17e96fb0beda8cd5ad308204e199af53308257.
//
// Solidity: event setAmountPerOnceEvent(uint256 amount_)
func (_Bank *BankFilterer) WatchSetAmountPerOnceEvent(opts *bind.WatchOpts, sink chan<- *BankSetAmountPerOnceEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "setAmountPerOnceEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankSetAmountPerOnceEvent)
				if err := _Bank.contract.UnpackLog(event, "setAmountPerOnceEvent", log); err != nil {
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

// ParseSetAmountPerOnceEvent is a log parse operation binding the contract event 0x368cfa2a9a5bcb34f51c0a687e17e96fb0beda8cd5ad308204e199af53308257.
//
// Solidity: event setAmountPerOnceEvent(uint256 amount_)
func (_Bank *BankFilterer) ParseSetAmountPerOnceEvent(log types.Log) (*BankSetAmountPerOnceEvent, error) {
	event := new(BankSetAmountPerOnceEvent)
	if err := _Bank.contract.UnpackLog(event, "setAmountPerOnceEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankSetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the Bank contract.
type BankSetReadyEventIterator struct {
	Event *BankSetReadyEvent // Event containing the contract specifics and raw log

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
func (it *BankSetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankSetReadyEvent)
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
		it.Event = new(BankSetReadyEvent)
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
func (it *BankSetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankSetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankSetReadyEvent represents a SetReadyEvent event raised by the Bank contract.
type BankSetReadyEvent struct {
	Done bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Bank *BankFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*BankSetReadyEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &BankSetReadyEventIterator{contract: _Bank.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Bank *BankFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *BankSetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankSetReadyEvent)
				if err := _Bank.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
// Solidity: event setReadyEvent(bool done)
func (_Bank *BankFilterer) ParseSetReadyEvent(log types.Log) (*BankSetReadyEvent, error) {
	event := new(BankSetReadyEvent)
	if err := _Bank.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankUpgradeEventIterator is returned from FilterUpgradeEvent and is used to iterate over the raw logs and unpacked data for UpgradeEvent events raised by the Bank contract.
type BankUpgradeEventIterator struct {
	Event *BankUpgradeEvent // Event containing the contract specifics and raw log

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
func (it *BankUpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankUpgradeEvent)
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
		it.Event = new(BankUpgradeEvent)
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
func (it *BankUpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankUpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankUpgradeEvent represents a UpgradeEvent event raised by the Bank contract.
type BankUpgradeEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpgradeEvent is a free log retrieval operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Bank *BankFilterer) FilterUpgradeEvent(opts *bind.FilterOpts) (*BankUpgradeEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return &BankUpgradeEventIterator{contract: _Bank.contract, event: "upgradeEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeEvent is a free log subscription operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Bank *BankFilterer) WatchUpgradeEvent(opts *bind.WatchOpts, sink chan<- *BankUpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankUpgradeEvent)
				if err := _Bank.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
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

// ParseUpgradeEvent is a log parse operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Bank *BankFilterer) ParseUpgradeEvent(log types.Log) (*BankUpgradeEvent, error) {
	event := new(BankUpgradeEvent)
	if err := _Bank.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BankUpgradeRegistryEventIterator is returned from FilterUpgradeRegistryEvent and is used to iterate over the raw logs and unpacked data for UpgradeRegistryEvent events raised by the Bank contract.
type BankUpgradeRegistryEventIterator struct {
	Event *BankUpgradeRegistryEvent // Event containing the contract specifics and raw log

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
func (it *BankUpgradeRegistryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankUpgradeRegistryEvent)
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
		it.Event = new(BankUpgradeRegistryEvent)
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
func (it *BankUpgradeRegistryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankUpgradeRegistryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankUpgradeRegistryEvent represents a UpgradeRegistryEvent event raised by the Bank contract.
type BankUpgradeRegistryEvent struct {
	Reg common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpgradeRegistryEvent is a free log retrieval operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Bank *BankFilterer) FilterUpgradeRegistryEvent(opts *bind.FilterOpts) (*BankUpgradeRegistryEventIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return &BankUpgradeRegistryEventIterator{contract: _Bank.contract, event: "upgradeRegistryEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeRegistryEvent is a free log subscription operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Bank *BankFilterer) WatchUpgradeRegistryEvent(opts *bind.WatchOpts, sink chan<- *BankUpgradeRegistryEvent) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankUpgradeRegistryEvent)
				if err := _Bank.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
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

// ParseUpgradeRegistryEvent is a log parse operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Bank *BankFilterer) ParseUpgradeRegistryEvent(log types.Log) (*BankUpgradeRegistryEvent, error) {
	event := new(BankUpgradeRegistryEvent)
	if err := _Bank.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
