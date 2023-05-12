package uniswapV3

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"math/rand"
	"time"
)

func (d *dex) provider() *types.EthProvider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(d.ps) > 0 {
		p := d.ps[r.Intn(len(d.ps))]
		return p
	}
	return &types.EthProvider{}
}
