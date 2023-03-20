package entity

import "time"

type ExchangeToken interface {
	Snapshot() ExchangeToken
}

type Token struct {
	Symbol   string
	Standard string
	Network  string

	Address    string  `bson:"-"`
	Decimals   uint64  `bson:"-"`
	HasExtraId bool    `bson:"-"`
	Native     bool    `bson:"-"`
	Min        float64 `bson:"-"`
	Max        float64 `bson:"-"`

	BlockTime           time.Duration `bson:"-"`
	Tag                 string        `bson:"-"`
	MinDeposit          float64       `bson:"-"`
	MinOrderSize        string        `bson:"-"`
	MaxOrderSize        string        `bson:"-"`
	MinWithdrawalSize   string        `bson:"-"`
	WithdrawalMinFee    string        `bson:"-"`
	OrderPrecision      int           `bson:"-"`
	WithdrawalPrecision int           `bson:"-"`
	SetChain            bool          `bson:"-"`

	ET ExchangeToken `bson:"-"`
}

func (t *Token) String() string {
	return t.Symbol + "-" + t.Standard + "-" + t.Network
}

func (t *Token) Equal(t2 *Token) bool {
	return t.String() == t2.String()
}

func (t *Token) Snapshot() *Token {
	var et ExchangeToken
	if t.ET != nil {
		et = t.ET.Snapshot()
	} else {
		et = nil
	}

	return &Token{
		Symbol:   t.Symbol,
		Standard: t.Standard,
		Network:  t.Network,

		Address:    t.Address,
		Decimals:   t.Decimals,
		HasExtraId: t.HasExtraId,
		Native:     t.Native,
		Min:        t.Min,
		Max:        t.Max,

		BlockTime:         t.BlockTime,
		Tag:               t.Tag,
		MinDeposit:        t.MinDeposit,
		MinOrderSize:      t.MinOrderSize,
		MaxOrderSize:      t.MaxOrderSize,
		MinWithdrawalSize: t.MinWithdrawalSize,
		WithdrawalMinFee:  t.WithdrawalMinFee,
		OrderPrecision:    t.OrderPrecision,
		SetChain:          t.SetChain,
		ET:                et,
	}
}
