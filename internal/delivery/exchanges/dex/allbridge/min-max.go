package allbridge

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
	"fmt"
)

func (ex *exchange) minAndMax(tEfa, feeRate float64) float64 {
	min := tEfa + (tEfa * feeRate * 2)
	min = min + (min * 0.1)
	return min * 2
}

func (ex *exchange) updateMinAndMax(p *entity.Pair, tl tokenList, exchangeFee, feeRate float64,
	tokensEfa map[string]float64) bool {
	agent := ex.agent(fmt.Sprintf("updateMinAndMax: %s", p.String()))
	var update bool
	if err := ex.updateMinForToken(tl, exchangeFee, feeRate, p.T1, tokensEfa); err != nil {
		ex.l.Debug(agent, err.Error())
	} else {
		update = true
	}
	if err := ex.updateMinForToken(tl, exchangeFee, feeRate, p.T2, tokensEfa); err != nil {
		ex.l.Debug(agent, err.Error())
	} else {
		update = true
	}
	return update
}
func (ex *exchange) updateMinForToken(tl tokenList, exchangeFee, feeRate float64, t *entity.Token, tokensEfa map[string]float64) error {
	var efa float64

	if tl.isTokenExists(t.Id) {
		efa = exchangeFee
	} else {
		ef, ok := tokensEfa[t.Id.String()]
		if ok {
			efa = ef
		} else {
			et := t.ET.(*types.EToken)
			e, err := ex.exs.GetByNID(et.ExtraExchange)
			if err != nil {
				return err
			}

			ep, err := ex.pairs.Get(e.Id(), t.Id.String(), et.OtherToken)
			if err != nil {
				return err
			}

			ef, _, err := e.(entity.EVMDex).ExchangeFeeAmount(t.Id, ep, exchangeFee)
			if err != nil {
				return err
			}
			efa = ef
		}
		tokensEfa[t.Id.String()] = efa
	}
	tokensEfa[t.Id.String()] = efa

	min := ex.minAndMax(efa, feeRate)
	t.Min = min
	return nil
}
