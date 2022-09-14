package uniswapv3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order_service/internal/delivery/exchanges/uniswap/v3/dto"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"os"
	"sync"
)

type tokens struct {
	Tokens []token `json:"tokens"`
}

func (u *UniSwapV3) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := u.agent("AddPairs")

	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	ts := &tokens{}

	b, err := os.ReadFile(u.cfg.TokensFile)
	if err != nil {
		u.l.Error(agent, err.Error())

		res, err := http.DefaultClient.Get(u.cfg.TokensUrl)
		if err != nil {
			return nil, err
		}
		b, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = os.WriteFile(u.cfg.TokensFile, b, os.ModePerm)
		if err != nil {
			u.l.Error(agent, err.Error())
		}
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

			if bt == ether {
				bt = wrappedETH
				btIsNative = true
			}
			if qt == ether {
				qt = wrappedETH
				qtIsNative = true
			}

			for _, t0 := range ts.Tokens {
				if t0.ChainId != u.chainId.Int64() {
					continue
				}

				if t0.Symbol == bt {
					for _, t1 := range ts.Tokens {
						if t1.ChainId != u.chainId.Int64() {
							continue
						}
						if t1.Symbol == qt {
							pair, err := u.highestLiquidPool(t0, t1)
							if err != nil {
								res.Failed = append(res.Failed, &entity.PairsErr{
									Pair: p.String(),
									Err:  err,
								})
								return
							}

							if btIsNative {
								pair.BT.Symbol = ether
							} else if qtIsNative {
								pair.QT.Symbol = ether
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
								approveErr1 = u.am.infinitApproves(pair.BT, routerV2, adds...)
								wg.Done()
							}()

							var approveErr2 []error
							wg.Add(1)
							go func() {
								approveErr2 = u.am.infinitApproves(pair.QT, routerV2, adds...)
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
							fmt.Sprintf("token `%s` for chain `%d` did not found in tokens list", qt, u.chainId.Int64()))),
					})
					return
				}
			}
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err: errors.Wrap(errors.ErrNotFound, errors.NewMesssage(
					fmt.Sprintf("token `%s` for chain `%d` did not found in tokens list", bt, u.chainId.Int64()))),
			})

		}(p, *ts)
	}
	pwg.Wait()
	return res, nil
}
