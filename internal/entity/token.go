package entity

type ExchangeToken interface {
	Snapshot() ExchangeToken
}

type Token struct {
	Symbol   string
	Standard string
	Network  string

	ContractAddress string  `bson:"-"`
	Decimals        uint64  `bson:"-"`
	HasExtraId      bool    `bson:"-"`
	Native          bool    `bson:"-"`
	Min             float64 `bson:"-"`
	Max             float64 `bson:"-"`

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

		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		HasExtraId:      t.HasExtraId,
		Native:          t.Native,
		Min:             t.Min,
		Max:             t.Max,

		ET: et,
	}
}
