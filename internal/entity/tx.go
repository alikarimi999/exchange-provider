package entity

import "github.com/ethereum/go-ethereum/core/types"

type TxType string

var (
	Evm TxType = "EVM"
)

type Developer struct {
	Function   string   `json:"function"`
	Contract   string   `json:"contract"`
	Parameters []string `json:"parameters"`
	Value      string   `json:"value"`
}

type Tx interface {
	Type() TxType
	Step() uint
	From() string
}

type EvmTx struct {
	Tx          *types.Transaction
	IsApproveTx bool
	CurrentStep uint
	Sender      string
	Developer   *Developer
}

func (e *EvmTx) Type() TxType { return Evm }
func (e *EvmTx) Step() uint   { return e.CurrentStep }
func (e *EvmTx) From() string { return e.Sender }
