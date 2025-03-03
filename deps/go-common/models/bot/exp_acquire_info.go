package bot

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ExperienceAcquireInfos struct {
	Key               string `bson:"key,omitempty" json:"key,omitempty"` //daily-mission: missionName-startDate-accountAddress, account distribute: contractAddress-transactionHash-accountAddress
	EventName         string `bson:"eventName,omitempty" json:"eventName,omitempty"`
	EventNameEn       string `bson:"eventNameEn,omitempty" json:"eventNameEn,omitempty"`
	MissionRewardDate string `bson:"missionRewardDate,omitempty" json:"missionRewardDate,omitempty"`
	CreatedAt         int64  `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	Experience        int    `bson:"experience,omitempty" json:"experience,omitempty"`
	Account           string `bson:"account,omitempty" json:"account,omitempty"`
	TokenID           int    `bson:"tokenId,omitempty" json:"tokenId,omitempty"`
}

func (r ExperienceAcquireInfos) Collection() string {
	return "experience-acquire-infos"
}

func (r ExperienceAcquireInfos) Table() string {
	return r.Collection()
}

func (r ExperienceAcquireInfos) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys:    bson.D{{"account", 1}},
			Options: options.Index().SetName("indexAccount"),
		},
		{
			Keys:    bson.D{{"tokenId", 6}},
			Options: options.Index().SetName("indexTokenId"),
		},
	}
}
