package dto

import (
	"exchange-provider/internal/entity"
)

type Swap struct {
	Id      uint64 `gorm:"primary_key"`
	UserId  int64
	OrderId int64
	Status  string // succed, failed

	Index int
	ExId  string

	InCoin  string
	InChain string

	OutCoin  string
	OutChain string

	Exchange string

	InAmount  string
	OutAmount string

	Fee         string
	FeeCurrency string

	FailedDesc string
	MetaData   jsonb
}

func SwapToDto(s *entity.Swap, r *entity.Route, index int) *Swap {
	if s == nil {
		return &Swap{}
	}
	if r == nil {
		r = &entity.Route{}
	}

	return &Swap{
		Id:   uint64(s.Id),
		ExId: s.ExId,

		Index: index,

		InCoin:  r.In.CoinId,
		InChain: r.In.ChainId,

		OutCoin:  r.Out.CoinId,
		OutChain: r.Out.ChainId,

		Exchange: r.Exchange,

		UserId:  s.UserId,
		OrderId: s.OrderId,
		Status:  string(s.Status),

		InAmount:  s.InAmount,
		OutAmount: s.OutAmount,

		Fee:         s.Fee,
		FeeCurrency: s.FeeCurrency,

		FailedDesc: s.FailedDesc,
		MetaData:   jsonb(s.MetaData),
	}
}

func (s *Swap) ToEntity() (*entity.Swap, *entity.Route, int) {
	if s.MetaData == nil {
		s.MetaData = make(jsonb)
	}

	swap := &entity.Swap{
		Id:   s.Id,
		ExId: s.ExId,

		UserId:  s.UserId,
		OrderId: s.OrderId,
		Status:  entity.ExOrderStatus(s.Status),

		InAmount:  s.InAmount,
		OutAmount: s.OutAmount,

		Fee:         s.Fee,
		FeeCurrency: s.FeeCurrency,

		FailedDesc: s.FailedDesc,
		MetaData:   entity.MetaData(s.MetaData),
	}

	route := &entity.Route{
		In:       &entity.Coin{CoinId: s.InCoin, ChainId: s.InChain},
		Out:      &entity.Coin{CoinId: s.OutCoin, ChainId: s.OutChain},
		Exchange: s.Exchange,
	}

	return swap, route, s.Index

}
