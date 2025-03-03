package m1

import (
	"context"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	errors "biyard.co/iuniverse-api/errors"
	heroesSongContestSubmission "biyard.co/iuniverse-api/models/heroesSongContestSubmission"
)

type HeroesSongContestAdminController struct {
	db                          *dynamodb.DynamoDB
	heroesSongContestSubmission *dynamodb.Table[heroesSongContestSubmission.HeroesSongContestSubmission]
	cfg                         config.AppConfig
}

func NewHeroesSongContestAdminController(ctx context.Context, cfg config.AppConfig) *HeroesSongContestAdminController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	heroesSongContestSubmission, err := dynamodb.NewTable[heroesSongContestSubmission.HeroesSongContestSubmission](ctx, db)

	if err != nil {
		logger.Critical(ctx, err)
	}

	return &HeroesSongContestAdminController{
		db:                          db,
		heroesSongContestSubmission: heroesSongContestSubmission,
		cfg:                         cfg,
	}
}

func (r *HeroesSongContestAdminController) OnInit(route base.PathRouter) {
	route.POST("/submissions", r.PutSubmissions)
	route.GET("/submissions", r.GetSubmissions)
}

type Submission struct {
	GoogleID    string `json:"googleId"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	MainNFTID   int64  `json:"mainNftId"`
	ImageURL    string `json:"imageUrl"`
	AudioURL    string `json:"audioUrl"`
	Lyrics      string `json:"lyrics"`
	SubmittedAt int64  `json:"submittedAt"`
}

type PutSubmissionsRequest struct {
	Submissions []Submission `json:"submissions"`
}

type PutSubmissionsResponse struct {
	Status string `json:"status"`
}

type GetSubmissionsRequest struct{}

type GetSubmissionsResponse struct {
	Status      string                                                    `json:"status"`
	Submissions []heroesSongContestSubmission.HeroesSongContestSubmission `json:"submissions"`
}

func (r *HeroesSongContestAdminController) PutSubmissions(ctx *base.Context, req *PutSubmissionsRequest) (*PutSubmissionsResponse, *base.Error) {
	for _, submission := range req.Submissions {
		_, e := r.PutSubmission(ctx, submission)

		if e != nil {
			logger.Errorf(ctx, "submission: %+v err: %+v", submission, e)

			return nil, e
		}
	}

	return &PutSubmissionsResponse{
		Status: "ok",
	}, nil
}

func (r *HeroesSongContestAdminController) PutSubmission(ctx *base.Context, submission Submission) (*PutSubmissionsResponse, *base.Error) {
	e := r.heroesSongContestSubmission.Insert(ctx, &heroesSongContestSubmission.HeroesSongContestSubmission{
		Key:         submission.GoogleID,
		Email:       submission.Email,
		Address:     submission.Address,
		Name:        submission.Name,
		Phone:       submission.Phone,
		MainNFTID:   submission.MainNFTID,
		AudioURL:    submission.AudioURL,
		ImageURL:    submission.ImageURL,
		Lyrics:      submission.Lyrics,
		SubmittedAt: submission.SubmittedAt,
		CreatedAt:   time.Now().Unix(),
	})

	if e != nil {
		logger.Errorf(ctx, "submission: %+v err: %+v", submission, e)
		return nil, errors.ErrInsertFailed
	}

	return &PutSubmissionsResponse{
		Status: "ok",
	}, nil
}

func (r *HeroesSongContestAdminController) GetSubmissions(ctx *base.Context, req *GetSubmissionsRequest) (*GetSubmissionsResponse, *base.Error) {
	doc, e := r.heroesSongContestSubmission.FindAll(ctx)

	if e != nil {
		logger.Errorf(ctx, "err: %+v", e)
		return nil, errors.ErrMsg
	}

	return &GetSubmissionsResponse{
		Status:      "ok",
		Submissions: doc,
	}, nil
}
