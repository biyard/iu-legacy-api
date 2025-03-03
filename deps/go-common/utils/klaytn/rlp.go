package klaytn

import (
	"bytes"
	"strings"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/rlp"
)

// DecodeRLP decodes a hex string to a transaction.
// This hex string should be a non-prefixed by 0x
func DecodeRLP(hex string) (*types.Transaction, error) {
	hex = strings.Replace(hex, "0x", "", 1)

	rlpTx := common.Hex2Bytes(hex)
	tx := types.Transaction{}
	err := rlp.Decode(bytes.NewReader(rlpTx), &tx)

	return &tx, err
}
