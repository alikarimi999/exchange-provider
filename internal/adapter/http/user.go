package http

import (
	"net/http"
	"order_service/internal/adapter/http/dto"
	"order_service/internal/app"
)

func (s *Server) GetPairsToUser(ctx Context) {
	userId, _ := ctx.GetKey("user_id")

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

			dp := dto.EntityPairToUserRequest(s.app.ApplySpread(p))
			dp.FeeRate = s.app.GetUserFee(userId.(int64))
			dp.MinBaseCoinDeposit, dp.MinQuoteCoinDeposit = s.app.GetMinPairDeposit(p.BC.Coin, p.QC.Coin)

			pairs[p.String()] = dp
		}
	}

	ps := []*dto.UserPair{}
	for _, p := range pairs {
		ps = append(ps, p)
	}

	ctx.JSON(http.StatusOK, ps)
}

func (s *Server) GetFeeToUser(ctx Context) {
	userId, _ := ctx.GetKey("user_id")

	ctx.JSON(http.StatusOK, s.app.GetUserFee(userId.(int64)))
}
