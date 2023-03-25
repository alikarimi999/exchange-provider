package http

import (
	"exchange-provider/internal/adapter/http/dto"
	sdto "exchange-provider/internal/delivery/exchanges/cex/swapspace/dto"
	"strconv"
	"strings"

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
	switch strings.Split(ex.Name(), "-")[0] {
	case "swapspace":
		req := &sdto.AddPairsRequest{}
		if err := ctx.Bind(req); err != nil {
			ctx.JSON(nil, err)
			return
		}

		res, err = s.app.AddPairs(ex, req)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}

	// case "kucoin":
	// 	req := &dto.KucoinAddPairsRequest{}
	// 	if err := ctx.Bind(req); err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}

	// 	if err := req.Validate(); err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}

	// 	kps := &kdto.AddPairsRequest{}
	// 	for _, p := range req.Pairs {
	// 		kps.Pairs = append(kps.Pairs, p.Map())
	// 	}

	// 	res, err = s.app.AddPairs(ex, kps)
	// 	if err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}

	case "uniswapv3", "panckakeswapv2":
		req := &edto.AddPairsRequest{}
		if err := ctx.Bind(req); err != nil {
			ctx.JSON(nil, err)
			return
		}

		res, err = s.app.AddPairs(ex, req)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}
	}

	ctx.JSON(dto.FromEntity(res), nil)
}

func (s *Server) GePairsToUser(ctx Context) {
	req := &dto.PaginatedPairsRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	pa := req.ToEntity()
	if err := s.pairs.GetPaginated(pa); err != nil {
		ctx.JSON(nil, err)
		return
	}

	totalPage := pa.Total / pa.PerPage
	if pa.Total%pa.PerPage > 0 {
		totalPage++
	}
	res := &dto.PaginatedPairsResp{
		PaginatedResponse: dto.PaginatedResponse{
			CurrentPage: pa.Page,
			PageSize:    int64(len(pa.Pairs)),
			TotalNum:    pa.Total,
			TotalPage:   totalPage,
		},
	}

	for _, p := range pa.Pairs {
		res.Pairs = append(res.Pairs, dto.PairFromEntity(p))
	}
	ctx.JSON(res, nil)
}

// func (s *Server) GetPairsToAdmin(ctx Context) {
// req := &dto.PaginatedPairsRequest{}
// 	if err := ctx.Bind(req); err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}
// 	if err := req.Validate(true); err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}

// 	pa := req.ToEntity()
// 	if err := s.pairs.GetPaginated(pa); err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}

// 	exs := make(map[string]entity.Exchange)
// 	ps := []*entity.Pair{}
// 	for _, p := range pa.Pairs {
// 		ex, ok := exs[p.Exchange]
// 		if !ok {
// 			var err error
// 			ex, err = s.app.GetExchange(p.Exchange)
// 			if err != nil {
// 				continue
// 			}
// 			exs[p.Exchange] = ex
// 		}
// 		p.T1.MinDeposit, p.T2.MinDeposit = s.app.GetMinPairDeposit(p.T1.String(), p.T2.String())
// 		p.SpreadRate = s.app.GetPairSpread(p.T1.Token, p.T2.Token)
// 		ps = append(ps, p)
// 	}

// 	pa.Pairs = ps
// 	ctx.JSON(dto.PairsResp(pa, true), nil)
// }

// func (s *Server) RemovePair(ctx Context) {
// 	req := &dto.RemovePairRequest{}
// 	if err := ctx.Bind(req); err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}

// 	t1, t2, err := req.Parse()
// 	if err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}

// 	ex, err := s.app.GetExchange(req.Exchange)
// 	if err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}

// 	err = s.app.RemovePair(ex, t1, t2, req.Force)
// 	if err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}

// 	req.Msg = "done"
// 	ctx.JSON(req, nil)
// }
