package shop

import (
	"math/big"

	common "github.com/klaytn/klaytn/common"
)

type ShopItem struct {
	Id              *big.Int       `abi:"id"`
	Price           *big.Int       `abi:"price"`
	TokenId         *big.Int       `abi:"tokenId"`
	Name            string         `abi:"name"`
	Image           string         `abi:"image"`
	Supply          *big.Int       `abi:"supply"`
	Likes           *big.Int       `abi:"likes"`
	Reports         *big.Int       `abi:"reports"`
	ContractAddress common.Address `abi:"contractAddress"`
	Creator         common.Address `abi:"creator"`
	Remaining       *big.Int       `abi:"remaining"`
	Level           uint8          `abi:"level"`
	Metadata        string         `abi:"metadata"`
}
