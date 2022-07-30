package entity

import (
	"encoding/json"
	"order_service/pkg/utils"
	"time"
)

type OrderStatus string

const (
	OrderStatusOpen OrderStatus = "open"

	OrderStatusWaitForDepositeConfirm OrderStatus = "wait_for_deposite_confirm"
	OrderStatusDepositeConfimred      OrderStatus = "deposite_confirmed"

	OrderStatusWaitForExchangeOrderConfirm OrderStatus = "wait_for_exchange_order_confirm"
	OrderStatusExchangeOrderConfirmed      OrderStatus = "exchange_order_confirmed"

	OrderStatusWaitForWithdrawalConfirm OrderStatus = "wait_for_withdrawal_confirm"
	OrderStatusWithdrawalConfirmed      OrderStatus = "withdrawal_confirmed"

	OrderStatusSecceed OrderStatus = "succeed"
	OrderStatusFailed  OrderStatus = "failed"
)

type UserOrder struct {
	Id            int64
	UserId        int64
	CreatedAt     int64
	Status        OrderStatus
	Deposite      *Deposit
	Exchange      string
	Withdrawal    *Withdrawal
	BC            *Coin
	QC            *Coin
	Side          string // buy or sell
	Size          string
	Funds         string
	ExchangeOrder *ExchangeOrder
	Broken        bool
	BrokeReason   string
}

func NewOrder(userId int64, address string, bc, qc *Coin, side, ex string) *UserOrder {
	w := &UserOrder{
		Id:        genOrderId(9),
		UserId:    userId,
		CreatedAt: time.Now().Unix(),
		Status:    OrderStatusOpen,
		Exchange:  ex,
		BC:        bc,
		QC:        qc,
		Side:      side,
		Deposite: &Deposit{
			UserId:   userId,
			Exchange: ex,
		},
		ExchangeOrder: &ExchangeOrder{
			UserId:   userId,
			Status:   "",
			Exchange: ex,
		},
		Withdrawal: &Withdrawal{
			UserId:   userId,
			Address:  address,
			Exchange: ex,
			Status:   "",
		},
	}

	w.Deposite.OrderId = w.Id
	w.Withdrawal.OrderId = w.Id
	w.ExchangeOrder.OrderId = w.Id
	return w
}

func genOrderId(l int) int64 {
	return utils.RandInt64(l)
}

func (o *UserOrder) AddDeposite(d *Deposit) {
	o.Deposite.Id = d.Id
	o.Deposite.Address = d.Address
	return
}

// implement stringer interface
func (o *UserOrder) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}
