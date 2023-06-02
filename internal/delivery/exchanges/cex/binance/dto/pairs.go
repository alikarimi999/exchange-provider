package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

type AddPairsRequest struct {
	Pairs []*Pair `json:"pairs"`
}
type Token struct {
	Coin        string `json:"coin"`
	Network     string `json:"network"`
	StableToken string `json:"stableToken"`

	BlockTime string `json:"blockTime"`
}

type Pair struct {
	BC          *EToken          `json:"bc"`
	QC          *EToken          `json:"qc"`
	IC          string           `json:"ic"`
	Enable      bool             `json:"enable"`
	Spreads     map[uint]float64 `json:"spreads"`
	FeeRate1    float64          `json:"feeRate1"`
	FeeRate2    float64          `json:"feeRate2"`
	ExchangeFee float64          `json:"exchangeFee"`
}

func (p Pair) String() string {
	return fmt.Sprintf("%s/%s", p.BC.String(), p.QC.String())
}

type EToken struct {
	entity.TokenId
	ContractAddress string  `json:"contractAddress"`
	Decimals        int     `json:"decimals"`
	Native          bool    `json:"native"`
	Min             float64 `json:"min"`
	Max             float64 `json:"max"`
	ET              Token   `json:"exchangeToken"`
}

func (t *EToken) String() string {
	return t.TokenId.String()
}

func (t *EToken) toEntity(fn func(Token, binance.Network) (entity.ExchangeToken, error),
	n binance.Network) (*entity.Token, error) {

	et, err := fn(t.ET, n)
	if err != nil {
		return nil, err
	}

	return &entity.Token{
		Id:              *t.ToUpper(),
		ContractAddress: t.ContractAddress,
		Decimals:        uint64(t.Decimals),
		Native:          t.Native,
		Min:             t.Min,
		Max:             t.Max,
		ET:              et,
	}, nil
}

func (p *Pair) ToEntity(fn func(Token, binance.Network) (entity.ExchangeToken, error), bc, qc binance.Network) (*entity.Pair, error) {
	if p.BC == nil {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("bc is required"))
	}
	if p.QC == nil {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("qc is required"))
	}
	t1, err := p.BC.toEntity(fn, bc)
	if err != nil {
		return nil, err
	}
	t2, err := p.QC.toEntity(fn, qc)
	if err != nil {
		return nil, err
	}

	ep := &entity.Pair{
		T1:          t1,
		T2:          t2,
		Enable:      p.Enable,
		FeeRate1:    p.FeeRate1,
		FeeRate2:    p.FeeRate2,
		ExchangeFee: p.ExchangeFee,
	}
	if p.Spreads == nil || len(p.Spreads) == 0 {
		ep.Spreads = make(map[uint]float64)
	} else {
		ep.Spreads = p.Spreads
	}

	return ep, nil
}
