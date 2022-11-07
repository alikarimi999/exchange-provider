package panckakeswapv2

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"math/rand"
	"time"
)

func (ex *Panckakeswapv2) provider() *types.Provider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(ex.ps) > 0 {
		p := ex.ps[r.Intn(len(ex.ps))]
		return p
	}
	return &types.Provider{}
}
