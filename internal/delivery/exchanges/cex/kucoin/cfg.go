package kucoin

import (
	"exchange-provider/internal/entity"
	"fmt"

	"github.com/Kucoin/kucoin-go-sdk"
)

type API struct {
	ApiKey        string `json:"apiKey"`
	ApiSecret     string `json:"apiSecret"`
	ApiPassphrase string `json:"apiPassphrase"`
}

type Config struct {
	Id       uint `json:"id"`
	Enable   bool `json:"enable"`
	ReadApi  *API `json:"readApi,omitempty"`
	WriteApi *API `json:"writeApi,omitempty"`

	CoinListUrl string `json:"coinListUrl"`

	ApiVersion string `json:"apiVersion"`
	ApiUrl     string `json:"apiUrl"`
	Message    string `json:"message"`
}

func (k *exchange) Configs() interface{} {
	return k.cfg
}

func (k *exchange) UpdateConfigs(cfgi interface{}, store entity.ExchangeStore) error {
	cfg, ok := cfgi.(*Config)
	if !ok {
		return fmt.Errorf("invalid configs")
	}
	cfg, err := cfg.validate()
	if err != nil {
		return err
	}
	cfg.Enable = k.cfg.Enable
	if cfg.Id != k.Id() {
		return fmt.Errorf("the id field is not mutable")
	}

	readApi := kucoin.NewApiService(
		kucoin.ApiBaseURIOption(cfg.ApiUrl),
		kucoin.ApiKeyOption(cfg.ReadApi.ApiKey),
		kucoin.ApiSecretOption(cfg.ReadApi.ApiSecret),
		kucoin.ApiPassPhraseOption(cfg.ReadApi.ApiPassphrase),
		kucoin.ApiKeyVersionOption(cfg.ApiVersion))

	writeApi := kucoin.NewApiService(
		kucoin.ApiBaseURIOption(cfg.ApiUrl),
		kucoin.ApiKeyOption(cfg.WriteApi.ApiKey),
		kucoin.ApiSecretOption(cfg.WriteApi.ApiSecret),
		kucoin.ApiPassPhraseOption(cfg.WriteApi.ApiPassphrase),
		kucoin.ApiKeyVersionOption(cfg.ApiVersion))

	if err := ping(readApi); err != nil {
		return err
	}
	if err := ping(writeApi); err != nil {
		return err
	}
	if err := store.UpdateConfigs(k, cfg); err != nil {
		return err
	}

	k.readApi = readApi
	k.writeApi = writeApi
	k.cfg = cfg

	return nil
}

func (cfg *Config) validate() (*Config, error) {

	if cfg.Id == 0 {
		return nil, fmt.Errorf("id is required")
	}
	if cfg.ReadApi.ApiKey == "" {
		return nil, fmt.Errorf("readApi.apiKey is required")
	}
	if cfg.ReadApi.ApiSecret == "" {
		return nil, fmt.Errorf("readApi.apiSecret is required")
	}
	if cfg.ReadApi.ApiPassphrase == "" {
		return nil, fmt.Errorf("readApi.apiPassphrase is required")
	}

	if cfg.WriteApi.ApiKey == "" {
		return nil, fmt.Errorf("writeApi.apiKey is required")
	}
	if cfg.WriteApi.ApiSecret == "" {
		return nil, fmt.Errorf("writeApi.apiSecret is required")
	}
	if cfg.WriteApi.ApiPassphrase == "" {
		return nil, fmt.Errorf("writeApi.apiPassphrase is required")
	}

	if cfg.ApiVersion == "" {
		cfg.ApiVersion = "2"
	}
	if cfg.ApiUrl == "" {
		cfg.ApiUrl = "https://api.kucoin.com"
	}

	return cfg, nil

}
