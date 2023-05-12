package exrepo

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *ExchangeRepo) decrypt(ex *Exchange) (entity.Exchange, error) {
	dec, err := utils.RSA_OAEP_Decrypt(ex.Configs, *r.prv)
	if err != nil {
		return nil, err
	}

	switch ex.Type {
	case entity.CEX:
		switch ex.Name {
		case "kucoin":
			cfg := &kucoin.Configs{}
			if err := bson.Unmarshal([]byte(dec), cfg); err != nil {
				return nil, err
			}
			cfg.Enable = ex.Enable
			return kucoin.NewKucoinExchange(cfg, r.pairs, r.l, true, r.repo, r.fee, r.spread)

			// case "swapspace":
			// 	cfg := &swapspace.Config{}
			// 	if err := bson.Unmarshal([]byte(dec), cfg); err != nil {
			// 		return nil, err
			// 	}
			// 	return swapspace.SwapSpace(cfg, r.repo, r.pairs, r.l)

		}
	case entity.EvmDEX:
		cfg := &evm.Config{}
		if err := bson.Unmarshal([]byte(dec), cfg); err != nil {
			return nil, err
		}
		cfg.Enable = ex.Enable
		return evm.NewEvmDex(cfg, r.repo, r.pairs, r.l, true)
	}
	return nil, errors.Wrap(errors.New(fmt.Sprintf("unkown exchange `%d`", ex.Id)))
}
