package fee

import (
	"context"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type data struct {
	Id     uint                      `bson:"_id"`
	Tables map[uint][]*entity.Spread `bson:"tables"`
}

type spreadTable struct {
	mux *sync.RWMutex
	ts  *data
	c   *mongo.Collection
}

func NewSpreadTable(db *mongo.Database) (entity.SpreadTable, error) {
	st := &spreadTable{
		mux: &sync.RWMutex{},
		ts: &data{
			Tables: make(map[uint][]*entity.Spread),
		},
		c: db.Collection("spread-table"),
	}
	if err := st.retrive(); err != nil {
		return nil, err
	}
	return st, nil
}

func (st *spreadTable) Add(tables map[uint][]*entity.Spread) (map[uint][]*entity.Spread, error) {
	st.mux.Lock()
	defer st.mux.Unlock()
	_, ok0 := tables[0]
	_, ok1 := st.ts.Tables[0]
	if !ok0 && !ok1 {
		return nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("you must add level 0 first"))
	}

	for lvl, t := range tables {
		if t == nil || len(t) == 0 {
			return nil, errors.Wrap(errors.ErrBadRequest,
				fmt.Errorf("level %d is empty", lvl))
		}
	}

	for lvl, t := range st.ts.Tables {
		if _, ok := tables[lvl]; !ok {
			tables[lvl] = t
		}
	}

	d := &data{
		Id:     0,
		Tables: tables,
	}
	upsert := true
	_, err := st.c.ReplaceOne(context.Background(), bson.D{{}},
		d, &options.ReplaceOptions{Upsert: &upsert})
	if err != nil {
		return nil, err
	}
	st.ts.Tables = tables
	return st.ts.Tables, nil
}

func (st *spreadTable) GetByPrice(lvl uint, price float64) float64 {
	st.mux.RLock()
	defer st.mux.RUnlock()
	ts, ok := st.ts.Tables[lvl]
	if !ok {
		ts = st.ts.Tables[0]
	}
	for _, s := range ts {
		if price >= s.Start && (price < s.End || s.End == 0) {
			return s.Rate
		}
	}

	return 0
}

func (st *spreadTable) Levels() []uint {
	st.mux.RLock()
	defer st.mux.RUnlock()
	levels := []uint{}
	for lvl := range st.ts.Tables {
		levels = append(levels, lvl)
	}
	return levels
}

func (st *spreadTable) Remove(levels []uint) error {
	st.mux.Lock()
	defer st.mux.Unlock()

	for _, l := range levels {
		if l == 0 {
			return errors.Wrap(errors.ErrBadRequest, fmt.Errorf("you cannot remove level 0"))
		}
	}

	ts := make(map[uint][]*entity.Spread)
	for k, v := range st.ts.Tables {
		ts[k] = v
	}
	for lvl := range ts {
		for _, l := range levels {
			if l == lvl {
				delete(ts, lvl)
			}
		}
	}

	upsert := true
	_, err := st.c.ReplaceOne(context.Background(), bson.D{{}},
		ts, &options.ReplaceOptions{Upsert: &upsert})
	if err != nil {
		return err
	}
	st.ts.Tables = ts
	return nil
}

func (st *spreadTable) GetAll() map[uint][]*entity.Spread {
	st.mux.RLock()
	defer st.mux.RUnlock()
	ts := make(map[uint][]*entity.Spread)
	for k, v := range st.ts.Tables {
		ts[k] = v
	}
	return ts
}

func (st *spreadTable) retrive() error {
	res := st.c.FindOne(context.Background(), bson.D{{}})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil
		}
		return res.Err()
	}

	d := &data{}
	if err := res.Decode(d); err != nil {
		return err
	}
	st.ts = d
	return nil
}
