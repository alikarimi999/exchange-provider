package app

import "exchange-provider/internal/entity"

func (o *OrderUseCase) GetMinPairDeposit(c1, c2 string) (minC1, minC2 float64) {
	return o.pc.PairMinDeposit(c1, c2)
}

func (o *OrderUseCase) ChangeMinDeposit(s *entity.PairMinDeposit) error {
	return o.pc.ChangeMinDeposit(s)
}

func (o *OrderUseCase) AllMinDeposit() []*entity.PairMinDeposit {
	return o.pc.AllMinDeposit()
}
