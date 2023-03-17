package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"sort"
)

func (s *Server) Tokens(ctx Context) {
	mts := make(map[string]*dto.Token)
	ts := []*dto.Token{}
	exs := s.app.AllExchanges()
	for _, ex := range exs {
		tokens := ex.Tokens()
		for _, t := range tokens {
			mts[t.TokenId+t.ChainId] = &dto.Token{
				TokenId: t.TokenId,
				ChainId: t.ChainId,
				LP:      ex.Id(),
			}
		}
	}

	for _, t := range mts {
		ts = append(ts, t)
	}
	sort.Slice(ts, func(i, j int) bool {
		ti := ts[i]
		tj := ts[j]
		return ti.ChainId < tj.ChainId
	})
	ctx.JSON(ts, nil)
}
