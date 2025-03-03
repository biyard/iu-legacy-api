package jira

import (
	"context"
	"net/http"
	"sync"
	"time"

	"biyard.co/common/logger"
	"biyard.co/common/utils/httpclient"
)

var once sync.Once
var cli *JiraAPI

type JiraAPI struct {
	ctx context.Context
	cli *httpclient.HTTPClient
}

func NewJiraAPI(ctx context.Context, endpoint, token string) *JiraAPI {
	once.Do(func() {
		_cli, err := httpclient.NewHttpClient(httpclient.ConstantBackoff(3, 1*time.Second), httpclient.Limiter(1000), httpclient.Headers(http.Header{
			"Content-Type":  []string{"application/json"},
			"Authorization": []string{"Basic " + token},
		}), httpclient.Endpoint(endpoint))

		if err != nil {
			logger.Criticalf(nil, "NewHttpClient error: %v", err)

		}
		cli = &JiraAPI{
			ctx: ctx,
			cli: _cli,
		}
	})

	return cli
}

func (r *JiraAPI) ToMention(id string) string {
	return "<@" + id + ">"
}

func (r *JiraAPI) ToLink(link, text string) string {
	return "<" + link + "|" + text + ">"
}

type JiraIssueResponse struct {
	Expand    string      `json:"expand"`
	ID        string      `json:"id"`
	Self      string      `json:"self"`
	Key       string      `json:"key"`
	MaxResult int         `json:"maxResults"`
	Total     int         `json:"total"`
	Issues    []JiraIssue `json:"issues"`
}

type JiraIssue struct {
	Expand string          `json:"expand"`
	ID     string          `json:"id"`
	Self   string          `json:"self"`
	Key    string          `json:"key"`
	Fields JiraIssueFields `json:"fields"`
}

type JiraIssueAssignee struct {
	Self        string `json:"self"`
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
}

type JiraIssueFields struct {
	Summary   string            `json:"summary"`
	DueDate   string            `json:"duedate"`
	Status    JiraIssueStatus   `json:"status"`
	Assignee  JiraIssueAssignee `json:"assignee"`
	StartDate string            `json:"customfield_10015"`
}

type JiraIssueStatus struct {
	Self        string `json:"self"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

type JiraIssueByJQLRequest struct {
	Expand       []string `json:"expand"`
	Fileds       []string `json:"fields"`
	JQL          string   `json:"jql"`
	MaxResult    int      `json:"maxResults"`
	StartAt      int      `json:"startAt"`
	FieldsByKeys bool     `json:"fieldsByKeys"`
}

func (r *JiraAPI) GetIssueByJQL(ctx context.Context, jql string) (*JiraIssueResponse, error) {
	var res JiraIssueResponse
	code, e := r.cli.Post(ctx, "/rest/api/3/search", &JiraIssueByJQLRequest{
		Expand:    []string{"names", "schema"},
		Fileds:    []string{"summary", "duedate", "status", "assignee", "customfield_10015"},
		JQL:       jql,
		MaxResult: 100,
		StartAt:   0,
	}, &res)
	if e != nil {
		logger.Errorf(ctx, "r.jira.Get error(%v): %v", code, e.Error())
		return nil, e
	}

	return &res, nil
}
