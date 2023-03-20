package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"sync"
	"time"
)

func (u *dex) Run() {
	u.l.Debug(fmt.Sprintf("%s.Run", u.Id()), "started")
}

func (u *dex) Name() string {
	return u.cfg.Name
}

func (u *dex) Id() string {
	return u.cfg.Name + "-" + u.cfg.Network
}

func (u *dex) Type() entity.ExType {
	return entity.EvmDEX
}
func (u *dex) Stop() {
	op := fmt.Sprintf("%s.Stop", u.Id())
	close(u.stopCh)
	u.stoppedAt = time.Now()
	u.l.Debug(string(op), "stopped")
}

func (u *dex) Configs() interface{} {
	u.cfg.Accounts, _ = u.wallet.AllAccounts(u.cfg.AccountCount)
	return u.cfg
}

func (u *dex) GetAllPairs() []*entity.Pair {
	agent := u.agent("GetAllPairs")

	ps := u.pairs.getAll()
	pairs := make([]*entity.Pair, len(ps))

	wg := sync.WaitGroup{}
	for i, p := range ps {
		wg.Add(1)
		go func(p types.Pair, i int) {
			defer wg.Done()
			newPair, err := u.PairWithPrice(p.T1, p.T2)
			if err != nil {
				u.l.Error(agent, err.Error())
				newPair = &p
			}

			pairs[i] = newPair.ToEntity(u.Id(), u.cfg.NativeToken, u.cfg.TokenStandard)
		}(p, i)
	}

	wg.Wait()
	return pairs
}

func (u *dex) Price(ps ...*entity.Pair) ([]*entity.Pair, error) {
	// if bc.ChainId != u.cfg.TokenStandard || qc.ChainId != u.cfg.TokenStandard {
	// 	return nil, fmt.Errorf("unexpected chain id %v and chain id %v", bc.ChainId, qc.ChainId)
	// }

	// p, err := u.pairs.get(bc.TokenId, qc.TokenId)
	// if err != nil {
	// 	return nil, err
	// }

	// p, err = u.PairWithPrice(p.T1, p.T2)
	// if err != nil {
	// 	return nil, err
	// }
	// return p.ToEntity(u.Id(), u.cfg.NativeToken, u.cfg.TokenStandard), nil
	return nil, nil
}

func (u *dex) Support(t1, t2 *entity.Token) bool {
	if t1.Standard != u.cfg.TokenStandard || t2.Standard != u.cfg.TokenStandard {
		return false
	}
	_, err := u.pairs.get(t1.Symbol, t2.Symbol)
	return err == nil
}

func (u *dex) RemovePair(t1, t2 *entity.Token) error {
	if t1.Standard != u.cfg.TokenStandard || t2.Standard != u.cfg.TokenStandard {
		return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}

	if p, err := u.pairs.get(t1.Symbol, t2.Symbol); err == nil {
		id := pairId(p.T1.Symbol, p.T2.Symbol)
		ps := u.v.GetStringSlice(fmt.Sprintf("%s.pairs", u.Id()))
		for i, p := range ps {
			if p == id {
				ps = append(ps[:i], ps[i+1:]...)
			}
		}
		u.v.Set(fmt.Sprintf("%s.pairs", u.Id()), ps)
		if err := u.v.WriteConfig(); err != nil {
			return err
		}
		return u.pairs.remove(t1.Symbol, t2.Symbol)
	}
	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}
