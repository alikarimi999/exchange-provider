package dto

import (
	"exchange-provider/internal/entity"
	"fmt"
)

type AddPairsRequest struct {
	Pairs []*Pair `json:"pairs"`
}
type Token struct {
	Currency  string `json:"currency"`
	ChainName string `json:"chainName"`
	Chain     string

	DepositAddress string `json:"depositAddress"`
	DepositTag     string `json:"depositTag"`

	BlockTime           string `json:"blockTime"`
	WithdrawalPrecision int    `json:"withdrawalPrecision"`
}

type Pair struct {
	BC *EToken `json:"bc"`
	QC *EToken `json:"qc"`
}

func (p Pair) String() string {
	return fmt.Sprintf("%s-%s-%s/%s-%s-%s", p.BC.Symbol, p.BC.Standard, p.BC.Network,
		p.QC.Symbol, p.QC.Standard, p.QC.Network)
}

type EToken struct {
	Symbol   string `json:"symbol"`
	Standard string `json:"standard"`
	Network  string `json:"network"`

	ContractAddress string `json:"contractAddress"`
	Decimals        int    `json:"decimals"`
	Native          bool   `json:"native"`
	ET              Token  `json:"exchangeToken"`
}

func (t *EToken) toEntity(fn func(Token) (entity.ExchangeToken, error)) (*entity.Token, error) {
	et, err := fn(t.ET)
	if err != nil {
		return nil, err
	}

	return &entity.Token{
		Symbol:   t.Symbol,
		Standard: t.Standard,
		Network:  t.Network,

		ContractAddress: t.ContractAddress,
		Decimals:        uint64(t.Decimals),
		Native:          t.Native,
		ET:              et,
	}, nil
}

func (p *Pair) ToEntity(fn func(Token) (entity.ExchangeToken, error)) (*entity.Pair, error) {
	t1, err := p.BC.toEntity(fn)
	if err != nil {
		return nil, err
	}
	t2, err := p.QC.toEntity(fn)
	if err != nil {
		return nil, err
	}
	return &entity.Pair{
		T1: t1,
		T2: t2,
	}, nil
}
