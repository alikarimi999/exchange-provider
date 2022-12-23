package dto

import (
	kdto "exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/internal/entity"
	"math/big"

	"exchange-provider/pkg/errors"
)

type Token struct {
	TokenId             string  `json:"tokenId"`
	ChainId             string  `json:"chainId"`
	ContractAddress     string  `json:"contract_address,omitempty"`
	Address             string  `json:"address,omitempty"`
	Tag                 string  `json:"tag,omitempty"`
	BlockTime           string  `json:"block_time,omitempty"`
	MinDeposit          float64 `json:"min_deposit"`
	MinOrderSize        string  `json:"min_order_size,omitempty"`
	MaxOrderSize        string  `json:"max_order_size,omitempty"`
	MinWithdrawalSize   string  `json:"min_withdrawal_size,omitempty"`
	MinWithdrawalFee    string  `json:"min_withdrawal_fee,omitempty"`
	OrderPrecision      int     `json:"order_precision,omitempty"`
	WithdrawalPrecision int     `json:"withdrawal_precision,omitempty"`
}

type kuPair struct {
	T1 *Token `json:"t1"`
	T2 *Token `json:"t2"`
}

func (p *kuPair) Map() *kdto.Pair {

	pair := &kdto.Pair{
		T1: &kdto.Token{
			TokenId:             p.T1.TokenId,
			ChainId:             p.T1.ChainId,
			WithdrawalPrecision: p.T1.WithdrawalPrecision,
		},
		T2: &kdto.Token{
			TokenId:             p.T2.TokenId,
			ChainId:             p.T2.ChainId,
			WithdrawalPrecision: p.T2.WithdrawalPrecision,
		},
	}

	pair.T1.BlockTime, _ = toTime(p.T1.BlockTime)
	pair.T2.BlockTime, _ = toTime(p.T2.BlockTime)

	return pair
}

type KucoinAddPairsRequest struct {
	Pairs []*kuPair `json:"pairs"`
}

// check there wasn't any zero values in the request
// if there was return an error that the value must set
func (r *KucoinAddPairsRequest) Validate() error {
	for _, p := range r.Pairs {
		if p.T1.TokenId == "" || p.T1.ChainId == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("token1 must have id"))
		}
		if p.T2.TokenId == "" || p.T2.ChainId == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("token2 must have id"))
		}

		if p.T1.WithdrawalPrecision == 0 || p.T2.WithdrawalPrecision == 0 {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("withdrawal_precision must be set"))
		}

		if _, err := toTime(p.T1.BlockTime); err != nil {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error()))
		}

		if _, err := toTime(p.T2.BlockTime); err != nil {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error()))
		}

	}
	return nil
}

type AdminPair struct {
	T1 *Token `json:"t1"`
	T2 *Token `json:"t2"`

	ContractAddress      string   `json:"contract_address,omitempty"`
	FeeTier              int64    `json:"fee_tier,omitempty"`
	Liquidity            *big.Int `json:"liquidity,omitempty"`
	Price                string   `json:"price,omitempty"`
	Price1               string   `json:"price1,omitempty"`
	Price2               string   `json:"price2,omitempty"`
	FeeCurrency          string   `json:"fee_currency"`
	ExchangeOrderFeeRate string   `json:"exchange_order_fee_rate,omitempty"`
	SpreadRate           string   `json:"spread_rate"`
}

func PairDTO(p *entity.Pair) *AdminPair {
	ap := &AdminPair{
		T1: &Token{
			TokenId: p.T1.TokenId,
			ChainId: p.T1.ChainId,

			ContractAddress:     p.T1.ContractAddress,
			Address:             p.T1.Address,
			Tag:                 p.T1.Tag,
			MinDeposit:          p.T1.MinDeposit,
			MinOrderSize:        p.T1.MinOrderSize,
			MaxOrderSize:        p.T1.MaxOrderSize,
			MinWithdrawalSize:   p.T1.MinWithdrawalSize,
			MinWithdrawalFee:    p.T1.WithdrawalMinFee,
			OrderPrecision:      p.T1.OrderPrecision,
			WithdrawalPrecision: p.T1.WithdrawalPrecision,
		},
		T2: &Token{
			TokenId: p.T2.TokenId,
			ChainId: p.T2.ChainId,

			ContractAddress:     p.T2.ContractAddress,
			Address:             p.T2.Address,
			Tag:                 p.T2.Tag,
			MinDeposit:          p.T2.MinDeposit,
			MinOrderSize:        p.T2.MinOrderSize,
			MaxOrderSize:        p.T2.MaxOrderSize,
			MinWithdrawalSize:   p.T2.MinWithdrawalSize,
			MinWithdrawalFee:    p.T2.WithdrawalMinFee,
			OrderPrecision:      p.T2.OrderPrecision,
			WithdrawalPrecision: p.T2.WithdrawalPrecision,
		},

		ContractAddress:      p.ContractAddress,
		FeeTier:              p.FeeTier,
		Liquidity:            p.Liquidity,
		Price1:               p.Price1,
		Price2:               p.Price2,
		SpreadRate:           p.SpreadRate,
		FeeCurrency:          p.FeeCurrency,
		ExchangeOrderFeeRate: p.OrderFeeRate,
	}

	if ap.Price1 == ap.Price2 {
		ap.Price = ap.Price1
		ap.Price1 = ""
		ap.Price2 = ""
	}
	return ap
}

type PairsErr struct {
	Pair string `json:"pair"`
	Err  string `json:"error"`
}
type AddPairsResult struct {
	Addedd []string    `json:"added_pairs"`
	Exs    []string    `json:"existed_pairs"`
	Failed []*PairsErr `json:"failed_pairs"`
	Error  string      `json:"error"`
}

func FromEntity(r *entity.AddPairsResult) *AddPairsResult {
	res := &AddPairsResult{
		Addedd: []string{},
		Exs:    []string{},
		Failed: []*PairsErr{},
	}
	for _, p := range r.Added {
		res.Addedd = append(res.Addedd, p.String())
	}
	res.Exs = append(res.Exs, r.Existed...)

	for _, p := range r.Failed {
		res.Failed = append(res.Failed, &PairsErr{
			Pair: p.Pair,
			Err:  p.Err.Error(),
		})
	}
	return res

}
