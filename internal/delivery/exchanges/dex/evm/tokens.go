package evm

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"sync"
)

type tokens struct {
	mux  *sync.RWMutex
	list map[string]*entity.Token
}

func newTokens() *tokens {
	return &tokens{
		mux:  &sync.RWMutex{},
		list: make(map[string]*entity.Token),
	}
}

func (ts *tokens) add(t *entity.Token) {
	ts.mux.Lock()
	defer ts.mux.Unlock()
	ts.list[t.String()] = t.Snapshot()
}

func (ts *tokens) exists(id string) bool {
	ts.mux.RLock()
	defer ts.mux.RUnlock()
	_, ok := ts.list[id]
	return ok
}

func (ts *tokens) get(id string) (*entity.Token, error) {
	ts.mux.RLock()
	defer ts.mux.RUnlock()
	t, ok := ts.list[id]
	if ok {
		return t.Snapshot(), nil
	}
	return nil, errors.Wrap(errors.ErrNotFound)
}
