package v1

import (
	"context"
	"strings"

	"biyard.co/common/base"
	"biyard.co/common/dto"
	"biyard.co/common/utils/klaytn"
	"biyard.co/iuniverse-api/config"
	"github.com/klaytn/klaytn/blockchain/types"
)

type FeeController struct {
	w                *klaytn.Wallet
	allowedContracts map[string]bool
	ctx              context.Context
}

func NewFeeController(ctx context.Context, cfg config.AppConfig) *FeeController {
	w, err := klaytn.NewKlaytnWallet(cfg.Feepayer.Address, cfg.Feepayer.PrivateKey)
	if err != nil {
		panic("invalid private key")
	}

	w.SetABI(cfg.ABI).SetEndpoint(cfg.Klaytn.Endpoint).SetChainID(int64(cfg.Klaytn.ChainID)).SetFeepayer(cfg.Feepayer.Address, cfg.Feepayer.PrivateKey)

	c := map[string]bool{}
	for _, contract := range cfg.Contracts {
		c[strings.ToLower(contract)] = true
	}

	return &FeeController{
		w:                w,
		allowedContracts: c,
		ctx:              ctx,
	}
}

func (r *FeeController) OnInit(route base.PathRouter) {
	route.POST("/tx", r.SignTransactionAsFeepayer)
}

func (r *FeeController) SignTransactionAsFeepayer(c *base.Context, req *dto.SignTransactionAsFeePayerRequest) (*dto.SignTransactionAsFeePayerResponse, *base.Error) {
	rlp, err := r.w.SignRawTransactionAsFeepayer(c, req.RLP, func(tx *types.Transaction) bool {
		return r.allowedContracts[strings.ToLower(tx.To().String())]
	})

	if err != nil {
		return nil, err
	}

	return &dto.SignTransactionAsFeePayerResponse{
		RawTransaction: rlp,
	}, nil
}
