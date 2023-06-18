package binance

import (
	"exchange-provider/internal/entity"
)

func (ex *exchange) calcPrice(p0, p1 float64, in, out entity.TokenId, p *entity.Pair) float64 {
	if p.EP.(*ExchangePair).HasIntermediaryCoin {
		if p.T1.String() == in.String() {
			return applyFee((applyFee(p0, p.EP.(*ExchangePair).BinanceFeeRate1) /
				p1), p.EP.(*ExchangePair).BinanceFeeRate2)
		} else {
			return 1 / applyFee(applyFee(p1, p.EP.(*ExchangePair).BinanceFeeRate2)/p0,
				p.EP.(*ExchangePair).BinanceFeeRate1)
		}
	}
	if p.T1.String() == in.String() {
		return applyFee(p0, p.EP.(*ExchangePair).BinanceFeeRate1)
	} else {
		return 1 / applyFee(1/p0, p.EP.(*ExchangePair).BinanceFeeRate1)
	}
}

func (ex *exchange) price(p *entity.Pair) (float64, float64, error) {

	ep := p.EP.(*ExchangePair)
	if ep.HasIntermediaryCoin {

		bc := p.T1.ET.(*Token).Coin
		qc := ep.IC1.Coin
		p0, err := ex.si.getPrice(bc, qc)
		if err != nil {
			return 0, 0, err
		}

		bc = p.T2.ET.(*Token).Coin
		qc = ep.IC2.Coin
		p1, err := ex.si.getPrice(bc, qc)
		if err != nil {
			return 0, 0, err
		}
		return p0, p1, nil
	}

	bc := p.T1.ET.(*Token).Coin
	qc := p.T2.ET.(*Token).Coin
	p0, err := ex.si.getPrice(bc, qc)
	if err != nil {
		return 0, 0, err
	}
	return p0, 0, nil
}

func applyFee(price, fee float64) float64 {
	return price - (price * fee)
}
