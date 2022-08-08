package http

import (
	"net/http"
	"order_service/internal/adapter/http/dto"
	"order_service/internal/app"
)

func (s *Server) GetPairsToUser(ctx Context) {
	req := &dto.GetPairsToUserRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	pairs := make(map[string]*dto.UserPair)
	exs := s.app.AllExchanges()
	for _, ex := range exs {
		if ex.CurrentStatus != app.ExchangeStatusActive {
			continue
		}
		ps := ex.GetAllPairs()
		for _, p := range ps {
			// if pair is exist in pairs, skip it
			if _, ok := pairs[p.String()]; ok {
				continue
			}
			pairs[p.String()] = dto.EntityPairToUserRequest(p)
		}
	}

	ctx.JSON(http.StatusOK, pairs)
}
