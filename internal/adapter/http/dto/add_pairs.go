package dto

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"time"
)

type coin struct {
	Id        string `json:"id"`
	Chain     string `json:"chain"`
	BlockTime string `json:"block_time"`
}
type coinConfig struct {
	SetChain  bool `json:"set_chain"`
	Precision int  `json:"precision"`
}

type exchangeConfig struct {
	BaseC  *coinConfig `json:"base_coin"`
	QuoteC *coinConfig `json:"quote_coin"`
}

type pair struct {
	BaseC  *coin                      `json:"base_coin"`
	QuoteC *coin                      `json:"quote_coin"`
	ExsC   map[string]*exchangeConfig `json:"exchanges_config"`
}

type AddPairsRequest struct {
	Pairs []*pair `json:"pairs"`
}

// check there wasn't any zero values in the request
// if there was return an error that the value must set
func (r *AddPairsRequest) Validate() error {
	for _, p := range r.Pairs {
		if p.BaseC.Id == "" || p.BaseC.Chain == "" || p.BaseC.BlockTime == "" {
			return errors.Wrap(errors.ErrBadRequest, "base coin must have id, chain and block time")
		}
		if p.QuoteC.Id == "" || p.QuoteC.Chain == "" || p.QuoteC.BlockTime == "" {
			return errors.Wrap(errors.ErrBadRequest, "quote coin must have id, chain and block time")
		}
		for ex, conf := range p.ExsC {
			if conf.BaseC.Precision == 0 {
				return errors.Wrap(errors.ErrBadRequest,
					fmt.Sprintf("base coin config for exchange '%s' must have 'precision' and 'set_chain'", ex))
			}
			if conf.QuoteC.Precision == 0 {
				return errors.Wrap(errors.ErrBadRequest,
					fmt.Sprintf("quote coin config for exchange '%s' must have 'precision' and 'set_chain'", ex))
			}
		}
	}
	return nil
}

func (r *pair) BaseCoin() (*entity.Coin, error) {
	c := &entity.Coin{
		Id: r.BaseC.Id,
		Chain: &entity.Chain{
			Id: r.BaseC.Chain,
		},
	}

	bt, err := toTime(r.BaseC.BlockTime)
	if err != nil {
		return nil, errors.Wrap(err, errors.ErrBadRequest, "block_time should follow the format 10s, 10m")
	}
	c.Chain.BlockTime = bt
	return c, nil
}

func (r *pair) QuoteCoin() (*entity.Coin, error) {
	c := &entity.Coin{
		Id: r.QuoteC.Id,
		Chain: &entity.Chain{
			Id: r.QuoteC.Chain,
		},
	}

	bt, err := toTime(r.QuoteC.BlockTime)
	if err != nil {
		return nil, err
	}
	c.Chain.BlockTime = bt
	return c, nil
}

func (req *pair) ExchangePairs(bc, qc *entity.Coin) map[string]*entity.ExchangePair {
	exchangePairs := map[string]*entity.ExchangePair{}
	for ex, conf := range req.ExsC {
		ep := &entity.ExchangePair{
			BC: &entity.CoinConfig{
				Coin:      bc,
				Precision: conf.BaseC.Precision,
				SetChain:  conf.BaseC.SetChain,
			},
			QC: &entity.CoinConfig{
				Coin:      qc,
				Precision: conf.QuoteC.Precision,
				SetChain:  conf.QuoteC.SetChain,
			},
		}

		exchangePairs[ex] = ep

	}
	return exchangePairs
}

func toTime(t string) (time.Duration, error) {
	return time.ParseDuration(t)

}
