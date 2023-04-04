package pairsRepo

import (
	"context"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sort"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type pairsRepo struct {
	mux *sync.RWMutex
	eps map[string]*exPairs
	c   *mongo.Collection
	l   logger.Logger
}

func PairsRepo(db *mongo.Database, l logger.Logger) entity.PairsRepo {
	pr := &pairsRepo{
		mux: &sync.RWMutex{},
		eps: make(map[string]*exPairs),
		c:   db.Collection("pairs"),
		l:   l,
	}
	if err := pr.retrievePairs(); err != nil {
		panic(err)
	}
	return pr
}

func (pr *pairsRepo) Add(ex entity.Exchange, ps ...*entity.Pair) error {
	agent := pr.agent("Add")
	exp := &exchangePairs{
		NID:  ex.NID(),
		ExId: ex.Id(),
	}

	pr.mux.Lock()
	defer pr.mux.Unlock()
	ep, ok := pr.eps[ex.NID()]
	if !ok {
		for _, p := range ps {
			exp.Pairs = append(exp.Pairs, pFromEntity(p))
		}

		_, err := pr.c.InsertOne(context.Background(), exp)
		if err != nil {
			pr.l.Error(agent, err.Error())
			return err
		}
		ep = newExPairs(ex)
		pr.eps[ex.NID()] = ep
		pr.sortEps()
		ep.add(ps...)
		for _, p := range ps {
			pr.l.Debug(agent, fmt.Sprintf("pair '%s' added to exchange '%s'",
				pairId(p.T1.String(), p.T2.String()), ex.NID()))

		}
	} else {
		ps2 := []*entity.Pair{}
		for _, p := range ps {
			if ep.exists(p.T1.String(), p.T2.String()) {
				continue
			}
			ps2 = append(ps2, p)
			exp.Pairs = append(exp.Pairs, pFromEntity(p))
		}

		_, err := pr.c.UpdateByID(context.Background(), ex.NID(),
			bson.M{"$push": bson.M{"pairs": bson.M{"$each": exp.Pairs}}})

		if err != nil {
			pr.l.Error(pr.agent("Add"), err.Error())
			return err
		}
		ep.add(ps2...)
		for _, p := range ps2 {
			pr.l.Debug(agent, fmt.Sprintf("pair '%s' added to exchange '%s'",
				pairId(p.T1.String(), p.T2.String()), ex.NID()))

		}
	}
	return nil
}

func (pr *pairsRepo) Get(exId uint, t1, t2 string) (*entity.Pair, bool) {
	pr.mux.RLock()
	defer pr.mux.RUnlock()
	for _, ep := range pr.eps {
		if ep.exId == exId {
			return ep.get(t1, t2)
		}
	}
	return nil, false
}

func (pr *pairsRepo) Exists(exId uint, t1, t2 string) bool {
	pr.mux.RLock()
	defer pr.mux.RUnlock()
	for _, ep := range pr.eps {
		if ep.exId == exId {
			return ep.exists(t1, t2)
		}
	}
	return false
}

func (pr *pairsRepo) Update(exId uint, p *entity.Pair) {
	pr.mux.Lock()
	defer pr.mux.Unlock()
	for _, ep := range pr.eps {
		if ep.exId == exId {
			ep.update(p)
		}
	}
}

func (pr *pairsRepo) Remove(exId uint, t1, t2 string) error {
	agent := pr.agent("Remove")
	pr.mux.Lock()
	defer pr.mux.Unlock()
	for _, ep := range pr.eps {
		if ep.exId == exId {
			_, err := pr.c.UpdateByID(context.Background(), ep.exNID,
				bson.M{"$pull": bson.M{"pairs": bson.M{"_id": pairId(t1, t2)}}})
			if err != nil {
				pr.l.Error(pr.agent("Remove"), err.Error())
				return err
			}
			ep.remove(t1, t2)
			pr.l.Debug(agent, fmt.Sprintf("pair '%s' removed from exchange '%s'",
				pairId(t1, t2), ep.exNID))

		}
	}
	return nil
}

func (pr *pairsRepo) RemoveAll(exId uint) error {
	agent := pr.agent("RemoveAll")
	pr.mux.Lock()
	defer pr.mux.Unlock()
	for nid, ep := range pr.eps {
		if ep.exId == exId {
			res, err := pr.c.DeleteOne(context.Background(), bson.D{{"_id", ep.exNID}})
			if err != nil {
				return err
			}
			delete(pr.eps, nid)
			if res.DeletedCount == 1 {
				pr.l.Debug(agent, fmt.Sprintf("all pairs of exchange '%s' removed", nid))
			}
		}
	}
	return nil
}

func (pr *pairsRepo) GetAll(exId uint) []*entity.Pair {
	pr.mux.Lock()
	defer pr.mux.Unlock()
	for _, ep := range pr.eps {
		if ep.exId == exId {
			return ep.getAll()
		}
	}
	return []*entity.Pair{}
}

func (pr *pairsRepo) sortEps() {
	m1 := pr.eps
	m2 := make(map[string]*exPairs)

	ids := []int{}
	for _, ep := range m1 {
		ids = append(ids, int(ep.exId))
	}
	sort.Ints(ids)

	for _, id := range ids {
		for name, ep := range pr.eps {
			if ep.exId == uint(id) {
				m2[name] = m1[name]
			}
		}
	}

	pr.eps = m2
}
