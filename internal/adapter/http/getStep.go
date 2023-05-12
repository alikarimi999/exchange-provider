package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strconv"
)

func (s *Server) GetStep(ctx Context) {
	id := ctx.Param("orderId")
	sParam := ctx.Param("step")

	if sParam == "" {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("need orderId parameter")))
	}
	step, err := strconv.Atoi(sParam)
	if err != nil {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
	}

	oId, err := dto.ParseId(id, entity.PrefOrder)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	ord, err := s.repo.Get(oId)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	switch ord.Type() {
	case entity.CEXOrder:
		ctx.JSON(dto.SingleStepResponse(ord))
		return
	case entity.EVMOrder:
		if step != 1 {
			ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("step out of range")))
			return
		}
		tx, err := s.app.GetMultiStep(ord, uint(step))
		if err != nil {
			ctx.JSON(nil, err)
			return
		}
		ctx.JSON(dto.MultiStep(ord, tx, step, step), nil)
		return
	}
}
