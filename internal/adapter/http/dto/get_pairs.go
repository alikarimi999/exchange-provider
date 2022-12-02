package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
)

type GetAllPairsRequest struct {
	Es []string `json:"exchanges"`
}

type Exchange struct {
	Pairs []*AdminPair `json:"pairs"`
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
	Coin1           string  `json:"coin1"`
	Coin2           string  `json:"coin2"`
	BuyPrice        string  `json:"buy_price,omitempty"`
	SellPrice       string  `json:"sell_price,omitempty"`
	FeeRate         string  `json:"fee_rate,omitempty"`
	BuyTransferFee  string  `json:"buy_transfer_fee,omitempty"`
	SellTransferFee string  `json:"sell_transfer_fee,omitempty"`
	MinDepositCoin1 float64 `json:"min_deposit_coin1,omitempty"`
	MinDepositCoin2 float64 `json:"min_deposit_coin2,omitempty"`
	Msg             string  `json:"message,omitempty"`
}

func EntityPairToUserRequest(p *entity.Pair, exTyp entity.ExType) *UserPair {
	pair := &UserPair{
		Coin1:     p.C1.String(),
		Coin2:     p.C2.String(),
		BuyPrice:  p.BestAsk,
		SellPrice: p.BestBid,
		FeeRate:   p.FeeRate,
	}
	if exTyp == entity.CEX {
		pair.BuyTransferFee = fmt.Sprintf("%s/%s", p.C1.WithdrawalMinFee, p.C1.String())
		pair.SellTransferFee = fmt.Sprintf("%s/%s", p.C2.WithdrawalMinFee, p.C2.String())
	}
	return pair
}

type GetPairsToUserResponse struct {
	Pairs []*UserPair `json:"pairs"`
}

type Pair struct {
	Coin1 *entity.Coin
	Coin2 *entity.Coin
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.Coin1.String(), p.Coin2.String())
}

type GetPairsToUserRequest struct {
	Pairs []*UserPair `json:"pairs"`
}

func (r *GetPairsToUserRequest) Parse() ([]*Pair, error) {
	pairs := []*Pair{}
	for _, p := range r.Pairs {
		bc, err := ParseCoin(p.Coin1)
		if err != nil {
			return nil, err
		}
		qc, err := ParseCoin(p.Coin2)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, &Pair{
			Coin1: bc,
			Coin2: qc,
		})
	}
	return pairs, nil
}
