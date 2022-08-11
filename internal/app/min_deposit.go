package app

import "order_service/internal/entity"

func (o *OrderUseCase) GetMinPairDeposit(bc, qc *entity.Coin) (minBc, minQc float64) {
	return o.sr.PairMinDeposit(bc, qc)
}

func (o *OrderUseCase) ChangeMinDeposit(bc, qc *entity.Coin, minBc, minQc float64) error {
	return o.sr.ChangeMinDeposit(bc, qc, minBc, minQc)
}

func (o *OrderUseCase) AllMinDeposit() []*entity.PairMinDeposit {
	return o.sr.AllMinDeposit()
}
