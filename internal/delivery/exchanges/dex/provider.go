package dex

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"math/rand"
	"time"
)

func (ex *dex) provider() *ts.EthProvider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(ex.cfg.Providers) > 0 {
		p := ex.cfg.Providers[r.Intn(len(ex.cfg.Providers))]
		return p
	}
	return &ts.EthProvider{}
}
