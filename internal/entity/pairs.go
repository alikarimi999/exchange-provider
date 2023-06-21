package entity

type ExchangePair interface {
	Snapshot() ExchangePair
}

type Pair struct {
	T1 *Token `json:"t1"`
	T2 *Token `json:"t2"`

	FeeRate1    float64 `json:"feeRate1"`
	FeeRate2    float64 `json:"feeRate2"`
	ExchangeFee float64 `json:"exchangeFee"`

	Spreads  map[uint]float64 `json:"spreads"`
	LP       uint             `json:"lp,omitempty"`
	Exchange string           `json:"exchange"`
	Enable   bool             `json:"enable"`
	EP       ExchangePair     `json:"ep"`
}

func (p *Pair) String() string {
	return p.T1.String() + "/" + p.T2.String()
}

func (p *Pair) Snapshot() *Pair {
	sp := make(map[uint]float64)
	if p.Spreads != nil {
		for k, v := range p.Spreads {
			sp[k] = v
		}
	}

	var ep ExchangePair
	if p.EP != nil {
		ep = p.EP.Snapshot()
	}
	return &Pair{
		T1:          p.T1.Snapshot(),
		T2:          p.T2.Snapshot(),
		FeeRate1:    p.FeeRate1,
		FeeRate2:    p.FeeRate2,
		ExchangeFee: p.ExchangeFee,
		Spreads:     sp,
		LP:          p.LP,
		Exchange:    p.Exchange,
		Enable:      p.Enable,
		EP:          ep,
	}
}

func (p *Pair) Spread(lvl uint) float64 {
	s, ok := p.Spreads[lvl]
	if ok {
		return s
	}
	return p.Spreads[0]
}
