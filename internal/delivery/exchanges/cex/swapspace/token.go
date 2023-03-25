package swapspace

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

type Token struct {
	Code       string
	Network    string
	HasExtraId bool
}

func (t *Token) Snapshot() entity.ExchangeToken {
	return &Token{
		Code:       t.Code,
		Network:    t.Network,
		HasExtraId: t.HasExtraId,
	}
}

func (ex *exchange) retrieveInOut(from, to *entity.Token) (p *entity.Pair, in, out *Token, err error) {
	p, ok := ex.pairs.Get(ex.Id(), from.String(), to.String())
	if !ok {
		return nil, nil, nil, errors.Wrap(errors.ErrNotFound)
	}

	if p.T1.Equal(from) {
		in = p.T1.ET.(*Token)
		out = p.T2.ET.(*Token)
	} else {
		in = p.T2.ET.(*Token)
		out = p.T1.ET.(*Token)
	}

	return p, in, out, nil
}
