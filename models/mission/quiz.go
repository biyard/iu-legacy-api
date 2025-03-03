package mission

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Quiz struct {
	Key                string     `bson:"key,omitempty" json:"key,omitempty"` //missionId
	MissionNameKo      string     `bson:"missionNameKo,omitempty" json:"missionNameKo,omitempty"`
	MissionNameEn      string     `bson:"missionNameEn,omitempty" json:"missionNameEn,omitempty"`
	MissionType        string     `bson:"missionType,omitempty" json:"missionType,omitempty"`
	Exp                int        `bson:"exp,omitempty" json:"exp,omitempty"`
	StartDate          int        `bson:"startDate,omitempty" json:"startDate,omitempty"`
	Questions          []string   `bson:"questions,omitempty" json:"questions,omitempty"`
	MultipleChoiceView [][]string `bson:"multipleChoiceView,omitempty" json:"multipleChoiceView,omitempty"`
	Answers            []int      `bson:"answers,omitempty" json:"answers,omitempty"`
	Cutline            int        `bson:"cutline,omitempty" json:"cutline,omitempty"`
	CreatedAt          int64      `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	StartId            int64      `bson:"startId,omitempty" json:"startId,omitempty"`

	IndexStartDate string `bson:"indexStartDate,omitempty" json:"indexStartDate,omitempty"` //lang-startDate

	Address string `bson:"address,omitempty" json:"address,omitempty"`
	Lang    string `bson:"lang,omitempty" json:"lang,omitempty"`
}

func (r Quiz) Collection() string {
	return "missionDatas"
}

func (r Quiz) Table() string {
	return r.Collection()
}

func (r Quiz) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys:    bson.D{{"indexStartDate", 1}},
			Options: options.Index().SetName("indexStartDate"),
		},
		{
			Keys:    bson.D{{"address", 1}},
			Options: options.Index().SetName("address"),
		},
	}
}
