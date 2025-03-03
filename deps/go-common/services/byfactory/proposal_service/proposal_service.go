package byfactory

import (
	"context"

	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/client"
)

type ProposalService struct {
	clientEth      *client.Client
	klaytnEndpoint string
	senderAddress  string
}

func NewProposalService(endpoint, sender string) *ProposalService {
	clientEth, clientErr := client.Dial(endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	return &ProposalService{
		clientEth:      clientEth,
		klaytnEndpoint: endpoint,
		senderAddress:  sender,
	}
}
