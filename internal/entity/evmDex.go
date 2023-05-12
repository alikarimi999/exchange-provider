package entity

type EVMDex interface {
	Exchange
	Chain() string
	CreateTx(Order) (Tx, error)
}
