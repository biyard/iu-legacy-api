package encrypt

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ManagerEncryt struct {
	Key      string `bson:"key,omitempty" json:"key,omitempty"`
	Type     string `bson:"type,omitempty" json:"type,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}

func (r ManagerEncryt) Collection() string {
	return "encrypts"
}

func (r ManagerEncryt) Table() string {
	return r.Collection()
}

func (r ManagerEncryt) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{"type", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	}
}
