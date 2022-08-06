package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"strings"
)

type GetAllPairsRequest struct {
	Exchanges []string `json:"exchange_names"`
}

type Exchange struct {
	Status string  `json:"status"`
	Pairs  []*Pair `json:"pairs"`
}

type GetAllPairsResponse struct {
	Exchanges map[string]*Exchange `json:"exchanges"`
	Messages  []string             `json:"messages"`
}

type GetPairRequest struct {
	Exchanges []string `json:"exchanges"`
	BC        string   `json:"base_coin"`  // combined with chain id  ex: BTC-BTC
	QC        string   `json:"quote_coin"` // combined with chain id  ex: USDT-TRC20
}

func (r *GetPairRequest) Parse() (bc, qc *entity.Coin, err error) {

	if r.BC == "" || r.QC == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("base coin and quote coin must be set"))
	}

	if len(r.Exchanges) == 0 {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("at least one exchange must be set"))
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

func ParseCoin(coin string) (*entity.Coin, error) {
	parts := strings.Split(coin, "-")
	if len(parts) != 2 {
		return nil, errors.Wrap(errors.ErrBadRequest,
			errors.NewMesssage("coin must be in format: <coin_id>-<chain_id>"))
	}

	return &entity.Coin{
		CoinId:  parts[0],
		ChainId: parts[1],
	}, nil
}

type GetPairResponse struct {
	Exchanges map[string]*Pair `json:"exchanges"`
	Messages  []string         `json:"messages"`
}
