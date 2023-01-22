package app

import (
	"exchange-provider/internal/entity"
)

// func (o *OrderUseCase) GetAllPairs()

func (o *OrderUseCase) AddPairs(ex entity.Exchange, data interface{}) (*entity.AddPairsResult, error) {
	return ex.AddPairs(data)
}

func (o *OrderUseCase) RemovePair(ex entity.Exchange, t1, t2 *entity.Token, force bool) error {

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

	return ex.RemovePair(t1, t2)
}
