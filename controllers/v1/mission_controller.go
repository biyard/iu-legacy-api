package v1

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
	"biyard.co/common/logger"
	"biyard.co/common/utils/klaytn"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/models/mission"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"go.mongodb.org/mongo-driver/bson"

	errors "biyard.co/iuniverse-api/errors"
)

var abi = `
[
	{
		"inputs": [
		  {
			"internalType": "address",
			"name": "contractAddress",
			"type": "address"
		  },
		  {
			"internalType": "uint256",
			"name": "tokenID",
			"type": "uint256"
		  },
		  {
			"internalType": "uint256",
			"name": "amount",
			"type": "uint256"
		  }
		],
		"name": "addNFTEXP",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  },
	  {
		"inputs": [
		  {
			"internalType": "address",
			"name": "userAddress",
			"type": "address"
		  },
		  {
			"components": [
			  {
				"internalType": "string",
				"name": "key",
				"type": "string"
			  },
			  {
				"internalType": "string",
				"name": "missionInfo",
				"type": "string"
			  },
			  {
				"internalType": "uint256",
				"name": "exp",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "progressDate",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "startDate",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "endDate",
				"type": "uint256"
			  }
			],
			"internalType": "struct ExperienceActivityData",
			"name": "data",
			"type": "tuple"
		  }
		],
		"name": "addAccountActivities",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  }
]
`

type MissionController struct {
	db                 *dynamodb.DynamoDB
	cfg                config.AppConfig
	quizs              *dynamodb.Table[mission.Quiz]
	indexQuizes        *dynamodb.Index[mission.Quiz]
	indexAddressQuizes *dynamodb.Index[mission.Quiz]

	feepayer *klaytn.Wallet
	owner    *klaytn.Wallet
}

func NewMissionController(ctx context.Context, cfg config.AppConfig) *MissionController {
	db := dynamodb.New(ctx, cfg.Database, cfg.Config.AWS)

	quizs, err := dynamodb.NewTable[mission.Quiz](ctx, db)
	indexQuizes, _ := dynamodb.NewIndex[mission.Quiz](ctx, db, "indexStartDate")
	indexAddressQuizes, _ := dynamodb.NewIndex[mission.Quiz](ctx, db, "address")

	if err != nil {
		logger.Critical(ctx, err)
	}

	f, err := klaytn.NewKlaytnWallet(cfg.Feepayer.Address, cfg.Feepayer.PrivateKey)
	if err != nil {
		logger.Critical(ctx, "invalid private key")
	}

	o, err := klaytn.NewKlaytnWallet(cfg.Owner.Address, cfg.Owner.PrivateKey)
	if err != nil {
		logger.Critical(ctx, "invalid private key")
	}

	o.SetABI(abi).SetEndpoint(cfg.Klaytn.Endpoint).SetChainID(int64(cfg.Klaytn.ChainID)).SetFeepayer(cfg.Feepayer.Address, cfg.Feepayer.PrivateKey)

	return &MissionController{
		db:                 db,
		quizs:              quizs,
		cfg:                cfg,
		indexQuizes:        indexQuizes,
		indexAddressQuizes: indexAddressQuizes,

		feepayer: f,
		owner:    o,
	}
}

type ListQuizMissionRequest struct {
	Address string `form:"address"`
}

type ListQuizMissionResponse struct {
	Status string     `json:"status"`
	Quizes []QuizData `json:"quizes"`
}

type ScoringQuizMissionRequest struct {
	Key     string `json:"key"`
	Address string `json:"address"`
	Answers []int  `json:"answers"`
}

type ScoringQuizMissionResponse struct {
	Status         string `json:"status"`
	CorrectAnswers int    `json:"correctanswers"`
}

type QuizData struct {
	Key                  string     `json:"key"`
	Address              string     `json:"address"`
	Completed            bool       `json:"completed"`
	MissionNameKo        string     `json:"missionNameKo"`
	MissionNameEn        string     `json:"missionNameEn"`
	MissionTypeKo        string     `json:"missionTypeKo"`
	MissionTypeEn        string     `json:"missionTypeEn"`
	Exp                  int        `json:"exp"`
	StartDate            int        `json:"startDate"`
	QuestionsKo          []string   `json:"questionsKo"`
	QuestionsEn          []string   `json:"questionsEn"`
	MultipleChoiceViewKo [][]string `json:"multipleChoiceViewKo"`
	MultipleChoiceViewEn [][]string `json:"multipleChoiceViewEn"`
	Cutline              int        `json:"cutline"`
}

type ExperienceActivityData struct {
	Key          string   `abi:"key"`
	MissionInfo  string   `abi:"missionInfo"`
	Exp          *big.Int `abi:"exp"`
	ProgressDate *big.Int `abi:"progressDate"`
	StartDate    *big.Int `abi:"startDate"`
	EndDate      *big.Int `abi:"endDate"`
}

func (r *MissionController) OnInit(route base.PathRouter) {
	route.GET("/quizes", r.listQuizMissions)
	route.POST("/quiz/score", r.scoringQuizMission)
}

func (r *MissionController) scoringQuizMission(ctx *base.Context, req *ScoringQuizMissionRequest) (*ScoringQuizMissionResponse, *base.Error) {
	quizInfo, err := r.quizs.FindOne(ctx, bson.D{{Key: "key", Value: req.Key}})

	if err != nil {
		logger.Errorf(ctx, "failed to query quiz data "+err.Error())
		return nil, base.ErrUnknown.WithMessage("failed to query quiz data")
	}

	userQuizInfo, _ := r.quizs.FindOne(ctx, bson.D{{Key: "key", Value: fmt.Sprintf("%s-%s", req.Address, quizInfo.MissionNameKo)}})
	if userQuizInfo != nil {
		return nil, base.ErrUnknown.WithMessage("already participate in quiz mission")
	}

	logger.Debugf(ctx, "quizInfo: %+v", quizInfo)

	correctAnswer := quizInfo.Answers

	successCount := 0

	for i := 0; i < len(correctAnswer); i++ {
		if correctAnswer[i] == req.Answers[i] {
			successCount += 1
		}
	}

	now := time.Now()
	sec := now.Unix()

	if quizInfo.Cutline > successCount {
		return &ScoringQuizMissionResponse{
			Status:         "failed",
			CorrectAnswers: successCount,
		}, nil
	}

	_, e1 := r.RewardXP(ctx, r.cfg, req.Address, quizInfo.MissionNameKo, quizInfo.Exp, int(sec))
	if e1 != nil {
		logger.Errorf(ctx, "failed to reward xp: %+v, error: %+v", quizInfo.Exp, e1)

		return nil, base.ErrUnknown.WithMessage("failed to reward xp")
	}

	e := r.quizs.Insert(ctx, &mission.Quiz{
		Key:           fmt.Sprintf("%s-%s", req.Address, quizInfo.MissionNameKo),
		Address:       req.Address,
		MissionNameKo: quizInfo.MissionNameKo,
		CreatedAt:     sec,
	})

	if e != nil {
		logger.Errorf(ctx, "failed to insert address mission data "+err.Error())
		return nil, base.ErrUnknown.WithMessage("failed to insert address mission data")
	}

	return &ScoringQuizMissionResponse{
		Status:         "success",
		CorrectAnswers: successCount,
	}, nil
}

func (r *MissionController) listQuizMissions(ctx *base.Context, req *ListQuizMissionRequest) (*ListQuizMissionResponse, *base.Error) {
	nowTime := time.Now().Add(time.Hour * time.Duration(9))

	prevDate := nowTime.AddDate(0, 0, -1)

	location, _ := time.LoadLocation("Asia/Seoul")

	date := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 10, 0, 0, 0, location)

	if nowTime.Hour() < 10 {
		date = time.Date(prevDate.Year(), prevDate.Month(), prevDate.Day(), 10, 0, 0, 0, location)
	}

	logger.Debugf(ctx, "quiz mission Time: %+v %+v %+v %+v", date.Hour(), date.Minute(), date.Second(), date.Unix())

	quizInfos, err := r.indexQuizes.Find(ctx, dynamodb.QueryRequest{
		Key: fmt.Sprintf("%s-%s", "ko", strconv.Itoa(int(date.Unix()))),
	})

	if err != nil {
		logger.Errorf(ctx, "failed to query quiz data"+err.Error())
		return nil, base.ErrUnknown.WithMessage("failed to query quiz data")
	}

	quizInfosEn, err := r.indexQuizes.Find(ctx, dynamodb.QueryRequest{
		Key: fmt.Sprintf("%s-%s", "en", strconv.Itoa(int(date.Unix()))),
	})

	if err != nil {
		logger.Errorf(ctx, "failed to query quiz data"+err.Error())
		return nil, base.ErrUnknown.WithMessage("failed to query quiz data")
	}

	completeMissions, err := r.indexAddressQuizes.Find(ctx, dynamodb.QueryRequest{
		Key: req.Address,
	})

	var addressMissions []string = []string{}

	if err == nil {
		for i := 0; i < len(completeMissions.Items); i++ {
			addressMissions = append(addressMissions, completeMissions.Items[i].MissionNameKo)
		}

		logger.Debugf(ctx, "addressMissionsData: %+v", addressMissions)
	} else {
		logger.Errorf(ctx, "failed to query complete mission data "+err.Error())
	}

	logger.Debugf(ctx, "QuizData: %+v", quizInfos.Items)

	var quizDatas []QuizData = []QuizData{}

	for i := 0; i < len(quizInfos.Items); i++ {
		completed := false
		items := quizInfos.Items[i]
		if contains(addressMissions, items.MissionNameKo) {
			completed = true
		}

		var quizDataEn mission.Quiz

		for j := 0; j < len(quizInfosEn.Items); j++ {
			if quizInfosEn.Items[j].MissionNameKo == quizInfos.Items[i].MissionNameKo {
				quizDataEn = quizInfosEn.Items[j]
				break
			}
		}

		quizDatas = append(quizDatas, QuizData{
			Key:                  items.Key,
			Address:              req.Address,
			Completed:            completed,
			MissionNameKo:        items.MissionNameKo,
			MissionNameEn:        quizDataEn.MissionNameEn,
			MissionTypeKo:        items.MissionType,
			MissionTypeEn:        quizDataEn.MissionType,
			Exp:                  items.Exp,
			StartDate:            items.StartDate,
			QuestionsKo:          items.Questions,
			QuestionsEn:          quizDataEn.Questions,
			MultipleChoiceViewKo: items.MultipleChoiceView,
			MultipleChoiceViewEn: quizDataEn.MultipleChoiceView,
			Cutline:              items.Cutline,
		})
	}

	return &ListQuizMissionResponse{
		Status: "ok",
		Quizes: quizDatas,
	}, nil
}

type SignTransactionResponse struct {
	Address   string         `json:"address"`
	TxReceipt *types.Receipt `json:"txReceipt"`
}

func (r *MissionController) RewardXP(ctx *base.Context, cfg config.AppConfig, account string, msg string, rewardXP int, finishedAt int) (*SignTransactionResponse, *base.Error) {
	receipt, err := r.owner.SendRawTransactionWithFeepayer(ctx, r.feepayer, r.cfg.ContractAddresses.AccountProfile, "addAccountActivities",
		common.HexToAddress(account),
		ExperienceActivityData{
			Key:          fmt.Sprintf("%v-%v", msg, time.Now().Unix()),
			MissionInfo:  msg,
			Exp:          big.NewInt(int64(rewardXP)),
			ProgressDate: big.NewInt(int64(finishedAt)),
			StartDate:    big.NewInt(0),
			EndDate:      big.NewInt(10000000000),
		},
	)

	if err != nil {
		return nil, errors.ErrTransactionSend
	}

	return &SignTransactionResponse{
		Address:   r.owner.Address.String(),
		TxReceipt: receipt,
	}, nil
}
