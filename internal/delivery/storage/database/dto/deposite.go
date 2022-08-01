package dto

import "order_service/internal/entity"

type Deposite struct {
	Id         int64
	OrderId    int64
	UserId     int64
	Exchange   string
	Volume     string
	Fullfilled bool
	Address    string
}

func DToDto(d *entity.Deposit) *Deposite {
	if d == nil {
		return &Deposite{}
	}

	return &Deposite{
		Id:         d.Id,
		UserId:     d.UserId,
		OrderId:    d.OrderId,
		Exchange:   d.Exchange,
		Volume:     d.Volume,
		Fullfilled: d.Fullfilled,
		Address:    d.Address,
	}
}

func (d *Deposite) ToEntity() *entity.Deposit {
	if d == nil {
		return &entity.Deposit{}
	}

	return &entity.Deposit{
		Id:         d.Id,
		UserId:     d.UserId,
		OrderId:    d.OrderId,
		Exchange:   d.Exchange,
		Volume:     d.Volume,
		Fullfilled: d.Fullfilled,
		Address:    d.Address,
	}
}
