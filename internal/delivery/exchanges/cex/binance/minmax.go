package binance

import (
	"exchange-provider/internal/entity"
	"math"
	"math/big"
	"strconv"

	"github.com/adshao/go-binance/v2"
)

func (k *exchange) minAndMax(p *entity.Pair, p0, p1,
	bcEFA, qcEFA float64, s0, s1 binance.Symbol) error {

	price := k.calcPrice(p0, p1, p.T1.Id, p.T2.Id, p)
	spread, err := k.spread(0, p, price)
	if err != nil {
		return err
	}

	if p.EP.(*ExchangePair).HasIntermediaryCoin {
		bc := p.T1.ET.(*Token)
		qc := p.EP.(*ExchangePair).IC1
		minBC0 := bcMin(p, s0, bc, qc, p0, spread)
		minQC0 := qcMin(p, s0, bc, qc, p0, spread)

		bc = p.T2.ET.(*Token)
		qc = p.EP.(*ExchangePair).IC2
		minBC1 := bcMin(p, s0, bc, qc, p1, spread)
		minQC1 := qcMin(p, s0, bc, qc, p1, spread)

		var minBC, minQC float64
		if minBC0*p0 >= minQC1 {
			minBC = minBC0
		} else {
			minBC = minQC1 / p0
		}

		if minBC1*p1 >= minQC0 {
			minQC = minBC1
		} else {
			minQC = minQC0 / p1
		}
		minBC = minBC + bcEFA
		minBC = minBC + (minBC * p.FeeRate2)
		minQC = minQC + qcEFA
		minQC = minQC + (minQC * p.FeeRate1)
		p.T1.Min, p.T2.Min = min(p, minBC, minQC)
		return nil
	}
	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)

	bcMa := bcMin(p, s0, bc, qc, price, spread) + bcEFA
	bcMa = bcMa + (bcMa * p.FeeRate2)

	qcMa := qcMin(p, s0, bc, qc, price, spread) + qcEFA
	qcMa = qcMa + (qcMa * p.FeeRate1)
	p.T1.Min, p.T2.Min = min(p, bcMa, qcMa)
	return nil
}

func min(p *entity.Pair, min0, min1 float64) (float64, float64) {
	m0, _ := strconv.ParseFloat(trim(big.NewFloat(min0+(min0*0.5)).Text('f', 12), p.T1.ET.(*Token).OrderPrecision), 64)
	m1, _ := strconv.ParseFloat(trim(big.NewFloat(min1+(min1*0.5)).Text('f', 12), p.T2.ET.(*Token).OrderPrecision), 64)
	return m0, m1
}

func bcMin(p *entity.Pair, s binance.Symbol, bc, qc *Token, price, spread float64) float64 {
	ps := (price - (price * spread))
	min0 := getNotionalMin(s) / ps
	min1 := (qc.MinWithdrawalFee + qc.MinWithdrawalSize) / ps
	min2 := getMinLS(s)
	return math.Max(math.Max(min0, min1), min2)
}

func qcMin(p *entity.Pair, s binance.Symbol, bc, qc *Token, price, spread float64) float64 {
	min0 := getNotionalMin(s)
	min1 := (bc.MinWithdrawalFee + bc.MinWithdrawalSize) * (price + (price * spread))
	return math.Max(min0, min1)
}

func getNotionalMin(s binance.Symbol) float64 {
	var min float64
	for _, fs := range s.Filters {
		if fs["filterType"] == "NOTIONAL" {
			minS, ok := fs["minNotional"].(string)
			if ok {
				min, _ = strconv.ParseFloat(minS, 64)
				break
			}
		}
	}
	return min
}

func getMinLS(s binance.Symbol) float64 {
	var min float64
	for _, fs := range s.Filters {
		if fs["filterType"] == "LOT_SIZE" {
			minS, ok := fs["minQty"].(string)
			if ok {
				min, _ = strconv.ParseFloat(minS, 64)
				break
			}
		}
	}
	return min
}
