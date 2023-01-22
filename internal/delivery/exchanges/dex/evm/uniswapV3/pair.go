package uniswapV3

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV3/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

func (d *dex) Pair(t1, t2 types.Token) (*types.Pair, error) {
	pool, err := d.pair(t1, t2)
	if err != nil {
		return nil, err
	}

	p, err := contracts.NewUniswapv3Pool(pool.Address, d.provider())
	if err != nil {
		return nil, err
	}
	slot, err := p.Slot0(nil)
	if err != nil {
		return nil, err
	}

	at0, err := p.Token0(nil)
	if err != nil {
		return nil, err
	}

	at1, err := p.Token1(nil)
	if err != nil {
		return nil, err
	}

	if at0 == t1.Address && at1 == t2.Address {
		price := NewPrice(t1, t2, Q192, new(big.Int).Mul(slot.SqrtPriceX96, slot.SqrtPriceX96))
		pool.Price1 = price.ToSignificant(10)
		pool.Price2 = price.Invert().ToSignificant(10)
	} else {
		price := NewPrice(t2, t1, Q192, new(big.Int).Mul(slot.SqrtPriceX96, slot.SqrtPriceX96))
		pool.Price2 = price.ToSignificant(10)
		pool.Price1 = price.Invert().ToSignificant(10)
	}

	return pool, nil

}

func (d *dex) pair(t1, t2 types.Token) (*types.Pair, error) {
	agent := d.agent("pairWithPrice")

	f, err := contracts.NewUniswapv3Factory(d.factory, d.provider())
	if err != nil {
		return nil, err
	}

	pairs := []*types.Pair{}
	mux := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for _, fee := range feeTiers {
		wg.Add(1)
		go func(fee *big.Int) {
			defer wg.Done()

			a, err := f.GetPool(nil, t1.Address, t2.Address, fee)
			if err != nil {
				d.l.Error(agent, err.Error())
				return
			} else if a == common.HexToAddress("0") {
				return
			}

			p, err := contracts.NewUniswapv3Pool(a, d.provider())
			if err != nil {
				d.l.Error(agent, err.Error())
				return
			}

			l, err := p.Liquidity(nil)
			if err != nil {
				d.l.Error(agent, err.Error())
				return
			}

			mux.Lock()
			pairs = append(pairs, &types.Pair{
				Address:   a,
				Liquidity: l,
				FeeTier:   fee,
			})
			mux.Unlock()
		}(fee)
	}
	wg.Wait()

	pool := &types.Pair{
		T1:        t1,
		T2:        t2,
		Liquidity: common.Big0,
		Address:   common.HexToAddress("0"),
	}

	for _, p := range pairs {
		if p.Liquidity.Cmp(pool.Liquidity) == +1 {
			pool.Address = p.Address
			pool.Liquidity = p.Liquidity
			pool.FeeTier = p.FeeTier
		}
	}
	if pool.Address == common.HexToAddress("0") || pool.Liquidity == big.NewInt(0) {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}
	return pool, nil
}
