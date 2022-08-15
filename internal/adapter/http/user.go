package http

import (
	"net/http"
	"order_service/internal/adapter/http/dto"
	"order_service/internal/app"
	"order_service/pkg/errors"
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
		if ex.CurrentStatus != app.ExchangeStatusActive {
			continue
		}
		if len(dps) == 0 {
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
		} else {
			for _, p := range dps {
				// if pair is exist in pairs, skip it
				if _, ok := pairs[p.String()]; ok {
					continue
				}

				ep, err := ex.GetPair(p.BC, p.QC)
				if err != nil {
					if errors.ErrorCode(err) == errors.ErrNotFound && i == lenExs-1 {
						resp.Pairs = append(resp.Pairs, &dto.UserPair{
							BC:  p.BC.String(),
							QC:  p.QC.String(),
							Msg: "pair not found",
						})
					}
					continue
				}

				dp := dto.EntityPairToUserRequest(s.app.ApplySpread(ep))
				dp.FeeRate = s.app.GetUserFee(userId.(int64))
				dp.MinBaseCoinDeposit, dp.MinQuoteCoinDeposit = s.app.GetMinPairDeposit(p.BC, p.QC)
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
