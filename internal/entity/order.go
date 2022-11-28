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

type Route struct {
	In       *Coin
	Out      *Coin
	Exchange string
}

type Order struct {
	Id     int64
	UserId int64

	CreatedAt int64
	Status    OrderStatus

	Deposit *Deposit

	Routes map[int]*Route
	Swaps  map[int]*Swap

	Withdrawal *Withdrawal

	SpreadVol  string
	SpreadRate string

	FailedCode int64
	FailedDesc string
	MetaData
}

func NewOrder(userId int64, wAddress, dAddress *Address, routes map[int]*Route) *Order {
	o := &Order{
		UserId:    userId,
		CreatedAt: time.Now().Unix(),
		Status:    OSNew,
		Routes:    routes,

		Deposit: &Deposit{
			UserId:   userId,
			Status:   "",
			Exchange: routes[0].Exchange,
			Address:  dAddress,
			Coin:     routes[0].In,
		},

		Swaps: make(map[int]*Swap),

		Withdrawal: &Withdrawal{
			UserId:   userId,
			Address:  wAddress,
			Exchange: routes[len(routes)-1].Exchange,
			Status:   "",
			Coin:     routes[len(routes)-1].Out,
		},
		MetaData: make(MetaData),
	}

	for i := range o.Routes {
		o.Swaps[i] = &Swap{
			UserId: o.UserId,
		}
	}

	return o
}

// implement stringer interface
func (o *Order) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}
