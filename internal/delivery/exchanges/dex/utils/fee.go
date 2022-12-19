package utils

import (
	"math"
	"math/big"
)

var EthDecimals = 18

func TxFee(gasPrice *big.Int, gasUsed uint64) string {
	price := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), big.NewFloat(math.Pow10(-EthDecimals)))
	fee := new(big.Float).Mul(price, big.NewFloat(float64(gasUsed)))
	return fee.Text('f', EthDecimals)
}
