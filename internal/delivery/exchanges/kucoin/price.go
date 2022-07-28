package kucoin

import (
	"fmt"
	"order_service/internal/entity"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) setPrice(p *entity.Pair) error {
	res, err := k.api.TickerLevel1(fmt.Sprintf("%s%s%s", p.BC.Coin.CoinId, pairDelimiter, p.QC.Coin.CoinId))
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	t := &kucoin.TickerLevel1Model{}
	err = res.ReadData(t)
	if err != nil {
		return err
	}

	p.Price = t.BestAsk
	return nil
}

func (k *kucoinExchange) setOrderFeeRate(p *entity.Pair) error {
	res, err := k.api.ActualFee(fmt.Sprintf("%s%s%s", p.BC.Coin.CoinId, pairDelimiter, p.QC.Coin.CoinId))
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	m := kucoin.TradeFeesResultModel{}
	err = res.ReadData(&m)
	if err != nil {
		return err
	}

	p.OrderFeeRate = m[0].TakerFeeRate

	return nil

}

func (k *kucoinExchange) setWithdrawalLimit(p *entity.Pair) error {
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
			break
		}
	}

	res, err = k.api.CurrencyV2(p.QC.Coin.CoinId, "")
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	m = kucoin.CurrencyV2Model{}
	err = res.ReadData(&m)
	if err != nil {
		return err
	}

	for _, c := range m.Chains {
		if c.ChainName == p.QC.ChainId {
			p.QC.MinWithdrawalSize = c.WithdrawalMinSize
			p.QC.WithdrawalMinFee = c.WithdrawalMinFee
			break
		}
	}

	p.BC.WithdrawalPrecision = int(m.Precision)
	p.QC.WithdrawalPrecision = int(m.Precision)

	return nil
}

func (k *kucoinExchange) setInfos(p *entity.Pair) error {
	err := k.setPrice(p)
	if err != nil {
		return err
	}

	err = k.setOrderFeeRate(p)
	if err != nil {
		return err
	}

	err = k.setWithdrawalLimit(p)
	if err != nil {
		return err
	}
	return nil
}
