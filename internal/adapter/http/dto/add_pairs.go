package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"time"
)

type coin struct {
	CoinId            string `json:"coin_id"`
	ChainId           string `json:"chain_id"`
	BlockTime         string `json:"block_time,omitempty"`
	MinOrderSize      string `json:"min_size,omitempty"`
	MaxOrderSize      string `json:"max_size,omitempty"`
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
		if p.BC.CoinId == "" || p.BC.ChainId == "" || p.BC.BlockTime == "" {
			return errors.Wrap(errors.ErrBadRequest, "base coin must have id, chain and block time")
		}
		if p.QC.CoinId == "" || p.QC.ChainId == "" || p.QC.BlockTime == "" {
			return errors.Wrap(errors.ErrBadRequest, "quote coin must have id, chain and block time")
		}

		if len(p.ExsC) == 0 {
			return errors.Wrap(errors.ErrBadRequest, "exchanges config must be set")
		}

	}
	return nil
}

func (r *exPair) BaseCoin() (*entity.Coin, error) {
	c := &entity.Coin{
		Id: r.BC.CoinId,
		Chain: &entity.Chain{
			Id: r.BC.ChainId,
		},
	}

	bt, err := toTime(r.BC.BlockTime)
	if err != nil {
		return nil, errors.Wrap(err, errors.ErrBadRequest, "block_time should follow the format 10s, 10m")
	}
	c.Chain.BlockTime = bt
	return c, nil
}

func (r *exPair) QuoteCoin() (*entity.Coin, error) {
	c := &entity.Coin{
		Id: r.QC.CoinId,
		Chain: &entity.Chain{
			Id: r.QC.ChainId,
		},
	}

	bt, err := toTime(r.QC.BlockTime)
	if err != nil {
		return nil, err
	}
	c.Chain.BlockTime = bt
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

	Price                string `json:"price"`
	FeeCurrency          string `json:"fee_currency"`
	ExchangeOrderFeeRate string `json:"exchange_order_fee_rate"`
}

func PairDTO(p *entity.Pair) *pair {
	return &pair{
		BC: &coin{
			CoinId:            p.BC.Coin.Id,
			ChainId:           p.BC.Coin.Chain.Id,
			BlockTime:         p.BC.Coin.Chain.BlockTime.String(),
			MinOrderSize:      p.BC.MinOrderSize,
			MaxOrderSize:      p.BC.MaxOrderSize,
			MinWithdrawalSize: p.BC.MinWithdrawalSize,
			WithdrawalMinFee:  p.BC.WithdrawalMinFee,
		},
		QC: &coin{
			CoinId:            p.QC.Coin.Id,
			ChainId:           p.QC.Coin.Chain.Id,
			BlockTime:         p.QC.Coin.Chain.BlockTime.String(),
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

func toTime(t string) (time.Duration, error) {
	return time.ParseDuration(t)

}
