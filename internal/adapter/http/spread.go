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
		T1     string  `json:"t1"`
		T2     string  `json:"t2"`
		Spread float64 `json:"spread"`
		Msg    string  `json:"message"`
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

		bc, err := dto.ParseToken(p.T1)
		if err != nil {
			p.Msg = err.Error()
			resp = append(resp, p)
			continue
		}

		qc, err := dto.ParseToken(p.T2)
		if err != nil {
			p.Msg = err.Error()
			resp = append(resp, p)
			continue
		}

		if p.Spread <= 0 || p.Spread >= 1 {
			p.Msg = "spread rate must be > 0 and < 1"
			resp = append(resp, p)
			continue
		}

		if err := s.app.ChangePairSpread(bc, qc, p.Spread); err != nil {
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
