package reward

import (
	"fmt"
	"math/big"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/common"

	v1 "biyard.co/iuniverse-webworker/controllers/v1"
)

type RewardController struct {
	wallet         *v1.WalletController
	feeEndpoint    string
	klaytnEndpoint string
	accountAddress string
	feePayer       string
	chainID        int
}

type ExperienceActivityData struct {
	Key          string   `abi:"key"`
	MissionInfo  string   `abi:"missionInfo"`
	Exp          *big.Int `abi:"exp"`
	ProgressDate *big.Int `abi:"progressDate"`
	StartDate    *big.Int `abi:"startDate"`
	EndDate      *big.Int `abi:"endDate"`
}

var abi = `
[
	{
		"inputs": [
		  {
			"internalType": "address",
			"name": "contractAddress",
			"type": "address"
		  },
		  {
			"internalType": "uint256",
			"name": "tokenID",
			"type": "uint256"
		  },
		  {
			"internalType": "uint256",
			"name": "amount",
			"type": "uint256"
		  }
		],
		"name": "addNFTEXP",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  },
	  {
		"inputs": [
		  {
			"internalType": "address",
			"name": "userAddress",
			"type": "address"
		  },
		  {
			"components": [
			  {
				"internalType": "string",
				"name": "key",
				"type": "string"
			  },
			  {
				"internalType": "string",
				"name": "missionInfo",
				"type": "string"
			  },
			  {
				"internalType": "uint256",
				"name": "exp",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "progressDate",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "startDate",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "endDate",
				"type": "uint256"
			  }
			],
			"internalType": "struct ExperienceActivityData",
			"name": "data",
			"type": "tuple"
		  }
		],
		"name": "addAccountActivities",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  }
]
`

func NewRewardController(owner, privateKey, feeEndpoint, klaytnEndpoint, accountAddress, feePayer string, chainID int) *RewardController {
	wallet := v1.NewWalletController(owner, privateKey, abi)
	return &RewardController{
		wallet:         wallet,
		feeEndpoint:    feeEndpoint,
		klaytnEndpoint: klaytnEndpoint,
		accountAddress: accountAddress,
		feePayer:       feePayer,
		chainID:        chainID,
	}
}

func (r *RewardController) RewardXP(ctx *base.Context, account string, msg string, rewardXP int, finishedAt int) (*v1.SignTransactionResponse, *base.Error) {
	res, err := r.wallet.UseFastHttp().SignTransaction(ctx, v1.SignTransactionRequest{
		FeeEndpoint:     r.feeEndpoint,
		Endpoint:        r.klaytnEndpoint,
		ContractAddress: r.accountAddress,
		Function:        "addAccountActivities",
		Params: []any{
			common.HexToAddress(account),
			ExperienceActivityData{
				Key:          fmt.Sprintf("%v-%v", msg, time.Now().Unix()),
				MissionInfo:  msg,
				Exp:          big.NewInt(int64(rewardXP)),
				ProgressDate: big.NewInt(int64(finishedAt)),
				StartDate:    big.NewInt(0),
				EndDate:      big.NewInt(10000000000),
			},
		},
		FeePayer: r.feePayer,
		ChainID:  int64(r.chainID),
	})

	if err != nil {
		logger.Errorf(ctx, "transaction send error: %+v", err)
		return nil, base.ErrUnknown
	}

	return res, nil
}
