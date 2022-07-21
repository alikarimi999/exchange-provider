package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"
	"time"
)

type withdrawalHandler struct {
	tickers []*chainTicker
	tracker *withdrawalTracker
	wg      *sync.WaitGroup
	l       logger.Logger
}

func newWithdrawalHandler(repo entity.OrderRepo, oc entity.OrderCache,
	wc entity.WithdrawalCache, exs map[string]entity.Exchange, l logger.Logger) *withdrawalHandler {

	w := &withdrawalHandler{
		wg:      &sync.WaitGroup{},
		tracker: newWithdrawalTracker(repo, oc, wc, exs, l),
		l:       l,
	}

	btc := &chainTicker{
		chain:       entity.ChainBTC,
		cache:       wc,
		ticker:      time.NewTicker(time.Minute * 10),
		tracker:     w.tracker,
		windowsSize: time.Minute * 10,
		l:           l,
	}
	tron := &chainTicker{
		chain:       entity.ChainTRC20,
		cache:       wc,
		ticker:      time.NewTicker(time.Second * 60),
		tracker:     w.tracker,
		windowsSize: time.Minute * 10,
		l:           l,
	}

	w.tickers = []*chainTicker{btc, tron}

	return w
}

func (h *withdrawalHandler) run(wg *sync.WaitGroup) {
	const agent = "Withdrawal-Handler.run"

	defer wg.Done()

	h.wg.Add(1)
	go h.tracker.run(h.wg)

	for _, ti := range h.tickers {
		h.wg.Add(1)
		go ti.tick(h.wg)
		h.l.Debug(agent, fmt.Sprintf("Started ticker for chain: '%s'", ti.chain))
	}

	h.wg.Wait()
}
