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

// MemberMetaData contains all meta data concerning the Member contract.
var MemberMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"stateAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"beforeUpgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"roleNames\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"rolePermissions\",\"type\":\"uint256[]\"}],\"name\":\"changeInitialDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"name\":\"changeMemberToGroupEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitAppDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityType\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structProposalSummary\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"emitProposalActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perm\",\"type\":\"uint256\"}],\"name\":\"grantPermissionsToGroupEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistryEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"groupNames\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"perms\",\"type\":\"uint256[]\"}],\"name\":\"addGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"name\":\"addMembersToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"groupNames\",\"type\":\"string[]\"}],\"name\":\"changeMemberBulkGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"name\":\"changeMemberGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"getBulkGroup\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getGroup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"perm\",\"type\":\"uint256\"}],\"name\":\"grantPermissionsToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"handleActivityHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"name\":\"hasGroupPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"perm\",\"type\":\"uint256\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listGroups\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"groupNames\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"permissions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"members\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"groups\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameExt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MemberABI is the input ABI used to generate the binding from.
// Deprecated: Use MemberMetaData.ABI instead.
var MemberABI = MemberMetaData.ABI

// MemberBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const MemberBinRuntime = ``

// Member is an auto generated Go binding around a Klaytn contract.
type Member struct {
	MemberCaller     // Read-only binding to the contract
	MemberTransactor // Write-only binding to the contract
	MemberFilterer   // Log filterer for contract events
}

// MemberCaller is an auto generated read-only Go binding around a Klaytn contract.
type MemberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemberTransactor is an auto generated write-only Go binding around a Klaytn contract.
type MemberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemberFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type MemberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemberSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type MemberSession struct {
	Contract     *Member           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemberCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type MemberCallerSession struct {
	Contract *MemberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MemberTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type MemberTransactorSession struct {
	Contract     *MemberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemberRaw is an auto generated low-level Go binding around a Klaytn contract.
type MemberRaw struct {
	Contract *Member // Generic contract binding to access the raw methods on
}

// MemberCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type MemberCallerRaw struct {
	Contract *MemberCaller // Generic read-only contract binding to access the raw methods on
}

// MemberTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type MemberTransactorRaw struct {
	Contract *MemberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMember creates a new instance of Member, bound to a specific deployed contract.
func NewMember(address common.Address, backend bind.ContractBackend) (*Member, error) {
	contract, err := bindMember(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Member{MemberCaller: MemberCaller{contract: contract}, MemberTransactor: MemberTransactor{contract: contract}, MemberFilterer: MemberFilterer{contract: contract}}, nil
}

// NewMemberCaller creates a new read-only instance of Member, bound to a specific deployed contract.
func NewMemberCaller(address common.Address, caller bind.ContractCaller) (*MemberCaller, error) {
	contract, err := bindMember(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemberCaller{contract: contract}, nil
}

// NewMemberTransactor creates a new write-only instance of Member, bound to a specific deployed contract.
func NewMemberTransactor(address common.Address, transactor bind.ContractTransactor) (*MemberTransactor, error) {
	contract, err := bindMember(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemberTransactor{contract: contract}, nil
}

// NewMemberFilterer creates a new log filterer instance of Member, bound to a specific deployed contract.
func NewMemberFilterer(address common.Address, filterer bind.ContractFilterer) (*MemberFilterer, error) {
	contract, err := bindMember(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemberFilterer{contract: contract}, nil
}

// bindMember binds a generic wrapper to an already deployed contract.
func bindMember(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MemberMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Member *MemberRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Member.Contract.MemberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Member *MemberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Member.Contract.MemberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Member *MemberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Member.Contract.MemberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Member *MemberCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Member.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Member *MemberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Member.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Member *MemberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Member.Contract.contract.Transact(opts, method, params...)
}

// GetBulkGroup is a free data retrieval call binding the contract method 0x1cb458e9.
//
// Solidity: function getBulkGroup(address[] users) view returns(string[])
func (_Member *MemberCaller) GetBulkGroup(opts *bind.CallOpts, users []common.Address) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Member.contract.Call(opts, out, "getBulkGroup", users)
	return *ret0, err
}

// GetBulkGroup is a free data retrieval call binding the contract method 0x1cb458e9.
//
// Solidity: function getBulkGroup(address[] users) view returns(string[])
func (_Member *MemberSession) GetBulkGroup(users []common.Address) ([]string, error) {
	return _Member.Contract.GetBulkGroup(&_Member.CallOpts, users)
}

// GetBulkGroup is a free data retrieval call binding the contract method 0x1cb458e9.
//
// Solidity: function getBulkGroup(address[] users) view returns(string[])
func (_Member *MemberCallerSession) GetBulkGroup(users []common.Address) ([]string, error) {
	return _Member.Contract.GetBulkGroup(&_Member.CallOpts, users)
}

// GetGroup is a free data retrieval call binding the contract method 0x53adce21.
//
// Solidity: function getGroup(address user) view returns(string)
func (_Member *MemberCaller) GetGroup(opts *bind.CallOpts, user common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Member.contract.Call(opts, out, "getGroup", user)
	return *ret0, err
}

// GetGroup is a free data retrieval call binding the contract method 0x53adce21.
//
// Solidity: function getGroup(address user) view returns(string)
func (_Member *MemberSession) GetGroup(user common.Address) (string, error) {
	return _Member.Contract.GetGroup(&_Member.CallOpts, user)
}

// GetGroup is a free data retrieval call binding the contract method 0x53adce21.
//
// Solidity: function getGroup(address user) view returns(string)
func (_Member *MemberCallerSession) GetGroup(user common.Address) (string, error) {
	return _Member.Contract.GetGroup(&_Member.CallOpts, user)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address)
func (_Member *MemberCaller) GetStateAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Member.contract.Call(opts, out, "getStateAddr")
	return *ret0, err
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address)
func (_Member *MemberSession) GetStateAddr() (common.Address, error) {
	return _Member.Contract.GetStateAddr(&_Member.CallOpts)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address)
func (_Member *MemberCallerSession) GetStateAddr() (common.Address, error) {
	return _Member.Contract.GetStateAddr(&_Member.CallOpts)
}

// HasGroupPermission is a free data retrieval call binding the contract method 0x6f4694bb.
//
// Solidity: function hasGroupPermission(address addr, string groupName) view returns(bool)
func (_Member *MemberCaller) HasGroupPermission(opts *bind.CallOpts, addr common.Address, groupName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Member.contract.Call(opts, out, "hasGroupPermission", addr, groupName)
	return *ret0, err
}

// HasGroupPermission is a free data retrieval call binding the contract method 0x6f4694bb.
//
// Solidity: function hasGroupPermission(address addr, string groupName) view returns(bool)
func (_Member *MemberSession) HasGroupPermission(addr common.Address, groupName string) (bool, error) {
	return _Member.Contract.HasGroupPermission(&_Member.CallOpts, addr, groupName)
}

// HasGroupPermission is a free data retrieval call binding the contract method 0x6f4694bb.
//
// Solidity: function hasGroupPermission(address addr, string groupName) view returns(bool)
func (_Member *MemberCallerSession) HasGroupPermission(addr common.Address, groupName string) (bool, error) {
	return _Member.Contract.HasGroupPermission(&_Member.CallOpts, addr, groupName)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d944fd9.
//
// Solidity: function hasPermission(address addr, uint256 perm) view returns(bool)
func (_Member *MemberCaller) HasPermission(opts *bind.CallOpts, addr common.Address, perm *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Member.contract.Call(opts, out, "hasPermission", addr, perm)
	return *ret0, err
}

// HasPermission is a free data retrieval call binding the contract method 0x8d944fd9.
//
// Solidity: function hasPermission(address addr, uint256 perm) view returns(bool)
func (_Member *MemberSession) HasPermission(addr common.Address, perm *big.Int) (bool, error) {
	return _Member.Contract.HasPermission(&_Member.CallOpts, addr, perm)
}

// HasPermission is a free data retrieval call binding the contract method 0x8d944fd9.
//
// Solidity: function hasPermission(address addr, uint256 perm) view returns(bool)
func (_Member *MemberCallerSession) HasPermission(addr common.Address, perm *big.Int) (bool, error) {
	return _Member.Contract.HasPermission(&_Member.CallOpts, addr, perm)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address user) view returns(bool)
func (_Member *MemberCaller) IsMember(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Member.contract.Call(opts, out, "isMember", user)
	return *ret0, err
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address user) view returns(bool)
func (_Member *MemberSession) IsMember(user common.Address) (bool, error) {
	return _Member.Contract.IsMember(&_Member.CallOpts, user)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address user) view returns(bool)
func (_Member *MemberCallerSession) IsMember(user common.Address) (bool, error) {
	return _Member.Contract.IsMember(&_Member.CallOpts, user)
}

// ListGroups is a free data retrieval call binding the contract method 0x9be35a4b.
//
// Solidity: function listGroups() view returns(string[] groupNames, uint256[] permissions)
func (_Member *MemberCaller) ListGroups(opts *bind.CallOpts) (struct {
	GroupNames  []string
	Permissions []*big.Int
}, error) {
	ret := new(struct {
		GroupNames  []string
		Permissions []*big.Int
	})
	out := ret
	err := _Member.contract.Call(opts, out, "listGroups")
	return *ret, err
}

// ListGroups is a free data retrieval call binding the contract method 0x9be35a4b.
//
// Solidity: function listGroups() view returns(string[] groupNames, uint256[] permissions)
func (_Member *MemberSession) ListGroups() (struct {
	GroupNames  []string
	Permissions []*big.Int
}, error) {
	return _Member.Contract.ListGroups(&_Member.CallOpts)
}

// ListGroups is a free data retrieval call binding the contract method 0x9be35a4b.
//
// Solidity: function listGroups() view returns(string[] groupNames, uint256[] permissions)
func (_Member *MemberCallerSession) ListGroups() (struct {
	GroupNames  []string
	Permissions []*big.Int
}, error) {
	return _Member.Contract.ListGroups(&_Member.CallOpts)
}

// ListMembers is a free data retrieval call binding the contract method 0xb6afd2ca.
//
// Solidity: function listMembers() view returns(address[] members, string[] groups)
func (_Member *MemberCaller) ListMembers(opts *bind.CallOpts) (struct {
	Members []common.Address
	Groups  []string
}, error) {
	ret := new(struct {
		Members []common.Address
		Groups  []string
	})
	out := ret
	err := _Member.contract.Call(opts, out, "listMembers")
	return *ret, err
}

// ListMembers is a free data retrieval call binding the contract method 0xb6afd2ca.
//
// Solidity: function listMembers() view returns(address[] members, string[] groups)
func (_Member *MemberSession) ListMembers() (struct {
	Members []common.Address
	Groups  []string
}, error) {
	return _Member.Contract.ListMembers(&_Member.CallOpts)
}

// ListMembers is a free data retrieval call binding the contract method 0xb6afd2ca.
//
// Solidity: function listMembers() view returns(address[] members, string[] groups)
func (_Member *MemberCallerSession) ListMembers() (struct {
	Members []common.Address
	Groups  []string
}, error) {
	return _Member.Contract.ListMembers(&_Member.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Member *MemberCaller) NameExt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Member.contract.Call(opts, out, "nameExt")
	return *ret0, err
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Member *MemberSession) NameExt() (string, error) {
	return _Member.Contract.NameExt(&_Member.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Member *MemberCallerSession) NameExt() (string, error) {
	return _Member.Contract.NameExt(&_Member.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Member *MemberCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Member.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Member *MemberSession) Ready() (bool, error) {
	return _Member.Contract.Ready(&_Member.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Member *MemberCallerSession) Ready() (bool, error) {
	return _Member.Contract.Ready(&_Member.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Member *MemberCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Member.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Member *MemberSession) Registry() (common.Address, error) {
	return _Member.Contract.Registry(&_Member.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Member *MemberCallerSession) Registry() (common.Address, error) {
	return _Member.Contract.Registry(&_Member.CallOpts)
}

// AddGroup is a paid mutator transaction binding the contract method 0x9e65a8a2.
//
// Solidity: function addGroup(string[] groupNames, uint256[] perms) returns()
func (_Member *MemberTransactor) AddGroup(opts *bind.TransactOpts, groupNames []string, perms []*big.Int) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "addGroup", groupNames, perms)
}

// AddGroup is a paid mutator transaction binding the contract method 0x9e65a8a2.
//
// Solidity: function addGroup(string[] groupNames, uint256[] perms) returns()
func (_Member *MemberSession) AddGroup(groupNames []string, perms []*big.Int) (*types.Transaction, error) {
	return _Member.Contract.AddGroup(&_Member.TransactOpts, groupNames, perms)
}

// AddGroup is a paid mutator transaction binding the contract method 0x9e65a8a2.
//
// Solidity: function addGroup(string[] groupNames, uint256[] perms) returns()
func (_Member *MemberTransactorSession) AddGroup(groupNames []string, perms []*big.Int) (*types.Transaction, error) {
	return _Member.Contract.AddGroup(&_Member.TransactOpts, groupNames, perms)
}

// AddMembersToGroup is a paid mutator transaction binding the contract method 0x57cd53ff.
//
// Solidity: function addMembersToGroup(address[] addrs, string groupName) returns()
func (_Member *MemberTransactor) AddMembersToGroup(opts *bind.TransactOpts, addrs []common.Address, groupName string) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "addMembersToGroup", addrs, groupName)
}

// AddMembersToGroup is a paid mutator transaction binding the contract method 0x57cd53ff.
//
// Solidity: function addMembersToGroup(address[] addrs, string groupName) returns()
func (_Member *MemberSession) AddMembersToGroup(addrs []common.Address, groupName string) (*types.Transaction, error) {
	return _Member.Contract.AddMembersToGroup(&_Member.TransactOpts, addrs, groupName)
}

// AddMembersToGroup is a paid mutator transaction binding the contract method 0x57cd53ff.
//
// Solidity: function addMembersToGroup(address[] addrs, string groupName) returns()
func (_Member *MemberTransactorSession) AddMembersToGroup(addrs []common.Address, groupName string) (*types.Transaction, error) {
	return _Member.Contract.AddMembersToGroup(&_Member.TransactOpts, addrs, groupName)
}

// ChangeMemberBulkGroup is a paid mutator transaction binding the contract method 0x07f85ee9.
//
// Solidity: function changeMemberBulkGroup(address[] addrs, string[] groupNames) returns()
func (_Member *MemberTransactor) ChangeMemberBulkGroup(opts *bind.TransactOpts, addrs []common.Address, groupNames []string) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "changeMemberBulkGroup", addrs, groupNames)
}

// ChangeMemberBulkGroup is a paid mutator transaction binding the contract method 0x07f85ee9.
//
// Solidity: function changeMemberBulkGroup(address[] addrs, string[] groupNames) returns()
func (_Member *MemberSession) ChangeMemberBulkGroup(addrs []common.Address, groupNames []string) (*types.Transaction, error) {
	return _Member.Contract.ChangeMemberBulkGroup(&_Member.TransactOpts, addrs, groupNames)
}

// ChangeMemberBulkGroup is a paid mutator transaction binding the contract method 0x07f85ee9.
//
// Solidity: function changeMemberBulkGroup(address[] addrs, string[] groupNames) returns()
func (_Member *MemberTransactorSession) ChangeMemberBulkGroup(addrs []common.Address, groupNames []string) (*types.Transaction, error) {
	return _Member.Contract.ChangeMemberBulkGroup(&_Member.TransactOpts, addrs, groupNames)
}

// ChangeMemberGroup is a paid mutator transaction binding the contract method 0xf45ce0dd.
//
// Solidity: function changeMemberGroup(address addr, string groupName) returns()
func (_Member *MemberTransactor) ChangeMemberGroup(opts *bind.TransactOpts, addr common.Address, groupName string) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "changeMemberGroup", addr, groupName)
}

// ChangeMemberGroup is a paid mutator transaction binding the contract method 0xf45ce0dd.
//
// Solidity: function changeMemberGroup(address addr, string groupName) returns()
func (_Member *MemberSession) ChangeMemberGroup(addr common.Address, groupName string) (*types.Transaction, error) {
	return _Member.Contract.ChangeMemberGroup(&_Member.TransactOpts, addr, groupName)
}

// ChangeMemberGroup is a paid mutator transaction binding the contract method 0xf45ce0dd.
//
// Solidity: function changeMemberGroup(address addr, string groupName) returns()
func (_Member *MemberTransactorSession) ChangeMemberGroup(addr common.Address, groupName string) (*types.Transaction, error) {
	return _Member.Contract.ChangeMemberGroup(&_Member.TransactOpts, addr, groupName)
}

// GrantPermissionsToGroup is a paid mutator transaction binding the contract method 0x3f6aec21.
//
// Solidity: function grantPermissionsToGroup(string groupName, uint256 perm) returns()
func (_Member *MemberTransactor) GrantPermissionsToGroup(opts *bind.TransactOpts, groupName string, perm *big.Int) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "grantPermissionsToGroup", groupName, perm)
}

// GrantPermissionsToGroup is a paid mutator transaction binding the contract method 0x3f6aec21.
//
// Solidity: function grantPermissionsToGroup(string groupName, uint256 perm) returns()
func (_Member *MemberSession) GrantPermissionsToGroup(groupName string, perm *big.Int) (*types.Transaction, error) {
	return _Member.Contract.GrantPermissionsToGroup(&_Member.TransactOpts, groupName, perm)
}

// GrantPermissionsToGroup is a paid mutator transaction binding the contract method 0x3f6aec21.
//
// Solidity: function grantPermissionsToGroup(string groupName, uint256 perm) returns()
func (_Member *MemberTransactorSession) GrantPermissionsToGroup(groupName string, perm *big.Int) (*types.Transaction, error) {
	return _Member.Contract.GrantPermissionsToGroup(&_Member.TransactOpts, groupName, perm)
}

// HandleActivityHook is a paid mutator transaction binding the contract method 0x9ecd71a7.
//
// Solidity: function handleActivityHook(uint256 activityFlag, bytes data) returns()
func (_Member *MemberTransactor) HandleActivityHook(opts *bind.TransactOpts, activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "handleActivityHook", activityFlag, data)
}

// HandleActivityHook is a paid mutator transaction binding the contract method 0x9ecd71a7.
//
// Solidity: function handleActivityHook(uint256 activityFlag, bytes data) returns()
func (_Member *MemberSession) HandleActivityHook(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _Member.Contract.HandleActivityHook(&_Member.TransactOpts, activityFlag, data)
}

// HandleActivityHook is a paid mutator transaction binding the contract method 0x9ecd71a7.
//
// Solidity: function handleActivityHook(uint256 activityFlag, bytes data) returns()
func (_Member *MemberTransactorSession) HandleActivityHook(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _Member.Contract.HandleActivityHook(&_Member.TransactOpts, activityFlag, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Member *MemberTransactor) Initialize(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "initialize", addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Member *MemberSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Member.Contract.Initialize(&_Member.TransactOpts, addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Member *MemberTransactorSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Member.Contract.Initialize(&_Member.TransactOpts, addrs)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Member *MemberTransactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "setForwarder", forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Member *MemberSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Member.Contract.SetForwarder(&_Member.TransactOpts, forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Member *MemberTransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Member.Contract.SetForwarder(&_Member.TransactOpts, forwarder)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Member *MemberTransactor) Upgrade(opts *bind.TransactOpts, newVersion common.Address) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "upgrade", newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Member *MemberSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Member.Contract.Upgrade(&_Member.TransactOpts, newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Member *MemberTransactorSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Member.Contract.Upgrade(&_Member.TransactOpts, newVersion)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Member *MemberTransactor) UpgradeRegistry(opts *bind.TransactOpts, reg common.Address) (*types.Transaction, error) {
	return _Member.contract.Transact(opts, "upgradeRegistry", reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Member *MemberSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Member.Contract.UpgradeRegistry(&_Member.TransactOpts, reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Member *MemberTransactorSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Member.Contract.UpgradeRegistry(&_Member.TransactOpts, reg)
}

// MemberBeforeUpgradeHookEventIterator is returned from FilterBeforeUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for BeforeUpgradeHookEvent events raised by the Member contract.
type MemberBeforeUpgradeHookEventIterator struct {
	Event *MemberBeforeUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *MemberBeforeUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberBeforeUpgradeHookEvent)
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
		it.Event = new(MemberBeforeUpgradeHookEvent)
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
func (it *MemberBeforeUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberBeforeUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberBeforeUpgradeHookEvent represents a BeforeUpgradeHookEvent event raised by the Member contract.
type MemberBeforeUpgradeHookEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBeforeUpgradeHookEvent is a free log retrieval operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Member *MemberFilterer) FilterBeforeUpgradeHookEvent(opts *bind.FilterOpts) (*MemberBeforeUpgradeHookEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &MemberBeforeUpgradeHookEventIterator{contract: _Member.contract, event: "beforeUpgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchBeforeUpgradeHookEvent is a free log subscription operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Member *MemberFilterer) WatchBeforeUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *MemberBeforeUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberBeforeUpgradeHookEvent)
				if err := _Member.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
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
func (_Member *MemberFilterer) ParseBeforeUpgradeHookEvent(log types.Log) (*MemberBeforeUpgradeHookEvent, error) {
	event := new(MemberBeforeUpgradeHookEvent)
	if err := _Member.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberChangeInitialDataEventIterator is returned from FilterChangeInitialDataEvent and is used to iterate over the raw logs and unpacked data for ChangeInitialDataEvent events raised by the Member contract.
type MemberChangeInitialDataEventIterator struct {
	Event *MemberChangeInitialDataEvent // Event containing the contract specifics and raw log

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
func (it *MemberChangeInitialDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberChangeInitialDataEvent)
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
		it.Event = new(MemberChangeInitialDataEvent)
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
func (it *MemberChangeInitialDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberChangeInitialDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberChangeInitialDataEvent represents a ChangeInitialDataEvent event raised by the Member contract.
type MemberChangeInitialDataEvent struct {
	RoleNames       []string
	RolePermissions []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChangeInitialDataEvent is a free log retrieval operation binding the contract event 0x522f55e057acece2e4e947f4c718020faf35ece752f5271a76a531a64f10a2c2.
//
// Solidity: event changeInitialDataEvent(string[] roleNames, uint256[] rolePermissions)
func (_Member *MemberFilterer) FilterChangeInitialDataEvent(opts *bind.FilterOpts) (*MemberChangeInitialDataEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return &MemberChangeInitialDataEventIterator{contract: _Member.contract, event: "changeInitialDataEvent", logs: logs, sub: sub}, nil
}

// WatchChangeInitialDataEvent is a free log subscription operation binding the contract event 0x522f55e057acece2e4e947f4c718020faf35ece752f5271a76a531a64f10a2c2.
//
// Solidity: event changeInitialDataEvent(string[] roleNames, uint256[] rolePermissions)
func (_Member *MemberFilterer) WatchChangeInitialDataEvent(opts *bind.WatchOpts, sink chan<- *MemberChangeInitialDataEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberChangeInitialDataEvent)
				if err := _Member.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
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

// ParseChangeInitialDataEvent is a log parse operation binding the contract event 0x522f55e057acece2e4e947f4c718020faf35ece752f5271a76a531a64f10a2c2.
//
// Solidity: event changeInitialDataEvent(string[] roleNames, uint256[] rolePermissions)
func (_Member *MemberFilterer) ParseChangeInitialDataEvent(log types.Log) (*MemberChangeInitialDataEvent, error) {
	event := new(MemberChangeInitialDataEvent)
	if err := _Member.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberChangeMemberToGroupEventIterator is returned from FilterChangeMemberToGroupEvent and is used to iterate over the raw logs and unpacked data for ChangeMemberToGroupEvent events raised by the Member contract.
type MemberChangeMemberToGroupEventIterator struct {
	Event *MemberChangeMemberToGroupEvent // Event containing the contract specifics and raw log

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
func (it *MemberChangeMemberToGroupEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberChangeMemberToGroupEvent)
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
		it.Event = new(MemberChangeMemberToGroupEvent)
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
func (it *MemberChangeMemberToGroupEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberChangeMemberToGroupEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberChangeMemberToGroupEvent represents a ChangeMemberToGroupEvent event raised by the Member contract.
type MemberChangeMemberToGroupEvent struct {
	Addr      common.Address
	GroupName string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangeMemberToGroupEvent is a free log retrieval operation binding the contract event 0xd07efed75b2c76d8ad977c3601c6cfd7c63144ad6c4cfb647df28df352e9d9a1.
//
// Solidity: event changeMemberToGroupEvent(address addr, string groupName)
func (_Member *MemberFilterer) FilterChangeMemberToGroupEvent(opts *bind.FilterOpts) (*MemberChangeMemberToGroupEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "changeMemberToGroupEvent")
	if err != nil {
		return nil, err
	}
	return &MemberChangeMemberToGroupEventIterator{contract: _Member.contract, event: "changeMemberToGroupEvent", logs: logs, sub: sub}, nil
}

// WatchChangeMemberToGroupEvent is a free log subscription operation binding the contract event 0xd07efed75b2c76d8ad977c3601c6cfd7c63144ad6c4cfb647df28df352e9d9a1.
//
// Solidity: event changeMemberToGroupEvent(address addr, string groupName)
func (_Member *MemberFilterer) WatchChangeMemberToGroupEvent(opts *bind.WatchOpts, sink chan<- *MemberChangeMemberToGroupEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "changeMemberToGroupEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberChangeMemberToGroupEvent)
				if err := _Member.contract.UnpackLog(event, "changeMemberToGroupEvent", log); err != nil {
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

// ParseChangeMemberToGroupEvent is a log parse operation binding the contract event 0xd07efed75b2c76d8ad977c3601c6cfd7c63144ad6c4cfb647df28df352e9d9a1.
//
// Solidity: event changeMemberToGroupEvent(address addr, string groupName)
func (_Member *MemberFilterer) ParseChangeMemberToGroupEvent(log types.Log) (*MemberChangeMemberToGroupEvent, error) {
	event := new(MemberChangeMemberToGroupEvent)
	if err := _Member.contract.UnpackLog(event, "changeMemberToGroupEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberEmitActivityEventIterator is returned from FilterEmitActivityEvent and is used to iterate over the raw logs and unpacked data for EmitActivityEvent events raised by the Member contract.
type MemberEmitActivityEventIterator struct {
	Event *MemberEmitActivityEvent // Event containing the contract specifics and raw log

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
func (it *MemberEmitActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberEmitActivityEvent)
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
		it.Event = new(MemberEmitActivityEvent)
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
func (it *MemberEmitActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberEmitActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberEmitActivityEvent represents a EmitActivityEvent event raised by the Member contract.
type MemberEmitActivityEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitActivityEvent is a free log retrieval operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Member *MemberFilterer) FilterEmitActivityEvent(opts *bind.FilterOpts) (*MemberEmitActivityEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return &MemberEmitActivityEventIterator{contract: _Member.contract, event: "emitActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitActivityEvent is a free log subscription operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Member *MemberFilterer) WatchEmitActivityEvent(opts *bind.WatchOpts, sink chan<- *MemberEmitActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberEmitActivityEvent)
				if err := _Member.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
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
func (_Member *MemberFilterer) ParseEmitActivityEvent(log types.Log) (*MemberEmitActivityEvent, error) {
	event := new(MemberEmitActivityEvent)
	if err := _Member.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberEmitAppDataEventIterator is returned from FilterEmitAppDataEvent and is used to iterate over the raw logs and unpacked data for EmitAppDataEvent events raised by the Member contract.
type MemberEmitAppDataEventIterator struct {
	Event *MemberEmitAppDataEvent // Event containing the contract specifics and raw log

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
func (it *MemberEmitAppDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberEmitAppDataEvent)
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
		it.Event = new(MemberEmitAppDataEvent)
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
func (it *MemberEmitAppDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberEmitAppDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberEmitAppDataEvent represents a EmitAppDataEvent event raised by the Member contract.
type MemberEmitAppDataEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitAppDataEvent is a free log retrieval operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Member *MemberFilterer) FilterEmitAppDataEvent(opts *bind.FilterOpts) (*MemberEmitAppDataEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return &MemberEmitAppDataEventIterator{contract: _Member.contract, event: "emitAppDataEvent", logs: logs, sub: sub}, nil
}

// WatchEmitAppDataEvent is a free log subscription operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Member *MemberFilterer) WatchEmitAppDataEvent(opts *bind.WatchOpts, sink chan<- *MemberEmitAppDataEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberEmitAppDataEvent)
				if err := _Member.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
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
func (_Member *MemberFilterer) ParseEmitAppDataEvent(log types.Log) (*MemberEmitAppDataEvent, error) {
	event := new(MemberEmitAppDataEvent)
	if err := _Member.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberEmitProposalActivityEventIterator is returned from FilterEmitProposalActivityEvent and is used to iterate over the raw logs and unpacked data for EmitProposalActivityEvent events raised by the Member contract.
type MemberEmitProposalActivityEventIterator struct {
	Event *MemberEmitProposalActivityEvent // Event containing the contract specifics and raw log

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
func (it *MemberEmitProposalActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberEmitProposalActivityEvent)
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
		it.Event = new(MemberEmitProposalActivityEvent)
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
func (it *MemberEmitProposalActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberEmitProposalActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberEmitProposalActivityEvent represents a EmitProposalActivityEvent event raised by the Member contract.
type MemberEmitProposalActivityEvent struct {
	ActivityType *big.Int
	Proposal     ProposalSummary
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitProposalActivityEvent is a free log retrieval operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Member *MemberFilterer) FilterEmitProposalActivityEvent(opts *bind.FilterOpts) (*MemberEmitProposalActivityEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return &MemberEmitProposalActivityEventIterator{contract: _Member.contract, event: "emitProposalActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitProposalActivityEvent is a free log subscription operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Member *MemberFilterer) WatchEmitProposalActivityEvent(opts *bind.WatchOpts, sink chan<- *MemberEmitProposalActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberEmitProposalActivityEvent)
				if err := _Member.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
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
func (_Member *MemberFilterer) ParseEmitProposalActivityEvent(log types.Log) (*MemberEmitProposalActivityEvent, error) {
	event := new(MemberEmitProposalActivityEvent)
	if err := _Member.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberGrantPermissionsToGroupEventIterator is returned from FilterGrantPermissionsToGroupEvent and is used to iterate over the raw logs and unpacked data for GrantPermissionsToGroupEvent events raised by the Member contract.
type MemberGrantPermissionsToGroupEventIterator struct {
	Event *MemberGrantPermissionsToGroupEvent // Event containing the contract specifics and raw log

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
func (it *MemberGrantPermissionsToGroupEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberGrantPermissionsToGroupEvent)
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
		it.Event = new(MemberGrantPermissionsToGroupEvent)
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
func (it *MemberGrantPermissionsToGroupEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberGrantPermissionsToGroupEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberGrantPermissionsToGroupEvent represents a GrantPermissionsToGroupEvent event raised by the Member contract.
type MemberGrantPermissionsToGroupEvent struct {
	GroupName string
	Perm      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGrantPermissionsToGroupEvent is a free log retrieval operation binding the contract event 0xf4e4761ed5c30bfc9ee24f3a7d20107f81fadbdbf1c8187f7b67d9c0057f4213.
//
// Solidity: event grantPermissionsToGroupEvent(string groupName, uint256 perm)
func (_Member *MemberFilterer) FilterGrantPermissionsToGroupEvent(opts *bind.FilterOpts) (*MemberGrantPermissionsToGroupEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "grantPermissionsToGroupEvent")
	if err != nil {
		return nil, err
	}
	return &MemberGrantPermissionsToGroupEventIterator{contract: _Member.contract, event: "grantPermissionsToGroupEvent", logs: logs, sub: sub}, nil
}

// WatchGrantPermissionsToGroupEvent is a free log subscription operation binding the contract event 0xf4e4761ed5c30bfc9ee24f3a7d20107f81fadbdbf1c8187f7b67d9c0057f4213.
//
// Solidity: event grantPermissionsToGroupEvent(string groupName, uint256 perm)
func (_Member *MemberFilterer) WatchGrantPermissionsToGroupEvent(opts *bind.WatchOpts, sink chan<- *MemberGrantPermissionsToGroupEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "grantPermissionsToGroupEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberGrantPermissionsToGroupEvent)
				if err := _Member.contract.UnpackLog(event, "grantPermissionsToGroupEvent", log); err != nil {
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

// ParseGrantPermissionsToGroupEvent is a log parse operation binding the contract event 0xf4e4761ed5c30bfc9ee24f3a7d20107f81fadbdbf1c8187f7b67d9c0057f4213.
//
// Solidity: event grantPermissionsToGroupEvent(string groupName, uint256 perm)
func (_Member *MemberFilterer) ParseGrantPermissionsToGroupEvent(log types.Log) (*MemberGrantPermissionsToGroupEvent, error) {
	event := new(MemberGrantPermissionsToGroupEvent)
	if err := _Member.contract.UnpackLog(event, "grantPermissionsToGroupEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the Member contract.
type MemberInitializeEventIterator struct {
	Event *MemberInitializeEvent // Event containing the contract specifics and raw log

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
func (it *MemberInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberInitializeEvent)
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
		it.Event = new(MemberInitializeEvent)
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
func (it *MemberInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberInitializeEvent represents a InitializeEvent event raised by the Member contract.
type MemberInitializeEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Member *MemberFilterer) FilterInitializeEvent(opts *bind.FilterOpts) (*MemberInitializeEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &MemberInitializeEventIterator{contract: _Member.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Member *MemberFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *MemberInitializeEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberInitializeEvent)
				if err := _Member.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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
func (_Member *MemberFilterer) ParseInitializeEvent(log types.Log) (*MemberInitializeEvent, error) {
	event := new(MemberInitializeEvent)
	if err := _Member.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberInitializeHookEventIterator is returned from FilterInitializeHookEvent and is used to iterate over the raw logs and unpacked data for InitializeHookEvent events raised by the Member contract.
type MemberInitializeHookEventIterator struct {
	Event *MemberInitializeHookEvent // Event containing the contract specifics and raw log

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
func (it *MemberInitializeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberInitializeHookEvent)
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
		it.Event = new(MemberInitializeHookEvent)
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
func (it *MemberInitializeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberInitializeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberInitializeHookEvent represents a InitializeHookEvent event raised by the Member contract.
type MemberInitializeHookEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeHookEvent is a free log retrieval operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Member *MemberFilterer) FilterInitializeHookEvent(opts *bind.FilterOpts) (*MemberInitializeHookEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return &MemberInitializeHookEventIterator{contract: _Member.contract, event: "initializeHookEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeHookEvent is a free log subscription operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Member *MemberFilterer) WatchInitializeHookEvent(opts *bind.WatchOpts, sink chan<- *MemberInitializeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberInitializeHookEvent)
				if err := _Member.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
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
func (_Member *MemberFilterer) ParseInitializeHookEvent(log types.Log) (*MemberInitializeHookEvent, error) {
	event := new(MemberInitializeHookEvent)
	if err := _Member.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberSetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the Member contract.
type MemberSetReadyEventIterator struct {
	Event *MemberSetReadyEvent // Event containing the contract specifics and raw log

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
func (it *MemberSetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberSetReadyEvent)
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
		it.Event = new(MemberSetReadyEvent)
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
func (it *MemberSetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberSetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberSetReadyEvent represents a SetReadyEvent event raised by the Member contract.
type MemberSetReadyEvent struct {
	Done bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Member *MemberFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*MemberSetReadyEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &MemberSetReadyEventIterator{contract: _Member.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Member *MemberFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *MemberSetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberSetReadyEvent)
				if err := _Member.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
func (_Member *MemberFilterer) ParseSetReadyEvent(log types.Log) (*MemberSetReadyEvent, error) {
	event := new(MemberSetReadyEvent)
	if err := _Member.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberUpgradeEventIterator is returned from FilterUpgradeEvent and is used to iterate over the raw logs and unpacked data for UpgradeEvent events raised by the Member contract.
type MemberUpgradeEventIterator struct {
	Event *MemberUpgradeEvent // Event containing the contract specifics and raw log

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
func (it *MemberUpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberUpgradeEvent)
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
		it.Event = new(MemberUpgradeEvent)
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
func (it *MemberUpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberUpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberUpgradeEvent represents a UpgradeEvent event raised by the Member contract.
type MemberUpgradeEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpgradeEvent is a free log retrieval operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Member *MemberFilterer) FilterUpgradeEvent(opts *bind.FilterOpts) (*MemberUpgradeEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return &MemberUpgradeEventIterator{contract: _Member.contract, event: "upgradeEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeEvent is a free log subscription operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Member *MemberFilterer) WatchUpgradeEvent(opts *bind.WatchOpts, sink chan<- *MemberUpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberUpgradeEvent)
				if err := _Member.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
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
func (_Member *MemberFilterer) ParseUpgradeEvent(log types.Log) (*MemberUpgradeEvent, error) {
	event := new(MemberUpgradeEvent)
	if err := _Member.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MemberUpgradeRegistryEventIterator is returned from FilterUpgradeRegistryEvent and is used to iterate over the raw logs and unpacked data for UpgradeRegistryEvent events raised by the Member contract.
type MemberUpgradeRegistryEventIterator struct {
	Event *MemberUpgradeRegistryEvent // Event containing the contract specifics and raw log

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
func (it *MemberUpgradeRegistryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemberUpgradeRegistryEvent)
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
		it.Event = new(MemberUpgradeRegistryEvent)
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
func (it *MemberUpgradeRegistryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemberUpgradeRegistryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemberUpgradeRegistryEvent represents a UpgradeRegistryEvent event raised by the Member contract.
type MemberUpgradeRegistryEvent struct {
	Reg common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpgradeRegistryEvent is a free log retrieval operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Member *MemberFilterer) FilterUpgradeRegistryEvent(opts *bind.FilterOpts) (*MemberUpgradeRegistryEventIterator, error) {

	logs, sub, err := _Member.contract.FilterLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return &MemberUpgradeRegistryEventIterator{contract: _Member.contract, event: "upgradeRegistryEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeRegistryEvent is a free log subscription operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Member *MemberFilterer) WatchUpgradeRegistryEvent(opts *bind.WatchOpts, sink chan<- *MemberUpgradeRegistryEvent) (event.Subscription, error) {

	logs, sub, err := _Member.contract.WatchLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemberUpgradeRegistryEvent)
				if err := _Member.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
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
func (_Member *MemberFilterer) ParseUpgradeRegistryEvent(log types.Log) (*MemberUpgradeRegistryEvent, error) {
	event := new(MemberUpgradeRegistryEvent)
	if err := _Member.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
