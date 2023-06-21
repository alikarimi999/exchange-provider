package allbridge

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

type chache struct {
	lastId uint64

	mux0  *sync.RWMutex
	list0 map[string]time.Time

	mux1  *sync.RWMutex
	list1 map[string]item

	t *time.Ticker
}

func newCache() *chache {
	return &chache{
		mux0:  &sync.RWMutex{},
		list0: make(map[string]time.Time),
		mux1:  &sync.RWMutex{},
		list1: make(map[string]item),
		t:     time.NewTicker(5 * time.Second),
	}
}

func (c *chache) run(ex *allBridge, stopCh <-chan struct{}) {
	agent := ex.agent("chache.run")
	for {
		select {
		case <-c.t.C:
			url, _ := url.Parse(ex.cfg.Url)
			v := url.Query()
			v.Set("page", "1")
			v.Set("limit", "20")
			url.RawQuery = v.Encode()

			if c.lastId == 0 {
				func() {
					req, _ := http.NewRequest(http.MethodGet, url.String(), nil)
					res, err := http.DefaultClient.Do(req)
					if err != nil {
						ex.l.Debug(agent, err.Error())
						return
					}
					defer res.Body.Close()

				}()
			}

		case <-stopCh:
			return
		}
	}
}

type transfers struct {
	Items []item `json:"items"`
	Meta  Meta   `json:"meta"`
}
type item struct {
	ID               string `json:"id"`
	Status           string `json:"status"`
	Timestamp        int64  `json:"timestamp"`
	FromChainSymbol  string `json:"fromChainSymbol"`
	ToChainSymbol    string `json:"toChainSymbol"`
	FromAmount       string `json:"fromAmount"`
	ToAmount         string `json:"toAmount"`
	FromTokenAddress string `json:"fromTokenAddress"`
	ToTokenAddress   string `json:"toTokenAddress"`
	FromAddress      string `json:"fromAddress"`
	ToAddress        string `json:"toAddress"`
	MessagingType    string `json:"messagingType"`
	downloadedAt     time.Time
}

type transferD struct {
	item
	Transactions []transactions `json:"transactions"`
}
type transactions struct {
	BlockTime   int64  `json:"blockTime"`
	ChainSymbol string `json:"chainSymbol"`
	TxID        string `json:"txId"`
	Type        string `json:"type"`
}

type Meta struct {
	ItemCount    int `json:"itemCount"`
	TotalItems   int `json:"totalItems"`
	ItemsPerPage int `json:"itemsPerPage"`
	TotalPages   int `json:"totalPages"`
	CurrentPage  int `json:"currentPage"`
}
