package google

import (
	"context"
	"net/http"
	"sync"
	"time"

	"google.golang.org/api/oauth2/v2"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/common/utils/httpclient"
)

type GoogleApi struct {
	ctx        context.Context
	cli        *httpclient.HTTPClient
	httpClient *http.Client
}

var cli *GoogleApi
var once sync.Once

func NewGoogleApi(ctx context.Context, endpoint string) *GoogleApi {
	once.Do(func() {
		cl, err := httpclient.NewHttpClient(httpclient.ConstantBackoff(3, 1*time.Second), httpclient.Limiter(1000), httpclient.Headers(http.Header{
			"Content-Type": []string{"application/json"},
		}), httpclient.Endpoint(endpoint))

		if err != nil {
			logger.Criticalf(nil, "NewHttpClient error: %v", err)

		}
		cli = &GoogleApi{
			ctx:        ctx,
			cli:        cl,
			httpClient: &http.Client{},
		}
	})

	return cli
}

// GetTokenInfo returns OpenID information of a GetTokenInfo
func (r *GoogleApi) GetTokenInfo(ctx *base.Context, req GetTokenInfoRequest) (*oauth2.Tokeninfo, *base.Error) {
	logger.Infof(ctx, "GetTokenInfo %+v", req)
	// var resp GetTokenInfoResponse
	// if code, err := r.cli.Get(ctx, "/tokeninfo", req, &resp); err != nil {
	// 	logger.Errorf(ctx, "GetTokenInfo error(%v): %v", code, err)
	// 	return nil, base.ErrUnknown.WithMessage(err.Error())
	// }
	oauth2Service, err := oauth2.New(r.httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(req.IDToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		logger.Errorf(ctx, "oauth2Service.Tokeninfo error: %v", err)
		return nil, base.ErrUnknown.WithMessage(err.Error())
	}
	logger.Infof(ctx, "tokenInfo: %+v", tokenInfo)
	return tokenInfo, nil
}
