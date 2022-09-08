package uniswapv3

import (
	"math"
	"math/big"
)

func computeTxFee(gasPrice *big.Int, gasUsed uint64) string {
	price := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), big.NewFloat(math.Pow10(-ethDecimals)))
	fee := new(big.Float).Mul(price, big.NewFloat(float64(gasUsed)))
	return fee.Text('f', ethDecimals)
}
