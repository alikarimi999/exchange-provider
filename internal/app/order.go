package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"strconv"
	"sync"

	"order_service/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type OrderUseCase struct {
	repo  entity.OrderRepo
	cache entity.OrderCache
	pc    entity.PairConfigs
	rc    *redis.Client
	DS    entity.DepositeService
	oh    *orderHandler
	wh    *withdrawalHandler
	fs    entity.FeeService

	exs *exStore
	l   logger.Logger
}

func NewOrderUseCase(rc *redis.Client, repo entity.OrderRepo, exRepo ExchangeRepo, pc entity.PairConfigs, oc entity.OrderCache,
	depo entity.DepositeService, fee entity.FeeService, l logger.Logger) *OrderUseCase {

	o := &OrderUseCase{
		repo:  repo,
		cache: oc,
		rc:    rc,
		DS:    depo,
		pc:    pc,
		exs:   newExStore(l, exRepo),
		fs:    fee,
		l:     l,
	}

	o.oh = newOrderHandler(o, repo, oc, pc, oc, fee, o.exs, l)
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

	d, err := u.DS.New(userId, o.Id, dc, ex)
	if err != nil {
		return nil, err
	}

	o.Deposite = d

	o.Deposite.Fullfilled = true

	o.Status = entity.OrderStatusWaitForDepositeConfirm

	if err := u.write(o); err != nil {
		err = errors.Wrap(err, fmt.Sprintf("userId: '%d'", userId), op, &ErrMsg{msg: "create order failed, internal error"})
		u.l.Error(string(op), err.Error())
		return nil, err
	}

	u.l.Debug(string(op), fmt.Sprintf("order %s  created", o.String()))
	return o, nil
}

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

	if status == ExchangeStatusDisable {
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

	minBc, minQc := u.pc.PairMinDeposit(o.BC, o.QC)
	vf, err := strconv.ParseFloat(vol, 64)
	if err != nil {
		return errors.Wrap(err, op, &ErrMsg{msg: "invalid volume"})
	}
	switch o.Side {
	case "buy":
		if vf < minQc {
			o.Broken = true
			o.BreakReason = BR_InsufficientDepositVolume
			u.l.Info(string(op), o.BreakReason)

			if err := u.write(o); err != nil {
				u.l.Error(string(op), err.Error())
				return err
			}
			return nil
		}

	case "sell":

		if vf < minBc {
			o.Broken = true
			o.BreakReason = BR_InsufficientDepositVolume
			u.l.Info(string(op), o.BreakReason)

			if err := u.write(o); err != nil {
				u.l.Error(string(op), err.Error())
				return err
			}
			return nil
		}

	default:
		return errors.Wrap(err, op, errors.NewMesssage(fmt.Sprintf("order side is %s not supported", o.Side)))
	}

	o.Status = entity.OrderStatusDepositeConfimred

	if err := u.write(o); err != nil {
		err = errors.Wrap(err, o.String(), op, errors.ErrInternal)
		u.l.Error(string(op), err.Error())
		return err
	}

	u.oh.handle(o)
	return nil

}
