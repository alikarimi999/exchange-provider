package kucoin

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"order_service/pkg/logger"
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
	op := errors.Op(fmt.Sprintf("%s.pairList.download", p.k.NID()))

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

func (pl *pairList) support(p *entity.Pair) (bool, error) {
	agent := fmt.Sprintf("%s.pairList.support", pl.k.NID())
	if len(pl.pairs) == 0 {
		pl.l.Debug(agent, "pairs not downloaded yet")
		if err := pl.download(); err != nil {
			return false, err
		}
	}

	pl.mux.Lock()
	defer pl.mux.Unlock()

	for _, pair := range pl.pairs {
		if pair.BaseCurrency == p.BC.CoinId && pair.QuoteCurrency == p.QC.CoinId {
			p.BC.MinOrderSize = pair.BaseMinSize
			p.BC.MaxOrderSize = pair.BaseMaxSize
			p.BC.OrderPrecision = calcPrecision(pair.BaseIncrement)
			p.QC.MinOrderSize = pair.QuoteMinSize
			p.QC.MaxOrderSize = pair.QuoteMaxSize
			p.QC.OrderPrecision = calcPrecision(pair.QuoteIncrement)

			p.FeeCurrency = pair.FeeCurrency

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
