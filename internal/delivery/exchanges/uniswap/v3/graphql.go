package uniswapv3

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type GetBestPoolGraph struct {
	Data struct {
		Pools []*poolGraph `json:"pools"`
	} `json:"data"`
}

func (p *poolGraph) ToPool() (*pool, error) {

	pool := &pool{
		id: p.Id,
		t0: &token{
			symbol:  p.T0.Symbol,
			address: common.HexToAddress(p.T0.Id),
		},
		t1: &token{
			symbol:  p.T1.Symbol,
			address: common.HexToAddress(p.T1.Id),
		},
		t0Price: p.T0Price,
		t1Price: p.T1Price,
	}
	t0d, err := strconv.Atoi(p.T0.Decimals)
	if err != nil {
		return nil, err
	}
	t1d, err := strconv.Atoi(p.T1.Decimals)
	if err != nil {
		return nil, err
	}

	feeTier, err := strconv.Atoi(p.FeeTier)
	if err != nil {
		return nil, err
	}

	pool.t0.decimals = t0d
	pool.t1.decimals = t1d
	pool.feeTier = feeTier

	return pool, nil
}

type tokenGraph struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals string `json:"decimals"`
}
type poolGraph struct {
	Id      string      `json:"id"`
	T0      *tokenGraph `json:"token0"`
	T1      *tokenGraph `json:"token1"`
	T0Price string      `json:"token0Price"`
	T1Price string      `json:"token1Price"`
	FeeTier string      `json:"feeTier"`
}
