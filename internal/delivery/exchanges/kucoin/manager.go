package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
)

func (k *kucoinExchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := fmt.Sprintf("%s.AddPairs", k.NID())

	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(agent, errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	cs := map[string]*kuCoin{}
	ps := []*pair{}
	for _, p := range req.Pairs {
		ps = append(ps, fromDto(p))
	}

	aps := []*pair{}
	for _, p := range ps {

		if k.exchangePairs.exists(p.BC, p.QC) {
			res.Existed = append(res.Existed, p.String())
			continue
		}

		ok, err := k.pls.support(p)
		if err != nil {
			return nil, err
		}
		if !ok {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  errors.Wrap(errors.ErrBadRequest, errors.New("pair not supported")),
			})
			continue
		}

		if err := k.setInfos(p); err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  err,
			})
			continue
		}

		k.v.Set(fmt.Sprintf("%s.pairs.%s", k.NID(), p.Id), p)
		if err := k.v.WriteConfig(); err != nil {
			k.l.Error(agent, err.Error())
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  err,
			})
			continue
		}
		aps = append(aps, p)
		res.Added = append(res.Added, entity.Pair{
			BC: &entity.PairCoin{
				Coin: &entity.Coin{CoinId: p.BC.CoinId, ChainId: p.BC.ChainId},
			},
			QC: &entity.PairCoin{
				Coin: &entity.Coin{CoinId: p.QC.CoinId, ChainId: p.QC.ChainId},
			},
		})
		cs[p.BC.CoinId+p.BC.ChainId] = p.BC.snapshot()

		cs[p.QC.CoinId+p.QC.ChainId] = p.QC.snapshot()
	}

	k.exchangePairs.add(aps...)
	k.supportedCoins.add(cs)

	return res, nil

}

func (k *kucoinExchange) GetAllPairs() []*entity.Pair {
	agent := fmt.Sprintf("%s.GetAllPairs", k.NID())

	pairs := []*entity.Pair{}
	ps := k.exchangePairs.snapshot()

	for _, p := range ps {
		pe := p.toEntity()
		if err := k.setPrice(pe); err != nil {
			k.l.Error(agent, err.Error())
			continue
		}
		if err := k.setOrderFeeRate(pe); err != nil {
			k.l.Error(agent, err.Error())
			continue
		}

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
	if k.exchangePairs.exists(coinFromEntity(bc), coinFromEntity(qc)) {
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
	return k.exchangePairs.exists(coinFromEntity(bc), coinFromEntity(qc))
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
	k.l.Debug(string(op), "ping was successful")

	k.l.Debug(string(op), "downloading pairs list from kucoin")

	// download pairs list from kucoin again
	if err := k.pls.download(); err != nil {
		return nil, errors.Wrap(string(op), err)
	}

	res := &entity.StartAgainResult{}
	// check if current pairs are still supported by kucoin
	ps := k.exchangePairs.snapshot()
	k.exchangePairs.purge()
	cs := k.supportedCoins.snapshot()
	k.supportedCoins.purge()
	newPs := []*pair{}
	newCs := map[string]*kuCoin{}
	for _, p := range ps {
		ok, err := k.pls.support(p)
		if err != nil {
			return nil, errors.Wrap(string(op), err)
		}

		if !ok {
			delete(k.v.Get(fmt.Sprintf("%s.pairs", k.NID())).(map[string]interface{}), strings.ToLower(p.Id))
			if err := k.v.WriteConfig(); err != nil {
				k.l.Error(string(op), err.Error())
			}
			res.Removed = append(res.Removed, &entity.PairsErr{
				Pair: p.String(),
				Err:  fmt.Errorf("pair is not supported by kucoin anymore so it will be removed"),
			})
			continue
		}

		if err := k.setInfos(p); err != nil {
			res.Removed = append(res.Removed, &entity.PairsErr{
				Pair: p.String(),
				Err:  fmt.Errorf("retrieving infos for pair failed due to error ( %s ) so it well be removed", err.Error()),
			})
			continue
		}
		newPs = append(newPs, p)
		newCs[p.BC.CoinId+p.BC.ChainId] = cs[p.BC.CoinId+p.BC.ChainId].snapshot()
		newCs[p.QC.CoinId+p.QC.ChainId] = cs[p.QC.CoinId+p.QC.ChainId].snapshot()

	}

	k.exchangePairs.add(newPs...)
	k.supportedCoins.add(newCs)

	k.l.Info(string(op), fmt.Sprintf("%d pairs were added", len(newPs)))
	k.l.Info(string(op), fmt.Sprintf("%d pairs were removed", len(ps)-len(newPs)))
	k.l.Info(string(op), fmt.Sprintf("exchange %s started again successfully", k.NID()))

	return res, nil
}
