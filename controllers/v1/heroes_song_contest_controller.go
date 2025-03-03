package v1

import (
	"context"
	"fmt"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/heroesSongContest"
	"go.mongodb.org/mongo-driver/bson"
)

type HeroesSongContestController struct {
	db                         *dynamodb.DynamoDB
	heroesSongContestCandidate *dynamodb.Table[heroesSongContest.HeroesSongContestCandidate]
	heroesSongContestLike      *dynamodb.Table[heroesSongContest.HeroesSongContestLike]
	cfg                        config.AppConfig
}

func NewHeroesSongContestController(ctx context.Context, cfg config.AppConfig) *HeroesSongContestController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	heroesSongContestCandidate, err := dynamodb.NewTable[heroesSongContest.HeroesSongContestCandidate](ctx, db)
	heroesSongContestLike, err := dynamodb.NewTable[heroesSongContest.HeroesSongContestLike](ctx, db)

	if err != nil {
		logger.Critical(ctx, err)
	}

	return &HeroesSongContestController{
		db:                         db,
		heroesSongContestCandidate: heroesSongContestCandidate,
		heroesSongContestLike:      heroesSongContestLike,
		cfg:                        cfg,
	}
}

func (r *HeroesSongContestController) OnInit(route base.PathRouter) {
	route.POST("/like", r.PutLike)
	route.POST("/play/count", r.IncreasePlayCount)
	route.GET("/candidates/all", r.ListCandidates)
	route.GET("/candidate", r.GetCandidate)
	route.GET("/like/candidates", r.getLikeCandidateIdList)
}

type PutLikeRequest struct {
	CandidateID string `json:"candidateId"`
	Address     string `json:"address"`
}

type PutLikeResponse struct {
	Status string `json:"status"`
}

type IncreasePlayCountRequest struct {
	CandidateID string `json:"candidateId"`
}

type IncreasePlayCountReponse struct {
	Status string `json:"status"`
}

type ListCandidatesRequest struct{}

type ListCandidatesResponse struct {
	Status     string                                         `json:"status"`
	Candidates []heroesSongContest.HeroesSongContestCandidate `json:"candidates"`
}

type GetCandidateRequest struct {
	ID string `form:"id"`
}

type GetCandidateReponse struct {
	Status    string                                       `json:"status"`
	Candidate heroesSongContest.HeroesSongContestCandidate `json:"candidate"`
}

type GetLikeCandidateIdListRequest struct {
	Address string `form:"address"`
}

type GetLikeCandidateIdListReponse struct {
	Status          string   `json:"status"`
	CandidateIDList []string `json:"candidateIdList"`
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func (r *HeroesSongContestController) PutLike(ctx *base.Context, req *PutLikeRequest) (*PutLikeResponse, *base.Error) {
	if req.Address == "" || req.CandidateID == "" {
		return nil, errors.ErrMissingRequiredParameters
	}

	like, err := r.heroesSongContestLike.FindOne(ctx, bson.D{{"key", fmt.Sprintf(req.Address)}})

	if like == nil {
		e := r.heroesSongContestLike.Insert(ctx, &heroesSongContest.HeroesSongContestLike{
			Key:             req.Address,
			CandidateIDList: []string{req.CandidateID},
			CreatedAt:       time.Now().Unix(),
		})

		if e != nil {
			logger.Errorf(ctx, "err: %+v", e)
			return nil, errors.ErrInsertFailed
		}
	} else {
		if contains(like.CandidateIDList, req.CandidateID) {
			return nil, errors.ErrUpdateFailed.WithMessage("Candidate ID already exists in the like list")
		}

		candidateIDList := append(like.CandidateIDList, req.CandidateID)

		e := r.heroesSongContestLike.UpdateOne(ctx, bson.D{{"key", req.Address}}, bson.D{{"$set", bson.D{{"candidateIdList", candidateIDList}}}})

		if e != nil {
			logger.Errorf(ctx, "err: %+v", e)
			return nil, errors.ErrUpdateFailed
		}
	}

	candidate, err := r.heroesSongContestCandidate.FindOne(ctx, bson.D{{"key", fmt.Sprintf(req.CandidateID)}})

	if candidate == nil {
		logger.Errorf(ctx, "err: %+v", err)
		return nil, errors.ErrNotExists.WithMessage("Candidate not found")
	}

	r.heroesSongContestCandidate.UpdateOne(ctx, bson.D{{"key", req.CandidateID}}, bson.D{{"$set", bson.D{{"likeCount", candidate.LikeCount + 1}}}})

	return &PutLikeResponse{
		Status: "ok",
	}, nil
}

func (r *HeroesSongContestController) IncreasePlayCount(ctx *base.Context, req *IncreasePlayCountRequest) (*IncreasePlayCountReponse, *base.Error) {
	if req.CandidateID == "" {
		return nil, errors.ErrMissingRequiredParameters
	}

	candidate, err := r.heroesSongContestCandidate.FindOne(ctx, bson.D{{"key", fmt.Sprintf(req.CandidateID)}})

	if candidate == nil {
		logger.Errorf(ctx, "err: %+v", err)
		return nil, errors.ErrNotExists.WithMessage("Candidate not found")
	}

	r.heroesSongContestCandidate.UpdateOne(ctx, bson.D{{"key", req.CandidateID}}, bson.D{{"$set", bson.D{{"playCount", candidate.PlayCount + 1}}}})

	return &IncreasePlayCountReponse{
		Status: "ok",
	}, nil
}

func (r *HeroesSongContestController) ListCandidates(ctx *base.Context, req *ListCandidatesRequest) (*ListCandidatesResponse, *base.Error) {
	doc, e := r.heroesSongContestCandidate.FindAll(ctx)

	if e != nil {
		logger.Errorf(ctx, "err: %+v", e)
		return nil, errors.ErrMsg
	}

	return &ListCandidatesResponse{
		Status:     "ok",
		Candidates: doc,
	}, nil
}

func (r *HeroesSongContestController) GetCandidate(ctx *base.Context, req *GetCandidateRequest) (*GetCandidateReponse, *base.Error) {
	candidate, err := r.heroesSongContestCandidate.FindOne(ctx, bson.D{{"key", fmt.Sprintf(req.ID)}})

	if candidate == nil {
		logger.Errorf(ctx, "err: %+v", err)
		return nil, errors.ErrNotExists
	}

	return &GetCandidateReponse{
		Status:    "ok",
		Candidate: *candidate,
	}, nil
}

func (r *HeroesSongContestController) getLikeCandidateIdList(ctx *base.Context, req *GetLikeCandidateIdListRequest) (*GetLikeCandidateIdListReponse, *base.Error) {
	like, _ := r.heroesSongContestLike.FindOne(ctx, bson.D{{"key", fmt.Sprintf(req.Address)}})

	if like == nil {
		return &GetLikeCandidateIdListReponse{
			Status:          "ok",
			CandidateIDList: []string{},
		}, nil
	}

	return &GetLikeCandidateIdListReponse{
		Status:          "ok",
		CandidateIDList: like.CandidateIDList,
	}, nil
}
