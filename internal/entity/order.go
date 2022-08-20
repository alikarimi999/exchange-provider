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

	OrderStatusSucceed OrderStatus = "succeed"
	OrderStatusFailed  OrderStatus = "failed"
)

type UserOrder struct {
	Id        int64
	UserId    int64
	Seq       int64
	CreatedAt int64
	Status    OrderStatus
	Exchange  string
	BC        *Coin
	QC        *Coin
	Side      string // buy or sell

	SpreadVol     string
	SpreadRate    string
	Deposite      *Deposit
	ExchangeOrder *ExchangeOrder
	Withdrawal    *Withdrawal
	Broken        bool
	BreakReason   string
}

func NewOrder(userId int64, address string, bc, qc *Coin, side, ex string) *UserOrder {
	w := &UserOrder{
		UserId:    userId,
		CreatedAt: time.Now().Unix(),
		Status:    OrderStatusOpen,
		Exchange:  ex,
		BC:        bc,
		QC:        qc,
		Side:      side,
		Deposite:  &Deposit{},
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

	w.Withdrawal.OrderId = w.Id
	w.ExchangeOrder.OrderId = w.Id
	return w
}

func (o *UserOrder) AddDepositeConfirm(id int64, vol string) {
	o.Status = OrderStatusDepositeConfimred
	o.Deposite = &Deposit{
		Id:         id,
		UserId:     o.UserId,
		OrderId:    o.Id,
		Exchange:   o.Exchange,
		Volume:     vol,
		Fullfilled: true,
		Address:    "",
	}

}

func genOrderId(l int) int64 {
	return utils.RandInt64(l)
}

// implement stringer interface
func (o *UserOrder) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}
