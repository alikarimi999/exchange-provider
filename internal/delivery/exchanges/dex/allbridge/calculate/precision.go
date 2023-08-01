package calculate

import (
	"math"
	"math/big"
)

var Allbridge_Precision int

func toSystemPrecision(amount *big.Float, decimals int) *big.Float {
	return convertAmountPrecision(amount, decimals, Allbridge_Precision)
}

func fromSystemPrecision(amount *big.Float, decimals int) *big.Float {
	return convertAmountPrecision(amount, Allbridge_Precision, decimals)
}

func convertAmountPrecision(amount *big.Float, decimalsFrom int, decimalsTo int) *big.Float {
	dif := decimalsTo - decimalsFrom
	powBase10 := new(big.Float).SetFloat64(math.Pow10(dif))
	return amount.Mul(amount, powBase10)
}
