package dto

import (
	kt "exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	et "exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Type        entity.OrderType
	ExchangeNID string
	Status      string
	Order       bson.Raw
}

func (o *Order) ToEntity() (entity.Order, error) {
	switch o.Type {
	case entity.CEXOrder:
		switch strings.Split(o.ExchangeNID, "-")[0] {
		case "kucoin":
			ko := &kt.Order{}
			if err := bson.Unmarshal(o.Order, ko); err != nil {
				return nil, err
			}
			return ko, nil
		}
	case entity.EVMOrder:
		eo := &et.Order{}
		if err := bson.Unmarshal(o.Order, eo); err != nil {
			return nil, err
		}
		return eo, nil
	}
	return nil, errors.Wrap(errors.ErrBadRequest)

}

func UoToDto(o entity.Order) *Order {
	id, _ := primitive.ObjectIDFromHex(o.ID().Id)
	raw, _ := bson.Marshal(o)
	return &Order{
		Id:          id,
		Type:        o.Type(),
		ExchangeNID: o.ExchangeNid(),
		Status:      o.STATUS().String(),
		Order:       raw,
	}
}
