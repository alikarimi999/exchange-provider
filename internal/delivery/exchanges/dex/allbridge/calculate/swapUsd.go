package calculate

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"math/big"
	"strconv"
)

func swapToVUsd(amount *big.Float, ti *types.TokenInfo, pi *types.PoolInfo) *big.Float {
	fs, _ := strconv.ParseFloat(ti.FeeShare, 64)
	av, _ := strconv.ParseFloat(pi.AValue, 64)
	dv, _ := strconv.ParseFloat(pi.DValue, 64)
	vb, _ := strconv.ParseFloat(pi.VUsdBalance, 64)
	tb, _ := strconv.ParseFloat(pi.TokenBalance, 64)

	fee := new(big.Float).Mul(amount, big.NewFloat(fs))
	amountWithoutFee := new(big.Float).Sub(amount, fee)
	inSystemPrecision := toSystemPrecision(amountWithoutFee, ti.Decimals)
	tokenBalance := new(big.Float).Add(big.NewFloat(tb), inSystemPrecision)
	vUsdNewAmount := getY(tokenBalance, big.NewFloat(av), big.NewFloat(dv))
	return new(big.Float).Sub(big.NewFloat(vb), vUsdNewAmount)
}

func swapFromVUsd(amount *big.Float, ti *types.TokenInfo, pi *types.PoolInfo) *big.Float {
	av, _ := strconv.ParseFloat(pi.AValue, 64)
	dv, _ := strconv.ParseFloat(pi.DValue, 64)
	vb, _ := strconv.ParseFloat(pi.VUsdBalance, 64)
	tb, _ := strconv.ParseFloat(pi.TokenBalance, 64)
	feeShare, _ := strconv.ParseFloat(ti.FeeShare, 64)

	vUsdBalance := new(big.Float).Add(amount, big.NewFloat(vb))
	newAmount := getY(vUsdBalance, big.NewFloat(av), big.NewFloat(dv))
	result := fromSystemPrecision(new(big.Float).Sub(big.NewFloat(tb), newAmount), ti.Decimals)
	fee := new(big.Float).Mul(result, big.NewFloat(feeShare))
	return new(big.Float).Sub(result, fee)
}
