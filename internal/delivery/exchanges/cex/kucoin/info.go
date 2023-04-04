package kucoin

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) orderFeeRate(bc, qc *token) float64 {
	res, err := k.readApi.ActualFee(bc.Currency + "-" + qc.Currency)
	if err := handleSDKErr(err, res); err != nil {
		k.l.Error("Kucoin.setOrderFeeRate", err.Error())
		return 0
	}

	m := kucoin.TradeFeesResultModel{}
	err = res.ReadData(&m)
	if err != nil {
		k.l.Error("Kucoin.setOrderFeeRate", err.Error())
		return 0
	}

	f, _ := strconv.ParseFloat(m[0].TakerFeeRate, 64)
	return f
}

type token struct {
	Currency        string `json:"currency"`
	ChainName       string `json:"chainName"`
	WalletPrecision string `json:"walletPrecision"`
	Chain           string `json:"chain"`
}

type tokens struct {
	Code string  `json:"code"`
	Msg  string  `json:"msg"`
	Data []token `json:"data"`
}

func (k *kucoinExchange) retreiveTokens() (*tokens, error) {
	url := "https://www.kucoin.com/_api/currency/currency/chain-info?lang=en_US"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	ts := &tokens{}
	if err := json.Unmarshal(b, ts); err != nil {
		return nil, err
	}
	if ts.Code != "200" {
		return nil, fmt.Errorf(ts.Msg)
	}
	return ts, nil
}

func (k *kucoinExchange) setWithdrawalLimit(et *entity.Token) error {
	t := et.ET.(*Token)
	res, err := k.readApi.CurrencyV2(t.Currency, "")
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	m := kucoin.CurrencyV2Model{}
	err = res.ReadData(&m)
	if err != nil {
		return err
	}

	for _, c := range m.Chains {
		if c.IsDepositEnabled && c.IsWithdrawEnabled && c.ChainName == string(t.ChainName) {
			t.ConfirmBlocks = c.Confirms
			t.MinWithdrawalSize, _ = strconv.ParseFloat(c.WithdrawalMinSize, 64)
			t.MinWithdrawalFee, _ = strconv.ParseFloat(c.WithdrawalMinFee, 64)
			et.ContractAddress = c.ContractAddress
			return nil
		}
	}

	ch := []string{}
	for _, c := range m.Chains {
		ch = append(ch, c.ChainName)
	}

	return errors.Wrap(errors.ErrBadRequest, errors.Op("Kucoin.setBCWithdrawalLimit"),
		errors.NewMesssage(fmt.Sprintf("coin %s with chain %s not supported by kucoin,supported chains for %s is %+v",
			t.Currency, t.ChainName, t.Currency, ch)))
}

func (k *kucoinExchange) setMinAndMax(p *entity.Pair) error {
	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)
	price, err := k.price(bc, qc)
	if err != nil {
		return err
	}

	minT1 := (qc.MinWithdrawalFee + qc.MinWithdrawalSize) * (1 / price)
	p.T1.Max = bc.MaxOrderSize
	minT2 := (bc.MinWithdrawalFee + bc.MinWithdrawalSize) * price
	p.T2.Max = qc.MaxOrderSize

	p.T1.Min = minT1 + (minT1 * 0.15)
	p.T2.Min = minT2 + (minT2 * 0.15)

	return nil
}

func (k *kucoinExchange) setInfos(p *entity.Pair) error {

	err := k.setWithdrawalLimit(p.T1)
	if err != nil {
		return err
	}

	err = k.setWithdrawalLimit(p.T2)
	if err != nil {
		return err
	}

	return k.setMinAndMax(p)
}
