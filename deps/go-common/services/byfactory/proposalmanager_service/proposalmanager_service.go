package byfactory

import (
	"context"
	"fmt"
	"math/big"

	"biyard.co/common/base"
	proposalContract "biyard.co/common/contracts/proposal"
	proposalManagerContract "biyard.co/common/contracts/proposal-manager"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
)

type ProposalManagerService struct {
	clientEth      *client.Client
	klaytnEndpoint string
	senderAddress  string
}

func NewProposalManagerService(endpoint, sender string) *ProposalManagerService {
	clientEth, clientErr := client.Dial(endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	return &ProposalManagerService{
		clientEth:      clientEth,
		klaytnEndpoint: endpoint,
		senderAddress:  sender,
	}
}

func (r *ProposalManagerService) GetPolicies(context *base.Context, proposalManagerAddress common.Address) ([]string, []string, error) {
	proposalManagerInstance, proposalManagerErr := proposalManagerContract.NewProposalManagerAppV2(proposalManagerAddress, r.clientEth)

	if proposalManagerErr != nil {
		return nil, nil, proposalManagerErr
	}

	policies, policiesErr := proposalManagerInstance.GetPolicies(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	})
	if policiesErr != nil {
		return nil, nil, policiesErr
	}

	policyValues := []string{}

	for _, v := range policies {
		value, err := proposalManagerInstance.GetPolicy(&bind.CallOpts{
			From: common.HexToAddress(r.senderAddress),
		}, v)
		if err != nil {
			return nil, nil, err
		}

		strValue := fmt.Sprintf("%s", value)
		policyValues = append(policyValues, strValue)
	}

	logger.Debugf(context, "policies: %+v policyValues: %+v", policies, policyValues)

	return policies, policyValues, nil
}

func (r *ProposalManagerService) GetVoterValue(context *base.Context, proposalAddress common.Address, proposalManagerAddress common.Address, voter common.Address, proposalId *big.Int) (string, string, error) {
	proposalInstance, proposalErr := proposalContract.NewProposal(proposalAddress, r.clientEth)
	proposalManagerInstance, proposalManagerErr := proposalManagerContract.NewProposalManagerAppV2(proposalManagerAddress, r.clientEth)

	if proposalErr != nil {
		return "", "", proposalErr
	}

	if proposalManagerErr != nil {
		return "", "", proposalManagerErr
	}

	voterValue, err := proposalInstance.GetVoterValueV2(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	}, voter, proposalId)
	if err != nil {
		return "", "", err
	}

	participation := "none"

	check, addrs, err := proposalManagerInstance.GetAcceptersV2(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	}, proposalId)
	if err != nil {
		return "", "", err
	}

	if check {
		isAccepter := false
		for _, v := range addrs {
			if v == voter {
				isAccepter = true
				break
			}
		}

		if isAccepter {
			participation = "accepter"
		} else {
			_, addrs, err := proposalManagerInstance.GetRejectersV2(&bind.CallOpts{
				From: common.HexToAddress(r.senderAddress),
			}, proposalId)
			if err != nil {
				// return nil, nil, err
			}

			isRejecter := false

			for _, v := range addrs {
				if v == voter {
					isRejecter = true
					break
				}
			}

			if isRejecter {
				participation = "rejecter"
			}
		}
	} else {
		participation = "disableSeeVote"
	}

	value := fmt.Sprintf("%s", voterValue)

	return value, participation, nil
}
