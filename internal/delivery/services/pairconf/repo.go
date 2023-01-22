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

	cmap "github.com/orcaman/concurrent-map"
)

type exchangePairs struct {
	*sync.RWMutex
	list []entity.Pair
	ex   entity.Exchange
}

type repo struct {
	list cmap.ConcurrentMap
	l    logger.Logger
}

func NewPairRepo(l logger.Logger) entity.PairRepo {
	return &repo{
		list: cmap.New(),
		l:    l,
	}
}

func (r *repo) Add(ex entity.Exchange, ps ...*entity.Pair) {
	agent := "repo.Add"
	epI, ok := r.list.Get(ex.Id())
	if !ok {
		ep := &exchangePairs{
			RWMutex: &sync.RWMutex{},
			ex:      ex,
		}
		for _, p := range ps {
			p.Exchange = ex.Id()
			ep.list = append(ep.list, *p)
			r.l.Debug(agent, fmt.Sprintf("'%s' added", p.String()))
		}
		sortList(ep.list)
		r.list.Set(ex.Id(), ep)
		return
	}
	ep := epI.(*exchangePairs)
	ep.Lock()
	for _, p := range ps {
		ep.list = append(ep.list, *p)
		r.l.Debug(agent, fmt.Sprintf("'%s' added", p.String()))
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
				return &p, nil
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

func (r *repo) GetPaginated(p *entity.Paginated) error {
	exs := r.exchanges(p)
	ps := []entity.Pair{}

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

func sortList(list []entity.Pair) {
	sort.Slice(list, func(i, j int) bool {
		pi := list[i]
		pj := list[j]
		return strings.Compare(pi.String(), pj.String()) == -1
	})
}
