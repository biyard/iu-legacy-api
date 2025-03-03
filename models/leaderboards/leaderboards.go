package leaderboards

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Leaderboard struct {
	Key              string `bson:"key,omitempty" json:"key,omitempty"`
	RankType         string `bson:"rankType,omitempty" json:"rankType,omitempty"`
	DailyCount       int    `bson:"dailyCount,omitempty" json:"dailyCount,omitempty"`
	Level            int    `bson:"lv,omitempty" json:"lv,omitempty"`
	Exp              int    `bson:"exp,omitempty" json:"exp,omitempty"`
	VotingCount      int    `bson:"votingCount,omitempty" json:"votingCount,omitempty"`
	Version          int    `bson:"version,omitempty" json:"version,omitempty"`
	CreatedAt        int64  `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt        int64  `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	IndexVersionType string `bson:"indexVersionType,omitempty" json:"indexVersionType,omitempty"`
}

func (r Leaderboard) Collection() string {
	return "leaderboards"
}

func (r Leaderboard) Table() string {
	return r.Collection()
}

func (r Leaderboard) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys:    bson.D{{"indexVersionType", 2}},
			Options: options.Index().SetName("sortingDailyCount"),
		},
		{
			Keys:    bson.D{{"indexVersionType", 3}},
			Options: options.Index().SetName("sortingLevel"),
		},
		{
			Keys:    bson.D{{"indexVersionType", 4}},
			Options: options.Index().SetName("sortingExp"),
		},
		{
			Keys:    bson.D{{"indexVersionType", 5}},
			Options: options.Index().SetName("sortingVotingCount"),
		},
	}
}
