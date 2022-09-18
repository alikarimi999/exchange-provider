package dto

import (
	"fmt"
	"exchange-provider/pkg/errors"
)

type AddPairsRequest struct {
	Pairs []*Pair `json:"pairs"`
}

func (req AddPairsRequest) Validate() error {
	for _, p := range req.Pairs {
		if p.BaseToken == "" || p.Quote_Token == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("Invalid pair token"))
		}
	}
	return nil
}

type Pair struct {
	BaseToken   string `json:"base_token"`
	Quote_Token string `json:"quote_token"`
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.BaseToken, p.Quote_Token)
}
