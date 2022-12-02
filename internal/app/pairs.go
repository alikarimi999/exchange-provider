package app

import (
	"exchange-provider/internal/entity"
)

func (o *OrderUseCase) AddPairs(ex entity.Exchange, data interface{}) (*entity.AddPairsResult, error) {
	return ex.AddPairs(data)
}

func (o *OrderUseCase) GetAllPairsByExchange(ex entity.Exchange) ([]*entity.Pair, error) {

	ps := ex.GetAllPairs()
	// set spread_rate
	for _, p := range ps {
		p.C1.MinDeposit, p.C2.MinDeposit = o.pc.PairMinDeposit(p.C1.Coin.String(), p.C2.Coin.String())
		p.SpreadRate = o.pc.GetPairSpread(p.C1.Coin, p.C2.Coin)
	}

	return ps, nil
}

func (o *OrderUseCase) RemovePair(ex entity.Exchange, bc, qc *entity.Coin, force bool) error {

	// if !force {
	// 	f1 := &entity.Filter{
	// 		Param:    "base_coin",
	// 		Operator: entity.FilterOperatorEqual,
	// 		Values:   []interface{}{bc.CoinId},
	// 	}

	// 	f2 := &entity.Filter{
	// 		Param:    "base_chain",
	// 		Operator: entity.FilterOperatorEqual,
	// 		Values:   []interface{}{bc.ChainId},
	// 	}

	// 	f3 := &entity.Filter{
	// 		Param:    "quote_coin",
	// 		Operator: entity.FilterOperatorEqual,
	// 		Values:   []interface{}{qc.CoinId},
	// 	}

	// 	f4 := &entity.Filter{
	// 		Param:    "quote_chain",
	// 		Operator: entity.FilterOperatorEqual,
	// 		Values:   []interface{}{qc.ChainId},
	// 	}

	// 	t, err := o.totalPendingOrders(ex, f1, f2, f3, f4)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if t > 0 {
	// 		return errors.Wrap(errors.ErrForbidden,
	// 			errors.NewMesssage(fmt.Sprintf("there are %d pending orders with %s/%s", t, bc.String(), qc.String())))
	// 	}

	// }

	return ex.RemovePair(bc, qc)
}
