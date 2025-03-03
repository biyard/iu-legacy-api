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

// ActivityDescription is an auto generated low-level Go binding around an user-defined struct.
type ActivityDescription struct {
	Flag        string
	Description string
}

// BizTemplate is an auto generated low-level Go binding around an user-defined struct.
type BizTemplate struct {
	Symbol     string
	Revision   *big.Int
	Extensions []ExtensionDeployment
}

// Bytecode is an auto generated low-level Go binding around an user-defined struct.
type Bytecode struct {
	Name     string
	Revision *big.Int
	Bytecode []byte
}

// Dao is an auto generated low-level Go binding around an user-defined struct.
type Dao struct {
	Name        string
	Description string
	Template    string
	Founders    []common.Address
	Registry    common.Address
}

// DaoRequest is an auto generated low-level Go binding around an user-defined struct.
type DaoRequest struct {
	Request     DaoUserRequest
	Approved    bool
	Sender      common.Address
	Processed   *big.Int
	CreationFee *big.Int
	Registry    common.Address
}

// DaoUserRequest is an auto generated low-level Go binding around an user-defined struct.
type DaoUserRequest struct {
	Name        string
	Description string
	Founders    []common.Address
	Template    string
	Extensions  []ExtensionRequest
}

// DeployArgumentDescription is an auto generated low-level Go binding around an user-defined struct.
type DeployArgumentDescription struct {
	Variable   string
	Parameters []string
}

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

// Extension is an auto generated low-level Go binding around an user-defined struct.
type Extension struct {
	Name                    string
	Revision                *big.Int
	Bytecode                []byte
	ReadActivities          []ActivityDescription
	WriteActivities         []ActivityDescription
	DecisionActivities      []ActivityDescription
	NeededPermissions       *big.Int
	HookActivityPermissions *big.Int
	Arguments               []DeployArgumentDescription
	Dependancies            []string
	ExtensionType           uint8
}

// ExtensionDeployment is an auto generated low-level Go binding around an user-defined struct.
type ExtensionDeployment struct {
	Name   string
	Params DeploymentParameterV1
}

// ExtensionRequest is an auto generated low-level Go binding around an user-defined struct.
type ExtensionRequest struct {
	Name       string
	Parameters DeploymentParameterV1
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

// DaoFactoryMetaData contains all meta data concerning the DaoFactory contract.
var DaoFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"state\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"layers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"}],\"internalType\":\"structExtensionDeployment[]\",\"name\":\"extensions\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structBizTemplate\",\"name\":\"ext\",\"type\":\"tuple\"}],\"name\":\"addBussinessTemplateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addContractEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addExtensionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addExtensionStateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"}],\"name\":\"allowQuoteEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFactory\",\"type\":\"address\"}],\"name\":\"beforeUpgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"state\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"layers\",\"type\":\"address[]\"}],\"name\":\"changeInitialDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"daoReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"deployDaoRegistryEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ext\",\"type\":\"address\"}],\"name\":\"deployExtensionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"rejectDaoRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setCreationFeeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setDaoRegistryBytecodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setDaoRegistryStateBytecodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"setLayersEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"storedAppDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"upgradeContractEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"upgradeExtensionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"upgradeExtensionStateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"state\",\"type\":\"address\"}],\"name\":\"upgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"useDeprecatedFunctionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"withdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"}],\"name\":\"addContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDao\",\"name\":\"dao\",\"type\":\"tuple\"}],\"name\":\"addDao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"readActivities\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"writeActivities\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"decisionActivities\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"neededPermissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookActivityPermissions\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"variable\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"parameters\",\"type\":\"string[]\"}],\"internalType\":\"structDeployArgumentDescription[]\",\"name\":\"arguments\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"dependancies\",\"type\":\"string[]\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structExtension\",\"name\":\"ext\",\"type\":\"tuple\"}],\"name\":\"addExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"extName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"}],\"name\":\"addExtensionState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quote\",\"type\":\"uint256\"}],\"name\":\"allowQuote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"approveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tmpl\",\"type\":\"string\"}],\"name\":\"delBizTemplate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"delDao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"extensionName\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"extStateAddr\",\"type\":\"address\"}],\"name\":\"deployExtension\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookPermissions\",\"type\":\"uint256\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structRegisteredExtension\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"stateAddress\",\"type\":\"address\"}],\"name\":\"deployExtensionFromExtension\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookPermissions\",\"type\":\"uint256\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structRegisteredExtension\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"deployExtensionState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"deployExtensionStateFromExtension\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"}],\"name\":\"deployStandardContractFromExtension\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"extension\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"readActivities\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"writeActivities\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"decisionActivities\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"neededPermissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookActivityPermissions\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"variable\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"parameters\",\"type\":\"string[]\"}],\"internalType\":\"structDeployArgumentDescription[]\",\"name\":\"arguments\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"dependancies\",\"type\":\"string[]\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structExtension\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getBizTemplate\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"}],\"internalType\":\"structExtensionDeployment[]\",\"name\":\"extensions\",\"type\":\"tuple[]\"}],\"internalType\":\"structBizTemplate\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"}],\"internalType\":\"structBytecode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getDao\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDao\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getDaoRequest\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"internalType\":\"structExtensionRequest[]\",\"name\":\"extensions\",\"type\":\"tuple[]\"}],\"internalType\":\"structDaoUserRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"processed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDaoRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getPagenationDaos\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDao[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listAllowedQuotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listBizTemplate\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"}],\"internalType\":\"structExtensionDeployment[]\",\"name\":\"extensions\",\"type\":\"tuple[]\"}],\"internalType\":\"structBizTemplate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listDaos\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDao[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listExtensions\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"readActivities\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"writeActivities\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"decisionActivities\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"neededPermissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookActivityPermissions\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"variable\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"parameters\",\"type\":\"string[]\"}],\"internalType\":\"structDeployArgumentDescription[]\",\"name\":\"arguments\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"dependancies\",\"type\":\"string[]\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structExtension[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listRequests\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"internalType\":\"structExtensionRequest[]\",\"name\":\"extensions\",\"type\":\"tuple[]\"}],\"internalType\":\"structDaoUserRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"processed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDaoRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listSenderDaos\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDao[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"listUserDaos\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDao[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"processRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"quoteOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"rejectRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"internalType\":\"structExtensionRequest[]\",\"name\":\"extensions\",\"type\":\"tuple[]\"}],\"internalType\":\"structDaoUserRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"requestDaoCreation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setCreationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"}],\"name\":\"setDaoRegistryBytecode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"}],\"name\":\"setDaoRegistryStateBytecode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"setLayers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"storedAppData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tmpName\",\"type\":\"string\"}],\"name\":\"tmpListDaos\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDao[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"depName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"}],\"name\":\"upgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"template\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"founders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"internalType\":\"structDao\",\"name\":\"dao\",\"type\":\"tuple\"}],\"name\":\"upgradeDao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"readActivities\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"writeActivities\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityDescription[]\",\"name\":\"decisionActivities\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"neededPermissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookActivityPermissions\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"variable\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"parameters\",\"type\":\"string[]\"}],\"internalType\":\"structDeployArgumentDescription[]\",\"name\":\"arguments\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"dependancies\",\"type\":\"string[]\"},{\"internalType\":\"enumExtensionType\",\"name\":\"extensionType\",\"type\":\"uint8\"}],\"internalType\":\"structExtension\",\"name\":\"ext\",\"type\":\"tuple\"}],\"name\":\"upgradeExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"extName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"}],\"name\":\"upgradeExtensionState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"state\",\"type\":\"address\"}],\"name\":\"upgradeHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"}],\"internalType\":\"structExtensionDeployment[]\",\"name\":\"extensions\",\"type\":\"tuple[]\"}],\"internalType\":\"structBizTemplate\",\"name\":\"tmpl\",\"type\":\"tuple\"}],\"name\":\"upsertBizTemplate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DaoFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use DaoFactoryMetaData.ABI instead.
var DaoFactoryABI = DaoFactoryMetaData.ABI

// DaoFactoryBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const DaoFactoryBinRuntime = ``

// DaoFactory is an auto generated Go binding around a Klaytn contract.
type DaoFactory struct {
	DaoFactoryCaller     // Read-only binding to the contract
	DaoFactoryTransactor // Write-only binding to the contract
	DaoFactoryFilterer   // Log filterer for contract events
}

// DaoFactoryCaller is an auto generated read-only Go binding around a Klaytn contract.
type DaoFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoFactoryTransactor is an auto generated write-only Go binding around a Klaytn contract.
type DaoFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoFactoryFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type DaoFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoFactorySession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type DaoFactorySession struct {
	Contract     *DaoFactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoFactoryCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type DaoFactoryCallerSession struct {
	Contract *DaoFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DaoFactoryTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type DaoFactoryTransactorSession struct {
	Contract     *DaoFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DaoFactoryRaw is an auto generated low-level Go binding around a Klaytn contract.
type DaoFactoryRaw struct {
	Contract *DaoFactory // Generic contract binding to access the raw methods on
}

// DaoFactoryCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type DaoFactoryCallerRaw struct {
	Contract *DaoFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DaoFactoryTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type DaoFactoryTransactorRaw struct {
	Contract *DaoFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDaoFactory creates a new instance of DaoFactory, bound to a specific deployed contract.
func NewDaoFactory(address common.Address, backend bind.ContractBackend) (*DaoFactory, error) {
	contract, err := bindDaoFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DaoFactory{DaoFactoryCaller: DaoFactoryCaller{contract: contract}, DaoFactoryTransactor: DaoFactoryTransactor{contract: contract}, DaoFactoryFilterer: DaoFactoryFilterer{contract: contract}}, nil
}

// NewDaoFactoryCaller creates a new read-only instance of DaoFactory, bound to a specific deployed contract.
func NewDaoFactoryCaller(address common.Address, caller bind.ContractCaller) (*DaoFactoryCaller, error) {
	contract, err := bindDaoFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaoFactoryCaller{contract: contract}, nil
}

// NewDaoFactoryTransactor creates a new write-only instance of DaoFactory, bound to a specific deployed contract.
func NewDaoFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DaoFactoryTransactor, error) {
	contract, err := bindDaoFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaoFactoryTransactor{contract: contract}, nil
}

// NewDaoFactoryFilterer creates a new log filterer instance of DaoFactory, bound to a specific deployed contract.
func NewDaoFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DaoFactoryFilterer, error) {
	contract, err := bindDaoFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaoFactoryFilterer{contract: contract}, nil
}

// bindDaoFactory binds a generic wrapper to an already deployed contract.
func bindDaoFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DaoFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoFactory *DaoFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DaoFactory.Contract.DaoFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoFactory *DaoFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoFactory.Contract.DaoFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoFactory *DaoFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoFactory.Contract.DaoFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoFactory *DaoFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DaoFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoFactory *DaoFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoFactory *DaoFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoFactory.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_DaoFactory *DaoFactoryCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_DaoFactory *DaoFactorySession) Balance() (*big.Int, error) {
	return _DaoFactory.Contract.Balance(&_DaoFactory.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_DaoFactory *DaoFactoryCallerSession) Balance() (*big.Int, error) {
	return _DaoFactory.Contract.Balance(&_DaoFactory.CallOpts)
}

// Extension is a free data retrieval call binding the contract method 0x7fbc6d62.
//
// Solidity: function extension(string name) view returns((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8))
func (_DaoFactory *DaoFactoryCaller) Extension(opts *bind.CallOpts, name string) (Extension, error) {
	var (
		ret0 = new(Extension)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "extension", name)
	return *ret0, err
}

// Extension is a free data retrieval call binding the contract method 0x7fbc6d62.
//
// Solidity: function extension(string name) view returns((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8))
func (_DaoFactory *DaoFactorySession) Extension(name string) (Extension, error) {
	return _DaoFactory.Contract.Extension(&_DaoFactory.CallOpts, name)
}

// Extension is a free data retrieval call binding the contract method 0x7fbc6d62.
//
// Solidity: function extension(string name) view returns((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8))
func (_DaoFactory *DaoFactoryCallerSession) Extension(name string) (Extension, error) {
	return _DaoFactory.Contract.Extension(&_DaoFactory.CallOpts, name)
}

// GetBizTemplate is a free data retrieval call binding the contract method 0xa8333e87.
//
// Solidity: function getBizTemplate(string name) view returns((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]), bool)
func (_DaoFactory *DaoFactoryCaller) GetBizTemplate(opts *bind.CallOpts, name string) (BizTemplate, bool, error) {
	var (
		ret0 = new(BizTemplate)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DaoFactory.contract.Call(opts, out, "getBizTemplate", name)
	return *ret0, *ret1, err
}

// GetBizTemplate is a free data retrieval call binding the contract method 0xa8333e87.
//
// Solidity: function getBizTemplate(string name) view returns((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]), bool)
func (_DaoFactory *DaoFactorySession) GetBizTemplate(name string) (BizTemplate, bool, error) {
	return _DaoFactory.Contract.GetBizTemplate(&_DaoFactory.CallOpts, name)
}

// GetBizTemplate is a free data retrieval call binding the contract method 0xa8333e87.
//
// Solidity: function getBizTemplate(string name) view returns((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]), bool)
func (_DaoFactory *DaoFactoryCallerSession) GetBizTemplate(name string) (BizTemplate, bool, error) {
	return _DaoFactory.Contract.GetBizTemplate(&_DaoFactory.CallOpts, name)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string depName) view returns((string,uint256,bytes))
func (_DaoFactory *DaoFactoryCaller) GetContract(opts *bind.CallOpts, depName string) (Bytecode, error) {
	var (
		ret0 = new(Bytecode)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "getContract", depName)
	return *ret0, err
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string depName) view returns((string,uint256,bytes))
func (_DaoFactory *DaoFactorySession) GetContract(depName string) (Bytecode, error) {
	return _DaoFactory.Contract.GetContract(&_DaoFactory.CallOpts, depName)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string depName) view returns((string,uint256,bytes))
func (_DaoFactory *DaoFactoryCallerSession) GetContract(depName string) (Bytecode, error) {
	return _DaoFactory.Contract.GetContract(&_DaoFactory.CallOpts, depName)
}

// GetDao is a free data retrieval call binding the contract method 0xf6934a2d.
//
// Solidity: function getDao(string name) view returns((string,string,string,address[],address), bool)
func (_DaoFactory *DaoFactoryCaller) GetDao(opts *bind.CallOpts, name string) (Dao, bool, error) {
	var (
		ret0 = new(Dao)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DaoFactory.contract.Call(opts, out, "getDao", name)
	return *ret0, *ret1, err
}

// GetDao is a free data retrieval call binding the contract method 0xf6934a2d.
//
// Solidity: function getDao(string name) view returns((string,string,string,address[],address), bool)
func (_DaoFactory *DaoFactorySession) GetDao(name string) (Dao, bool, error) {
	return _DaoFactory.Contract.GetDao(&_DaoFactory.CallOpts, name)
}

// GetDao is a free data retrieval call binding the contract method 0xf6934a2d.
//
// Solidity: function getDao(string name) view returns((string,string,string,address[],address), bool)
func (_DaoFactory *DaoFactoryCallerSession) GetDao(name string) (Dao, bool, error) {
	return _DaoFactory.Contract.GetDao(&_DaoFactory.CallOpts, name)
}

// GetDaoRequest is a free data retrieval call binding the contract method 0xeb0e9681.
//
// Solidity: function getDaoRequest(string name) view returns(((string,string,address[],string,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]),bool,address,uint256,uint256,address))
func (_DaoFactory *DaoFactoryCaller) GetDaoRequest(opts *bind.CallOpts, name string) (DaoRequest, error) {
	var (
		ret0 = new(DaoRequest)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "getDaoRequest", name)
	return *ret0, err
}

// GetDaoRequest is a free data retrieval call binding the contract method 0xeb0e9681.
//
// Solidity: function getDaoRequest(string name) view returns(((string,string,address[],string,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]),bool,address,uint256,uint256,address))
func (_DaoFactory *DaoFactorySession) GetDaoRequest(name string) (DaoRequest, error) {
	return _DaoFactory.Contract.GetDaoRequest(&_DaoFactory.CallOpts, name)
}

// GetDaoRequest is a free data retrieval call binding the contract method 0xeb0e9681.
//
// Solidity: function getDaoRequest(string name) view returns(((string,string,address[],string,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]),bool,address,uint256,uint256,address))
func (_DaoFactory *DaoFactoryCallerSession) GetDaoRequest(name string) (DaoRequest, error) {
	return _DaoFactory.Contract.GetDaoRequest(&_DaoFactory.CallOpts, name)
}

// GetPagenationDaos is a free data retrieval call binding the contract method 0x216cdc53.
//
// Solidity: function getPagenationDaos(uint256 page, uint256 pageSize) view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCaller) GetPagenationDaos(opts *bind.CallOpts, page *big.Int, pageSize *big.Int) ([]Dao, error) {
	var (
		ret0 = new([]Dao)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "getPagenationDaos", page, pageSize)
	return *ret0, err
}

// GetPagenationDaos is a free data retrieval call binding the contract method 0x216cdc53.
//
// Solidity: function getPagenationDaos(uint256 page, uint256 pageSize) view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactorySession) GetPagenationDaos(page *big.Int, pageSize *big.Int) ([]Dao, error) {
	return _DaoFactory.Contract.GetPagenationDaos(&_DaoFactory.CallOpts, page, pageSize)
}

// GetPagenationDaos is a free data retrieval call binding the contract method 0x216cdc53.
//
// Solidity: function getPagenationDaos(uint256 page, uint256 pageSize) view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCallerSession) GetPagenationDaos(page *big.Int, pageSize *big.Int) ([]Dao, error) {
	return _DaoFactory.Contract.GetPagenationDaos(&_DaoFactory.CallOpts, page, pageSize)
}

// ListAllowedQuotes is a free data retrieval call binding the contract method 0x59df49b5.
//
// Solidity: function listAllowedQuotes() view returns(address[], uint256[])
func (_DaoFactory *DaoFactoryCaller) ListAllowedQuotes(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DaoFactory.contract.Call(opts, out, "listAllowedQuotes")
	return *ret0, *ret1, err
}

// ListAllowedQuotes is a free data retrieval call binding the contract method 0x59df49b5.
//
// Solidity: function listAllowedQuotes() view returns(address[], uint256[])
func (_DaoFactory *DaoFactorySession) ListAllowedQuotes() ([]common.Address, []*big.Int, error) {
	return _DaoFactory.Contract.ListAllowedQuotes(&_DaoFactory.CallOpts)
}

// ListAllowedQuotes is a free data retrieval call binding the contract method 0x59df49b5.
//
// Solidity: function listAllowedQuotes() view returns(address[], uint256[])
func (_DaoFactory *DaoFactoryCallerSession) ListAllowedQuotes() ([]common.Address, []*big.Int, error) {
	return _DaoFactory.Contract.ListAllowedQuotes(&_DaoFactory.CallOpts)
}

// ListBizTemplate is a free data retrieval call binding the contract method 0x5953bbdb.
//
// Solidity: function listBizTemplate() view returns((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[])[])
func (_DaoFactory *DaoFactoryCaller) ListBizTemplate(opts *bind.CallOpts) ([]BizTemplate, error) {
	var (
		ret0 = new([]BizTemplate)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "listBizTemplate")
	return *ret0, err
}

// ListBizTemplate is a free data retrieval call binding the contract method 0x5953bbdb.
//
// Solidity: function listBizTemplate() view returns((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[])[])
func (_DaoFactory *DaoFactorySession) ListBizTemplate() ([]BizTemplate, error) {
	return _DaoFactory.Contract.ListBizTemplate(&_DaoFactory.CallOpts)
}

// ListBizTemplate is a free data retrieval call binding the contract method 0x5953bbdb.
//
// Solidity: function listBizTemplate() view returns((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[])[])
func (_DaoFactory *DaoFactoryCallerSession) ListBizTemplate() ([]BizTemplate, error) {
	return _DaoFactory.Contract.ListBizTemplate(&_DaoFactory.CallOpts)
}

// ListDaos is a free data retrieval call binding the contract method 0xd9678af2.
//
// Solidity: function listDaos() view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCaller) ListDaos(opts *bind.CallOpts) ([]Dao, error) {
	var (
		ret0 = new([]Dao)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "listDaos")
	return *ret0, err
}

// ListDaos is a free data retrieval call binding the contract method 0xd9678af2.
//
// Solidity: function listDaos() view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactorySession) ListDaos() ([]Dao, error) {
	return _DaoFactory.Contract.ListDaos(&_DaoFactory.CallOpts)
}

// ListDaos is a free data retrieval call binding the contract method 0xd9678af2.
//
// Solidity: function listDaos() view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCallerSession) ListDaos() ([]Dao, error) {
	return _DaoFactory.Contract.ListDaos(&_DaoFactory.CallOpts)
}

// ListExtensions is a free data retrieval call binding the contract method 0x5e76155f.
//
// Solidity: function listExtensions() view returns((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8)[])
func (_DaoFactory *DaoFactoryCaller) ListExtensions(opts *bind.CallOpts) ([]Extension, error) {
	var (
		ret0 = new([]Extension)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "listExtensions")
	return *ret0, err
}

// ListExtensions is a free data retrieval call binding the contract method 0x5e76155f.
//
// Solidity: function listExtensions() view returns((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8)[])
func (_DaoFactory *DaoFactorySession) ListExtensions() ([]Extension, error) {
	return _DaoFactory.Contract.ListExtensions(&_DaoFactory.CallOpts)
}

// ListExtensions is a free data retrieval call binding the contract method 0x5e76155f.
//
// Solidity: function listExtensions() view returns((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8)[])
func (_DaoFactory *DaoFactoryCallerSession) ListExtensions() ([]Extension, error) {
	return _DaoFactory.Contract.ListExtensions(&_DaoFactory.CallOpts)
}

// ListRequests is a free data retrieval call binding the contract method 0x53d5b816.
//
// Solidity: function listRequests() view returns(((string,string,address[],string,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]),bool,address,uint256,uint256,address)[])
func (_DaoFactory *DaoFactoryCaller) ListRequests(opts *bind.CallOpts) ([]DaoRequest, error) {
	var (
		ret0 = new([]DaoRequest)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "listRequests")
	return *ret0, err
}

// ListRequests is a free data retrieval call binding the contract method 0x53d5b816.
//
// Solidity: function listRequests() view returns(((string,string,address[],string,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]),bool,address,uint256,uint256,address)[])
func (_DaoFactory *DaoFactorySession) ListRequests() ([]DaoRequest, error) {
	return _DaoFactory.Contract.ListRequests(&_DaoFactory.CallOpts)
}

// ListRequests is a free data retrieval call binding the contract method 0x53d5b816.
//
// Solidity: function listRequests() view returns(((string,string,address[],string,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]),bool,address,uint256,uint256,address)[])
func (_DaoFactory *DaoFactoryCallerSession) ListRequests() ([]DaoRequest, error) {
	return _DaoFactory.Contract.ListRequests(&_DaoFactory.CallOpts)
}

// ListSenderDaos is a free data retrieval call binding the contract method 0x6a7c0a1f.
//
// Solidity: function listSenderDaos() view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCaller) ListSenderDaos(opts *bind.CallOpts) ([]Dao, error) {
	var (
		ret0 = new([]Dao)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "listSenderDaos")
	return *ret0, err
}

// ListSenderDaos is a free data retrieval call binding the contract method 0x6a7c0a1f.
//
// Solidity: function listSenderDaos() view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactorySession) ListSenderDaos() ([]Dao, error) {
	return _DaoFactory.Contract.ListSenderDaos(&_DaoFactory.CallOpts)
}

// ListSenderDaos is a free data retrieval call binding the contract method 0x6a7c0a1f.
//
// Solidity: function listSenderDaos() view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCallerSession) ListSenderDaos() ([]Dao, error) {
	return _DaoFactory.Contract.ListSenderDaos(&_DaoFactory.CallOpts)
}

// ListUserDaos is a free data retrieval call binding the contract method 0x15aa288b.
//
// Solidity: function listUserDaos(address addr) view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCaller) ListUserDaos(opts *bind.CallOpts, addr common.Address) ([]Dao, error) {
	var (
		ret0 = new([]Dao)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "listUserDaos", addr)
	return *ret0, err
}

// ListUserDaos is a free data retrieval call binding the contract method 0x15aa288b.
//
// Solidity: function listUserDaos(address addr) view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactorySession) ListUserDaos(addr common.Address) ([]Dao, error) {
	return _DaoFactory.Contract.ListUserDaos(&_DaoFactory.CallOpts, addr)
}

// ListUserDaos is a free data retrieval call binding the contract method 0x15aa288b.
//
// Solidity: function listUserDaos(address addr) view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCallerSession) ListUserDaos(addr common.Address) ([]Dao, error) {
	return _DaoFactory.Contract.ListUserDaos(&_DaoFactory.CallOpts, addr)
}

// QuoteOf is a free data retrieval call binding the contract method 0xcab1ef67.
//
// Solidity: function quoteOf(address addr) view returns(uint256)
func (_DaoFactory *DaoFactoryCaller) QuoteOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "quoteOf", addr)
	return *ret0, err
}

// QuoteOf is a free data retrieval call binding the contract method 0xcab1ef67.
//
// Solidity: function quoteOf(address addr) view returns(uint256)
func (_DaoFactory *DaoFactorySession) QuoteOf(addr common.Address) (*big.Int, error) {
	return _DaoFactory.Contract.QuoteOf(&_DaoFactory.CallOpts, addr)
}

// QuoteOf is a free data retrieval call binding the contract method 0xcab1ef67.
//
// Solidity: function quoteOf(address addr) view returns(uint256)
func (_DaoFactory *DaoFactoryCallerSession) QuoteOf(addr common.Address) (*big.Int, error) {
	return _DaoFactory.Contract.QuoteOf(&_DaoFactory.CallOpts, addr)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_DaoFactory *DaoFactoryCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_DaoFactory *DaoFactorySession) Ready() (bool, error) {
	return _DaoFactory.Contract.Ready(&_DaoFactory.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_DaoFactory *DaoFactoryCallerSession) Ready() (bool, error) {
	return _DaoFactory.Contract.Ready(&_DaoFactory.CallOpts)
}

// TmpListDaos is a free data retrieval call binding the contract method 0x7a3a890f.
//
// Solidity: function tmpListDaos(string tmpName) view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCaller) TmpListDaos(opts *bind.CallOpts, tmpName string) ([]Dao, error) {
	var (
		ret0 = new([]Dao)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "tmpListDaos", tmpName)
	return *ret0, err
}

// TmpListDaos is a free data retrieval call binding the contract method 0x7a3a890f.
//
// Solidity: function tmpListDaos(string tmpName) view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactorySession) TmpListDaos(tmpName string) ([]Dao, error) {
	return _DaoFactory.Contract.TmpListDaos(&_DaoFactory.CallOpts, tmpName)
}

// TmpListDaos is a free data retrieval call binding the contract method 0x7a3a890f.
//
// Solidity: function tmpListDaos(string tmpName) view returns((string,string,string,address[],address)[])
func (_DaoFactory *DaoFactoryCallerSession) TmpListDaos(tmpName string) ([]Dao, error) {
	return _DaoFactory.Contract.TmpListDaos(&_DaoFactory.CallOpts, tmpName)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DaoFactory *DaoFactoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DaoFactory.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DaoFactory *DaoFactorySession) Version() (string, error) {
	return _DaoFactory.Contract.Version(&_DaoFactory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DaoFactory *DaoFactoryCallerSession) Version() (string, error) {
	return _DaoFactory.Contract.Version(&_DaoFactory.CallOpts)
}

// AddContract is a paid mutator transaction binding the contract method 0x152993a8.
//
// Solidity: function addContract(string depName, bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactor) AddContract(opts *bind.TransactOpts, depName string, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "addContract", depName, bytecode)
}

// AddContract is a paid mutator transaction binding the contract method 0x152993a8.
//
// Solidity: function addContract(string depName, bytes bytecode) returns()
func (_DaoFactory *DaoFactorySession) AddContract(depName string, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.AddContract(&_DaoFactory.TransactOpts, depName, bytecode)
}

// AddContract is a paid mutator transaction binding the contract method 0x152993a8.
//
// Solidity: function addContract(string depName, bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactorSession) AddContract(depName string, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.AddContract(&_DaoFactory.TransactOpts, depName, bytecode)
}

// AddDao is a paid mutator transaction binding the contract method 0xa78272a2.
//
// Solidity: function addDao((string,string,string,address[],address) dao) returns()
func (_DaoFactory *DaoFactoryTransactor) AddDao(opts *bind.TransactOpts, dao Dao) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "addDao", dao)
}

// AddDao is a paid mutator transaction binding the contract method 0xa78272a2.
//
// Solidity: function addDao((string,string,string,address[],address) dao) returns()
func (_DaoFactory *DaoFactorySession) AddDao(dao Dao) (*types.Transaction, error) {
	return _DaoFactory.Contract.AddDao(&_DaoFactory.TransactOpts, dao)
}

// AddDao is a paid mutator transaction binding the contract method 0xa78272a2.
//
// Solidity: function addDao((string,string,string,address[],address) dao) returns()
func (_DaoFactory *DaoFactoryTransactorSession) AddDao(dao Dao) (*types.Transaction, error) {
	return _DaoFactory.Contract.AddDao(&_DaoFactory.TransactOpts, dao)
}

// AddExtension is a paid mutator transaction binding the contract method 0x16ec7c20.
//
// Solidity: function addExtension((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8) ext) returns()
func (_DaoFactory *DaoFactoryTransactor) AddExtension(opts *bind.TransactOpts, ext Extension) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "addExtension", ext)
}

// AddExtension is a paid mutator transaction binding the contract method 0x16ec7c20.
//
// Solidity: function addExtension((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8) ext) returns()
func (_DaoFactory *DaoFactorySession) AddExtension(ext Extension) (*types.Transaction, error) {
	return _DaoFactory.Contract.AddExtension(&_DaoFactory.TransactOpts, ext)
}

// AddExtension is a paid mutator transaction binding the contract method 0x16ec7c20.
//
// Solidity: function addExtension((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8) ext) returns()
func (_DaoFactory *DaoFactoryTransactorSession) AddExtension(ext Extension) (*types.Transaction, error) {
	return _DaoFactory.Contract.AddExtension(&_DaoFactory.TransactOpts, ext)
}

// AddExtensionState is a paid mutator transaction binding the contract method 0x07fca25e.
//
// Solidity: function addExtensionState(string extName, bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactor) AddExtensionState(opts *bind.TransactOpts, extName string, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "addExtensionState", extName, bytecode)
}

// AddExtensionState is a paid mutator transaction binding the contract method 0x07fca25e.
//
// Solidity: function addExtensionState(string extName, bytes bytecode) returns()
func (_DaoFactory *DaoFactorySession) AddExtensionState(extName string, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.AddExtensionState(&_DaoFactory.TransactOpts, extName, bytecode)
}

// AddExtensionState is a paid mutator transaction binding the contract method 0x07fca25e.
//
// Solidity: function addExtensionState(string extName, bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactorSession) AddExtensionState(extName string, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.AddExtensionState(&_DaoFactory.TransactOpts, extName, bytecode)
}

// AllowQuote is a paid mutator transaction binding the contract method 0x56285575.
//
// Solidity: function allowQuote(address addr, uint256 quote) returns()
func (_DaoFactory *DaoFactoryTransactor) AllowQuote(opts *bind.TransactOpts, addr common.Address, quote *big.Int) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "allowQuote", addr, quote)
}

// AllowQuote is a paid mutator transaction binding the contract method 0x56285575.
//
// Solidity: function allowQuote(address addr, uint256 quote) returns()
func (_DaoFactory *DaoFactorySession) AllowQuote(addr common.Address, quote *big.Int) (*types.Transaction, error) {
	return _DaoFactory.Contract.AllowQuote(&_DaoFactory.TransactOpts, addr, quote)
}

// AllowQuote is a paid mutator transaction binding the contract method 0x56285575.
//
// Solidity: function allowQuote(address addr, uint256 quote) returns()
func (_DaoFactory *DaoFactoryTransactorSession) AllowQuote(addr common.Address, quote *big.Int) (*types.Transaction, error) {
	return _DaoFactory.Contract.AllowQuote(&_DaoFactory.TransactOpts, addr, quote)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x1e0715df.
//
// Solidity: function approveRequest(string name) returns()
func (_DaoFactory *DaoFactoryTransactor) ApproveRequest(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "approveRequest", name)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x1e0715df.
//
// Solidity: function approveRequest(string name) returns()
func (_DaoFactory *DaoFactorySession) ApproveRequest(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.ApproveRequest(&_DaoFactory.TransactOpts, name)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x1e0715df.
//
// Solidity: function approveRequest(string name) returns()
func (_DaoFactory *DaoFactoryTransactorSession) ApproveRequest(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.ApproveRequest(&_DaoFactory.TransactOpts, name)
}

// DelBizTemplate is a paid mutator transaction binding the contract method 0x7d0d5e6c.
//
// Solidity: function delBizTemplate(string tmpl) returns()
func (_DaoFactory *DaoFactoryTransactor) DelBizTemplate(opts *bind.TransactOpts, tmpl string) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "delBizTemplate", tmpl)
}

// DelBizTemplate is a paid mutator transaction binding the contract method 0x7d0d5e6c.
//
// Solidity: function delBizTemplate(string tmpl) returns()
func (_DaoFactory *DaoFactorySession) DelBizTemplate(tmpl string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DelBizTemplate(&_DaoFactory.TransactOpts, tmpl)
}

// DelBizTemplate is a paid mutator transaction binding the contract method 0x7d0d5e6c.
//
// Solidity: function delBizTemplate(string tmpl) returns()
func (_DaoFactory *DaoFactoryTransactorSession) DelBizTemplate(tmpl string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DelBizTemplate(&_DaoFactory.TransactOpts, tmpl)
}

// DelDao is a paid mutator transaction binding the contract method 0x2bb752cf.
//
// Solidity: function delDao(string name) returns()
func (_DaoFactory *DaoFactoryTransactor) DelDao(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "delDao", name)
}

// DelDao is a paid mutator transaction binding the contract method 0x2bb752cf.
//
// Solidity: function delDao(string name) returns()
func (_DaoFactory *DaoFactorySession) DelDao(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DelDao(&_DaoFactory.TransactOpts, name)
}

// DelDao is a paid mutator transaction binding the contract method 0x2bb752cf.
//
// Solidity: function delDao(string name) returns()
func (_DaoFactory *DaoFactoryTransactorSession) DelDao(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DelDao(&_DaoFactory.TransactOpts, name)
}

// DeployExtension is a paid mutator transaction binding the contract method 0x71f7d903.
//
// Solidity: function deployExtension(address registry, string extensionName, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, address extStateAddr) returns((string,uint256,address,uint256,uint256,uint8))
func (_DaoFactory *DaoFactoryTransactor) DeployExtension(opts *bind.TransactOpts, registry common.Address, extensionName string, params DeploymentParameterV1, extStateAddr common.Address) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "deployExtension", registry, extensionName, params, extStateAddr)
}

// DeployExtension is a paid mutator transaction binding the contract method 0x71f7d903.
//
// Solidity: function deployExtension(address registry, string extensionName, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, address extStateAddr) returns((string,uint256,address,uint256,uint256,uint8))
func (_DaoFactory *DaoFactorySession) DeployExtension(registry common.Address, extensionName string, params DeploymentParameterV1, extStateAddr common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployExtension(&_DaoFactory.TransactOpts, registry, extensionName, params, extStateAddr)
}

// DeployExtension is a paid mutator transaction binding the contract method 0x71f7d903.
//
// Solidity: function deployExtension(address registry, string extensionName, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, address extStateAddr) returns((string,uint256,address,uint256,uint256,uint8))
func (_DaoFactory *DaoFactoryTransactorSession) DeployExtension(registry common.Address, extensionName string, params DeploymentParameterV1, extStateAddr common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployExtension(&_DaoFactory.TransactOpts, registry, extensionName, params, extStateAddr)
}

// DeployExtensionFromExtension is a paid mutator transaction binding the contract method 0xe5477e0d.
//
// Solidity: function deployExtensionFromExtension(string name, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, address stateAddress) returns((string,uint256,address,uint256,uint256,uint8))
func (_DaoFactory *DaoFactoryTransactor) DeployExtensionFromExtension(opts *bind.TransactOpts, name string, params DeploymentParameterV1, stateAddress common.Address) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "deployExtensionFromExtension", name, params, stateAddress)
}

// DeployExtensionFromExtension is a paid mutator transaction binding the contract method 0xe5477e0d.
//
// Solidity: function deployExtensionFromExtension(string name, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, address stateAddress) returns((string,uint256,address,uint256,uint256,uint8))
func (_DaoFactory *DaoFactorySession) DeployExtensionFromExtension(name string, params DeploymentParameterV1, stateAddress common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployExtensionFromExtension(&_DaoFactory.TransactOpts, name, params, stateAddress)
}

// DeployExtensionFromExtension is a paid mutator transaction binding the contract method 0xe5477e0d.
//
// Solidity: function deployExtensionFromExtension(string name, (string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]) params, address stateAddress) returns((string,uint256,address,uint256,uint256,uint8))
func (_DaoFactory *DaoFactoryTransactorSession) DeployExtensionFromExtension(name string, params DeploymentParameterV1, stateAddress common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployExtensionFromExtension(&_DaoFactory.TransactOpts, name, params, stateAddress)
}

// DeployExtensionState is a paid mutator transaction binding the contract method 0xbca39057.
//
// Solidity: function deployExtensionState(string name) returns(address)
func (_DaoFactory *DaoFactoryTransactor) DeployExtensionState(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "deployExtensionState", name)
}

// DeployExtensionState is a paid mutator transaction binding the contract method 0xbca39057.
//
// Solidity: function deployExtensionState(string name) returns(address)
func (_DaoFactory *DaoFactorySession) DeployExtensionState(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployExtensionState(&_DaoFactory.TransactOpts, name)
}

// DeployExtensionState is a paid mutator transaction binding the contract method 0xbca39057.
//
// Solidity: function deployExtensionState(string name) returns(address)
func (_DaoFactory *DaoFactoryTransactorSession) DeployExtensionState(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployExtensionState(&_DaoFactory.TransactOpts, name)
}

// DeployExtensionStateFromExtension is a paid mutator transaction binding the contract method 0xc32a45c0.
//
// Solidity: function deployExtensionStateFromExtension(string name) returns(address)
func (_DaoFactory *DaoFactoryTransactor) DeployExtensionStateFromExtension(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "deployExtensionStateFromExtension", name)
}

// DeployExtensionStateFromExtension is a paid mutator transaction binding the contract method 0xc32a45c0.
//
// Solidity: function deployExtensionStateFromExtension(string name) returns(address)
func (_DaoFactory *DaoFactorySession) DeployExtensionStateFromExtension(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployExtensionStateFromExtension(&_DaoFactory.TransactOpts, name)
}

// DeployExtensionStateFromExtension is a paid mutator transaction binding the contract method 0xc32a45c0.
//
// Solidity: function deployExtensionStateFromExtension(string name) returns(address)
func (_DaoFactory *DaoFactoryTransactorSession) DeployExtensionStateFromExtension(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployExtensionStateFromExtension(&_DaoFactory.TransactOpts, name)
}

// DeployStandardContractFromExtension is a paid mutator transaction binding the contract method 0xdbd7b930.
//
// Solidity: function deployStandardContractFromExtension(string depName) returns(address)
func (_DaoFactory *DaoFactoryTransactor) DeployStandardContractFromExtension(opts *bind.TransactOpts, depName string) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "deployStandardContractFromExtension", depName)
}

// DeployStandardContractFromExtension is a paid mutator transaction binding the contract method 0xdbd7b930.
//
// Solidity: function deployStandardContractFromExtension(string depName) returns(address)
func (_DaoFactory *DaoFactorySession) DeployStandardContractFromExtension(depName string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployStandardContractFromExtension(&_DaoFactory.TransactOpts, depName)
}

// DeployStandardContractFromExtension is a paid mutator transaction binding the contract method 0xdbd7b930.
//
// Solidity: function deployStandardContractFromExtension(string depName) returns(address)
func (_DaoFactory *DaoFactoryTransactorSession) DeployStandardContractFromExtension(depName string) (*types.Transaction, error) {
	return _DaoFactory.Contract.DeployStandardContractFromExtension(&_DaoFactory.TransactOpts, depName)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_DaoFactory *DaoFactoryTransactor) Initialize(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "initialize", addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_DaoFactory *DaoFactorySession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.Initialize(&_DaoFactory.TransactOpts, addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_DaoFactory *DaoFactoryTransactorSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.Initialize(&_DaoFactory.TransactOpts, addrs)
}

// ProcessRequest is a paid mutator transaction binding the contract method 0x80a7edea.
//
// Solidity: function processRequest(string name) returns()
func (_DaoFactory *DaoFactoryTransactor) ProcessRequest(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "processRequest", name)
}

// ProcessRequest is a paid mutator transaction binding the contract method 0x80a7edea.
//
// Solidity: function processRequest(string name) returns()
func (_DaoFactory *DaoFactorySession) ProcessRequest(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.ProcessRequest(&_DaoFactory.TransactOpts, name)
}

// ProcessRequest is a paid mutator transaction binding the contract method 0x80a7edea.
//
// Solidity: function processRequest(string name) returns()
func (_DaoFactory *DaoFactoryTransactorSession) ProcessRequest(name string) (*types.Transaction, error) {
	return _DaoFactory.Contract.ProcessRequest(&_DaoFactory.TransactOpts, name)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x72a8d706.
//
// Solidity: function rejectRequest(string name, string reason) returns()
func (_DaoFactory *DaoFactoryTransactor) RejectRequest(opts *bind.TransactOpts, name string, reason string) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "rejectRequest", name, reason)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x72a8d706.
//
// Solidity: function rejectRequest(string name, string reason) returns()
func (_DaoFactory *DaoFactorySession) RejectRequest(name string, reason string) (*types.Transaction, error) {
	return _DaoFactory.Contract.RejectRequest(&_DaoFactory.TransactOpts, name, reason)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x72a8d706.
//
// Solidity: function rejectRequest(string name, string reason) returns()
func (_DaoFactory *DaoFactoryTransactorSession) RejectRequest(name string, reason string) (*types.Transaction, error) {
	return _DaoFactory.Contract.RejectRequest(&_DaoFactory.TransactOpts, name, reason)
}

// RequestDaoCreation is a paid mutator transaction binding the contract method 0xd58a731c.
//
// Solidity: function requestDaoCreation((string,string,address[],string,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]) req) payable returns()
func (_DaoFactory *DaoFactoryTransactor) RequestDaoCreation(opts *bind.TransactOpts, req DaoUserRequest) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "requestDaoCreation", req)
}

// RequestDaoCreation is a paid mutator transaction binding the contract method 0xd58a731c.
//
// Solidity: function requestDaoCreation((string,string,address[],string,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]) req) payable returns()
func (_DaoFactory *DaoFactorySession) RequestDaoCreation(req DaoUserRequest) (*types.Transaction, error) {
	return _DaoFactory.Contract.RequestDaoCreation(&_DaoFactory.TransactOpts, req)
}

// RequestDaoCreation is a paid mutator transaction binding the contract method 0xd58a731c.
//
// Solidity: function requestDaoCreation((string,string,address[],string,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]) req) payable returns()
func (_DaoFactory *DaoFactoryTransactorSession) RequestDaoCreation(req DaoUserRequest) (*types.Transaction, error) {
	return _DaoFactory.Contract.RequestDaoCreation(&_DaoFactory.TransactOpts, req)
}

// SetCreationFee is a paid mutator transaction binding the contract method 0xb7d86225.
//
// Solidity: function setCreationFee(uint256 fee) returns()
func (_DaoFactory *DaoFactoryTransactor) SetCreationFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "setCreationFee", fee)
}

// SetCreationFee is a paid mutator transaction binding the contract method 0xb7d86225.
//
// Solidity: function setCreationFee(uint256 fee) returns()
func (_DaoFactory *DaoFactorySession) SetCreationFee(fee *big.Int) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetCreationFee(&_DaoFactory.TransactOpts, fee)
}

// SetCreationFee is a paid mutator transaction binding the contract method 0xb7d86225.
//
// Solidity: function setCreationFee(uint256 fee) returns()
func (_DaoFactory *DaoFactoryTransactorSession) SetCreationFee(fee *big.Int) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetCreationFee(&_DaoFactory.TransactOpts, fee)
}

// SetDaoRegistryBytecode is a paid mutator transaction binding the contract method 0xd3b1ad78.
//
// Solidity: function setDaoRegistryBytecode(bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactor) SetDaoRegistryBytecode(opts *bind.TransactOpts, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "setDaoRegistryBytecode", bytecode)
}

// SetDaoRegistryBytecode is a paid mutator transaction binding the contract method 0xd3b1ad78.
//
// Solidity: function setDaoRegistryBytecode(bytes bytecode) returns()
func (_DaoFactory *DaoFactorySession) SetDaoRegistryBytecode(bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetDaoRegistryBytecode(&_DaoFactory.TransactOpts, bytecode)
}

// SetDaoRegistryBytecode is a paid mutator transaction binding the contract method 0xd3b1ad78.
//
// Solidity: function setDaoRegistryBytecode(bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactorSession) SetDaoRegistryBytecode(bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetDaoRegistryBytecode(&_DaoFactory.TransactOpts, bytecode)
}

// SetDaoRegistryStateBytecode is a paid mutator transaction binding the contract method 0x18cf0c0f.
//
// Solidity: function setDaoRegistryStateBytecode(bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactor) SetDaoRegistryStateBytecode(opts *bind.TransactOpts, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "setDaoRegistryStateBytecode", bytecode)
}

// SetDaoRegistryStateBytecode is a paid mutator transaction binding the contract method 0x18cf0c0f.
//
// Solidity: function setDaoRegistryStateBytecode(bytes bytecode) returns()
func (_DaoFactory *DaoFactorySession) SetDaoRegistryStateBytecode(bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetDaoRegistryStateBytecode(&_DaoFactory.TransactOpts, bytecode)
}

// SetDaoRegistryStateBytecode is a paid mutator transaction binding the contract method 0x18cf0c0f.
//
// Solidity: function setDaoRegistryStateBytecode(bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactorSession) SetDaoRegistryStateBytecode(bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetDaoRegistryStateBytecode(&_DaoFactory.TransactOpts, bytecode)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_DaoFactory *DaoFactoryTransactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "setForwarder", forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_DaoFactory *DaoFactorySession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetForwarder(&_DaoFactory.TransactOpts, forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_DaoFactory *DaoFactoryTransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetForwarder(&_DaoFactory.TransactOpts, forwarder)
}

// SetLayers is a paid mutator transaction binding the contract method 0x8acc75e5.
//
// Solidity: function setLayers(address[] addrs) returns()
func (_DaoFactory *DaoFactoryTransactor) SetLayers(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "setLayers", addrs)
}

// SetLayers is a paid mutator transaction binding the contract method 0x8acc75e5.
//
// Solidity: function setLayers(address[] addrs) returns()
func (_DaoFactory *DaoFactorySession) SetLayers(addrs []common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetLayers(&_DaoFactory.TransactOpts, addrs)
}

// SetLayers is a paid mutator transaction binding the contract method 0x8acc75e5.
//
// Solidity: function setLayers(address[] addrs) returns()
func (_DaoFactory *DaoFactoryTransactorSession) SetLayers(addrs []common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.SetLayers(&_DaoFactory.TransactOpts, addrs)
}

// StoredAppData is a paid mutator transaction binding the contract method 0x57e65538.
//
// Solidity: function storedAppData(uint256 activityFlag, bytes data) returns()
func (_DaoFactory *DaoFactoryTransactor) StoredAppData(opts *bind.TransactOpts, activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "storedAppData", activityFlag, data)
}

// StoredAppData is a paid mutator transaction binding the contract method 0x57e65538.
//
// Solidity: function storedAppData(uint256 activityFlag, bytes data) returns()
func (_DaoFactory *DaoFactorySession) StoredAppData(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.StoredAppData(&_DaoFactory.TransactOpts, activityFlag, data)
}

// StoredAppData is a paid mutator transaction binding the contract method 0x57e65538.
//
// Solidity: function storedAppData(uint256 activityFlag, bytes data) returns()
func (_DaoFactory *DaoFactoryTransactorSession) StoredAppData(activityFlag *big.Int, data []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.StoredAppData(&_DaoFactory.TransactOpts, activityFlag, data)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_DaoFactory *DaoFactoryTransactor) Upgrade(opts *bind.TransactOpts, newVersion common.Address) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "upgrade", newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_DaoFactory *DaoFactorySession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.Upgrade(&_DaoFactory.TransactOpts, newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_DaoFactory *DaoFactoryTransactorSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.Upgrade(&_DaoFactory.TransactOpts, newVersion)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xcf26984b.
//
// Solidity: function upgradeContract(string depName, uint256 version, bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactor) UpgradeContract(opts *bind.TransactOpts, depName string, version *big.Int, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "upgradeContract", depName, version, bytecode)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xcf26984b.
//
// Solidity: function upgradeContract(string depName, uint256 version, bytes bytecode) returns()
func (_DaoFactory *DaoFactorySession) UpgradeContract(depName string, version *big.Int, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeContract(&_DaoFactory.TransactOpts, depName, version, bytecode)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xcf26984b.
//
// Solidity: function upgradeContract(string depName, uint256 version, bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactorSession) UpgradeContract(depName string, version *big.Int, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeContract(&_DaoFactory.TransactOpts, depName, version, bytecode)
}

// UpgradeDao is a paid mutator transaction binding the contract method 0xae6d0563.
//
// Solidity: function upgradeDao((string,string,string,address[],address) dao) returns()
func (_DaoFactory *DaoFactoryTransactor) UpgradeDao(opts *bind.TransactOpts, dao Dao) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "upgradeDao", dao)
}

// UpgradeDao is a paid mutator transaction binding the contract method 0xae6d0563.
//
// Solidity: function upgradeDao((string,string,string,address[],address) dao) returns()
func (_DaoFactory *DaoFactorySession) UpgradeDao(dao Dao) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeDao(&_DaoFactory.TransactOpts, dao)
}

// UpgradeDao is a paid mutator transaction binding the contract method 0xae6d0563.
//
// Solidity: function upgradeDao((string,string,string,address[],address) dao) returns()
func (_DaoFactory *DaoFactoryTransactorSession) UpgradeDao(dao Dao) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeDao(&_DaoFactory.TransactOpts, dao)
}

// UpgradeExtension is a paid mutator transaction binding the contract method 0x38e1cd28.
//
// Solidity: function upgradeExtension((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8) ext) returns()
func (_DaoFactory *DaoFactoryTransactor) UpgradeExtension(opts *bind.TransactOpts, ext Extension) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "upgradeExtension", ext)
}

// UpgradeExtension is a paid mutator transaction binding the contract method 0x38e1cd28.
//
// Solidity: function upgradeExtension((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8) ext) returns()
func (_DaoFactory *DaoFactorySession) UpgradeExtension(ext Extension) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeExtension(&_DaoFactory.TransactOpts, ext)
}

// UpgradeExtension is a paid mutator transaction binding the contract method 0x38e1cd28.
//
// Solidity: function upgradeExtension((string,uint256,bytes,(string,string)[],(string,string)[],(string,string)[],uint256,uint256,(string,string[])[],string[],uint8) ext) returns()
func (_DaoFactory *DaoFactoryTransactorSession) UpgradeExtension(ext Extension) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeExtension(&_DaoFactory.TransactOpts, ext)
}

// UpgradeExtensionState is a paid mutator transaction binding the contract method 0x56fe87a3.
//
// Solidity: function upgradeExtensionState(string extName, uint256 version, bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactor) UpgradeExtensionState(opts *bind.TransactOpts, extName string, version *big.Int, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "upgradeExtensionState", extName, version, bytecode)
}

// UpgradeExtensionState is a paid mutator transaction binding the contract method 0x56fe87a3.
//
// Solidity: function upgradeExtensionState(string extName, uint256 version, bytes bytecode) returns()
func (_DaoFactory *DaoFactorySession) UpgradeExtensionState(extName string, version *big.Int, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeExtensionState(&_DaoFactory.TransactOpts, extName, version, bytecode)
}

// UpgradeExtensionState is a paid mutator transaction binding the contract method 0x56fe87a3.
//
// Solidity: function upgradeExtensionState(string extName, uint256 version, bytes bytecode) returns()
func (_DaoFactory *DaoFactoryTransactorSession) UpgradeExtensionState(extName string, version *big.Int, bytecode []byte) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeExtensionState(&_DaoFactory.TransactOpts, extName, version, bytecode)
}

// UpgradeHook is a paid mutator transaction binding the contract method 0x1d261060.
//
// Solidity: function upgradeHook(address state) returns()
func (_DaoFactory *DaoFactoryTransactor) UpgradeHook(opts *bind.TransactOpts, state common.Address) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "upgradeHook", state)
}

// UpgradeHook is a paid mutator transaction binding the contract method 0x1d261060.
//
// Solidity: function upgradeHook(address state) returns()
func (_DaoFactory *DaoFactorySession) UpgradeHook(state common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeHook(&_DaoFactory.TransactOpts, state)
}

// UpgradeHook is a paid mutator transaction binding the contract method 0x1d261060.
//
// Solidity: function upgradeHook(address state) returns()
func (_DaoFactory *DaoFactoryTransactorSession) UpgradeHook(state common.Address) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpgradeHook(&_DaoFactory.TransactOpts, state)
}

// UpsertBizTemplate is a paid mutator transaction binding the contract method 0x06ec5e67.
//
// Solidity: function upsertBizTemplate((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]) tmpl) returns()
func (_DaoFactory *DaoFactoryTransactor) UpsertBizTemplate(opts *bind.TransactOpts, tmpl BizTemplate) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "upsertBizTemplate", tmpl)
}

// UpsertBizTemplate is a paid mutator transaction binding the contract method 0x06ec5e67.
//
// Solidity: function upsertBizTemplate((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]) tmpl) returns()
func (_DaoFactory *DaoFactorySession) UpsertBizTemplate(tmpl BizTemplate) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpsertBizTemplate(&_DaoFactory.TransactOpts, tmpl)
}

// UpsertBizTemplate is a paid mutator transaction binding the contract method 0x06ec5e67.
//
// Solidity: function upsertBizTemplate((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]) tmpl) returns()
func (_DaoFactory *DaoFactoryTransactorSession) UpsertBizTemplate(tmpl BizTemplate) (*types.Transaction, error) {
	return _DaoFactory.Contract.UpsertBizTemplate(&_DaoFactory.TransactOpts, tmpl)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x5a6b26ba.
//
// Solidity: function withdrawal(address addr, uint256 value) returns()
func (_DaoFactory *DaoFactoryTransactor) Withdrawal(opts *bind.TransactOpts, addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _DaoFactory.contract.Transact(opts, "withdrawal", addr, value)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x5a6b26ba.
//
// Solidity: function withdrawal(address addr, uint256 value) returns()
func (_DaoFactory *DaoFactorySession) Withdrawal(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _DaoFactory.Contract.Withdrawal(&_DaoFactory.TransactOpts, addr, value)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x5a6b26ba.
//
// Solidity: function withdrawal(address addr, uint256 value) returns()
func (_DaoFactory *DaoFactoryTransactorSession) Withdrawal(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _DaoFactory.Contract.Withdrawal(&_DaoFactory.TransactOpts, addr, value)
}

// DaoFactoryAddBussinessTemplateEventIterator is returned from FilterAddBussinessTemplateEvent and is used to iterate over the raw logs and unpacked data for AddBussinessTemplateEvent events raised by the DaoFactory contract.
type DaoFactoryAddBussinessTemplateEventIterator struct {
	Event *DaoFactoryAddBussinessTemplateEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryAddBussinessTemplateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryAddBussinessTemplateEvent)
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
		it.Event = new(DaoFactoryAddBussinessTemplateEvent)
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
func (it *DaoFactoryAddBussinessTemplateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryAddBussinessTemplateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryAddBussinessTemplateEvent represents a AddBussinessTemplateEvent event raised by the DaoFactory contract.
type DaoFactoryAddBussinessTemplateEvent struct {
	Ext BizTemplate
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddBussinessTemplateEvent is a free log retrieval operation binding the contract event 0xa053a289024ee0aebe9fc9e1bd26a8ddc5707295ec57cd0ea8b7e5b3bfe63091.
//
// Solidity: event addBussinessTemplateEvent((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]) ext)
func (_DaoFactory *DaoFactoryFilterer) FilterAddBussinessTemplateEvent(opts *bind.FilterOpts) (*DaoFactoryAddBussinessTemplateEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "addBussinessTemplateEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryAddBussinessTemplateEventIterator{contract: _DaoFactory.contract, event: "addBussinessTemplateEvent", logs: logs, sub: sub}, nil
}

// WatchAddBussinessTemplateEvent is a free log subscription operation binding the contract event 0xa053a289024ee0aebe9fc9e1bd26a8ddc5707295ec57cd0ea8b7e5b3bfe63091.
//
// Solidity: event addBussinessTemplateEvent((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]) ext)
func (_DaoFactory *DaoFactoryFilterer) WatchAddBussinessTemplateEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryAddBussinessTemplateEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "addBussinessTemplateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryAddBussinessTemplateEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "addBussinessTemplateEvent", log); err != nil {
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

// ParseAddBussinessTemplateEvent is a log parse operation binding the contract event 0xa053a289024ee0aebe9fc9e1bd26a8ddc5707295ec57cd0ea8b7e5b3bfe63091.
//
// Solidity: event addBussinessTemplateEvent((string,uint256,(string,(string[],uint256[],bool[],address[],string[][],uint256[][],address[][],bool[][]))[]) ext)
func (_DaoFactory *DaoFactoryFilterer) ParseAddBussinessTemplateEvent(log types.Log) (*DaoFactoryAddBussinessTemplateEvent, error) {
	event := new(DaoFactoryAddBussinessTemplateEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "addBussinessTemplateEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryAddContractEventIterator is returned from FilterAddContractEvent and is used to iterate over the raw logs and unpacked data for AddContractEvent events raised by the DaoFactory contract.
type DaoFactoryAddContractEventIterator struct {
	Event *DaoFactoryAddContractEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryAddContractEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryAddContractEvent)
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
		it.Event = new(DaoFactoryAddContractEvent)
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
func (it *DaoFactoryAddContractEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryAddContractEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryAddContractEvent represents a AddContractEvent event raised by the DaoFactory contract.
type DaoFactoryAddContractEvent struct {
	Name     string
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddContractEvent is a free log retrieval operation binding the contract event 0xda1bb1ae0ac2e35d2763bfc263cc0510b896f281b6d938212f3081fe55ed1a7c.
//
// Solidity: event addContractEvent(string name, address operator)
func (_DaoFactory *DaoFactoryFilterer) FilterAddContractEvent(opts *bind.FilterOpts) (*DaoFactoryAddContractEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "addContractEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryAddContractEventIterator{contract: _DaoFactory.contract, event: "addContractEvent", logs: logs, sub: sub}, nil
}

// WatchAddContractEvent is a free log subscription operation binding the contract event 0xda1bb1ae0ac2e35d2763bfc263cc0510b896f281b6d938212f3081fe55ed1a7c.
//
// Solidity: event addContractEvent(string name, address operator)
func (_DaoFactory *DaoFactoryFilterer) WatchAddContractEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryAddContractEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "addContractEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryAddContractEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "addContractEvent", log); err != nil {
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

// ParseAddContractEvent is a log parse operation binding the contract event 0xda1bb1ae0ac2e35d2763bfc263cc0510b896f281b6d938212f3081fe55ed1a7c.
//
// Solidity: event addContractEvent(string name, address operator)
func (_DaoFactory *DaoFactoryFilterer) ParseAddContractEvent(log types.Log) (*DaoFactoryAddContractEvent, error) {
	event := new(DaoFactoryAddContractEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "addContractEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryAddExtensionEventIterator is returned from FilterAddExtensionEvent and is used to iterate over the raw logs and unpacked data for AddExtensionEvent events raised by the DaoFactory contract.
type DaoFactoryAddExtensionEventIterator struct {
	Event *DaoFactoryAddExtensionEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryAddExtensionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryAddExtensionEvent)
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
		it.Event = new(DaoFactoryAddExtensionEvent)
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
func (it *DaoFactoryAddExtensionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryAddExtensionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryAddExtensionEvent represents a AddExtensionEvent event raised by the DaoFactory contract.
type DaoFactoryAddExtensionEvent struct {
	Name     string
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddExtensionEvent is a free log retrieval operation binding the contract event 0x257deca14dc58c55eefa667f07843a2b5da52c6bb339111f4fad8b24f05071c7.
//
// Solidity: event addExtensionEvent(string name, address operator)
func (_DaoFactory *DaoFactoryFilterer) FilterAddExtensionEvent(opts *bind.FilterOpts) (*DaoFactoryAddExtensionEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "addExtensionEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryAddExtensionEventIterator{contract: _DaoFactory.contract, event: "addExtensionEvent", logs: logs, sub: sub}, nil
}

// WatchAddExtensionEvent is a free log subscription operation binding the contract event 0x257deca14dc58c55eefa667f07843a2b5da52c6bb339111f4fad8b24f05071c7.
//
// Solidity: event addExtensionEvent(string name, address operator)
func (_DaoFactory *DaoFactoryFilterer) WatchAddExtensionEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryAddExtensionEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "addExtensionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryAddExtensionEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "addExtensionEvent", log); err != nil {
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

// ParseAddExtensionEvent is a log parse operation binding the contract event 0x257deca14dc58c55eefa667f07843a2b5da52c6bb339111f4fad8b24f05071c7.
//
// Solidity: event addExtensionEvent(string name, address operator)
func (_DaoFactory *DaoFactoryFilterer) ParseAddExtensionEvent(log types.Log) (*DaoFactoryAddExtensionEvent, error) {
	event := new(DaoFactoryAddExtensionEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "addExtensionEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryAddExtensionStateEventIterator is returned from FilterAddExtensionStateEvent and is used to iterate over the raw logs and unpacked data for AddExtensionStateEvent events raised by the DaoFactory contract.
type DaoFactoryAddExtensionStateEventIterator struct {
	Event *DaoFactoryAddExtensionStateEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryAddExtensionStateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryAddExtensionStateEvent)
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
		it.Event = new(DaoFactoryAddExtensionStateEvent)
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
func (it *DaoFactoryAddExtensionStateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryAddExtensionStateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryAddExtensionStateEvent represents a AddExtensionStateEvent event raised by the DaoFactory contract.
type DaoFactoryAddExtensionStateEvent struct {
	Name     string
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddExtensionStateEvent is a free log retrieval operation binding the contract event 0x55f2ef85a2dad56945a1a12578feb70732a67aeea619371c4577c7b0f899dde6.
//
// Solidity: event addExtensionStateEvent(string name, address operator)
func (_DaoFactory *DaoFactoryFilterer) FilterAddExtensionStateEvent(opts *bind.FilterOpts) (*DaoFactoryAddExtensionStateEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "addExtensionStateEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryAddExtensionStateEventIterator{contract: _DaoFactory.contract, event: "addExtensionStateEvent", logs: logs, sub: sub}, nil
}

// WatchAddExtensionStateEvent is a free log subscription operation binding the contract event 0x55f2ef85a2dad56945a1a12578feb70732a67aeea619371c4577c7b0f899dde6.
//
// Solidity: event addExtensionStateEvent(string name, address operator)
func (_DaoFactory *DaoFactoryFilterer) WatchAddExtensionStateEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryAddExtensionStateEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "addExtensionStateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryAddExtensionStateEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "addExtensionStateEvent", log); err != nil {
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

// ParseAddExtensionStateEvent is a log parse operation binding the contract event 0x55f2ef85a2dad56945a1a12578feb70732a67aeea619371c4577c7b0f899dde6.
//
// Solidity: event addExtensionStateEvent(string name, address operator)
func (_DaoFactory *DaoFactoryFilterer) ParseAddExtensionStateEvent(log types.Log) (*DaoFactoryAddExtensionStateEvent, error) {
	event := new(DaoFactoryAddExtensionStateEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "addExtensionStateEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryAllowQuoteEventIterator is returned from FilterAllowQuoteEvent and is used to iterate over the raw logs and unpacked data for AllowQuoteEvent events raised by the DaoFactory contract.
type DaoFactoryAllowQuoteEventIterator struct {
	Event *DaoFactoryAllowQuoteEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryAllowQuoteEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryAllowQuoteEvent)
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
		it.Event = new(DaoFactoryAllowQuoteEvent)
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
func (it *DaoFactoryAllowQuoteEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryAllowQuoteEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryAllowQuoteEvent represents a AllowQuoteEvent event raised by the DaoFactory contract.
type DaoFactoryAllowQuoteEvent struct {
	Operator common.Address
	Addr     common.Address
	Quote    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowQuoteEvent is a free log retrieval operation binding the contract event 0x4852588adc652efe5cc552dee0c2e0709ae2bfc1993c40915af86d6b6884e7f9.
//
// Solidity: event allowQuoteEvent(address operator, address addr, uint256 quote)
func (_DaoFactory *DaoFactoryFilterer) FilterAllowQuoteEvent(opts *bind.FilterOpts) (*DaoFactoryAllowQuoteEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "allowQuoteEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryAllowQuoteEventIterator{contract: _DaoFactory.contract, event: "allowQuoteEvent", logs: logs, sub: sub}, nil
}

// WatchAllowQuoteEvent is a free log subscription operation binding the contract event 0x4852588adc652efe5cc552dee0c2e0709ae2bfc1993c40915af86d6b6884e7f9.
//
// Solidity: event allowQuoteEvent(address operator, address addr, uint256 quote)
func (_DaoFactory *DaoFactoryFilterer) WatchAllowQuoteEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryAllowQuoteEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "allowQuoteEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryAllowQuoteEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "allowQuoteEvent", log); err != nil {
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

// ParseAllowQuoteEvent is a log parse operation binding the contract event 0x4852588adc652efe5cc552dee0c2e0709ae2bfc1993c40915af86d6b6884e7f9.
//
// Solidity: event allowQuoteEvent(address operator, address addr, uint256 quote)
func (_DaoFactory *DaoFactoryFilterer) ParseAllowQuoteEvent(log types.Log) (*DaoFactoryAllowQuoteEvent, error) {
	event := new(DaoFactoryAllowQuoteEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "allowQuoteEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryBeforeUpgradeHookEventIterator is returned from FilterBeforeUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for BeforeUpgradeHookEvent events raised by the DaoFactory contract.
type DaoFactoryBeforeUpgradeHookEventIterator struct {
	Event *DaoFactoryBeforeUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryBeforeUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryBeforeUpgradeHookEvent)
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
		it.Event = new(DaoFactoryBeforeUpgradeHookEvent)
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
func (it *DaoFactoryBeforeUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryBeforeUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryBeforeUpgradeHookEvent represents a BeforeUpgradeHookEvent event raised by the DaoFactory contract.
type DaoFactoryBeforeUpgradeHookEvent struct {
	NewFactory common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBeforeUpgradeHookEvent is a free log retrieval operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newFactory)
func (_DaoFactory *DaoFactoryFilterer) FilterBeforeUpgradeHookEvent(opts *bind.FilterOpts) (*DaoFactoryBeforeUpgradeHookEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryBeforeUpgradeHookEventIterator{contract: _DaoFactory.contract, event: "beforeUpgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchBeforeUpgradeHookEvent is a free log subscription operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newFactory)
func (_DaoFactory *DaoFactoryFilterer) WatchBeforeUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryBeforeUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryBeforeUpgradeHookEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
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
// Solidity: event beforeUpgradeHookEvent(address newFactory)
func (_DaoFactory *DaoFactoryFilterer) ParseBeforeUpgradeHookEvent(log types.Log) (*DaoFactoryBeforeUpgradeHookEvent, error) {
	event := new(DaoFactoryBeforeUpgradeHookEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryChangeInitialDataEventIterator is returned from FilterChangeInitialDataEvent and is used to iterate over the raw logs and unpacked data for ChangeInitialDataEvent events raised by the DaoFactory contract.
type DaoFactoryChangeInitialDataEventIterator struct {
	Event *DaoFactoryChangeInitialDataEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryChangeInitialDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryChangeInitialDataEvent)
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
		it.Event = new(DaoFactoryChangeInitialDataEvent)
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
func (it *DaoFactoryChangeInitialDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryChangeInitialDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryChangeInitialDataEvent represents a ChangeInitialDataEvent event raised by the DaoFactory contract.
type DaoFactoryChangeInitialDataEvent struct {
	State   common.Address
	Manager common.Address
	Layers  []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangeInitialDataEvent is a free log retrieval operation binding the contract event 0x135a605ee368dd19332630115f8b314652b55414d6e98baad12dd2cc6d1b7a3d.
//
// Solidity: event changeInitialDataEvent(address state, address manager, address[] layers)
func (_DaoFactory *DaoFactoryFilterer) FilterChangeInitialDataEvent(opts *bind.FilterOpts) (*DaoFactoryChangeInitialDataEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryChangeInitialDataEventIterator{contract: _DaoFactory.contract, event: "changeInitialDataEvent", logs: logs, sub: sub}, nil
}

// WatchChangeInitialDataEvent is a free log subscription operation binding the contract event 0x135a605ee368dd19332630115f8b314652b55414d6e98baad12dd2cc6d1b7a3d.
//
// Solidity: event changeInitialDataEvent(address state, address manager, address[] layers)
func (_DaoFactory *DaoFactoryFilterer) WatchChangeInitialDataEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryChangeInitialDataEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryChangeInitialDataEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
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

// ParseChangeInitialDataEvent is a log parse operation binding the contract event 0x135a605ee368dd19332630115f8b314652b55414d6e98baad12dd2cc6d1b7a3d.
//
// Solidity: event changeInitialDataEvent(address state, address manager, address[] layers)
func (_DaoFactory *DaoFactoryFilterer) ParseChangeInitialDataEvent(log types.Log) (*DaoFactoryChangeInitialDataEvent, error) {
	event := new(DaoFactoryChangeInitialDataEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryDaoReadyEventIterator is returned from FilterDaoReadyEvent and is used to iterate over the raw logs and unpacked data for DaoReadyEvent events raised by the DaoFactory contract.
type DaoFactoryDaoReadyEventIterator struct {
	Event *DaoFactoryDaoReadyEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryDaoReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryDaoReadyEvent)
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
		it.Event = new(DaoFactoryDaoReadyEvent)
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
func (it *DaoFactoryDaoReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryDaoReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryDaoReadyEvent represents a DaoReadyEvent event raised by the DaoFactory contract.
type DaoFactoryDaoReadyEvent struct {
	Name     string
	Registry common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDaoReadyEvent is a free log retrieval operation binding the contract event 0x873f2c944c570d1e2bdd79d434aa9d0fe0898425993ddd509fbebb9a7c4c5209.
//
// Solidity: event daoReadyEvent(string name, address registry)
func (_DaoFactory *DaoFactoryFilterer) FilterDaoReadyEvent(opts *bind.FilterOpts) (*DaoFactoryDaoReadyEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "daoReadyEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryDaoReadyEventIterator{contract: _DaoFactory.contract, event: "daoReadyEvent", logs: logs, sub: sub}, nil
}

// WatchDaoReadyEvent is a free log subscription operation binding the contract event 0x873f2c944c570d1e2bdd79d434aa9d0fe0898425993ddd509fbebb9a7c4c5209.
//
// Solidity: event daoReadyEvent(string name, address registry)
func (_DaoFactory *DaoFactoryFilterer) WatchDaoReadyEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryDaoReadyEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "daoReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryDaoReadyEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "daoReadyEvent", log); err != nil {
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

// ParseDaoReadyEvent is a log parse operation binding the contract event 0x873f2c944c570d1e2bdd79d434aa9d0fe0898425993ddd509fbebb9a7c4c5209.
//
// Solidity: event daoReadyEvent(string name, address registry)
func (_DaoFactory *DaoFactoryFilterer) ParseDaoReadyEvent(log types.Log) (*DaoFactoryDaoReadyEvent, error) {
	event := new(DaoFactoryDaoReadyEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "daoReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryDeployDaoRegistryEventIterator is returned from FilterDeployDaoRegistryEvent and is used to iterate over the raw logs and unpacked data for DeployDaoRegistryEvent events raised by the DaoFactory contract.
type DaoFactoryDeployDaoRegistryEventIterator struct {
	Event *DaoFactoryDeployDaoRegistryEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryDeployDaoRegistryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryDeployDaoRegistryEvent)
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
		it.Event = new(DaoFactoryDeployDaoRegistryEvent)
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
func (it *DaoFactoryDeployDaoRegistryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryDeployDaoRegistryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryDeployDaoRegistryEvent represents a DeployDaoRegistryEvent event raised by the DaoFactory contract.
type DaoFactoryDeployDaoRegistryEvent struct {
	Name            string
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeployDaoRegistryEvent is a free log retrieval operation binding the contract event 0x41e2b03bcea94d6fe49454a370f737f3c9aa0cfee946f674c5ede49bd6257d54.
//
// Solidity: event deployDaoRegistryEvent(string name, address contractAddress)
func (_DaoFactory *DaoFactoryFilterer) FilterDeployDaoRegistryEvent(opts *bind.FilterOpts) (*DaoFactoryDeployDaoRegistryEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "deployDaoRegistryEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryDeployDaoRegistryEventIterator{contract: _DaoFactory.contract, event: "deployDaoRegistryEvent", logs: logs, sub: sub}, nil
}

// WatchDeployDaoRegistryEvent is a free log subscription operation binding the contract event 0x41e2b03bcea94d6fe49454a370f737f3c9aa0cfee946f674c5ede49bd6257d54.
//
// Solidity: event deployDaoRegistryEvent(string name, address contractAddress)
func (_DaoFactory *DaoFactoryFilterer) WatchDeployDaoRegistryEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryDeployDaoRegistryEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "deployDaoRegistryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryDeployDaoRegistryEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "deployDaoRegistryEvent", log); err != nil {
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

// ParseDeployDaoRegistryEvent is a log parse operation binding the contract event 0x41e2b03bcea94d6fe49454a370f737f3c9aa0cfee946f674c5ede49bd6257d54.
//
// Solidity: event deployDaoRegistryEvent(string name, address contractAddress)
func (_DaoFactory *DaoFactoryFilterer) ParseDeployDaoRegistryEvent(log types.Log) (*DaoFactoryDeployDaoRegistryEvent, error) {
	event := new(DaoFactoryDeployDaoRegistryEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "deployDaoRegistryEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryDeployExtensionEventIterator is returned from FilterDeployExtensionEvent and is used to iterate over the raw logs and unpacked data for DeployExtensionEvent events raised by the DaoFactory contract.
type DaoFactoryDeployExtensionEventIterator struct {
	Event *DaoFactoryDeployExtensionEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryDeployExtensionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryDeployExtensionEvent)
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
		it.Event = new(DaoFactoryDeployExtensionEvent)
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
func (it *DaoFactoryDeployExtensionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryDeployExtensionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryDeployExtensionEvent represents a DeployExtensionEvent event raised by the DaoFactory contract.
type DaoFactoryDeployExtensionEvent struct {
	Name string
	Ext  common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeployExtensionEvent is a free log retrieval operation binding the contract event 0x0181ce0a3a8916dc2645f16fed1c4a0faa7364b32ecefef12a62ff7070941fc1.
//
// Solidity: event deployExtensionEvent(string name, address ext)
func (_DaoFactory *DaoFactoryFilterer) FilterDeployExtensionEvent(opts *bind.FilterOpts) (*DaoFactoryDeployExtensionEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "deployExtensionEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryDeployExtensionEventIterator{contract: _DaoFactory.contract, event: "deployExtensionEvent", logs: logs, sub: sub}, nil
}

// WatchDeployExtensionEvent is a free log subscription operation binding the contract event 0x0181ce0a3a8916dc2645f16fed1c4a0faa7364b32ecefef12a62ff7070941fc1.
//
// Solidity: event deployExtensionEvent(string name, address ext)
func (_DaoFactory *DaoFactoryFilterer) WatchDeployExtensionEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryDeployExtensionEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "deployExtensionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryDeployExtensionEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "deployExtensionEvent", log); err != nil {
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

// ParseDeployExtensionEvent is a log parse operation binding the contract event 0x0181ce0a3a8916dc2645f16fed1c4a0faa7364b32ecefef12a62ff7070941fc1.
//
// Solidity: event deployExtensionEvent(string name, address ext)
func (_DaoFactory *DaoFactoryFilterer) ParseDeployExtensionEvent(log types.Log) (*DaoFactoryDeployExtensionEvent, error) {
	event := new(DaoFactoryDeployExtensionEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "deployExtensionEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the DaoFactory contract.
type DaoFactoryInitializeEventIterator struct {
	Event *DaoFactoryInitializeEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryInitializeEvent)
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
		it.Event = new(DaoFactoryInitializeEvent)
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
func (it *DaoFactoryInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryInitializeEvent represents a InitializeEvent event raised by the DaoFactory contract.
type DaoFactoryInitializeEvent struct {
	Addrs []common.Address
	Ready bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0x2430551ceb4aab64f507476c57d0e80446a521f54684d4f5664a563619e0b9b9.
//
// Solidity: event initializeEvent(address[] addrs, bool ready)
func (_DaoFactory *DaoFactoryFilterer) FilterInitializeEvent(opts *bind.FilterOpts) (*DaoFactoryInitializeEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryInitializeEventIterator{contract: _DaoFactory.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0x2430551ceb4aab64f507476c57d0e80446a521f54684d4f5664a563619e0b9b9.
//
// Solidity: event initializeEvent(address[] addrs, bool ready)
func (_DaoFactory *DaoFactoryFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryInitializeEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryInitializeEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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

// ParseInitializeEvent is a log parse operation binding the contract event 0x2430551ceb4aab64f507476c57d0e80446a521f54684d4f5664a563619e0b9b9.
//
// Solidity: event initializeEvent(address[] addrs, bool ready)
func (_DaoFactory *DaoFactoryFilterer) ParseInitializeEvent(log types.Log) (*DaoFactoryInitializeEvent, error) {
	event := new(DaoFactoryInitializeEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryInitializeHookEventIterator is returned from FilterInitializeHookEvent and is used to iterate over the raw logs and unpacked data for InitializeHookEvent events raised by the DaoFactory contract.
type DaoFactoryInitializeHookEventIterator struct {
	Event *DaoFactoryInitializeHookEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryInitializeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryInitializeHookEvent)
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
		it.Event = new(DaoFactoryInitializeHookEvent)
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
func (it *DaoFactoryInitializeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryInitializeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryInitializeHookEvent represents a InitializeHookEvent event raised by the DaoFactory contract.
type DaoFactoryInitializeHookEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeHookEvent is a free log retrieval operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_DaoFactory *DaoFactoryFilterer) FilterInitializeHookEvent(opts *bind.FilterOpts) (*DaoFactoryInitializeHookEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryInitializeHookEventIterator{contract: _DaoFactory.contract, event: "initializeHookEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeHookEvent is a free log subscription operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_DaoFactory *DaoFactoryFilterer) WatchInitializeHookEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryInitializeHookEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryInitializeHookEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
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
func (_DaoFactory *DaoFactoryFilterer) ParseInitializeHookEvent(log types.Log) (*DaoFactoryInitializeHookEvent, error) {
	event := new(DaoFactoryInitializeHookEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryRejectDaoRequestEventIterator is returned from FilterRejectDaoRequestEvent and is used to iterate over the raw logs and unpacked data for RejectDaoRequestEvent events raised by the DaoFactory contract.
type DaoFactoryRejectDaoRequestEventIterator struct {
	Event *DaoFactoryRejectDaoRequestEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryRejectDaoRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryRejectDaoRequestEvent)
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
		it.Event = new(DaoFactoryRejectDaoRequestEvent)
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
func (it *DaoFactoryRejectDaoRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryRejectDaoRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryRejectDaoRequestEvent represents a RejectDaoRequestEvent event raised by the DaoFactory contract.
type DaoFactoryRejectDaoRequestEvent struct {
	Name   string
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRejectDaoRequestEvent is a free log retrieval operation binding the contract event 0x7f6b704344d0588f7971c7b78c632ca0a282909956856b23d6ae2d8fb4fba831.
//
// Solidity: event rejectDaoRequestEvent(string name, string reason)
func (_DaoFactory *DaoFactoryFilterer) FilterRejectDaoRequestEvent(opts *bind.FilterOpts) (*DaoFactoryRejectDaoRequestEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "rejectDaoRequestEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryRejectDaoRequestEventIterator{contract: _DaoFactory.contract, event: "rejectDaoRequestEvent", logs: logs, sub: sub}, nil
}

// WatchRejectDaoRequestEvent is a free log subscription operation binding the contract event 0x7f6b704344d0588f7971c7b78c632ca0a282909956856b23d6ae2d8fb4fba831.
//
// Solidity: event rejectDaoRequestEvent(string name, string reason)
func (_DaoFactory *DaoFactoryFilterer) WatchRejectDaoRequestEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryRejectDaoRequestEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "rejectDaoRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryRejectDaoRequestEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "rejectDaoRequestEvent", log); err != nil {
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

// ParseRejectDaoRequestEvent is a log parse operation binding the contract event 0x7f6b704344d0588f7971c7b78c632ca0a282909956856b23d6ae2d8fb4fba831.
//
// Solidity: event rejectDaoRequestEvent(string name, string reason)
func (_DaoFactory *DaoFactoryFilterer) ParseRejectDaoRequestEvent(log types.Log) (*DaoFactoryRejectDaoRequestEvent, error) {
	event := new(DaoFactoryRejectDaoRequestEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "rejectDaoRequestEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactorySetCreationFeeEventIterator is returned from FilterSetCreationFeeEvent and is used to iterate over the raw logs and unpacked data for SetCreationFeeEvent events raised by the DaoFactory contract.
type DaoFactorySetCreationFeeEventIterator struct {
	Event *DaoFactorySetCreationFeeEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactorySetCreationFeeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactorySetCreationFeeEvent)
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
		it.Event = new(DaoFactorySetCreationFeeEvent)
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
func (it *DaoFactorySetCreationFeeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactorySetCreationFeeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactorySetCreationFeeEvent represents a SetCreationFeeEvent event raised by the DaoFactory contract.
type DaoFactorySetCreationFeeEvent struct {
	Operator common.Address
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetCreationFeeEvent is a free log retrieval operation binding the contract event 0xcb3cba1daa3013658c59be32c3b876a1f16905b023a550d0e63fb0dee096f19d.
//
// Solidity: event setCreationFeeEvent(address operator, uint256 fee)
func (_DaoFactory *DaoFactoryFilterer) FilterSetCreationFeeEvent(opts *bind.FilterOpts) (*DaoFactorySetCreationFeeEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "setCreationFeeEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactorySetCreationFeeEventIterator{contract: _DaoFactory.contract, event: "setCreationFeeEvent", logs: logs, sub: sub}, nil
}

// WatchSetCreationFeeEvent is a free log subscription operation binding the contract event 0xcb3cba1daa3013658c59be32c3b876a1f16905b023a550d0e63fb0dee096f19d.
//
// Solidity: event setCreationFeeEvent(address operator, uint256 fee)
func (_DaoFactory *DaoFactoryFilterer) WatchSetCreationFeeEvent(opts *bind.WatchOpts, sink chan<- *DaoFactorySetCreationFeeEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "setCreationFeeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactorySetCreationFeeEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "setCreationFeeEvent", log); err != nil {
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

// ParseSetCreationFeeEvent is a log parse operation binding the contract event 0xcb3cba1daa3013658c59be32c3b876a1f16905b023a550d0e63fb0dee096f19d.
//
// Solidity: event setCreationFeeEvent(address operator, uint256 fee)
func (_DaoFactory *DaoFactoryFilterer) ParseSetCreationFeeEvent(log types.Log) (*DaoFactorySetCreationFeeEvent, error) {
	event := new(DaoFactorySetCreationFeeEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "setCreationFeeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactorySetDaoRegistryBytecodeEventIterator is returned from FilterSetDaoRegistryBytecodeEvent and is used to iterate over the raw logs and unpacked data for SetDaoRegistryBytecodeEvent events raised by the DaoFactory contract.
type DaoFactorySetDaoRegistryBytecodeEventIterator struct {
	Event *DaoFactorySetDaoRegistryBytecodeEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactorySetDaoRegistryBytecodeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactorySetDaoRegistryBytecodeEvent)
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
		it.Event = new(DaoFactorySetDaoRegistryBytecodeEvent)
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
func (it *DaoFactorySetDaoRegistryBytecodeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactorySetDaoRegistryBytecodeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactorySetDaoRegistryBytecodeEvent represents a SetDaoRegistryBytecodeEvent event raised by the DaoFactory contract.
type DaoFactorySetDaoRegistryBytecodeEvent struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetDaoRegistryBytecodeEvent is a free log retrieval operation binding the contract event 0x5c81d0aa81a01914b4c925f7534f928b3f5540533e6564c1dc56fe6e5bf4ecc4.
//
// Solidity: event setDaoRegistryBytecodeEvent(address operator)
func (_DaoFactory *DaoFactoryFilterer) FilterSetDaoRegistryBytecodeEvent(opts *bind.FilterOpts) (*DaoFactorySetDaoRegistryBytecodeEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "setDaoRegistryBytecodeEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactorySetDaoRegistryBytecodeEventIterator{contract: _DaoFactory.contract, event: "setDaoRegistryBytecodeEvent", logs: logs, sub: sub}, nil
}

// WatchSetDaoRegistryBytecodeEvent is a free log subscription operation binding the contract event 0x5c81d0aa81a01914b4c925f7534f928b3f5540533e6564c1dc56fe6e5bf4ecc4.
//
// Solidity: event setDaoRegistryBytecodeEvent(address operator)
func (_DaoFactory *DaoFactoryFilterer) WatchSetDaoRegistryBytecodeEvent(opts *bind.WatchOpts, sink chan<- *DaoFactorySetDaoRegistryBytecodeEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "setDaoRegistryBytecodeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactorySetDaoRegistryBytecodeEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "setDaoRegistryBytecodeEvent", log); err != nil {
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

// ParseSetDaoRegistryBytecodeEvent is a log parse operation binding the contract event 0x5c81d0aa81a01914b4c925f7534f928b3f5540533e6564c1dc56fe6e5bf4ecc4.
//
// Solidity: event setDaoRegistryBytecodeEvent(address operator)
func (_DaoFactory *DaoFactoryFilterer) ParseSetDaoRegistryBytecodeEvent(log types.Log) (*DaoFactorySetDaoRegistryBytecodeEvent, error) {
	event := new(DaoFactorySetDaoRegistryBytecodeEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "setDaoRegistryBytecodeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactorySetDaoRegistryStateBytecodeEventIterator is returned from FilterSetDaoRegistryStateBytecodeEvent and is used to iterate over the raw logs and unpacked data for SetDaoRegistryStateBytecodeEvent events raised by the DaoFactory contract.
type DaoFactorySetDaoRegistryStateBytecodeEventIterator struct {
	Event *DaoFactorySetDaoRegistryStateBytecodeEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactorySetDaoRegistryStateBytecodeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactorySetDaoRegistryStateBytecodeEvent)
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
		it.Event = new(DaoFactorySetDaoRegistryStateBytecodeEvent)
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
func (it *DaoFactorySetDaoRegistryStateBytecodeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactorySetDaoRegistryStateBytecodeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactorySetDaoRegistryStateBytecodeEvent represents a SetDaoRegistryStateBytecodeEvent event raised by the DaoFactory contract.
type DaoFactorySetDaoRegistryStateBytecodeEvent struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetDaoRegistryStateBytecodeEvent is a free log retrieval operation binding the contract event 0x506b6d6b9acb4f72c20c3bf484e68c0e2397beb9fb4e85e543b780fabae62de3.
//
// Solidity: event setDaoRegistryStateBytecodeEvent(address operator)
func (_DaoFactory *DaoFactoryFilterer) FilterSetDaoRegistryStateBytecodeEvent(opts *bind.FilterOpts) (*DaoFactorySetDaoRegistryStateBytecodeEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "setDaoRegistryStateBytecodeEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactorySetDaoRegistryStateBytecodeEventIterator{contract: _DaoFactory.contract, event: "setDaoRegistryStateBytecodeEvent", logs: logs, sub: sub}, nil
}

// WatchSetDaoRegistryStateBytecodeEvent is a free log subscription operation binding the contract event 0x506b6d6b9acb4f72c20c3bf484e68c0e2397beb9fb4e85e543b780fabae62de3.
//
// Solidity: event setDaoRegistryStateBytecodeEvent(address operator)
func (_DaoFactory *DaoFactoryFilterer) WatchSetDaoRegistryStateBytecodeEvent(opts *bind.WatchOpts, sink chan<- *DaoFactorySetDaoRegistryStateBytecodeEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "setDaoRegistryStateBytecodeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactorySetDaoRegistryStateBytecodeEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "setDaoRegistryStateBytecodeEvent", log); err != nil {
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

// ParseSetDaoRegistryStateBytecodeEvent is a log parse operation binding the contract event 0x506b6d6b9acb4f72c20c3bf484e68c0e2397beb9fb4e85e543b780fabae62de3.
//
// Solidity: event setDaoRegistryStateBytecodeEvent(address operator)
func (_DaoFactory *DaoFactoryFilterer) ParseSetDaoRegistryStateBytecodeEvent(log types.Log) (*DaoFactorySetDaoRegistryStateBytecodeEvent, error) {
	event := new(DaoFactorySetDaoRegistryStateBytecodeEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "setDaoRegistryStateBytecodeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactorySetLayersEventIterator is returned from FilterSetLayersEvent and is used to iterate over the raw logs and unpacked data for SetLayersEvent events raised by the DaoFactory contract.
type DaoFactorySetLayersEventIterator struct {
	Event *DaoFactorySetLayersEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactorySetLayersEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactorySetLayersEvent)
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
		it.Event = new(DaoFactorySetLayersEvent)
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
func (it *DaoFactorySetLayersEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactorySetLayersEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactorySetLayersEvent represents a SetLayersEvent event raised by the DaoFactory contract.
type DaoFactorySetLayersEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetLayersEvent is a free log retrieval operation binding the contract event 0x84d60135ac749ffa638f873730272f2fbf40626bd74039558a06f829ca9ceacc.
//
// Solidity: event setLayersEvent(address[] addrs)
func (_DaoFactory *DaoFactoryFilterer) FilterSetLayersEvent(opts *bind.FilterOpts) (*DaoFactorySetLayersEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "setLayersEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactorySetLayersEventIterator{contract: _DaoFactory.contract, event: "setLayersEvent", logs: logs, sub: sub}, nil
}

// WatchSetLayersEvent is a free log subscription operation binding the contract event 0x84d60135ac749ffa638f873730272f2fbf40626bd74039558a06f829ca9ceacc.
//
// Solidity: event setLayersEvent(address[] addrs)
func (_DaoFactory *DaoFactoryFilterer) WatchSetLayersEvent(opts *bind.WatchOpts, sink chan<- *DaoFactorySetLayersEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "setLayersEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactorySetLayersEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "setLayersEvent", log); err != nil {
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

// ParseSetLayersEvent is a log parse operation binding the contract event 0x84d60135ac749ffa638f873730272f2fbf40626bd74039558a06f829ca9ceacc.
//
// Solidity: event setLayersEvent(address[] addrs)
func (_DaoFactory *DaoFactoryFilterer) ParseSetLayersEvent(log types.Log) (*DaoFactorySetLayersEvent, error) {
	event := new(DaoFactorySetLayersEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "setLayersEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactorySetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the DaoFactory contract.
type DaoFactorySetReadyEventIterator struct {
	Event *DaoFactorySetReadyEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactorySetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactorySetReadyEvent)
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
		it.Event = new(DaoFactorySetReadyEvent)
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
func (it *DaoFactorySetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactorySetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactorySetReadyEvent represents a SetReadyEvent event raised by the DaoFactory contract.
type DaoFactorySetReadyEvent struct {
	R   bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool r)
func (_DaoFactory *DaoFactoryFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*DaoFactorySetReadyEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactorySetReadyEventIterator{contract: _DaoFactory.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool r)
func (_DaoFactory *DaoFactoryFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *DaoFactorySetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactorySetReadyEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
func (_DaoFactory *DaoFactoryFilterer) ParseSetReadyEvent(log types.Log) (*DaoFactorySetReadyEvent, error) {
	event := new(DaoFactorySetReadyEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryStoredAppDataEventIterator is returned from FilterStoredAppDataEvent and is used to iterate over the raw logs and unpacked data for StoredAppDataEvent events raised by the DaoFactory contract.
type DaoFactoryStoredAppDataEventIterator struct {
	Event *DaoFactoryStoredAppDataEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryStoredAppDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryStoredAppDataEvent)
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
		it.Event = new(DaoFactoryStoredAppDataEvent)
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
func (it *DaoFactoryStoredAppDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryStoredAppDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryStoredAppDataEvent represents a StoredAppDataEvent event raised by the DaoFactory contract.
type DaoFactoryStoredAppDataEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStoredAppDataEvent is a free log retrieval operation binding the contract event 0xf54e1e77b326c519703e2e17fea0bb6600009edd9597e65edd161c9960094862.
//
// Solidity: event storedAppDataEvent(uint256 activityFlag, bytes data)
func (_DaoFactory *DaoFactoryFilterer) FilterStoredAppDataEvent(opts *bind.FilterOpts) (*DaoFactoryStoredAppDataEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "storedAppDataEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryStoredAppDataEventIterator{contract: _DaoFactory.contract, event: "storedAppDataEvent", logs: logs, sub: sub}, nil
}

// WatchStoredAppDataEvent is a free log subscription operation binding the contract event 0xf54e1e77b326c519703e2e17fea0bb6600009edd9597e65edd161c9960094862.
//
// Solidity: event storedAppDataEvent(uint256 activityFlag, bytes data)
func (_DaoFactory *DaoFactoryFilterer) WatchStoredAppDataEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryStoredAppDataEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "storedAppDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryStoredAppDataEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "storedAppDataEvent", log); err != nil {
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
func (_DaoFactory *DaoFactoryFilterer) ParseStoredAppDataEvent(log types.Log) (*DaoFactoryStoredAppDataEvent, error) {
	event := new(DaoFactoryStoredAppDataEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "storedAppDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryUpgradeContractEventIterator is returned from FilterUpgradeContractEvent and is used to iterate over the raw logs and unpacked data for UpgradeContractEvent events raised by the DaoFactory contract.
type DaoFactoryUpgradeContractEventIterator struct {
	Event *DaoFactoryUpgradeContractEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryUpgradeContractEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryUpgradeContractEvent)
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
		it.Event = new(DaoFactoryUpgradeContractEvent)
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
func (it *DaoFactoryUpgradeContractEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryUpgradeContractEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryUpgradeContractEvent represents a UpgradeContractEvent event raised by the DaoFactory contract.
type DaoFactoryUpgradeContractEvent struct {
	Name     string
	Operator common.Address
	Version  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradeContractEvent is a free log retrieval operation binding the contract event 0x52edb208bd40695023a7d7f0efc05a0fe145f3e6c299182682ae4d02262ffa60.
//
// Solidity: event upgradeContractEvent(string name, address operator, uint256 version)
func (_DaoFactory *DaoFactoryFilterer) FilterUpgradeContractEvent(opts *bind.FilterOpts) (*DaoFactoryUpgradeContractEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "upgradeContractEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryUpgradeContractEventIterator{contract: _DaoFactory.contract, event: "upgradeContractEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeContractEvent is a free log subscription operation binding the contract event 0x52edb208bd40695023a7d7f0efc05a0fe145f3e6c299182682ae4d02262ffa60.
//
// Solidity: event upgradeContractEvent(string name, address operator, uint256 version)
func (_DaoFactory *DaoFactoryFilterer) WatchUpgradeContractEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryUpgradeContractEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "upgradeContractEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryUpgradeContractEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "upgradeContractEvent", log); err != nil {
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

// ParseUpgradeContractEvent is a log parse operation binding the contract event 0x52edb208bd40695023a7d7f0efc05a0fe145f3e6c299182682ae4d02262ffa60.
//
// Solidity: event upgradeContractEvent(string name, address operator, uint256 version)
func (_DaoFactory *DaoFactoryFilterer) ParseUpgradeContractEvent(log types.Log) (*DaoFactoryUpgradeContractEvent, error) {
	event := new(DaoFactoryUpgradeContractEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "upgradeContractEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryUpgradeEventIterator is returned from FilterUpgradeEvent and is used to iterate over the raw logs and unpacked data for UpgradeEvent events raised by the DaoFactory contract.
type DaoFactoryUpgradeEventIterator struct {
	Event *DaoFactoryUpgradeEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryUpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryUpgradeEvent)
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
		it.Event = new(DaoFactoryUpgradeEvent)
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
func (it *DaoFactoryUpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryUpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryUpgradeEvent represents a UpgradeEvent event raised by the DaoFactory contract.
type DaoFactoryUpgradeEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpgradeEvent is a free log retrieval operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_DaoFactory *DaoFactoryFilterer) FilterUpgradeEvent(opts *bind.FilterOpts) (*DaoFactoryUpgradeEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryUpgradeEventIterator{contract: _DaoFactory.contract, event: "upgradeEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeEvent is a free log subscription operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_DaoFactory *DaoFactoryFilterer) WatchUpgradeEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryUpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryUpgradeEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
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
func (_DaoFactory *DaoFactoryFilterer) ParseUpgradeEvent(log types.Log) (*DaoFactoryUpgradeEvent, error) {
	event := new(DaoFactoryUpgradeEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryUpgradeExtensionEventIterator is returned from FilterUpgradeExtensionEvent and is used to iterate over the raw logs and unpacked data for UpgradeExtensionEvent events raised by the DaoFactory contract.
type DaoFactoryUpgradeExtensionEventIterator struct {
	Event *DaoFactoryUpgradeExtensionEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryUpgradeExtensionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryUpgradeExtensionEvent)
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
		it.Event = new(DaoFactoryUpgradeExtensionEvent)
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
func (it *DaoFactoryUpgradeExtensionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryUpgradeExtensionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryUpgradeExtensionEvent represents a UpgradeExtensionEvent event raised by the DaoFactory contract.
type DaoFactoryUpgradeExtensionEvent struct {
	Name     string
	Operator common.Address
	Version  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradeExtensionEvent is a free log retrieval operation binding the contract event 0x21cadf28592892fe417855870e1d5e38c2826383cf3cc30dd367e16f1e1e8921.
//
// Solidity: event upgradeExtensionEvent(string name, address operator, uint256 version)
func (_DaoFactory *DaoFactoryFilterer) FilterUpgradeExtensionEvent(opts *bind.FilterOpts) (*DaoFactoryUpgradeExtensionEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "upgradeExtensionEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryUpgradeExtensionEventIterator{contract: _DaoFactory.contract, event: "upgradeExtensionEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeExtensionEvent is a free log subscription operation binding the contract event 0x21cadf28592892fe417855870e1d5e38c2826383cf3cc30dd367e16f1e1e8921.
//
// Solidity: event upgradeExtensionEvent(string name, address operator, uint256 version)
func (_DaoFactory *DaoFactoryFilterer) WatchUpgradeExtensionEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryUpgradeExtensionEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "upgradeExtensionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryUpgradeExtensionEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "upgradeExtensionEvent", log); err != nil {
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

// ParseUpgradeExtensionEvent is a log parse operation binding the contract event 0x21cadf28592892fe417855870e1d5e38c2826383cf3cc30dd367e16f1e1e8921.
//
// Solidity: event upgradeExtensionEvent(string name, address operator, uint256 version)
func (_DaoFactory *DaoFactoryFilterer) ParseUpgradeExtensionEvent(log types.Log) (*DaoFactoryUpgradeExtensionEvent, error) {
	event := new(DaoFactoryUpgradeExtensionEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "upgradeExtensionEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryUpgradeExtensionStateEventIterator is returned from FilterUpgradeExtensionStateEvent and is used to iterate over the raw logs and unpacked data for UpgradeExtensionStateEvent events raised by the DaoFactory contract.
type DaoFactoryUpgradeExtensionStateEventIterator struct {
	Event *DaoFactoryUpgradeExtensionStateEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryUpgradeExtensionStateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryUpgradeExtensionStateEvent)
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
		it.Event = new(DaoFactoryUpgradeExtensionStateEvent)
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
func (it *DaoFactoryUpgradeExtensionStateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryUpgradeExtensionStateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryUpgradeExtensionStateEvent represents a UpgradeExtensionStateEvent event raised by the DaoFactory contract.
type DaoFactoryUpgradeExtensionStateEvent struct {
	Name     string
	Operator common.Address
	Version  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradeExtensionStateEvent is a free log retrieval operation binding the contract event 0x3aed65a18ecfd4f00013a88faa0cfb1a2df5ee24a657eef46a45c5064131dfbe.
//
// Solidity: event upgradeExtensionStateEvent(string name, address operator, uint256 version)
func (_DaoFactory *DaoFactoryFilterer) FilterUpgradeExtensionStateEvent(opts *bind.FilterOpts) (*DaoFactoryUpgradeExtensionStateEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "upgradeExtensionStateEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryUpgradeExtensionStateEventIterator{contract: _DaoFactory.contract, event: "upgradeExtensionStateEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeExtensionStateEvent is a free log subscription operation binding the contract event 0x3aed65a18ecfd4f00013a88faa0cfb1a2df5ee24a657eef46a45c5064131dfbe.
//
// Solidity: event upgradeExtensionStateEvent(string name, address operator, uint256 version)
func (_DaoFactory *DaoFactoryFilterer) WatchUpgradeExtensionStateEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryUpgradeExtensionStateEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "upgradeExtensionStateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryUpgradeExtensionStateEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "upgradeExtensionStateEvent", log); err != nil {
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

// ParseUpgradeExtensionStateEvent is a log parse operation binding the contract event 0x3aed65a18ecfd4f00013a88faa0cfb1a2df5ee24a657eef46a45c5064131dfbe.
//
// Solidity: event upgradeExtensionStateEvent(string name, address operator, uint256 version)
func (_DaoFactory *DaoFactoryFilterer) ParseUpgradeExtensionStateEvent(log types.Log) (*DaoFactoryUpgradeExtensionStateEvent, error) {
	event := new(DaoFactoryUpgradeExtensionStateEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "upgradeExtensionStateEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryUpgradeHookEventIterator is returned from FilterUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for UpgradeHookEvent events raised by the DaoFactory contract.
type DaoFactoryUpgradeHookEventIterator struct {
	Event *DaoFactoryUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryUpgradeHookEvent)
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
		it.Event = new(DaoFactoryUpgradeHookEvent)
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
func (it *DaoFactoryUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryUpgradeHookEvent represents a UpgradeHookEvent event raised by the DaoFactory contract.
type DaoFactoryUpgradeHookEvent struct {
	State common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgradeHookEvent is a free log retrieval operation binding the contract event 0x3615fad934cb4c7adca8f74fba8c44a2fc2dd7634c87976c6631fae7091edb17.
//
// Solidity: event upgradeHookEvent(address state)
func (_DaoFactory *DaoFactoryFilterer) FilterUpgradeHookEvent(opts *bind.FilterOpts) (*DaoFactoryUpgradeHookEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "upgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryUpgradeHookEventIterator{contract: _DaoFactory.contract, event: "upgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeHookEvent is a free log subscription operation binding the contract event 0x3615fad934cb4c7adca8f74fba8c44a2fc2dd7634c87976c6631fae7091edb17.
//
// Solidity: event upgradeHookEvent(address state)
func (_DaoFactory *DaoFactoryFilterer) WatchUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "upgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryUpgradeHookEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "upgradeHookEvent", log); err != nil {
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

// ParseUpgradeHookEvent is a log parse operation binding the contract event 0x3615fad934cb4c7adca8f74fba8c44a2fc2dd7634c87976c6631fae7091edb17.
//
// Solidity: event upgradeHookEvent(address state)
func (_DaoFactory *DaoFactoryFilterer) ParseUpgradeHookEvent(log types.Log) (*DaoFactoryUpgradeHookEvent, error) {
	event := new(DaoFactoryUpgradeHookEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "upgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryUseDeprecatedFunctionEventIterator is returned from FilterUseDeprecatedFunctionEvent and is used to iterate over the raw logs and unpacked data for UseDeprecatedFunctionEvent events raised by the DaoFactory contract.
type DaoFactoryUseDeprecatedFunctionEventIterator struct {
	Event *DaoFactoryUseDeprecatedFunctionEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryUseDeprecatedFunctionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryUseDeprecatedFunctionEvent)
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
		it.Event = new(DaoFactoryUseDeprecatedFunctionEvent)
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
func (it *DaoFactoryUseDeprecatedFunctionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryUseDeprecatedFunctionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryUseDeprecatedFunctionEvent represents a UseDeprecatedFunctionEvent event raised by the DaoFactory contract.
type DaoFactoryUseDeprecatedFunctionEvent struct {
	Name   string
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUseDeprecatedFunctionEvent is a free log retrieval operation binding the contract event 0x9f38ba4e9c1373070d508d4e05a8530771a534450d09d31e396e3dcb0505cd23.
//
// Solidity: event useDeprecatedFunctionEvent(string name, address sender)
func (_DaoFactory *DaoFactoryFilterer) FilterUseDeprecatedFunctionEvent(opts *bind.FilterOpts) (*DaoFactoryUseDeprecatedFunctionEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "useDeprecatedFunctionEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryUseDeprecatedFunctionEventIterator{contract: _DaoFactory.contract, event: "useDeprecatedFunctionEvent", logs: logs, sub: sub}, nil
}

// WatchUseDeprecatedFunctionEvent is a free log subscription operation binding the contract event 0x9f38ba4e9c1373070d508d4e05a8530771a534450d09d31e396e3dcb0505cd23.
//
// Solidity: event useDeprecatedFunctionEvent(string name, address sender)
func (_DaoFactory *DaoFactoryFilterer) WatchUseDeprecatedFunctionEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryUseDeprecatedFunctionEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "useDeprecatedFunctionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryUseDeprecatedFunctionEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "useDeprecatedFunctionEvent", log); err != nil {
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

// ParseUseDeprecatedFunctionEvent is a log parse operation binding the contract event 0x9f38ba4e9c1373070d508d4e05a8530771a534450d09d31e396e3dcb0505cd23.
//
// Solidity: event useDeprecatedFunctionEvent(string name, address sender)
func (_DaoFactory *DaoFactoryFilterer) ParseUseDeprecatedFunctionEvent(log types.Log) (*DaoFactoryUseDeprecatedFunctionEvent, error) {
	event := new(DaoFactoryUseDeprecatedFunctionEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "useDeprecatedFunctionEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaoFactoryWithdrawalEventIterator is returned from FilterWithdrawalEvent and is used to iterate over the raw logs and unpacked data for WithdrawalEvent events raised by the DaoFactory contract.
type DaoFactoryWithdrawalEventIterator struct {
	Event *DaoFactoryWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *DaoFactoryWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoFactoryWithdrawalEvent)
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
		it.Event = new(DaoFactoryWithdrawalEvent)
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
func (it *DaoFactoryWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoFactoryWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoFactoryWithdrawalEvent represents a WithdrawalEvent event raised by the DaoFactory contract.
type DaoFactoryWithdrawalEvent struct {
	Addr    common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalEvent is a free log retrieval operation binding the contract event 0xdb08d597821b1c166560f377fe9bc04596133de2c5425997a1281594ca871cb9.
//
// Solidity: event withdrawalEvent(address addr, uint256 balance)
func (_DaoFactory *DaoFactoryFilterer) FilterWithdrawalEvent(opts *bind.FilterOpts) (*DaoFactoryWithdrawalEventIterator, error) {

	logs, sub, err := _DaoFactory.contract.FilterLogs(opts, "withdrawalEvent")
	if err != nil {
		return nil, err
	}
	return &DaoFactoryWithdrawalEventIterator{contract: _DaoFactory.contract, event: "withdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawalEvent is a free log subscription operation binding the contract event 0xdb08d597821b1c166560f377fe9bc04596133de2c5425997a1281594ca871cb9.
//
// Solidity: event withdrawalEvent(address addr, uint256 balance)
func (_DaoFactory *DaoFactoryFilterer) WatchWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *DaoFactoryWithdrawalEvent) (event.Subscription, error) {

	logs, sub, err := _DaoFactory.contract.WatchLogs(opts, "withdrawalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoFactoryWithdrawalEvent)
				if err := _DaoFactory.contract.UnpackLog(event, "withdrawalEvent", log); err != nil {
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

// ParseWithdrawalEvent is a log parse operation binding the contract event 0xdb08d597821b1c166560f377fe9bc04596133de2c5425997a1281594ca871cb9.
//
// Solidity: event withdrawalEvent(address addr, uint256 balance)
func (_DaoFactory *DaoFactoryFilterer) ParseWithdrawalEvent(log types.Log) (*DaoFactoryWithdrawalEvent, error) {
	event := new(DaoFactoryWithdrawalEvent)
	if err := _DaoFactory.contract.UnpackLog(event, "withdrawalEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
