package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"

	"order_service/pkg/errors"
)

// handle user orders in multiple steps
// steps:
// 1. open a new order in exchange to exchange user provided coin to requested coin
// 2. track the exchange order
// 3. if exchange order succeed, withdrawal the requested coin to user provided address
// 4. if the exchange return the withdrawal id, add the withdrawal to withdrawal cache and withdrawal handler proccess will track it's status

type orderHandler struct {
	repo     entity.OrderRepo
	oc       entity.OrderCache
	wc       entity.WithdrawalCache
	exs      map[string]entity.Exchange
	eTracker *exOrderTracker
	oCh      chan *entity.UserOrder
	fee      entity.FeeService

	l logger.Logger
}

func newOrderHandler(repo entity.OrderRepo, oc entity.OrderCache, wc entity.WithdrawalCache, fee entity.FeeService, exs map[string]entity.Exchange, l logger.Logger) *orderHandler {
	oh := &orderHandler{
		repo:     repo,
		oc:       oc,
		wc:       wc,
		exs:      exs,
		eTracker: newExOrderTracker(oc, l),
		oCh:      make(chan *entity.UserOrder, 1024),
		fee:      fee,

		l: l,
	}
	return oh
}

func (o *orderHandler) run(wg *sync.WaitGroup) {
	defer wg.Done()
	const op = errors.Op("User-Order-Handler.run")

	go o.eTracker.run()

	for {
		select {
		case order := <-o.oCh:

			go func(ord *entity.UserOrder) {
				o.l.Debug(string(op), fmt.Sprintf("handle order: '%d' for user: '%d'", ord.Id, ord.UserId))
				ex := o.exs[ord.Exchange]

				// 1. open a new order in exchange to exchange user provided coin to requested coin
				id, err := ex.Exchange(ord.ProvideCoin, ord.RequestCoin, ord.Deposite.Volume)
				if err != nil {

					ord.Broken = true
					ord.BrokeReason = fmt.Sprintf("unable to create order in exchange: %s", err.Error())

					o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())

					if err := o.oc.Update(ord); err != nil {
						o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
					}
					return

				}

				ord.ExchangeOrder.Id = id
				ord.Status = entity.OrderStatusWaitForExchangeOrderConfirm
				if err = o.oc.Update(ord); err != nil {
					o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
					return
				}

				ef := &exTrackerFeed{
					userId:  ord.UserId,
					orderId: ord.Id,
					eo:      ord.ExchangeOrder,
					ex:      ex,
					succeed: make(chan bool),
				}
				go o.eTracker.track(ef)
				o.l.Debug(string(op), fmt.Sprintf("order: '%d' for user: '%d' is waiting for exchange order: '%s' confirmation", ord.Id, ord.UserId, ord.ExchangeOrder.Id))

				if <-ef.succeed {
					switch ord.ExchangeOrder.Status {
					case entity.ExOrderSucceed:

						ord.Status = entity.OrderStatusExchangeOrderConfirmed
						o.l.Debug(string(op), fmt.Sprintf("order: '%d' for user: '%d' exchange order: '%s' confirmed", ord.Id, ord.UserId, ord.ExchangeOrder.Id))
						if err = o.oc.Update(ord); err != nil {
							o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
							return
						}

						switch ord.ExchangeOrder.Side {
						case "buy":
							ord.Withdrawal.Total = ord.ExchangeOrder.Size

						case "sell":
							ord.Withdrawal.Total = ord.ExchangeOrder.Funds
						}

						r, f, err := o.fee.ApplyFee(ord.UserId, ord.Withdrawal.Total)
						if err != nil {
							ord.Broken = true
							ord.BrokeReason = fmt.Sprintf("unable to apply fee: %s", err.Error())

							o.l.Error(string(op), errors.Wrap(err, op,
								fmt.Sprintf("orderId: '%d', userId: '%d'", ord.Id, ord.UserId)).Error())

							if err := o.oc.Update(ord); err != nil {
								o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
							}
							return

						}

						o.l.Debug(string(op), fmt.Sprintf("order: %d user: %d , transferring '%s' %v to '%s'", ord.Id, ord.UserId, r, ord.RequestCoin, ord.Withdrawal.Address))
						ord.Withdrawal.Fee = f
						ord.Withdrawal.Id, err = ex.Withdrawal(ord.RequestCoin, ord.Withdrawal.Address, r)
						if err != nil {
							ord.Broken = true
							ord.BrokeReason = fmt.Sprintf("unable to create withdrawal in exchange: %s", err.Error())

							o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())

							if err := o.oc.Update(ord); err != nil {
								o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
							}
							return

						}

						ord.Withdrawal.Status = entity.WithdrawalPending
						ord.Status = entity.OrderStatusWaitForWithdrawalConfirm
						o.l.Debug(string(op), fmt.Sprintf("order: '%d' for user: '%d' withdrawal order: '%s' created", ord.Id, ord.UserId, ord.Withdrawal.Id))
						if err = o.oc.Update(ord); err != nil {
							o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
							return
						}

						// add to withdrawal cache
						// and wait for withdrawal confirm
						if err := o.wc.AddPendingWithdrawal(ord.Withdrawal); err != nil {
							o.l.Error(string(op), errors.Wrap(err, op, ord.Withdrawal.String()).Error())
						}
						o.l.Debug(string(op), fmt.Sprintf("order: '%d' for user: '%d' is waiting for withdrawal: '%s' to confirm", ord.Id, ord.UserId, ord.Withdrawal.Id))
						return

					case entity.ExOrderPending:
						o.handle(ord)
						return

					default:

						ord.Broken = true
						ord.BrokeReason = "exchange order failed"

						o.l.Error(string(op), fmt.Sprintf("order: '%d' for user: '%d' exchange order: '%s' has status: '%s'", ord.Id, ord.UserId, ord.ExchangeOrder.Id, ord.ExchangeOrder.Status))

						if err := o.oc.Update(ord); err != nil {
							o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
						}
						return

					}
				}

				ord.Broken = true
				ord.BrokeReason = "exchange order tracking failed"

				o.l.Error(string(op), fmt.Sprintf("order: '%d' for user: '%d' exchange order: '%s' tracking failed for unknown reason.", ord.Id, ord.UserId, ord.ExchangeOrder.Id))

				if err := o.oc.Update(ord); err != nil {
					o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
				}
				return

			}(order)

		}
	}

}

func (h *orderHandler) handle(o *entity.UserOrder) {
	h.oCh <- o
}
