package uniswapv3

import (
	"encoding/json"
	"io"
	"net/http"
	"order_service/internal/delivery/exchanges/uniswap/v3/dto"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"os"
)

type tokens struct {
	Tokens []*token `json:"tokens"`
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

start:
	for _, p := range req.Pairs {
		bt := p.BaseToken
		qt := p.Quote_Token
		if bt == ether {
			bt = wrappedETH
		}
		if qt == ether {
			qt = wrappedETH
		}

		if u.pairs.exist(bt, qt) {
			res.Existed = append(res.Existed, p.String())
			continue
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
							break start
						}

						u.pairs.add(*pair)
						res.Added = append(res.Added, pair.String())
						break start
					}
				}
			}
		}
	}

	return res, nil
}
