package entity

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type EvmOrder struct {
	Id        string
	UserId    uint64 `bson:"userId"`
	Status    OrderStatus
	Steps     map[uint]*EvmStep
	Sender    common.Address
	Receiver  common.Address
	AmountIn  float64 `bson:"amountIn"`
	FeeRate   float64 `bson:"feeRate"`
	CreatedAt int64   `bson:"createdAt"`
}

func NewEvmOrder(userId uint64, steps map[uint]*EvmStep, sender, receiver common.Address,
	amountIn, feeRate float64) *EvmOrder {

	return &EvmOrder{
		UserId:    userId,
		Steps:     steps,
		Status:    ONew,
		Sender:    sender,
		Receiver:  receiver,
		AmountIn:  amountIn,
		FeeRate:   feeRate,
		CreatedAt: time.Now().UTC().Unix(),
	}
}

func (o *EvmOrder) ID() string      { return o.Id }
func (o *EvmOrder) SetId(id string) { o.Id = id }
func (o *EvmOrder) Type() OrderType { return EVMOrder }
