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
		from = p.T1
		to = p.T2
	} else {
		from = p.T2
		to = p.T1
	}
	return p, from.ET.(*Token), to.ET.(*Token), nil
}
