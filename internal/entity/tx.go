package entity

import "math/big"

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
}

type EvmTx struct {
	Network     string
	TxData      []byte
	IsApproveTx bool
	CurrentStep uint
	From        string
	To          string
	Value       *big.Int
	Developer   *Developer
}

func (e *EvmTx) Type() TxType { return Evm }
func (e *EvmTx) Step() uint   { return e.CurrentStep }
