package dto

import "order_service/internal/entity"

type Deposit struct {
	Id         int64  `json:"deposit_id"`
	UserId     int64  `json:"user_id"`
	OrderId    int64  `json:"order_id"`
	Exchange   string `json:"exchange"`
	Volume     string `json:"volume"`
	Fullfilled bool   `json:"fullfilled"`
	Address    string `json:"address"`
}

func DFromEntity(d *entity.Deposit) *Deposit {
	return &Deposit{
		Id:         d.Id,
		UserId:     d.UserId,
		OrderId:    d.OrderId,
		Exchange:   d.Exchange,
		Volume:     d.Volume,
		Fullfilled: d.Fullfilled,
		Address:    d.Address,
	}
}
