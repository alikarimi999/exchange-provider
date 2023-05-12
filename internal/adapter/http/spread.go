package http

import (
	"exchange-provider/internal/entity"
)

func (s *Server) AddSpread(ctx Context) {
	req := struct {
		Tables map[uint][]*entity.Spread `json:"tables"`
		Msg    string                    `json:"msg"`
	}{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	st, err := s.spread.Add(req.Tables)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	req.Tables = st
	req.Msg = "done"
	ctx.JSON(req, nil)
}

func (s *Server) GetAll(ctx Context) {
	ctx.JSON(s.spread.GetAll(), nil)
}

func (s *Server) RemoveSpread(ctx Context) {
	req := struct {
		Levels []uint `json:"levels"`
		Msg    string `json:"msg"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	err := s.spread.Remove(req.Levels)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	req.Msg = "done"
	ctx.JSON(req, nil)

}
