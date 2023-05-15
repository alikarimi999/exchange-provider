package kucoin

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) orderFeeRat(p *entity.Pair) error {
	ep := p.EP.(*ExchangePair)
	if ep.HasIntermediaryCoin {
		var (
			f1, f2 float64
			err    error
		)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			bc := p.T1.ET.(*Token)
			qc := ep.IC1
			f1, err = k.orderFeeRate(bc.Currency, qc.Currency)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			bc := p.T2.ET.(*Token)
			qc := ep.IC1
			f2, err = k.orderFeeRate(bc.Currency, qc.Currency)
		}()
		wg.Wait()
		if err != nil {
			return err
		}

		ep.KucoinFeeRate1 = f1
		ep.KucoinFeeRate2 = f2
		return nil
	}
	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)
	f0, err := k.orderFeeRate(bc.Currency, qc.Currency)
	ep.KucoinFeeRate1 = f0
	return err
}

func (k *kucoinExchange) orderFeeRate(bc, qc string) (float64, error) {
	agent := k.agent("orderFeeRate")
	res, err := k.readApi.ActualFee(bc + "-" + qc)
	if err := handleSDKErr(err, res); err != nil {
		k.l.Debug(agent, err.Error())
		return 0, err
	}

	m := kucoin.TradeFeesResultModel{}
	err = res.ReadData(&m)
	if err != nil {
		k.l.Debug(agent, err.Error())
		return 0, err
	}

	f, err := strconv.ParseFloat(m[0].TakerFeeRate, 64)
	if err != nil {
		k.l.Debug(agent, err.Error())
		return 0, err
	}
	return f, nil
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

func (k *kucoinExchange) isDipositAndWithdrawEnable(t *Token) (bool, bool, error) {
	agent := k.agent("isDipositAndWithdrawEnable")

	res, err := k.readApi.CurrencyV2(t.Currency, t.Chain)
	if err := handleSDKErr(err, res); err != nil {
		k.l.Error(agent, err.Error())
		return false, false, err
	}

	m := kucoin.CurrencyV2Model{}
	err = res.ReadData(&m)
	if err != nil {
		k.l.Error(agent, err.Error())
		return false, false, err
	}

	for _, c := range m.Chains {
		if c.ChainName == string(t.ChainName) {
			t.ConfirmBlocks = c.Confirms
			t.MinWithdrawalSize, _ = strconv.ParseFloat(c.WithdrawalMinSize, 64)
			t.MinWithdrawalFee, _ = strconv.ParseFloat(c.WithdrawalMinFee, 64)
			return c.IsDepositEnabled, c.IsWithdrawEnabled, nil
		}
	}
	return false, false, errors.Wrap(errors.ErrNotFound)
}

func (k *kucoinExchange) setWithdrawalLimit(et *entity.Token) error {
	agent := k.agent("setWithdrawalLimit")
	t := et.ET.(*Token)
	res, err := k.readApi.CurrencyV2(t.Currency, t.Chain)
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	m := kucoin.CurrencyV2Model{}
	if err := res.ReadData(&m); err != nil {
		return handleSDKErr(err, res)
	}

	for _, c := range m.Chains {
		if c.IsDepositEnabled && c.IsWithdrawEnabled && c.ChainName == string(t.ChainName) {
			t.ConfirmBlocks = c.Confirms
			t.MinWithdrawalSize, _ = strconv.ParseFloat(c.WithdrawalMinSize, 64)
			t.MinWithdrawalFee, _ = strconv.ParseFloat(c.WithdrawalMinFee, 64)
			et.ContractAddress = c.ContractAddress

			res, err = k.readApi.WithdrawalQuotas(t.Currency, t.Chain)
			if err := handleSDKErr(err, res); err != nil {
				k.l.Debug(agent, err.Error())
				return err
			}

			wq := &kucoin.WithdrawalQuotasModel{}
			if err := res.ReadData(wq); err != nil {
				k.l.Debug(agent, err.Error())
				return handleSDKErr(err, res)
			}
			t.WithdrawalPrecision = int(wq.Precision)
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

func (k *kucoinExchange) getAddress(t *Token) error {
	agent := k.agent("getAddress")
	res, err := k.readApi.DepositAddresses(t.Currency, t.Chain)
	if err := handleSDKErr(err, res); err != nil {
		k.l.Debug(agent, err.Error())
		return err
	}
	da := &kucoin.DepositAddressModel{}
	if err := res.ReadData(da); err != nil {
		k.l.Debug(agent, err.Error())
		return handleSDKErr(err, res)
	}
	t.DepositAddress = da.Address
	t.DepositTag = da.Memo
	return nil
}

func (k *kucoinExchange) setInfos(p *entity.Pair) error {
	var err error
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = k.setWithdrawalLimit(p.T1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = k.setWithdrawalLimit(p.T2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = k.getAddress(p.T1.ET.(*Token))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = k.getAddress(p.T2.ET.(*Token))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = k.orderFeeRat(p)
	}()

	wg.Wait()
	if err != nil {
		return err
	}

	if err := k.minAndMax(p); err != nil {
		return err
	}
	return nil
}
