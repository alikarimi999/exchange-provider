package swapspace

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"time"
)

const (
	baseUrl = "https://api.swapspace.co/api/v2"
)

type exchange struct {
	*Config
	repo   entity.OrderRepo
	tokens *tokenList

	stopCh chan struct{}
	l      logger.Logger
}

func SwapSpace(cfg *Config, repo entity.OrderRepo, l logger.Logger) (entity.Cex, error) {

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	ex := &exchange{
		Config: cfg,
		tokens: newTokenList(),
		repo:   repo,
		stopCh: make(chan struct{}),
		l:      l,
	}
	return ex, ex.getCurrencies()
}

func (ex *exchange) Id() uint         { return ex.Config.Id }
func (*exchange) Name() string        { return entity.SwapSpace }
func (*exchange) Type() entity.ExType { return entity.CEX }
func (ex *exchange) Remove()          { ex.stopCh <- struct{}{} }
func (ex *exchange) Run() {
	t := time.NewTicker(6 * time.Hour)
	for {
		select {
		case <-t.C:
			ex.getCurrencies()
		case <-ex.stopCh:
			return
		}
	}
}
