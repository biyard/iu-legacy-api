package firebase

import (
	"context"
	"sync"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type Firebase struct {
	ctx context.Context
	app *firebase.App
}

var cli *Firebase
var once sync.Once

func NewFirebase(ctx context.Context, credentialPath string) *Firebase {
	once.Do(func() {
		app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(credentialPath))
		if err != nil {
			logger.Criticalf(nil, "error initializing app: %v\n", err)
		}
		if err != nil {
			logger.Criticalf(nil, "NewHttpClient error: %v", err)

		}
		cli = &Firebase{
			ctx: ctx,
			app: app,
		}
	})

	return cli
}

// GetTokenInfo returns OpenID information of a GetTokenInfo
func (r *Firebase) GetTokenInfo(ctx context.Context, req string) (*auth.Token, *base.Error) {
	logger.Debugf(ctx, "GetTokenInfo %+v", req)
	cli, err := r.app.Auth(ctx)
	if err != nil {
		logger.Errorf(ctx, "error getting Auth client: %v\n", err)
		return nil, base.ErrFirebaseAuthClient
	}

	token, err := cli.VerifyIDToken(ctx, req)
	if err != nil {
		logger.Errorf(ctx, "error verifying ID token: %v\n", err)
		return nil, base.ErrFirebaseVerifyIDToken
	}
	logger.Debugf(ctx, "tokenInfo: %+v", token)

	return token, nil
}
