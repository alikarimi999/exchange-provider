package utils

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type amQueue struct {
	mux *sync.Mutex
	ts  []string
}

func newAMQueue() *amQueue {
	return &amQueue{mux: &sync.Mutex{}}
}

func (q *amQueue) uid(t types.Token, owner, spender common.Address, chainId int64) string {
	return fmt.Sprintf("%s-%s-%s-%d", t.Symbol, owner, spender, chainId)
}

func (a *amQueue) exists(t types.Token, owner, spender common.Address, chainId int64) bool {
	a.mux.Lock()
	defer a.mux.Unlock()
	uid := a.uid(t, owner, spender, chainId)
	for _, id := range a.ts {
		if uid == id {
			return true
		}
	}

	return false
}

func (a *amQueue) add(t types.Token, owner, spender common.Address, chainId int64) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.ts = append(a.ts, a.uid(t, owner, spender, chainId))
}

func (a *amQueue) remove(t types.Token, owner, spender common.Address, chainId int64) {
	a.mux.Lock()
	defer a.mux.Unlock()

	uid := a.uid(t, owner, spender, chainId)
	for i, id := range a.ts {
		if uid == id {
			a.ts = append(a.ts[:i], a.ts[i+1:]...)
			break
		}
	}
}
