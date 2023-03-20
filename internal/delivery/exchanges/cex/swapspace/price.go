package swapspace

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"net/http"
	"net/url"
)

type estimateAmount struct {
	Partner      string  `json:"partner"`
	FromAmount   float64 `json:"fromAmount"`
	ToAmount     float64 `json:"toAmount"`
	FromCurrency string  `json:"fromCurrency"`
	FromNetwork  string  `json:"fromNetwork"`
	ToCurrency   string  `json:"toCurrency"`
	ToNetwork    string  `json:"toNetwork"`
	SupportRate  float64 `json:"supportRate"`
	Duration     string  `json:"duration"`
	Fixed        bool    `json:"fixed"`
	Min          float64 `json:"min"`
	Max          float64 `json:"max"`
	Exists       bool    `json:"exists"`
	ID           string  `json:"id"`
}

func (ex *exchange) EstimateAmountOut(from, to *entity.Token, amount float64) (float64, float64, error) {
	p, in, out, err := ex.retrieveInOut(from, to)
	if err != nil {
		return 0, 0, err
	}

	eAmounts, err := ex.estimateAmounts(in, out, amount)
	if err != nil {
		return 0, 0, err
	}

	ea, max, min, err := ex.price(eAmounts, amount)
	if err != nil {
		if from.Equal(p.T1) {
			if min != p.T1.Min || max != p.T1.Max {
				p.T1.Min = min
				p.T1.Max = max
				ex.pairs.Update(ex.Id(), p)
			} else if min != p.T2.Min || max != p.T2.Max {
				p.T2.Min = min
				p.T2.Max = max
				ex.pairs.Update(ex.Id(), p)
			}
		}
		return 0, min, err
	}

	return ea.ToAmount, 0, nil
}

func (ex *exchange) price(eAmounts []*estimateAmount, amount float64) (esa *estimateAmount, max, min float64, err error) {
	eas := []*estimateAmount{}
	for _, eAmount := range eAmounts {
		if eAmount.Exists && eAmount.ToAmount > 0 && eAmount.SupportRate >= 2 {
			eas = append(eas, eAmount)
		}
	}

	var amountF float64 = 0
	eaMin := &estimateAmount{}
	if len(eas) > 0 {
		for _, ea := range eas {
			if (ea.Min == 0 || amount >= ea.Min) &&
				(ea.Max == 0 || amount <= ea.Max) && ea.ToAmount > amountF {
				amountF = ea.ToAmount
				eaMin = ea
			}
		}
		return eaMin, 0, 0, nil
	}

	eas = []*estimateAmount{}
	for _, ea := range eAmounts {
		if ea.Exists && ea.SupportRate >= 2 {
			eas = append(eas, ea)
		}
	}

	eaMin = eas[0]
	eaMax := eas[0]
	for _, ea := range eas {
		if ea.Min < eaMin.Min {
			eaMin = ea
		}
	}

	for _, ea := range eas {
		if ea.Max == 0 {
			eaMax = ea
			break
		}
		if ea.Max > eaMax.Max {
			eaMax = ea
		}
	}

	return nil, eaMax.Max, eaMin.Min, errors.Wrap(errors.ErrNotFound)
}

func (ex *exchange) estimateAmounts(in, out *Token, amountIn float64) ([]*estimateAmount, error) {
	urlStr, _ := url.JoinPath(baseUrl, "/amounts")

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("fromCurrency", in.Code)
	v.Add("fromNetwork", in.Network)
	v.Add("toCurrency", out.Code)
	v.Add("toNetwork", out.Network)
	v.Add("amount", fmt.Sprintf("%v", amountIn))
	v.Add("fixed", "true")
	v.Add("float", "true")
	u.RawQuery = v.Encode()

	b, err := ex.request(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	eAmounts := []*estimateAmount{}
	if err := json.Unmarshal(b, &eAmounts); err != nil {
		return nil, err
	}

	return eAmounts, nil
}
