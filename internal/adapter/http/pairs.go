package http

import (
	"exchange-provider/internal/adapter/http/dto"
	udto "exchange-provider/internal/delivery/exchanges/dex/dto"
	"exchange-provider/internal/delivery/exchanges/dex/multichain"
	kdto "exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/internal/entity"
	"sync"
)

func (s *Server) AddPairs(ctx Context) {
	nid := ctx.Param("id")
	ex, err := s.app.GetExchange(nid)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	res := &entity.AddPairsResult{}

	switch ex.Name() {
	case "kucoin":
		req := &dto.KucoinAddPairsRequest{}
		if err := ctx.Bind(req); err != nil {
			ctx.JSON(nil, err)
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(nil, err)
			return
		}

		kps := &kdto.AddPairsRequest{}
		for _, p := range req.Pairs {
			kps.Pairs = append(kps.Pairs, p.Map())
		}

		res, err = s.app.AddPairs(ex, kps)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}

	case "uniswapv3", "panckakeswapv2":
		req := &udto.AddPairsRequest{}
		if err := ctx.Bind(req); err != nil {
			ctx.JSON(nil, err)
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(nil, err)
			return
		}

		res, err = s.app.AddPairs(ex, req)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}

	case "multichain":
		req := &multichain.AddPairsRequest{}
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

func (s *Server) GetPairsToAdmin(ctx Context) {
	req := &dto.PaginatedPairsRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}
	if err := req.Validate(true); err != nil {
		ctx.JSON(nil, err)
		return
	}

	pa := req.ToEntity()
	if err := s.pairs.GetPaginated(pa); err != nil {
		ctx.JSON(nil, err)
		return
	}

	exs := make(map[string]entity.Exchange)
	ps := []*entity.Pair{}
	wg := &sync.WaitGroup{}
	for _, p := range pa.Pairs {
		ex, ok := exs[p.Exchange]
		if !ok {
			var err error
			ex, err = s.app.GetExchange(p.Exchange)
			if err != nil {
				continue
			}
			exs[p.Exchange] = ex
		}
		wg.Add(1)
		go func(p *entity.Pair, ex entity.Exchange) {
			defer wg.Done()
			if req.Price {
				var err error
				p, err = ex.Price(p.T1.Token, p.T2.Token)
				if err != nil {
					return
				}
			}

			p.T1.MinDeposit, p.T2.MinDeposit = s.app.GetMinPairDeposit(p.T1.String(), p.T2.String())
			p.SpreadRate = s.app.GetPairSpread(p.T1.Token, p.T2.Token)
			ps = append(ps, p)
		}(p, ex)
	}
	wg.Wait()
	pa.Pairs = ps
	ctx.JSON(dto.PairsResp(pa, true), nil)
}

func (s *Server) RemovePair(ctx Context) {
	req := &dto.RemovePairRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	t1, t2, err := req.Parse()
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	ex, err := s.app.GetExchange(req.Exchange)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	err = s.app.RemovePair(ex, t1, t2, req.Force)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	req.Msg = "done"
	ctx.JSON(req, nil)
}
