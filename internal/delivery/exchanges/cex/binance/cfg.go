package binance

import (
	"exchange-provider/internal/entity"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

type API struct {
	ApiKey    string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
}

type Config struct {
	Id     uint `json:"id"`
	Enable bool `json:"enable"`

	Api     *API   `json:"api,omitempty"`
	Message string `json:"message,omitempty"`
}

func (ex *exchange) UpdateConfigs(cfgi interface{}, store entity.ExchangeStore) error {
	cfg, ok := cfgi.(*Config)
	if !ok {
		return fmt.Errorf("invalid config")
	}
	cfg, err := cfg.validate()
	if err != nil {
		return err
	}

	if cfg.Id != ex.cfg.Id {
		return fmt.Errorf("the id field is not mutable")
	}

	c := binance.NewClient(cfg.Api.ApiKey, cfg.Api.ApiSecret)
	if err := ping(c); err != nil {
		return err
	}
	cfg.Enable = ex.cfg.Enable
	if err := store.UpdateConfigs(ex, cfg); err != nil {
		return err
	}
	ex.cfg = cfg
	ex.c = c
	return nil
}

func (cfg *Config) validate() (*Config, error) {

	if cfg.Id == 0 {
		return nil, fmt.Errorf("id is required")
	}

	if cfg.Api == nil {
		return nil, fmt.Errorf("api is required")
	}

	if cfg.Api.ApiKey == "" {
		return nil, fmt.Errorf("apiKey is required")
	}

	if cfg.Api.ApiSecret == "" {
		return nil, fmt.Errorf("apiSecret is required")
	}
	return cfg, nil
}
