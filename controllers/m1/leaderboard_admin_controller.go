package m1

import (
	"context"

	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/leaderboards"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
)

type LeaderboardAdminController struct {
	db                *dynamodb.DynamoDB
	leaderboards      *dynamodb.Table[leaderboards.Leaderboard]
	leaderboardAdmins *dynamodb.Table[leaderboards.LeaderboardAdmin]
}

func NewLeaderboardAdminController(ctx context.Context, cfg config.AppConfig) *LeaderboardAdminController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)

	leaderboardInfos, err := dynamodb.NewTable[leaderboards.Leaderboard](ctx, db)
	if err != nil {
		logger.Critical(ctx, err)
	}

	leaderboardAdmins, err := dynamodb.NewTable[leaderboards.LeaderboardAdmin](ctx, db)
	if err != nil {
		logger.Critical(ctx, err)
	}

	return &LeaderboardAdminController{
		db:                db,
		leaderboards:      leaderboardInfos,
		leaderboardAdmins: leaderboardAdmins,
	}
}

func (r *LeaderboardAdminController) OnInit(route base.PathRouter) {
	route.POST("/migration", r.migration)
}

type PostMigrationRequest struct {
	Version int `json:"version"`
}

type PostMigrationResponse struct {
	Status string `json:"status"`
}

func (r *LeaderboardAdminController) migration(ctx *base.Context, req *PostMigrationRequest) (*PostMigrationResponse, *base.Error) {
	leaderboardData, err := r.leaderboards.FindAll(ctx)

	if err != nil {
		logger.Errorf(ctx, "err: %+v", err)
		return nil, errors.ErrGetAllDataFailed
	}

	for i := 0; i < len(leaderboardData); i++ {
		e2 := r.leaderboardAdmins.Insert(ctx, &leaderboards.LeaderboardAdmin{
			Key:              leaderboardData[i].Key,
			RankType:         leaderboardData[i].RankType,
			DailyCount:       leaderboardData[i].DailyCount,
			Level:            leaderboardData[i].Level,
			Exp:              leaderboardData[i].Exp,
			VotingCount:      leaderboardData[i].VotingCount,
			Version:          leaderboardData[i].Version,
			CreatedAt:        leaderboardData[i].CreatedAt,
			UpdatedAt:        leaderboardData[i].UpdatedAt,
			IndexVersionType: leaderboardData[i].IndexVersionType,
		})

		if e2 != nil {
			logger.Errorf(ctx, "err: %+v", e2)
			return nil, errors.ErrInsertFailed
		}
	}

	return &PostMigrationResponse{
		Status: "ok",
	}, nil
}
