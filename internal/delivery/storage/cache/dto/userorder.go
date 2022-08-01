package dto

import (
	"encoding/json"
	"order_service/internal/entity"
)

type UserOrder struct {
	Id         int64
	UserId     int64
	CreatedAt  int64
	Status     string
	Deposite   *deposite
	Exchange   string
	Withdrawal *Withdrawal
	BC         string
	BChain     string
	QC         string
	QChain     string

	Side          string
	ExchangeOrder *exchangeOrder
	Broken        bool
	BreakeReason  string
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
		BC:            u.BC.CoinId,
		BChain:        u.BC.ChainId,
		QC:            u.QC.CoinId,
		QChain:        u.QC.ChainId,
		Side:          u.Side,
		ExchangeOrder: eoToDto(u.ExchangeOrder),
		Broken:        u.Broken,
		BreakeReason:  u.BreakReason,
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
		BC: &entity.Coin{
			CoinId:  u.BC,
			ChainId: u.BChain,
		},

		QC: &entity.Coin{
			CoinId:  u.QC,
			ChainId: u.QChain,
		},
		Side:          u.Side,
		ExchangeOrder: u.ExchangeOrder.ToEntity(),
		Broken:        u.Broken,
		BreakReason:   u.BreakeReason,
	}
}

func (u *UserOrder) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
