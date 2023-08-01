package types

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"time"
)

const (
	OCreated           = "Created"
	ODepositTxIdSetted = "DepositTxIdSetted"
	ODepositeConfimred = "DepositeConfimred"

	OFirstSwapCompleted  = "FirstSwapCompleted"
	OSecondSwapCompleted = "SecondSwapCompleted"

	OWithdrawalTracking  = "WithdrawalTracking"
	OWithdrawalConfirmed = "WithdrawalConfirmed"

	ODepositFailed    = "DepositFailed"
	OFirstSwapFailed  = "FirstSwapFailed"
	OSecondSwapFailed = "SecondSwapFailed"
	OWithdrawalFailed = "WithdrawalFailed"
	OFailed           = "Failed"
	OExpired          = "Expired"
)

type Order struct {
	*entity.ObjectId
	UserID string
	Status string
	ExNid  string
	ExLp   uint

	ApiKey string
	BusId  uint
	Level  uint

	In  entity.TokenId
	Out entity.TokenId

	SetAmountIn               float64
	EstimateAmountOut         float64
	EstimateExchangeFeeAmount float64
	EstimateFeeAmount         float64
	InitialPrice              float64
	ExecutedPrice             float64

	Deposit    Deposit
	Swaps      map[int]*Swap
	Withdrawal *Withdrawal

	SpreadRate   float64
	SpreadAmount float64
	FeeRate      float64
	ExchangeFee  float64

	ExchangeFeeAmount    float64
	FeeAmount            float64
	FeeAndSpreadCurrency entity.TokenId

	Sender entity.Address

	FailedDesc string
	CreatedAT  int64
	UpdatedAt  int64
	ExpireAt   int64
}

func (o *Order) ID() *entity.ObjectId { return o.ObjectId }
func (o *Order) SetId(id string) {
	o.ObjectId = &entity.ObjectId{Prefix: entity.PrefOrder, Id: id}
}
func (o *Order) Type() entity.OrderType { return entity.CEXOrder }
func (o *Order) STATUS() entity.OrderStatus {

	switch o.Status {
	case OCreated:
		return entity.OCreated
	case ODepositFailed, OFirstSwapFailed, OSecondSwapFailed, OWithdrawalFailed, OFailed:
		return entity.OFailed
	case OWithdrawalConfirmed:
		return entity.OCompleted
	case OExpired:
		return entity.OExpired
	default:
		return entity.OPending
	}
}
func (o *Order) ExchangeNid() string { return o.ExNid }
func (o *Order) Update()             { o.UpdatedAt = time.Now().Unix() }
func (o *Order) UserId() string      { return o.UserID }
func (o *Order) CreatedAt() int64    { return o.CreatedAT }
func (o *Order) StepsCount() uint    { return 1 }
func (o *Order) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *Order) Expire() bool {
	if o.Status == OExpired ||
		(o.Status == OCreated && o.ExpireAt > 0 && time.Now().Unix() > o.ExpireAt) {
		o.Status = OExpired
		return true
	}
	return false
}
