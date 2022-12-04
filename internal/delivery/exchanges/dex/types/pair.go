package types

import (
	"exchange-provider/internal/entity"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var Delimiter string = "/"

type Pair struct {
	Address common.Address
	T1      Token `json:"t1"`
	T2      Token `json:"t2"`

	BaseIsZero bool

	Price     string
	Liquidity *big.Int
	FeeTier   *big.Int
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s%s%s", p.T1.String(), Delimiter, p.T2.String())
}

func (p *Pair) ToEntity(native, chainId string) *entity.Pair {

	pair := &entity.Pair{
		C1: p.T1.ToEntity(chainId),
		C2: p.T2.ToEntity(chainId),

		ContractAddress: p.Address.String(),

		Liquidity:   p.Liquidity,
		BestAsk:     p.Price,
		BestBid:     p.Price,
		FeeCurrency: native,
	}

	if p.Address == common.HexToAddress("0x") {
		pair.ContractAddress = ""
	}

	if p.FeeTier != nil {
		pair.FeeTier = p.FeeTier.Int64()
	}
	return pair
}
