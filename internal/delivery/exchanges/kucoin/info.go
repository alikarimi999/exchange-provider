package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) setPrice(p *entity.Pair) error {
	res, err := k.api.TickerLevel1(fmt.Sprintf("%s%s%s", p.C1.Coin.CoinId, pairDelimiter, p.C2.Coin.CoinId))
	if err := handleSDKErr(err, res); err != nil {
		k.l.Error(fmt.Sprintf("%s.setPrice", k.Id()), err.Error())
		return err
	}

	t := &kucoin.TickerLevel1Model{}
	err = res.ReadData(t)
	if err != nil {
		k.l.Error("Kucoin.setPrice", err.Error())
		return err
	}

	p.Price1 = t.BestAsk
	p.Price2 = t.BestBid

	return nil
}

func (k *kucoinExchange) setOrderFeeRate(p *entity.Pair) error {
	res, err := k.api.ActualFee(fmt.Sprintf("%s%s%s", p.C1.Coin.CoinId, pairDelimiter, p.C2.Coin.CoinId))
	if err := handleSDKErr(err, res); err != nil {
		k.l.Error("Kucoin.setOrderFeeRate", err.Error())
		return err
	}

	m := kucoin.TradeFeesResultModel{}
	err = res.ReadData(&m)
	if err != nil {
		k.l.Error("Kucoin.setOrderFeeRate", err.Error())
		return err
	}

	p.OrderFeeRate = m[0].TakerFeeRate

	return nil

}

func (k *kucoinExchange) setBCWithdrawalLimit(p *pair) error {
	res, err := k.api.CurrencyV2(p.BC.CoinId, "")
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	m := kucoin.CurrencyV2Model{}
	err = res.ReadData(&m)
	if err != nil {
		return err
	}

	for _, c := range m.Chains {
		if c.ChainName == p.BC.ChainId {
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
			p.BC.CoinId, p.BC.ChainId, p.BC.CoinId, ch)))
}

func (k *kucoinExchange) setQCWithdrawalLimit(p *pair) error {
	res, err := k.api.CurrencyV2(p.QC.CoinId, "")
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	m := kucoin.CurrencyV2Model{}
	err = res.ReadData(&m)
	if err != nil {
		return err
	}

	for _, c := range m.Chains {
		if c.ChainName == p.QC.ChainId {
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
			p.QC.CoinId, p.QC.ChainId, p.QC.CoinId, ch)))
}

func (k *kucoinExchange) setAddress(pc *kuCoin) error {
	op := errors.Op(fmt.Sprintf("%s.setChain", k.Id()))
	var coin string
	var chain string
	if pc.needChain {
		coin = pc.CoinId
		chain = pc.ChainId
	} else {
		coin = pc.CoinId
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
