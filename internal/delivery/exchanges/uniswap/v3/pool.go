package uniswapv3

import (
	"math"
	"math/big"
	"order_service/internal/delivery/exchanges/uniswap/v3/contracts"
	"order_service/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

func (u *UniSwapV3) bestPool(bt, qt *token) (*pair, error) {

	pool, err := u.highestLiquidPool(bt.address, qt.address)
	if err != nil {
		return nil, err
	}

	p, err := contracts.NewUniswapv3Pool(pool.address, u.dp)
	if err != nil {
		return nil, err
	}
	slot, err := p.Slot0(nil)
	if err != nil {
		return nil, err
	}

	t0, err := p.Token0(nil)
	if err != nil {
		return nil, err
	}

	t1, err := p.Token1(nil)
	if err != nil {
		return nil, err
	}

	pf, price := price(slot.SqrtPriceX96, bt.decimals, qt.decimals)

	if t0 == qt.address && t1 == bt.address {
		price = new(big.Float).Quo(big.NewFloat(1), pf).Text('f', pricePrec)
	} else {
		pool.baseIsZero = true
	}

	pool.bt = bt
	pool.qt = qt
	pool.price = price

	return pool, nil

}

func (u *UniSwapV3) highestLiquidPool(t0, t1 common.Address) (*pair, error) {

	f := u.factory

	pool := &pair{
		address:   common.HexToAddress("0"),
		liquidity: big.NewInt(0),
	}

	for _, fee := range feeTiers {
		a, err := f.GetPool(nil, t0, t1, fee)
		if err != nil || a == common.HexToAddress("0") {
			continue
		}

		p, err := contracts.NewUniswapv3Pool(a, u.dp)
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

func price(SqrtPriceX96 *big.Int, d0, d1 int) (*big.Float, string) {
	dom := new(big.Int).Exp(SqrtPriceX96, big.NewInt(2), nil)
	num := new(big.Int).Exp(big.NewInt(2), big.NewInt(192), nil)
	price := new(big.Float).Quo(new(big.Float).SetInt(dom), new(big.Float).SetInt(num))
	price = new(big.Float).Mul(price, big.NewFloat(math.Pow10(int(math.Abs(float64(d0-d1))))))
	return price, price.Text('f', pricePrec)
}
