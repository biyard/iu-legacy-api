package md

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Md struct {
	Key       string `bson:"key,omitempty" json:"key,omitempty"`
	KoreaMd   MdData `bson:"ko" json:"ko"`
	EnglishMd MdData `bson:"en" json:"en"`
	StartTime int64  `bson:"startTime" json:"startTime"`
	EndTime   int64  `bson:"endTime" json:"endTime"`
}

type MdData struct {
	Title    string `bson:"title" json:"title"`
	Contents string `bson:"contents" json:"contents"`
}

func (r Md) Collection() string {
	return "mds"
}

func (r Md) Table() string {
	return r.Collection()
}

func (r Md) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{"key", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	}
}
