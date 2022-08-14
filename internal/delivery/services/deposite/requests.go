package deposite

import (
	"io"
	"net/http"
	"order_service/pkg/errors"
	"order_service/pkg/utils"
)

func (d *depositeService) setTxIdRequest(body io.Reader) (*http.Request, error) {
	d.mux.Lock()
	defer d.mux.Unlock()
	u := d.c.Url
	p := d.c.STP
	if u == "" || p == "" {
		return nil, errors.Wrap(errors.NewMesssage("url or path is empty"))
	}
	return request(u, p, body)
}
func (d *depositeService) newDepositRequest(body io.Reader) (*http.Request, error) {
	d.mux.Lock()
	defer d.mux.Unlock()
	u := d.c.Url
	p := d.c.NDP
	if u == "" || p == "" {
		return nil, errors.Wrap(errors.NewMesssage("url or path is empty"))
	}
	return request(u, p, body)
}

func request(url, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, utils.JoinSafe(url, path), body)
	if err != nil {
		return nil, errors.Wrap(errors.NewMesssage(err.Error()))
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
