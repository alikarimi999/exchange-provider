package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func (u *OrderUseCase) NewEvmOrder(userId string, sender, reveiver common.Address,
	amountIn float64, route *entity.Route) (*entity.EvmOrder, error) {
	const op = errors.Op("OrderUsecase.NewEvmDexOrder")

	ex, err := u.exs.get(route.Exchange)
	if err != nil {
		return nil, err
	}

	sf := u.fs.GetUserFee(userId)
	f, _ := strconv.ParseFloat(sf, 64)
	o := entity.NewEvmOrder(userId, make(map[uint]*entity.EvmStep), sender, reveiver, amountIn, f)
	if err := ex.(entity.EVMDex).SetStpes(o, route); err != nil {
		return nil, err
	}
	if err := u.write(o); err != nil {
		u.l.Error(string(op), err.Error())
		return nil, err
	}

	return o, nil
}

func (u *OrderUseCase) NewCexOrder(userId string, wa *entity.Address,
	routes map[int]*entity.Route) (*entity.CexOrder, error) {

	const op = errors.Op("OrderUsecase.NewCexOrder")

	ex, err := u.exs.get(routes[0].Exchange)
	if err != nil {
		return nil, err
	}

	dc := routes[0].In
	da, err := ex.(entity.Cex).GetAddress(dc)
	if err != nil {
		return nil, err
	}

	o := entity.NewOrder(userId, wa, da, routes)

	if err := u.write(o); err != nil {
		u.l.Error(string(op), err.Error())
		return nil, errors.Wrap(err, op, errors.NewMesssage("create order failed, internal error"))
	}

	return o, nil
}
