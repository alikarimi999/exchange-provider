package pairconf

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

type exchangePairs struct {
	*sync.RWMutex
	list         []*entity.Pair
	pricePending atomic.Int32
	ex           entity.Exchange
}

type repo struct {
	list cmap.ConcurrentMap
	t    *time.Ticker
	l    logger.Logger
}

func NewPairRepo(l logger.Logger, t *time.Ticker) entity.PairRepo {
	r := &repo{
		list: cmap.New(),
		t:    t,
		l:    l,
	}
	go r.run()
	return r
}

func (r *repo) run() {
	for range r.t.C {
		exs := r.list.Keys()
		for _, ex := range exs {
			ePairsI, _ := r.list.Get(ex)
			exPairs := ePairsI.(*exchangePairs)
			if exPairs.pricePending.Load() == 1 {
				continue
			}

			go func(exPairs *exchangePairs) {
				exPairs.pricePending.Store(1)
				defer exPairs.pricePending.Store(0)
				ps := []*entity.Pair{}
				exPairs.RLock()
				for _, p := range exPairs.list {
					ps = append(ps, p.Snapshot())
				}
				exPairs.RUnlock()

				ps, err := exPairs.ex.Price(ps...)
				if err != nil {
					return
				}

				exPairs.Lock()
				for _, p := range ps {
					for _, pl := range exPairs.list {
						if p.Equal(pl) {
							pl.Price1 = p.Price1
							pl.Price2 = p.Price2
						}
					}
				}
				exPairs.Unlock()

			}(exPairs)
		}
	}
}

func (r *repo) Add(ex entity.Exchange, ps ...*entity.Pair) {
	epI, ok := r.list.Get(ex.Id())
	if !ok {
		ep := &exchangePairs{
			RWMutex: &sync.RWMutex{},
			ex:      ex,
		}
		for _, p := range ps {
			p.Exchange = ex.Id()
			ep.list = append(ep.list, p.Snapshot())
		}
		sortList(ep.list)
		r.list.Set(ex.Id(), ep)
		return
	}
	ep := epI.(*exchangePairs)
	ep.Lock()
	for _, p := range ps {
		ep.list = append(ep.list, p.Snapshot())
	}
	sortList(ep.list)
	ep.Unlock()
}

func (r *repo) Get(ex string, t1, t2 *entity.Token) (*entity.Pair, error) {
	exPairI, ok := r.list.Get(ex)
	if ok {
		exPair := exPairI.(*exchangePairs)
		exPair.RLock()
		defer exPair.RUnlock()
		for _, p := range exPair.list {
			if (p.T1.Equal(t1) && p.T2.Equal(t2)) || (p.T1.Equal(t2) && p.T2.Equal(t1)) {
				return p.Snapshot(), nil
			}
		}
	}
	return nil, errors.Wrap(errors.ErrNotFound)
}

func (r *repo) Exists(ex string, t1, t2 *entity.Token) bool {
	exPairI, ok := r.list.Get(ex)
	if ok {
		exPair := exPairI.(*exchangePairs)
		exPair.RLock()
		defer exPair.RUnlock()
		for _, p := range exPair.list {
			if (p.T1.Equal(t1) && p.T2.Equal(t2)) || (p.T1.Equal(t2) && p.T2.Equal(t1)) {
				return true
			}
		}
	}
	return false
}

func (r *repo) Remove(ex string, t1, t2 *entity.Token) error {
	exPairI, ok := r.list.Get(ex)
	if ok {
		exPair := exPairI.(*exchangePairs)
		exPair.Lock()
		defer exPair.Unlock()
		for i, p := range exPair.list {
			if (p.T1.Equal(t1) && p.T2.Equal(t2)) || (p.T1.Equal(t2) && p.T2.Equal(t1)) {
				exPair.list = append(exPair.list[:i], exPair.list[i+1:]...)
				return nil
			}
		}
	}
	return errors.Wrap(errors.ErrNotFound)
}

func (r *repo) RemoveExchange(ex string) error {
	r.list.Remove(ex)
	return nil
}

func (r *repo) GetPaginated(p *entity.Paginated) error {
	exs := r.exchanges(p)
	ps := []*entity.Pair{}

	end := p.PerPage * p.Page
	start := end - p.PerPage

	for _, ex := range exs {
		ePairsI, ok := r.list.Get(ex)
		if ok {
			ePairs := ePairsI.(*exchangePairs)
			ePairs.RLock()
			p.Total += int64(len(ePairs.list))
			ps = append(ps, ePairs.list...)
			ePairs.RUnlock()
		}
	}

	if int(start) < len(ps) {
		if len(ps) < int(end) {
			end = int64(len(ps))
		}
		for _, ep := range ps[start:end] {
			p.Pairs = append(p.Pairs, ep.Snapshot())
		}
	}
	return nil
}

func (r *repo) exchanges(p *entity.Paginated) []string {
	exs := []string{}
	if len(p.Filters) == 0 || len(p.Filters[0].Values) == 0 {
		exs = append(exs, r.list.Keys()...)
		sort.Strings(exs)
		return exs
	}
	for _, e := range p.Filters[0].Values {
		exs = append(exs, e.(string))
	}
	sort.Strings(exs)
	return exs
}

func pairId(t1, t2 string) string {
	if strings.Compare(t1, t2) == -1 {
		return fmt.Sprintf("%s%s%s", t1, types.Delimiter, t2)
	}
	return fmt.Sprintf("%s%s%s", t2, types.Delimiter, t1)
}

func sortList(list []*entity.Pair) {
	sort.Slice(list, func(i, j int) bool {
		pi := list[i]
		pj := list[j]
		return strings.Compare(pi.String(), pj.String()) == -1
	})
}
