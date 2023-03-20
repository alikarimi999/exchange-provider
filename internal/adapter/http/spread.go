package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/pkg/errors"
)

func (s *Server) GetAllPairsSpread(ctx Context) {
	ctx.JSON(s.app.GetAllPairsSpread(), nil)
}

func (s *Server) ChangePairSpread(ctx Context) {

	type pair struct {
		T1     dto.Token `json:"t1"`
		T2     dto.Token `json:"t2"`
		Spread float64   `json:"spread"`
		Msg    string    `json:"message"`
	}

	req := struct {
		Pairs []*pair `json:"pairs"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	resp := []*pair{}
	for _, p := range req.Pairs {

		t1 := p.T1.ToEntity()
		t2 := p.T2.ToEntity()

		if p.Spread <= 0 || p.Spread >= 1 {
			p.Msg = "spread rate must be > 0 and < 1"
			resp = append(resp, p)
			continue
		}

		if err := s.app.ChangePairSpread(t1, t2, p.Spread); err != nil {
			p.Msg = err.Error()
			resp = append(resp, p)
			continue
		}

		p.Msg = "done"
		resp = append(resp, p)
	}
	ctx.JSON(resp, nil)
}

func (s *Server) GetDefaultSpread(ctx Context) {
	ctx.JSON(s.app.GetDefaultSpread(), nil)
}

func (s *Server) ChangeDefaultSpread(ctx Context) {
	req := struct {
		Spread float64 `json:"default_spread_rate"`
		Msg    string  `json:"message"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if req.Spread <= 0 || req.Spread >= 1 {
		err := errors.Wrap(errors.ErrBadRequest,
			errors.New("default spread rate must be > 0 and < 1"))
		ctx.JSON(nil, err)
		return
	}

	if err := s.app.ChangeDefaultSpread(req.Spread); err != nil {
		ctx.JSON(nil, err)
		return
	}
	req.Msg = "done"
	ctx.JSON(req, nil)
}
