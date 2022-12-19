package numbers

import (
	"fmt"
	"math"
	"math/big"
)

func BigIntToFloatString(bn *big.Int, decimal int) string {
	bf := new(big.Float).Quo(new(big.Float).SetInt(bn), big.NewFloat(math.Pow10(decimal)))
	return bf.Text('f', decimal)
}

func FloatStringToBigInt(s string, decimals int) (*big.Int, error) {
	bf, err := StringToBigFloat(s)
	if err != nil {
		return nil, err
	}
	bn, _ := new(big.Float).Mul(bf, big.NewFloat(math.Pow10(decimals))).Int(nil)
	return bn, nil

}

func StringToBigFloat(s string) (*big.Float, error) {
	if s == "" {
		return big.NewFloat(0), nil
	}

	bf, ok := new(big.Float).SetString(s)
	if ok {
		return bf, nil
	}
	return nil, fmt.Errorf("numbers: invalid string")

}
