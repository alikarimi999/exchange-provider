package swapspace

import "exchange-provider/pkg/errors"

type Config struct {
	Id      uint   `json:"id"`
	ApiKey  string `json:"apiKey"`
	Message string
}

func (cfg *Config) Validate() error {
	if cfg.Id == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("id is required"))

	}
	if cfg.ApiKey == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("apiKey is required"))
	}
	return nil
}

func (ex *exchange) Configs() interface{} {
	return ex.Config
}
