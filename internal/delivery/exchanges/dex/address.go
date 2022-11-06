package dex

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (u *dex) GetAddress(c *entity.Coin) (*entity.Address, error) {
	if c.ChainId != u.cfg.TokenStandard {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	_, err := u.tokens.get(c.CoinId)
	if err != nil {
		return nil, err
	}

	addr, err := u.wallet.RandAddress()
	if err != nil {
		return nil, err
	}
	return &entity.Address{Addr: addr.String()}, nil
}
