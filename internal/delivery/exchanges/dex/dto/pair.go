package dto

import (
	"exchange-provider/pkg/errors"
	"fmt"
)

type AddPairsRequest struct {
	Pairs []*Pair `json:"pairs"`
}

func (req AddPairsRequest) Validate() error {
	for _, p := range req.Pairs {
		if p.C1 == "" || p.C2 == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("Invalid pair token"))
		}
	}
	return nil
}

type Pair struct {
	C1 string `json:"coin1"`
	C2 string `json:"coin2"`
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.C1, p.C2)
}
