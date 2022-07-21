package dto

import (
	"encoding/json"
	"order_service/internal/entity"
)

type UserOrder struct {
	Id            int64
	UserId        int64
	CreatedAt     int64
	Status        string
	Deposite      *deposite
	Exchange      string
	Withdrawal    *Withdrawal
	RequestCoin   string
	RequestNet    string
	ProvideCoin   string
	ProvideNet    string
	ExchangeOrder *exchangeOrder
	Broken        bool
	BrokeReason   string
}

func ToDTO(u *entity.UserOrder) *UserOrder {
	return &UserOrder{
		Id:            u.Id,
		UserId:        u.UserId,
		CreatedAt:     u.CreatedAt,
		Status:        string(u.Status),
		Deposite:      dToDto(u.Deposite),
		Exchange:      u.Exchange,
		Withdrawal:    OWToDTO(u.Withdrawal),
		RequestCoin:   u.RequestCoin.Symbol,
		RequestNet:    string(u.RequestCoin.Chain),
		ProvideCoin:   u.ProvideCoin.Symbol,
		ProvideNet:    string(u.ProvideCoin.Chain),
		ExchangeOrder: eoToDto(u.ExchangeOrder),
		Broken:        u.Broken,
		BrokeReason:   u.BrokeReason,
	}
}

func (u *UserOrder) ToEntity() *entity.UserOrder {
	return &entity.UserOrder{
		Id:         u.Id,
		UserId:     u.UserId,
		CreatedAt:  u.CreatedAt,
		Status:     entity.OrderStatus(u.Status),
		Deposite:   u.Deposite.ToEntity(),
		Exchange:   u.Exchange,
		Withdrawal: u.Withdrawal.ToEntity(),
		RequestCoin: entity.Coin{
			Symbol: u.RequestCoin,
			Chain:  entity.Chain(u.RequestNet),
		},

		ProvideCoin: entity.Coin{
			Symbol: u.ProvideCoin,
			Chain:  entity.Chain(u.ProvideNet),
		},

		ExchangeOrder: u.ExchangeOrder.ToEntity(),
		Broken:        u.Broken,
		BrokeReason:   u.BrokeReason,
	}
}

func (u *UserOrder) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
