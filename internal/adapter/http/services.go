package http

func (s *Server) ChangeDepositServiceConfig(cfg interface{}) error {
	return s.app.DS.ChangeConfigs(cfg)
}

func (s *Server) GetDepositServiceConfig() interface{} {
	return s.app.DS.Configs()
}
