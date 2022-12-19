package dto

import (
	"exchange-provider/internal/entity"
)

type Withdrawal struct {
	Id      uint64 `gorm:"primary_key"`
	OrderId int64

	Status  string
	Address string
	Tag     string

	Token string

	Unwrapped bool

	Volume      string
	Fee         string
	FeeCurrency string

	TxId       string
	FailedDesc string
}

func WToDto(w *entity.Withdrawal) *Withdrawal {

	if w == nil {
		return &Withdrawal{}
	}

	return &Withdrawal{
		Id:      w.Id,
		OrderId: w.OrderId,

		Address: w.Addr,
		Tag:     w.Tag,

		Token:     w.Token.String(),
		Unwrapped: w.Unwrapped,

		Volume:      w.Volume,
		Fee:         w.Fee,
		FeeCurrency: w.FeeCurrency,

		TxId:       w.TxId,
		Status:     string(w.Status),
		FailedDesc: w.FailedDesc,
	}
}

func (w *Withdrawal) ToEntity() (*entity.Withdrawal, error) {
	ew := &entity.Withdrawal{
		Id:      w.Id,
		OrderId: w.OrderId,

		Address:   &entity.Address{Addr: w.Address, Tag: w.Tag},
		Unwrapped: w.Unwrapped,

		Volume:      w.Volume,
		Fee:         w.Fee,
		FeeCurrency: w.FeeCurrency,

		TxId:       w.TxId,
		Status:     w.Status,
		FailedDesc: w.FailedDesc,
	}

	t, c, err := parseToken(w.Token)
	if err != nil {
		return nil, err
	}
	ew.Token = &entity.Token{TokenId: t, ChainId: c}

	return ew, nil
}
