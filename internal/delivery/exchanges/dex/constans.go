package dex

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var max96 = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 96), common.Big1)
var erc20TransferSignature = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

var ethDecimals = 18
