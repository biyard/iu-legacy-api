package roulette_user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RouletteUser struct {
	Key     string `bson:"key,omitempty" json:"key,omitempty"`
	Account string `bson:"account,omitempty" json:"account,omitempty"`

	Product         string `bson:"product,omitempty" json:"product,omitempty"`
	Count           int    `bson:"count,omitempty" json:"count,omitempty"`
	ParticipantType string `bson:"type,omitempty" json:"type,omitempty"`
	IsAction        int    `bson:"isAction,omitempty" json:"isAction,omitempty"` // 1: complete, 2: not complete
}

func (r RouletteUser) Collection() string {
	return "prizes"
}

func (r RouletteUser) Table() string {
	return r.Collection()
}

func (r RouletteUser) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "account", Value: 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	}
}
