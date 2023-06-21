package kucoin

import (
	"exchange-provider/internal/entity"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *exchange) setOrderFeeRate(p *entity.Pair) error {
	ep := p.EP.(*ExchangePair)
	if ep.HasIntermediaryCoin {
		bc := p.T1.ET.(*Token)
		qc := ep.IC1
		f1, err := k.si.getFeeRate(bc.Currency, qc.Currency)
		if err != nil {
			return err
		}
		bc = p.T2.ET.(*Token)
		qc = ep.IC2
		f2, err := k.si.getFeeRate(bc.Currency, qc.Currency)
		if err != nil {
			return err
		}

		ep.KucoinFeeRate1 = f1
		ep.KucoinFeeRate2 = f2
		return nil
	}
	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)
	f0, err := k.si.getFeeRate(bc.Currency, qc.Currency)
	ep.KucoinFeeRate1 = f0
	return err
}

func (k *exchange) getAddress(t *Token) error {
	res, err := k.readApi.DepositAddresses(t.Currency, t.Chain)
	if err := handleSDKErr(err, res); err != nil {
		return err
	}
	da := &kucoin.DepositAddressModel{}
	if err := res.ReadData(da); err != nil {
		return handleSDKErr(err, res)
	}
	t.DepositAddress = da.Address
	t.DepositTag = da.Memo
	return nil
}

func (ex *exchange) setPairInfos(p *entity.Pair) error {
	var bc, qc *Token
	ep := p.EP.(*ExchangePair)
	if ep.HasIntermediaryCoin {
		bc = p.T1.ET.(*Token)
		qc = ep.IC1
		if err := ex.si.setPairInfos(bc, qc); err != nil {
			return err
		}
		bc = p.T2.ET.(*Token)
		qc = ep.IC2
		return ex.si.setPairInfos(bc, qc)
	}
	bc = p.T1.ET.(*Token)
	qc = p.T2.ET.(*Token)
	return ex.si.setPairInfos(bc, qc)
}
func (ex *exchange) checkStable(p *entity.Pair) error {
	bc := p.T1.ET.(*Token).Currency
	qc := p.T1.ET.(*Token).StableToken
	if bc != qc {
		if _, err := ex.si.getSymbol(bc, qc); err != nil {
			return err
		}
	}

	bc = p.T2.ET.(*Token).Currency
	qc = p.T2.ET.(*Token).StableToken
	if bc != qc {
		if _, err := ex.si.getSymbol(bc, qc); err != nil {
			return err
		}
	}
	return nil
}

func (k *exchange) setInfos(p *entity.Pair) error {
	if err := k.setPairInfos(p); err != nil {
		return err
	}

	if err := k.checkStable(p); err != nil {
		return err
	}

	if err := k.setTokenInfos(p.T1); err != nil {
		return err
	}

	if err := k.setTokenInfos(p.T2); err != nil {
		return err
	}

	if err := k.minAndMax(p); err != nil {
		return err
	}
	return nil
}
