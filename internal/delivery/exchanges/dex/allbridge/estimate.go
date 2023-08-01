package allbridge

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/calculate"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

func (d *exchange) EstimateAmountOut(in, out entity.TokenId, amountIn float64, lvl uint,
	opts interface{}) ([]*entity.EstimateAmount, error) {

	p, err := d.pairs.Get(d.Id(), in.String(), out.String())
	if err != nil {
		return nil, err
	}
	es0 := &entity.EstimateAmount{
		P:           p,
		FeeCurrency: in,
		AmountIn:    amountIn, ExchangeFee: d.cfg.ExchangeFee, FeeRate: d.cfg.FeeRate,
	}
	ess := []*entity.EstimateAmount{es0}

	if p.T1.String() == in.String() {
		min := p.T1.Min
		max := p.T1.Max
		if (min != 0 && amountIn < min) || (max != 0 && amountIn > max) {
			return ess, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(fmt.Sprintf("min is %f and max is %f", min, max)))
		}

	} else {
		min := p.T2.Min
		max := p.T2.Max
		if (min != 0 && amountIn < min) || (max != 0 && amountIn > max) {
			return ess, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(fmt.Sprintf("min is %f and max is %f", min, max)))

		}
	}

	ss, err := d.estimate(in, out, amountIn, lvl, es0)
	if err != nil {
		return nil, err
	}
	ls := len(ss) - 1
	es0.AmountOut = ss[ls][len(ss[ls])-1].amountOut
	es0.Data = ss
	es1 := &entity.EstimateAmount{FeeCurrency: out,
		AmountIn: es0.AmountOut, ExchangeFee: d.cfg.ExchangeFee, FeeRate: d.cfg.FeeRate}
	ss1, err := d.estimate(out, in, es0.AmountOut, lvl, es1)
	if err == nil {
		ls = len(ss1) - 1
		amIn := ss1[ls][len(ss1[ls])-1].amountOut

		amIn = (amountIn * es0.AmountOut) / amIn
		amInPlusFee := (amIn / (1 - es1.FeeRate))
		feeAmount := amInPlusFee - amIn
		amIn = amInPlusFee + es1.ExchangeFeeAmount

		es1.AmountIn = amIn
		es1.AmountOut = amountIn
		es1.FeeAmount = feeAmount
		ess = append(ess, es1)

	} else {
		d.l.Debug(d.agent("EstimateAmountOut"), err.Error())
	}

	return ess, err
}

type route struct {
	in        entity.TokenId
	out       entity.TokenId
	ex        entity.Exchange
	amountIn  float64
	amountOut float64
	es        *entity.EstimateAmount
}
type steps map[int][]*route

func (a *exchange) estimate(in, out entity.TokenId, amount float64,
	lvl uint, es *entity.EstimateAmount) (steps, error) {
	ss := make(steps)
	if a.tl.isTokenExists(in) && a.tl.isTokenExists(out) {
		es.InUsd = 1
		es.OutUsd = 1
		es.ExchangeFeeAmount = es.ExchangeFee
		amount = amount - es.ExchangeFeeAmount
		es.FeeAmount = amount * es.FeeRate
		amount = amount - es.FeeAmount

		r, err := a.internalEstimate(in, out, amount, lvl)
		if err != nil {
			return nil, err
		}

		ss[0] = append(ss[0], r)
		return ss, nil

	} else if !a.tl.isTokenExists(in) && a.tl.isTokenExists(out) {
		r0, err := a.externalEstimate(in, entity.TokenId{}, amount, lvl, true)
		if err != nil {
			return nil, err
		}
		ss[0] = append(ss[0], r0)
		es.InUsd = r0.es.InUsd
		es.OutUsd = 1
		es.ExchangeFee = r0.es.ExchangeFee
		es.ExchangeFeeAmount = r0.es.ExchangeFeeAmount
		es.FeeRate = r0.es.FeeRate
		es.FeeAmount = r0.es.FeeAmount
		r1, err := a.internalEstimate(r0.out, out, r0.amountOut, lvl)
		if err != nil {
			return nil, err
		}
		ss[0] = append(ss[0], r1)
		return ss, nil

	} else if a.tl.isTokenExists(in) && !a.tl.isTokenExists(out) {
		es.InUsd = 1
		es.ExchangeFeeAmount = es.ExchangeFee
		amount = amount - es.ExchangeFee
		es.FeeAmount = amount * es.FeeRate
		amount = amount - es.FeeAmount

		ts, err := a.tl.tokensInNetwork(out.Network)
		if err != nil {
			return nil, err
		}

		for _, tt := range ts {
			r0, err := a.internalEstimate(in, tt.TokenId, amount, lvl)
			if err != nil {
				continue
			}
			ss[0] = append(ss[0], r0)
			break
		}
		if len(ss) != 1 {
			return nil, errors.Wrap(errors.ErrInternal,
				errors.NewMesssage("unable to estimate this pair"))
		}

		r1, err := a.externalEstimate(ss[0][0].out, out, ss[0][0].amountOut, lvl, false)
		if err != nil {
			return nil, err
		}
		ss[1] = append(ss[1], r1)
		es.OutUsd = r1.es.OutUsd
		return ss, nil

	} else {
		r0, err := a.externalEstimate(in, in, amount, lvl, true)
		if err != nil {
			return nil, err
		}

		ss[0] = append(ss[0], r0)
		es.InUsd = r0.es.InUsd
		es.ExchangeFee = r0.es.ExchangeFee
		es.ExchangeFeeAmount = r0.es.ExchangeFeeAmount
		es.FeeRate = r0.es.FeeRate
		es.FeeAmount = r0.es.FeeAmount

		ts, err := a.tl.tokensInNetwork(out.Network)
		if err != nil {
			return nil, err
		}

		for _, t := range ts {
			r1, err := a.internalEstimate(r0.out, t.TokenId, r0.amountOut, lvl)
			if err != nil {
				continue
			} else {
				ss[0] = append(ss[0], r1)
				break
			}
		}
		if len(ss[0]) != 2 {
			return nil, errors.Wrap(errors.ErrInternal,
				errors.NewMesssage("unable to estimate this pair"))
		}

		r2, err := a.externalEstimate(ss[0][1].out, out, ss[0][1].amountOut, lvl, false)
		if err != nil {
			return nil, err
		}

		ss[1] = append(ss[1], r2)
		es.OutUsd = r2.es.OutUsd
		return ss, nil
	}
}

func (a *exchange) externalEstimate(in, out entity.TokenId, amount float64, lvl uint, In bool) (*route, error) {
	ts, err := a.tl.tokensInNetwork(in.Network)
	if err != nil {
		return nil, err
	}
	opts := &evm.EstimateOpts{
		CustomizeFee: true,
		ExchangeFee:  a.cfg.ExchangeFee,
		FeeRate:      a.cfg.FeeRate,
	}

	if In {
		for _, tt := range ts {
			out := entity.TokenId{
				Symbol:   tt.Symbol,
				Standard: tt.Standard,
				Network:  tt.Network,
			}
			exs := a.exs.GetAll()
			estimates := []*estimate{}
			for _, ex := range exs {
				if ex.Type() == entity.EvmDEX && ex.IsEnable() {
					dex := ex.(entity.EVMDex)
					if dex.Network() == in.Network && (dex.Standard() == in.Standard ||
						in.Symbol == in.Standard) {
						es, err := dex.EstimateAmountOut(in, out, amount, lvl, opts)
						if err == nil && es[0].AmountOut > 0 {
							estimates = append(estimates, &estimate{
								ex:             ex,
								EstimateAmount: es[0],
							})
						}
					}
				}
			}
			if len(estimates) == 0 {
				continue
			}

			es0 := &estimate{
				EstimateAmount: &entity.EstimateAmount{},
			}

			for _, es := range estimates {
				if es.EstimateAmount != nil && es.AmountOut > es0.AmountOut {
					es0 = es
				}
			}

			return &route{
				in:        in,
				out:       out,
				ex:        es0.ex,
				amountIn:  amount,
				amountOut: es0.AmountOut,
				es:        es0.EstimateAmount,
			}, nil
		}
	} else {
		exs := a.exs.GetAll()
		estimates := []*estimate{}
		for _, ex := range exs {
			if ex.Type() == entity.EvmDEX && ex.IsEnable() {
				dex := ex.(entity.EVMDex)
				if dex.Network() == in.Network && (dex.Standard() == in.Standard ||
					in.Symbol == in.Standard) {
					es, err := dex.EstimateAmountOut(in, out, amount, lvl, opts)
					if err == nil && es[0].AmountOut > 0 {
						estimates = append(estimates, &estimate{
							ex:             ex,
							EstimateAmount: es[0],
						})
					}
				}
			}
		}
		if len(estimates) == 0 {
			return nil, errors.Wrap(errors.ErrInternal, errors.NewMesssage("unable to estimate this pair"))
		}

		es0 := &estimate{
			EstimateAmount: &entity.EstimateAmount{},
		}

		for _, es := range estimates {
			if es.EstimateAmount != nil && es.AmountOut > es0.AmountOut {
				es0 = es
			}
		}

		return &route{
			in:        in,
			out:       out,
			ex:        es0.ex,
			amountIn:  amount,
			amountOut: es0.AmountOut,
			es:        es0.EstimateAmount,
		}, nil
	}
	return nil, errors.Wrap(errors.ErrInternal, errors.NewMesssage("unable to estimate this pair"))
}
func (ex *exchange) internalEstimate(in, out entity.TokenId, amount float64,
	lvl uint) (*route, error) {
	In, err := ex.tl.getTokenInfo(in)
	if err != nil {
		return nil, err
	}
	Out, err := ex.tl.getTokenInfo(out)
	if err != nil {
		return nil, err
	}

	ps, err := getPoolInfo([]*types.TokenInfo{In, Out})
	if err != nil {
		return nil, err
	}

	amountOut := calculate.Estimate(amount, In, Out, ps[In.Chain][In.PoolAddress],
		ps[Out.Chain][Out.PoolAddress])

	return &route{
		in:        in,
		out:       out,
		ex:        ex,
		amountIn:  amount,
		amountOut: amountOut,
	}, nil

}

type estimate struct {
	ex entity.Exchange
	*entity.EstimateAmount
}
