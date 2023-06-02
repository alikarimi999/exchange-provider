package binance

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/cex/binance/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/adshao/go-binance/v2"
)

func getBcQcWcFeeRate(o *types.Order, p *entity.Pair,
	i int) (bc, qc, wc *Token) {
	if i == 0 {
		if len(o.Swaps) == 2 {
			if o.Swaps[0].In.String() == p.T1.Id.String() {
				bc = p.T1.ET.(*Token)
				qc = p.EP.(*ExchangePair).IC1
			} else {
				bc = p.T2.ET.(*Token)
				qc = p.EP.(*ExchangePair).IC2
			}
		} else {
			bc = p.T1.ET.(*Token)
			qc = p.T2.ET.(*Token)
			if o.Swaps[0].Side == binance.SideTypeSell {
				wc = qc
			} else {
				wc = bc
			}
		}
		o.Swaps[0].InAmountRequested = o.Deposit.Amount
	} else {

		o.Swaps[1].InAmountRequested = o.Swaps[0].OutAmount
		if o.Swaps[1].Out.String() == p.T2.String() {
			bc = p.T2.ET.(*Token)
			qc = p.EP.(*ExchangePair).IC2
		} else {
			bc = p.T1.ET.(*Token)
			qc = p.EP.(*ExchangePair).IC1
		}
		wc = bc
	}
	return
}

func (ex *exchange) setOrderFeeRate(p *entity.Pair) error {
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
			f1, err = ex.orderFeeRate(bc.Coin, qc.Coin)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			bc := p.T2.ET.(*Token)
			qc := ep.IC2
			f2, err = ex.orderFeeRate(bc.Coin, qc.Coin)
		}()
		wg.Wait()

		if err != nil {
			return err
		}

		ep.BinanceFeeRate1 = f1
		ep.BinanceFeeRate2 = f2
		return nil
	}

	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)
	f0, err := ex.orderFeeRate(bc.Coin, qc.Coin)
	ep.BinanceFeeRate1 = f0
	return err
}

func (ex *exchange) orderFeeRate(bc, qc string) (float64, error) {
	res, err := ex.c.NewTradeFeeService().Symbol(bc + qc).Do(context.Background())
	if err != nil {
		return 0, err
	}

	if len(res) == 0 {
		return 0, errors.Wrap(errors.ErrNotFound, fmt.Errorf("order fee rate not found for pair %s/%s", bc, qc))
	}

	return strconv.ParseFloat(res[0].TakerCommission, 64)
}

func (ex *exchange) exchangeFeeAmount(out *entity.Token, p *entity.Pair) (float64, error) {
	var (
		qcDollar float64
	)

	if out.ET.(*Token).Coin == out.ET.(*Token).StableToken {
		qcDollar = 1
	} else {
		qd, err := ex.ticker(out.ET.(*Token).Coin, out.ET.(*Token).StableToken)
		if err != nil {
			return 0, err
		}
		qcDollar = qd
	}

	return p.ExchangeFee / qcDollar, nil
}
