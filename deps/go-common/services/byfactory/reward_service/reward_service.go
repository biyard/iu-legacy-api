package byfactory

import (
	"context"

	"biyard.co/common/base"
	rewardContract "biyard.co/common/contracts/reward"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
)

type RewardService struct {
	clientEth      *client.Client
	klaytnEndpoint string
	senderAddress  string
}

func NewRewardService(endpoint, sender string) *RewardService {
	clientEth, clientErr := client.Dial(endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	return &RewardService{
		clientEth:      clientEth,
		klaytnEndpoint: endpoint,
		senderAddress:  sender,
	}
}

func (r *RewardService) GetStandardMTAddress(context *base.Context, rewardAddress common.Address) (common.Address, error) {
	rewardInstance, rewardErr := rewardContract.NewReward(rewardAddress, r.clientEth)

	if rewardErr != nil {
		return common.HexToAddress("0x0000000000000000000000000000000000000000"), rewardErr
	}

	standardMtAddr, mtErr := rewardInstance.GetStandardMTAddress(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	})

	if mtErr != nil {
		return common.HexToAddress("0x0000000000000000000000000000000000000000"), mtErr
	}

	return standardMtAddr, nil
}
