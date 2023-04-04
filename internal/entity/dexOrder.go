package entity

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type DexOrder struct {
	*ObjectId
	UserId    string
	Status    string
	Steps     map[uint]*Step
	Sender    common.Address
	Receiver  common.Address
	AmountIn  float64
	FeeRate   float64
	CreatedAt int64
}

func NewDexOrder(userId string, steps map[uint]*Step, sender, receiver common.Address,
	amountIn, feeRate float64) *DexOrder {

	return &DexOrder{
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

func (o *DexOrder) ID() *ObjectId   { return o.ObjectId }
func (o *DexOrder) SetId(id string) { o.ObjectId = &ObjectId{Prefix: PrefOrder, Id: id} }
func (o *DexOrder) Type() OrderType { return EVMOrder }
