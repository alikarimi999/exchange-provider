package exrepo

import (
	"encoding/json"
	"fmt"
	"exchange-provider/internal/delivery/exchanges/kucoin"
	uniswapv3 "exchange-provider/internal/delivery/exchanges/uniswap/v3"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
)

func (r *ExchangeRepo) decrypt(ex *Exchange) (entity.Exchange, error) {

	jb := make(jsonb)

	dec, err := utils.RSA_OAEP_Decrypt(ex.Configs, *r.prv)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(dec), &jb); err != nil {
		return nil, err
	}

	switch ex.Name {
	case "kucoin":
		key, ok := jb["api_key"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid API key: %v", jb["api_key"])
		}
		secret, ok := jb["api_secret"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid API secret: %v", jb["api_secret"])
		}
		passphrase, ok := jb["api_passphrase"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid API passphrase: %v", jb["api_passphrase"])
		}

		cfg := &kucoin.Configs{
			ApiKey:        key,
			ApiSecret:     secret,
			ApiPassphrase: passphrase,
		}

		return kucoin.NewKucoinExchange(cfg, r.rc, r.v, r.l, true)

	case "uniswapv3":

		m, ok := jb["mnemonic"].(string)
		if !ok {
			return nil, errors.Wrap(errors.New(fmt.Sprintf("`%+v` does not have a mnemonic", ex)))
		}

		cfg := &uniswapv3.Config{
			Mnemonic: m,
		}
		return uniswapv3.NewExchange(cfg, r.rc, r.v, r.l, true)

	}
	return nil, errors.Wrap(errors.New(fmt.Sprintf("unkown exchange `%s`", ex.Name)))
}
