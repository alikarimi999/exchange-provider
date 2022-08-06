package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

type RemovePairRequest struct {
	Exchange string `json:"exchange"`
	BC       string `json:"base_coin"`  // combined with chain id  ex: BTC-BTC
	QC       string `json:"quote_coin"` // combined with chain id  ex: USDT-TRC20
	Force    bool   `json:"force"`
}

func (r *RemovePairRequest) Parse() (bc, qc *entity.Coin, err error) {

	if r.BC == "" || r.QC == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("base coin and quote coin must be set"))
	}

	if r.Exchange == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("exchanges must be set"))
	}

	bc, err = ParseCoin(r.BC)
	if err != nil {
		return nil, nil, err
	}

	qc, err = ParseCoin(r.QC)
	if err != nil {
		return nil, nil, err
	}

	return bc, qc, nil
}
