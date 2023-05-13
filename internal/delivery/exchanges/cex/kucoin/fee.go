package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/internal/entity"
)

func getBcQcWcFeeRate(o *types.Order, p *entity.Pair,
	i int) (bc, qc, wc *Token) {
	var kucoinFeeRate float64
	if i == 0 {
		if len(o.Swaps) == 2 {
			if o.Swaps[0].In.String() == p.T1.Id.String() {
				bc = p.T1.ET.(*Token)
				qc = p.EP.(*ExchangePair).IC1
				kucoinFeeRate = p.EP.(*ExchangePair).KucoinFeeRate1
			} else {
				bc = p.T2.ET.(*Token)
				qc = p.EP.(*ExchangePair).IC2
				kucoinFeeRate = p.EP.(*ExchangePair).KucoinFeeRate2
			}
		} else {
			bc = p.T1.ET.(*Token)
			qc = p.T2.ET.(*Token)
			if o.Swaps[0].Side == sellSide {
				wc = qc
			} else {
				wc = bc
			}
			kucoinFeeRate = p.EP.(*ExchangePair).KucoinFeeRate1
		}
		kucoinFeeRate = kucoinFeeRate + (kucoinFeeRate * 0.001)
		if o.Swaps[0].Side == buySide {
			o.Swaps[0].InAmountRequested = o.Deposit.Amount - (o.Deposit.Amount * kucoinFeeRate)
		} else {
			o.Swaps[0].InAmountRequested = o.Deposit.Amount
		}
	} else {
		o.Swaps[1].InAmountRequested = o.Swaps[0].OutAmount - o.Swaps[0].KucoinFee
		if o.Swaps[1].Out.String() == p.T2.String() {
			bc = p.T2.ET.(*Token)
			qc = p.EP.(*ExchangePair).IC2
			kucoinFeeRate = p.EP.(*ExchangePair).KucoinFeeRate2
		} else {
			bc = p.T1.ET.(*Token)
			qc = p.EP.(*ExchangePair).IC1
			kucoinFeeRate = p.EP.(*ExchangePair).KucoinFeeRate1
		}

		kucoinFeeRate = kucoinFeeRate + (kucoinFeeRate * 0.001)
		feeAmount := o.Swaps[1].InAmountRequested * kucoinFeeRate
		o.Swaps[1].InAmountRequested = o.Swaps[1].InAmountRequested - feeAmount
		wc = bc
	}
	return
}

func (k *kucoinExchange) exchangeFeeAmount(out *entity.Token, p *entity.Pair) (float64, error) {
	var (
		qcDollar float64
	)

	if out.ET.(*Token).Currency == out.ET.(*Token).StableToken {
		qcDollar = 1
	} else {
		qd, err := k.stableTicker(out.ET.(*Token).Currency, out.ET.(*Token).StableToken)
		if err != nil {
			return 0, err
		}
		qcDollar = qd
	}

	return p.ExchangeFee / qcDollar, nil
}
