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
		if p.BT == "" || p.QT == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("Invalid pair token"))
		}
	}
	return nil
}

type Pair struct {
	BT string `json:"base_token"`
	QT string `json:"quote_token"`
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.BT, p.QT)
}
