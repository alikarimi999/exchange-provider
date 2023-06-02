package entity

type ExType string

const (
	CEX      ExType = "CEX"
	EvmDEX   ExType = "EvmDEX"
	CrossDex ExType = "CrossDex"
)

type EstimateAmount struct {
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
}

type Exchange interface {
	Id() uint
	Name() string
	NID() string
	EnableDisable(enable bool)
	IsEnable() bool
	Type() ExType
	NewOrder(interface{}, *APIToken) (Order, error)
	EstimateAmountOut(t1, t2 TokenId, amount float64, lvl uint) (*EstimateAmount, error)
	AddPairs(data interface{}) (*AddPairsResult, error)
	RemovePair(t1, t2 TokenId) error
	Configs() interface{}
	Remove()
}
