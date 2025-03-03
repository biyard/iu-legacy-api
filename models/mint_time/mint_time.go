package mint_time

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Time struct {
	Key                string `bson:"key,omitempty" json:"key,omitempty"`
	WhitelistStartTime int64  `bson:"whitelistStartTime,omitempty" json:"whitelistStartTime,omitempty"`
	WhitelistEndTime   int64  `bson:"whitelistEndTime,omitempty" json:"whitelistEndTime,omitempty"`
	PublicStartTime    int64  `bson:"publicStartTime,omitempty" json:"publicStartTime,omitempty"`
	PublicEndTime      int64  `bson:"publicEndTime,omitempty" json:"publicEndTime,omitempty"`
}

func (r Time) Collection() string {
	return "times"
}

func (r Time) Table() string {
	return r.Collection()
}

func (r Time) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{"key", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	}
}
