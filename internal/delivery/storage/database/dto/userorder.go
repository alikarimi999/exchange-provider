package dto

import (
	"order_service/internal/entity"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId        int64 `gorm:"primaryKey"`
	Status        string
	Deposite      *Deposite `gorm:"foreignKey:OrderId,UserId"`
	Exchange      string
	Withdrawal    *Withdrawal `gorm:"foreignKey:OrderId,UserId"`
	RequestCoin   string
	RequestChain  string
	ProvideCoin   string
	ProvideChain  string
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
		RequestCoin:   uo.RequestCoin.Id,
		RequestChain:  uo.RequestCoin.Chain.Id,
		ProvideCoin:   uo.ProvideCoin.Id,
		ProvideChain:  uo.ProvideCoin.Chain.Id,
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
		RequestCoin:   &entity.Coin{Id: o.RequestCoin, Chain: &entity.Chain{Id: o.RequestChain}},
		ProvideCoin:   &entity.Coin{Id: o.ProvideCoin, Chain: &entity.Chain{Id: o.ProvideChain}},
		ExchangeOrder: o.ExchangeOrder.ToEntity(),
	}
}
