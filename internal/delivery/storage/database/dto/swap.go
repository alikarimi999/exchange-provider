package dto

import (
	"exchange-provider/internal/entity"
	"fmt"
	"strings"
)

type Swap struct {
	Id      uint64 `gorm:"primary_key"`
	OrderId int64
	Status  string // succed, failed

	Index int
	TxId  string

	In        string
	Out       string
	InAmount  string
	OutAmount string

	Exchange string

	Fee         string
	FeeCurrency string

	FailedDesc string
	MetaData   jsonb
}

func swapToDto(s *entity.Swap, r *entity.Route, index int) *Swap {
	if s == nil {
		return &Swap{Index: index}
	}
	if r == nil {
		r = &entity.Route{}
	}

	return &Swap{
		Id:   uint64(s.Id),
		TxId: s.TxId,

		Index: index,

		In:  r.In.String(),
		Out: r.Out.String(),

		Exchange: r.Exchange,

		OrderId: s.OrderId,
		Status:  string(s.Status),

		InAmount:  s.InAmount,
		OutAmount: s.OutAmount,

		Fee:         s.Fee,
		FeeCurrency: s.FeeCurrency,

		FailedDesc: s.FailedDesc,
		MetaData:   jsonb(s.MetaData),
	}
}

func (s *Swap) ToEntity() (*entity.Swap, *entity.Route, int, error) {
	if s.MetaData == nil {
		s.MetaData = make(jsonb)
	}

	swap := &entity.Swap{
		Id:   s.Id,
		TxId: s.TxId,

		OrderId: s.OrderId,
		Status:  s.Status,

		InAmount:  s.InAmount,
		OutAmount: s.OutAmount,

		Fee:         s.Fee,
		FeeCurrency: s.FeeCurrency,

		FailedDesc: s.FailedDesc,
		MetaData:   entity.MetaData(s.MetaData),
	}

	in1, in2, err := parseToken(s.In)
	if err != nil {
		return nil, nil, 0, err
	}
	out1, out2, err := parseToken(s.Out)
	if err != nil {
		return nil, nil, 0, err
	}

	route := &entity.Route{
		In:       &entity.Token{TokenId: in1, ChainId: in2},
		Out:      &entity.Token{TokenId: out1, ChainId: out2},
		Exchange: s.Exchange,
	}

	return swap, route, s.Index, nil

}

func parseToken(t string) (string, string, error) {
	if t == "" {
		return "", "", nil
	}
	ss := strings.Split(t, "-")
	if len(ss) != 2 {
		return "", "", fmt.Errorf("corrupted token string %s", t)
	}
	return ss[0], ss[1], nil
}
