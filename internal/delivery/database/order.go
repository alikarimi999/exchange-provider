package database

import (
	"context"
	"exchange-provider/internal/delivery/database/dto"
	"exchange-provider/internal/entity"
	"fmt"

	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDb struct {
	orders *mongo.Collection
	l      logger.Logger
}

func NewOrderRepo(db *mongo.Database, l logger.Logger) (entity.OrderRepo, error) {
	return &mongoDb{
		orders: db.Collection("orders"),
		l:      l,
	}, nil
}

func (m *mongoDb) Add(order entity.Order) error {
	agent := m.agent("Add")

	id := primitive.NewObjectID()
	order.SetId(id.Hex())
	o := dto.UoToDto(order)
	_, err := m.orders.InsertOne(context.Background(), o)
	if err != nil {
		m.l.Error(agent, fmt.Sprintf("( %s ) ( %s )", order.String(), err.Error()))
		return errors.Wrap(errors.ErrInternal)
	}
	return nil
}

func (m *mongoDb) Get(id *entity.ObjectId) (entity.Order, error) {
	agent := m.agent("Get")
	oId, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	r := m.orders.FindOne(context.Background(), bson.M{"_id": oId})
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
		m.l.Debug(agent, err.Error())
		return nil, err
	}

	if o.Status == entity.OCreated.String() && eo.Expire() {
		o.Status = entity.OExpired.String()
		o = dto.UoToDto(eo)
		_, err := m.orders.ReplaceOne(context.Background(), bson.M{"_id": oId}, o)
		if err != nil {
			m.l.Debug(agent, err.Error())
			return nil, errors.Wrap(errors.ErrInternal)
		}
	}

	return eo, nil
}

func (m *mongoDb) Update(order entity.Order) error {
	agent := m.agent("Update")

	id := order.ID()
	if id != nil {
		ids, _ := primitive.ObjectIDFromHex(order.ID().Id)
		o := dto.UoToDto(order)
		order.Update()
		_, err := m.orders.ReplaceOne(context.Background(), bson.M{"_id": ids}, o)
		if err != nil {
			m.l.Error(agent, fmt.Sprintf("( %s ) ( %s )", order.String(), err.Error()))
			return err
		}

		return nil
	}
	return m.Add(order)
}

func (m *mongoDb) GetWithFilter(key string, value interface{}) ([]entity.Order, error) {
	agent := m.agent("GetWithFilter")

	cur, err := m.orders.Find(context.Background(), bson.M{key: value})
	if err != nil {
		m.l.Debug(agent, err.Error())
		return nil, cur.Err()
	} else if cur.Err() != nil {
		m.l.Debug(agent, cur.Err().Error())
		return nil, cur.Err()
	}

	os := []*dto.Order{}
	if err := cur.All(context.Background(), &os); err != nil {
		return nil, err
	}

	if len(os) == 0 {
		return nil, errors.Wrap(errors.ErrNotFound)
	}

	eos := []entity.Order{}
	for _, o := range os {
		eo, err := o.ToEntity()
		if err != nil {
			m.l.Debug(agent, cur.Err().Error())
			continue
		}
		eos = append(eos, eo)
	}
	return eos, nil
}

// check if any deposit has this tx_id
func (m *mongoDb) TxIdExists(txId string) (bool, error) {
	return m.checkDocumentByTxID(txId)
}

func (m *mongoDb) checkDocumentByTxID(txId string) (bool, error) {
	agent := m.agent("checkDocumentByTtxID")
	filter := bson.M{
		"$or": []bson.M{
			{"order.steps.0.srctxid": txId},
			{"order.steps.1.srctxid": txId},
			{"order.deposit.txid": txId},
		},
	}
	err := m.orders.FindOne(context.Background(), filter, nil).Err()
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		m.l.Error(agent, err.Error())
		return false, err
	}
	return true, nil
}
