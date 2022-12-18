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

func ParseToken(t string) (*entity.Token, error) {
	parts := strings.Split(t, "-")
	if len(parts) != 2 {
		return nil, errors.Wrap(errors.ErrBadRequest,
			errors.NewMesssage("token must be in format: <tokenId>-<chainId>"))
	}

	return &entity.Token{
		TokenId: strings.ToUpper(parts[0]),
		ChainId: strings.ToUpper(parts[1]),
	}, nil
}

type UserPair struct {
	T1               string  `json:"t1"`
	T2               string  `json:"t2"`
	Price1           string  `json:"price1,omitempty"`
	Price2           string  `json:"price2,omitempty"`
	FeeRate          string  `json:"fee_rate,omitempty"`
	TransferFee1     string  `json:"transfer_fee1,omitempty"`
	TransferFee2     string  `json:"transfer_fee2,omitempty"`
	MinDepositToken1 float64 `json:"min_deposit_token1,omitempty"`
	MinDepositToken2 float64 `json:"min_deposit_token2,omitempty"`
	Msg              string  `json:"message,omitempty"`
}

func EntityPairToUserRequest(p *entity.Pair, exTyp entity.ExType) *UserPair {
	pair := &UserPair{
		T1:      p.T1.String(),
		T2:      p.T2.String(),
		Price1:  p.Price1,
		Price2:  p.Price2,
		FeeRate: p.FeeRate,
	}
	if exTyp == entity.CEX {
		pair.TransferFee1 = fmt.Sprintf("%s/%s", p.T2.WithdrawalMinFee, p.T2.String())
		pair.TransferFee2 = fmt.Sprintf("%s/%s", p.T1.WithdrawalMinFee, p.T1.String())
	}
	return pair
}

type GetPairsToUserResponse struct {
	Pairs []*UserPair `json:"pairs"`
}

type Pair struct {
	T1 *entity.Token
	T2 *entity.Token
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.T1.String(), p.T2.String())
}

type GetPairsToUserRequest struct {
	Pairs []*UserPair `json:"pairs"`
}

func (r *GetPairsToUserRequest) Parse() ([]*Pair, error) {
	pairs := []*Pair{}
	for _, p := range r.Pairs {
		bc, err := ParseToken(p.T1)
		if err != nil {
			return nil, err
		}
		qc, err := ParseToken(p.T2)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, &Pair{
			T1: bc,
			T2: qc,
		})
	}
	return pairs, nil
}
