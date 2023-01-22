package http

import (
	"exchange-provider/internal/entity"
)

func (s *Server) GetMinPairDeposit(ctx Context) {
	req := struct {
		T1 string `json:"t1"`
		T2 string `json:"t2"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	t1, t2 := s.app.GetMinPairDeposit(req.T1, req.T2)
	ctx.JSON(struct {
		MinT1 float64 `json:"minT1"`
		MinT2 float64 `json:"minT2"`
	}{
		MinT1: t1,
		MinT2: t2,
	}, nil)

}

func (s *Server) GetAllMinDeposit(ctx Context) {
	ctx.JSON(s.app.AllMinDeposit(), nil)
}

func (s *Server) ChangeMinDeposit(ctx Context) {
	req := struct {
		T1    string  `json:"t1"`
		MinT1 float64 `json:"minT1"`
		T2    string  `json:"t2"`
		MinT2 float64 `json:"minT2"`
		Msg   string  `json:"message"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if err := s.app.ChangeMinDeposit(&entity.PairMinDeposit{
		C1: &entity.CoinMinDeposit{Coin: req.T1, Min: req.MinT1},
		C2: &entity.CoinMinDeposit{Coin: req.T2, Min: req.MinT2}},
	); err != nil {
		ctx.JSON(nil, err)
		return
	}

	req.Msg = "done"
	ctx.JSON(req, nil)
}
