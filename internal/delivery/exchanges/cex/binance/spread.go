package binance

import "exchange-provider/internal/entity"

func (ex *exchange) spread(lvl uint, p *entity.Pair, price float64) (float64, error) {
	if p.T1.Id.Symbol == p.T2.Id.Symbol {
		return 0, nil
	}

	var s float64
	t1 := p.T1.ET.(*Token)
	t2 := p.T2.ET.(*Token)
	if len(p.Spreads) > 0 {
		s = p.Spread(lvl)
	} else if t2.Coin == p.T2.ET.(*Token).StableToken {
		s = ex.st.GetByPrice(lvl, price)
	} else if t1.Coin == p.T1.ET.(*Token).StableToken {
		s = ex.st.GetByPrice(lvl, 1)
	} else {
		bcDollar, err := ex.ticker(t1.Coin, p.T1.ET.(*Token).StableToken)
		if err != nil {
			return 0, err
		}
		s = ex.st.GetByPrice(lvl, bcDollar)
	}

	return s, nil
}
