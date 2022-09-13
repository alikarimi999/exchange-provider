package uniswapv3

import (
	"math/big"
	"order_service/internal/delivery/exchanges/uniswap/v3/contracts"
	"order_service/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

func (u *UniSwapV3) setBestPrice(bt, qt token) (*pair, error) {

	pool, err := u.highestLiquidPool(bt, qt)
	if err != nil {
		return nil, err
	}

	p, err := contracts.NewUniswapv3Pool(pool.address, u.Provider)
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

	f := u.factory

	pool := &pair{
		BT:        bt,
		QT:        qt,
		address:   common.HexToAddress("0"),
		liquidity: big.NewInt(0),
	}

	for _, fee := range feeTiers {
		a, err := f.GetPool(nil, bt.Address, qt.Address, fee)
		if err != nil || a == common.HexToAddress("0") {
			continue
		}

		p, err := contracts.NewUniswapv3Pool(a, u.Provider)
		if err != nil {
			continue
		}

		l, err := p.Liquidity(nil)
		if err != nil {
			continue
		}

		if l.Cmp(pool.liquidity) == +1 {
			pool.address = a
			pool.liquidity = l
			pool.feeTier = fee
		}

	}
	if pool.address == common.HexToAddress("0") || pool.liquidity == big.NewInt(0) {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}
	return pool, nil
}
