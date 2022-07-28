package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

type coin struct {
	CoinId            string `json:"coin_id"`
	ChainId           string `json:"chain_id"`
	MinOrderSize      string `json:"min_order_size,omitempty"`
	MaxOrderSize      string `json:"max_order_size,omitempty"`
	MinWithdrawalSize string `json:"min_withdrawal_size,omitempty"`
	WithdrawalMinFee  string `json:"withdrawal_min_fee,omitempty"`
}
type coinConfig struct {
	SetChain bool `json:"set_chain"`
}

type exchangeConfig struct {
	BaseC  *coinConfig `json:"base_coin"`
	QuoteC *coinConfig `json:"quote_coin"`
}

type exPair struct {
	BC   *coin                      `json:"base_coin"`
	QC   *coin                      `json:"quote_coin"`
	ExsC map[string]*exchangeConfig `json:"exchanges_config"`
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

		if len(p.ExsC) == 0 {
			return errors.Wrap(errors.ErrBadRequest, "exchanges config must be set")
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
	for ex, conf := range req.ExsC {
		ep := &entity.Pair{
			BC: &entity.PairCoin{
				Coin:     bc,
				SetChain: conf.BaseC.SetChain,
			},
			QC: &entity.PairCoin{
				Coin:     qc,
				SetChain: conf.QuoteC.SetChain,
			},
		}

		exchangePairs[ex] = ep

	}
	return exchangePairs
}

type pair struct {
	BC *coin `json:"base_coin"`
	QC *coin `json:"quote_coin"`

	Price                string `json:"price,omitempty"`
	FeeCurrency          string `json:"fee_currency,omitempty"`
	ExchangeOrderFeeRate string `json:"exchange_order_fee_rate,omitempty"`
}

func PairDTO(p *entity.Pair) *pair {
	return &pair{
		BC: &coin{
			CoinId:            p.BC.Coin.CoinId,
			ChainId:           p.BC.Coin.ChainId,
			MinOrderSize:      p.BC.MinOrderSize,
			MaxOrderSize:      p.BC.MaxOrderSize,
			MinWithdrawalSize: p.BC.MinWithdrawalSize,
			WithdrawalMinFee:  p.BC.WithdrawalMinFee,
		},
		QC: &coin{
			CoinId:            p.QC.Coin.CoinId,
			ChainId:           p.QC.Coin.ChainId,
			MinOrderSize:      p.QC.MinOrderSize,
			MaxOrderSize:      p.QC.MaxOrderSize,
			MinWithdrawalSize: p.QC.MinWithdrawalSize,
			WithdrawalMinFee:  p.QC.WithdrawalMinFee,
		},
		Price:                p.Price,
		FeeCurrency:          p.FeeCurrency,
		ExchangeOrderFeeRate: p.OrderFeeRate,
	}
}

type AddPairsErr struct {
	Pair *pair  `json:"pair"`
	Err  string `json:"error"`
}
type AddPairsResult struct {
	Addedd []*pair        `json:"added_pairs"`
	Exs    []*pair        `json:"existed_pairs"`
	Failed []*AddPairsErr `json:"failed_pairs"`
}

func FromEntity(r *entity.AddPairsResult) *AddPairsResult {
	res := &AddPairsResult{
		Addedd: []*pair{},
		Exs:    []*pair{},
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
