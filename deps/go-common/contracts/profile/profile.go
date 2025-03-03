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

// DeploymentParameterV1 is an auto generated low-level Go binding around an user-defined struct.
type DeploymentParameterV1 struct {
	Strings     []string
	Numbers     []*big.Int
	Bools       []bool
	Addresses   []common.Address
	Strings2d   [][]string
	Numbers2d   [][]*big.Int
	Addresses2d [][]common.Address
	Bools2d     [][]bool
}

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

// ProfileMetaData contains all meta data concerning the Profile contract.
var ProfileMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"stateAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"beforeUpgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"changeNicknameEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"completeSetupEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitAppDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityType\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structProposalSummary\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"emitProposalActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"indexed\":false,\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"injectDependancyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setProfileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistryEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNickName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNickNameList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStandardNftAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"injectDependancy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameExt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setNickname\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updateNickname\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProfileABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfileMetaData.ABI instead.
var ProfileABI = ProfileMetaData.ABI

// ProfileBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ProfileBinRuntime = ``

// Profile is an auto generated Go binding around a Klaytn contract.
type Profile struct {
	ProfileCaller     // Read-only binding to the contract
	ProfileTransactor // Write-only binding to the contract
	ProfileFilterer   // Log filterer for contract events
}

// ProfileCaller is an auto generated read-only Go binding around a Klaytn contract.
type ProfileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ProfileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ProfileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ProfileSession struct {
	Contract     *Profile          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfileCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ProfileCallerSession struct {
	Contract *ProfileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProfileTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ProfileTransactorSession struct {
	Contract     *ProfileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProfileRaw is an auto generated low-level Go binding around a Klaytn contract.
type ProfileRaw struct {
	Contract *Profile // Generic contract binding to access the raw methods on
}

// ProfileCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ProfileCallerRaw struct {
	Contract *ProfileCaller // Generic read-only contract binding to access the raw methods on
}

// ProfileTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ProfileTransactorRaw struct {
	Contract *ProfileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfile creates a new instance of Profile, bound to a specific deployed contract.
func NewProfile(address common.Address, backend bind.ContractBackend) (*Profile, error) {
	contract, err := bindProfile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Profile{ProfileCaller: ProfileCaller{contract: contract}, ProfileTransactor: ProfileTransactor{contract: contract}, ProfileFilterer: ProfileFilterer{contract: contract}}, nil
}

// NewProfileCaller creates a new read-only instance of Profile, bound to a specific deployed contract.
func NewProfileCaller(address common.Address, caller bind.ContractCaller) (*ProfileCaller, error) {
	contract, err := bindProfile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileCaller{contract: contract}, nil
}

// NewProfileTransactor creates a new write-only instance of Profile, bound to a specific deployed contract.
func NewProfileTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfileTransactor, error) {
	contract, err := bindProfile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileTransactor{contract: contract}, nil
}

// NewProfileFilterer creates a new log filterer instance of Profile, bound to a specific deployed contract.
func NewProfileFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfileFilterer, error) {
	contract, err := bindProfile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfileFilterer{contract: contract}, nil
}

// bindProfile binds a generic wrapper to an already deployed contract.
func bindProfile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProfileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.ProfileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transact(opts, method, params...)
}

// GetNickName is a free data retrieval call binding the contract method 0xead0327d.
//
// Solidity: function getNickName(address user) view returns(string)
func (_Profile *ProfileCaller) GetNickName(opts *bind.CallOpts, user common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Profile.contract.Call(opts, out, "getNickName", user)
	return *ret0, err
}

// GetNickName is a free data retrieval call binding the contract method 0xead0327d.
//
// Solidity: function getNickName(address user) view returns(string)
func (_Profile *ProfileSession) GetNickName(user common.Address) (string, error) {
	return _Profile.Contract.GetNickName(&_Profile.CallOpts, user)
}

// GetNickName is a free data retrieval call binding the contract method 0xead0327d.
//
// Solidity: function getNickName(address user) view returns(string)
func (_Profile *ProfileCallerSession) GetNickName(user common.Address) (string, error) {
	return _Profile.Contract.GetNickName(&_Profile.CallOpts, user)
}

// GetNickNameList is a free data retrieval call binding the contract method 0xf63ed9d0.
//
// Solidity: function getNickNameList() view returns(string[])
func (_Profile *ProfileCaller) GetNickNameList(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Profile.contract.Call(opts, out, "getNickNameList")
	return *ret0, err
}

// GetNickNameList is a free data retrieval call binding the contract method 0xf63ed9d0.
//
// Solidity: function getNickNameList() view returns(string[])
func (_Profile *ProfileSession) GetNickNameList() ([]string, error) {
	return _Profile.Contract.GetNickNameList(&_Profile.CallOpts)
}

// GetNickNameList is a free data retrieval call binding the contract method 0xf63ed9d0.
//
// Solidity: function getNickNameList() view returns(string[])
func (_Profile *ProfileCallerSession) GetNickNameList() ([]string, error) {
	return _Profile.Contract.GetNickNameList(&_Profile.CallOpts)
}

// GetProfile is a free data retrieval call binding the contract method 0x0f53a470.
//
// Solidity: function getProfile(address user) view returns(string)
func (_Profile *ProfileCaller) GetProfile(opts *bind.CallOpts, user common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Profile.contract.Call(opts, out, "getProfile", user)
	return *ret0, err
}

// GetProfile is a free data retrieval call binding the contract method 0x0f53a470.
//
// Solidity: function getProfile(address user) view returns(string)
func (_Profile *ProfileSession) GetProfile(user common.Address) (string, error) {
	return _Profile.Contract.GetProfile(&_Profile.CallOpts, user)
}

// GetProfile is a free data retrieval call binding the contract method 0x0f53a470.
//
// Solidity: function getProfile(address user) view returns(string)
func (_Profile *ProfileCallerSession) GetProfile(user common.Address) (string, error) {
	return _Profile.Contract.GetProfile(&_Profile.CallOpts, user)
}

// GetStandardNftAddress is a free data retrieval call binding the contract method 0xf8ce6f04.
//
// Solidity: function getStandardNftAddress() view returns(address)
func (_Profile *ProfileCaller) GetStandardNftAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Profile.contract.Call(opts, out, "getStandardNftAddress")
	return *ret0, err
}

// GetStandardNftAddress is a free data retrieval call binding the contract method 0xf8ce6f04.
//
// Solidity: function getStandardNftAddress() view returns(address)
func (_Profile *ProfileSession) GetStandardNftAddress() (common.Address, error) {
	return _Profile.Contract.GetStandardNftAddress(&_Profile.CallOpts)
}

// GetStandardNftAddress is a free data retrieval call binding the contract method 0xf8ce6f04.
//
// Solidity: function getStandardNftAddress() view returns(address)
func (_Profile *ProfileCallerSession) GetStandardNftAddress() (common.Address, error) {
	return _Profile.Contract.GetStandardNftAddress(&_Profile.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Profile *ProfileCaller) NameExt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Profile.contract.Call(opts, out, "nameExt")
	return *ret0, err
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Profile *ProfileSession) NameExt() (string, error) {
	return _Profile.Contract.NameExt(&_Profile.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Profile *ProfileCallerSession) NameExt() (string, error) {
	return _Profile.Contract.NameExt(&_Profile.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Profile *ProfileCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Profile.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Profile *ProfileSession) Ready() (bool, error) {
	return _Profile.Contract.Ready(&_Profile.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Profile *ProfileCallerSession) Ready() (bool, error) {
	return _Profile.Contract.Ready(&_Profile.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Profile *ProfileCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Profile.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Profile *ProfileSession) Registry() (common.Address, error) {
	return _Profile.Contract.Registry(&_Profile.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Profile *ProfileCallerSession) Registry() (common.Address, error) {
	return _Profile.Contract.Registry(&_Profile.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Profile *ProfileTransactor) Initialize(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "initialize", addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Profile *ProfileSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Profile.Contract.Initialize(&_Profile.TransactOpts, addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Profile *ProfileTransactorSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Profile.Contract.Initialize(&_Profile.TransactOpts, addrs)
}

// InjectDependancy is a paid mutator transaction binding the contract method 0x35d07ec4.
//
// Solidity: function injectDependancy(string depName, address addr) returns()
func (_Profile *ProfileTransactor) InjectDependancy(opts *bind.TransactOpts, depName string, addr common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "injectDependancy", depName, addr)
}

// InjectDependancy is a paid mutator transaction binding the contract method 0x35d07ec4.
//
// Solidity: function injectDependancy(string depName, address addr) returns()
func (_Profile *ProfileSession) InjectDependancy(depName string, addr common.Address) (*types.Transaction, error) {
	return _Profile.Contract.InjectDependancy(&_Profile.TransactOpts, depName, addr)
}

// InjectDependancy is a paid mutator transaction binding the contract method 0x35d07ec4.
//
// Solidity: function injectDependancy(string depName, address addr) returns()
func (_Profile *ProfileTransactorSession) InjectDependancy(depName string, addr common.Address) (*types.Transaction, error) {
	return _Profile.Contract.InjectDependancy(&_Profile.TransactOpts, depName, addr)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Profile *ProfileTransactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setForwarder", forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Profile *ProfileSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Profile.Contract.SetForwarder(&_Profile.TransactOpts, forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Profile *ProfileTransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Profile.Contract.SetForwarder(&_Profile.TransactOpts, forwarder)
}

// SetNickname is a paid mutator transaction binding the contract method 0x1c5d9faa.
//
// Solidity: function setNickname(string name) returns()
func (_Profile *ProfileTransactor) SetNickname(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setNickname", name)
}

// SetNickname is a paid mutator transaction binding the contract method 0x1c5d9faa.
//
// Solidity: function setNickname(string name) returns()
func (_Profile *ProfileSession) SetNickname(name string) (*types.Transaction, error) {
	return _Profile.Contract.SetNickname(&_Profile.TransactOpts, name)
}

// SetNickname is a paid mutator transaction binding the contract method 0x1c5d9faa.
//
// Solidity: function setNickname(string name) returns()
func (_Profile *ProfileTransactorSession) SetNickname(name string) (*types.Transaction, error) {
	return _Profile.Contract.SetNickname(&_Profile.TransactOpts, name)
}

// SetProfile is a paid mutator transaction binding the contract method 0xb2ff3913.
//
// Solidity: function setProfile(address user, address nftContract, uint256 tokenId) returns()
func (_Profile *ProfileTransactor) SetProfile(opts *bind.TransactOpts, user common.Address, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setProfile", user, nftContract, tokenId)
}

// SetProfile is a paid mutator transaction binding the contract method 0xb2ff3913.
//
// Solidity: function setProfile(address user, address nftContract, uint256 tokenId) returns()
func (_Profile *ProfileSession) SetProfile(user common.Address, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.SetProfile(&_Profile.TransactOpts, user, nftContract, tokenId)
}

// SetProfile is a paid mutator transaction binding the contract method 0xb2ff3913.
//
// Solidity: function setProfile(address user, address nftContract, uint256 tokenId) returns()
func (_Profile *ProfileTransactorSession) SetProfile(user common.Address, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.SetProfile(&_Profile.TransactOpts, user, nftContract, tokenId)
}

// UpdateNickname is a paid mutator transaction binding the contract method 0xbd2e3b17.
//
// Solidity: function updateNickname(string name) returns()
func (_Profile *ProfileTransactor) UpdateNickname(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "updateNickname", name)
}

// UpdateNickname is a paid mutator transaction binding the contract method 0xbd2e3b17.
//
// Solidity: function updateNickname(string name) returns()
func (_Profile *ProfileSession) UpdateNickname(name string) (*types.Transaction, error) {
	return _Profile.Contract.UpdateNickname(&_Profile.TransactOpts, name)
}

// UpdateNickname is a paid mutator transaction binding the contract method 0xbd2e3b17.
//
// Solidity: function updateNickname(string name) returns()
func (_Profile *ProfileTransactorSession) UpdateNickname(name string) (*types.Transaction, error) {
	return _Profile.Contract.UpdateNickname(&_Profile.TransactOpts, name)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Profile *ProfileTransactor) Upgrade(opts *bind.TransactOpts, newVersion common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "upgrade", newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Profile *ProfileSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Profile.Contract.Upgrade(&_Profile.TransactOpts, newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Profile *ProfileTransactorSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Profile.Contract.Upgrade(&_Profile.TransactOpts, newVersion)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Profile *ProfileTransactor) UpgradeRegistry(opts *bind.TransactOpts, reg common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "upgradeRegistry", reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Profile *ProfileSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Profile.Contract.UpgradeRegistry(&_Profile.TransactOpts, reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Profile *ProfileTransactorSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Profile.Contract.UpgradeRegistry(&_Profile.TransactOpts, reg)
}

// ProfileBeforeUpgradeHookEventIterator is returned from FilterBeforeUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for BeforeUpgradeHookEvent events raised by the Profile contract.
type ProfileBeforeUpgradeHookEventIterator struct {
	Event *ProfileBeforeUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *ProfileBeforeUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileBeforeUpgradeHookEvent)
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
		it.Event = new(ProfileBeforeUpgradeHookEvent)
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
func (it *ProfileBeforeUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileBeforeUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileBeforeUpgradeHookEvent represents a BeforeUpgradeHookEvent event raised by the Profile contract.
type ProfileBeforeUpgradeHookEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBeforeUpgradeHookEvent is a free log retrieval operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Profile *ProfileFilterer) FilterBeforeUpgradeHookEvent(opts *bind.FilterOpts) (*ProfileBeforeUpgradeHookEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileBeforeUpgradeHookEventIterator{contract: _Profile.contract, event: "beforeUpgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchBeforeUpgradeHookEvent is a free log subscription operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Profile *ProfileFilterer) WatchBeforeUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *ProfileBeforeUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileBeforeUpgradeHookEvent)
				if err := _Profile.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseBeforeUpgradeHookEvent(log types.Log) (*ProfileBeforeUpgradeHookEvent, error) {
	event := new(ProfileBeforeUpgradeHookEvent)
	if err := _Profile.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileChangeNicknameEventIterator is returned from FilterChangeNicknameEvent and is used to iterate over the raw logs and unpacked data for ChangeNicknameEvent events raised by the Profile contract.
type ProfileChangeNicknameEventIterator struct {
	Event *ProfileChangeNicknameEvent // Event containing the contract specifics and raw log

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
func (it *ProfileChangeNicknameEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileChangeNicknameEvent)
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
		it.Event = new(ProfileChangeNicknameEvent)
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
func (it *ProfileChangeNicknameEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileChangeNicknameEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileChangeNicknameEvent represents a ChangeNicknameEvent event raised by the Profile contract.
type ProfileChangeNicknameEvent struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChangeNicknameEvent is a free log retrieval operation binding the contract event 0xc658ad3bf83d318cba5be2d3c29e5c7df77f32a706e1b1860ac037be52db8786.
//
// Solidity: event changeNicknameEvent(string name)
func (_Profile *ProfileFilterer) FilterChangeNicknameEvent(opts *bind.FilterOpts) (*ProfileChangeNicknameEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "changeNicknameEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileChangeNicknameEventIterator{contract: _Profile.contract, event: "changeNicknameEvent", logs: logs, sub: sub}, nil
}

// WatchChangeNicknameEvent is a free log subscription operation binding the contract event 0xc658ad3bf83d318cba5be2d3c29e5c7df77f32a706e1b1860ac037be52db8786.
//
// Solidity: event changeNicknameEvent(string name)
func (_Profile *ProfileFilterer) WatchChangeNicknameEvent(opts *bind.WatchOpts, sink chan<- *ProfileChangeNicknameEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "changeNicknameEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileChangeNicknameEvent)
				if err := _Profile.contract.UnpackLog(event, "changeNicknameEvent", log); err != nil {
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

// ParseChangeNicknameEvent is a log parse operation binding the contract event 0xc658ad3bf83d318cba5be2d3c29e5c7df77f32a706e1b1860ac037be52db8786.
//
// Solidity: event changeNicknameEvent(string name)
func (_Profile *ProfileFilterer) ParseChangeNicknameEvent(log types.Log) (*ProfileChangeNicknameEvent, error) {
	event := new(ProfileChangeNicknameEvent)
	if err := _Profile.contract.UnpackLog(event, "changeNicknameEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileCompleteSetupEventIterator is returned from FilterCompleteSetupEvent and is used to iterate over the raw logs and unpacked data for CompleteSetupEvent events raised by the Profile contract.
type ProfileCompleteSetupEventIterator struct {
	Event *ProfileCompleteSetupEvent // Event containing the contract specifics and raw log

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
func (it *ProfileCompleteSetupEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileCompleteSetupEvent)
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
		it.Event = new(ProfileCompleteSetupEvent)
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
func (it *ProfileCompleteSetupEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileCompleteSetupEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileCompleteSetupEvent represents a CompleteSetupEvent event raised by the Profile contract.
type ProfileCompleteSetupEvent struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCompleteSetupEvent is a free log retrieval operation binding the contract event 0xefb63b2500e859546a2f900fe8b455236a475f013dc4d91dfa1ee50ddbd78f61.
//
// Solidity: event completeSetupEvent()
func (_Profile *ProfileFilterer) FilterCompleteSetupEvent(opts *bind.FilterOpts) (*ProfileCompleteSetupEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "completeSetupEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileCompleteSetupEventIterator{contract: _Profile.contract, event: "completeSetupEvent", logs: logs, sub: sub}, nil
}

// WatchCompleteSetupEvent is a free log subscription operation binding the contract event 0xefb63b2500e859546a2f900fe8b455236a475f013dc4d91dfa1ee50ddbd78f61.
//
// Solidity: event completeSetupEvent()
func (_Profile *ProfileFilterer) WatchCompleteSetupEvent(opts *bind.WatchOpts, sink chan<- *ProfileCompleteSetupEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "completeSetupEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileCompleteSetupEvent)
				if err := _Profile.contract.UnpackLog(event, "completeSetupEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseCompleteSetupEvent(log types.Log) (*ProfileCompleteSetupEvent, error) {
	event := new(ProfileCompleteSetupEvent)
	if err := _Profile.contract.UnpackLog(event, "completeSetupEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileEmitActivityEventIterator is returned from FilterEmitActivityEvent and is used to iterate over the raw logs and unpacked data for EmitActivityEvent events raised by the Profile contract.
type ProfileEmitActivityEventIterator struct {
	Event *ProfileEmitActivityEvent // Event containing the contract specifics and raw log

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
func (it *ProfileEmitActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileEmitActivityEvent)
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
		it.Event = new(ProfileEmitActivityEvent)
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
func (it *ProfileEmitActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileEmitActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileEmitActivityEvent represents a EmitActivityEvent event raised by the Profile contract.
type ProfileEmitActivityEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitActivityEvent is a free log retrieval operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Profile *ProfileFilterer) FilterEmitActivityEvent(opts *bind.FilterOpts) (*ProfileEmitActivityEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileEmitActivityEventIterator{contract: _Profile.contract, event: "emitActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitActivityEvent is a free log subscription operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Profile *ProfileFilterer) WatchEmitActivityEvent(opts *bind.WatchOpts, sink chan<- *ProfileEmitActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileEmitActivityEvent)
				if err := _Profile.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseEmitActivityEvent(log types.Log) (*ProfileEmitActivityEvent, error) {
	event := new(ProfileEmitActivityEvent)
	if err := _Profile.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileEmitAppDataEventIterator is returned from FilterEmitAppDataEvent and is used to iterate over the raw logs and unpacked data for EmitAppDataEvent events raised by the Profile contract.
type ProfileEmitAppDataEventIterator struct {
	Event *ProfileEmitAppDataEvent // Event containing the contract specifics and raw log

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
func (it *ProfileEmitAppDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileEmitAppDataEvent)
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
		it.Event = new(ProfileEmitAppDataEvent)
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
func (it *ProfileEmitAppDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileEmitAppDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileEmitAppDataEvent represents a EmitAppDataEvent event raised by the Profile contract.
type ProfileEmitAppDataEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitAppDataEvent is a free log retrieval operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Profile *ProfileFilterer) FilterEmitAppDataEvent(opts *bind.FilterOpts) (*ProfileEmitAppDataEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileEmitAppDataEventIterator{contract: _Profile.contract, event: "emitAppDataEvent", logs: logs, sub: sub}, nil
}

// WatchEmitAppDataEvent is a free log subscription operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Profile *ProfileFilterer) WatchEmitAppDataEvent(opts *bind.WatchOpts, sink chan<- *ProfileEmitAppDataEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileEmitAppDataEvent)
				if err := _Profile.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseEmitAppDataEvent(log types.Log) (*ProfileEmitAppDataEvent, error) {
	event := new(ProfileEmitAppDataEvent)
	if err := _Profile.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileEmitProposalActivityEventIterator is returned from FilterEmitProposalActivityEvent and is used to iterate over the raw logs and unpacked data for EmitProposalActivityEvent events raised by the Profile contract.
type ProfileEmitProposalActivityEventIterator struct {
	Event *ProfileEmitProposalActivityEvent // Event containing the contract specifics and raw log

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
func (it *ProfileEmitProposalActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileEmitProposalActivityEvent)
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
		it.Event = new(ProfileEmitProposalActivityEvent)
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
func (it *ProfileEmitProposalActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileEmitProposalActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileEmitProposalActivityEvent represents a EmitProposalActivityEvent event raised by the Profile contract.
type ProfileEmitProposalActivityEvent struct {
	ActivityType *big.Int
	Proposal     ProposalSummary
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitProposalActivityEvent is a free log retrieval operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Profile *ProfileFilterer) FilterEmitProposalActivityEvent(opts *bind.FilterOpts) (*ProfileEmitProposalActivityEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileEmitProposalActivityEventIterator{contract: _Profile.contract, event: "emitProposalActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitProposalActivityEvent is a free log subscription operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Profile *ProfileFilterer) WatchEmitProposalActivityEvent(opts *bind.WatchOpts, sink chan<- *ProfileEmitProposalActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileEmitProposalActivityEvent)
				if err := _Profile.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseEmitProposalActivityEvent(log types.Log) (*ProfileEmitProposalActivityEvent, error) {
	event := new(ProfileEmitProposalActivityEvent)
	if err := _Profile.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the Profile contract.
type ProfileInitializeEventIterator struct {
	Event *ProfileInitializeEvent // Event containing the contract specifics and raw log

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
func (it *ProfileInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileInitializeEvent)
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
		it.Event = new(ProfileInitializeEvent)
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
func (it *ProfileInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileInitializeEvent represents a InitializeEvent event raised by the Profile contract.
type ProfileInitializeEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Profile *ProfileFilterer) FilterInitializeEvent(opts *bind.FilterOpts) (*ProfileInitializeEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileInitializeEventIterator{contract: _Profile.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Profile *ProfileFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *ProfileInitializeEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileInitializeEvent)
				if err := _Profile.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseInitializeEvent(log types.Log) (*ProfileInitializeEvent, error) {
	event := new(ProfileInitializeEvent)
	if err := _Profile.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileInitializeHookEventIterator is returned from FilterInitializeHookEvent and is used to iterate over the raw logs and unpacked data for InitializeHookEvent events raised by the Profile contract.
type ProfileInitializeHookEventIterator struct {
	Event *ProfileInitializeHookEvent // Event containing the contract specifics and raw log

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
func (it *ProfileInitializeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileInitializeHookEvent)
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
		it.Event = new(ProfileInitializeHookEvent)
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
func (it *ProfileInitializeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileInitializeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileInitializeHookEvent represents a InitializeHookEvent event raised by the Profile contract.
type ProfileInitializeHookEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeHookEvent is a free log retrieval operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Profile *ProfileFilterer) FilterInitializeHookEvent(opts *bind.FilterOpts) (*ProfileInitializeHookEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileInitializeHookEventIterator{contract: _Profile.contract, event: "initializeHookEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeHookEvent is a free log subscription operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Profile *ProfileFilterer) WatchInitializeHookEvent(opts *bind.WatchOpts, sink chan<- *ProfileInitializeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileInitializeHookEvent)
				if err := _Profile.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseInitializeHookEvent(log types.Log) (*ProfileInitializeHookEvent, error) {
	event := new(ProfileInitializeHookEvent)
	if err := _Profile.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileInjectDependancyEventIterator is returned from FilterInjectDependancyEvent and is used to iterate over the raw logs and unpacked data for InjectDependancyEvent events raised by the Profile contract.
type ProfileInjectDependancyEventIterator struct {
	Event *ProfileInjectDependancyEvent // Event containing the contract specifics and raw log

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
func (it *ProfileInjectDependancyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileInjectDependancyEvent)
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
		it.Event = new(ProfileInjectDependancyEvent)
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
func (it *ProfileInjectDependancyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileInjectDependancyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileInjectDependancyEvent represents a InjectDependancyEvent event raised by the Profile contract.
type ProfileInjectDependancyEvent struct {
	Params  DeploymentParameterV1
	DepName string
	Addr    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInjectDependancyEvent is a free log retrieval operation binding the contract event 0xfe4f1a2874ac9abcaeccffebbdb9d6d5cf2192246cf4d1e8dd66e55d6309c493.
//
// Solidity: event injectDependancyEvent((string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, string depName, address addr)
func (_Profile *ProfileFilterer) FilterInjectDependancyEvent(opts *bind.FilterOpts) (*ProfileInjectDependancyEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "injectDependancyEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileInjectDependancyEventIterator{contract: _Profile.contract, event: "injectDependancyEvent", logs: logs, sub: sub}, nil
}

// WatchInjectDependancyEvent is a free log subscription operation binding the contract event 0xfe4f1a2874ac9abcaeccffebbdb9d6d5cf2192246cf4d1e8dd66e55d6309c493.
//
// Solidity: event injectDependancyEvent((string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, string depName, address addr)
func (_Profile *ProfileFilterer) WatchInjectDependancyEvent(opts *bind.WatchOpts, sink chan<- *ProfileInjectDependancyEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "injectDependancyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileInjectDependancyEvent)
				if err := _Profile.contract.UnpackLog(event, "injectDependancyEvent", log); err != nil {
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

// ParseInjectDependancyEvent is a log parse operation binding the contract event 0xfe4f1a2874ac9abcaeccffebbdb9d6d5cf2192246cf4d1e8dd66e55d6309c493.
//
// Solidity: event injectDependancyEvent((string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, string depName, address addr)
func (_Profile *ProfileFilterer) ParseInjectDependancyEvent(log types.Log) (*ProfileInjectDependancyEvent, error) {
	event := new(ProfileInjectDependancyEvent)
	if err := _Profile.contract.UnpackLog(event, "injectDependancyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileSetProfileEventIterator is returned from FilterSetProfileEvent and is used to iterate over the raw logs and unpacked data for SetProfileEvent events raised by the Profile contract.
type ProfileSetProfileEventIterator struct {
	Event *ProfileSetProfileEvent // Event containing the contract specifics and raw log

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
func (it *ProfileSetProfileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetProfileEvent)
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
		it.Event = new(ProfileSetProfileEvent)
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
func (it *ProfileSetProfileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetProfileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetProfileEvent represents a SetProfileEvent event raised by the Profile contract.
type ProfileSetProfileEvent struct {
	User        common.Address
	NftContract common.Address
	TokenId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetProfileEvent is a free log retrieval operation binding the contract event 0x1d53b1c44f59e949550995390c981d8bcad729c99caf9e4b973c5f83ebcc1924.
//
// Solidity: event setProfileEvent(address user, address nftContract, uint256 tokenId)
func (_Profile *ProfileFilterer) FilterSetProfileEvent(opts *bind.FilterOpts) (*ProfileSetProfileEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "setProfileEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileSetProfileEventIterator{contract: _Profile.contract, event: "setProfileEvent", logs: logs, sub: sub}, nil
}

// WatchSetProfileEvent is a free log subscription operation binding the contract event 0x1d53b1c44f59e949550995390c981d8bcad729c99caf9e4b973c5f83ebcc1924.
//
// Solidity: event setProfileEvent(address user, address nftContract, uint256 tokenId)
func (_Profile *ProfileFilterer) WatchSetProfileEvent(opts *bind.WatchOpts, sink chan<- *ProfileSetProfileEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "setProfileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetProfileEvent)
				if err := _Profile.contract.UnpackLog(event, "setProfileEvent", log); err != nil {
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

// ParseSetProfileEvent is a log parse operation binding the contract event 0x1d53b1c44f59e949550995390c981d8bcad729c99caf9e4b973c5f83ebcc1924.
//
// Solidity: event setProfileEvent(address user, address nftContract, uint256 tokenId)
func (_Profile *ProfileFilterer) ParseSetProfileEvent(log types.Log) (*ProfileSetProfileEvent, error) {
	event := new(ProfileSetProfileEvent)
	if err := _Profile.contract.UnpackLog(event, "setProfileEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileSetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the Profile contract.
type ProfileSetReadyEventIterator struct {
	Event *ProfileSetReadyEvent // Event containing the contract specifics and raw log

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
func (it *ProfileSetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetReadyEvent)
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
		it.Event = new(ProfileSetReadyEvent)
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
func (it *ProfileSetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetReadyEvent represents a SetReadyEvent event raised by the Profile contract.
type ProfileSetReadyEvent struct {
	Done bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Profile *ProfileFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*ProfileSetReadyEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileSetReadyEventIterator{contract: _Profile.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Profile *ProfileFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *ProfileSetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetReadyEvent)
				if err := _Profile.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseSetReadyEvent(log types.Log) (*ProfileSetReadyEvent, error) {
	event := new(ProfileSetReadyEvent)
	if err := _Profile.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileUpgradeEventIterator is returned from FilterUpgradeEvent and is used to iterate over the raw logs and unpacked data for UpgradeEvent events raised by the Profile contract.
type ProfileUpgradeEventIterator struct {
	Event *ProfileUpgradeEvent // Event containing the contract specifics and raw log

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
func (it *ProfileUpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileUpgradeEvent)
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
		it.Event = new(ProfileUpgradeEvent)
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
func (it *ProfileUpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileUpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileUpgradeEvent represents a UpgradeEvent event raised by the Profile contract.
type ProfileUpgradeEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpgradeEvent is a free log retrieval operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Profile *ProfileFilterer) FilterUpgradeEvent(opts *bind.FilterOpts) (*ProfileUpgradeEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileUpgradeEventIterator{contract: _Profile.contract, event: "upgradeEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeEvent is a free log subscription operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Profile *ProfileFilterer) WatchUpgradeEvent(opts *bind.WatchOpts, sink chan<- *ProfileUpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileUpgradeEvent)
				if err := _Profile.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseUpgradeEvent(log types.Log) (*ProfileUpgradeEvent, error) {
	event := new(ProfileUpgradeEvent)
	if err := _Profile.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProfileUpgradeRegistryEventIterator is returned from FilterUpgradeRegistryEvent and is used to iterate over the raw logs and unpacked data for UpgradeRegistryEvent events raised by the Profile contract.
type ProfileUpgradeRegistryEventIterator struct {
	Event *ProfileUpgradeRegistryEvent // Event containing the contract specifics and raw log

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
func (it *ProfileUpgradeRegistryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileUpgradeRegistryEvent)
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
		it.Event = new(ProfileUpgradeRegistryEvent)
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
func (it *ProfileUpgradeRegistryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileUpgradeRegistryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileUpgradeRegistryEvent represents a UpgradeRegistryEvent event raised by the Profile contract.
type ProfileUpgradeRegistryEvent struct {
	Reg common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpgradeRegistryEvent is a free log retrieval operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Profile *ProfileFilterer) FilterUpgradeRegistryEvent(opts *bind.FilterOpts) (*ProfileUpgradeRegistryEventIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return &ProfileUpgradeRegistryEventIterator{contract: _Profile.contract, event: "upgradeRegistryEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeRegistryEvent is a free log subscription operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Profile *ProfileFilterer) WatchUpgradeRegistryEvent(opts *bind.WatchOpts, sink chan<- *ProfileUpgradeRegistryEvent) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileUpgradeRegistryEvent)
				if err := _Profile.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseUpgradeRegistryEvent(log types.Log) (*ProfileUpgradeRegistryEvent, error) {
	event := new(ProfileUpgradeRegistryEvent)
	if err := _Profile.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
