package dto

import "fmt"

type AddPairsRequest struct {
	Pairs []*Pair `json:"pairs"`
}

type Pair struct {
	BaseToken   string `json:"base_token"`
	Quote_Token string `json:"quote_token"`
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.BaseToken, p.Quote_Token)
}
