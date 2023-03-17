package swapspace

import (
	"bytes"
	"encoding/json"
	"exchange-provider/pkg/errors"
	"io"
	"net/http"
)

func (p *exchange) request(method, url string, data interface{}) ([]byte, error) {
	agent := p.agent("request")

	var (
		body io.Reader
		b    []byte
	)

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			p.l.Error(agent, err.Error())
			return nil, errors.Wrap(errors.ErrInternal, err)
		}
		body = bytes.NewBuffer(b)
	} else {
		body = nil
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		p.l.Error(agent, err.Error())
		return nil, errors.Wrap(errors.ErrInternal, err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", p.ApiKey)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		p.l.Error(agent, err.Error())
		return nil, errors.Wrap(errors.ErrInternal, err)
	}

	b, err = io.ReadAll(rsp.Body)
	if err != nil {
		p.l.Error(agent, err.Error())
		return nil, errors.Wrap(errors.ErrInternal, err)
	}

	switch rsp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return b, nil
	case http.StatusBadRequest:
		return nil, errors.Wrap(errors.ErrBadRequest, errors.New(string(b)))
	default:
		return nil, errors.Wrap(rsp.Status)
	}

}
