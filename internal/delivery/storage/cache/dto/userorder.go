package dto

import (
	"encoding/json"
	"order_service/internal/entity"
)

type UserOrder struct {
	Id         int64
	UserId     int64
	Seq        int64
	CreatedAt  int64
	Status     string
	Deposite   *deposite
	Exchange   string
	Withdrawal *Withdrawal
	BC         string
	BChain     string
	QC         string
	QChain     string

	Side string

	SpreadRate string
	SpreadVol  string

	ExchangeOrder *exchangeOrder

	FaileCode  int64
	FailedDesc string
	entity.MetaData
}

func ToDTO(u *entity.UserOrder) *UserOrder {
	return &UserOrder{
		Id:         u.Id,
		UserId:     u.UserId,
		Seq:        u.Seq,
		CreatedAt:  u.CreatedAt,
		Status:     string(u.Status),
		Deposite:   DToDto(u.Deposit),
		Exchange:   u.Exchange,
		Withdrawal: OWToDTO(u.Withdrawal),
		BC:         u.BC.CoinId,
		BChain:     u.BC.ChainId,
		QC:         u.QC.CoinId,
		QChain:     u.QC.ChainId,
		Side:       u.Side,

		SpreadRate: u.SpreadRate,
		SpreadVol:  u.SpreadVol,

		ExchangeOrder: eoToDto(u.ExchangeOrder),

		FaileCode:  u.FailedCode,
		FailedDesc: u.FailedDesc,
		MetaData:   u.MetaData,
	}
}

func (u *UserOrder) ToEntity() *entity.UserOrder {
	if u.MetaData == nil {
		u.MetaData = make(entity.MetaData)
	}
	return &entity.UserOrder{
		Id:         u.Id,
		UserId:     u.UserId,
		Seq:        u.Seq,
		CreatedAt:  u.CreatedAt,
		Status:     entity.OrderStatus(u.Status),
		Deposit:    u.Deposite.ToEntity(),
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

		Side: u.Side,

		SpreadRate: u.SpreadRate,
		SpreadVol:  u.SpreadVol,

		ExchangeOrder: u.ExchangeOrder.ToEntity(),

		FailedCode: u.FaileCode,
		FailedDesc: u.FailedDesc,
		MetaData:   u.MetaData,
	}
}

func (u *UserOrder) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
