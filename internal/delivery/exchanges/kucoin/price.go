package kucoin

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) setPrice(p *entity.Pair) error {
	res, err := k.api.TickerLevel1(fmt.Sprintf("%s%s%s", p.BC.Coin.CoinId, pairDelimiter, p.QC.Coin.CoinId))
	if err := handleSDKErr(err, res); err != nil {
		k.l.Error("Kucoin.setPrice", err.Error())
		return err
	}

	t := &kucoin.TickerLevel1Model{}
	err = res.ReadData(t)
	if err != nil {
		k.l.Error("Kucoin.setPrice", err.Error())
		return err
	}

	p.BestAsk = t.BestAsk
	p.BestBid = t.BestBid

	return nil
}

func (k *kucoinExchange) setOrderFeeRate(p *entity.Pair) error {
	res, err := k.api.ActualFee(fmt.Sprintf("%s%s%s", p.BC.Coin.CoinId, pairDelimiter, p.QC.Coin.CoinId))
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

func (k *kucoinExchange) setBCWithdrawalLimit(p *entity.Pair) error {
	res, err := k.api.CurrencyV2(p.BC.Coin.CoinId, "")
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
			p.BC.MinWithdrawalSize = c.WithdrawalMinSize
			p.BC.WithdrawalMinFee = c.WithdrawalMinFee
			p.BC.WithdrawalPrecision = int(m.Precision)
			return nil
		}
	}

	ch := []string{}
	for _, c := range m.Chains {
		ch = append(ch, c.ChainName)
	}

	return errors.Wrap(errors.ErrBadRequest, errors.Op("Kucoin.setBCWithdrawalLimit"),
		errors.NewMesssage(fmt.Sprintf("coin %s with chain %s not supported by kucoin,supported chains for %s is %+v", p.BC.Coin.CoinId, p.BC.ChainId, p.BC.CoinId, ch)))
}

func (k *kucoinExchange) setQCWithdrawalLimit(p *entity.Pair) error {
	res, err := k.api.CurrencyV2(p.QC.Coin.CoinId, "")
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
			p.QC.MinWithdrawalSize = c.WithdrawalMinSize
			p.QC.WithdrawalMinFee = c.WithdrawalMinFee
			p.QC.WithdrawalPrecision = int(m.Precision)
			return nil
		}
	}

	ch := []string{}
	for _, c := range m.Chains {
		ch = append(ch, c.ChainName)
	}

	return errors.Wrap(errors.ErrBadRequest, errors.Op("Kucoin.setBCWithdrawalLimit"),
		errors.NewMesssage(fmt.Sprintf("coin %s with chain %s not supported by kucoin,supported chains for %s is %+v", p.QC.Coin.CoinId, p.QC.ChainId, p.QC.CoinId, ch)))
}

func (k *kucoinExchange) setChain(pc *entity.PairCoin) error {
	const op = errors.Op("Kucoin.setChain")
	var coin string
	var chain string
	if pc.SetChain {
		coin = pc.CoinId
		chain = pc.ChainId
	} else {
		coin = pc.CoinId
		chain = ""
	}

	res, err := k.api.DepositAddresses(coin, chain)
	if pc.SetChain && res != nil && res.Code == "900014" && res.Message == "Invalid chainId" {
		pc.SetChain = false
		return k.setChain(pc)
	}

	if err = handleSDKErr(err, res); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (k *kucoinExchange) setInfos(p *entity.Pair) error {

	p.BC.SetChain = true
	p.QC.SetChain = true

	if err := k.setChain(p.BC); err != nil {
		return err
	}

	if err := k.setChain(p.QC); err != nil {
		return err
	}

	err := k.setPrice(p)
	if err != nil {
		return err
	}

	err = k.setOrderFeeRate(p)
	if err != nil {
		return err
	}

	err = k.setBCWithdrawalLimit(p)
	if err != nil {
		return err
	}

	err = k.setQCWithdrawalLimit(p)
	if err != nil {
		return err
	}

	return nil
}
