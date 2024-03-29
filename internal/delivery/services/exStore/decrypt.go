package store

import (
	"exchange-provider/internal/delivery/exchanges/cex/binance"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *exchangeRepo) decrypt(ex *Exchange, lastUpdate time.Time) (entity.Exchange, error) {
	dec, err := utils.RSA_OAEP_Decrypt(ex.Configs, *r.prv)
	if err != nil {
		return nil, err
	}

	switch ex.Type {
	case entity.CEX:
		switch ex.Name {
		case "kucoin":
			cfg := &kucoin.Config{}
			if err := bson.Unmarshal([]byte(dec), cfg); err != nil {
				return nil, err
			}
			cfg.Enable = ex.Enable
			return kucoin.NewExchange(cfg, r.pairs, r.l, true, lastUpdate, r.repo, r.fee, r.spread)
		case "binance":
			cfg := &binance.Config{}
			if err := bson.Unmarshal([]byte(dec), cfg); err != nil {
				return nil, err
			}
			cfg.Enable = ex.Enable
			return binance.NewExchange(cfg, r.repo, r.pairs, r.spread, r.l, lastUpdate, true)
		}
	case entity.EvmDEX:
		cfg := &evm.Config{}
		if err := bson.Unmarshal([]byte(dec), cfg); err != nil {
			return nil, err
		}
		cfg.Enable = ex.Enable
		for _, p := range ex.Providers {
			cfg.Providers = p
		}
		return evm.NewEvmDex(cfg, r.repo, r.pairs, r.l)

	case entity.CrossDex:
		if ex.Name == "allbridge" {
			cfg := &allbridge.Config{}
			if err := bson.Unmarshal([]byte(dec), cfg); err != nil {
				return nil, err
			}

			for n, ps := range ex.Providers {
				cfg.Networks[n].Provider = ps[0]
			}

			cfg.Enable = ex.Enable
			return allbridge.NewExchange(cfg, r.exs, r.repo, r.exs, r.pairs, r.l, true)
		}
	}
	return nil, errors.Wrap(errors.New(fmt.Sprintf("unkown exchange `%d`", ex.Id)))
}
