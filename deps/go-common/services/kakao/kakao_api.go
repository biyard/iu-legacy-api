package kakao

import (
	"context"
	"net/http"
	"sync"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/common/utils/httpclient"
)

type KakaoAPI struct {
	ctx context.Context
	cli *httpclient.HTTPClient
}

var cli *KakaoAPI
var once sync.Once

func NewKakaoAPI(ctx context.Context, endpoint string) *KakaoAPI {
	once.Do(func() {
		kcli, err := httpclient.NewHttpClient(httpclient.ConstantBackoff(3, 1*time.Second), httpclient.Limiter(1000), httpclient.Headers(http.Header{
			"Content-Type": []string{"application/x-www-form-urlencoded;charset=utf-8"},
		}), httpclient.Endpoint(endpoint))

		if err != nil {
			logger.Criticalf(nil, "NewHttpClient error: %v", err)

		}
		cli = &KakaoAPI{
			ctx: ctx,
			cli: kcli,
		}
	})

	return cli
}

func (r *KakaoAPI) GetAccessToken(ctx context.Context, req KakaoOAuthTokenRequest) (*KakaoOAuthTokenResponse, *base.Error) {
	var res KakaoOAuthTokenResponse
	code, e := r.cli.Post(ctx, "/oauth/token", req, &res)
	if e != nil {
		logger.Errorf(ctx, "r.kakao.Post error: %v", e)
		return nil, base.ErrKakaoInvalieCode
	}

	logger.Debugf(ctx, "KakaoOAuthTokenStatus %d %+v", code, e, res)

	return &res, nil
}

func (r *KakaoAPI) GetTokenInfo(ctx context.Context, req KakaoOpenIDTokenInfoRequest) (*KakaoOpenIDTokenInfoResponse, *base.Error) {
	logger.Debugf(ctx, "GetTokenInfo request: %+v", req)
	var res KakaoOpenIDTokenInfoResponse
	code, err := r.cli.Post(ctx, "/oauth/tokeninfo", req, &res)
	if err != nil {
		logger.Errorf(ctx, "failed to get token info %v", err)
		return nil, base.ErrKakaoTokenInfo
	}
	logger.Debugf(ctx, "KakaoOAuthTokenInforResponse %d %v", code, res)

	return &res, nil
}
