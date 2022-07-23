package app

import (
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"
)

type tickerList struct {
	mu   sync.Mutex
	list map[string]*chainTicker
}

type withdrawalHandler struct {
	tList    *tickerList
	tickerCh chan *chainTicker
	tracker  *withdrawalTracker
	wg       *sync.WaitGroup
	l        logger.Logger
}

func newWithdrawalHandler(repo entity.OrderRepo, oc entity.OrderCache,
	wc entity.WithdrawalCache, exs map[string]entity.Exchange, l logger.Logger) *withdrawalHandler {

	w := &withdrawalHandler{
		tList: &tickerList{
			mu:   sync.Mutex{},
			list: map[string]*chainTicker{},
		},
		tickerCh: make(chan *chainTicker),
		tracker:  newWithdrawalTracker(repo, oc, wc, exs, l),
		wg:       &sync.WaitGroup{},

		l: l,
	}

	return w
}

func (h *withdrawalHandler) run(wg *sync.WaitGroup) {
	const agent = "Withdrawal-Handler.run"

	defer wg.Done()

	h.wg.Add(1)
	go h.tracker.run(h.wg)

	for ti := range h.tickerCh {
		h.wg.Add(1)
		go ti.tick(h.wg)
	}

	h.wg.Wait()
}

func (h *withdrawalHandler) addChainTickers(chains []*entity.Chain) {
	// check if chain exists
	h.tList.mu.Lock()
	defer h.tList.mu.Unlock()

	for _, chain := range chains {
		if _, ok := h.tList.list[chain.Id]; !ok {
			ti := h.newChainTicker(chain)
			h.tList.list[chain.Id] = ti
			h.tickerCh <- ti
		}

	}

}

func (h *withdrawalHandler) removeTicker(chainId string) {
	ti := h.tList.list[chainId]
	ti.stop()
	delete(h.tList.list, chainId)
}

func (h *withdrawalHandler) isTickerRunning(chainId string) bool {
	_, ok := h.tList.list[chainId]
	return ok
}
