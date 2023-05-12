package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"
)

type ExStore struct {
	repo      ExchangeRepo
	mux       *sync.RWMutex
	exchanges map[string]entity.Exchange
	l         logger.Logger
}

func newExStore(l logger.Logger, exRepo ExchangeRepo, exs []entity.Exchange) *ExStore {
	s := &ExStore{
		repo:      exRepo,
		mux:       &sync.RWMutex{},
		exchanges: make(map[string]entity.Exchange),
		l:         l,
	}

	for _, ex := range exs {
		s.exchanges[ex.NID()] = ex
		l.Debug("exStore.add", fmt.Sprintf("lp '%d' added", ex.Id()))

	}
	return s
}

func (a *ExStore) get(id uint) (entity.Exchange, error) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	for _, ex := range a.exchanges {
		if ex.Id() == id {
			return ex, nil
		}
	}
	err := fmt.Errorf("lp '%d' not found", id)
	return nil, errors.Wrap(errors.ErrNotFound, err,
		errors.NewMesssage(err.Error()))
}

func (a *ExStore) getByNID(name string) (entity.Exchange, error) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	ex, ok := a.exchanges[name]
	if !ok {
		err := fmt.Errorf("lp '%s' not found", ex.NID())
		return nil, errors.Wrap(errors.ErrNotFound, err,
			errors.NewMesssage(err.Error()))
	}
	return ex, nil
}

func (a *ExStore) addExchange(ex entity.Exchange) error {
	a.mux.Lock()
	defer a.mux.Unlock()

	if err := a.repo.Add(ex); err != nil {
		return err
	}

	a.exchanges[ex.NID()] = ex
	a.l.Debug("exStore.addExchange", fmt.Sprintf("lp '%d' added", ex.Id()))
	return nil
}

func (a *ExStore) exists(id uint) bool {
	a.mux.RLock()
	defer a.mux.RUnlock()
	for _, ex := range a.exchanges {
		if ex.Id() == id {
			return true
		}
	}
	return false
}

func (a *ExStore) getAll() []entity.Exchange {
	a.mux.RLock()
	defer a.mux.RUnlock()

	var exs []entity.Exchange
	for _, ex := range a.exchanges {
		exs = append(exs, ex)
	}

	return exs
}

func (a *ExStore) enableDisable(exId uint, enable bool) error {
	a.mux.RLock()
	defer a.mux.RUnlock()
	for _, ex := range a.exchanges {
		if ex.Id() == exId {
			if ex.IsEnable() == enable {
				if enable {
					return errors.Wrap(errors.ErrBadRequest, fmt.Errorf("lp is enable"))
				} else {
					return errors.Wrap(errors.ErrBadRequest, fmt.Errorf("lp is disable"))
				}
			}
			if err := a.repo.EnableDisable(exId, enable); err != nil {
				return err
			}
			ex.EnableDisable(enable)
			return nil
		}
	}
	return errors.Wrap(errors.ErrNotFound, fmt.Errorf("lp not found"))
}
func (a *ExStore) enableDisableAll(enable bool) error {
	a.mux.RLock()
	defer a.mux.RUnlock()
	if err := a.repo.EnableDisableAll(enable); err != nil {
		return err
	}
	for _, ex := range a.exchanges {
		ex.EnableDisable(enable)
	}
	return nil
}

func (a *ExStore) RemoveAll() error {
	a.mux.Lock()
	defer a.mux.Unlock()
	if err := a.repo.RemoveAll(); err != nil {
		return err
	}
	for _, ex := range a.exchanges {
		ex.Remove()
	}
	a.exchanges = make(map[string]entity.Exchange)
	return nil
}

func (a *ExStore) Remove(id uint) error {
	a.mux.Lock()
	defer a.mux.Unlock()
	for _, ex := range a.exchanges {
		if ex.Id() == id {
			if err := a.repo.Remove(ex); err != nil {
				return err
			}
			ex.Remove()
			delete(a.exchanges, ex.NID())
			return nil
		}
	}

	err := fmt.Errorf("lp '%d' not found", id)
	return errors.Wrap(errors.ErrNotFound, err,
		errors.NewMesssage(err.Error()))
}
