package byfactory

import (
	"context"

	"biyard.co/common/base"
	registrycontract "biyard.co/common/contracts/dao-registry"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
)

type DaoRegistryService struct {
	clientEth      *client.Client
	klaytnEndpoint string
	senderAddress  string
}

func NewDaoRegistryService(endpoint, sender string) *DaoRegistryService {
	clientEth, clientErr := client.Dial(endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	return &DaoRegistryService{
		clientEth:      clientEth,
		klaytnEndpoint: endpoint,
		senderAddress:  sender,
	}
}

func (r *DaoRegistryService) AddressOfExtension(context *base.Context, registryAddress common.Address, extensionName string) (extensionAddress common.Address, err error) {
	registryInstance, registryErr := registrycontract.NewDaoRegistry(registryAddress, r.clientEth)

	if registryErr != nil {
		return common.HexToAddress(""), registryErr
	}

	extensionAddr, extensionErr := registryInstance.AddressOfExtension(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	}, extensionName)

	if extensionErr != nil {
		return common.HexToAddress(""), extensionErr
	}

	return extensionAddr, nil
}

func (r *DaoRegistryService) GetExtensions(context *base.Context, registryAddress common.Address) ([]registrycontract.RegisteredExtension, error) {
	registryInstance, registryErr := registrycontract.NewDaoRegistry(registryAddress, r.clientEth)

	if registryErr != nil {
		return nil, registryErr
	}

	extensions, err := registryInstance.Extensions(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	})
	if err != nil {
		return nil, err
	}

	return extensions, nil
}

func (r *DaoRegistryService) GetExtension(context *base.Context, registryAddress common.Address, extensionName string) (registrycontract.RegisteredExtension, error) {
	registryInstance, registryErr := registrycontract.NewDaoRegistry(registryAddress, r.clientEth)

	if registryErr != nil {
		return registrycontract.RegisteredExtension{}, registryErr
	}

	extensions, err := registryInstance.Extensions(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	})
	if err != nil {
		return registrycontract.RegisteredExtension{}, err
	}

	check := false

	dt := registrycontract.RegisteredExtension{}
	for _, v := range extensions {
		if v.Name == extensionName {
			check = true
			dt = v
			break
		}
	}

	if !check {
		return registrycontract.RegisteredExtension{}, nil
	}

	return dt, nil
}
