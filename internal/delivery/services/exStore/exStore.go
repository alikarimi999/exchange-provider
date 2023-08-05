package store

import (
	"crypto/rsa"
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type exStore struct {
	repo      *exchangeRepo
	mux       *sync.RWMutex
	exchanges map[string]entity.Exchange
	l         logger.Logger
}

func NewExchangeStore(db *mongo.Database, ws app.WalletStore, pairs entity.PairsRepo,
	repo entity.OrderRepo, fee entity.FeeTable, spread entity.SpreadTable,
	l logger.Logger, prvKey *rsa.PrivateKey, lastUpdate time.Time) (entity.ExchangeStore, error) {

	s := &exStore{
		repo:      newExchangeRepo(db, ws, pairs, repo, fee, spread, l, prvKey),
		mux:       &sync.RWMutex{},
		exchanges: make(map[string]entity.Exchange),
		l:         l,
	}
	s.repo.exs = s
	exs, err := s.repo.getAll(lastUpdate)
	if err != nil {
		return nil, err
	}
	for _, ex := range exs {
		s.exchanges[ex.NID()] = ex
		l.Debug("exStore.add", fmt.Sprintf("lp '%d' added", ex.Id()))

	}
	return s, nil
}

func (a *exStore) Get(id uint) (entity.Exchange, error) {
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

func (a *exStore) GetByNID(name string) (entity.Exchange, error) {
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

func (a *exStore) AddExchange(ex entity.Exchange) error {
	a.mux.Lock()
	defer a.mux.Unlock()

	if err := a.repo.add(ex); err != nil {
		return err
	}

	a.exchanges[ex.NID()] = ex
	a.l.Debug("exStore.addExchange", fmt.Sprintf("lp '%d' added", ex.Id()))
	return nil
}

func (a *exStore) Exists(id uint) bool {
	a.mux.RLock()
	defer a.mux.RUnlock()
	for _, ex := range a.exchanges {
		if ex.Id() == id {
			return true
		}
	}
	return false
}

func (a *exStore) GetAll() []entity.Exchange {
	a.mux.RLock()
	defer a.mux.RUnlock()

	var exs []entity.Exchange
	for _, ex := range a.exchanges {
		exs = append(exs, ex)
	}

	return exs
}

func (a *exStore) GetAllMap() map[string]entity.Exchange {
	a.mux.RLock()
	defer a.mux.RUnlock()

	m := make(map[string]entity.Exchange)
	for _, ex := range a.exchanges {
		m[ex.NID()] = ex
	}
	return m
}

func (a *exStore) EnableDisable(exId uint, enable bool) error {
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
			if err := a.repo.enableDisable(exId, enable); err != nil {
				return err
			}
			ex.EnableDisable(enable)
			return nil
		}
	}
	return errors.Wrap(errors.ErrNotFound, fmt.Errorf("lp not found"))
}
func (a *exStore) EnableDisableAll(enable bool) error {
	a.mux.RLock()
	defer a.mux.RUnlock()
	if err := a.repo.enableDisableAll(enable); err != nil {
		return err
	}
	for _, ex := range a.exchanges {
		ex.EnableDisable(enable)
	}
	return nil
}

func (a *exStore) RemoveAll() error {
	a.mux.Lock()
	defer a.mux.Unlock()
	if err := a.repo.removeAll(); err != nil {
		return err
	}
	for _, ex := range a.exchanges {
		ex.Remove()
	}
	a.exchanges = make(map[string]entity.Exchange)
	return nil
}

func (a *exStore) Remove(id uint) error {
	a.mux.Lock()
	defer a.mux.Unlock()
	for _, ex := range a.exchanges {
		if ex.Id() == id {
			if err := a.repo.eemove(ex); err != nil {
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
