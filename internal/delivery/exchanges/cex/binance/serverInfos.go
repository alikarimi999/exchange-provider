package binance

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
)

type serverInfos struct {
	pMux   *sync.RWMutex
	prices map[string]float64

	cMux  *sync.RWMutex
	coins map[string]binance.Network

	sMux    *sync.RWMutex
	symbols map[string]binance.Symbol

	tMux      *sync.RWMutex
	tradeFees map[string]float64

	t *time.Ticker
}

func newServerInfos(ex *exchange) (*serverInfos, error) {
	ss, err := ex.downloadSymbols()
	if err != nil {
		return nil, err
	}
	symbols := make(map[string]binance.Symbol)
	for _, s := range ss {
		symbols[s.Symbol] = s
	}

	cs, err := ex.downloadCoins()
	if err != nil {
		return nil, err
	}
	coins := make(map[string]binance.Network)
	for _, c := range cs {
		for _, n := range c.NetworkList {
			coins[c.Coin+n.Network] = n
		}
	}

	ps, err := ex.downloadPrices()
	if err != nil {
		return nil, err
	}
	prices := make(map[string]float64)
	for _, p := range ps {
		prices[p.Symbol], _ = strconv.ParseFloat(p.Price, 64)
	}

	ts, err := ex.c.NewTradeFeeService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	tradeFees := make(map[string]float64)
	for _, t := range ts {
		tradeFees[t.Symbol], _ = strconv.ParseFloat(t.TakerCommission, 64)
	}

	return &serverInfos{
		pMux:   &sync.RWMutex{},
		prices: prices,

		cMux:  &sync.RWMutex{},
		coins: coins,

		sMux:    &sync.RWMutex{},
		symbols: symbols,

		tMux:      &sync.RWMutex{},
		tradeFees: tradeFees,

		t: time.NewTicker(25 * time.Second),
	}, nil
}

func (si *serverInfos) run(ex *exchange, stopCh chan struct{}) {
	agent := ex.agent("serverInfos.run")
	for {
		select {
		case <-si.t.C:

			ps, err := ex.downloadPrices()
			if err == nil {
				prices := make(map[string]float64)
				for _, p := range ps {
					prices[p.Symbol], _ = strconv.ParseFloat(p.Price, 64)
				}
				si.pMux.Lock()
				si.prices = prices
				si.pMux.Unlock()
			} else {
				ex.l.Error(agent, err.Error())
			}

			cs, err := ex.downloadCoins()
			if err == nil {
				coins := make(map[string]binance.Network)
				for _, c := range cs {
					for _, n := range c.NetworkList {
						coins[c.Coin+n.Network] = n
					}
				}
				si.cMux.Lock()
				si.coins = coins
				si.cMux.Unlock()
			} else {
				ex.l.Error(agent, err.Error())
			}

			ss, err := ex.downloadSymbols()
			if err == nil {
				symbols := make(map[string]binance.Symbol)
				for _, s := range ss {
					symbols[s.Symbol] = s
				}
				si.sMux.Lock()
				si.symbols = symbols
				si.sMux.Unlock()
			} else {
				ex.l.Error(agent, err.Error())
			}

			ts, err := ex.c.NewTradeFeeService().Do(context.Background())
			if err == nil {
				tradeFees := make(map[string]float64)
				for _, t := range ts {
					tradeFees[t.Symbol], _ = strconv.ParseFloat(t.TakerCommission, 64)
				}
				si.tMux.Lock()
				si.tradeFees = tradeFees
				si.tMux.Unlock()
			} else {
				ex.l.Error(agent, err.Error())
			}

		case <-stopCh:
			fmt.Println("stop")
			return
		}
	}
}

func (si *serverInfos) getPrice(bc, qc string) (float64, error) {
	si.pMux.RLock()
	defer si.pMux.RUnlock()
	p, ok := si.prices[bc+qc]
	if !ok {
		return 0, fmt.Errorf("getPrice: symbol %s/%s not found", bc, qc)
	}
	return p, nil
}

func (si *serverInfos) getSymbol(bc, qc string) (binance.Symbol, error) {
	si.sMux.RLock()
	defer si.sMux.RUnlock()
	smb, ok := si.symbols[bc+qc]
	if !ok {
		return binance.Symbol{}, fmt.Errorf("getSymbol: symbol %s/%s not found", bc, qc)
	}
	return smb, nil
}

func (si *serverInfos) getCoin(coin, network string) (binance.Network, error) {
	si.cMux.RLock()
	defer si.cMux.RUnlock()
	c, ok := si.coins[coin+network]
	if !ok {
		return binance.Network{}, fmt.Errorf("getCoin: coin %s-%s not found", coin, network)
	}
	return c, nil
}

func (si *serverInfos) getTradeFee(bc, qc string) (float64, error) {
	si.tMux.RLock()
	defer si.tMux.RUnlock()
	t, ok := si.tradeFees[bc+qc]
	if !ok {
		return 0, fmt.Errorf("getTradeFee: symbol %s/%s not found", bc, qc)
	}
	return t, nil
}
