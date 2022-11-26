package entity

type PairConfigs interface {
	GetDefaultSpread() string
	ChangeDefaultSpread(s float64) error
	GetPairSpread(bc, qc *Coin) string
	ChangePairSpread(bc, qc *Coin, s float64) error
	GetAllPairsSpread() map[string]float64
	ApplySpread(bc, qc *Coin, vol string) (appliedVol, spreadVol, spreadRate string, err error)

	PairMinDeposit(c1, c2 string) (float64, float64)

	ChangeMinDeposit(...*PairMinDeposit) error
	AllMinDeposit() []*PairMinDeposit
}

type PairMinDeposit struct {
	C1 *CoinMinDeposit
	C2 *CoinMinDeposit
}

type CoinMinDeposit struct {
	Coin string
	Min  float64
}
