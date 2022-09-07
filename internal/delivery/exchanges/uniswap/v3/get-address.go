package uniswapv3

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (u *UniSwapV3) GetAddress(c *entity.Coin) (*entity.Address, error) {
	if c.ChainId != chainId {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}
	addr, err := u.wallet.RandAddress()
	if err != nil {
		return nil, err
	}
	return &entity.Address{Addr: addr.String()}, nil
}
