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

func dToDto(d *entity.Deposite) *deposite {
	return &deposite{
		Id:         d.Id,
		Exchange:   d.Exchange,
		Volume:     d.Volume,
		Fullfilled: d.Fullfilled,
		Address:    d.Address,
	}
}

func (d *deposite) ToEntity() *entity.Deposite {
	return &entity.Deposite{
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
