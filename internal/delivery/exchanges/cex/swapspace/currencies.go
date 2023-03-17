package swapspace

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

type tokenList struct {
	*sync.RWMutex
	list []*token
}

func newTokenList() *tokenList {
	return &tokenList{
		RWMutex: &sync.RWMutex{},
	}
}

func (ex *exchange) getCurrencies() error {
	agent := ex.agent("getCurrencies")

	urlStr, _ := url.JoinPath(baseUrl, "currencies")
	b, err := ex.request(http.MethodGet, urlStr, nil)
	if err != nil {
		return err
	}

	tokens := []*token{}
	if err := json.Unmarshal(b, &tokens); err != nil {
		return err
	}

	ex.tokens.Lock()
	ex.tokens.list = tokens
	ex.tokens.Unlock()
	ex.l.Debug(agent, fmt.Sprintf("%d currencies available", len(tokens)))
	return nil
}
