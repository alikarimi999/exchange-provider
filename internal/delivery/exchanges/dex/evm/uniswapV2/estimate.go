package uniswapV2

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"math"
	"math/big"
)

func (d *dex) EstimateAmountOut(in, out *types.Token, amount float64) (float64, uint64, error) {
	con, _ := contracts.NewContracts(d.contract, d.provider())
	amountIn, _ := big.NewFloat(0).Mul(big.NewFloat(amount), big.NewFloat(math.Pow10(in.Decimals))).Int(nil)
	res, err := con.EstimateAmountOut(nil, d.router, in.Address, out.Address, amountIn, 2)
	if err != nil {
		return 0, 0, err
	}
	amountOut, _ := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.AmountOut), big.NewFloat(math.Pow10(out.Decimals))).Float64()
	return amountOut, res.Fee.Uint64(), nil
}
