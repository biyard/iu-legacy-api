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

// StandardMTMetaData contains all meta data concerning the StandardMT contract.
var StandardMTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"changeTokenUriEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"u\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ptokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pi\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vtokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"vi\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ctokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ci\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"migrateOperatorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintSBTEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"changeTokenUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTokenUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"u\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ptokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pi\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vtokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"vi\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ctokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ci\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"migrateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StandardMTABI is the input ABI used to generate the binding from.
// Deprecated: Use StandardMTMetaData.ABI instead.
var StandardMTABI = StandardMTMetaData.ABI

// StandardMTBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const StandardMTBinRuntime = ``

// StandardMT is an auto generated Go binding around a Klaytn contract.
type StandardMT struct {
	StandardMTCaller     // Read-only binding to the contract
	StandardMTTransactor // Write-only binding to the contract
	StandardMTFilterer   // Log filterer for contract events
}

// StandardMTCaller is an auto generated read-only Go binding around a Klaytn contract.
type StandardMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardMTTransactor is an auto generated write-only Go binding around a Klaytn contract.
type StandardMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardMTFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type StandardMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardMTSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type StandardMTSession struct {
	Contract     *StandardMT       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardMTCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type StandardMTCallerSession struct {
	Contract *StandardMTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StandardMTTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type StandardMTTransactorSession struct {
	Contract     *StandardMTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StandardMTRaw is an auto generated low-level Go binding around a Klaytn contract.
type StandardMTRaw struct {
	Contract *StandardMT // Generic contract binding to access the raw methods on
}

// StandardMTCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type StandardMTCallerRaw struct {
	Contract *StandardMTCaller // Generic read-only contract binding to access the raw methods on
}

// StandardMTTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type StandardMTTransactorRaw struct {
	Contract *StandardMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardMT creates a new instance of StandardMT, bound to a specific deployed contract.
func NewStandardMT(address common.Address, backend bind.ContractBackend) (*StandardMT, error) {
	contract, err := bindStandardMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardMT{StandardMTCaller: StandardMTCaller{contract: contract}, StandardMTTransactor: StandardMTTransactor{contract: contract}, StandardMTFilterer: StandardMTFilterer{contract: contract}}, nil
}

// NewStandardMTCaller creates a new read-only instance of StandardMT, bound to a specific deployed contract.
func NewStandardMTCaller(address common.Address, caller bind.ContractCaller) (*StandardMTCaller, error) {
	contract, err := bindStandardMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardMTCaller{contract: contract}, nil
}

// NewStandardMTTransactor creates a new write-only instance of StandardMT, bound to a specific deployed contract.
func NewStandardMTTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardMTTransactor, error) {
	contract, err := bindStandardMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardMTTransactor{contract: contract}, nil
}

// NewStandardMTFilterer creates a new log filterer instance of StandardMT, bound to a specific deployed contract.
func NewStandardMTFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardMTFilterer, error) {
	contract, err := bindStandardMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardMTFilterer{contract: contract}, nil
}

// bindStandardMT binds a generic wrapper to an already deployed contract.
func bindStandardMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StandardMTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardMT *StandardMTRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardMT.Contract.StandardMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardMT *StandardMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardMT.Contract.StandardMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardMT *StandardMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardMT.Contract.StandardMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardMT *StandardMTCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardMT *StandardMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardMT *StandardMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardMT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_StandardMT *StandardMTCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardMT.contract.Call(opts, out, "balanceOf", account, id)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_StandardMT *StandardMTSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _StandardMT.Contract.BalanceOf(&_StandardMT.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_StandardMT *StandardMTCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _StandardMT.Contract.BalanceOf(&_StandardMT.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_StandardMT *StandardMTCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _StandardMT.contract.Call(opts, out, "balanceOfBatch", accounts, ids)
	return *ret0, err
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_StandardMT *StandardMTSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _StandardMT.Contract.BalanceOfBatch(&_StandardMT.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_StandardMT *StandardMTCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _StandardMT.Contract.BalanceOfBatch(&_StandardMT.CallOpts, accounts, ids)
}

// GetTokenUri is a free data retrieval call binding the contract method 0x8ad91345.
//
// Solidity: function getTokenUri(uint256 id) view returns(string)
func (_StandardMT *StandardMTCaller) GetTokenUri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StandardMT.contract.Call(opts, out, "getTokenUri", id)
	return *ret0, err
}

// GetTokenUri is a free data retrieval call binding the contract method 0x8ad91345.
//
// Solidity: function getTokenUri(uint256 id) view returns(string)
func (_StandardMT *StandardMTSession) GetTokenUri(id *big.Int) (string, error) {
	return _StandardMT.Contract.GetTokenUri(&_StandardMT.CallOpts, id)
}

// GetTokenUri is a free data retrieval call binding the contract method 0x8ad91345.
//
// Solidity: function getTokenUri(uint256 id) view returns(string)
func (_StandardMT *StandardMTCallerSession) GetTokenUri(id *big.Int) (string, error) {
	return _StandardMT.Contract.GetTokenUri(&_StandardMT.CallOpts, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_StandardMT *StandardMTCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StandardMT.contract.Call(opts, out, "isApprovedForAll", account, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_StandardMT *StandardMTSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _StandardMT.Contract.IsApprovedForAll(&_StandardMT.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_StandardMT *StandardMTCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _StandardMT.Contract.IsApprovedForAll(&_StandardMT.CallOpts, account, operator)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_StandardMT *StandardMTCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StandardMT.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_StandardMT *StandardMTSession) Ready() (bool, error) {
	return _StandardMT.Contract.Ready(&_StandardMT.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_StandardMT *StandardMTCallerSession) Ready() (bool, error) {
	return _StandardMT.Contract.Ready(&_StandardMT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StandardMT *StandardMTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StandardMT.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StandardMT *StandardMTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StandardMT.Contract.SupportsInterface(&_StandardMT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StandardMT *StandardMTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StandardMT.Contract.SupportsInterface(&_StandardMT.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_StandardMT *StandardMTCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StandardMT.contract.Call(opts, out, "uri", tokenId)
	return *ret0, err
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_StandardMT *StandardMTSession) Uri(tokenId *big.Int) (string, error) {
	return _StandardMT.Contract.Uri(&_StandardMT.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_StandardMT *StandardMTCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _StandardMT.Contract.Uri(&_StandardMT.CallOpts, tokenId)
}

// ChangeTokenUri is a paid mutator transaction binding the contract method 0xa20e08ca.
//
// Solidity: function changeTokenUri(string uri) returns()
func (_StandardMT *StandardMTTransactor) ChangeTokenUri(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _StandardMT.contract.Transact(opts, "changeTokenUri", uri)
}

// ChangeTokenUri is a paid mutator transaction binding the contract method 0xa20e08ca.
//
// Solidity: function changeTokenUri(string uri) returns()
func (_StandardMT *StandardMTSession) ChangeTokenUri(uri string) (*types.Transaction, error) {
	return _StandardMT.Contract.ChangeTokenUri(&_StandardMT.TransactOpts, uri)
}

// ChangeTokenUri is a paid mutator transaction binding the contract method 0xa20e08ca.
//
// Solidity: function changeTokenUri(string uri) returns()
func (_StandardMT *StandardMTTransactorSession) ChangeTokenUri(uri string) (*types.Transaction, error) {
	return _StandardMT.Contract.ChangeTokenUri(&_StandardMT.TransactOpts, uri)
}

// Initialize is a paid mutator transaction binding the contract method 0x76e0a647.
//
// Solidity: function initialize(string u, uint256 ptokenId, string pi, uint256 vtokenId, string vi, uint256 ctokenId, string ci, address operator, address ) returns()
func (_StandardMT *StandardMTTransactor) Initialize(opts *bind.TransactOpts, u string, ptokenId *big.Int, pi string, vtokenId *big.Int, vi string, ctokenId *big.Int, ci string, operator common.Address, arg8 common.Address) (*types.Transaction, error) {
	return _StandardMT.contract.Transact(opts, "initialize", u, ptokenId, pi, vtokenId, vi, ctokenId, ci, operator, arg8)
}

// Initialize is a paid mutator transaction binding the contract method 0x76e0a647.
//
// Solidity: function initialize(string u, uint256 ptokenId, string pi, uint256 vtokenId, string vi, uint256 ctokenId, string ci, address operator, address ) returns()
func (_StandardMT *StandardMTSession) Initialize(u string, ptokenId *big.Int, pi string, vtokenId *big.Int, vi string, ctokenId *big.Int, ci string, operator common.Address, arg8 common.Address) (*types.Transaction, error) {
	return _StandardMT.Contract.Initialize(&_StandardMT.TransactOpts, u, ptokenId, pi, vtokenId, vi, ctokenId, ci, operator, arg8)
}

// Initialize is a paid mutator transaction binding the contract method 0x76e0a647.
//
// Solidity: function initialize(string u, uint256 ptokenId, string pi, uint256 vtokenId, string vi, uint256 ctokenId, string ci, address operator, address ) returns()
func (_StandardMT *StandardMTTransactorSession) Initialize(u string, ptokenId *big.Int, pi string, vtokenId *big.Int, vi string, ctokenId *big.Int, ci string, operator common.Address, arg8 common.Address) (*types.Transaction, error) {
	return _StandardMT.Contract.Initialize(&_StandardMT.TransactOpts, u, ptokenId, pi, vtokenId, vi, ctokenId, ci, operator, arg8)
}

// MigrateOperator is a paid mutator transaction binding the contract method 0x6379808f.
//
// Solidity: function migrateOperator(address newOperator) returns()
func (_StandardMT *StandardMTTransactor) MigrateOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _StandardMT.contract.Transact(opts, "migrateOperator", newOperator)
}

// MigrateOperator is a paid mutator transaction binding the contract method 0x6379808f.
//
// Solidity: function migrateOperator(address newOperator) returns()
func (_StandardMT *StandardMTSession) MigrateOperator(newOperator common.Address) (*types.Transaction, error) {
	return _StandardMT.Contract.MigrateOperator(&_StandardMT.TransactOpts, newOperator)
}

// MigrateOperator is a paid mutator transaction binding the contract method 0x6379808f.
//
// Solidity: function migrateOperator(address newOperator) returns()
func (_StandardMT *StandardMTTransactorSession) MigrateOperator(newOperator common.Address) (*types.Transaction, error) {
	return _StandardMT.Contract.MigrateOperator(&_StandardMT.TransactOpts, newOperator)
}

// MintSBT is a paid mutator transaction binding the contract method 0x8883915b.
//
// Solidity: function mintSBT(address user, uint256 tokenId, uint256 quantity, bytes data) returns()
func (_StandardMT *StandardMTTransactor) MintSBT(opts *bind.TransactOpts, user common.Address, tokenId *big.Int, quantity *big.Int, data []byte) (*types.Transaction, error) {
	return _StandardMT.contract.Transact(opts, "mintSBT", user, tokenId, quantity, data)
}

// MintSBT is a paid mutator transaction binding the contract method 0x8883915b.
//
// Solidity: function mintSBT(address user, uint256 tokenId, uint256 quantity, bytes data) returns()
func (_StandardMT *StandardMTSession) MintSBT(user common.Address, tokenId *big.Int, quantity *big.Int, data []byte) (*types.Transaction, error) {
	return _StandardMT.Contract.MintSBT(&_StandardMT.TransactOpts, user, tokenId, quantity, data)
}

// MintSBT is a paid mutator transaction binding the contract method 0x8883915b.
//
// Solidity: function mintSBT(address user, uint256 tokenId, uint256 quantity, bytes data) returns()
func (_StandardMT *StandardMTTransactorSession) MintSBT(user common.Address, tokenId *big.Int, quantity *big.Int, data []byte) (*types.Transaction, error) {
	return _StandardMT.Contract.MintSBT(&_StandardMT.TransactOpts, user, tokenId, quantity, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_StandardMT *StandardMTTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _StandardMT.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_StandardMT *StandardMTSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _StandardMT.Contract.SafeBatchTransferFrom(&_StandardMT.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_StandardMT *StandardMTTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _StandardMT.Contract.SafeBatchTransferFrom(&_StandardMT.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_StandardMT *StandardMTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _StandardMT.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_StandardMT *StandardMTSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _StandardMT.Contract.SafeTransferFrom(&_StandardMT.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_StandardMT *StandardMTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _StandardMT.Contract.SafeTransferFrom(&_StandardMT.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StandardMT *StandardMTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _StandardMT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StandardMT *StandardMTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StandardMT.Contract.SetApprovalForAll(&_StandardMT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StandardMT *StandardMTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StandardMT.Contract.SetApprovalForAll(&_StandardMT.TransactOpts, operator, approved)
}

// StandardMTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StandardMT contract.
type StandardMTApprovalForAllIterator struct {
	Event *StandardMTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StandardMTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardMTApprovalForAll)
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
		it.Event = new(StandardMTApprovalForAll)
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
func (it *StandardMTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardMTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardMTApprovalForAll represents a ApprovalForAll event raised by the StandardMT contract.
type StandardMTApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_StandardMT *StandardMTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*StandardMTApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StandardMT.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StandardMTApprovalForAllIterator{contract: _StandardMT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_StandardMT *StandardMTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StandardMTApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StandardMT.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardMTApprovalForAll)
				if err := _StandardMT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_StandardMT *StandardMTFilterer) ParseApprovalForAll(log types.Log) (*StandardMTApprovalForAll, error) {
	event := new(StandardMTApprovalForAll)
	if err := _StandardMT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardMTTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the StandardMT contract.
type StandardMTTransferBatchIterator struct {
	Event *StandardMTTransferBatch // Event containing the contract specifics and raw log

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
func (it *StandardMTTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardMTTransferBatch)
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
		it.Event = new(StandardMTTransferBatch)
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
func (it *StandardMTTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardMTTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardMTTransferBatch represents a TransferBatch event raised by the StandardMT contract.
type StandardMTTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_StandardMT *StandardMTFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*StandardMTTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardMT.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardMTTransferBatchIterator{contract: _StandardMT.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_StandardMT *StandardMTFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *StandardMTTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardMT.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardMTTransferBatch)
				if err := _StandardMT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_StandardMT *StandardMTFilterer) ParseTransferBatch(log types.Log) (*StandardMTTransferBatch, error) {
	event := new(StandardMTTransferBatch)
	if err := _StandardMT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardMTTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the StandardMT contract.
type StandardMTTransferSingleIterator struct {
	Event *StandardMTTransferSingle // Event containing the contract specifics and raw log

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
func (it *StandardMTTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardMTTransferSingle)
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
		it.Event = new(StandardMTTransferSingle)
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
func (it *StandardMTTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardMTTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardMTTransferSingle represents a TransferSingle event raised by the StandardMT contract.
type StandardMTTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_StandardMT *StandardMTFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*StandardMTTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardMT.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardMTTransferSingleIterator{contract: _StandardMT.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_StandardMT *StandardMTFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *StandardMTTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardMT.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardMTTransferSingle)
				if err := _StandardMT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_StandardMT *StandardMTFilterer) ParseTransferSingle(log types.Log) (*StandardMTTransferSingle, error) {
	event := new(StandardMTTransferSingle)
	if err := _StandardMT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardMTURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the StandardMT contract.
type StandardMTURIIterator struct {
	Event *StandardMTURI // Event containing the contract specifics and raw log

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
func (it *StandardMTURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardMTURI)
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
		it.Event = new(StandardMTURI)
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
func (it *StandardMTURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardMTURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardMTURI represents a URI event raised by the StandardMT contract.
type StandardMTURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_StandardMT *StandardMTFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*StandardMTURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StandardMT.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &StandardMTURIIterator{contract: _StandardMT.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_StandardMT *StandardMTFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *StandardMTURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StandardMT.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardMTURI)
				if err := _StandardMT.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_StandardMT *StandardMTFilterer) ParseURI(log types.Log) (*StandardMTURI, error) {
	event := new(StandardMTURI)
	if err := _StandardMT.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardMTChangeTokenUriEventIterator is returned from FilterChangeTokenUriEvent and is used to iterate over the raw logs and unpacked data for ChangeTokenUriEvent events raised by the StandardMT contract.
type StandardMTChangeTokenUriEventIterator struct {
	Event *StandardMTChangeTokenUriEvent // Event containing the contract specifics and raw log

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
func (it *StandardMTChangeTokenUriEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardMTChangeTokenUriEvent)
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
		it.Event = new(StandardMTChangeTokenUriEvent)
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
func (it *StandardMTChangeTokenUriEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardMTChangeTokenUriEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardMTChangeTokenUriEvent represents a ChangeTokenUriEvent event raised by the StandardMT contract.
type StandardMTChangeTokenUriEvent struct {
	Uri string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChangeTokenUriEvent is a free log retrieval operation binding the contract event 0x4b451ab0a02e4ff0b97c018ededfd0cf4919376795b77640cf39815dabcb69b5.
//
// Solidity: event changeTokenUriEvent(string uri)
func (_StandardMT *StandardMTFilterer) FilterChangeTokenUriEvent(opts *bind.FilterOpts) (*StandardMTChangeTokenUriEventIterator, error) {

	logs, sub, err := _StandardMT.contract.FilterLogs(opts, "changeTokenUriEvent")
	if err != nil {
		return nil, err
	}
	return &StandardMTChangeTokenUriEventIterator{contract: _StandardMT.contract, event: "changeTokenUriEvent", logs: logs, sub: sub}, nil
}

// WatchChangeTokenUriEvent is a free log subscription operation binding the contract event 0x4b451ab0a02e4ff0b97c018ededfd0cf4919376795b77640cf39815dabcb69b5.
//
// Solidity: event changeTokenUriEvent(string uri)
func (_StandardMT *StandardMTFilterer) WatchChangeTokenUriEvent(opts *bind.WatchOpts, sink chan<- *StandardMTChangeTokenUriEvent) (event.Subscription, error) {

	logs, sub, err := _StandardMT.contract.WatchLogs(opts, "changeTokenUriEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardMTChangeTokenUriEvent)
				if err := _StandardMT.contract.UnpackLog(event, "changeTokenUriEvent", log); err != nil {
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

// ParseChangeTokenUriEvent is a log parse operation binding the contract event 0x4b451ab0a02e4ff0b97c018ededfd0cf4919376795b77640cf39815dabcb69b5.
//
// Solidity: event changeTokenUriEvent(string uri)
func (_StandardMT *StandardMTFilterer) ParseChangeTokenUriEvent(log types.Log) (*StandardMTChangeTokenUriEvent, error) {
	event := new(StandardMTChangeTokenUriEvent)
	if err := _StandardMT.contract.UnpackLog(event, "changeTokenUriEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardMTInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the StandardMT contract.
type StandardMTInitializeEventIterator struct {
	Event *StandardMTInitializeEvent // Event containing the contract specifics and raw log

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
func (it *StandardMTInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardMTInitializeEvent)
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
		it.Event = new(StandardMTInitializeEvent)
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
func (it *StandardMTInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardMTInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardMTInitializeEvent represents a InitializeEvent event raised by the StandardMT contract.
type StandardMTInitializeEvent struct {
	U        string
	PtokenId *big.Int
	Pi       string
	VtokenId *big.Int
	Vi       string
	CtokenId *big.Int
	Ci       string
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0xfddac5803a189a8f4c226b42e477d1b20d32a1bf825dbce945d34a9cbe276316.
//
// Solidity: event initializeEvent(string u, uint256 ptokenId, string pi, uint256 vtokenId, string vi, uint256 ctokenId, string ci, address operator)
func (_StandardMT *StandardMTFilterer) FilterInitializeEvent(opts *bind.FilterOpts) (*StandardMTInitializeEventIterator, error) {

	logs, sub, err := _StandardMT.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &StandardMTInitializeEventIterator{contract: _StandardMT.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0xfddac5803a189a8f4c226b42e477d1b20d32a1bf825dbce945d34a9cbe276316.
//
// Solidity: event initializeEvent(string u, uint256 ptokenId, string pi, uint256 vtokenId, string vi, uint256 ctokenId, string ci, address operator)
func (_StandardMT *StandardMTFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *StandardMTInitializeEvent) (event.Subscription, error) {

	logs, sub, err := _StandardMT.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardMTInitializeEvent)
				if err := _StandardMT.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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

// ParseInitializeEvent is a log parse operation binding the contract event 0xfddac5803a189a8f4c226b42e477d1b20d32a1bf825dbce945d34a9cbe276316.
//
// Solidity: event initializeEvent(string u, uint256 ptokenId, string pi, uint256 vtokenId, string vi, uint256 ctokenId, string ci, address operator)
func (_StandardMT *StandardMTFilterer) ParseInitializeEvent(log types.Log) (*StandardMTInitializeEvent, error) {
	event := new(StandardMTInitializeEvent)
	if err := _StandardMT.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardMTMigrateOperatorEventIterator is returned from FilterMigrateOperatorEvent and is used to iterate over the raw logs and unpacked data for MigrateOperatorEvent events raised by the StandardMT contract.
type StandardMTMigrateOperatorEventIterator struct {
	Event *StandardMTMigrateOperatorEvent // Event containing the contract specifics and raw log

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
func (it *StandardMTMigrateOperatorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardMTMigrateOperatorEvent)
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
		it.Event = new(StandardMTMigrateOperatorEvent)
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
func (it *StandardMTMigrateOperatorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardMTMigrateOperatorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardMTMigrateOperatorEvent represents a MigrateOperatorEvent event raised by the StandardMT contract.
type StandardMTMigrateOperatorEvent struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMigrateOperatorEvent is a free log retrieval operation binding the contract event 0x099d2c6de184bc25866ad9bf894425f5af26d7970ba0154ab19feddaae24a78c.
//
// Solidity: event migrateOperatorEvent(address newOperator)
func (_StandardMT *StandardMTFilterer) FilterMigrateOperatorEvent(opts *bind.FilterOpts) (*StandardMTMigrateOperatorEventIterator, error) {

	logs, sub, err := _StandardMT.contract.FilterLogs(opts, "migrateOperatorEvent")
	if err != nil {
		return nil, err
	}
	return &StandardMTMigrateOperatorEventIterator{contract: _StandardMT.contract, event: "migrateOperatorEvent", logs: logs, sub: sub}, nil
}

// WatchMigrateOperatorEvent is a free log subscription operation binding the contract event 0x099d2c6de184bc25866ad9bf894425f5af26d7970ba0154ab19feddaae24a78c.
//
// Solidity: event migrateOperatorEvent(address newOperator)
func (_StandardMT *StandardMTFilterer) WatchMigrateOperatorEvent(opts *bind.WatchOpts, sink chan<- *StandardMTMigrateOperatorEvent) (event.Subscription, error) {

	logs, sub, err := _StandardMT.contract.WatchLogs(opts, "migrateOperatorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardMTMigrateOperatorEvent)
				if err := _StandardMT.contract.UnpackLog(event, "migrateOperatorEvent", log); err != nil {
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
func (_StandardMT *StandardMTFilterer) ParseMigrateOperatorEvent(log types.Log) (*StandardMTMigrateOperatorEvent, error) {
	event := new(StandardMTMigrateOperatorEvent)
	if err := _StandardMT.contract.UnpackLog(event, "migrateOperatorEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardMTMintSBTEventIterator is returned from FilterMintSBTEvent and is used to iterate over the raw logs and unpacked data for MintSBTEvent events raised by the StandardMT contract.
type StandardMTMintSBTEventIterator struct {
	Event *StandardMTMintSBTEvent // Event containing the contract specifics and raw log

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
func (it *StandardMTMintSBTEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardMTMintSBTEvent)
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
		it.Event = new(StandardMTMintSBTEvent)
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
func (it *StandardMTMintSBTEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardMTMintSBTEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardMTMintSBTEvent represents a MintSBTEvent event raised by the StandardMT contract.
type StandardMTMintSBTEvent struct {
	User     common.Address
	TokenId  *big.Int
	Quantity *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintSBTEvent is a free log retrieval operation binding the contract event 0x4f02a77788758705ce2cfd3292867693f24689b5f615dbe4f4499026e4b09cfb.
//
// Solidity: event mintSBTEvent(address user, uint256 tokenId, uint256 quantity, bytes data)
func (_StandardMT *StandardMTFilterer) FilterMintSBTEvent(opts *bind.FilterOpts) (*StandardMTMintSBTEventIterator, error) {

	logs, sub, err := _StandardMT.contract.FilterLogs(opts, "mintSBTEvent")
	if err != nil {
		return nil, err
	}
	return &StandardMTMintSBTEventIterator{contract: _StandardMT.contract, event: "mintSBTEvent", logs: logs, sub: sub}, nil
}

// WatchMintSBTEvent is a free log subscription operation binding the contract event 0x4f02a77788758705ce2cfd3292867693f24689b5f615dbe4f4499026e4b09cfb.
//
// Solidity: event mintSBTEvent(address user, uint256 tokenId, uint256 quantity, bytes data)
func (_StandardMT *StandardMTFilterer) WatchMintSBTEvent(opts *bind.WatchOpts, sink chan<- *StandardMTMintSBTEvent) (event.Subscription, error) {

	logs, sub, err := _StandardMT.contract.WatchLogs(opts, "mintSBTEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardMTMintSBTEvent)
				if err := _StandardMT.contract.UnpackLog(event, "mintSBTEvent", log); err != nil {
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

// ParseMintSBTEvent is a log parse operation binding the contract event 0x4f02a77788758705ce2cfd3292867693f24689b5f615dbe4f4499026e4b09cfb.
//
// Solidity: event mintSBTEvent(address user, uint256 tokenId, uint256 quantity, bytes data)
func (_StandardMT *StandardMTFilterer) ParseMintSBTEvent(log types.Log) (*StandardMTMintSBTEvent, error) {
	event := new(StandardMTMintSBTEvent)
	if err := _StandardMT.contract.UnpackLog(event, "mintSBTEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardMTSetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the StandardMT contract.
type StandardMTSetReadyEventIterator struct {
	Event *StandardMTSetReadyEvent // Event containing the contract specifics and raw log

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
func (it *StandardMTSetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardMTSetReadyEvent)
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
		it.Event = new(StandardMTSetReadyEvent)
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
func (it *StandardMTSetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardMTSetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardMTSetReadyEvent represents a SetReadyEvent event raised by the StandardMT contract.
type StandardMTSetReadyEvent struct {
	R   bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool r)
func (_StandardMT *StandardMTFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*StandardMTSetReadyEventIterator, error) {

	logs, sub, err := _StandardMT.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &StandardMTSetReadyEventIterator{contract: _StandardMT.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool r)
func (_StandardMT *StandardMTFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *StandardMTSetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _StandardMT.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardMTSetReadyEvent)
				if err := _StandardMT.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
func (_StandardMT *StandardMTFilterer) ParseSetReadyEvent(log types.Log) (*StandardMTSetReadyEvent, error) {
	event := new(StandardMTSetReadyEvent)
	if err := _StandardMT.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
