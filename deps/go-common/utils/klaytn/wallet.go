package klaytn

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/crypto"
	"github.com/klaytn/klaytn/rlp"
)

type FilterFunc func(*types.Transaction) bool

type Wallet struct {
	Address  common.Address
	privKey  *ecdsa.PrivateKey
	abi      *abi.ABI
	gasPrice *big.Int
	chainID  int64
	endpoint string
	cli      *RPC

	feepayer *Wallet
}

func NewKlaytnWallet(address, privKey string) (*Wallet, error) {
	privKey = strings.Replace(privKey, "0x", "", 1)
	key, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	gasPrice, _ := hexutil.DecodeBig("0xae9f7bcc00")

	return &Wallet{
		Address:  common.HexToAddress(address),
		privKey:  key,
		gasPrice: gasPrice,
	}, nil
}

func (r *Wallet) SetABI(obj string) *Wallet {
	parsed, _ := abi.JSON(strings.NewReader(obj))
	r.abi = &parsed

	return r
}

func (r *Wallet) SetChainID(chainID int64) *Wallet {
	r.chainID = chainID

	return r
}

func (r *Wallet) SetEndpoint(endpoint string) *Wallet {
	r.endpoint = endpoint
	r.cli = NewRPC(context.Background(), endpoint)

	return r
}

func (r *Wallet) SetFeepayer(address, privateKey string) *Wallet {
	r.feepayer, _ = NewKlaytnWallet(address, privateKey)
	return r
}

func (r *Wallet) SignRawTransactionAsFeepayer(ctx *base.Context, hex string, filter ...FilterFunc) (string, *base.Error) {
	tx, err := DecodeRLP(hex)
	if err != nil {
		logger.Errorf(ctx, "Failed to decode transaction: %v", err)
		return "", base.ErrKlaytnFailedToDecodeTx
	}

	if len(filter) > 0 && !filter[0](tx) {
		logger.Errorf(ctx, "Transaction is filtered out: %v", tx)
		return "", base.ErrKlaytnFilteredTransaction
	}

	// It is based on klaytn/blockchain/types/transaction_signing_test.go:481
	signer := types.LatestSignerForChainID(tx.ChainId())
	tx.SignFeePayer(signer, r.privKey)
	// hash, err := signer.HashFeePayer(tx)
	// feePayerSig, err := types.NewTxSignaturesWithValues(signer, tx, hash, []*ecdsa.PrivateKey{r.privKey})
	// if err != nil {
	// 	logger.Errorf(ctx, "Failed to sign fee payer signature: %v", err)
	// 	return "", base.ErrKlaytnFailedToSignTx
	// }

	// tx.SetFeePayerSignatures(feePayerSig)
	txRLP, err := rlp.EncodeToBytes(tx)
	if err != nil {
		logger.Errorf(ctx, "Failed to encode transaction: %v", err)
		return "", base.ErrKlaytnFailedToEncodeTx
	}

	return hexutil.Encode(txRLP), nil
}
func (r *Wallet) Call(ctx *base.Context, contractAddress string, out any, method string, args ...any) *base.Error {
	d, err := r.abi.Pack(method, args...)
	if err != nil {
		logger.Errorf(ctx, "Failed to pack abi: %v", err)
		return base.ErrKlaytnFailedToEncodeTx
	}

	str, err := r.cli.Call(ctx, r.Address.String(), contractAddress, hexutil.Encode(d))
	if err != nil {
		logger.Errorf(ctx, "Failed to call: %v", err)
		return base.ErrKlaytnFailedToCall
	}

	data := hexutil.MustDecode(str)

	r.abi.Unpack(out, method, data)

	return nil
}

func (r *Wallet) SendRawTransactionWithFeepayer(ctx *base.Context, feepayer *Wallet, contractAddress string, method string, args ...any) (*types.Receipt, *base.Error) {
	tx, err := r.SignTransaction(ctx, contractAddress, method, args...)
	if err != nil {
		return nil, err
	}

	tx, _ = feepayer.SignRawTransactionAsFeepayer(ctx, tx)
	logger.Debugf(ctx, "rlp: %s", tx)
	receipt, e := r.SendRawTransaction(ctx, tx)
	if e != nil {
		logger.Errorf(ctx, "Failed to send transaction: %v", e)
		return nil, base.ErrKlaytnFailedToSendTx
	}

	return receipt, nil
}

func (r *Wallet) SignTransaction(ctx *base.Context, contractAddress string, method string, args ...any) (string, *base.Error) {
	to := common.HexToAddress(contractAddress)
	var data []byte

	// logger.Errorf(ctx, "abi data: %v", r.abi)
	logger.Debugf(ctx, "%s(%v)", method, args)
	d, err := r.abi.Pack(method, args...)
	if err != nil {
		logger.Errorf(ctx, "Failed to pack abi: %v", err)
		return "", base.ErrKlaytnFailedToEncodeTx.WithMessage(err.Error())
	}
	data = d

	// FIXME: support non-delegated transaction
	fp := r.feepayer.Address
	nonce, err := r.cli.GetNonce(ctx, r.Address.Hex())
	if err != nil {
		logger.Errorf(ctx, "Failed to get nonce: %v", err)
		return "", base.ErrKlaytnFailedToGetNonce.WithMessage(err.Error())
	}

	tx, err := types.NewTransactionWithMap(types.TxTypeFeeDelegatedSmartContractExecution, map[types.TxValueKeyType]interface{}{
		types.TxValueKeyNonce:    nonce,
		types.TxValueKeyTo:       to,
		types.TxValueKeyFrom:     r.Address,
		types.TxValueKeyAmount:   common.Big0,
		types.TxValueKeyData:     data,
		types.TxValueKeyGasLimit: uint64(700000),
		types.TxValueKeyGasPrice: r.gasPrice,
		types.TxValueKeyFeePayer: fp,
	})
	logger.Debugf(ctx, "transaction map: %+v", map[types.TxValueKeyType]interface{}{
		types.TxValueKeyNonce:    nonce,
		types.TxValueKeyTo:       to,
		types.TxValueKeyFrom:     r.Address,
		types.TxValueKeyAmount:   common.Big0,
		types.TxValueKeyData:     data,
		types.TxValueKeyGasLimit: uint64(700000),
		types.TxValueKeyGasPrice: r.gasPrice,
		types.TxValueKeyFeePayer: fp,
	})
	if err != nil {
		logger.Errorf(ctx, "Failed to create transaction: %v", err)
		return "", base.ErrKlaytnFailedToEncodeTx.WithMessage(err.Error())
	}

	e := tx.Sign(types.LatestSignerForChainID(big.NewInt(r.chainID)), r.privKey)
	if e != nil {
		logger.Errorf(ctx, "Failed to sign: %v", e)
		return "", base.ErrKlaytnFailedToEncodeTx.WithMessage(e.Error())
	}

	logger.Debugf(ctx, "transaction: %+v", tx)
	rawTransaction, _ := rlp.EncodeToBytes(tx)
	logger.Debugf(ctx, "rawTransaction: %+v", rawTransaction)
	encodedTx := hexutil.Encode(rawTransaction)
	logger.Debugf(ctx, "encodedTx: %+v", encodedTx)

	return encodedTx, nil
}

func (r *Wallet) SendRawTransaction(ctx *base.Context, tx string) (*types.Receipt, error) {
	return r.cli.SendRawTransaction(ctx, tx)
}
