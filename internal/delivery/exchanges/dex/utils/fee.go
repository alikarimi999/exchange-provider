package utils

import (
	"math"
	"math/big"
)

func TxFee(gasPrice *big.Int, gasUsed uint64) string {
	price := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), big.NewFloat(math.Pow10(-18)))
	fee := new(big.Float).Mul(price, big.NewFloat(float64(gasUsed)))
	return fee.Text('f', 18)
}
