package binance

import (
	"context"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/adshao/go-binance/v2"
)

func (ex *exchange) isDipositAndWithdrawEnable(t *Token) (bool, bool, error) {
	n, ok := ex.cl.getCoin(t.Coin, t.Network)
	if !ok {
		return false, false, errors.Wrap(errors.ErrInternal)
	}
	t.MinWithdrawalFee, _ = strconv.ParseFloat(n.WithdrawFee, 64)
	return n.DepositEnable, n.WithdrawEnable, nil
}

func (ex *exchange) getSymbol(bc, qc string) (*binance.Symbol, error) {
	inf, err := ex.c.NewExchangeInfoService().Symbol(bc + qc).Do(context.Background())
	if err != nil {
		return nil, err
	}
	if len(inf.Symbols) == 0 {
		return nil, fmt.Errorf("pair %s/%s not found", bc, qc)
	}
	return &inf.Symbols[0], nil
}

func (ex *exchange) getPairSymbols(p *entity.Pair) (*binance.Symbol, *binance.Symbol, error) {
	if p.EP.(*ExchangePair).HasIntermediaryCoin {
		bc := p.T1.ET.(*Token)
		qc := p.EP.(*ExchangePair).IC1
		s0, err := ex.getSymbol(bc.Coin, qc.Coin)
		if err != nil {
			return nil, nil, err
		}

		bc = p.T2.ET.(*Token)
		qc = p.EP.(*ExchangePair).IC2
		s1, err := ex.getSymbol(bc.Coin, qc.Coin)
		if err != nil {
			return nil, nil, err
		}
		return s0, s1, nil
	}
	s, err := ex.getSymbol(p.T1.ET.(*Token).Coin, p.T2.ET.(*Token).Coin)
	return s, nil, err
}

func (ex *exchange) setPairsInfos(p *entity.Pair) (*binance.Symbol, *binance.Symbol, error) {
	var bc, qc *Token
	if p.EP.(*ExchangePair).HasIntermediaryCoin {
		bc = p.T1.ET.(*Token)
		qc = p.EP.(*ExchangePair).IC1
		s0, err := ex.getSymbol(bc.Coin, qc.Coin)
		if err != nil {
			return nil, nil, err
		}

		pb, pq, err := getPrecision(bc.Coin, qc.Coin, s0)
		if err != nil {
			return nil, nil, err
		}
		bc.OrderPrecision = pb
		qc.OrderPrecision = pq

		bc = p.T2.ET.(*Token)
		qc = p.EP.(*ExchangePair).IC2
		s1, err := ex.getSymbol(bc.Coin, qc.Coin)
		if err != nil {
			return nil, nil, err
		}

		pb, pq, err = getPrecision(bc.Coin, qc.Coin, s1)
		if err != nil {
			return nil, nil, err
		}

		bc.OrderPrecision = pb
		qc.OrderPrecision = pq
		return s0, s1, nil
	}
	bc = p.T1.ET.(*Token)
	qc = p.T2.ET.(*Token)
	inf, err := ex.c.NewExchangeInfoService().Symbol(bc.Coin + qc.Coin).Do(context.Background())
	if err != nil {
		return nil, nil, err
	}
	if len(inf.Symbols) == 0 {
		return nil, nil, fmt.Errorf("pair %s/%s not found", bc.Coin, qc.Coin)
	}
	pb, pq, err := getPrecision(bc.Coin, qc.Coin, &inf.Symbols[0])
	if err != nil {
		return nil, nil, err
	}

	bc.OrderPrecision = pb
	qc.OrderPrecision = pq
	return &inf.Symbols[0], nil, nil
}

func getPrecision(bc, qc string, s *binance.Symbol) (int, int, error) {
	for _, fs := range s.Filters {
		if fs["filterType"] == "LOT_SIZE" {
			ss := strings.Split(fs["stepSize"].(string), ".")
			if len(ss) != 2 {
				return 0, 0, fmt.Errorf("invalid step size for pair %s/%s", bc, qc)
			}
			if strings.Contains(ss[0], "1") {
				i, _ := strconv.Atoi(ss[0])
				if i > 1 {
					return 0, 0, fmt.Errorf("invalid step size for pair %s/%s", bc, qc)
				}
				return 0, s.QuoteAssetPrecision, nil
			} else if strings.Contains(ss[1], "1") {
				return strings.IndexAny(ss[1], "1") + 1, s.QuoteAssetPrecision, nil
			} else {
				return s.BaseAssetPrecision, s.QuoteAssetPrecision, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("step size not found for pair %s/%s", bc, qc)
}

func (ex *exchange) getAddress(t *Token) error {
	da, err := ex.c.NewGetDepositAddressService().Coin(t.Coin).
		Network(t.Network).Do(context.Background())
	if err != nil {
		return err
	}
	t.DepositAddress = da.Address
	t.DepositTag = da.Tag
	return nil
}

func (ex *exchange) infos(p *entity.Pair) error {
	var err1, err2, err3, err4 error
	var bEFA, qEFA float64

	var s0, s1 *binance.Symbol
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		s0, s1, err1 = ex.setPairsInfos(p)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err2 = ex.setOrderFeeRate(p)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bEFA, err3 = ex.exchangeFeeAmount(p.T1, p)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		qEFA, err4 = ex.exchangeFeeAmount(p.T2, p)
	}()

	wg.Wait()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	if err3 != nil {
		return err3
	}
	if err4 != nil {
		return err4
	}

	p0, p1, err := ex.price(p)
	if err != nil {
		return err
	}
	return ex.minAndMax(p, p0, p1, bEFA, qEFA, s0, s1)
}
