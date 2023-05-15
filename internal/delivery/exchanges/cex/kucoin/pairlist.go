package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"strconv"
	"strings"

	"github.com/Kucoin/kucoin-go-sdk"
)

type pairList struct {
	k       *kucoinExchange
	symbols []*kucoin.SymbolModel
	l       logger.Logger
}

func newPairList(k *kucoinExchange, api *kucoin.ApiService, l logger.Logger) *pairList {
	pl := &pairList{
		k:       k,
		symbols: make([]*kucoin.SymbolModel, 0),
		l:       l,
	}
	return pl
}

func (p *pairList) downloadList() error {
	agent := p.k.agent("pairList.downloadList")
	res, err := p.k.readApi.Symbols("")
	if err := handleSDKErr(err, res); err != nil {
		return err
	}

	pairs := []*kucoin.SymbolModel{}
	if err := res.ReadData(&pairs); err != nil {
		return err
	}
	p.symbols = pairs
	p.l.Debug(agent, fmt.Sprintf("'%d' pairs downloaded", len(pairs)))
	return nil
}

func (k *kucoinExchange) support(p *entity.Pair) error {
	var bc, qc *Token
	ep := p.EP.(*ExchangePair)
	if ep.HasIntermediaryCoin {
		bc = p.T1.ET.(*Token)
		qc = ep.IC1
		if err := k.pls.support(bc, qc); err != nil {
			return err
		}
		bc = p.T2.ET.(*Token)
		qc = ep.IC2
		return k.pls.support(bc, qc)
	}
	bc = p.T1.ET.(*Token)
	qc = p.T2.ET.(*Token)
	return k.pls.support(bc, qc)
}

func (pl *pairList) support(bc, qc *Token) error {
	for _, pair := range pl.symbols {
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
			bc.MinOrderSize, _ = strconv.ParseFloat(pair.BaseMinSize, 64)
			bc.MaxOrderSize, _ = strconv.ParseFloat(pair.BaseMaxSize, 64)
			bc.OrderPrecision = calcPrecision(pair.BaseIncrement)
			qc.MinOrderSize, _ = strconv.ParseFloat(pair.QuoteMinSize, 64)
			qc.MaxOrderSize, _ = strconv.ParseFloat(pair.QuoteMaxSize, 64)
			qc.OrderPrecision = calcPrecision(pair.QuoteIncrement)
			return nil
		}
	}
	return errors.Wrap(errors.ErrNotFound, fmt.Errorf("kucoin does not support pair '%s/%s'",
		bc.Currency, qc.Currency))
}

func calcPrecision(s string) int {
	ss := strings.Split(s, ".")
	if len(ss) == 1 {
		return 0
	}
	return len(ss[1])
}
