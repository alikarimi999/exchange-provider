package dto

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
	"fmt"
)

type AddPairsRequest struct {
	Pairs []*Pair `json:"pairs"`
}

type Pair struct {
	T1          *EToken `json:"t1"`
	T2          *EToken `json:"t2"`
	Enable      bool    `json:"enable"`
	FeeRate1    float64 `json:"feeRate1"`
	FeeRate2    float64 `json:"feeRate2"`
	ExchangeFee float64 `json:"exchangeFee"`
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.T1.String(), p.T2.String())
}

func (t *EToken) toEntity() (*entity.Token, error) {
	if err := t.Token.Check(); err != nil {
		return nil, err
	}

	return &entity.Token{
		Id:              t.TokenId,
		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		Native:          t.Native,
		Min:             t.Min,
		Max:             t.Max,
		ET:              &types.EToken{},
	}, nil
}

func (p *Pair) ToEntity() (*entity.Pair, error) {
	t1, err := p.T1.toEntity()
	if err != nil {
		return nil, err
	}

	t2, err := p.T2.toEntity()
	if err != nil {
		return nil, err
	}
	return &entity.Pair{
		T1:          t1,
		T2:          t2,
		Enable:      p.Enable,
		FeeRate1:    p.FeeRate1,
		FeeRate2:    p.FeeRate2,
		ExchangeFee: p.ExchangeFee,
		EP:          &types.ExchangePair{},
	}, nil
}
