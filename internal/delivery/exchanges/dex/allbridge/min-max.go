package allbridge

import "exchange-provider/internal/entity"

func (ex *exchange) minAndMax(t *entity.Token, p *entity.Pair,
	dex entity.EVMDex, tokenEfa map[string]float64) (float64, error) {
	tEfa, ok := tokenEfa[t.String()]
	if !ok {
		efa, _, err := dex.ExchangeFeeAmount(t.Id, p, ex.cfg.ExchangeFee)
		if err != nil {
			return 0, err
		}
		tokenEfa[t.String()] = efa
		tEfa = efa
	}

	min := tEfa + (tEfa * ex.cfg.FeeRate * 2)
	min = min + (min * 0.1)

	return min, nil
}
