package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"
)

type exStore struct {
	repo      ExchangeRepo
	mux       *sync.RWMutex
	exchanges map[string]entity.Exchange
	l         logger.Logger
}

func newExStore(l logger.Logger, exRepo ExchangeRepo) *exStore {
	s := &exStore{
		repo:      exRepo,
		mux:       &sync.RWMutex{},
		exchanges: make(map[string]entity.Exchange),
		l:         l,
	}

	exs, err := s.repo.GetAll()
	if err != nil {
		s.l.Error("exStore", err.Error())
		return s
	}

	for _, ex := range exs {
		s.exchanges[ex.Name()] = ex
		if ex.Type() == entity.CEX {
			go ex.(entity.Cex).Run()
		}
		l.Debug("exStore.add", fmt.Sprintf("exchange '%d' added", ex.Id()))

	}
	return s
}

func (a *exStore) get(id uint) (entity.Exchange, error) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	for _, ex := range a.exchanges {
		if ex.Id() == id {
			return ex, nil
		}
	}
	return nil, errors.Wrap(errors.ErrNotFound)
}

func (a *exStore) getByName(name string) (entity.Exchange, error) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	ex, ok := a.exchanges[name]
	if !ok {
		return nil, errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage(fmt.Sprintf("exchange %s not found", ex.Name())))
	}
	return ex, nil
}

func (a *exStore) addExchange(ex entity.Exchange) error {
	a.mux.Lock()
	defer a.mux.Unlock()

	if err := a.repo.Add(ex); err != nil {
		return err
	}

	a.exchanges[ex.Name()] = ex
	if ex.Type() == entity.CEX {
		go ex.(entity.Cex).Run()
	}
	a.l.Debug("exStore.add", fmt.Sprintf("exchange '%d' added", ex.Id()))
	return nil
}

func (a *exStore) exists(id uint) bool {
	a.mux.RLock()
	defer a.mux.RUnlock()
	for _, ex := range a.exchanges {
		if ex.Id() == id {
			return true
		}
	}
	return false
}

func (a *exStore) getAll() []entity.Exchange {
	a.mux.RLock()
	defer a.mux.RUnlock()

	var exs []entity.Exchange
	for _, ex := range a.exchanges {
		exs = append(exs, ex)
	}

	return exs
}

func (a *exStore) remove(id uint) error {
	a.mux.Lock()
	defer a.mux.Unlock()
	for _, ex := range a.exchanges {
		if ex.Id() == id {
			if err := a.repo.Remove(ex); err != nil {
				return err
			}
			ex.Remove()
			delete(a.exchanges, ex.Name())
		}
	}

	return fmt.Errorf("exchange %d not found", id)
}
