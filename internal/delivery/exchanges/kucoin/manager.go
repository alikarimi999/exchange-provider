package kucoin

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (k *kucoinExchange) AddPairs(pairs []*entity.Pair) (*entity.AddPairsResult, error) {

	res := &entity.AddPairsResult{}

	ps := []*pair{}
	cs := map[string]*withdrawalCoin{}

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
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p,
				Err:  errors.Wrap(errors.ErrBadRequest, errors.New("pair not supported")),
			})
			continue
		}

		err = k.setInfos(p)
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p,
				Err:  err,
			})
			continue
		}

		pa := fromEntity(p)
		ps = append(ps, pa)
		res.Added = append(res.Added, p)
		cs[p.BC.CoinId+p.BC.ChainId] = &withdrawalCoin{
			precision: p.BC.WithdrawalPrecision,
			needChain: p.BC.SetChain,
		}

		cs[p.QC.CoinId+p.QC.ChainId] = &withdrawalCoin{
			precision: p.QC.WithdrawalPrecision,
			needChain: p.QC.SetChain,
		}
	}

	k.exchangePairs.add(ps)
	k.withdrawalCoins.add(cs)

	return res, nil

}

func (k *kucoinExchange) GetAllPairs() []*entity.Pair {
	pairs := []*entity.Pair{}
	ps := k.exchangePairs.snapshot()

	for _, p := range ps {
		pe := p.toEntity()
		k.setPrice(pe)
		k.setOrderFeeRate(pe)

		pairs = append(pairs, pe)
	}
	return pairs
}

func (k *kucoinExchange) GetPair(bc, qc *entity.Coin) (*entity.Pair, error) {
	p, err := k.exchangePairs.get(bc, qc)
	if err != nil {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}
	pe := p.toEntity()
	if err := k.setPrice(pe); err != nil {
		return nil, err
	}
	if err := k.setOrderFeeRate(pe); err != nil {
		return nil, err
	}
	return pe, nil
}

func (k *kucoinExchange) RemovePair(bc, qc *entity.Coin) error {
	if k.exchangePairs.exists(bc, qc) {
		k.exchangePairs.remove(bc, qc)
		return nil
	}
	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}

func (k *kucoinExchange) Support(bc, qc *entity.Coin) bool {
	return k.exchangePairs.exists(bc, qc)
}
