package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"

	"order_service/pkg/errors"
)

type OrderUseCase struct {
	repo  entity.OrderRepo
	cache entity.OrderCache
	ds    entity.DepositeService
	exs   map[string]entity.Exchange
	oh    *orderHandler
	wh    *withdrawalHandler

	l logger.Logger
}

func NewOrderUseCase(repo entity.OrderRepo, oc entity.OrderCache, wc entity.WithdrawalCache,
	depo entity.DepositeService, fee entity.FeeService, exs map[string]entity.Exchange, l logger.Logger) *OrderUseCase {

	o := &OrderUseCase{
		repo:  repo,
		cache: oc,
		ds:    depo,
		exs:   exs,
		oh:    newOrderHandler(repo, oc, wc, fee, exs, l),
		wh:    newWithdrawalHandler(repo, oc, wc, exs, l),
		l:     l,
	}
	return o
}

func (o *OrderUseCase) Run(wg *sync.WaitGroup) {
	const agent = "Order-UseCase"
	defer wg.Done()
	w := &sync.WaitGroup{}
	w.Add(1)
	go o.oh.run(w)
	w.Add(1)
	go o.wh.run(w)

	o.l.Debug(agent, "started")

	w.Wait()

}

// user request an order
// steps:
// 1. create an order
// 2. send a request to the deposite service to create a deposite
// 3. get the deposite id as response and add it to the order
// 4. add the order to the cache
func (u *OrderUseCase) NewUserOrder(userId int64, address string, rCoin, pCoin entity.Coin) (*entity.UserOrder, error) {
	const op = errors.Op("Order-Usecase.NewUserOrder")
	o := entity.NewOrder(userId, address, rCoin, pCoin)
	d, err := u.ds.New(userId, o.Id, pCoin)
	if err != nil {
		err = errors.Wrap(err, op, fmt.Sprintf("userId: '%d',rCoin: %+v , pCoin: %+v ", userId, rCoin, pCoin),
			&ErrMsg{msg: "create deposite failed, internal error"})
		u.l.Error(string(op), err.Error())

		// remove the order from the cache
		if err := u.cache.Delete(userId, o.Id); err != nil {
			u.l.Error(string(op), fmt.Sprintf("orderId: '%d' userId: '%d'", o.Id, o.UserId))
		}
		return nil, err
	}

	o.AddDeposite(d)
	o.Status = entity.OrderStatusWaitForDepositeConfirm

	if err = u.cache.Add(o); err != nil {
		err = errors.Wrap(err, fmt.Sprintf("userId: '%d'", userId), op, &ErrMsg{msg: "create order failed, internal error"})
		u.l.Error(string(op), err.Error())
		return nil, err
	}
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

	cos, ce := u.GetAllPendingUserOrders(userId)
	if ce != nil {
		switch errors.ErrorCode(ce) {
		case errors.ErrNotFound:
			// ignore
		default:
			ce = errors.Wrap(ce, op, fmt.Sprintf("userId: '%d'", userId),
				&ErrMsg{msg: "get orders failed, internal error"})
			u.l.Error(string(op), ce.Error())
			return nil, ce
		}
	}

	pos, pe := u.GetAllClosedUserOrders(userId)
	if pe != nil {
		switch errors.ErrorCode(pe) {
		case errors.ErrNotFound:
			// ignore
		default:
			pe = errors.Wrap(pe, op, fmt.Sprintf("userId: '%d'", userId),
				&ErrMsg{msg: "get orders failed, internal error"})
			u.l.Error(string(op), pe.Error())
			return nil, pe
		}
	}
	if errors.ErrorCode(ce) == errors.ErrNotFound && errors.ErrorCode(pe) == errors.ErrNotFound {
		return nil, errors.Wrap(errors.ErrNotFound, op,
			&ErrMsg{msg: "orders not found"})
	}

	return append(cos, pos...), nil

}

// retrieve all orders from the permenant db
func (u *OrderUseCase) GetAllClosedUserOrders(userId int64) ([]*entity.UserOrder, error) {
	const op = errors.Op("Order-UseCase.GetAllUserOrders")

	os, err := u.repo.GetAll(userId)
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

// retrieve all orders from the cache
func (u *OrderUseCase) GetAllPendingUserOrders(userId int64) ([]*entity.UserOrder, error) {
	const op = errors.Op("Order-UseCase.GetAllPendingUserOrders")
	os, err := u.cache.GetAll(userId)
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

	o, err := u.cache.Get(userId, orderId)
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

	if o.Status != entity.OrderStatusWaitForDepositeConfirm {
		return errors.Wrap(err, op, &ErrMsg{msg: "order status is not waiting for deposite confirmation"})
	}
	o.Deposite.Volume = vol
	o.Deposite.Fullfilled = true
	o.Status = entity.OrderStatusDepositeConfimred

	if err := u.cache.Update(o); err != nil {
		err = errors.Wrap(err, o.String(), op, errors.ErrInternal)
		u.l.Error(string(op), err.Error())
		return err
	}

	u.oh.handle(o)
	return nil

}
