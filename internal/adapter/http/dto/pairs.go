package dto

import (
	"exchange-provider/internal/entity"
)

const (
	EnableCmd  string = "enable"
	DisableCmd string = "disable"
	RemoveCmd  string = "remove"
)

type Token struct {
	entity.TokenId
	StableToken string `json:"stableToken"`
	Min         Number `json:"min"`
	Max         Number `json:"max"`
}

func TokenFromEntity(et *entity.Token) Token {
	return Token{
		TokenId: entity.TokenId{
			Symbol:   et.Id.Symbol,
			Standard: et.Id.Standard,
			Network:  et.Id.Network,
		},
		Min: Number(et.Min),
		Max: Number(et.Max),
	}
}

type Pair struct {
	T1          Token            `json:"t1"`
	T2          Token            `json:"t2"`
	Enable      bool             `json:"enable"`
	FeeRate1    float64          `json:"feeRate1"`
	FeeRate2    float64          `json:"feeRate2"`
	ExchangeFee float64          `json:"exchangeFee"`
	Spreads     map[uint]float64 `json:"spreads,omitempty"`
	LP          uint             `json:"lp"`
}

func (p *Pair) Update(ep *entity.Pair, acceptZero bool) {
	if acceptZero {
		ep.T1.Min = float64(p.T1.Min)
		ep.T1.Max = float64(p.T1.Max)
		ep.T2.Min = float64(p.T2.Min)
		ep.T2.Max = float64(p.T2.Max)
		ep.FeeRate1 = p.FeeRate1
		ep.FeeRate2 = p.FeeRate2
		ep.ExchangeFee = p.ExchangeFee
	} else {
		if p.T1.Min > 0 {
			ep.T1.Min = float64(p.T1.Min)
		}
		if p.T1.Max > 0 {
			ep.T1.Max = float64(p.T1.Max)
		}
		if p.T2.Min > 0 {
			ep.T2.Min = float64(p.T2.Min)
		}
		if p.T2.Max > 0 {
			ep.T2.Max = float64(p.T2.Max)
		}
		if p.FeeRate1 > 0 {
			ep.FeeRate1 = p.FeeRate1
		}
		if p.FeeRate2 > 0 {
			ep.FeeRate2 = p.FeeRate2
		}
		if p.ExchangeFee > 0 {
			ep.ExchangeFee = p.ExchangeFee
		}
	}
	if p.T1.StableToken != "" {
		ep.T1.StableToken = p.T1.StableToken
	}

	if p.T2.StableToken != "" {
		ep.T2.StableToken = p.T2.StableToken
	}

	for k, v := range p.Spreads {
		ep.Spreads[k] = v
	}
}

type Pairs []struct {
	T1 entity.TokenId `json:"t1"`
	T2 entity.TokenId `json:"t2"`
	LP uint           `json:"lp"`
}

type UpdatePairReq struct {
	AcceptZero bool   `json:"acceptZero"`
	Pairs      []Pair `json:"pairs"`
}

type PairsRes []struct {
	Pair string `json:"pair"`
	Msg  string `json:"msg"`
}
type LpsRes []struct {
	Lp  uint   `json:"lp"`
	Msg string `json:"msg"`
}
type PairsRequest struct {
	Cmd   string `json:"cmd"`
	All   bool   `json:"all"`
	Pairs `json:"pairs"`
}

type LpsRequest struct {
	Cmd string `json:"cmd"`
	All bool   `json:"all"`
	Lps []uint `json:"lps"`
}

type CmdResp struct {
	PairsRes `json:"pairs,omitempty"`
	LpsRes   `json:"lps,omitempty"`
}
