package byfactory

import (
	"context"
	"fmt"

	"biyard.co/common/base"
	standardTokenContract "biyard.co/common/contracts/standard-token"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
)

type StandardTokenService struct {
	clientEth      *client.Client
	klaytnEndpoint string
	senderAddress  string
}

func NewStandardTokenService(endpoint, sender string) *StandardTokenService {
	clientEth, clientErr := client.Dial(endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	return &StandardTokenService{
		clientEth:      clientEth,
		klaytnEndpoint: endpoint,
		senderAddress:  sender,
	}
}

func (r *StandardTokenService) GetBalance(context *base.Context, standardTokenAddress common.Address) (string, error) {
	standardTokenInstance, standardTokenErr := standardTokenContract.NewStandardToken(standardTokenAddress, r.clientEth)

	if standardTokenErr != nil {
		return "", standardTokenErr
	}

	dt, err := standardTokenInstance.BalanceOf(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	}, standardTokenAddress)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", dt), nil
}
