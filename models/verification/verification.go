package verification

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Verification struct {
	Key             string `bson:"key,omitempty" json:"key,omitempty"`
	VerificationKey string `bson:"verificationKey,omitempty" json:"verificationKey,omitempty"`
	IsUsed          string `bson:"isUsed,omitempty" json:"isUsed,omitempty"`
}

func (r Verification) Collection() string {
	return "verifications"
}

func (r Verification) Table() string {
	return r.Collection()
}

func (r Verification) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "verificationKey", Value: 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	}
}
