package pairsRepo

import (
	"exchange-provider/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type token struct {
	Symbol   string
	Standard string
	Network  string

	ContractAddress string  `bson:"contractAddress,omitempty"`
	Decimals        uint64  `bson:"decimals,omitempty"`
	HasExtraId      bool    `bson:"hasExtraId,omitempty"`
	Native          bool    `bson:"native,omitempty"`
	Min             float64 `bson:"min,omitempty"`
	Max             float64 `bson:"max,omitempty"`
	ET              bson.Raw
}

func fromEntity(t *entity.Token) *token {
	et, _ := bson.Marshal(t.ET)
	return &token{
		Symbol:          t.Symbol,
		Standard:        t.Standard,
		Network:         t.Network,
		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		HasExtraId:      t.HasExtraId,
		Native:          t.Native,
		Min:             t.Min,
		Max:             t.Max,
		ET:              et,
	}
}

func (t token) toEntity(fn func(bson.Raw) entity.ExchangeToken) *entity.Token {
	return &entity.Token{
		Symbol:          t.Symbol,
		Standard:        t.Standard,
		Network:         t.Network,
		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		HasExtraId:      t.HasExtraId,
		Native:          t.Native,
		Min:             t.Min,
		Max:             t.Max,
		ET:              fn(t.ET),
	}
}
