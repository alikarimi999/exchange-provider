package evm

import (
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"os"
)

type tokens struct {
	Tokens []types.Token `json:"tokens"`
}

func (d *EvmDex) Tokens() []*entity.Token {
	ts := []*entity.Token{}
	for _, t := range d.ts.Tokens {
		ts = append(ts, t.ToToken())
	}
	return ts
}

func (d *EvmDex) retreiveTokens() error {
	b, err := os.ReadFile(d.TokensFile)
	if err != nil {
		return err
	}

	ts := &tokens{}
	if err := json.Unmarshal(b, ts); err != nil {
		return err
	}

	d.ts = ts
	return nil
}

func (d *EvmDex) get(t string) (*types.Token, error) {
	native := t == d.NativeToken
	if native {
		t = d.WrappedNativeToken
	}
	for _, t0 := range d.ts.Tokens {
		if t0.ChainId == d.ChainId && t0.Symbol == t {
			if native {
				t0.Symbol = d.NativeToken
				t0.Native = true
			}
			return t0.SnapShot(), nil
		}
	}
	return nil, errors.Wrap(errors.ErrNotFound)
}
