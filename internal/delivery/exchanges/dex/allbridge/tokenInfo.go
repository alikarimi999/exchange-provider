package allbridge

import (
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"io"
	"net/http"
)

type tokenList map[string]types.Chain

func (tl tokenList) getTokenInfo(t entity.TokenId) (*types.TokenInfo, error) {
	c, ok := tl[t.Network]
	if ok {
		for _, t0 := range c.Tokens {
			if t.Symbol == t0.Symbol {
				return t0, nil
			}
		}
	}

	return nil, errors.Wrap(errors.ErrNotFound,
		errors.NewMesssage(fmt.Sprintf("token %s not found", t.String())))
}

func (tl tokenList) isTokenExists(t entity.TokenId) bool {
	c, ok := tl[t.Network]
	if ok {
		for _, t0 := range c.Tokens {
			if t.Symbol == t0.Symbol {
				return true
			}
		}
	}
	return false
}

func (tl tokenList) tokensInNetwork(n string) ([]*types.TokenInfo, error) {
	c, ok := tl[n]
	if ok {
		return c.Tokens, nil
	}
	return nil, errors.Wrap(errors.ErrNotFound,
		errors.NewMesssage(fmt.Sprintf("network %s not found", n)))
}

func getTokenInfo(ns mapCfgNetwork) (map[string]types.Chain, error) {
	url := "https://core.api.allbridgecoreapi.net/token-info"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	tokens := make(map[string]types.Chain)
	if err := json.Unmarshal(b, &tokens); err != nil {
		return nil, err
	}

	chains := make(map[string]types.Chain)

	for id, c := range tokens {
		if s, ok := ns[id]; ok {
			for _, t := range c.Tokens {
				t.Chain = id
				t.Network = s.Network
				t.Standard = s.Standard
				t.ChainId = c.ChainID
				t.TransferTime = make(map[string]types.TransferTime)
			}
			for nid, n := range ns {
				for _, t := range c.Tokens {
					tt, ok := c.TransferTime[nid]
					if ok {
						t.TransferTime[n.Network] = tt
					}
				}
			}
			chains[s.Network] = c

		}
	}
	return chains, nil
}
