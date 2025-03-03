package m1

import (
	"context"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/md"
	"go.mongodb.org/mongo-driver/bson"
)

type MdController struct {
	mds *dynamodb.Table[md.Md]
	cfg config.AppConfig
}

func NewMdController(ctx context.Context, cfg config.AppConfig) *MdController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	mds, err := dynamodb.NewTable[md.Md](ctx, db)
	if err != nil {
		panic(err)
	}

	return &MdController{
		mds: mds,
		cfg: cfg,
	}
}

func (r *MdController) OnInit(route base.PathRouter) {
	route.POST("", r.PostMdData)
}

type MdPostRequest struct {
	Key                string `json:"key"`
	KoreanTitle        string `json:"koTitle"`
	KoreanDescription  string `json:"koDescription"`
	EnglishTitle       string `json:"enTitle"`
	EnglishDescription string `json:"enDescription"`
	StartTime          int64  `json:"startTime"`
	EndTime            int64  `json:"endTime"`
}

type MdPostResponse struct {
	Status string `json:"status"`
}

func (r *MdController) PostMdData(ctx *base.Context, req MdPostRequest) (*MdPostResponse, *base.Error) {
	dt, _ := r.mds.FindOne(ctx, bson.D{{"key", req.Key}})

	if dt != nil {
		logger.Debugf(ctx, "already exists: %+v", dt)

		e2 := r.mds.UpdateOne(ctx, bson.D{{"key", req.Key}}, bson.D{{"$set", bson.D{{"ko", md.MdData{
			Title:    req.KoreanTitle,
			Contents: req.KoreanDescription,
		}}}}})
		e3 := r.mds.UpdateOne(ctx, bson.D{{"key", req.Key}}, bson.D{{"$set", bson.D{{"en", md.MdData{
			Title:    req.EnglishTitle,
			Contents: req.EnglishDescription,
		}}}}})
		e4 := r.mds.UpdateOne(ctx, bson.D{{"key", req.Key}}, bson.D{{"$set", bson.D{{"startTime", req.StartTime}}}})
		e5 := r.mds.UpdateOne(ctx, bson.D{{"key", req.Key}}, bson.D{{"$set", bson.D{{"endTime", req.EndTime}}}})

		if e2 != nil || e3 != nil || e4 != nil || e5 != nil {
			logger.Errorf(ctx, "koreanTitle: %+v koreanDes: %+v englishTitle: %+v englishDes: %+v startTime: %+v endTime: %+v err: %+v %+v %+v %+v", req.KoreanTitle, req.KoreanDescription, req.EnglishTitle, req.EnglishDescription, req.StartTime, req.EndTime, e2, e3, e4, e5)
			return nil, errors.ErrUpdateFailed
		}
	} else {
		e2 := r.mds.Insert(ctx, &md.Md{
			Key: req.Key,
			KoreaMd: md.MdData{
				Title:    req.KoreanTitle,
				Contents: req.KoreanDescription,
			},
			EnglishMd: md.MdData{
				Title:    req.EnglishTitle,
				Contents: req.EnglishDescription,
			},
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		})

		if e2 != nil {
			logger.Errorf(ctx, "koreanTitle: %+v koreanDes: %+v englishTitle: %+v englishDes: %+v startTime: %+v endTime: %+v err: %+v", req.KoreanTitle, req.KoreanDescription, req.EnglishTitle, req.EnglishDescription, req.StartTime, req.EndTime, e2)
			return nil, errors.ErrInsertFailed
		}
	}

	return &MdPostResponse{
		Status: "ok",
	}, nil
}
