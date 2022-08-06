package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) AddPairs(ex *Exchange, pairs []*entity.Pair) (*entity.AddPairsResult, error) {

	return ex.AddPairs(pairs)
}

func (o *OrderUseCase) GetAllPairs(ex *Exchange) ([]*entity.Pair, error) {
	return o.setFee(ex.GetAllPairs()...), nil
}

func (o *OrderUseCase) GetPair(ex entity.Exchange, bc, qc *entity.Coin) (*entity.Pair, error) {
	p, err := ex.GetPair(bc, qc)
	if err != nil {
		return nil, err
	}

	return o.setFee(p)[0], nil
}

func (o *OrderUseCase) setFee(ps ...*entity.Pair) []*entity.Pair {
	f := o.fs.GetFee()
	for _, p := range ps {
		p.Fee = f
	}
	return ps
}

func (o *OrderUseCase) RemovePair(ex entity.Exchange, bc, qc *entity.Coin, force bool) error {

	if !force {
		f1 := &entity.Filter{
			Param:    "base_coin",
			Operator: entity.FilterOperatorEqual,
			Values:   []interface{}{bc.CoinId},
		}

		f2 := &entity.Filter{
			Param:    "base_chain",
			Operator: entity.FilterOperatorEqual,
			Values:   []interface{}{bc.ChainId},
		}

		f3 := &entity.Filter{
			Param:    "quote_coin",
			Operator: entity.FilterOperatorEqual,
			Values:   []interface{}{qc.CoinId},
		}

		f4 := &entity.Filter{
			Param:    "quote_chain",
			Operator: entity.FilterOperatorEqual,
			Values:   []interface{}{qc.ChainId},
		}

		t, err := o.totalPendingOrders(ex, f1, f2, f3, f4)
		if err != nil {
			return err
		}
		if t > 0 {
			return errors.Wrap(errors.ErrForbidden,
				errors.NewMesssage(fmt.Sprintf("there are %d pending orders with %s/%s", t, bc.String(), qc.String())))
		}

	}

	return ex.RemovePair(bc, qc)
}
