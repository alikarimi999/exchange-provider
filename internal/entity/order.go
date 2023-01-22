package entity

type OrderType uint

const (
	CEXOrder OrderType = iota
	EVMOrder
)

type Order interface {
	ID() string
	SetId(string)
	Type() OrderType
}
