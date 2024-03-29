package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
)

type OrderUseCase struct {
	repo entity.OrderRepo
	fs   entity.FeeTable
	WalletStore
	exs entity.ExchangeStore
	l   logger.Logger
}

func NewOrderUseCase(repo entity.OrderRepo, exs entity.ExchangeStore,
	ws WalletStore, fee entity.FeeTable, l logger.Logger) *OrderUseCase {

	o := &OrderUseCase{
		repo:        repo,
		WalletStore: ws,
		exs:         exs,
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
