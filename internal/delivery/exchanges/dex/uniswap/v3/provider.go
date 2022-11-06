package uniswapv3

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"math/rand"
	"time"
)

func (ex *UniswapV3) provider() *types.Provider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(ex.Ps) > 0 {
		p := ex.Ps[r.Intn(len(ex.Ps))]
		return p
	}
	return &types.Provider{}
}
