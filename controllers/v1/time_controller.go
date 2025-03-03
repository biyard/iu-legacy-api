package v1

import (
	"context"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/contracts"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/mint_time"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
	"go.mongodb.org/mongo-driver/bson"
)

type TimeController struct {
	times     *dynamodb.Table[mint_time.Time]
	clientEth *client.Client
	cfg       config.AppConfig
}

func NewTimeController(ctx context.Context, cfg config.AppConfig) *TimeController {
	clientEth, clientErr := client.Dial(cfg.Klaytn.Endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	times, err := dynamodb.NewTable[mint_time.Time](ctx, db)
	if err != nil {
		panic(err)
	}

	return &TimeController{
		cfg:       cfg,
		clientEth: clientEth,
		times:     times,
	}
}

func (r *TimeController) OnInit(route base.PathRouter) {
	route.GET("", r.TimeByAddress)
}

type TimeAddressRequest struct {
	Sender string `form:"sender"`
}

type TimeAddressResponse struct {
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

func (r *TimeController) TimeByAddress(ctx *base.Context, req TimeAddressRequest) (*TimeAddressResponse, *base.Error) {
	now := time.Now()
	sec := now.Unix()

	closedTokenClaimInstance, closedTokenClaimErr := contracts.NewClosedTokenClaim(common.HexToAddress(r.cfg.ContractAddresses.ClosedTokenClaim), r.clientEth)

	if closedTokenClaimErr != nil {
		logger.Debugf(ctx, "err: %+v addr: %+v", closedTokenClaimErr, common.HexToAddress(r.cfg.ContractAddresses.ClosedTokenClaim))
		return nil, errors.ErrGetClosedTokenClaimInfoFailed
	}

	groups, err := closedTokenClaimInstance.AllowedOf(&bind.CallOpts{}, common.HexToAddress(req.Sender))
	if err != nil {
		logger.Debugf(ctx, "sender: %+v", common.HexToAddress(req.Sender))
		return nil, errors.ErrAllowedError
	}

	logger.Debugf(ctx, "groups: %+v len: %+v", groups, len(groups))

	dt, _ := r.times.FindOne(ctx, bson.D{{"key", "time"}})

	if dt == nil {
		return nil, errors.ErrNotExists
	}

	if len(groups) >= 1 && sec <= dt.WhitelistEndTime {
		return &TimeAddressResponse{
			StartTime: dt.WhitelistStartTime,
			EndTime:   dt.WhitelistEndTime,
		}, nil
	}

	return &TimeAddressResponse{
		StartTime: dt.PublicStartTime,
		EndTime:   dt.PublicEndTime,
	}, nil
}
