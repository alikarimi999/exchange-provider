package swapspace

import (
	"exchange-provider/internal/entity"
)

func (ex *exchange) checkPair(p *entity.Pair) error {
	min, max, err := ex.minANDmax(p.T1.ET.(*Token), p.T2.ET.(*Token))
	if err != nil {
		return err
	}
	p.T1.Min = min
	p.T1.Max = max

	min, max, err = ex.minANDmax(p.T2.ET.(*Token), p.T1.ET.(*Token))
	if err != nil {
		return err
	}
	p.T2.Min = min
	p.T2.Max = max
	return nil
}

func (ex *exchange) minANDmax(in, out *Token) (float64, float64, error) {
	eAmounts, err := ex.estimateAmounts(in, out, 0.1)
	if err != nil {
		return 0, 0, err
	}

	eas := []*estimateAmount{}
	for _, ea := range eAmounts {
		if ea.Exists && ea.SupportRate >= 2 {
			eas = append(eas, ea)
		}
	}

	min := eas[0].Min
	max := eas[1].Max

	for _, ea := range eas {
		if ea.Min < min {
			min = ea.Min
		}

	}

	for _, ea := range eas {
		if ea.Max == 0 {
			max = 0
			break
		}

		if ea.Max > max {
			max = ea.Max
		}
	}

	return min, max, nil
}
