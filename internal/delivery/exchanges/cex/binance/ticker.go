package binance

import (
	"context"
	"fmt"
	"strconv"
)

func (ex *exchange) ticker(bc, qc string) (float64, error) {
	tickers, err := ex.c.NewListSymbolTickerService().Symbol(bc + qc).Do(context.Background())
	if err != nil {
		return 0, err
	}

	if len(tickers) == 0 {
		return 0, fmt.Errorf("ticker not fount for pair %s/%s", bc, qc)
	}

	price, _ := strconv.ParseFloat(tickers[0].LastPrice, 64)
	return price, nil
}
