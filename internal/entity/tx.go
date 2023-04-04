package entity

import "github.com/ethereum/go-ethereum/core/types"

type TxType string

var (
	Evm TxType = "EVM"
)

type Tx interface {
	Type() TxType
}

type EvmTx struct {
	Tx          *types.Transaction
	IsApproveTx bool
}

func (e *EvmTx) Type() TxType { return Evm }
