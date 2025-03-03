package klaytn

import (
	"context"
	"errors"
	"strings"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/common/utils/httpclient"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common/hexutil"
)

type JSONRPCRequest struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type JSONRPCResponse[T any] struct {
	Version string `json:"jsonrpc"`
	Result  T      `json:"result"`
	ID      int    `json:"id"`
}

type RPC struct {
	ctx      context.Context
	cli      []*httpclient.HTTPClient
	selected int
	endpoint string
}

type JSONCallParams struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Input string `json:"input"`
}

func NewRPC(ctx context.Context, endpoint string, opts ...httpclient.ClientOption) *RPC {
	cli := make([]*httpclient.HTTPClient, 0)
	endpoints := strings.Split(endpoint, ",")
	for _, endpoint := range endpoints {
		cli = append(cli, httpclient.DefaultClient(endpoint).WithOpts(opts...))
	}

	return &RPC{
		ctx:      ctx,
		cli:      cli,
		selected: 0,
		endpoint: endpoint,
	}
}

func (r *RPC) Send(ctx *base.Context, req *JSONRPCRequest, res any, retry ...int) (int, error) {
	code, err := r.cli[r.selected].Post(ctx, "", req, res)
	if err != nil {
		logger.Errorf(ctx, "%v(%v) error: %v", r.cli[r.selected].Endpoint(), code, err)
		r.selected = (r.selected + 1) % len(r.cli)
		if len(retry) > 0 && retry[0] > 0 {
			return r.Send(ctx, req, res, retry[0]-1)
		}

		return code, err
	}

	return code, err
}

func (r *RPC) Call(ctx *base.Context, from string, to string, input string) (string, error) {
	res := &JSONRPCResponse[string]{}
	_, err := r.cli[r.selected].Post(ctx, "", &JSONRPCRequest{
		Version: "2.0",
		Method:  "klay_call",
		Params: []any{JSONCallParams{
			From:  from,
			To:    to,
			Input: input,
		}, "latest"},
		ID: 1,
	}, res)

	if err != nil {
		r.selected = (r.selected + 1) % len(r.cli)
		return "", err
	}

	logger.Debugf(ctx, "JSONRPCResponse: %+v %+v", res, err)

	return res.Result, err
}

func (r *RPC) SendRawTransaction(ctx *base.Context, hex string) (*types.Receipt, error) {
	txHash := ""
	res := &JSONRPCResponse[string]{}
	_, err := r.Send(ctx, &JSONRPCRequest{
		Version: "2.0",
		Method:  "klay_sendRawTransaction",
		Params:  []string{hex},
		ID:      1,
	}, res)
	if err != nil {
		return nil, err
	}
	txHash = res.Result

	logger.Debugf(ctx, "txHash: %s", txHash)

	return r.GetTransactionReceipt(txHash)
}

func (r *RPC) GetTransactionReceipt(txHash string) (*types.Receipt, error) {
	res := &JSONRPCResponse[*types.Receipt]{}
	time.Sleep(1 * time.Second)

	for i := 0; i < 30; i++ {
		r.Send(base.NewBaseContext(), &JSONRPCRequest{
			Version: "2.0",
			Method:  "klay_getTransactionReceipt",
			Params:  []string{txHash},
			ID:      1,
		}, res)
		if res.Result != nil {
			return res.Result, nil
		}
		time.Sleep(500 * time.Millisecond)
	}

	return nil, errors.New("not found")
}

func (r *RPC) GetNonce(ctx *base.Context, address string) (uint64, error) {
	res := &JSONRPCResponse[string]{}

	logger.Debugs(ctx, JSONRPCRequest{
		Version: "2.0",
		Method:  "klay_getTransactionCount",
		Params:  []string{address, "latest"},
		ID:      1,
	})

	_, err := r.Send(ctx, &JSONRPCRequest{
		Version: "2.0",
		Method:  "klay_getTransactionCount",
		Params:  []string{address, "latest"},
		ID:      1,
	}, res)
	if err != nil {
		return 0, err
	}

	nonce, _ := hexutil.DecodeUint64(res.Result)

	return nonce, nil
}

func (r *RPC) BlockNumber(ctx *base.Context) (string, error) {
	res := &JSONRPCResponse[string]{}
	logger.Debugf(ctx, "endpoint: %v", r.cli)

	_, err := r.Send(ctx, &JSONRPCRequest{
		Version: "2.0",
		Method:  "klay_blockNumber",
		Params:  []string{},
		ID:      1,
	}, res)
	if err != nil {
		return "0x0", err
	}

	return res.Result, err
}

type Log struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}

func (r *RPC) GetLogs(ctx *base.Context, fromBlock, toBlock string, addresses []string, topics []string) ([]Log, error) {
	res := &JSONRPCResponse[[]Log]{}

	_, err := r.Send(ctx, &JSONRPCRequest{
		Version: "2.0",
		Method:  "klay_getLogs",
		Params: []map[string]interface{}{
			{
				"fromBlock": fromBlock,
				"toBlock":   toBlock,
				"address":   addresses,
				// "topics":    topics,
			},
		},
		ID: 1,
	}, res)

	return res.Result, err
}
