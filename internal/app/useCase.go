package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"

	"github.com/go-redis/redis/v9"
)

type OrderUseCase struct {
	repo  entity.OrderRepo
	cache entity.OrderCache
	pc    entity.PairConfigs
	rc    *redis.Client
	oh    *orderHandler
	pairs entity.PairRepo
	wh    *withdrawalHandler
	fs    entity.FeeService

	WalletStore
	exs *exStore
	l   logger.Logger
}

func NewOrderUseCase(pairs entity.PairRepo, rc *redis.Client, repo entity.OrderRepo, exRepo ExchangeRepo, ws WalletStore,
	pc entity.PairConfigs, oc entity.OrderCache, fee entity.FeeService, l logger.Logger) *OrderUseCase {

	o := &OrderUseCase{
		repo:        repo,
		cache:       oc,
		rc:          rc,
		pc:          pc,
		pairs:       pairs,
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
