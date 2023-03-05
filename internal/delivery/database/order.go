package database

import (
	"context"
	"exchange-provider/internal/delivery/database/dto"
	"exchange-provider/internal/entity"

	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDb struct {
	orders *mongo.Collection
	c      *cache
	l      logger.Logger
}

func NewOrderRepo(db *mongo.Database, l logger.Logger) (entity.OrderRepo, error) {
	m := &mongoDb{
		orders: db.Collection("orders"),
		c:      newCache(db),
		l:      l,
	}
	return m, m.retrivePendingWithd()
}

func (m *mongoDb) Add(order entity.Order) error {
	agent := m.agent("Add")

	id := primitive.NewObjectID()
	order.SetId(id.Hex())
	o, err := dto.UoToDto(order)
	if err != nil {
		m.l.Error(agent, err.Error())
		return err
	}
	_, err = m.orders.InsertOne(context.Background(), o)
	if err != nil {
		m.l.Error(agent, err.Error())
		return err
	}
	return nil
}

func (m *mongoDb) Get(id *entity.ObjectId) (entity.Order, error) {
	agent := m.agent("Get")
	oId, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	r := m.orders.FindOne(context.Background(), bson.D{{"_id", oId}})
	if r.Err() != nil {
		if r.Err() == mongo.ErrNoDocuments {
			return nil, errors.Wrap(errors.ErrNotFound)
		}
		m.l.Error(agent, r.Err().Error())
		return nil, r.Err()
	}

	o := &dto.Order{}
	if err := r.Decode(o); err != nil {
		return nil, err
	}
	eo, err := o.ToEntity()
	if err != nil {
		m.l.Error(agent, r.Err().Error())
		return nil, err
	}
	return eo, nil
}

func (m *mongoDb) GetAll(UserId string) ([]entity.Order, error) {
	agent := m.agent("GetAll")

	osDTO := []*dto.Order{}
	cur, err := m.orders.Find(context.Background(), bson.D{{"userid", UserId}})
	if err != nil {
		m.l.Error(agent, err.Error())
		return nil, err
	}

	if err := cur.All(context.Background(), &osDTO); err != nil {
		m.l.Error(agent, err.Error())
		return nil, err
	}

	os := []entity.Order{}
	for _, o := range osDTO {
		eo, err := o.ToEntity()
		if err != nil {
			continue
		}
		os = append(os, eo)
	}
	return os, nil
}

func (m *mongoDb) Update(order entity.Order) error {
	agent := m.agent("Update")

	id, _ := primitive.ObjectIDFromHex(order.ID().Id)
	o, err := dto.UoToDto(order)
	if err != nil {
		m.l.Error(agent, err.Error())
		return err
	}
	res, err := m.orders.ReplaceOne(context.Background(), bson.D{{"_id", id}}, o)
	if err != nil {
		m.l.Error(agent, err.Error())
		return err
	}
	if res.ModifiedCount == 0 {
		return errors.Wrap(errors.ErrNotFound)
	}

	return nil
}

// check if any deposit has this tx_id
func (m *mongoDb) TxIdExists(txId string) (bool, error) {
	agent := m.agent("CheckTxId")

	res := m.orders.FindOne(context.Background(), bson.D{{"order.deposit.txid", txId}})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		m.l.Error(agent, res.Err().Error())
		return false, res.Err()
	}
	return true, nil
}
