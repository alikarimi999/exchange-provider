package entity

import "github.com/ethereum/go-ethereum/core/types"

type EVMDex interface {
	Exchange
	SetStpes(*EvmOrder, *Route) error
	GetStep(o *EvmOrder, step uint) (*types.Transaction, error)
}
