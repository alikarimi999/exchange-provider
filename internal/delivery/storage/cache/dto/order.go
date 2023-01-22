package dto

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

type Order struct {
	Type  entity.OrderType
	Order json.RawMessage
}

func ToDto(o entity.Order) (*Order, error) {
	b, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return &Order{
		Type:  o.Type(),
		Order: b,
	}, nil
}

func (o *Order) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}

func (o *Order) ToEntity() (entity.Order, error) {
	switch o.Type {
	case entity.CEXOrder:
		co := &entity.CexOrder{}
		if err := json.Unmarshal(o.Order, co); err != nil {
			return nil, err
		}
		return co, nil
	case entity.EVMOrder:
		eo := &entity.EvmOrder{}
		if err := json.Unmarshal(o.Order, eo); err != nil {
			return nil, err
		}
		return eo, nil
	default:
		return nil, errors.Wrap(errors.ErrBadRequest)
	}
}
