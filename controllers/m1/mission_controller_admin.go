package m1

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	errors "biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/mission"
	"go.mongodb.org/mongo-driver/bson"
)

type MissionAdminController struct {
	db    *dynamodb.DynamoDB
	quizs *dynamodb.Table[mission.Quiz]
}

func NewMissionAdminController(ctx context.Context, cfg config.AppConfig) *MissionAdminController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)

	quizs, err := dynamodb.NewTable[mission.Quiz](ctx, db)

	if err != nil {
		logger.Critical(ctx, err)
	}

	return &MissionAdminController{
		db:    db,
		quizs: quizs,
	}
}

func (r *MissionAdminController) OnInit(route base.PathRouter) {
	route.POST("/quiz", r.addQuizMission)
}

type AddQuizMissionRequest struct {
	MissionNameKo        string     `json:"missionNameKo"`
	MissionNameEn        string     `json:"missionNameEn"`
	MissionTypeKo        string     `json:"missionTypeKo"`
	MissionTypeEn        string     `json:"missionTypeEn"`
	QuestionsKo          []string   `json:"questionsKo"`
	QuestionsEn          []string   `json:"questionsEn"`
	MultipleChoiceViewKo [][]string `json:"multipleChoiceViewKo"`
	MultipleChoiceViewEn [][]string `json:"multipleChoiceViewEn"`
	Answers              []int      `json:"answers"`
	Exp                  int        `json:"exp"`
	StartDate            int        `json:"startDate"`
	Cutline              int        `json:"cutline"`
}

type AddQuizMissionResponse struct {
	Status               string     `json:"status"`
	MissionNameKo        string     `json:"missionNameKo"`
	MissionNameEn        string     `json:"missionNameEn"`
	MissionTypeKo        string     `json:"missionTypeKo"`
	MissionTypeEn        string     `json:"missionTypeEn"`
	QuestionsKo          []string   `json:"questionsKo"`
	QuestionsEn          []string   `json:"questionsEn"`
	MultipleChoiceViewKo [][]string `json:"multipleChoiceViewKo"`
	MultipleChoiceViewEn [][]string `json:"multipleChoiceViewEn"`
	Answers              []int      `json:"answers"`
	Exp                  int        `json:"exp"`
	StartDate            int        `json:"startDate"`
	Cutline              int        `json:"cutline"`
}

func (r *MissionAdminController) addQuizMission(ctx *base.Context, req *AddQuizMissionRequest) (*AddQuizMissionResponse, *base.Error) {
	now := time.Now()
	sec := now.Unix()

	s, _ := r.quizs.FindOne(ctx, bson.D{{Key: "key", Value: "startNumber"}})

	startNum := 0
	if s != nil {
		startNum = int(s.StartId)

		err := r.quizs.Insert(ctx, &mission.Quiz{
			Key:       "startNumber",
			StartId:   int64(startNum),
			CreatedAt: sec,
		})

		if err != nil {
			logger.Errorf(ctx, "start number insert failed err: %+v", err)
			return nil, errors.ErrInsertStartNumberFailed
		}
	}

	e := r.quizs.Insert(ctx, &mission.Quiz{
		Key:                fmt.Sprintf("%+v", startNum),
		MissionNameKo:      req.MissionNameKo,
		MissionNameEn:      req.MissionNameEn,
		MissionType:        req.MissionTypeKo,
		Exp:                req.Exp,
		StartDate:          req.StartDate,
		Questions:          req.QuestionsKo,
		MultipleChoiceView: req.MultipleChoiceViewKo,
		Answers:            req.Answers,
		Cutline:            req.Cutline,
		CreatedAt:          sec,
		IndexStartDate:     fmt.Sprintf("%s-%s", "ko", strconv.Itoa(req.StartDate)),
		Lang:               "ko",
	})

	if e != nil {
		logger.Errorf(ctx, "quizData insert Failed %+v %+v %+v %+v %+v %+v %+v %+v %+v", req.MissionNameKo, req.MissionTypeKo, req.Exp, req.StartDate, req.QuestionsKo, req.MultipleChoiceViewKo, req.Answers, req.Cutline)
		return nil, errors.ErrUpdateFailed
	}

	startNum += 1

	e2 := r.quizs.Insert(ctx, &mission.Quiz{
		Key:                fmt.Sprintf("%+v", startNum),
		MissionNameKo:      req.MissionNameKo,
		MissionNameEn:      req.MissionNameEn,
		MissionType:        req.MissionTypeEn,
		Exp:                req.Exp,
		StartDate:          req.StartDate,
		Questions:          req.QuestionsEn,
		MultipleChoiceView: req.MultipleChoiceViewEn,
		Answers:            req.Answers,
		Cutline:            req.Cutline,
		CreatedAt:          sec,
		IndexStartDate:     fmt.Sprintf("%s-%s", "en", strconv.Itoa(req.StartDate)),
		Lang:               "en",
	})

	if e2 != nil {
		logger.Errorf(ctx, "quizData insert Failed %+v %+v %+v %+v %+v %+v %+v %+v %+v", req.MissionNameEn, req.MissionTypeEn, req.Exp, req.StartDate, req.QuestionsEn, req.MultipleChoiceViewEn, req.Answers, req.Cutline)
		return nil, errors.ErrUpdateFailed
	}

	startNum += 1

	e3 := r.quizs.UpdateOne(ctx, bson.D{{"key", "startNumber"}}, bson.D{{"$set", bson.D{{"startId", int64(startNum)}}}})

	if e3 != nil {
		logger.Errorf(ctx, "start number update failed err: %+v", e3)
		return nil, errors.ErrUpdateStartNumberFailed
	}

	return &AddQuizMissionResponse{
		Status:               "ok",
		MissionNameKo:        req.MissionNameKo,
		MissionNameEn:        req.MissionNameEn,
		MissionTypeKo:        req.MissionTypeKo,
		MissionTypeEn:        req.MissionTypeEn,
		QuestionsKo:          req.QuestionsKo,
		QuestionsEn:          req.QuestionsEn,
		MultipleChoiceViewKo: req.MultipleChoiceViewKo,
		MultipleChoiceViewEn: req.MultipleChoiceViewEn,
		Answers:              req.Answers,
		Exp:                  req.Exp,
		StartDate:            req.StartDate,
		Cutline:              req.Cutline,
	}, nil
}
