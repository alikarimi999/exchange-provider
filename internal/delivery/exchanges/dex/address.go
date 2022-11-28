package dex

import (
	"exchange-provider/internal/entity"
)

func (u *dex) GetAddress(c *entity.Coin) (*entity.Address, error) {
	addr, err := u.wallet.RandAddress()
	if err != nil {
		return nil, err
	}
	return &entity.Address{Addr: addr.String()}, nil
}
