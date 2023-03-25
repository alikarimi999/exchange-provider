package app

import (
	"exchange-provider/internal/entity"
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
	return p
}
