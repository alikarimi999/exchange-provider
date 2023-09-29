package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (u *OrderUseCase) Allowance(tId entity.TokenId, owner string) (*entity.Allowance, error) {
	exs := u.exs.GetAll()
	for _, ex := range exs {
		if ex.Type() == entity.EvmDEX {
			if t := ex.(entity.EVMDex).GetToken(tId); t != nil {
				return ex.(entity.EVMDex).Allowance(t, owner)
			}
		}
	}
	return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("token not found"))
}
