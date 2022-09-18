package uniswapv3

import (
	"math/big"
	"exchange-provider/internal/delivery/exchanges/uniswap/v3/contracts"
	"exchange-provider/pkg/errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

func (u *UniSwapV3) setBestPrice(bt, qt token) (*pair, error) {

	pool, err := u.highestLiquidPool(bt, qt)
	if err != nil {
		return nil, err
	}

	p, err := contracts.NewUniswapv3Pool(pool.address, u.provider)
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
		pool.baseIsZero = true
	}

	pool.price = price.ToSignificant(10)

	return pool, nil

}

func (u *UniSwapV3) highestLiquidPool(bt, qt token) (*pair, error) {
	agent := u.agent("highestLiquidPool")
	f := u.factory

	pairs := []*pair{}
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

			p, err := contracts.NewUniswapv3Pool(a, u.provider)
			if err != nil {
				u.l.Error(agent, err.Error())
				return
			}

			l, err := p.Liquidity(nil)
			if err != nil {
				u.l.Error(agent, err.Error())
				return
			}

			pairs = append(pairs, &pair{
				address:   a,
				liquidity: l,
				feeTier:   fee,
			})
		}(fee)
	}
	wg.Wait()

	pool := &pair{
		BT:        bt,
		QT:        qt,
		liquidity: common.Big0,
		address:   common.HexToAddress("0"),
	}

	for _, p := range pairs {
		if p.liquidity.Cmp(pool.liquidity) == +1 {
			pool.address = p.address
			pool.liquidity = p.liquidity
			pool.feeTier = p.feeTier
		}
	}
	if pool.address == common.HexToAddress("0") || pool.liquidity == big.NewInt(0) {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}
	return pool, nil
}
