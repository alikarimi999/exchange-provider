package binance

import (
	"context"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/adshao/go-binance/v2"
)

func (ex *exchange) isDipositAndWithdrawEnable(t *Token) (bool, bool, error) {
	n, err := ex.si.getCoin(t.Coin, t.Network)
	if err != nil {
		return false, false, errors.Wrap(errors.ErrInternal)
	}
	t.MinWithdrawalFee, _ = strconv.ParseFloat(n.WithdrawFee, 64)
	return n.DepositEnable, n.WithdrawEnable, nil
}

func (ex *exchange) getPairSymbols(p *entity.Pair) (binance.Symbol, binance.Symbol, error) {
	if p.EP.(*ExchangePair).HasIntermediaryCoin {
		bc := p.T1.ET.(*Token)
		qc := p.EP.(*ExchangePair).IC1
		s0, err := ex.si.getSymbol(bc.Coin, qc.Coin)
		if err != nil {
			return binance.Symbol{}, binance.Symbol{}, err
		}

		bc = p.T2.ET.(*Token)
		qc = p.EP.(*ExchangePair).IC2
		s1, err := ex.si.getSymbol(bc.Coin, qc.Coin)
		if err != nil {
			return binance.Symbol{}, binance.Symbol{}, err
		}

		return s0, s1, nil
	}
	bc := p.T1.ET.(*Token).Coin
	qc := p.T2.ET.(*Token).Coin
	s, err := ex.si.getSymbol(bc, qc)
	if err != nil {
		return binance.Symbol{}, binance.Symbol{}, err
	}
	return s, binance.Symbol{}, nil

}

func (ex *exchange) downloadSymbols() ([]binance.Symbol, error) {
	infos, err := ex.c.NewExchangeInfoService().Permissions("SPOT").Do(context.Background())
	if err != nil {
		return nil, err
	}

	ss := []binance.Symbol{}
	for _, s := range infos.Symbols {
		if s.Status == "TRADING" && s.IsSpotTradingAllowed {
			var hasMarket bool
			for _, ot := range s.OrderTypes {
				if ot == "MARKET" {
					hasMarket = true
					break
				}
			}
			if hasMarket {
				ss = append(ss, s)
			}
		}
	}
	return ss, nil
}

func (ex *exchange) downloadPrices() ([]*binance.SymbolPrice, error) {
	return ex.c.NewListPricesService().Do(context.Background())
}

func isPairEnable(p *entity.Pair, s0, s1 binance.Symbol) bool {
	if p.EP.(*ExchangePair).HasIntermediaryCoin {
		if s0.IsSpotTradingAllowed && s1.IsSpotTradingAllowed {
			return true
		} else {
			return false
		}
	}
	return s0.IsSpotTradingAllowed
}

func setPairsInfos(p *entity.Pair, s0, s1 binance.Symbol) error {
	var bc, qc *Token
	if p.EP.(*ExchangePair).HasIntermediaryCoin {
		bc = p.T1.ET.(*Token)
		qc = p.EP.(*ExchangePair).IC2

		pb, pq, err := getPrecision(bc.Coin, qc.Coin, s0)
		if err != nil {
			return err
		}
		bc.OrderPrecision = pb
		qc.OrderPrecision = pq

		bc = p.T2.ET.(*Token)
		qc = p.EP.(*ExchangePair).IC2
		pb, pq, err = getPrecision(bc.Coin, qc.Coin, s1)
		if err != nil {
			return err
		}

		bc.OrderPrecision = pb
		qc.OrderPrecision = pq
		return nil
	}

	bc = p.T1.ET.(*Token)
	qc = p.T2.ET.(*Token)
	pb, pq, err := getPrecision(bc.Coin, qc.Coin, s0)
	if err != nil {
		return err
	}

	bc.OrderPrecision = pb
	qc.OrderPrecision = pq
	return nil
}

func getPrecision(bc, qc string, s binance.Symbol) (int, int, error) {
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

func (ex *exchange) downloadCoins() ([]*binance.CoinInfo, error) {
	return ex.c.NewGetAllCoinsInfoService().Do(context.Background())
}

func (ex *exchange) infos(p *entity.Pair) error {

	s0, s1, err := ex.getPairSymbols(p)
	if err != nil {
		return err
	}

	if err := setPairsInfos(p, s0, s1); err != nil {
		return err
	}

	bEFA, _, err := ex.exchangeFeeAmount(p.T1, p)
	if err != nil {
		return err
	}

	qEFA, _, err := ex.exchangeFeeAmount(p.T2, p)
	if err != nil {
		return err
	}

	p0, p1, err := ex.price(p)
	if err != nil {
		return err
	}

	return ex.minAndMax(p, p0, p1, bEFA, qEFA, s0, s1)
}
