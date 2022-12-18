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

	Fee            string
	FeeCurrency    string
	SpreadRate     string
	SpreadVol      string
	SpreadCurrency string

	FailedCode int64
	FailedDesc string
	MetaData   jsonb
}

func (o *Order) ToEntity() (*entity.Order, error) {
	if o.MetaData == nil {
		o.MetaData = make(jsonb)
	}
	if o.Deposit == nil {
		o.Deposit = new(Deposit)
	}

	if o.Withdrawal == nil {
		o.Withdrawal = new(Withdrawal)
	}

	d, err := o.Deposit.ToEntity()
	if err != nil {
		return nil, err
	}

	w, err := o.Withdrawal.ToEntity()
	if err != nil {
		return nil, err
	}

	order := &entity.Order{
		Id:        int64(o.ID),
		CreatedAt: o.CreatedAt.Unix(),
		UserId:    o.UserId,
		Status:    entity.OrderStatus(o.Status),

		Routes:     make(map[int]*entity.Route),
		Deposit:    d,
		Swaps:      make(map[int]*entity.Swap),
		Withdrawal: w,

		Fee:         o.Fee,
		FeeCurrency: o.FeeCurrency,

		SpreadRate:     o.SpreadRate,
		SpreadVol:      o.SpreadVol,
		SpreadCurrency: o.SpreadCurrency,

		FailedCode: o.FailedCode,
		FailedDesc: o.FailedDesc,
		MetaData:   entity.MetaData(o.MetaData),
	}

	for _, swap := range o.Swaps {
		s, r, i, err := swap.ToEntity()
		if err != nil {
			return nil, err
		}
		order.Swaps[i] = s
		order.Routes[i] = r
	}

	return order, nil
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
		UserId:     o.UserId,
		Status:     string(o.Status),
		Deposit:    DToDto(o.Deposit),
		Withdrawal: WToDto(o.Withdrawal),

		Fee:         o.Fee,
		FeeCurrency: o.FeeCurrency,

		SpreadRate:     o.SpreadRate,
		SpreadVol:      o.SpreadVol,
		SpreadCurrency: o.SpreadCurrency,

		FailedCode: o.FailedCode,
		FailedDesc: o.FailedDesc,
		MetaData:   jsonb(o.MetaData),
	}
	for i, r := range o.Routes {
		order.Swaps = append(order.Swaps, swapToDto(o.Swaps[i], r, i))
	}
	return order
}
