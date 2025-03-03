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

// IDataProposalParameter is an auto generated low-level Go binding around an user-defined struct.
type IDataProposalParameter struct {
	ProposalType   uint8
	Proposer       common.Address
	Contents       string
	DidFinished    bool
	DidPass        bool
	SubmittedAt    *big.Int
	FinishedAt     *big.Int
	AcceptedPowers *big.Int
	RejectedPowers *big.Int
	FinalizeVp     *big.Int
}

// IPolicyDataProposedPoicy is an auto generated low-level Go binding around an user-defined struct.
type IPolicyDataProposedPoicy struct {
	Name  string
	Value *big.Int
}

// IPolicyDataV2ProposedPoicy is an auto generated low-level Go binding around an user-defined struct.
type IPolicyDataV2ProposedPoicy struct {
	Name  string
	Value *big.Int
}

// ProposalParameterV2 is an auto generated low-level Go binding around an user-defined struct.
type ProposalParameterV2 struct {
	ProposalType   uint8
	Proposer       common.Address
	Contents       string
	DidFinished    bool
	DidPass        bool
	SubmittedAt    *big.Int
	FinishedAt     *big.Int
	AcceptedPowers *big.Int
	RejectedPowers *big.Int
	FinalizeVp     *big.Int
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

// ProposalManagerAppV2MetaData contains all meta data concerning the ProposalManagerAppV2 contract.
var ProposalManagerAppV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"stateAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"beforeUpgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enableSeeVote\",\"type\":\"bool\"}],\"name\":\"changeInitialDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitAppDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityType\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structProposalSummary\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"emitProposalActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistryEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposalContractAddress\",\"type\":\"address\"}],\"name\":\"contractListProposals\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposalContractAddress\",\"type\":\"address\"}],\"name\":\"contractListProposalsV2\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getAccepters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getAcceptersV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getComment\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getCommentV2\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getPagenationProposals\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getParticipants\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getParticipantsV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolicies\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"}],\"name\":\"getPolicy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIData.ProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structIData.ProposalParameter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"getProposalStructure\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIData.ProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structIData.ProposalParameter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStructureV2\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameterV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalV2\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameterV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getProposedPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structIPolicyData.ProposedPoicy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposedPolicyV2\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structIPolicyDataV2.ProposedPoicy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getRejecters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getRejectersV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateAddr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalProposals\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getVoterValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getVoterValueV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameExt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProposalManagerAppV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalManagerAppV2MetaData.ABI instead.
var ProposalManagerAppV2ABI = ProposalManagerAppV2MetaData.ABI

// ProposalManagerAppV2BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ProposalManagerAppV2BinRuntime = ``

// ProposalManagerAppV2 is an auto generated Go binding around a Klaytn contract.
type ProposalManagerAppV2 struct {
	ProposalManagerAppV2Caller     // Read-only binding to the contract
	ProposalManagerAppV2Transactor // Write-only binding to the contract
	ProposalManagerAppV2Filterer   // Log filterer for contract events
}

// ProposalManagerAppV2Caller is an auto generated read-only Go binding around a Klaytn contract.
type ProposalManagerAppV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalManagerAppV2Transactor is an auto generated write-only Go binding around a Klaytn contract.
type ProposalManagerAppV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalManagerAppV2Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ProposalManagerAppV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalManagerAppV2Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ProposalManagerAppV2Session struct {
	Contract     *ProposalManagerAppV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProposalManagerAppV2CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ProposalManagerAppV2CallerSession struct {
	Contract *ProposalManagerAppV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ProposalManagerAppV2TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ProposalManagerAppV2TransactorSession struct {
	Contract     *ProposalManagerAppV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ProposalManagerAppV2Raw is an auto generated low-level Go binding around a Klaytn contract.
type ProposalManagerAppV2Raw struct {
	Contract *ProposalManagerAppV2 // Generic contract binding to access the raw methods on
}

// ProposalManagerAppV2CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ProposalManagerAppV2CallerRaw struct {
	Contract *ProposalManagerAppV2Caller // Generic read-only contract binding to access the raw methods on
}

// ProposalManagerAppV2TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ProposalManagerAppV2TransactorRaw struct {
	Contract *ProposalManagerAppV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewProposalManagerAppV2 creates a new instance of ProposalManagerAppV2, bound to a specific deployed contract.
func NewProposalManagerAppV2(address common.Address, backend bind.ContractBackend) (*ProposalManagerAppV2, error) {
	contract, err := bindProposalManagerAppV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2{ProposalManagerAppV2Caller: ProposalManagerAppV2Caller{contract: contract}, ProposalManagerAppV2Transactor: ProposalManagerAppV2Transactor{contract: contract}, ProposalManagerAppV2Filterer: ProposalManagerAppV2Filterer{contract: contract}}, nil
}

// NewProposalManagerAppV2Caller creates a new read-only instance of ProposalManagerAppV2, bound to a specific deployed contract.
func NewProposalManagerAppV2Caller(address common.Address, caller bind.ContractCaller) (*ProposalManagerAppV2Caller, error) {
	contract, err := bindProposalManagerAppV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2Caller{contract: contract}, nil
}

// NewProposalManagerAppV2Transactor creates a new write-only instance of ProposalManagerAppV2, bound to a specific deployed contract.
func NewProposalManagerAppV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ProposalManagerAppV2Transactor, error) {
	contract, err := bindProposalManagerAppV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2Transactor{contract: contract}, nil
}

// NewProposalManagerAppV2Filterer creates a new log filterer instance of ProposalManagerAppV2, bound to a specific deployed contract.
func NewProposalManagerAppV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ProposalManagerAppV2Filterer, error) {
	contract, err := bindProposalManagerAppV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2Filterer{contract: contract}, nil
}

// bindProposalManagerAppV2 binds a generic wrapper to an already deployed contract.
func bindProposalManagerAppV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProposalManagerAppV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalManagerAppV2 *ProposalManagerAppV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProposalManagerAppV2.Contract.ProposalManagerAppV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalManagerAppV2 *ProposalManagerAppV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.ProposalManagerAppV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalManagerAppV2 *ProposalManagerAppV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.ProposalManagerAppV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProposalManagerAppV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalManagerAppV2 *ProposalManagerAppV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalManagerAppV2 *ProposalManagerAppV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.contract.Transact(opts, method, params...)
}

// ContractListProposals is a free data retrieval call binding the contract method 0x2ece9311.
//
// Solidity: function contractListProposals(address proposalContractAddress) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) ContractListProposals(opts *bind.CallOpts, proposalContractAddress common.Address) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "contractListProposals", proposalContractAddress)
	return *ret0, err
}

// ContractListProposals is a free data retrieval call binding the contract method 0x2ece9311.
//
// Solidity: function contractListProposals(address proposalContractAddress) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) ContractListProposals(proposalContractAddress common.Address) ([]string, error) {
	return _ProposalManagerAppV2.Contract.ContractListProposals(&_ProposalManagerAppV2.CallOpts, proposalContractAddress)
}

// ContractListProposals is a free data retrieval call binding the contract method 0x2ece9311.
//
// Solidity: function contractListProposals(address proposalContractAddress) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) ContractListProposals(proposalContractAddress common.Address) ([]string, error) {
	return _ProposalManagerAppV2.Contract.ContractListProposals(&_ProposalManagerAppV2.CallOpts, proposalContractAddress)
}

// ContractListProposalsV2 is a free data retrieval call binding the contract method 0xd22beafc.
//
// Solidity: function contractListProposalsV2(address proposalContractAddress) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) ContractListProposalsV2(opts *bind.CallOpts, proposalContractAddress common.Address) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "contractListProposalsV2", proposalContractAddress)
	return *ret0, err
}

// ContractListProposalsV2 is a free data retrieval call binding the contract method 0xd22beafc.
//
// Solidity: function contractListProposalsV2(address proposalContractAddress) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) ContractListProposalsV2(proposalContractAddress common.Address) ([]string, error) {
	return _ProposalManagerAppV2.Contract.ContractListProposalsV2(&_ProposalManagerAppV2.CallOpts, proposalContractAddress)
}

// ContractListProposalsV2 is a free data retrieval call binding the contract method 0xd22beafc.
//
// Solidity: function contractListProposalsV2(address proposalContractAddress) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) ContractListProposalsV2(proposalContractAddress common.Address) ([]string, error) {
	return _ProposalManagerAppV2.Contract.ContractListProposalsV2(&_ProposalManagerAppV2.CallOpts, proposalContractAddress)
}

// GetAccepters is a free data retrieval call binding the contract method 0x1d774ba8.
//
// Solidity: function getAccepters(string uuid_) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetAccepters(opts *bind.CallOpts, uuid_ string) (bool, []common.Address, error) {
	var (
		ret0 = new(bool)
		ret1 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getAccepters", uuid_)
	return *ret0, *ret1, err
}

// GetAccepters is a free data retrieval call binding the contract method 0x1d774ba8.
//
// Solidity: function getAccepters(string uuid_) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetAccepters(uuid_ string) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetAccepters(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetAccepters is a free data retrieval call binding the contract method 0x1d774ba8.
//
// Solidity: function getAccepters(string uuid_) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetAccepters(uuid_ string) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetAccepters(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetAcceptersV2 is a free data retrieval call binding the contract method 0x95f27b7b.
//
// Solidity: function getAcceptersV2(uint256 proposalId) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetAcceptersV2(opts *bind.CallOpts, proposalId *big.Int) (bool, []common.Address, error) {
	var (
		ret0 = new(bool)
		ret1 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getAcceptersV2", proposalId)
	return *ret0, *ret1, err
}

// GetAcceptersV2 is a free data retrieval call binding the contract method 0x95f27b7b.
//
// Solidity: function getAcceptersV2(uint256 proposalId) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetAcceptersV2(proposalId *big.Int) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetAcceptersV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetAcceptersV2 is a free data retrieval call binding the contract method 0x95f27b7b.
//
// Solidity: function getAcceptersV2(uint256 proposalId) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetAcceptersV2(proposalId *big.Int) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetAcceptersV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetComment is a free data retrieval call binding the contract method 0xcbbaa589.
//
// Solidity: function getComment(string uuid_) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetComment(opts *bind.CallOpts, uuid_ string) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getComment", uuid_)
	return *ret0, err
}

// GetComment is a free data retrieval call binding the contract method 0xcbbaa589.
//
// Solidity: function getComment(string uuid_) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetComment(uuid_ string) ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetComment(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetComment is a free data retrieval call binding the contract method 0xcbbaa589.
//
// Solidity: function getComment(string uuid_) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetComment(uuid_ string) ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetComment(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetCommentV2 is a free data retrieval call binding the contract method 0x0454d213.
//
// Solidity: function getCommentV2(uint256 proposalId) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetCommentV2(opts *bind.CallOpts, proposalId *big.Int) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getCommentV2", proposalId)
	return *ret0, err
}

// GetCommentV2 is a free data retrieval call binding the contract method 0x0454d213.
//
// Solidity: function getCommentV2(uint256 proposalId) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetCommentV2(proposalId *big.Int) ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetCommentV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetCommentV2 is a free data retrieval call binding the contract method 0x0454d213.
//
// Solidity: function getCommentV2(uint256 proposalId) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetCommentV2(proposalId *big.Int) ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetCommentV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetPagenationProposals is a free data retrieval call binding the contract method 0x39972d5f.
//
// Solidity: function getPagenationProposals(uint256 page, uint256 pageSize) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetPagenationProposals(opts *bind.CallOpts, page *big.Int, pageSize *big.Int) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getPagenationProposals", page, pageSize)
	return *ret0, err
}

// GetPagenationProposals is a free data retrieval call binding the contract method 0x39972d5f.
//
// Solidity: function getPagenationProposals(uint256 page, uint256 pageSize) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetPagenationProposals(page *big.Int, pageSize *big.Int) ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetPagenationProposals(&_ProposalManagerAppV2.CallOpts, page, pageSize)
}

// GetPagenationProposals is a free data retrieval call binding the contract method 0x39972d5f.
//
// Solidity: function getPagenationProposals(uint256 page, uint256 pageSize) view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetPagenationProposals(page *big.Int, pageSize *big.Int) ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetPagenationProposals(&_ProposalManagerAppV2.CallOpts, page, pageSize)
}

// GetParticipants is a free data retrieval call binding the contract method 0x7e90def7.
//
// Solidity: function getParticipants(string uuid_) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetParticipants(opts *bind.CallOpts, uuid_ string) (bool, []common.Address, error) {
	var (
		ret0 = new(bool)
		ret1 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getParticipants", uuid_)
	return *ret0, *ret1, err
}

// GetParticipants is a free data retrieval call binding the contract method 0x7e90def7.
//
// Solidity: function getParticipants(string uuid_) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetParticipants(uuid_ string) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetParticipants(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetParticipants is a free data retrieval call binding the contract method 0x7e90def7.
//
// Solidity: function getParticipants(string uuid_) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetParticipants(uuid_ string) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetParticipants(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetParticipantsV2 is a free data retrieval call binding the contract method 0x71166f71.
//
// Solidity: function getParticipantsV2(uint256 proposalId) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetParticipantsV2(opts *bind.CallOpts, proposalId *big.Int) (bool, []common.Address, error) {
	var (
		ret0 = new(bool)
		ret1 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getParticipantsV2", proposalId)
	return *ret0, *ret1, err
}

// GetParticipantsV2 is a free data retrieval call binding the contract method 0x71166f71.
//
// Solidity: function getParticipantsV2(uint256 proposalId) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetParticipantsV2(proposalId *big.Int) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetParticipantsV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetParticipantsV2 is a free data retrieval call binding the contract method 0x71166f71.
//
// Solidity: function getParticipantsV2(uint256 proposalId) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetParticipantsV2(proposalId *big.Int) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetParticipantsV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetPolicies(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getPolicies")
	return *ret0, err
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetPolicies() ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetPolicies(&_ProposalManagerAppV2.CallOpts)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetPolicies() ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetPolicies(&_ProposalManagerAppV2.CallOpts)
}

// GetPolicy is a free data retrieval call binding the contract method 0x60dd5f90.
//
// Solidity: function getPolicy(string policyName_) view returns(uint256)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetPolicy(opts *bind.CallOpts, policyName_ string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getPolicy", policyName_)
	return *ret0, err
}

// GetPolicy is a free data retrieval call binding the contract method 0x60dd5f90.
//
// Solidity: function getPolicy(string policyName_) view returns(uint256)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetPolicy(policyName_ string) (*big.Int, error) {
	return _ProposalManagerAppV2.Contract.GetPolicy(&_ProposalManagerAppV2.CallOpts, policyName_)
}

// GetPolicy is a free data retrieval call binding the contract method 0x60dd5f90.
//
// Solidity: function getPolicy(string policyName_) view returns(uint256)
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetPolicy(policyName_ string) (*big.Int, error) {
	return _ProposalManagerAppV2.Contract.GetPolicy(&_ProposalManagerAppV2.CallOpts, policyName_)
}

// GetProposal is a free data retrieval call binding the contract method 0x25d3a09d.
//
// Solidity: function getProposal(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetProposal(opts *bind.CallOpts, uuid string) (IDataProposalParameter, error) {
	var (
		ret0 = new(IDataProposalParameter)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getProposal", uuid)
	return *ret0, err
}

// GetProposal is a free data retrieval call binding the contract method 0x25d3a09d.
//
// Solidity: function getProposal(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetProposal(uuid string) (IDataProposalParameter, error) {
	return _ProposalManagerAppV2.Contract.GetProposal(&_ProposalManagerAppV2.CallOpts, uuid)
}

// GetProposal is a free data retrieval call binding the contract method 0x25d3a09d.
//
// Solidity: function getProposal(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetProposal(uuid string) (IDataProposalParameter, error) {
	return _ProposalManagerAppV2.Contract.GetProposal(&_ProposalManagerAppV2.CallOpts, uuid)
}

// GetProposalStructure is a free data retrieval call binding the contract method 0xfef2293f.
//
// Solidity: function getProposalStructure(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetProposalStructure(opts *bind.CallOpts, uuid string) (IDataProposalParameter, error) {
	var (
		ret0 = new(IDataProposalParameter)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getProposalStructure", uuid)
	return *ret0, err
}

// GetProposalStructure is a free data retrieval call binding the contract method 0xfef2293f.
//
// Solidity: function getProposalStructure(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetProposalStructure(uuid string) (IDataProposalParameter, error) {
	return _ProposalManagerAppV2.Contract.GetProposalStructure(&_ProposalManagerAppV2.CallOpts, uuid)
}

// GetProposalStructure is a free data retrieval call binding the contract method 0xfef2293f.
//
// Solidity: function getProposalStructure(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetProposalStructure(uuid string) (IDataProposalParameter, error) {
	return _ProposalManagerAppV2.Contract.GetProposalStructure(&_ProposalManagerAppV2.CallOpts, uuid)
}

// GetProposalStructureV2 is a free data retrieval call binding the contract method 0x3cd7c90b.
//
// Solidity: function getProposalStructureV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetProposalStructureV2(opts *bind.CallOpts, proposalId *big.Int) (ProposalParameterV2, error) {
	var (
		ret0 = new(ProposalParameterV2)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getProposalStructureV2", proposalId)
	return *ret0, err
}

// GetProposalStructureV2 is a free data retrieval call binding the contract method 0x3cd7c90b.
//
// Solidity: function getProposalStructureV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetProposalStructureV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _ProposalManagerAppV2.Contract.GetProposalStructureV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetProposalStructureV2 is a free data retrieval call binding the contract method 0x3cd7c90b.
//
// Solidity: function getProposalStructureV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetProposalStructureV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _ProposalManagerAppV2.Contract.GetProposalStructureV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetProposalV2 is a free data retrieval call binding the contract method 0xc6d61e6b.
//
// Solidity: function getProposalV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetProposalV2(opts *bind.CallOpts, proposalId *big.Int) (ProposalParameterV2, error) {
	var (
		ret0 = new(ProposalParameterV2)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getProposalV2", proposalId)
	return *ret0, err
}

// GetProposalV2 is a free data retrieval call binding the contract method 0xc6d61e6b.
//
// Solidity: function getProposalV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetProposalV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _ProposalManagerAppV2.Contract.GetProposalV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetProposalV2 is a free data retrieval call binding the contract method 0xc6d61e6b.
//
// Solidity: function getProposalV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetProposalV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _ProposalManagerAppV2.Contract.GetProposalV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetProposedPolicy is a free data retrieval call binding the contract method 0x44ecebb1.
//
// Solidity: function getProposedPolicy(string uuid_) view returns((string,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetProposedPolicy(opts *bind.CallOpts, uuid_ string) (IPolicyDataProposedPoicy, error) {
	var (
		ret0 = new(IPolicyDataProposedPoicy)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getProposedPolicy", uuid_)
	return *ret0, err
}

// GetProposedPolicy is a free data retrieval call binding the contract method 0x44ecebb1.
//
// Solidity: function getProposedPolicy(string uuid_) view returns((string,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetProposedPolicy(uuid_ string) (IPolicyDataProposedPoicy, error) {
	return _ProposalManagerAppV2.Contract.GetProposedPolicy(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetProposedPolicy is a free data retrieval call binding the contract method 0x44ecebb1.
//
// Solidity: function getProposedPolicy(string uuid_) view returns((string,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetProposedPolicy(uuid_ string) (IPolicyDataProposedPoicy, error) {
	return _ProposalManagerAppV2.Contract.GetProposedPolicy(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetProposedPolicyV2 is a free data retrieval call binding the contract method 0x03ec9978.
//
// Solidity: function getProposedPolicyV2(uint256 proposalId) view returns((string,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetProposedPolicyV2(opts *bind.CallOpts, proposalId *big.Int) (IPolicyDataV2ProposedPoicy, error) {
	var (
		ret0 = new(IPolicyDataV2ProposedPoicy)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getProposedPolicyV2", proposalId)
	return *ret0, err
}

// GetProposedPolicyV2 is a free data retrieval call binding the contract method 0x03ec9978.
//
// Solidity: function getProposedPolicyV2(uint256 proposalId) view returns((string,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetProposedPolicyV2(proposalId *big.Int) (IPolicyDataV2ProposedPoicy, error) {
	return _ProposalManagerAppV2.Contract.GetProposedPolicyV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetProposedPolicyV2 is a free data retrieval call binding the contract method 0x03ec9978.
//
// Solidity: function getProposedPolicyV2(uint256 proposalId) view returns((string,uint256))
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetProposedPolicyV2(proposalId *big.Int) (IPolicyDataV2ProposedPoicy, error) {
	return _ProposalManagerAppV2.Contract.GetProposedPolicyV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetRejecters is a free data retrieval call binding the contract method 0x2149e6c6.
//
// Solidity: function getRejecters(string uuid_) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetRejecters(opts *bind.CallOpts, uuid_ string) (bool, []common.Address, error) {
	var (
		ret0 = new(bool)
		ret1 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getRejecters", uuid_)
	return *ret0, *ret1, err
}

// GetRejecters is a free data retrieval call binding the contract method 0x2149e6c6.
//
// Solidity: function getRejecters(string uuid_) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetRejecters(uuid_ string) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetRejecters(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetRejecters is a free data retrieval call binding the contract method 0x2149e6c6.
//
// Solidity: function getRejecters(string uuid_) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetRejecters(uuid_ string) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetRejecters(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetRejectersV2 is a free data retrieval call binding the contract method 0x2350bdc7.
//
// Solidity: function getRejectersV2(uint256 proposalId) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetRejectersV2(opts *bind.CallOpts, proposalId *big.Int) (bool, []common.Address, error) {
	var (
		ret0 = new(bool)
		ret1 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getRejectersV2", proposalId)
	return *ret0, *ret1, err
}

// GetRejectersV2 is a free data retrieval call binding the contract method 0x2350bdc7.
//
// Solidity: function getRejectersV2(uint256 proposalId) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetRejectersV2(proposalId *big.Int) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetRejectersV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetRejectersV2 is a free data retrieval call binding the contract method 0x2350bdc7.
//
// Solidity: function getRejectersV2(uint256 proposalId) view returns(bool, address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetRejectersV2(proposalId *big.Int) (bool, []common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetRejectersV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetStateAddr(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getStateAddr")
	return *ret0, err
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetStateAddr() ([]common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetStateAddr(&_ProposalManagerAppV2.CallOpts)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetStateAddr() ([]common.Address, error) {
	return _ProposalManagerAppV2.Contract.GetStateAddr(&_ProposalManagerAppV2.CallOpts)
}

// GetTotalProposals is a free data retrieval call binding the contract method 0x1a5007dd.
//
// Solidity: function getTotalProposals() view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetTotalProposals(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getTotalProposals")
	return *ret0, err
}

// GetTotalProposals is a free data retrieval call binding the contract method 0x1a5007dd.
//
// Solidity: function getTotalProposals() view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetTotalProposals() ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetTotalProposals(&_ProposalManagerAppV2.CallOpts)
}

// GetTotalProposals is a free data retrieval call binding the contract method 0x1a5007dd.
//
// Solidity: function getTotalProposals() view returns(string[])
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetTotalProposals() ([]string, error) {
	return _ProposalManagerAppV2.Contract.GetTotalProposals(&_ProposalManagerAppV2.CallOpts)
}

// GetVoterValue is a free data retrieval call binding the contract method 0x1f3f882a.
//
// Solidity: function getVoterValue(string uuid_) view returns(uint256)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetVoterValue(opts *bind.CallOpts, uuid_ string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getVoterValue", uuid_)
	return *ret0, err
}

// GetVoterValue is a free data retrieval call binding the contract method 0x1f3f882a.
//
// Solidity: function getVoterValue(string uuid_) view returns(uint256)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetVoterValue(uuid_ string) (*big.Int, error) {
	return _ProposalManagerAppV2.Contract.GetVoterValue(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetVoterValue is a free data retrieval call binding the contract method 0x1f3f882a.
//
// Solidity: function getVoterValue(string uuid_) view returns(uint256)
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetVoterValue(uuid_ string) (*big.Int, error) {
	return _ProposalManagerAppV2.Contract.GetVoterValue(&_ProposalManagerAppV2.CallOpts, uuid_)
}

// GetVoterValueV2 is a free data retrieval call binding the contract method 0x9b636bfc.
//
// Solidity: function getVoterValueV2(uint256 proposalId) view returns(uint256)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) GetVoterValueV2(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "getVoterValueV2", proposalId)
	return *ret0, err
}

// GetVoterValueV2 is a free data retrieval call binding the contract method 0x9b636bfc.
//
// Solidity: function getVoterValueV2(uint256 proposalId) view returns(uint256)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) GetVoterValueV2(proposalId *big.Int) (*big.Int, error) {
	return _ProposalManagerAppV2.Contract.GetVoterValueV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// GetVoterValueV2 is a free data retrieval call binding the contract method 0x9b636bfc.
//
// Solidity: function getVoterValueV2(uint256 proposalId) view returns(uint256)
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) GetVoterValueV2(proposalId *big.Int) (*big.Int, error) {
	return _ProposalManagerAppV2.Contract.GetVoterValueV2(&_ProposalManagerAppV2.CallOpts, proposalId)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) NameExt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "nameExt")
	return *ret0, err
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) NameExt() (string, error) {
	return _ProposalManagerAppV2.Contract.NameExt(&_ProposalManagerAppV2.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) NameExt() (string, error) {
	return _ProposalManagerAppV2.Contract.NameExt(&_ProposalManagerAppV2.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) Ready() (bool, error) {
	return _ProposalManagerAppV2.Contract.Ready(&_ProposalManagerAppV2.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) Ready() (bool, error) {
	return _ProposalManagerAppV2.Contract.Ready(&_ProposalManagerAppV2.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Caller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProposalManagerAppV2.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) Registry() (common.Address, error) {
	return _ProposalManagerAppV2.Contract.Registry(&_ProposalManagerAppV2.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ProposalManagerAppV2 *ProposalManagerAppV2CallerSession) Registry() (common.Address, error) {
	return _ProposalManagerAppV2.Contract.Registry(&_ProposalManagerAppV2.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2Transactor) Initialize(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.contract.Transact(opts, "initialize", addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.Initialize(&_ProposalManagerAppV2.TransactOpts, addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2TransactorSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.Initialize(&_ProposalManagerAppV2.TransactOpts, addrs)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2Transactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.contract.Transact(opts, "setForwarder", forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.SetForwarder(&_ProposalManagerAppV2.TransactOpts, forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2TransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.SetForwarder(&_ProposalManagerAppV2.TransactOpts, forwarder)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2Transactor) Upgrade(opts *bind.TransactOpts, newVersion common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.contract.Transact(opts, "upgrade", newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.Upgrade(&_ProposalManagerAppV2.TransactOpts, newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2TransactorSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.Upgrade(&_ProposalManagerAppV2.TransactOpts, newVersion)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2Transactor) UpgradeRegistry(opts *bind.TransactOpts, reg common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.contract.Transact(opts, "upgradeRegistry", reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2Session) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.UpgradeRegistry(&_ProposalManagerAppV2.TransactOpts, reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_ProposalManagerAppV2 *ProposalManagerAppV2TransactorSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _ProposalManagerAppV2.Contract.UpgradeRegistry(&_ProposalManagerAppV2.TransactOpts, reg)
}

// ProposalManagerAppV2BeforeUpgradeHookEventIterator is returned from FilterBeforeUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for BeforeUpgradeHookEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2BeforeUpgradeHookEventIterator struct {
	Event *ProposalManagerAppV2BeforeUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2BeforeUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2BeforeUpgradeHookEvent)
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
		it.Event = new(ProposalManagerAppV2BeforeUpgradeHookEvent)
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
func (it *ProposalManagerAppV2BeforeUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2BeforeUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2BeforeUpgradeHookEvent represents a BeforeUpgradeHookEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2BeforeUpgradeHookEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBeforeUpgradeHookEvent is a free log retrieval operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterBeforeUpgradeHookEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2BeforeUpgradeHookEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2BeforeUpgradeHookEventIterator{contract: _ProposalManagerAppV2.contract, event: "beforeUpgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchBeforeUpgradeHookEvent is a free log subscription operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchBeforeUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2BeforeUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2BeforeUpgradeHookEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
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
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseBeforeUpgradeHookEvent(log types.Log) (*ProposalManagerAppV2BeforeUpgradeHookEvent, error) {
	event := new(ProposalManagerAppV2BeforeUpgradeHookEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalManagerAppV2ChangeInitialDataEventIterator is returned from FilterChangeInitialDataEvent and is used to iterate over the raw logs and unpacked data for ChangeInitialDataEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2ChangeInitialDataEventIterator struct {
	Event *ProposalManagerAppV2ChangeInitialDataEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2ChangeInitialDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2ChangeInitialDataEvent)
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
		it.Event = new(ProposalManagerAppV2ChangeInitialDataEvent)
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
func (it *ProposalManagerAppV2ChangeInitialDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2ChangeInitialDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2ChangeInitialDataEvent represents a ChangeInitialDataEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2ChangeInitialDataEvent struct {
	EnableSeeVote bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChangeInitialDataEvent is a free log retrieval operation binding the contract event 0xd42ce7a4847e33e6fb2eff528c77b46a56e58e687bfb166f964d3e0ec4a34f0e.
//
// Solidity: event changeInitialDataEvent(bool enableSeeVote)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterChangeInitialDataEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2ChangeInitialDataEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2ChangeInitialDataEventIterator{contract: _ProposalManagerAppV2.contract, event: "changeInitialDataEvent", logs: logs, sub: sub}, nil
}

// WatchChangeInitialDataEvent is a free log subscription operation binding the contract event 0xd42ce7a4847e33e6fb2eff528c77b46a56e58e687bfb166f964d3e0ec4a34f0e.
//
// Solidity: event changeInitialDataEvent(bool enableSeeVote)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchChangeInitialDataEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2ChangeInitialDataEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2ChangeInitialDataEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
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

// ParseChangeInitialDataEvent is a log parse operation binding the contract event 0xd42ce7a4847e33e6fb2eff528c77b46a56e58e687bfb166f964d3e0ec4a34f0e.
//
// Solidity: event changeInitialDataEvent(bool enableSeeVote)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseChangeInitialDataEvent(log types.Log) (*ProposalManagerAppV2ChangeInitialDataEvent, error) {
	event := new(ProposalManagerAppV2ChangeInitialDataEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalManagerAppV2EmitActivityEventIterator is returned from FilterEmitActivityEvent and is used to iterate over the raw logs and unpacked data for EmitActivityEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2EmitActivityEventIterator struct {
	Event *ProposalManagerAppV2EmitActivityEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2EmitActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2EmitActivityEvent)
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
		it.Event = new(ProposalManagerAppV2EmitActivityEvent)
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
func (it *ProposalManagerAppV2EmitActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2EmitActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2EmitActivityEvent represents a EmitActivityEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2EmitActivityEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitActivityEvent is a free log retrieval operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterEmitActivityEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2EmitActivityEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2EmitActivityEventIterator{contract: _ProposalManagerAppV2.contract, event: "emitActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitActivityEvent is a free log subscription operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchEmitActivityEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2EmitActivityEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2EmitActivityEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
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
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseEmitActivityEvent(log types.Log) (*ProposalManagerAppV2EmitActivityEvent, error) {
	event := new(ProposalManagerAppV2EmitActivityEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalManagerAppV2EmitAppDataEventIterator is returned from FilterEmitAppDataEvent and is used to iterate over the raw logs and unpacked data for EmitAppDataEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2EmitAppDataEventIterator struct {
	Event *ProposalManagerAppV2EmitAppDataEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2EmitAppDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2EmitAppDataEvent)
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
		it.Event = new(ProposalManagerAppV2EmitAppDataEvent)
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
func (it *ProposalManagerAppV2EmitAppDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2EmitAppDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2EmitAppDataEvent represents a EmitAppDataEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2EmitAppDataEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitAppDataEvent is a free log retrieval operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterEmitAppDataEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2EmitAppDataEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2EmitAppDataEventIterator{contract: _ProposalManagerAppV2.contract, event: "emitAppDataEvent", logs: logs, sub: sub}, nil
}

// WatchEmitAppDataEvent is a free log subscription operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchEmitAppDataEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2EmitAppDataEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2EmitAppDataEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
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
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseEmitAppDataEvent(log types.Log) (*ProposalManagerAppV2EmitAppDataEvent, error) {
	event := new(ProposalManagerAppV2EmitAppDataEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalManagerAppV2EmitProposalActivityEventIterator is returned from FilterEmitProposalActivityEvent and is used to iterate over the raw logs and unpacked data for EmitProposalActivityEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2EmitProposalActivityEventIterator struct {
	Event *ProposalManagerAppV2EmitProposalActivityEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2EmitProposalActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2EmitProposalActivityEvent)
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
		it.Event = new(ProposalManagerAppV2EmitProposalActivityEvent)
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
func (it *ProposalManagerAppV2EmitProposalActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2EmitProposalActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2EmitProposalActivityEvent represents a EmitProposalActivityEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2EmitProposalActivityEvent struct {
	ActivityType *big.Int
	Proposal     ProposalSummary
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitProposalActivityEvent is a free log retrieval operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterEmitProposalActivityEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2EmitProposalActivityEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2EmitProposalActivityEventIterator{contract: _ProposalManagerAppV2.contract, event: "emitProposalActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitProposalActivityEvent is a free log subscription operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchEmitProposalActivityEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2EmitProposalActivityEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2EmitProposalActivityEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
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
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseEmitProposalActivityEvent(log types.Log) (*ProposalManagerAppV2EmitProposalActivityEvent, error) {
	event := new(ProposalManagerAppV2EmitProposalActivityEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalManagerAppV2InitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2InitializeEventIterator struct {
	Event *ProposalManagerAppV2InitializeEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2InitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2InitializeEvent)
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
		it.Event = new(ProposalManagerAppV2InitializeEvent)
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
func (it *ProposalManagerAppV2InitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2InitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2InitializeEvent represents a InitializeEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2InitializeEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterInitializeEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2InitializeEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2InitializeEventIterator{contract: _ProposalManagerAppV2.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2InitializeEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2InitializeEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseInitializeEvent(log types.Log) (*ProposalManagerAppV2InitializeEvent, error) {
	event := new(ProposalManagerAppV2InitializeEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalManagerAppV2InitializeHookEventIterator is returned from FilterInitializeHookEvent and is used to iterate over the raw logs and unpacked data for InitializeHookEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2InitializeHookEventIterator struct {
	Event *ProposalManagerAppV2InitializeHookEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2InitializeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2InitializeHookEvent)
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
		it.Event = new(ProposalManagerAppV2InitializeHookEvent)
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
func (it *ProposalManagerAppV2InitializeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2InitializeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2InitializeHookEvent represents a InitializeHookEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2InitializeHookEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeHookEvent is a free log retrieval operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterInitializeHookEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2InitializeHookEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2InitializeHookEventIterator{contract: _ProposalManagerAppV2.contract, event: "initializeHookEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeHookEvent is a free log subscription operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchInitializeHookEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2InitializeHookEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2InitializeHookEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
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
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseInitializeHookEvent(log types.Log) (*ProposalManagerAppV2InitializeHookEvent, error) {
	event := new(ProposalManagerAppV2InitializeHookEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalManagerAppV2SetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2SetReadyEventIterator struct {
	Event *ProposalManagerAppV2SetReadyEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2SetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2SetReadyEvent)
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
		it.Event = new(ProposalManagerAppV2SetReadyEvent)
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
func (it *ProposalManagerAppV2SetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2SetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2SetReadyEvent represents a SetReadyEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2SetReadyEvent struct {
	Done bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2SetReadyEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2SetReadyEventIterator{contract: _ProposalManagerAppV2.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2SetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2SetReadyEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseSetReadyEvent(log types.Log) (*ProposalManagerAppV2SetReadyEvent, error) {
	event := new(ProposalManagerAppV2SetReadyEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalManagerAppV2UpgradeEventIterator is returned from FilterUpgradeEvent and is used to iterate over the raw logs and unpacked data for UpgradeEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2UpgradeEventIterator struct {
	Event *ProposalManagerAppV2UpgradeEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2UpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2UpgradeEvent)
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
		it.Event = new(ProposalManagerAppV2UpgradeEvent)
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
func (it *ProposalManagerAppV2UpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2UpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2UpgradeEvent represents a UpgradeEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2UpgradeEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpgradeEvent is a free log retrieval operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterUpgradeEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2UpgradeEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2UpgradeEventIterator{contract: _ProposalManagerAppV2.contract, event: "upgradeEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeEvent is a free log subscription operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchUpgradeEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2UpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2UpgradeEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
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
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseUpgradeEvent(log types.Log) (*ProposalManagerAppV2UpgradeEvent, error) {
	event := new(ProposalManagerAppV2UpgradeEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalManagerAppV2UpgradeRegistryEventIterator is returned from FilterUpgradeRegistryEvent and is used to iterate over the raw logs and unpacked data for UpgradeRegistryEvent events raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2UpgradeRegistryEventIterator struct {
	Event *ProposalManagerAppV2UpgradeRegistryEvent // Event containing the contract specifics and raw log

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
func (it *ProposalManagerAppV2UpgradeRegistryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalManagerAppV2UpgradeRegistryEvent)
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
		it.Event = new(ProposalManagerAppV2UpgradeRegistryEvent)
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
func (it *ProposalManagerAppV2UpgradeRegistryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalManagerAppV2UpgradeRegistryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalManagerAppV2UpgradeRegistryEvent represents a UpgradeRegistryEvent event raised by the ProposalManagerAppV2 contract.
type ProposalManagerAppV2UpgradeRegistryEvent struct {
	Reg common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpgradeRegistryEvent is a free log retrieval operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) FilterUpgradeRegistryEvent(opts *bind.FilterOpts) (*ProposalManagerAppV2UpgradeRegistryEventIterator, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.FilterLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalManagerAppV2UpgradeRegistryEventIterator{contract: _ProposalManagerAppV2.contract, event: "upgradeRegistryEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeRegistryEvent is a free log subscription operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) WatchUpgradeRegistryEvent(opts *bind.WatchOpts, sink chan<- *ProposalManagerAppV2UpgradeRegistryEvent) (event.Subscription, error) {

	logs, sub, err := _ProposalManagerAppV2.contract.WatchLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalManagerAppV2UpgradeRegistryEvent)
				if err := _ProposalManagerAppV2.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
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
func (_ProposalManagerAppV2 *ProposalManagerAppV2Filterer) ParseUpgradeRegistryEvent(log types.Log) (*ProposalManagerAppV2UpgradeRegistryEvent, error) {
	event := new(ProposalManagerAppV2UpgradeRegistryEvent)
	if err := _ProposalManagerAppV2.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
