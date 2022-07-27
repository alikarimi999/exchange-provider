package kucoin

import "order_service/pkg/errors"

func (k *kucoinExchange) Configs() interface{} {
	return k.cfg
}

func validateConfigs(cfgi interface{}) (*Configs, error) {
	if cfgi == nil {
		return nil, errors.Wrap(errors.New("configs is nil"))
	}

	cfg, ok := cfgi.(*Configs)
	if !ok {
		return nil, errors.Wrap(errors.New("invalid configs"))
	}
	if cfg.ApiKey == "" {
		return nil, errors.New("api key is required")
	}
	if cfg.ApiSecret == "" {
		return nil, errors.New("api secret is required")
	}
	if cfg.ApiPassphrase == "" {
		return nil, errors.New("api passphrase is required")
	}
	if cfg.ApiVersion == "" {
		cfg.ApiVersion = "2"
	}
	if cfg.ApiUrl == "" {
		cfg.ApiUrl = "https://api.kucoin.com"
	}

	return cfg, nil

}
