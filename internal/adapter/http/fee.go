package http

import (
	"exchange-provider/pkg/errors"
)

func (s *Server) ChangeDefaultFee(ctx Context) {
	req := struct {
		DefaultFeeRate float64 `json:"default_fee_rate"`
		Msg            string  `json:"message"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if req.DefaultFeeRate <= 0 || req.DefaultFeeRate >= 1 {
		err := errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("default fee rate must be > 0 and < 1"))
		ctx.JSON(nil, err)
		return
	}

	if err := s.fee.ChangeDefaultFee(req.DefaultFeeRate); err != nil {
		ctx.JSON(nil, err)
		return
	}
	req.Msg = "done"
	ctx.JSON(req, nil)
}

func (s *Server) GetDefaultFee(ctx Context) {
	ctx.JSON(
		struct {
			D float64 `json:"defaultFee"`
		}{
			D: s.fee.GetDefaultFee(),
		}, nil)
}

func (s *Server) GetFees(ctx Context) {
	ctx.JSON(s.fee.GetAllBusFees(), nil)
}

func (s *Server) ChangeUserFee(ctx Context) {
	type userFee struct {
		UserId string  `json:"userId"`
		F      float64 `json:"feeRate"`
		M      string  `json:"message"`
	}
	req := struct {
		Users []*userFee `json:"users"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	resp := struct {
		Users []*userFee `json:"users"`
	}{}

	for _, u := range req.Users {
		if u.UserId == "" || u.F == 0 {
			u.M = "user id or fee rate is empty"
			resp.Users = append(resp.Users, u)
			continue
		}

		if err := s.fee.UpdateBusFee(u.UserId, u.F); err != nil {
			u.M = err.Error()
			resp.Users = append(resp.Users, u)
			continue
		}
		u.M = "fee rate changed"
		resp.Users = append(resp.Users, u)
	}
	ctx.JSON(resp, nil)
}
