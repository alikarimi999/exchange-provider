package dto

import (
	"encoding/json"
	"order_service/internal/entity"
)

type deposite struct {
	Id         int64
	Exchange   string
	Volume     string
	Fullfilled bool
	Address    string
}

func dToDto(d *entity.Deposit) *deposite {
	return &deposite{
		Id:         d.Id,
		Exchange:   d.Exchange,
		Volume:     d.Volume,
		Fullfilled: d.Fullfilled,
		Address:    d.Address,
	}
}

func (d *deposite) ToEntity() *entity.Deposit {
	return &entity.Deposit{
		Id:         d.Id,
		Exchange:   d.Exchange,
		Volume:     d.Volume,
		Fullfilled: d.Fullfilled,
		Address:    d.Address,
	}
}

func (d *deposite) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}
