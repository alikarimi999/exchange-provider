package entity

type PairConfigs interface {
	GetDefaultSpread() string
	ChangeDefaultSpread(s float64) error
	GetPairSpread(bc, qc *Coin) string
	ChangePairSpread(bc, qc *Coin, s float64) error
	GetAllPairsSpread() map[string]float64
	ApplySpread(bc, qc *Coin, size string) (remainder, rate string, err error)

	PairMinDeposit(bc, qc *Coin) (minBc, minQc float64)

	ChangeMinDeposit(bc, qc *Coin, minBc, minQc float64) error
	AllMinDeposit() []*PairMinDeposit
}

type PairMinDeposit struct {
	Pair         string
	MinBaseCoin  float64
	MinQouteCoin float64
}
