package byfactory

import (
	"context"
	"fmt"

	"biyard.co/common/base"
	communityContract "biyard.co/common/contracts/community"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
)

type CommunityService struct {
	clientEth      *client.Client
	klaytnEndpoint string
	senderAddress  string
}

func NewCommunityService(endpoint, sender string) *CommunityService {
	clientEth, clientErr := client.Dial(endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	return &CommunityService{
		clientEth:      clientEth,
		klaytnEndpoint: endpoint,
		senderAddress:  sender,
	}
}

func (r *CommunityService) GetProposalType(context *base.Context, communityAddress common.Address) ([]string, []string, error) {
	communityInstance, communityErr := communityContract.NewCommunityProposal(communityAddress, r.clientEth)

	if communityErr != nil {
		return nil, nil, communityErr
	}

	types, values, err := communityInstance.ProposalTypeList(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	})
	if err != nil {
		return nil, nil, err
	}

	typeValues := []string{}

	for _, v := range values {
		val := fmt.Sprintf("%s", v)
		typeValues = append(typeValues, val)
	}

	return types, typeValues, nil
}
