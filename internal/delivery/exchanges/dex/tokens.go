package dex

import "exchange-provider/internal/entity"

func (d *dex) Tokens() []*entity.Token {
	ts := []*entity.Token{}
	for _, t := range d.tokens.getAll() {
		ts = append(ts, t.ToToken())
	}
	return ts
}
