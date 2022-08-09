package http

import (
	"fmt"
	"order_service/internal/adapter/http/dto"
)

func (s *Server) GetAllPairsSpread(ctx Context) {
	ctx.JSON(200, s.app.GetAllPairsSpread())
}

func (s *Server) ChangePairSpread(ctx Context) {

	type pair struct {
		BC     string  `json:"base_coin"`
		QC     string  `json:"quote_coin"`
		Spread float64 `json:"spread"`
		Msg    string  `json:"msg"`
	}

	req := struct {
		Pairs []*pair `json:"pairs"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	resp := []*pair{}
	for _, p := range req.Pairs {

		bc, err := dto.ParseCoin(p.BC)
		if err != nil {
			p.Msg = err.Error()
			resp = append(resp, p)
			continue
		}

		qc, err := dto.ParseCoin(p.QC)
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

		p.Msg = fmt.Sprintf("spread changed to %f", p.Spread)
		resp = append(resp, p)

	}
	ctx.JSON(200, resp)
}

func (s *Server) GetDefaultSpread(ctx Context) {
	ctx.JSON(200, s.app.GetDefaultSpread())
}

func (s *Server) ChangeDefaultSpread(ctx Context) {
	req := struct {
		Spread float64 `json:"default_spread_rate"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	if req.Spread <= 0 || req.Spread >= 1 {
		ctx.JSON(500, "default spread rate must be > 0 and < 1")
		return
	}

	if err := s.app.ChangeDefaultSpread(req.Spread); err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, fmt.Sprintf("default spread rate changed to %f", req.Spread))
}
