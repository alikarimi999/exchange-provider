package http

import (
	"exchange-provider/internal/entity"
)

func (s *Server) routing(in, out *entity.Token) (map[int]*entity.Route, error) {
	ex, err := s.app.SelectExchangeByPair(in, out)
	if err != nil {
		return nil, err
	}
	routes := make(map[int]*entity.Route)
	routes[0] = &entity.Route{In: in, Out: out, Exchange: ex.Id(), ExType: ex.Type()}
	return routes, nil
}
