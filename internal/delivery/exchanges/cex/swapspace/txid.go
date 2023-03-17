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
	Status string `json:"status"`
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
		s, err := ex.exchangeStatus(oid)
		if err != nil {
			return true, err
		}

		switch s {
		case "waiting", "confirming", "exchanging", "sending", "verifying":
			time.Sleep(time.Duration(di) * time.Minute)
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
		o.FailedCode = entity.FCExOrdFailed
		o.FailedDesc = err.Error()
	}
}

func (ex *exchange) exchangeStatus(id string) (string, error) {
	agent := ex.agent("exchangeStatus")

	urlStr, _ := url.JoinPath(baseUrl, "exchange", id)
	b, err := ex.request(http.MethodGet, urlStr, nil)
	if err != nil {
		return "", err
	}
	res := &exStatus{}
	if err := json.Unmarshal(b, res); err != nil {
		ex.l.Error(agent, err.Error())
		return "", err
	}
	return res.Status, nil
}
