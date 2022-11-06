package dex

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func hashToAddress(h common.Hash) common.Address {
	return common.BytesToAddress(h[:])
}

func txFee(gasPrice *big.Int, gasUsed uint64) string {
	price := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), big.NewFloat(math.Pow10(-18)))
	fee := new(big.Float).Mul(price, big.NewFloat(float64(gasUsed)))
	return fee.Text('f', 18)
}
