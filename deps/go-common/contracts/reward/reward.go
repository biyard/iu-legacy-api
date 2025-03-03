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

// RewardMetaData contains all meta data concerning the Reward contract.
var RewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"stateAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"insertProposalType\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"proposalReward\",\"type\":\"uint256[]\"}],\"name\":\"addProposalTypeRewardFeeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"beforeUpgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"indexed\":false,\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"changeInitialDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"changeUriEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"completeSetupEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitAppDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityType\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structProposalSummary\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"emitProposalActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proposalType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalTypeFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"gtRewardV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"indexed\":false,\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"injectDependancyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"activityType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"sbtRewardEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistryEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"changeUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGtRewardParams\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposalType\",\"type\":\"string\"}],\"name\":\"getProposalTypeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSbtRewardParams\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStandardMTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateAddr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposalType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"proposalTypeFee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"gtReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposalType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"proposalTypeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"gtRewardV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"handleActivityHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"injectDependancy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameExt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"activityType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"sbtReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RewardABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardMetaData.ABI instead.
var RewardABI = RewardMetaData.ABI

// RewardBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const RewardBinRuntime = ``

// Reward is an auto generated Go binding around a Klaytn contract.
type Reward struct {
	RewardCaller     // Read-only binding to the contract
	RewardTransactor // Write-only binding to the contract
	RewardFilterer   // Log filterer for contract events
}

// RewardCaller is an auto generated read-only Go binding around a Klaytn contract.
type RewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardTransactor is an auto generated write-only Go binding around a Klaytn contract.
type RewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type RewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type RewardSession struct {
	Contract     *Reward           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type RewardCallerSession struct {
	Contract *RewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RewardTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type RewardTransactorSession struct {
	Contract     *RewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardRaw is an auto generated low-level Go binding around a Klaytn contract.
type RewardRaw struct {
	Contract *Reward // Generic contract binding to access the raw methods on
}

// RewardCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type RewardCallerRaw struct {
	Contract *RewardCaller // Generic read-only contract binding to access the raw methods on
}

// RewardTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type RewardTransactorRaw struct {
	Contract *RewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReward creates a new instance of Reward, bound to a specific deployed contract.
func NewReward(address common.Address, backend bind.ContractBackend) (*Reward, error) {
	contract, err := bindReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reward{RewardCaller: RewardCaller{contract: contract}, RewardTransactor: RewardTransactor{contract: contract}, RewardFilterer: RewardFilterer{contract: contract}}, nil
}

// NewRewardCaller creates a new read-only instance of Reward, bound to a specific deployed contract.
func NewRewardCaller(address common.Address, caller bind.ContractCaller) (*RewardCaller, error) {
	contract, err := bindReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardCaller{contract: contract}, nil
}

// NewRewardTransactor creates a new write-only instance of Reward, bound to a specific deployed contract.
func NewRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardTransactor, error) {
	contract, err := bindReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardTransactor{contract: contract}, nil
}

// NewRewardFilterer creates a new log filterer instance of Reward, bound to a specific deployed contract.
func NewRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardFilterer, error) {
	contract, err := bindReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardFilterer{contract: contract}, nil
}

// bindReward binds a generic wrapper to an already deployed contract.
func bindReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reward *RewardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reward.Contract.RewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reward *RewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.Contract.RewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reward *RewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reward.Contract.RewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reward *RewardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reward *RewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reward *RewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reward.Contract.contract.Transact(opts, method, params...)
}

// GetGtRewardParams is a free data retrieval call binding the contract method 0xb9a57482.
//
// Solidity: function getGtRewardParams() view returns(uint256[], bool[])
func (_Reward *RewardCaller) GetGtRewardParams(opts *bind.CallOpts) ([]*big.Int, []bool, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Reward.contract.Call(opts, out, "getGtRewardParams")
	return *ret0, *ret1, err
}

// GetGtRewardParams is a free data retrieval call binding the contract method 0xb9a57482.
//
// Solidity: function getGtRewardParams() view returns(uint256[], bool[])
func (_Reward *RewardSession) GetGtRewardParams() ([]*big.Int, []bool, error) {
	return _Reward.Contract.GetGtRewardParams(&_Reward.CallOpts)
}

// GetGtRewardParams is a free data retrieval call binding the contract method 0xb9a57482.
//
// Solidity: function getGtRewardParams() view returns(uint256[], bool[])
func (_Reward *RewardCallerSession) GetGtRewardParams() ([]*big.Int, []bool, error) {
	return _Reward.Contract.GetGtRewardParams(&_Reward.CallOpts)
}

// GetProposalTypeReward is a free data retrieval call binding the contract method 0xa726476e.
//
// Solidity: function getProposalTypeReward(string proposalType) view returns(uint256)
func (_Reward *RewardCaller) GetProposalTypeReward(opts *bind.CallOpts, proposalType string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getProposalTypeReward", proposalType)
	return *ret0, err
}

// GetProposalTypeReward is a free data retrieval call binding the contract method 0xa726476e.
//
// Solidity: function getProposalTypeReward(string proposalType) view returns(uint256)
func (_Reward *RewardSession) GetProposalTypeReward(proposalType string) (*big.Int, error) {
	return _Reward.Contract.GetProposalTypeReward(&_Reward.CallOpts, proposalType)
}

// GetProposalTypeReward is a free data retrieval call binding the contract method 0xa726476e.
//
// Solidity: function getProposalTypeReward(string proposalType) view returns(uint256)
func (_Reward *RewardCallerSession) GetProposalTypeReward(proposalType string) (*big.Int, error) {
	return _Reward.Contract.GetProposalTypeReward(&_Reward.CallOpts, proposalType)
}

// GetSbtRewardParams is a free data retrieval call binding the contract method 0x224869db.
//
// Solidity: function getSbtRewardParams() view returns(uint256[][], bool[])
func (_Reward *RewardCaller) GetSbtRewardParams(opts *bind.CallOpts) ([][]*big.Int, []bool, error) {
	var (
		ret0 = new([][]*big.Int)
		ret1 = new([]bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Reward.contract.Call(opts, out, "getSbtRewardParams")
	return *ret0, *ret1, err
}

// GetSbtRewardParams is a free data retrieval call binding the contract method 0x224869db.
//
// Solidity: function getSbtRewardParams() view returns(uint256[][], bool[])
func (_Reward *RewardSession) GetSbtRewardParams() ([][]*big.Int, []bool, error) {
	return _Reward.Contract.GetSbtRewardParams(&_Reward.CallOpts)
}

// GetSbtRewardParams is a free data retrieval call binding the contract method 0x224869db.
//
// Solidity: function getSbtRewardParams() view returns(uint256[][], bool[])
func (_Reward *RewardCallerSession) GetSbtRewardParams() ([][]*big.Int, []bool, error) {
	return _Reward.Contract.GetSbtRewardParams(&_Reward.CallOpts)
}

// GetStandardMTAddress is a free data retrieval call binding the contract method 0x71e2ee2d.
//
// Solidity: function getStandardMTAddress() view returns(address)
func (_Reward *RewardCaller) GetStandardMTAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getStandardMTAddress")
	return *ret0, err
}

// GetStandardMTAddress is a free data retrieval call binding the contract method 0x71e2ee2d.
//
// Solidity: function getStandardMTAddress() view returns(address)
func (_Reward *RewardSession) GetStandardMTAddress() (common.Address, error) {
	return _Reward.Contract.GetStandardMTAddress(&_Reward.CallOpts)
}

// GetStandardMTAddress is a free data retrieval call binding the contract method 0x71e2ee2d.
//
// Solidity: function getStandardMTAddress() view returns(address)
func (_Reward *RewardCallerSession) GetStandardMTAddress() (common.Address, error) {
	return _Reward.Contract.GetStandardMTAddress(&_Reward.CallOpts)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_Reward *RewardCaller) GetStateAddr(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getStateAddr")
	return *ret0, err
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_Reward *RewardSession) GetStateAddr() ([]common.Address, error) {
	return _Reward.Contract.GetStateAddr(&_Reward.CallOpts)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_Reward *RewardCallerSession) GetStateAddr() ([]common.Address, error) {
	return _Reward.Contract.GetStateAddr(&_Reward.CallOpts)
}

// GetUri is a free data retrieval call binding the contract method 0xda0544aa.
//
// Solidity: function getUri(uint256 id) view returns(string)
func (_Reward *RewardCaller) GetUri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getUri", id)
	return *ret0, err
}

// GetUri is a free data retrieval call binding the contract method 0xda0544aa.
//
// Solidity: function getUri(uint256 id) view returns(string)
func (_Reward *RewardSession) GetUri(id *big.Int) (string, error) {
	return _Reward.Contract.GetUri(&_Reward.CallOpts, id)
}

// GetUri is a free data retrieval call binding the contract method 0xda0544aa.
//
// Solidity: function getUri(uint256 id) view returns(string)
func (_Reward *RewardCallerSession) GetUri(id *big.Int) (string, error) {
	return _Reward.Contract.GetUri(&_Reward.CallOpts, id)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Reward *RewardCaller) NameExt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "nameExt")
	return *ret0, err
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Reward *RewardSession) NameExt() (string, error) {
	return _Reward.Contract.NameExt(&_Reward.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Reward *RewardCallerSession) NameExt() (string, error) {
	return _Reward.Contract.NameExt(&_Reward.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Reward *RewardCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Reward *RewardSession) Ready() (bool, error) {
	return _Reward.Contract.Ready(&_Reward.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Reward *RewardCallerSession) Ready() (bool, error) {
	return _Reward.Contract.Ready(&_Reward.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Reward *RewardCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Reward *RewardSession) Registry() (common.Address, error) {
	return _Reward.Contract.Registry(&_Reward.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Reward *RewardCallerSession) Registry() (common.Address, error) {
	return _Reward.Contract.Registry(&_Reward.CallOpts)
}

// ChangeUri is a paid mutator transaction binding the contract method 0xeb1f9f6e.
//
// Solidity: function changeUri(string uri) returns()
func (_Reward *RewardTransactor) ChangeUri(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "changeUri", uri)
}

// ChangeUri is a paid mutator transaction binding the contract method 0xeb1f9f6e.
//
// Solidity: function changeUri(string uri) returns()
func (_Reward *RewardSession) ChangeUri(uri string) (*types.Transaction, error) {
	return _Reward.Contract.ChangeUri(&_Reward.TransactOpts, uri)
}

// ChangeUri is a paid mutator transaction binding the contract method 0xeb1f9f6e.
//
// Solidity: function changeUri(string uri) returns()
func (_Reward *RewardTransactorSession) ChangeUri(uri string) (*types.Transaction, error) {
	return _Reward.Contract.ChangeUri(&_Reward.TransactOpts, uri)
}

// GtReward is a paid mutator transaction binding the contract method 0x481928c9.
//
// Solidity: function gtReward(string proposalType, uint256 proposalTypeFee, string uuid) returns()
func (_Reward *RewardTransactor) GtReward(opts *bind.TransactOpts, proposalType string, proposalTypeFee *big.Int, uuid string) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "gtReward", proposalType, proposalTypeFee, uuid)
}

// GtReward is a paid mutator transaction binding the contract method 0x481928c9.
//
// Solidity: function gtReward(string proposalType, uint256 proposalTypeFee, string uuid) returns()
func (_Reward *RewardSession) GtReward(proposalType string, proposalTypeFee *big.Int, uuid string) (*types.Transaction, error) {
	return _Reward.Contract.GtReward(&_Reward.TransactOpts, proposalType, proposalTypeFee, uuid)
}

// GtReward is a paid mutator transaction binding the contract method 0x481928c9.
//
// Solidity: function gtReward(string proposalType, uint256 proposalTypeFee, string uuid) returns()
func (_Reward *RewardTransactorSession) GtReward(proposalType string, proposalTypeFee *big.Int, uuid string) (*types.Transaction, error) {
	return _Reward.Contract.GtReward(&_Reward.TransactOpts, proposalType, proposalTypeFee, uuid)
}

// GtRewardV2 is a paid mutator transaction binding the contract method 0xf78cbd61.
//
// Solidity: function gtRewardV2(string proposalType, uint256 proposalTypeFee, uint256 proposalId) returns()
func (_Reward *RewardTransactor) GtRewardV2(opts *bind.TransactOpts, proposalType string, proposalTypeFee *big.Int, proposalId *big.Int) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "gtRewardV2", proposalType, proposalTypeFee, proposalId)
}

// GtRewardV2 is a paid mutator transaction binding the contract method 0xf78cbd61.
//
// Solidity: function gtRewardV2(string proposalType, uint256 proposalTypeFee, uint256 proposalId) returns()
func (_Reward *RewardSession) GtRewardV2(proposalType string, proposalTypeFee *big.Int, proposalId *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.GtRewardV2(&_Reward.TransactOpts, proposalType, proposalTypeFee, proposalId)
}

// GtRewardV2 is a paid mutator transaction binding the contract method 0xf78cbd61.
//
// Solidity: function gtRewardV2(string proposalType, uint256 proposalTypeFee, uint256 proposalId) returns()
func (_Reward *RewardTransactorSession) GtRewardV2(proposalType string, proposalTypeFee *big.Int, proposalId *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.GtRewardV2(&_Reward.TransactOpts, proposalType, proposalTypeFee, proposalId)
}

// HandleActivityHook is a paid mutator transaction binding the contract method 0x9ecd71a7.
//
// Solidity: function handleActivityHook(uint256 activityFlag, bytes data) returns()
func (_Reward *RewardTransactor) HandleActivityHook(opts *bind.TransactOpts, activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "handleActivityHook", activityFlag, data)
}

// HandleActivityHook is a paid mutator transaction binding the contract method 0x9ecd71a7.
//
// Solidity: function handleActivityHook(uint256 activityFlag, bytes data) returns()
func (_Reward *RewardSession) HandleActivityHook(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _Reward.Contract.HandleActivityHook(&_Reward.TransactOpts, activityFlag, data)
}

// HandleActivityHook is a paid mutator transaction binding the contract method 0x9ecd71a7.
//
// Solidity: function handleActivityHook(uint256 activityFlag, bytes data) returns()
func (_Reward *RewardTransactorSession) HandleActivityHook(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _Reward.Contract.HandleActivityHook(&_Reward.TransactOpts, activityFlag, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Reward *RewardTransactor) Initialize(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "initialize", addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Reward *RewardSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Reward.Contract.Initialize(&_Reward.TransactOpts, addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Reward *RewardTransactorSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Reward.Contract.Initialize(&_Reward.TransactOpts, addrs)
}

// InjectDependancy is a paid mutator transaction binding the contract method 0x35d07ec4.
//
// Solidity: function injectDependancy(string depName, address addr) returns()
func (_Reward *RewardTransactor) InjectDependancy(opts *bind.TransactOpts, depName string, addr common.Address) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "injectDependancy", depName, addr)
}

// InjectDependancy is a paid mutator transaction binding the contract method 0x35d07ec4.
//
// Solidity: function injectDependancy(string depName, address addr) returns()
func (_Reward *RewardSession) InjectDependancy(depName string, addr common.Address) (*types.Transaction, error) {
	return _Reward.Contract.InjectDependancy(&_Reward.TransactOpts, depName, addr)
}

// InjectDependancy is a paid mutator transaction binding the contract method 0x35d07ec4.
//
// Solidity: function injectDependancy(string depName, address addr) returns()
func (_Reward *RewardTransactorSession) InjectDependancy(depName string, addr common.Address) (*types.Transaction, error) {
	return _Reward.Contract.InjectDependancy(&_Reward.TransactOpts, depName, addr)
}

// SbtReward is a paid mutator transaction binding the contract method 0x04ff4f7b.
//
// Solidity: function sbtReward(string activityType, address addr) returns()
func (_Reward *RewardTransactor) SbtReward(opts *bind.TransactOpts, activityType string, addr common.Address) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "sbtReward", activityType, addr)
}

// SbtReward is a paid mutator transaction binding the contract method 0x04ff4f7b.
//
// Solidity: function sbtReward(string activityType, address addr) returns()
func (_Reward *RewardSession) SbtReward(activityType string, addr common.Address) (*types.Transaction, error) {
	return _Reward.Contract.SbtReward(&_Reward.TransactOpts, activityType, addr)
}

// SbtReward is a paid mutator transaction binding the contract method 0x04ff4f7b.
//
// Solidity: function sbtReward(string activityType, address addr) returns()
func (_Reward *RewardTransactorSession) SbtReward(activityType string, addr common.Address) (*types.Transaction, error) {
	return _Reward.Contract.SbtReward(&_Reward.TransactOpts, activityType, addr)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Reward *RewardTransactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "setForwarder", forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Reward *RewardSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Reward.Contract.SetForwarder(&_Reward.TransactOpts, forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Reward *RewardTransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Reward.Contract.SetForwarder(&_Reward.TransactOpts, forwarder)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Reward *RewardTransactor) Upgrade(opts *bind.TransactOpts, newVersion common.Address) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "upgrade", newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Reward *RewardSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Reward.Contract.Upgrade(&_Reward.TransactOpts, newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Reward *RewardTransactorSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Reward.Contract.Upgrade(&_Reward.TransactOpts, newVersion)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Reward *RewardTransactor) UpgradeRegistry(opts *bind.TransactOpts, reg common.Address) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "upgradeRegistry", reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Reward *RewardSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Reward.Contract.UpgradeRegistry(&_Reward.TransactOpts, reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Reward *RewardTransactorSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Reward.Contract.UpgradeRegistry(&_Reward.TransactOpts, reg)
}

// RewardAddProposalTypeRewardFeeEventIterator is returned from FilterAddProposalTypeRewardFeeEvent and is used to iterate over the raw logs and unpacked data for AddProposalTypeRewardFeeEvent events raised by the Reward contract.
type RewardAddProposalTypeRewardFeeEventIterator struct {
	Event *RewardAddProposalTypeRewardFeeEvent // Event containing the contract specifics and raw log

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
func (it *RewardAddProposalTypeRewardFeeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardAddProposalTypeRewardFeeEvent)
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
		it.Event = new(RewardAddProposalTypeRewardFeeEvent)
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
func (it *RewardAddProposalTypeRewardFeeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardAddProposalTypeRewardFeeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardAddProposalTypeRewardFeeEvent represents a AddProposalTypeRewardFeeEvent event raised by the Reward contract.
type RewardAddProposalTypeRewardFeeEvent struct {
	InsertProposalType []string
	ProposalReward     []*big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAddProposalTypeRewardFeeEvent is a free log retrieval operation binding the contract event 0x3c987937f1988035912128c9705f07a172a891d7ba8d2680e0cfe8790b8fbed8.
//
// Solidity: event addProposalTypeRewardFeeEvent(string[] insertProposalType, uint256[] proposalReward)
func (_Reward *RewardFilterer) FilterAddProposalTypeRewardFeeEvent(opts *bind.FilterOpts) (*RewardAddProposalTypeRewardFeeEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "addProposalTypeRewardFeeEvent")
	if err != nil {
		return nil, err
	}
	return &RewardAddProposalTypeRewardFeeEventIterator{contract: _Reward.contract, event: "addProposalTypeRewardFeeEvent", logs: logs, sub: sub}, nil
}

// WatchAddProposalTypeRewardFeeEvent is a free log subscription operation binding the contract event 0x3c987937f1988035912128c9705f07a172a891d7ba8d2680e0cfe8790b8fbed8.
//
// Solidity: event addProposalTypeRewardFeeEvent(string[] insertProposalType, uint256[] proposalReward)
func (_Reward *RewardFilterer) WatchAddProposalTypeRewardFeeEvent(opts *bind.WatchOpts, sink chan<- *RewardAddProposalTypeRewardFeeEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "addProposalTypeRewardFeeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardAddProposalTypeRewardFeeEvent)
				if err := _Reward.contract.UnpackLog(event, "addProposalTypeRewardFeeEvent", log); err != nil {
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

// ParseAddProposalTypeRewardFeeEvent is a log parse operation binding the contract event 0x3c987937f1988035912128c9705f07a172a891d7ba8d2680e0cfe8790b8fbed8.
//
// Solidity: event addProposalTypeRewardFeeEvent(string[] insertProposalType, uint256[] proposalReward)
func (_Reward *RewardFilterer) ParseAddProposalTypeRewardFeeEvent(log types.Log) (*RewardAddProposalTypeRewardFeeEvent, error) {
	event := new(RewardAddProposalTypeRewardFeeEvent)
	if err := _Reward.contract.UnpackLog(event, "addProposalTypeRewardFeeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardBeforeUpgradeHookEventIterator is returned from FilterBeforeUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for BeforeUpgradeHookEvent events raised by the Reward contract.
type RewardBeforeUpgradeHookEventIterator struct {
	Event *RewardBeforeUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *RewardBeforeUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardBeforeUpgradeHookEvent)
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
		it.Event = new(RewardBeforeUpgradeHookEvent)
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
func (it *RewardBeforeUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardBeforeUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardBeforeUpgradeHookEvent represents a BeforeUpgradeHookEvent event raised by the Reward contract.
type RewardBeforeUpgradeHookEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBeforeUpgradeHookEvent is a free log retrieval operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Reward *RewardFilterer) FilterBeforeUpgradeHookEvent(opts *bind.FilterOpts) (*RewardBeforeUpgradeHookEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &RewardBeforeUpgradeHookEventIterator{contract: _Reward.contract, event: "beforeUpgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchBeforeUpgradeHookEvent is a free log subscription operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Reward *RewardFilterer) WatchBeforeUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *RewardBeforeUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardBeforeUpgradeHookEvent)
				if err := _Reward.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseBeforeUpgradeHookEvent(log types.Log) (*RewardBeforeUpgradeHookEvent, error) {
	event := new(RewardBeforeUpgradeHookEvent)
	if err := _Reward.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardChangeInitialDataEventIterator is returned from FilterChangeInitialDataEvent and is used to iterate over the raw logs and unpacked data for ChangeInitialDataEvent events raised by the Reward contract.
type RewardChangeInitialDataEventIterator struct {
	Event *RewardChangeInitialDataEvent // Event containing the contract specifics and raw log

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
func (it *RewardChangeInitialDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardChangeInitialDataEvent)
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
		it.Event = new(RewardChangeInitialDataEvent)
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
func (it *RewardChangeInitialDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardChangeInitialDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardChangeInitialDataEvent represents a ChangeInitialDataEvent event raised by the Reward contract.
type RewardChangeInitialDataEvent struct {
	Params DeploymentParameterV1
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChangeInitialDataEvent is a free log retrieval operation binding the contract event 0xad736d523628f64b8c26f8391e499889bef63b9ec3ad1f32f6b5dbfeeba69853.
//
// Solidity: event changeInitialDataEvent((string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params)
func (_Reward *RewardFilterer) FilterChangeInitialDataEvent(opts *bind.FilterOpts) (*RewardChangeInitialDataEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return &RewardChangeInitialDataEventIterator{contract: _Reward.contract, event: "changeInitialDataEvent", logs: logs, sub: sub}, nil
}

// WatchChangeInitialDataEvent is a free log subscription operation binding the contract event 0xad736d523628f64b8c26f8391e499889bef63b9ec3ad1f32f6b5dbfeeba69853.
//
// Solidity: event changeInitialDataEvent((string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params)
func (_Reward *RewardFilterer) WatchChangeInitialDataEvent(opts *bind.WatchOpts, sink chan<- *RewardChangeInitialDataEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardChangeInitialDataEvent)
				if err := _Reward.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
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

// ParseChangeInitialDataEvent is a log parse operation binding the contract event 0xad736d523628f64b8c26f8391e499889bef63b9ec3ad1f32f6b5dbfeeba69853.
//
// Solidity: event changeInitialDataEvent((string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params)
func (_Reward *RewardFilterer) ParseChangeInitialDataEvent(log types.Log) (*RewardChangeInitialDataEvent, error) {
	event := new(RewardChangeInitialDataEvent)
	if err := _Reward.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardChangeUriEventIterator is returned from FilterChangeUriEvent and is used to iterate over the raw logs and unpacked data for ChangeUriEvent events raised by the Reward contract.
type RewardChangeUriEventIterator struct {
	Event *RewardChangeUriEvent // Event containing the contract specifics and raw log

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
func (it *RewardChangeUriEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardChangeUriEvent)
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
		it.Event = new(RewardChangeUriEvent)
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
func (it *RewardChangeUriEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardChangeUriEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardChangeUriEvent represents a ChangeUriEvent event raised by the Reward contract.
type RewardChangeUriEvent struct {
	Uri string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChangeUriEvent is a free log retrieval operation binding the contract event 0xf443af1af0017f68c37a41bb076d5f194366929571bf0d688a2b66653f49ec28.
//
// Solidity: event changeUriEvent(string uri)
func (_Reward *RewardFilterer) FilterChangeUriEvent(opts *bind.FilterOpts) (*RewardChangeUriEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "changeUriEvent")
	if err != nil {
		return nil, err
	}
	return &RewardChangeUriEventIterator{contract: _Reward.contract, event: "changeUriEvent", logs: logs, sub: sub}, nil
}

// WatchChangeUriEvent is a free log subscription operation binding the contract event 0xf443af1af0017f68c37a41bb076d5f194366929571bf0d688a2b66653f49ec28.
//
// Solidity: event changeUriEvent(string uri)
func (_Reward *RewardFilterer) WatchChangeUriEvent(opts *bind.WatchOpts, sink chan<- *RewardChangeUriEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "changeUriEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardChangeUriEvent)
				if err := _Reward.contract.UnpackLog(event, "changeUriEvent", log); err != nil {
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

// ParseChangeUriEvent is a log parse operation binding the contract event 0xf443af1af0017f68c37a41bb076d5f194366929571bf0d688a2b66653f49ec28.
//
// Solidity: event changeUriEvent(string uri)
func (_Reward *RewardFilterer) ParseChangeUriEvent(log types.Log) (*RewardChangeUriEvent, error) {
	event := new(RewardChangeUriEvent)
	if err := _Reward.contract.UnpackLog(event, "changeUriEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardCompleteSetupEventIterator is returned from FilterCompleteSetupEvent and is used to iterate over the raw logs and unpacked data for CompleteSetupEvent events raised by the Reward contract.
type RewardCompleteSetupEventIterator struct {
	Event *RewardCompleteSetupEvent // Event containing the contract specifics and raw log

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
func (it *RewardCompleteSetupEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardCompleteSetupEvent)
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
		it.Event = new(RewardCompleteSetupEvent)
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
func (it *RewardCompleteSetupEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardCompleteSetupEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardCompleteSetupEvent represents a CompleteSetupEvent event raised by the Reward contract.
type RewardCompleteSetupEvent struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCompleteSetupEvent is a free log retrieval operation binding the contract event 0xefb63b2500e859546a2f900fe8b455236a475f013dc4d91dfa1ee50ddbd78f61.
//
// Solidity: event completeSetupEvent()
func (_Reward *RewardFilterer) FilterCompleteSetupEvent(opts *bind.FilterOpts) (*RewardCompleteSetupEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "completeSetupEvent")
	if err != nil {
		return nil, err
	}
	return &RewardCompleteSetupEventIterator{contract: _Reward.contract, event: "completeSetupEvent", logs: logs, sub: sub}, nil
}

// WatchCompleteSetupEvent is a free log subscription operation binding the contract event 0xefb63b2500e859546a2f900fe8b455236a475f013dc4d91dfa1ee50ddbd78f61.
//
// Solidity: event completeSetupEvent()
func (_Reward *RewardFilterer) WatchCompleteSetupEvent(opts *bind.WatchOpts, sink chan<- *RewardCompleteSetupEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "completeSetupEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardCompleteSetupEvent)
				if err := _Reward.contract.UnpackLog(event, "completeSetupEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseCompleteSetupEvent(log types.Log) (*RewardCompleteSetupEvent, error) {
	event := new(RewardCompleteSetupEvent)
	if err := _Reward.contract.UnpackLog(event, "completeSetupEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardEmitActivityEventIterator is returned from FilterEmitActivityEvent and is used to iterate over the raw logs and unpacked data for EmitActivityEvent events raised by the Reward contract.
type RewardEmitActivityEventIterator struct {
	Event *RewardEmitActivityEvent // Event containing the contract specifics and raw log

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
func (it *RewardEmitActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardEmitActivityEvent)
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
		it.Event = new(RewardEmitActivityEvent)
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
func (it *RewardEmitActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardEmitActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardEmitActivityEvent represents a EmitActivityEvent event raised by the Reward contract.
type RewardEmitActivityEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitActivityEvent is a free log retrieval operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Reward *RewardFilterer) FilterEmitActivityEvent(opts *bind.FilterOpts) (*RewardEmitActivityEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return &RewardEmitActivityEventIterator{contract: _Reward.contract, event: "emitActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitActivityEvent is a free log subscription operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Reward *RewardFilterer) WatchEmitActivityEvent(opts *bind.WatchOpts, sink chan<- *RewardEmitActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardEmitActivityEvent)
				if err := _Reward.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseEmitActivityEvent(log types.Log) (*RewardEmitActivityEvent, error) {
	event := new(RewardEmitActivityEvent)
	if err := _Reward.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardEmitAppDataEventIterator is returned from FilterEmitAppDataEvent and is used to iterate over the raw logs and unpacked data for EmitAppDataEvent events raised by the Reward contract.
type RewardEmitAppDataEventIterator struct {
	Event *RewardEmitAppDataEvent // Event containing the contract specifics and raw log

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
func (it *RewardEmitAppDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardEmitAppDataEvent)
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
		it.Event = new(RewardEmitAppDataEvent)
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
func (it *RewardEmitAppDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardEmitAppDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardEmitAppDataEvent represents a EmitAppDataEvent event raised by the Reward contract.
type RewardEmitAppDataEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitAppDataEvent is a free log retrieval operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Reward *RewardFilterer) FilterEmitAppDataEvent(opts *bind.FilterOpts) (*RewardEmitAppDataEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return &RewardEmitAppDataEventIterator{contract: _Reward.contract, event: "emitAppDataEvent", logs: logs, sub: sub}, nil
}

// WatchEmitAppDataEvent is a free log subscription operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Reward *RewardFilterer) WatchEmitAppDataEvent(opts *bind.WatchOpts, sink chan<- *RewardEmitAppDataEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardEmitAppDataEvent)
				if err := _Reward.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseEmitAppDataEvent(log types.Log) (*RewardEmitAppDataEvent, error) {
	event := new(RewardEmitAppDataEvent)
	if err := _Reward.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardEmitProposalActivityEventIterator is returned from FilterEmitProposalActivityEvent and is used to iterate over the raw logs and unpacked data for EmitProposalActivityEvent events raised by the Reward contract.
type RewardEmitProposalActivityEventIterator struct {
	Event *RewardEmitProposalActivityEvent // Event containing the contract specifics and raw log

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
func (it *RewardEmitProposalActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardEmitProposalActivityEvent)
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
		it.Event = new(RewardEmitProposalActivityEvent)
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
func (it *RewardEmitProposalActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardEmitProposalActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardEmitProposalActivityEvent represents a EmitProposalActivityEvent event raised by the Reward contract.
type RewardEmitProposalActivityEvent struct {
	ActivityType *big.Int
	Proposal     ProposalSummary
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitProposalActivityEvent is a free log retrieval operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Reward *RewardFilterer) FilterEmitProposalActivityEvent(opts *bind.FilterOpts) (*RewardEmitProposalActivityEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return &RewardEmitProposalActivityEventIterator{contract: _Reward.contract, event: "emitProposalActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitProposalActivityEvent is a free log subscription operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Reward *RewardFilterer) WatchEmitProposalActivityEvent(opts *bind.WatchOpts, sink chan<- *RewardEmitProposalActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardEmitProposalActivityEvent)
				if err := _Reward.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseEmitProposalActivityEvent(log types.Log) (*RewardEmitProposalActivityEvent, error) {
	event := new(RewardEmitProposalActivityEvent)
	if err := _Reward.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardGtRewardV2EventIterator is returned from FilterGtRewardV2Event and is used to iterate over the raw logs and unpacked data for GtRewardV2Event events raised by the Reward contract.
type RewardGtRewardV2EventIterator struct {
	Event *RewardGtRewardV2Event // Event containing the contract specifics and raw log

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
func (it *RewardGtRewardV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardGtRewardV2Event)
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
		it.Event = new(RewardGtRewardV2Event)
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
func (it *RewardGtRewardV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardGtRewardV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardGtRewardV2Event represents a GtRewardV2Event event raised by the Reward contract.
type RewardGtRewardV2Event struct {
	ProposalType    string
	ProposalTypeFee *big.Int
	ProposalId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterGtRewardV2Event is a free log retrieval operation binding the contract event 0xe4bec885286e0dbf51e646a906affbd638ff54cc8a1c9cc103cfbb74ad21eb1f.
//
// Solidity: event gtRewardV2Event(string proposalType, uint256 proposalTypeFee, uint256 proposalId)
func (_Reward *RewardFilterer) FilterGtRewardV2Event(opts *bind.FilterOpts) (*RewardGtRewardV2EventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "gtRewardV2Event")
	if err != nil {
		return nil, err
	}
	return &RewardGtRewardV2EventIterator{contract: _Reward.contract, event: "gtRewardV2Event", logs: logs, sub: sub}, nil
}

// WatchGtRewardV2Event is a free log subscription operation binding the contract event 0xe4bec885286e0dbf51e646a906affbd638ff54cc8a1c9cc103cfbb74ad21eb1f.
//
// Solidity: event gtRewardV2Event(string proposalType, uint256 proposalTypeFee, uint256 proposalId)
func (_Reward *RewardFilterer) WatchGtRewardV2Event(opts *bind.WatchOpts, sink chan<- *RewardGtRewardV2Event) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "gtRewardV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardGtRewardV2Event)
				if err := _Reward.contract.UnpackLog(event, "gtRewardV2Event", log); err != nil {
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

// ParseGtRewardV2Event is a log parse operation binding the contract event 0xe4bec885286e0dbf51e646a906affbd638ff54cc8a1c9cc103cfbb74ad21eb1f.
//
// Solidity: event gtRewardV2Event(string proposalType, uint256 proposalTypeFee, uint256 proposalId)
func (_Reward *RewardFilterer) ParseGtRewardV2Event(log types.Log) (*RewardGtRewardV2Event, error) {
	event := new(RewardGtRewardV2Event)
	if err := _Reward.contract.UnpackLog(event, "gtRewardV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the Reward contract.
type RewardInitializeEventIterator struct {
	Event *RewardInitializeEvent // Event containing the contract specifics and raw log

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
func (it *RewardInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardInitializeEvent)
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
		it.Event = new(RewardInitializeEvent)
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
func (it *RewardInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardInitializeEvent represents a InitializeEvent event raised by the Reward contract.
type RewardInitializeEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Reward *RewardFilterer) FilterInitializeEvent(opts *bind.FilterOpts) (*RewardInitializeEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &RewardInitializeEventIterator{contract: _Reward.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Reward *RewardFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *RewardInitializeEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardInitializeEvent)
				if err := _Reward.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseInitializeEvent(log types.Log) (*RewardInitializeEvent, error) {
	event := new(RewardInitializeEvent)
	if err := _Reward.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardInitializeHookEventIterator is returned from FilterInitializeHookEvent and is used to iterate over the raw logs and unpacked data for InitializeHookEvent events raised by the Reward contract.
type RewardInitializeHookEventIterator struct {
	Event *RewardInitializeHookEvent // Event containing the contract specifics and raw log

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
func (it *RewardInitializeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardInitializeHookEvent)
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
		it.Event = new(RewardInitializeHookEvent)
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
func (it *RewardInitializeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardInitializeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardInitializeHookEvent represents a InitializeHookEvent event raised by the Reward contract.
type RewardInitializeHookEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeHookEvent is a free log retrieval operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Reward *RewardFilterer) FilterInitializeHookEvent(opts *bind.FilterOpts) (*RewardInitializeHookEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return &RewardInitializeHookEventIterator{contract: _Reward.contract, event: "initializeHookEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeHookEvent is a free log subscription operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Reward *RewardFilterer) WatchInitializeHookEvent(opts *bind.WatchOpts, sink chan<- *RewardInitializeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardInitializeHookEvent)
				if err := _Reward.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseInitializeHookEvent(log types.Log) (*RewardInitializeHookEvent, error) {
	event := new(RewardInitializeHookEvent)
	if err := _Reward.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardInjectDependancyEventIterator is returned from FilterInjectDependancyEvent and is used to iterate over the raw logs and unpacked data for InjectDependancyEvent events raised by the Reward contract.
type RewardInjectDependancyEventIterator struct {
	Event *RewardInjectDependancyEvent // Event containing the contract specifics and raw log

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
func (it *RewardInjectDependancyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardInjectDependancyEvent)
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
		it.Event = new(RewardInjectDependancyEvent)
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
func (it *RewardInjectDependancyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardInjectDependancyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardInjectDependancyEvent represents a InjectDependancyEvent event raised by the Reward contract.
type RewardInjectDependancyEvent struct {
	Params  DeploymentParameterV1
	DepName string
	Addr    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInjectDependancyEvent is a free log retrieval operation binding the contract event 0xfe4f1a2874ac9abcaeccffebbdb9d6d5cf2192246cf4d1e8dd66e55d6309c493.
//
// Solidity: event injectDependancyEvent((string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, string depName, address addr)
func (_Reward *RewardFilterer) FilterInjectDependancyEvent(opts *bind.FilterOpts) (*RewardInjectDependancyEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "injectDependancyEvent")
	if err != nil {
		return nil, err
	}
	return &RewardInjectDependancyEventIterator{contract: _Reward.contract, event: "injectDependancyEvent", logs: logs, sub: sub}, nil
}

// WatchInjectDependancyEvent is a free log subscription operation binding the contract event 0xfe4f1a2874ac9abcaeccffebbdb9d6d5cf2192246cf4d1e8dd66e55d6309c493.
//
// Solidity: event injectDependancyEvent((string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, string depName, address addr)
func (_Reward *RewardFilterer) WatchInjectDependancyEvent(opts *bind.WatchOpts, sink chan<- *RewardInjectDependancyEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "injectDependancyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardInjectDependancyEvent)
				if err := _Reward.contract.UnpackLog(event, "injectDependancyEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseInjectDependancyEvent(log types.Log) (*RewardInjectDependancyEvent, error) {
	event := new(RewardInjectDependancyEvent)
	if err := _Reward.contract.UnpackLog(event, "injectDependancyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardSbtRewardEventIterator is returned from FilterSbtRewardEvent and is used to iterate over the raw logs and unpacked data for SbtRewardEvent events raised by the Reward contract.
type RewardSbtRewardEventIterator struct {
	Event *RewardSbtRewardEvent // Event containing the contract specifics and raw log

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
func (it *RewardSbtRewardEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardSbtRewardEvent)
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
		it.Event = new(RewardSbtRewardEvent)
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
func (it *RewardSbtRewardEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardSbtRewardEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardSbtRewardEvent represents a SbtRewardEvent event raised by the Reward contract.
type RewardSbtRewardEvent struct {
	ActivityType string
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSbtRewardEvent is a free log retrieval operation binding the contract event 0x1b25c502e76978296f37f39ac27fa1b273c8f6f1353e35c84d08c428c5b8747b.
//
// Solidity: event sbtRewardEvent(string activityType, address addr)
func (_Reward *RewardFilterer) FilterSbtRewardEvent(opts *bind.FilterOpts) (*RewardSbtRewardEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "sbtRewardEvent")
	if err != nil {
		return nil, err
	}
	return &RewardSbtRewardEventIterator{contract: _Reward.contract, event: "sbtRewardEvent", logs: logs, sub: sub}, nil
}

// WatchSbtRewardEvent is a free log subscription operation binding the contract event 0x1b25c502e76978296f37f39ac27fa1b273c8f6f1353e35c84d08c428c5b8747b.
//
// Solidity: event sbtRewardEvent(string activityType, address addr)
func (_Reward *RewardFilterer) WatchSbtRewardEvent(opts *bind.WatchOpts, sink chan<- *RewardSbtRewardEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "sbtRewardEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardSbtRewardEvent)
				if err := _Reward.contract.UnpackLog(event, "sbtRewardEvent", log); err != nil {
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

// ParseSbtRewardEvent is a log parse operation binding the contract event 0x1b25c502e76978296f37f39ac27fa1b273c8f6f1353e35c84d08c428c5b8747b.
//
// Solidity: event sbtRewardEvent(string activityType, address addr)
func (_Reward *RewardFilterer) ParseSbtRewardEvent(log types.Log) (*RewardSbtRewardEvent, error) {
	event := new(RewardSbtRewardEvent)
	if err := _Reward.contract.UnpackLog(event, "sbtRewardEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardSetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the Reward contract.
type RewardSetReadyEventIterator struct {
	Event *RewardSetReadyEvent // Event containing the contract specifics and raw log

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
func (it *RewardSetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardSetReadyEvent)
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
		it.Event = new(RewardSetReadyEvent)
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
func (it *RewardSetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardSetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardSetReadyEvent represents a SetReadyEvent event raised by the Reward contract.
type RewardSetReadyEvent struct {
	Done bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Reward *RewardFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*RewardSetReadyEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &RewardSetReadyEventIterator{contract: _Reward.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Reward *RewardFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *RewardSetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardSetReadyEvent)
				if err := _Reward.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseSetReadyEvent(log types.Log) (*RewardSetReadyEvent, error) {
	event := new(RewardSetReadyEvent)
	if err := _Reward.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardUpgradeEventIterator is returned from FilterUpgradeEvent and is used to iterate over the raw logs and unpacked data for UpgradeEvent events raised by the Reward contract.
type RewardUpgradeEventIterator struct {
	Event *RewardUpgradeEvent // Event containing the contract specifics and raw log

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
func (it *RewardUpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardUpgradeEvent)
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
		it.Event = new(RewardUpgradeEvent)
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
func (it *RewardUpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardUpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardUpgradeEvent represents a UpgradeEvent event raised by the Reward contract.
type RewardUpgradeEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpgradeEvent is a free log retrieval operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Reward *RewardFilterer) FilterUpgradeEvent(opts *bind.FilterOpts) (*RewardUpgradeEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return &RewardUpgradeEventIterator{contract: _Reward.contract, event: "upgradeEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeEvent is a free log subscription operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Reward *RewardFilterer) WatchUpgradeEvent(opts *bind.WatchOpts, sink chan<- *RewardUpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardUpgradeEvent)
				if err := _Reward.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseUpgradeEvent(log types.Log) (*RewardUpgradeEvent, error) {
	event := new(RewardUpgradeEvent)
	if err := _Reward.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardUpgradeRegistryEventIterator is returned from FilterUpgradeRegistryEvent and is used to iterate over the raw logs and unpacked data for UpgradeRegistryEvent events raised by the Reward contract.
type RewardUpgradeRegistryEventIterator struct {
	Event *RewardUpgradeRegistryEvent // Event containing the contract specifics and raw log

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
func (it *RewardUpgradeRegistryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardUpgradeRegistryEvent)
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
		it.Event = new(RewardUpgradeRegistryEvent)
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
func (it *RewardUpgradeRegistryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardUpgradeRegistryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardUpgradeRegistryEvent represents a UpgradeRegistryEvent event raised by the Reward contract.
type RewardUpgradeRegistryEvent struct {
	Reg common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpgradeRegistryEvent is a free log retrieval operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Reward *RewardFilterer) FilterUpgradeRegistryEvent(opts *bind.FilterOpts) (*RewardUpgradeRegistryEventIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return &RewardUpgradeRegistryEventIterator{contract: _Reward.contract, event: "upgradeRegistryEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeRegistryEvent is a free log subscription operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Reward *RewardFilterer) WatchUpgradeRegistryEvent(opts *bind.WatchOpts, sink chan<- *RewardUpgradeRegistryEvent) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardUpgradeRegistryEvent)
				if err := _Reward.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
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
func (_Reward *RewardFilterer) ParseUpgradeRegistryEvent(log types.Log) (*RewardUpgradeRegistryEvent, error) {
	event := new(RewardUpgradeRegistryEvent)
	if err := _Reward.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
