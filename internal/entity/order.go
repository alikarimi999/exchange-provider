package entity

import (
	"encoding/json"
	"sort"
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
	In       *Token
	Out      *Token
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

	Fee         string
	FeeCurrency string

	SpreadVol      string
	SpreadRate     string
	SpreadCurrency string

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
			Status:  "",
			Address: dAddress,
			Token:   routes[0].In,
		},

		Swaps: make(map[int]*Swap),

		Withdrawal: &Withdrawal{
			Address: wAddress,
			Status:  "",
			Token:   routes[len(routes)-1].Out,
		},
		MetaData: make(MetaData),
	}

	for i := range routes {
		o.Swaps[i] = &Swap{Status: SwapPending}
	}

	return o
}

func (o *Order) SortedRoutes() []*Route {
	indexes := []int{}
	for i := range o.Routes {
		indexes = append(indexes, i)
	}

	sort.Ints(indexes)
	routes := []*Route{}
	for _, i := range indexes {
		routes = append(routes, o.Routes[i])
	}
	return routes
}

// implement stringer interface
func (o *Order) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}
