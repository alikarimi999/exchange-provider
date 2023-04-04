package dto

import (
	"exchange-provider/internal/entity"
)

type OrderStep struct {
	OrderId     string `json:"orderId"`
	Type        string `json:"type,omitempty"`
	TotalSteps  int    `json:"totalSteps"`
	CurrentStep int    `json:"currentStep"`
}

type SingleStep struct {
	*OrderStep
	Duration string `json:"duration"`
	Token    Token  `json:"token"`
	Address  string `json:"address"`
	Tag      string `json:"tag"`

	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
	ExpireAt  int64 `json:"expireAt"`
}

func SingleStepResponse(o *entity.CexOrder) *SingleStep {
	return &SingleStep{
		OrderStep: &OrderStep{OrderId: o.ObjectId.String(), CurrentStep: 1, TotalSteps: 1},
		Duration:  o.Swaps[0].Duration,
		Token:     tokenFromEntity(o.Routes[0].In, false),
		Address:   o.Deposit.Address.Addr,
		Tag:       o.Deposit.Address.Tag,

		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
		ExpireAt:  o.Deposit.ExpireAt,
	}
}

type multiStep struct {
	*OrderStep
	Transaction interface{} `json:"transaction"`
}

func MultiStep(oId, sender string, tx entity.Tx, step, steps int) *multiStep {
	ms := &multiStep{
		OrderStep: &OrderStep{OrderId: oId, CurrentStep: step, TotalSteps: steps},
	}
	switch tx.Type() {
	case entity.Evm:
		ms.Type = string(tx.Type())
		ms.Transaction = evmTx(tx, sender)
	}
	return ms
}
