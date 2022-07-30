package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"order_service/pkg/logger"
	"sync"
)

type exStore struct {
	mux      *sync.Mutex
	exs      map[string]entity.Exchange
	addCh    chan entity.Exchange
	updateCh chan entity.Exchange
	l        logger.Logger
}

func newExStore(l logger.Logger) *exStore {
	return &exStore{
		mux:      &sync.Mutex{},
		exs:      make(map[string]entity.Exchange),
		addCh:    make(chan entity.Exchange),
		updateCh: make(chan entity.Exchange),
		l:        l,
	}
}

func (a *exStore) get(exchange string) (entity.Exchange, error) {
	a.mux.Lock()
	defer a.mux.Unlock()
	ex, ok := a.exs[exchange]
	if !ok {
		return nil, errors.Wrap(errors.ErrNotFound)
	}
	return ex, nil
}

func (a *exStore) start(wg *sync.WaitGroup) {
	const agent = "Exchange-Sotore.start"
	defer wg.Done()
	a.l.Debug(agent, "started")

	for {
		select {
		case ex := <-a.addCh:

			a.mux.Lock()
			a.exs[ex.ID()] = ex
			a.mux.Unlock()

			a.l.Debug(agent, fmt.Sprintf("exchange %s added", ex.ID()))
			wg.Add(1)
			go ex.Run(wg)

		case ex := <-a.updateCh:
			a.l.Debug(agent, fmt.Sprintf("exchange %s updated", ex.ID()))
			// wg.Add(1)
			// go ex.Run(wg)

		}
	}
}

func (a *exStore) add(exs ...entity.Exchange) {
	for _, ex := range exs {
		a.addCh <- ex
	}
}

func (a *exStore) update(exs ...entity.Exchange) {
	for _, ex := range exs {
		a.updateCh <- ex
	}
}

func (a *exStore) exists(exchange string) bool {
	_, ok := a.exs[exchange]
	return ok
}

func (a *exStore) getAll() []string {
	a.mux.Lock()
	defer a.mux.Unlock()

	var exs []string
	for ex := range a.exs {
		exs = append(exs, ex)
	}

	return exs
}
