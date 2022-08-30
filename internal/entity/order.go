package entity

import (
	"encoding/json"
	"time"
)

type OrderStatus string

const (
	OrderStatusOpen OrderStatus = ""

	OrderStatusWaitForDepositeConfirm OrderStatus = "wait_for_deposite_confirm"
	OrderStatusDepositeConfimred      OrderStatus = "deposite_confirmed"
	OsDepositFailed                   OrderStatus = "deposit_failed"

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
	Deposit       *Deposit
	ExchangeOrder *ExchangeOrder
	Withdrawal    *Withdrawal
	Broken        bool
	BreakReason   string
}

func NewOrder(userId int64, wAddress, dAddress *Address, bc, qc *Coin, side string, ex string) *UserOrder {
	w := &UserOrder{
		UserId:    userId,
		CreatedAt: time.Now().Unix(),
		Status:    OrderStatusOpen,
		Exchange:  ex,
		BC:        bc,
		QC:        qc,
		Side:      side,
		Deposit: &Deposit{
			UserId:   userId,
			Status:   "",
			Exchange: ex,
			Address:  dAddress,
		},

		ExchangeOrder: &ExchangeOrder{
			UserId:   userId,
			Status:   "",
			Exchange: ex,
		},
		Withdrawal: &Withdrawal{
			UserId:   userId,
			Address:  wAddress,
			Exchange: ex,
			Status:   "",
		},
	}

	if side == "buy" {
		w.Deposit.Coin = qc
	} else {
		w.Deposit.Coin = bc
	}

	w.Withdrawal.OrderId = w.Id
	w.ExchangeOrder.OrderId = w.Id
	return w
}

// implement stringer interface
func (o *UserOrder) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}
