package slack

import (
	"context"
	"net/http"
	"sync"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/common/utils/httpclient"
)

var once sync.Once
var cli *SlackAPI

type SlackAPI struct {
	ctx context.Context
	cli *httpclient.HTTPClient
}

func NewSlackAPI(ctx context.Context, endpoint string) *SlackAPI {
	once.Do(func() {
		_cli, err := httpclient.NewHttpClient(httpclient.ConstantBackoff(3, 1*time.Second), httpclient.Limiter(1000), httpclient.Headers(http.Header{
			"Content-Type": []string{"application/json"},
		}), httpclient.Endpoint(endpoint))

		if err != nil {
			logger.Criticalf(nil, "NewHttpClient error: %v", err)

		}
		cli = &SlackAPI{
			ctx: ctx,
			cli: _cli,
		}
	})

	return cli
}

func (r *SlackAPI) SendMessage(msg string) *base.Error {
	code, err := r.cli.Post(r.ctx, "", map[string]string{
		"text": msg,
	}, nil)
	if err != nil {
		logger.Errorf(r.ctx, "failed to send message %v", err)
		return base.ErrUnknown.WithMessage("failed to send message; " + err.Error())
	}

	logger.Debugf(r.ctx, "SlackSendMessageStatus %d %+v", code, err)

	return nil
}

func (r *SlackAPI) SendMessageByMarkdown(msg any) *base.Error {
	code, err := r.cli.Post(r.ctx, "", map[string]any{
		"text": map[string]any{
			"type": "mrkdwn",
			"text": msg,
		},
	}, nil)
	if err != nil {
		logger.Errorf(r.ctx, "failed to send message %v", err)
		return base.ErrUnknown.WithMessage("failed to send message; " + err.Error())
	}

	logger.Debugf(r.ctx, "SlackSendMessageStatus %d %+v", code, err)

	return nil
}
