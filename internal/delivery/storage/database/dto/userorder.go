package dto

import (
	"order_service/internal/entity"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId     int64
	Seq        int64
	Status     string
	Deposit    *Deposit `gorm:"foreignKey:OrderId"`
	Exchange   string
	Withdrawal *Withdrawal `gorm:"foreignKey:OrderId"`

	BaseCoin  string
	BaseChain string

	QuoteCoin  string
	QuoteChain string

	Side string

	SpreadRate    string
	SpreadVol     string
	ExchangeOrder *ExchangeOrder `gorm:"foreignKey:OrderId"`

	FailedCode int64
	FailedDesc string
}

func UoToDto(uo *entity.UserOrder) *Order {
	if uo == nil {
		return &Order{}
	}

	return &Order{
		Model: gorm.Model{
			ID:        uint(uo.Id),
			CreatedAt: time.Unix(uo.CreatedAt, 0),
		},
		UserId:     uo.UserId,
		Seq:        uo.Seq,
		Status:     string(uo.Status),
		Deposit:    DToDto(uo.Deposit),
		Exchange:   uo.Exchange,
		Withdrawal: WToDto(uo.Withdrawal),
		BaseCoin:   uo.BC.CoinId,
		BaseChain:  uo.BC.ChainId,
		QuoteCoin:  uo.QC.CoinId,
		QuoteChain: uo.QC.ChainId,
		Side:       uo.Side,

		SpreadRate:    uo.SpreadRate,
		SpreadVol:     uo.SpreadVol,
		ExchangeOrder: EToDto(uo.ExchangeOrder),

		FailedCode: uo.FailedCode,
		FailedDesc: uo.FailedDesc,
	}
}

func (o *Order) ToEntity() *entity.UserOrder {

	if o == nil {
		return &entity.UserOrder{}
	}

	return &entity.UserOrder{
		Id:         int64(o.ID),
		CreatedAt:  o.CreatedAt.Unix(),
		UserId:     o.UserId,
		Seq:        o.Seq,
		Status:     entity.OrderStatus(o.Status),
		Deposit:    o.Deposit.ToEntity(),
		Exchange:   o.Exchange,
		Withdrawal: o.Withdrawal.ToEntity(),
		BC:         &entity.Coin{CoinId: o.BaseCoin, ChainId: o.BaseChain},
		QC:         &entity.Coin{CoinId: o.QuoteCoin, ChainId: o.QuoteChain},
		Side:       o.Side,

		SpreadRate:    o.SpreadRate,
		SpreadVol:     o.SpreadVol,
		ExchangeOrder: o.ExchangeOrder.ToEntity(),

		FailedCode: o.FailedCode,
		FailedDesc: o.FailedDesc,
	}
}
