package http

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (s *Server) GetMinPairDeposit(ctx Context) {
	req := struct {
		C1 string `json:"coin1"`
		C2 string `json:"coin2"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	c1, c2 := s.app.GetMinPairDeposit(req.C1, req.C2)

	ctx.JSON(200, struct {
		MinC1 float64 `json:"min_coin1"`
		MinC2 float64 `json:"min_coin2"`
	}{
		MinC1: c1,
		MinC2: c2,
	})

}

func (s *Server) GetAllMinDeposit(ctx Context) {
	ctx.JSON(200, s.app.AllMinDeposit())
}

func (s *Server) ChangeMinDeposit(ctx Context) {
	req := struct {
		C1    string  `json:"coin1"`
		MinC1 float64 `json:"min_coin1"`
		C2    string  `json:"coin2"`
		MinC2 float64 `json:"min_coin2"`
		Msg   string  `json:"message,omitempty"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	if req.MinC1 <= 0 || req.MinC2 <= 0 {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("min deposit must be greater than 0")))
		return
	}

	if err := s.app.ChangeMinDeposit(&entity.PairMinDeposit{
		C1: &entity.CoinMinDeposit{Coin: req.C1, Min: req.MinC1},
		C2: &entity.CoinMinDeposit{Coin: req.C2, Min: req.MinC2}},
	); err != nil {
		handlerErr(ctx, err)
		return
	}

	req.Msg = "change was successful"
	ctx.JSON(200, req)
}
