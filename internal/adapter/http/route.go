package http

import (
	"exchange-provider/internal/entity"
)

func (s *Server) routing(in, out *entity.Token) (map[int]*entity.Route, error) {
	routes := make(map[int]*entity.Route)

	if ex, err := s.app.SelectExchangeByPair(in, out); err == nil {
		routes[0] = &entity.Route{In: in, Out: out}
		routes[0].Exchange = ex.Id()
		return routes, nil

	} else if s.cf.lowerEq(in.ChainId, out.ChainId) {
		routes[0] = &entity.Route{In: in, Out: &entity.Token{TokenId: out.TokenId, ChainId: in.ChainId}}
		ex, err := s.app.SelectExchangeByPair(routes[0].In, routes[0].Out)
		if err == nil {
			routes[0].Exchange = ex.Id()
			routes[1] = &entity.Route{In: routes[0].Out, Out: out}
			ex, err := s.app.SelectExchangeByPair(routes[1].In, routes[1].Out)
			if err == nil {
				routes[1].Exchange = ex.Id()
			}
			return routes, nil
		}
	}

	routes[0] = &entity.Route{In: in, Out: &entity.Token{TokenId: in.TokenId, ChainId: out.ChainId}}
	routes[1] = &entity.Route{In: routes[0].Out, Out: out}
	ex, err := s.app.SelectExchangeByPair(routes[0].In, routes[0].Out)
	if err != nil {
		return nil, err
	}
	routes[0].Exchange = ex.Id()
	ex, err = s.app.SelectExchangeByPair(routes[1].In, routes[1].Out)
	if err != nil {
		return nil, err
	}
	routes[1].Exchange = ex.Id()

	return routes, nil
}
