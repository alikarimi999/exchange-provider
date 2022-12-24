package dex

import (
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/dex/dto"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"os"
	"sync"
)

type tokens struct {
	Tokens []types.Token `json:"tokens"`
}

func (d *dex) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := d.agent("AddPairs")

	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	pwg := &sync.WaitGroup{}

	ps := d.v.GetStringSlice(fmt.Sprintf("%s.pairs", d.Id()))
	pMux := &sync.Mutex{}
	for _, dp := range req.Pairs {
		if d.pairs.exist(dp.T1, dp.T2) {
			d.l.Debug(agent, fmt.Sprintf("pair %s already exists", dp.String()))
			res.Existed = append(res.Existed, dp.String())
			continue
		}

		pwg.Add(1)
		go func(p *dto.Pair) {
			defer pwg.Done()
			if err := d.addPair(p.T1, p.T2); err != nil {
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
				return
			}
			res.Added = append(res.Added, entity.Pair{
				T1: &entity.PairCoin{
					Token: &entity.Token{TokenId: p.T1, ChainId: d.cfg.TokenStandard},
				},
				T2: &entity.PairCoin{
					Token: &entity.Token{TokenId: p.T2, ChainId: d.cfg.TokenStandard},
				},
			})
			pMux.Lock()
			ps = append(ps, p.String())
			pMux.Unlock()
		}(dp)

	}
	pwg.Wait()
	d.v.Set(fmt.Sprintf("%s.pairs", d.Id()), ps)
	if err := d.v.WriteConfig(); err != nil {
		d.l.Error(agent, err.Error())
	}
	return res, nil
}

func (d *dex) addPair(t1 string, t2 string) error {
	ts := &tokens{}

	b, err := os.ReadFile(d.cfg.TokensFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, ts); err != nil {
		return err
	}

	var btIsNative bool
	var qtIsNative bool

	if t1 == d.cfg.NativeToken {
		t1 = d.cfg.wrapNative()
		btIsNative = true
	}
	if t2 == d.cfg.NativeToken {
		t2 = d.cfg.wrapNative()
		qtIsNative = true
	}

	for _, t0 := range ts.Tokens {
		if t0.ChainId != int64(d.cfg.ChainId) {
			continue
		}

		if t0.Symbol == t1 {
			for _, t1 := range ts.Tokens {
				if t1.ChainId != int64(d.cfg.ChainId) {
					continue
				}
				if t1.Symbol == t2 {
					pair, err := d.Pair(t0, t1)
					if err != nil {
						return err
					}

					if btIsNative {
						pair.T1.Symbol = d.cfg.NativeToken
						pair.T1.Native = true
					} else if qtIsNative {
						pair.T2.Symbol = d.cfg.NativeToken
						pair.T2.Native = true
					}

					// check pair's tokens allowance
					adds, err := d.wallet.AllAddresses(d.cfg.AccountCount)
					if err != nil {
						return err
					}

					wg := &sync.WaitGroup{}
					var approveErr1 []error
					wg.Add(1)
					go func() {
						approveErr1 = d.am.InfinitApproves(pair.T1, d.cfg.Router, adds...)
						wg.Done()
					}()

					var approveErr2 []error
					wg.Add(1)
					go func() {
						approveErr2 = d.am.InfinitApproves(pair.T2, d.cfg.Router, adds...)
						wg.Done()
					}()

					wg.Wait()
					if len(approveErr1) > 0 {
						e := errors.New("")
						for _, err := range approveErr1 {
							e = errors.New(fmt.Sprintf("%s\n%s", e.Error(), err.Error()))
						}
						return e
					}

					if approveErr2 != nil {
						e := errors.New("")
						for _, err := range approveErr1 {
							e = errors.New(fmt.Sprintf("%s\n%s", e.Error(), err.Error()))
						}
						return e
					}

					d.pairs.add(*pair)
					d.tokens.add(pair.T1, pair.T2)
					return nil
				}

			}
			return errors.New(
				fmt.Sprintf("token `%s` for chain `%d` did not found in tokens list", t2, d.cfg.ChainId))
		}
	}
	return errors.New(
		fmt.Sprintf("token `%s` for chain `%d` did not found in tokens list", t1, d.cfg.ChainId))

}
