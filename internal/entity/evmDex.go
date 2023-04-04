package entity

type EVMDex interface {
	Exchange
	Chain() string
	SetStpes(*DexOrder, *Route) error
	GetStep(o *DexOrder, step uint) (Tx, error)
}
