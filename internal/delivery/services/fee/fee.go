package fee

import (
	"context"
	"exchange-provider/internal/entity"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type fTable struct {
	Id         primitive.ObjectID `bson:"_id"`
	DefaultFee float64
	List       map[string]float64
}

type feeTable struct {
	mux *sync.RWMutex
	c   *mongo.Collection
	fTable
}

func NewFeeTable(db *mongo.Database) (entity.FeeTable, error) {

	f := &feeTable{
		c:   db.Collection("fee-table"),
		mux: &sync.RWMutex{},
		fTable: fTable{
			DefaultFee: 0,
			List:       make(map[string]float64),
		},
	}

	return f, nil
}

func (f *feeTable) GetDefaultFee() float64 {
	f.mux.RLock()
	defer f.mux.RUnlock()
	return f.DefaultFee
}

func (f *feeTable) ChangeDefaultFee(fee float64) error {
	t := fTable{
		DefaultFee: fee,
		List:       f.List,
	}
	_, err := f.c.ReplaceOne(context.Background(), bson.D{{}}, t)
	if err != nil {
		return err
	}

	f.mux.Lock()
	f.DefaultFee = fee
	f.mux.Unlock()
	return nil
}

func (f *feeTable) GetBusFee(busId string) float64 {
	f.mux.RLock()
	defer f.mux.RUnlock()
	fee := f.List[busId]
	if fee == 0 {
		fee = f.DefaultFee
	}
	return fee
}

func (f *feeTable) UpdateBusFee(busId string, fee float64) error {
	f.mux.RLock()
	list := make(map[string]float64)
	for k, v := range f.List {
		list[k] = v
	}
	list[busId] = fee
	t := fTable{
		DefaultFee: f.DefaultFee,
		List:       list,
	}
	f.mux.RUnlock()

	_, err := f.c.InsertOne(context.Background(), t)
	if err != nil {
		return err
	}

	f.mux.Lock()
	f.List = list
	f.mux.Unlock()
	return nil
}

func (f *feeTable) GetAllBusFees() map[string]float64 {
	f.mux.RLock()
	defer f.mux.RUnlock()
	list := make(map[string]float64)
	for k, v := range f.List {
		list[k] = v
	}
	return list
}
