package v1

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	leaderboards "biyard.co/iuniverse-api/models/leaderboards"
	"go.mongodb.org/mongo-driver/bson"
)

type LeaderboardController struct {
	cfg config.AppConfig

	db *dynamodb.DynamoDB

	indexDailyCount  *dynamodb.Index[leaderboards.Leaderboard]
	indexLevel       *dynamodb.Index[leaderboards.Leaderboard]
	indexExperience  *dynamodb.Index[leaderboards.Leaderboard]
	indexVotingCount *dynamodb.Index[leaderboards.Leaderboard]
	leaderboardInfos *dynamodb.Table[leaderboards.Leaderboard]
}

func NewLeaderboardController(ctx context.Context, cfg config.AppConfig) *LeaderboardController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	indexDailyCount, _ := dynamodb.NewIndex[leaderboards.Leaderboard](ctx, db, "sortingDailyCount")
	indexLevel, _ := dynamodb.NewIndex[leaderboards.Leaderboard](ctx, db, "sortingLevel")
	indexExperience, _ := dynamodb.NewIndex[leaderboards.Leaderboard](ctx, db, "sortingExp")
	indexVotingCount, _ := dynamodb.NewIndex[leaderboards.Leaderboard](ctx, db, "sortingVotingCount")

	leaderboardInfos, err := dynamodb.NewTable[leaderboards.Leaderboard](ctx, db)
	if err != nil {
		logger.Critical(ctx, err)
	}

	return &LeaderboardController{
		cfg:              cfg,
		db:               db,
		indexDailyCount:  indexDailyCount,
		indexLevel:       indexLevel,
		indexExperience:  indexExperience,
		indexVotingCount: indexVotingCount,
		leaderboardInfos: leaderboardInfos,
	}
}

func (r *LeaderboardController) OnInit(route base.PathRouter) {
	route.GET("/rank/daily", r.getDailyRank)
	route.GET("/rank/voting", r.getVotingRank)
	route.GET("/rank/experience", r.getExperienceRank)
	route.GET("/rank/level", r.getLevelRank)
}

type GetLeaderboardRequest struct {
	Version   int `form:"version"`
	RankCount int `form:"rankCount"`
}

type GetLeaderboardResponse struct {
	Status      string                     `json:"status"`
	Leaderboard []leaderboards.Leaderboard `json:"leaderboard"`
}

type GetDailyRankResponse struct {
	Status      string      `json:"status"`
	Leaderboard []DailyRank `json:"leaderboard"`
	UpdatedAt   int         `json:"updatedAt"`
}

type GetVotingResponse struct {
	Status      string       `json:"status"`
	Leaderboard []VotingRank `json:"leaderboard"`
	UpdatedAt   int          `json:"updatedAt"`
}

type GetExperienceResponse struct {
	Status      string           `json:"status"`
	Leaderboard []ExperienceRank `json:"leaderboard"`
	UpdatedAt   int              `json:"updatedAt"`
}

type GetLevelResponse struct {
	Status      string      `json:"status"`
	Leaderboard []LevelRank `json:"leaderboard"`
	UpdatedAt   int         `json:"updatedAt"`
}

type DailyRank struct {
	Version        int    `json:"version"`
	Rank           int    `json:"rank"`
	DailyCount     int    `json:"dailyCount"`
	AccountAddress string `json:"accountAddress"`
}

type VotingRank struct {
	Version        int    `json:"version"`
	Rank           int    `json:"rank"`
	VotingCount    int    `json:"votingCount"`
	AccountAddress string `json:"accountAddress"`
}

type ExperienceRank struct {
	Version        int    `json:"version"`
	Rank           int    `json:"rank"`
	Experience     int    `json:"experience"`
	NftNum         int    `json:"nftNum"`
	AccountAddress string `json:"accountAddress"`
}

type LevelRank struct {
	Version        int    `json:"version"`
	Rank           int    `json:"rank"`
	Level          int    `json:"level"`
	NftNum         int    `json:"nftNum"`
	AccountAddress string `json:"accountAddress"`
}

func (r *LeaderboardController) getDailyRank(c *base.Context, req *GetLeaderboardRequest) (*GetDailyRankResponse, *base.Error) {
	logger.Debugf(c, "get Daily Rank Api")

	timeData, _ := r.leaderboardInfos.FindOne(c, bson.D{{"key", fmt.Sprintf("DMR")}})

	t := 0

	if timeData != nil {
		t = int(timeData.UpdatedAt)
	} else {
		t = 0
	}

	leaderboardInfos, err := r.indexDailyCount.Find(c, dynamodb.QueryRequest{
		Key:            fmt.Sprintf("%+v-DMR", strconv.Itoa(req.Version)),
		DisableSortKey: true,
	})

	if err != nil {
		logger.Errorf(c, "DailyRank Query Failed")
		return nil, base.ErrUnknown.WithMessage("failed to query daily rank" + err.Error())
	}

	n := 0 //정렬 갯수 선정

	for i := 0; i < len(leaderboardInfos.Items)-1; i++ {
		if leaderboardInfos.Items[i].DailyCount != leaderboardInfos.Items[i+1].DailyCount && i >= req.RankCount {
			n = i + 1
			break
		}
	}

	if n <= req.RankCount {
		n = len(leaderboardInfos.Items)
	}

	items := leaderboardInfos.Items[:n]

	sort.SliceStable(items, func(i, j int) bool {
		if items[i].DailyCount > items[j].DailyCount {
			return true
		} else if items[i].DailyCount == items[j].DailyCount {
			if items[i].CreatedAt > items[j].CreatedAt {
				return false
			} else {
				return true
			}
		} else {
			return false
		}
	})

	leaderboard := make([]DailyRank, 0)

	for i := 0; i < len(items); i++ {
		if i == req.RankCount {
			break
		}

		d := strings.Split(items[i].Key, "#")

		leaderboard = append(leaderboard, DailyRank{
			Version:        req.Version,
			Rank:           i + 1,
			DailyCount:     items[i].DailyCount,
			AccountAddress: d[3],
		})
	}

	return &GetDailyRankResponse{
		Status:      "ok",
		Leaderboard: leaderboard,
		UpdatedAt:   t,
	}, nil
}

func (r *LeaderboardController) getVotingRank(c *base.Context, req *GetLeaderboardRequest) (*GetVotingResponse, *base.Error) {
	logger.Debugf(c, "get Voting Rank Api")

	timeData, _ := r.leaderboardInfos.FindOne(c, bson.D{{"key", fmt.Sprintf("VR")}})

	t := 0

	if timeData != nil {
		t = int(timeData.UpdatedAt)
	} else {
		t = 0
	}

	leaderboardInfos, err := r.indexVotingCount.Find(c, dynamodb.QueryRequest{
		Key:            fmt.Sprintf("%+v-VR", strconv.Itoa(req.Version)),
		DisableSortKey: true,
	})

	if err != nil {
		logger.Errorf(c, "VotingRank Query Failed")
		return nil, base.ErrUnknown.WithMessage("failed to query voting rank" + err.Error())
	}

	n := 0 //정렬 갯수 선정

	for i := 0; i < len(leaderboardInfos.Items)-1; i++ {
		if leaderboardInfos.Items[i].VotingCount != leaderboardInfos.Items[i+1].VotingCount && i >= req.RankCount {
			n = i + 1
			break
		}
	}

	if n <= req.RankCount {
		n = len(leaderboardInfos.Items)
	}

	items := leaderboardInfos.Items[:n]

	sort.SliceStable(items, func(i, j int) bool {
		if items[i].VotingCount > items[j].VotingCount {
			return true
		} else if items[i].VotingCount == items[j].VotingCount {
			if items[i].CreatedAt > items[j].CreatedAt {
				return false
			} else {
				return true
			}
		} else {
			return false
		}
	})

	leaderboard := make([]VotingRank, 0)

	for i := 0; i < len(items); i++ {
		if i == req.RankCount {
			break
		}

		d := strings.Split(items[i].Key, "#")

		leaderboard = append(leaderboard, VotingRank{
			Version:        req.Version,
			Rank:           i + 1,
			VotingCount:    items[i].VotingCount,
			AccountAddress: d[3],
		})
	}

	return &GetVotingResponse{
		Status:      "ok",
		Leaderboard: leaderboard,
		UpdatedAt:   t,
	}, nil
}

func (r *LeaderboardController) getExperienceRank(c *base.Context, req *GetLeaderboardRequest) (*GetExperienceResponse, *base.Error) {
	logger.Debugf(c, "get Experience Rank Api")

	timeData, _ := r.leaderboardInfos.FindOne(c, bson.D{{"key", fmt.Sprintf("ER")}})

	t := 0

	if timeData != nil {
		t = int(timeData.UpdatedAt)
	} else {
		t = 0
	}

	leaderboardInfos, err := r.indexExperience.Find(c, dynamodb.QueryRequest{
		Key:            fmt.Sprintf("%+v-ER", strconv.Itoa(req.Version)),
		DisableSortKey: true,
	})

	if err != nil {
		logger.Errorf(c, "ExperienceRank Query Failed")
		return nil, base.ErrUnknown.WithMessage("failed to query experience rank" + err.Error())
	}

	n := 0 //정렬 갯수 선정

	for i := 0; i < len(leaderboardInfos.Items)-1; i++ {
		if leaderboardInfos.Items[i].Exp != leaderboardInfos.Items[i+1].Exp && i >= req.RankCount {
			n = i + 1
			break
		}
	}

	if n <= req.RankCount {
		n = len(leaderboardInfos.Items)
	}

	items := leaderboardInfos.Items[:n]

	sort.SliceStable(items, func(i, j int) bool {
		if items[i].Exp > items[j].Exp {
			return true
		} else if items[i].Exp == items[j].Exp {
			if items[i].CreatedAt > items[j].CreatedAt {
				return false
			} else {
				return true
			}
		} else {
			return false
		}
	})

	leaderboard := make([]ExperienceRank, 0)

	for i := 0; i < len(items); i++ {
		if i == req.RankCount {
			break
		}

		d := strings.Split(items[i].Key, "#")
		nftNum, _ := strconv.Atoi(d[3])

		leaderboard = append(leaderboard, ExperienceRank{
			Version:        req.Version,
			Rank:           i + 1,
			Experience:     items[i].Exp,
			NftNum:         nftNum,
			AccountAddress: d[4],
		})
	}

	return &GetExperienceResponse{
		Status:      "ok",
		Leaderboard: leaderboard,
		UpdatedAt:   t,
	}, nil
}

func (r *LeaderboardController) getLevelRank(c *base.Context, req *GetLeaderboardRequest) (*GetLevelResponse, *base.Error) {
	logger.Debugf(c, "get Level Rank Api")

	timeData, _ := r.leaderboardInfos.FindOne(c, bson.D{{"key", fmt.Sprintf("LR")}})

	t := 0

	if timeData != nil {
		t = int(timeData.UpdatedAt)
	} else {
		t = 0
	}

	leaderboardInfos, err := r.indexLevel.Find(c, dynamodb.QueryRequest{
		Key:            fmt.Sprintf("%+v-LR", strconv.Itoa(req.Version)),
		DisableSortKey: true,
	})

	if err != nil {
		logger.Errorf(c, "LevelRank Query Failed")
		return nil, base.ErrUnknown.WithMessage("failed to query level rank" + err.Error())
	}

	n := 0 //정렬 갯수 선정

	for i := 0; i < len(leaderboardInfos.Items)-1; i++ {
		if leaderboardInfos.Items[i].Level != leaderboardInfos.Items[i+1].Level && i >= req.RankCount {
			n = i + 1
			break
		}
	}

	if n <= req.RankCount {
		n = len(leaderboardInfos.Items)
	}

	items := leaderboardInfos.Items[:n]

	sort.SliceStable(items, func(i, j int) bool {
		if items[i].Level > items[j].Level {
			return true
		} else if items[i].Level == items[j].Level {
			if items[i].CreatedAt > items[j].CreatedAt {
				return false
			} else {
				return true
			}
		} else {
			return false
		}
	})

	leaderboard := make([]LevelRank, 0)

	for i := 0; i < len(items); i++ {
		if i == req.RankCount {
			break
		}

		d := strings.Split(items[i].Key, "#")
		nftNum, _ := strconv.Atoi(d[3])

		leaderboard = append(leaderboard, LevelRank{
			Version:        req.Version,
			Rank:           i + 1,
			Level:          items[i].Level,
			NftNum:         nftNum,
			AccountAddress: d[4],
		})
	}

	return &GetLevelResponse{
		Status:      "ok",
		Leaderboard: leaderboard,
		UpdatedAt:   t,
	}, nil
}
