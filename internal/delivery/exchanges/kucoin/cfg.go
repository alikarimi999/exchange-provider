package kucoin

import "exchange-provider/pkg/errors"

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
	if cfg.ApiKey == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("api key is required"))
	}
	if cfg.ApiSecret == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("api secret is required"))
	}
	if cfg.ApiPassphrase == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("api passphrase is required"))
	}
	if cfg.ApiVersion == "" {
		cfg.ApiVersion = "2"
	}
	if cfg.ApiUrl == "" {
		cfg.ApiUrl = "https://api.kucoin.com"
	}

	return cfg, nil

}
