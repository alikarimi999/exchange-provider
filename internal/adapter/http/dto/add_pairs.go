package dto

import (
	kdto "order_service/internal/delivery/exchanges/kucoin/dto"
	"order_service/internal/entity"

	"order_service/pkg/errors"
)

type Coin struct {
	CoinId              string  `json:"coin_id"`
	ChainId             string  `json:"chain_id"`
	Address             string  `json:"address"`
	Tag                 string  `json:"tag"`
	BlockTime           string  `json:"block_time"`
	MinDeposit          float64 `json:"min_deposit"`
	MinOrderSize        string  `json:"min_order_size"`
	MaxOrderSize        string  `json:"max_order_size"`
	MinWithdrawalSize   string  `json:"min_withdrawal_size"`
	MinWithdrawalFee    string  `json:"min_withdrawal_fee"`
	OrderPrecision      int     `json:"order_precision"`
	WithdrawalPrecision int     `json:"withdrawal_precision"`
	SetChain            bool    `json:"set_chain"`
}

type kuPair struct {
	BC *Coin `json:"base_coin"`
	QC *Coin `json:"quote_coin"`
}

func (p *kuPair) Map() *kdto.Pair {
	pair := &kdto.Pair{
		BC: &kdto.Coin{
			CoinId:              p.BC.CoinId,
			ChainId:             p.BC.ChainId,
			WithdrawalPrecision: p.BC.WithdrawalPrecision,
		},
		QC: &kdto.Coin{
			CoinId:              p.QC.CoinId,
			ChainId:             p.QC.ChainId,
			WithdrawalPrecision: p.QC.WithdrawalPrecision,
		},
	}

	pair.BC.BlockTime, _ = toTime(p.BC.BlockTime)
	pair.QC.BlockTime, _ = toTime(p.QC.BlockTime)

	return pair
}

type KucoinAddPairsRequest struct {
	Pairs []*kuPair `json:"pairs"`
}

// check there wasn't any zero values in the request
// if there was return an error that the value must set
func (r *KucoinAddPairsRequest) Validate() error {
	for _, p := range r.Pairs {
		if p.BC.CoinId == "" || p.BC.ChainId == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("base coin must have id"))
		}
		if p.QC.CoinId == "" || p.QC.ChainId == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("quote coin must have id"))
		}

		if p.BC.WithdrawalPrecision == 0 || p.QC.WithdrawalPrecision == 0 {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("withdrawal_precision must be set"))
		}

		if _, err := toTime(p.BC.BlockTime); err != nil {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error()))
		}

		if _, err := toTime(p.QC.BlockTime); err != nil {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error()))
		}

	}
	return nil
}

func (r *kuPair) BaseCoin() (*entity.Coin, error) {
	c := &entity.Coin{
		CoinId:  r.BC.CoinId,
		ChainId: r.BC.ChainId,
	}

	return c, nil
}

func (r *kuPair) QuoteCoin() (*entity.Coin, error) {
	c := &entity.Coin{
		CoinId:  r.QC.CoinId,
		ChainId: r.QC.ChainId,
	}

	return c, nil
}

type AdminPair struct {
	BC *Coin `json:"base_coin"`
	QC *Coin `json:"quote_coin"`

	BestAskPrice         string `json:"best_ask_price"`
	BestBidPrice         string `json:"best_bid_price"`
	FeeCurrency          string `json:"fee_currency"`
	ExchangeOrderFeeRate string `json:"exchange_order_fee_rate"`
	SpreadRate           string `json:"spread_rate"`
}

func PairDTO(p *entity.Pair) *AdminPair {
	return &AdminPair{
		BC: &Coin{
			CoinId:              p.BC.Coin.CoinId,
			ChainId:             p.BC.Coin.ChainId,
			Address:             p.BC.Address,
			Tag:                 p.BC.Tag,
			BlockTime:           p.BC.BlockTime.String(),
			MinDeposit:          p.BC.MinDeposit,
			MinOrderSize:        p.BC.MinOrderSize,
			MaxOrderSize:        p.BC.MaxOrderSize,
			MinWithdrawalSize:   p.BC.MinWithdrawalSize,
			MinWithdrawalFee:    p.BC.WithdrawalMinFee,
			OrderPrecision:      p.BC.OrderPrecision,
			WithdrawalPrecision: p.BC.WithdrawalPrecision,
			SetChain:            p.BC.SetChain,
		},
		QC: &Coin{
			CoinId:              p.QC.Coin.CoinId,
			ChainId:             p.QC.Coin.ChainId,
			Address:             p.QC.Address,
			Tag:                 p.QC.Tag,
			BlockTime:           p.BC.BlockTime.String(),
			MinDeposit:          p.QC.MinDeposit,
			MinOrderSize:        p.QC.MinOrderSize,
			MaxOrderSize:        p.QC.MaxOrderSize,
			MinWithdrawalSize:   p.QC.MinWithdrawalSize,
			MinWithdrawalFee:    p.QC.WithdrawalMinFee,
			OrderPrecision:      p.QC.OrderPrecision,
			WithdrawalPrecision: p.QC.WithdrawalPrecision,
			SetChain:            p.QC.SetChain,
		},
		BestAskPrice:         p.BestAsk,
		BestBidPrice:         p.BestBid,
		SpreadRate:           p.SpreadRate,
		FeeCurrency:          p.FeeCurrency,
		ExchangeOrderFeeRate: p.OrderFeeRate,
	}
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
	res.Addedd = append(res.Addedd, r.Added...)
	res.Exs = append(res.Exs, r.Existed...)

	for _, p := range r.Failed {
		res.Failed = append(res.Failed, &PairsErr{
			Pair: p.Pair,
			Err:  p.Err.Error(),
		})
	}
	return res

}
