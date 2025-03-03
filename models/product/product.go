package product

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RouletteProduct struct {
	Key                           string   `bson:"key,omitempty" json:"key,omitempty"`
	Index                         int      `bson:"index,omitempty" json:"index,omitempty"`
	ProductList                   []string `bson:"productList,omitempty" json:"productList,omitempty"` // product의 인덱스 리스트로 삽입
	ProductCount                  []int    `bson:"productCount,omitempty" json:"productCount,omitempty"`
	ProductRouletteOnsiteSequence []string `bson:"productOnsiteSequence,omitempty" json:"productOnsiteSequence,omitempty"`
	ProductRouletteOnlineSequence []string `bson:"productOnlineSequence,omitempty" json:"productOnlineSequence,omitempty"`
}

func (r RouletteProduct) Collection() string {
	return "products"
}

func (r RouletteProduct) Table() string {
	return r.Collection()
}

func (r RouletteProduct) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{"index", 1},
				{"productList", 1},
			},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	}
}
