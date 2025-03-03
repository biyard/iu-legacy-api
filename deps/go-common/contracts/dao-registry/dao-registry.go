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

// RegisteredExtension is an auto generated low-level Go binding around an user-defined struct.
type RegisteredExtension struct {
	Name            string
	Revision        *big.Int
	Addr            common.Address
	Permissions     *big.Int
	HookPermissions *big.Int
	ExtensionType   uint8
}

// DaoRegistryMetaData contains all meta data concerning the DaoRegistry contract.
var DaoRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"state\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"}],\"name\":\"MetaTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"changeActivityHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookPermissions\",\"type\":\"uint256\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structRegisteredExtension\",\"name\":\"ext\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDelete\",\"type\":\"bool\"}],\"name\":\"changeExtensionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"daoManagerAddress\",\"type\":\"address\"}],\"name\":\"changeInitialDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"changeProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"extName\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"indexed\":false,\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"stateDeployCheck\",\"type\":\"bool\"}],\"name\":\"deployExtensionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"extName\",\"type\":\"string\"}],\"name\":\"deployExtensionStateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"}],\"name\":\"deployStandardContractEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"setDaoFactoryAddressEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"check\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"storedAppDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"upgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"state\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"}],\"name\":\"upgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdrawalEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC712_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookPermissions\",\"type\":\"uint256\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structRegisteredExtension\",\"name\":\"ext\",\"type\":\"tuple\"}],\"name\":\"addExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"addressOfExtension\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"perm\",\"type\":\"uint256\"}],\"name\":\"checkExtensionPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"extName\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"stateDeployCheck\",\"type\":\"bool\"}],\"name\":\"deployExtension\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookPermissions\",\"type\":\"uint256\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structRegisteredExtension\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"extName\",\"type\":\"string\"}],\"name\":\"deployExtensionState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"}],\"name\":\"deployStandardContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"callee\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"executeMetaTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensions\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookPermissions\",\"type\":\"uint256\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structRegisteredExtension[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"}],\"name\":\"finishVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeperator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"internalType\":\"structProposalSummary\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateAddr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"}],\"name\":\"getTypedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"handleActivity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"name\":\"listExtensionsByType\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"internalType\":\"structProposalSummary[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"names\",\"type\":\"string[]\"}],\"name\":\"removeExtensions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"setDaoFactoryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"check\",\"type\":\"bool\"}],\"name\":\"setReady\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"storedAppData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"internalType\":\"structProposalSummary\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"submitProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"toTypedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookPermissions\",\"type\":\"uint256\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structRegisteredExtension\",\"name\":\"newExt\",\"type\":\"tuple\"}],\"name\":\"upgradeExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"state\",\"type\":\"address\"}],\"name\":\"upgradeHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DaoRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DaoRegistryMetaData.ABI instead.
var DaoRegistryABI = DaoRegistryMetaData.ABI

// DaoRegistryBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const DaoRegistryBinRuntime = ``

// DaoRegistry is an auto generated Go binding around a Klaytn contract.
type DaoRegistry struct {
	DaoRegistryCaller     // Read-only binding to the contract
	DaoRegistryTransactor // Write-only binding to the contract
	DaoRegistryFilterer   // Log filterer for contract events
}

// DaoRegistryCaller is an auto generated read-only Go binding around a Klaytn contract.
type DaoRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoRegistryTransactor is an auto generated write-only Go binding around a Klaytn contract.
type DaoRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoRegistryFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type DaoRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoRegistrySession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type DaoRegistrySession struct {
	Contract     *DaoRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoRegistryCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type DaoRegistryCallerSession struct {
	Contract *DaoRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DaoRegistryTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type DaoRegistryTransactorSession struct {
	Contract     *DaoRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DaoRegistryRaw is an auto generated low-level Go binding around a Klaytn contract.
type DaoRegistryRaw struct {
	Contract *DaoRegistry // Generic contract binding to access the raw methods on
}

// DaoRegistryCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type DaoRegistryCallerRaw struct {
	Contract *DaoRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DaoRegistryTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type DaoRegistryTransactorRaw struct {
	Contract *DaoRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDaoRegistry creates a new instance of DaoRegistry, bound to a specific deployed contract.
func NewDaoRegistry(address common.Address, backend bind.ContractBackend) (*DaoRegistry, error) {
	contract, err := bindDaoRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DaoRegistry{DaoRegistryCaller: DaoRegistryCaller{contract: contract}, DaoRegistryTransactor: DaoRegistryTransactor{contract: contract}, DaoRegistryFilterer: DaoRegistryFilterer{contract: contract}}, nil
}

// NewDaoRegistryCaller creates a new read-only instance of DaoRegistry, bound to a specific deployed contract.
func NewDaoRegistryCaller(address common.Address, caller bind.ContractCaller) (*DaoRegistryCaller, error) {
	contract, err := bindDaoRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaoRegistryCaller{contract: contract}, nil
}

// NewDaoRegistryTransactor creates a new write-only instance of DaoRegistry, bound to a specific deployed contract.
func NewDaoRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DaoRegistryTransactor, error) {
	contract, err := bindDaoRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaoRegistryTransactor{contract: contract}, nil
}

// NewDaoRegistryFilterer creates a new log filterer instance of DaoRegistry, bound to a specific deployed contract.
func NewDaoRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DaoRegistryFilterer, error) {
	contract, err := bindDaoRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaoRegistryFilterer{contract: contract}, nil
}

// bindDaoRegistry binds a generic wrapper to an already deployed contract.
func bindDaoRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DaoRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoRegistry *DaoRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DaoRegistry.Contract.DaoRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoRegistry *DaoRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoRegistry.Contract.DaoRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoRegistry *DaoRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoRegistry.Contract.DaoRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoRegistry *DaoRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DaoRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoRegistry *DaoRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoRegistry *DaoRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoRegistry.Contract.contract.Transact(opts, method, params...)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_DaoRegistry *DaoRegistryCaller) ERC712VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "ERC712_VERSION")
	return *ret0, err
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_DaoRegistry *DaoRegistrySession) ERC712VERSION() (string, error) {
	return _DaoRegistry.Contract.ERC712VERSION(&_DaoRegistry.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_DaoRegistry *DaoRegistryCallerSession) ERC712VERSION() (string, error) {
	return _DaoRegistry.Contract.ERC712VERSION(&_DaoRegistry.CallOpts)
}

// AddressOfExtension is a free data retrieval call binding the contract method 0x1df6c5f2.
//
// Solidity: function addressOfExtension(string name_) view returns(address)
func (_DaoRegistry *DaoRegistryCaller) AddressOfExtension(opts *bind.CallOpts, name_ string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "addressOfExtension", name_)
	return *ret0, err
}

// AddressOfExtension is a free data retrieval call binding the contract method 0x1df6c5f2.
//
// Solidity: function addressOfExtension(string name_) view returns(address)
func (_DaoRegistry *DaoRegistrySession) AddressOfExtension(name_ string) (common.Address, error) {
	return _DaoRegistry.Contract.AddressOfExtension(&_DaoRegistry.CallOpts, name_)
}

// AddressOfExtension is a free data retrieval call binding the contract method 0x1df6c5f2.
//
// Solidity: function addressOfExtension(string name_) view returns(address)
func (_DaoRegistry *DaoRegistryCallerSession) AddressOfExtension(name_ string) (common.Address, error) {
	return _DaoRegistry.Contract.AddressOfExtension(&_DaoRegistry.CallOpts, name_)
}

// CheckExtensionPermission is a free data retrieval call binding the contract method 0xffa8c73d.
//
// Solidity: function checkExtensionPermission(address addr, uint256 perm) view returns(bool)
func (_DaoRegistry *DaoRegistryCaller) CheckExtensionPermission(opts *bind.CallOpts, addr common.Address, perm *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "checkExtensionPermission", addr, perm)
	return *ret0, err
}

// CheckExtensionPermission is a free data retrieval call binding the contract method 0xffa8c73d.
//
// Solidity: function checkExtensionPermission(address addr, uint256 perm) view returns(bool)
func (_DaoRegistry *DaoRegistrySession) CheckExtensionPermission(addr common.Address, perm *big.Int) (bool, error) {
	return _DaoRegistry.Contract.CheckExtensionPermission(&_DaoRegistry.CallOpts, addr, perm)
}

// CheckExtensionPermission is a free data retrieval call binding the contract method 0xffa8c73d.
//
// Solidity: function checkExtensionPermission(address addr, uint256 perm) view returns(bool)
func (_DaoRegistry *DaoRegistryCallerSession) CheckExtensionPermission(addr common.Address, perm *big.Int) (bool, error) {
	return _DaoRegistry.Contract.CheckExtensionPermission(&_DaoRegistry.CallOpts, addr, perm)
}

// Extensions is a free data retrieval call binding the contract method 0x24f159c2.
//
// Solidity: function extensions() view returns((string,uint256,address,uint256,uint256,uint8)[])
func (_DaoRegistry *DaoRegistryCaller) Extensions(opts *bind.CallOpts) ([]RegisteredExtension, error) {
	var (
		ret0 = new([]RegisteredExtension)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "extensions")
	return *ret0, err
}

// Extensions is a free data retrieval call binding the contract method 0x24f159c2.
//
// Solidity: function extensions() view returns((string,uint256,address,uint256,uint256,uint8)[])
func (_DaoRegistry *DaoRegistrySession) Extensions() ([]RegisteredExtension, error) {
	return _DaoRegistry.Contract.Extensions(&_DaoRegistry.CallOpts)
}

// Extensions is a free data retrieval call binding the contract method 0x24f159c2.
//
// Solidity: function extensions() view returns((string,uint256,address,uint256,uint256,uint8)[])
func (_DaoRegistry *DaoRegistryCallerSession) Extensions() ([]RegisteredExtension, error) {
	return _DaoRegistry.Contract.Extensions(&_DaoRegistry.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_DaoRegistry *DaoRegistryCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_DaoRegistry *DaoRegistrySession) GetBalance() (*big.Int, error) {
	return _DaoRegistry.Contract.GetBalance(&_DaoRegistry.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_DaoRegistry *DaoRegistryCallerSession) GetBalance() (*big.Int, error) {
	return _DaoRegistry.Contract.GetBalance(&_DaoRegistry.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 id)
func (_DaoRegistry *DaoRegistryCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "getChainId")
	return *ret0, err
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 id)
func (_DaoRegistry *DaoRegistrySession) GetChainId() (*big.Int, error) {
	return _DaoRegistry.Contract.GetChainId(&_DaoRegistry.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 id)
func (_DaoRegistry *DaoRegistryCallerSession) GetChainId() (*big.Int, error) {
	return _DaoRegistry.Contract.GetChainId(&_DaoRegistry.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_DaoRegistry *DaoRegistryCaller) GetDomainSeperator(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "getDomainSeperator")
	return *ret0, err
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_DaoRegistry *DaoRegistrySession) GetDomainSeperator() ([32]byte, error) {
	return _DaoRegistry.Contract.GetDomainSeperator(&_DaoRegistry.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_DaoRegistry *DaoRegistryCallerSession) GetDomainSeperator() ([32]byte, error) {
	return _DaoRegistry.Contract.GetDomainSeperator(&_DaoRegistry.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address addr) view returns(uint256)
func (_DaoRegistry *DaoRegistryCaller) GetNonce(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "getNonce", addr)
	return *ret0, err
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address addr) view returns(uint256)
func (_DaoRegistry *DaoRegistrySession) GetNonce(addr common.Address) (*big.Int, error) {
	return _DaoRegistry.Contract.GetNonce(&_DaoRegistry.CallOpts, addr)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address addr) view returns(uint256)
func (_DaoRegistry *DaoRegistryCallerSession) GetNonce(addr common.Address) (*big.Int, error) {
	return _DaoRegistry.Contract.GetNonce(&_DaoRegistry.CallOpts, addr)
}

// GetProposal is a free data retrieval call binding the contract method 0xf95ac8c6.
//
// Solidity: function getProposal(string proposalAppName, uint256 proposalId) view returns((uint256,address,string,string,string,string,uint256,uint16,uint256))
func (_DaoRegistry *DaoRegistryCaller) GetProposal(opts *bind.CallOpts, proposalAppName string, proposalId *big.Int) (ProposalSummary, error) {
	var (
		ret0 = new(ProposalSummary)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "getProposal", proposalAppName, proposalId)
	return *ret0, err
}

// GetProposal is a free data retrieval call binding the contract method 0xf95ac8c6.
//
// Solidity: function getProposal(string proposalAppName, uint256 proposalId) view returns((uint256,address,string,string,string,string,uint256,uint16,uint256))
func (_DaoRegistry *DaoRegistrySession) GetProposal(proposalAppName string, proposalId *big.Int) (ProposalSummary, error) {
	return _DaoRegistry.Contract.GetProposal(&_DaoRegistry.CallOpts, proposalAppName, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf95ac8c6.
//
// Solidity: function getProposal(string proposalAppName, uint256 proposalId) view returns((uint256,address,string,string,string,string,uint256,uint16,uint256))
func (_DaoRegistry *DaoRegistryCallerSession) GetProposal(proposalAppName string, proposalId *big.Int) (ProposalSummary, error) {
	return _DaoRegistry.Contract.GetProposal(&_DaoRegistry.CallOpts, proposalAppName, proposalId)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_DaoRegistry *DaoRegistryCaller) GetStateAddr(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "getStateAddr")
	return *ret0, err
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_DaoRegistry *DaoRegistrySession) GetStateAddr() ([]common.Address, error) {
	return _DaoRegistry.Contract.GetStateAddr(&_DaoRegistry.CallOpts)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_DaoRegistry *DaoRegistryCallerSession) GetStateAddr() ([]common.Address, error) {
	return _DaoRegistry.Contract.GetStateAddr(&_DaoRegistry.CallOpts)
}

// GetTypedMessageHash is a free data retrieval call binding the contract method 0x55494c3e.
//
// Solidity: function getTypedMessageHash(address userAddress, bytes functionSignature) view returns(bytes32)
func (_DaoRegistry *DaoRegistryCaller) GetTypedMessageHash(opts *bind.CallOpts, userAddress common.Address, functionSignature []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "getTypedMessageHash", userAddress, functionSignature)
	return *ret0, err
}

// GetTypedMessageHash is a free data retrieval call binding the contract method 0x55494c3e.
//
// Solidity: function getTypedMessageHash(address userAddress, bytes functionSignature) view returns(bytes32)
func (_DaoRegistry *DaoRegistrySession) GetTypedMessageHash(userAddress common.Address, functionSignature []byte) ([32]byte, error) {
	return _DaoRegistry.Contract.GetTypedMessageHash(&_DaoRegistry.CallOpts, userAddress, functionSignature)
}

// GetTypedMessageHash is a free data retrieval call binding the contract method 0x55494c3e.
//
// Solidity: function getTypedMessageHash(address userAddress, bytes functionSignature) view returns(bytes32)
func (_DaoRegistry *DaoRegistryCallerSession) GetTypedMessageHash(userAddress common.Address, functionSignature []byte) ([32]byte, error) {
	return _DaoRegistry.Contract.GetTypedMessageHash(&_DaoRegistry.CallOpts, userAddress, functionSignature)
}

// ListExtensionsByType is a free data retrieval call binding the contract method 0xc4d2c25d.
//
// Solidity: function listExtensionsByType(uint8 extensionType) view returns(string[])
func (_DaoRegistry *DaoRegistryCaller) ListExtensionsByType(opts *bind.CallOpts, extensionType uint8) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "listExtensionsByType", extensionType)
	return *ret0, err
}

// ListExtensionsByType is a free data retrieval call binding the contract method 0xc4d2c25d.
//
// Solidity: function listExtensionsByType(uint8 extensionType) view returns(string[])
func (_DaoRegistry *DaoRegistrySession) ListExtensionsByType(extensionType uint8) ([]string, error) {
	return _DaoRegistry.Contract.ListExtensionsByType(&_DaoRegistry.CallOpts, extensionType)
}

// ListExtensionsByType is a free data retrieval call binding the contract method 0xc4d2c25d.
//
// Solidity: function listExtensionsByType(uint8 extensionType) view returns(string[])
func (_DaoRegistry *DaoRegistryCallerSession) ListExtensionsByType(extensionType uint8) ([]string, error) {
	return _DaoRegistry.Contract.ListExtensionsByType(&_DaoRegistry.CallOpts, extensionType)
}

// ListProposals is a free data retrieval call binding the contract method 0x26c7eb1c.
//
// Solidity: function listProposals() view returns((uint256,address,string,string,string,string,uint256,uint16,uint256)[])
func (_DaoRegistry *DaoRegistryCaller) ListProposals(opts *bind.CallOpts) ([]ProposalSummary, error) {
	var (
		ret0 = new([]ProposalSummary)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "listProposals")
	return *ret0, err
}

// ListProposals is a free data retrieval call binding the contract method 0x26c7eb1c.
//
// Solidity: function listProposals() view returns((uint256,address,string,string,string,string,uint256,uint16,uint256)[])
func (_DaoRegistry *DaoRegistrySession) ListProposals() ([]ProposalSummary, error) {
	return _DaoRegistry.Contract.ListProposals(&_DaoRegistry.CallOpts)
}

// ListProposals is a free data retrieval call binding the contract method 0x26c7eb1c.
//
// Solidity: function listProposals() view returns((uint256,address,string,string,string,string,uint256,uint16,uint256)[])
func (_DaoRegistry *DaoRegistryCallerSession) ListProposals() ([]ProposalSummary, error) {
	return _DaoRegistry.Contract.ListProposals(&_DaoRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DaoRegistry *DaoRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DaoRegistry *DaoRegistrySession) Name() (string, error) {
	return _DaoRegistry.Contract.Name(&_DaoRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DaoRegistry *DaoRegistryCallerSession) Name() (string, error) {
	return _DaoRegistry.Contract.Name(&_DaoRegistry.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_DaoRegistry *DaoRegistryCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_DaoRegistry *DaoRegistrySession) Ready() (bool, error) {
	return _DaoRegistry.Contract.Ready(&_DaoRegistry.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_DaoRegistry *DaoRegistryCallerSession) Ready() (bool, error) {
	return _DaoRegistry.Contract.Ready(&_DaoRegistry.CallOpts)
}

// ToTypedMessageHash is a free data retrieval call binding the contract method 0x1034b1c0.
//
// Solidity: function toTypedMessageHash(bytes32 messageHash) view returns(bytes32)
func (_DaoRegistry *DaoRegistryCaller) ToTypedMessageHash(opts *bind.CallOpts, messageHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "toTypedMessageHash", messageHash)
	return *ret0, err
}

// ToTypedMessageHash is a free data retrieval call binding the contract method 0x1034b1c0.
//
// Solidity: function toTypedMessageHash(bytes32 messageHash) view returns(bytes32)
func (_DaoRegistry *DaoRegistrySession) ToTypedMessageHash(messageHash [32]byte) ([32]byte, error) {
	return _DaoRegistry.Contract.ToTypedMessageHash(&_DaoRegistry.CallOpts, messageHash)
}

// ToTypedMessageHash is a free data retrieval call binding the contract method 0x1034b1c0.
//
// Solidity: function toTypedMessageHash(bytes32 messageHash) view returns(bytes32)
func (_DaoRegistry *DaoRegistryCallerSession) ToTypedMessageHash(messageHash [32]byte) ([32]byte, error) {
	return _DaoRegistry.Contract.ToTypedMessageHash(&_DaoRegistry.CallOpts, messageHash)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DaoRegistry *DaoRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DaoRegistry.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DaoRegistry *DaoRegistrySession) Version() (string, error) {
	return _DaoRegistry.Contract.Version(&_DaoRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DaoRegistry *DaoRegistryCallerSession) Version() (string, error) {
	return _DaoRegistry.Contract.Version(&_DaoRegistry.CallOpts)
}

// AddExtension is a paid mutator transaction binding the contract method 0x17d611f8.
//
// Solidity: function addExtension((string,uint256,address,uint256,uint256,uint8) ext) returns()
func (_DaoRegistry *DaoRegistryTransactor) AddExtension(opts *bind.TransactOpts, ext RegisteredExtension) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "addExtension", ext)
}

// AddExtension is a paid mutator transaction binding the contract method 0x17d611f8.
//
// Solidity: function addExtension((string,uint256,address,uint256,uint256,uint8) ext) returns()
func (_DaoRegistry *DaoRegistrySession) AddExtension(ext RegisteredExtension) (*types.Transaction, error) {
	return _DaoRegistry.Contract.AddExtension(&_DaoRegistry.TransactOpts, ext)
}

// AddExtension is a paid mutator transaction binding the contract method 0x17d611f8.
//
// Solidity: function addExtension((string,uint256,address,uint256,uint256,uint8) ext) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) AddExtension(ext RegisteredExtension) (*types.Transaction, error) {
	return _DaoRegistry.Contract.AddExtension(&_DaoRegistry.TransactOpts, ext)
}

// DeployExtension is a paid mutator transaction binding the contract method 0xb7dd8377.
//
// Solidity: function deployExtension(string extName, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, bool stateDeployCheck) returns((string,uint256,address,uint256,uint256,uint8))
func (_DaoRegistry *DaoRegistryTransactor) DeployExtension(opts *bind.TransactOpts, extName string, params DeploymentParameterV1, stateDeployCheck bool) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "deployExtension", extName, params, stateDeployCheck)
}

// DeployExtension is a paid mutator transaction binding the contract method 0xb7dd8377.
//
// Solidity: function deployExtension(string extName, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, bool stateDeployCheck) returns((string,uint256,address,uint256,uint256,uint8))
func (_DaoRegistry *DaoRegistrySession) DeployExtension(extName string, params DeploymentParameterV1, stateDeployCheck bool) (*types.Transaction, error) {
	return _DaoRegistry.Contract.DeployExtension(&_DaoRegistry.TransactOpts, extName, params, stateDeployCheck)
}

// DeployExtension is a paid mutator transaction binding the contract method 0xb7dd8377.
//
// Solidity: function deployExtension(string extName, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, bool stateDeployCheck) returns((string,uint256,address,uint256,uint256,uint8))
func (_DaoRegistry *DaoRegistryTransactorSession) DeployExtension(extName string, params DeploymentParameterV1, stateDeployCheck bool) (*types.Transaction, error) {
	return _DaoRegistry.Contract.DeployExtension(&_DaoRegistry.TransactOpts, extName, params, stateDeployCheck)
}

// DeployExtensionState is a paid mutator transaction binding the contract method 0xbca39057.
//
// Solidity: function deployExtensionState(string extName) returns(address)
func (_DaoRegistry *DaoRegistryTransactor) DeployExtensionState(opts *bind.TransactOpts, extName string) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "deployExtensionState", extName)
}

// DeployExtensionState is a paid mutator transaction binding the contract method 0xbca39057.
//
// Solidity: function deployExtensionState(string extName) returns(address)
func (_DaoRegistry *DaoRegistrySession) DeployExtensionState(extName string) (*types.Transaction, error) {
	return _DaoRegistry.Contract.DeployExtensionState(&_DaoRegistry.TransactOpts, extName)
}

// DeployExtensionState is a paid mutator transaction binding the contract method 0xbca39057.
//
// Solidity: function deployExtensionState(string extName) returns(address)
func (_DaoRegistry *DaoRegistryTransactorSession) DeployExtensionState(extName string) (*types.Transaction, error) {
	return _DaoRegistry.Contract.DeployExtensionState(&_DaoRegistry.TransactOpts, extName)
}

// DeployStandardContract is a paid mutator transaction binding the contract method 0x9b4843e0.
//
// Solidity: function deployStandardContract(string depName) returns(address)
func (_DaoRegistry *DaoRegistryTransactor) DeployStandardContract(opts *bind.TransactOpts, depName string) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "deployStandardContract", depName)
}

// DeployStandardContract is a paid mutator transaction binding the contract method 0x9b4843e0.
//
// Solidity: function deployStandardContract(string depName) returns(address)
func (_DaoRegistry *DaoRegistrySession) DeployStandardContract(depName string) (*types.Transaction, error) {
	return _DaoRegistry.Contract.DeployStandardContract(&_DaoRegistry.TransactOpts, depName)
}

// DeployStandardContract is a paid mutator transaction binding the contract method 0x9b4843e0.
//
// Solidity: function deployStandardContract(string depName) returns(address)
func (_DaoRegistry *DaoRegistryTransactorSession) DeployStandardContract(depName string) (*types.Transaction, error) {
	return _DaoRegistry.Contract.DeployStandardContract(&_DaoRegistry.TransactOpts, depName)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0xeae52114.
//
// Solidity: function executeMetaTransaction(string callee, address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_DaoRegistry *DaoRegistryTransactor) ExecuteMetaTransaction(opts *bind.TransactOpts, callee string, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "executeMetaTransaction", callee, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0xeae52114.
//
// Solidity: function executeMetaTransaction(string callee, address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_DaoRegistry *DaoRegistrySession) ExecuteMetaTransaction(callee string, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _DaoRegistry.Contract.ExecuteMetaTransaction(&_DaoRegistry.TransactOpts, callee, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0xeae52114.
//
// Solidity: function executeMetaTransaction(string callee, address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_DaoRegistry *DaoRegistryTransactorSession) ExecuteMetaTransaction(callee string, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _DaoRegistry.Contract.ExecuteMetaTransaction(&_DaoRegistry.TransactOpts, callee, userAddress, functionSignature, sigR, sigS, sigV)
}

// FinishVoting is a paid mutator transaction binding the contract method 0x510240a3.
//
// Solidity: function finishVoting(string proposalAppName, uint256 proposalId, uint16 voteStatus) returns()
func (_DaoRegistry *DaoRegistryTransactor) FinishVoting(opts *bind.TransactOpts, proposalAppName string, proposalId *big.Int, voteStatus uint16) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "finishVoting", proposalAppName, proposalId, voteStatus)
}

// FinishVoting is a paid mutator transaction binding the contract method 0x510240a3.
//
// Solidity: function finishVoting(string proposalAppName, uint256 proposalId, uint16 voteStatus) returns()
func (_DaoRegistry *DaoRegistrySession) FinishVoting(proposalAppName string, proposalId *big.Int, voteStatus uint16) (*types.Transaction, error) {
	return _DaoRegistry.Contract.FinishVoting(&_DaoRegistry.TransactOpts, proposalAppName, proposalId, voteStatus)
}

// FinishVoting is a paid mutator transaction binding the contract method 0x510240a3.
//
// Solidity: function finishVoting(string proposalAppName, uint256 proposalId, uint16 voteStatus) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) FinishVoting(proposalAppName string, proposalId *big.Int, voteStatus uint16) (*types.Transaction, error) {
	return _DaoRegistry.Contract.FinishVoting(&_DaoRegistry.TransactOpts, proposalAppName, proposalId, voteStatus)
}

// HandleActivity is a paid mutator transaction binding the contract method 0xa18e1d8a.
//
// Solidity: function handleActivity(uint256 activityFlag, bytes data) returns()
func (_DaoRegistry *DaoRegistryTransactor) HandleActivity(opts *bind.TransactOpts, activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "handleActivity", activityFlag, data)
}

// HandleActivity is a paid mutator transaction binding the contract method 0xa18e1d8a.
//
// Solidity: function handleActivity(uint256 activityFlag, bytes data) returns()
func (_DaoRegistry *DaoRegistrySession) HandleActivity(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _DaoRegistry.Contract.HandleActivity(&_DaoRegistry.TransactOpts, activityFlag, data)
}

// HandleActivity is a paid mutator transaction binding the contract method 0xa18e1d8a.
//
// Solidity: function handleActivity(uint256 activityFlag, bytes data) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) HandleActivity(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _DaoRegistry.Contract.HandleActivity(&_DaoRegistry.TransactOpts, activityFlag, data)
}

// RemoveExtensions is a paid mutator transaction binding the contract method 0x95b196aa.
//
// Solidity: function removeExtensions(string[] names) returns()
func (_DaoRegistry *DaoRegistryTransactor) RemoveExtensions(opts *bind.TransactOpts, names []string) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "removeExtensions", names)
}

// RemoveExtensions is a paid mutator transaction binding the contract method 0x95b196aa.
//
// Solidity: function removeExtensions(string[] names) returns()
func (_DaoRegistry *DaoRegistrySession) RemoveExtensions(names []string) (*types.Transaction, error) {
	return _DaoRegistry.Contract.RemoveExtensions(&_DaoRegistry.TransactOpts, names)
}

// RemoveExtensions is a paid mutator transaction binding the contract method 0x95b196aa.
//
// Solidity: function removeExtensions(string[] names) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) RemoveExtensions(names []string) (*types.Transaction, error) {
	return _DaoRegistry.Contract.RemoveExtensions(&_DaoRegistry.TransactOpts, names)
}

// SetDaoFactoryAddress is a paid mutator transaction binding the contract method 0xf4d343b8.
//
// Solidity: function setDaoFactoryAddress(address factory) returns()
func (_DaoRegistry *DaoRegistryTransactor) SetDaoFactoryAddress(opts *bind.TransactOpts, factory common.Address) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "setDaoFactoryAddress", factory)
}

// SetDaoFactoryAddress is a paid mutator transaction binding the contract method 0xf4d343b8.
//
// Solidity: function setDaoFactoryAddress(address factory) returns()
func (_DaoRegistry *DaoRegistrySession) SetDaoFactoryAddress(factory common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.SetDaoFactoryAddress(&_DaoRegistry.TransactOpts, factory)
}

// SetDaoFactoryAddress is a paid mutator transaction binding the contract method 0xf4d343b8.
//
// Solidity: function setDaoFactoryAddress(address factory) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) SetDaoFactoryAddress(factory common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.SetDaoFactoryAddress(&_DaoRegistry.TransactOpts, factory)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_DaoRegistry *DaoRegistryTransactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "setForwarder", forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_DaoRegistry *DaoRegistrySession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.SetForwarder(&_DaoRegistry.TransactOpts, forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.SetForwarder(&_DaoRegistry.TransactOpts, forwarder)
}

// SetReady is a paid mutator transaction binding the contract method 0xdb3b35f6.
//
// Solidity: function setReady(bool check) returns()
func (_DaoRegistry *DaoRegistryTransactor) SetReady(opts *bind.TransactOpts, check bool) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "setReady", check)
}

// SetReady is a paid mutator transaction binding the contract method 0xdb3b35f6.
//
// Solidity: function setReady(bool check) returns()
func (_DaoRegistry *DaoRegistrySession) SetReady(check bool) (*types.Transaction, error) {
	return _DaoRegistry.Contract.SetReady(&_DaoRegistry.TransactOpts, check)
}

// SetReady is a paid mutator transaction binding the contract method 0xdb3b35f6.
//
// Solidity: function setReady(bool check) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) SetReady(check bool) (*types.Transaction, error) {
	return _DaoRegistry.Contract.SetReady(&_DaoRegistry.TransactOpts, check)
}

// StoredAppData is a paid mutator transaction binding the contract method 0x57e65538.
//
// Solidity: function storedAppData(uint256 activityFlag, bytes data) returns()
func (_DaoRegistry *DaoRegistryTransactor) StoredAppData(opts *bind.TransactOpts, activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "storedAppData", activityFlag, data)
}

// StoredAppData is a paid mutator transaction binding the contract method 0x57e65538.
//
// Solidity: function storedAppData(uint256 activityFlag, bytes data) returns()
func (_DaoRegistry *DaoRegistrySession) StoredAppData(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _DaoRegistry.Contract.StoredAppData(&_DaoRegistry.TransactOpts, activityFlag, data)
}

// StoredAppData is a paid mutator transaction binding the contract method 0x57e65538.
//
// Solidity: function storedAppData(uint256 activityFlag, bytes data) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) StoredAppData(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _DaoRegistry.Contract.StoredAppData(&_DaoRegistry.TransactOpts, activityFlag, data)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xef1008d1.
//
// Solidity: function submitProposal((uint256,address,string,string,string,string,uint256,uint16,uint256) proposal) returns()
func (_DaoRegistry *DaoRegistryTransactor) SubmitProposal(opts *bind.TransactOpts, proposal ProposalSummary) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "submitProposal", proposal)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xef1008d1.
//
// Solidity: function submitProposal((uint256,address,string,string,string,string,uint256,uint16,uint256) proposal) returns()
func (_DaoRegistry *DaoRegistrySession) SubmitProposal(proposal ProposalSummary) (*types.Transaction, error) {
	return _DaoRegistry.Contract.SubmitProposal(&_DaoRegistry.TransactOpts, proposal)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xef1008d1.
//
// Solidity: function submitProposal((uint256,address,string,string,string,string,uint256,uint16,uint256) proposal) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) SubmitProposal(proposal ProposalSummary) (*types.Transaction, error) {
	return _DaoRegistry.Contract.SubmitProposal(&_DaoRegistry.TransactOpts, proposal)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newRegistry) returns()
func (_DaoRegistry *DaoRegistryTransactor) Upgrade(opts *bind.TransactOpts, newRegistry common.Address) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "upgrade", newRegistry)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newRegistry) returns()
func (_DaoRegistry *DaoRegistrySession) Upgrade(newRegistry common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.Upgrade(&_DaoRegistry.TransactOpts, newRegistry)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newRegistry) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) Upgrade(newRegistry common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.Upgrade(&_DaoRegistry.TransactOpts, newRegistry)
}

// UpgradeExtension is a paid mutator transaction binding the contract method 0x1f388f7c.
//
// Solidity: function upgradeExtension((string,uint256,address,uint256,uint256,uint8) newExt) returns()
func (_DaoRegistry *DaoRegistryTransactor) UpgradeExtension(opts *bind.TransactOpts, newExt RegisteredExtension) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "upgradeExtension", newExt)
}

// UpgradeExtension is a paid mutator transaction binding the contract method 0x1f388f7c.
//
// Solidity: function upgradeExtension((string,uint256,address,uint256,uint256,uint8) newExt) returns()
func (_DaoRegistry *DaoRegistrySession) UpgradeExtension(newExt RegisteredExtension) (*types.Transaction, error) {
	return _DaoRegistry.Contract.UpgradeExtension(&_DaoRegistry.TransactOpts, newExt)
}

// UpgradeExtension is a paid mutator transaction binding the contract method 0x1f388f7c.
//
// Solidity: function upgradeExtension((string,uint256,address,uint256,uint256,uint8) newExt) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) UpgradeExtension(newExt RegisteredExtension) (*types.Transaction, error) {
	return _DaoRegistry.Contract.UpgradeExtension(&_DaoRegistry.TransactOpts, newExt)
}

// UpgradeHook is a paid mutator transaction binding the contract method 0x1d261060.
//
// Solidity: function upgradeHook(address state) returns()
func (_DaoRegistry *DaoRegistryTransactor) UpgradeHook(opts *bind.TransactOpts, state common.Address) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "upgradeHook", state)
}

// UpgradeHook is a paid mutator transaction binding the contract method 0x1d261060.
//
// Solidity: function upgradeHook(address state) returns()
func (_DaoRegistry *DaoRegistrySession) UpgradeHook(state common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.UpgradeHook(&_DaoRegistry.TransactOpts, state)
}

// UpgradeHook is a paid mutator transaction binding the contract method 0x1d261060.
//
// Solidity: function upgradeHook(address state) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) UpgradeHook(state common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.UpgradeHook(&_DaoRegistry.TransactOpts, state)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55e4e759.
//
// Solidity: function voteProposal(string proposalAppName, uint256 proposalId) returns()
func (_DaoRegistry *DaoRegistryTransactor) VoteProposal(opts *bind.TransactOpts, proposalAppName string, proposalId *big.Int) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "voteProposal", proposalAppName, proposalId)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55e4e759.
//
// Solidity: function voteProposal(string proposalAppName, uint256 proposalId) returns()
func (_DaoRegistry *DaoRegistrySession) VoteProposal(proposalAppName string, proposalId *big.Int) (*types.Transaction, error) {
	return _DaoRegistry.Contract.VoteProposal(&_DaoRegistry.TransactOpts, proposalAppName, proposalId)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55e4e759.
//
// Solidity: function voteProposal(string proposalAppName, uint256 proposalId) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) VoteProposal(proposalAppName string, proposalId *big.Int) (*types.Transaction, error) {
	return _DaoRegistry.Contract.VoteProposal(&_DaoRegistry.TransactOpts, proposalAppName, proposalId)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x96131049.
//
// Solidity: function withdrawal(address addr) returns()
func (_DaoRegistry *DaoRegistryTransactor) Withdrawal(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DaoRegistry.contract.Transact(opts, "withdrawal", addr)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x96131049.
//
// Solidity: function withdrawal(address addr) returns()
func (_DaoRegistry *DaoRegistrySession) Withdrawal(addr common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.Withdrawal(&_DaoRegistry.TransactOpts, addr)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x96131049.
//
// Solidity: function withdrawal(address addr) returns()
func (_DaoRegistry *DaoRegistryTransactorSession) Withdrawal(addr common.Address) (*types.Transaction, error) {
	return _DaoRegistry.Contract.Withdrawal(&_DaoRegistry.TransactOpts, addr)
}

// DaoRegistryMetaTransactionExecutedIterator is returned from FilterMetaTransactionExecuted and is used to iterate over the raw logs and unpacked data for MetaTransactionExecuted events raised by the DaoRegistry contract.
type DaoRegistryMetaTransactionExecutedIterator struct {
	Event *DaoRegistryMetaTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *DaoRegistryMetaTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryMetaTransactionExecuted)
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
		it.Event = new(DaoRegistryMetaTransactionExecuted)
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
func (it *DaoRegistryMetaTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryMetaTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryMetaTransactionExecuted represents a MetaTransactionExecuted event raised by the DaoRegistry contract.
type DaoRegistryMetaTransactionExecuted struct {
	UserAddress       common.Address
	RelayerAddress    common.Address
	FunctionSignature []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMetaTransactionExecuted is a free log retrieval operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_DaoRegistry *DaoRegistryFilterer) FilterMetaTransactionExecuted(opts *bind.FilterOpts) (*DaoRegistryMetaTransactionExecutedIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryMetaTransactionExecutedIterator{contract: _DaoRegistry.contract, event: "MetaTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchMetaTransactionExecuted is a free log subscription operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_DaoRegistry *DaoRegistryFilterer) WatchMetaTransactionExecuted(opts *bind.WatchOpts, sink chan<- *DaoRegistryMetaTransactionExecuted) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryMetaTransactionExecuted)
				if err := _DaoRegistry.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
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

// ParseMetaTransactionExecuted is a log parse operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_DaoRegistry *DaoRegistryFilterer) ParseMetaTransactionExecuted(log types.Log) (*DaoRegistryMetaTransactionExecuted, error) {
	event := new(DaoRegistryMetaTransactionExecuted)
	if err := _DaoRegistry.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryAddressEventIterator is returned from FilterAddressEvent and is used to iterate over the raw logs and unpacked data for AddressEvent events raised by the DaoRegistry contract.
type DaoRegistryAddressEventIterator struct {
	Event *DaoRegistryAddressEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryAddressEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryAddressEvent)
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
		it.Event = new(DaoRegistryAddressEvent)
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
func (it *DaoRegistryAddressEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryAddressEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryAddressEvent represents a AddressEvent event raised by the DaoRegistry contract.
type DaoRegistryAddressEvent struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressEvent is a free log retrieval operation binding the contract event 0x80b9c575d6f7a1b5952b944d4757a95fa037d4101e21beda363682b355041615.
//
// Solidity: event addressEvent(address addr)
func (_DaoRegistry *DaoRegistryFilterer) FilterAddressEvent(opts *bind.FilterOpts) (*DaoRegistryAddressEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "addressEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryAddressEventIterator{contract: _DaoRegistry.contract, event: "addressEvent", logs: logs, sub: sub}, nil
}

// WatchAddressEvent is a free log subscription operation binding the contract event 0x80b9c575d6f7a1b5952b944d4757a95fa037d4101e21beda363682b355041615.
//
// Solidity: event addressEvent(address addr)
func (_DaoRegistry *DaoRegistryFilterer) WatchAddressEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryAddressEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "addressEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryAddressEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "addressEvent", log); err != nil {
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

// ParseAddressEvent is a log parse operation binding the contract event 0x80b9c575d6f7a1b5952b944d4757a95fa037d4101e21beda363682b355041615.
//
// Solidity: event addressEvent(address addr)
func (_DaoRegistry *DaoRegistryFilterer) ParseAddressEvent(log types.Log) (*DaoRegistryAddressEvent, error) {
	event := new(DaoRegistryAddressEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "addressEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryChangeActivityHookEventIterator is returned from FilterChangeActivityHookEvent and is used to iterate over the raw logs and unpacked data for ChangeActivityHookEvent events raised by the DaoRegistry contract.
type DaoRegistryChangeActivityHookEventIterator struct {
	Event *DaoRegistryChangeActivityHookEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryChangeActivityHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryChangeActivityHookEvent)
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
		it.Event = new(DaoRegistryChangeActivityHookEvent)
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
func (it *DaoRegistryChangeActivityHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryChangeActivityHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryChangeActivityHookEvent represents a ChangeActivityHookEvent event raised by the DaoRegistry contract.
type DaoRegistryChangeActivityHookEvent struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChangeActivityHookEvent is a free log retrieval operation binding the contract event 0x6975dc8be73a61254519c7b7f288900a594956872fb85607085dc0029126475a.
//
// Solidity: event changeActivityHookEvent(address addr)
func (_DaoRegistry *DaoRegistryFilterer) FilterChangeActivityHookEvent(opts *bind.FilterOpts) (*DaoRegistryChangeActivityHookEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "changeActivityHookEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryChangeActivityHookEventIterator{contract: _DaoRegistry.contract, event: "changeActivityHookEvent", logs: logs, sub: sub}, nil
}

// WatchChangeActivityHookEvent is a free log subscription operation binding the contract event 0x6975dc8be73a61254519c7b7f288900a594956872fb85607085dc0029126475a.
//
// Solidity: event changeActivityHookEvent(address addr)
func (_DaoRegistry *DaoRegistryFilterer) WatchChangeActivityHookEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryChangeActivityHookEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "changeActivityHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryChangeActivityHookEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "changeActivityHookEvent", log); err != nil {
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

// ParseChangeActivityHookEvent is a log parse operation binding the contract event 0x6975dc8be73a61254519c7b7f288900a594956872fb85607085dc0029126475a.
//
// Solidity: event changeActivityHookEvent(address addr)
func (_DaoRegistry *DaoRegistryFilterer) ParseChangeActivityHookEvent(log types.Log) (*DaoRegistryChangeActivityHookEvent, error) {
	event := new(DaoRegistryChangeActivityHookEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "changeActivityHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryChangeExtensionEventIterator is returned from FilterChangeExtensionEvent and is used to iterate over the raw logs and unpacked data for ChangeExtensionEvent events raised by the DaoRegistry contract.
type DaoRegistryChangeExtensionEventIterator struct {
	Event *DaoRegistryChangeExtensionEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryChangeExtensionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryChangeExtensionEvent)
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
		it.Event = new(DaoRegistryChangeExtensionEvent)
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
func (it *DaoRegistryChangeExtensionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryChangeExtensionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryChangeExtensionEvent represents a ChangeExtensionEvent event raised by the DaoRegistry contract.
type DaoRegistryChangeExtensionEvent struct {
	Ext      RegisteredExtension
	IsDelete bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeExtensionEvent is a free log retrieval operation binding the contract event 0x044b3aafcefd1ba9c05bc331c56e264426c07ba48d5cd314c573bd6a1bbdb0c6.
//
// Solidity: event changeExtensionEvent((string,uint256,address,uint256,uint256,uint8) ext, bool isDelete)
func (_DaoRegistry *DaoRegistryFilterer) FilterChangeExtensionEvent(opts *bind.FilterOpts) (*DaoRegistryChangeExtensionEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "changeExtensionEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryChangeExtensionEventIterator{contract: _DaoRegistry.contract, event: "changeExtensionEvent", logs: logs, sub: sub}, nil
}

// WatchChangeExtensionEvent is a free log subscription operation binding the contract event 0x044b3aafcefd1ba9c05bc331c56e264426c07ba48d5cd314c573bd6a1bbdb0c6.
//
// Solidity: event changeExtensionEvent((string,uint256,address,uint256,uint256,uint8) ext, bool isDelete)
func (_DaoRegistry *DaoRegistryFilterer) WatchChangeExtensionEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryChangeExtensionEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "changeExtensionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryChangeExtensionEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "changeExtensionEvent", log); err != nil {
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

// ParseChangeExtensionEvent is a log parse operation binding the contract event 0x044b3aafcefd1ba9c05bc331c56e264426c07ba48d5cd314c573bd6a1bbdb0c6.
//
// Solidity: event changeExtensionEvent((string,uint256,address,uint256,uint256,uint8) ext, bool isDelete)
func (_DaoRegistry *DaoRegistryFilterer) ParseChangeExtensionEvent(log types.Log) (*DaoRegistryChangeExtensionEvent, error) {
	event := new(DaoRegistryChangeExtensionEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "changeExtensionEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryChangeInitialDataEventIterator is returned from FilterChangeInitialDataEvent and is used to iterate over the raw logs and unpacked data for ChangeInitialDataEvent events raised by the DaoRegistry contract.
type DaoRegistryChangeInitialDataEventIterator struct {
	Event *DaoRegistryChangeInitialDataEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryChangeInitialDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryChangeInitialDataEvent)
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
		it.Event = new(DaoRegistryChangeInitialDataEvent)
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
func (it *DaoRegistryChangeInitialDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryChangeInitialDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryChangeInitialDataEvent represents a ChangeInitialDataEvent event raised by the DaoRegistry contract.
type DaoRegistryChangeInitialDataEvent struct {
	Name              string
	Value             string
	DaoManagerAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChangeInitialDataEvent is a free log retrieval operation binding the contract event 0x4a5fda6b7feb3c529c4eec78b1eb7022705b12ce800bbd5116112ec06ed2e0eb.
//
// Solidity: event changeInitialDataEvent(string name, string value, address daoManagerAddress)
func (_DaoRegistry *DaoRegistryFilterer) FilterChangeInitialDataEvent(opts *bind.FilterOpts) (*DaoRegistryChangeInitialDataEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryChangeInitialDataEventIterator{contract: _DaoRegistry.contract, event: "changeInitialDataEvent", logs: logs, sub: sub}, nil
}

// WatchChangeInitialDataEvent is a free log subscription operation binding the contract event 0x4a5fda6b7feb3c529c4eec78b1eb7022705b12ce800bbd5116112ec06ed2e0eb.
//
// Solidity: event changeInitialDataEvent(string name, string value, address daoManagerAddress)
func (_DaoRegistry *DaoRegistryFilterer) WatchChangeInitialDataEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryChangeInitialDataEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryChangeInitialDataEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
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

// ParseChangeInitialDataEvent is a log parse operation binding the contract event 0x4a5fda6b7feb3c529c4eec78b1eb7022705b12ce800bbd5116112ec06ed2e0eb.
//
// Solidity: event changeInitialDataEvent(string name, string value, address daoManagerAddress)
func (_DaoRegistry *DaoRegistryFilterer) ParseChangeInitialDataEvent(log types.Log) (*DaoRegistryChangeInitialDataEvent, error) {
	event := new(DaoRegistryChangeInitialDataEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryChangeProposalEventIterator is returned from FilterChangeProposalEvent and is used to iterate over the raw logs and unpacked data for ChangeProposalEvent events raised by the DaoRegistry contract.
type DaoRegistryChangeProposalEventIterator struct {
	Event *DaoRegistryChangeProposalEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryChangeProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryChangeProposalEvent)
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
		it.Event = new(DaoRegistryChangeProposalEvent)
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
func (it *DaoRegistryChangeProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryChangeProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryChangeProposalEvent represents a ChangeProposalEvent event raised by the DaoRegistry contract.
type DaoRegistryChangeProposalEvent struct {
	ProposalAppName string
	ProposalId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChangeProposalEvent is a free log retrieval operation binding the contract event 0xd9d44f830057eef403b108eae33b230dd4785675e9a85291086e65c363f69e45.
//
// Solidity: event changeProposalEvent(string proposalAppName, uint256 proposalId)
func (_DaoRegistry *DaoRegistryFilterer) FilterChangeProposalEvent(opts *bind.FilterOpts) (*DaoRegistryChangeProposalEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "changeProposalEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryChangeProposalEventIterator{contract: _DaoRegistry.contract, event: "changeProposalEvent", logs: logs, sub: sub}, nil
}

// WatchChangeProposalEvent is a free log subscription operation binding the contract event 0xd9d44f830057eef403b108eae33b230dd4785675e9a85291086e65c363f69e45.
//
// Solidity: event changeProposalEvent(string proposalAppName, uint256 proposalId)
func (_DaoRegistry *DaoRegistryFilterer) WatchChangeProposalEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryChangeProposalEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "changeProposalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryChangeProposalEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "changeProposalEvent", log); err != nil {
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

// ParseChangeProposalEvent is a log parse operation binding the contract event 0xd9d44f830057eef403b108eae33b230dd4785675e9a85291086e65c363f69e45.
//
// Solidity: event changeProposalEvent(string proposalAppName, uint256 proposalId)
func (_DaoRegistry *DaoRegistryFilterer) ParseChangeProposalEvent(log types.Log) (*DaoRegistryChangeProposalEvent, error) {
	event := new(DaoRegistryChangeProposalEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "changeProposalEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryDeployExtensionEventIterator is returned from FilterDeployExtensionEvent and is used to iterate over the raw logs and unpacked data for DeployExtensionEvent events raised by the DaoRegistry contract.
type DaoRegistryDeployExtensionEventIterator struct {
	Event *DaoRegistryDeployExtensionEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryDeployExtensionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryDeployExtensionEvent)
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
		it.Event = new(DaoRegistryDeployExtensionEvent)
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
func (it *DaoRegistryDeployExtensionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryDeployExtensionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryDeployExtensionEvent represents a DeployExtensionEvent event raised by the DaoRegistry contract.
type DaoRegistryDeployExtensionEvent struct {
	ExtName          string
	Params           DeploymentParameterV1
	StateDeployCheck bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeployExtensionEvent is a free log retrieval operation binding the contract event 0x2cbff4a1a22f8d98aa9efb71f99141c55b23b623e5f7662b5950ef20a5af9d65.
//
// Solidity: event deployExtensionEvent(string extName, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, bool stateDeployCheck)
func (_DaoRegistry *DaoRegistryFilterer) FilterDeployExtensionEvent(opts *bind.FilterOpts) (*DaoRegistryDeployExtensionEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "deployExtensionEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryDeployExtensionEventIterator{contract: _DaoRegistry.contract, event: "deployExtensionEvent", logs: logs, sub: sub}, nil
}

// WatchDeployExtensionEvent is a free log subscription operation binding the contract event 0x2cbff4a1a22f8d98aa9efb71f99141c55b23b623e5f7662b5950ef20a5af9d65.
//
// Solidity: event deployExtensionEvent(string extName, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, bool stateDeployCheck)
func (_DaoRegistry *DaoRegistryFilterer) WatchDeployExtensionEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryDeployExtensionEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "deployExtensionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryDeployExtensionEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "deployExtensionEvent", log); err != nil {
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

// ParseDeployExtensionEvent is a log parse operation binding the contract event 0x2cbff4a1a22f8d98aa9efb71f99141c55b23b623e5f7662b5950ef20a5af9d65.
//
// Solidity: event deployExtensionEvent(string extName, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, bool stateDeployCheck)
func (_DaoRegistry *DaoRegistryFilterer) ParseDeployExtensionEvent(log types.Log) (*DaoRegistryDeployExtensionEvent, error) {
	event := new(DaoRegistryDeployExtensionEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "deployExtensionEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryDeployExtensionStateEventIterator is returned from FilterDeployExtensionStateEvent and is used to iterate over the raw logs and unpacked data for DeployExtensionStateEvent events raised by the DaoRegistry contract.
type DaoRegistryDeployExtensionStateEventIterator struct {
	Event *DaoRegistryDeployExtensionStateEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryDeployExtensionStateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryDeployExtensionStateEvent)
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
		it.Event = new(DaoRegistryDeployExtensionStateEvent)
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
func (it *DaoRegistryDeployExtensionStateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryDeployExtensionStateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryDeployExtensionStateEvent represents a DeployExtensionStateEvent event raised by the DaoRegistry contract.
type DaoRegistryDeployExtensionStateEvent struct {
	ExtName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeployExtensionStateEvent is a free log retrieval operation binding the contract event 0xf7a0f1a2094d678b6d82a255598859e9aea7822f22689999784e28934b005548.
//
// Solidity: event deployExtensionStateEvent(string extName)
func (_DaoRegistry *DaoRegistryFilterer) FilterDeployExtensionStateEvent(opts *bind.FilterOpts) (*DaoRegistryDeployExtensionStateEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "deployExtensionStateEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryDeployExtensionStateEventIterator{contract: _DaoRegistry.contract, event: "deployExtensionStateEvent", logs: logs, sub: sub}, nil
}

// WatchDeployExtensionStateEvent is a free log subscription operation binding the contract event 0xf7a0f1a2094d678b6d82a255598859e9aea7822f22689999784e28934b005548.
//
// Solidity: event deployExtensionStateEvent(string extName)
func (_DaoRegistry *DaoRegistryFilterer) WatchDeployExtensionStateEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryDeployExtensionStateEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "deployExtensionStateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryDeployExtensionStateEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "deployExtensionStateEvent", log); err != nil {
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

// ParseDeployExtensionStateEvent is a log parse operation binding the contract event 0xf7a0f1a2094d678b6d82a255598859e9aea7822f22689999784e28934b005548.
//
// Solidity: event deployExtensionStateEvent(string extName)
func (_DaoRegistry *DaoRegistryFilterer) ParseDeployExtensionStateEvent(log types.Log) (*DaoRegistryDeployExtensionStateEvent, error) {
	event := new(DaoRegistryDeployExtensionStateEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "deployExtensionStateEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryDeployStandardContractEventIterator is returned from FilterDeployStandardContractEvent and is used to iterate over the raw logs and unpacked data for DeployStandardContractEvent events raised by the DaoRegistry contract.
type DaoRegistryDeployStandardContractEventIterator struct {
	Event *DaoRegistryDeployStandardContractEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryDeployStandardContractEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryDeployStandardContractEvent)
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
		it.Event = new(DaoRegistryDeployStandardContractEvent)
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
func (it *DaoRegistryDeployStandardContractEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryDeployStandardContractEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryDeployStandardContractEvent represents a DeployStandardContractEvent event raised by the DaoRegistry contract.
type DaoRegistryDeployStandardContractEvent struct {
	DepName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeployStandardContractEvent is a free log retrieval operation binding the contract event 0xcd7374809e92e6c83a5c99dc3fce9f817496813960940d54665b25fc51f9ab40.
//
// Solidity: event deployStandardContractEvent(string depName)
func (_DaoRegistry *DaoRegistryFilterer) FilterDeployStandardContractEvent(opts *bind.FilterOpts) (*DaoRegistryDeployStandardContractEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "deployStandardContractEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryDeployStandardContractEventIterator{contract: _DaoRegistry.contract, event: "deployStandardContractEvent", logs: logs, sub: sub}, nil
}

// WatchDeployStandardContractEvent is a free log subscription operation binding the contract event 0xcd7374809e92e6c83a5c99dc3fce9f817496813960940d54665b25fc51f9ab40.
//
// Solidity: event deployStandardContractEvent(string depName)
func (_DaoRegistry *DaoRegistryFilterer) WatchDeployStandardContractEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryDeployStandardContractEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "deployStandardContractEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryDeployStandardContractEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "deployStandardContractEvent", log); err != nil {
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

// ParseDeployStandardContractEvent is a log parse operation binding the contract event 0xcd7374809e92e6c83a5c99dc3fce9f817496813960940d54665b25fc51f9ab40.
//
// Solidity: event deployStandardContractEvent(string depName)
func (_DaoRegistry *DaoRegistryFilterer) ParseDeployStandardContractEvent(log types.Log) (*DaoRegistryDeployStandardContractEvent, error) {
	event := new(DaoRegistryDeployStandardContractEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "deployStandardContractEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistrySetDaoFactoryAddressEventIterator is returned from FilterSetDaoFactoryAddressEvent and is used to iterate over the raw logs and unpacked data for SetDaoFactoryAddressEvent events raised by the DaoRegistry contract.
type DaoRegistrySetDaoFactoryAddressEventIterator struct {
	Event *DaoRegistrySetDaoFactoryAddressEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistrySetDaoFactoryAddressEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistrySetDaoFactoryAddressEvent)
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
		it.Event = new(DaoRegistrySetDaoFactoryAddressEvent)
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
func (it *DaoRegistrySetDaoFactoryAddressEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistrySetDaoFactoryAddressEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistrySetDaoFactoryAddressEvent represents a SetDaoFactoryAddressEvent event raised by the DaoRegistry contract.
type DaoRegistrySetDaoFactoryAddressEvent struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetDaoFactoryAddressEvent is a free log retrieval operation binding the contract event 0x6d6b0d54ff761f9657cbf411a1a8836202e9a0f0a46c352223a8e635251e82a5.
//
// Solidity: event setDaoFactoryAddressEvent(address factory)
func (_DaoRegistry *DaoRegistryFilterer) FilterSetDaoFactoryAddressEvent(opts *bind.FilterOpts) (*DaoRegistrySetDaoFactoryAddressEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "setDaoFactoryAddressEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistrySetDaoFactoryAddressEventIterator{contract: _DaoRegistry.contract, event: "setDaoFactoryAddressEvent", logs: logs, sub: sub}, nil
}

// WatchSetDaoFactoryAddressEvent is a free log subscription operation binding the contract event 0x6d6b0d54ff761f9657cbf411a1a8836202e9a0f0a46c352223a8e635251e82a5.
//
// Solidity: event setDaoFactoryAddressEvent(address factory)
func (_DaoRegistry *DaoRegistryFilterer) WatchSetDaoFactoryAddressEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistrySetDaoFactoryAddressEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "setDaoFactoryAddressEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistrySetDaoFactoryAddressEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "setDaoFactoryAddressEvent", log); err != nil {
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

// ParseSetDaoFactoryAddressEvent is a log parse operation binding the contract event 0x6d6b0d54ff761f9657cbf411a1a8836202e9a0f0a46c352223a8e635251e82a5.
//
// Solidity: event setDaoFactoryAddressEvent(address factory)
func (_DaoRegistry *DaoRegistryFilterer) ParseSetDaoFactoryAddressEvent(log types.Log) (*DaoRegistrySetDaoFactoryAddressEvent, error) {
	event := new(DaoRegistrySetDaoFactoryAddressEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "setDaoFactoryAddressEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistrySetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the DaoRegistry contract.
type DaoRegistrySetReadyEventIterator struct {
	Event *DaoRegistrySetReadyEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistrySetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistrySetReadyEvent)
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
		it.Event = new(DaoRegistrySetReadyEvent)
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
func (it *DaoRegistrySetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistrySetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistrySetReadyEvent represents a SetReadyEvent event raised by the DaoRegistry contract.
type DaoRegistrySetReadyEvent struct {
	Check bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool check)
func (_DaoRegistry *DaoRegistryFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*DaoRegistrySetReadyEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistrySetReadyEventIterator{contract: _DaoRegistry.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool check)
func (_DaoRegistry *DaoRegistryFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistrySetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistrySetReadyEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
// Solidity: event setReadyEvent(bool check)
func (_DaoRegistry *DaoRegistryFilterer) ParseSetReadyEvent(log types.Log) (*DaoRegistrySetReadyEvent, error) {
	event := new(DaoRegistrySetReadyEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryStoredAppDataEventIterator is returned from FilterStoredAppDataEvent and is used to iterate over the raw logs and unpacked data for StoredAppDataEvent events raised by the DaoRegistry contract.
type DaoRegistryStoredAppDataEventIterator struct {
	Event *DaoRegistryStoredAppDataEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryStoredAppDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryStoredAppDataEvent)
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
		it.Event = new(DaoRegistryStoredAppDataEvent)
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
func (it *DaoRegistryStoredAppDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryStoredAppDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryStoredAppDataEvent represents a StoredAppDataEvent event raised by the DaoRegistry contract.
type DaoRegistryStoredAppDataEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStoredAppDataEvent is a free log retrieval operation binding the contract event 0xf54e1e77b326c519703e2e17fea0bb6600009edd9597e65edd161c9960094862.
//
// Solidity: event storedAppDataEvent(uint256 activityFlag, bytes data)
func (_DaoRegistry *DaoRegistryFilterer) FilterStoredAppDataEvent(opts *bind.FilterOpts) (*DaoRegistryStoredAppDataEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "storedAppDataEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryStoredAppDataEventIterator{contract: _DaoRegistry.contract, event: "storedAppDataEvent", logs: logs, sub: sub}, nil
}

// WatchStoredAppDataEvent is a free log subscription operation binding the contract event 0xf54e1e77b326c519703e2e17fea0bb6600009edd9597e65edd161c9960094862.
//
// Solidity: event storedAppDataEvent(uint256 activityFlag, bytes data)
func (_DaoRegistry *DaoRegistryFilterer) WatchStoredAppDataEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryStoredAppDataEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "storedAppDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryStoredAppDataEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "storedAppDataEvent", log); err != nil {
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

// ParseStoredAppDataEvent is a log parse operation binding the contract event 0xf54e1e77b326c519703e2e17fea0bb6600009edd9597e65edd161c9960094862.
//
// Solidity: event storedAppDataEvent(uint256 activityFlag, bytes data)
func (_DaoRegistry *DaoRegistryFilterer) ParseStoredAppDataEvent(log types.Log) (*DaoRegistryStoredAppDataEvent, error) {
	event := new(DaoRegistryStoredAppDataEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "storedAppDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryUpgradeEventIterator is returned from FilterUpgradeEvent and is used to iterate over the raw logs and unpacked data for UpgradeEvent events raised by the DaoRegistry contract.
type DaoRegistryUpgradeEventIterator struct {
	Event *DaoRegistryUpgradeEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryUpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryUpgradeEvent)
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
		it.Event = new(DaoRegistryUpgradeEvent)
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
func (it *DaoRegistryUpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryUpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryUpgradeEvent represents a UpgradeEvent event raised by the DaoRegistry contract.
type DaoRegistryUpgradeEvent struct {
	NewRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpgradeEvent is a free log retrieval operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newRegistry)
func (_DaoRegistry *DaoRegistryFilterer) FilterUpgradeEvent(opts *bind.FilterOpts) (*DaoRegistryUpgradeEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryUpgradeEventIterator{contract: _DaoRegistry.contract, event: "upgradeEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeEvent is a free log subscription operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newRegistry)
func (_DaoRegistry *DaoRegistryFilterer) WatchUpgradeEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryUpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryUpgradeEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
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
// Solidity: event upgradeEvent(address newRegistry)
func (_DaoRegistry *DaoRegistryFilterer) ParseUpgradeEvent(log types.Log) (*DaoRegistryUpgradeEvent, error) {
	event := new(DaoRegistryUpgradeEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryUpgradeHookEventIterator is returned from FilterUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for UpgradeHookEvent events raised by the DaoRegistry contract.
type DaoRegistryUpgradeHookEventIterator struct {
	Event *DaoRegistryUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryUpgradeHookEvent)
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
		it.Event = new(DaoRegistryUpgradeHookEvent)
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
func (it *DaoRegistryUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryUpgradeHookEvent represents a UpgradeHookEvent event raised by the DaoRegistry contract.
type DaoRegistryUpgradeHookEvent struct {
	State common.Address
	Ready bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgradeHookEvent is a free log retrieval operation binding the contract event 0xa95ae080cb61f0eb5969d5d0e0b62b35fb591abc2731c20a81485b2f4f08595a.
//
// Solidity: event upgradeHookEvent(address state, bool ready)
func (_DaoRegistry *DaoRegistryFilterer) FilterUpgradeHookEvent(opts *bind.FilterOpts) (*DaoRegistryUpgradeHookEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "upgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryUpgradeHookEventIterator{contract: _DaoRegistry.contract, event: "upgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeHookEvent is a free log subscription operation binding the contract event 0xa95ae080cb61f0eb5969d5d0e0b62b35fb591abc2731c20a81485b2f4f08595a.
//
// Solidity: event upgradeHookEvent(address state, bool ready)
func (_DaoRegistry *DaoRegistryFilterer) WatchUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "upgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryUpgradeHookEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "upgradeHookEvent", log); err != nil {
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

// ParseUpgradeHookEvent is a log parse operation binding the contract event 0xa95ae080cb61f0eb5969d5d0e0b62b35fb591abc2731c20a81485b2f4f08595a.
//
// Solidity: event upgradeHookEvent(address state, bool ready)
func (_DaoRegistry *DaoRegistryFilterer) ParseUpgradeHookEvent(log types.Log) (*DaoRegistryUpgradeHookEvent, error) {
	event := new(DaoRegistryUpgradeHookEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "upgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoRegistryWithdrawalEventIterator is returned from FilterWithdrawalEvent and is used to iterate over the raw logs and unpacked data for WithdrawalEvent events raised by the DaoRegistry contract.
type DaoRegistryWithdrawalEventIterator struct {
	Event *DaoRegistryWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *DaoRegistryWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRegistryWithdrawalEvent)
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
		it.Event = new(DaoRegistryWithdrawalEvent)
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
func (it *DaoRegistryWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRegistryWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRegistryWithdrawalEvent represents a WithdrawalEvent event raised by the DaoRegistry contract.
type DaoRegistryWithdrawalEvent struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalEvent is a free log retrieval operation binding the contract event 0x881406214596259b5d24cf97917f72ed1f2081c2e45799ab7e56e24ba38be1c5.
//
// Solidity: event withdrawalEvent(address addr)
func (_DaoRegistry *DaoRegistryFilterer) FilterWithdrawalEvent(opts *bind.FilterOpts) (*DaoRegistryWithdrawalEventIterator, error) {

	logs, sub, err := _DaoRegistry.contract.FilterLogs(opts, "withdrawalEvent")
	if err != nil {
		return nil, err
	}
	return &DaoRegistryWithdrawalEventIterator{contract: _DaoRegistry.contract, event: "withdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawalEvent is a free log subscription operation binding the contract event 0x881406214596259b5d24cf97917f72ed1f2081c2e45799ab7e56e24ba38be1c5.
//
// Solidity: event withdrawalEvent(address addr)
func (_DaoRegistry *DaoRegistryFilterer) WatchWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *DaoRegistryWithdrawalEvent) (event.Subscription, error) {

	logs, sub, err := _DaoRegistry.contract.WatchLogs(opts, "withdrawalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRegistryWithdrawalEvent)
				if err := _DaoRegistry.contract.UnpackLog(event, "withdrawalEvent", log); err != nil {
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

// ParseWithdrawalEvent is a log parse operation binding the contract event 0x881406214596259b5d24cf97917f72ed1f2081c2e45799ab7e56e24ba38be1c5.
//
// Solidity: event withdrawalEvent(address addr)
func (_DaoRegistry *DaoRegistryFilterer) ParseWithdrawalEvent(log types.Log) (*DaoRegistryWithdrawalEvent, error) {
	event := new(DaoRegistryWithdrawalEvent)
	if err := _DaoRegistry.contract.UnpackLog(event, "withdrawalEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
