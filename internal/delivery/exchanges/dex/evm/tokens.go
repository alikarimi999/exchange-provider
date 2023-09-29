package evm

import (
	"exchange-provider/internal/entity"
	"sync"
)

type tokensList struct {
	mux  *sync.RWMutex
	list map[string]*entity.Token
}

func newTokenList() *tokensList {
	return &tokensList{
		mux:  &sync.RWMutex{},
		list: make(map[string]*entity.Token),
	}
}

func (tl *tokensList) add(t *entity.Token) {
	tl.mux.Lock()
	defer tl.mux.Unlock()
	tl.list[t.String()] = t
}

func (tl *tokensList) get(tId entity.TokenId) *entity.Token {
	tl.mux.RLock()
	defer tl.mux.RUnlock()
	return tl.list[tId.String()]
}
