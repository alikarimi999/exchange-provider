package http

// func (s *Server) GetPairsToUser(ctx Context) {
// 	// userId, _ := ctx.GetKey("user_id")
// 	req := &dto.PaginatedPairsRequest{}
// 	if err := ctx.Bind(req); err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}
// 	if err := req.Validate(false); err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}

// 	pa := req.ToEntity()
// 	if err := s.pairs.GetPaginated(pa); err != nil {
// 		ctx.JSON(nil, err)
// 		return
// 	}

// 	ps := []*entity.Pair{}
// 	for _, p := range pa.Pairs {
// 		p.FeeRate = s.app.GetDefaultFee()
// 		ps = append(ps, p)
// 	}

// 	pa.Pairs = ps
// 	ctx.JSON(dto.PairsResp(pa, false), nil)
// }
