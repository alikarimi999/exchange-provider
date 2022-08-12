package kucoin

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"strings"
)

func (k *kucoinExchange) AddPairs(pairs []*entity.Pair) (*entity.AddPairsResult, error) {
	agent := fmt.Sprintf("%s.AddPairs", k.NID())
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

		k.v.Set(fmt.Sprintf("%s.pairs.%s", k.NID(), pa.Id), pa)
		if err := k.v.WriteConfig(); err != nil {
			k.l.Error(agent, err.Error())
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p,
				Err:  err,
			})
			continue
		}
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
		id := bc.CoinId + bc.ChainId + qc.CoinId + qc.ChainId
		delete(k.v.Get(fmt.Sprintf("%s.pairs", k.NID())).(map[string]interface{}), strings.ToLower(id))
		if err := k.v.WriteConfig(); err != nil {
			return err
		}

		k.exchangePairs.remove(id)
		return nil
	}
	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}

func (k *kucoinExchange) Support(bc, qc *entity.Coin) bool {
	return k.exchangePairs.exists(bc, qc)
}

func (k *kucoinExchange) StartAgain() (*entity.StartAgainResult, error) {
	op := errors.Op(fmt.Sprintf("%s.StartAgain", k.NID()))
	k.stopCh = make(chan struct{})

	k.l.Debug(string(op), "starting again")
	k.l.Debug(string(op), fmt.Sprintf("stopped at %s", k.stopedAt.Format("2006-01-02 15:04:05")))

	// check if ping is successful
	if err := k.ping(); err != nil {
		return nil, errors.Wrap(string(op), err)
	}
	k.l.Debug(string(op), fmt.Sprintf("ping was successful"))

	k.l.Debug(string(op), "downloading pairs list from kucoin")

	// download pairs list from kucoin again
	if err := k.pls.download(); err != nil {
		return nil, errors.Wrap(string(op), err)
	}

	res := &entity.StartAgainResult{}
	// check if current pairs are still supported by kucoin
	ps := k.exchangePairs.snapshot()
	k.exchangePairs.purge()
	cs := k.withdrawalCoins.snapshot()
	k.withdrawalCoins.purge()
	newPs := []*pair{}
	newCs := map[string]*withdrawalCoin{}
	for _, p := range ps {
		pe := p.toEntity()
		ok, err := k.pls.support(pe)
		if err != nil {
			return nil, errors.Wrap(string(op), err)
		}

		if !ok {
			delete(k.v.Get(fmt.Sprintf("%s.pairs", k.NID())).(map[string]interface{}), strings.ToLower(p.Id))
			if err := k.v.WriteConfig(); err != nil {
				k.l.Error(string(op), err.Error())
			}
			res.Removed = append(res.Removed, &entity.PairsErr{
				Pair: pe,
				Err:  fmt.Errorf("pair is not supported by kucoin anymore so it will be removed"),
			})
			continue
		}

		if err := k.setInfos(pe); err != nil {
			res.Removed = append(res.Removed, &entity.PairsErr{
				Pair: pe,
				Err:  fmt.Errorf("retrieving infos for pair failed due to error ( %s ) so it well be removed", err.Error()),
			})
			continue
		}
		newPs = append(newPs, fromEntity(pe))
		newCs[pe.BC.CoinId+pe.BC.ChainId] = &withdrawalCoin{
			precision: cs[pe.BC.CoinId+pe.BC.ChainId].precision,
			needChain: pe.BC.SetChain,
		}

		newCs[pe.QC.CoinId+pe.QC.ChainId] = &withdrawalCoin{
			precision: cs[pe.QC.CoinId+pe.QC.ChainId].precision,
			needChain: pe.QC.SetChain,
		}

	}

	k.exchangePairs.add(newPs)
	k.withdrawalCoins.add(newCs)

	k.l.Info(string(op), fmt.Sprintf("%d pairs were added", len(newPs)))
	k.l.Info(string(op), fmt.Sprintf("%d pairs were removed", len(ps)-len(newPs)))
	k.l.Info(string(op), fmt.Sprintf("exchange %s started again successfully", k.NID()))

	return res, nil
}
