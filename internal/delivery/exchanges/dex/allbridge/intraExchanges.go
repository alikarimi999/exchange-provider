package allbridge

import (
	"exchange-provider/internal/entity"
	"sync"
)

type intraExchanges struct {
	mux  *sync.RWMutex
	list map[string][]entity.Exchange
}

func (i *intraExchanges) add(e entity.Exchange) {
	switch e.Type() {
	case entity.EvmDEX:
		ex := e.(entity.EVMDex)
		i.mux.Lock()
		i.list[ex.Chain()] = append(i.list[ex.Chain()], e)
		i.mux.Unlock()
	}
}
