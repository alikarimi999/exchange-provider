package swapspace

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
)

const (
	baseUrl = "https://api.swapspace.co/api/v2"
)

type exchange struct {
	*Config
	repo  entity.OrderRepo
	pairs entity.PairsRepo
	l     logger.Logger
}

func SwapSpace(cfg *Config, repo entity.OrderRepo,
	pr entity.PairsRepo, l logger.Logger) (entity.Cex, error) {

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	ex := &exchange{
		Config: cfg,
		repo:   repo,
		pairs:  pr,
		l:      l,
	}
	return ex, nil
}

func (ex *exchange) Id() uint  { return ex.Config.Id }
func (*exchange) Name() string { return "swapspace" }
func (ex *exchange) NID() string {
	return fmt.Sprintf("%s-%d", ex.Name(), ex.Id())
}
func (*exchange) Type() entity.ExType { return entity.CEX }
func (ex *exchange) Remove()          {}
func (ex *exchange) Run()             {}
