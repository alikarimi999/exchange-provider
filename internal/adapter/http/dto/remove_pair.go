package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

type RemovePairRequest struct {
	Exchange string `json:"exchange"`
	T1       string `json:"t1"` // combined with chain id  ex: BTC-BTC
	T2       string `json:"t2"` // combined with chain id  ex: USDT-TRC20
	Force    bool   `json:"force"`
	Msg      string `json:"message"`
}

func (r *RemovePairRequest) Parse() (t1, t2 *entity.Token, err error) {

	if r.T1 == "" || r.T2 == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("t1 and quote t2 must be set"))
	}

	if r.Exchange == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("exchanges must be set"))
	}

	t1, err = ParseToken(r.T1)
	if err != nil {
		return nil, nil, err
	}

	t2, err = ParseToken(r.T2)
	if err != nil {
		return nil, nil, err
	}

	return t1, t2, nil
}
