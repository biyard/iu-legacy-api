package m1

import (
	"context"
	"fmt"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/mint_time"
	"github.com/klaytn/klaytn/client"
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
	route.POST("", r.SetTime)
}

type TimeRequest struct {
	WhitelistStartTime int64 `json:"whitelistStartTime"`
	WhitelistEndTime   int64 `json:"whitelistEndTime"`
	PublicStartTime    int64 `json:"publicStartTime"`
	PublicEndTime      int64 `json:"publicEndTime"`
}

type TimeResponse struct {
	Status string `json:"status"`
}

func (r *TimeController) SetTime(ctx *base.Context, req TimeRequest) (*TimeResponse, *base.Error) {
	dt, _ := r.times.FindOne(ctx, bson.D{{"key", "time"}})

	if dt != nil {
		logger.Debugf(ctx, "already exists: %+v", dt)

		e2 := r.times.UpdateOne(ctx, bson.D{{"key", fmt.Sprintf("%s", "time")}}, bson.D{{"$set", bson.D{{"whitelistStartTime", req.WhitelistStartTime}}}})
		e3 := r.times.UpdateOne(ctx, bson.D{{"key", fmt.Sprintf("%s", "time")}}, bson.D{{"$set", bson.D{{"whitelistEndTime", req.WhitelistEndTime}}}})
		e4 := r.times.UpdateOne(ctx, bson.D{{"key", fmt.Sprintf("%s", "time")}}, bson.D{{"$set", bson.D{{"publicStartTime", req.PublicStartTime}}}})
		e5 := r.times.UpdateOne(ctx, bson.D{{"key", fmt.Sprintf("%s", "time")}}, bson.D{{"$set", bson.D{{"publicEndTime", req.PublicEndTime}}}})

		if e2 != nil || e3 != nil || e4 != nil || e5 != nil {
			logger.Errorf(ctx, "wlStart: %+v wlEnd: %+v pubStart: %+v pubEnd: %+v err: %+v %+v %+v %+v", req.WhitelistStartTime, req.WhitelistEndTime, req.PublicStartTime, req.PublicEndTime, e2, e3, e4, e5)
			return nil, errors.ErrUpdateFailed
		}
	} else {
		e2 := r.times.Insert(ctx, &mint_time.Time{
			Key:                fmt.Sprintf("%s", "time"),
			WhitelistStartTime: req.WhitelistStartTime,
			WhitelistEndTime:   req.WhitelistEndTime,
			PublicStartTime:    req.PublicStartTime,
			PublicEndTime:      req.PublicEndTime,
		})

		if e2 != nil {
			logger.Errorf(ctx, "wlStart: %+v wlEnd: %+v pubStart: %+v pubEnd: %+v", req.WhitelistStartTime, req.WhitelistEndTime, req.PublicStartTime, req.PublicEndTime)
			return nil, errors.ErrInsertFailed
		}
	}

	return &TimeResponse{
		Status: "ok",
	}, nil
}
