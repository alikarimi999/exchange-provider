package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"fmt"
)

func (s *Server) GenerateAPIToken(ctx Context) {
	req := &dto.CreateApiReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if req.BusId == 0 {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, fmt.Errorf("busId cannot be '0'")))
		return
	}

	if uint(len(req.Ips)) > s.api.MaxIps() {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("only %d ips allowed", s.api.MaxIps())))
		return
	}

	if req.CheckIp && len(req.Ips) == 0 {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("you have to add at least one ip")))
		return
	}

	id := utils.RandString(32)
	at := &entity.APIToken{
		Id:      utils.Hash(id),
		BusName: req.BusName,
		BusId:   req.BusId,
		Level:   req.Level,
		Ips:     req.Ips,
		Write:   req.Write,
		CheckIp: req.CheckIp,
	}

	if err := s.api.AddApiToken(at); err != nil {
		ctx.JSON(nil, err)
		return
	}
	req.Key = id
	ctx.JSON(req, nil)
}

func (s *Server) AddIP(ctx Context) {
	req := &dto.CreateApiReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}
	if req.Key == "" {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("apiKey is required")))
		return
	}
	if len(req.Ips) == 0 {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("ips is required")))
		return
	}

	at, err := s.api.Get(utils.Hash(req.Key))
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	if uint(len(at.Ips)+len(req.Ips)) > s.api.MaxIps() {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("you can only add %d IPs", s.api.MaxIps())))
		return
	}

	at.Ips = append(at.Ips, req.Ips...)
	if err := s.api.Update(at); err != nil {
		ctx.JSON(nil, err)
		return
	}

	req.BusId = at.BusId
	req.Ips = at.Ips
	req.Level = at.Level
	req.Write = at.Write
	req.CheckIp = at.CheckIp
	ctx.JSON(req, nil)
}

func (s *Server) RemoveIp(ctx Context) {
	req := &dto.CreateApiReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if req.Key == "" {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("apiKey is required")))
		return
	}
	if len(req.Ips) == 0 {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("ips is required")))
		return
	}

	at, err := s.api.Get(utils.Hash(req.Key))
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	for _, ip := range req.Ips {
		for i := 0; i < len(at.Ips); i++ {
			if ip == at.Ips[i] {
				at.Ips = append(at.Ips[:i], at.Ips[i+1:]...)
				i--
			}
		}
	}

	if at.CheckIp && len(at.Ips) == 0 {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, fmt.Errorf("you cannot remove all ips")))
		return
	}

	if err := s.api.Update(at); err != nil {
		ctx.JSON(nil, err)
		return
	}

	req.BusId = at.BusId
	req.Ips = at.Ips
	req.Level = at.Level
	req.Write = at.Write
	req.CheckIp = at.CheckIp
	ctx.JSON(req, nil)
}

func (s *Server) UpdateLevel(ctx Context) {
	req := &dto.CreateApiReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}
	if req.Key == "" {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("apiKey is required")))
		return
	}

	at, err := s.api.Get(utils.Hash(req.Key))
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	if req.Level == at.Level {
		ctx.JSON(nil, fmt.Errorf("apiKey level already is %d", at.Level))
		return
	}
	at.Level = req.Level
	if err := s.api.Update(at); err != nil {
		ctx.JSON(nil, err)
		return
	}

	req.BusId = at.BusId
	req.Ips = at.Ips
	req.Level = at.Level
	req.Write = at.Write
	req.CheckIp = at.CheckIp
	ctx.JSON(req, nil)
}

func (s *Server) Remove(ctx Context) {
	id := ctx.Param("id")
	if err := s.api.Remove(utils.Hash(id)); err != nil {
		ctx.JSON(nil, err)
		return
	}
	ctx.JSON(struct {
		Msg string `json:"msg"`
	}{Msg: "done"}, nil)
}
