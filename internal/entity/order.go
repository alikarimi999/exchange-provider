package entity

import (
	"encoding/json"
	"time"
)

type OrderStatus string

const (
	OSNew OrderStatus = ""

	OSTxIdSetted        OrderStatus = "txId_setted"
	OSDepositeConfimred OrderStatus = "deposite_confirmed"

	OSWaitForExchangeOrderConfirm OrderStatus = "wait_for_exchange_order_confirm"
	OSExchangeOrderConfirmed      OrderStatus = "exchange_order_confirmed"

	OSWaitForWithdrawalConfirm OrderStatus = "wait_for_withdrawal_confirm"
	OSWithdrawalConfirmed      OrderStatus = "withdrawal_confirmed"

	OSSucceed OrderStatus = "succeed"
	OSFailed  OrderStatus = "failed"
)

const (
	FCInternalError int64 = iota + 1
	FCDepositFailed
	FCExOrdFailed
	FCWithdFailed
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
	FailedCode    int64
	FailedDesc    string
}

func NewOrder(userId int64, wAddress, dAddress *Address, bc, qc *Coin, side string, ex string) *UserOrder {
	w := &UserOrder{
		UserId:    userId,
		CreatedAt: time.Now().Unix(),
		Status:    OSNew,
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
			Side:     side,
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
		w.Withdrawal.Coin = bc
	} else {
		w.Deposit.Coin = bc
		w.Withdrawal.Coin = qc
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
