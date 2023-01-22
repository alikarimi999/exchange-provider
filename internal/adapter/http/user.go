package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/entity"
	"sync"
)

func (s *Server) GetPairsToUser(ctx Context) {
	userId, _ := ctx.GetKey("user_id")
	req := &dto.PaginatedPairsRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if err := req.Validate(false); err != nil {
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
	pmux := &sync.Mutex{}
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
			p.FeeRate = s.app.GetUserFee(uint64(userId.(int64)))
			pmux.Lock()
			ps = append(ps, p)
			pmux.Unlock()
		}(p, ex)
	}
	wg.Wait()
	pa.Pairs = ps
	ctx.JSON(dto.PairsResp(pa, false), nil)
}
