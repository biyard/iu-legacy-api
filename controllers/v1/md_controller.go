package v1

import (
	"context"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/md"
	"biyard.co/iuniverse-api/models/md_account"
	"go.mongodb.org/mongo-driver/bson"
)

type MdController struct {
	mds        *dynamodb.Table[md.Md]
	mdAccounts *dynamodb.Table[md_account.MdAccount]
	cfg        config.AppConfig
}

func NewMdController(ctx context.Context, cfg config.AppConfig) *MdController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	mds, err := dynamodb.NewTable[md.Md](ctx, db)
	mdAccounts, err := dynamodb.NewTable[md_account.MdAccount](ctx, db)
	if err != nil {
		panic(err)
	}

	return &MdController{
		mds:        mds,
		mdAccounts: mdAccounts,
		cfg:        cfg,
	}
}

func (r *MdController) OnInit(route base.PathRouter) {
	route.GET("", r.GetMdData)
	route.GET("/accounts", r.GetMdAccountData)
	route.POST("/accounts", r.PostMdAccountData)
}

type MdGetRequest struct {
	Key string `form:"key"`
}

type MdGetResponse struct {
	Key                string `json:"key"`
	KoreanTitle        string `json:"koTitle"`
	KoreanDescription  string `json:"koDescription"`
	EnglishTitle       string `json:"enTitle"`
	EnglishDescription string `json:"enDescription"`
	StartTime          int64  `json:"startTime"`
	EndTime            int64  `json:"endTime"`
}

type MdPostCheckRequest struct {
	Address string `json:"address"`
	Check   int    `json:"check"` // 1: no, 2: yes
}

type MdPostCheckResponse struct {
	Status string `json:"status"`
}

type MdGetCheckRequest struct {
	Address string `json:"address"`
}

type MdGetCheckResponse struct {
	Address string `json:"address"`
	Check   bool   `json:"isModal"`
}

func (r *MdController) GetMdData(ctx *base.Context, req MdGetRequest) (*MdGetResponse, *base.Error) {
	dt, _ := r.mds.FindOne(ctx, bson.D{{"key", req.Key}})

	if dt == nil {
		return nil, errors.ErrNotExists
	}

	return &MdGetResponse{
		Key:                req.Key,
		KoreanTitle:        dt.KoreaMd.Title,
		KoreanDescription:  dt.KoreaMd.Contents,
		EnglishTitle:       dt.EnglishMd.Title,
		EnglishDescription: dt.EnglishMd.Contents,
		StartTime:          dt.StartTime,
		EndTime:            dt.EndTime,
	}, nil
}

func (r *MdController) PostMdAccountData(ctx *base.Context, req MdPostCheckRequest) (*MdPostCheckResponse, *base.Error) {
	dt, _ := r.mdAccounts.FindOne(ctx, bson.D{{"key", req.Address}})

	if dt != nil {
		logger.Debugf(ctx, "already exists: %+v", dt)

		e := r.mdAccounts.UpdateOne(ctx, bson.D{{"key", req.Address}}, bson.D{{"$set", bson.D{{"account", req.Address}}}})
		e2 := r.mdAccounts.UpdateOne(ctx, bson.D{{"key", req.Address}}, bson.D{{"$set", bson.D{{"check", req.Check}}}})

		if e != nil || e2 != nil {
			logger.Errorf(ctx, "account: %+v check: %+v err: %+v %+v", req.Address, req.Check, e, e2)
			return nil, errors.ErrUpdateFailed
		}
	} else {
		e := r.mdAccounts.Insert(ctx, &md_account.MdAccount{
			Key:     req.Address,
			Account: req.Address,
			Check:   req.Check,
		})

		if e != nil {
			logger.Errorf(ctx, "account: %+v check: %+v err: %+v", req.Address, req.Check, e)
			return nil, errors.ErrInsertFailed
		}
	}

	return &MdPostCheckResponse{
		Status: "ok",
	}, nil
}

func (r *MdController) GetMdAccountData(ctx *base.Context, req MdGetCheckRequest) (*MdGetCheckResponse, *base.Error) {
	dt, _ := r.mdAccounts.FindOne(ctx, bson.D{{"key", req.Address}})

	if dt == nil {
		return nil, errors.ErrNotExists
	}

	check := false

	if dt.Check == 1 {
		check = false
	} else {
		check = true
	}

	return &MdGetCheckResponse{
		Address: req.Address,
		Check:   check,
	}, nil
}
