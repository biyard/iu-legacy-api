package m1

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/common/utils/klaytn"
	"biyard.co/iuniverse-api/config"
	errors "biyard.co/iuniverse-api/errors"
	shop "biyard.co/iuniverse-api/models/shop"
	"biyard.co/iuniverse-api/services/canister"
	"biyard.co/iuniverse-api/services/discord"

	"biyard.co/common/cache/redis"
	a "github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
)

type IcpController struct {
	feepayer   *klaytn.Wallet
	owner      *klaytn.Wallet
	contract   string
	clientEth  *client.Client
	cfg        config.AppConfig
	cli        *klaytn.RPC
	abi        *a.ABI
	daoChannel *discord.DiscordAPI
	icp        *canister.CanisterAPI
	rcli       *redis.Redis
	proc       *redis.Document[bool]

	handles []func(*base.Context, *NotifyProposalRequest)
}

var proposalTypes = []string{"[인천히 이벤트/행사 연계 제안]"}

var shopAbi = `
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
		"internalType": "address",
		"name": "contractAddress",
		"type": "address"
		},
		{"internalType": "uint256", "name": "tokenID", "type": "uint256"}
	],
	"name": "getNFTEXP",
	"outputs": [
		{"internalType": "uint256", "name": "", "type": "uint256"},
		{"internalType": "uint256", "name": "", "type": "uint256"},
		{"internalType": "uint256", "name": "", "type": "uint256"}
	],
	"stateMutability": "view",
	"type": "function"
	},
	{
		"inputs": [
			{"internalType": "uint256", "name": "tokenId", "type": "uint256"}
		],
		"name": "ownerOf",
		"outputs": [
			{"internalType": "address", "name": "", "type": "address"}
		],
		"stateMutability": "view",
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
	},
	{
		"inputs": [
		  {
			"internalType": "address",
			"name": "userAddress",
			"type": "address"
		  },
		  {
			"internalType": "uint256",
			"name": "exp",
			"type": "uint256"
		  }
		],
		"name": "subAccountEXP",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  }
]
`

func NewIcpController(ctx context.Context, cfg config.AppConfig) *IcpController {
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

	cli := klaytn.NewRPC(ctx, cfg.Klaytn.Endpoint)
	parsed, err := a.JSON(strings.NewReader(shopAbi))

	daoChannel, err := discord.NewDiscordAPI(ctx, discord.DiscordAPIConfig{
		BotToken:  cfg.Discord.BotToken,
		ChannelID: cfg.Discord.Channels.DAO,
	})
	if err != nil {
		logger.Critical(ctx, "failed to create discord api: %+v", err)
	}

	icp_cli, err := canister.NewCanisterAPI(ctx, cfg.ICP.Endpoint)
	if err != nil {
		logger.Critical(ctx, "failed to create ICP api: %+v", err)
	}

	red, err := redis.NewRedis(ctx, cfg.Redis)
	logger.Debugf(ctx, "redis: %+v ", red)

	if err != (*base.Error)(nil) {
		logger.Critical(ctx, "failed to create redis: %+v", err)
	}

	obj := &IcpController{
		cfg:        cfg,
		feepayer:   f,
		owner:      o,
		contract:   cfg.ContractAddresses.NFT,
		clientEth:  clientEth,
		cli:        cli,
		abi:        &parsed,
		daoChannel: daoChannel,
		rcli:       red,
		icp:        icp_cli,
		proc:       redis.NewDocument[bool](red),
	}

	obj.handles = []func(*base.Context, *NotifyProposalRequest){obj.SubmitProposal, obj.FinishProposal}

	return obj
}

func (r *IcpController) OnInit(route base.PathRouter) {
	route.POST("/bridge/nft", r.Bridge)
	route.GET("/shop/item", r.getShopItem)
	route.GET("/nft/:token-id/level", r.GetNFTLevel)
	route.POST("/discord/proposals", r.NotifyProposal)
	route.GET("/nft/owner/:address/:token-id", r.GetOwner)
}

type GetOwnerRequest struct {
	TokenID uint   `uri:"token-id"`
	Address string `uri:"address"`
}

type GetOwnerResponse struct {
	Verified bool `json:"verified"`
}

func (r *IcpController) GetOwner(ctx *base.Context, req *GetOwnerRequest) (*GetOwnerResponse, *base.Error) {
	var out common.Address

	r.call(ctx, &out, r.cfg.ContractAddresses.NFT, "ownerOf", big.NewInt(int64(req.TokenID)))

	return &GetOwnerResponse{Verified: strings.ToLower(out.String()) == strings.ToLower(req.Address)}, nil
}

type NotifyProposalRequest struct {
	// Action: 1:"create" or 2:"finished"
	Action     int  `json:"action"`
	ProposalID uint `json:"proposalId"`

	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Deadline       int64    `json:"deadline"`
	ProposalType   uint     `json:"proposalType"`
	Result         *string  `json:"result"`
	Rewards        []Reward `json:"rewards,omitempty"`
	ProposerReward *Reward  `json:"proposerReward,omitempty"`
}

type Reward struct {
	Address string `json:"address"`
	Amount  uint64 `json:"amount"`
}

type NotifyProposalResponse struct {
	Status string `json:"status"`
}

type ExperienceActivityData struct {
	Key          string   `abi:"key"`
	MissionInfo  string   `abi:"missionInfo"`
	Exp          *big.Int `abi:"exp"`
	ProgressDate *big.Int `abi:"progressDate"`
	StartDate    *big.Int `abi:"startDate"`
	EndDate      *big.Int `abi:"endDate"`
}

type RewardXPResponse struct {
	Address   string         `json:"address"`
	TxReceipt *types.Receipt `json:"txReceipt"`
}

func (r *IcpController) RewardXP(ctx *base.Context, account string, finishedAt int64, rewardXP uint64, msg string) (*RewardXPResponse, *base.Error) {
	tx, _ := r.owner.SignTransaction(
		ctx,
		r.cfg.ContractAddresses.AccountProfile,
		"addAccountActivities",
		common.HexToAddress(account),
		ExperienceActivityData{
			Key:          fmt.Sprintf("dao-%v", time.Now().Unix()),
			MissionInfo:  msg,
			Exp:          big.NewInt(int64(rewardXP)),
			ProgressDate: big.NewInt(int64(finishedAt)),
			StartDate:    big.NewInt(0),
			EndDate:      big.NewInt(10000000000),
		},
	)
	tx, _ = r.feepayer.SignRawTransactionAsFeepayer(ctx, tx)
	logger.Debugf(ctx, "rlp: %s", tx)
	receipt, _ := r.owner.SendRawTransaction(ctx, tx)
	_ = receipt

	return &RewardXPResponse{
		Address:   account,
		TxReceipt: receipt,
	}, nil
}

func (r *IcpController) FinishProposal(ctx *base.Context, req *NotifyProposalRequest) {
	result := "기각 / Rejected"
	reward := "1 EXP"

	if *req.Result == "ACCEPTED" {
		result = "채택 / Yes"
		reward = "Vote Yes : VP x 10 EXP\nVote No : VP x 5 EXP"
	} else if *req.Result == "REJECTED" {
		result = "거절 / No"
		reward = "Vote Yes : VP x 5 EXP\nVote No : VP x 10 EXP"
	}

	_, err := r.RewardXP(ctx, req.ProposerReward.Address, req.Deadline, req.ProposerReward.Amount, "DAO Proposal Reward")
	if err != nil {
		logger.Errorf(ctx, "failed to reward xp: %+v, error: %+v", reward, err)
	}

	for _, reward := range req.Rewards {
		_, err := r.RewardXP(ctx, reward.Address, req.Deadline, reward.Amount, "DAO Voting Reward")
		if err != nil {
			logger.Errorf(ctx, "failed to reward xp: %+v, error: %+v", reward, err)
		}
	}

	zone, _ := time.LoadLocation("Asia/Seoul")
	msg := discord.DiscordEmbedRequest{
		Title: "[제안 종료 / Proposal closed]",
		Description: strings.Join([]string{
			"* 제목 / Title",
			proposalTypes[req.ProposalType] + " " + req.Title,
			"",
			"* 투표 종료일(Vote deadline)",
			time.Unix(req.Deadline, 0).In(zone).Format("2006-01-02 15:04:05") + " (KST)",
			"",
			"* 투표 결과 / Voting results",
			result,
			"",
			"* 보상 EXP / Rewards EXP",
			reward,
			"",
		}, "\n"),
		Buttons: []discord.DiscordButton{
			{
				Label: "결과 확인 / Go to result",
				URL:   fmt.Sprintf("%+v/dao/proposal/%v", r.cfg.Endpoint, req.ProposalID),
				Emoji: "📊",
			},
		},
	}
	mid, e := r.daoChannel.SendMessage(ctx, msg)
	if e != nil {
		logger.Errorf(ctx, "failed to send message: %+v", e)
	}
	logger.Debugf(ctx, "message id: %s", mid)

}

func (r *IcpController) SubmitProposal(ctx *base.Context, req *NotifyProposalRequest) {
	logger.Debugf(ctx, "SendProposalMessage: %+v", req)
	zone, _ := time.LoadLocation("Asia/Seoul")

	msg := discord.DiscordEmbedRequest{
		Title: "[새로운 제안 등록 / New proposal submitted]",
		Description: strings.Join([]string{
			"* 제목 / Title",
			proposalTypes[req.ProposalType] + " " + req.Title,
			"",
			"* 투표 종료일(Vote deadline)",
			time.Unix(req.Deadline, 0).In(zone).Format("2006-01-02 15:04:05") + " (KST)",
			"",
			"* 내용(Detail)",
			req.Description,
			"",
			"※ 투표는 인천히어로즈 홀더만 참여 가능합니다.",
		}, "\n"),
		Buttons: []discord.DiscordButton{
			{
				Label: "투표하러 가기 / Go to vote",
				URL:   fmt.Sprintf("%v/dao/proposal/%v", r.cfg.Endpoint, req.ProposalID),
				Emoji: "🗳",
			},
		},
	}

	mid, err := r.daoChannel.SendMessage(ctx, msg)
	if err != nil {
		logger.Errorf(ctx, "failed to send message: %+v", err)
	}
	logger.Debugf(ctx, "message id: %s", mid)

}

func (r *IcpController) NotifyProposal(ctx *base.Context, req *NotifyProposalRequest) (*NotifyProposalResponse, *base.Error) {
	key := fmt.Sprintf("%v##%v", req.Action, req.ProposalID)

	lock := r.rcli.NewLock(key)
	logger.Infof(ctx, "NotifyProposal: %+v", req)
	lock.Lock()
	logger.Infof(ctx, "Lock: %+v", req)
	defer lock.Unlock()

	if doc, err := r.proc.Get(ctx, key); err == nil && *doc == true {
		logger.Debugf(ctx, "Already notified: %+v", req)
		return &NotifyProposalResponse{Status: "Ok"}, nil
	}
	r.proc.Cache(ctx, key, true)

	logger.Infof(ctx, "Initialized NotifyProposal start: %+v", req)
	r.handles[req.Action-1](ctx, req)

	return &NotifyProposalResponse{Status: "Ok"}, nil
}

type GetNFTLevelRequest struct {
	TokenID uint `uri:"token-id"`
}

type GetNFTLevelResponse struct {
	Level uint `json:"level"`
}

// GetNFTLevel returns a level of a NFT token by token ID.
func (r *IcpController) GetNFTLevel(ctx *base.Context, req *GetNFTLevelRequest) (*GetNFTLevelResponse, *base.Error) {
	var (
		level  *big.Int
		elevel *big.Int
		exp    *big.Int
	)
	out := []any{&level, &elevel, &exp}

	r.call(ctx, &out, r.cfg.ContractAddresses.EXP, "getNFTEXP", common.HexToAddress(r.cfg.ContractAddresses.NFT), big.NewInt(int64(req.TokenID)))

	logger.Debugf(ctx, "out: %+v", out)

	return &GetNFTLevelResponse{Level: uint(level.Int64())}, nil
}

type BridgeRequest struct {
	TokenID   uint   `json:"tokenId"`
	Address   string `json:"address"`
	Signature string `json:"signature"`
}

type BridgeResponse struct {
	Status string `json:"status"`
}

type GetPublicShopItemRequest struct{}

type GetPublicShopItemResponse struct {
	Status string        `json:"status"`
	Item   shop.ShopItem `json:"item"`
}

func (r *IcpController) call(ctx *base.Context, out any, contractAddress string, function string, params ...any) *base.Error {
	cli := r.cli

	data, err := r.abi.Pack(function, params...)
	if err != nil {
		logger.Errorf(ctx, "Failed to pack abi: %v", err)
		return base.ErrKlaytnFailedToEncodeTx.WithMessage(err.Error())
	}

	logger.Debugf(ctx, "Hex string: %+v", hexutil.Encode(data))

	str, err := cli.Call(ctx, r.feepayer.Address.String(), contractAddress, hexutil.Encode(data))

	if err != nil {
		logger.Errorf(ctx, "Failed to call: %v", err)
		return base.ErrKlaytnFailedToCall.WithMessage(err.Error())
	}
	if str == "" {
		logger.Errorf(ctx, "empty response: %v", str)
		return base.ErrKlaytnFailedToCall.WithMessage("empty response")
	}

	d, err := hexutil.Decode(str)
	if err != nil {
		logger.Errorf(ctx, "Failed to decode: %v", err)
		return base.ErrKlaytnFailedToDecodeTx.WithMessage(err.Error())
	}

	logger.Debugf(ctx, "Decoded: %+v", d)

	if err := r.abi.Unpack(out, function, d); err != nil {
		logger.Errorf(ctx, "Failed to unpack: %v", err)
		return base.ErrKlaytnFailedToDecodeTx.WithMessage(err.Error())
	}

	return nil

}

func (r *IcpController) getShopItem(ctx *base.Context, req *GetPublicShopItemRequest) (*GetPublicShopItemResponse, *base.Error) {
	p := &shop.ShopItem{}
	err := r.feepayer.Call(ctx, r.cfg.ContractAddresses.Shop, p, "getShopItem", big.NewInt(5))
	if err != nil {
		logger.Errorf(ctx, "get data failed: %+v", err)
		return nil, errors.ErrContractCallFailed
	}

	return &GetPublicShopItemResponse{
		Status: "success",
		Item:   *p,
	}, nil
}

// Bridge is a function that transfer NFT token to contract address or owner address (temporary owner).
func (r *IcpController) Bridge(ctx *base.Context, req *BridgeRequest) (*BridgeResponse, *base.Error) {
	tx, _ := r.owner.SignTransaction(ctx, r.contract, "safeTransferFrom", common.HexToAddress(req.Address), r.owner.Address, big.NewInt(int64(req.TokenID)))
	tx, _ = r.feepayer.SignRawTransactionAsFeepayer(ctx, tx)
	logger.Debugf(ctx, "rlp: %s", tx)
	receipt, _ := r.owner.SendRawTransaction(ctx, tx)
	_ = receipt

	return &BridgeResponse{Status: "Ok"}, nil
}
