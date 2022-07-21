package dto

import "order_service/internal/entity"

type Deposite struct {
	Id         int64
	UserId     int64
	OrderId    int64
	Exchange   string
	Volume     string
	Fullfilled bool
	Address    string
}

func DFromEntity(d *entity.Deposite) *Deposite {
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
