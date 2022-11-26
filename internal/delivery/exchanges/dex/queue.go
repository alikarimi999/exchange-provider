package dex

// import (
// 	"exchange-provider/internal/delivery/exchanges/dex/types"
// 	"fmt"
// 	"sync"

// 	"github.com/ethereum/go-ethereum/common"
// )

// type amQueue struct {
// 	mux *sync.Mutex
// 	ts  []string
// }

// func newAMQueue() *amQueue {
// 	return &amQueue{mux: &sync.Mutex{}}
// }

// func (q *amQueue) uid(t types.Token, owner, spender common.Address) string {
// 	return fmt.Sprintf("%s-%s-%s", t.Symbol, owner, spender)
// }

// func (a *amQueue) exists(t types.Token, owner, spender common.Address) bool {
// 	a.mux.Lock()
// 	defer a.mux.Unlock()
// 	uid := a.uid(t, owner, spender)
// 	for _, id := range a.ts {
// 		if uid == id {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (a *amQueue) add(t types.Token, owner, spender common.Address) {
// 	a.mux.Lock()
// 	defer a.mux.Unlock()
// 	a.ts = append(a.ts, a.uid(t, owner, spender))
// }

// func (a *amQueue) remove(t types.Token, owner, spender common.Address) {
// 	a.mux.Lock()
// 	defer a.mux.Unlock()

// 	uid := a.uid(t, owner, spender)
// 	for i, id := range a.ts {
// 		if uid == id {
// 			a.ts = append(a.ts[:i], a.ts[i+1:]...)
// 			break
// 		}
// 	}
// }
