package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"

	"exchange-provider/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type OrderUseCase struct {
	repo  entity.OrderRepo
	cache entity.OrderCache
	pc    entity.PairConfigs
	rc    *redis.Client
	oh    *orderHandler

	wh *withdrawalHandler
	fs entity.FeeService

	WalletStore
	exs *exStore
	l   logger.Logger
}

func NewOrderUseCase(rc *redis.Client, repo entity.OrderRepo, exRepo ExchangeRepo, ws WalletStore,
	pc entity.PairConfigs, oc entity.OrderCache, fee entity.FeeService, l logger.Logger) *OrderUseCase {

	o := &OrderUseCase{
		repo:        repo,
		cache:       oc,
		rc:          rc,
		pc:          pc,
		WalletStore: ws,
		exs:         newExStore(l, exRepo),
		fs:          fee,
		l:           l,
	}

	o.oh = newOrderHandler(o, repo, oc, pc, oc, fee, o.exs, l)
	o.wh = newWithdrawalHandler(o, repo, oc, oc, o.exs, l)
	return o
}

func (o *OrderUseCase) Run() {
	const agent = "Order-UseCase"

	go o.wh.handle()
	go o.wh.tracker.run()
	o.l.Debug(agent, "started")
}

func (u *OrderUseCase) NewOrder(userId int64, wa *entity.Address,
	routes map[int]*entity.Route) (*entity.Order, error) {

	const op = errors.Op("Order-Usecase.NewUserOrder")

	ex, err := u.GetExchange(routes[0].Exchange)
	if err != nil {
		return nil, err
	}

	dc := routes[0].In
	da, err := ex.GetAddress(dc)
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
