package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
)

type OrderUseCase struct {
	repo entity.OrderRepo
	fs   entity.FeeService
	WalletStore
	exs *exStore
	l   logger.Logger
}

func NewOrderUseCase(repo entity.OrderRepo, exRepo ExchangeRepo, ws WalletStore,
	fee entity.FeeService, l logger.Logger) *OrderUseCase {

	o := &OrderUseCase{
		repo:        repo,
		WalletStore: ws,
		exs:         newExStore(l, exRepo),
		fs:          fee,
		l:           l,
	}
	return o
}

func (o *OrderUseCase) Run() {
	const agent = "Order-UseCase"

	// go o.wh.handle()
	// go o.wh.tracker.run()
	o.l.Debug(agent, "started")
}
