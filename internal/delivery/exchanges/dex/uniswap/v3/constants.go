package uniswapv3

import "math/big"

const (
	tier0 = 100
	tier1 = 500
	tier2 = 3000
	tier3 = 10000
)

var feeTiers = []*big.Int{big.NewInt(tier0), big.NewInt(tier1), big.NewInt(tier2), big.NewInt(tier3)}

var Q96 = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)
var Q192 = new(big.Int).Exp(Q96, big.NewInt(2), nil)
