package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"net/http"
	"sync"
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
	pmux := &sync.RWMutex{}
	wg := &sync.WaitGroup{}

	exs := s.app.AllExchanges()
	lenExs := len(exs)
	for i, ex := range exs {
		if len(dps) == 0 {
			wg.Add(1)
			go func(ex entity.Exchange) {
				defer wg.Done()

				ps := ex.GetAllPairs()
				for _, p := range ps {
					// if pair is exist in pairs, skip it
					pmux.RLock()
					if _, ok := pairs[p.String()]; ok {
						pmux.RUnlock()
						continue
					}
					pmux.RUnlock()

					dp := dto.EntityPairToUserRequest(s.app.ApplySpread(p), ex.Type())
					dp.FeeRate = s.app.GetUserFee(userId.(int64))
					dp.MinDepositToken1, dp.MinDepositToken2 = s.app.GetMinPairDeposit(p.T1.String(), p.T2.String())

					pmux.Lock()
					pairs[p.String()] = dp
					pmux.Unlock()
				}
			}(ex)
		} else {
			for _, p := range dps {
				// if pair is exist in pairs, skip it
				if _, ok := pairs[p.String()]; ok {
					continue
				}

				ep, err := ex.GetPair(p.T1, p.T2)
				if err != nil {
					if errors.ErrorCode(err) == errors.ErrNotFound && i == lenExs-1 {
						resp.Pairs = append(resp.Pairs, &dto.UserPair{
							T1:  p.T1.String(),
							T2:  p.T2.String(),
							Msg: "pair not found",
						})
					}
					continue
				}

				dp := dto.EntityPairToUserRequest(s.app.ApplySpread(ep), ex.Type())
				dp.FeeRate = s.app.GetUserFee(userId.(int64))
				dp.MinDepositToken1, dp.MinDepositToken2 = s.app.GetMinPairDeposit(p.T1.String(), p.T2.String())
				pairs[p.String()] = dp

			}
		}
	}
	wg.Wait()

	for _, p := range pairs {
		resp.Pairs = append(resp.Pairs, p)
	}

	ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetFeeToUser(ctx Context) {
	userId, _ := ctx.GetKey("user_id")

	ctx.JSON(http.StatusOK, s.app.GetUserFee(userId.(int64)))
}
