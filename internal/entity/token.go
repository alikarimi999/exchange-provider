package entity

type ExchangeToken interface {
	Snapshot() ExchangeToken
}

type Token struct {
	Symbol   string
	Standard string
	Network  string

	Address    string  `bson:"address,omitempty"`
	Decimals   uint64  `bson:"decimals,omitempty"`
	HasExtraId bool    `bson:"hasExtraId,omitempty"`
	Native     bool    `bson:"native,omitempty"`
	Min        float64 `bson:"min,omitempty"`
	Max        float64 `bson:"max,omitempty"`

	ET ExchangeToken `bson:"et,omitempty"`
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

		ET: et,
	}
}
