package dto

import (
	"order_service/internal/entity"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId     int64 `gorm:"primaryKey"`
	Status     string
	Deposite   *Deposite `gorm:"foreignKey:OrderId,UserId"`
	Exchange   string
	Withdrawal *Withdrawal `gorm:"foreignKey:OrderId,UserId"`

	BaseCoin  string
	BaseChain string

	QuoteCoin  string
	QuoteChain string

	Side          string
	ExchangeOrder *ExchangeOrder `gorm:"foreignKey:OrderId,UserId"`

	Broken      bool
	BreakReason string
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
		UserId:        uo.UserId,
		Status:        string(uo.Status),
		Deposite:      DToDto(uo.Deposite),
		Exchange:      uo.Exchange,
		Withdrawal:    WToDto(uo.Withdrawal),
		BaseCoin:      uo.BC.CoinId,
		BaseChain:     uo.BC.ChainId,
		QuoteCoin:     uo.QC.CoinId,
		QuoteChain:    uo.QC.ChainId,
		Side:          uo.Side,
		ExchangeOrder: EToDto(uo.ExchangeOrder),
		Broken:        uo.Broken,
		BreakReason:   uo.BreakReason,
	}
}

func (o *Order) ToEntity() *entity.UserOrder {

	if o == nil {
		return &entity.UserOrder{}
	}

	return &entity.UserOrder{
		Id:            int64(o.ID),
		CreatedAt:     o.CreatedAt.Unix(),
		UserId:        o.UserId,
		Status:        entity.OrderStatus(o.Status),
		Deposite:      o.Deposite.ToEntity(),
		Exchange:      o.Exchange,
		Withdrawal:    o.Withdrawal.ToEntity(),
		BC:            &entity.Coin{CoinId: o.BaseCoin, ChainId: o.BaseChain},
		QC:            &entity.Coin{CoinId: o.QuoteCoin, ChainId: o.QuoteChain},
		Side:          o.Side,
		ExchangeOrder: o.ExchangeOrder.ToEntity(),
		Broken:        o.Broken,
		BreakReason:   o.BreakReason,
	}
}
