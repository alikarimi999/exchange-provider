package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
)

type exStore struct {
	repo      ExchangeRepo
	exchanges map[uint]entity.Exchange
	l         logger.Logger
}

func newExStore(l logger.Logger, exRepo ExchangeRepo) *exStore {
	s := &exStore{
		repo:      exRepo,
		exchanges: make(map[uint]entity.Exchange),
		l:         l,
	}

	exs, err := s.repo.GetAll()
	if err != nil {
		s.l.Error("exStore", err.Error())
		return s
	}

	for _, ex := range exs {
		s.exchanges[ex.Id()] = ex
		if ex.Type() == entity.CEX {
			go ex.(entity.Cex).Run()
		}
		l.Debug("exStore.add", fmt.Sprintf("exchange '%s' added", ex.Id()))

	}
	return s
}

func (a *exStore) get(id uint) (entity.Exchange, error) {
	if _, ok := a.exchanges[id]; ok {
		return a.exchanges[id], nil
	}
	return nil, errors.Wrap(errors.ErrNotFound)
}

func (a *exStore) AddExchange(ex entity.Exchange) error {
	if err := a.repo.Add(ex); err != nil {
		return err
	}
	a.exchanges[ex.Id()] = ex
	if ex.Type() == entity.CEX {
		go ex.(entity.Cex).Run()
	}
	a.l.Debug("exStore.add", fmt.Sprintf("exchange '%s' added", ex.Id()))
	return nil
}

func (a *exStore) exists(id uint) bool {
	if _, ok := a.exchanges[id]; ok {
		return true
	}
	return false
}

func (a *exStore) getAll() []entity.Exchange {
	var exs []entity.Exchange
	for _, ex := range a.exchanges {
		exs = append(exs, ex)
	}

	return exs
}

func (a *exStore) getByNames(names ...string) []entity.Exchange {
	if len(names) == 0 {
		return a.getAll()
	}

	exs := []entity.Exchange{}
	for _, ex := range a.exchanges {
		for _, name := range names {
			if ex.Name() == name {
				exs = append(exs, ex)
			}
		}
	}
	return exs
}

func (a *exStore) remove(id uint) error {
	if ex, ok := a.exchanges[id]; ok {
		if err := a.repo.Remove(ex); err != nil {
			return err
		}

		ex.Remove()
		delete(a.exchanges, id)
		return nil
	}

	return fmt.Errorf("exchange %d not found", id)
}
