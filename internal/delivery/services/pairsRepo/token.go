package pairsRepo

import (
	"exchange-provider/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type token struct {
	Symbol   string
	Standard string
	Network  string

	StableToken     string  `bson:"stableToken,omitempty"`
	ContractAddress string  `bson:"contractAddress,omitempty"`
	Decimals        uint64  `bson:"decimals,omitempty"`
	Native          bool    `bson:"native,omitempty"`
	Min             float64 `bson:"min,omitempty"`
	Max             float64 `bson:"max,omitempty"`
	ET              bson.Raw
}

func fromEntity(t *entity.Token) *token {
	et, _ := bson.Marshal(t.ET)
	return &token{

		Symbol:          t.Id.Symbol,
		Standard:        t.Id.Standard,
		Network:         t.Id.Network,
		StableToken:     t.StableToken,
		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		Native:          t.Native,
		Min:             t.Min,
		Max:             t.Max,
		ET:              et,
	}
}

func (t token) toEntity(fn func(bson.Raw) entity.ExchangeToken) *entity.Token {
	return &entity.Token{
		Id: entity.TokenId{
			Symbol:   t.Symbol,
			Standard: t.Standard,
			Network:  t.Network,
		},
		StableToken:     t.StableToken,
		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		Native:          t.Native,
		Min:             t.Min,
		Max:             t.Max,
		ET:              fn(t.ET),
	}
}
