package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"sync"

	"exchange-provider/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type OrderUseCase struct {
	repo  entity.OrderRepo
	cache entity.OrderCache
	pc    entity.PairConfigs
	rc    *redis.Client
	oh    *orderHandler
	dh    *depositHandler
	wh    *withdrawalHandler
	fs    entity.FeeService

	exs *exStore
	l   logger.Logger
}

func NewOrderUseCase(rc *redis.Client, repo entity.OrderRepo, exRepo ExchangeRepo,
	pc entity.PairConfigs, oc entity.OrderCache, fee entity.FeeService, l logger.Logger) *OrderUseCase {

	o := &OrderUseCase{
		repo:  repo,
		cache: oc,
		rc:    rc,
		pc:    pc,
		exs:   newExStore(l, exRepo),
		fs:    fee,
		l:     l,
	}

	o.oh = newOrderHandler(o, repo, oc, pc, oc, fee, o.exs, l)
	o.dh = newDepositHandler(o)
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
	go o.dh.handle(wg)
	wg.Add(1)
	go o.wh.tracker.run(wg)

	wg.Add(1)
	go o.exs.start(w)

	o.l.Debug(agent, "started")

	w.Wait()

}

func (u *OrderUseCase) NewUserOrder(userId int64, wa *entity.Address, bc, qc *entity.Coin, side string, ex entity.Exchange) (*entity.UserOrder, error) {
	const op = errors.Op("Order-Usecase.NewUserOrder")

	var dc *entity.Coin
	if side == "buy" {
		dc = qc
	} else {
		dc = bc
	}

	da, err := ex.GetAddress(dc)
	if err != nil {
		return nil, err
	}

	o := entity.NewOrder(userId, wa, da, bc, qc, side, ex.NID())

	if err := u.write(o); err != nil {
		return nil, errors.Wrap(err, op, errors.NewMesssage("create order failed, internal error"))
	}

	// u.l.Debug(string(op), fmt.Sprintf("order (%s) created", o.String()))
	return o, nil
}
