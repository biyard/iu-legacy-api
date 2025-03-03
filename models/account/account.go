package account

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountType int

const (
	Google AccountType = iota + 1
	Kakao
)

func (r AccountType) String() string {
	switch r {
	case Google:
		return "google"
	case Kakao:
		return "kakao"
	default:
		return "unknown"
	}
}

// Account is a struct that contains account information.
type Account struct {
	Key string `bson:"key,omitempty" json:"key,omitempty"`
	CreatedAt int64 `bson:"createdAt,omitempty" json:"createdAt,omitempty"`

	ID      string      `bson:"id,omitempty" json:"id,omitempty"`
	Type    AccountType `bson:"type,omitempty" json:"type,omitempty"`
	Nonce   string      `bson:"nonce,omitempty" json:"nonce,omitempty"`
	Address string      `bson:"address,omitempty" json:"address,omitempty"`
}

func (r Account) Collection() string {
	return "accounts"
}

func (r Account) Table() string {
	return r.Collection()
}

func (r Account) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{"id", 1},
				{"type", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
		{
			Keys: bson.D{
				{"nonce", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
		{
			Keys: bson.D{
				{"address", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	}
}
