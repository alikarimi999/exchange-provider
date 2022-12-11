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

func (p *Panckakeswapv2) PairWithPrice(in, out types.Token) (*types.Pair, error) {

	con, err := contracts.NewContract(p.router, p.provider())
	if err != nil {
		return nil, err
	}

	amountIn := big.NewInt(int64(math.Pow10(6)))
	amounts, err := con.GetAmountsOut(nil, amountIn, []common.Address{in.Address, out.Address})
	if err != nil {
		return nil, err
	}

	pair := &types.Pair{
		T1: in,
		T2: out,
	}

	inf := big.NewFloat(0).SetInt(amountIn)
	outf := big.NewFloat(0).SetInt(amounts[1])
	pair.Price1 = big.NewFloat(0).Quo(outf, inf).String()
	pair.Price2 = big.NewFloat(0).Quo(inf, outf).String()
	return pair, nil
}
