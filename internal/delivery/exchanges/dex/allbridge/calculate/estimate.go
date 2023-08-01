package calculate

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"math"
	"math/big"
)

func Estimate(amount float64, srcToken, dstToken *types.TokenInfo,
	srcPool, dstPool *types.PoolInfo) float64 {
	amountToSend := new(big.Float).Mul(big.NewFloat(amount),
		big.NewFloat(math.Pow10(srcToken.Decimals)))

	vUsd := swapToVUsd(amountToSend, srcToken, srcPool)
	result := swapFromVUsd(vUsd, dstToken, dstPool)
	amountToRecievd, _ := new(big.Float).Quo(result, big.NewFloat(math.Pow10(dstToken.Decimals))).Float64()

	return amountToRecievd
}
