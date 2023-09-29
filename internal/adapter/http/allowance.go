package http

import (
	"exchange-provider/internal/adapter/http/dto"
)

func (s *Server) Allowance(ctx Context) {
	req := &dto.AllowanceReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	a, err := s.app.Allowance(*req.Token.ToUpper(), req.Owner)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	ctx.JSON(dto.AllowanceRes{
		Token: dto.AToken{
			TokenId:  a.Token.Id,
			Address:  a.Token.ContractAddress,
			Decimals: a.Token.Decimals,
		},
		Owner:   a.Owner,
		Spender: a.Spender,
		Amount:  a.Amount.String(),
	}, nil)
}
