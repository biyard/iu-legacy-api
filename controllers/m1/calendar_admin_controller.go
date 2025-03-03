package m1

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	errors "biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/calendar"
	"go.mongodb.org/mongo-driver/bson"
)

type CalendarAdminController struct {
	db        *dynamodb.DynamoDB
	calendars *dynamodb.Table[calendar.Calendar]
}

func NewCalendarAdminController(ctx context.Context, cfg config.AppConfig) *CalendarAdminController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)

	calendars, err := dynamodb.NewTable[calendar.Calendar](ctx, db)

	if err != nil {
		logger.Critical(ctx, err)
	}

	return &CalendarAdminController{
		db:        db,
		calendars: calendars,
	}
}

func (r *CalendarAdminController) OnInit(route base.PathRouter) {
	route.POST("", r.addCalendarInfos)
}

type AddCalendarInfoRequest struct {
	Year    int      `json:"year"`
	Month   int      `json:"month"`
	Title   string   `json:"title"`
	TitleEn string   `json:"titleEn"`
	Events  []string `json:"events"`
}

type AddCalendarInfoResponse struct {
	Status string `json:"status"`
}

func (r *CalendarAdminController) addCalendarInfos(ctx *base.Context, req *AddCalendarInfoRequest) (*AddCalendarInfoResponse, *base.Error) {
	now := time.Now()
	sec := now.Unix()

	calendarData, _ := r.calendars.FindOne(ctx, bson.D{{"key", fmt.Sprintf("%s-%s", strconv.Itoa(req.Year), strconv.Itoa(req.Month))}})

	if calendarData == nil {
		e := r.calendars.Insert(ctx, &calendar.Calendar{
			Key:       fmt.Sprintf("%s-%s", strconv.Itoa(req.Year), strconv.Itoa(req.Month)),
			Year:      int64(req.Year),
			YearStr:   strconv.Itoa(req.Year),
			Month:     int64(req.Month),
			Title:     req.Title,
			TitleEn:   req.TitleEn,
			Events:    req.Events,
			CreatedAt: sec,
		})

		if e != nil {
			logger.Errorf(ctx, "year: %+v month: %+v title: %+v events: %+v err: %+v", req.Year, req.Month, req.Title, req.Events, e)
			return nil, errors.ErrInsertFailed
		}
	} else {
		e := r.calendars.UpdateOne(ctx, bson.D{{"key", fmt.Sprintf("%s-%s", strconv.Itoa(req.Year), strconv.Itoa(req.Month))}}, bson.D{{"$set", bson.D{{"title", req.Title}, {"titleEn", req.TitleEn}, {"events", req.Events}, {"createdAt", sec}}}})

		if e != nil {
			logger.Errorf(ctx, "calendarData update Failed %+v %+v %+v", req.Title, req.Events, e)
			return nil, errors.ErrUpdateFailed
		}
	}

	return &AddCalendarInfoResponse{
		Status: "ok",
	}, nil
}
