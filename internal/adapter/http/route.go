package http

import (
	"exchange-provider/internal/entity"
	"fmt"
)

func (s *Server) routing(in, out *entity.Coin) (map[int]*entity.Route, error) {
	routes := make(map[int]*entity.Route)
	if in.CoinId == out.CoinId || in.ChainId == out.ChainId {

		routes[0] = &entity.Route{In: in, Out: out}
		ex, err := s.app.SelectExchangeByPair(in, out)
		if err != nil {
			return nil, err
		}
		routes[0].Exchange = ex.NID()
		return routes, nil

	} else if s.cf.lowerEq(in.ChainId, out.ChainId) {

		routes[0] = &entity.Route{In: in, Out: &entity.Coin{CoinId: out.CoinId, ChainId: in.ChainId}}

		ex, err := s.app.SelectExchangeByPair(routes[0].In, routes[0].Out)
		if err == nil {
			routes[0].Exchange = ex.NID()
			routes[1] = &entity.Route{In: routes[0].Out, Out: out}
			ex, err := s.app.SelectExchangeByPair(routes[1].In, routes[1].Out)
			if err == nil {
				routes[1].Exchange = ex.NID()
			}
			return routes, nil
		}

	}

	routes[0] = &entity.Route{In: in, Out: &entity.Coin{CoinId: in.CoinId, ChainId: out.ChainId}}
	routes[1] = &entity.Route{In: routes[0].Out, Out: out}
	ex, err := s.app.SelectExchangeByPair(routes[0].In, routes[0].Out)
	if err != nil {
		return nil, fmt.Errorf("NoRouteFound")
	}
	routes[0].Exchange = ex.NID()
	ex, err = s.app.SelectExchangeByPair(routes[1].In, routes[1].Out)
	if err != nil {
		return nil, fmt.Errorf("NoRouteFound")
	}
	routes[1].Exchange = ex.NID()
	return routes, nil
}
