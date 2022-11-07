package uniswapv3

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	tier0 = 100
	tier1 = 500
	tier2 = 3000
	tier3 = 10000
)

var feeTiers = []*big.Int{big.NewInt(tier0), big.NewInt(tier1), big.NewInt(tier2), big.NewInt(tier3)}
var erc20TransferSignature = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

var Q96 = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)
var Q192 = new(big.Int).Exp(Q96, big.NewInt(2), nil)
