package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"math/rand"
	"time"
)

func (ex *Chain) provider() *types.Provider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(ex.Providers) > 0 {
		p := ex.Providers[r.Intn(len(ex.Providers))]
		return p
	}
	return &types.Provider{}
}
