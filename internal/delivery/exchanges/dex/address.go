package dex

import (
	"exchange-provider/internal/entity"
)

func (u *dex) GetAddress(c *entity.Token) (*entity.Address, error) {
	addr, err := u.wallet.RandAddress(u.cfg.AccountCount)
	if err != nil {
		return nil, err
	}
	return &entity.Address{Addr: addr.String()}, nil
}
