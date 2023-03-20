package exrepo

import (
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin"
	"exchange-provider/internal/delivery/exchanges/cex/swapspace"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
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
		rKey := jb["read.apiKey"].(string)
		rSecret := jb["read.apiSecret"].(string)
		rPassphrase := jb["read.apiPassphrase"].(string)

		wKey := jb["write.apiKey"].(string)
		wSecret := jb["write.apiSecret"].(string)
		wPassphrase := jb["write.apiPassphrase"].(string)

		cfg := &kucoin.Configs{
			ReadApi: &kucoin.API{
				ApiKey:        rKey,
				ApiSecret:     rSecret,
				ApiPassphrase: rPassphrase,
			},
			WriteApi: &kucoin.API{
				ApiKey:        wKey,
				ApiSecret:     wSecret,
				ApiPassphrase: wPassphrase,
			},
		}

		return kucoin.NewKucoinExchange(cfg, r.v, r.l, true, r.repo, r.pc, r.fee)

	case "swapspace":
		apiKey := jb["apiKey"].(string)
		cfg := &swapspace.Config{ApiKey: apiKey, Id: ex.Id}
		return swapspace.SwapSpace(cfg, r.repo, r.pairs, r.l)

	case "uniswapv3", "uniswapv2", "panckakeswapv2":

		hk, ok := jb["hexKey"].(string)
		if !ok {
			return nil, errors.Wrap(errors.New(fmt.Sprintf("`%+v` does not have mnemonic paramether", ex)))
		}
		n, ok := jb["network"].(string)
		if !ok {
			return nil, errors.Wrap(errors.New(fmt.Sprintf("`%+v` does not have network paramether", ex)))
		}

		cfg := &evm.Config{
			HexKey:  hk,
			Name:    ex.Name,
			Network: n,
		}
		return evm.NewEvmDex(cfg, r.v, r.l, true)

		// case "multichain":
		// 	m, ok := jb["mnemonic"].(string)
		// 	if !ok {
		// 		return nil, errors.Wrap(errors.New(fmt.Sprintf("`%+v` does not have mnemonic paramether", ex)))
		// 	}

		// 	cfg := &multichain.Config{Name: ex.Id, Mnemonic: m}
		// 	return multichain.NewMultichain(cfg, r.WalletStore, r.v, r.l, true)

	}
	return nil, errors.Wrap(errors.New(fmt.Sprintf("unkown exchange `%d`", ex.Id)))
}
