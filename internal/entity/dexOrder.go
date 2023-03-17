package entity

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type EvmOrder struct {
	*ObjectId
	UserId    string
	Status    string
	Steps     map[uint]*EvmStep
	Sender    common.Address
	Receiver  common.Address
	AmountIn  float64
	FeeRate   float64
	CreatedAt int64
}

func NewEvmOrder(userId string, steps map[uint]*EvmStep, sender, receiver common.Address,
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

func (o *EvmOrder) ID() *ObjectId   { return o.ObjectId }
func (o *EvmOrder) SetId(id string) { o.ObjectId = &ObjectId{Prefix: PrefOrder, Id: id} }
func (o *EvmOrder) Type() OrderType { return EVMOrder }
