package allbridge

type Config struct {
	Id   uint
	Name string
}

type allBridge struct {
	cfg *Config
	exs *intraExchanges
	ps  *pairs
}

func (a *allBridge) Id() uint     { return a.cfg.Id }
func (a *allBridge) Name() string { return a.cfg.Name }
