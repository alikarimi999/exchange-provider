package app

import (
	"exchange-provider/internal/entity"
	"strconv"
)

func (o *OrderUseCase) GetDefaultSpread() string {
	return o.pc.GetDefaultSpread()
}

func (o *OrderUseCase) ChangeDefaultSpread(s float64) error {
	return o.pc.ChangeDefaultSpread(s)
}

func (o *OrderUseCase) GetPairSpread(bc, qc *entity.Token) string {
	return o.pc.GetPairSpread(bc, qc)
}

func (o *OrderUseCase) ChangePairSpread(bc, qc *entity.Token, s float64) error {
	return o.pc.ChangePairSpread(bc, qc, s)
}

func (o *OrderUseCase) GetAllPairsSpread() map[string]float64 {
	return o.pc.GetAllPairsSpread()
}

func (o *OrderUseCase) ApplySpread(p *entity.Pair) *entity.Pair {
	rate := o.pc.GetPairSpread(p.T1.Token, p.T2.Token)

	r, _ := strconv.ParseFloat(rate, 64)
	p1, _ := strconv.ParseFloat(p.Price1, 64)
	p2, _ := strconv.ParseFloat(p.Price2, 64)

	p.Price1 = strconv.FormatFloat(p1*(1+r), 'f', 6, 64)
	p.Price2 = strconv.FormatFloat(p2*(1+r), 'f', 6, 64)

	return p
}
