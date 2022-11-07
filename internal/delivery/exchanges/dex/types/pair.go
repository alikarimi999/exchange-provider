package types

import (
	"exchange-provider/internal/entity"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

var Delimiter string = "/"

type Pair struct {
	Address common.Address
	BT      Token `json:"bt"`
	QT      Token `json:"qt"`

	BaseIsZero bool

	Price     string
	Liquidity *big.Int
	FeeTier   *big.Int
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s%s%s", p.BT.String(), Delimiter, p.QT.String())
}

func (p *Pair) ToEntity(native, standard string, blockTime time.Duration) *entity.Pair {

	pair := &entity.Pair{
		BC: p.BT.ToEntity(standard, blockTime),
		QC: p.QT.ToEntity(standard, blockTime),

		ContractAddress: p.Address.String(),

		Liquidity:   p.Liquidity,
		BestAsk:     p.Price,
		BestBid:     p.Price,
		FeeCurrency: native,
	}

	if p.FeeTier != nil {
		pair.FeeTier = p.FeeTier.Int64()
	}
	return pair
}
