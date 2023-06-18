package http

import (
	"exchange-provider/internal/adapter/http/dto"
	bdto "exchange-provider/internal/delivery/exchanges/cex/binance/dto"
	kdto "exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"

	edto "exchange-provider/internal/delivery/exchanges/dex/evm/dto"
	"exchange-provider/internal/entity"
)

func (s *Server) AddPairs(ctx Context) {
	ids := ctx.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	ex, err := s.app.GetExchange(uint(id))
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	res := &entity.AddPairsResult{}
	var req interface{}
	switch ex.Type() {
	case entity.CEX:
		switch ex.Name() {
		case "kucoin":
			req = &kdto.AddPairsRequest{}
		case "binance":
			req = &bdto.AddPairsRequest{}
		}
	case entity.EvmDEX:
		req = &edto.AddPairsRequest{}
	}

	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}
	res, err = ex.AddPairs(req)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	ctx.JSON(dto.FromEntity(res), nil)
}

func (s *Server) GetPairs(ctx Context) {
	api := ctx.GetApi()
	admin := api == nil
	req := &dto.PaginatedReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	pa := req.Map()
	if err := s.pairs.GetPaginated(pa, admin); err != nil {
		fmt.Println(err)
		ctx.JSON(nil, err)
		return
	}

	var totalPage int64
	if pa.Page != 0 && pa.PerPage != 0 {
		totalPage = pa.Total / pa.PerPage
		if pa.Total%pa.PerPage > 0 {
			totalPage++
		}
	}
	res := &dto.PaginatedPairsResp{
		PaginatedResponse: dto.PaginatedResponse{
			CurrentPage: pa.Page,
			PageSize:    int64(len(pa.Pairs)),
			TotalNum:    pa.Total,
			TotalPage:   totalPage,
		},
	}

	if admin {
		res.Pairs = pa.Pairs
	} else {
		ps := []dto.Pair{}
		for _, p := range pa.Pairs {
			ps = append(ps, dto.Pair{
				T1:          dto.TokenFromEntity(p.T1),
				T2:          dto.TokenFromEntity(p.T2),
				Enable:      p.Enable,
				FeeRate1:    p.FeeRate1,
				FeeRate2:    p.FeeRate2,
				ExchangeFee: p.ExchangeFee,
				LP:          p.LP,
			})
		}
		res.Pairs = ps
	}
	ctx.JSON(res, nil)
}

func (s *Server) UpdatePairs(ctx Context) {
	req := &dto.UpdatePairReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	res := &dto.CmdResp{}
	for _, p := range req.Pairs {
		resp := struct {
			Pair string "json:\"pair\""
			Msg  string "json:\"msg\""
		}{
			Pair: fmt.Sprintf("%s/%s", p.T1.String(), p.T2.String()),
		}

		ex, err := s.app.GetExchange(p.LP)
		if err != nil {
			resp.Msg = err.Error()
			res.PairsRes = append(res.PairsRes, resp)
			continue
		}

		p.T1.ToUpper()
		p.T2.ToUpper()
		dp, err := s.pairs.Get(ex.Id(), p.T1.String(), p.T2.String())
		if err != nil {
			resp.Msg = err.Error()
			res.PairsRes = append(res.PairsRes, resp)
			continue
		}
		p.Update(dp, ex.Name(), req.AcceptZero)
		if err := s.pairs.Update(ex.Id(), dp); err != nil {
			resp.Msg = err.Error()
			res.PairsRes = append(res.PairsRes, resp)
			continue
		}
		resp.Msg = "done"
		res.PairsRes = append(res.PairsRes, resp)
	}
	ctx.JSON(res, nil)
}

func (s *Server) CommandPairs(ctx Context) {
	req := &dto.PairsRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if req.Cmd != "enable" && req.Cmd != "disable" && req.Cmd != "remove" {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("cmd '%s' is not supported", req.Cmd)))
		return
	}

	if req.All {
		err := s.pairs.UpdateAll(req.Cmd)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}
		ctx.JSON(struct {
			Msg string "json:\"msg\""
		}{Msg: "done"}, nil)
		return
	}

	res := &dto.CmdResp{}
	for _, p := range req.Pairs {
		resp := struct {
			Pair string "json:\"pair\""
			Msg  string "json:\"msg\""
		}{
			Pair: fmt.Sprintf("%s/%s", p.T1.String(), p.T2.String()),
		}
		ex, err := s.app.GetExchange(p.LP)
		if err != nil {
			resp.Msg = err.Error()
			res.PairsRes = append(res.PairsRes, resp)
			continue
		}
		p.T1.ToUpper()
		p.T2.ToUpper()

		switch req.Cmd {
		case dto.RemoveCmd:
			err := s.app.RemovePair(ex, p.T1, p.T2)
			if err != nil {
				resp.Msg = err.Error()
			} else {
				resp.Msg = "done"
			}
		case dto.EnableCmd:
			ep, err := s.pairs.Get(ex.Id(), p.T1.String(), p.T2.String())
			if err != nil {
				resp.Msg = err.Error()
				break
			}
			if ep.Enable {
				resp.Msg = "pair is enable"
				break
			}
			ep.Enable = true
			s.pairs.Update(ex.Id(), ep)
			resp.Msg = "done"
		case dto.DisableCmd:
			ep, err := s.pairs.Get(ex.Id(), p.T1.String(), p.T2.String())
			if err != nil {
				resp.Msg = err.Error()
				break
			}
			if !ep.Enable {
				resp.Msg = "pair is disable"
				break
			}
			ep.Enable = false
			s.pairs.Update(ex.Id(), ep)
			resp.Msg = "done"
		}
		res.PairsRes = append(res.PairsRes, resp)
	}

	ctx.JSON(res, nil)
}
