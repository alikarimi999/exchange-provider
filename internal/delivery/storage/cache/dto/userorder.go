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
	RequestChain  string
	ProvideCoin   string
	ProvideChain  string
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
		RequestCoin:   u.RequestCoin.Id,
		RequestChain:  u.RequestCoin.Chain.Id,
		ProvideCoin:   u.ProvideCoin.Id,
		ProvideChain:  u.ProvideCoin.Chain.Id,
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
		RequestCoin: &entity.Coin{
			Id:    u.RequestCoin,
			Chain: &entity.Chain{Id: u.RequestChain},
		},

		ProvideCoin: &entity.Coin{
			Id:    u.ProvideCoin,
			Chain: &entity.Chain{Id: u.ProvideChain},
		},

		ExchangeOrder: u.ExchangeOrder.ToEntity(),
		Broken:        u.Broken,
		BrokeReason:   u.BrokeReason,
	}
}

func (u *UserOrder) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
