package exrepo

import (
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/dex"
	"exchange-provider/internal/delivery/exchanges/kucoin"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"fmt"
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

	case "uniswapv3", "panckakeswapv2", "multichain":

		m, ok := jb["mnemonic"].(string)
		if !ok {
			return nil, errors.Wrap(errors.New(fmt.Sprintf("`%+v` does not have mnemonic paramether", ex)))
		}
		n, ok := jb["network"].(string)
		if !ok {
			return nil, errors.Wrap(errors.New(fmt.Sprintf("`%+v` does not have network paramether", ex)))
		}

		cfg := &dex.Config{
			Mnemonic: m,
			Name:     ex.Name,
			Network:  n,
		}
		return dex.NewDEX(cfg, r.rc, r.v, r.l, true)

	}
	return nil, errors.Wrap(errors.New(fmt.Sprintf("unkown exchange `%s`", ex.Name)))
}
