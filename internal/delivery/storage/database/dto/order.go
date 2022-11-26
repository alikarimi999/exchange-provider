package dto

import (
	"exchange-provider/internal/entity"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId int64
	Status string

	Deposit    *Deposit    `gorm:"foreignKey:OrderId"`
	Swaps      []*Swap     `gorm:"foreignKey:OrderId"`
	Withdrawal *Withdrawal `gorm:"foreignKey:OrderId"`

	SpreadRate string
	SpreadVol  string

	FailedCode int64
	FailedDesc string
	MetaData   jsonb
}

func (o *Order) ToEntity() *entity.Order {
	if o.MetaData == nil {
		o.MetaData = make(jsonb)
	}
	if o.Deposit == nil {
		o.Deposit = new(Deposit)
	}
	if o.Withdrawal == nil {
		o.Withdrawal = new(Withdrawal)
	}

	order := &entity.Order{
		Id:        int64(o.ID),
		CreatedAt: o.CreatedAt.Unix(),
		UserId:    o.UserId,
		Status:    entity.OrderStatus(o.Status),

		Routes: make(map[int]*entity.Route),

		Deposit:    o.Deposit.ToEntity(),
		Swaps:      make(map[int]*entity.Swap),
		Withdrawal: o.Withdrawal.ToEntity(),

		SpreadRate: o.SpreadRate,
		SpreadVol:  o.SpreadVol,

		FailedCode: o.FailedCode,
		FailedDesc: o.FailedDesc,
		MetaData:   entity.MetaData(o.MetaData),
	}
	for _, swap := range o.Swaps {
		s, r, i := swap.ToEntity()
		order.Swaps[i] = s
		order.Routes[i] = r
	}
	return order
}

func UoToDto(o *entity.Order) *Order {
	if o == nil {
		return &Order{}
	}

	order := &Order{
		Model: gorm.Model{
			ID:        uint(o.Id),
			CreatedAt: time.Unix(o.CreatedAt, 0),
		},
		UserId:  o.UserId,
		Status:  string(o.Status),
		Deposit: DToDto(o.Deposit),

		Withdrawal: WToDto(o.Withdrawal),

		SpreadRate: o.SpreadRate,
		SpreadVol:  o.SpreadVol,

		FailedCode: o.FailedCode,
		FailedDesc: o.FailedDesc,
		MetaData:   jsonb(o.MetaData),
	}
	for i, s := range o.Swaps {
		order.Swaps = append(order.Swaps, SwapToDto(s, o.Routes[i], i))
	}
	return order
}
