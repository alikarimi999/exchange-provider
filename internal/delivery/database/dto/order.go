package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Type  entity.OrderType
	Order bson.Raw
}

func (o *Order) ToEntity() (entity.Order, error) {
	switch o.Type {
	case entity.CEXOrder:
		cOrder := &entity.CexOrder{}
		if err := bson.Unmarshal(o.Order, cOrder); err != nil {
			return nil, err
		}
		return cOrder, nil
	case entity.EVMOrder:
		eOrder := &entity.EvmOrder{}
		if err := bson.Unmarshal(o.Order, eOrder); err != nil {
			return nil, err
		}
		return eOrder, nil
	default:
		return nil, errors.Wrap(errors.ErrBadRequest)
	}
}

func UoToDto(o entity.Order) (*Order, error) {
	id, err := primitive.ObjectIDFromHex(o.ID().Id)
	if err != nil {
		return nil, err
	}
	raw, err := bson.Marshal(o)
	if err != nil {
		return nil, err
	}

	return &Order{
		Id:    id,
		Type:  o.Type(),
		Order: raw,
	}, nil

}
