package v1

import (
	"context"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/encrypt"
)

type EncryptController struct {
	db       *dynamodb.DynamoDB
	encrypts *dynamodb.Table[encrypt.ManagerEncryt]
}

func NewEncryptController(ctx context.Context, cfg config.AppConfig) *EncryptController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	encrypts, err := dynamodb.NewTable[encrypt.ManagerEncryt](ctx, db)
	if err != nil {
		panic(err)
	}

	logger.Debugf(ctx, "encrypt: %+v", encrypts)

	return &EncryptController{
		db:       db,
		encrypts: encrypts,
	}
}

type PasswordRequest struct {
	Password string `json:"password"`
}

type PasswordResponse struct {
	Password string `json:"password"`
}

type GetPasswordRequest struct{}

func (r *EncryptController) OnInit(route base.PathRouter) {
	route.POST("/", r.sendPasswordData)
	// route.GET("/", r.checkPasswordData)
	// route.PUT("/", r.changePasswordData)
}

func (r *EncryptController) sendPasswordData(c *base.Context, req *PasswordRequest) (*PasswordResponse, *base.Error) {
	err := r.encrypts.Insert(c, &encrypt.ManagerEncryt{
		Key:      "manager",
		Password: req.Password,
		Type:     "manager",
	})
	if err != nil {
		logger.Errorf(c, "encrypt.ManagerEncryt error: %v", err)
		return nil, errors.ErrInsertFailed
	}

	return &PasswordResponse{
		Password: req.Password,
	}, nil
}

// func (r *EncryptController) checkPasswordData(c *base.Context, req *PasswordRequest) (*PasswordResponse, *base.Error) {
// 	d, e := r.encrypts.FindOne(c, bson.D{{Key: "key", Value: "roulette"}})

// 	return &PasswordResponse{
// 		Password: req.Password,
// 	}, nil
// }
