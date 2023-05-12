package pairsRepo

import (
	"context"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"sort"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pairsRepo struct {
	mux *sync.RWMutex
	eps map[string]*exPairs
	c   *mongo.Collection
	l   logger.Logger
}

func PairsRepo(db *mongo.Database, l logger.Logger) (*pairsRepo, error) {
	pr := &pairsRepo{
		mux: &sync.RWMutex{},
		eps: make(map[string]*exPairs),
		c:   db.Collection("pairs"),
		l:   l,
	}
	if err := pr.retrievePairs(); err != nil {
		return nil, err
	}
	return pr, nil
}

func (pr *pairsRepo) AddExchanges(exs []entity.Exchange) {
	for _, ex := range exs {
		ep, ok := pr.eps[ex.NID()]
		if ok {
			ep.ex = ex
		}
	}
}

func (pr *pairsRepo) Add(ex entity.Exchange, ps ...*entity.Pair) error {
	agent := pr.agent("Add")

	pr.mux.Lock()
	defer pr.mux.Unlock()
	ep, ok := pr.eps[ex.NID()]
	if !ok {
		exp := &exchangePairs{
			NID:    ex.NID(),
			ExId:   ex.Id(),
			ExType: ex.Type(),
		}

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
		psi := []interface{}{}
		for _, p := range ps {
			if ep.exists(p.T1.String(), p.T2.String()) {
				continue
			}
			ps2 = append(ps2, p)
			psi = append(psi, pFromEntity(p))
		}

		update := bson.M{"$push": bson.M{"pairs": bson.M{"$each": psi}}}
		res, err := pr.c.UpdateByID(context.Background(), ex.NID(), update)
		if res.MatchedCount == 0 {
			return fmt.Errorf("mongo response MatchedCount is 0")
		}

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

func (pr *pairsRepo) Get(exId uint, t1, t2 string) (*entity.Pair, error) {
	pr.mux.RLock()
	defer pr.mux.RUnlock()
	for _, ep := range pr.eps {
		if ep.exId == exId {
			p, err := ep.get(t1, t2)
			if err != nil {
				return nil, err
			}
			return p, nil
		}
	}
	return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage(
		fmt.Sprintf("exchange '%d' not found", exId)))
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

func (pr *pairsRepo) Update(exId uint, p *entity.Pair) error {
	agent := pr.agent("Update")
	pr.mux.RLock()
	defer pr.mux.RUnlock()
	for _, ep := range pr.eps {
		if ep.exId == exId {
			update := bson.M{"$set": bson.M{"pairs.$[elem]": pFromEntity(p)}}
			arrayFilter := options.Update().SetArrayFilters(options.ArrayFilters{
				Filters: []interface{}{bson.M{"elem.id": p.String()}},
			})
			res, err := pr.c.UpdateByID(context.Background(), ep.exNID, update, arrayFilter)
			if err != nil {
				return err
			}
			if res.MatchedCount == 0 {
				return fmt.Errorf("mongo response MatchedCount is 0")
			}
			ep.update(p)
			pr.l.Debug(agent, fmt.Sprintf("pair '%s' updated in exchange '%s'",
				p.String(), ep.exNID))
			return nil
		}
	}
	return errors.Wrap(errors.ErrUnknown)
}

func (pr *pairsRepo) Remove(exId uint, t1, t2 string, hard bool) error {
	agent := pr.agent("Remove")
	pr.mux.RLock()
	defer pr.mux.RUnlock()
	for _, ep := range pr.eps {
		if ep.exId == exId {
			if ep.exists(t1, t2) {
				if hard {
					res, err := pr.c.UpdateByID(context.Background(), ep.exNID,
						bson.M{"$pull": bson.M{"pairs": bson.M{"id": pairId(t1, t2)}}})
					if err != nil {
						return err
					}
					if res.MatchedCount == 0 {
						return fmt.Errorf("mongo response MatchedCount is 0")
					}
				}
				ep.remove(t1, t2)
				pr.l.Debug(agent, fmt.Sprintf("pair '%s' removed from exchange '%s'",
					pairId(t1, t2), ep.exNID))
				return nil
			}
		}
	}

	return errors.Wrap(errors.ErrNotFound,
		fmt.Errorf("pair '%s/%s' not found in exchange '%d'", t1, t2, exId))
}

func (pr *pairsRepo) UpdateAll(cmd string) error {

	var enable bool
	switch cmd {
	case "remove":
		return pr.RemoveAllExchanges()
	case "enable":
		enable = true
	case "disable":
		enable = false
	default:
		return errors.Wrap(errors.ErrBadRequest, fmt.Errorf("cmd '%s' is invalid", cmd))
	}

	pr.mux.RLock()
	defer pr.mux.RUnlock()

	update := bson.M{"$set": bson.M{"pairs.$[elem].enable": enable}}
	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem.id": bson.M{"$exists": true}}},
	})

	_, err := pr.c.UpdateMany(context.Background(), bson.M{}, update, updateOptions)
	if err != nil {
		return err
	}
	for _, ep := range pr.eps {
		ep.enableDisableAll(enable)
	}
	return nil
}

func (pr *pairsRepo) RemoveAllExchanges() error {
	pr.mux.Lock()
	defer pr.mux.Unlock()
	_, err := pr.c.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	pr.eps = make(map[string]*exPairs)
	return nil
}

func (pr *pairsRepo) RemoveAll(exId uint, hard bool) error {
	agent := pr.agent("RemoveAll")
	pr.mux.Lock()
	defer pr.mux.Unlock()
	for nid, ep := range pr.eps {
		if ep.exId == exId {
			if hard {
				res, err := pr.c.DeleteOne(context.Background(), bson.D{{"_id", ep.exNID}})
				if err != nil {
					return err
				}
				delete(pr.eps, nid)
				if res.DeletedCount == 1 {
					pr.l.Debug(agent, fmt.Sprintf("all pairs of exchange '%s' removed from memory and database", nid))
				}
			} else {
				delete(pr.eps, nid)
				pr.l.Debug(agent, fmt.Sprintf("all pairs of exchange '%s' removed from memory", nid))
			}

		}
	}
	return nil
}

func (pr *pairsRepo) GetAll(exId uint) []*entity.Pair {
	pr.mux.RLock()
	defer pr.mux.RUnlock()
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
