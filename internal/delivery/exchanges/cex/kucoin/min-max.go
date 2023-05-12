package kucoin

// func (k *kucoinExchange) setMinMax(p *entity.Pair, price float64) error {

// 	et1 := p.T1.ET.(*Token)
// 	et2 := p.T2.ET.(*Token)
// 	var bcDollar, qcDollar float64
// 	if et2.Currency == usdt {
// 		qcDollar = 1
// 		bd, err := k.usdtTicker(et1.Currency)
// 		if err != nil {
// 			return err
// 		}
// 		bcDollar = bd
// 	} else {
// 		qd, err := k.usdtTicker(et2.Currency)
// 		if err != nil {
// 			return err
// 		}
// 		qcDollar = qd
// 		if et1.Currency == usdt {
// 			bcDollar = 1
// 		} else {
// 			bd, err := k.usdtTicker(et1.Currency)
// 			if err != nil {
// 				return err
// 			}
// 			bcDollar = bd
// 		}
// 	}

// 	ls := k.st.Levels()
// 	var l uint
// 	for _, lv := range ls {
// 		if lv > l {
// 			l = lv
// 		}
// 	}

// 	s, _ := k.spread(l, p, price)
// 	min1 := et2.MinWithdrawalFee + et2.MinWithdrawalSize
// 	min1 = (min1 + (min1 * p.FeeRate2) + (p.ExchangeFee / qcDollar))
// 	min1 = min1 / (price - (price * s))

// 	s, _ = k.spread(0, p, price)
// 	min2 := (et1.MinWithdrawalFee + et1.MinWithdrawalSize) * (price + (price * s))
// 	min2 = (p.ExchangeFee / bcDollar) * (min2 * p.FeeRate1) * price

// 	p.T1.Min = min1
// 	p.T2.Min = min2
// 	return nil
// }
