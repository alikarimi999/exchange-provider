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

	Price1    string
	Price2    string
	Liquidity *big.Int
	FeeTier   *big.Int
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s%s%s", p.T1.String(), Delimiter, p.T2.String())
}

func (p *Pair) ToEntity(exchange, native, chainId string) *entity.Pair {

	pair := &entity.Pair{
		T1:              p.T1.ToEntity(chainId),
		T2:              p.T2.ToEntity(chainId),
		Exchange:        exchange,
		ContractAddress: p.Address.String(),

		Liquidity:   p.Liquidity,
		Price1:      p.Price1,
		Price2:      p.Price2,
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
