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
		if p.T1 == "" || p.T2 == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("Invalid pair token"))
		}
	}
	return nil
}

type Pair struct {
	T1 string `json:"t1"`
	T2 string `json:"t2"`
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.T1, p.T2)
}
