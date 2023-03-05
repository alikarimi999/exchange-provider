package entity

type OrderType uint

const (
	CEXOrder OrderType = iota
	EVMOrder
)

type Order interface {
	ID() *ObjectId
	SetId(string)
	Type() OrderType
}
