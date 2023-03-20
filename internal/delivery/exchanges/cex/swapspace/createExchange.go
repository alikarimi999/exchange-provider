package swapspace

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type createExchangeReq struct {
	Partner       string  `json:"partner"`
	FromCurrency  string  `json:"fromCurrency"`
	FromNetwork   string  `json:"fromNetwork"`
	ToCurrency    string  `json:"toCurrency"`
	ToNetwork     string  `json:"toNetwork"`
	Address       string  `json:"address"`
	ExtraID       string  `json:"extraId"`
	Amount        float64 `json:"amount"`
	Fixed         bool    `json:"fixed"`
	Refund        string  `json:"refund"`
	RefundExtraId string  `json:"refundExtraId"`
	RateId        string  `json:"rateId"`
	UserIp        string  `json:"userIp"`
}

type exchangeResponse struct {
	ID         string `json:"id"`
	Timestamps struct {
		CreatedAt string `json:"createdAt"`
		ExpiresAt string `json:"expiresAt"`
	} `json:"timestamps"`
	From struct {
		Address         string `json:"address"`
		ExtraID         string `json:"extraId"`
		TransactionHash string `json:"transactionHash"`
	} `json:"from"`
	To struct {
		TransactionHash string `json:"transactionHash"`
	} `json:"to"`
	Rate          float64 `json:"rate"`
	Status        string  `json:"status"`
	Confirmations int     `json:"confirmations"`
}

func (ex *exchange) SetDepositddress(o *entity.CexOrder) error {
	return ex.createExchange(o)
}

func (ex *exchange) createExchange(o *entity.CexOrder) error {
	from := o.Routes[0].In
	to := o.Routes[0].Out

	p, in, out, err := ex.retrieveInOut(from, to)
	if err != nil {
		return err
	}

	eAmounts, err := ex.estimateAmounts(in, out, o.Deposit.Volume)
	if err != nil {
		return err
	}

	ea, max, min, err := ex.price(eAmounts, o.Deposit.Volume)
	if err != nil {
		if from.Equal(p.T1) {
			if min != p.T1.Min || max != p.T1.Max {
				p.T1.Min = min
				p.T1.Max = max
				ex.pairs.Update(ex.Id(), p)
				return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("min amount is %f", min)))
			}
		} else if min != p.T2.Min || max != p.T2.Max {
			p.T2.Min = min
			p.T2.Max = max
			ex.pairs.Update(ex.Id(), p)
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("min amount is %f", min)))
		}
		return err
	}

	cer := &createExchangeReq{
		Partner:       ea.Partner,
		FromCurrency:  in.Code,
		FromNetwork:   in.Network,
		ToCurrency:    out.Code,
		ToNetwork:     out.Network,
		Address:       o.Withdrawal.Addr,
		ExtraID:       o.Withdrawal.Address.Tag,
		Amount:        o.Deposit.Volume,
		Fixed:         ea.Fixed,
		Refund:        o.Refund.Addr,
		RefundExtraId: o.Refund.Tag,
		RateId:        "",
		UserIp:        "0.0.0.0",
	}

	url, _ := url.JoinPath(baseUrl, "/exchange")

	b, err := ex.request(http.MethodPost, url, cer)
	if err != nil {
		return err
	}

	er := &exchangeResponse{}
	if err := json.Unmarshal(b, er); err != nil {
		return err
	}

	t, err := time.Parse(time.RFC3339, fmt.Sprintf("%sZ", er.Timestamps.ExpiresAt))
	if err == nil {
		o.Deposit.ExpireAt = t.Unix()
	}

	o.Deposit.Address.Addr = er.From.Address
	o.Deposit.Address.Tag = er.From.ExtraID
	o.Swaps[0].Duration = ea.Duration
	o.MetaData["id_in_swapspace"] = er.ID
	return nil
}
