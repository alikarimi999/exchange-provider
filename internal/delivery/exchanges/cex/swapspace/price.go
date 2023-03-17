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

func (ex *exchange) EstimateAmountOut(t1, t2 *entity.Token, amount float64) (float64, float64, error) {
	ea, min, err := ex.price(&pair{t1: fromEntity(t1),
		t2: fromEntity(t2)}, amount)
	if err != nil {
		return 0, min, err
	}
	return ea.ToAmount, 0, nil
}

func (ex *exchange) price(p *pair, amount float64) (*estimateAmount, float64, error) {
	eAmounts, err := ex.estimateAmounts(p, amount)
	if err != nil {
		return nil, 0, err
	}

	eas := []*estimateAmount{}
	for _, eAmount := range eAmounts {
		if eAmount.Exists && eAmount.ToAmount > 0 && eAmount.SupportRate >= 2 {
			eas = append(eas, eAmount)
		}
	}

	var amountF float64 = 0
	eAmount := &estimateAmount{}
	if len(eas) > 0 {
		for _, ea := range eas {
			if (ea.Min == 0 || amount >= ea.Min) &&
				(ea.Max == 0 || amount <= ea.Max) && ea.ToAmount > amountF {
				amountF = ea.ToAmount
				eAmount = ea
			}
		}
		return eAmount, 0, nil
	}

	eas = []*estimateAmount{}
	for _, ea := range eAmounts {
		if ea.Exists && ea.SupportRate >= 2 {
			eas = append(eas, ea)
		}
	}

	eAmount = eas[0]
	for _, ea := range eas {
		if ea.Min < eAmount.Min {
			eAmount = ea
		}
	}

	return nil, eAmount.Min, errors.Wrap(errors.ErrNotFound)
}

func (ex *exchange) estimateAmounts(p *pair, amountIn float64) ([]*estimateAmount, error) {
	urlStr, _ := url.JoinPath(baseUrl, "/amounts")

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("fromCurrency", p.t1.Code)
	v.Add("fromNetwork", p.t1.Network)
	v.Add("toCurrency", p.t2.Code)
	v.Add("toNetwork", p.t2.Network)
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
