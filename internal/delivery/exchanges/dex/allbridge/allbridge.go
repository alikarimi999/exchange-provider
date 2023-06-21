package allbridge

import (
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
)

type Config struct {
	Id     uint
	Url    string
	Name   string
	Enable bool
}

type allBridge struct {
	cfg *Config
	exs app.ExchangeRepo
	ps  *pairs

	l logger.Logger
}

func (a *allBridge) Id() uint                  { return a.cfg.Id }
func (a *allBridge) Name() string              { return "allbridge" }
func (a *allBridge) NID() string               { return fmt.Sprintf("%s-%d", a.Name(), a.cfg.Id) }
func (a *allBridge) EnableDisable(enable bool) { a.cfg.Enable = enable }
func (a *allBridge) IsEnable() bool            { return a.cfg.Enable }
func (a *allBridge) Type() entity.ExType       { return entity.CrossDex }
func (a *allBridge) NewOrder(interface{}, *entity.APIToken) (entity.Order, error)
func (a *allBridge) EstimateAmountOut(t1, t2 entity.TokenId, amount float64, lvl uint) (*entity.EstimateAmount, error)
func (a *allBridge) AddPairs(data interface{}) (*entity.AddPairsResult, error)
func (a *allBridge) RemovePair(t1, t2 entity.TokenId) error
func (a *allBridge) Configs() interface{}
func (a *allBridge) Remove()
