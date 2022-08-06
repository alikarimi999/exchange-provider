package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"order_service/pkg/logger"
	"sync"
	"time"
)

const (
	ExchangeStatusActive   = "active"
	ExchangeStatusDeactive = "deactive"
	ExchangeStatusDisable  = "disable"
)

type Exchange struct {
	entity.Exchange
	CurrentStatus  string
	PreviousStatus string
	LastChangeTime time.Time
}

type exStore struct {
	mux *sync.Mutex

	actives   map[string]*Exchange
	deactives map[string]*Exchange
	disableds map[string]*Exchange

	addCh chan *Exchange

	l logger.Logger
}

func newExStore(l logger.Logger) *exStore {
	return &exStore{
		mux: &sync.Mutex{},

		actives:   make(map[string]*Exchange),
		deactives: make(map[string]*Exchange),
		disableds: make(map[string]*Exchange),

		addCh: make(chan *Exchange),
		l:     l,
	}
}

func (a *exStore) activate(nid string) error {
	ex, err := a.get(nid)
	if err != nil {
		return err
	}
	a.mux.Lock()
	defer a.mux.Unlock()
	switch ex.CurrentStatus {
	case ExchangeStatusActive:
		return errors.Wrap(errors.NewMesssage(fmt.Sprintf("exchange %s already active", nid)))
	case ExchangeStatusDeactive:
		delete(a.deactives, nid)
		ex.PreviousStatus = ex.CurrentStatus
		ex.CurrentStatus = ExchangeStatusActive
		ex.LastChangeTime = time.Now()

		a.actives[nid] = ex
		return nil
	case ExchangeStatusDisable:
		delete(a.disableds, nid)
		ex.PreviousStatus = ex.CurrentStatus
		ex.CurrentStatus = ExchangeStatusActive
		ex.LastChangeTime = time.Now()
		a.add(ex)
	}

	return nil
}

func (a *exStore) deactivate(nid string) error {
	ex, err := a.get(nid)
	if err != nil {
		return err
	}

	a.mux.Lock()
	defer a.mux.Unlock()
	switch ex.CurrentStatus {
	case ExchangeStatusActive:
		delete(a.actives, nid)
		ex.PreviousStatus = ex.CurrentStatus
		ex.CurrentStatus = ExchangeStatusDeactive
		ex.LastChangeTime = time.Now()
		a.deactives[nid] = ex
		return nil
	case ExchangeStatusDeactive:
		return errors.Wrap(errors.NewMesssage(fmt.Sprintf("exchange %s already inactive", nid)))
	case ExchangeStatusDisable:
		delete(a.disableds, nid)
		ex.PreviousStatus = ex.CurrentStatus
		ex.CurrentStatus = ExchangeStatusDeactive
		ex.LastChangeTime = time.Now()
		a.add(ex)
	}

	return nil
}

func (a *exStore) disable(nid string) error {
	ex, err := a.get(nid)
	if err != nil {
		return err
	}

	a.mux.Lock()
	defer a.mux.Unlock()
	switch ex.CurrentStatus {
	case ExchangeStatusActive:
		delete(a.actives, nid)
		ex.Stop()
		ex.PreviousStatus = ex.CurrentStatus
		ex.CurrentStatus = ExchangeStatusDisable
		ex.LastChangeTime = time.Now()
		a.disableds[nid] = ex
		return nil
	case ExchangeStatusDeactive:
		delete(a.deactives, nid)
		ex.Stop()
		ex.PreviousStatus = ex.CurrentStatus
		ex.CurrentStatus = ExchangeStatusDisable
		ex.LastChangeTime = time.Now()
		a.disableds[nid] = ex
		return nil
	case ExchangeStatusDisable:
		return errors.Wrap(errors.NewMesssage(fmt.Sprintf("exchange %s already disabled", nid)))
	}

	return nil
}

func (a *exStore) get(nid string) (*Exchange, error) {
	a.mux.Lock()
	defer a.mux.Unlock()

	_, ok := a.actives[nid]
	if ok {
		return a.actives[nid], nil
	}

	_, ok = a.deactives[nid]
	if ok {
		return a.deactives[nid], nil
	}

	_, ok = a.disableds[nid]
	if ok {
		return a.disableds[nid], nil
	}

	return nil, errors.Wrap(errors.ErrNotFound)
}

func (a *exStore) getActives() []*Exchange {
	a.mux.Lock()
	defer a.mux.Unlock()

	var exs []*Exchange
	for _, ex := range a.actives {
		exs = append(exs, ex)
	}

	return exs
}

func (a *exStore) getDeactives() []*Exchange {
	a.mux.Lock()
	defer a.mux.Unlock()

	var exs []*Exchange
	for _, ex := range a.deactives {
		exs = append(exs, ex)
	}

	return exs
}

func (a *exStore) start(wg *sync.WaitGroup) {
	const agent = "Exchange-Sotore.start"
	defer wg.Done()
	a.l.Debug(agent, "started")

	for {
		select {
		case ex := <-a.addCh:

			a.mux.Lock()
			switch ex.CurrentStatus {
			case ExchangeStatusActive:
				a.actives[ex.NID()] = ex
			case ExchangeStatusDeactive:
				a.deactives[ex.NID()] = ex
			case ExchangeStatusDisable:
				a.disableds[ex.NID()] = ex
				continue
			default:
				continue
			}

			a.mux.Unlock()
			wg.Add(1)
			go ex.Run(wg)
			a.l.Info(agent, fmt.Sprintf("exchange %s started", ex.NID()))
		}
	}
}

func (a *exStore) add(ex *Exchange) {
	a.addCh <- ex
}

func (a *exStore) exists(nid string) (bool, string) {

	_, ok := a.actives[nid]
	if ok {
		return true, ExchangeStatusActive
	}

	_, ok = a.deactives[nid]
	if ok {
		return true, ExchangeStatusDeactive
	}

	_, ok = a.disableds[nid]
	if ok {
		return true, ExchangeStatusDisable
	}

	return false, ""
}

func (a *exStore) getAllList() []string {
	a.mux.Lock()
	defer a.mux.Unlock()

	var exs []string
	for nid := range a.actives {
		exs = append(exs, nid)
	}
	for nid := range a.deactives {
		exs = append(exs, nid)
	}
	for nid := range a.disableds {
		exs = append(exs, nid)
	}
	return exs
}

func (a *exStore) getAllNames() []string {
	a.mux.Lock()
	defer a.mux.Unlock()

	var exs []string
	for _, ex := range a.actives {
		exs = append(exs, ex.Name())
	}
	for _, ex := range a.deactives {
		exs = append(exs, ex.Name())
	}
	for _, ex := range a.disableds {
		exs = append(exs, ex.Name())
	}
	return exs

}

func (a *exStore) all(names ...string) []*Exchange {
	if len(names) == 0 {
		names = a.getAllNames()
	}

	a.mux.Lock()
	defer a.mux.Unlock()

	exs := []*Exchange{}

	for _, ex := range a.actives {
		for _, name := range names {
			if ex.Name() == name {
				exs = append(exs, ex)
			}
		}
	}

	for _, ex := range a.deactives {
		for _, name := range names {
			if ex.Name() == name {
				exs = append(exs, ex)
			}
		}
	}

	for _, ex := range a.disableds {
		for _, name := range names {
			if ex.Name() == name {
				exs = append(exs, ex)
			}
		}
	}

	return exs
}

func (a *exStore) remove(nid string) error {
	a.mux.Lock()
	defer a.mux.Unlock()

	_, ok := a.actives[nid]
	if ok {
		delete(a.actives, nid)
		return nil
	}

	_, ok = a.deactives[nid]
	if ok {
		delete(a.deactives, nid)
		return nil
	}
	_, ok = a.disableds[nid]
	if ok {
		delete(a.disableds, nid)
		return nil
	}
	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage(fmt.Sprintf("exchange %s not found", nid)))
}
