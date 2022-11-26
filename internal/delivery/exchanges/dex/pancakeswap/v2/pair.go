package panckakeswapv2

import (
	"exchange-provider/internal/delivery/exchanges/dex/pancakeswap/v2/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (p *Panckakeswapv2) Pair(bt, qt types.Token) (*types.Pair, error) {
	return &types.Pair{T1: bt, T2: qt}, nil
}

func (p *Panckakeswapv2) PairWithPrice(bt, qt types.Token) (*types.Pair, error) {

	con, err := contracts.NewContract(p.router, p.provider())
	if err != nil {
		return nil, err
	}

	amounts, err := con.GetAmountsOut(nil, big.NewInt(int64(math.Pow10(6))), []common.Address{bt.Address, qt.Address})
	if err != nil {
		return nil, err
	}

	pair := &types.Pair{
		T1: bt,
		T2: qt,
	}

	pair.Price = big.NewInt(0).Div(amounts[1], big.NewInt(int64(math.Pow10(6)))).String()

	return pair, nil
}
