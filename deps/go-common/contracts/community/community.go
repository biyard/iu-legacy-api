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

// CommunityProposalMetaData contains all meta data concerning the CommunityProposal contract.
var CommunityProposalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"strings2d\",\"type\":\"string[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"numbers2d\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"addresses2d\",\"type\":\"address[][]\"},{\"internalType\":\"bool[][]\",\"name\":\"bools2d\",\"type\":\"bool[][]\"}],\"internalType\":\"structDeploymentParameterV1\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"stateAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"actionAfterFinishVoteEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"beforeUpgradeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"publicComment\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enableComment\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"proposalTypes\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"proposalValues\",\"type\":\"uint256[]\"}],\"name\":\"changeInitialDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"proposaltypes\",\"type\":\"string[]\"}],\"name\":\"changeProposalTypeEvnet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityFlag\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitAppDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activityType\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voteAppName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"voteStatus\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structProposalSummary\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"emitProposalActivityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initializeHookEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contents_\",\"type\":\"string\"}],\"name\":\"setCommentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contents_\",\"type\":\"string\"}],\"name\":\"setCommentV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structProposalParameterV2\",\"name\":\"pro\",\"type\":\"tuple\"}],\"name\":\"setProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"name\":\"setReadyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"summary_\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fullContents_\",\"type\":\"string\"}],\"name\":\"submitCommunityProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"summary_\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fullContents_\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proposalType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"votingApp\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maximumVotePower\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"options\",\"type\":\"string[]\"}],\"name\":\"submitCommunityProposalV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistryEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"accept\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comments\",\"type\":\"string\"}],\"name\":\"voteCommunityProposalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"actionAfterFinishVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"convertProposalIdToUuid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getAccepters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getAcceptersV2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getComment\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getCommentV2\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getParticipantsV2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolicies\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyName_\",\"type\":\"string\"}],\"name\":\"getPolicy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"getProposalStructure\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStructureV2\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameterV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getProposalType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposalType\",\"type\":\"string\"}],\"name\":\"getProposalTypeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalTypeV2\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalV2\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameterV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getRejecters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getRejectersV2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateAddr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"}],\"name\":\"getVoterValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getVoterValueV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listProposals\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameExt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalTypeList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contents_\",\"type\":\"string\"}],\"name\":\"setComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contents_\",\"type\":\"string\"}],\"name\":\"setCommentV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contents\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"didFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didPass\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectedPowers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeVp\",\"type\":\"uint256\"}],\"internalType\":\"structProposalParameterV2\",\"name\":\"pro\",\"type\":\"tuple\"}],\"name\":\"setProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fullContents_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"}],\"name\":\"submitCommunityProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fullContents_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposalType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"votingApp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"baseNumVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumVotePower\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"options\",\"type\":\"string[]\"}],\"name\":\"submitProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"updateProposalStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reg\",\"type\":\"address\"}],\"name\":\"upgradeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"accept\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comments\",\"type\":\"string\"}],\"name\":\"voteCommunityProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CommunityProposalABI is the input ABI used to generate the binding from.
// Deprecated: Use CommunityProposalMetaData.ABI instead.
var CommunityProposalABI = CommunityProposalMetaData.ABI

// CommunityProposalBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CommunityProposalBinRuntime = ``

// CommunityProposal is an auto generated Go binding around a Klaytn contract.
type CommunityProposal struct {
	CommunityProposalCaller     // Read-only binding to the contract
	CommunityProposalTransactor // Write-only binding to the contract
	CommunityProposalFilterer   // Log filterer for contract events
}

// CommunityProposalCaller is an auto generated read-only Go binding around a Klaytn contract.
type CommunityProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommunityProposalTransactor is an auto generated write-only Go binding around a Klaytn contract.
type CommunityProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommunityProposalFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type CommunityProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommunityProposalSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type CommunityProposalSession struct {
	Contract     *CommunityProposal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CommunityProposalCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type CommunityProposalCallerSession struct {
	Contract *CommunityProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CommunityProposalTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type CommunityProposalTransactorSession struct {
	Contract     *CommunityProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CommunityProposalRaw is an auto generated low-level Go binding around a Klaytn contract.
type CommunityProposalRaw struct {
	Contract *CommunityProposal // Generic contract binding to access the raw methods on
}

// CommunityProposalCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type CommunityProposalCallerRaw struct {
	Contract *CommunityProposalCaller // Generic read-only contract binding to access the raw methods on
}

// CommunityProposalTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type CommunityProposalTransactorRaw struct {
	Contract *CommunityProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommunityProposal creates a new instance of CommunityProposal, bound to a specific deployed contract.
func NewCommunityProposal(address common.Address, backend bind.ContractBackend) (*CommunityProposal, error) {
	contract, err := bindCommunityProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommunityProposal{CommunityProposalCaller: CommunityProposalCaller{contract: contract}, CommunityProposalTransactor: CommunityProposalTransactor{contract: contract}, CommunityProposalFilterer: CommunityProposalFilterer{contract: contract}}, nil
}

// NewCommunityProposalCaller creates a new read-only instance of CommunityProposal, bound to a specific deployed contract.
func NewCommunityProposalCaller(address common.Address, caller bind.ContractCaller) (*CommunityProposalCaller, error) {
	contract, err := bindCommunityProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommunityProposalCaller{contract: contract}, nil
}

// NewCommunityProposalTransactor creates a new write-only instance of CommunityProposal, bound to a specific deployed contract.
func NewCommunityProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*CommunityProposalTransactor, error) {
	contract, err := bindCommunityProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommunityProposalTransactor{contract: contract}, nil
}

// NewCommunityProposalFilterer creates a new log filterer instance of CommunityProposal, bound to a specific deployed contract.
func NewCommunityProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*CommunityProposalFilterer, error) {
	contract, err := bindCommunityProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommunityProposalFilterer{contract: contract}, nil
}

// bindCommunityProposal binds a generic wrapper to an already deployed contract.
func bindCommunityProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommunityProposalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommunityProposal *CommunityProposalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CommunityProposal.Contract.CommunityProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommunityProposal *CommunityProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommunityProposal.Contract.CommunityProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommunityProposal *CommunityProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommunityProposal.Contract.CommunityProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommunityProposal *CommunityProposalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CommunityProposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommunityProposal *CommunityProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommunityProposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommunityProposal *CommunityProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommunityProposal.Contract.contract.Transact(opts, method, params...)
}

// ConvertProposalIdToUuid is a free data retrieval call binding the contract method 0x08b2054b.
//
// Solidity: function convertProposalIdToUuid(uint256 proposalId) view returns(string)
func (_CommunityProposal *CommunityProposalCaller) ConvertProposalIdToUuid(opts *bind.CallOpts, proposalId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "convertProposalIdToUuid", proposalId)
	return *ret0, err
}

// ConvertProposalIdToUuid is a free data retrieval call binding the contract method 0x08b2054b.
//
// Solidity: function convertProposalIdToUuid(uint256 proposalId) view returns(string)
func (_CommunityProposal *CommunityProposalSession) ConvertProposalIdToUuid(proposalId *big.Int) (string, error) {
	return _CommunityProposal.Contract.ConvertProposalIdToUuid(&_CommunityProposal.CallOpts, proposalId)
}

// ConvertProposalIdToUuid is a free data retrieval call binding the contract method 0x08b2054b.
//
// Solidity: function convertProposalIdToUuid(uint256 proposalId) view returns(string)
func (_CommunityProposal *CommunityProposalCallerSession) ConvertProposalIdToUuid(proposalId *big.Int) (string, error) {
	return _CommunityProposal.Contract.ConvertProposalIdToUuid(&_CommunityProposal.CallOpts, proposalId)
}

// GetAccepters is a free data retrieval call binding the contract method 0x1d774ba8.
//
// Solidity: function getAccepters(string uuid_) view returns(address[])
func (_CommunityProposal *CommunityProposalCaller) GetAccepters(opts *bind.CallOpts, uuid_ string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getAccepters", uuid_)
	return *ret0, err
}

// GetAccepters is a free data retrieval call binding the contract method 0x1d774ba8.
//
// Solidity: function getAccepters(string uuid_) view returns(address[])
func (_CommunityProposal *CommunityProposalSession) GetAccepters(uuid_ string) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetAccepters(&_CommunityProposal.CallOpts, uuid_)
}

// GetAccepters is a free data retrieval call binding the contract method 0x1d774ba8.
//
// Solidity: function getAccepters(string uuid_) view returns(address[])
func (_CommunityProposal *CommunityProposalCallerSession) GetAccepters(uuid_ string) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetAccepters(&_CommunityProposal.CallOpts, uuid_)
}

// GetAcceptersV2 is a free data retrieval call binding the contract method 0x95f27b7b.
//
// Solidity: function getAcceptersV2(uint256 proposalId) view returns(address[])
func (_CommunityProposal *CommunityProposalCaller) GetAcceptersV2(opts *bind.CallOpts, proposalId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getAcceptersV2", proposalId)
	return *ret0, err
}

// GetAcceptersV2 is a free data retrieval call binding the contract method 0x95f27b7b.
//
// Solidity: function getAcceptersV2(uint256 proposalId) view returns(address[])
func (_CommunityProposal *CommunityProposalSession) GetAcceptersV2(proposalId *big.Int) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetAcceptersV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetAcceptersV2 is a free data retrieval call binding the contract method 0x95f27b7b.
//
// Solidity: function getAcceptersV2(uint256 proposalId) view returns(address[])
func (_CommunityProposal *CommunityProposalCallerSession) GetAcceptersV2(proposalId *big.Int) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetAcceptersV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetComment is a free data retrieval call binding the contract method 0xcbbaa589.
//
// Solidity: function getComment(string uuid_) view returns(string[])
func (_CommunityProposal *CommunityProposalCaller) GetComment(opts *bind.CallOpts, uuid_ string) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getComment", uuid_)
	return *ret0, err
}

// GetComment is a free data retrieval call binding the contract method 0xcbbaa589.
//
// Solidity: function getComment(string uuid_) view returns(string[])
func (_CommunityProposal *CommunityProposalSession) GetComment(uuid_ string) ([]string, error) {
	return _CommunityProposal.Contract.GetComment(&_CommunityProposal.CallOpts, uuid_)
}

// GetComment is a free data retrieval call binding the contract method 0xcbbaa589.
//
// Solidity: function getComment(string uuid_) view returns(string[])
func (_CommunityProposal *CommunityProposalCallerSession) GetComment(uuid_ string) ([]string, error) {
	return _CommunityProposal.Contract.GetComment(&_CommunityProposal.CallOpts, uuid_)
}

// GetCommentV2 is a free data retrieval call binding the contract method 0x0454d213.
//
// Solidity: function getCommentV2(uint256 proposalId) view returns(string[])
func (_CommunityProposal *CommunityProposalCaller) GetCommentV2(opts *bind.CallOpts, proposalId *big.Int) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getCommentV2", proposalId)
	return *ret0, err
}

// GetCommentV2 is a free data retrieval call binding the contract method 0x0454d213.
//
// Solidity: function getCommentV2(uint256 proposalId) view returns(string[])
func (_CommunityProposal *CommunityProposalSession) GetCommentV2(proposalId *big.Int) ([]string, error) {
	return _CommunityProposal.Contract.GetCommentV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetCommentV2 is a free data retrieval call binding the contract method 0x0454d213.
//
// Solidity: function getCommentV2(uint256 proposalId) view returns(string[])
func (_CommunityProposal *CommunityProposalCallerSession) GetCommentV2(proposalId *big.Int) ([]string, error) {
	return _CommunityProposal.Contract.GetCommentV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetParticipants is a free data retrieval call binding the contract method 0x7e90def7.
//
// Solidity: function getParticipants(string uuid_) view returns(address[])
func (_CommunityProposal *CommunityProposalCaller) GetParticipants(opts *bind.CallOpts, uuid_ string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getParticipants", uuid_)
	return *ret0, err
}

// GetParticipants is a free data retrieval call binding the contract method 0x7e90def7.
//
// Solidity: function getParticipants(string uuid_) view returns(address[])
func (_CommunityProposal *CommunityProposalSession) GetParticipants(uuid_ string) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetParticipants(&_CommunityProposal.CallOpts, uuid_)
}

// GetParticipants is a free data retrieval call binding the contract method 0x7e90def7.
//
// Solidity: function getParticipants(string uuid_) view returns(address[])
func (_CommunityProposal *CommunityProposalCallerSession) GetParticipants(uuid_ string) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetParticipants(&_CommunityProposal.CallOpts, uuid_)
}

// GetParticipantsV2 is a free data retrieval call binding the contract method 0x71166f71.
//
// Solidity: function getParticipantsV2(uint256 proposalId) view returns(address[])
func (_CommunityProposal *CommunityProposalCaller) GetParticipantsV2(opts *bind.CallOpts, proposalId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getParticipantsV2", proposalId)
	return *ret0, err
}

// GetParticipantsV2 is a free data retrieval call binding the contract method 0x71166f71.
//
// Solidity: function getParticipantsV2(uint256 proposalId) view returns(address[])
func (_CommunityProposal *CommunityProposalSession) GetParticipantsV2(proposalId *big.Int) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetParticipantsV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetParticipantsV2 is a free data retrieval call binding the contract method 0x71166f71.
//
// Solidity: function getParticipantsV2(uint256 proposalId) view returns(address[])
func (_CommunityProposal *CommunityProposalCallerSession) GetParticipantsV2(proposalId *big.Int) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetParticipantsV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns(string[])
func (_CommunityProposal *CommunityProposalCaller) GetPolicies(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getPolicies")
	return *ret0, err
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns(string[])
func (_CommunityProposal *CommunityProposalSession) GetPolicies() ([]string, error) {
	return _CommunityProposal.Contract.GetPolicies(&_CommunityProposal.CallOpts)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns(string[])
func (_CommunityProposal *CommunityProposalCallerSession) GetPolicies() ([]string, error) {
	return _CommunityProposal.Contract.GetPolicies(&_CommunityProposal.CallOpts)
}

// GetPolicy is a free data retrieval call binding the contract method 0x60dd5f90.
//
// Solidity: function getPolicy(string policyName_) view returns(uint256)
func (_CommunityProposal *CommunityProposalCaller) GetPolicy(opts *bind.CallOpts, policyName_ string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getPolicy", policyName_)
	return *ret0, err
}

// GetPolicy is a free data retrieval call binding the contract method 0x60dd5f90.
//
// Solidity: function getPolicy(string policyName_) view returns(uint256)
func (_CommunityProposal *CommunityProposalSession) GetPolicy(policyName_ string) (*big.Int, error) {
	return _CommunityProposal.Contract.GetPolicy(&_CommunityProposal.CallOpts, policyName_)
}

// GetPolicy is a free data retrieval call binding the contract method 0x60dd5f90.
//
// Solidity: function getPolicy(string policyName_) view returns(uint256)
func (_CommunityProposal *CommunityProposalCallerSession) GetPolicy(policyName_ string) (*big.Int, error) {
	return _CommunityProposal.Contract.GetPolicy(&_CommunityProposal.CallOpts, policyName_)
}

// GetProposal is a free data retrieval call binding the contract method 0x25d3a09d.
//
// Solidity: function getProposal(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalCaller) GetProposal(opts *bind.CallOpts, uuid string) (ProposalParameter, error) {
	var (
		ret0 = new(ProposalParameter)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getProposal", uuid)
	return *ret0, err
}

// GetProposal is a free data retrieval call binding the contract method 0x25d3a09d.
//
// Solidity: function getProposal(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalSession) GetProposal(uuid string) (ProposalParameter, error) {
	return _CommunityProposal.Contract.GetProposal(&_CommunityProposal.CallOpts, uuid)
}

// GetProposal is a free data retrieval call binding the contract method 0x25d3a09d.
//
// Solidity: function getProposal(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalCallerSession) GetProposal(uuid string) (ProposalParameter, error) {
	return _CommunityProposal.Contract.GetProposal(&_CommunityProposal.CallOpts, uuid)
}

// GetProposalStructure is a free data retrieval call binding the contract method 0xfef2293f.
//
// Solidity: function getProposalStructure(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalCaller) GetProposalStructure(opts *bind.CallOpts, uuid string) (ProposalParameter, error) {
	var (
		ret0 = new(ProposalParameter)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getProposalStructure", uuid)
	return *ret0, err
}

// GetProposalStructure is a free data retrieval call binding the contract method 0xfef2293f.
//
// Solidity: function getProposalStructure(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalSession) GetProposalStructure(uuid string) (ProposalParameter, error) {
	return _CommunityProposal.Contract.GetProposalStructure(&_CommunityProposal.CallOpts, uuid)
}

// GetProposalStructure is a free data retrieval call binding the contract method 0xfef2293f.
//
// Solidity: function getProposalStructure(string uuid) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalCallerSession) GetProposalStructure(uuid string) (ProposalParameter, error) {
	return _CommunityProposal.Contract.GetProposalStructure(&_CommunityProposal.CallOpts, uuid)
}

// GetProposalStructureV2 is a free data retrieval call binding the contract method 0x3cd7c90b.
//
// Solidity: function getProposalStructureV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalCaller) GetProposalStructureV2(opts *bind.CallOpts, proposalId *big.Int) (ProposalParameterV2, error) {
	var (
		ret0 = new(ProposalParameterV2)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getProposalStructureV2", proposalId)
	return *ret0, err
}

// GetProposalStructureV2 is a free data retrieval call binding the contract method 0x3cd7c90b.
//
// Solidity: function getProposalStructureV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalSession) GetProposalStructureV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _CommunityProposal.Contract.GetProposalStructureV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetProposalStructureV2 is a free data retrieval call binding the contract method 0x3cd7c90b.
//
// Solidity: function getProposalStructureV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalCallerSession) GetProposalStructureV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _CommunityProposal.Contract.GetProposalStructureV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetProposalType is a free data retrieval call binding the contract method 0x9bcfd974.
//
// Solidity: function getProposalType(string uuid_) view returns(string)
func (_CommunityProposal *CommunityProposalCaller) GetProposalType(opts *bind.CallOpts, uuid_ string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getProposalType", uuid_)
	return *ret0, err
}

// GetProposalType is a free data retrieval call binding the contract method 0x9bcfd974.
//
// Solidity: function getProposalType(string uuid_) view returns(string)
func (_CommunityProposal *CommunityProposalSession) GetProposalType(uuid_ string) (string, error) {
	return _CommunityProposal.Contract.GetProposalType(&_CommunityProposal.CallOpts, uuid_)
}

// GetProposalType is a free data retrieval call binding the contract method 0x9bcfd974.
//
// Solidity: function getProposalType(string uuid_) view returns(string)
func (_CommunityProposal *CommunityProposalCallerSession) GetProposalType(uuid_ string) (string, error) {
	return _CommunityProposal.Contract.GetProposalType(&_CommunityProposal.CallOpts, uuid_)
}

// GetProposalTypeFee is a free data retrieval call binding the contract method 0x3a94243f.
//
// Solidity: function getProposalTypeFee(string proposalType) view returns(uint256)
func (_CommunityProposal *CommunityProposalCaller) GetProposalTypeFee(opts *bind.CallOpts, proposalType string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getProposalTypeFee", proposalType)
	return *ret0, err
}

// GetProposalTypeFee is a free data retrieval call binding the contract method 0x3a94243f.
//
// Solidity: function getProposalTypeFee(string proposalType) view returns(uint256)
func (_CommunityProposal *CommunityProposalSession) GetProposalTypeFee(proposalType string) (*big.Int, error) {
	return _CommunityProposal.Contract.GetProposalTypeFee(&_CommunityProposal.CallOpts, proposalType)
}

// GetProposalTypeFee is a free data retrieval call binding the contract method 0x3a94243f.
//
// Solidity: function getProposalTypeFee(string proposalType) view returns(uint256)
func (_CommunityProposal *CommunityProposalCallerSession) GetProposalTypeFee(proposalType string) (*big.Int, error) {
	return _CommunityProposal.Contract.GetProposalTypeFee(&_CommunityProposal.CallOpts, proposalType)
}

// GetProposalTypeV2 is a free data retrieval call binding the contract method 0x40eb05dd.
//
// Solidity: function getProposalTypeV2(uint256 proposalId) view returns(string)
func (_CommunityProposal *CommunityProposalCaller) GetProposalTypeV2(opts *bind.CallOpts, proposalId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getProposalTypeV2", proposalId)
	return *ret0, err
}

// GetProposalTypeV2 is a free data retrieval call binding the contract method 0x40eb05dd.
//
// Solidity: function getProposalTypeV2(uint256 proposalId) view returns(string)
func (_CommunityProposal *CommunityProposalSession) GetProposalTypeV2(proposalId *big.Int) (string, error) {
	return _CommunityProposal.Contract.GetProposalTypeV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetProposalTypeV2 is a free data retrieval call binding the contract method 0x40eb05dd.
//
// Solidity: function getProposalTypeV2(uint256 proposalId) view returns(string)
func (_CommunityProposal *CommunityProposalCallerSession) GetProposalTypeV2(proposalId *big.Int) (string, error) {
	return _CommunityProposal.Contract.GetProposalTypeV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetProposalV2 is a free data retrieval call binding the contract method 0xc6d61e6b.
//
// Solidity: function getProposalV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalCaller) GetProposalV2(opts *bind.CallOpts, proposalId *big.Int) (ProposalParameterV2, error) {
	var (
		ret0 = new(ProposalParameterV2)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getProposalV2", proposalId)
	return *ret0, err
}

// GetProposalV2 is a free data retrieval call binding the contract method 0xc6d61e6b.
//
// Solidity: function getProposalV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalSession) GetProposalV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _CommunityProposal.Contract.GetProposalV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetProposalV2 is a free data retrieval call binding the contract method 0xc6d61e6b.
//
// Solidity: function getProposalV2(uint256 proposalId) view returns((uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256))
func (_CommunityProposal *CommunityProposalCallerSession) GetProposalV2(proposalId *big.Int) (ProposalParameterV2, error) {
	return _CommunityProposal.Contract.GetProposalV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetRejecters is a free data retrieval call binding the contract method 0x2149e6c6.
//
// Solidity: function getRejecters(string uuid_) view returns(address[])
func (_CommunityProposal *CommunityProposalCaller) GetRejecters(opts *bind.CallOpts, uuid_ string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getRejecters", uuid_)
	return *ret0, err
}

// GetRejecters is a free data retrieval call binding the contract method 0x2149e6c6.
//
// Solidity: function getRejecters(string uuid_) view returns(address[])
func (_CommunityProposal *CommunityProposalSession) GetRejecters(uuid_ string) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetRejecters(&_CommunityProposal.CallOpts, uuid_)
}

// GetRejecters is a free data retrieval call binding the contract method 0x2149e6c6.
//
// Solidity: function getRejecters(string uuid_) view returns(address[])
func (_CommunityProposal *CommunityProposalCallerSession) GetRejecters(uuid_ string) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetRejecters(&_CommunityProposal.CallOpts, uuid_)
}

// GetRejectersV2 is a free data retrieval call binding the contract method 0x2350bdc7.
//
// Solidity: function getRejectersV2(uint256 proposalId) view returns(address[])
func (_CommunityProposal *CommunityProposalCaller) GetRejectersV2(opts *bind.CallOpts, proposalId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getRejectersV2", proposalId)
	return *ret0, err
}

// GetRejectersV2 is a free data retrieval call binding the contract method 0x2350bdc7.
//
// Solidity: function getRejectersV2(uint256 proposalId) view returns(address[])
func (_CommunityProposal *CommunityProposalSession) GetRejectersV2(proposalId *big.Int) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetRejectersV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetRejectersV2 is a free data retrieval call binding the contract method 0x2350bdc7.
//
// Solidity: function getRejectersV2(uint256 proposalId) view returns(address[])
func (_CommunityProposal *CommunityProposalCallerSession) GetRejectersV2(proposalId *big.Int) ([]common.Address, error) {
	return _CommunityProposal.Contract.GetRejectersV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_CommunityProposal *CommunityProposalCaller) GetStateAddr(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getStateAddr")
	return *ret0, err
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_CommunityProposal *CommunityProposalSession) GetStateAddr() ([]common.Address, error) {
	return _CommunityProposal.Contract.GetStateAddr(&_CommunityProposal.CallOpts)
}

// GetStateAddr is a free data retrieval call binding the contract method 0x12567bd1.
//
// Solidity: function getStateAddr() view returns(address[])
func (_CommunityProposal *CommunityProposalCallerSession) GetStateAddr() ([]common.Address, error) {
	return _CommunityProposal.Contract.GetStateAddr(&_CommunityProposal.CallOpts)
}

// GetVoterValue is a free data retrieval call binding the contract method 0x1f3f882a.
//
// Solidity: function getVoterValue(string uuid_) view returns(uint256)
func (_CommunityProposal *CommunityProposalCaller) GetVoterValue(opts *bind.CallOpts, uuid_ string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getVoterValue", uuid_)
	return *ret0, err
}

// GetVoterValue is a free data retrieval call binding the contract method 0x1f3f882a.
//
// Solidity: function getVoterValue(string uuid_) view returns(uint256)
func (_CommunityProposal *CommunityProposalSession) GetVoterValue(uuid_ string) (*big.Int, error) {
	return _CommunityProposal.Contract.GetVoterValue(&_CommunityProposal.CallOpts, uuid_)
}

// GetVoterValue is a free data retrieval call binding the contract method 0x1f3f882a.
//
// Solidity: function getVoterValue(string uuid_) view returns(uint256)
func (_CommunityProposal *CommunityProposalCallerSession) GetVoterValue(uuid_ string) (*big.Int, error) {
	return _CommunityProposal.Contract.GetVoterValue(&_CommunityProposal.CallOpts, uuid_)
}

// GetVoterValueV2 is a free data retrieval call binding the contract method 0x9b636bfc.
//
// Solidity: function getVoterValueV2(uint256 proposalId) view returns(uint256)
func (_CommunityProposal *CommunityProposalCaller) GetVoterValueV2(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "getVoterValueV2", proposalId)
	return *ret0, err
}

// GetVoterValueV2 is a free data retrieval call binding the contract method 0x9b636bfc.
//
// Solidity: function getVoterValueV2(uint256 proposalId) view returns(uint256)
func (_CommunityProposal *CommunityProposalSession) GetVoterValueV2(proposalId *big.Int) (*big.Int, error) {
	return _CommunityProposal.Contract.GetVoterValueV2(&_CommunityProposal.CallOpts, proposalId)
}

// GetVoterValueV2 is a free data retrieval call binding the contract method 0x9b636bfc.
//
// Solidity: function getVoterValueV2(uint256 proposalId) view returns(uint256)
func (_CommunityProposal *CommunityProposalCallerSession) GetVoterValueV2(proposalId *big.Int) (*big.Int, error) {
	return _CommunityProposal.Contract.GetVoterValueV2(&_CommunityProposal.CallOpts, proposalId)
}

// ListProposals is a free data retrieval call binding the contract method 0x26c7eb1c.
//
// Solidity: function listProposals() view returns(string[])
func (_CommunityProposal *CommunityProposalCaller) ListProposals(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "listProposals")
	return *ret0, err
}

// ListProposals is a free data retrieval call binding the contract method 0x26c7eb1c.
//
// Solidity: function listProposals() view returns(string[])
func (_CommunityProposal *CommunityProposalSession) ListProposals() ([]string, error) {
	return _CommunityProposal.Contract.ListProposals(&_CommunityProposal.CallOpts)
}

// ListProposals is a free data retrieval call binding the contract method 0x26c7eb1c.
//
// Solidity: function listProposals() view returns(string[])
func (_CommunityProposal *CommunityProposalCallerSession) ListProposals() ([]string, error) {
	return _CommunityProposal.Contract.ListProposals(&_CommunityProposal.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_CommunityProposal *CommunityProposalCaller) NameExt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "nameExt")
	return *ret0, err
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_CommunityProposal *CommunityProposalSession) NameExt() (string, error) {
	return _CommunityProposal.Contract.NameExt(&_CommunityProposal.CallOpts)
}

// NameExt is a free data retrieval call binding the contract method 0xe8a1e296.
//
// Solidity: function nameExt() view returns(string)
func (_CommunityProposal *CommunityProposalCallerSession) NameExt() (string, error) {
	return _CommunityProposal.Contract.NameExt(&_CommunityProposal.CallOpts)
}

// ProposalTypeList is a free data retrieval call binding the contract method 0xc4c9aedd.
//
// Solidity: function proposalTypeList() view returns(string[], uint256[])
func (_CommunityProposal *CommunityProposalCaller) ProposalTypeList(opts *bind.CallOpts) ([]string, []*big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _CommunityProposal.contract.Call(opts, out, "proposalTypeList")
	return *ret0, *ret1, err
}

// ProposalTypeList is a free data retrieval call binding the contract method 0xc4c9aedd.
//
// Solidity: function proposalTypeList() view returns(string[], uint256[])
func (_CommunityProposal *CommunityProposalSession) ProposalTypeList() ([]string, []*big.Int, error) {
	return _CommunityProposal.Contract.ProposalTypeList(&_CommunityProposal.CallOpts)
}

// ProposalTypeList is a free data retrieval call binding the contract method 0xc4c9aedd.
//
// Solidity: function proposalTypeList() view returns(string[], uint256[])
func (_CommunityProposal *CommunityProposalCallerSession) ProposalTypeList() ([]string, []*big.Int, error) {
	return _CommunityProposal.Contract.ProposalTypeList(&_CommunityProposal.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_CommunityProposal *CommunityProposalCaller) Ready(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_CommunityProposal *CommunityProposalSession) Ready() (bool, error) {
	return _CommunityProposal.Contract.Ready(&_CommunityProposal.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() view returns(bool)
func (_CommunityProposal *CommunityProposalCallerSession) Ready() (bool, error) {
	return _CommunityProposal.Contract.Ready(&_CommunityProposal.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_CommunityProposal *CommunityProposalCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CommunityProposal.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_CommunityProposal *CommunityProposalSession) Registry() (common.Address, error) {
	return _CommunityProposal.Contract.Registry(&_CommunityProposal.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_CommunityProposal *CommunityProposalCallerSession) Registry() (common.Address, error) {
	return _CommunityProposal.Contract.Registry(&_CommunityProposal.CallOpts)
}

// ActionAfterFinishVote is a paid mutator transaction binding the contract method 0xb3428d6f.
//
// Solidity: function actionAfterFinishVote(uint256 proposalId) returns()
func (_CommunityProposal *CommunityProposalTransactor) ActionAfterFinishVote(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "actionAfterFinishVote", proposalId)
}

// ActionAfterFinishVote is a paid mutator transaction binding the contract method 0xb3428d6f.
//
// Solidity: function actionAfterFinishVote(uint256 proposalId) returns()
func (_CommunityProposal *CommunityProposalSession) ActionAfterFinishVote(proposalId *big.Int) (*types.Transaction, error) {
	return _CommunityProposal.Contract.ActionAfterFinishVote(&_CommunityProposal.TransactOpts, proposalId)
}

// ActionAfterFinishVote is a paid mutator transaction binding the contract method 0xb3428d6f.
//
// Solidity: function actionAfterFinishVote(uint256 proposalId) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) ActionAfterFinishVote(proposalId *big.Int) (*types.Transaction, error) {
	return _CommunityProposal.Contract.ActionAfterFinishVote(&_CommunityProposal.TransactOpts, proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_CommunityProposal *CommunityProposalTransactor) Initialize(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "initialize", addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_CommunityProposal *CommunityProposalSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _CommunityProposal.Contract.Initialize(&_CommunityProposal.TransactOpts, addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addrs) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) Initialize(addrs []common.Address) (*types.Transaction, error) {
	return _CommunityProposal.Contract.Initialize(&_CommunityProposal.TransactOpts, addrs)
}

// SetComment is a paid mutator transaction binding the contract method 0x328e0b3f.
//
// Solidity: function setComment(string uuid_, string contents_) returns()
func (_CommunityProposal *CommunityProposalTransactor) SetComment(opts *bind.TransactOpts, uuid_ string, contents_ string) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "setComment", uuid_, contents_)
}

// SetComment is a paid mutator transaction binding the contract method 0x328e0b3f.
//
// Solidity: function setComment(string uuid_, string contents_) returns()
func (_CommunityProposal *CommunityProposalSession) SetComment(uuid_ string, contents_ string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SetComment(&_CommunityProposal.TransactOpts, uuid_, contents_)
}

// SetComment is a paid mutator transaction binding the contract method 0x328e0b3f.
//
// Solidity: function setComment(string uuid_, string contents_) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) SetComment(uuid_ string, contents_ string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SetComment(&_CommunityProposal.TransactOpts, uuid_, contents_)
}

// SetCommentV2 is a paid mutator transaction binding the contract method 0xe14ffe6a.
//
// Solidity: function setCommentV2(uint256 proposalId, string contents_) returns()
func (_CommunityProposal *CommunityProposalTransactor) SetCommentV2(opts *bind.TransactOpts, proposalId *big.Int, contents_ string) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "setCommentV2", proposalId, contents_)
}

// SetCommentV2 is a paid mutator transaction binding the contract method 0xe14ffe6a.
//
// Solidity: function setCommentV2(uint256 proposalId, string contents_) returns()
func (_CommunityProposal *CommunityProposalSession) SetCommentV2(proposalId *big.Int, contents_ string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SetCommentV2(&_CommunityProposal.TransactOpts, proposalId, contents_)
}

// SetCommentV2 is a paid mutator transaction binding the contract method 0xe14ffe6a.
//
// Solidity: function setCommentV2(uint256 proposalId, string contents_) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) SetCommentV2(proposalId *big.Int, contents_ string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SetCommentV2(&_CommunityProposal.TransactOpts, proposalId, contents_)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_CommunityProposal *CommunityProposalTransactor) SetForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "setForwarder", forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_CommunityProposal *CommunityProposalSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SetForwarder(&_CommunityProposal.TransactOpts, forwarder)
}

// SetForwarder is a paid mutator transaction binding the contract method 0xb9998a24.
//
// Solidity: function setForwarder(address forwarder) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) SetForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SetForwarder(&_CommunityProposal.TransactOpts, forwarder)
}

// SetProposal is a paid mutator transaction binding the contract method 0xef037285.
//
// Solidity: function setProposal(uint256 proposalId, (uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256) pro) returns()
func (_CommunityProposal *CommunityProposalTransactor) SetProposal(opts *bind.TransactOpts, proposalId *big.Int, pro ProposalParameterV2) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "setProposal", proposalId, pro)
}

// SetProposal is a paid mutator transaction binding the contract method 0xef037285.
//
// Solidity: function setProposal(uint256 proposalId, (uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256) pro) returns()
func (_CommunityProposal *CommunityProposalSession) SetProposal(proposalId *big.Int, pro ProposalParameterV2) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SetProposal(&_CommunityProposal.TransactOpts, proposalId, pro)
}

// SetProposal is a paid mutator transaction binding the contract method 0xef037285.
//
// Solidity: function setProposal(uint256 proposalId, (uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256) pro) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) SetProposal(proposalId *big.Int, pro ProposalParameterV2) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SetProposal(&_CommunityProposal.TransactOpts, proposalId, pro)
}

// SubmitCommunityProposal is a paid mutator transaction binding the contract method 0x26ddd617.
//
// Solidity: function submitCommunityProposal(string uuid_, string summary_, string fullContents_, string proposalType, uint256 finishedAt) returns()
func (_CommunityProposal *CommunityProposalTransactor) SubmitCommunityProposal(opts *bind.TransactOpts, uuid_ string, summary_ string, fullContents_ string, proposalType string, finishedAt *big.Int) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "submitCommunityProposal", uuid_, summary_, fullContents_, proposalType, finishedAt)
}

// SubmitCommunityProposal is a paid mutator transaction binding the contract method 0x26ddd617.
//
// Solidity: function submitCommunityProposal(string uuid_, string summary_, string fullContents_, string proposalType, uint256 finishedAt) returns()
func (_CommunityProposal *CommunityProposalSession) SubmitCommunityProposal(uuid_ string, summary_ string, fullContents_ string, proposalType string, finishedAt *big.Int) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SubmitCommunityProposal(&_CommunityProposal.TransactOpts, uuid_, summary_, fullContents_, proposalType, finishedAt)
}

// SubmitCommunityProposal is a paid mutator transaction binding the contract method 0x26ddd617.
//
// Solidity: function submitCommunityProposal(string uuid_, string summary_, string fullContents_, string proposalType, uint256 finishedAt) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) SubmitCommunityProposal(uuid_ string, summary_ string, fullContents_ string, proposalType string, finishedAt *big.Int) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SubmitCommunityProposal(&_CommunityProposal.TransactOpts, uuid_, summary_, fullContents_, proposalType, finishedAt)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x0359ce16.
//
// Solidity: function submitProposal(string uuid_, string summary_, string fullContents_, string proposalType, string votingApp, uint256 baseNumVote, uint256 finishedAt, uint256 maximumVotePower, string[] options) returns()
func (_CommunityProposal *CommunityProposalTransactor) SubmitProposal(opts *bind.TransactOpts, uuid_ string, summary_ string, fullContents_ string, proposalType string, votingApp string, baseNumVote *big.Int, finishedAt *big.Int, maximumVotePower *big.Int, options []string) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "submitProposal", uuid_, summary_, fullContents_, proposalType, votingApp, baseNumVote, finishedAt, maximumVotePower, options)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x0359ce16.
//
// Solidity: function submitProposal(string uuid_, string summary_, string fullContents_, string proposalType, string votingApp, uint256 baseNumVote, uint256 finishedAt, uint256 maximumVotePower, string[] options) returns()
func (_CommunityProposal *CommunityProposalSession) SubmitProposal(uuid_ string, summary_ string, fullContents_ string, proposalType string, votingApp string, baseNumVote *big.Int, finishedAt *big.Int, maximumVotePower *big.Int, options []string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SubmitProposal(&_CommunityProposal.TransactOpts, uuid_, summary_, fullContents_, proposalType, votingApp, baseNumVote, finishedAt, maximumVotePower, options)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x0359ce16.
//
// Solidity: function submitProposal(string uuid_, string summary_, string fullContents_, string proposalType, string votingApp, uint256 baseNumVote, uint256 finishedAt, uint256 maximumVotePower, string[] options) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) SubmitProposal(uuid_ string, summary_ string, fullContents_ string, proposalType string, votingApp string, baseNumVote *big.Int, finishedAt *big.Int, maximumVotePower *big.Int, options []string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.SubmitProposal(&_CommunityProposal.TransactOpts, uuid_, summary_, fullContents_, proposalType, votingApp, baseNumVote, finishedAt, maximumVotePower, options)
}

// UpdateProposalStatus is a paid mutator transaction binding the contract method 0xc327c51d.
//
// Solidity: function updateProposalStatus(string uuid) returns()
func (_CommunityProposal *CommunityProposalTransactor) UpdateProposalStatus(opts *bind.TransactOpts, uuid string) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "updateProposalStatus", uuid)
}

// UpdateProposalStatus is a paid mutator transaction binding the contract method 0xc327c51d.
//
// Solidity: function updateProposalStatus(string uuid) returns()
func (_CommunityProposal *CommunityProposalSession) UpdateProposalStatus(uuid string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.UpdateProposalStatus(&_CommunityProposal.TransactOpts, uuid)
}

// UpdateProposalStatus is a paid mutator transaction binding the contract method 0xc327c51d.
//
// Solidity: function updateProposalStatus(string uuid) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) UpdateProposalStatus(uuid string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.UpdateProposalStatus(&_CommunityProposal.TransactOpts, uuid)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_CommunityProposal *CommunityProposalTransactor) Upgrade(opts *bind.TransactOpts, newVersion common.Address) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "upgrade", newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_CommunityProposal *CommunityProposalSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _CommunityProposal.Contract.Upgrade(&_CommunityProposal.TransactOpts, newVersion)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newVersion) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) Upgrade(newVersion common.Address) (*types.Transaction, error) {
	return _CommunityProposal.Contract.Upgrade(&_CommunityProposal.TransactOpts, newVersion)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_CommunityProposal *CommunityProposalTransactor) UpgradeRegistry(opts *bind.TransactOpts, reg common.Address) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "upgradeRegistry", reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_CommunityProposal *CommunityProposalSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _CommunityProposal.Contract.UpgradeRegistry(&_CommunityProposal.TransactOpts, reg)
}

// UpgradeRegistry is a paid mutator transaction binding the contract method 0xc349e289.
//
// Solidity: function upgradeRegistry(address reg) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) UpgradeRegistry(reg common.Address) (*types.Transaction, error) {
	return _CommunityProposal.Contract.UpgradeRegistry(&_CommunityProposal.TransactOpts, reg)
}

// VoteCommunityProposal is a paid mutator transaction binding the contract method 0x0ea07810.
//
// Solidity: function voteCommunityProposal(string uuid, bool accept, uint256 value, string comments) returns()
func (_CommunityProposal *CommunityProposalTransactor) VoteCommunityProposal(opts *bind.TransactOpts, uuid string, accept bool, value *big.Int, comments string) (*types.Transaction, error) {
	return _CommunityProposal.contract.Transact(opts, "voteCommunityProposal", uuid, accept, value, comments)
}

// VoteCommunityProposal is a paid mutator transaction binding the contract method 0x0ea07810.
//
// Solidity: function voteCommunityProposal(string uuid, bool accept, uint256 value, string comments) returns()
func (_CommunityProposal *CommunityProposalSession) VoteCommunityProposal(uuid string, accept bool, value *big.Int, comments string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.VoteCommunityProposal(&_CommunityProposal.TransactOpts, uuid, accept, value, comments)
}

// VoteCommunityProposal is a paid mutator transaction binding the contract method 0x0ea07810.
//
// Solidity: function voteCommunityProposal(string uuid, bool accept, uint256 value, string comments) returns()
func (_CommunityProposal *CommunityProposalTransactorSession) VoteCommunityProposal(uuid string, accept bool, value *big.Int, comments string) (*types.Transaction, error) {
	return _CommunityProposal.Contract.VoteCommunityProposal(&_CommunityProposal.TransactOpts, uuid, accept, value, comments)
}

// CommunityProposalActionAfterFinishVoteEventIterator is returned from FilterActionAfterFinishVoteEvent and is used to iterate over the raw logs and unpacked data for ActionAfterFinishVoteEvent events raised by the CommunityProposal contract.
type CommunityProposalActionAfterFinishVoteEventIterator struct {
	Event *CommunityProposalActionAfterFinishVoteEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalActionAfterFinishVoteEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalActionAfterFinishVoteEvent)
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
		it.Event = new(CommunityProposalActionAfterFinishVoteEvent)
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
func (it *CommunityProposalActionAfterFinishVoteEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalActionAfterFinishVoteEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalActionAfterFinishVoteEvent represents a ActionAfterFinishVoteEvent event raised by the CommunityProposal contract.
type CommunityProposalActionAfterFinishVoteEvent struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionAfterFinishVoteEvent is a free log retrieval operation binding the contract event 0xe8e16a570e998c815cb483f1c4671b2c796abcd2f4e5bd0a705ce4b95c1a1936.
//
// Solidity: event actionAfterFinishVoteEvent(uint256 proposalId)
func (_CommunityProposal *CommunityProposalFilterer) FilterActionAfterFinishVoteEvent(opts *bind.FilterOpts) (*CommunityProposalActionAfterFinishVoteEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "actionAfterFinishVoteEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalActionAfterFinishVoteEventIterator{contract: _CommunityProposal.contract, event: "actionAfterFinishVoteEvent", logs: logs, sub: sub}, nil
}

// WatchActionAfterFinishVoteEvent is a free log subscription operation binding the contract event 0xe8e16a570e998c815cb483f1c4671b2c796abcd2f4e5bd0a705ce4b95c1a1936.
//
// Solidity: event actionAfterFinishVoteEvent(uint256 proposalId)
func (_CommunityProposal *CommunityProposalFilterer) WatchActionAfterFinishVoteEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalActionAfterFinishVoteEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "actionAfterFinishVoteEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalActionAfterFinishVoteEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "actionAfterFinishVoteEvent", log); err != nil {
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

// ParseActionAfterFinishVoteEvent is a log parse operation binding the contract event 0xe8e16a570e998c815cb483f1c4671b2c796abcd2f4e5bd0a705ce4b95c1a1936.
//
// Solidity: event actionAfterFinishVoteEvent(uint256 proposalId)
func (_CommunityProposal *CommunityProposalFilterer) ParseActionAfterFinishVoteEvent(log types.Log) (*CommunityProposalActionAfterFinishVoteEvent, error) {
	event := new(CommunityProposalActionAfterFinishVoteEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "actionAfterFinishVoteEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalBeforeUpgradeHookEventIterator is returned from FilterBeforeUpgradeHookEvent and is used to iterate over the raw logs and unpacked data for BeforeUpgradeHookEvent events raised by the CommunityProposal contract.
type CommunityProposalBeforeUpgradeHookEventIterator struct {
	Event *CommunityProposalBeforeUpgradeHookEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalBeforeUpgradeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalBeforeUpgradeHookEvent)
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
		it.Event = new(CommunityProposalBeforeUpgradeHookEvent)
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
func (it *CommunityProposalBeforeUpgradeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalBeforeUpgradeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalBeforeUpgradeHookEvent represents a BeforeUpgradeHookEvent event raised by the CommunityProposal contract.
type CommunityProposalBeforeUpgradeHookEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBeforeUpgradeHookEvent is a free log retrieval operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_CommunityProposal *CommunityProposalFilterer) FilterBeforeUpgradeHookEvent(opts *bind.FilterOpts) (*CommunityProposalBeforeUpgradeHookEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalBeforeUpgradeHookEventIterator{contract: _CommunityProposal.contract, event: "beforeUpgradeHookEvent", logs: logs, sub: sub}, nil
}

// WatchBeforeUpgradeHookEvent is a free log subscription operation binding the contract event 0xabb34268785307464f0f2b2cda2295a4b3e4df2103b7d09c2253247e9c4b987d.
//
// Solidity: event beforeUpgradeHookEvent(address newVersion)
func (_CommunityProposal *CommunityProposalFilterer) WatchBeforeUpgradeHookEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalBeforeUpgradeHookEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "beforeUpgradeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalBeforeUpgradeHookEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseBeforeUpgradeHookEvent(log types.Log) (*CommunityProposalBeforeUpgradeHookEvent, error) {
	event := new(CommunityProposalBeforeUpgradeHookEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "beforeUpgradeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalChangeInitialDataEventIterator is returned from FilterChangeInitialDataEvent and is used to iterate over the raw logs and unpacked data for ChangeInitialDataEvent events raised by the CommunityProposal contract.
type CommunityProposalChangeInitialDataEventIterator struct {
	Event *CommunityProposalChangeInitialDataEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalChangeInitialDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalChangeInitialDataEvent)
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
		it.Event = new(CommunityProposalChangeInitialDataEvent)
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
func (it *CommunityProposalChangeInitialDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalChangeInitialDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalChangeInitialDataEvent represents a ChangeInitialDataEvent event raised by the CommunityProposal contract.
type CommunityProposalChangeInitialDataEvent struct {
	PublicComment  bool
	EnableComment  bool
	ProposalTypes  []string
	ProposalValues []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangeInitialDataEvent is a free log retrieval operation binding the contract event 0x2db2a17c48c13c8de6f756efadf9a0a244b6b95f192b5faefc952ebd83219489.
//
// Solidity: event changeInitialDataEvent(bool publicComment, bool enableComment, string[] proposalTypes, uint256[] proposalValues)
func (_CommunityProposal *CommunityProposalFilterer) FilterChangeInitialDataEvent(opts *bind.FilterOpts) (*CommunityProposalChangeInitialDataEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalChangeInitialDataEventIterator{contract: _CommunityProposal.contract, event: "changeInitialDataEvent", logs: logs, sub: sub}, nil
}

// WatchChangeInitialDataEvent is a free log subscription operation binding the contract event 0x2db2a17c48c13c8de6f756efadf9a0a244b6b95f192b5faefc952ebd83219489.
//
// Solidity: event changeInitialDataEvent(bool publicComment, bool enableComment, string[] proposalTypes, uint256[] proposalValues)
func (_CommunityProposal *CommunityProposalFilterer) WatchChangeInitialDataEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalChangeInitialDataEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "changeInitialDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalChangeInitialDataEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
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

// ParseChangeInitialDataEvent is a log parse operation binding the contract event 0x2db2a17c48c13c8de6f756efadf9a0a244b6b95f192b5faefc952ebd83219489.
//
// Solidity: event changeInitialDataEvent(bool publicComment, bool enableComment, string[] proposalTypes, uint256[] proposalValues)
func (_CommunityProposal *CommunityProposalFilterer) ParseChangeInitialDataEvent(log types.Log) (*CommunityProposalChangeInitialDataEvent, error) {
	event := new(CommunityProposalChangeInitialDataEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "changeInitialDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalChangeProposalTypeEvnetIterator is returned from FilterChangeProposalTypeEvnet and is used to iterate over the raw logs and unpacked data for ChangeProposalTypeEvnet events raised by the CommunityProposal contract.
type CommunityProposalChangeProposalTypeEvnetIterator struct {
	Event *CommunityProposalChangeProposalTypeEvnet // Event containing the contract specifics and raw log

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
func (it *CommunityProposalChangeProposalTypeEvnetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalChangeProposalTypeEvnet)
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
		it.Event = new(CommunityProposalChangeProposalTypeEvnet)
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
func (it *CommunityProposalChangeProposalTypeEvnetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalChangeProposalTypeEvnetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalChangeProposalTypeEvnet represents a ChangeProposalTypeEvnet event raised by the CommunityProposal contract.
type CommunityProposalChangeProposalTypeEvnet struct {
	Proposaltypes []string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChangeProposalTypeEvnet is a free log retrieval operation binding the contract event 0x483e3e7fbac6aa08df034442c9b7952b9449aa337b6dd66827570fdb57125087.
//
// Solidity: event changeProposalTypeEvnet(string[] proposaltypes)
func (_CommunityProposal *CommunityProposalFilterer) FilterChangeProposalTypeEvnet(opts *bind.FilterOpts) (*CommunityProposalChangeProposalTypeEvnetIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "changeProposalTypeEvnet")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalChangeProposalTypeEvnetIterator{contract: _CommunityProposal.contract, event: "changeProposalTypeEvnet", logs: logs, sub: sub}, nil
}

// WatchChangeProposalTypeEvnet is a free log subscription operation binding the contract event 0x483e3e7fbac6aa08df034442c9b7952b9449aa337b6dd66827570fdb57125087.
//
// Solidity: event changeProposalTypeEvnet(string[] proposaltypes)
func (_CommunityProposal *CommunityProposalFilterer) WatchChangeProposalTypeEvnet(opts *bind.WatchOpts, sink chan<- *CommunityProposalChangeProposalTypeEvnet) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "changeProposalTypeEvnet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalChangeProposalTypeEvnet)
				if err := _CommunityProposal.contract.UnpackLog(event, "changeProposalTypeEvnet", log); err != nil {
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

// ParseChangeProposalTypeEvnet is a log parse operation binding the contract event 0x483e3e7fbac6aa08df034442c9b7952b9449aa337b6dd66827570fdb57125087.
//
// Solidity: event changeProposalTypeEvnet(string[] proposaltypes)
func (_CommunityProposal *CommunityProposalFilterer) ParseChangeProposalTypeEvnet(log types.Log) (*CommunityProposalChangeProposalTypeEvnet, error) {
	event := new(CommunityProposalChangeProposalTypeEvnet)
	if err := _CommunityProposal.contract.UnpackLog(event, "changeProposalTypeEvnet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalEmitActivityEventIterator is returned from FilterEmitActivityEvent and is used to iterate over the raw logs and unpacked data for EmitActivityEvent events raised by the CommunityProposal contract.
type CommunityProposalEmitActivityEventIterator struct {
	Event *CommunityProposalEmitActivityEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalEmitActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalEmitActivityEvent)
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
		it.Event = new(CommunityProposalEmitActivityEvent)
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
func (it *CommunityProposalEmitActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalEmitActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalEmitActivityEvent represents a EmitActivityEvent event raised by the CommunityProposal contract.
type CommunityProposalEmitActivityEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitActivityEvent is a free log retrieval operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_CommunityProposal *CommunityProposalFilterer) FilterEmitActivityEvent(opts *bind.FilterOpts) (*CommunityProposalEmitActivityEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalEmitActivityEventIterator{contract: _CommunityProposal.contract, event: "emitActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitActivityEvent is a free log subscription operation binding the contract event 0x7249c2b0e62796f46f0c3adaf23c3822033e8e31d1ebd7b83b78f89fa9374568.
//
// Solidity: event emitActivityEvent(uint256 activityFlag, bytes data)
func (_CommunityProposal *CommunityProposalFilterer) WatchEmitActivityEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalEmitActivityEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "emitActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalEmitActivityEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseEmitActivityEvent(log types.Log) (*CommunityProposalEmitActivityEvent, error) {
	event := new(CommunityProposalEmitActivityEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "emitActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalEmitAppDataEventIterator is returned from FilterEmitAppDataEvent and is used to iterate over the raw logs and unpacked data for EmitAppDataEvent events raised by the CommunityProposal contract.
type CommunityProposalEmitAppDataEventIterator struct {
	Event *CommunityProposalEmitAppDataEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalEmitAppDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalEmitAppDataEvent)
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
		it.Event = new(CommunityProposalEmitAppDataEvent)
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
func (it *CommunityProposalEmitAppDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalEmitAppDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalEmitAppDataEvent represents a EmitAppDataEvent event raised by the CommunityProposal contract.
type CommunityProposalEmitAppDataEvent struct {
	ActivityFlag *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitAppDataEvent is a free log retrieval operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_CommunityProposal *CommunityProposalFilterer) FilterEmitAppDataEvent(opts *bind.FilterOpts) (*CommunityProposalEmitAppDataEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalEmitAppDataEventIterator{contract: _CommunityProposal.contract, event: "emitAppDataEvent", logs: logs, sub: sub}, nil
}

// WatchEmitAppDataEvent is a free log subscription operation binding the contract event 0x821f268b753e7e702b1f95e4b73107c90f40b13f7ab2531db1d07d5977e3aaab.
//
// Solidity: event emitAppDataEvent(uint256 activityFlag, bytes data)
func (_CommunityProposal *CommunityProposalFilterer) WatchEmitAppDataEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalEmitAppDataEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "emitAppDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalEmitAppDataEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseEmitAppDataEvent(log types.Log) (*CommunityProposalEmitAppDataEvent, error) {
	event := new(CommunityProposalEmitAppDataEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "emitAppDataEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalEmitProposalActivityEventIterator is returned from FilterEmitProposalActivityEvent and is used to iterate over the raw logs and unpacked data for EmitProposalActivityEvent events raised by the CommunityProposal contract.
type CommunityProposalEmitProposalActivityEventIterator struct {
	Event *CommunityProposalEmitProposalActivityEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalEmitProposalActivityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalEmitProposalActivityEvent)
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
		it.Event = new(CommunityProposalEmitProposalActivityEvent)
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
func (it *CommunityProposalEmitProposalActivityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalEmitProposalActivityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalEmitProposalActivityEvent represents a EmitProposalActivityEvent event raised by the CommunityProposal contract.
type CommunityProposalEmitProposalActivityEvent struct {
	ActivityType *big.Int
	Proposal     ProposalSummary
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmitProposalActivityEvent is a free log retrieval operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_CommunityProposal *CommunityProposalFilterer) FilterEmitProposalActivityEvent(opts *bind.FilterOpts) (*CommunityProposalEmitProposalActivityEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalEmitProposalActivityEventIterator{contract: _CommunityProposal.contract, event: "emitProposalActivityEvent", logs: logs, sub: sub}, nil
}

// WatchEmitProposalActivityEvent is a free log subscription operation binding the contract event 0x9706d432aa2f8e2c9c60af85d6675ffbf994a7e81896bc00983868797946128c.
//
// Solidity: event emitProposalActivityEvent(uint256 activityType, (uint256,address,string,string,string,string,uint256,uint16,uint256) proposal)
func (_CommunityProposal *CommunityProposalFilterer) WatchEmitProposalActivityEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalEmitProposalActivityEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "emitProposalActivityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalEmitProposalActivityEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseEmitProposalActivityEvent(log types.Log) (*CommunityProposalEmitProposalActivityEvent, error) {
	event := new(CommunityProposalEmitProposalActivityEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "emitProposalActivityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the CommunityProposal contract.
type CommunityProposalInitializeEventIterator struct {
	Event *CommunityProposalInitializeEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalInitializeEvent)
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
		it.Event = new(CommunityProposalInitializeEvent)
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
func (it *CommunityProposalInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalInitializeEvent represents a InitializeEvent event raised by the CommunityProposal contract.
type CommunityProposalInitializeEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_CommunityProposal *CommunityProposalFilterer) FilterInitializeEvent(opts *bind.FilterOpts) (*CommunityProposalInitializeEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalInitializeEventIterator{contract: _CommunityProposal.contract, event: "initializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0x130f519910ef25785b691c5ab683158b7b5660bd72fde72fca865e0251bac414.
//
// Solidity: event initializeEvent(address[] addrs)
func (_CommunityProposal *CommunityProposalFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalInitializeEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "initializeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalInitializeEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "initializeEvent", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseInitializeEvent(log types.Log) (*CommunityProposalInitializeEvent, error) {
	event := new(CommunityProposalInitializeEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "initializeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalInitializeHookEventIterator is returned from FilterInitializeHookEvent and is used to iterate over the raw logs and unpacked data for InitializeHookEvent events raised by the CommunityProposal contract.
type CommunityProposalInitializeHookEventIterator struct {
	Event *CommunityProposalInitializeHookEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalInitializeHookEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalInitializeHookEvent)
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
		it.Event = new(CommunityProposalInitializeHookEvent)
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
func (it *CommunityProposalInitializeHookEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalInitializeHookEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalInitializeHookEvent represents a InitializeHookEvent event raised by the CommunityProposal contract.
type CommunityProposalInitializeHookEvent struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializeHookEvent is a free log retrieval operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_CommunityProposal *CommunityProposalFilterer) FilterInitializeHookEvent(opts *bind.FilterOpts) (*CommunityProposalInitializeHookEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalInitializeHookEventIterator{contract: _CommunityProposal.contract, event: "initializeHookEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeHookEvent is a free log subscription operation binding the contract event 0x8906ba50e434fc1c744ae930f96960ce691ab1d380bf737d626a8d707dd28238.
//
// Solidity: event initializeHookEvent(address[] addrs)
func (_CommunityProposal *CommunityProposalFilterer) WatchInitializeHookEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalInitializeHookEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "initializeHookEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalInitializeHookEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseInitializeHookEvent(log types.Log) (*CommunityProposalInitializeHookEvent, error) {
	event := new(CommunityProposalInitializeHookEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "initializeHookEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalSetCommentEventIterator is returned from FilterSetCommentEvent and is used to iterate over the raw logs and unpacked data for SetCommentEvent events raised by the CommunityProposal contract.
type CommunityProposalSetCommentEventIterator struct {
	Event *CommunityProposalSetCommentEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalSetCommentEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalSetCommentEvent)
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
		it.Event = new(CommunityProposalSetCommentEvent)
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
func (it *CommunityProposalSetCommentEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalSetCommentEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalSetCommentEvent represents a SetCommentEvent event raised by the CommunityProposal contract.
type CommunityProposalSetCommentEvent struct {
	Uuid     string
	Contents string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetCommentEvent is a free log retrieval operation binding the contract event 0xa281126857901973627e8850dc48b9bdf42ae5157d9d7f88e1d197bc71b02ca8.
//
// Solidity: event setCommentEvent(string uuid_, string contents_)
func (_CommunityProposal *CommunityProposalFilterer) FilterSetCommentEvent(opts *bind.FilterOpts) (*CommunityProposalSetCommentEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "setCommentEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalSetCommentEventIterator{contract: _CommunityProposal.contract, event: "setCommentEvent", logs: logs, sub: sub}, nil
}

// WatchSetCommentEvent is a free log subscription operation binding the contract event 0xa281126857901973627e8850dc48b9bdf42ae5157d9d7f88e1d197bc71b02ca8.
//
// Solidity: event setCommentEvent(string uuid_, string contents_)
func (_CommunityProposal *CommunityProposalFilterer) WatchSetCommentEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalSetCommentEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "setCommentEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalSetCommentEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "setCommentEvent", log); err != nil {
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

// ParseSetCommentEvent is a log parse operation binding the contract event 0xa281126857901973627e8850dc48b9bdf42ae5157d9d7f88e1d197bc71b02ca8.
//
// Solidity: event setCommentEvent(string uuid_, string contents_)
func (_CommunityProposal *CommunityProposalFilterer) ParseSetCommentEvent(log types.Log) (*CommunityProposalSetCommentEvent, error) {
	event := new(CommunityProposalSetCommentEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "setCommentEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalSetCommentV2EventIterator is returned from FilterSetCommentV2Event and is used to iterate over the raw logs and unpacked data for SetCommentV2Event events raised by the CommunityProposal contract.
type CommunityProposalSetCommentV2EventIterator struct {
	Event *CommunityProposalSetCommentV2Event // Event containing the contract specifics and raw log

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
func (it *CommunityProposalSetCommentV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalSetCommentV2Event)
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
		it.Event = new(CommunityProposalSetCommentV2Event)
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
func (it *CommunityProposalSetCommentV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalSetCommentV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalSetCommentV2Event represents a SetCommentV2Event event raised by the CommunityProposal contract.
type CommunityProposalSetCommentV2Event struct {
	ProposalId *big.Int
	Contents   string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetCommentV2Event is a free log retrieval operation binding the contract event 0x522c02a42eaa2e08cd549b81c0c94e432e8e0e7332d2852c3bf9793cdf7b3975.
//
// Solidity: event setCommentV2Event(uint256 proposalId, string contents_)
func (_CommunityProposal *CommunityProposalFilterer) FilterSetCommentV2Event(opts *bind.FilterOpts) (*CommunityProposalSetCommentV2EventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "setCommentV2Event")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalSetCommentV2EventIterator{contract: _CommunityProposal.contract, event: "setCommentV2Event", logs: logs, sub: sub}, nil
}

// WatchSetCommentV2Event is a free log subscription operation binding the contract event 0x522c02a42eaa2e08cd549b81c0c94e432e8e0e7332d2852c3bf9793cdf7b3975.
//
// Solidity: event setCommentV2Event(uint256 proposalId, string contents_)
func (_CommunityProposal *CommunityProposalFilterer) WatchSetCommentV2Event(opts *bind.WatchOpts, sink chan<- *CommunityProposalSetCommentV2Event) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "setCommentV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalSetCommentV2Event)
				if err := _CommunityProposal.contract.UnpackLog(event, "setCommentV2Event", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseSetCommentV2Event(log types.Log) (*CommunityProposalSetCommentV2Event, error) {
	event := new(CommunityProposalSetCommentV2Event)
	if err := _CommunityProposal.contract.UnpackLog(event, "setCommentV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalSetProposalEventIterator is returned from FilterSetProposalEvent and is used to iterate over the raw logs and unpacked data for SetProposalEvent events raised by the CommunityProposal contract.
type CommunityProposalSetProposalEventIterator struct {
	Event *CommunityProposalSetProposalEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalSetProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalSetProposalEvent)
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
		it.Event = new(CommunityProposalSetProposalEvent)
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
func (it *CommunityProposalSetProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalSetProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalSetProposalEvent represents a SetProposalEvent event raised by the CommunityProposal contract.
type CommunityProposalSetProposalEvent struct {
	ProposalId *big.Int
	Pro        ProposalParameterV2
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetProposalEvent is a free log retrieval operation binding the contract event 0x48e7e67d793f172da13ca586476a4961735c1252978a42d0afa1b3f67f330fc0.
//
// Solidity: event setProposalEvent(uint256 proposalId, (uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256) pro)
func (_CommunityProposal *CommunityProposalFilterer) FilterSetProposalEvent(opts *bind.FilterOpts) (*CommunityProposalSetProposalEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "setProposalEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalSetProposalEventIterator{contract: _CommunityProposal.contract, event: "setProposalEvent", logs: logs, sub: sub}, nil
}

// WatchSetProposalEvent is a free log subscription operation binding the contract event 0x48e7e67d793f172da13ca586476a4961735c1252978a42d0afa1b3f67f330fc0.
//
// Solidity: event setProposalEvent(uint256 proposalId, (uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256) pro)
func (_CommunityProposal *CommunityProposalFilterer) WatchSetProposalEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalSetProposalEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "setProposalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalSetProposalEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "setProposalEvent", log); err != nil {
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

// ParseSetProposalEvent is a log parse operation binding the contract event 0x48e7e67d793f172da13ca586476a4961735c1252978a42d0afa1b3f67f330fc0.
//
// Solidity: event setProposalEvent(uint256 proposalId, (uint8,address,string,bool,bool,uint256,uint256,uint256,uint256,uint256) pro)
func (_CommunityProposal *CommunityProposalFilterer) ParseSetProposalEvent(log types.Log) (*CommunityProposalSetProposalEvent, error) {
	event := new(CommunityProposalSetProposalEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "setProposalEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalSetReadyEventIterator is returned from FilterSetReadyEvent and is used to iterate over the raw logs and unpacked data for SetReadyEvent events raised by the CommunityProposal contract.
type CommunityProposalSetReadyEventIterator struct {
	Event *CommunityProposalSetReadyEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalSetReadyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalSetReadyEvent)
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
		it.Event = new(CommunityProposalSetReadyEvent)
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
func (it *CommunityProposalSetReadyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalSetReadyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalSetReadyEvent represents a SetReadyEvent event raised by the CommunityProposal contract.
type CommunityProposalSetReadyEvent struct {
	Done bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetReadyEvent is a free log retrieval operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_CommunityProposal *CommunityProposalFilterer) FilterSetReadyEvent(opts *bind.FilterOpts) (*CommunityProposalSetReadyEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalSetReadyEventIterator{contract: _CommunityProposal.contract, event: "setReadyEvent", logs: logs, sub: sub}, nil
}

// WatchSetReadyEvent is a free log subscription operation binding the contract event 0xef4ad5c5d8d806adca5a37b004e802d773a83a01b26294b4c0346da5c4145bfe.
//
// Solidity: event setReadyEvent(bool done)
func (_CommunityProposal *CommunityProposalFilterer) WatchSetReadyEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalSetReadyEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "setReadyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalSetReadyEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseSetReadyEvent(log types.Log) (*CommunityProposalSetReadyEvent, error) {
	event := new(CommunityProposalSetReadyEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "setReadyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalSubmitCommunityProposalEventIterator is returned from FilterSubmitCommunityProposalEvent and is used to iterate over the raw logs and unpacked data for SubmitCommunityProposalEvent events raised by the CommunityProposal contract.
type CommunityProposalSubmitCommunityProposalEventIterator struct {
	Event *CommunityProposalSubmitCommunityProposalEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalSubmitCommunityProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalSubmitCommunityProposalEvent)
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
		it.Event = new(CommunityProposalSubmitCommunityProposalEvent)
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
func (it *CommunityProposalSubmitCommunityProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalSubmitCommunityProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalSubmitCommunityProposalEvent represents a SubmitCommunityProposalEvent event raised by the CommunityProposal contract.
type CommunityProposalSubmitCommunityProposalEvent struct {
	Uuid         string
	Summary      string
	FullContents string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSubmitCommunityProposalEvent is a free log retrieval operation binding the contract event 0x551e1165d121823bcaa8e0d2c2cde8103aacd27c581d769d39b52729de4947af.
//
// Solidity: event submitCommunityProposalEvent(string uuid_, string summary_, string fullContents_)
func (_CommunityProposal *CommunityProposalFilterer) FilterSubmitCommunityProposalEvent(opts *bind.FilterOpts) (*CommunityProposalSubmitCommunityProposalEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "submitCommunityProposalEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalSubmitCommunityProposalEventIterator{contract: _CommunityProposal.contract, event: "submitCommunityProposalEvent", logs: logs, sub: sub}, nil
}

// WatchSubmitCommunityProposalEvent is a free log subscription operation binding the contract event 0x551e1165d121823bcaa8e0d2c2cde8103aacd27c581d769d39b52729de4947af.
//
// Solidity: event submitCommunityProposalEvent(string uuid_, string summary_, string fullContents_)
func (_CommunityProposal *CommunityProposalFilterer) WatchSubmitCommunityProposalEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalSubmitCommunityProposalEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "submitCommunityProposalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalSubmitCommunityProposalEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "submitCommunityProposalEvent", log); err != nil {
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

// ParseSubmitCommunityProposalEvent is a log parse operation binding the contract event 0x551e1165d121823bcaa8e0d2c2cde8103aacd27c581d769d39b52729de4947af.
//
// Solidity: event submitCommunityProposalEvent(string uuid_, string summary_, string fullContents_)
func (_CommunityProposal *CommunityProposalFilterer) ParseSubmitCommunityProposalEvent(log types.Log) (*CommunityProposalSubmitCommunityProposalEvent, error) {
	event := new(CommunityProposalSubmitCommunityProposalEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "submitCommunityProposalEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalSubmitCommunityProposalV2EventIterator is returned from FilterSubmitCommunityProposalV2Event and is used to iterate over the raw logs and unpacked data for SubmitCommunityProposalV2Event events raised by the CommunityProposal contract.
type CommunityProposalSubmitCommunityProposalV2EventIterator struct {
	Event *CommunityProposalSubmitCommunityProposalV2Event // Event containing the contract specifics and raw log

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
func (it *CommunityProposalSubmitCommunityProposalV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalSubmitCommunityProposalV2Event)
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
		it.Event = new(CommunityProposalSubmitCommunityProposalV2Event)
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
func (it *CommunityProposalSubmitCommunityProposalV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalSubmitCommunityProposalV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalSubmitCommunityProposalV2Event represents a SubmitCommunityProposalV2Event event raised by the CommunityProposal contract.
type CommunityProposalSubmitCommunityProposalV2Event struct {
	Uuid             string
	Summary          string
	FullContents     string
	ProposalType     string
	VotingApp        string
	FinishedAt       *big.Int
	MaximumVotePower *big.Int
	Options          []string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSubmitCommunityProposalV2Event is a free log retrieval operation binding the contract event 0xd02c3aa11a773c862f0b7f09947c1a0fed70d92f84bf66b703f0b82f99e90f9a.
//
// Solidity: event submitCommunityProposalV2Event(string uuid_, string summary_, string fullContents_, string proposalType, string votingApp, uint256 finishedAt, uint256 maximumVotePower, string[] options)
func (_CommunityProposal *CommunityProposalFilterer) FilterSubmitCommunityProposalV2Event(opts *bind.FilterOpts) (*CommunityProposalSubmitCommunityProposalV2EventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "submitCommunityProposalV2Event")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalSubmitCommunityProposalV2EventIterator{contract: _CommunityProposal.contract, event: "submitCommunityProposalV2Event", logs: logs, sub: sub}, nil
}

// WatchSubmitCommunityProposalV2Event is a free log subscription operation binding the contract event 0xd02c3aa11a773c862f0b7f09947c1a0fed70d92f84bf66b703f0b82f99e90f9a.
//
// Solidity: event submitCommunityProposalV2Event(string uuid_, string summary_, string fullContents_, string proposalType, string votingApp, uint256 finishedAt, uint256 maximumVotePower, string[] options)
func (_CommunityProposal *CommunityProposalFilterer) WatchSubmitCommunityProposalV2Event(opts *bind.WatchOpts, sink chan<- *CommunityProposalSubmitCommunityProposalV2Event) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "submitCommunityProposalV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalSubmitCommunityProposalV2Event)
				if err := _CommunityProposal.contract.UnpackLog(event, "submitCommunityProposalV2Event", log); err != nil {
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

// ParseSubmitCommunityProposalV2Event is a log parse operation binding the contract event 0xd02c3aa11a773c862f0b7f09947c1a0fed70d92f84bf66b703f0b82f99e90f9a.
//
// Solidity: event submitCommunityProposalV2Event(string uuid_, string summary_, string fullContents_, string proposalType, string votingApp, uint256 finishedAt, uint256 maximumVotePower, string[] options)
func (_CommunityProposal *CommunityProposalFilterer) ParseSubmitCommunityProposalV2Event(log types.Log) (*CommunityProposalSubmitCommunityProposalV2Event, error) {
	event := new(CommunityProposalSubmitCommunityProposalV2Event)
	if err := _CommunityProposal.contract.UnpackLog(event, "submitCommunityProposalV2Event", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalUpgradeEventIterator is returned from FilterUpgradeEvent and is used to iterate over the raw logs and unpacked data for UpgradeEvent events raised by the CommunityProposal contract.
type CommunityProposalUpgradeEventIterator struct {
	Event *CommunityProposalUpgradeEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalUpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalUpgradeEvent)
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
		it.Event = new(CommunityProposalUpgradeEvent)
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
func (it *CommunityProposalUpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalUpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalUpgradeEvent represents a UpgradeEvent event raised by the CommunityProposal contract.
type CommunityProposalUpgradeEvent struct {
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpgradeEvent is a free log retrieval operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_CommunityProposal *CommunityProposalFilterer) FilterUpgradeEvent(opts *bind.FilterOpts) (*CommunityProposalUpgradeEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalUpgradeEventIterator{contract: _CommunityProposal.contract, event: "upgradeEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeEvent is a free log subscription operation binding the contract event 0x168ee9ab85e117bd3b0065af51eedfc1884a638a66152d4e79b2e059d705dda7.
//
// Solidity: event upgradeEvent(address newVersion)
func (_CommunityProposal *CommunityProposalFilterer) WatchUpgradeEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalUpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "upgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalUpgradeEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseUpgradeEvent(log types.Log) (*CommunityProposalUpgradeEvent, error) {
	event := new(CommunityProposalUpgradeEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "upgradeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalUpgradeRegistryEventIterator is returned from FilterUpgradeRegistryEvent and is used to iterate over the raw logs and unpacked data for UpgradeRegistryEvent events raised by the CommunityProposal contract.
type CommunityProposalUpgradeRegistryEventIterator struct {
	Event *CommunityProposalUpgradeRegistryEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalUpgradeRegistryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalUpgradeRegistryEvent)
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
		it.Event = new(CommunityProposalUpgradeRegistryEvent)
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
func (it *CommunityProposalUpgradeRegistryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalUpgradeRegistryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalUpgradeRegistryEvent represents a UpgradeRegistryEvent event raised by the CommunityProposal contract.
type CommunityProposalUpgradeRegistryEvent struct {
	Reg common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpgradeRegistryEvent is a free log retrieval operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_CommunityProposal *CommunityProposalFilterer) FilterUpgradeRegistryEvent(opts *bind.FilterOpts) (*CommunityProposalUpgradeRegistryEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalUpgradeRegistryEventIterator{contract: _CommunityProposal.contract, event: "upgradeRegistryEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeRegistryEvent is a free log subscription operation binding the contract event 0x5f547d7b14a389802a72a60aac756fb3cddae162c923210d779a2bc8947f704d.
//
// Solidity: event upgradeRegistryEvent(address reg)
func (_CommunityProposal *CommunityProposalFilterer) WatchUpgradeRegistryEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalUpgradeRegistryEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "upgradeRegistryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalUpgradeRegistryEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
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
func (_CommunityProposal *CommunityProposalFilterer) ParseUpgradeRegistryEvent(log types.Log) (*CommunityProposalUpgradeRegistryEvent, error) {
	event := new(CommunityProposalUpgradeRegistryEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "upgradeRegistryEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CommunityProposalVoteCommunityProposalEventIterator is returned from FilterVoteCommunityProposalEvent and is used to iterate over the raw logs and unpacked data for VoteCommunityProposalEvent events raised by the CommunityProposal contract.
type CommunityProposalVoteCommunityProposalEventIterator struct {
	Event *CommunityProposalVoteCommunityProposalEvent // Event containing the contract specifics and raw log

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
func (it *CommunityProposalVoteCommunityProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityProposalVoteCommunityProposalEvent)
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
		it.Event = new(CommunityProposalVoteCommunityProposalEvent)
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
func (it *CommunityProposalVoteCommunityProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityProposalVoteCommunityProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityProposalVoteCommunityProposalEvent represents a VoteCommunityProposalEvent event raised by the CommunityProposal contract.
type CommunityProposalVoteCommunityProposalEvent struct {
	Uuid     string
	Accept   bool
	Value    *big.Int
	Comments string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoteCommunityProposalEvent is a free log retrieval operation binding the contract event 0x4b3d6d4ecb28cfaa303f862355aa36cefdcabf3d592c9a58fe7b838661d36e31.
//
// Solidity: event voteCommunityProposalEvent(string uuid, bool accept, uint256 value, string comments)
func (_CommunityProposal *CommunityProposalFilterer) FilterVoteCommunityProposalEvent(opts *bind.FilterOpts) (*CommunityProposalVoteCommunityProposalEventIterator, error) {

	logs, sub, err := _CommunityProposal.contract.FilterLogs(opts, "voteCommunityProposalEvent")
	if err != nil {
		return nil, err
	}
	return &CommunityProposalVoteCommunityProposalEventIterator{contract: _CommunityProposal.contract, event: "voteCommunityProposalEvent", logs: logs, sub: sub}, nil
}

// WatchVoteCommunityProposalEvent is a free log subscription operation binding the contract event 0x4b3d6d4ecb28cfaa303f862355aa36cefdcabf3d592c9a58fe7b838661d36e31.
//
// Solidity: event voteCommunityProposalEvent(string uuid, bool accept, uint256 value, string comments)
func (_CommunityProposal *CommunityProposalFilterer) WatchVoteCommunityProposalEvent(opts *bind.WatchOpts, sink chan<- *CommunityProposalVoteCommunityProposalEvent) (event.Subscription, error) {

	logs, sub, err := _CommunityProposal.contract.WatchLogs(opts, "voteCommunityProposalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityProposalVoteCommunityProposalEvent)
				if err := _CommunityProposal.contract.UnpackLog(event, "voteCommunityProposalEvent", log); err != nil {
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

// ParseVoteCommunityProposalEvent is a log parse operation binding the contract event 0x4b3d6d4ecb28cfaa303f862355aa36cefdcabf3d592c9a58fe7b838661d36e31.
//
// Solidity: event voteCommunityProposalEvent(string uuid, bool accept, uint256 value, string comments)
func (_CommunityProposal *CommunityProposalFilterer) ParseVoteCommunityProposalEvent(log types.Log) (*CommunityProposalVoteCommunityProposalEvent, error) {
	event := new(CommunityProposalVoteCommunityProposalEvent)
	if err := _CommunityProposal.contract.UnpackLog(event, "voteCommunityProposalEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
