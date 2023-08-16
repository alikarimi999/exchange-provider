package entity

type ExType string

const (
	CEX      ExType = "CEX"
	EvmDEX   ExType = "EvmDEX"
	CrossDex ExType = "CrossDex"
)

type EstimateAmount struct {
	InUsd             float64
	OutUsd            float64
	Price             float64
	AmountIn          float64
	AmountOut         float64
	SpreadRate        float64
	FeeRate           float64
	FeeAmount         float64
	ExchangeFee       float64
	ExchangeFeeAmount float64
	FeeCurrency       TokenId
	P                 *Pair
	Data              interface{}
}

type Exchange interface {
	Id() uint
	Name() string
	NID() string
	EnableDisable(enable bool)
	IsEnable() bool
	Type() ExType
	NewOrder(interface{}, *APIToken) (Order, error)
	SetTxId(o Order, txId string) error
	EstimateAmountOut(t1, t2 TokenId, amount float64, lvl uint, opts interface{}) ([]*EstimateAmount, error)
	AddPairs(data interface{}) (*AddPairsResult, error)
	RemovePair(t1, t2 TokenId) error
	Configs() interface{}
	UpdateConfigs(cfgi interface{}, store ExchangeStore) error
	Remove()
	UpdateStatus(Order) error
}
