package uniswapv3

import (
	"exchange-provider/internal/entity"
	"fmt"
	"sync"
	"time"
)

func (u *dex) Stop() {
	op := fmt.Sprintf("%s.Stop", u.NID())
	close(u.stopCh)
	u.stoppedAt = time.Now()
	u.l.Debug(string(op), "stopped")
}

func (u *dex) Configs() interface{} {
	u.cfg.Id = u.NID()
	u.cfg.Accounts, _ = u.wallet.AllAccounts()
	return u.cfg
}

func (u *dex) GetAllPairs() []*entity.Pair {
	agent := u.agent("GetAllPairs")

	ps := u.pairs.getAll()
	pairs := []*entity.Pair{}

	wg := sync.WaitGroup{}
	for _, p := range ps {
		wg.Add(1)
		go func(p pair) {
			defer wg.Done()
			newPair, err := u.setBestPrice(p.BT, p.QT)
			if err != nil {
				u.l.Error(agent, err.Error())
				return
			}

			pairs = append(pairs, newPair.ToEntity(u))
		}(p)
	}

	wg.Wait()
	return pairs
}

func (u *dex) StartAgain() (*entity.StartAgainResult, error) {
	agent := u.agent("StartAgain")
	u.l.Debug(agent, "start again")

	for _, p := range u.cfg.Providers {
		if err := p.ping(); err != nil {
			return nil, err
		}
	}

	u.stopCh = make(chan struct{})
	return &entity.StartAgainResult{}, nil
}

func (u *dex) GetPair(bc, qc *entity.Coin) (*entity.Pair, error) {
	if bc.ChainId != u.cfg.TokenStandard || qc.ChainId != u.cfg.TokenStandard {
		return nil, fmt.Errorf("unexpected chain id %v and chain id %v", bc.ChainId, qc.ChainId)
	}

	p, err := u.pairs.get(bc.CoinId, qc.CoinId)
	if err != nil {
		return nil, err
	}

	p, err = u.setBestPrice(p.BT, p.QT)
	if err != nil {
		return nil, err
	}
	return p.ToEntity(u), nil
}
