package md_account

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MdAccount struct {
	Key     string `bson:"key,omitempty" json:"key,omitempty"`
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Check   int    `bson:"check,omitempty" json:"check,omitempty"`
}

func (r MdAccount) Collection() string {
	return "md_accounts"
}

func (r MdAccount) Table() string {
	return r.Collection()
}

func (r MdAccount) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{"key", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	}
}
