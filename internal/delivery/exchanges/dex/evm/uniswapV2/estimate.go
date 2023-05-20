package uniswapV2

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/pkg/errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (d *dex) EstimateAmountOut(in, out *types.Token, amount float64) (float64, uint64, error) {
	con, _ := contracts.NewContracts(d.contract, d.provider())
	amountIn, _ := big.NewFloat(0).Mul(big.NewFloat(amount),
		big.NewFloat(math.Pow10(int(in.Decimals)))).Int(nil)
	res, err := con.EstimateAmountOut(nil, d.priceProvider, d.router, common.HexToAddress(in.ContractAddress),
		common.HexToAddress(out.ContractAddress), amountIn, 2)
	if err != nil {
		return 0, 0, errors.Wrap(errors.ErrInternal, err)
	}
	amountOut, _ := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(res.AmountOut),
		big.NewFloat(math.Pow10(int(out.Decimals)))).Float64()

	return amountOut, res.Fee.Uint64(), nil
}
