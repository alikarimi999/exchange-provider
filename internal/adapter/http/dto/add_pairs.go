package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

type Coin struct {
	CoinId              string `json:"coin_id"`
	ChainId             string `json:"chain_id"`
	MinOrderSize        string `json:"min_order_size,omitempty"`
	MaxOrderSize        string `json:"max_order_size,omitempty"`
	MinWithdrawalSize   string `json:"min_withdrawal_size,omitempty"`
	MinWithdrawalFee    string `json:"min_withdrawal_fee,omitempty"`
	OrderPrecision      int    `json:"order_precision,omitempty"`
	WithdrawalPrecision int    `json:"withdrawal_precision,omitempty"`
	SetChain            bool   `json:"set_chain,omitempty"`
}
type coinConfig struct {
	WithdrawalPrecision int `json:"withdrawal_precision,omitempty"`
}

type ExchangeConfig struct {
	BC *coinConfig `json:"base_coin"`
	QC *coinConfig `json:"quote_coin"`
}
type exPair struct {
	BC        *Coin                      `json:"base_coin"`
	QC        *Coin                      `json:"quote_coin"`
	Exchanges map[string]*ExchangeConfig `json:"exchange_config"`
}

type AddPairsRequest struct {
	Pairs []*exPair `json:"pairs"`
}

// check there wasn't any zero values in the request
// if there was return an error that the value must set
func (r *AddPairsRequest) Validate() error {
	for _, p := range r.Pairs {
		if p.BC.CoinId == "" || p.BC.ChainId == "" {
			return errors.Wrap(errors.ErrBadRequest, "base coin must have id")
		}
		if p.QC.CoinId == "" || p.QC.ChainId == "" {
			return errors.Wrap(errors.ErrBadRequest, "quote coin must have id")
		}

		if len(p.Exchanges) == 0 {
			return errors.Wrap(errors.ErrBadRequest, "at least one exchange must be set")
		}

	}
	return nil
}

func (r *exPair) BaseCoin() (*entity.Coin, error) {
	c := &entity.Coin{
		CoinId:  r.BC.CoinId,
		ChainId: r.BC.ChainId,
	}

	return c, nil
}

func (r *exPair) QuoteCoin() (*entity.Coin, error) {
	c := &entity.Coin{
		CoinId:  r.QC.CoinId,
		ChainId: r.QC.ChainId,
	}

	return c, nil
}

func (req *exPair) ExchangePairs(bc, qc *entity.Coin) map[string]*entity.Pair {
	exchangePairs := map[string]*entity.Pair{}
	for id, conf := range req.Exchanges {
		ep := &entity.Pair{
			BC: &entity.PairCoin{
				Coin:                bc,
				WithdrawalPrecision: conf.BC.WithdrawalPrecision,
			},

			QC: &entity.PairCoin{
				Coin:                qc,
				WithdrawalPrecision: conf.QC.WithdrawalPrecision,
			},
		}

		exchangePairs[id] = ep

	}
	return exchangePairs
}

type Pair struct {
	BC *Coin `json:"base_coin"`
	QC *Coin `json:"quote_coin"`

	BestAskPrice         string `json:"best_ask_price,omitempty"`
	BestBidPrice         string `json:"best_bid_price,omitempty"`
	FeeCurrency          string `json:"fee_currency,omitempty"`
	ExchangeOrderFeeRate string `json:"exchange_order_fee_rate,omitempty"`
	FeeRate              string `json:"fee_rate,omitempty"`
}

func PairDTO(p *entity.Pair) *Pair {
	return &Pair{
		BC: &Coin{
			CoinId:              p.BC.Coin.CoinId,
			ChainId:             p.BC.Coin.ChainId,
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
		FeeCurrency:          p.FeeCurrency,
		ExchangeOrderFeeRate: p.OrderFeeRate,
		FeeRate:              p.Fee,
	}
}

type AddPairsErr struct {
	Pair *Pair  `json:"pair"`
	Err  string `json:"error"`
}
type AddPairsResult struct {
	Addedd []*Pair        `json:"added_pairs"`
	Exs    []*Pair        `json:"existed_pairs"`
	Failed []*AddPairsErr `json:"failed_pairs"`
}

func FromEntity(r *entity.AddPairsResult) *AddPairsResult {
	res := &AddPairsResult{
		Addedd: []*Pair{},
		Exs:    []*Pair{},
		Failed: []*AddPairsErr{},
	}
	for _, p := range r.Added {
		res.Addedd = append(res.Addedd, PairDTO(p))
	}
	for _, p := range r.Existed {
		res.Exs = append(res.Exs, PairDTO(p))
	}
	for _, p := range r.Failed {
		res.Failed = append(res.Failed, &AddPairsErr{
			Pair: PairDTO(p.Pair),
			Err:  p.Err.Error(),
		})
	}
	return res

}

type AddPairsResponse struct {
	Exchanges map[string]*AddPairsResult `json:"exchanges"`
}
