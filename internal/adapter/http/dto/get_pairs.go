package dto

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"strings"
)

type GetAllPairsRequest struct {
	Names []string `json:"exchange_names"`
}

type Exchange struct {
	Status string       `json:"status"`
	Pairs  []*AdminPair `json:"pairs"`
}

type GetAllPairsResponse struct {
	Exchanges map[string]*Exchange `json:"exchanges"`
	Messages  []string             `json:"messages"`
}

type GetPairsToUserRequest struct {
	BC string `json:"base_coin"`  // combined with chain id  ex: BTC-BTC
	QC string `json:"quote_coin"` // combined with chain id  ex: USDT-TRC20
}

func (r *GetPairsToUserRequest) Parse() (bc, qc *entity.Coin, err error) {

	if r.BC == "" || r.QC == "" {
		return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("base coin and quote coin must be set"))
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

type UserPair struct {
	BC                  string  `json:"base_coin"`
	QC                  string  `json:"quote_coin"`
	BuyPrice            string  `json:"buy_price"`
	SellPrice           string  `json:"sell_price"`
	FeeRate             string  `json:"fee_rate"`
	BuyTransferFee      string  `json:"buy_transfer_fee"`
	SellTransferFee     string  `json:"sell_transfer_fee"`
	MinBaseCoinDeposit  float64 `json:"min_base_coin_deposit"`
	MinQuoteCoinDeposit float64 `json:"min_quote_coin_deposit"`
}

func EntityPairToUserRequest(p *entity.Pair) *UserPair {
	return &UserPair{
		BC:              p.BC.String(),
		QC:              p.QC.String(),
		BuyPrice:        p.BestAsk,
		SellPrice:       p.BestBid,
		FeeRate:         p.FeeRate,
		BuyTransferFee:  fmt.Sprintf("%s/%s", p.BC.WithdrawalMinFee, p.BC.String()),
		SellTransferFee: fmt.Sprintf("%s/%s", p.QC.WithdrawalMinFee, p.QC.String()),
	}
}

type GetPairsToUserResponse struct {
	Pairs []*AdminPair `json:"pairs"`
}
