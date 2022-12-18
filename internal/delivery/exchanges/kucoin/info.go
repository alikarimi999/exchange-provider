package kucoin

import (
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) getPrice(p *pair) (string, string, error) {
	res, err := k.api.TickerLevel1(p.BC.TokenId + "-" + p.QC.TokenId)
	if err := handleSDKErr(err, res); err != nil {
		k.l.Error(fmt.Sprintf("%s.setPrice", k.Id()), err.Error())
		return "", "", err
	}

	t := &kucoin.TickerLevel1Model{}
	err = res.ReadData(t)
	if err != nil {
		k.l.Error("Kucoin.setPrice", err.Error())
		return "", "", err
	}

	f, err := numbers.StringToBigFloat(t.BestBid)
	if err != nil {
		return "", "", err
	}

	return t.BestAsk, new(big.Float).Quo(big.NewFloat(1), f).String(), nil
}

func (k *kucoinExchange) orderFeeRate(p *pair) string {
	res, err := k.api.ActualFee(p.BC.TokenId + "-" + p.QC.TokenId)
	if err := handleSDKErr(err, res); err != nil {
		k.l.Error("Kucoin.setOrderFeeRate", err.Error())
		return ""
	}

	m := kucoin.TradeFeesResultModel{}
	err = res.ReadData(&m)
	if err != nil {
		k.l.Error("Kucoin.setOrderFeeRate", err.Error())
		return ""
	}

	return m[0].TakerFeeRate

}

func (k *kucoinExchange) setBCWithdrawalLimit(p *pair) error {
	res, err := k.api.CurrencyV2(p.BC.TokenId, "")
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	m := kucoin.CurrencyV2Model{}
	err = res.ReadData(&m)
	if err != nil {
		return err
	}

	for _, c := range m.Chains {
		if c.ChainName == string(p.BC.Standard) {
			p.BC.ConfirmBlocks = c.Confirms
			p.BC.minWithdrawalSize = c.WithdrawalMinSize
			p.BC.minWithdrawalFee = c.WithdrawalMinFee
			return nil
		}
	}

	ch := []string{}
	for _, c := range m.Chains {
		ch = append(ch, c.ChainName)
	}

	return errors.Wrap(errors.ErrBadRequest, errors.Op("Kucoin.setBCWithdrawalLimit"),
		errors.NewMesssage(fmt.Sprintf("coin %s with chain %s not supported by kucoin,supported chains for %s is %+v",
			p.BC.TokenId, p.BC.Standard, p.BC.TokenId, ch)))
}

func (k *kucoinExchange) setQCWithdrawalLimit(p *pair) error {
	res, err := k.api.CurrencyV2(p.QC.TokenId, "")
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	m := kucoin.CurrencyV2Model{}
	err = res.ReadData(&m)
	if err != nil {
		return err
	}

	for _, c := range m.Chains {
		if c.ChainName == string(p.QC.Standard) {
			p.QC.ConfirmBlocks = c.Confirms
			p.QC.minWithdrawalSize = c.WithdrawalMinSize
			p.QC.minWithdrawalFee = c.WithdrawalMinFee
			return nil
		}
	}

	ch := []string{}
	for _, c := range m.Chains {
		ch = append(ch, c.ChainName)
	}

	return errors.Wrap(errors.ErrBadRequest, errors.Op("Kucoin.setBCWithdrawalLimit"),
		errors.NewMesssage(fmt.Sprintf("coin %s with chain %s not supported by kucoin,supported chains for %s is %+v",
			p.QC.TokenId, p.QC.Standard, p.QC.TokenId, ch)))
}

func (k *kucoinExchange) setAddress(pc *kuToken) error {
	op := errors.Op(fmt.Sprintf("%s.setChain", k.Id()))
	var coin string
	var chain string
	if pc.needChain {
		coin = pc.TokenId
		chain = string(pc.Standard)
	} else {
		coin = pc.TokenId
		chain = ""
	}

	res, err := k.api.DepositAddresses(coin, chain)
	if pc.needChain && res != nil && res.Code == "900014" && res.Message == "Invalid chainId" {
		pc.needChain = false
		return k.setAddress(pc)
	}
	a := &kucoin.DepositAddressModel{}
	if err := res.ReadData(a); err != nil {
		return err
	}
	pc.address = a.Address
	pc.tag = a.Memo

	k.l.Debug(string(op), fmt.Sprintf("`%s` address downloaded `%s:%s`", pc.String(), pc.address, pc.tag))

	if err = handleSDKErr(err, res); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (k *kucoinExchange) setInfos(p *pair) error {

	p.BC.needChain = true
	p.QC.needChain = true

	if err := k.setAddress(p.BC); err != nil {
		return err
	}

	if err := k.setAddress(p.QC); err != nil {
		return err
	}

	err := k.setBCWithdrawalLimit(p)
	if err != nil {
		return err
	}

	err = k.setQCWithdrawalLimit(p)
	if err != nil {
		return err
	}

	return nil
}
