package swapspace

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/try"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type exStatus struct {
	From struct {
		Amount float64 `json:"amount"`
	} `json:"from"`
	To struct {
		Amount float64 `json:"amount"`
	} `json:"to"`
	Status string `json:"status"`
	ID     string `json:"id"`
	Error  bool   `json:"error"`
}

func (ex *exchange) TxIdSetted(o *entity.CexOrder) {
	ex.trackExchange(o)
	ex.repo.Update(o)
}

func (ex *exchange) trackExchange(o *entity.CexOrder) {
	ds := o.Swaps[0].Duration
	d := strings.Split(ds, "-")
	di, _ := strconv.Atoi(d[0])
	if di == 0 {
		di = 5
	}
	oid := o.MetaData["id_in_swapspace"].(string)

	err := try.Do(15, func(attempt uint64) (retry bool, err error) {
		res, err := ex.exchangeStatus(oid)
		if err != nil {
			return true, err
		}

		if res.Error {
			return true, fmt.Errorf("")
		}
		o.Withdrawal.Volume = fmt.Sprintf("%v", res.To.Amount)
		switch res.Status {
		case "waiting", "confirming", "exchanging", "sending", "verifying":
			time.Sleep(time.Second)
			return true, fmt.Errorf("")
		case "finished":
			o.Status = entity.OSucceeded
		case "failed":
			o.Status = entity.OFailed
		case "refunded":
			o.Status = entity.ORefunded
		case "expired":
			o.Status = entity.OExpired
		}
		return false, nil
	})
	if err != nil {
		o.Status = entity.OFailed
		o.FailedCode = entity.FCExOrdFailed
		o.FailedDesc = err.Error()
	}
}

func (ex *exchange) exchangeStatus(id string) (*exStatus, error) {
	agent := ex.agent("exchangeStatus")

	urlStr, _ := url.JoinPath(baseUrl, "exchange", id)
	b, err := ex.request(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}
	res := &exStatus{}
	if err := json.Unmarshal(b, res); err != nil {
		ex.l.Error(agent, err.Error())
		return nil, err
	}
	return res, nil
}
