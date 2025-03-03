package v1

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/common/services/firebase"
	"biyard.co/common/services/kakao"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/account"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/crypto/sha3"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthController struct {
	redirectUri map[string]bool
	kakao       *kakao.KakaoAPI
	firebase    *firebase.Firebase

	cfg config.AppConfig

	// db       *mongodb.MongoDB
	// accounts *mongodb.Collection[account.Account]
	db       *dynamodb.DynamoDB
	accounts *dynamodb.Table[account.Account]
}

func NewAuthController(ctx context.Context, cfg config.AppConfig) *AuthController {
	// db := mongodb.New(ctx, cfg.Database.Endpoint, cfg.Database.Name)
	// accounts, err := mongodb.NewCollection[account.Account](ctx, db)
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	accounts, err := dynamodb.NewTable[account.Account](ctx, db)
	if err != nil {
		panic(err)
	}

	logger.Debugf(ctx, "accounts: %+v", accounts)
	m := map[string]bool{}
	for _, v := range cfg.OAuth.RedirectURI {
		m[v] = true
	}

	return &AuthController{
		redirectUri: m,
		kakao:       kakao.NewKakaoAPI(ctx, cfg.OAuth.Kakao.Endpoint),
		firebase:    firebase.NewFirebase(ctx, cfg.OAuth.Google.CredentialFile),
		cfg:         cfg,
		db:          db,
		accounts:    accounts,
	}
}

func (r *AuthController) OnInit(route base.PathRouter) {
	route.GET("", r.CallbackOAuth)
	route.POST("/address", r.NotifyAddress)
}

type NotifyAddressRequest struct {
	Address string `json:"address"`
	ID      string `json:"id"`
}

type NotifyAddressResponse struct {
	Status string `json:"status"`
}

// NotifyAddress sets the address of the user.
func (r *AuthController) NotifyAddress(ctx *base.Context, req NotifyAddressRequest) (*NotifyAddressResponse, *base.Error) {
	logger.Infof(ctx, "NotifyAddress %+v", req)
	err := r.accounts.UpdateOne(ctx, bson.D{{"key", req.ID}}, bson.D{{"$set", bson.D{{"address", req.Address}}}})
	if err != nil {
		logger.Errorf(ctx, "accounts.UpdateOne error: %v", err)
		return nil, errors.ErrAccountAlreadyExists
	}

	return &NotifyAddressResponse{Status: "OK"}, nil
}

type CallbackOAuthRequest struct {
	Provider    string `form:"provider"`
	Code        string `form:"code"`
	RedirectURI string `form:"redirect-uri"`
}

type CallbackOAuthResponse struct {
	AddressHint    string `json:"addressHint"`
	PrivateKeyHint string `json:"privateKeyHint"`
	RestoreKey     string `json:"restoreKey"`
	Address        string `json:"address,omitempty"`
	ID             string `json:"id"`
}

// CallbackOAuth provides a callback for Kakao OAuth.
func (r *AuthController) CallbackOAuth(ctx *base.Context, req CallbackOAuthRequest) (*CallbackOAuthResponse, *base.Error) {
	logger.Infof(ctx, "CallbackKakaoOAuth %+v", req)
	handle := r.getHintsByGoogle

	if req.Provider == account.Kakao.String() {
		handle = r.getHintsByKakao
	}
	addressHint, privateKeyHint, address, nonce, err := handle(ctx, req)

	logger.Debugf(ctx, "Address Hint %+v", addressHint)
	logger.Debugf(ctx, "PrivateKey Hint %+v", privateKeyHint)

	return &CallbackOAuthResponse{
		AddressHint:    addressHint,
		PrivateKeyHint: privateKeyHint,
		RestoreKey:     r.cfg.Hint.Account.PublicKey,
		Address:        address,
		ID:             nonce,
	}, err
}

// KakaoOAuthToken provides a token for Kakao OAuth.
func (r *AuthController) getHintsByKakao(ctx *base.Context, req CallbackOAuthRequest) (addressHint, privateKeyHint, address, nonce string, err *base.Error) {
	// resp, e := r.kakao.GetAccessToken(ctx, kakao.KakaoOAuthTokenRequest{
	// 	GrantType:   "authorization_code",
	// 	Code:        req.Code,
	// 	RedirectUri: req.RedirectURI,
	// 	ClientId:    r.cfg.OAuth.Kakao.ClientID,
	// })
	// if e != nil {
	// 	err = e
	// 	return
	// }

	idTokenResp, e := r.kakao.GetTokenInfo(ctx, kakao.KakaoOpenIDTokenInfoRequest{
		IDToken: req.Code,
	})
	if e != nil {
		err = e
		return
	}

	logger.Debugf(ctx, "KakaoOpenIDTokenInfoResponse %+v", idTokenResp)

	return r.getHints(ctx, account.Kakao, idTokenResp.Issuer, idTokenResp.Subject)
}

func (r *AuthController) getHintsByGoogle(ctx *base.Context, req CallbackOAuthRequest) (addressHint, privateKeyHint, address, nonce string, err *base.Error) {
	info, e := r.firebase.GetTokenInfo(ctx, req.Code)
	if e != nil {
		err = e
		return
	}

	return r.getHints(ctx, account.Google, info.Audience, (info.Firebase.Identities["google.com"].([]any))[0].(string))
}

func (r *AuthController) getHints(ctx *base.Context, typ account.AccountType, issuer, subject string) (string, string, string, string, *base.Error) {
	logger.Debugf(ctx, "getHints issuer: %+v", issuer)
	addressHint := fmt.Sprintf("%s-%s", typ.String(), subject)
	address := ""
	key := ""

	acc, _ := r.accounts.FindOne(ctx, bson.D{{"key", fmt.Sprintf("%s#%s", typ, subject)}})
	nonce := ""
	if acc != nil {
		nonce = acc.Nonce
		key = acc.Key
	}
	logger.Debugf(ctx, "Nonce %+v", nonce)
	unit := 64
	rnd := make([]byte, unit)
	if nonce == "" {
		if n, e := rand.Read(rnd); e != nil || n != unit {
			logger.Errorf(ctx, "rand.Read error: %v", e)
			return "", "", "", "", errors.ErrRandomGeneration
		}
		nonce = base64.StdEncoding.EncodeToString(rnd)
	} else {
		address = acc.Address
		rnd, _ = base64.StdEncoding.DecodeString(nonce)
	}

	sha3Hash := sha3.NewKeccak512()
	logger.Debugf(ctx, "hints: %s %s %s", r.cfg.Hint.Secret, nonce, addressHint)

	hint := sha3Hash.Sum([]byte(r.cfg.Hint.Secret + nonce + addressHint))

	privateKeyHint := hexutil.Encode(hint)

	if acc == nil || acc.Nonce == "" {
		key = fmt.Sprintf("%s#%s", typ, subject)
		dt := account.Account{
			Key:       key,
			CreatedAt: time.Now().Unix(),
			ID:        subject,
			Type:      typ,
			Nonce:     nonce,
		}
		r.accounts.Insert(ctx, &dt)
	}

	return addressHint, privateKeyHint, address, key, nil
}
