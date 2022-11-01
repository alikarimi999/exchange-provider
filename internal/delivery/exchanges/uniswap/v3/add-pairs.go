package uniswapv3

import (
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/uniswap/v3/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"os"
	"sync"
)

type tokens struct {
	Tokens []Token `json:"tokens"`
}

func (d *dex) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := d.agent("AddPairs")

	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	pwg := &sync.WaitGroup{}

	ps := d.v.GetStringSlice(fmt.Sprintf("%s.pairs", d.NID()))
	for _, dp := range req.Pairs {
		if d.pairs.exist(dp.BT, dp.QT) {
			d.l.Debug(agent, fmt.Sprintf("pair %s already exists", dp.String()))
			res.Existed = append(res.Existed, dp.String())
			continue
		}

		pwg.Add(1)
		go func(p *dto.Pair) {
			defer pwg.Done()
			if err := d.addPair(p.BT, p.QT); err != nil {
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
				return
			}
			res.Added = append(res.Added, entity.Pair{
				BC: &entity.PairCoin{
					Coin: &entity.Coin{CoinId: p.BT, ChainId: d.cfg.TokenStandard},
				},
				QC: &entity.PairCoin{
					Coin: &entity.Coin{CoinId: p.QT, ChainId: d.cfg.TokenStandard},
				},
			})
			ps = append(ps, p.String())
		}(dp)

	}
	pwg.Wait()
	d.v.Set(fmt.Sprintf("%s.pairs", d.NID()), ps)
	if err := d.v.WriteConfig(); err != nil {
		d.l.Error(agent, err.Error())
	}
	return res, nil
}

func (d *dex) addPair(bt string, qt string) error {
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

	if bt == d.cfg.NativeToken {
		bt = d.cfg.wrapNative()
		btIsNative = true
	}
	if qt == d.cfg.NativeToken {
		qt = d.cfg.wrapNative()
		qtIsNative = true
	}

	for _, t0 := range ts.Tokens {
		if t0.ChainId != int64(d.cfg.ChianId) {
			continue
		}

		if t0.Symbol == bt {
			for _, t1 := range ts.Tokens {
				if t1.ChainId != int64(d.cfg.ChianId) {
					continue
				}
				if t1.Symbol == qt {
					pair, err := d.pairWithPrice(t0, t1)
					if err != nil {
						return err
					}

					if btIsNative {
						pair.BT.Symbol = d.cfg.NativeToken
						pair.BT.Native = true
					} else if qtIsNative {
						pair.QT.Symbol = d.cfg.NativeToken
						pair.QT.Native = true
					}

					// check pair's tokens allowance
					adds, err := d.wallet.AllAddresses()
					if err != nil {
						return err
					}

					wg := &sync.WaitGroup{}
					var approveErr1 []error
					wg.Add(1)
					go func() {
						approveErr1 = d.am.infinitApproves(pair.BT, d.cfg.Router, adds...)
						wg.Done()
					}()

					var approveErr2 []error
					wg.Add(1)
					go func() {
						approveErr2 = d.am.infinitApproves(pair.QT, d.cfg.Router, adds...)
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
					d.tokens.add(pair.BT, pair.QT)
					return nil
				}

			}
			return errors.New(
				fmt.Sprintf("token `%s` for chain `%d` did not found in tokens list", qt, d.cfg.ChianId))
		}
	}
	return errors.New(
		fmt.Sprintf("token `%s` for chain `%d` did not found in tokens list", bt, d.cfg.ChianId))

}
