package m1

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/common/utils/klaytn"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/roulette_user.go"
	"biyard.co/iuniverse-api/models/verification"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"go.mongodb.org/mongo-driver/bson"

	"biyard.co/common/base"
)

// charset use random string
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type VerificationController struct {
	db            *dynamodb.DynamoDB
	verifications *dynamodb.Table[verification.Verification]
	users         *dynamodb.Table[roulette_user.RouletteUser]
	endpoint      string
	feeEndpoint   string
	contract      []string
	chainID       int

	feepayer *klaytn.Wallet
	owner    *klaytn.Wallet
}

func NewVerificationController(ctx context.Context, cfg config.AppConfig) *VerificationController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	verifications, err := dynamodb.NewTable[verification.Verification](ctx, db)
	if err != nil {
		panic(err)
	}

	users, err := dynamodb.NewTable[roulette_user.RouletteUser](ctx, db)
	if err != nil {
		panic(err)
	}

	logger.Debugf(ctx, "verifications: %+v", verifications)

	f, err := klaytn.NewKlaytnWallet(cfg.Feepayer.Address, cfg.Feepayer.PrivateKey)
	if err != nil {
		logger.Critical(ctx, "invalid private key")
	}

	o, err := klaytn.NewKlaytnWallet(cfg.Owner.Address, cfg.Owner.PrivateKey)
	if err != nil {
		logger.Critical(ctx, "invalid private key")
	}

	o.SetABI(cfg.ABI).SetEndpoint(cfg.Klaytn.Endpoint).SetChainID(int64(cfg.Klaytn.ChainID)).SetFeepayer(cfg.Feepayer.Address, cfg.Feepayer.PrivateKey)

	return &VerificationController{
		db:            db,
		verifications: verifications,
		users:         users,
		endpoint:      cfg.Klaytn.Endpoint,
		feeEndpoint:   cfg.Services.Endpoint,
		contract:      cfg.Contracts,
		chainID:       cfg.Klaytn.ChainID,
		feepayer:      f,
		owner:         o,
	}
}

type GetVerificationDataRequest struct{}

type GetVerificationDataResponse struct {
	Echo string `json:"result"`
}

type GetTokenDataRequest struct {
	Address         string `json:"address"`
	VerificationKey string `json:"verificationKey"`
}

type SignTransactionResponse struct {
	Address   string         `json:"address"`
	TxReceipt *types.Receipt `json:"txReceipt"`
}

type GetTokenDataResponse struct {
	Echo *SignTransactionResponse `json:"result"`
}

func (r *VerificationController) OnInit(route base.PathRouter) {
	route.POST("", r.createCode)
	route.POST("/verify", r.verifyCode)
	route.POST("/reset", r.resetPlay)
	route.POST("/action", r.actionPlay)
}

func (r *VerificationController) actionPlay(c *base.Context, req *GetTokenDataRequest) (*GetVerificationDataResponse, *base.Error) {
	dt, err := r.verifications.FindOne(c, bson.D{{Key: "key", Value: req.VerificationKey}})
	if err != nil {
		logger.Errorf(c, "verification.FindOne error: %v", err)
		return nil, errors.ErrVerificationNotExists
	}

	if dt.IsUsed == "used" {
		return nil, errors.ErrVerificationIsUsed
	}

	_, e := r.users.FindOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", req.Address, "onsite")}})
	_, e2 := r.users.FindOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", req.Address, "online")}})

	if e != nil && e2 != nil {
		logger.Errorf(c, "accounts.FindOne error: %v", err)
		return nil, errors.ErrUserEmpty
	}

	errUpdate := r.verifications.UpdateOne(c, bson.D{{"key", req.VerificationKey}}, bson.D{{"$set", bson.D{{"isUsed", "used"}}}})
	if errUpdate != nil {
		logger.Errorf(c, "accounts.UpdateOne error: %v", err)
		return nil, errors.ErrVerificationAlreadyExists
	}

	if e == nil {
		errUpdateUser := r.users.UpdateOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", req.Address, "onsite")}}, bson.D{{"$set", bson.D{{"isAction", 1}}}})
		if errUpdateUser != nil {
			logger.Errorf(c, "accounts.UpdateOne error: %v", err)
			return nil, errors.ErrFailedToUpdateAddress
		}
	} else if e2 == nil {
		errUpdateUser := r.users.UpdateOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", req.Address, "online")}}, bson.D{{"$set", bson.D{{"isAction", 1}}}})
		if errUpdateUser != nil {
			logger.Errorf(c, "accounts.UpdateOne error: %v", err)
			return nil, errors.ErrFailedToUpdateAddress
		}
	}

	return &GetVerificationDataResponse{
		Echo: "successed",
	}, nil
}

// unused, used
func (r *VerificationController) createCode(c *base.Context, req *GetVerificationDataRequest) (*GetVerificationDataResponse, *base.Error) {
	str := stringWithCharset(16, charset)

	err := r.verifications.Insert(c, &verification.Verification{
		Key:             str,
		VerificationKey: str,
		IsUsed:          "unused",
	})
	if err != nil {
		logger.Errorf(c, "verification.InsertOne error: %v", err)
		return nil, errors.ErrVerificationAlreadyExists
	}

	return &GetVerificationDataResponse{
		Echo: str,
	}, nil
}

func (r *VerificationController) resetPlay(c *base.Context, req *GetTokenDataRequest) (*GetVerificationDataResponse, *base.Error) {
	dt, err := r.verifications.FindOne(c, bson.D{{Key: "key", Value: req.VerificationKey}})
	if err != nil {
		logger.Errorf(c, "verification.FindOne error: %v", err)
		return nil, errors.ErrVerificationNotExists
	}

	if dt.IsUsed == "used" {
		return nil, errors.ErrVerificationIsUsed
	}

	_, e := r.users.FindOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s", req.Address)}})

	if e == nil {
		err := r.users.DeleteOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s", req.Address)}})
		if err != nil {
			logger.Errorf(c, "users.deleteOne error: %v", err)
			return &GetVerificationDataResponse{
				Echo: "failed",
			}, errors.ErrDeleteFailed
		}
	}

	errUpdate := r.verifications.UpdateOne(c, bson.D{{"key", req.VerificationKey}}, bson.D{{"$set", bson.D{{"isUsed", "used"}}}})
	if errUpdate != nil {
		logger.Errorf(c, "accounts.UpdateOne error: %v", err)
		return nil, errors.ErrVerificationAlreadyExists
	}

	return &GetVerificationDataResponse{
		Echo: "successed",
	}, nil
}

func (r *VerificationController) verifyCode(c *base.Context, req *GetTokenDataRequest) (*GetTokenDataResponse, *base.Error) {
	key := req.VerificationKey
	address := req.Address

	dt, err := r.verifications.FindOne(c, bson.D{{Key: "key", Value: key}})
	if err != nil {
		logger.Errorf(c, "verification.FindOne error: %v", err)
		return nil, errors.ErrVerificationNotExists
	}

	if dt.IsUsed == "used" {
		return nil, errors.ErrVerificationIsUsed
	}

	receipt, err := r.owner.SendRawTransactionWithFeepayer(c, r.feepayer, r.contract[1], "mintSBT", common.HexToAddress(address), big.NewInt(1), big.NewInt(1), []byte(""))
	_ = receipt

	return &GetTokenDataResponse{
		Echo: &SignTransactionResponse{
			Address:   r.owner.Address.String(),
			TxReceipt: receipt,
		},
	}, nil
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// stringWithCharset return of random string
func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
