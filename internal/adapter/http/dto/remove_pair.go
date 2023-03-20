package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

type RemovePairRequest struct {
	Exchange string `json:"exchange"`
	T1       Token  `json:"t1"` // combined with chain id  ex: BTC-BTC
	T2       Token  `json:"t2"` // combined with chain id  ex: USDT-TRC20
	Force    bool   `json:"force"`
	Msg      string `json:"message"`
}

func (r *RemovePairRequest) Parse() (t1, t2 *entity.Token, err error) {

	if r.T1.Symbol == "" || r.T1.Standard == "" || r.T1.Network == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("t1 is invalid"))
	}
	if r.T2.Symbol == "" || r.T2.Standard == "" || r.T2.Network == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("t2 is invalid"))
	}

	if r.Exchange == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("exchanges must be set"))
	}

	return r.T1.ToEntity(), r.T2.ToEntity(), nil
}
