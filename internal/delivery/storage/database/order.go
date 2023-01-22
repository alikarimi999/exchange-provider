package database

import (
	"context"
	"exchange-provider/internal/delivery/storage/database/dto"
	"exchange-provider/internal/entity"

	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDb struct {
	orders *mongo.Collection
	l      logger.Logger
}

func NewUserRepo(db *mongo.Database, l logger.Logger) entity.OrderRepo {
	return &MongoDb{
		orders: db.Collection("orders"),
		l:      l,
	}
}

func (m *MongoDb) Add(order entity.Order) error {
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

func (m *MongoDb) Get(id string) (entity.Order, error) {
	agent := m.agent("Get")
	oId, err := primitive.ObjectIDFromHex(id)
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
	r.Decode(o)
	eo, err := o.ToEntity()
	if err != nil {
		m.l.Error(agent, r.Err().Error())
		return nil, err
	}
	return eo, nil
}

func (m *MongoDb) GetAll(UserId uint64) ([]entity.Order, error) {
	agent := m.agent("GetAll")

	osDTO := []*dto.Order{}
	cur, err := m.orders.Find(context.Background(), bson.D{{"userId", UserId}})
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

func (m *MongoDb) Update(order entity.Order) error {
	agent := m.agent("Update")

	id, _ := primitive.ObjectIDFromHex(order.ID())
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
func (m *MongoDb) TxIdExists(txId string) (bool, error) {
	agent := m.agent("CheckTxId")

	res := m.orders.FindOne(context.Background(), bson.D{{"order.deposit.txId", txId}})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		m.l.Error(agent, res.Err().Error())
		return false, res.Err()
	}
	return true, nil
}
