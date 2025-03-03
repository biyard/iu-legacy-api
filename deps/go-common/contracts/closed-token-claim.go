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

// ClosedTokenClaimMetaData contains all meta data concerning the ClosedTokenClaim contract.
var ClosedTokenClaimMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"AddPartner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicLimit\",\"type\":\"uint256\"}],\"name\":\"SetClaimLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"open\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPartner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endPartner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPublic\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endPublic\",\"type\":\"uint256\"}],\"name\":\"SetupMinting\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"partnerName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"addPartner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"partner\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addPartnerWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allSlots\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"allowedOf\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"otp\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"claimToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfCandidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"partnerName\",\"type\":\"string\"}],\"name\":\"partnerSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainedSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"setAllowDuplicatedClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"partnerLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicLimit\",\"type\":\"uint256\"}],\"name\":\"setClaimLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startPartner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPartner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPublic\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPublic\",\"type\":\"uint256\"}],\"name\":\"setTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"open\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPartner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPartner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPublic\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPublic\",\"type\":\"uint256\"}],\"name\":\"setupMinting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"updateContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClosedTokenClaimABI is the input ABI used to generate the binding from.
// Deprecated: Use ClosedTokenClaimMetaData.ABI instead.
var ClosedTokenClaimABI = ClosedTokenClaimMetaData.ABI

// ClosedTokenClaimBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ClosedTokenClaimBinRuntime = ``

// ClosedTokenClaim is an auto generated Go binding around a Klaytn contract.
type ClosedTokenClaim struct {
	ClosedTokenClaimCaller     // Read-only binding to the contract
	ClosedTokenClaimTransactor // Write-only binding to the contract
	ClosedTokenClaimFilterer   // Log filterer for contract events
}

// ClosedTokenClaimCaller is an auto generated read-only Go binding around a Klaytn contract.
type ClosedTokenClaimCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClosedTokenClaimTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ClosedTokenClaimTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClosedTokenClaimFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ClosedTokenClaimFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClosedTokenClaimSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ClosedTokenClaimSession struct {
	Contract     *ClosedTokenClaim // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClosedTokenClaimCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ClosedTokenClaimCallerSession struct {
	Contract *ClosedTokenClaimCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ClosedTokenClaimTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ClosedTokenClaimTransactorSession struct {
	Contract     *ClosedTokenClaimTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ClosedTokenClaimRaw is an auto generated low-level Go binding around a Klaytn contract.
type ClosedTokenClaimRaw struct {
	Contract *ClosedTokenClaim // Generic contract binding to access the raw methods on
}

// ClosedTokenClaimCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ClosedTokenClaimCallerRaw struct {
	Contract *ClosedTokenClaimCaller // Generic read-only contract binding to access the raw methods on
}

// ClosedTokenClaimTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ClosedTokenClaimTransactorRaw struct {
	Contract *ClosedTokenClaimTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClosedTokenClaim creates a new instance of ClosedTokenClaim, bound to a specific deployed contract.
func NewClosedTokenClaim(address common.Address, backend bind.ContractBackend) (*ClosedTokenClaim, error) {
	contract, err := bindClosedTokenClaim(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClosedTokenClaim{ClosedTokenClaimCaller: ClosedTokenClaimCaller{contract: contract}, ClosedTokenClaimTransactor: ClosedTokenClaimTransactor{contract: contract}, ClosedTokenClaimFilterer: ClosedTokenClaimFilterer{contract: contract}}, nil
}

// NewClosedTokenClaimCaller creates a new read-only instance of ClosedTokenClaim, bound to a specific deployed contract.
func NewClosedTokenClaimCaller(address common.Address, caller bind.ContractCaller) (*ClosedTokenClaimCaller, error) {
	contract, err := bindClosedTokenClaim(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClosedTokenClaimCaller{contract: contract}, nil
}

// NewClosedTokenClaimTransactor creates a new write-only instance of ClosedTokenClaim, bound to a specific deployed contract.
func NewClosedTokenClaimTransactor(address common.Address, transactor bind.ContractTransactor) (*ClosedTokenClaimTransactor, error) {
	contract, err := bindClosedTokenClaim(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClosedTokenClaimTransactor{contract: contract}, nil
}

// NewClosedTokenClaimFilterer creates a new log filterer instance of ClosedTokenClaim, bound to a specific deployed contract.
func NewClosedTokenClaimFilterer(address common.Address, filterer bind.ContractFilterer) (*ClosedTokenClaimFilterer, error) {
	contract, err := bindClosedTokenClaim(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClosedTokenClaimFilterer{contract: contract}, nil
}

// bindClosedTokenClaim binds a generic wrapper to an already deployed contract.
func bindClosedTokenClaim(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClosedTokenClaimMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClosedTokenClaim *ClosedTokenClaimRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClosedTokenClaim.Contract.ClosedTokenClaimCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClosedTokenClaim *ClosedTokenClaimRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.ClosedTokenClaimTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClosedTokenClaim *ClosedTokenClaimRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.ClosedTokenClaimTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClosedTokenClaim *ClosedTokenClaimCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClosedTokenClaim.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClosedTokenClaim *ClosedTokenClaimTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClosedTokenClaim *ClosedTokenClaimTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.contract.Transact(opts, method, params...)
}

// AllSlots is a free data retrieval call binding the contract method 0xdded2326.
//
// Solidity: function allSlots() view returns(string[], uint256[])
func (_ClosedTokenClaim *ClosedTokenClaimCaller) AllSlots(opts *bind.CallOpts) ([]string, []*big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ClosedTokenClaim.contract.Call(opts, out, "allSlots")
	return *ret0, *ret1, err
}

// AllSlots is a free data retrieval call binding the contract method 0xdded2326.
//
// Solidity: function allSlots() view returns(string[], uint256[])
func (_ClosedTokenClaim *ClosedTokenClaimSession) AllSlots() ([]string, []*big.Int, error) {
	return _ClosedTokenClaim.Contract.AllSlots(&_ClosedTokenClaim.CallOpts)
}

// AllSlots is a free data retrieval call binding the contract method 0xdded2326.
//
// Solidity: function allSlots() view returns(string[], uint256[])
func (_ClosedTokenClaim *ClosedTokenClaimCallerSession) AllSlots() ([]string, []*big.Int, error) {
	return _ClosedTokenClaim.Contract.AllSlots(&_ClosedTokenClaim.CallOpts)
}

// AllowedOf is a free data retrieval call binding the contract method 0x5dbff683.
//
// Solidity: function allowedOf(address addr) view returns(string[])
func (_ClosedTokenClaim *ClosedTokenClaimCaller) AllowedOf(opts *bind.CallOpts, addr common.Address) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ClosedTokenClaim.contract.Call(opts, out, "allowedOf", addr)
	return *ret0, err
}

// AllowedOf is a free data retrieval call binding the contract method 0x5dbff683.
//
// Solidity: function allowedOf(address addr) view returns(string[])
func (_ClosedTokenClaim *ClosedTokenClaimSession) AllowedOf(addr common.Address) ([]string, error) {
	return _ClosedTokenClaim.Contract.AllowedOf(&_ClosedTokenClaim.CallOpts, addr)
}

// AllowedOf is a free data retrieval call binding the contract method 0x5dbff683.
//
// Solidity: function allowedOf(address addr) view returns(string[])
func (_ClosedTokenClaim *ClosedTokenClaimCallerSession) AllowedOf(addr common.Address) ([]string, error) {
	return _ClosedTokenClaim.Contract.AllowedOf(&_ClosedTokenClaim.CallOpts, addr)
}

// IsClaimed is a free data retrieval call binding the contract method 0x57c9ca14.
//
// Solidity: function isClaimed() view returns(uint256, bool)
func (_ClosedTokenClaim *ClosedTokenClaimCaller) IsClaimed(opts *bind.CallOpts) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ClosedTokenClaim.contract.Call(opts, out, "isClaimed")
	return *ret0, *ret1, err
}

// IsClaimed is a free data retrieval call binding the contract method 0x57c9ca14.
//
// Solidity: function isClaimed() view returns(uint256, bool)
func (_ClosedTokenClaim *ClosedTokenClaimSession) IsClaimed() (*big.Int, bool, error) {
	return _ClosedTokenClaim.Contract.IsClaimed(&_ClosedTokenClaim.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0x57c9ca14.
//
// Solidity: function isClaimed() view returns(uint256, bool)
func (_ClosedTokenClaim *ClosedTokenClaimCallerSession) IsClaimed() (*big.Int, bool, error) {
	return _ClosedTokenClaim.Contract.IsClaimed(&_ClosedTokenClaim.CallOpts)
}

// NumberOfCandidates is a free data retrieval call binding the contract method 0x9c84bb27.
//
// Solidity: function numberOfCandidates() view returns(uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCaller) NumberOfCandidates(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClosedTokenClaim.contract.Call(opts, out, "numberOfCandidates")
	return *ret0, err
}

// NumberOfCandidates is a free data retrieval call binding the contract method 0x9c84bb27.
//
// Solidity: function numberOfCandidates() view returns(uint256)
func (_ClosedTokenClaim *ClosedTokenClaimSession) NumberOfCandidates() (*big.Int, error) {
	return _ClosedTokenClaim.Contract.NumberOfCandidates(&_ClosedTokenClaim.CallOpts)
}

// NumberOfCandidates is a free data retrieval call binding the contract method 0x9c84bb27.
//
// Solidity: function numberOfCandidates() view returns(uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCallerSession) NumberOfCandidates() (*big.Int, error) {
	return _ClosedTokenClaim.Contract.NumberOfCandidates(&_ClosedTokenClaim.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClosedTokenClaim *ClosedTokenClaimCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClosedTokenClaim.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClosedTokenClaim *ClosedTokenClaimSession) Owner() (common.Address, error) {
	return _ClosedTokenClaim.Contract.Owner(&_ClosedTokenClaim.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClosedTokenClaim *ClosedTokenClaimCallerSession) Owner() (common.Address, error) {
	return _ClosedTokenClaim.Contract.Owner(&_ClosedTokenClaim.CallOpts)
}

// PartnerSlots is a free data retrieval call binding the contract method 0x9510b8e7.
//
// Solidity: function partnerSlots(string partnerName) view returns(uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCaller) PartnerSlots(opts *bind.CallOpts, partnerName string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClosedTokenClaim.contract.Call(opts, out, "partnerSlots", partnerName)
	return *ret0, err
}

// PartnerSlots is a free data retrieval call binding the contract method 0x9510b8e7.
//
// Solidity: function partnerSlots(string partnerName) view returns(uint256)
func (_ClosedTokenClaim *ClosedTokenClaimSession) PartnerSlots(partnerName string) (*big.Int, error) {
	return _ClosedTokenClaim.Contract.PartnerSlots(&_ClosedTokenClaim.CallOpts, partnerName)
}

// PartnerSlots is a free data retrieval call binding the contract method 0x9510b8e7.
//
// Solidity: function partnerSlots(string partnerName) view returns(uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCallerSession) PartnerSlots(partnerName string) (*big.Int, error) {
	return _ClosedTokenClaim.Contract.PartnerSlots(&_ClosedTokenClaim.CallOpts, partnerName)
}

// RemainedSlot is a free data retrieval call binding the contract method 0xddeffb92.
//
// Solidity: function remainedSlot() view returns(uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCaller) RemainedSlot(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClosedTokenClaim.contract.Call(opts, out, "remainedSlot")
	return *ret0, err
}

// RemainedSlot is a free data retrieval call binding the contract method 0xddeffb92.
//
// Solidity: function remainedSlot() view returns(uint256)
func (_ClosedTokenClaim *ClosedTokenClaimSession) RemainedSlot() (*big.Int, error) {
	return _ClosedTokenClaim.Contract.RemainedSlot(&_ClosedTokenClaim.CallOpts)
}

// RemainedSlot is a free data retrieval call binding the contract method 0xddeffb92.
//
// Solidity: function remainedSlot() view returns(uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCallerSession) RemainedSlot() (*big.Int, error) {
	return _ClosedTokenClaim.Contract.RemainedSlot(&_ClosedTokenClaim.CallOpts)
}

// Time is a free data retrieval call binding the contract method 0x16ada547.
//
// Solidity: function time() view returns(uint256, uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCaller) Time(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ClosedTokenClaim.contract.Call(opts, out, "time")
	return *ret0, *ret1, err
}

// Time is a free data retrieval call binding the contract method 0x16ada547.
//
// Solidity: function time() view returns(uint256, uint256)
func (_ClosedTokenClaim *ClosedTokenClaimSession) Time() (*big.Int, *big.Int, error) {
	return _ClosedTokenClaim.Contract.Time(&_ClosedTokenClaim.CallOpts)
}

// Time is a free data retrieval call binding the contract method 0x16ada547.
//
// Solidity: function time() view returns(uint256, uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCallerSession) Time() (*big.Int, *big.Int, error) {
	return _ClosedTokenClaim.Contract.Time(&_ClosedTokenClaim.CallOpts)
}

// TimeAll is a free data retrieval call binding the contract method 0xa01a7fa6.
//
// Solidity: function timeAll() view returns(uint256, uint256, uint256, uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCaller) TimeAll(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _ClosedTokenClaim.contract.Call(opts, out, "timeAll")
	return *ret0, *ret1, *ret2, *ret3, err
}

// TimeAll is a free data retrieval call binding the contract method 0xa01a7fa6.
//
// Solidity: function timeAll() view returns(uint256, uint256, uint256, uint256)
func (_ClosedTokenClaim *ClosedTokenClaimSession) TimeAll() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ClosedTokenClaim.Contract.TimeAll(&_ClosedTokenClaim.CallOpts)
}

// TimeAll is a free data retrieval call binding the contract method 0xa01a7fa6.
//
// Solidity: function timeAll() view returns(uint256, uint256, uint256, uint256)
func (_ClosedTokenClaim *ClosedTokenClaimCallerSession) TimeAll() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ClosedTokenClaim.Contract.TimeAll(&_ClosedTokenClaim.CallOpts)
}

// AddPartner is a paid mutator transaction binding the contract method 0x9a3df92a.
//
// Solidity: function addPartner(string partnerName, uint256 supply) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) AddPartner(opts *bind.TransactOpts, partnerName string, supply *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "addPartner", partnerName, supply)
}

// AddPartner is a paid mutator transaction binding the contract method 0x9a3df92a.
//
// Solidity: function addPartner(string partnerName, uint256 supply) returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) AddPartner(partnerName string, supply *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.AddPartner(&_ClosedTokenClaim.TransactOpts, partnerName, supply)
}

// AddPartner is a paid mutator transaction binding the contract method 0x9a3df92a.
//
// Solidity: function addPartner(string partnerName, uint256 supply) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) AddPartner(partnerName string, supply *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.AddPartner(&_ClosedTokenClaim.TransactOpts, partnerName, supply)
}

// AddPartnerWhitelist is a paid mutator transaction binding the contract method 0x2efed898.
//
// Solidity: function addPartnerWhitelist(string partner, address[] addrs) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) AddPartnerWhitelist(opts *bind.TransactOpts, partner string, addrs []common.Address) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "addPartnerWhitelist", partner, addrs)
}

// AddPartnerWhitelist is a paid mutator transaction binding the contract method 0x2efed898.
//
// Solidity: function addPartnerWhitelist(string partner, address[] addrs) returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) AddPartnerWhitelist(partner string, addrs []common.Address) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.AddPartnerWhitelist(&_ClosedTokenClaim.TransactOpts, partner, addrs)
}

// AddPartnerWhitelist is a paid mutator transaction binding the contract method 0x2efed898.
//
// Solidity: function addPartnerWhitelist(string partner, address[] addrs) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) AddPartnerWhitelist(partner string, addrs []common.Address) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.AddPartnerWhitelist(&_ClosedTokenClaim.TransactOpts, partner, addrs)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x70e55757.
//
// Solidity: function claimToken(string otp, bytes32 hash) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) ClaimToken(opts *bind.TransactOpts, otp string, hash [32]byte) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "claimToken", otp, hash)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x70e55757.
//
// Solidity: function claimToken(string otp, bytes32 hash) returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) ClaimToken(otp string, hash [32]byte) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.ClaimToken(&_ClosedTokenClaim.TransactOpts, otp, hash)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x70e55757.
//
// Solidity: function claimToken(string otp, bytes32 hash) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) ClaimToken(otp string, hash [32]byte) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.ClaimToken(&_ClosedTokenClaim.TransactOpts, otp, hash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) RenounceOwnership() (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.RenounceOwnership(&_ClosedTokenClaim.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.RenounceOwnership(&_ClosedTokenClaim.TransactOpts)
}

// SetAllowDuplicatedClaim is a paid mutator transaction binding the contract method 0xd859dd6b.
//
// Solidity: function setAllowDuplicatedClaim(bool allow) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) SetAllowDuplicatedClaim(opts *bind.TransactOpts, allow bool) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "setAllowDuplicatedClaim", allow)
}

// SetAllowDuplicatedClaim is a paid mutator transaction binding the contract method 0xd859dd6b.
//
// Solidity: function setAllowDuplicatedClaim(bool allow) returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) SetAllowDuplicatedClaim(allow bool) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.SetAllowDuplicatedClaim(&_ClosedTokenClaim.TransactOpts, allow)
}

// SetAllowDuplicatedClaim is a paid mutator transaction binding the contract method 0xd859dd6b.
//
// Solidity: function setAllowDuplicatedClaim(bool allow) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) SetAllowDuplicatedClaim(allow bool) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.SetAllowDuplicatedClaim(&_ClosedTokenClaim.TransactOpts, allow)
}

// SetClaimLimit is a paid mutator transaction binding the contract method 0x940312a9.
//
// Solidity: function setClaimLimit(uint256 partnerLimit, uint256 publicLimit) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) SetClaimLimit(opts *bind.TransactOpts, partnerLimit *big.Int, publicLimit *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "setClaimLimit", partnerLimit, publicLimit)
}

// SetClaimLimit is a paid mutator transaction binding the contract method 0x940312a9.
//
// Solidity: function setClaimLimit(uint256 partnerLimit, uint256 publicLimit) returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) SetClaimLimit(partnerLimit *big.Int, publicLimit *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.SetClaimLimit(&_ClosedTokenClaim.TransactOpts, partnerLimit, publicLimit)
}

// SetClaimLimit is a paid mutator transaction binding the contract method 0x940312a9.
//
// Solidity: function setClaimLimit(uint256 partnerLimit, uint256 publicLimit) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) SetClaimLimit(partnerLimit *big.Int, publicLimit *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.SetClaimLimit(&_ClosedTokenClaim.TransactOpts, partnerLimit, publicLimit)
}

// SetTime is a paid mutator transaction binding the contract method 0xf938c417.
//
// Solidity: function setTime(uint256 startPartner, uint256 endPartner, uint256 startPublic, uint256 endPublic) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) SetTime(opts *bind.TransactOpts, startPartner *big.Int, endPartner *big.Int, startPublic *big.Int, endPublic *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "setTime", startPartner, endPartner, startPublic, endPublic)
}

// SetTime is a paid mutator transaction binding the contract method 0xf938c417.
//
// Solidity: function setTime(uint256 startPartner, uint256 endPartner, uint256 startPublic, uint256 endPublic) returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) SetTime(startPartner *big.Int, endPartner *big.Int, startPublic *big.Int, endPublic *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.SetTime(&_ClosedTokenClaim.TransactOpts, startPartner, endPartner, startPublic, endPublic)
}

// SetTime is a paid mutator transaction binding the contract method 0xf938c417.
//
// Solidity: function setTime(uint256 startPartner, uint256 endPartner, uint256 startPublic, uint256 endPublic) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) SetTime(startPartner *big.Int, endPartner *big.Int, startPublic *big.Int, endPublic *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.SetTime(&_ClosedTokenClaim.TransactOpts, startPartner, endPartner, startPublic, endPublic)
}

// SetupMinting is a paid mutator transaction binding the contract method 0x59cf1bbd.
//
// Solidity: function setupMinting(uint256 tokenId, uint256 open, uint256 startPartner, uint256 endPartner, uint256 startPublic, uint256 endPublic) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) SetupMinting(opts *bind.TransactOpts, tokenId *big.Int, open *big.Int, startPartner *big.Int, endPartner *big.Int, startPublic *big.Int, endPublic *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "setupMinting", tokenId, open, startPartner, endPartner, startPublic, endPublic)
}

// SetupMinting is a paid mutator transaction binding the contract method 0x59cf1bbd.
//
// Solidity: function setupMinting(uint256 tokenId, uint256 open, uint256 startPartner, uint256 endPartner, uint256 startPublic, uint256 endPublic) returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) SetupMinting(tokenId *big.Int, open *big.Int, startPartner *big.Int, endPartner *big.Int, startPublic *big.Int, endPublic *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.SetupMinting(&_ClosedTokenClaim.TransactOpts, tokenId, open, startPartner, endPartner, startPublic, endPublic)
}

// SetupMinting is a paid mutator transaction binding the contract method 0x59cf1bbd.
//
// Solidity: function setupMinting(uint256 tokenId, uint256 open, uint256 startPartner, uint256 endPartner, uint256 startPublic, uint256 endPublic) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) SetupMinting(tokenId *big.Int, open *big.Int, startPartner *big.Int, endPartner *big.Int, startPublic *big.Int, endPublic *big.Int) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.SetupMinting(&_ClosedTokenClaim.TransactOpts, tokenId, open, startPartner, endPartner, startPublic, endPublic)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.TransferOwnership(&_ClosedTokenClaim.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.TransferOwnership(&_ClosedTokenClaim.TransactOpts, newOwner)
}

// UpdateContractAddress is a paid mutator transaction binding the contract method 0xbdf3e088.
//
// Solidity: function updateContractAddress(address minter) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactor) UpdateContractAddress(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _ClosedTokenClaim.contract.Transact(opts, "updateContractAddress", minter)
}

// UpdateContractAddress is a paid mutator transaction binding the contract method 0xbdf3e088.
//
// Solidity: function updateContractAddress(address minter) returns()
func (_ClosedTokenClaim *ClosedTokenClaimSession) UpdateContractAddress(minter common.Address) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.UpdateContractAddress(&_ClosedTokenClaim.TransactOpts, minter)
}

// UpdateContractAddress is a paid mutator transaction binding the contract method 0xbdf3e088.
//
// Solidity: function updateContractAddress(address minter) returns()
func (_ClosedTokenClaim *ClosedTokenClaimTransactorSession) UpdateContractAddress(minter common.Address) (*types.Transaction, error) {
	return _ClosedTokenClaim.Contract.UpdateContractAddress(&_ClosedTokenClaim.TransactOpts, minter)
}

// ClosedTokenClaimAddPartnerIterator is returned from FilterAddPartner and is used to iterate over the raw logs and unpacked data for AddPartner events raised by the ClosedTokenClaim contract.
type ClosedTokenClaimAddPartnerIterator struct {
	Event *ClosedTokenClaimAddPartner // Event containing the contract specifics and raw log

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
func (it *ClosedTokenClaimAddPartnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosedTokenClaimAddPartner)
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
		it.Event = new(ClosedTokenClaimAddPartner)
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
func (it *ClosedTokenClaimAddPartnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosedTokenClaimAddPartnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosedTokenClaimAddPartner represents a AddPartner event raised by the ClosedTokenClaim contract.
type ClosedTokenClaimAddPartner struct {
	Name   string
	Supply *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddPartner is a free log retrieval operation binding the contract event 0x82d149b9acc00260cde0e3266b502e3f8c4fe337a2579c38c899b6e9ee14ee87.
//
// Solidity: event AddPartner(string name, uint256 supply)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) FilterAddPartner(opts *bind.FilterOpts) (*ClosedTokenClaimAddPartnerIterator, error) {

	logs, sub, err := _ClosedTokenClaim.contract.FilterLogs(opts, "AddPartner")
	if err != nil {
		return nil, err
	}
	return &ClosedTokenClaimAddPartnerIterator{contract: _ClosedTokenClaim.contract, event: "AddPartner", logs: logs, sub: sub}, nil
}

// WatchAddPartner is a free log subscription operation binding the contract event 0x82d149b9acc00260cde0e3266b502e3f8c4fe337a2579c38c899b6e9ee14ee87.
//
// Solidity: event AddPartner(string name, uint256 supply)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) WatchAddPartner(opts *bind.WatchOpts, sink chan<- *ClosedTokenClaimAddPartner) (event.Subscription, error) {

	logs, sub, err := _ClosedTokenClaim.contract.WatchLogs(opts, "AddPartner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosedTokenClaimAddPartner)
				if err := _ClosedTokenClaim.contract.UnpackLog(event, "AddPartner", log); err != nil {
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

// ParseAddPartner is a log parse operation binding the contract event 0x82d149b9acc00260cde0e3266b502e3f8c4fe337a2579c38c899b6e9ee14ee87.
//
// Solidity: event AddPartner(string name, uint256 supply)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) ParseAddPartner(log types.Log) (*ClosedTokenClaimAddPartner, error) {
	event := new(ClosedTokenClaimAddPartner)
	if err := _ClosedTokenClaim.contract.UnpackLog(event, "AddPartner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ClosedTokenClaimOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ClosedTokenClaim contract.
type ClosedTokenClaimOwnershipTransferredIterator struct {
	Event *ClosedTokenClaimOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClosedTokenClaimOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosedTokenClaimOwnershipTransferred)
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
		it.Event = new(ClosedTokenClaimOwnershipTransferred)
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
func (it *ClosedTokenClaimOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosedTokenClaimOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosedTokenClaimOwnershipTransferred represents a OwnershipTransferred event raised by the ClosedTokenClaim contract.
type ClosedTokenClaimOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClosedTokenClaimOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ClosedTokenClaim.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClosedTokenClaimOwnershipTransferredIterator{contract: _ClosedTokenClaim.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClosedTokenClaimOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ClosedTokenClaim.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosedTokenClaimOwnershipTransferred)
				if err := _ClosedTokenClaim.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) ParseOwnershipTransferred(log types.Log) (*ClosedTokenClaimOwnershipTransferred, error) {
	event := new(ClosedTokenClaimOwnershipTransferred)
	if err := _ClosedTokenClaim.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ClosedTokenClaimSetClaimLimitIterator is returned from FilterSetClaimLimit and is used to iterate over the raw logs and unpacked data for SetClaimLimit events raised by the ClosedTokenClaim contract.
type ClosedTokenClaimSetClaimLimitIterator struct {
	Event *ClosedTokenClaimSetClaimLimit // Event containing the contract specifics and raw log

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
func (it *ClosedTokenClaimSetClaimLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosedTokenClaimSetClaimLimit)
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
		it.Event = new(ClosedTokenClaimSetClaimLimit)
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
func (it *ClosedTokenClaimSetClaimLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosedTokenClaimSetClaimLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosedTokenClaimSetClaimLimit represents a SetClaimLimit event raised by the ClosedTokenClaim contract.
type ClosedTokenClaimSetClaimLimit struct {
	PartnerLimit *big.Int
	PublicLimit  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetClaimLimit is a free log retrieval operation binding the contract event 0x29ccccbef07ba9132979e670e8950bdd8ed522c328c7c73bbbb42d5352d5ba78.
//
// Solidity: event SetClaimLimit(uint256 partnerLimit, uint256 publicLimit)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) FilterSetClaimLimit(opts *bind.FilterOpts) (*ClosedTokenClaimSetClaimLimitIterator, error) {

	logs, sub, err := _ClosedTokenClaim.contract.FilterLogs(opts, "SetClaimLimit")
	if err != nil {
		return nil, err
	}
	return &ClosedTokenClaimSetClaimLimitIterator{contract: _ClosedTokenClaim.contract, event: "SetClaimLimit", logs: logs, sub: sub}, nil
}

// WatchSetClaimLimit is a free log subscription operation binding the contract event 0x29ccccbef07ba9132979e670e8950bdd8ed522c328c7c73bbbb42d5352d5ba78.
//
// Solidity: event SetClaimLimit(uint256 partnerLimit, uint256 publicLimit)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) WatchSetClaimLimit(opts *bind.WatchOpts, sink chan<- *ClosedTokenClaimSetClaimLimit) (event.Subscription, error) {

	logs, sub, err := _ClosedTokenClaim.contract.WatchLogs(opts, "SetClaimLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosedTokenClaimSetClaimLimit)
				if err := _ClosedTokenClaim.contract.UnpackLog(event, "SetClaimLimit", log); err != nil {
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

// ParseSetClaimLimit is a log parse operation binding the contract event 0x29ccccbef07ba9132979e670e8950bdd8ed522c328c7c73bbbb42d5352d5ba78.
//
// Solidity: event SetClaimLimit(uint256 partnerLimit, uint256 publicLimit)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) ParseSetClaimLimit(log types.Log) (*ClosedTokenClaimSetClaimLimit, error) {
	event := new(ClosedTokenClaimSetClaimLimit)
	if err := _ClosedTokenClaim.contract.UnpackLog(event, "SetClaimLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ClosedTokenClaimSetupMintingIterator is returned from FilterSetupMinting and is used to iterate over the raw logs and unpacked data for SetupMinting events raised by the ClosedTokenClaim contract.
type ClosedTokenClaimSetupMintingIterator struct {
	Event *ClosedTokenClaimSetupMinting // Event containing the contract specifics and raw log

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
func (it *ClosedTokenClaimSetupMintingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosedTokenClaimSetupMinting)
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
		it.Event = new(ClosedTokenClaimSetupMinting)
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
func (it *ClosedTokenClaimSetupMintingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosedTokenClaimSetupMintingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosedTokenClaimSetupMinting represents a SetupMinting event raised by the ClosedTokenClaim contract.
type ClosedTokenClaimSetupMinting struct {
	TokenId      *big.Int
	Open         *big.Int
	StartPartner *big.Int
	EndPartner   *big.Int
	StartPublic  *big.Int
	EndPublic    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetupMinting is a free log retrieval operation binding the contract event 0xa63cb8d9e9ef07df4a349837d408f52bfac8d4545f5e65ff27293ab85481abb1.
//
// Solidity: event SetupMinting(uint256 tokenId, uint256 open, uint256 startPartner, uint256 endPartner, uint256 startPublic, uint256 endPublic)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) FilterSetupMinting(opts *bind.FilterOpts) (*ClosedTokenClaimSetupMintingIterator, error) {

	logs, sub, err := _ClosedTokenClaim.contract.FilterLogs(opts, "SetupMinting")
	if err != nil {
		return nil, err
	}
	return &ClosedTokenClaimSetupMintingIterator{contract: _ClosedTokenClaim.contract, event: "SetupMinting", logs: logs, sub: sub}, nil
}

// WatchSetupMinting is a free log subscription operation binding the contract event 0xa63cb8d9e9ef07df4a349837d408f52bfac8d4545f5e65ff27293ab85481abb1.
//
// Solidity: event SetupMinting(uint256 tokenId, uint256 open, uint256 startPartner, uint256 endPartner, uint256 startPublic, uint256 endPublic)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) WatchSetupMinting(opts *bind.WatchOpts, sink chan<- *ClosedTokenClaimSetupMinting) (event.Subscription, error) {

	logs, sub, err := _ClosedTokenClaim.contract.WatchLogs(opts, "SetupMinting")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosedTokenClaimSetupMinting)
				if err := _ClosedTokenClaim.contract.UnpackLog(event, "SetupMinting", log); err != nil {
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

// ParseSetupMinting is a log parse operation binding the contract event 0xa63cb8d9e9ef07df4a349837d408f52bfac8d4545f5e65ff27293ab85481abb1.
//
// Solidity: event SetupMinting(uint256 tokenId, uint256 open, uint256 startPartner, uint256 endPartner, uint256 startPublic, uint256 endPublic)
func (_ClosedTokenClaim *ClosedTokenClaimFilterer) ParseSetupMinting(log types.Log) (*ClosedTokenClaimSetupMinting, error) {
	event := new(ClosedTokenClaimSetupMinting)
	if err := _ClosedTokenClaim.contract.UnpackLog(event, "SetupMinting", log); err != nil {
		return nil, err
	}
	return event, nil
}
