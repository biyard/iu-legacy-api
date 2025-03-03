package byfactory

import (
	"context"

	"biyard.co/common/base"
	factorycontract "biyard.co/common/contracts/dao-factory"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
)

type DaoFactoryService struct {
	clientEth      *client.Client
	klaytnEndpoint string
	senderAddress  string
}

func NewDaoFactoryService(endpoint, sender string) *DaoFactoryService {
	clientEth, clientErr := client.Dial(endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	return &DaoFactoryService{
		clientEth:      clientEth,
		klaytnEndpoint: endpoint,
		senderAddress:  sender,
	}
}

func (r *DaoFactoryService) GetDao(context *base.Context, daoName string, factoryAddress common.Address) (registryAddress common.Address, isExists bool, err error) {
	factoryInstance, daoFactoryErr := factorycontract.NewDaoFactory(factoryAddress, r.clientEth)

	if daoFactoryErr != nil {
		return common.HexToAddress(""), false, daoFactoryErr
	}

	daoRes, isExists, daoResErr := factoryInstance.GetDao(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	}, daoName)

	if daoResErr != nil {
		return common.HexToAddress(""), false, daoResErr
	}

	return daoRes.Registry, isExists, nil
}
