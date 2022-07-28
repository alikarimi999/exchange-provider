package kucoin

import (
	"fmt"
	"order_service/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) ChangeAccount(cfgi interface{}) error {
	const agent = "Kucoin-Exchange.ChangeAccount"

	cfg, err := validateConfigs(cfgi)
	if err != nil {
		return errors.Wrap(agent, err)
	}

	k1 := &kucoinExchange{
		cfg:             cfg,
		l:               k.l,
		exchangePairs:   k.exchangePairs,
		withdrawalCoins: k.withdrawalCoins,
	}

	k1.api = kucoin.NewApiService(
		kucoin.ApiBaseURIOption(cfg.ApiUrl),
		kucoin.ApiKeyOption(cfg.ApiKey),
		kucoin.ApiSecretOption(cfg.ApiSecret),
		kucoin.ApiPassPhraseOption(cfg.ApiPassphrase),
		kucoin.ApiKeyVersionOption(cfg.ApiVersion),
	)

	if err := k1.ping(); err != nil {
		err = errors.New(fmt.Sprintf("change account failed due to error: %s", err.Error()))
		k1.l.Info(agent, err.Error())
		return errors.Wrap(agent, err)
	}

	k1.l.Debug(agent, "ping was successful")

	k.api = k1.api
	k.cfg = cfg
	k.ot.api = k1.api
	k.wa.api = k1.api
	k.pls.api = k1.api

	return nil

}
