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

// BiyardNFTMetaData contains all meta data concerning the BiyardNFT contract.
var BiyardNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"whitelist\",\"type\":\"string[]\"}],\"name\":\"AddToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"whitelisted\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"SetClaimLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start_token_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"last_token_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useWhitelist\",\"type\":\"bool\"}],\"name\":\"SetMintInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"partner\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"}],\"name\":\"SetPartnerQuota\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"when\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"partner\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addPartnerWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addTrustOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"allowedOf\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"delTrustedParty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"}],\"name\":\"executeCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"historyOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"holders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isTrusted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listQuotas\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listTrustOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listTurstedParties\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"listWhitelistAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfCandidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeTrustOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"setAllowDuplicatedClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setClaimLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"belong\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"}],\"name\":\"setPartnerQuota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"belongs\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quotas\",\"type\":\"uint256[]\"}],\"name\":\"setPartnerQuotas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTrustedParty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useWhitelist\",\"type\":\"bool\"}],\"name\":\"setupMintInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BiyardNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use BiyardNFTMetaData.ABI instead.
var BiyardNFTABI = BiyardNFTMetaData.ABI

// BiyardNFTBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BiyardNFTBinRuntime = ``

// BiyardNFT is an auto generated Go binding around a Klaytn contract.
type BiyardNFT struct {
	BiyardNFTCaller     // Read-only binding to the contract
	BiyardNFTTransactor // Write-only binding to the contract
	BiyardNFTFilterer   // Log filterer for contract events
}

// BiyardNFTCaller is an auto generated read-only Go binding around a Klaytn contract.
type BiyardNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiyardNFTTransactor is an auto generated write-only Go binding around a Klaytn contract.
type BiyardNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiyardNFTFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type BiyardNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiyardNFTSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type BiyardNFTSession struct {
	Contract     *BiyardNFT        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiyardNFTCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type BiyardNFTCallerSession struct {
	Contract *BiyardNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BiyardNFTTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type BiyardNFTTransactorSession struct {
	Contract     *BiyardNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BiyardNFTRaw is an auto generated low-level Go binding around a Klaytn contract.
type BiyardNFTRaw struct {
	Contract *BiyardNFT // Generic contract binding to access the raw methods on
}

// BiyardNFTCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type BiyardNFTCallerRaw struct {
	Contract *BiyardNFTCaller // Generic read-only contract binding to access the raw methods on
}

// BiyardNFTTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type BiyardNFTTransactorRaw struct {
	Contract *BiyardNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBiyardNFT creates a new instance of BiyardNFT, bound to a specific deployed contract.
func NewBiyardNFT(address common.Address, backend bind.ContractBackend) (*BiyardNFT, error) {
	contract, err := bindBiyardNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BiyardNFT{BiyardNFTCaller: BiyardNFTCaller{contract: contract}, BiyardNFTTransactor: BiyardNFTTransactor{contract: contract}, BiyardNFTFilterer: BiyardNFTFilterer{contract: contract}}, nil
}

// NewBiyardNFTCaller creates a new read-only instance of BiyardNFT, bound to a specific deployed contract.
func NewBiyardNFTCaller(address common.Address, caller bind.ContractCaller) (*BiyardNFTCaller, error) {
	contract, err := bindBiyardNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BiyardNFTCaller{contract: contract}, nil
}

// NewBiyardNFTTransactor creates a new write-only instance of BiyardNFT, bound to a specific deployed contract.
func NewBiyardNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*BiyardNFTTransactor, error) {
	contract, err := bindBiyardNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BiyardNFTTransactor{contract: contract}, nil
}

// NewBiyardNFTFilterer creates a new log filterer instance of BiyardNFT, bound to a specific deployed contract.
func NewBiyardNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*BiyardNFTFilterer, error) {
	contract, err := bindBiyardNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BiyardNFTFilterer{contract: contract}, nil
}

// bindBiyardNFT binds a generic wrapper to an already deployed contract.
func bindBiyardNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BiyardNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BiyardNFT *BiyardNFTRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BiyardNFT.Contract.BiyardNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BiyardNFT *BiyardNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiyardNFT.Contract.BiyardNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BiyardNFT *BiyardNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BiyardNFT.Contract.BiyardNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BiyardNFT *BiyardNFTCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BiyardNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BiyardNFT *BiyardNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiyardNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BiyardNFT *BiyardNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BiyardNFT.Contract.contract.Transact(opts, method, params...)
}

// AllowedOf is a free data retrieval call binding the contract method 0x5dbff683.
//
// Solidity: function allowedOf(address addr) view returns(string[])
func (_BiyardNFT *BiyardNFTCaller) AllowedOf(opts *bind.CallOpts, addr common.Address) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "allowedOf", addr)
	return *ret0, err
}

// AllowedOf is a free data retrieval call binding the contract method 0x5dbff683.
//
// Solidity: function allowedOf(address addr) view returns(string[])
func (_BiyardNFT *BiyardNFTSession) AllowedOf(addr common.Address) ([]string, error) {
	return _BiyardNFT.Contract.AllowedOf(&_BiyardNFT.CallOpts, addr)
}

// AllowedOf is a free data retrieval call binding the contract method 0x5dbff683.
//
// Solidity: function allowedOf(address addr) view returns(string[])
func (_BiyardNFT *BiyardNFTCallerSession) AllowedOf(addr common.Address) ([]string, error) {
	return _BiyardNFT.Contract.AllowedOf(&_BiyardNFT.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BiyardNFT *BiyardNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BiyardNFT *BiyardNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BiyardNFT.Contract.BalanceOf(&_BiyardNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BiyardNFT *BiyardNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BiyardNFT.Contract.BalanceOf(&_BiyardNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BiyardNFT *BiyardNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BiyardNFT *BiyardNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BiyardNFT.Contract.GetApproved(&_BiyardNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BiyardNFT *BiyardNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BiyardNFT.Contract.GetApproved(&_BiyardNFT.CallOpts, tokenId)
}

// HistoryOf is a free data retrieval call binding the contract method 0x3a11b65a.
//
// Solidity: function historyOf(uint256 tokenId) view returns(address[])
func (_BiyardNFT *BiyardNFTCaller) HistoryOf(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "historyOf", tokenId)
	return *ret0, err
}

// HistoryOf is a free data retrieval call binding the contract method 0x3a11b65a.
//
// Solidity: function historyOf(uint256 tokenId) view returns(address[])
func (_BiyardNFT *BiyardNFTSession) HistoryOf(tokenId *big.Int) ([]common.Address, error) {
	return _BiyardNFT.Contract.HistoryOf(&_BiyardNFT.CallOpts, tokenId)
}

// HistoryOf is a free data retrieval call binding the contract method 0x3a11b65a.
//
// Solidity: function historyOf(uint256 tokenId) view returns(address[])
func (_BiyardNFT *BiyardNFTCallerSession) HistoryOf(tokenId *big.Int) ([]common.Address, error) {
	return _BiyardNFT.Contract.HistoryOf(&_BiyardNFT.CallOpts, tokenId)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() view returns(address[])
func (_BiyardNFT *BiyardNFTCaller) Holders(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "holders")
	return *ret0, err
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() view returns(address[])
func (_BiyardNFT *BiyardNFTSession) Holders() ([]common.Address, error) {
	return _BiyardNFT.Contract.Holders(&_BiyardNFT.CallOpts)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() view returns(address[])
func (_BiyardNFT *BiyardNFTCallerSession) Holders() ([]common.Address, error) {
	return _BiyardNFT.Contract.Holders(&_BiyardNFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BiyardNFT *BiyardNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BiyardNFT *BiyardNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BiyardNFT.Contract.IsApprovedForAll(&_BiyardNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BiyardNFT *BiyardNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BiyardNFT.Contract.IsApprovedForAll(&_BiyardNFT.CallOpts, owner, operator)
}

// IsTrusted is a free data retrieval call binding the contract method 0x96d64879.
//
// Solidity: function isTrusted(address addr) view returns(bool)
func (_BiyardNFT *BiyardNFTCaller) IsTrusted(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "isTrusted", addr)
	return *ret0, err
}

// IsTrusted is a free data retrieval call binding the contract method 0x96d64879.
//
// Solidity: function isTrusted(address addr) view returns(bool)
func (_BiyardNFT *BiyardNFTSession) IsTrusted(addr common.Address) (bool, error) {
	return _BiyardNFT.Contract.IsTrusted(&_BiyardNFT.CallOpts, addr)
}

// IsTrusted is a free data retrieval call binding the contract method 0x96d64879.
//
// Solidity: function isTrusted(address addr) view returns(bool)
func (_BiyardNFT *BiyardNFTCallerSession) IsTrusted(addr common.Address) (bool, error) {
	return _BiyardNFT.Contract.IsTrusted(&_BiyardNFT.CallOpts, addr)
}

// ListQuotas is a free data retrieval call binding the contract method 0xc8938403.
//
// Solidity: function listQuotas() view returns(string[], uint256[])
func (_BiyardNFT *BiyardNFTCaller) ListQuotas(opts *bind.CallOpts) ([]string, []*big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BiyardNFT.contract.Call(opts, out, "listQuotas")
	return *ret0, *ret1, err
}

// ListQuotas is a free data retrieval call binding the contract method 0xc8938403.
//
// Solidity: function listQuotas() view returns(string[], uint256[])
func (_BiyardNFT *BiyardNFTSession) ListQuotas() ([]string, []*big.Int, error) {
	return _BiyardNFT.Contract.ListQuotas(&_BiyardNFT.CallOpts)
}

// ListQuotas is a free data retrieval call binding the contract method 0xc8938403.
//
// Solidity: function listQuotas() view returns(string[], uint256[])
func (_BiyardNFT *BiyardNFTCallerSession) ListQuotas() ([]string, []*big.Int, error) {
	return _BiyardNFT.Contract.ListQuotas(&_BiyardNFT.CallOpts)
}

// ListTrustOperators is a free data retrieval call binding the contract method 0x3fa1e1c4.
//
// Solidity: function listTrustOperators() view returns(address[])
func (_BiyardNFT *BiyardNFTCaller) ListTrustOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "listTrustOperators")
	return *ret0, err
}

// ListTrustOperators is a free data retrieval call binding the contract method 0x3fa1e1c4.
//
// Solidity: function listTrustOperators() view returns(address[])
func (_BiyardNFT *BiyardNFTSession) ListTrustOperators() ([]common.Address, error) {
	return _BiyardNFT.Contract.ListTrustOperators(&_BiyardNFT.CallOpts)
}

// ListTrustOperators is a free data retrieval call binding the contract method 0x3fa1e1c4.
//
// Solidity: function listTrustOperators() view returns(address[])
func (_BiyardNFT *BiyardNFTCallerSession) ListTrustOperators() ([]common.Address, error) {
	return _BiyardNFT.Contract.ListTrustOperators(&_BiyardNFT.CallOpts)
}

// ListTurstedParties is a free data retrieval call binding the contract method 0xd6652b75.
//
// Solidity: function listTurstedParties() view returns(address[])
func (_BiyardNFT *BiyardNFTCaller) ListTurstedParties(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "listTurstedParties")
	return *ret0, err
}

// ListTurstedParties is a free data retrieval call binding the contract method 0xd6652b75.
//
// Solidity: function listTurstedParties() view returns(address[])
func (_BiyardNFT *BiyardNFTSession) ListTurstedParties() ([]common.Address, error) {
	return _BiyardNFT.Contract.ListTurstedParties(&_BiyardNFT.CallOpts)
}

// ListTurstedParties is a free data retrieval call binding the contract method 0xd6652b75.
//
// Solidity: function listTurstedParties() view returns(address[])
func (_BiyardNFT *BiyardNFTCallerSession) ListTurstedParties() ([]common.Address, error) {
	return _BiyardNFT.Contract.ListTurstedParties(&_BiyardNFT.CallOpts)
}

// ListWhitelistAddresses is a free data retrieval call binding the contract method 0xb79dfdbd.
//
// Solidity: function listWhitelistAddresses(uint256 i) view returns(address[])
func (_BiyardNFT *BiyardNFTCaller) ListWhitelistAddresses(opts *bind.CallOpts, i *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "listWhitelistAddresses", i)
	return *ret0, err
}

// ListWhitelistAddresses is a free data retrieval call binding the contract method 0xb79dfdbd.
//
// Solidity: function listWhitelistAddresses(uint256 i) view returns(address[])
func (_BiyardNFT *BiyardNFTSession) ListWhitelistAddresses(i *big.Int) ([]common.Address, error) {
	return _BiyardNFT.Contract.ListWhitelistAddresses(&_BiyardNFT.CallOpts, i)
}

// ListWhitelistAddresses is a free data retrieval call binding the contract method 0xb79dfdbd.
//
// Solidity: function listWhitelistAddresses(uint256 i) view returns(address[])
func (_BiyardNFT *BiyardNFTCallerSession) ListWhitelistAddresses(i *big.Int) ([]common.Address, error) {
	return _BiyardNFT.Contract.ListWhitelistAddresses(&_BiyardNFT.CallOpts, i)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BiyardNFT *BiyardNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BiyardNFT *BiyardNFTSession) Name() (string, error) {
	return _BiyardNFT.Contract.Name(&_BiyardNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BiyardNFT *BiyardNFTCallerSession) Name() (string, error) {
	return _BiyardNFT.Contract.Name(&_BiyardNFT.CallOpts)
}

// NumberOfCandidates is a free data retrieval call binding the contract method 0x9c84bb27.
//
// Solidity: function numberOfCandidates() view returns(uint256)
func (_BiyardNFT *BiyardNFTCaller) NumberOfCandidates(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "numberOfCandidates")
	return *ret0, err
}

// NumberOfCandidates is a free data retrieval call binding the contract method 0x9c84bb27.
//
// Solidity: function numberOfCandidates() view returns(uint256)
func (_BiyardNFT *BiyardNFTSession) NumberOfCandidates() (*big.Int, error) {
	return _BiyardNFT.Contract.NumberOfCandidates(&_BiyardNFT.CallOpts)
}

// NumberOfCandidates is a free data retrieval call binding the contract method 0x9c84bb27.
//
// Solidity: function numberOfCandidates() view returns(uint256)
func (_BiyardNFT *BiyardNFTCallerSession) NumberOfCandidates() (*big.Int, error) {
	return _BiyardNFT.Contract.NumberOfCandidates(&_BiyardNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BiyardNFT *BiyardNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BiyardNFT *BiyardNFTSession) Owner() (common.Address, error) {
	return _BiyardNFT.Contract.Owner(&_BiyardNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BiyardNFT *BiyardNFTCallerSession) Owner() (common.Address, error) {
	return _BiyardNFT.Contract.Owner(&_BiyardNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BiyardNFT *BiyardNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BiyardNFT *BiyardNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BiyardNFT.Contract.OwnerOf(&_BiyardNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BiyardNFT *BiyardNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BiyardNFT.Contract.OwnerOf(&_BiyardNFT.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BiyardNFT *BiyardNFTCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BiyardNFT *BiyardNFTSession) Paused() (bool, error) {
	return _BiyardNFT.Contract.Paused(&_BiyardNFT.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BiyardNFT *BiyardNFTCallerSession) Paused() (bool, error) {
	return _BiyardNFT.Contract.Paused(&_BiyardNFT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BiyardNFT *BiyardNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BiyardNFT *BiyardNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BiyardNFT.Contract.SupportsInterface(&_BiyardNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BiyardNFT *BiyardNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BiyardNFT.Contract.SupportsInterface(&_BiyardNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BiyardNFT *BiyardNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BiyardNFT *BiyardNFTSession) Symbol() (string, error) {
	return _BiyardNFT.Contract.Symbol(&_BiyardNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BiyardNFT *BiyardNFTCallerSession) Symbol() (string, error) {
	return _BiyardNFT.Contract.Symbol(&_BiyardNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BiyardNFT *BiyardNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BiyardNFT *BiyardNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BiyardNFT.Contract.TokenURI(&_BiyardNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BiyardNFT *BiyardNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BiyardNFT.Contract.TokenURI(&_BiyardNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BiyardNFT *BiyardNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BiyardNFT.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BiyardNFT *BiyardNFTSession) TotalSupply() (*big.Int, error) {
	return _BiyardNFT.Contract.TotalSupply(&_BiyardNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BiyardNFT *BiyardNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _BiyardNFT.Contract.TotalSupply(&_BiyardNFT.CallOpts)
}

// AddPartnerWhitelist is a paid mutator transaction binding the contract method 0x2efed898.
//
// Solidity: function addPartnerWhitelist(string partner, address[] addrs) returns()
func (_BiyardNFT *BiyardNFTTransactor) AddPartnerWhitelist(opts *bind.TransactOpts, partner string, addrs []common.Address) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "addPartnerWhitelist", partner, addrs)
}

// AddPartnerWhitelist is a paid mutator transaction binding the contract method 0x2efed898.
//
// Solidity: function addPartnerWhitelist(string partner, address[] addrs) returns()
func (_BiyardNFT *BiyardNFTSession) AddPartnerWhitelist(partner string, addrs []common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.AddPartnerWhitelist(&_BiyardNFT.TransactOpts, partner, addrs)
}

// AddPartnerWhitelist is a paid mutator transaction binding the contract method 0x2efed898.
//
// Solidity: function addPartnerWhitelist(string partner, address[] addrs) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) AddPartnerWhitelist(partner string, addrs []common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.AddPartnerWhitelist(&_BiyardNFT.TransactOpts, partner, addrs)
}

// AddTrustOperator is a paid mutator transaction binding the contract method 0xe91eeb7a.
//
// Solidity: function addTrustOperator(address addr) returns()
func (_BiyardNFT *BiyardNFTTransactor) AddTrustOperator(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "addTrustOperator", addr)
}

// AddTrustOperator is a paid mutator transaction binding the contract method 0xe91eeb7a.
//
// Solidity: function addTrustOperator(address addr) returns()
func (_BiyardNFT *BiyardNFTSession) AddTrustOperator(addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.AddTrustOperator(&_BiyardNFT.TransactOpts, addr)
}

// AddTrustOperator is a paid mutator transaction binding the contract method 0xe91eeb7a.
//
// Solidity: function addTrustOperator(address addr) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) AddTrustOperator(addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.AddTrustOperator(&_BiyardNFT.TransactOpts, addr)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.Approve(&_BiyardNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.Approve(&_BiyardNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.Burn(&_BiyardNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.Burn(&_BiyardNFT.TransactOpts, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() payable returns()
func (_BiyardNFT *BiyardNFTTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() payable returns()
func (_BiyardNFT *BiyardNFTSession) Claim() (*types.Transaction, error) {
	return _BiyardNFT.Contract.Claim(&_BiyardNFT.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() payable returns()
func (_BiyardNFT *BiyardNFTTransactorSession) Claim() (*types.Transaction, error) {
	return _BiyardNFT.Contract.Claim(&_BiyardNFT.TransactOpts)
}

// DelTrustedParty is a paid mutator transaction binding the contract method 0xe3e486c5.
//
// Solidity: function delTrustedParty(address addr) returns()
func (_BiyardNFT *BiyardNFTTransactor) DelTrustedParty(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "delTrustedParty", addr)
}

// DelTrustedParty is a paid mutator transaction binding the contract method 0xe3e486c5.
//
// Solidity: function delTrustedParty(address addr) returns()
func (_BiyardNFT *BiyardNFTSession) DelTrustedParty(addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.DelTrustedParty(&_BiyardNFT.TransactOpts, addr)
}

// DelTrustedParty is a paid mutator transaction binding the contract method 0xe3e486c5.
//
// Solidity: function delTrustedParty(address addr) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) DelTrustedParty(addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.DelTrustedParty(&_BiyardNFT.TransactOpts, addr)
}

// ExecuteCode is a paid mutator transaction binding the contract method 0x3144c859.
//
// Solidity: function executeCode(bytes functionSignature) returns()
func (_BiyardNFT *BiyardNFTTransactor) ExecuteCode(opts *bind.TransactOpts, functionSignature []byte) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "executeCode", functionSignature)
}

// ExecuteCode is a paid mutator transaction binding the contract method 0x3144c859.
//
// Solidity: function executeCode(bytes functionSignature) returns()
func (_BiyardNFT *BiyardNFTSession) ExecuteCode(functionSignature []byte) (*types.Transaction, error) {
	return _BiyardNFT.Contract.ExecuteCode(&_BiyardNFT.TransactOpts, functionSignature)
}

// ExecuteCode is a paid mutator transaction binding the contract method 0x3144c859.
//
// Solidity: function executeCode(bytes functionSignature) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) ExecuteCode(functionSignature []byte) (*types.Transaction, error) {
	return _BiyardNFT.Contract.ExecuteCode(&_BiyardNFT.TransactOpts, functionSignature)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BiyardNFT *BiyardNFTTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BiyardNFT *BiyardNFTSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BiyardNFT.Contract.OnERC721Received(&_BiyardNFT.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BiyardNFT *BiyardNFTTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BiyardNFT.Contract.OnERC721Received(&_BiyardNFT.TransactOpts, arg0, arg1, arg2, arg3)
}

// RemoveTrustOperator is a paid mutator transaction binding the contract method 0x49adfdf5.
//
// Solidity: function removeTrustOperator(address addr) returns()
func (_BiyardNFT *BiyardNFTTransactor) RemoveTrustOperator(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "removeTrustOperator", addr)
}

// RemoveTrustOperator is a paid mutator transaction binding the contract method 0x49adfdf5.
//
// Solidity: function removeTrustOperator(address addr) returns()
func (_BiyardNFT *BiyardNFTSession) RemoveTrustOperator(addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.RemoveTrustOperator(&_BiyardNFT.TransactOpts, addr)
}

// RemoveTrustOperator is a paid mutator transaction binding the contract method 0x49adfdf5.
//
// Solidity: function removeTrustOperator(address addr) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) RemoveTrustOperator(addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.RemoveTrustOperator(&_BiyardNFT.TransactOpts, addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BiyardNFT *BiyardNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BiyardNFT *BiyardNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _BiyardNFT.Contract.RenounceOwnership(&_BiyardNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BiyardNFT *BiyardNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BiyardNFT.Contract.RenounceOwnership(&_BiyardNFT.TransactOpts)
}

// RequestWhitelist is a paid mutator transaction binding the contract method 0x6e1e9349.
//
// Solidity: function requestWhitelist() returns()
func (_BiyardNFT *BiyardNFTTransactor) RequestWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "requestWhitelist")
}

// RequestWhitelist is a paid mutator transaction binding the contract method 0x6e1e9349.
//
// Solidity: function requestWhitelist() returns()
func (_BiyardNFT *BiyardNFTSession) RequestWhitelist() (*types.Transaction, error) {
	return _BiyardNFT.Contract.RequestWhitelist(&_BiyardNFT.TransactOpts)
}

// RequestWhitelist is a paid mutator transaction binding the contract method 0x6e1e9349.
//
// Solidity: function requestWhitelist() returns()
func (_BiyardNFT *BiyardNFTTransactorSession) RequestWhitelist() (*types.Transaction, error) {
	return _BiyardNFT.Contract.RequestWhitelist(&_BiyardNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SafeTransferFrom(&_BiyardNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SafeTransferFrom(&_BiyardNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BiyardNFT *BiyardNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BiyardNFT *BiyardNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SafeTransferFrom0(&_BiyardNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SafeTransferFrom0(&_BiyardNFT.TransactOpts, from, to, tokenId, data)
}

// SetAllowDuplicatedClaim is a paid mutator transaction binding the contract method 0xd859dd6b.
//
// Solidity: function setAllowDuplicatedClaim(bool allow) returns()
func (_BiyardNFT *BiyardNFTTransactor) SetAllowDuplicatedClaim(opts *bind.TransactOpts, allow bool) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "setAllowDuplicatedClaim", allow)
}

// SetAllowDuplicatedClaim is a paid mutator transaction binding the contract method 0xd859dd6b.
//
// Solidity: function setAllowDuplicatedClaim(bool allow) returns()
func (_BiyardNFT *BiyardNFTSession) SetAllowDuplicatedClaim(allow bool) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetAllowDuplicatedClaim(&_BiyardNFT.TransactOpts, allow)
}

// SetAllowDuplicatedClaim is a paid mutator transaction binding the contract method 0xd859dd6b.
//
// Solidity: function setAllowDuplicatedClaim(bool allow) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SetAllowDuplicatedClaim(allow bool) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetAllowDuplicatedClaim(&_BiyardNFT.TransactOpts, allow)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BiyardNFT *BiyardNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BiyardNFT *BiyardNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetApprovalForAll(&_BiyardNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetApprovalForAll(&_BiyardNFT.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_BiyardNFT *BiyardNFTTransactor) SetBaseURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "setBaseURI", uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_BiyardNFT *BiyardNFTSession) SetBaseURI(uri string) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetBaseURI(&_BiyardNFT.TransactOpts, uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SetBaseURI(uri string) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetBaseURI(&_BiyardNFT.TransactOpts, uri)
}

// SetClaimLimit is a paid mutator transaction binding the contract method 0x020a0ff5.
//
// Solidity: function setClaimLimit(uint256 limit) returns()
func (_BiyardNFT *BiyardNFTTransactor) SetClaimLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "setClaimLimit", limit)
}

// SetClaimLimit is a paid mutator transaction binding the contract method 0x020a0ff5.
//
// Solidity: function setClaimLimit(uint256 limit) returns()
func (_BiyardNFT *BiyardNFTSession) SetClaimLimit(limit *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetClaimLimit(&_BiyardNFT.TransactOpts, limit)
}

// SetClaimLimit is a paid mutator transaction binding the contract method 0x020a0ff5.
//
// Solidity: function setClaimLimit(uint256 limit) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SetClaimLimit(limit *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetClaimLimit(&_BiyardNFT.TransactOpts, limit)
}

// SetPartnerQuota is a paid mutator transaction binding the contract method 0xaf93affd.
//
// Solidity: function setPartnerQuota(string belong, uint256 quota) returns()
func (_BiyardNFT *BiyardNFTTransactor) SetPartnerQuota(opts *bind.TransactOpts, belong string, quota *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "setPartnerQuota", belong, quota)
}

// SetPartnerQuota is a paid mutator transaction binding the contract method 0xaf93affd.
//
// Solidity: function setPartnerQuota(string belong, uint256 quota) returns()
func (_BiyardNFT *BiyardNFTSession) SetPartnerQuota(belong string, quota *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetPartnerQuota(&_BiyardNFT.TransactOpts, belong, quota)
}

// SetPartnerQuota is a paid mutator transaction binding the contract method 0xaf93affd.
//
// Solidity: function setPartnerQuota(string belong, uint256 quota) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SetPartnerQuota(belong string, quota *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetPartnerQuota(&_BiyardNFT.TransactOpts, belong, quota)
}

// SetPartnerQuotas is a paid mutator transaction binding the contract method 0xa73b5d7c.
//
// Solidity: function setPartnerQuotas(string[] belongs, uint256[] quotas) returns()
func (_BiyardNFT *BiyardNFTTransactor) SetPartnerQuotas(opts *bind.TransactOpts, belongs []string, quotas []*big.Int) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "setPartnerQuotas", belongs, quotas)
}

// SetPartnerQuotas is a paid mutator transaction binding the contract method 0xa73b5d7c.
//
// Solidity: function setPartnerQuotas(string[] belongs, uint256[] quotas) returns()
func (_BiyardNFT *BiyardNFTSession) SetPartnerQuotas(belongs []string, quotas []*big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetPartnerQuotas(&_BiyardNFT.TransactOpts, belongs, quotas)
}

// SetPartnerQuotas is a paid mutator transaction binding the contract method 0xa73b5d7c.
//
// Solidity: function setPartnerQuotas(string[] belongs, uint256[] quotas) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SetPartnerQuotas(belongs []string, quotas []*big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetPartnerQuotas(&_BiyardNFT.TransactOpts, belongs, quotas)
}

// SetTrustedParty is a paid mutator transaction binding the contract method 0x67b151e6.
//
// Solidity: function setTrustedParty(address addr) returns()
func (_BiyardNFT *BiyardNFTTransactor) SetTrustedParty(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "setTrustedParty", addr)
}

// SetTrustedParty is a paid mutator transaction binding the contract method 0x67b151e6.
//
// Solidity: function setTrustedParty(address addr) returns()
func (_BiyardNFT *BiyardNFTSession) SetTrustedParty(addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetTrustedParty(&_BiyardNFT.TransactOpts, addr)
}

// SetTrustedParty is a paid mutator transaction binding the contract method 0x67b151e6.
//
// Solidity: function setTrustedParty(address addr) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SetTrustedParty(addr common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetTrustedParty(&_BiyardNFT.TransactOpts, addr)
}

// SetupMintInfo is a paid mutator transaction binding the contract method 0xa92fd1f1.
//
// Solidity: function setupMintInfo(uint256 startTimestamp, uint256 startTokenId, uint256 lastTokenId, uint256 price, bool useWhitelist) returns()
func (_BiyardNFT *BiyardNFTTransactor) SetupMintInfo(opts *bind.TransactOpts, startTimestamp *big.Int, startTokenId *big.Int, lastTokenId *big.Int, price *big.Int, useWhitelist bool) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "setupMintInfo", startTimestamp, startTokenId, lastTokenId, price, useWhitelist)
}

// SetupMintInfo is a paid mutator transaction binding the contract method 0xa92fd1f1.
//
// Solidity: function setupMintInfo(uint256 startTimestamp, uint256 startTokenId, uint256 lastTokenId, uint256 price, bool useWhitelist) returns()
func (_BiyardNFT *BiyardNFTSession) SetupMintInfo(startTimestamp *big.Int, startTokenId *big.Int, lastTokenId *big.Int, price *big.Int, useWhitelist bool) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetupMintInfo(&_BiyardNFT.TransactOpts, startTimestamp, startTokenId, lastTokenId, price, useWhitelist)
}

// SetupMintInfo is a paid mutator transaction binding the contract method 0xa92fd1f1.
//
// Solidity: function setupMintInfo(uint256 startTimestamp, uint256 startTokenId, uint256 lastTokenId, uint256 price, bool useWhitelist) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) SetupMintInfo(startTimestamp *big.Int, startTokenId *big.Int, lastTokenId *big.Int, price *big.Int, useWhitelist bool) (*types.Transaction, error) {
	return _BiyardNFT.Contract.SetupMintInfo(&_BiyardNFT.TransactOpts, startTimestamp, startTokenId, lastTokenId, price, useWhitelist)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.TransferFrom(&_BiyardNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BiyardNFT.Contract.TransferFrom(&_BiyardNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BiyardNFT *BiyardNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BiyardNFT *BiyardNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.TransferOwnership(&_BiyardNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BiyardNFT *BiyardNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BiyardNFT.Contract.TransferOwnership(&_BiyardNFT.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BiyardNFT *BiyardNFTTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiyardNFT.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BiyardNFT *BiyardNFTSession) Withdraw() (*types.Transaction, error) {
	return _BiyardNFT.Contract.Withdraw(&_BiyardNFT.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BiyardNFT *BiyardNFTTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BiyardNFT.Contract.Withdraw(&_BiyardNFT.TransactOpts)
}

// BiyardNFTAddToWhitelistIterator is returned from FilterAddToWhitelist and is used to iterate over the raw logs and unpacked data for AddToWhitelist events raised by the BiyardNFT contract.
type BiyardNFTAddToWhitelistIterator struct {
	Event *BiyardNFTAddToWhitelist // Event containing the contract specifics and raw log

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
func (it *BiyardNFTAddToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTAddToWhitelist)
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
		it.Event = new(BiyardNFTAddToWhitelist)
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
func (it *BiyardNFTAddToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTAddToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTAddToWhitelist represents a AddToWhitelist event raised by the BiyardNFT contract.
type BiyardNFTAddToWhitelist struct {
	Addr      common.Address
	Whitelist []string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddToWhitelist is a free log retrieval operation binding the contract event 0x37e97b37de4586ebfdefaeaf1702e2e5fbef8c2245ccd04fe76337133d8f5a6c.
//
// Solidity: event AddToWhitelist(address addr, string[] whitelist)
func (_BiyardNFT *BiyardNFTFilterer) FilterAddToWhitelist(opts *bind.FilterOpts) (*BiyardNFTAddToWhitelistIterator, error) {

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "AddToWhitelist")
	if err != nil {
		return nil, err
	}
	return &BiyardNFTAddToWhitelistIterator{contract: _BiyardNFT.contract, event: "AddToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddToWhitelist is a free log subscription operation binding the contract event 0x37e97b37de4586ebfdefaeaf1702e2e5fbef8c2245ccd04fe76337133d8f5a6c.
//
// Solidity: event AddToWhitelist(address addr, string[] whitelist)
func (_BiyardNFT *BiyardNFTFilterer) WatchAddToWhitelist(opts *bind.WatchOpts, sink chan<- *BiyardNFTAddToWhitelist) (event.Subscription, error) {

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "AddToWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTAddToWhitelist)
				if err := _BiyardNFT.contract.UnpackLog(event, "AddToWhitelist", log); err != nil {
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

// ParseAddToWhitelist is a log parse operation binding the contract event 0x37e97b37de4586ebfdefaeaf1702e2e5fbef8c2245ccd04fe76337133d8f5a6c.
//
// Solidity: event AddToWhitelist(address addr, string[] whitelist)
func (_BiyardNFT *BiyardNFTFilterer) ParseAddToWhitelist(log types.Log) (*BiyardNFTAddToWhitelist, error) {
	event := new(BiyardNFTAddToWhitelist)
	if err := _BiyardNFT.contract.UnpackLog(event, "AddToWhitelist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BiyardNFT contract.
type BiyardNFTApprovalIterator struct {
	Event *BiyardNFTApproval // Event containing the contract specifics and raw log

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
func (it *BiyardNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTApproval)
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
		it.Event = new(BiyardNFTApproval)
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
func (it *BiyardNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTApproval represents a Approval event raised by the BiyardNFT contract.
type BiyardNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BiyardNFT *BiyardNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BiyardNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BiyardNFTApprovalIterator{contract: _BiyardNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BiyardNFT *BiyardNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BiyardNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTApproval)
				if err := _BiyardNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BiyardNFT *BiyardNFTFilterer) ParseApproval(log types.Log) (*BiyardNFTApproval, error) {
	event := new(BiyardNFTApproval)
	if err := _BiyardNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the BiyardNFT contract.
type BiyardNFTApprovalForAllIterator struct {
	Event *BiyardNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BiyardNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTApprovalForAll)
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
		it.Event = new(BiyardNFTApprovalForAll)
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
func (it *BiyardNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTApprovalForAll represents a ApprovalForAll event raised by the BiyardNFT contract.
type BiyardNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BiyardNFT *BiyardNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BiyardNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BiyardNFTApprovalForAllIterator{contract: _BiyardNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BiyardNFT *BiyardNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BiyardNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTApprovalForAll)
				if err := _BiyardNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BiyardNFT *BiyardNFTFilterer) ParseApprovalForAll(log types.Log) (*BiyardNFTApprovalForAll, error) {
	event := new(BiyardNFTApprovalForAll)
	if err := _BiyardNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the BiyardNFT contract.
type BiyardNFTClaimIterator struct {
	Event *BiyardNFTClaim // Event containing the contract specifics and raw log

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
func (it *BiyardNFTClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTClaim)
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
		it.Event = new(BiyardNFTClaim)
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
func (it *BiyardNFTClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTClaim represents a Claim event raised by the BiyardNFT contract.
type BiyardNFTClaim struct {
	Whitelisted string
	Addr        common.Address
	TokenId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x72f6bf1bec276ec35226dd14c6846b2a9e3932751965ef5a606e3fb3b7572dc5.
//
// Solidity: event Claim(string whitelisted, address addr, uint256 tokenId)
func (_BiyardNFT *BiyardNFTFilterer) FilterClaim(opts *bind.FilterOpts) (*BiyardNFTClaimIterator, error) {

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &BiyardNFTClaimIterator{contract: _BiyardNFT.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x72f6bf1bec276ec35226dd14c6846b2a9e3932751965ef5a606e3fb3b7572dc5.
//
// Solidity: event Claim(string whitelisted, address addr, uint256 tokenId)
func (_BiyardNFT *BiyardNFTFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *BiyardNFTClaim) (event.Subscription, error) {

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTClaim)
				if err := _BiyardNFT.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x72f6bf1bec276ec35226dd14c6846b2a9e3932751965ef5a606e3fb3b7572dc5.
//
// Solidity: event Claim(string whitelisted, address addr, uint256 tokenId)
func (_BiyardNFT *BiyardNFTFilterer) ParseClaim(log types.Log) (*BiyardNFTClaim, error) {
	event := new(BiyardNFTClaim)
	if err := _BiyardNFT.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BiyardNFT contract.
type BiyardNFTOwnershipTransferredIterator struct {
	Event *BiyardNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BiyardNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTOwnershipTransferred)
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
		it.Event = new(BiyardNFTOwnershipTransferred)
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
func (it *BiyardNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTOwnershipTransferred represents a OwnershipTransferred event raised by the BiyardNFT contract.
type BiyardNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BiyardNFT *BiyardNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BiyardNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BiyardNFTOwnershipTransferredIterator{contract: _BiyardNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BiyardNFT *BiyardNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BiyardNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTOwnershipTransferred)
				if err := _BiyardNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BiyardNFT *BiyardNFTFilterer) ParseOwnershipTransferred(log types.Log) (*BiyardNFTOwnershipTransferred, error) {
	event := new(BiyardNFTOwnershipTransferred)
	if err := _BiyardNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BiyardNFT contract.
type BiyardNFTPausedIterator struct {
	Event *BiyardNFTPaused // Event containing the contract specifics and raw log

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
func (it *BiyardNFTPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTPaused)
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
		it.Event = new(BiyardNFTPaused)
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
func (it *BiyardNFTPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTPaused represents a Paused event raised by the BiyardNFT contract.
type BiyardNFTPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BiyardNFT *BiyardNFTFilterer) FilterPaused(opts *bind.FilterOpts) (*BiyardNFTPausedIterator, error) {

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BiyardNFTPausedIterator{contract: _BiyardNFT.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BiyardNFT *BiyardNFTFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BiyardNFTPaused) (event.Subscription, error) {

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTPaused)
				if err := _BiyardNFT.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BiyardNFT *BiyardNFTFilterer) ParsePaused(log types.Log) (*BiyardNFTPaused, error) {
	event := new(BiyardNFTPaused)
	if err := _BiyardNFT.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTSetClaimLimitIterator is returned from FilterSetClaimLimit and is used to iterate over the raw logs and unpacked data for SetClaimLimit events raised by the BiyardNFT contract.
type BiyardNFTSetClaimLimitIterator struct {
	Event *BiyardNFTSetClaimLimit // Event containing the contract specifics and raw log

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
func (it *BiyardNFTSetClaimLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTSetClaimLimit)
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
		it.Event = new(BiyardNFTSetClaimLimit)
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
func (it *BiyardNFTSetClaimLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTSetClaimLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTSetClaimLimit represents a SetClaimLimit event raised by the BiyardNFT contract.
type BiyardNFTSetClaimLimit struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetClaimLimit is a free log retrieval operation binding the contract event 0xd19d6413488251a5ba14804d3fc507b2b90109b0be980306e6639b443cab016c.
//
// Solidity: event SetClaimLimit(uint256 limit)
func (_BiyardNFT *BiyardNFTFilterer) FilterSetClaimLimit(opts *bind.FilterOpts) (*BiyardNFTSetClaimLimitIterator, error) {

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "SetClaimLimit")
	if err != nil {
		return nil, err
	}
	return &BiyardNFTSetClaimLimitIterator{contract: _BiyardNFT.contract, event: "SetClaimLimit", logs: logs, sub: sub}, nil
}

// WatchSetClaimLimit is a free log subscription operation binding the contract event 0xd19d6413488251a5ba14804d3fc507b2b90109b0be980306e6639b443cab016c.
//
// Solidity: event SetClaimLimit(uint256 limit)
func (_BiyardNFT *BiyardNFTFilterer) WatchSetClaimLimit(opts *bind.WatchOpts, sink chan<- *BiyardNFTSetClaimLimit) (event.Subscription, error) {

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "SetClaimLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTSetClaimLimit)
				if err := _BiyardNFT.contract.UnpackLog(event, "SetClaimLimit", log); err != nil {
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

// ParseSetClaimLimit is a log parse operation binding the contract event 0xd19d6413488251a5ba14804d3fc507b2b90109b0be980306e6639b443cab016c.
//
// Solidity: event SetClaimLimit(uint256 limit)
func (_BiyardNFT *BiyardNFTFilterer) ParseSetClaimLimit(log types.Log) (*BiyardNFTSetClaimLimit, error) {
	event := new(BiyardNFTSetClaimLimit)
	if err := _BiyardNFT.contract.UnpackLog(event, "SetClaimLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTSetMintInfoIterator is returned from FilterSetMintInfo and is used to iterate over the raw logs and unpacked data for SetMintInfo events raised by the BiyardNFT contract.
type BiyardNFTSetMintInfoIterator struct {
	Event *BiyardNFTSetMintInfo // Event containing the contract specifics and raw log

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
func (it *BiyardNFTSetMintInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTSetMintInfo)
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
		it.Event = new(BiyardNFTSetMintInfo)
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
func (it *BiyardNFTSetMintInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTSetMintInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTSetMintInfo represents a SetMintInfo event raised by the BiyardNFT contract.
type BiyardNFTSetMintInfo struct {
	StartTimestamp *big.Int
	StartTokenId   *big.Int
	LastTokenId    *big.Int
	Price          *big.Int
	UseWhitelist   bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetMintInfo is a free log retrieval operation binding the contract event 0x2ad710cf1b4f8c2808c192c201e0a7a237c5f57060a30192054c4b95023b47ab.
//
// Solidity: event SetMintInfo(uint256 start_timestamp, uint256 start_token_id, uint256 last_token_id, uint256 price, bool useWhitelist)
func (_BiyardNFT *BiyardNFTFilterer) FilterSetMintInfo(opts *bind.FilterOpts) (*BiyardNFTSetMintInfoIterator, error) {

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "SetMintInfo")
	if err != nil {
		return nil, err
	}
	return &BiyardNFTSetMintInfoIterator{contract: _BiyardNFT.contract, event: "SetMintInfo", logs: logs, sub: sub}, nil
}

// WatchSetMintInfo is a free log subscription operation binding the contract event 0x2ad710cf1b4f8c2808c192c201e0a7a237c5f57060a30192054c4b95023b47ab.
//
// Solidity: event SetMintInfo(uint256 start_timestamp, uint256 start_token_id, uint256 last_token_id, uint256 price, bool useWhitelist)
func (_BiyardNFT *BiyardNFTFilterer) WatchSetMintInfo(opts *bind.WatchOpts, sink chan<- *BiyardNFTSetMintInfo) (event.Subscription, error) {

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "SetMintInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTSetMintInfo)
				if err := _BiyardNFT.contract.UnpackLog(event, "SetMintInfo", log); err != nil {
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

// ParseSetMintInfo is a log parse operation binding the contract event 0x2ad710cf1b4f8c2808c192c201e0a7a237c5f57060a30192054c4b95023b47ab.
//
// Solidity: event SetMintInfo(uint256 start_timestamp, uint256 start_token_id, uint256 last_token_id, uint256 price, bool useWhitelist)
func (_BiyardNFT *BiyardNFTFilterer) ParseSetMintInfo(log types.Log) (*BiyardNFTSetMintInfo, error) {
	event := new(BiyardNFTSetMintInfo)
	if err := _BiyardNFT.contract.UnpackLog(event, "SetMintInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTSetPartnerQuotaIterator is returned from FilterSetPartnerQuota and is used to iterate over the raw logs and unpacked data for SetPartnerQuota events raised by the BiyardNFT contract.
type BiyardNFTSetPartnerQuotaIterator struct {
	Event *BiyardNFTSetPartnerQuota // Event containing the contract specifics and raw log

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
func (it *BiyardNFTSetPartnerQuotaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTSetPartnerQuota)
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
		it.Event = new(BiyardNFTSetPartnerQuota)
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
func (it *BiyardNFTSetPartnerQuotaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTSetPartnerQuotaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTSetPartnerQuota represents a SetPartnerQuota event raised by the BiyardNFT contract.
type BiyardNFTSetPartnerQuota struct {
	Partner string
	Quota   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetPartnerQuota is a free log retrieval operation binding the contract event 0x2eb88213695caabc15f5393d14d9ccf3e68f9b1b40eee185b442283c6212bede.
//
// Solidity: event SetPartnerQuota(string partner, uint256 quota)
func (_BiyardNFT *BiyardNFTFilterer) FilterSetPartnerQuota(opts *bind.FilterOpts) (*BiyardNFTSetPartnerQuotaIterator, error) {

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "SetPartnerQuota")
	if err != nil {
		return nil, err
	}
	return &BiyardNFTSetPartnerQuotaIterator{contract: _BiyardNFT.contract, event: "SetPartnerQuota", logs: logs, sub: sub}, nil
}

// WatchSetPartnerQuota is a free log subscription operation binding the contract event 0x2eb88213695caabc15f5393d14d9ccf3e68f9b1b40eee185b442283c6212bede.
//
// Solidity: event SetPartnerQuota(string partner, uint256 quota)
func (_BiyardNFT *BiyardNFTFilterer) WatchSetPartnerQuota(opts *bind.WatchOpts, sink chan<- *BiyardNFTSetPartnerQuota) (event.Subscription, error) {

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "SetPartnerQuota")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTSetPartnerQuota)
				if err := _BiyardNFT.contract.UnpackLog(event, "SetPartnerQuota", log); err != nil {
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

// ParseSetPartnerQuota is a log parse operation binding the contract event 0x2eb88213695caabc15f5393d14d9ccf3e68f9b1b40eee185b442283c6212bede.
//
// Solidity: event SetPartnerQuota(string partner, uint256 quota)
func (_BiyardNFT *BiyardNFTFilterer) ParseSetPartnerQuota(log types.Log) (*BiyardNFTSetPartnerQuota, error) {
	event := new(BiyardNFTSetPartnerQuota)
	if err := _BiyardNFT.contract.UnpackLog(event, "SetPartnerQuota", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BiyardNFT contract.
type BiyardNFTTransferIterator struct {
	Event *BiyardNFTTransfer // Event containing the contract specifics and raw log

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
func (it *BiyardNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTTransfer)
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
		it.Event = new(BiyardNFTTransfer)
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
func (it *BiyardNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTTransfer represents a Transfer event raised by the BiyardNFT contract.
type BiyardNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BiyardNFT *BiyardNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BiyardNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BiyardNFTTransferIterator{contract: _BiyardNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BiyardNFT *BiyardNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BiyardNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTTransfer)
				if err := _BiyardNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BiyardNFT *BiyardNFTFilterer) ParseTransfer(log types.Log) (*BiyardNFTTransfer, error) {
	event := new(BiyardNFTTransfer)
	if err := _BiyardNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BiyardNFT contract.
type BiyardNFTUnpausedIterator struct {
	Event *BiyardNFTUnpaused // Event containing the contract specifics and raw log

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
func (it *BiyardNFTUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTUnpaused)
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
		it.Event = new(BiyardNFTUnpaused)
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
func (it *BiyardNFTUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTUnpaused represents a Unpaused event raised by the BiyardNFT contract.
type BiyardNFTUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BiyardNFT *BiyardNFTFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BiyardNFTUnpausedIterator, error) {

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BiyardNFTUnpausedIterator{contract: _BiyardNFT.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BiyardNFT *BiyardNFTFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BiyardNFTUnpaused) (event.Subscription, error) {

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTUnpaused)
				if err := _BiyardNFT.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BiyardNFT *BiyardNFTFilterer) ParseUnpaused(log types.Log) (*BiyardNFTUnpaused, error) {
	event := new(BiyardNFTUnpaused)
	if err := _BiyardNFT.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BiyardNFTWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the BiyardNFT contract.
type BiyardNFTWithdrawalIterator struct {
	Event *BiyardNFTWithdrawal // Event containing the contract specifics and raw log

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
func (it *BiyardNFTWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiyardNFTWithdrawal)
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
		it.Event = new(BiyardNFTWithdrawal)
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
func (it *BiyardNFTWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiyardNFTWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiyardNFTWithdrawal represents a Withdrawal event raised by the BiyardNFT contract.
type BiyardNFTWithdrawal struct {
	Amount *big.Int
	When   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xbf2ed60bd5b5965d685680c01195c9514e4382e28e3a5a2d2d5244bf59411b93.
//
// Solidity: event Withdrawal(uint256 amount, uint256 when)
func (_BiyardNFT *BiyardNFTFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*BiyardNFTWithdrawalIterator, error) {

	logs, sub, err := _BiyardNFT.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &BiyardNFTWithdrawalIterator{contract: _BiyardNFT.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xbf2ed60bd5b5965d685680c01195c9514e4382e28e3a5a2d2d5244bf59411b93.
//
// Solidity: event Withdrawal(uint256 amount, uint256 when)
func (_BiyardNFT *BiyardNFTFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *BiyardNFTWithdrawal) (event.Subscription, error) {

	logs, sub, err := _BiyardNFT.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiyardNFTWithdrawal)
				if err := _BiyardNFT.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xbf2ed60bd5b5965d685680c01195c9514e4382e28e3a5a2d2d5244bf59411b93.
//
// Solidity: event Withdrawal(uint256 amount, uint256 when)
func (_BiyardNFT *BiyardNFTFilterer) ParseWithdrawal(log types.Log) (*BiyardNFTWithdrawal, error) {
	event := new(BiyardNFTWithdrawal)
	if err := _BiyardNFT.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}
