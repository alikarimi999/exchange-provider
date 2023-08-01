package entity

type CrossDEX interface {
	Exchange
	CreateTx(o Order, step int) (Tx, error)
}
