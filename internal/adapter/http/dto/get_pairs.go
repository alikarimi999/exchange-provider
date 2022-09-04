package dto

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"strings"
)

type GetAllPairsRequest struct {
	Es []string `json:"exchanges"`
}

type Exchange struct {
	Status string       `json:"status"`
	Pairs  []*AdminPair `json:"pairs"`
}

type GetAllPairsResponse struct {
	Exchanges map[string]*Exchange `json:"exchanges"`
	Messages  []string             `json:"messages"`
}

func ParseCoin(coin string) (*entity.Coin, error) {
	parts := strings.Split(coin, "-")
	if len(parts) != 2 {
		return nil, errors.Wrap(errors.ErrBadRequest,
			errors.NewMesssage("coin must be in format: <coin_id>-<chain_id>"))
	}

	return &entity.Coin{
		CoinId:  strings.ToUpper(parts[0]),
		ChainId: strings.ToUpper(parts[1]),
	}, nil
}

type UserPair struct {
	BC                  string  `json:"base_coin"`
	QC                  string  `json:"quote_coin"`
	BuyPrice            string  `json:"buy_price,omitempty"`
	SellPrice           string  `json:"sell_price,omitempty"`
	FeeRate             string  `json:"fee_rate,omitempty"`
	BuyTransferFee      string  `json:"buy_transfer_fee,omitempty"`
	SellTransferFee     string  `json:"sell_transfer_fee,omitempty"`
	MinBaseCoinDeposit  float64 `json:"min_base_coin_deposit,omitempty"`
	MinQuoteCoinDeposit float64 `json:"min_quote_coin_deposit,omitempty"`
	Msg                 string  `json:"message,omitempty"`
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
	Pairs []*UserPair `json:"pairs"`
}

type Pair struct {
	BC *entity.Coin
	QC *entity.Coin
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.BC.String(), p.QC.String())
}

type GetPairsToUserRequest struct {
	Pairs []*UserPair `json:"pairs"`
}

func (r *GetPairsToUserRequest) Parse() ([]*Pair, error) {
	pairs := []*Pair{}
	for _, p := range r.Pairs {
		bc, err := ParseCoin(p.BC)
		if err != nil {
			return nil, err
		}
		qc, err := ParseCoin(p.QC)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, &Pair{
			BC: bc,
			QC: qc,
		})
	}
	return pairs, nil
}
