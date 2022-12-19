package utils

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"math/rand"
	"time"
)

func (am *ApproveManager) provider() *ts.EthProvider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(am.ps) > 0 {
		p := am.ps[r.Intn(len(am.ps))]
		return p
	}
	return &ts.EthProvider{}
}
