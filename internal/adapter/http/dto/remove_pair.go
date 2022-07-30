package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

type RemovePairRequest struct {
	Exchanges []string `json:"exchanges"`
	BC        string   `json:"base_coin"`  // combined with chain id  ex: BTC-BTC
	QC        string   `json:"quote_coin"` // combined with chain id  ex: USDT-TRC20
}

func (r *RemovePairRequest) Parse() (bc, qc *entity.Coin, err error) {

	if r.BC == "" || r.QC == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("base coin and quote coin must be set"))
	}

	if len(r.Exchanges) == 0 {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("at least one exchange must be set"))
	}

	if len(r.Exchanges) == 1 && r.Exchanges[0] == "*" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("* is not allowed in remove pair request"))
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

type RemovePairResponse struct {
	Exchanges map[string]string `json:"exchanges"`
}
