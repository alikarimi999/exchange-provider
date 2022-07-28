package kucoin

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (k *kucoinExchange) AddPairs(pairs []*entity.Pair) (*entity.AddPairsResult, error) {

	res := &entity.AddPairsResult{}

	ps := []*pair{}
	for _, p := range pairs {

		if k.exchangePairs.exists(p.BC.Coin, p.QC.Coin) {
			res.Existed = append(res.Existed, p)
			continue
		}

		ok, err := k.pls.support(p)
		if err != nil {
			return nil, err
		}
		if !ok {
			res.Failed = append(res.Failed, &entity.AddPairsErr{
				Pair: p,
				Err:  errors.Wrap(errors.ErrBadRequest, errors.New("pair not supported")),
			})
			continue
		}

		err = k.setInfos(p)
		if err != nil {
			res.Failed = append(res.Failed, &entity.AddPairsErr{
				Pair: p,
				Err:  err,
			})
			continue
		}

		ps = append(ps, fromEntity(p))
		res.Added = append(res.Added, p)
	}

	k.exchangePairs.add(ps)

	cs := map[string]*withdrawalCoin{}
	for _, p := range pairs {
		cs[p.BC.Id+p.BC.Chain.Id] = &withdrawalCoin{
			needChain: p.BC.SetChain,
		}

		cs[p.QC.Id+p.QC.Chain.Id] = &withdrawalCoin{
			needChain: p.QC.SetChain,
		}
	}

	k.withdrawalCoins.add(cs)

	return res, nil

}

func (k *kucoinExchange) GetPairs() []*entity.Pair {
	pairs := []*entity.Pair{}
	ps := k.exchangePairs.snapshot()
	for _, p := range ps {
		pairs = append(pairs, p.toEntity())
	}
	return pairs
}

func (k *kucoinExchange) Support(c1, c2 *entity.Coin) bool {
	return k.exchangePairs.exists(c1, c2)
}
