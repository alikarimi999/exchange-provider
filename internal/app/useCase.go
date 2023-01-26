package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
)

type OrderUseCase struct {
	repo  entity.OrderRepo
	pc    entity.PairConfigs
	oh    *orderHandler
	pairs entity.PairRepo
	wh    *withdrawalHandler
	fs    entity.FeeService

	WalletStore
	exs *exStore
	l   logger.Logger
}

func NewOrderUseCase(pairs entity.PairRepo, repo entity.OrderRepo, exRepo ExchangeRepo, ws WalletStore,
	pc entity.PairConfigs, fee entity.FeeService, l logger.Logger) *OrderUseCase {

	o := &OrderUseCase{
		repo:        repo,
		pc:          pc,
		pairs:       pairs,
		WalletStore: ws,
		exs:         newExStore(l, exRepo),
		fs:          fee,
		l:           l,
	}

	o.oh = newOrderHandler(o, repo, pc, fee, o.exs, l)
	o.wh = newWithdrawalHandler(o, repo, o.exs, l)
	return o
}

func (o *OrderUseCase) Run() {
	const agent = "Order-UseCase"

	go o.wh.handle()
	go o.wh.tracker.run()
	o.l.Debug(agent, "started")
}
