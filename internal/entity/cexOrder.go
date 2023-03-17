package entity

import (
	"sort"
	"time"
)

const (
	ONew               = "new"
	OPending           = "pending"
	OConfimDeposit     = "confirming deposit"
	ODepositeConfimred = "deposit confirmed"

	OWaitForSwapConfirm = "confirming swap"
	OSwapConfirmed      = "swap confirmed"

	OWaitForWithdrawalConfirm = "confirming withdraw"
	OWithdrawalConfirmed      = "withdraw confirmed"

	OSucceeded = "succeeded"
	ORefunding = "refunding"
	ORefunded  = "refunded"
	OExpired   = "expired"
	OFailed    = "failed"
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
	Exchange uint
	ExType
}

type CexOrder struct {
	*ObjectId
	UserId string
	Status string

	Deposit *Deposit
	Refund  Address

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
	CreatedAt int64
	UpdatedAt int64
}

func NewCexOrder(userId string, refund, reciver Address,
	routes map[int]*Route, amount float64) *CexOrder {

	t := time.Now().Unix()
	o := &CexOrder{
		UserId: userId,
		Status: ONew,
		Routes: routes,

		Deposit: &Deposit{
			Status: "",
			Token:  routes[0].In,
			Volume: amount,
		},
		Refund: refund,
		Swaps:  make(map[int]*Swap),

		Withdrawal: &Withdrawal{
			Address: reciver,
			Status:  "",
			Token:   routes[len(routes)-1].Out,
		},
		MetaData:  make(MetaData),
		CreatedAt: t,
		UpdatedAt: t,
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
func (o *CexOrder) ID() *ObjectId { return o.ObjectId }

func (o *CexOrder) Type() OrderType {
	return CEXOrder
}

func (o *CexOrder) SetId(id string) {
	o.ObjectId = &ObjectId{Prefix: PrefOrder, Id: id}
}
