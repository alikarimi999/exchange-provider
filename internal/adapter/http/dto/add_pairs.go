package dto

import (
	kdto "exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/internal/entity"
	"math/big"

	"exchange-provider/pkg/errors"
)

type Coin struct {
	CoinId              string  `json:"coin_id"`
	ChainId             string  `json:"chain_id"`
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
	SetChain            bool    `json:"set_chain,omitempty"`
}

type kuPair struct {
	C1 *Coin `json:"coin1"`
	C2 *Coin `json:"coin2"`
}

func (p *kuPair) Map() *kdto.Pair {

	pair := &kdto.Pair{
		C1: &kdto.Coin{
			CoinId:              p.C1.CoinId,
			ChainId:             p.C1.ChainId,
			WithdrawalPrecision: p.C1.WithdrawalPrecision,
		},
		C2: &kdto.Coin{
			CoinId:              p.C2.CoinId,
			ChainId:             p.C2.ChainId,
			WithdrawalPrecision: p.C2.WithdrawalPrecision,
		},
	}

	pair.C1.BlockTime, _ = toTime(p.C1.BlockTime)
	pair.C2.BlockTime, _ = toTime(p.C2.BlockTime)

	return pair
}

type KucoinAddPairsRequest struct {
	Pairs []*kuPair `json:"pairs"`
}

// check there wasn't any zero values in the request
// if there was return an error that the value must set
func (r *KucoinAddPairsRequest) Validate() error {
	for _, p := range r.Pairs {
		if p.C1.CoinId == "" || p.C1.ChainId == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("base coin must have id"))
		}
		if p.C2.CoinId == "" || p.C2.ChainId == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("quote coin must have id"))
		}

		if p.C1.WithdrawalPrecision == 0 || p.C2.WithdrawalPrecision == 0 {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("withdrawal_precision must be set"))
		}

		if _, err := toTime(p.C1.BlockTime); err != nil {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error()))
		}

		if _, err := toTime(p.C2.BlockTime); err != nil {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error()))
		}

	}
	return nil
}

type AdminPair struct {
	C1 *Coin `json:"coin1"`
	C2 *Coin `json:"coin2"`

	ContractAddress      string   `json:"contract_address,omitempty"`
	FeeTier              int64    `json:"fee_tier,omitempty"`
	Liquidity            *big.Int `json:"liquidity,omitempty"`
	Price                string   `json:"price,omitempty"`
	BuyPrice             string   `json:"buy_price,omitempty"`
	SellPrice            string   `json:"sell_price,omitempty"`
	FeeCurrency          string   `json:"fee_currency"`
	ExchangeOrderFeeRate string   `json:"exchange_order_fee_rate,omitempty"`
	SpreadRate           string   `json:"spread_rate"`
}

func PairDTO(p *entity.Pair) *AdminPair {
	ap := &AdminPair{
		C1: &Coin{
			CoinId:              p.C1.Coin.CoinId,
			ChainId:             p.C1.Coin.ChainId,
			ContractAddress:     p.C1.ContractAddress,
			Address:             p.C1.Address,
			Tag:                 p.C1.Tag,
			MinDeposit:          p.C1.MinDeposit,
			MinOrderSize:        p.C1.MinOrderSize,
			MaxOrderSize:        p.C1.MaxOrderSize,
			MinWithdrawalSize:   p.C1.MinWithdrawalSize,
			MinWithdrawalFee:    p.C1.WithdrawalMinFee,
			OrderPrecision:      p.C1.OrderPrecision,
			WithdrawalPrecision: p.C1.WithdrawalPrecision,
			SetChain:            p.C1.SetChain,
		},
		C2: &Coin{
			CoinId:              p.C2.Coin.CoinId,
			ChainId:             p.C2.Coin.ChainId,
			ContractAddress:     p.C2.ContractAddress,
			Address:             p.C2.Address,
			Tag:                 p.C2.Tag,
			MinDeposit:          p.C2.MinDeposit,
			MinOrderSize:        p.C2.MinOrderSize,
			MaxOrderSize:        p.C2.MaxOrderSize,
			MinWithdrawalSize:   p.C2.MinWithdrawalSize,
			MinWithdrawalFee:    p.C2.WithdrawalMinFee,
			OrderPrecision:      p.C2.OrderPrecision,
			WithdrawalPrecision: p.C2.WithdrawalPrecision,
			SetChain:            p.C2.SetChain,
		},

		ContractAddress:      p.ContractAddress,
		FeeTier:              p.FeeTier,
		Liquidity:            p.Liquidity,
		BuyPrice:             p.Price1,
		SellPrice:            p.Price2,
		SpreadRate:           p.SpreadRate,
		FeeCurrency:          p.FeeCurrency,
		ExchangeOrderFeeRate: p.OrderFeeRate,
	}

	if ap.BuyPrice == ap.SellPrice {
		ap.Price = ap.BuyPrice
		ap.BuyPrice = ""
		ap.SellPrice = ""
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
