package v1

import (
	"context"
	"fmt"
	"strconv"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
	"biyard.co/iuniverse-api/models/index"
	"biyard.co/iuniverse-api/models/product"
	"biyard.co/iuniverse-api/models/roulette_user.go"
	"go.mongodb.org/mongo-driver/bson"
)

type UserController struct {
	db       *dynamodb.DynamoDB
	users    *dynamodb.Table[roulette_user.RouletteUser]
	indexes  *dynamodb.Table[index.RouletteIndex]
	products *dynamodb.Table[product.RouletteProduct]
	cfg      config.AppConfig
}

type GetProductRequest struct {
	ParticipantType string `json:"type"` // onsite,online
	Address         string `json:"address"`
	ProductName     string `json:"product"`
}

type GetProductResponse struct {
	Address         string `json:"address"`
	ProductName     string `json:"product"`
	Count           int    `json:"count"`
	ParticipantType string `json:"type"`
	IsAction        bool   `json:"isAction"`
}

type GetProductDtRequest struct {
	Address         string `form:"address"`
	ParticipantType string `form:"type"`
}

type PostDataRequest struct {
	Code                          string   `json:"code"`
	ProductList                   []string `json:"productList"`
	ProductCount                  []int    `json:"productCount"`
	ProductRouletteOnsiteSequence []string `json:"productOnsiteSequence"` // roullet 배치 순서(현장)
	ProductRouletteOnlineSequence []string `json:"productOnlineSequence"` // roullet 배치 순서(온라인)
}

type PostDataResponse struct {
	Index                         int      `json:"index"`
	ProductList                   []string `json:"productList"`
	ProductCount                  []int    `json:"productCount"`
	ProductRouletteOnsiteSequence []string `json:"productOnsiteSequence"` // roullet 배치 순서(현장)
	ProductRouletteOnlineSequence []string `json:"productOnlineSequence"` // roullet 배치 순서(온라인)
}

type GetDataRequest struct{}

func NewRoulletUserController(ctx context.Context, cfg config.AppConfig) *UserController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)
	users, err := dynamodb.NewTable[roulette_user.RouletteUser](ctx, db)
	indexes, err := dynamodb.NewTable[index.RouletteIndex](ctx, db)
	products, err := dynamodb.NewTable[product.RouletteProduct](ctx, db)
	if err != nil {
		panic(err)
	}

	logger.Debugf(ctx, "users: %+v", users)
	logger.Debugf(ctx, "indexes: %+v", indexes)
	logger.Debugf(ctx, "products: %+v", products)

	return &UserController{
		db:       db,
		users:    users,
		indexes:  indexes,
		products: products,
		cfg:      cfg,
	}
}

func (r *UserController) OnInit(route base.PathRouter) {
	route.POST("/roulette", r.sendData)
	route.GET("/roulette", r.getData)
	route.POST("/user", r.sendRouletteData)
	route.GET("/user", r.getRouletteData)
}

func (r *UserController) sendData(c *base.Context, req *PostDataRequest) (*PostDataResponse, *base.Error) {
	if req.Code != r.cfg.Authorization.Code {
		logger.Errorf(c, "invalid code")
		return nil, errors.ErrInvalidCode
	}

	ind := 1
	d, e := r.indexes.FindOne(c, bson.D{{Key: "key", Value: "roulette"}})
	if e == nil || d != nil {
		ind = d.MaxIndex + 1
		r.indexes.UpdateOne(c, bson.D{{"key", "roulette"}}, bson.D{{"$set", bson.D{{"maxIndex", ind}}}})
	} else {
		r.indexes.Insert(c, &index.RouletteIndex{
			Key:      "roulette",
			MaxIndex: ind,
			Type:     "roulette",
		})
	}

	err := r.products.Insert(c, &product.RouletteProduct{
		Key:                           fmt.Sprintf("%s-%s", strconv.Itoa(ind), "roulette"),
		Index:                         ind,
		ProductList:                   req.ProductList,
		ProductCount:                  req.ProductCount,
		ProductRouletteOnsiteSequence: req.ProductRouletteOnsiteSequence,
		ProductRouletteOnlineSequence: req.ProductRouletteOnlineSequence,
	})
	if err != nil {
		logger.Errorf(c, "product.InsertOne error: %v", err)
		return nil, errors.ErrInsertFailed
	}

	return &PostDataResponse{
		Index:                         ind,
		ProductList:                   req.ProductList,
		ProductCount:                  req.ProductCount,
		ProductRouletteOnsiteSequence: req.ProductRouletteOnsiteSequence,
		ProductRouletteOnlineSequence: req.ProductRouletteOnlineSequence,
	}, nil
}

func (r *UserController) getData(c *base.Context, req *GetDataRequest) (*PostDataResponse, *base.Error) {
	ind := 0
	d, err := r.indexes.FindOne(c, bson.D{{Key: "key", Value: "roulette"}})
	if err == nil {
		ind = d.MaxIndex
	}

	dt, err := r.products.FindOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", strconv.Itoa(ind), "roulette")}})
	if err != nil {
		return &PostDataResponse{
			Index:                         0,
			ProductList:                   []string{},
			ProductCount:                  []int{},
			ProductRouletteOnsiteSequence: []string{},
			ProductRouletteOnlineSequence: []string{},
		}, nil
	}

	return &PostDataResponse{
		Index:                         dt.Index,
		ProductList:                   dt.ProductList,
		ProductCount:                  dt.ProductCount,
		ProductRouletteOnsiteSequence: dt.ProductRouletteOnsiteSequence,
		ProductRouletteOnlineSequence: dt.ProductRouletteOnlineSequence,
	}, nil
}

func (r *UserController) sendRouletteData(c *base.Context, req *GetProductRequest) (*GetProductResponse, *base.Error) {
	d, e2 := r.indexes.FindOne(c, bson.D{{Key: "key", Value: "roulette"}})

	if e2 != nil {
		logger.Debugf(c, "roulette data is not exists")
		return nil, errors.ErrProductEmpty
	}

	dt, e3 := r.products.FindOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", strconv.Itoa(d.MaxIndex), "roulette")}})

	if e3 != nil {
		logger.Debugf(c, "roulette data is not exists")
		return nil, errors.ErrProductEmpty
	}

	productsList := dt.ProductList
	productsCount := dt.ProductCount

	_, e := r.users.FindOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", req.Address, req.ParticipantType)}})

	if e == nil {
		logger.Debugf(c, "user is already play")
		return nil, errors.ErrAlreadyPlay
	}

	if req.ParticipantType == "onsite" {
		onlineDt, e := r.users.FindOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", req.Address, "online")}})

		if e == nil {
			onlineInd := -1

			logger.Debugs(c, productsList)
			logger.Debugs(c, onlineDt.Product)
			onlineInd = indexOf(onlineDt.Product, productsList)
			if onlineInd == -1 && onlineDt.Product != "꽝" {
				logger.Errorf(c, "product is not exists!")
				return nil, errors.ErrProductEmpty
			}

			if onlineInd != -1 {
				productsCount[onlineInd] += 1
			}
		}
	}

	ind := -1

	ind = indexOf(req.ProductName, productsList)
	if (ind == -1 || ((productsCount[ind] - 1) < 0)) && (req.ProductName != "꽝") {
		logger.Errorf(c, "product is not exists!")
		return nil, errors.ErrProductEmpty
	}

	if ind != -1 {
		productsCount[ind] -= 1
	}

	err := r.users.Insert(c, &roulette_user.RouletteUser{
		Key:     fmt.Sprintf("%s-%s", req.Address, req.ParticipantType),
		Account: req.Address,

		Product:         req.ProductName,
		Count:           1,
		ParticipantType: req.ParticipantType,
		IsAction:        2,
	})
	if err != nil {
		logger.Errorf(c, "user.InsertOne error: %v", err)
		return nil, errors.ErrUserAlreadyExists
	}

	r.products.UpdateOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", strconv.Itoa(d.MaxIndex), "roulette")}}, bson.D{{Key: "$set", Value: bson.D{{Key: "productCount", Value: productsCount}}}})

	return &GetProductResponse{
		Address:         req.Address,
		ProductName:     req.ProductName,
		Count:           1,
		ParticipantType: req.ParticipantType,
		IsAction:        false,
	}, nil
}

func (r *UserController) getRouletteData(c *base.Context, req *GetProductDtRequest) (*GetProductResponse, *base.Error) {
	dt, e := r.users.FindOne(c, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", req.Address, req.ParticipantType)}})

	if e != nil {
		logger.Debugf(c, "user is not exists")
		return &GetProductResponse{
			Address:         req.Address,
			ProductName:     "",
			Count:           0,
			ParticipantType: "",
		}, nil
	}

	isAction := false
	if dt.IsAction == 1 {
		isAction = true
	}

	return &GetProductResponse{
		Address:         dt.Account,
		ProductName:     dt.Product,
		Count:           dt.Count,
		ParticipantType: dt.ParticipantType,
		IsAction:        isAction,
	}, nil
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 // not found.
}
