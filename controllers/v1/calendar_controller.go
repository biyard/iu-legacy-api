package v1

import (
	"context"
	"fmt"
	"sort"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/models/calendar"
)

type CalendarController struct {
	cfg config.AppConfig

	db            *dynamodb.DynamoDB
	indexCalendar *dynamodb.Index[calendar.Calendar]
}

func NewCalendarController(ctx context.Context, cfg config.AppConfig) *CalendarController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	indexCalendar, err := dynamodb.NewIndex[calendar.Calendar](ctx, db, "yearIndex")
	if err != nil {
		logger.Critical(ctx, err)
	}

	return &CalendarController{
		cfg:           cfg,
		db:            db,
		indexCalendar: indexCalendar,
	}
}

func (r *CalendarController) OnInit(route base.PathRouter) {
	route.GET("/infos", r.getCallendarInfos)
}

type GetCalendarInfoRequest struct {
	Year string `form:"year"`
}

type GetCalendarInfoResponse struct {
	Status string              `json:"status"`
	Items  []calendar.Calendar `json:"items"`
}

func (r *CalendarController) getCallendarInfos(c *base.Context, req *GetCalendarInfoRequest) (*GetCalendarInfoResponse, *base.Error) {
	calendars, err := r.indexCalendar.Find(c, dynamodb.QueryRequest{
		Key: fmt.Sprintf("%s", req.Year),
	})

	if err != nil {
		logger.Errorf(c, "Calendar Data query Failed: %+v", err)
		return nil, base.ErrUnknown.WithMessage("failed to query calendar" + err.Error())
	}

	logger.Debugf(c, "calendar data: %+v", calendars.Items)

	items := calendars.Items

	sort.SliceStable(items, func(i, j int) bool {
		if items[i].Month < items[j].Month {
			return true
		} else {
			return false
		}
	})

	return &GetCalendarInfoResponse{
		Status: "ok",
		Items:  items,
	}, nil
}
