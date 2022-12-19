package entity

import (
	"encoding/json"
	"sort"
	"time"
)

const (
	OSTxIdSetted        string = "txId_setted"
	OSDepositeConfimred string = "deposite_confirmed"

	OSWaitForSwapConfirm string = "wait_for_swap_confirm"
	OSSwapConfirmed      string = "swap_confirmed"

	OSWaitForWithdrawalConfirm string = "wait_for_withdrawal_confirm"
	OSWithdrawalConfirmed      string = "withdrawal_confirmed"

	OSSucceed string = "succeed"
	OSFailed  string = "failed"
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
	Status    string

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
		Status:    "",
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
		o.Swaps[i] = &Swap{}
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
