package http

import (
	"fmt"
	"exchange-provider/pkg/errors"
)

func (s *Server) ChangeDefaultFee(ctx Context) {
	req := struct {
		DefaultFeeRate float64 `json:"default_fee_rate"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	if req.DefaultFeeRate <= 0 || req.DefaultFeeRate >= 1 {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("default fee rate must be > 0 and < 1")))
		return
	}

	if err := s.app.ChangeDefaultFee(req.DefaultFeeRate); err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, fmt.Sprintf("fee rate changed to %f", req.DefaultFeeRate))
}

func (s *Server) GetDefaultFee(ctx Context) {
	ctx.JSON(200, fmt.Sprintf("fee is %s", s.app.GetDefaultFee()))
}

func (s *Server) GetUsersFee(ctx Context) {
	req := struct {
		Users []int64 `json:"users"`
		All   bool    `json:"all"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	if req.All {
		ctx.JSON(200, s.app.GetAllUsersFee())
		return
	}

	resp := make(map[int64]string)

	for _, userId := range req.Users {
		resp[userId] = s.app.GetUserFee(userId)
	}

	ctx.JSON(200, resp)
}

func (s *Server) ChangeUserFee(ctx Context) {
	type userFee struct {
		Id int64   `json:"user_id"`
		F  float64 `json:"fee_rate"`
		M  string  `json:"msg"`
	}
	req := struct {
		Users []*userFee `json:"users"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	resp := struct {
		Users []*userFee `json:"users"`
	}{}

	for _, u := range req.Users {
		if u.Id == 0 || u.F == 0 {
			u.M = "user id or fee rate is empty"
			resp.Users = append(resp.Users, u)
			continue
		}

		if err := s.app.ChangeUserFee(u.Id, u.F); err != nil {
			u.M = err.Error()
			resp.Users = append(resp.Users, u)
			continue
		}
		u.M = "fee rate changed"
		resp.Users = append(resp.Users, u)
	}
	ctx.JSON(200, resp)
}
