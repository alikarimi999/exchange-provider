package types

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"time"
)

type Route struct {
	In        entity.TokenId
	Out       entity.TokenId
	AmountIn  float64
	AmountOut float64

	EstimateAmountOut float64
	ExchangeNid       string
	NeedApprove       bool
}

type Step struct {
	Status     string
	Receiver   string
	SrcTxId    string
	DstTxId    string
	AmountIn   float64
	AmountOut  float64
	Routes     map[int]*Route
	FailedDesc string
}

type Order struct {
	*entity.ObjectId
	UserID string
	Status entity.OrderStatus
	ExNid  string

	AllBridgeId string

	In  entity.TokenId
	Out entity.TokenId

	ApiKey string
	BusId  uint
	Level  uint

	Nonce string
	Steps map[int]*Step

	Sender   string
	Receiver string

	AmountIn          float64
	EstimateAmountOut float64

	FeeRate           float64
	FeeAmount         float64
	ExchangeFee       float64
	ExchangeFeeAmount float64

	FeeCurrency entity.TokenId
	CreatedAT   int64
	UpdatedAt   int64
	ExpireAt    int64
}

func (o *Order) ID() *entity.ObjectId       { return o.ObjectId }
func (o *Order) SetId(id string)            { o.ObjectId = &entity.ObjectId{Prefix: entity.PrefOrder, Id: id} }
func (o *Order) Type() entity.OrderType     { return entity.EVMOrder }
func (o *Order) STATUS() entity.OrderStatus { return o.Status }
func (o *Order) StepsStatus() interface{} {
	ss := make(OrderStatus)
	for i, s := range o.Steps {
		ss[i+1] = &StepStatus{Status: s.Status, TxId: s.SrcTxId, FailedDesc: s.FailedDesc,
			AmountIn: Number(s.AmountIn), AmountOut: Number(s.AmountOut)}
	}
	return ss
}
func (o *Order) ExchangeNid() string { return o.ExNid }
func (o *Order) UserId() string      { return o.UserID }
func (o *Order) CreatedAt() int64    { return o.CreatedAT }
func (o *Order) Update()             { o.UpdatedAt = time.Now().Unix() }
func (o *Order) StepsCount() uint    { return uint(len(o.Steps)) }
func (o *Order) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *Order) Expire() bool {
	if o.Status == entity.OExpired ||
		(o.Status == entity.OCreated && o.ExpireAt > 0 && time.Now().Unix() > o.ExpireAt) {
		o.Status = entity.OExpired
		return true
	}
	return false
}
