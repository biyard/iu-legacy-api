package sdk

import (
	"context"

	"biyard.co/common/dto"
	"biyard.co/common/utils/httpclient"
)

type FeeAPI struct {
	ctx context.Context
	cli *httpclient.HTTPClient
}

func NewFeeAPI(ctx context.Context, endpoint string, opts ...httpclient.ClientOption) *FeeAPI {
	return &FeeAPI{
		ctx: ctx,
		cli: httpclient.DefaultClient(endpoint).WithOpts(opts...),
	}
}

func (r *FeeAPI) SignTransactionAsFeepayer(ctx context.Context, req *dto.SignTransactionAsFeePayerRequest) (*dto.SignTransactionAsFeePayerResponse, error) {
	resp := &dto.SignTransactionAsFeePayerResponse{}
	_, err := r.cli.Post(ctx, "/v1/fee/tx", req, resp)

	return resp, err
}
