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

	if err := s.app.ChangeDefaultFee(req.DefaultFeeRate); err != nil {
		ctx.JSON(nil, err)
		return
	}
	req.Msg = "done"
	ctx.JSON(req, nil)
}

func (s *Server) GetDefaultFee(ctx Context) {
	ctx.JSON(
		struct {
			D string `json:"defaultFee"`
		}{
			D: s.app.GetDefaultFee(),
		}, nil)
}

func (s *Server) GetUsersFee(ctx Context) {
	req := struct {
		Users []uint64 `json:"users"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if len(req.Users) == 0 {
		ctx.JSON(s.app.GetAllUsersFee(), nil)
		return
	}

	resp := make(map[uint64]string)

	for _, userId := range req.Users {
		resp[userId] = s.app.GetUserFee(userId)
	}

	ctx.JSON(resp, nil)
}

func (s *Server) ChangeUserFee(ctx Context) {
	type userFee struct {
		Id uint64  `json:"user_id"`
		F  float64 `json:"fee_rate"`
		M  string  `json:"message"`
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
	ctx.JSON(resp, nil)
}
