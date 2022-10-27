package uniswapv3

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var factory = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
var routerV2 = common.HexToAddress("0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45")
var erc20TransferSignature = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

const (
	// chainId     = "ERC20"
	ethDecimals = 18
	// ether       = "ETH"
	wrappedETH = "WETH"
)

const (
	tier0 = 100
	tier1 = 500
	tier2 = 3000
	tier3 = 10000
)

var max96 = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 96), common.Big1)
var feeTiers = []*big.Int{big.NewInt(tier0), big.NewInt(tier1), big.NewInt(tier2), big.NewInt(tier3)}
var Q96 = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)
var Q192 = new(big.Int).Exp(Q96, big.NewInt(2), nil)
