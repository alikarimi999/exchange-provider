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

func (u *dex) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := u.agent("AddPairs")

	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	ts := &tokens{}

	b, err := os.ReadFile(u.cfg.TokensFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, ts); err != nil {
		return nil, err
	}

	res := &entity.AddPairsResult{}

	pwg := &sync.WaitGroup{}

	for _, p := range req.Pairs {

		pwg.Add(1)
		go func(p *dto.Pair, ts tokens) {
			defer pwg.Done()
			var btIsNative bool
			var qtIsNative bool
			bt := p.BaseToken
			qt := p.Quote_Token

			if u.pairs.exist(bt, qt) {
				u.l.Debug(agent, fmt.Sprintf("pair %s already exists", p.String()))
				res.Existed = append(res.Existed, p.String())
				return
			}

			if bt == u.cfg.NativeToken {
				bt = u.cfg.wrapNative()
				btIsNative = true
			}
			if qt == u.cfg.NativeToken {
				qt = u.cfg.wrapNative()
				qtIsNative = true
			}

			for _, t0 := range ts.Tokens {
				if t0.ChainId != int64(u.cfg.ChianId) {
					continue
				}

				if t0.Symbol == bt {
					for _, t1 := range ts.Tokens {
						if t1.ChainId != int64(u.cfg.ChianId) {
							continue
						}
						if t1.Symbol == qt {
							pair, err := u.pairWithPrice(t0, t1)
							if err != nil {
								res.Failed = append(res.Failed, &entity.PairsErr{
									Pair: p.String(),
									Err:  err,
								})
								return
							}

							if btIsNative {
								pair.BT.Symbol = u.cfg.NativeToken
								pair.BT.Native = true
							} else if qtIsNative {
								pair.QT.Symbol = u.cfg.NativeToken
								pair.QT.Native = true
							}

							// check pair's tokens allowance
							adds, err := u.wallet.AllAddresses()
							if err != nil {
								u.l.Error(agent, err.Error())
								res.Failed = append(res.Failed, &entity.PairsErr{
									Pair: p.String(),
									Err:  err,
								})
							}

							wg := &sync.WaitGroup{}
							var approveErr1 []error
							wg.Add(1)
							go func() {
								approveErr1 = u.am.infinitApproves(pair.BT, u.cfg.Router, adds...)
								wg.Done()
							}()

							var approveErr2 []error
							wg.Add(1)
							go func() {
								approveErr2 = u.am.infinitApproves(pair.QT, u.cfg.Router, adds...)
								wg.Done()
							}()

							wg.Wait()
							if len(approveErr1) > 0 {
								for _, err := range approveErr1 {
									u.l.Error(agent, err.Error())
								}
								res.Failed = append(res.Failed, &entity.PairsErr{
									Pair: p.String(),
									Err:  errors.Wrap(fmt.Sprintf("%s", approveErr1)),
								})
								return
							}

							if approveErr2 != nil {
								for _, err := range approveErr2 {
									u.l.Error(agent, err.Error())
								}
								res.Failed = append(res.Failed, &entity.PairsErr{
									Pair: p.String(),
									Err:  errors.Wrap(fmt.Sprintf("%s", approveErr2)),
								})
								return
							}

							u.v.Set(fmt.Sprintf("%s.pairs.%s", u.NID(), pairId(pair.BT.Symbol, pair.QT.Symbol)), pair)
							if err := u.v.WriteConfig(); err != nil {
								u.l.Error(agent, err.Error())
								res.Failed = append(res.Failed, &entity.PairsErr{
									Pair: p.String(),
									Err:  err,
								})
								return
							}

							u.pairs.add(*pair)
							u.tokens.add(pair.BT, pair.QT)
							res.Added = append(res.Added, pair.String())
							u.l.Debug(agent, fmt.Sprintf("pair %s added", pair.String()))
							return
						}

					}
					res.Failed = append(res.Failed, &entity.PairsErr{
						Pair: p.String(),
						Err: errors.Wrap(errors.ErrNotFound, errors.NewMesssage(
							fmt.Sprintf("token `%s` for chain `%d` did not found in tokens list", qt, u.cfg.ChianId))),
					})
					return
				}
			}
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err: errors.Wrap(errors.ErrNotFound, errors.NewMesssage(
					fmt.Sprintf("token `%s` for chain `%d` did not found in tokens list", bt, u.cfg.ChianId))),
			})

		}(p, *ts)
	}
	pwg.Wait()
	return res, nil
}
