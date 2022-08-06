package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"

	"order_service/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type OrderUseCase struct {
	repo  entity.OrderRepo
	cache entity.OrderCache
	rc    *redis.Client
	ds    entity.DepositeService
	oh    *orderHandler
	wh    *withdrawalHandler
	fs    entity.FeeService

	exs *exStore
	l   logger.Logger
}

func NewOrderUseCase(rc *redis.Client, repo entity.OrderRepo, oc entity.OrderCache,
	depo entity.DepositeService, fee entity.FeeService, l logger.Logger) *OrderUseCase {

	o := &OrderUseCase{
		repo:  repo,
		cache: oc,
		rc:    rc,
		ds:    depo,
		exs:   newExStore(l),
		fs:    fee,

		l: l,
	}

	o.oh = newOrderHandler(o, repo, oc, oc, fee, o.exs, l)
	o.wh = newWithdrawalHandler(o, repo, oc, oc, o.exs, l)
	return o
}

func (o *OrderUseCase) Run(wg *sync.WaitGroup) {
	const agent = "Order-UseCase"
	defer wg.Done()
	w := &sync.WaitGroup{}
	w.Add(1)
	go o.oh.run(w)
	w.Add(1)
	go o.wh.handle(w)

	wg.Add(1)
	go o.wh.tracker.run(wg)

	wg.Add(1)
	go o.exs.start(w)

	o.l.Debug(agent, "started")

	w.Wait()

}

// user request an order
// steps:
// 1. create an order
// 2. send a request to the deposite service to create a deposite
// 3. get the deposite id as response and add it to the order
// 4. add the order to the cache
func (u *OrderUseCase) NewUserOrder(userId int64, address string, bc, qc *entity.Coin, side, ex string) (*entity.UserOrder, error) {
	const op = errors.Op("Order-Usecase.NewUserOrder")

	o := entity.NewOrder(userId, address, bc, qc, side, ex)

	if err := u.write(o); err != nil {
		return nil, errors.Wrap(err, op, &ErrMsg{msg: "create order failed, internal error"})
	}

	var dc *entity.Coin
	if side == "buy" {
		dc = qc
	} else {
		dc = bc
	}

	d, err := u.ds.New(userId, o.Id, dc, ex)
	if err != nil {
		switch errors.ErrorCode(err) {
		case errors.ErrNotFound:
			err = errors.Wrap(err, op, &ErrMsg{msg: fmt.Sprintf("coin %s chain %s not found in deposit service", dc.CoinId, dc.ChainId)})
			u.l.Debug(string(op), err.Error())

		default:
			err = errors.Wrap(err, op, fmt.Sprintf("userId: '%d',quote_coin: %+v , base_oin: %+v ", userId, bc, qc),
				&ErrMsg{msg: "create deposite failed, internal error"})
			u.l.Error(string(op), err.Error())

			// remove the order from the cache
			if err := u.cache.Delete(userId, o.Id); err != nil {
				u.l.Error(string(op), fmt.Sprintf("orderId: '%d' userId: '%d'", o.Id, o.UserId))
			}
			return nil, err
		}
	}

	o.Deposite = d

	o.Deposite.Fullfilled = true

	o.Status = entity.OrderStatusWaitForDepositeConfirm

	if err := u.write(o); err != nil {
		err = errors.Wrap(err, fmt.Sprintf("userId: '%d'", userId), op, &ErrMsg{msg: "create order failed, internal error"})
		u.l.Error(string(op), err.Error())
		return nil, err
	}

	// if err = u.cache.Add(o); err != nil {
	// 	err = errors.Wrap(err, fmt.Sprintf("userId: '%d'", userId), op, &ErrMsg{msg: "create order failed, internal error"})
	// 	u.l.Error(string(op), err.Error())
	// 	return nil, err
	// }
	u.l.Debug(string(op), fmt.Sprintf("userId: '%d', orderId: '%d'. depositeId: '%d' created", o.UserId, o.Id, o.Deposite.Id))
	return o, nil
}

// retrieve the order from the cache
// if retrive from cache failed try permenant db
func (u *OrderUseCase) GetUserOrder(userId, orderId int64) (*entity.UserOrder, error) {
	const op = errors.Op("Order-UseCase.GetUserOrder")
	o, err := u.cache.Get(userId, orderId)
	if err != nil {
		switch errors.ErrorCode(err) {
		case errors.ErrNotFound:
			err = errors.Wrap(err, op, &ErrMsg{msg: "order not found"})
			u.l.Debug(string(op), err.Error())
		default:
			err = errors.Wrap(err, op, fmt.Sprintf("userId: '%d' , orderId: '%d' ", userId, orderId), &ErrMsg{msg: "get order failed, internal error"})
			u.l.Error(string(op), err.Error())
		}

		o, err = u.repo.Get(userId, orderId)
		if err != nil {
			switch errors.ErrorCode(err) {
			case errors.ErrNotFound:
				err = errors.Wrap(err, op, &ErrMsg{msg: "order not found"})
				u.l.Debug(string(op), err.Error())
				return nil, err
			default:
				err = errors.Wrap(err, op, fmt.Sprintf("userId: '%d' , orderId: '%d' ", userId, orderId),
					&ErrMsg{msg: "get order failed, internal error"})
				u.l.Error(string(op), err.Error())
				return nil, err
			}
		}
	}
	return o, nil
}

func (u *OrderUseCase) GetAllUserOrders(userId int64) ([]*entity.UserOrder, error) {
	const op = errors.Op("Order-UseCase.GetAllUserOrders")

	os := []*entity.UserOrder{{UserId: userId}}
	if err := u.read(&os); err != nil {
		switch errors.ErrorCode(err) {
		case errors.ErrNotFound:
			err = errors.Wrap(err, op, &ErrMsg{msg: "orders not found"})
			u.l.Debug(string(op), err.Error())
		default:
			err = errors.Wrap(err, op, fmt.Sprintf("userId: '%d'", userId), &ErrMsg{msg: "get orders failed, internal error"})
			u.l.Error(string(op), err.Error())
		}

		return nil, err
	}
	return os, nil

}

func (u *OrderUseCase) GetAllClosedUserOrders(userId int64) ([]*entity.UserOrder, error) {
	const op = errors.Op("Order-UseCase.GetAllUserOrders")

	os := []*entity.UserOrder{{UserId: userId}}
	if err := u.read(userId); err != nil {
		switch errors.ErrorCode(err) {
		case errors.ErrNotFound:
			err = errors.Wrap(err, op, &ErrMsg{msg: "orders not found"})
			u.l.Debug(string(op), err.Error())
		default:
			err = errors.Wrap(err, op, fmt.Sprintf("userId: '%d'", userId), &ErrMsg{msg: "get orders failed, internal error"})
			u.l.Error(string(op), err.Error())
		}

		return nil, err
	}
	return os, nil
}

// retrieve all orders from the cache
func (u *OrderUseCase) GetAllPendingUserOrders(userId int64) ([]*entity.UserOrder, error) {
	const op = errors.Op("Order-UseCase.GetAllPendingUserOrders")
	os := []*entity.UserOrder{}
	err := u.read(userId)
	if err != nil {
		switch errors.ErrorCode(err) {
		case errors.ErrNotFound:
			err = errors.Wrap(err, op, &ErrMsg{msg: "orders not found"})
			u.l.Debug(string(op), err.Error())
		default:
			err = errors.Wrap(err, op, fmt.Sprintf("userId: '%d'", userId), &ErrMsg{msg: "get orders failed, internal error"})
			u.l.Error(string(op), err.Error())
		}

		return nil, err
	}
	return os, nil
}

// it will be called after the `deposite.confirmed` event is received
// steps:
// 1. retrive the order from the cache
// 2. update the order's deposite volume and status
// 3. update the order's state on the cache.
// 4. send the order to the order handler
func (u *OrderUseCase) SetDepositeVolume(userId, orderId, depositeId int64, vol string) error {
	const op = errors.Op("Order-UseCase.SetDepositeVolume")

	o := &entity.UserOrder{Id: orderId, UserId: userId}

	err := u.read(o)
	if err != nil {
		switch errors.ErrorCode(err) {
		case errors.ErrNotFound:
			err = errors.Wrap(err, op, &ErrMsg{msg: "order not found"})
			u.l.Debug(string(op), err.Error())
			return err
		default:
			err = errors.Wrap(err, op, fmt.Sprintf("userId: '%d' , orderId: '%d' , depositeId: '%d' vol: '%s' ", userId, orderId,
				depositeId, vol), &ErrMsg{msg: "get order failed, internal error"})

			u.l.Error(string(op), err.Error())
			return err
		}
	}

	ok, status := u.exs.exists(o.Exchange)
	if !ok {
		err = errors.Wrap(errors.ErrBadRequest, op, errors.NewMesssage(fmt.Sprintf("exchange: '%s' not supported by this service", o.Exchange)))
		u.l.Error(string(op), err.Error())
		o.Broken = true
		o.BreakReason = errors.ErrorMsg(err)

		if er := u.write(o); err != nil {
			u.l.Error(string(op), err.Error())
			err = fmt.Errorf("%s, %s", err.Error(), er.Error())
			return err
		}

		return err
	}

	if status == ExchangeStatusDisabled {
		err = errors.Wrap(errors.ErrBadRequest, op, errors.NewMesssage(fmt.Sprintf("exchange: '%s' is disabled", o.Exchange)))
		u.l.Error(string(op), err.Error())
		o.Broken = true
		o.BreakReason = errors.ErrorMsg(err)

		if er := u.write(o); err != nil {
			u.l.Error(string(op), err.Error())
			err = fmt.Errorf("%s, %s", err.Error(), er.Error())
			return err
		}

		return err
	}

	if o.Status != entity.OrderStatusWaitForDepositeConfirm {
		return errors.Wrap(err, op, &ErrMsg{msg: "order status is not waiting for deposite confirmation"})
	}

	o.Deposite.Id = depositeId
	o.Deposite.Volume = vol

	switch o.Side {
	case "buy":
		o.Funds = vol
	case "sell":
		o.Size = vol
	default:
		return errors.Wrap(err, op, errors.NewMesssage(fmt.Sprintf("order side is %s not supported", o.Side)))
	}

	if err := u.write(o); err != nil {
		err = errors.Wrap(err, o.String(), op, errors.ErrInternal)
		u.l.Error(string(op), err.Error())
		return err
	}

	u.oh.handle(o)
	return nil

}
