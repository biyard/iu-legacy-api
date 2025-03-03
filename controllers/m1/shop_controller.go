package m1

import (
	"context"
	"math/big"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/common/utils/klaytn"
	"biyard.co/iuniverse-api/config"
	errors "biyard.co/iuniverse-api/errors"
	shop "biyard.co/iuniverse-api/models/shop"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/client"
	common "github.com/klaytn/klaytn/common"
)

var abi = `
[
	{
		"inputs": [
		  {
			"internalType": "uint256",
			"name": "itemId",
			"type": "uint256"
		  }
		],
		"name": "getShopItem",
		"outputs": [
		  {
			"components": [
			  {
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "price",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			  },
			  {
				"internalType": "string",
				"name": "name",
				"type": "string"
			  },
			  {
				"internalType": "string",
				"name": "image",
				"type": "string"
			  },
			  {
				"internalType": "uint256",
				"name": "supply",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "likes",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "reports",
				"type": "uint256"
			  },
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "address",
				"name": "creator",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "remaining",
				"type": "uint256"
			  },
			  {
				"internalType": "uint8",
				"name": "level",
				"type": "uint8"
			  },
			  {
				"internalType": "string",
				"name": "metadata",
				"type": "string"
			  }
			],
			"internalType": "struct Item",
			"name": "",
			"type": "tuple"
		  }
		],
		"stateMutability": "view",
		"type": "function"
	  },
	  {
		"inputs": [
		  {
			"internalType": "uint256",
			"name": "itemId",
			"type": "uint256"
		  },
		  {
			"components": [
			  {
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "price",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			  },
			  {
				"internalType": "string",
				"name": "name",
				"type": "string"
			  },
			  {
				"internalType": "string",
				"name": "image",
				"type": "string"
			  },
			  {
				"internalType": "uint256",
				"name": "supply",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "likes",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "reports",
				"type": "uint256"
			  },
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "address",
				"name": "creator",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "remaining",
				"type": "uint256"
			  },
			  {
				"internalType": "uint8",
				"name": "level",
				"type": "uint8"
			  },
			  {
				"internalType": "string",
				"name": "metadata",
				"type": "string"
			  }
			],
			"internalType": "struct Item",
			"name": "item",
			"type": "tuple"
		  }
		],
		"name": "updateItem",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  }
]
`

type ShopController struct {
	cfg               config.AppConfig
	clientEth         *client.Client
	lastChangeTokenID int64

	feepayer *klaytn.Wallet
	owner    *klaytn.Wallet
}

func NewShopAdminController(ctx context.Context, cfg config.AppConfig) *ShopController {

	clientEth, clientErr := client.Dial(cfg.Klaytn.Endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	f, err := klaytn.NewKlaytnWallet(cfg.Feepayer.Address, cfg.Feepayer.PrivateKey)
	if err != nil {
		logger.Critical(ctx, "invalid private key")
	}

	o, err := klaytn.NewKlaytnWallet(cfg.Owner.Address, cfg.Owner.PrivateKey)
	if err != nil {
		logger.Critical(ctx, "invalid private key")
	}

	o.SetABI(abi).SetEndpoint(cfg.Klaytn.Endpoint).SetChainID(int64(cfg.Klaytn.ChainID)).SetFeepayer(cfg.Feepayer.Address, cfg.Feepayer.PrivateKey)

	return &ShopController{
		cfg:       cfg,
		clientEth: clientEth,
		feepayer:  f,
		owner:     o,
	}
}

func (r *ShopController) OnInit(route base.PathRouter) {
	route.POST("", r.changeShopDataInfo)
	route.GET("/data", r.GetShopItem)
}

type ChangeShopDataInfoRequest struct {
	Id              int64  `json:"itemId"`
	Price           int64  `json:"price"`
	TokenId         int64  `json:"tokenId"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	Supply          int64  `json:"supply"`
	Likes           int64  `json:"likes"`
	Reports         int64  `json:"reports"`
	ContractAddress string `json:"contractAddress"`
	Creator         string `json:"creator"`
	Remaining       int64  `json:"remaining"`
	Level           int64  `json:"level"`
	Metadata        string `json:"metadata"`
}

type ChangeShopDataInfoResponse struct {
	Status string `json:"status"`
}

type GetShopItemRequest struct {
	Id int64 `form:"item-id"`
}

type GetShopItemResponse struct {
	Status string        `json:"status"`
	Item   shop.ShopItem `json:"item"`
}

func (r *ShopController) GetShopItem(ctx *base.Context, req *GetShopItemRequest) (*GetShopItemResponse, *base.Error) {
	p := &shop.ShopItem{}
	err := r.feepayer.Call(ctx, r.cfg.ContractAddresses.Shop, p, "getShopItem", big.NewInt(5))
	if err != nil {
		logger.Errorf(ctx, "get data failed: %+v", err)
		return nil, errors.ErrContractCallFailed
	}

	return &GetShopItemResponse{
		Status: "success",
		Item:   *p,
	}, nil
}

func (r *ShopController) changeShopDataInfo(ctx *base.Context, req *ChangeShopDataInfoRequest) (*ChangeShopDataInfoResponse, *base.Error) {
	// FIXME: use atomic
	if r.lastChangeTokenID >= req.TokenId {
		return &ChangeShopDataInfoResponse{
			Status: "success",
		}, nil
	}
	r.lastChangeTokenID = req.TokenId

	_, err := changeShopData(ctx, r, req)
	_ = err
	// if err != nil {
	// 	logger.Errorf(ctx, "transaction send error: %+v", err)
	// 	return nil, errors.ErrTransactionSend
	// }

	return &ChangeShopDataInfoResponse{
		Status: "success",
	}, nil
}

type ChangeShopDataResponse struct {
	Address   string         `json:"address"`
	TxReceipt *types.Receipt `json:"txReceipt"`
}

func changeShopData(ctx *base.Context, r *ShopController, item *ChangeShopDataInfoRequest) (*ChangeShopDataResponse, *base.Error) {
	tx, _ := r.owner.SignTransaction(
		ctx,
		r.cfg.ContractAddresses.Shop,
		"addAccountActivities",
		big.NewInt(int64(item.Id)),
		shop.ShopItem{
			Id:              big.NewInt(int64(item.Id)),
			Price:           big.NewInt(int64(item.Price)),
			TokenId:         big.NewInt(int64(item.TokenId)),
			Name:            item.Name,
			Image:           item.Image,
			Supply:          big.NewInt(int64(item.Supply)),
			Likes:           big.NewInt(int64(item.Likes)),
			Reports:         big.NewInt(int64(item.Reports)),
			ContractAddress: common.HexToAddress(item.ContractAddress),
			Creator:         common.HexToAddress(item.Creator),
			Remaining:       big.NewInt(int64(item.Remaining)),
			Level:           uint8(item.Level),
			Metadata:        item.Metadata,
		},
	)
	tx, _ = r.feepayer.SignRawTransactionAsFeepayer(ctx, tx)
	logger.Debugf(ctx, "rlp: %s", tx)
	receipt, _ := r.owner.SendRawTransaction(ctx, tx)
	_ = receipt

	return &ChangeShopDataResponse{
		Address:   r.owner.Address.String(),
		TxReceipt: receipt,
	}, nil
}
