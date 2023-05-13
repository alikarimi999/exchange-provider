package kucoin

import (
	"exchange-provider/internal/entity"
)

func (k *kucoinExchange) spread(lvl uint, p *entity.Pair, price float64) (float64, error) {
	var s float64
	t1 := p.T1.ET.(*Token)
	t2 := p.T2.ET.(*Token)
	if len(p.Spreads) > 0 {
		s = p.Spread(lvl)
	} else if t2.Currency == p.T2.ET.(*Token).StableToken {
		s = k.st.GetByPrice(lvl, price)
	} else if t1.Currency == p.T1.ET.(*Token).StableToken {
		s = k.st.GetByPrice(lvl, 1)
	} else {
		bcDollar, err := k.stableTicker(t1.Currency, p.T1.ET.(*Token).StableToken)
		if err != nil {
			return 0, err
		}
		s = k.st.GetByPrice(lvl, bcDollar)
	}

	return s, nil
}
