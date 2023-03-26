package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"

	"github.com/Kucoin/kucoin-go-sdk"
)

// func (k *kucoinExchange) orderFeeRate(p *pair) string {
// 	res, err := k.readApi.ActualFee(p.BC.TokenId + "-" + p.QC.TokenId)
// 	if err := handleSDKErr(err, res); err != nil {
// 		k.l.Error("Kucoin.setOrderFeeRate", err.Error())
// 		return ""
// 	}

// 	m := kucoin.TradeFeesResultModel{}
// 	err = res.ReadData(&m)
// 	if err != nil {
// 		k.l.Error("Kucoin.setOrderFeeRate", err.Error())
// 		return ""
// 	}

// 	return m[0].TakerFeeRate

// }

func (k *kucoinExchange) setWithdrawalLimit(t *Token) error {
	res, err := k.readApi.CurrencyV2(t.TokenId, "")
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	m := kucoin.CurrencyV2Model{}
	err = res.ReadData(&m)
	if err != nil {
		return err
	}

	for _, c := range m.Chains {
		if c.ChainName == string(t.Network) {
			t.ConfirmBlocks = c.Confirms
			t.MinWithdrawalSize, _ = strconv.ParseFloat(c.WithdrawalMinSize, 64)
			t.MinWithdrawalFee, _ = strconv.ParseFloat(c.WithdrawalMinFee, 64)
			return nil
		}
	}

	ch := []string{}
	for _, c := range m.Chains {
		ch = append(ch, c.ChainName)
	}

	return errors.Wrap(errors.ErrBadRequest, errors.Op("Kucoin.setBCWithdrawalLimit"),
		errors.NewMesssage(fmt.Sprintf("coin %s with chain %s not supported by kucoin,supported chains for %s is %+v",
			t.TokenId, t.Network, t.TokenId, ch)))
}

func (k *kucoinExchange) setAddress(t *Token) error {
	op := errors.Op(fmt.Sprintf("%s.setChain", k.Name()))
	var coin string
	var chain string
	if t.NeedChain {
		coin = t.TokenId
		chain = string(t.Network)
	} else {
		coin = t.TokenId
		chain = ""
	}

	res, err := k.readApi.DepositAddresses(coin, chain)
	if t.NeedChain && res != nil && res.Code == "900014" && res.Message == "Invalid chainId" {
		t.NeedChain = false
		return k.setAddress(t)
	}
	a := &kucoin.DepositAddressModel{}
	if err := res.ReadData(a); err != nil {
		return err
	}
	t.Address = a.Address
	t.Tag = a.Memo

	if err = handleSDKErr(err, res); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
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

	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)

	bc.NeedChain = true
	qc.NeedChain = true

	if err := k.setAddress(bc); err != nil {
		return err
	}

	if err := k.setAddress(qc); err != nil {
		return err
	}

	err := k.setWithdrawalLimit(bc)
	if err != nil {
		return err
	}

	err = k.setWithdrawalLimit(qc)
	if err != nil {
		return err
	}

	return k.setMinAndMax(p)
}
