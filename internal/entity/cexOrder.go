package entity

import (
	"sort"
	"time"
)

type OrderStatus uint

const (
	ONew OrderStatus = iota
	OConfimDeposit
	ODepositeConfimred

	OWaitForSwapConfirm
	OSwapConfirmed

	OWaitForWithdrawalConfirm
	OWithdrawalConfirmed

	Oucceeded
	OFailed
)

func (s OrderStatus) String() string {
	switch s {
	case ONew:
		return ""
	case OConfimDeposit:
		return "confirming transaction"
	case ODepositeConfimred:
		return "deposit confimred"
	case OWaitForSwapConfirm:
		return "cinfirmin swap"
	case OSwapConfirmed:
		return "swap confirmed"
	case OWaitForWithdrawalConfirm:
		return "confirming withdraw"
	case OWithdrawalConfirmed:
		return "withdrawal confirmed"
	case Oucceeded:
		return "succeeded"
	case OFailed:
		return "failed"
	default:
		return ""
	}
}

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
	ExType   `bson:"exType"`
}

type CexOrder struct {
	Id        string
	UserId    uint64 `bson:"userId"`
	CreatedAt int64  `bson:"createdAt"`
	Status    OrderStatus

	Deposit *Deposit

	Routes map[int]*Route
	Swaps  map[int]*Swap

	Withdrawal *Withdrawal

	Fee         string
	FeeCurrency string `bson:"feeCurrency"`

	SpreadVol      string `bson:"spreadVol"`
	SpreadRate     string `bson:"spreadRate"`
	SpreadCurrency string `bson:"spreadCurrency"`

	FailedCode int64  `bson:"failedCode"`
	FailedDesc string `bson:"failedDesc"`
	MetaData   `bson:"metaData"`
}

func NewOrder(userId uint64, wAddress, dAddress *Address, routes map[int]*Route) *CexOrder {
	o := &CexOrder{
		UserId:    userId,
		CreatedAt: time.Now().UTC().Unix(),
		Status:    ONew,
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

func (o *CexOrder) SortedRoutes() []*Route {
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
func (o *CexOrder) ID() string { return o.Id }

func (o *CexOrder) Type() OrderType {
	return CEXOrder
}

func (o *CexOrder) SetId(id string) {
	o.Id = id
	o.Deposit.OrderId = id
	for _, s := range o.Swaps {
		s.OrderId = id
	}
	o.Withdrawal.OrderId = id
}
