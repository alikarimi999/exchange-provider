package uniswapV2

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV2/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (d *dex) Pair(in, out types.Token) (*types.Pair, error) {
	con, err := contracts.NewContract(d.router, d.provider())
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
