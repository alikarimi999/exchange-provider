package kucoin

import (
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
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

func (pl *pairList) support(p *pair) (bool, error) {
	agent := fmt.Sprintf("%s.pairList.support", pl.k.Name())
	if len(pl.pairs) == 0 {
		pl.l.Debug(agent, "pairs not downloaded yet")
		if err := pl.download(); err != nil {
			return false, err
		}
	}

	pl.mux.Lock()
	defer pl.mux.Unlock()

	for _, pair := range pl.pairs {
		if pair.BaseCurrency == p.BC.TokenId && pair.QuoteCurrency == p.QC.TokenId {
			p.BC.minOrderSize = pair.BaseMinSize
			p.BC.maxOrderSize = pair.BaseMaxSize
			p.BC.orderPrecision = calcPrecision(pair.BaseIncrement)
			p.QC.minOrderSize = pair.QuoteMinSize
			p.QC.maxOrderSize = pair.QuoteMaxSize
			p.QC.orderPrecision = calcPrecision(pair.QuoteIncrement)

			p.feeCurrency = pair.FeeCurrency

			return true, nil
		} else if pair.BaseCurrency == p.QC.TokenId && pair.QuoteCurrency == p.BC.TokenId {
			x := p.QC
			p.QC = p.BC
			p.BC = x

			p.BC.minOrderSize = pair.BaseMinSize
			p.BC.maxOrderSize = pair.BaseMaxSize
			p.BC.orderPrecision = calcPrecision(pair.BaseIncrement)
			p.QC.minOrderSize = pair.QuoteMinSize
			p.QC.maxOrderSize = pair.QuoteMaxSize
			p.QC.orderPrecision = calcPrecision(pair.QuoteIncrement)

			p.feeCurrency = pair.FeeCurrency

			return true, nil
		}
	}

	return false, nil
}

func calcPrecision(s string) int {
	ss := strings.Split(s, ".")
	if len(ss) == 1 {
		return 0
	}
	return len(ss[1])
}
