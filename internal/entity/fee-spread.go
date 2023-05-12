package entity

type FeeTable interface {
	GetBusFee(busId string) float64
	UpdateBusFee(busId string, feeRate float64) error
	ChangeDefaultFee(float64) error
	GetDefaultFee() float64
	GetAllBusFees() map[string]float64
}

type Spread struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	Rate  float64 `json:"rate"`
}
type SpreadTable interface {
	Add(map[uint][]*Spread) (map[uint][]*Spread, error)
	Remove(levels []uint) error
	Levels() []uint
	GetByPrice(lvl uint, price float64) float64
	GetAll() map[uint][]*Spread
}
