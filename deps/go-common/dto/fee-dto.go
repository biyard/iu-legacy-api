package dto

type SignTransactionAsFeePayerRequest struct {
	RLP string `json:"rlp"`
}

type SignTransactionAsFeePayerResponse struct {
	RawTransaction string `json:"rawTransaction"`
}
