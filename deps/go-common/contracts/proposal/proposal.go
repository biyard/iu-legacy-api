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

// ProposalParameter is an auto generated low-level Go binding around an user-defined struct.
type ProposalParameter struct {
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

// ProposedPoicy is an auto generated low-level Go binding around an user-defined struct.
type ProposedPoicy struct {
	Name  string
	Value *big.Int
}

// ProposalMetaData contains all meta data concerning the Proposal contract.
var ProposalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"stateAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"addAcceptersV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"addParticipantsV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"addRejectersV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"beforeUpgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[7]\",\"name\":\"policies\",\"type\":\"string[7]\"},{\"indexed\":false,\"internalType\":\"uint256[7]\",\"name\":\"values\",\"type\":\"uint256[7]\"}],\"name\":\"changeInitialDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitAppDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityType\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structProposalSummary\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"emitProposalActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"check\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"policyValue\",\"type\":\"uint256\"}],\"name\":\"insertPolicyDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"check\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"insertProposalDataV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contents_\",\"type\":\"string\"}],\"name\":\"setCommentV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setProposalParticipantsV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"}],\"name\":\"setSummaryEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setUuidAddressV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nvp\",\"type\":\"uint256\"}],\"name\":\"setVoterPowerV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"updatePolicyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistryEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"addAccepters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"addAcceptersV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"addParticipants\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"addParticipantsV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"addRejecters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"addRejectersV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName\",\"type\":\"string\"}],\"name\":\"checkPolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"checkProposalIndex\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"checkProposalIndexV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName\",\"type\":\"string\"}],\"name\":\"checkProposalPolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"checkuuid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"checkuuidV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getAccepters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getAcceptersV2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getComment\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getCommentV2\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getCvp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getCvpV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getParticipantsV2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolicies\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"}],\"name\":\"getPolicy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getPolicyValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getProposalParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalParticipantsV2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"getProposalStructure\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStructureV2\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameterV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalV2\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameterV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getProposedPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structProposedPoicy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposedPolicyV2\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structProposedPoicy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getRejecters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getRejectersV2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateAddr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSummary\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"getUuidAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getUuidAddressV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getVoterResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getVoterValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getVoterValueV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"check\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"policyValue\",\"type\":\"uint256\"}],\"name\":\"insertPolicyData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"check\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"insertProposalData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"check\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"insertProposalDataV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameExt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"participantsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"participantsLengthV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contents_\",\"type\":\"string\"}],\"name\":\"setComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contents_\",\"type\":\"string\"}],\"name\":\"setCommentV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"setPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"policyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setPolicyData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setProposalParticipants\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setProposalParticipantsV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"}],\"name\":\"setSummary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setUuidAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setUuidAddressV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nvp\",\"type\":\"uint256\"}],\"name\":\"setVoterPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nvp\",\"type\":\"uint256\"}],\"name\":\"setVoterPowerV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"updatePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updatePolicyData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProposalABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalMetaData.ABI instead.
var ProposalABI = ProposalMetaData.ABI

// ProposalBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ProposalBinRuntime = ``

// Proposal is an auto generated Go binding around a Klaytn contract.
type Proposal struct {
	ProposalCaller     // Read-only binding to the contract
	ProposalTransactor // Write-only binding to the contract
	ProposalFilterer   // Log filterer for contract events
}

// ProposalCaller is an auto generated read-only Go binding around a Klaytn contract.
type ProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ProposalSession struct {
	Contract     *Proposal         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ProposalCallerSession struct {
	Contract *ProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProposalTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ProposalTransactorSession struct {
	Contract     *ProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProposalRaw is an auto generated low-level Go binding around a Klaytn contract.
type ProposalRaw struct {
	Contract *Proposal // Generic contract binding to access the raw methods on
}

// ProposalCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ProposalCallerRaw struct {
	Contract *ProposalCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ProposalTransactorRaw struct {
	Contract *ProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposal creates a new instance of Proposal, bound to a specific deployed contract.
func NewProposal(address common.Address, backend bind.ContractBackend) (*Proposal, error) {
	contract, err := bindProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proposal{ProposalCaller: ProposalCaller{contract: contract}, ProposalTransactor: ProposalTransactor{contract: contract}, ProposalFilterer: ProposalFilterer{contract: contract}}, nil
}

// NewProposalCaller creates a new read-only instance of Proposal, bound to a specific deployed contract.
func NewProposalCaller(address common.Address, caller bind.ContractCaller) (*ProposalCaller, error) {
	contract, err := bindProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalCaller{contract: contract}, nil
}

// NewProposalTransactor creates a new write-only instance of Proposal, bound to a specific deployed contract.
func NewProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalTransactor, error) {
	contract, err := bindProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalTransactor{contract: contract}, nil
}

// NewProposalFilterer creates a new log filterer instance of Proposal, bound to a specific deployed contract.
func NewProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalFilterer, error) {
	contract, err := bindProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalFilterer{contract: contract}, nil
}

// bindProposal binds a generic wrapper to an already deployed contract.
func bindProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProposalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposal *ProposalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proposal.Contract.ProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposal *ProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.Contract.ProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposal *ProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposal.Contract.ProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposal *ProposalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposal *ProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposal *ProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposal.Contract.contract.Transact(opts, method, params...)
}

// CheckPolicy is a free data retrieval call binding the contract method 0x60cd3bd6.
//
// Solidity: function checkPolicy(string policyName) view returns(bool)
func (_Proposal *ProposalCaller) CheckPolicy(opts *bind.CallOpts, policyName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "checkPolicy", policyName)
	return *ret0, err
}

// CheckPolicy is a free data retrieval call binding the contract method 0x60cd3bd6.
//
// Solidity: function checkPolicy(string policyName) view returns(bool)
func (_Proposal *ProposalSession) CheckPolicy(policyName string) (bool, error) {
	return _Proposal.Contract.CheckPolicy(&_Proposal.CallOpts, policyName)
}

// CheckPolicy is a free data retrieval call binding the contract method 0x60cd3bd6.
//
// Solidity: function checkPolicy(string policyName) view returns(bool)
func (_Proposal *ProposalCallerSession) CheckPolicy(policyName string) (bool, error) {
	return _Proposal.Contract.CheckPolicy(&_Proposal.CallOpts, policyName)
}

// CheckProposalIndex is a free data retrieval call binding the contract method 0x4c4fe4aa.
//
// Solidity: function checkProposalIndex(string uuid) view returns(bool)
func (_Proposal *ProposalCaller) CheckProposalIndex(opts *bind.CallOpts, uuid string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "checkProposalIndex", uuid)
	return *ret0, err
}

// CheckProposalIndex is a free data retrieval call binding the contract method 0x4c4fe4aa.
//
// Solidity: function checkProposalIndex(string uuid) view returns(bool)
func (_Proposal *ProposalSession) CheckProposalIndex(uuid string) (bool, error) {
	return _Proposal.Contract.CheckProposalIndex(&_Proposal.CallOpts, uuid)
}

// CheckProposalIndex is a free data retrieval call binding the contract method 0x4c4fe4aa.
//
// Solidity: function checkProposalIndex(string uuid) view returns(bool)
func (_Proposal *ProposalCallerSession) CheckProposalIndex(uuid string) (bool, error) {
	return _Proposal.Contract.CheckProposalIndex(&_Proposal.CallOpts, uuid)
}

// CheckProposalIndexV2 is a free data retrieval call binding the contract method 0xcab9a900.
//
// Solidity: function checkProposalIndexV2(uint256 proposalId) view returns(bool)
func (_Proposal *ProposalCaller) CheckProposalIndexV2(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "checkProposalIndexV2", proposalId)
	return *ret0, err
}

// CheckProposalIndexV2 is a free data retrieval call binding the contract method 0xcab9a900.
//
// Solidity: function checkProposalIndexV2(uint256 proposalId) view returns(bool)
func (_Proposal *ProposalSession) CheckProposalIndexV2(proposalId *big.Int) (bool, error) {
	return _Proposal.Contract.CheckProposalIndexV2(&_Proposal.CallOpts, proposalId)
}

// CheckProposalIndexV2 is a free data retrieval call binding the contract method 0xcab9a900.
//
// Solidity: function checkProposalIndexV2(uint256 proposalId) view returns(bool)
func (_Proposal *ProposalCallerSession) CheckProposalIndexV2(proposalId *big.Int) (bool, error) {
	return _Proposal.Contract.CheckProposalIndexV2(&_Proposal.CallOpts, proposalId)
}

// CheckProposalPolicy is a free data retrieval call binding the contract method 0x7521ca59.
//
// Solidity: function checkProposalPolicy(string policyName) view returns(bool)
func (_Proposal *ProposalCaller) CheckProposalPolicy(opts *bind.CallOpts, policyName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "checkProposalPolicy", policyName)
	return *ret0, err
}

// CheckProposalPolicy is a free data retrieval call binding the contract method 0x7521ca59.
//
// Solidity: function checkProposalPolicy(string policyName) view returns(bool)
func (_Proposal *ProposalSession) CheckProposalPolicy(policyName string) (bool, error) {
	return _Proposal.Contract.CheckProposalPolicy(&_Proposal.CallOpts, policyName)
}

// CheckProposalPolicy is a free data retrieval call binding the contract method 0x7521ca59.
//
// Solidity: function checkProposalPolicy(string policyName) view returns(bool)
func (_Proposal *ProposalCallerSession) CheckProposalPolicy(policyName string) (bool, error) {
	return _Proposal.Contract.CheckProposalPolicy(&_Proposal.CallOpts, policyName)
}

// Checkuuid is a free data retrieval call binding the contract method 0x6f202a8f.
//
// Solidity: function checkuuid(string uuid) view returns(bool)
func (_Proposal *ProposalCaller) Checkuuid(opts *bind.CallOpts, uuid string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "checkuuid", uuid)
	return *ret0, err
}

// Checkuuid is a free data retrieval call binding the contract method 0x6f202a8f.
//
// Solidity: function checkuuid(string uuid) view returns(bool)
func (_Proposal *ProposalSession) Checkuuid(uuid string) (bool, error) {
	return _Proposal.Contract.Checkuuid(&_Proposal.CallOpts, uuid)
}

// Checkuuid is a free data retrieval call binding the contract method 0x6f202a8f.
//
// Solidity: function checkuuid(string uuid) view returns(bool)
func (_Proposal *ProposalCallerSession) Checkuuid(uuid string) (bool, error) {
	return _Proposal.Contract.Checkuuid(&_Proposal.CallOpts, uuid)
}

// CheckuuidV2 is a free data retrieval call binding the contract method 0xc8d0ca4c.
//
// Solidity: function checkuuidV2(uint256 proposalId) view returns(bool)
func (_Proposal *ProposalCaller) CheckuuidV2(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "checkuuidV2", proposalId)
	return *ret0, err
}

// CheckuuidV2 is a free data retrieval call binding the contract method 0xc8d0ca4c.
//
// Solidity: function checkuuidV2(uint256 proposalId) view returns(bool)
func (_Proposal *ProposalSession) CheckuuidV2(proposalId *big.Int) (bool, error) {
	return _Proposal.Contract.CheckuuidV2(&_Proposal.CallOpts, proposalId)
}

// CheckuuidV2 is a free data retrieval call binding the contract method 0xc8d0ca4c.
//
// Solidity: function checkuuidV2(uint256 proposalId) view returns(bool)
func (_Proposal *ProposalCallerSession) CheckuuidV2(proposalId *big.Int) (bool, error) {
	return _Proposal.Contract.CheckuuidV2(&_Proposal.CallOpts, proposalId)
}

// GetAccepters is a free data retrieval call binding the contract method 0x1d774ba8.
//
// Solidity: function getAccepters(string uuid_) view returns(address[])
func (_Proposal *ProposalCaller) GetAccepters(opts *bind.CallOpts, uuid_ string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getAccepters", uuid_)
	return *ret0, err
}

// GetAccepters is a free data retrieval call binding the contract method 0x1d774ba8.
//
// Solidity: function getAccepters(string uuid_) view returns(address[])
func (_Proposal *ProposalSession) GetAccepters(uuid_ string) ([]common.Address, error) {
	return _Proposal.Contract.GetAccepters(&_Proposal.CallOpts, uuid_)
}

// GetAccepters is a free data retrieval call binding the contract method 0x1d774ba8.
//
// Solidity: function getAccepters(string uuid_) view returns(address[])
func (_Proposal *ProposalCallerSession) GetAccepters(uuid_ string) ([]common.Address, error) {
	return _Proposal.Contract.GetAccepters(&_Proposal.CallOpts, uuid_)
}

// GetAcceptersV2 is a free data retrieval call binding the contract method 0x95f27b7b.
//
// Solidity: function getAcceptersV2(uint256 proposalId) view returns(address[])
func (_Proposal *ProposalCaller) GetAcceptersV2(opts *bind.CallOpts, proposalId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getAcceptersV2", proposalId)
	return *ret0, err
}

// GetAcceptersV2 is a free data retrieval call binding the contract method 0x95f27b7b.
//
// Solidity: function getAcceptersV2(uint256 proposalId) view returns(address[])
func (_Proposal *ProposalSession) GetAcceptersV2(proposalId *big.Int) ([]common.Address, error) {
	return _Proposal.Contract.GetAcceptersV2(&_Proposal.CallOpts, proposalId)
}

// GetAcceptersV2 is a free data retrieval call binding the contract method 0x95f27b7b.
//
// Solidity: function getAcceptersV2(uint256 proposalId) view returns(address[])
func (_Proposal *ProposalCallerSession) GetAcceptersV2(proposalId *big.Int) ([]common.Address, error) {
	return _Proposal.Contract.GetAcceptersV2(&_Proposal.CallOpts, proposalId)
}

// GetComment is a free data retrieval call binding the contract method 0xcbbaa589.
//
// Solidity: function getComment(string uuid_) view returns(string[])
func (_Proposal *ProposalCaller) GetComment(opts *bind.CallOpts, uuid_ string) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getComment", uuid_)
	return *ret0, err
}

// GetComment is a free data retrieval call binding the contract method 0xcbbaa589.
//
// Solidity: function getComment(string uuid_) view returns(string[])
func (_Proposal *ProposalSession) GetComment(uuid_ string) ([]string, error) {
	return _Proposal.Contract.GetComment(&_Proposal.CallOpts, uuid_)
}

// GetComment is a free data retrieval call binding the contract method 0xcbbaa589.
//
// Solidity: function getComment(string uuid_) view returns(string[])
func (_Proposal *ProposalCallerSession) GetComment(uuid_ string) ([]string, error) {
	return _Proposal.Contract.GetComment(&_Proposal.CallOpts, uuid_)
}

// GetCommentV2 is a free data retrieval call binding the contract method 0x0454d213.
//
// Solidity: function getCommentV2(uint256 proposalId) view returns(string[])
func (_Proposal *ProposalCaller) GetCommentV2(opts *bind.CallOpts, proposalId *big.Int) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getCommentV2", proposalId)
	return *ret0, err
}

// GetCommentV2 is a free data retrieval call binding the contract method 0x0454d213.
//
// Solidity: function getCommentV2(uint256 proposalId) view returns(string[])
func (_Proposal *ProposalSession) GetCommentV2(proposalId *big.Int) ([]string, error) {
	return _Proposal.Contract.GetCommentV2(&_Proposal.CallOpts, proposalId)
}

// GetCommentV2 is a free data retrieval call binding the contract method 0x0454d213.
//
// Solidity: function getCommentV2(uint256 proposalId) view returns(string[])
func (_Proposal *ProposalCallerSession) GetCommentV2(proposalId *big.Int) ([]string, error) {
	return _Proposal.Contract.GetCommentV2(&_Proposal.CallOpts, proposalId)
}

// GetCvp is a free data retrieval call binding the contract method 0x96f94f64.
//
// Solidity: function getCvp(string uuid, address voter) view returns(uint256)
func (_Proposal *ProposalCaller) GetCvp(opts *bind.CallOpts, uuid string, voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getCvp", uuid, voter)
	return *ret0, err
}

// GetCvp is a free data retrieval call binding the contract method 0x96f94f64.
//
// Solidity: function getCvp(string uuid, address voter) view returns(uint256)
func (_Proposal *ProposalSession) GetCvp(uuid string, voter common.Address) (*big.Int, error) {
	return _Proposal.Contract.GetCvp(&_Proposal.CallOpts, uuid, voter)
}

// GetCvp is a free data retrieval call binding the contract method 0x96f94f64.
//
// Solidity: function getCvp(string uuid, address voter) view returns(uint256)
func (_Proposal *ProposalCallerSession) GetCvp(uuid string, voter common.Address) (*big.Int, error) {
	return _Proposal.Contract.GetCvp(&_Proposal.CallOpts, uuid, voter)
}

// GetCvpV2 is a free data retrieval call binding the contract method 0x0977710a.
//
// Solidity: function getCvpV2(uint256 proposalId, address voter) view returns(uint256)
func (_Proposal *ProposalCaller) GetCvpV2(opts *bind.CallOpts, proposalId *big.Int, voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getCvpV2", proposalId, voter)
	return *ret0, err
}

// GetCvpV2 is a free data retrieval call binding the contract method 0x0977710a.
//
// Solidity: function getCvpV2(uint256 proposalId, address voter) view returns(uint256)
func (_Proposal *ProposalSession) GetCvpV2(proposalId *big.Int, voter common.Address) (*big.Int, error) {
	return _Proposal.Contract.GetCvpV2(&_Proposal.CallOpts, proposalId, voter)
}

// GetCvpV2 is a free data retrieval call binding the contract method 0x0977710a.
//
// Solidity: function getCvpV2(uint256 proposalId, address voter) view returns(uint256)
func (_Proposal *ProposalCallerSession) GetCvpV2(proposalId *big.Int, voter common.Address) (*big.Int, error) {
	return _Proposal.Contract.GetCvpV2(&_Proposal.CallOpts, proposalId, voter)
}

// GetParticipants is a free data retrieval call binding the contract method 0x7e90def7.
//
// Solidity: function getParticipants(string uuid_) view returns(address[])
func (_Proposal *ProposalCaller) GetParticipants(opts *bind.CallOpts, uuid_ string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getParticipants", uuid_)
	return *ret0, err
}

// GetParticipants is a free data retrieval call binding the contract method 0x7e90def7.
//
// Solidity: function getParticipants(string uuid_) view returns(address[])
func (_Proposal *ProposalSession) GetParticipants(uuid_ string) ([]common.Address, error) {
	return _Proposal.Contract.GetParticipants(&_Proposal.CallOpts, uuid_)
}

// GetParticipants is a free data retrieval call binding the contract method 0x7e90def7.
//
// Solidity: function getParticipants(string uuid_) view returns(address[])
func (_Proposal *ProposalCallerSession) GetParticipants(uuid_ string) ([]common.Address, error) {
	return _Proposal.Contract.GetParticipants(&_Proposal.CallOpts, uuid_)
}

// GetParticipantsV2 is a free data retrieval call binding the contract method 0x71166f71.
//
// Solidity: function getParticipantsV2(uint256 proposalId) view returns(address[])
func (_Proposal *ProposalCaller) GetParticipantsV2(opts *bind.CallOpts, proposalId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getParticipantsV2", proposalId)
	return *ret0, err
}

// GetParticipantsV2 is a free data retrieval call binding the contract method 0x71166f71.
//
// Solidity: function getParticipantsV2(uint256 proposalId) view returns(address[])
func (_Proposal *ProposalSession) GetParticipantsV2(proposalId *big.Int) ([]common.Address, error) {
	return _Proposal.Contract.GetParticipantsV2(&_Proposal.CallOpts, proposalId)
}

// GetParticipantsV2 is a free data retrieval call binding the contract method 0x71166f71.
//
// Solidity: function getParticipantsV2(uint256 proposalId) view returns(address[])
func (_Proposal *ProposalCallerSession) GetParticipantsV2(proposalId *big.Int) ([]common.Address, error) {
	return _Proposal.Contract.GetParticipantsV2(&_Proposal.CallOpts, proposalId)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns(string[])
func (_Proposal *ProposalCaller) GetPolicies(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getPolicies")
	return *ret0, err
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns(string[])
func (_Proposal *ProposalSession) GetPolicies() ([]string, error) {
	return _Proposal.Contract.GetPolicies(&_Proposal.CallOpts)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns(string[])
func (_Proposal *ProposalCallerSession) GetPolicies() ([]string, error) {
	return _Proposal.Contract.GetPolicies(&_Proposal.CallOpts)
}

// GetPolicy is a free data retrieval call binding the contract method 0x60dd5f90.
//
// Solidity: function getPolicy(string policyName_) view returns(uint256)
func (_Proposal *ProposalCaller) GetPolicy(opts *bind.CallOpts, policyName_ string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getPolicy", policyName_)
	return *ret0, err
}

// GetPolicy is a free data retrieval call binding the contract method 0x60dd5f90.
//
// Solidity: function getPolicy(string policyName_) view returns(uint256)
func (_Proposal *ProposalSession) GetPolicy(policyName_ string) (*big.Int, error) {
	return _Proposal.Contract.GetPolicy(&_Proposal.CallOpts, policyName_)
}

// GetPolicy is a free data retrieval call binding the contract method 0x60dd5f90.
//
// Solidity: function getPolicy(string policyName_) view returns(uint256)
func (_Proposal *ProposalCallerSession) GetPolicy(policyName_ string) (*big.Int, error) {
	return _Proposal.Contract.GetPolicy(&_Proposal.CallOpts, policyName_)
}

// GetPolicyValue is a free data retrieval call binding the contract method 0x5b55320c.
//
// Solidity: function getPolicyValue(string key) view returns(uint256)
func (_Proposal *ProposalCaller) GetPolicyValue(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getPolicyValue", key)
	return *ret0, err
}

// GetPolicyValue is a free data retrieval call binding the contract method 0x5b55320c.
//
// Solidity: function getPolicyValue(string key) view returns(uint256)
func (_Proposal *ProposalSession) GetPolicyValue(key string) (*big.Int, error) {
	return _Proposal.Contract.GetPolicyValue(&_Proposal.CallOpts, key)
}

// GetPolicyValue is a free data retrieval call binding the contract method 0x5b55320c.
//
// Solidity: function getPolicyValue(string key) view returns(uint256)
func (_Proposal *ProposalCallerSession) GetPolicyValue(key string) (*big.Int, error) {
	return _Proposal.Contract.GetPolicyValue(&_Proposal.CallOpts, key)
}

// GetProposal is a free data retrieval call binding the contract method 0x25d3a09d.
//
// Solidity: function getProposal(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalCaller) GetProposal(opts *bind.CallOpts, uuid string) (ProposalParameter, error) {
	var (
		ret0 = new(ProposalParameter)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getProposal", uuid)
	return *ret0, err
}

// GetProposal is a free data retrieval call binding the contract method 0x25d3a09d.
//
// Solidity: function getProposal(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalSession) GetProposal(uuid string) (ProposalParameter, error) {
	return _Proposal.Contract.GetProposal(&_Proposal.CallOpts, uuid)
}

// GetProposal is a free data retrieval call binding the contract method 0x25d3a09d.
//
// Solidity: function getProposal(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalCallerSession) GetProposal(uuid string) (ProposalParameter, error) {
	return _Proposal.Contract.GetProposal(&_Proposal.CallOpts, uuid)
}

// GetProposalParticipants is a free data retrieval call binding the contract method 0x3a024e2b.
//
// Solidity: function getProposalParticipants(string policyName_, string uuid_) view returns(address[])
func (_Proposal *ProposalCaller) GetProposalParticipants(opts *bind.CallOpts, policyName_ string, uuid_ string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getProposalParticipants", policyName_, uuid_)
	return *ret0, err
}

// GetProposalParticipants is a free data retrieval call binding the contract method 0x3a024e2b.
//
// Solidity: function getProposalParticipants(string policyName_, string uuid_) view returns(address[])
func (_Proposal *ProposalSession) GetProposalParticipants(policyName_ string, uuid_ string) ([]common.Address, error) {
	return _Proposal.Contract.GetProposalParticipants(&_Proposal.CallOpts, policyName_, uuid_)
}

// GetProposalParticipants is a free data retrieval call binding the contract method 0x3a024e2b.
//
// Solidity: function getProposalParticipants(string policyName_, string uuid_) view returns(address[])
func (_Proposal *ProposalCallerSession) GetProposalParticipants(policyName_ string, uuid_ string) ([]common.Address, error) {
	return _Proposal.Contract.GetProposalParticipants(&_Proposal.CallOpts, policyName_, uuid_)
}

// GetProposalParticipantsV2 is a free data retrieval call binding the contract method 0x7480e3fc.
//
// Solidity: function getProposalParticipantsV2(string policyName_, uint256 proposalId) view returns(address[])
func (_Proposal *ProposalCaller) GetProposalParticipantsV2(opts *bind.CallOpts, policyName_ string, proposalId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getProposalParticipantsV2", policyName_, proposalId)
	return *ret0, err
}

// GetProposalParticipantsV2 is a free data retrieval call binding the contract method 0x7480e3fc.
//
// Solidity: function getProposalParticipantsV2(string policyName_, uint256 proposalId) view returns(address[])
func (_Proposal *ProposalSession) GetProposalParticipantsV2(policyName_ string, proposalId *big.Int) ([]common.Address, error) {
	return _Proposal.Contract.GetProposalParticipantsV2(&_Proposal.CallOpts, policyName_, proposalId)
}

// GetProposalParticipantsV2 is a free data retrieval call binding the contract method 0x7480e3fc.
//
// Solidity: function getProposalParticipantsV2(string policyName_, uint256 proposalId) view returns(address[])
func (_Proposal *ProposalCallerSession) GetProposalParticipantsV2(policyName_ string, proposalId *big.Int) ([]common.Address, error) {
	return _Proposal.Contract.GetProposalParticipantsV2(&_Proposal.CallOpts, policyName_, proposalId)
}

// GetProposalStructure is a free data retrieval call binding the contract method 0xfef2293f.
//
// Solidity: function getProposalStructure(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalCaller) GetProposalStructure(opts *bind.CallOpts, uuid string) (ProposalParameter, error) {
	var (
		ret0 = new(ProposalParameter)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getProposalStructure", uuid)
	return *ret0, err
}

// GetProposalStructure is a free data retrieval call binding the contract method 0xfef2293f.
//
// Solidity: function getProposalStructure(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalSession) GetProposalStructure(uuid string) (ProposalParameter, error) {
	return _Proposal.Contract.GetProposalStructure(&_Proposal.CallOpts, uuid)
}

// GetProposalStructure is a free data retrieval call binding the contract method 0xfef2293f.
//
// Solidity: function getProposalStructure(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalCallerSession) GetProposalStructure(uuid string) (ProposalParameter, error) {
	return _Proposal.Contract.GetProposalStructure(&_Proposal.CallOpts, uuid)
}

// GetProposalStructureV2 is a free data retrieval call binding the contract method 0x3cd7c90b.
//
// Solidity: function getProposalStructureV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalCaller) GetProposalStructureV2(opts *bind.CallOpts, proposalId *big.Int) (ProposalParameterV2, error) {
	var (
		ret0 = new(ProposalParameterV2)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getProposalStructureV2", proposalId)
	return *ret0, err
}

// GetProposalStructureV2 is a free data retrieval call binding the contract method 0x3cd7c90b.
//
// Solidity: function getProposalStructureV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalSession) GetProposalStructureV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _Proposal.Contract.GetProposalStructureV2(&_Proposal.CallOpts, proposalId)
}

// GetProposalStructureV2 is a free data retrieval call binding the contract method 0x3cd7c90b.
//
// Solidity: function getProposalStructureV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalCallerSession) GetProposalStructureV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _Proposal.Contract.GetProposalStructureV2(&_Proposal.CallOpts, proposalId)
}

// GetProposalV2 is a free data retrieval call binding the contract method 0xc6d61e6b.
//
// Solidity: function getProposalV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalCaller) GetProposalV2(opts *bind.CallOpts, proposalId *big.Int) (ProposalParameterV2, error) {
	var (
		ret0 = new(ProposalParameterV2)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getProposalV2", proposalId)
	return *ret0, err
}

// GetProposalV2 is a free data retrieval call binding the contract method 0xc6d61e6b.
//
// Solidity: function getProposalV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalSession) GetProposalV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _Proposal.Contract.GetProposalV2(&_Proposal.CallOpts, proposalId)
}

// GetProposalV2 is a free data retrieval call binding the contract method 0xc6d61e6b.
//
// Solidity: function getProposalV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_Proposal *ProposalCallerSession) GetProposalV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _Proposal.Contract.GetProposalV2(&_Proposal.CallOpts, proposalId)
}

// GetProposedPolicy is a free data retrieval call binding the contract method 0x44ecebb1.
//
// Solidity: function getProposedPolicy(string uuid_) view returns((string,uint256))
func (_Proposal *ProposalCaller) GetProposedPolicy(opts *bind.CallOpts, uuid_ string) (ProposedPoicy, error) {
	var (
		ret0 = new(ProposedPoicy)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getProposedPolicy", uuid_)
	return *ret0, err
}

// GetProposedPolicy is a free data retrieval call binding the contract method 0x44ecebb1.
//
// Solidity: function getProposedPolicy(string uuid_) view returns((string,uint256))
func (_Proposal *ProposalSession) GetProposedPolicy(uuid_ string) (ProposedPoicy, error) {
	return _Proposal.Contract.GetProposedPolicy(&_Proposal.CallOpts, uuid_)
}

// GetProposedPolicy is a free data retrieval call binding the contract method 0x44ecebb1.
//
// Solidity: function getProposedPolicy(string uuid_) view returns((string,uint256))
func (_Proposal *ProposalCallerSession) GetProposedPolicy(uuid_ string) (ProposedPoicy, error) {
	return _Proposal.Contract.GetProposedPolicy(&_Proposal.CallOpts, uuid_)
}

// GetProposedPolicyV2 is a free data retrieval call binding the contract method 0x03ec9978.
//
// Solidity: function getProposedPolicyV2(uint256 proposalId) view returns((string,uint256))
func (_Proposal *ProposalCaller) GetProposedPolicyV2(opts *bind.CallOpts, proposalId *big.Int) (ProposedPoicy, error) {
	var (
		ret0 = new(ProposedPoicy)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getProposedPolicyV2", proposalId)
	return *ret0, err
}

// GetProposedPolicyV2 is a free data retrieval call binding the contract method 0x03ec9978.
//
// Solidity: function getProposedPolicyV2(uint256 proposalId) view returns((string,uint256))
func (_Proposal *ProposalSession) GetProposedPolicyV2(proposalId *big.Int) (ProposedPoicy, error) {
	return _Proposal.Contract.GetProposedPolicyV2(&_Proposal.CallOpts, proposalId)
}

// GetProposedPolicyV2 is a free data retrieval call binding the contract method 0x03ec9978.
//
// Solidity: function getProposedPolicyV2(uint256 proposalId) view returns((string,uint256))
func (_Proposal *ProposalCallerSession) GetProposedPolicyV2(proposalId *big.Int) (ProposedPoicy, error) {
	return _Proposal.Contract.GetProposedPolicyV2(&_Proposal.CallOpts, proposalId)
}

// GetRejecters is a free data retrieval call binding the contract method 0x2149e6c6.
//
// Solidity: function getRejecters(string uuid_) view returns(address[])
func (_Proposal *ProposalCaller) GetRejecters(opts *bind.CallOpts, uuid_ string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getRejecters", uuid_)
	return *ret0, err
}

// GetRejecters is a free data retrieval call binding the contract method 0x2149e6c6.
//
// Solidity: function getRejecters(string uuid_) view returns(address[])
func (_Proposal *ProposalSession) GetRejecters(uuid_ string) ([]common.Address, error) {
	return _Proposal.Contract.GetRejecters(&_Proposal.CallOpts, uuid_)
}

// GetRejecters is a free data retrieval call binding the contract method 0x2149e6c6.
//
// Solidity: function getRejecters(string uuid_) view returns(address[])
func (_Proposal *ProposalCallerSession) GetRejecters(uuid_ string) ([]common.Address, error) {
	return _Proposal.Contract.GetRejecters(&_Proposal.CallOpts, uuid_)
}

// GetRejectersV2 is a free data retrieval call binding the contract method 0x2350bdc7.
//
// Solidity: function getRejectersV2(uint256 proposalId) view returns(address[])
func (_Proposal *ProposalCaller) GetRejectersV2(opts *bind.CallOpts, proposalId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getRejectersV2", proposalId)
	return *ret0, err
}

// GetRejectersV2 is a free data retrieval call binding the contract method 0x2350bdc7.
//
// Solidity: function getRejectersV2(uint256 proposalId) view returns(address[])
func (_Proposal *ProposalSession) GetRejectersV2(proposalId *big.Int) ([]common.Address, error) {
	return _Proposal.Contract.GetRejectersV2(&_Proposal.CallOpts, proposalId)
}

// GetRejectersV2 is a free data retrieval call binding the contract method 0x2350bdc7.
//
// Solidity: function getRejectersV2(uint256 proposalId) view returns(address[])
func (_Proposal *ProposalCallerSession) GetRejectersV2(proposalId *big.Int) ([]common.Address, error) {
	return _Proposal.Contract.GetRejectersV2(&_Proposal.CallOpts, proposalId)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_Proposal *ProposalCaller) GetStateAddr(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getStateAddr")
	return *ret0, err
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_Proposal *ProposalSession) GetStateAddr() ([]common.Address, error) {
	return _Proposal.Contract.GetStateAddr(&_Proposal.CallOpts)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_Proposal *ProposalCallerSession) GetStateAddr() ([]common.Address, error) {
	return _Proposal.Contract.GetStateAddr(&_Proposal.CallOpts)
}

// GetSummary is a free data retrieval call binding the contract method 0x4051ddac.
//
// Solidity: function getSummary() view returns(string[])
func (_Proposal *ProposalCaller) GetSummary(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getSummary")
	return *ret0, err
}

// GetSummary is a free data retrieval call binding the contract method 0x4051ddac.
//
// Solidity: function getSummary() view returns(string[])
func (_Proposal *ProposalSession) GetSummary() ([]string, error) {
	return _Proposal.Contract.GetSummary(&_Proposal.CallOpts)
}

// GetSummary is a free data retrieval call binding the contract method 0x4051ddac.
//
// Solidity: function getSummary() view returns(string[])
func (_Proposal *ProposalCallerSession) GetSummary() ([]string, error) {
	return _Proposal.Contract.GetSummary(&_Proposal.CallOpts)
}

// GetUuidAddress is a free data retrieval call binding the contract method 0xfeeff48a.
//
// Solidity: function getUuidAddress(string uuid) view returns(address)
func (_Proposal *ProposalCaller) GetUuidAddress(opts *bind.CallOpts, uuid string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getUuidAddress", uuid)
	return *ret0, err
}

// GetUuidAddress is a free data retrieval call binding the contract method 0xfeeff48a.
//
// Solidity: function getUuidAddress(string uuid) view returns(address)
func (_Proposal *ProposalSession) GetUuidAddress(uuid string) (common.Address, error) {
	return _Proposal.Contract.GetUuidAddress(&_Proposal.CallOpts, uuid)
}

// GetUuidAddress is a free data retrieval call binding the contract method 0xfeeff48a.
//
// Solidity: function getUuidAddress(string uuid) view returns(address)
func (_Proposal *ProposalCallerSession) GetUuidAddress(uuid string) (common.Address, error) {
	return _Proposal.Contract.GetUuidAddress(&_Proposal.CallOpts, uuid)
}

// GetUuidAddressV2 is a free data retrieval call binding the contract method 0x203bc900.
//
// Solidity: function getUuidAddressV2(uint256 proposalId) view returns(address)
func (_Proposal *ProposalCaller) GetUuidAddressV2(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getUuidAddressV2", proposalId)
	return *ret0, err
}

// GetUuidAddressV2 is a free data retrieval call binding the contract method 0x203bc900.
//
// Solidity: function getUuidAddressV2(uint256 proposalId) view returns(address)
func (_Proposal *ProposalSession) GetUuidAddressV2(proposalId *big.Int) (common.Address, error) {
	return _Proposal.Contract.GetUuidAddressV2(&_Proposal.CallOpts, proposalId)
}

// GetUuidAddressV2 is a free data retrieval call binding the contract method 0x203bc900.
//
// Solidity: function getUuidAddressV2(uint256 proposalId) view returns(address)
func (_Proposal *ProposalCallerSession) GetUuidAddressV2(proposalId *big.Int) (common.Address, error) {
	return _Proposal.Contract.GetUuidAddressV2(&_Proposal.CallOpts, proposalId)
}

// GetVoterResult is a free data retrieval call binding the contract method 0x8c4f62c1.
//
// Solidity: function getVoterResult(uint256 proposalId, address voter) view returns(bool)
func (_Proposal *ProposalCaller) GetVoterResult(opts *bind.CallOpts, proposalId *big.Int, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getVoterResult", proposalId, voter)
	return *ret0, err
}

// GetVoterResult is a free data retrieval call binding the contract method 0x8c4f62c1.
//
// Solidity: function getVoterResult(uint256 proposalId, address voter) view returns(bool)
func (_Proposal *ProposalSession) GetVoterResult(proposalId *big.Int, voter common.Address) (bool, error) {
	return _Proposal.Contract.GetVoterResult(&_Proposal.CallOpts, proposalId, voter)
}

// GetVoterResult is a free data retrieval call binding the contract method 0x8c4f62c1.
//
// Solidity: function getVoterResult(uint256 proposalId, address voter) view returns(bool)
func (_Proposal *ProposalCallerSession) GetVoterResult(proposalId *big.Int, voter common.Address) (bool, error) {
	return _Proposal.Contract.GetVoterResult(&_Proposal.CallOpts, proposalId, voter)
}

// GetVoterValue is a free data retrieval call binding the contract method 0x3f0c0a67.
//
// Solidity: function getVoterValue(address sender, string uuid_) view returns(uint256)
func (_Proposal *ProposalCaller) GetVoterValue(opts *bind.CallOpts, sender common.Address, uuid_ string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getVoterValue", sender, uuid_)
	return *ret0, err
}

// GetVoterValue is a free data retrieval call binding the contract method 0x3f0c0a67.
//
// Solidity: function getVoterValue(address sender, string uuid_) view returns(uint256)
func (_Proposal *ProposalSession) GetVoterValue(sender common.Address, uuid_ string) (*big.Int, error) {
	return _Proposal.Contract.GetVoterValue(&_Proposal.CallOpts, sender, uuid_)
}

// GetVoterValue is a free data retrieval call binding the contract method 0x3f0c0a67.
//
// Solidity: function getVoterValue(address sender, string uuid_) view returns(uint256)
func (_Proposal *ProposalCallerSession) GetVoterValue(sender common.Address, uuid_ string) (*big.Int, error) {
	return _Proposal.Contract.GetVoterValue(&_Proposal.CallOpts, sender, uuid_)
}

// GetVoterValueV2 is a free data retrieval call binding the contract method 0x4c558a6a.
//
// Solidity: function getVoterValueV2(address sender, uint256 proposalId) view returns(uint256)
func (_Proposal *ProposalCaller) GetVoterValueV2(opts *bind.CallOpts, sender common.Address, proposalId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "getVoterValueV2", sender, proposalId)
	return *ret0, err
}

// GetVoterValueV2 is a free data retrieval call binding the contract method 0x4c558a6a.
//
// Solidity: function getVoterValueV2(address sender, uint256 proposalId) view returns(uint256)
func (_Proposal *ProposalSession) GetVoterValueV2(sender common.Address, proposalId *big.Int) (*big.Int, error) {
	return _Proposal.Contract.GetVoterValueV2(&_Proposal.CallOpts, sender, proposalId)
}

// GetVoterValueV2 is a free data retrieval call binding the contract method 0x4c558a6a.
//
// Solidity: function getVoterValueV2(address sender, uint256 proposalId) view returns(uint256)
func (_Proposal *ProposalCallerSession) GetVoterValueV2(sender common.Address, proposalId *big.Int) (*big.Int, error) {
	return _Proposal.Contract.GetVoterValueV2(&_Proposal.CallOpts, sender, proposalId)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Proposal *ProposalCaller) NameExt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "nameExt")
	return *ret0, err
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Proposal *ProposalSession) NameExt() (string, error) {
	return _Proposal.Contract.NameExt(&_Proposal.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_Proposal *ProposalCallerSession) NameExt() (string, error) {
	return _Proposal.Contract.NameExt(&_Proposal.CallOpts)
}

// ParticipantsLength is a free data retrieval call binding the contract method 0x01f0ed1b.
//
// Solidity: function participantsLength(string uuid) view returns(uint256)
func (_Proposal *ProposalCaller) ParticipantsLength(opts *bind.CallOpts, uuid string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "participantsLength", uuid)
	return *ret0, err
}

// ParticipantsLength is a free data retrieval call binding the contract method 0x01f0ed1b.
//
// Solidity: function participantsLength(string uuid) view returns(uint256)
func (_Proposal *ProposalSession) ParticipantsLength(uuid string) (*big.Int, error) {
	return _Proposal.Contract.ParticipantsLength(&_Proposal.CallOpts, uuid)
}

// ParticipantsLength is a free data retrieval call binding the contract method 0x01f0ed1b.
//
// Solidity: function participantsLength(string uuid) view returns(uint256)
func (_Proposal *ProposalCallerSession) ParticipantsLength(uuid string) (*big.Int, error) {
	return _Proposal.Contract.ParticipantsLength(&_Proposal.CallOpts, uuid)
}

// ParticipantsLengthV2 is a free data retrieval call binding the contract method 0x3590a89d.
//
// Solidity: function participantsLengthV2(uint256 proposalId) view returns(uint256)
func (_Proposal *ProposalCaller) ParticipantsLengthV2(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "participantsLengthV2", proposalId)
	return *ret0, err
}

// ParticipantsLengthV2 is a free data retrieval call binding the contract method 0x3590a89d.
//
// Solidity: function participantsLengthV2(uint256 proposalId) view returns(uint256)
func (_Proposal *ProposalSession) ParticipantsLengthV2(proposalId *big.Int) (*big.Int, error) {
	return _Proposal.Contract.ParticipantsLengthV2(&_Proposal.CallOpts, proposalId)
}

// ParticipantsLengthV2 is a free data retrieval call binding the contract method 0x3590a89d.
//
// Solidity: function participantsLengthV2(uint256 proposalId) view returns(uint256)
func (_Proposal *ProposalCallerSession) ParticipantsLengthV2(proposalId *big.Int) (*big.Int, error) {
	return _Proposal.Contract.ParticipantsLengthV2(&_Proposal.CallOpts, proposalId)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Proposal *ProposalCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Proposal *ProposalSession) Ready() (bool, error) {
	return _Proposal.Contract.Ready(&_Proposal.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_Proposal *ProposalCallerSession) Ready() (bool, error) {
	return _Proposal.Contract.Ready(&_Proposal.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Proposal *ProposalCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Proposal.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Proposal *ProposalSession) Registry() (common.Address, error) {
	return _Proposal.Contract.Registry(&_Proposal.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Proposal *ProposalCallerSession) Registry() (common.Address, error) {
	return _Proposal.Contract.Registry(&_Proposal.CallOpts)
}

// AddAccepters is a paid mutator transaction binding the contract method 0x2f48fc0d.
//
// Solidity: function addAccepters(string uuid, address voter) returns()
func (_Proposal *ProposalTransactor) AddAccepters(opts *bind.TransactOpts, uuid string, voter common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "addAccepters", uuid, voter)
}

// AddAccepters is a paid mutator transaction binding the contract method 0x2f48fc0d.
//
// Solidity: function addAccepters(string uuid, address voter) returns()
func (_Proposal *ProposalSession) AddAccepters(uuid string, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddAccepters(&_Proposal.TransactOpts, uuid, voter)
}

// AddAccepters is a paid mutator transaction binding the contract method 0x2f48fc0d.
//
// Solidity: function addAccepters(string uuid, address voter) returns()
func (_Proposal *ProposalTransactorSession) AddAccepters(uuid string, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddAccepters(&_Proposal.TransactOpts, uuid, voter)
}

// AddAcceptersV2 is a paid mutator transaction binding the contract method 0xcd39c41a.
//
// Solidity: function addAcceptersV2(uint256 proposalId, address voter) returns()
func (_Proposal *ProposalTransactor) AddAcceptersV2(opts *bind.TransactOpts, proposalId *big.Int, voter common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "addAcceptersV2", proposalId, voter)
}

// AddAcceptersV2 is a paid mutator transaction binding the contract method 0xcd39c41a.
//
// Solidity: function addAcceptersV2(uint256 proposalId, address voter) returns()
func (_Proposal *ProposalSession) AddAcceptersV2(proposalId *big.Int, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddAcceptersV2(&_Proposal.TransactOpts, proposalId, voter)
}

// AddAcceptersV2 is a paid mutator transaction binding the contract method 0xcd39c41a.
//
// Solidity: function addAcceptersV2(uint256 proposalId, address voter) returns()
func (_Proposal *ProposalTransactorSession) AddAcceptersV2(proposalId *big.Int, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddAcceptersV2(&_Proposal.TransactOpts, proposalId, voter)
}

// AddParticipants is a paid mutator transaction binding the contract method 0xd4f757e7.
//
// Solidity: function addParticipants(string uuid, address voter) returns()
func (_Proposal *ProposalTransactor) AddParticipants(opts *bind.TransactOpts, uuid string, voter common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "addParticipants", uuid, voter)
}

// AddParticipants is a paid mutator transaction binding the contract method 0xd4f757e7.
//
// Solidity: function addParticipants(string uuid, address voter) returns()
func (_Proposal *ProposalSession) AddParticipants(uuid string, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddParticipants(&_Proposal.TransactOpts, uuid, voter)
}

// AddParticipants is a paid mutator transaction binding the contract method 0xd4f757e7.
//
// Solidity: function addParticipants(string uuid, address voter) returns()
func (_Proposal *ProposalTransactorSession) AddParticipants(uuid string, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddParticipants(&_Proposal.TransactOpts, uuid, voter)
}

// AddParticipantsV2 is a paid mutator transaction binding the contract method 0x4473b38b.
//
// Solidity: function addParticipantsV2(uint256 proposalId, address voter) returns()
func (_Proposal *ProposalTransactor) AddParticipantsV2(opts *bind.TransactOpts, proposalId *big.Int, voter common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "addParticipantsV2", proposalId, voter)
}

// AddParticipantsV2 is a paid mutator transaction binding the contract method 0x4473b38b.
//
// Solidity: function addParticipantsV2(uint256 proposalId, address voter) returns()
func (_Proposal *ProposalSession) AddParticipantsV2(proposalId *big.Int, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddParticipantsV2(&_Proposal.TransactOpts, proposalId, voter)
}

// AddParticipantsV2 is a paid mutator transaction binding the contract method 0x4473b38b.
//
// Solidity: function addParticipantsV2(uint256 proposalId, address voter) returns()
func (_Proposal *ProposalTransactorSession) AddParticipantsV2(proposalId *big.Int, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddParticipantsV2(&_Proposal.TransactOpts, proposalId, voter)
}

// AddRejecters is a paid mutator transaction binding the contract method 0xe87c0761.
//
// Solidity: function addRejecters(string uuid, address voter) returns()
func (_Proposal *ProposalTransactor) AddRejecters(opts *bind.TransactOpts, uuid string, voter common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "addRejecters", uuid, voter)
}

// AddRejecters is a paid mutator transaction binding the contract method 0xe87c0761.
//
// Solidity: function addRejecters(string uuid, address voter) returns()
func (_Proposal *ProposalSession) AddRejecters(uuid string, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddRejecters(&_Proposal.TransactOpts, uuid, voter)
}

// AddRejecters is a paid mutator transaction binding the contract method 0xe87c0761.
//
// Solidity: function addRejecters(string uuid, address voter) returns()
func (_Proposal *ProposalTransactorSession) AddRejecters(uuid string, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddRejecters(&_Proposal.TransactOpts, uuid, voter)
}

// AddRejectersV2 is a paid mutator transaction binding the contract method 0x2cda1088.
//
// Solidity: function addRejectersV2(uint256 proposalId, address voter) returns()
func (_Proposal *ProposalTransactor) AddRejectersV2(opts *bind.TransactOpts, proposalId *big.Int, voter common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "addRejectersV2", proposalId, voter)
}

// AddRejectersV2 is a paid mutator transaction binding the contract method 0x2cda1088.
//
// Solidity: function addRejectersV2(uint256 proposalId, address voter) returns()
func (_Proposal *ProposalSession) AddRejectersV2(proposalId *big.Int, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddRejectersV2(&_Proposal.TransactOpts, proposalId, voter)
}

// AddRejectersV2 is a paid mutator transaction binding the contract method 0x2cda1088.
//
// Solidity: function addRejectersV2(uint256 proposalId, address voter) returns()
func (_Proposal *ProposalTransactorSession) AddRejectersV2(proposalId *big.Int, voter common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.AddRejectersV2(&_Proposal.TransactOpts, proposalId, voter)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Proposal *ProposalTransactor) Initialize(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "initialize", addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Proposal *ProposalSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts, addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_Proposal *ProposalTransactorSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts, addrs)
}

// InsertPolicyData is a paid mutator transaction binding the contract method 0x6d41d8d2.
//
// Solidity: function insertPolicyData(string policyName, bool check, uint256 policyValue) returns()
func (_Proposal *ProposalTransactor) InsertPolicyData(opts *bind.TransactOpts, policyName string, check bool, policyValue *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "insertPolicyData", policyName, check, policyValue)
}

// InsertPolicyData is a paid mutator transaction binding the contract method 0x6d41d8d2.
//
// Solidity: function insertPolicyData(string policyName, bool check, uint256 policyValue) returns()
func (_Proposal *ProposalSession) InsertPolicyData(policyName string, check bool, policyValue *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.InsertPolicyData(&_Proposal.TransactOpts, policyName, check, policyValue)
}

// InsertPolicyData is a paid mutator transaction binding the contract method 0x6d41d8d2.
//
// Solidity: function insertPolicyData(string policyName, bool check, uint256 policyValue) returns()
func (_Proposal *ProposalTransactorSession) InsertPolicyData(policyName string, check bool, policyValue *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.InsertPolicyData(&_Proposal.TransactOpts, policyName, check, policyValue)
}

// InsertProposalData is a paid mutator transaction binding the contract method 0x318d9402.
//
// Solidity: function insertProposalData(string uuid, bool check, uint256 size) returns()
func (_Proposal *ProposalTransactor) InsertProposalData(opts *bind.TransactOpts, uuid string, check bool, size *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "insertProposalData", uuid, check, size)
}

// InsertProposalData is a paid mutator transaction binding the contract method 0x318d9402.
//
// Solidity: function insertProposalData(string uuid, bool check, uint256 size) returns()
func (_Proposal *ProposalSession) InsertProposalData(uuid string, check bool, size *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.InsertProposalData(&_Proposal.TransactOpts, uuid, check, size)
}

// InsertProposalData is a paid mutator transaction binding the contract method 0x318d9402.
//
// Solidity: function insertProposalData(string uuid, bool check, uint256 size) returns()
func (_Proposal *ProposalTransactorSession) InsertProposalData(uuid string, check bool, size *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.InsertProposalData(&_Proposal.TransactOpts, uuid, check, size)
}

// InsertProposalDataV2 is a paid mutator transaction binding the contract method 0x36ab565d.
//
// Solidity: function insertProposalDataV2(uint256 proposalId, bool check, uint256 size) returns()
func (_Proposal *ProposalTransactor) InsertProposalDataV2(opts *bind.TransactOpts, proposalId *big.Int, check bool, size *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "insertProposalDataV2", proposalId, check, size)
}

// InsertProposalDataV2 is a paid mutator transaction binding the contract method 0x36ab565d.
//
// Solidity: function insertProposalDataV2(uint256 proposalId, bool check, uint256 size) returns()
func (_Proposal *ProposalSession) InsertProposalDataV2(proposalId *big.Int, check bool, size *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.InsertProposalDataV2(&_Proposal.TransactOpts, proposalId, check, size)
}

// InsertProposalDataV2 is a paid mutator transaction binding the contract method 0x36ab565d.
//
// Solidity: function insertProposalDataV2(uint256 proposalId, bool check, uint256 size) returns()
func (_Proposal *ProposalTransactorSession) InsertProposalDataV2(proposalId *big.Int, check bool, size *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.InsertProposalDataV2(&_Proposal.TransactOpts, proposalId, check, size)
}

// SetComment is a paid mutator transaction binding the contract method 0x328e0b3f.
//
// Solidity: function setComment(string uuid_, string contents_) returns()
func (_Proposal *ProposalTransactor) SetComment(opts *bind.TransactOpts, uuid_ string, contents_ string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setComment", uuid_, contents_)
}

// SetComment is a paid mutator transaction binding the contract method 0x328e0b3f.
//
// Solidity: function setComment(string uuid_, string contents_) returns()
func (_Proposal *ProposalSession) SetComment(uuid_ string, contents_ string) (*types.Transaction, error) {
	return _Proposal.Contract.SetComment(&_Proposal.TransactOpts, uuid_, contents_)
}

// SetComment is a paid mutator transaction binding the contract method 0x328e0b3f.
//
// Solidity: function setComment(string uuid_, string contents_) returns()
func (_Proposal *ProposalTransactorSession) SetComment(uuid_ string, contents_ string) (*types.Transaction, error) {
	return _Proposal.Contract.SetComment(&_Proposal.TransactOpts, uuid_, contents_)
}

// SetCommentV2 is a paid mutator transaction binding the contract method 0xe14ffe6a.
//
// Solidity: function setCommentV2(uint256 proposalId, string contents_) returns()
func (_Proposal *ProposalTransactor) SetCommentV2(opts *bind.TransactOpts, proposalId *big.Int, contents_ string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setCommentV2", proposalId, contents_)
}

// SetCommentV2 is a paid mutator transaction binding the contract method 0xe14ffe6a.
//
// Solidity: function setCommentV2(uint256 proposalId, string contents_) returns()
func (_Proposal *ProposalSession) SetCommentV2(proposalId *big.Int, contents_ string) (*types.Transaction, error) {
	return _Proposal.Contract.SetCommentV2(&_Proposal.TransactOpts, proposalId, contents_)
}

// SetCommentV2 is a paid mutator transaction binding the contract method 0xe14ffe6a.
//
// Solidity: function setCommentV2(uint256 proposalId, string contents_) returns()
func (_Proposal *ProposalTransactorSession) SetCommentV2(proposalId *big.Int, contents_ string) (*types.Transaction, error) {
	return _Proposal.Contract.SetCommentV2(&_Proposal.TransactOpts, proposalId, contents_)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Proposal *ProposalTransactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setForwarder", forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Proposal *ProposalSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetForwarder(&_Proposal.TransactOpts, forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_Proposal *ProposalTransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetForwarder(&_Proposal.TransactOpts, forwarder)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x10130877.
//
// Solidity: function setPolicy(string policyName_, uint256 value_) returns()
func (_Proposal *ProposalTransactor) SetPolicy(opts *bind.TransactOpts, policyName_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setPolicy", policyName_, value_)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x10130877.
//
// Solidity: function setPolicy(string policyName_, uint256 value_) returns()
func (_Proposal *ProposalSession) SetPolicy(policyName_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetPolicy(&_Proposal.TransactOpts, policyName_, value_)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x10130877.
//
// Solidity: function setPolicy(string policyName_, uint256 value_) returns()
func (_Proposal *ProposalTransactorSession) SetPolicy(policyName_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetPolicy(&_Proposal.TransactOpts, policyName_, value_)
}

// SetPolicyData is a paid mutator transaction binding the contract method 0x42821064.
//
// Solidity: function setPolicyData(address sender, string policyName, uint256 value) returns()
func (_Proposal *ProposalTransactor) SetPolicyData(opts *bind.TransactOpts, sender common.Address, policyName string, value *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setPolicyData", sender, policyName, value)
}

// SetPolicyData is a paid mutator transaction binding the contract method 0x42821064.
//
// Solidity: function setPolicyData(address sender, string policyName, uint256 value) returns()
func (_Proposal *ProposalSession) SetPolicyData(sender common.Address, policyName string, value *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetPolicyData(&_Proposal.TransactOpts, sender, policyName, value)
}

// SetPolicyData is a paid mutator transaction binding the contract method 0x42821064.
//
// Solidity: function setPolicyData(address sender, string policyName, uint256 value) returns()
func (_Proposal *ProposalTransactorSession) SetPolicyData(sender common.Address, policyName string, value *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetPolicyData(&_Proposal.TransactOpts, sender, policyName, value)
}

// SetProposalParticipants is a paid mutator transaction binding the contract method 0xeb85c691.
//
// Solidity: function setProposalParticipants(string policyName_, string uuid_, address addr) returns()
func (_Proposal *ProposalTransactor) SetProposalParticipants(opts *bind.TransactOpts, policyName_ string, uuid_ string, addr common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setProposalParticipants", policyName_, uuid_, addr)
}

// SetProposalParticipants is a paid mutator transaction binding the contract method 0xeb85c691.
//
// Solidity: function setProposalParticipants(string policyName_, string uuid_, address addr) returns()
func (_Proposal *ProposalSession) SetProposalParticipants(policyName_ string, uuid_ string, addr common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetProposalParticipants(&_Proposal.TransactOpts, policyName_, uuid_, addr)
}

// SetProposalParticipants is a paid mutator transaction binding the contract method 0xeb85c691.
//
// Solidity: function setProposalParticipants(string policyName_, string uuid_, address addr) returns()
func (_Proposal *ProposalTransactorSession) SetProposalParticipants(policyName_ string, uuid_ string, addr common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetProposalParticipants(&_Proposal.TransactOpts, policyName_, uuid_, addr)
}

// SetProposalParticipantsV2 is a paid mutator transaction binding the contract method 0x49a91b6f.
//
// Solidity: function setProposalParticipantsV2(string policyName_, uint256 proposalId, address addr) returns()
func (_Proposal *ProposalTransactor) SetProposalParticipantsV2(opts *bind.TransactOpts, policyName_ string, proposalId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setProposalParticipantsV2", policyName_, proposalId, addr)
}

// SetProposalParticipantsV2 is a paid mutator transaction binding the contract method 0x49a91b6f.
//
// Solidity: function setProposalParticipantsV2(string policyName_, uint256 proposalId, address addr) returns()
func (_Proposal *ProposalSession) SetProposalParticipantsV2(policyName_ string, proposalId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetProposalParticipantsV2(&_Proposal.TransactOpts, policyName_, proposalId, addr)
}

// SetProposalParticipantsV2 is a paid mutator transaction binding the contract method 0x49a91b6f.
//
// Solidity: function setProposalParticipantsV2(string policyName_, uint256 proposalId, address addr) returns()
func (_Proposal *ProposalTransactorSession) SetProposalParticipantsV2(policyName_ string, proposalId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetProposalParticipantsV2(&_Proposal.TransactOpts, policyName_, proposalId, addr)
}

// SetSummary is a paid mutator transaction binding the contract method 0x89514f6e.
//
// Solidity: function setSummary(string summary) returns()
func (_Proposal *ProposalTransactor) SetSummary(opts *bind.TransactOpts, summary string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setSummary", summary)
}

// SetSummary is a paid mutator transaction binding the contract method 0x89514f6e.
//
// Solidity: function setSummary(string summary) returns()
func (_Proposal *ProposalSession) SetSummary(summary string) (*types.Transaction, error) {
	return _Proposal.Contract.SetSummary(&_Proposal.TransactOpts, summary)
}

// SetSummary is a paid mutator transaction binding the contract method 0x89514f6e.
//
// Solidity: function setSummary(string summary) returns()
func (_Proposal *ProposalTransactorSession) SetSummary(summary string) (*types.Transaction, error) {
	return _Proposal.Contract.SetSummary(&_Proposal.TransactOpts, summary)
}

// SetUuidAddress is a paid mutator transaction binding the contract method 0xf21d92a7.
//
// Solidity: function setUuidAddress(string uuid, address addr) returns()
func (_Proposal *ProposalTransactor) SetUuidAddress(opts *bind.TransactOpts, uuid string, addr common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setUuidAddress", uuid, addr)
}

// SetUuidAddress is a paid mutator transaction binding the contract method 0xf21d92a7.
//
// Solidity: function setUuidAddress(string uuid, address addr) returns()
func (_Proposal *ProposalSession) SetUuidAddress(uuid string, addr common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetUuidAddress(&_Proposal.TransactOpts, uuid, addr)
}

// SetUuidAddress is a paid mutator transaction binding the contract method 0xf21d92a7.
//
// Solidity: function setUuidAddress(string uuid, address addr) returns()
func (_Proposal *ProposalTransactorSession) SetUuidAddress(uuid string, addr common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetUuidAddress(&_Proposal.TransactOpts, uuid, addr)
}

// SetUuidAddressV2 is a paid mutator transaction binding the contract method 0x7950aeef.
//
// Solidity: function setUuidAddressV2(uint256 proposalId, address addr) returns()
func (_Proposal *ProposalTransactor) SetUuidAddressV2(opts *bind.TransactOpts, proposalId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setUuidAddressV2", proposalId, addr)
}

// SetUuidAddressV2 is a paid mutator transaction binding the contract method 0x7950aeef.
//
// Solidity: function setUuidAddressV2(uint256 proposalId, address addr) returns()
func (_Proposal *ProposalSession) SetUuidAddressV2(proposalId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetUuidAddressV2(&_Proposal.TransactOpts, proposalId, addr)
}

// SetUuidAddressV2 is a paid mutator transaction binding the contract method 0x7950aeef.
//
// Solidity: function setUuidAddressV2(uint256 proposalId, address addr) returns()
func (_Proposal *ProposalTransactorSession) SetUuidAddressV2(proposalId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetUuidAddressV2(&_Proposal.TransactOpts, proposalId, addr)
}

// SetVoterPower is a paid mutator transaction binding the contract method 0xb4db1e49.
//
// Solidity: function setVoterPower(string uuid, address voter, uint256 nvp) returns()
func (_Proposal *ProposalTransactor) SetVoterPower(opts *bind.TransactOpts, uuid string, voter common.Address, nvp *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setVoterPower", uuid, voter, nvp)
}

// SetVoterPower is a paid mutator transaction binding the contract method 0xb4db1e49.
//
// Solidity: function setVoterPower(string uuid, address voter, uint256 nvp) returns()
func (_Proposal *ProposalSession) SetVoterPower(uuid string, voter common.Address, nvp *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetVoterPower(&_Proposal.TransactOpts, uuid, voter, nvp)
}

// SetVoterPower is a paid mutator transaction binding the contract method 0xb4db1e49.
//
// Solidity: function setVoterPower(string uuid, address voter, uint256 nvp) returns()
func (_Proposal *ProposalTransactorSession) SetVoterPower(uuid string, voter common.Address, nvp *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetVoterPower(&_Proposal.TransactOpts, uuid, voter, nvp)
}

// SetVoterPowerV2 is a paid mutator transaction binding the contract method 0xffa5f2bb.
//
// Solidity: function setVoterPowerV2(uint256 proposalId, address voter, uint256 nvp) returns()
func (_Proposal *ProposalTransactor) SetVoterPowerV2(opts *bind.TransactOpts, proposalId *big.Int, voter common.Address, nvp *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setVoterPowerV2", proposalId, voter, nvp)
}

// SetVoterPowerV2 is a paid mutator transaction binding the contract method 0xffa5f2bb.
//
// Solidity: function setVoterPowerV2(uint256 proposalId, address voter, uint256 nvp) returns()
func (_Proposal *ProposalSession) SetVoterPowerV2(proposalId *big.Int, voter common.Address, nvp *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetVoterPowerV2(&_Proposal.TransactOpts, proposalId, voter, nvp)
}

// SetVoterPowerV2 is a paid mutator transaction binding the contract method 0xffa5f2bb.
//
// Solidity: function setVoterPowerV2(uint256 proposalId, address voter, uint256 nvp) returns()
func (_Proposal *ProposalTransactorSession) SetVoterPowerV2(proposalId *big.Int, voter common.Address, nvp *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetVoterPowerV2(&_Proposal.TransactOpts, proposalId, voter, nvp)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0xabee6533.
//
// Solidity: function updatePolicy(string policyName_, uint256 value_) returns()
func (_Proposal *ProposalTransactor) UpdatePolicy(opts *bind.TransactOpts, policyName_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "updatePolicy", policyName_, value_)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0xabee6533.
//
// Solidity: function updatePolicy(string policyName_, uint256 value_) returns()
func (_Proposal *ProposalSession) UpdatePolicy(policyName_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.UpdatePolicy(&_Proposal.TransactOpts, policyName_, value_)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0xabee6533.
//
// Solidity: function updatePolicy(string policyName_, uint256 value_) returns()
func (_Proposal *ProposalTransactorSession) UpdatePolicy(policyName_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.UpdatePolicy(&_Proposal.TransactOpts, policyName_, value_)
}

// UpdatePolicyData is a paid mutator transaction binding the contract method 0x3b02862f.
//
// Solidity: function updatePolicyData(string policyName, uint256 value) returns()
func (_Proposal *ProposalTransactor) UpdatePolicyData(opts *bind.TransactOpts, policyName string, value *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "updatePolicyData", policyName, value)
}

// UpdatePolicyData is a paid mutator transaction binding the contract method 0x3b02862f.
//
// Solidity: function updatePolicyData(string policyName, uint256 value) returns()
func (_Proposal *ProposalSession) UpdatePolicyData(policyName string, value *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.UpdatePolicyData(&_Proposal.TransactOpts, policyName, value)
}

// UpdatePolicyData is a paid mutator transaction binding the contract method 0x3b02862f.
//
// Solidity: function updatePolicyData(string policyName, uint256 value) returns()
func (_Proposal *ProposalTransactorSession) UpdatePolicyData(policyName string, value *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.UpdatePolicyData(&_Proposal.TransactOpts, policyName, value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Proposal *ProposalTransactor) Upgrade(opts *bind.TransactOpts, newVersion common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "upgrade", newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Proposal *ProposalSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Upgrade(&_Proposal.TransactOpts, newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_Proposal *ProposalTransactorSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Upgrade(&_Proposal.TransactOpts, newVersion)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Proposal *ProposalTransactor) UpgradeRegistry(opts *bind.TransactOpts, reg common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "upgradeRegistry", reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Proposal *ProposalSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.UpgradeRegistry(&_Proposal.TransactOpts, reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_Proposal *ProposalTransactorSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.UpgradeRegistry(&_Proposal.TransactOpts, reg)
}

// ProposalAddAcceptersV2EventIterator is returned from FilterAddAcceptersV2Event and is used to iterate over the raw logs and unpacked data for AddAcceptersV2Event events raised by the Proposal contract.
type ProposalAddAcceptersV2EventIterator struct {
	Event *ProposalAddAcceptersV2Event // Event containing the contract specifics and raw log

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
func (it *ProposalAddAcceptersV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalAddAcceptersV2Event)
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
		it.Event = new(ProposalAddAcceptersV2Event)
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
func (it *ProposalAddAcceptersV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalAddAcceptersV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalAddAcceptersV2Event represents a AddAcceptersV2Event event raised by the Proposal contract.
type ProposalAddAcceptersV2Event struct {
	ProposalId *big.Int
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddAcceptersV2Event is a free log retrieval operation binding the contract event 0x8acc10293c18e549c93faf91c78af2e2c34cca703f684a567fd7f71fd7bb9c08.
//
// Solidity: event addAcceptersV2Event(uint256 proposalId, address voter)
func (_Proposal *ProposalFilterer) FilterAddAcceptersV2Event(opts *bind.FilterOpts) (*ProposalAddAcceptersV2EventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "addAcceptersV2Event")
	if err != nil {
		return nil, err
	}
	return &ProposalAddAcceptersV2EventIterator{contract: _Proposal.contract, event: "addAcceptersV2Event", logs: logs, sub: sub}, nil
}

// WatchAddAcceptersV2Event is a free log subscription operation binding the contract event 0x8acc10293c18e549c93faf91c78af2e2c34cca703f684a567fd7f71fd7bb9c08.
//
// Solidity: event addAcceptersV2Event(uint256 proposalId, address voter)
func (_Proposal *ProposalFilterer) WatchAddAcceptersV2Event(opts *bind.WatchOpts, sink chan<- *ProposalAddAcceptersV2Event) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "addAcceptersV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalAddAcceptersV2Event)
				if err := _Proposal.contract.UnpackLog(event, "addAcceptersV2Event", log); err != nil {
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

// ParseAddAcceptersV2Event is a log parse operation binding the contract event 0x8acc10293c18e549c93faf91c78af2e2c34cca703f684a567fd7f71fd7bb9c08.
//
// Solidity: event addAcceptersV2Event(uint256 proposalId, address voter)
func (_Proposal *ProposalFilterer) ParseAddAcceptersV2Event(log types.Log) (*ProposalAddAcceptersV2Event, error) {
	event := new(ProposalAddAcceptersV2Event)
	if err := _Proposal.contract.UnpackLog(event, "addAcceptersV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalAddParticipantsV2EventIterator is returned from FilterAddParticipantsV2Event and is used to iterate over the raw logs and unpacked data for AddParticipantsV2Event events raised by the Proposal contract.
type ProposalAddParticipantsV2EventIterator struct {
	Event *ProposalAddParticipantsV2Event // Event containing the contract specifics and raw log

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
func (it *ProposalAddParticipantsV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalAddParticipantsV2Event)
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
		it.Event = new(ProposalAddParticipantsV2Event)
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
func (it *ProposalAddParticipantsV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalAddParticipantsV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalAddParticipantsV2Event represents a AddParticipantsV2Event event raised by the Proposal contract.
type ProposalAddParticipantsV2Event struct {
	ProposalId *big.Int
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddParticipantsV2Event is a free log retrieval operation binding the contract event 0x55356fcbcee543cf47110f771aac20077c3baacf70efe182790c4cb57ba76c92.
//
// Solidity: event addParticipantsV2Event(uint256 proposalId, address voter)
func (_Proposal *ProposalFilterer) FilterAddParticipantsV2Event(opts *bind.FilterOpts) (*ProposalAddParticipantsV2EventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "addParticipantsV2Event")
	if err != nil {
		return nil, err
	}
	return &ProposalAddParticipantsV2EventIterator{contract: _Proposal.contract, event: "addParticipantsV2Event", logs: logs, sub: sub}, nil
}

// WatchAddParticipantsV2Event is a free log subscription operation binding the contract event 0x55356fcbcee543cf47110f771aac20077c3baacf70efe182790c4cb57ba76c92.
//
// Solidity: event addParticipantsV2Event(uint256 proposalId, address voter)
func (_Proposal *ProposalFilterer) WatchAddParticipantsV2Event(opts *bind.WatchOpts, sink chan<- *ProposalAddParticipantsV2Event) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "addParticipantsV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalAddParticipantsV2Event)
				if err := _Proposal.contract.UnpackLog(event, "addParticipantsV2Event", log); err != nil {
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

// ParseAddParticipantsV2Event is a log parse operation binding the contract event 0x55356fcbcee543cf47110f771aac20077c3baacf70efe182790c4cb57ba76c92.
//
// Solidity: event addParticipantsV2Event(uint256 proposalId, address voter)
func (_Proposal *ProposalFilterer) ParseAddParticipantsV2Event(log types.Log) (*ProposalAddParticipantsV2Event, error) {
	event := new(ProposalAddParticipantsV2Event)
	if err := _Proposal.contract.UnpackLog(event, "addParticipantsV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalAddRejectersV2EventIterator is returned from FilterAddRejectersV2Event and is used to iterate over the raw logs and unpacked data for AddRejectersV2Event events raised by the Proposal contract.
type ProposalAddRejectersV2EventIterator struct {
	Event *ProposalAddRejectersV2Event // Event containing the contract specifics and raw log

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
func (it *ProposalAddRejectersV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalAddRejectersV2Event)
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
		it.Event = new(ProposalAddRejectersV2Event)
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
func (it *ProposalAddRejectersV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalAddRejectersV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalAddRejectersV2Event represents a AddRejectersV2Event event raised by the Proposal contract.
type ProposalAddRejectersV2Event struct {
	ProposalId *big.Int
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddRejectersV2Event is a free log retrieval operation binding the contract event 0xf7b123f72b03e8292817fa1447010662eaa785b030062c08879ca3eb21b37c6c.
//
// Solidity: event addRejectersV2Event(uint256 proposalId, address voter)
func (_Proposal *ProposalFilterer) FilterAddRejectersV2Event(opts *bind.FilterOpts) (*ProposalAddRejectersV2EventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "addRejectersV2Event")
	if err != nil {
		return nil, err
	}
	return &ProposalAddRejectersV2EventIterator{contract: _Proposal.contract, event: "addRejectersV2Event", logs: logs, sub: sub}, nil
}

// WatchAddRejectersV2Event is a free log subscription operation binding the contract event 0xf7b123f72b03e8292817fa1447010662eaa785b030062c08879ca3eb21b37c6c.
//
// Solidity: event addRejectersV2Event(uint256 proposalId, address voter)
func (_Proposal *ProposalFilterer) WatchAddRejectersV2Event(opts *bind.WatchOpts, sink chan<- *ProposalAddRejectersV2Event) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "addRejectersV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalAddRejectersV2Event)
				if err := _Proposal.contract.UnpackLog(event, "addRejectersV2Event", log); err != nil {
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

// ParseAddRejectersV2Event is a log parse operation binding the contract event 0xf7b123f72b03e8292817fa1447010662eaa785b030062c08879ca3eb21b37c6c.
//
// Solidity: event addRejectersV2Event(uint256 proposalId, address voter)
func (_Proposal *ProposalFilterer) ParseAddRejectersV2Event(log types.Log) (*ProposalAddRejectersV2Event, error) {
	event := new(ProposalAddRejectersV2Event)
	if err := _Proposal.contract.UnpackLog(event, "addRejectersV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalBeforeUpgradeHookEventIterator is returned from FilterBeforeUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for BeforeUpgradeHookEvent events raised by the Proposal contract.
type ProposalBeforeUpgradeHookEventIterator struct {
	Event *ProposalBeforeUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *ProposalBeforeUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalBeforeUpgradeHookEvent)
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
		it.Event = new(ProposalBeforeUpgradeHookEvent)
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
func (it *ProposalBeforeUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalBeforeUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalBeforeUpgradeHookEvent represents a BeforeUpgradeHookEvent event raised by the Proposal contract.
type ProposalBeforeUpgradeHookEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBeforeUpgradeHookEvent is a free log retrieval operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Proposal *ProposalFilterer) FilterBeforeUpgradeHookEvent(opts *bind.FilterOpts) (*ProposalBeforeUpgradeHookEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalBeforeUpgradeHookEventIterator{contract: _Proposal.contract, event: "beforeUpgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchBeforeUpgradeHookEvent is a free log subscription operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_Proposal *ProposalFilterer) WatchBeforeUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *ProposalBeforeUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalBeforeUpgradeHookEvent)
				if err := _Proposal.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseBeforeUpgradeHookEvent(log types.Log) (*ProposalBeforeUpgradeHookEvent, error) {
	event := new(ProposalBeforeUpgradeHookEvent)
	if err := _Proposal.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalChangeInitialDataEventIterator is returned from FilterChangeInitialDataEvent and is used to iterate over the raw logs and unpacked data for ChangeInitialDataEvent events raised by the Proposal contract.
type ProposalChangeInitialDataEventIterator struct {
	Event *ProposalChangeInitialDataEvent // Event containing the contract specifics and raw log

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
func (it *ProposalChangeInitialDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalChangeInitialDataEvent)
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
		it.Event = new(ProposalChangeInitialDataEvent)
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
func (it *ProposalChangeInitialDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalChangeInitialDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalChangeInitialDataEvent represents a ChangeInitialDataEvent event raised by the Proposal contract.
type ProposalChangeInitialDataEvent struct {
	Policies [7]string
	Values   [7]*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeInitialDataEvent is a free log retrieval operation binding the contract event 0x008e9f87e0ef0ca5d6d2712b11163a8fa0552135775129ef812f4a8787afbd9c.
//
// Solidity: event changeInitialDataEvent(string[7] policies, uint256[7] values)
func (_Proposal *ProposalFilterer) FilterChangeInitialDataEvent(opts *bind.FilterOpts) (*ProposalChangeInitialDataEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalChangeInitialDataEventIterator{contract: _Proposal.contract, event: "changeInitialDataEvent", logs: logs, sub: sub}, nil
}

// WatchChangeInitialDataEvent is a free log subscription operation binding the contract event 0x008e9f87e0ef0ca5d6d2712b11163a8fa0552135775129ef812f4a8787afbd9c.
//
// Solidity: event changeInitialDataEvent(string[7] policies, uint256[7] values)
func (_Proposal *ProposalFilterer) WatchChangeInitialDataEvent(opts *bind.WatchOpts, sink chan<- *ProposalChangeInitialDataEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalChangeInitialDataEvent)
				if err := _Proposal.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
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

// ParseChangeInitialDataEvent is a log parse operation binding the contract event 0x008e9f87e0ef0ca5d6d2712b11163a8fa0552135775129ef812f4a8787afbd9c.
//
// Solidity: event changeInitialDataEvent(string[7] policies, uint256[7] values)
func (_Proposal *ProposalFilterer) ParseChangeInitialDataEvent(log types.Log) (*ProposalChangeInitialDataEvent, error) {
	event := new(ProposalChangeInitialDataEvent)
	if err := _Proposal.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalEmitActivityEventIterator is returned from FilterEmitActivityEvent and is used to iterate over the raw logs and unpacked data for EmitActivityEvent events raised by the Proposal contract.
type ProposalEmitActivityEventIterator struct {
	Event *ProposalEmitActivityEvent // Event containing the contract specifics and raw log

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
func (it *ProposalEmitActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalEmitActivityEvent)
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
		it.Event = new(ProposalEmitActivityEvent)
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
func (it *ProposalEmitActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalEmitActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalEmitActivityEvent represents a EmitActivityEvent event raised by the Proposal contract.
type ProposalEmitActivityEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitActivityEvent is a free log retrieval operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Proposal *ProposalFilterer) FilterEmitActivityEvent(opts *bind.FilterOpts) (*ProposalEmitActivityEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalEmitActivityEventIterator{contract: _Proposal.contract, event: "emitActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitActivityEvent is a free log subscription operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_Proposal *ProposalFilterer) WatchEmitActivityEvent(opts *bind.WatchOpts, sink chan<- *ProposalEmitActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalEmitActivityEvent)
				if err := _Proposal.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseEmitActivityEvent(log types.Log) (*ProposalEmitActivityEvent, error) {
	event := new(ProposalEmitActivityEvent)
	if err := _Proposal.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalEmitAppDataEventIterator is returned from FilterEmitAppDataEvent and is used to iterate over the raw logs and unpacked data for EmitAppDataEvent events raised by the Proposal contract.
type ProposalEmitAppDataEventIterator struct {
	Event *ProposalEmitAppDataEvent // Event containing the contract specifics and raw log

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
func (it *ProposalEmitAppDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalEmitAppDataEvent)
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
		it.Event = new(ProposalEmitAppDataEvent)
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
func (it *ProposalEmitAppDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalEmitAppDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalEmitAppDataEvent represents a EmitAppDataEvent event raised by the Proposal contract.
type ProposalEmitAppDataEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitAppDataEvent is a free log retrieval operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Proposal *ProposalFilterer) FilterEmitAppDataEvent(opts *bind.FilterOpts) (*ProposalEmitAppDataEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalEmitAppDataEventIterator{contract: _Proposal.contract, event: "emitAppDataEvent", logs: logs, sub: sub}, nil
}

// WatchEmitAppDataEvent is a free log subscription operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_Proposal *ProposalFilterer) WatchEmitAppDataEvent(opts *bind.WatchOpts, sink chan<- *ProposalEmitAppDataEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalEmitAppDataEvent)
				if err := _Proposal.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseEmitAppDataEvent(log types.Log) (*ProposalEmitAppDataEvent, error) {
	event := new(ProposalEmitAppDataEvent)
	if err := _Proposal.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalEmitProposalActivityEventIterator is returned from FilterEmitProposalActivityEvent and is used to iterate over the raw logs and unpacked data for EmitProposalActivityEvent events raised by the Proposal contract.
type ProposalEmitProposalActivityEventIterator struct {
	Event *ProposalEmitProposalActivityEvent // Event containing the contract specifics and raw log

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
func (it *ProposalEmitProposalActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalEmitProposalActivityEvent)
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
		it.Event = new(ProposalEmitProposalActivityEvent)
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
func (it *ProposalEmitProposalActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalEmitProposalActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalEmitProposalActivityEvent represents a EmitProposalActivityEvent event raised by the Proposal contract.
type ProposalEmitProposalActivityEvent struct {
	ActivityType *big.Int
	Proposal     ProposalSummary
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitProposalActivityEvent is a free log retrieval operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Proposal *ProposalFilterer) FilterEmitProposalActivityEvent(opts *bind.FilterOpts) (*ProposalEmitProposalActivityEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalEmitProposalActivityEventIterator{contract: _Proposal.contract, event: "emitProposalActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitProposalActivityEvent is a free log subscription operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_Proposal *ProposalFilterer) WatchEmitProposalActivityEvent(opts *bind.WatchOpts, sink chan<- *ProposalEmitProposalActivityEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalEmitProposalActivityEvent)
				if err := _Proposal.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseEmitProposalActivityEvent(log types.Log) (*ProposalEmitProposalActivityEvent, error) {
	event := new(ProposalEmitProposalActivityEvent)
	if err := _Proposal.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the Proposal contract.
type ProposalInitializeEventIterator struct {
	Event *ProposalInitializeEvent // Event containing the contract specifics and raw log

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
func (it *ProposalInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalInitializeEvent)
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
		it.Event = new(ProposalInitializeEvent)
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
func (it *ProposalInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalInitializeEvent represents a InitializeEvent event raised by the Proposal contract.
type ProposalInitializeEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Proposal *ProposalFilterer) FilterInitializeEvent(opts *bind.FilterOpts) (*ProposalInitializeEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalInitializeEventIterator{contract: _Proposal.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_Proposal *ProposalFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *ProposalInitializeEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalInitializeEvent)
				if err := _Proposal.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseInitializeEvent(log types.Log) (*ProposalInitializeEvent, error) {
	event := new(ProposalInitializeEvent)
	if err := _Proposal.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalInitializeHookEventIterator is returned from FilterInitializeHookEvent and is used to iterate over the raw logs and unpacked data for InitializeHookEvent events raised by the Proposal contract.
type ProposalInitializeHookEventIterator struct {
	Event *ProposalInitializeHookEvent // Event containing the contract specifics and raw log

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
func (it *ProposalInitializeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalInitializeHookEvent)
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
		it.Event = new(ProposalInitializeHookEvent)
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
func (it *ProposalInitializeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalInitializeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalInitializeHookEvent represents a InitializeHookEvent event raised by the Proposal contract.
type ProposalInitializeHookEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeHookEvent is a free log retrieval operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Proposal *ProposalFilterer) FilterInitializeHookEvent(opts *bind.FilterOpts) (*ProposalInitializeHookEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalInitializeHookEventIterator{contract: _Proposal.contract, event: "initializeHookEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeHookEvent is a free log subscription operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_Proposal *ProposalFilterer) WatchInitializeHookEvent(opts *bind.WatchOpts, sink chan<- *ProposalInitializeHookEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalInitializeHookEvent)
				if err := _Proposal.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseInitializeHookEvent(log types.Log) (*ProposalInitializeHookEvent, error) {
	event := new(ProposalInitializeHookEvent)
	if err := _Proposal.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalInsertPolicyDataEventIterator is returned from FilterInsertPolicyDataEvent and is used to iterate over the raw logs and unpacked data for InsertPolicyDataEvent events raised by the Proposal contract.
type ProposalInsertPolicyDataEventIterator struct {
	Event *ProposalInsertPolicyDataEvent // Event containing the contract specifics and raw log

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
func (it *ProposalInsertPolicyDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalInsertPolicyDataEvent)
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
		it.Event = new(ProposalInsertPolicyDataEvent)
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
func (it *ProposalInsertPolicyDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalInsertPolicyDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalInsertPolicyDataEvent represents a InsertPolicyDataEvent event raised by the Proposal contract.
type ProposalInsertPolicyDataEvent struct {
	PolicyName  string
	Check       bool
	PolicyValue *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInsertPolicyDataEvent is a free log retrieval operation binding the contract event 0xbeadd3cf0e080e98396b359c37e6deba49234ccbb5f7bff3d9ea782cc1547665.
//
// Solidity: event insertPolicyDataEvent(string policyName, bool check, uint256 policyValue)
func (_Proposal *ProposalFilterer) FilterInsertPolicyDataEvent(opts *bind.FilterOpts) (*ProposalInsertPolicyDataEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "insertPolicyDataEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalInsertPolicyDataEventIterator{contract: _Proposal.contract, event: "insertPolicyDataEvent", logs: logs, sub: sub}, nil
}

// WatchInsertPolicyDataEvent is a free log subscription operation binding the contract event 0xbeadd3cf0e080e98396b359c37e6deba49234ccbb5f7bff3d9ea782cc1547665.
//
// Solidity: event insertPolicyDataEvent(string policyName, bool check, uint256 policyValue)
func (_Proposal *ProposalFilterer) WatchInsertPolicyDataEvent(opts *bind.WatchOpts, sink chan<- *ProposalInsertPolicyDataEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "insertPolicyDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalInsertPolicyDataEvent)
				if err := _Proposal.contract.UnpackLog(event, "insertPolicyDataEvent", log); err != nil {
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

// ParseInsertPolicyDataEvent is a log parse operation binding the contract event 0xbeadd3cf0e080e98396b359c37e6deba49234ccbb5f7bff3d9ea782cc1547665.
//
// Solidity: event insertPolicyDataEvent(string policyName, bool check, uint256 policyValue)
func (_Proposal *ProposalFilterer) ParseInsertPolicyDataEvent(log types.Log) (*ProposalInsertPolicyDataEvent, error) {
	event := new(ProposalInsertPolicyDataEvent)
	if err := _Proposal.contract.UnpackLog(event, "insertPolicyDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalInsertProposalDataV2EventIterator is returned from FilterInsertProposalDataV2Event and is used to iterate over the raw logs and unpacked data for InsertProposalDataV2Event events raised by the Proposal contract.
type ProposalInsertProposalDataV2EventIterator struct {
	Event *ProposalInsertProposalDataV2Event // Event containing the contract specifics and raw log

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
func (it *ProposalInsertProposalDataV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalInsertProposalDataV2Event)
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
		it.Event = new(ProposalInsertProposalDataV2Event)
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
func (it *ProposalInsertProposalDataV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalInsertProposalDataV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalInsertProposalDataV2Event represents a InsertProposalDataV2Event event raised by the Proposal contract.
type ProposalInsertProposalDataV2Event struct {
	ProposalId *big.Int
	Check      bool
	Size       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInsertProposalDataV2Event is a free log retrieval operation binding the contract event 0xf048d2f26e7b235094cecaba2912b28f0fbdf2d7c0bb25f6bd504ce94dc8bdbb.
//
// Solidity: event insertProposalDataV2Event(uint256 proposalId, bool check, uint256 size)
func (_Proposal *ProposalFilterer) FilterInsertProposalDataV2Event(opts *bind.FilterOpts) (*ProposalInsertProposalDataV2EventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "insertProposalDataV2Event")
	if err != nil {
		return nil, err
	}
	return &ProposalInsertProposalDataV2EventIterator{contract: _Proposal.contract, event: "insertProposalDataV2Event", logs: logs, sub: sub}, nil
}

// WatchInsertProposalDataV2Event is a free log subscription operation binding the contract event 0xf048d2f26e7b235094cecaba2912b28f0fbdf2d7c0bb25f6bd504ce94dc8bdbb.
//
// Solidity: event insertProposalDataV2Event(uint256 proposalId, bool check, uint256 size)
func (_Proposal *ProposalFilterer) WatchInsertProposalDataV2Event(opts *bind.WatchOpts, sink chan<- *ProposalInsertProposalDataV2Event) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "insertProposalDataV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalInsertProposalDataV2Event)
				if err := _Proposal.contract.UnpackLog(event, "insertProposalDataV2Event", log); err != nil {
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

// ParseInsertProposalDataV2Event is a log parse operation binding the contract event 0xf048d2f26e7b235094cecaba2912b28f0fbdf2d7c0bb25f6bd504ce94dc8bdbb.
//
// Solidity: event insertProposalDataV2Event(uint256 proposalId, bool check, uint256 size)
func (_Proposal *ProposalFilterer) ParseInsertProposalDataV2Event(log types.Log) (*ProposalInsertProposalDataV2Event, error) {
	event := new(ProposalInsertProposalDataV2Event)
	if err := _Proposal.contract.UnpackLog(event, "insertProposalDataV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalSetCommentV2EventIterator is returned from FilterSetCommentV2Event and is used to iterate over the raw logs and unpacked data for SetCommentV2Event events raised by the Proposal contract.
type ProposalSetCommentV2EventIterator struct {
	Event *ProposalSetCommentV2Event // Event containing the contract specifics and raw log

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
func (it *ProposalSetCommentV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalSetCommentV2Event)
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
		it.Event = new(ProposalSetCommentV2Event)
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
func (it *ProposalSetCommentV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalSetCommentV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalSetCommentV2Event represents a SetCommentV2Event event raised by the Proposal contract.
type ProposalSetCommentV2Event struct {
	ProposalId *big.Int
	Contents   string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetCommentV2Event is a free log retrieval operation binding the contract event 0x522c02a42eaa2e08cd549b81c0c94e432e8e0e7332d2852c3bf9793cdf7b3975.
//
// Solidity: event setCommentV2Event(uint256 proposalId, string contents_)
func (_Proposal *ProposalFilterer) FilterSetCommentV2Event(opts *bind.FilterOpts) (*ProposalSetCommentV2EventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "setCommentV2Event")
	if err != nil {
		return nil, err
	}
	return &ProposalSetCommentV2EventIterator{contract: _Proposal.contract, event: "setCommentV2Event", logs: logs, sub: sub}, nil
}

// WatchSetCommentV2Event is a free log subscription operation binding the contract event 0x522c02a42eaa2e08cd549b81c0c94e432e8e0e7332d2852c3bf9793cdf7b3975.
//
// Solidity: event setCommentV2Event(uint256 proposalId, string contents_)
func (_Proposal *ProposalFilterer) WatchSetCommentV2Event(opts *bind.WatchOpts, sink chan<- *ProposalSetCommentV2Event) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "setCommentV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalSetCommentV2Event)
				if err := _Proposal.contract.UnpackLog(event, "setCommentV2Event", log); err != nil {
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

// ParseSetCommentV2Event is a log parse operation binding the contract event 0x522c02a42eaa2e08cd549b81c0c94e432e8e0e7332d2852c3bf9793cdf7b3975.
//
// Solidity: event setCommentV2Event(uint256 proposalId, string contents_)
func (_Proposal *ProposalFilterer) ParseSetCommentV2Event(log types.Log) (*ProposalSetCommentV2Event, error) {
	event := new(ProposalSetCommentV2Event)
	if err := _Proposal.contract.UnpackLog(event, "setCommentV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalSetProposalParticipantsV2EventIterator is returned from FilterSetProposalParticipantsV2Event and is used to iterate over the raw logs and unpacked data for SetProposalParticipantsV2Event events raised by the Proposal contract.
type ProposalSetProposalParticipantsV2EventIterator struct {
	Event *ProposalSetProposalParticipantsV2Event // Event containing the contract specifics and raw log

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
func (it *ProposalSetProposalParticipantsV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalSetProposalParticipantsV2Event)
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
		it.Event = new(ProposalSetProposalParticipantsV2Event)
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
func (it *ProposalSetProposalParticipantsV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalSetProposalParticipantsV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalSetProposalParticipantsV2Event represents a SetProposalParticipantsV2Event event raised by the Proposal contract.
type ProposalSetProposalParticipantsV2Event struct {
	PolicyName string
	ProposalId *big.Int
	Addr       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetProposalParticipantsV2Event is a free log retrieval operation binding the contract event 0x335375dd4ba02441061a7f7bcb5c27850b58206d9dc390373a682ba7f2198298.
//
// Solidity: event setProposalParticipantsV2Event(string policyName_, uint256 proposalId, address addr)
func (_Proposal *ProposalFilterer) FilterSetProposalParticipantsV2Event(opts *bind.FilterOpts) (*ProposalSetProposalParticipantsV2EventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "setProposalParticipantsV2Event")
	if err != nil {
		return nil, err
	}
	return &ProposalSetProposalParticipantsV2EventIterator{contract: _Proposal.contract, event: "setProposalParticipantsV2Event", logs: logs, sub: sub}, nil
}

// WatchSetProposalParticipantsV2Event is a free log subscription operation binding the contract event 0x335375dd4ba02441061a7f7bcb5c27850b58206d9dc390373a682ba7f2198298.
//
// Solidity: event setProposalParticipantsV2Event(string policyName_, uint256 proposalId, address addr)
func (_Proposal *ProposalFilterer) WatchSetProposalParticipantsV2Event(opts *bind.WatchOpts, sink chan<- *ProposalSetProposalParticipantsV2Event) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "setProposalParticipantsV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalSetProposalParticipantsV2Event)
				if err := _Proposal.contract.UnpackLog(event, "setProposalParticipantsV2Event", log); err != nil {
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

// ParseSetProposalParticipantsV2Event is a log parse operation binding the contract event 0x335375dd4ba02441061a7f7bcb5c27850b58206d9dc390373a682ba7f2198298.
//
// Solidity: event setProposalParticipantsV2Event(string policyName_, uint256 proposalId, address addr)
func (_Proposal *ProposalFilterer) ParseSetProposalParticipantsV2Event(log types.Log) (*ProposalSetProposalParticipantsV2Event, error) {
	event := new(ProposalSetProposalParticipantsV2Event)
	if err := _Proposal.contract.UnpackLog(event, "setProposalParticipantsV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalSetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the Proposal contract.
type ProposalSetReadyEventIterator struct {
	Event *ProposalSetReadyEvent // Event containing the contract specifics and raw log

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
func (it *ProposalSetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalSetReadyEvent)
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
		it.Event = new(ProposalSetReadyEvent)
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
func (it *ProposalSetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalSetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalSetReadyEvent represents a SetReadyEvent event raised by the Proposal contract.
type ProposalSetReadyEvent struct {
	Done bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Proposal *ProposalFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*ProposalSetReadyEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalSetReadyEventIterator{contract: _Proposal.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_Proposal *ProposalFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *ProposalSetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalSetReadyEvent)
				if err := _Proposal.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseSetReadyEvent(log types.Log) (*ProposalSetReadyEvent, error) {
	event := new(ProposalSetReadyEvent)
	if err := _Proposal.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalSetSummaryEventIterator is returned from FilterSetSummaryEvent and is used to iterate over the raw logs and unpacked data for SetSummaryEvent events raised by the Proposal contract.
type ProposalSetSummaryEventIterator struct {
	Event *ProposalSetSummaryEvent // Event containing the contract specifics and raw log

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
func (it *ProposalSetSummaryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalSetSummaryEvent)
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
		it.Event = new(ProposalSetSummaryEvent)
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
func (it *ProposalSetSummaryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalSetSummaryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalSetSummaryEvent represents a SetSummaryEvent event raised by the Proposal contract.
type ProposalSetSummaryEvent struct {
	Summary string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetSummaryEvent is a free log retrieval operation binding the contract event 0x84bcb3c5ceb5df029a1d269ea5f54189785acf258ff5fb02e9abe50584c705ae.
//
// Solidity: event setSummaryEvent(string summary)
func (_Proposal *ProposalFilterer) FilterSetSummaryEvent(opts *bind.FilterOpts) (*ProposalSetSummaryEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "setSummaryEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalSetSummaryEventIterator{contract: _Proposal.contract, event: "setSummaryEvent", logs: logs, sub: sub}, nil
}

// WatchSetSummaryEvent is a free log subscription operation binding the contract event 0x84bcb3c5ceb5df029a1d269ea5f54189785acf258ff5fb02e9abe50584c705ae.
//
// Solidity: event setSummaryEvent(string summary)
func (_Proposal *ProposalFilterer) WatchSetSummaryEvent(opts *bind.WatchOpts, sink chan<- *ProposalSetSummaryEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "setSummaryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalSetSummaryEvent)
				if err := _Proposal.contract.UnpackLog(event, "setSummaryEvent", log); err != nil {
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

// ParseSetSummaryEvent is a log parse operation binding the contract event 0x84bcb3c5ceb5df029a1d269ea5f54189785acf258ff5fb02e9abe50584c705ae.
//
// Solidity: event setSummaryEvent(string summary)
func (_Proposal *ProposalFilterer) ParseSetSummaryEvent(log types.Log) (*ProposalSetSummaryEvent, error) {
	event := new(ProposalSetSummaryEvent)
	if err := _Proposal.contract.UnpackLog(event, "setSummaryEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalSetUuidAddressV2EventIterator is returned from FilterSetUuidAddressV2Event and is used to iterate over the raw logs and unpacked data for SetUuidAddressV2Event events raised by the Proposal contract.
type ProposalSetUuidAddressV2EventIterator struct {
	Event *ProposalSetUuidAddressV2Event // Event containing the contract specifics and raw log

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
func (it *ProposalSetUuidAddressV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalSetUuidAddressV2Event)
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
		it.Event = new(ProposalSetUuidAddressV2Event)
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
func (it *ProposalSetUuidAddressV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalSetUuidAddressV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalSetUuidAddressV2Event represents a SetUuidAddressV2Event event raised by the Proposal contract.
type ProposalSetUuidAddressV2Event struct {
	ProposalId *big.Int
	Addr       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetUuidAddressV2Event is a free log retrieval operation binding the contract event 0xfbaae34ae34012d80151ef9dfb58f9e7e004d3ebb7ab113c8f100c6b79da712c.
//
// Solidity: event setUuidAddressV2Event(uint256 proposalId, address addr)
func (_Proposal *ProposalFilterer) FilterSetUuidAddressV2Event(opts *bind.FilterOpts) (*ProposalSetUuidAddressV2EventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "setUuidAddressV2Event")
	if err != nil {
		return nil, err
	}
	return &ProposalSetUuidAddressV2EventIterator{contract: _Proposal.contract, event: "setUuidAddressV2Event", logs: logs, sub: sub}, nil
}

// WatchSetUuidAddressV2Event is a free log subscription operation binding the contract event 0xfbaae34ae34012d80151ef9dfb58f9e7e004d3ebb7ab113c8f100c6b79da712c.
//
// Solidity: event setUuidAddressV2Event(uint256 proposalId, address addr)
func (_Proposal *ProposalFilterer) WatchSetUuidAddressV2Event(opts *bind.WatchOpts, sink chan<- *ProposalSetUuidAddressV2Event) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "setUuidAddressV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalSetUuidAddressV2Event)
				if err := _Proposal.contract.UnpackLog(event, "setUuidAddressV2Event", log); err != nil {
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

// ParseSetUuidAddressV2Event is a log parse operation binding the contract event 0xfbaae34ae34012d80151ef9dfb58f9e7e004d3ebb7ab113c8f100c6b79da712c.
//
// Solidity: event setUuidAddressV2Event(uint256 proposalId, address addr)
func (_Proposal *ProposalFilterer) ParseSetUuidAddressV2Event(log types.Log) (*ProposalSetUuidAddressV2Event, error) {
	event := new(ProposalSetUuidAddressV2Event)
	if err := _Proposal.contract.UnpackLog(event, "setUuidAddressV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalSetVoterPowerV2EventIterator is returned from FilterSetVoterPowerV2Event and is used to iterate over the raw logs and unpacked data for SetVoterPowerV2Event events raised by the Proposal contract.
type ProposalSetVoterPowerV2EventIterator struct {
	Event *ProposalSetVoterPowerV2Event // Event containing the contract specifics and raw log

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
func (it *ProposalSetVoterPowerV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalSetVoterPowerV2Event)
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
		it.Event = new(ProposalSetVoterPowerV2Event)
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
func (it *ProposalSetVoterPowerV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalSetVoterPowerV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalSetVoterPowerV2Event represents a SetVoterPowerV2Event event raised by the Proposal contract.
type ProposalSetVoterPowerV2Event struct {
	ProposalId *big.Int
	Voter      common.Address
	Nvp        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetVoterPowerV2Event is a free log retrieval operation binding the contract event 0xa11af0302de1424d896f4ae50da766ee949a6e616d7211779b7aa13c3a82c21d.
//
// Solidity: event setVoterPowerV2Event(uint256 proposalId, address voter, uint256 nvp)
func (_Proposal *ProposalFilterer) FilterSetVoterPowerV2Event(opts *bind.FilterOpts) (*ProposalSetVoterPowerV2EventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "setVoterPowerV2Event")
	if err != nil {
		return nil, err
	}
	return &ProposalSetVoterPowerV2EventIterator{contract: _Proposal.contract, event: "setVoterPowerV2Event", logs: logs, sub: sub}, nil
}

// WatchSetVoterPowerV2Event is a free log subscription operation binding the contract event 0xa11af0302de1424d896f4ae50da766ee949a6e616d7211779b7aa13c3a82c21d.
//
// Solidity: event setVoterPowerV2Event(uint256 proposalId, address voter, uint256 nvp)
func (_Proposal *ProposalFilterer) WatchSetVoterPowerV2Event(opts *bind.WatchOpts, sink chan<- *ProposalSetVoterPowerV2Event) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "setVoterPowerV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalSetVoterPowerV2Event)
				if err := _Proposal.contract.UnpackLog(event, "setVoterPowerV2Event", log); err != nil {
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

// ParseSetVoterPowerV2Event is a log parse operation binding the contract event 0xa11af0302de1424d896f4ae50da766ee949a6e616d7211779b7aa13c3a82c21d.
//
// Solidity: event setVoterPowerV2Event(uint256 proposalId, address voter, uint256 nvp)
func (_Proposal *ProposalFilterer) ParseSetVoterPowerV2Event(log types.Log) (*ProposalSetVoterPowerV2Event, error) {
	event := new(ProposalSetVoterPowerV2Event)
	if err := _Proposal.contract.UnpackLog(event, "setVoterPowerV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalUpdatePolicyEventIterator is returned from FilterUpdatePolicyEvent and is used to iterate over the raw logs and unpacked data for UpdatePolicyEvent events raised by the Proposal contract.
type ProposalUpdatePolicyEventIterator struct {
	Event *ProposalUpdatePolicyEvent // Event containing the contract specifics and raw log

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
func (it *ProposalUpdatePolicyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalUpdatePolicyEvent)
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
		it.Event = new(ProposalUpdatePolicyEvent)
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
func (it *ProposalUpdatePolicyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalUpdatePolicyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalUpdatePolicyEvent represents a UpdatePolicyEvent event raised by the Proposal contract.
type ProposalUpdatePolicyEvent struct {
	PolicyName string
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdatePolicyEvent is a free log retrieval operation binding the contract event 0x3aa3ca4a0dc63ad7490d619117c294ce617268d8364b3ce4c9fe2dd107b82a3d.
//
// Solidity: event updatePolicyEvent(string policyName_, uint256 value_)
func (_Proposal *ProposalFilterer) FilterUpdatePolicyEvent(opts *bind.FilterOpts) (*ProposalUpdatePolicyEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "updatePolicyEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalUpdatePolicyEventIterator{contract: _Proposal.contract, event: "updatePolicyEvent", logs: logs, sub: sub}, nil
}

// WatchUpdatePolicyEvent is a free log subscription operation binding the contract event 0x3aa3ca4a0dc63ad7490d619117c294ce617268d8364b3ce4c9fe2dd107b82a3d.
//
// Solidity: event updatePolicyEvent(string policyName_, uint256 value_)
func (_Proposal *ProposalFilterer) WatchUpdatePolicyEvent(opts *bind.WatchOpts, sink chan<- *ProposalUpdatePolicyEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "updatePolicyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalUpdatePolicyEvent)
				if err := _Proposal.contract.UnpackLog(event, "updatePolicyEvent", log); err != nil {
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

// ParseUpdatePolicyEvent is a log parse operation binding the contract event 0x3aa3ca4a0dc63ad7490d619117c294ce617268d8364b3ce4c9fe2dd107b82a3d.
//
// Solidity: event updatePolicyEvent(string policyName_, uint256 value_)
func (_Proposal *ProposalFilterer) ParseUpdatePolicyEvent(log types.Log) (*ProposalUpdatePolicyEvent, error) {
	event := new(ProposalUpdatePolicyEvent)
	if err := _Proposal.contract.UnpackLog(event, "updatePolicyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalUpgradeEventIterator is returned from FilterUpgradeEvent and is used to iterate over the raw logs and unpacked data for UpgradeEvent events raised by the Proposal contract.
type ProposalUpgradeEventIterator struct {
	Event *ProposalUpgradeEvent // Event containing the contract specifics and raw log

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
func (it *ProposalUpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalUpgradeEvent)
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
		it.Event = new(ProposalUpgradeEvent)
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
func (it *ProposalUpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalUpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalUpgradeEvent represents a UpgradeEvent event raised by the Proposal contract.
type ProposalUpgradeEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpgradeEvent is a free log retrieval operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Proposal *ProposalFilterer) FilterUpgradeEvent(opts *bind.FilterOpts) (*ProposalUpgradeEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalUpgradeEventIterator{contract: _Proposal.contract, event: "upgradeEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeEvent is a free log subscription operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_Proposal *ProposalFilterer) WatchUpgradeEvent(opts *bind.WatchOpts, sink chan<- *ProposalUpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalUpgradeEvent)
				if err := _Proposal.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseUpgradeEvent(log types.Log) (*ProposalUpgradeEvent, error) {
	event := new(ProposalUpgradeEvent)
	if err := _Proposal.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProposalUpgradeRegistryEventIterator is returned from FilterUpgradeRegistryEvent and is used to iterate over the raw logs and unpacked data for UpgradeRegistryEvent events raised by the Proposal contract.
type ProposalUpgradeRegistryEventIterator struct {
	Event *ProposalUpgradeRegistryEvent // Event containing the contract specifics and raw log

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
func (it *ProposalUpgradeRegistryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalUpgradeRegistryEvent)
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
		it.Event = new(ProposalUpgradeRegistryEvent)
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
func (it *ProposalUpgradeRegistryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalUpgradeRegistryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalUpgradeRegistryEvent represents a UpgradeRegistryEvent event raised by the Proposal contract.
type ProposalUpgradeRegistryEvent struct {
	Reg common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpgradeRegistryEvent is a free log retrieval operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Proposal *ProposalFilterer) FilterUpgradeRegistryEvent(opts *bind.FilterOpts) (*ProposalUpgradeRegistryEventIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return &ProposalUpgradeRegistryEventIterator{contract: _Proposal.contract, event: "upgradeRegistryEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeRegistryEvent is a free log subscription operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_Proposal *ProposalFilterer) WatchUpgradeRegistryEvent(opts *bind.WatchOpts, sink chan<- *ProposalUpgradeRegistryEvent) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalUpgradeRegistryEvent)
				if err := _Proposal.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseUpgradeRegistryEvent(log types.Log) (*ProposalUpgradeRegistryEvent, error) {
	event := new(ProposalUpgradeRegistryEvent)
	if err := _Proposal.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
