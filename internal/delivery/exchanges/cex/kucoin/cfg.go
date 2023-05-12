package kucoin

import (
	"exchange-provider/pkg/errors"
)

type API struct {
	ApiKey        string `json:"apiKey"`
	ApiSecret     string `json:"apiSecret"`
	ApiPassphrase string `json:"apiPassphrase"`
}

type Configs struct {
	Id       uint `json:"id"`
	Enable   bool `json:"enable"`
	ReadApi  *API `json:"readApi,omitempty"`
	WriteApi *API `json:"writeApi,omitempty"`

	ApiVersion string `json:"apiVersion"`
	ApiUrl     string `json:"apiUrl"`
	Message    string `json:"message"`
}

func (k *kucoinExchange) Configs() interface{} {
	return k.cfg
}

func validateConfigs(cfgi interface{}) (*Configs, error) {
	if cfgi == nil {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("configs is nil"))
	}

	cfg, ok := cfgi.(*Configs)
	if !ok {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid configs"))
	}

	if cfg.Id == 0 {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("id is required"))
	}
	if cfg.ReadApi.ApiKey == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("readApi.apiKey is required"))
	}
	if cfg.ReadApi.ApiSecret == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("readApi.apiSecret is required"))
	}
	if cfg.ReadApi.ApiPassphrase == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("readApi.apiPassphrase is required"))
	}

	if cfg.WriteApi.ApiKey == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("writeApi.apiKey is required"))
	}
	if cfg.WriteApi.ApiSecret == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("writeApi.apiSecret is required"))
	}
	if cfg.WriteApi.ApiPassphrase == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("writeApi.apiPassphrase is required"))
	}

	if cfg.ApiVersion == "" {
		cfg.ApiVersion = "2"
	}
	if cfg.ApiUrl == "" {
		cfg.ApiUrl = "https://api.kucoin.com"
	}

	return cfg, nil

}
