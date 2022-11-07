package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

func (u *dex) Run(wg *sync.WaitGroup) {
	defer wg.Done()

}

func (u *dex) Type() entity.ExType {
	return entity.DEX
}
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
		go func(p types.Pair) {
			defer wg.Done()
			newPair, err := u.PairWithPrice(p.BT, p.QT)
			if err != nil {
				u.l.Error(agent, err.Error())
				return
			}

			pairs = append(pairs, newPair.ToEntity(u.cfg.NativeToken, u.cfg.TokenStandard, u.cfg.BlockTime))
		}(p)
	}

	wg.Wait()
	return pairs
}

func (u *dex) StartAgain() (*entity.StartAgainResult, error) {
	agent := u.agent("StartAgain")
	u.l.Debug(agent, "start again")

	for _, p := range u.cfg.Providers {
		if err := p.Ping(); err != nil {
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

	p, err = u.PairWithPrice(p.BT, p.QT)
	if err != nil {
		return nil, err
	}
	return p.ToEntity(u.cfg.NativeToken, u.cfg.TokenStandard, u.cfg.BlockTime), nil
}

func (u *dex) Support(bc, qc *entity.Coin) bool {
	if bc.ChainId != u.cfg.TokenStandard || qc.ChainId != u.cfg.TokenStandard {
		return false
	}
	_, err := u.pairs.get(bc.CoinId, qc.CoinId)
	return err == nil
}

func (u *dex) RemovePair(bc, qc *entity.Coin) error {
	if bc.ChainId != u.cfg.TokenStandard || qc.ChainId != u.cfg.TokenStandard {
		return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}

	if u.pairs.existsExactly(bc.CoinId, qc.CoinId) {
		id := pairId(bc.CoinId, qc.CoinId)
		delete(u.v.Get(fmt.Sprintf("%s.pairs", u.NID())).(map[string]interface{}), strings.ToLower(id))
		if err := u.v.WriteConfig(); err != nil {
			return err
		}
		return u.pairs.remove(bc.CoinId, qc.CoinId)
	}
	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}
