package kucoin

import (
	"exchange-provider/internal/entity"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Kucoin/kucoin-go-sdk"
)

type serverInfos struct {
	sMux    *sync.RWMutex
	symbols map[string]*kucoin.SymbolModel

	tMux   *sync.RWMutex
	tokens map[string]*kucoinToken

	pMux   *sync.RWMutex
	prices map[string]float64

	fMux     *sync.RWMutex
	feeRates map[string]float64

	t *time.Ticker
}

func newServerInfos(ex *exchange, api *kucoin.ApiService) (*serverInfos, error) {
	ss, err := ex.downloadList()
	if err != nil {
		return nil, err
	}
	symbols := make(map[string]*kucoin.SymbolModel)
	for _, s := range ss {
		symbols[s.Symbol] = s
	}

	ts, err := ex.downloadTokens()
	if err != nil {
		return nil, err
	}
	tokens := make(map[string]*kucoinToken)
	for _, t := range ts.Data {
		tokens[t.Currency+t.ChainName] = t
	}

	ps, err := ex.downloadPrices()
	if err != nil {
		return nil, err
	}

	prices := make(map[string]float64)
	feeRates := make(map[string]float64)
	for _, p := range ps {
		prices[p.Symbol], _ = strconv.ParseFloat(p.Last, 64)
		tf, _ := strconv.ParseFloat(p.TakerFeeRate, 64)
		tfc, _ := strconv.ParseFloat(p.TakerCoefficient, 64)
		feeRates[p.Symbol] = tf * tfc
	}

	return &serverInfos{
		sMux:    &sync.RWMutex{},
		symbols: symbols,

		tMux:   &sync.RWMutex{},
		tokens: tokens,

		pMux:   &sync.RWMutex{},
		prices: prices,

		fMux:     &sync.RWMutex{},
		feeRates: feeRates,

		t: time.NewTicker(25 * time.Second),
	}, nil

}

func (si *serverInfos) run(ex *exchange, stopCh chan struct{}) {
	agent := ex.agent("serverInfos.run")
	for {
		select {
		case <-si.t.C:
			ss, err := ex.downloadList()
			if err == nil {
				symbols := make(map[string]*kucoin.SymbolModel)
				for _, s := range ss {
					symbols[s.Symbol] = s
				}
				si.sMux.Lock()
				si.symbols = symbols
				si.sMux.Unlock()

			} else {
				ex.l.Error(agent, err.Error())
			}

			ts, err := ex.downloadTokens()
			if err == nil {
				tokens := make(map[string]*kucoinToken)
				for _, t := range ts.Data {
					tokens[t.Currency+t.ChainName] = t
				}
				si.tMux.Lock()
				si.tokens = tokens
				si.tMux.Unlock()
			} else {
				ex.l.Error(agent, err.Error())
			}

			ps, err := ex.downloadPrices()
			if err == nil {
				prices := make(map[string]float64)
				feeRates := make(map[string]float64)
				for _, p := range ps {
					prices[p.Symbol], _ = strconv.ParseFloat(p.Last, 64)
					tf, _ := strconv.ParseFloat(p.TakerFeeRate, 64)
					tfc, _ := strconv.ParseFloat(p.TakerCoefficient, 64)
					feeRates[p.Symbol] = tf * tfc
				}
				si.pMux.Lock()
				si.prices = prices
				si.pMux.Unlock()

				si.fMux.Lock()
				si.feeRates = feeRates
				si.fMux.Unlock()

			} else {
				ex.l.Error(agent, err.Error())
			}
		case <-stopCh:
			return
		}
	}
}

func (ex *exchange) downloadList() ([]*kucoin.SymbolModel, error) {
	res, err := ex.readApi.Symbols("")
	if err := handleSDKErr(err, res); err != nil {
		return nil, err
	}

	pairs := []*kucoin.SymbolModel{}
	if err := res.ReadData(&pairs); err != nil {
		return nil, err
	}

	return pairs, nil
}

func (si *serverInfos) getSymbol(bc, qc string) (*kucoin.SymbolModel, error) {
	si.sMux.RLock()
	defer si.sMux.RUnlock()
	s, ok := si.symbols[bc+"-"+qc]
	if !ok {
		return nil, fmt.Errorf("symbol %s/%s not found", bc, qc)
	}
	return s, nil
}

func (si *serverInfos) getToken(currency string, chainName string) (*kucoinToken, error) {
	si.tMux.RLock()
	defer si.tMux.RUnlock()
	t, ok := si.tokens[currency+chainName]
	if !ok {
		return nil, fmt.Errorf("token %s-%s not found", currency, chainName)
	}
	return t, nil
}

func (si *serverInfos) getPrice(bc, qc string) (float64, error) {
	si.pMux.RLock()
	defer si.pMux.RUnlock()
	p, ok := si.prices[bc+"-"+qc]
	if !ok {
		return 0, fmt.Errorf("symbol %s/%s not found", bc, qc)
	}
	return p, nil
}

func (si *serverInfos) getFeeRate(bc, qc string) (float64, error) {
	si.fMux.RLock()
	defer si.fMux.RUnlock()
	f, ok := si.feeRates[bc+"-"+qc]
	if !ok {
		return 0, fmt.Errorf("symbol %s/%s not found", bc, qc)
	}
	return f, nil
}

func (pl *serverInfos) setPairInfos(bc, qc *Token) error {
	pl.sMux.RLock()
	defer pl.sMux.RUnlock()
	s, ok := pl.symbols[bc.Currency+"-"+qc.Currency]
	if ok {
		if !s.EnableTrading {
			return fmt.Errorf("symbol %s/%s is not enable in kucoin", bc.Currency, qc.Currency)
		}
		bc.MinOrderSize, _ = strconv.ParseFloat(s.BaseMinSize, 64)
		bc.MaxOrderSize, _ = strconv.ParseFloat(s.BaseMaxSize, 64)
		bc.OrderPrecision = calcPrecision(s.BaseIncrement)
		qc.MinOrderSize, _ = strconv.ParseFloat(s.QuoteMinSize, 64)
		qc.MaxOrderSize, _ = strconv.ParseFloat(s.QuoteMaxSize, 64)
		qc.OrderPrecision = calcPrecision(s.QuoteIncrement)
		return nil
	}

	return fmt.Errorf("symbol %s/%s not found",
		bc.Currency, qc.Currency)
}

func (ex *exchange) downloadPrices() ([]*kucoin.TickerModel, error) {
	res, err := ex.readApi.Tickers()
	if err != nil {
		return nil, err
	}

	ts := kucoin.TickersResponseModel{}
	err = res.ReadData(&ts)
	return ts.Tickers, err
}

func (ex *exchange) isPairEnabled(p *entity.Pair) bool {
	ep := p.EP.(*ExchangePair)
	if ep.HasIntermediaryCoin {
		bc := p.T1.ET.(*Token).Currency
		qc := ep.IC1.Currency
		s, err := ex.si.getSymbol(bc, qc)
		if err != nil {
			return false
		}
		if !s.EnableTrading {
			return false
		}

		bc = p.T2.ET.(*Token).Currency
		qc = ep.IC2.Currency
		s, err = ex.si.getSymbol(bc, qc)
		if err != nil {
			return false
		}
		return s.EnableTrading
	}
	s, err := ex.si.getSymbol(p.T1.ET.(*Token).Currency, p.T2.ET.(*Token).Currency)
	if err != nil {
		return false
	}
	return s.EnableTrading
}

func calcPrecision(s string) int {
	ss := strings.Split(s, ".")
	if len(ss) == 1 {
		return 0
	}
	return len(ss[1])
}
