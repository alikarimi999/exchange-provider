package entity

type OrderType uint

const (
	CEXOrder OrderType = iota
	EVMOrder
)

type OrderStatus string

func (os OrderStatus) String() string {
	return string(os)
}

const (
	OCreated   OrderStatus = "created"
	OPending   OrderStatus = "pending"
	OCompleted OrderStatus = "completed"
	OExpired   OrderStatus = "expired"
	OFailed    OrderStatus = "failed"
)

type Order interface {
	ID() *ObjectId
	SetId(string)
	Type() OrderType
	STATUS() OrderStatus
	ExchangeNid() string
	Update()
	Steps() uint
	Expire() bool
	String() string
	UserId() string
	CreatedAt() int64
}
