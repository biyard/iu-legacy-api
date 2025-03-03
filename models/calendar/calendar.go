package calendar

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Calendar struct {
	Key string `bson:"key,omitempty" json:"key,omitempty"`

	Year      int64    `bson:"year,omitempty" json:"year,omitempty"`
	YearStr   string   `bson:"yearStr,omitempty" json:"yearStr,omitempty"`
	Month     int64    `bson:"month,omitempty" json:"month,omitempty"`
	Title     string   `bson:"title,omitempty" json:"title,omitempty"`
	TitleEn   string   `bson:"titleEn,omitempty" json:"titleEn,omitempty"`
	Events    []string `bson:"events,omitempty" json:"events,omitempty"`
	CreatedAt int64    `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func (r Calendar) Collection() string {
	return "calendars"
}

func (r Calendar) Table() string {
	return r.Collection()
}

func (r Calendar) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys:    bson.D{{"yearStr", 1}},
			Options: options.Index().SetName("yearIndex"),
		},
	}
}
