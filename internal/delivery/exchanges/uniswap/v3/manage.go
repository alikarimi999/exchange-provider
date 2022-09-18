package uniswapv3

import (
	"fmt"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"sync"
	"time"
)

func (u *UniSwapV3) Stop() {
	op := fmt.Sprintf("%s.Stop", u.NID())
	close(u.stopCh)
	u.stoppedAt = time.Now()
	u.l.Debug(string(op), "stopped")
}

func (u *UniSwapV3) Configs() interface{} {
	u.cfg.Id = u.accountId
	u.cfg.Accounts, _ = u.wallet.AllAccounts()
	u.cfg.DefaultProvider = u.provider.URL
	u.cfg.BackupProviders = u.backupProvidersURL
	return u.cfg
}

func (u *UniSwapV3) GetAllPairs() []*entity.Pair {
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

func (u *UniSwapV3) StartAgain() (*entity.StartAgainResult, error) {
	agent := u.agent("StartAgain")
	u.l.Debug(agent, "start again")

	if err := u.pingProvider(); err != nil {
		return nil, errors.Wrap(errors.Op(agent), err)
	}

	u.stopCh = make(chan struct{})
	return &entity.StartAgainResult{}, nil
}

func (u *UniSwapV3) GetPair(bc, qc *entity.Coin) (*entity.Pair, error) {
	if bc.ChainId != chainId || qc.ChainId != chainId {
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
