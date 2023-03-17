package dto

import (
	"exchange-provider/internal/entity"

	"github.com/ethereum/go-ethereum/core/types"
)

type StepType string

const (
	evmStepType StepType = "EVM"
)

type OrderStep struct {
	OrderId     string   `json:"orderId"`
	Type        StepType `json:"type,omitempty"`
	TotalSteps  int      `json:"totalSteps"`
	CurrentStep int      `json:"currentStep"`
}

type SingleStep struct {
	*OrderStep
	Token   string `json:"token"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

func SingleStepResponse(o *entity.CexOrder) *SingleStep {
	return &SingleStep{
		OrderStep: &OrderStep{OrderId: o.ObjectId.String(), CurrentStep: 1, TotalSteps: 1},
		Token:     o.Routes[0].In.String(),
		Address:   o.Deposit.Address.Addr,
		Tag:       o.Deposit.Address.Tag,
	}
}

type multiStep struct {
	*OrderStep
	IsApproveTx bool        `json:"isApproveTx"`
	Transaction interface{} `json:"transaction"`
}

func MultiStep(oId, sender string, tx interface{}, step, steps int, isApprove bool) *multiStep {
	ms := &multiStep{
		OrderStep:   &OrderStep{OrderId: oId, CurrentStep: step, TotalSteps: steps},
		IsApproveTx: isApprove,
	}
	switch t := tx.(type) {
	case *types.Transaction:
		ms.Type = evmStepType
		ms.Transaction = evmTx(t, sender)
	}
	return ms
}
