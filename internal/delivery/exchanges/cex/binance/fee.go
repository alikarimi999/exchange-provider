package binance

import (
	"exchange-provider/internal/delivery/exchanges/cex/binance/types"
	"exchange-provider/internal/entity"

	"github.com/adshao/go-binance/v2"
)

func getBcQcWcFeeRate(o *types.Order, p *entity.Pair,
	i int) (bc, qc, wc *Token, feeRate float64) {

	if i == 0 {
		if len(o.Swaps) == 2 {
			if o.Swaps[0].In.String() == p.T1.Id.String() {
				bc = p.T1.ET.(*Token)
				qc = p.EP.(*ExchangePair).IC1
				feeRate = p.EP.(*ExchangePair).BinanceFeeRate1
			} else {
				bc = p.T2.ET.(*Token)
				qc = p.EP.(*ExchangePair).IC2
				feeRate = p.EP.(*ExchangePair).BinanceFeeRate2
			}
		} else {
			bc = p.T1.ET.(*Token)
			qc = p.T2.ET.(*Token)
			feeRate = p.EP.(*ExchangePair).BinanceFeeRate1
			if o.Swaps[0].Side == binance.SideTypeSell {
				wc = qc
			} else {
				wc = bc
			}
		}
		o.Swaps[0].InAmountRequested = o.Deposit.Amount
	} else {

		o.Swaps[1].InAmountRequested = o.Swaps[0].OutAmount
		if o.Swaps[1].Out.String() == p.T2.String() {
			bc = p.T2.ET.(*Token)
			qc = p.EP.(*ExchangePair).IC2
			feeRate = p.EP.(*ExchangePair).BinanceFeeRate2
		} else {
			bc = p.T1.ET.(*Token)
			qc = p.EP.(*ExchangePair).IC1
			feeRate = p.EP.(*ExchangePair).BinanceFeeRate1
		}
		wc = bc
	}
	return
}

func (ex *exchange) setOrderFeeRate(p *entity.Pair) error {
	if p.T1.Id.Symbol == p.T2.Id.Symbol {
		return nil
	}
	ep := p.EP.(*ExchangePair)
	if ep.HasIntermediaryCoin {

		bc := p.T1.ET.(*Token)
		qc := ep.IC1
		f1, err := ex.si.getTradeFee(bc.Coin, qc.Coin)
		if err != nil {
			return err
		}

		bc = p.T2.ET.(*Token)
		qc = ep.IC2
		f2, err := ex.si.getTradeFee(bc.Coin, qc.Coin)
		if err != nil {
			return err
		}

		ep.BinanceFeeRate1 = f1
		ep.BinanceFeeRate2 = f2
		return nil
	}

	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)
	f0, err := ex.si.getTradeFee(bc.Coin, qc.Coin)
	if err != nil {
		return err
	}
	ep.BinanceFeeRate1 = f0
	return err
}

func (ex *exchange) tokensEfa(t *entity.Token, p *entity.Pair,
	tokensEfa map[string]float64) (float64, error) {
	efa, ok := tokensEfa[t.Id.String()]
	if ok {
		return efa, nil
	}
	efa, _, err := ex.exchangeFeeAmount(t, p)
	if err != nil {
		return 0, err
	}
	tokensEfa[t.Id.String()] = efa
	return efa, nil
}
func (ex *exchange) exchangeFeeAmount(t *entity.Token, p *entity.Pair) (float64, float64, error) {
	var (
		qcDollar float64
	)

	if t.ET.(*Token).Coin == t.ET.(*Token).StableToken {
		qcDollar = 1
	} else {
		bc := t.ET.(*Token).Coin
		qc := t.ET.(*Token).StableToken
		qd, err := ex.si.getPrice(bc, qc)
		if err != nil {
			return 0, 0, err
		}
		qcDollar = qd
	}

	return p.ExchangeFee / qcDollar, qcDollar, nil
}
