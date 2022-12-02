package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/pkg/errors"
	"net/http"
)

func (s *Server) GetPairsToUser(ctx Context) {
	userId, _ := ctx.GetKey("user_id")
	req := &dto.GetPairsToUserRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp := &dto.GetPairsToUserResponse{}

	var dps []*dto.Pair
	var err error
	if len(req.Pairs) > 0 {
		dps, err = req.Parse()
		if err != nil {
			handlerErr(ctx, err)
			return
		}
	}

	pairs := make(map[string]*dto.UserPair)
	exs := s.app.AllExchanges()
	lenExs := len(exs)
	for i, ex := range exs {
		if len(dps) == 0 {
			ps := ex.GetAllPairs()
			for _, p := range ps {
				// if pair is exist in pairs, skip it
				if _, ok := pairs[p.String()]; ok {
					continue
				}

				dp := dto.EntityPairToUserRequest(s.app.ApplySpread(p), ex.Type())
				dp.FeeRate = s.app.GetUserFee(userId.(int64))
				dp.MinDepositCoin1, dp.MinDepositCoin2 = s.app.GetMinPairDeposit(p.C1.Coin.String(), p.C2.Coin.String())

				pairs[p.String()] = dp
			}
		} else {
			for _, p := range dps {
				// if pair is exist in pairs, skip it
				if _, ok := pairs[p.String()]; ok {
					continue
				}

				ep, err := ex.GetPair(p.Coin1, p.Coin2)
				if err != nil {
					if errors.ErrorCode(err) == errors.ErrNotFound && i == lenExs-1 {
						resp.Pairs = append(resp.Pairs, &dto.UserPair{
							Coin1: p.Coin1.String(),
							Coin2: p.Coin2.String(),
							Msg:   "pair not found",
						})
					}
					continue
				}

				dp := dto.EntityPairToUserRequest(s.app.ApplySpread(ep), ex.Type())
				dp.FeeRate = s.app.GetUserFee(userId.(int64))
				dp.MinDepositCoin1, dp.MinDepositCoin2 = s.app.GetMinPairDeposit(p.Coin1.String(), p.Coin2.String())
				pairs[p.String()] = dp

			}
		}
	}

	for _, p := range pairs {
		resp.Pairs = append(resp.Pairs, p)
	}

	ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetFeeToUser(ctx Context) {
	userId, _ := ctx.GetKey("user_id")

	ctx.JSON(http.StatusOK, s.app.GetUserFee(userId.(int64)))
}
