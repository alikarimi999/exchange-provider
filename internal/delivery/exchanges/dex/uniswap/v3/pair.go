package uniswapv3

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/uniswap/v3/contracts"
	"exchange-provider/pkg/errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

func (u *UniswapV3) PairWithPrice(bt, qt types.Token) (*types.Pair, error) {

	pool, err := u.Pair(bt, qt)
	if err != nil {
		return nil, err
	}

	p, err := contracts.NewUniswapv3Pool(pool.Address, u.provider())
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

	price := NewPrice(bt, qt, Q192, new(big.Int).Mul(slot.SqrtPriceX96, slot.SqrtPriceX96))

	if at0 == qt.Address && at1 == bt.Address {
		price = NewPrice(bt, qt, new(big.Int).Mul(slot.SqrtPriceX96, slot.SqrtPriceX96), Q192)
	} else {
		pool.BaseIsZero = true
	}

	pool.Price1 = price.ToSignificant(10)
	pool.Price2 = price.Invert().ToSignificant(10)

	return pool, nil

}

func (u *UniswapV3) Pair(bt, qt types.Token) (*types.Pair, error) {
	agent := u.agent("pairWithPrice")

	f, err := contracts.NewUniswapv3Factory(u.factory, u.provider())
	if err != nil {
		return nil, err
	}

	pairs := []*types.Pair{}
	wg := &sync.WaitGroup{}
	for _, fee := range feeTiers {
		wg.Add(1)
		go func(fee *big.Int) {
			defer wg.Done()

			a, err := f.GetPool(nil, bt.Address, qt.Address, fee)
			if err != nil {
				u.l.Error(agent, err.Error())
				return
			} else if a == common.HexToAddress("0") {
				return
			}

			p, err := contracts.NewUniswapv3Pool(a, u.provider())
			if err != nil {
				u.l.Error(agent, err.Error())
				return
			}

			l, err := p.Liquidity(nil)
			if err != nil {
				u.l.Error(agent, err.Error())
				return
			}

			pairs = append(pairs, &types.Pair{
				Address:   a,
				Liquidity: l,
				FeeTier:   fee,
			})
		}(fee)
	}
	wg.Wait()

	pool := &types.Pair{
		T1:        bt,
		T2:        qt,
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
