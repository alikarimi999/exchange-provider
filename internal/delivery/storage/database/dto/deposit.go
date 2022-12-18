package dto

import "exchange-provider/internal/entity"

type Deposit struct {
	Id      int64
	OrderId int64

	Status string
	Token  string

	TxId   string
	Volume string

	Address string
	Tag     string

	FailedDesc string
}

func DToDto(d *entity.Deposit) *Deposit {
	if d == nil {
		return &Deposit{}
	}

	return &Deposit{
		Id:      d.Id,
		OrderId: d.OrderId,

		Status: d.Status,
		Token:  d.Token.String(),

		TxId:   d.TxId,
		Volume: d.Volume,

		Address: d.Addr,
		Tag:     d.Tag,

		FailedDesc: d.FailedDesc,
	}
}

func (d *Deposit) ToEntity() (*entity.Deposit, error) {
	ed := &entity.Deposit{
		Id:      d.Id,
		OrderId: d.OrderId,

		Status: d.Status,
		TxId:   d.TxId,
		Volume: d.Volume,

		Address: &entity.Address{Addr: d.Address, Tag: d.Tag},

		FailedDesc: d.FailedDesc,
	}

	t, c, err := parseToken(d.Token)
	if err != nil {
		return nil, err
	}
	ed.Token = &entity.Token{TokenId: t, ChainId: c}
	return ed, nil
}
