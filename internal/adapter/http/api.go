package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"fmt"
	"net"
	"strconv"
	"strings"
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

	if req.CheckIp && len(req.Ips) == 0 {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("you have to add at least one ip")))
		return
	}
	for _, ip := range req.Ips {
		if !isValidIP(ip) {
			ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
				fmt.Errorf("ip %s is not valid", ip)))
			return
		}
	}

	if uint(len(req.Ips)) > s.api.MaxIps() {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("only %d ips are allowed", s.api.MaxIps())))
		return
	}

	id := fmt.Sprintf("%s_%s", s.api.ApiPrefix(), utils.RandString(32))
	at := &entity.APIToken{
		Id:      utils.Hash(id),
		BusName: req.BusName,
		BusId:   req.BusId,
		Level:   req.Level,
		Write:   req.Write,
		CheckIp: req.CheckIp,
	}

	for _, ip := range req.Ips {
		if isIn(ip, at.Ips) {
			ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
				fmt.Errorf("ip %s is repetitive", ip)))
			return
		}
		at.Ips = append(at.Ips, ip)
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

	for _, ip := range req.Ips {
		if !isValidIP(ip) {
			ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
				fmt.Errorf("ip %s is not valid", ip)))
			return
		}
	}

	for _, ip := range req.Ips {
		if isIn(ip, at.Ips) {
			ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
				fmt.Errorf("ip %s is repetitive", ip)))
			return
		}
		at.Ips = append(at.Ips, ip)
	}

	if uint(len(at.Ips)) > s.api.MaxIps() {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("only %d ips are allowed", s.api.MaxIps())))
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

func (s *Server) GetApi(ctx Context) {
	id := ctx.Param("id")
	if strings.Contains(id, s.api.ApiPrefix()) {
		ctx.JSON(s.api.Get(utils.Hash(id)))
		return
	}
	busId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, fmt.Errorf("'%s' is invalid", id)))
		return
	}
	ctx.JSON(s.api.GetByBusId(uint(busId)))
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
		if !isIn(ip, at.Ips) {
			ctx.JSON(nil, errors.Wrap(errors.ErrNotFound,
				fmt.Errorf("ip %s not exists", ip)))
			return
		}
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

func isIn(s string, ss []string) bool {
	for _, si := range ss {
		if si == s {
			return true
		}
	}
	return false
}
func isValidIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}
