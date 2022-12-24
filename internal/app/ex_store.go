package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"
)

type exStore struct {
	mux *sync.Mutex

	repo      ExchangeRepo
	exchanges map[string]entity.Exchange

	l logger.Logger
}

func newExStore(l logger.Logger, exRepo ExchangeRepo) *exStore {
	s := &exStore{
		mux:       &sync.Mutex{},
		repo:      exRepo,
		exchanges: make(map[string]entity.Exchange),
		l:         l,
	}

	exs, err := s.repo.GetAll()
	if err != nil {
		s.l.Error("exStore", err.Error())
		return s
	}

	for _, ex := range exs {
		s.exchanges[ex.Id()] = ex
		ex.Run()
		l.Debug("exStore.add", fmt.Sprintf("exchange '%s' added", ex.Id()))

	}
	return s
}

func (a *exStore) get(nid string) (entity.Exchange, error) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if _, ok := a.exchanges[nid]; ok {
		return a.exchanges[nid], nil
	}

	return nil, errors.Wrap(errors.ErrNotFound)
}

func (a *exStore) add(ex entity.Exchange) error {
	if err := a.repo.Add(ex); err != nil {
		return err
	}
	a.exchanges[ex.Id()] = ex
	go ex.Run()
	a.l.Debug("exStore.add", fmt.Sprintf("exchange '%s' added", ex.Id()))
	return nil
}

func (a *exStore) exists(nid string) bool {

	if _, ok := a.exchanges[nid]; ok {
		return true
	}
	return false
}

func (a *exStore) getAll() []entity.Exchange {
	a.mux.Lock()
	defer a.mux.Unlock()

	var exs []entity.Exchange
	for _, ex := range a.exchanges {
		exs = append(exs, ex)
	}

	return exs
}

func (a *exStore) all(names ...string) []entity.Exchange {
	if len(names) == 0 {
		return a.getAll()
	}

	a.mux.Lock()
	defer a.mux.Unlock()

	exs := []entity.Exchange{}

	for _, ex := range a.exchanges {
		for _, name := range names {
			if ex.Id() == name {
				exs = append(exs, ex)
			}
		}
	}

	return exs
}

func (a *exStore) remove(nid string) error {
	a.mux.Lock()
	defer a.mux.Unlock()

	ex, ok := a.exchanges[nid]
	if ok {
		if err := a.repo.Remove(ex); err != nil {
			return err
		}
		ex.Stop()
		delete(a.exchanges, nid)
		return nil
	}

	return fmt.Errorf("exchange %s not found", nid)
}
