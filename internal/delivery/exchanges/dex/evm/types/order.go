package types

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Order struct {
	*entity.ObjectId
	UserID string
	Status entity.OrderStatus
	ExNid  string
	In     entity.TokenId
	Out    entity.TokenId

	ApiKey string
	BusId  uint
	Level  uint

	NeedApprove bool
	Approved    bool
	Sender      common.Address
	Receiver    common.Address
	AmountIn    float64
	CreatedAT   int64
	UpdatedAt   int64
	ExpireAt    int64
}

func (o *Order) ID() *entity.ObjectId       { return o.ObjectId }
func (o *Order) SetId(id string)            { o.ObjectId = &entity.ObjectId{Prefix: entity.PrefOrder, Id: id} }
func (o *Order) Type() entity.OrderType     { return entity.EVMOrder }
func (o *Order) STATUS() entity.OrderStatus { return o.Status }
func (o *Order) ExchangeNid() string        { return o.ExNid }
func (o *Order) UserId() string             { return o.UserID }
func (o *Order) CreatedAt() int64           { return o.CreatedAT }
func (o *Order) Update()                    { o.UpdatedAt = time.Now().Unix() }
func (o *Order) Steps() uint                { return 1 }
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
