package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/Kucoin/kucoin-go-sdk"
)

type pairList struct {
	k     *kucoinExchange
	mux   *sync.Mutex
	api   *kucoin.ApiService
	pairs []*kucoin.SymbolModel

	l logger.Logger
}

func newPairList(k *kucoinExchange, api *kucoin.ApiService, l logger.Logger) *pairList {
	return &pairList{
		k:     k,
		mux:   &sync.Mutex{},
		api:   api,
		pairs: make([]*kucoin.SymbolModel, 0),
		l:     l,
	}
}

func (p *pairList) download() error {
	op := errors.Op(fmt.Sprintf("%s.pairList.download", p.k.Name()))

	res, err := p.api.Symbols("")
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	pairs := []*kucoin.SymbolModel{}
	if err := res.ReadData(&pairs); err != nil {
		return err
	}

	p.mux.Lock()
	defer p.mux.Unlock()
	p.pairs = pairs

	p.l.Debug(string(op), fmt.Sprintf("%d pairs downloaded", len(pairs)))
	return nil
}

func (pl *pairList) support(p *entity.Pair, fromDB bool) error {
	agent := fmt.Sprintf("%s.pairList.support", pl.k.Name())
	if len(pl.pairs) == 0 {
		if err := pl.download(); err != nil {
			return err
		}
		pl.l.Debug(agent, "pairs list downloaded")
	}

	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)

	pl.mux.Lock()
	defer pl.mux.Unlock()

	for _, pair := range pl.pairs {
		if !pair.EnableTrading {
			continue
		}
		symbol := strings.Split(pair.Symbol, "-")
		if len(symbol) != 2 {
			continue
		}
		bSymbol := symbol[0]
		qSymbol := symbol[1]
		if bSymbol == bc.Currency && qSymbol == qc.Currency {
			if fromDB {
				return nil
			}
			bc.MinOrderSize, _ = strconv.ParseFloat(pair.BaseMinSize, 64)
			bc.MaxOrderSize, _ = strconv.ParseFloat(pair.BaseMaxSize, 64)
			bc.OrderPrecision = calcPrecision(pair.BaseIncrement)
			qc.MinOrderSize, _ = strconv.ParseFloat(pair.QuoteMinSize, 64)
			qc.MaxOrderSize, _ = strconv.ParseFloat(pair.QuoteMaxSize, 64)
			qc.OrderPrecision = calcPrecision(pair.QuoteIncrement)

			return nil
		} else if bSymbol == qc.Currency && qSymbol == bc.Currency {
			tx := p.T1
			t1 := p.T2
			t2 := tx

			bc := t1.ET.(*Token)
			qc := t2.ET.(*Token)

			bc.MinOrderSize, _ = strconv.ParseFloat(pair.BaseMinSize, 64)
			bc.MaxOrderSize, _ = strconv.ParseFloat(pair.BaseMaxSize, 64)
			bc.OrderPrecision = calcPrecision(pair.BaseIncrement)
			qc.MinOrderSize, _ = strconv.ParseFloat(pair.QuoteMinSize, 64)
			qc.MaxOrderSize, _ = strconv.ParseFloat(pair.QuoteMaxSize, 64)
			qc.OrderPrecision = calcPrecision(pair.QuoteIncrement)

			p.T1 = t1
			p.T2 = t2
			return nil
		}
	}
	return errors.New(fmt.Sprintf("pair '%s' not supported", p.String()))
}

func calcPrecision(s string) int {
	ss := strings.Split(s, ".")
	if len(ss) == 1 {
		return 0
	}
	return len(ss[1])
}
