package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func (u *OrderUseCase) NewOrder(userId string, sender, refund, reciever entity.Address,
	in, out *entity.Token, amount float64, lp uint) (entity.Order, error) {

	if refund.Addr == "" {
		refund = sender
	}
	routes := make(map[int]*entity.Route)
	var err error
	if lp > 0 {
		ex, err := u.exs.get(lp)
		if err != nil {
			return nil, err
		}

		routes[0] = &entity.Route{
			In:       in,
			Out:      out,
			Exchange: ex.Name(),
			ExType:   ex.Type(),
		}
	} else {
		routes, err = u.routing(in, out, amount)
		if err != nil {
			return nil, err
		}
	}
	ex, _ := u.exs.getByName(routes[0].Exchange)
	switch ex.Type() {
	case entity.EvmDEX:
		return u.newEvmOrder(userId, common.HexToAddress(refund.Addr),
			common.HexToAddress(reciever.Addr), amount, routes[0])
	default:
		if refund.Addr == "" {
			refund = sender
		}
		return u.newCexOrder(userId, refund, reciever, amount, routes)
	}
}

func (u *OrderUseCase) newEvmOrder(userId string, sender, reciever common.Address,
	amountIn float64, route *entity.Route) (*entity.EvmOrder, error) {
	const op = errors.Op("OrderUsecase.NewEvmDexOrder")

	ex, err := u.exs.getByName(route.Exchange)
	if err != nil {
		return nil, err
	}

	sf := u.fs.GetUserFee(userId)
	f, _ := strconv.ParseFloat(sf, 64)
	o := entity.NewEvmOrder(userId, make(map[uint]*entity.EvmStep), sender, reciever, amountIn, f)
	if err := ex.(entity.EVMDex).SetStpes(o, route); err != nil {
		return nil, err
	}
	if err := u.write(o); err != nil {
		u.l.Error(string(op), err.Error())
		return nil, err
	}

	return o, nil
}

func (u *OrderUseCase) newCexOrder(userId string, refund, reciever entity.Address,
	amount float64, routes map[int]*entity.Route) (*entity.CexOrder, error) {

	const op = errors.Op("OrderUsecase.NewCexOrder")
	ex, err := u.exs.getByName(routes[0].Exchange)
	if err != nil {
		return nil, err
	}

	o := entity.NewCexOrder(userId, refund, reciever, routes, amount)
	if err := ex.(entity.Cex).SetDepositddress(o); err != nil {
		return nil, err
	}
	if err := u.write(o); err != nil {
		u.l.Error(string(op), err.Error())
		return nil, errors.Wrap(err, op, errors.NewMesssage("create order failed, internal error"))
	}

	return o, nil
}
