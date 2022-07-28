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

	BC     string
	BChain string

	QC     string
	QChain string

	Side          string
	ExchangeOrder *ExchangeOrder `gorm:"foreignKey:OrderId,UserId"`
}

func UoToDto(uo *entity.UserOrder) *Order {
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
		BC:            uo.BC.CoinId,
		BChain:        uo.BC.ChainId,
		QC:            uo.QC.CoinId,
		QChain:        uo.QC.ChainId,
		Side:          uo.Side,
		ExchangeOrder: EToDto(uo.ExchangeOrder),
	}
}

func (o *Order) ToEntity() *entity.UserOrder {
	return &entity.UserOrder{
		Id:            int64(o.ID),
		CreatedAt:     o.CreatedAt.Unix(),
		UserId:        o.UserId,
		Status:        entity.OrderStatus(o.Status),
		Deposite:      o.Deposite.ToEntity(),
		Exchange:      o.Exchange,
		Withdrawal:    o.Withdrawal.ToEntity(),
		BC:            &entity.Coin{CoinId: o.BC, ChainId: o.BChain},
		QC:            &entity.Coin{CoinId: o.QC, ChainId: o.QChain},
		Side:          o.Side,
		ExchangeOrder: o.ExchangeOrder.ToEntity(),
	}
}
