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

func (u *dex) Name() string {
	return u.cfg.Name
}

func (u *dex) Id() string {
	return u.cfg.Name + "-" + u.cfg.Network
}

func (u *dex) Type() entity.ExType {
	return entity.DEX
}
func (u *dex) Stop() {
	op := fmt.Sprintf("%s.Stop", u.Id())
	close(u.stopCh)
	u.stoppedAt = time.Now()
	u.l.Debug(string(op), "stopped")
}

func (u *dex) Configs() interface{} {
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
			newPair, err := u.PairWithPrice(p.T1, p.T2)
			if err != nil {
				u.l.Error(agent, err.Error())
				return
			}

			pairs = append(pairs, newPair.ToEntity(u.cfg.NativeToken, u.cfg.chainId))
		}(p)
	}

	wg.Wait()
	return pairs
}

func (u *dex) GetPair(bc, qc *entity.Token) (*entity.Pair, error) {
	if bc.ChainId != u.cfg.chainId || qc.ChainId != u.cfg.chainId {
		return nil, fmt.Errorf("unexpected chain id %v and chain id %v", bc.ChainId, qc.ChainId)
	}

	p, err := u.pairs.get(bc.TokenId, qc.TokenId)
	if err != nil {
		return nil, err
	}

	p, err = u.PairWithPrice(p.T1, p.T2)
	if err != nil {
		return nil, err
	}
	return p.ToEntity(u.cfg.NativeToken, u.cfg.chainId), nil
}

func (u *dex) Support(bc, qc *entity.Token) bool {
	if bc.ChainId != u.cfg.chainId || qc.ChainId != u.cfg.chainId {
		return false
	}
	_, err := u.pairs.get(bc.TokenId, qc.TokenId)
	return err == nil
}

func (u *dex) RemovePair(t1, t2 *entity.Token) error {
	if t1.ChainId != u.cfg.chainId || t2.ChainId != u.cfg.chainId {
		return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}

	if p, err := u.pairs.get(t1.TokenId, t2.TokenId); err == nil {
		id := pairId(p.T1.Symbol, p.T2.Symbol)
		delete(u.v.Get(fmt.Sprintf("%s.pairs", u.Id())).(map[string]interface{}), strings.ToLower(id))
		if err := u.v.WriteConfig(); err != nil {
			return err
		}
		return u.pairs.remove(t1.TokenId, t2.TokenId)
	}
	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}
