package panckakeswapv2

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"math/rand"
	"time"
)

func (ex *Panckakeswapv2) provider() *ts.EthProvider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(ex.ps) > 0 {
		p := ex.ps[r.Intn(len(ex.ps))]
		return p
	}
	return &ts.EthProvider{}
}
