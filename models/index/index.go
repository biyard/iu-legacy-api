package index

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RouletteIndex struct {
	Key      string `bson:"key,omitempty" json:"key,omitempty"`
	Type     string `bson:"type,omitempty" json:"type,omitempty"`
	MaxIndex int    `bson:"maxIndex,omitempty" json:"maxIndex,omitempty"`
}

func (r RouletteIndex) Collection() string {
	return "indexes"
}

func (r RouletteIndex) Table() string {
	return r.Collection()
}

func (r RouletteIndex) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{"key", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
		{
			Keys: bson.D{
				{"type", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
		{
			Keys: bson.D{
				{"maxIndex", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	}
}
