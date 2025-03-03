package byfactory

import (
	"context"
	"fmt"

	"biyard.co/common/base"
	bankContract "biyard.co/common/contracts/bank"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
)

type BankService struct {
	clientEth      *client.Client
	klaytnEndpoint string
	senderAddress  string
}

func NewBankService(endpoint, sender string) *BankService {
	clientEth, clientErr := client.Dial(endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	return &BankService{
		clientEth:      clientEth,
		klaytnEndpoint: endpoint,
		senderAddress:  sender,
	}
}

func (r *BankService) GetRoles(context *base.Context, bankAddress common.Address) ([]string, []string, error) {
	bankInstance, bankErr := bankContract.NewBank(bankAddress, r.clientEth)

	if bankErr != nil {
		return nil, nil, bankErr
	}

	roles, values, roleErr := bankInstance.GetRoleList(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	})

	if roleErr != nil {
		return nil, nil, roleErr
	}

	roleValues := []string{}

	for _, v := range values {
		val := fmt.Sprintf("%s", v)
		roleValues = append(roleValues, val)
	}

	return roles, roleValues, nil
}

func (r *BankService) GetStandardTokenAddress(context *base.Context, bankAddress common.Address) (common.Address, error) {
	bankInstance, bankErr := bankContract.NewBank(bankAddress, r.clientEth)

	if bankErr != nil {
		return common.HexToAddress("0x0000000000000000000000000000000000000000"), bankErr
	}

	addr, err := bankInstance.GetStandardTokenAddress(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	})
	if err != nil {
		return common.HexToAddress("0x0000000000000000000000000000000000000000"), err
	}

	return addr, nil
}
