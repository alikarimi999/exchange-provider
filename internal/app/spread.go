package app

import (
	"order_service/internal/entity"
	"strconv"
)

func (o *OrderUseCase) GetDefaultSpread() string {
	return o.pc.GetDefaultSpread()
}

func (o *OrderUseCase) ChangeDefaultSpread(s float64) error {
	return o.pc.ChangeDefaultSpread(s)
}

func (o *OrderUseCase) GetPairSpread(bc, qc *entity.Coin) string {
	return o.pc.GetPairSpread(bc, qc)
}

func (o *OrderUseCase) ChangePairSpread(bc, qc *entity.Coin, s float64) error {
	return o.pc.ChangePairSpread(bc, qc, s)
}

func (o *OrderUseCase) GetAllPairsSpread() map[string]float64 {
	return o.pc.GetAllPairsSpread()
}

func (o *OrderUseCase) ApplySpread(p *entity.Pair) *entity.Pair {
	rate := o.pc.GetPairSpread(p.BC.Coin, p.QC.Coin)

	r, _ := strconv.ParseFloat(rate, 64)
	bestAsk, _ := strconv.ParseFloat(p.BestAsk, 64)
	bestBid, _ := strconv.ParseFloat(p.BestBid, 64)

	p.BestAsk = strconv.FormatFloat(bestAsk*(1+r), 'f', 6, 64)
	p.BestBid = strconv.FormatFloat(bestBid*(1-r), 'f', 6, 64)

	return p
}
